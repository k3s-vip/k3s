package spegel

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync/atomic"
	"time"

	"github.com/k3s-io/k3s/pkg/clientaccess"
	"github.com/k3s-io/k3s/pkg/daemons/config"
	"github.com/k3s-io/k3s/pkg/util"
	"github.com/k3s-io/k3s/pkg/util/errors"
	"github.com/k3s-io/k3s/pkg/util/wait"
	"github.com/k3s-io/k3s/pkg/version"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/sirupsen/logrus"
	routing "github.com/spegel-org/spegel/pkg/routing/libp2p"
	"golang.org/x/sync/errgroup"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/labels"
	nodeutil "k8s.io/kubernetes/pkg/controller/util/node"
)

type Bootstrapper interface {
	routing.Bootstrapper
	Run(ctx context.Context, id peer.AddrInfo) error
}

// explicit interface checks
var _ Bootstrapper = &selfBootstrapper{}
var _ Bootstrapper = &notSelfBootstrapper{}
var _ Bootstrapper = &agentBootstrapper{}
var _ Bootstrapper = &serverBootstrapper{}
var _ Bootstrapper = &chainingBootstrapper{}

type selfBootstrapper struct {
	id atomic.Pointer[peer.AddrInfo]
}

// NewSelfBootstrapper returns a stub p2p bootstrapper that just returns its own ID
func NewSelfBootstrapper() Bootstrapper {
	return &selfBootstrapper{}
}

func (s *selfBootstrapper) Run(ctx context.Context, id peer.AddrInfo) error {
	s.id.Store(&id)
	return waitForDone(ctx)
}

func (s *selfBootstrapper) Get(ctx context.Context) ([]peer.AddrInfo, error) {
	id := s.id.Load()
	if id == nil {
		return nil, errors.New("p2p peer not ready")
	}
	return []peer.AddrInfo{*id}, nil
}

type notSelfBootstrapper struct {
	id atomic.Pointer[peer.AddrInfo]
	b  Bootstrapper
}

// NewNotSelfBootstrapper wraps an existing bootstrapper,
// and will never return a list of peers containing only itself.
// This is best used to prevent the spegel router from considering
// itself as "Ready" when it has no peers.
func NewNotSelfBootstrapper(b Bootstrapper) Bootstrapper {
	return &notSelfBootstrapper{
		b: b,
	}
}

func (ns *notSelfBootstrapper) Run(ctx context.Context, id peer.AddrInfo) error {
	ns.id.Store(&id)
	return ns.b.Run(ctx, id)
}

func (ns *notSelfBootstrapper) Get(ctx context.Context) ([]peer.AddrInfo, error) {
	peers, err := ns.b.Get(ctx)
	id := ns.id.Load()
	if err == nil && len(peers) == 1 && id != nil && peers[0].ID == id.ID {
		return nil, nil
	}
	return peers, err
}

type agentBootstrapper struct {
	server     string
	token      string
	clientCert string
	clientKey  string
	kubeConfig string
	info       atomic.Pointer[clientaccess.Info]
}

// NewAgentBootstrapper returns a p2p bootstrapper that retrieves a peer address from its server
func NewAgentBootstrapper(server, token, dataDir string) Bootstrapper {
	return &agentBootstrapper{
		clientCert: filepath.Join(dataDir, "agent", "client-kubelet.crt"),
		clientKey:  filepath.Join(dataDir, "agent", "client-kubelet.key"),
		kubeConfig: filepath.Join(dataDir, "agent", "kubelet.kubeconfig"),
		server:     server,
		token:      token,
	}
}

func (c *agentBootstrapper) Run(ctx context.Context, id peer.AddrInfo) error {
	if len(id.Addrs) == 0 {
		return errors.New("local peer has no addresses")
	}
	if c.server != "" && c.token != "" {
		withCert := clientaccess.WithClientCertificate(c.clientCert, c.clientKey)
		info, err := clientaccess.ParseAndValidateToken(c.server, c.token, withCert)
		if err != nil {
			return errors.WithMessage(err, "failed to validate join token")
		}
		c.info.Store(info)
	}

	client, err := util.GetClientSet(c.kubeConfig)
	if err != nil {
		return errors.WithMessage(err, "failed to create kubernetes client")
	}

	go wait.PollUntilContextCancel(ctx, 1*time.Second, true, func(ctx context.Context) (bool, error) {
		nodeName := os.Getenv("NODE_NAME")
		if nodeName == "" {
			return false, nil
		}
		b, err := id.MarshalJSON()
		if err != nil {
			return false, err
		}
		addresses := string(b)
		address := fmt.Sprintf("%s/p2p/%s", id.Addrs[0].String(), id.ID.String())

		patch := util.NewPatchList()
		patcher := util.NewPatcher[*v1.Node](client.CoreV1().Nodes())

		patch.Add(addresses, "metadata", "annotations", P2pMulAddrAnnotation)
		patch.Add(address, "metadata", "annotations", P2pAddressAnnotation)
		patch.Add("true", "metadata", "labels", P2pEnabledLabel)

		if _, err = patcher.Patch(ctx, patch, nodeName); err != nil {
			logrus.Debugf("Failed to update P2P address annotations and labels: %v", err)
			return false, nil
		}
		logrus.Infof("Node P2P address annotations and labels added: %s %s", address, addresses)
		return true, nil
	})
	return waitForDone(ctx)
}

func (c *agentBootstrapper) Get(ctx context.Context) ([]peer.AddrInfo, error) {
	if c.server == "" || c.token == "" {
		return nil, errors.New("cannot get addresses without server and token")
	}

	info := c.info.Load()
	if info == nil {
		return nil, errors.New("client not ready")
	}

	addr, err := info.Get("/v1-"+version.Program+"/p2p", clientaccess.WithHeader("Accept", "application/json"))
	if err != nil {
		return nil, err
	}

	// If the response cannot be decoded as a JSON list of addresses, fall back
	// to using it as a legacy single-address response.
	addrs := []string{}
	if err := json.Unmarshal(addr, &addrs); err != nil {
		addrs = append(addrs, string(addr))
	}

	addrInfos := []peer.AddrInfo{}
	for _, addr := range addrs {
		if addrInfo, err := peer.AddrInfoFromString(addr); err == nil {
			addrInfos = append(addrInfos, *addrInfo)
		}
	}
	return addrInfos, nil
}

type serverBootstrapper struct {
	controlConfig *config.Control
}

// NewServerBootstrapper returns a p2p bootstrapper that returns an address from the Kubernetes node list
func NewServerBootstrapper(controlConfig *config.Control) Bootstrapper {
	return &serverBootstrapper{
		controlConfig: controlConfig,
	}
}

func (s *serverBootstrapper) Run(ctx context.Context, id peer.AddrInfo) error {
	return waitForDone(ctx)
}

func (s *serverBootstrapper) Get(ctx context.Context) ([]peer.AddrInfo, error) {
	if s.controlConfig.Runtime.Core == nil {
		return nil, util.ErrCoreNotReady
	}
	nodeName := os.Getenv("NODE_NAME")
	if nodeName == "" {
		return nil, errors.New("node name not set")
	}

	nodes := s.controlConfig.Runtime.Core.Core().V1().Node()
	if !nodes.Informer().HasSynced() {
		return nil, errors.New("node cache informer not synced")
	}

	labelSelector := labels.Set{P2pEnabledLabel: "true"}.AsSelector()
	nodeList, err := nodes.Cache().List(labelSelector)
	if err != nil {
		return nil, err
	}

	addrs := []peer.AddrInfo{}
	for _, node := range nodeList {
		if node.Name == nodeName {
			// don't return our own address
			continue
		}
		if find, condition := nodeutil.GetNodeCondition(&node.Status, v1.NodeReady); find == -1 || condition.Status != v1.ConditionTrue {
			// don't return the address of a not-ready node
			continue
		}
		if val, ok := node.Annotations[P2pMulAddrAnnotation]; ok {
			info := &peer.AddrInfo{}
			if err := info.UnmarshalJSON([]byte(val)); err == nil {
				addrs = append(addrs, *info)
			}
		}
		if val, ok := node.Annotations[P2pAddressAnnotation]; ok {
			for _, addr := range strings.Split(val, ",") {
				if info, err := peer.AddrInfoFromString(addr); err == nil {
					addrs = append(addrs, *info)
				}
			}
		}
	}
	return addrs, nil
}

type chainingBootstrapper struct {
	bootstrappers []Bootstrapper
}

// NewChainingBootstrapper returns a p2p bootstrapper that passes through to a list of bootstrappers.
// Addressess are returned from the first bootstrapper that returns successfully, to prevent infinite
// recursion if this chains through an agentBootstrapper when the server URL is set to an external
// loadbalancer that points back at this node.
func NewChainingBootstrapper(bootstrappers ...Bootstrapper) Bootstrapper {
	return &chainingBootstrapper{
		bootstrappers: bootstrappers,
	}
}

func (c *chainingBootstrapper) Run(ctx context.Context, id peer.AddrInfo) error {
	eg, ctx := errgroup.WithContext(ctx)
	for i := range c.bootstrappers {
		b := c.bootstrappers[i]
		eg.Go(func() error {
			return b.Run(ctx, id)
		})
	}
	return eg.Wait()
}

func (c *chainingBootstrapper) Get(ctx context.Context) ([]peer.AddrInfo, error) {
	errs := []error{}
	for i := range c.bootstrappers {
		b := c.bootstrappers[i]
		as, err := b.Get(ctx)
		if err != nil {
			errs = append(errs, err)
		} else if len(as) != 0 {
			return as, nil
		}
	}
	return nil, errors.Join(errs...)
}

func waitForDone(ctx context.Context) error {
	<-ctx.Done()
	if err := ctx.Err(); err != nil && !errors.Is(err, context.Canceled) {
		return err
	}
	return nil
}
