package agent

import (
	"context"
	"fmt"
	"net"
	"os"
	"strings"
	"time"

	"github.com/k3s-io/k3s/pkg/agent/config"
	"github.com/k3s-io/k3s/pkg/agent/proxy"
	agentutil "github.com/k3s-io/k3s/pkg/agent/util"
	daemonconfig "github.com/k3s-io/k3s/pkg/daemons/config"
	"github.com/k3s-io/k3s/pkg/daemons/executor"
	"github.com/k3s-io/k3s/pkg/signals"
	"github.com/k3s-io/k3s/pkg/util"
	pkgerrors "github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/component-base/logs"
	logsv1 "k8s.io/component-base/logs/api/v1"
	_ "k8s.io/component-base/metrics/prometheus/restclient" // for client metric registration
	_ "k8s.io/component-base/metrics/prometheus/version"    // for version metric registration
	kubeletconfig "k8s.io/kubelet/config/v1beta1"
	"k8s.io/kubernetes/pkg/util/taints"
	utilsnet "k8s.io/utils/net"
	utilsptr "k8s.io/utils/ptr"
	"sigs.k8s.io/yaml"
)

func Agent(ctx context.Context, nodeConfig *daemonconfig.Node, proxy proxy.Proxy) error {
	logs.InitLogs()
	defer logs.FlushLogs()

	go func() {
		<-executor.CRIReadyChan()
		if err := startKubelet(ctx, &nodeConfig.AgentConfig); err != nil {
			signals.RequestShutdown(pkgerrors.WithMessage(err, "failed to start kubelet"))
		}
	}()

	go func() {
		if !config.KubeProxyDisabled(ctx, nodeConfig, proxy) {
			if err := startKubeProxy(ctx, &nodeConfig.AgentConfig); err != nil {
				signals.RequestShutdown(pkgerrors.WithMessage(err, "failed to start kube-proxy"))
			}
		}
	}()

	return nil
}

func startKubeProxy(ctx context.Context, cfg *daemonconfig.Agent) error {
	argsMap := kubeProxyArgs(cfg)
	args := util.GetArgs(argsMap, cfg.ExtraKubeProxyArgs)
	logrus.Infof("Running kube-proxy %s", daemonconfig.ArgString(args))
	return executor.KubeProxy(ctx, args)
}

func startKubelet(ctx context.Context, cfg *daemonconfig.Agent) error {
	argsMap, defaultConfig, err := kubeletArgsAndConfig(cfg)
	if err != nil {
		return pkgerrors.WithMessage(err, "prepare default configuration drop-in")
	}

	extraArgs, err := extractConfigArgs(cfg.KubeletConfig, cfg.ExtraKubeletArgs)
	if err != nil {
		return pkgerrors.WithMessage(err, "prepare user configuration drop-ins")
	}

	if err := writeKubeletConfig(cfg.KubeletConfig, defaultConfig); err != nil {
		return pkgerrors.WithMessage(err, "generate default kubelet configuration drop-in")
	}

	args := util.GetArgs(argsMap, extraArgs)
	logrus.Infof("Running kubelet %s", daemonconfig.ArgString(args))

	return executor.Kubelet(ctx, args)
}

// ImageCredProvAvailable checks to see if the kubelet image credential provider bin dir and config
// files exist and are of the correct types. This is exported so that it may be used by downstream projects.
func ImageCredProvAvailable(cfg *daemonconfig.Agent) bool {
	if info, err := os.Stat(cfg.ImageCredProvBinDir); err != nil || !info.IsDir() {
		logrus.Debugf("Kubelet image credential provider bin directory check failed: %v", err)
		return false
	}
	if info, err := os.Stat(cfg.ImageCredProvConfig); err != nil || info.IsDir() {
		logrus.Debugf("Kubelet image credential provider config file check failed: %v", err)
		return false
	}
	return true
}

// extractConfigArgs strips out any --config flags from the
// provided args list, and if set, copies the content of the file or dir into
// the target drop-in directory.
func extractConfigArgs(path string, extraArgs []string) ([]string, error) {
	args := make([]string, 0, len(extraArgs))
	strippedArgs := map[string]string{}
	var skipVal bool
	for i := range extraArgs {
		if skipVal {
			skipVal = false
			continue
		}

		var val string
		key := strings.TrimPrefix(extraArgs[i], "--")
		if k, v, ok := strings.Cut(key, "="); ok {
			// key=val pair
			key = k
			val = v
		} else if len(extraArgs) > i+1 {
			// key in this arg, value in next arg
			val = extraArgs[i+1]
			skipVal = true
		}

		switch key {
		case "config":
			if val == "" {
				return nil, fmt.Errorf("value required for kubelet-arg --%s", key)
			}
			strippedArgs[key] = val
		default:
			args = append(args, extraArgs[i])
		}
	}

	// copy the config file into our managed config dir, unless its already in there
	if strippedArgs["config"] != "" && !strings.HasPrefix(strippedArgs["config"], path) {
		src := strippedArgs["config"]
		dest := path
		if err := agentutil.CopyFile(src, dest, false); err != nil {
			return nil, pkgerrors.WithMessagef(err, "copy config %q into managed drop-in dir %q", src, dest)
		}
	}
	return args, nil
}

// writeKubeletConfig marshals the provided KubeletConfiguration object into a
// drop-in config file in the target drop-in directory.
func writeKubeletConfig(path string, config *kubeletconfig.KubeletConfiguration) error {
	b, err := yaml.Marshal(config)
	if err != nil {
		return err
	}
	return os.WriteFile(path, b, 0600)
}

func defaultKubeletConfig(cfg *daemonconfig.Agent) (*kubeletconfig.KubeletConfiguration, error) {
	bindAddress := "127.0.0.1"
	if utilsnet.IsIPv6(net.ParseIP([]string{cfg.NodeIP}[0])) {
		bindAddress = "::1"
	}

	defaultConfig := &kubeletconfig.KubeletConfiguration{
		TypeMeta: metav1.TypeMeta{
			APIVersion: "kubelet.config.k8s.io/v1beta1",
			Kind:       "KubeletConfiguration",
		},
		CPUManagerReconcilePeriod:        metav1.Duration{Duration: time.Second * 10},
		CgroupDriver:                     "cgroupfs",
		ClusterDomain:                    cfg.ClusterDomain,
		EvictionPressureTransitionPeriod: metav1.Duration{Duration: time.Minute * 5},
		FailSwapOn:                       utilsptr.To(false),
		FileCheckFrequency:               metav1.Duration{Duration: time.Second * 20},
		HTTPCheckFrequency:               metav1.Duration{Duration: time.Second * 20},
		HealthzBindAddress:               bindAddress,
		ImageMinimumGCAge:                metav1.Duration{Duration: time.Minute * 2},
		NodeStatusReportFrequency:        metav1.Duration{Duration: time.Minute * 5},
		NodeStatusUpdateFrequency:        metav1.Duration{Duration: time.Second * 10},
		ProtectKernelDefaults:            cfg.ProtectKernelDefaults,
		ReadOnlyPort:                     0,
		RuntimeRequestTimeout:            metav1.Duration{Duration: time.Minute * 2},
		StreamingConnectionIdleTimeout:   metav1.Duration{Duration: time.Hour * 4},
		SyncFrequency:                    metav1.Duration{Duration: time.Minute},
		VolumeStatsAggPeriod:             metav1.Duration{Duration: time.Minute},
		EvictionHard: map[string]string{
			"imagefs.available": "15%",
			"nodefs.available":  "15%",
		},
		EvictionMinimumReclaim: map[string]string{
			"imagefs.available": "20%",
			"nodefs.available":  "20%",
		},
		Authentication: kubeletconfig.KubeletAuthentication{
			Anonymous: kubeletconfig.KubeletAnonymousAuthentication{
				Enabled: utilsptr.To(false),
			},
			Webhook: kubeletconfig.KubeletWebhookAuthentication{
				Enabled:  utilsptr.To(true),
				CacheTTL: metav1.Duration{Duration: time.Minute * 2},
			},
		},
		Authorization: kubeletconfig.KubeletAuthorization{
			Mode: kubeletconfig.KubeletAuthorizationModeWebhook,
			Webhook: kubeletconfig.KubeletWebhookAuthorization{
				CacheAuthorizedTTL:   metav1.Duration{Duration: time.Minute * 5},
				CacheUnauthorizedTTL: metav1.Duration{Duration: time.Second * 30},
			},
		},
		Logging: logsv1.LoggingConfiguration{
			Format:         "text",
			Verbosity:      logsv1.VerbosityLevel(cfg.VLevel),
			FlushFrequency: time.Second * 5,
		},
	}

	if cfg.ListenAddress != "" {
		defaultConfig.Address = cfg.ListenAddress
	}

	if cfg.ClientCA != "" {
		defaultConfig.Authentication.X509.ClientCAFile = cfg.ClientCA
	}

	if cfg.ServingKubeletCert != "" && cfg.ServingKubeletKey != "" {
		defaultConfig.TLSCertFile = cfg.ServingKubeletCert
		defaultConfig.TLSPrivateKeyFile = cfg.ServingKubeletKey
	}

	for _, addr := range cfg.ClusterDNSs {
		defaultConfig.ClusterDNS = append(defaultConfig.ClusterDNS, addr.String())
	}

	if cfg.ResolvConf != "" {
		defaultConfig.ResolverConfig = utilsptr.To(cfg.ResolvConf)
	}

	if cfg.PodManifests != "" && defaultConfig.StaticPodPath == "" {
		defaultConfig.StaticPodPath = cfg.PodManifests
	}
	if err := os.MkdirAll(defaultConfig.StaticPodPath, 0750); err != nil {
		return nil, pkgerrors.WithMessagef(err, "failed to create static pod manifest dir %s", defaultConfig.StaticPodPath)
	}

	t, _, err := taints.ParseTaints(cfg.NodeTaints)
	if err != nil {
		return nil, pkgerrors.WithMessage(err, "failed to parse node taints")
	}

	defaultConfig.RegisterWithTaints = t
	logsv1.VModuleConfigurationPflag(&defaultConfig.Logging.VModule).Set(cfg.VModule)

	return defaultConfig, nil
}
