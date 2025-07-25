package deps

import (
	"bytes"
	"crypto"
	"crypto/rand"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io/fs"
	"net"
	"os"
	"path/filepath"
	"strings"
	"text/template"
	"time"

	"github.com/k3s-io/k3s/pkg/clientaccess"
	"github.com/k3s-io/k3s/pkg/cloudprovider"
	"github.com/k3s-io/k3s/pkg/daemons/config"
	"github.com/k3s-io/k3s/pkg/passwd"
	"github.com/k3s-io/k3s/pkg/secretsencrypt"
	"github.com/k3s-io/k3s/pkg/util"
	"github.com/k3s-io/k3s/pkg/version"
	certutil "github.com/rancher/dynamiclistener/cert"
	"github.com/sirupsen/logrus"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/sets"
	apiserverconfigv1 "k8s.io/apiserver/pkg/apis/config/v1"
	apiserverv1beta1 "k8s.io/apiserver/pkg/apis/apiserver/v1beta1"
	"k8s.io/apiserver/pkg/authentication/user"
	"k8s.io/client-go/util/keyutil"
)

const (
	ipsecTokenSize = 48

	RequestHeaderCN = "system:auth-proxy"
)

var kubeconfigTemplate = template.Must(template.New("kubeconfig").Parse(`apiVersion: v1
clusters:
- cluster:
    server: {{.URL}}
    certificate-authority: {{.CACert}}
  name: local
contexts:
- context:
    cluster: local
    namespace: default
    user: user
  name: Default
current-context: Default
kind: Config
preferences: {}
users:
- name: user
  user:
    client-certificate: {{.ClientCert}}
    client-key: {{.ClientKey}}
`))

func migratePassword(p *passwd.Passwd) error {
	server, _ := p.Pass("server")
	node, _ := p.Pass("node")
	if server == "" && node != "" {
		return p.EnsureUser("server", version.Program+":server", node)
	}
	return nil
}

func KubeConfig(dest, url, caCert, clientCert, clientKey string) error {
	data := struct {
		URL        string
		CACert     string
		ClientCert string
		ClientKey  string
	}{
		URL:        url,
		CACert:     caCert,
		ClientCert: clientCert,
		ClientKey:  clientKey,
	}

	// cis-1.24 and newer require kubeconfigs to be 0600
	output, err := os.OpenFile(dest, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		return err
	}
	defer output.Close()

	return kubeconfigTemplate.Execute(output, &data)
}

// CreateRuntimeCertFiles is responsible for filling out all the
// .crt and .key filenames for a ControlRuntime.
func CreateRuntimeCertFiles(config *config.Control) {
	runtime := config.Runtime
	runtime.ClientCA = filepath.Join(config.DataDir, "tls", "client-ca.crt")
	runtime.ClientCAKey = filepath.Join(config.DataDir, "tls", "client-ca.key")
	runtime.ServerCA = filepath.Join(config.DataDir, "tls", "server-ca.crt")
	runtime.ServerCAKey = filepath.Join(config.DataDir, "tls", "server-ca.key")
	runtime.RequestHeaderCA = filepath.Join(config.DataDir, "tls", "request-header-ca.crt")
	runtime.RequestHeaderCAKey = filepath.Join(config.DataDir, "tls", "request-header-ca.key")
	runtime.IPSECKey = filepath.Join(config.DataDir, "cred", "ipsec.psk")

	runtime.ServiceKey = filepath.Join(config.DataDir, "tls", "service.key")
	runtime.PasswdFile = filepath.Join(config.DataDir, "cred", "passwd")
	runtime.NodePasswdFile = filepath.Join(config.DataDir, "cred", "node-passwd")

	runtime.SigningClientCA = filepath.Join(config.DataDir, "tls", "client-ca.nochain.crt")
	runtime.SigningServerCA = filepath.Join(config.DataDir, "tls", "server-ca.nochain.crt")
	runtime.ServiceCurrentKey = filepath.Join(config.DataDir, "tls", "service.current.key")

	runtime.KubeConfigAdmin = filepath.Join(config.DataDir, "cred", "admin.kubeconfig")
	runtime.KubeConfigSupervisor = filepath.Join(config.DataDir, "cred", "supervisor.kubeconfig")
	runtime.KubeConfigController = filepath.Join(config.DataDir, "cred", "controller.kubeconfig")
	runtime.KubeConfigScheduler = filepath.Join(config.DataDir, "cred", "scheduler.kubeconfig")
	runtime.KubeConfigAPIServer = filepath.Join(config.DataDir, "cred", "api-server.kubeconfig")
	runtime.KubeConfigCloudController = filepath.Join(config.DataDir, "cred", "cloud-controller.kubeconfig")

	runtime.ClientAdminCert = filepath.Join(config.DataDir, "tls", "client-admin.crt")
	runtime.ClientAdminKey = filepath.Join(config.DataDir, "tls", "client-admin.key")
	runtime.ClientSupervisorCert = filepath.Join(config.DataDir, "tls", "client-supervisor.crt")
	runtime.ClientSupervisorKey = filepath.Join(config.DataDir, "tls", "client-supervisor.key")
	runtime.ClientControllerCert = filepath.Join(config.DataDir, "tls", "client-controller.crt")
	runtime.ClientControllerKey = filepath.Join(config.DataDir, "tls", "client-controller.key")
	runtime.ClientCloudControllerCert = filepath.Join(config.DataDir, "tls", "client-"+version.Program+"-cloud-controller.crt")
	runtime.ClientCloudControllerKey = filepath.Join(config.DataDir, "tls", "client-"+version.Program+"-cloud-controller.key")
	runtime.ClientSchedulerCert = filepath.Join(config.DataDir, "tls", "client-scheduler.crt")
	runtime.ClientSchedulerKey = filepath.Join(config.DataDir, "tls", "client-scheduler.key")
	runtime.ClientKubeAPICert = filepath.Join(config.DataDir, "tls", "client-kube-apiserver.crt")
	runtime.ClientKubeAPIKey = filepath.Join(config.DataDir, "tls", "client-kube-apiserver.key")
	runtime.ClientKubeProxyCert = filepath.Join(config.DataDir, "tls", "client-kube-proxy.crt")
	runtime.ClientKubeProxyKey = filepath.Join(config.DataDir, "tls", "client-kube-proxy.key")
	runtime.ClientK3sControllerCert = filepath.Join(config.DataDir, "tls", "client-"+version.Program+"-controller.crt")
	runtime.ClientK3sControllerKey = filepath.Join(config.DataDir, "tls", "client-"+version.Program+"-controller.key")

	runtime.ServingKubeAPICert = filepath.Join(config.DataDir, "tls", "serving-kube-apiserver.crt")
	runtime.ServingKubeAPIKey = filepath.Join(config.DataDir, "tls", "serving-kube-apiserver.key")

	runtime.ServingKubeSchedulerCert = filepath.Join(config.DataDir, "tls", "kube-scheduler", "kube-scheduler.crt")
	runtime.ServingKubeSchedulerKey = filepath.Join(config.DataDir, "tls", "kube-scheduler", "kube-scheduler.key")

	runtime.ServingKubeControllerCert = filepath.Join(config.DataDir, "tls", "kube-controller-manager", "kube-controller-manager.crt")
	runtime.ServingKubeControllerKey = filepath.Join(config.DataDir, "tls", "kube-controller-manager", "kube-controller-manager.key")

	runtime.ClientKubeletKey = filepath.Join(config.DataDir, "tls", "client-kubelet.key")
	runtime.ServingKubeletKey = filepath.Join(config.DataDir, "tls", "serving-kubelet.key")

	runtime.EgressSelectorConfig = filepath.Join(config.DataDir, "etc", "egress-selector-config.yaml")
	runtime.CloudControllerConfig = filepath.Join(config.DataDir, "etc", "cloud-config.yaml")

	runtime.ClientAuthProxyCert = filepath.Join(config.DataDir, "tls", "client-auth-proxy.crt")
	runtime.ClientAuthProxyKey = filepath.Join(config.DataDir, "tls", "client-auth-proxy.key")

	runtime.ETCDServerCA = filepath.Join(config.DataDir, "tls", "etcd", "server-ca.crt")
	runtime.ETCDServerCAKey = filepath.Join(config.DataDir, "tls", "etcd", "server-ca.key")
	runtime.ETCDPeerCA = filepath.Join(config.DataDir, "tls", "etcd", "peer-ca.crt")
	runtime.ETCDPeerCAKey = filepath.Join(config.DataDir, "tls", "etcd", "peer-ca.key")
	runtime.ServerETCDCert = filepath.Join(config.DataDir, "tls", "etcd", "server-client.crt")
	runtime.ServerETCDKey = filepath.Join(config.DataDir, "tls", "etcd", "server-client.key")
	runtime.PeerServerClientETCDCert = filepath.Join(config.DataDir, "tls", "etcd", "peer-server-client.crt")
	runtime.PeerServerClientETCDKey = filepath.Join(config.DataDir, "tls", "etcd", "peer-server-client.key")
	runtime.ClientETCDCert = filepath.Join(config.DataDir, "tls", "etcd", "client.crt")
	runtime.ClientETCDKey = filepath.Join(config.DataDir, "tls", "etcd", "client.key")

	if config.EncryptSecrets {
		runtime.EncryptionConfig = filepath.Join(config.DataDir, "cred", "encryption-config.json")
		runtime.EncryptionHash = filepath.Join(config.DataDir, "cred", "encryption-state.json")
	}
}

// GenServerDeps is responsible for generating the cluster dependencies
// needed to successfully bootstrap a cluster.
func GenServerDeps(config *config.Control) error {
	runtime := config.Runtime

	if err := cleanupLegacyCerts(config); err != nil {
		return err
	}

	if err := genCerts(config); err != nil {
		return err
	}

	if err := genServiceAccount(runtime); err != nil {
		return err
	}

	if err := genUsers(config); err != nil {
		return err
	}

	if err := genEncryptedNetworkInfo(config); err != nil {
		return err
	}

	if err := genEncryptionConfigAndState(config); err != nil {
		return err
	}

	if err := genEgressSelectorConfig(config); err != nil {
		return err
	}

	if err := genCloudConfig(config); err != nil {
		return err
	}

	return readTokens(runtime)
}

func readTokens(runtime *config.ControlRuntime) error {
	tokens, err := passwd.Read(runtime.PasswdFile)
	if err != nil {
		return err
	}

	if nodeToken, ok := tokens.Pass("node"); ok {
		runtime.AgentToken = "node:" + nodeToken
	}
	if serverToken, ok := tokens.Pass("server"); ok {
		runtime.ServerToken = "server:" + serverToken
	}

	return nil
}

func getNodePass(config *config.Control, serverPass string) string {
	if config.AgentToken == "" {
		if _, passwd, ok := clientaccess.ParseUsernamePassword(serverPass); ok {
			return passwd
		}
		return serverPass
	}
	return config.AgentToken
}

func genUsers(config *config.Control) error {
	runtime := config.Runtime
	passwd, err := passwd.Read(runtime.PasswdFile)
	if err != nil {
		return err
	}

	if err := migratePassword(passwd); err != nil {
		return err
	}

	// if no token is provided on bootstrap, we generate a random token
	serverPass, err := getServerPass(passwd, config)
	if err != nil {
		return err
	}

	nodePass := getNodePass(config, serverPass)

	if err := passwd.EnsureUser("node", version.Program+":agent", nodePass); err != nil {
		return err
	}

	if err := passwd.EnsureUser("server", version.Program+":server", serverPass); err != nil {
		return err
	}

	return passwd.Write(runtime.PasswdFile)
}

func genEncryptedNetworkInfo(controlConfig *config.Control) error {
	runtime := controlConfig.Runtime
	if s, err := os.Stat(runtime.IPSECKey); err == nil && s.Size() > 0 {
		psk, err := os.ReadFile(runtime.IPSECKey)
		if err != nil {
			return err
		}
		controlConfig.IPSECPSK = strings.TrimSpace(string(psk))
		return nil
	}

	psk, err := util.Random(ipsecTokenSize)
	if err != nil {
		return err
	}

	controlConfig.IPSECPSK = psk
	return os.WriteFile(runtime.IPSECKey, []byte(psk+"\n"), 0600)
}

func getServerPass(passwd *passwd.Passwd, config *config.Control) (string, error) {
	var err error

	serverPass := config.Token
	if serverPass == "" {
		serverPass, _ = passwd.Pass("server")
	}
	if serverPass == "" {
		serverPass, err = util.Random(16)
		if err != nil {
			return "", err
		}
	}

	return serverPass, nil
}

func genCerts(config *config.Control) error {
	if err := genClientCerts(config); err != nil {
		return err
	}
	if err := genServerCerts(config); err != nil {
		return err
	}
	if err := genRequestHeaderCerts(config); err != nil {
		return err
	}
	return genETCDCerts(config)
}

func getSigningCertFactory(regen bool, altNames *certutil.AltNames, extKeyUsage []x509.ExtKeyUsage, caCertFile, caKeyFile string) signedCertFactory {
	return func(commonName string, organization []string, certFile, keyFile string) (bool, error) {
		return createClientCertKey(regen, commonName, organization, altNames, extKeyUsage, caCertFile, caKeyFile, certFile, keyFile)
	}
}

func genClientCerts(config *config.Control) error {
	runtime := config.Runtime
	regen, err := createSigningCertKey(version.Program+"-client", runtime.ClientCA, runtime.ClientCAKey)
	if err != nil {
		return err
	}

	certs, err := certutil.CertsFromFile(runtime.ClientCA)
	if err != nil {
		return err
	}

	// If our CA certs are signed by a root or intermediate CA, ClientCA will contain a chain.
	// The controller-manager's signer wants just a single cert, not a full chain; so create a file
	// that is guaranteed to contain only a single certificate.
	if err := certutil.WriteCert(runtime.SigningClientCA, certutil.EncodeCertPEM(certs[0])); err != nil {
		return err
	}

	factory := getSigningCertFactory(regen, nil, []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth}, runtime.ClientCA, runtime.ClientCAKey)

	var certGen bool

	apiEndpoint := fmt.Sprintf("https://%s:%d", config.Loopback(true), config.APIServerPort)

	certGen, err = factory("system:admin", []string{user.SystemPrivilegedGroup}, runtime.ClientAdminCert, runtime.ClientAdminKey)
	if err != nil {
		return err
	}
	if certGen {
		if err := KubeConfig(runtime.KubeConfigAdmin, apiEndpoint, runtime.ServerCA, runtime.ClientAdminCert, runtime.ClientAdminKey); err != nil {
			return err
		}
	}

	certGen, err = factory("system:"+version.Program+"-supervisor", []string{user.SystemPrivilegedGroup}, runtime.ClientSupervisorCert, runtime.ClientSupervisorKey)
	if err != nil {
		return err
	}
	if certGen {
		if err := KubeConfig(runtime.KubeConfigSupervisor, apiEndpoint, runtime.ServerCA, runtime.ClientSupervisorCert, runtime.ClientSupervisorKey); err != nil {
			return err
		}
	}

	certGen, err = factory(user.KubeControllerManager, nil, runtime.ClientControllerCert, runtime.ClientControllerKey)
	if err != nil {
		return err
	}
	if certGen {
		if err := KubeConfig(runtime.KubeConfigController, apiEndpoint, runtime.ServerCA, runtime.ClientControllerCert, runtime.ClientControllerKey); err != nil {
			return err
		}
	}

	certGen, err = factory(user.KubeScheduler, nil, runtime.ClientSchedulerCert, runtime.ClientSchedulerKey)
	if err != nil {
		return err
	}
	if certGen {
		if err := KubeConfig(runtime.KubeConfigScheduler, apiEndpoint, runtime.ServerCA, runtime.ClientSchedulerCert, runtime.ClientSchedulerKey); err != nil {
			return err
		}
	}

	certGen, err = factory(user.APIServerUser, []string{user.SystemPrivilegedGroup}, runtime.ClientKubeAPICert, runtime.ClientKubeAPIKey)
	if err != nil {
		return err
	}
	if certGen {
		if err := KubeConfig(runtime.KubeConfigAPIServer, apiEndpoint, runtime.ServerCA, runtime.ClientKubeAPICert, runtime.ClientKubeAPIKey); err != nil {
			return err
		}
	}

	if _, _, err := certutil.LoadOrGenerateKeyFile(runtime.ClientKubeProxyKey, regen); err != nil {
		return err
	}

	if _, _, err := certutil.LoadOrGenerateKeyFile(runtime.ClientK3sControllerKey, regen); err != nil {
		return err
	}

	if _, _, err := certutil.LoadOrGenerateKeyFile(runtime.ClientKubeletKey, regen); err != nil {
		return err
	}

	certGen, err = factory(version.Program+"-cloud-controller-manager", nil, runtime.ClientCloudControllerCert, runtime.ClientCloudControllerKey)
	if err != nil {
		return err
	}
	if certGen {
		if err := KubeConfig(runtime.KubeConfigCloudController, apiEndpoint, runtime.ServerCA, runtime.ClientCloudControllerCert, runtime.ClientCloudControllerKey); err != nil {
			return err
		}
	}

	return nil
}

func genServerCerts(config *config.Control) error {
	runtime := config.Runtime
	regen, err := createServerSigningCertKey(config)
	if err != nil {
		return err
	}

	altNames := &certutil.AltNames{
		DNSNames: []string{"kubernetes", "kubernetes.default", "kubernetes.default.svc", "kubernetes.default.svc." + config.ClusterDomain},
	}

	addSANs(altNames, config.SANs)

	if _, err := createClientCertKey(regen, "kube-apiserver", nil,
		altNames, []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		runtime.ServerCA, runtime.ServerCAKey,
		runtime.ServingKubeAPICert, runtime.ServingKubeAPIKey); err != nil {
		return err
	}

	if _, _, err := certutil.LoadOrGenerateKeyFile(runtime.ServingKubeletKey, regen); err != nil {
		return err
	}

	altNames = &certutil.AltNames{}
	addSANs(altNames, []string{"localhost", "127.0.0.1", "::1"})

	if _, err := createClientCertKey(regen, "kube-scheduler", nil,
		altNames, []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		runtime.ServerCA, runtime.ServerCAKey,
		runtime.ServingKubeSchedulerCert, runtime.ServingKubeSchedulerKey); err != nil {
		return err
	}

	if _, err := createClientCertKey(regen, "kube-controller-manager", nil,
		altNames, []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		runtime.ServerCA, runtime.ServerCAKey,
		runtime.ServingKubeControllerCert, runtime.ServingKubeControllerKey); err != nil {
		return err
	}

	return nil
}

func genETCDCerts(config *config.Control) error {
	runtime := config.Runtime
	regen, err := createSigningCertKey("etcd-server", runtime.ETCDServerCA, runtime.ETCDServerCAKey)
	if err != nil {
		return err
	}

	altNames := &certutil.AltNames{
		DNSNames: []string{"kine.sock"},
	}

	addSANs(altNames, config.SANs)

	if _, err := createClientCertKey(regen, "etcd-client", nil,
		nil, []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth},
		runtime.ETCDServerCA, runtime.ETCDServerCAKey,
		runtime.ClientETCDCert, runtime.ClientETCDKey); err != nil {
		return err
	}

	regen, err = createSigningCertKey("etcd-peer", runtime.ETCDPeerCA, runtime.ETCDPeerCAKey)
	if err != nil {
		return err
	}

	if _, err := createClientCertKey(regen, "etcd-peer", nil,
		altNames, []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth, x509.ExtKeyUsageClientAuth},
		runtime.ETCDPeerCA, runtime.ETCDPeerCAKey,
		runtime.PeerServerClientETCDCert, runtime.PeerServerClientETCDKey); err != nil {
		return err
	}

	if config.DisableETCD {
		return nil
	}

	if _, err := createClientCertKey(regen, "etcd-server", nil,
		altNames, []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth, x509.ExtKeyUsageClientAuth},
		runtime.ETCDServerCA, runtime.ETCDServerCAKey,
		runtime.ServerETCDCert, runtime.ServerETCDKey); err != nil {
		return err
	}

	return nil
}

func genRequestHeaderCerts(config *config.Control) error {
	runtime := config.Runtime
	regen, err := createSigningCertKey(version.Program+"-request-header", runtime.RequestHeaderCA, runtime.RequestHeaderCAKey)
	if err != nil {
		return err
	}

	if _, err := createClientCertKey(regen, RequestHeaderCN, nil,
		nil, []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth},
		runtime.RequestHeaderCA, runtime.RequestHeaderCAKey,
		runtime.ClientAuthProxyCert, runtime.ClientAuthProxyKey); err != nil {
		return err
	}

	return nil
}

type signedCertFactory = func(commonName string, organization []string, certFile, keyFile string) (bool, error)

func createServerSigningCertKey(config *config.Control) (bool, error) {
	runtime := config.Runtime
	TokenCA := filepath.Join(config.DataDir, "tls", "token-ca.crt")
	TokenCAKey := filepath.Join(config.DataDir, "tls", "token-ca.key")

	if exists(TokenCA, TokenCAKey) && !exists(runtime.ServerCA) && !exists(runtime.ServerCAKey) {
		logrus.Infof("Upgrading token-ca files to server-ca")
		if err := os.Link(TokenCA, runtime.ServerCA); err != nil {
			return false, err
		}
		if err := os.Link(TokenCAKey, runtime.ServerCAKey); err != nil {
			return false, err
		}
		return true, nil
	}
	regen, err := createSigningCertKey(version.Program+"-server", runtime.ServerCA, runtime.ServerCAKey)
	if err != nil {
		return regen, err
	}

	// If our CA certs are signed by a root or intermediate CA, ServerCA will contain a chain.
	// The controller-manager's signer wants just a single cert, not a full chain; so create a file
	// that is guaranteed to contain only a single certificate.
	certs, err := certutil.CertsFromFile(runtime.ServerCA)
	if err != nil {
		return regen, err
	}

	if err := certutil.WriteCert(runtime.SigningServerCA, certutil.EncodeCertPEM(certs[0])); err != nil {
		return regen, err
	}

	return regen, nil
}

func addSANs(altNames *certutil.AltNames, sans []string) {
	for _, san := range sans {
		ip := net.ParseIP(san)
		if ip == nil {
			altNames.DNSNames = append(altNames.DNSNames, san)
		} else {
			altNames.IPs = append(altNames.IPs, ip)
		}
	}
}

func fieldsChanged(certFile string, commonName string, organization []string, sans *certutil.AltNames, caCertFile string) bool {
	if sans == nil {
		sans = &certutil.AltNames{}
	}

	certificates, err := certutil.CertsFromFile(certFile)
	if err != nil || len(certificates) == 0 {
		return false
	}

	if certificates[0].Subject.CommonName != commonName {
		return true
	}

	if !sets.NewString(certificates[0].Subject.Organization...).Equal(sets.NewString(organization...)) {
		return true
	}

	if !sets.NewString(certificates[0].DNSNames...).HasAll(sans.DNSNames...) {
		return true
	}

	ips := sets.NewString()
	for _, ip := range certificates[0].IPAddresses {
		ips.Insert(ip.String())
	}

	for _, ip := range sans.IPs {
		if !ips.Has(ip.String()) {
			return true
		}
	}

	caCertificates, err := certutil.CertsFromFile(caCertFile)
	if err != nil || len(caCertificates) == 0 {
		return false
	}

	return !bytes.Equal(certificates[0].AuthorityKeyId, caCertificates[0].SubjectKeyId)
}

func createClientCertKey(regen bool, commonName string, organization []string, altNames *certutil.AltNames, extKeyUsage []x509.ExtKeyUsage, caCertFile, caKeyFile, certFile, keyFile string) (bool, error) {
	// check for reasons to renew the certificate even if not manually requested.
	regen = regen || expired(certFile) || fieldsChanged(certFile, commonName, organization, altNames, caCertFile)

	if !regen {
		if exists(certFile, keyFile) {
			return false, nil
		}
	}

	caKey, err := certutil.PrivateKeyFromFile(caKeyFile)
	if err != nil {
		return false, err
	}

	caCerts, err := certutil.CertsFromFile(caCertFile)
	if err != nil {
		return false, err
	}

	keyBytes, _, err := certutil.LoadOrGenerateKeyFile(keyFile, regen)
	if err != nil {
		return false, err
	}

	key, err := certutil.ParsePrivateKeyPEM(keyBytes)
	if err != nil {
		return false, err
	}

	cfg := certutil.Config{
		CommonName:   commonName,
		Organization: organization,
		Usages:       extKeyUsage,
	}
	if altNames != nil {
		cfg.AltNames = *altNames
	}
	cert, err := certutil.NewSignedCert(cfg, key.(crypto.Signer), caCerts[0], caKey.(crypto.Signer))
	if err != nil {
		return false, err
	}

	return true, certutil.WriteCert(certFile, util.EncodeCertsPEM(cert, caCerts))
}

func cleanupLegacyCerts(config *config.Control) error {
	// remove legacy certs that are no longer used
	legacyCerts := []string{
		config.Runtime.ClientKubeProxyCert,
		config.Runtime.ClientK3sControllerCert,
	}
	for _, cert := range legacyCerts {
		if err := os.Remove(cert); err != nil && !errors.Is(err, os.ErrNotExist) {
			return err
		}
	}
	return nil
}

func exists(files ...string) bool {
	for _, file := range files {
		if _, err := os.Stat(file); err != nil {
			return false
		}
	}
	return true
}

func genServiceAccount(runtime *config.ControlRuntime) error {
	if _, err := os.Stat(runtime.ServiceKey); err != nil {
		if !errors.Is(err, fs.ErrNotExist) {
			return err
		}
		key, err := certutil.NewPrivateKey()
		if err != nil {
			return err
		}
		if err := certutil.WriteKey(runtime.ServiceKey, certutil.EncodePrivateKeyPEM(key)); err != nil {
			return err
		}
	}

	// When rotating the ServiceAccount signing key, it is necessary to keep the old keys in ServiceKey so that
	// old ServiceAccount tokens can be validated during the switchover process. The first key in the file
	// should be the current key used to sign ServiceAccount tokens; others are old keys used for verification
	// only. Create a file containing just the first key in the list, which will be used to configure the
	// signing controller.
	key, err := keyutil.PrivateKeyFromFile(runtime.ServiceKey)
	if err != nil {
		return err
	}

	keyData, err := keyutil.MarshalPrivateKeyToPEM(key)
	if err != nil {
		return err
	}

	return certutil.WriteKey(runtime.ServiceCurrentKey, keyData)
}

func createSigningCertKey(prefix, certFile, keyFile string) (bool, error) {
	if exists(certFile, keyFile) {
		return false, nil
	}

	caKeyBytes, _, err := certutil.LoadOrGenerateKeyFile(keyFile, false)
	if err != nil {
		return false, err
	}

	caKey, err := certutil.ParsePrivateKeyPEM(caKeyBytes)
	if err != nil {
		return false, err
	}

	cfg := certutil.Config{
		CommonName: fmt.Sprintf("%s-ca@%d", prefix, time.Now().Unix()),
	}

	cert, err := certutil.NewSelfSignedCACert(cfg, caKey.(crypto.Signer))
	if err != nil {
		return false, err
	}

	if err := certutil.WriteCert(certFile, certutil.EncodeCertPEM(cert)); err != nil {
		return false, err
	}
	return true, nil
}

func expired(certFile string) bool {
	certificates, err := certutil.CertsFromFile(certFile)
	if err != nil {
		return false
	}
	return certutil.IsCertExpired(certificates[0], config.CertificateRenewDays)
}

func genEncryptionConfigAndState(controlConfig *config.Control) error {
	runtime := controlConfig.Runtime
	if !controlConfig.EncryptSecrets {
		return nil
	}
	var keyName string
	switch controlConfig.EncryptProvider {
	case secretsencrypt.AESCBCProvider:
		keyName = "aescbckey"
	case secretsencrypt.SecretBoxProvider:
		keyName = "secretboxkey"
	default:
		return fmt.Errorf("unsupported secrets-encryption-key-type %s", controlConfig.EncryptProvider)
	}
	if s, err := os.Stat(runtime.EncryptionConfig); err == nil && s.Size() > 0 {
		// On upgrade from older versions, the encryption hash may not exist, create it
		if _, err := os.Stat(runtime.EncryptionHash); errors.Is(err, os.ErrNotExist) {
			curEncryptionByte, err := os.ReadFile(runtime.EncryptionConfig)
			if err != nil {
				return err
			}
			encryptionConfigHash := sha256.Sum256(curEncryptionByte)
			ann := "start-" + hex.EncodeToString(encryptionConfigHash[:])
			return os.WriteFile(controlConfig.Runtime.EncryptionHash, []byte(ann), 0600)
		}
		return nil
	}

	keyByte := make([]byte, secretsencrypt.KeySize)
	if _, err := rand.Read(keyByte); err != nil {
		return err
	}
	newKey := []apiserverconfigv1.Key{
		{
			Name:   keyName,
			Secret: base64.StdEncoding.EncodeToString(keyByte),
		},
	}
	var provider []apiserverconfigv1.ProviderConfiguration
	if controlConfig.EncryptProvider == secretsencrypt.AESCBCProvider {
		provider = []apiserverconfigv1.ProviderConfiguration{
			{
				AESCBC: &apiserverconfigv1.AESConfiguration{
					Keys: newKey,
				},
			},
			{
				Identity: &apiserverconfigv1.IdentityConfiguration{},
			},
		}
	} else if controlConfig.EncryptProvider == secretsencrypt.SecretBoxProvider {
		provider = []apiserverconfigv1.ProviderConfiguration{
			{
				Secretbox: &apiserverconfigv1.SecretboxConfiguration{
					Keys: newKey,
				},
			},
			{
				Identity: &apiserverconfigv1.IdentityConfiguration{},
			},
		}
	}

	encConfig := apiserverconfigv1.EncryptionConfiguration{
		TypeMeta: metav1.TypeMeta{
			Kind:       "EncryptionConfiguration",
			APIVersion: "apiserver.config.k8s.io/v1",
		},
		Resources: []apiserverconfigv1.ResourceConfiguration{
			{
				Resources: []string{"secrets"},
				Providers: provider,
			},
		},
	}
	b, err := json.Marshal(encConfig)
	if err != nil {
		return err
	}
	if err := util.AtomicWrite(runtime.EncryptionConfig, b, 0600); err != nil {
		return err
	}
	encryptionConfigHash := sha256.Sum256(b)
	ann := "start-" + hex.EncodeToString(encryptionConfigHash[:])
	return os.WriteFile(controlConfig.Runtime.EncryptionHash, []byte(ann), 0600)
}

func genEgressSelectorConfig(controlConfig *config.Control) error {
	var clusterConn apiserverv1beta1.Connection

	if controlConfig.EgressSelectorMode == config.EgressSelectorModeDisabled {
		clusterConn = apiserverv1beta1.Connection{
			ProxyProtocol: apiserverv1beta1.ProtocolDirect,
		}
	} else {
		clusterConn = apiserverv1beta1.Connection{
			ProxyProtocol: apiserverv1beta1.ProtocolHTTPConnect,
			Transport: &apiserverv1beta1.Transport{
				TCP: &apiserverv1beta1.TCPTransport{
					URL: fmt.Sprintf("https://%s:%d", controlConfig.BindAddressOrLoopback(false, true), controlConfig.SupervisorPort),
					TLSConfig: &apiserverv1beta1.TLSConfig{
						CABundle:   controlConfig.Runtime.ServerCA,
						ClientKey:  controlConfig.Runtime.ClientKubeAPIKey,
						ClientCert: controlConfig.Runtime.ClientKubeAPICert,
					},
				},
			},
		}
	}

	egressConfig := apiserverv1beta1.EgressSelectorConfiguration{
		TypeMeta: metav1.TypeMeta{
			Kind:       "EgressSelectorConfiguration",
			APIVersion: "apiserver.k8s.io/v1beta1",
		},
		EgressSelections: []apiserverv1beta1.EgressSelection{
			{
				Name:       "cluster",
				Connection: clusterConn,
			},
		},
	}

	b, err := json.Marshal(egressConfig)
	if err != nil {
		return err
	}
	return os.WriteFile(controlConfig.Runtime.EgressSelectorConfig, b, 0600)
}

func genCloudConfig(controlConfig *config.Control) error {
	cloudConfig := cloudprovider.Config{
		LBDefaultPriorityClassName: cloudprovider.DefaultLBPriorityClassName,
		LBEnabled:                  !controlConfig.DisableServiceLB,
		LBNamespace:                controlConfig.ServiceLBNamespace,
		LBImage:                    cloudprovider.DefaultLBImage,
		Rootless:                   controlConfig.Rootless,
		NodeEnabled:                !controlConfig.DisableCCM,
	}
	if controlConfig.SystemDefaultRegistry != "" {
		cloudConfig.LBImage = controlConfig.SystemDefaultRegistry + "/" + cloudConfig.LBImage
	}
	b, err := json.Marshal(cloudConfig)
	if err != nil {
		return err
	}
	return os.WriteFile(controlConfig.Runtime.CloudControllerConfig, b, 0600)
}
