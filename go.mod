module github.com/k3s-io/k3s

go 1.24.11

godebug gotypesalias=0

exclude golang.org/x/tools v0.38.0 // https://github.com/golang/go/issues/75888

replace (
	github.com/Mirantis/cri-dockerd => github.com/k3s-vip/cri-dockerd v0.4.1-k34vip0
	github.com/cloudnativelabs/kube-router/v2 => github.com/k3s-vip/kube-router/v2 v2.6.3-vip
	github.com/containerd/containerd/v2 => github.com/k3s-vip/containerd/v2 v2.1.6-upgrade
	github.com/containerd/nri => github.com/containerd/nri v0.9.0 // #161
	github.com/distribution/reference => github.com/distribution/reference v0.5.0
	github.com/docker/docker => github.com/docker/docker v27.5.1+incompatible
	github.com/google/cadvisor => github.com/google/cadvisor v0.54.1
	github.com/google/cel-go => github.com/google/cel-go v0.26.1
	github.com/opencontainers/runc => github.com/k3s-vip/runc v1.3.4-vip0
	github.com/rancher/dynamiclistener => github.com/k3s-vip/dynamiclistener v0.6.4-vip0
	github.com/spegel-org/spegel => github.com/k3s-vip/spegel v0.6.0-vip8
	go.etcd.io/etcd/api/v3 => github.com/k3s-vip/etcd/api/v3 v3.6.7-vip
	go.etcd.io/etcd/client/pkg/v3 => github.com/k3s-vip/etcd/client/pkg/v3 v3.6.7-vip
	go.etcd.io/etcd/client/v3 => github.com/k3s-vip/etcd/client/v3 v3.6.7-vip
	go.etcd.io/etcd/etcdutl/v3 => github.com/k3s-vip/etcd/etcdutl/v3 v3.6.7-vip
	go.etcd.io/etcd/pkg/v3 => github.com/k3s-vip/etcd/pkg/v3 v3.6.7-vip
	go.etcd.io/etcd/server/v3 => github.com/k3s-vip/etcd/server/v3 v3.6.7-vip
	go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc => go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc v0.60.0
	k8s.io/api => github.com/k3s-vip/kubernetes/staging/src/k8s.io/api v1.34.4-vip
	k8s.io/apiextensions-apiserver => github.com/k3s-vip/kubernetes/staging/src/k8s.io/apiextensions-apiserver v1.34.4-vip
	k8s.io/apimachinery => github.com/k3s-vip/kubernetes/staging/src/k8s.io/apimachinery v1.34.4-vip
	k8s.io/apiserver => github.com/k3s-vip/kubernetes/staging/src/k8s.io/apiserver v1.34.4-vip
	k8s.io/cli-runtime => github.com/k3s-vip/kubernetes/staging/src/k8s.io/cli-runtime v1.34.4-vip
	k8s.io/client-go => github.com/k3s-vip/kubernetes/staging/src/k8s.io/client-go v1.34.4-vip
	k8s.io/cloud-provider => github.com/k3s-vip/kubernetes/staging/src/k8s.io/cloud-provider v1.34.4-vip
	k8s.io/cluster-bootstrap => github.com/k3s-vip/kubernetes/staging/src/k8s.io/cluster-bootstrap v1.34.4-vip
	k8s.io/code-generator => github.com/k3s-vip/kubernetes/staging/src/k8s.io/code-generator v1.34.4-vip
	k8s.io/component-base => github.com/k3s-vip/kubernetes/staging/src/k8s.io/component-base v1.34.4-vip
	k8s.io/component-helpers => github.com/k3s-vip/kubernetes/staging/src/k8s.io/component-helpers v1.34.4-vip
	k8s.io/controller-manager => github.com/k3s-vip/kubernetes/staging/src/k8s.io/controller-manager v1.34.4-vip
	k8s.io/cri-api => github.com/k3s-vip/kubernetes/staging/src/k8s.io/cri-api v1.34.4-vip
	k8s.io/cri-client => github.com/k3s-vip/kubernetes/staging/src/k8s.io/cri-client v1.34.4-vip
	k8s.io/csi-translation-lib => github.com/k3s-vip/kubernetes/staging/src/k8s.io/csi-translation-lib v1.34.4-vip
	k8s.io/dynamic-resource-allocation => github.com/k3s-vip/kubernetes/staging/src/k8s.io/dynamic-resource-allocation v1.34.4-vip
	k8s.io/endpointslice => github.com/k3s-vip/kubernetes/staging/src/k8s.io/endpointslice v1.34.4-vip
	k8s.io/externaljwt => github.com/k3s-vip/kubernetes/staging/src/k8s.io/externaljwt v1.34.4-vip
	k8s.io/klog/v2 => github.com/k3s-vip/klog/v2 latest // main
	k8s.io/kms => github.com/k3s-vip/kubernetes/staging/src/k8s.io/kms v1.34.4-vip
	k8s.io/kube-aggregator => github.com/k3s-vip/kubernetes/staging/src/k8s.io/kube-aggregator v1.34.4-vip
	k8s.io/kube-controller-manager => github.com/k3s-vip/kubernetes/staging/src/k8s.io/kube-controller-manager v1.34.4-vip
	k8s.io/kube-openapi => k8s.io/kube-openapi v0.0.0-20250710124328-f3f2b991d03b
	k8s.io/kube-proxy => github.com/k3s-vip/kubernetes/staging/src/k8s.io/kube-proxy v1.34.4-vip
	k8s.io/kube-scheduler => github.com/k3s-vip/kubernetes/staging/src/k8s.io/kube-scheduler v1.34.4-vip
	k8s.io/kubectl => github.com/k3s-vip/kubernetes/staging/src/k8s.io/kubectl v1.34.4-vip
	k8s.io/kubelet => github.com/k3s-vip/kubernetes/staging/src/k8s.io/kubelet v1.34.4-vip
	k8s.io/kubernetes => github.com/k3s-vip/kubernetes v1.34.4-vip
	k8s.io/metrics => github.com/k3s-vip/kubernetes/staging/src/k8s.io/metrics v1.34.4-vip
	k8s.io/mount-utils => github.com/k3s-vip/kubernetes/staging/src/k8s.io/mount-utils v1.34.4-vip
	k8s.io/node-api => github.com/k3s-vip/kubernetes/staging/src/k8s.io/node-api v1.34.4-vip
	k8s.io/pod-security-admission => github.com/k3s-vip/kubernetes/staging/src/k8s.io/pod-security-admission v1.34.4-vip
	k8s.io/sample-apiserver => github.com/k3s-vip/kubernetes/staging/src/k8s.io/sample-apiserver v1.34.4-vip
	k8s.io/sample-cli-plugin => github.com/k3s-vip/kubernetes/staging/src/k8s.io/sample-cli-plugin v1.34.4-vip
	k8s.io/sample-controller => github.com/k3s-vip/kubernetes/staging/src/k8s.io/sample-controller v1.34.4-vip
	sigs.k8s.io/cri-tools => github.com/k3s-vip/cri-tools v1.34.4-vip
)

require (
	github.com/Microsoft/hcsshim latest // stable
	github.com/Mirantis/cri-dockerd latest // replace
	github.com/blang/semver/v4 v4.0.0
	github.com/cloudnativelabs/kube-router/v2 latest // replace
	github.com/containerd/cgroups/v3 v3.1.0 // for runtime-spec v1.2.1
	github.com/containerd/containerd/api latest // stable
	github.com/containerd/containerd/v2 latest // replace
	github.com/containerd/errdefs latest // stable
	github.com/containerd/fuse-overlayfs-snapshotter/v2 v2.1.6 // stable
	github.com/containerd/stargz-snapshotter v0.18.1 // latest
	github.com/containerd/zfs/v2 latest // stable
	github.com/coreos/go-iptables v0.8.0
	github.com/coreos/go-systemd/v22 v22.6.0
	github.com/distribution/reference latest // replace
	github.com/docker/docker latest // replace
	github.com/dustin/go-humanize v1.0.1
	github.com/dunglas/httpsfv v1.1.0
	github.com/erikdubbelboer/gspt v0.0.0-20190125194910-e68493906b83
	github.com/flannel-io/flannel latest // stable
	github.com/fsnotify/fsnotify v1.9.0
	github.com/go-logr/logr v1.4.3
	github.com/go-logr/stdr v1.2.3-0.20220714215716-96bad1d688c5
	github.com/go-test/deep v1.0.7
	github.com/google/cadvisor v0.54.1 // replace
	github.com/google/go-containerregistry latest // stable
	github.com/google/uuid latest // stable
	github.com/gorilla/mux v1.8.1
	github.com/gorilla/websocket v1.5.4-0.20250319132907-e064f32e3674
	github.com/inetaf/tcpproxy latest // stable
	github.com/ipfs/go-ds-leveldb v0.5.0
	github.com/ipfs/go-log/v2 v2.9.0
	github.com/joho/godotenv latest // stable
	github.com/json-iterator/go v1.1.12
	github.com/k3s-io/helm-controller latest // stable
	github.com/k3s-io/kine v0 // stable
	github.com/klauspost/compress v1.18.3
	github.com/klauspost/crc32 v1.3.0 // indirect
	github.com/libp2p/go-libp2p v0.46.0
	github.com/minio/minio-go/v7 latest // stable
	github.com/mwitkow/go-http-dialer v0.0.0-20161116154839-378f744fb2b8
	github.com/natefinch/lumberjack v2.0.0+incompatible
	github.com/onsi/ginkgo/v2 v2.25.1
	github.com/onsi/gomega v1.38.1
	github.com/opencontainers/cgroups v0.0.5 // stable
	github.com/opencontainers/runc latest // replace
	github.com/opencontainers/selinux v1.13.1 // latest
	github.com/otiai10/copy v1.7.0
	github.com/pkg/errors v0.9.1
	github.com/prometheus/client_golang latest // stable
	github.com/prometheus/common v0 // stable
	github.com/rancher/dynamiclistener v0 // replace
	github.com/rancher/lasso v0.2.5 // latest
	github.com/rancher/permissions latest // stable
	github.com/rancher/remotedialer latest // stable
	github.com/rancher/wharfie latest // stable
	github.com/rancher/wrangler/v3 latest // stable
	github.com/robfig/cron/v3 v3.0.1
	github.com/rootless-containers/rootlesskit/v2 v2.3.6 // stable
	github.com/sirupsen/logrus v1.9.4
	github.com/spegel-org/spegel latest // replace
	github.com/spf13/pflag v1.0.10
	github.com/stretchr/testify v1.11.1
	github.com/urfave/cli/v2 v2.27.7
	github.com/vishvananda/netlink latest // stable
	github.com/yl2chen/cidranger v1.0.2
	go.etcd.io/etcd/api/v3 latest // replace
	go.etcd.io/etcd/client/pkg/v3 latest // replace
	go.etcd.io/etcd/client/v3 latest // replace
	go.etcd.io/etcd/etcdutl/v3 latest // replace
	go.etcd.io/etcd/server/v3 latest // replace
	golang.org/x/crypto v0.45.0
	golang.org/x/mod v0.30.0
	golang.org/x/net v0.47.0
	golang.org/x/sync v0.19.0
	golang.org/x/sys v0.39.0
	google.golang.org/grpc v1.77.0
	gopkg.in/yaml.v2 latest // stable
	k8s.io/api v0.34.4 // replace
	k8s.io/apimachinery v0.34.4 // replace
	k8s.io/apiserver v0.34.4 // replace
	k8s.io/cli-runtime v0.34.4 // replace
	k8s.io/client-go v0.34.4 // replace
	k8s.io/cloud-provider v0.34.4 // replace
	k8s.io/cluster-bootstrap v0.34.4 // replace
	k8s.io/component-base v0.34.4 // replace
	k8s.io/component-helpers v0.34.4 // replace
	k8s.io/cri-api v0.34.4 // replace
	k8s.io/cri-client v0.34.4 // replace
	k8s.io/klog/v2 latest // replace
	k8s.io/kube-proxy v0.34.4 // replace
	k8s.io/kubectl v0.34.4 // replace
	k8s.io/kubelet v0.34.4 // replace
	k8s.io/kubernetes v1.34.4 // replace
	k8s.io/utils latest // stable
	sigs.k8s.io/cri-tools v1.34.0 // replace
	sigs.k8s.io/yaml latest // stable
)

require (
	github.com/antlr4-go/antlr/v4 v4.13.1 // indirect
	github.com/bahlo/generic-list-go v0.2.0 // indirect
	github.com/buger/jsonparser v1.1.1 // indirect
	github.com/fxamacker/cbor/v2 v2.9.0 // indirect
	github.com/go-task/slim-sprig/v3 latest // stable
	github.com/invopop/jsonschema v0.13.0 // indirect
	github.com/knqyf263/go-plugin latest // stable
	github.com/philhofer/fwd latest // stable
	github.com/stretchr/objx v0.5.2 // indirect
	github.com/tetratelabs/wazero latest // stable
	github.com/tinylib/msgp latest // stable
	github.com/wk8/go-ordered-map/v2 v2.1.8 // indirect
	github.com/x448/float16 v0.8.4 // indirect
	gopkg.in/evanphx/json-patch.v4 latest // stable
)

require (
	cel.dev/expr v0.24.0 // indirect
	dario.cat/mergo latest // stable
	filippo.io/edwards25519 latest // stable
	github.com/Azure/go-ansiterm v0.0.0-20250102033503-faa5f7b0171c // indirect
	github.com/JeffAshton/win_pdh v0.0.0-20161109143554-76bb4ee9f0ab // indirect
	github.com/MakeNowJust/heredoc v1.0.0 // indirect
	github.com/Masterminds/semver/v3 v3.4.0 // indirect
	github.com/Microsoft/go-winio v0.6.2 // indirect
	github.com/Microsoft/hnslib v0.1.1 // indirect
	github.com/NYTimes/gziphandler v1.1.1 // indirect
	github.com/Rican7/retry v0.3.1 // indirect
	github.com/antithesishq/antithesis-sdk-go v0.5.0 // indirect
	github.com/armon/circbuf v0.0.0-20190214190532-5111143e8da2 // indirect
	github.com/avast/retry-go/v4 v4.7.0 // indirect
	github.com/benbjohnson/clock v1.3.5 // indirect
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/blang/semver v3.5.1+incompatible // indirect
	github.com/bronze1man/goStrongswanVici v0.0.0-20231128135937-211cef3b0b20 // indirect
	github.com/cenkalti/backoff/v5 latest // stable
	github.com/cespare/xxhash/v2 v2.3.0 // indirect
	github.com/chai2010/gettext-go v1.0.2 // indirect
	github.com/checkpoint-restore/checkpointctl v1.3.0 // indirect
	github.com/checkpoint-restore/go-criu/v7 latest // stable
	github.com/cilium/ebpf v0.17.3 // indirect
	github.com/container-storage-interface/spec v1.9.0 // indirect
	github.com/containerd/btrfs/v2 latest // stable
	github.com/containerd/console latest // stable
	github.com/containerd/continuity latest // stable
	github.com/containerd/errdefs/pkg latest // stable
	github.com/containerd/fifo latest // stable
	github.com/containerd/go-cni latest // stable
	github.com/containerd/go-runc latest // stable
	github.com/containerd/imgcrypt/v2 latest // stable
	github.com/containerd/log latest // stable
	github.com/containerd/nri latest // replace
	github.com/containerd/otelttrpc latest // stable
	github.com/containerd/platforms latest // stable
	github.com/containerd/plugin latest // stable
	github.com/containerd/stargz-snapshotter/estargz v0.18.1 // latest
	github.com/containerd/ttrpc latest // stable
	github.com/containerd/typeurl/v2 latest // stable
	github.com/containernetworking/cni latest // stable
	github.com/containernetworking/plugins latest // stable
	github.com/containers/ocicrypt latest // stable
	github.com/coreos/go-oidc v2.3.0+incompatible // indirect
	github.com/coreos/go-semver v0.3.1 // indirect
	github.com/cpuguy83/go-md2man/v2 v2.0.7 // indirect
	github.com/cyphar/filepath-securejoin v0.5.2 // latest
	github.com/davecgh/go-spew v1.1.2-0.20180830191138-d8f796af33cc // indirect
	github.com/davidlazar/go-crypto v0.0.0-20200604182044-b73af7476f6c // indirect
	github.com/decred/dcrd/dcrec/secp256k1/v4 v4.4.0 // indirect
	github.com/docker/cli latest // stable
	github.com/docker/distribution latest // stable
	github.com/docker/docker-credential-helpers latest // stable
	github.com/docker/go-connections latest // stable
	github.com/docker/go-events latest // stable
	github.com/docker/go-metrics latest // stable
	github.com/docker/go-units latest // stable
	github.com/emicklei/go-restful v2.16.0+incompatible // indirect
	github.com/emicklei/go-restful/v3 v3.13.0 // indirect
	github.com/euank/go-kmsg-parser v2.0.0+incompatible // indirect
	github.com/evanphx/json-patch v5.9.11+incompatible // indirect
	github.com/exponent-io/jsonpath v0.0.0-20210407135951-1de76d718b3f // indirect
	github.com/expr-lang/expr latest // stable
	github.com/fatih/camelcase v1.0.0 // indirect
	github.com/felixge/httpsnoop v1.0.4 // indirect
	github.com/filecoin-project/go-clock v0.1.0 // indirect
	github.com/flynn/noise v1.1.0 // indirect
	github.com/gammazero/deque v1.2.0 // indirect
	github.com/ghodss/yaml v1.0.0 // indirect
	github.com/go-errors/errors v1.4.2 // indirect
	github.com/go-ini/ini v1.67.0 // latest
	github.com/go-jose/go-jose/v4 latest // stable
	github.com/go-openapi/jsonpointer v0.21.2 // latest
	github.com/go-openapi/jsonreference v0.21.0 // latest
	github.com/go-openapi/swag v0.23.1 // latest
	github.com/go-sql-driver/mysql latest // stable
	github.com/godbus/dbus/v5 v5.1.0 // indirect
	github.com/gofrs/flock v0.12.1 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang-jwt/jwt/v5 latest // stable
	github.com/golang/groupcache v0.0.0-20241129210726-2c02b8208cf8 // indirect
	github.com/golang/protobuf v1.5.4 // indirect
	github.com/golang/snappy v0.0.5-0.20231225225746-43d5d4cd4e0e // indirect
	github.com/google/btree latest // stable
	github.com/google/cel-go latest // replace
	github.com/google/gnostic-models latest // stable
	github.com/google/go-cmp latest // stable
	github.com/google/go-tpm latest // stable
	github.com/google/gopacket latest // stable
	github.com/google/pprof latest // stable
	github.com/gregjones/httpcache v0.0.0-20190611155906-901d90724c79 // indirect
	github.com/grpc-ecosystem/go-grpc-middleware/providers/prometheus latest // stable
	github.com/grpc-ecosystem/go-grpc-middleware/v2 latest // stable
	github.com/grpc-ecosystem/go-grpc-prometheus latest // stable
	github.com/grpc-ecosystem/grpc-gateway/v2 latest // stable
	github.com/guillaumemichel/reservedpool v0.3.0 // indirect
	github.com/hanwen/go-fuse/v2 v2.8.0 // indirect
	github.com/hashicorp/go-cleanhttp v0.5.2 // indirect
	github.com/hashicorp/go-retryablehttp v0.7.8 // indirect
	github.com/hashicorp/go-version v1.7.0 // indirect
	github.com/hashicorp/golang-lru v1.0.2 // indirect
	github.com/hashicorp/golang-lru/arc/v2 v2.0.7 // indirect
	github.com/hashicorp/golang-lru/v2 v2.0.7 // indirect
	github.com/huin/goupnp v1.3.0 // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/intel/goresctrl v0.8.0 // indirect
	github.com/ipfs/boxo v0.35.2 // indirect
	github.com/ipfs/go-block-format v0.2.3 // indirect
	github.com/ipfs/go-cid v0.6.0 // indirect
	github.com/ipfs/go-datastore v0.9.0 // indirect
	github.com/ipfs/go-test v0.2.3 // indirect
	github.com/ipld/go-ipld-prime v0.21.0 // indirect
	github.com/jackc/pgerrcode v0.0.0-20240316143900-6e2875d9b438 // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20240606120523-5a60cdf6a761 // indirect
	github.com/jackc/pgx/v5 v5.8.0 // indirect
	github.com/jackc/puddle/v2 latest // stable
	github.com/jackpal/go-nat-pmp v1.0.2 // indirect
	github.com/jbenet/go-temp-err-catcher v0.1.0 // indirect
	github.com/jonboulle/clockwork v0.5.0 // indirect
	github.com/josharian/intern v1.0.0 // indirect
	github.com/josharian/native v1.1.0 // indirect
	github.com/karrick/godirwalk v1.17.0 // indirect
	github.com/klauspost/cpuid/v2 v2.3.0 // indirect
	github.com/koron/go-ssdp v0.0.6 // indirect
	github.com/kylelemons/godebug v1.1.0 // indirect
	github.com/libopenstorage/openstorage v1.0.0 // indirect
	github.com/libp2p/go-buffer-pool v0.1.0 // indirect
	github.com/libp2p/go-cidranger v1.1.0 // indirect
	github.com/libp2p/go-flow-metrics v0.3.0 // indirect
	github.com/libp2p/go-libp2p-asn-util v0.4.1 // indirect
	github.com/libp2p/go-libp2p-kad-dht v0.36.0 // indirect
	github.com/libp2p/go-libp2p-kbucket v0.8.0 // indirect
	github.com/libp2p/go-libp2p-record v0.3.1 // indirect
	github.com/libp2p/go-libp2p-routing-helpers v0.7.5 // indirect
	github.com/libp2p/go-msgio v0.3.0 // indirect
	github.com/libp2p/go-netroute v0.3.0 // indirect
	github.com/libp2p/go-reuseport v0.4.0 // indirect
	github.com/libp2p/go-yamux/v5 latest // stable
	github.com/liggitt/tabwriter v0.0.0-20181228230101-89fcab3d43de // indirect
	github.com/lithammer/dedent v1.1.0 // indirect
	github.com/mailru/easyjson v0.9.1 // indirect
	github.com/marten-seemann/tcp v0.0.0-20210406111302-dfbc87cc63fd // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/mattn/go-sqlite3 latest // stable
	github.com/mdlayher/genetlink v1.3.2 // indirect
	github.com/mdlayher/netlink v1.7.2 // indirect
	github.com/mdlayher/socket v0.5.1 // indirect
	github.com/mdlayher/vsock v1.2.1 // indirect
	github.com/miekg/dns v1.1.69 // indirect
	github.com/miekg/pkcs11 v1.1.1 // indirect
	github.com/mikioh/tcpinfo v0.0.0-20190314235526-30a79bb1804b // indirect
	github.com/mikioh/tcpopt v0.0.0-20190314235656-172688c1accc // indirect
	github.com/minio/crc64nvme latest // stable
	github.com/minio/highwayhash latest // stable
	github.com/minio/md5-simd latest // stable
	github.com/minio/sha256-simd latest // stable
	github.com/mistifyio/go-zfs v2.1.2-0.20190413222219-f784269be439+incompatible // indirect
	github.com/mistifyio/go-zfs/v3 v3.0.1 // indirect
	github.com/mitchellh/go-homedir v1.1.0 // indirect
	github.com/mitchellh/go-wordwrap v1.0.1 // indirect
	github.com/moby/docker-image-spec latest // stable
	github.com/moby/ipvs latest // stable
	github.com/moby/locker latest // stable
	github.com/moby/spdystream latest // stable
	github.com/moby/sys/capability latest // stable
	github.com/moby/sys/mountinfo latest // stable
	github.com/moby/sys/reexec latest // stable
	github.com/moby/sys/sequential latest // stable
	github.com/moby/sys/signal latest // stable
	github.com/moby/sys/symlink latest // stable
	github.com/moby/sys/user latest // stable
	github.com/moby/sys/userns latest // stable
	github.com/moby/term latest // stable
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.3-0.20250322232337-35a7c28c31ee // indirect
	github.com/mohae/deepcopy v0.0.0-20170929034955-c48cc78d4826 // indirect
	github.com/monochromegane/go-gitignore v0.0.0-20200626010858-205db1a8cc00 // indirect
	github.com/morikuni/aec v1.0.0 // indirect
	github.com/mr-tron/base58 v1.2.0 // indirect
	github.com/multiformats/go-base32 v0.1.0 // indirect
	github.com/multiformats/go-base36 v0.2.0 // indirect
	github.com/multiformats/go-multiaddr v0.16.1 // indirect
	github.com/multiformats/go-multiaddr-dns v0.4.1 // indirect
	github.com/multiformats/go-multiaddr-fmt v0.1.0 // indirect
	github.com/multiformats/go-multibase v0.2.0 // indirect
	github.com/multiformats/go-multicodec v0.10.0 // indirect
	github.com/multiformats/go-multihash v0.2.3 // indirect
	github.com/multiformats/go-multistream v0.6.1 // indirect
	github.com/multiformats/go-varint v0.1.0 // indirect
	github.com/munnerz/goautoneg v0.0.0-20191010083416-a7dc8b61c822 // indirect
	github.com/mxk/go-flowrate v0.0.0-20140419014527-cca7078d478f // indirect
	github.com/nats-io/jsm.go v0.3.0 // indirect
	github.com/nats-io/jwt/v2 v2.8.0 // indirect
	github.com/nats-io/nats-server/v2 v2.12.2 // indirect
	github.com/nats-io/nats.go v1.48.0 // indirect
	github.com/nats-io/nkeys v0.4.11 // indirect
	github.com/nats-io/nuid v1.0.1 // indirect
	github.com/opencontainers/go-digest latest // stable
	github.com/opencontainers/image-spec latest // stable
	github.com/opencontainers/runtime-spec v1.2.1 // stable
	github.com/opencontainers/runtime-tools v0.9.1-0.20251111083745-e5b454202754 // for runtime-spec v1.2.1
	github.com/pbnjay/memory v0.0.0-20210728143218-7b4eea64cf58 // indirect
	github.com/pelletier/go-toml/v2 v2.2.4 // indirect
	github.com/peterbourgon/diskv v2.0.1+incompatible // indirect
	github.com/petermattis/goid latest // stable
	github.com/pierrec/lz4 v2.6.0+incompatible // indirect
	github.com/pion/datachannel v1.5.10 // pion/transport/v3
	github.com/pion/dtls/v2 latest // stable
	github.com/pion/dtls/v3 v3.0.9 // pion/transport/v3
	github.com/pion/ice/v4 v4.1.0 // pion/transport/v3
	github.com/pion/interceptor latest // stable
	github.com/pion/logging latest // stable
	github.com/pion/mdns/v2 latest // stable
	github.com/pion/randutil latest // stable
	github.com/pion/rtcp latest // stable
	github.com/pion/rtp latest // stable
	github.com/pion/sctp v1.9.0 // pion/transport/v3
	github.com/pion/sdp/v3 latest // stable
	github.com/pion/srtp/v3 v3.0.9 // pion/transport/v3
	github.com/pion/stun latest // stable
	github.com/pion/stun/v3 v3.1.0 // pion/transport/v3
	github.com/pion/transport/v2 latest // stable
	github.com/pion/transport/v3 latest // stable
	github.com/pion/turn/v4 v4.1.3 // pion/transport/v3
	github.com/pion/webrtc/v4 v4.2.1 // pion/transport/v3
	github.com/pmezard/go-difflib v1.0.1-0.20181226105442-5d4384ee4fb2 // indirect
	github.com/polydawn/refmt v0.89.0 // indirect
	github.com/pquerna/cachecontrol v0.1.0 // indirect
	github.com/probe-lab/go-libdht v0.4.0 // indirect
	github.com/prometheus/client_model latest // stable
	github.com/prometheus/procfs latest // stable
	github.com/quic-go/qpack v0.6.0 // indirect
	github.com/quic-go/quic-go v0.57.1 // indirect
	github.com/quic-go/webtransport-go v0.9.0 // indirect
	github.com/rs/xid v1.6.0 // indirect
	github.com/russross/blackfriday/v2 v2.1.0 // indirect
	github.com/sasha-s/go-deadlock latest // stable
	github.com/shengdoushi/base58 v1.0.0 // indirect
	github.com/smallstep/pkcs7 latest // stable
	github.com/soheilhy/cmux v0.1.5 // indirect
	github.com/spaolacci/murmur3 v1.1.0 // indirect
	github.com/spf13/cobra v1.9.1 // indirect
	github.com/stefanberger/go-pkcs11uri v0.0.0-20230803200340-78284954bff6 // indirect
	github.com/stoewer/go-strcase v1.3.1 // indirect
	github.com/syndtr/goleveldb v1.0.0 // indirect
	github.com/tchap/go-patricia/v2 v2.3.2 // indirect
	github.com/tidwall/btree v1.8.1 // indirect
	github.com/tmc/grpc-websocket-proxy v0.0.0-20220101234140-673ab2c3ae75 // indirect
	github.com/vbatts/tar-split v0.12.1 // indirect
	github.com/vishvananda/netns latest // stable
	github.com/whyrusleeping/go-keyspace v0.0.0-20160322163242-5b898ac5add1 // indirect
	github.com/wlynxg/anet latest // stable
	github.com/xiang90/probing v0.0.0-20221125231312-a49e3df8f510 // indirect
	github.com/xlab/treeprint v1.2.0 // indirect
	github.com/xrash/smetrics v0.0.0-20240521201337-686a1a2994c1 // indirect
	go.etcd.io/bbolt latest // stable
	go.etcd.io/etcd/pkg/v3 latest // replace
	go.etcd.io/raft/v3 latest // stable
	go.opencensus.io v0.24.0 // indirect
	go.opentelemetry.io/auto/sdk latest // stable
	go.opentelemetry.io/contrib/instrumentation/github.com/emicklei/go-restful/otelrestful v0.44.0 // indirect
	go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc v0.60.0 // indirect
	go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp v0.63.0 // indirect
	go.opentelemetry.io/otel v1.38.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlptrace v1.38.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc v1.38.0 // indirect
	go.opentelemetry.io/otel/metric v1.38.0 // indirect
	go.opentelemetry.io/otel/sdk v1.38.0 // indirect
	go.opentelemetry.io/otel/trace v1.38.0 // indirect
	go.opentelemetry.io/proto/otlp v1.7.1 // indirect
	go.uber.org/automaxprocs latest // stable
	go.uber.org/dig latest // stable
	go.uber.org/fx latest // stable
	go.uber.org/mock latest // stable
	go.uber.org/multierr latest // stable
	go.uber.org/zap latest // stable
	go.yaml.in/yaml/v2 latest // stable
	go.yaml.in/yaml/v3 latest // stable
	golang.org/x/exp v0.0.0-20251023183803-a4bb9ffd2546 // indirect
	golang.org/x/oauth2 v0.32.0 // indirect
	golang.org/x/telemetry master // latest
	golang.org/x/term v0.37.0 // indirect
	golang.org/x/text v0.31.0 // indirect
	golang.org/x/time v0.14.0 // indirect
	golang.org/x/tools v0.39.0 // indirect
	golang.zx2c4.com/wireguard v0.0.0-20231211153847-12269c276173 // indirect
	golang.zx2c4.com/wireguard/wgctrl v0.0.0-20241231184526-a9ab2273dd10 // indirect
	gonum.org/v1/gonum v0.16.0 // indirect
	google.golang.org/genproto/googleapis/api latest // stable
	google.golang.org/genproto/googleapis/rpc latest // stable
	google.golang.org/protobuf v1.36.10 // indirect
	gopkg.in/go-jose/go-jose.v2 latest // stable
	gopkg.in/inf.v0 latest // stable
	gopkg.in/natefinch/lumberjack.v2 latest // stable
	gopkg.in/yaml.v3 latest // stable
	k8s.io/apiextensions-apiserver v0.34.4 // replace
	k8s.io/code-generator v0.34.4 // replace
	k8s.io/controller-manager v0.34.4 // replace
	k8s.io/csi-translation-lib v0.34.4 // replace
	k8s.io/dynamic-resource-allocation v0.34.4 // replace
	k8s.io/endpointslice v0.34.4 // replace
	k8s.io/externaljwt v0.34.4 // replace
	k8s.io/gengo latest // stable
	k8s.io/gengo/v2 latest // stable
	k8s.io/kms v0.34.4 // replace
	k8s.io/kube-aggregator v0.34.4 // replace
	k8s.io/kube-controller-manager v0.34.4 // replace
	k8s.io/kube-openapi v0.0.0-20250710124328-f3f2b991d03b // replace
	k8s.io/kube-scheduler v0.34.4 // replace
	k8s.io/metrics v0.34.4 // replace
	k8s.io/mount-utils v0.34.4 // replace
	k8s.io/pod-security-admission v0.34.4 // replace
	lukechampine.com/blake3 v1.4.1 // indirect
	sigs.k8s.io/apiserver-network-proxy/konnectivity-client v0.34.0 // stable
	sigs.k8s.io/json latest // stable
	sigs.k8s.io/knftables latest // stable
	sigs.k8s.io/kustomize/api v0.20.1 // indirect
	sigs.k8s.io/kustomize/kustomize/v5 v5.7.1 // indirect
	sigs.k8s.io/kustomize/kyaml v0.20.1 // indirect
	sigs.k8s.io/randfill latest // stable
	sigs.k8s.io/structured-merge-diff/v6 latest // stable
	tags.cncf.io/container-device-interface v1.0.1 // indirect
	tags.cncf.io/container-device-interface/specs-go v1.0.0 // indirect
)
