module github.com/k3s-io/k3s

go 1.23.6

godebug gotypesalias=0

exclude golang.org/x/tools v0.38.0 // https://github.com/golang/go/issues/75888

replace (
	github.com/Microsoft/hcsshim => github.com/Microsoft/hcsshim v0.11.8
	github.com/Mirantis/cri-dockerd => github.com/k3s-vip/cri-dockerd v0.4.1-k29vip0
	github.com/cloudnativelabs/kube-router/v2 => github.com/k3s-vip/kube-router/v2 v2.6.3-vip
	github.com/containerd/containerd => github.com/k3s-vip/containerd v1.7.30-upgrade0
	github.com/containerd/nri => github.com/containerd/nri v0.9.0 // #161
	github.com/containerd/stargz-snapshotter => github.com/containerd/stargz-snapshotter v0.15.1
	github.com/distribution/reference => github.com/distribution/reference v0.5.0
	github.com/docker/docker => github.com/docker/docker v27.5.1+incompatible
	github.com/flannel-io/flannel => github.com/k3s-vip/flannel v0.28.1-vip
	github.com/google/cadvisor => github.com/google/cadvisor v0.54.1
	github.com/google/cel-go => github.com/google/cel-go v0.17.8
	github.com/k3s-io/kine => github.com/k3s-vip/kine v0.14.11-vip
	github.com/kubernetes-sigs/cri-tools => github.com/k3s-vip/cri-tools v1.29.15-eol
	github.com/opencontainers/runc => github.com/k3s-vip/runc v1.3.4-vip0
	github.com/rancher/dynamiclistener => github.com/k3s-vip/dynamiclistener v0.6.4-downgrade0
	github.com/rancher/wrangler => github.com/rancher/wrangler v1.1.1-0.20230818201331-3604a6be798d
	github.com/spegel-org/spegel => github.com/k3s-vip/spegel v0.6.0-downgrade8
	go.etcd.io/etcd/api/v3 => github.com/k3s-vip/etcd/api/v3 v3.6.7-vip
	go.etcd.io/etcd/client/pkg/v3 => github.com/k3s-vip/etcd/client/pkg/v3 v3.6.7-vip
	go.etcd.io/etcd/client/v3 => github.com/k3s-vip/etcd/client/v3 v3.6.7-vip
	go.etcd.io/etcd/etcdutl/v3 => github.com/k3s-vip/etcd/etcdutl/v3 v3.6.7-vip
	go.etcd.io/etcd/pkg/v3 => github.com/k3s-vip/etcd/pkg/v3 v3.6.7-vip
	go.etcd.io/etcd/server/v3 => github.com/k3s-vip/etcd/server/v3 v3.6.7-vip
	go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc => go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc v0.60.0
	k8s.io/api => github.com/k3s-vip/kubernetes/staging/src/k8s.io/api v1.29.15-vip
	k8s.io/apiextensions-apiserver => github.com/k3s-vip/kubernetes/staging/src/k8s.io/apiextensions-apiserver v1.29.15-vip
	k8s.io/apimachinery => github.com/k3s-vip/kubernetes/staging/src/k8s.io/apimachinery v1.29.15-vip
	k8s.io/apiserver => github.com/k3s-vip/kubernetes/staging/src/k8s.io/apiserver v1.29.15-vip
	k8s.io/cli-runtime => github.com/k3s-vip/kubernetes/staging/src/k8s.io/cli-runtime v1.29.15-vip
	k8s.io/client-go => github.com/k3s-vip/kubernetes/staging/src/k8s.io/client-go v1.29.15-vip
	k8s.io/cloud-provider => github.com/k3s-vip/kubernetes/staging/src/k8s.io/cloud-provider v1.29.15-vip
	k8s.io/cluster-bootstrap => github.com/k3s-vip/kubernetes/staging/src/k8s.io/cluster-bootstrap v1.29.15-vip
	k8s.io/code-generator => github.com/k3s-vip/kubernetes/staging/src/k8s.io/code-generator v1.29.15-vip
	k8s.io/component-base => github.com/k3s-vip/kubernetes/staging/src/k8s.io/component-base v1.29.15-vip
	k8s.io/component-helpers => github.com/k3s-vip/kubernetes/staging/src/k8s.io/component-helpers v1.29.15-vip
	k8s.io/controller-manager => github.com/k3s-vip/kubernetes/staging/src/k8s.io/controller-manager v1.29.15-vip
	k8s.io/cri-api => github.com/k3s-vip/kubernetes/staging/src/k8s.io/cri-api v1.29.15-vip
	k8s.io/csi-translation-lib => github.com/k3s-vip/kubernetes/staging/src/k8s.io/csi-translation-lib v1.29.15-vip
	k8s.io/dynamic-resource-allocation => github.com/k3s-vip/kubernetes/staging/src/k8s.io/dynamic-resource-allocation v1.29.15-vip
	k8s.io/endpointslice => github.com/k3s-vip/kubernetes/staging/src/k8s.io/endpointslice v1.29.15-vip
	k8s.io/klog/v2 => github.com/k3s-vip/klog/v2 latest // main
	k8s.io/kms => github.com/k3s-vip/kubernetes/staging/src/k8s.io/kms v1.29.15-vip
	k8s.io/kube-aggregator => github.com/k3s-vip/kubernetes/staging/src/k8s.io/kube-aggregator v1.29.15-vip
	k8s.io/kube-controller-manager => github.com/k3s-vip/kubernetes/staging/src/k8s.io/kube-controller-manager v1.29.15-vip
	k8s.io/kube-openapi => k8s.io/kube-openapi v0.0.0-20250701173324-9bd5c66d9911
	k8s.io/kube-proxy => github.com/k3s-vip/kubernetes/staging/src/k8s.io/kube-proxy v1.29.15-vip
	k8s.io/kube-scheduler => github.com/k3s-vip/kubernetes/staging/src/k8s.io/kube-scheduler v1.29.15-vip
	k8s.io/kubectl => github.com/k3s-vip/kubernetes/staging/src/k8s.io/kubectl v1.29.15-vip
	k8s.io/kubelet => github.com/k3s-vip/kubernetes/staging/src/k8s.io/kubelet v1.29.15-vip
	k8s.io/kubernetes => github.com/k3s-vip/kubernetes v1.29.15-vip
	k8s.io/legacy-cloud-providers => github.com/k3s-vip/kubernetes/staging/src/k8s.io/legacy-cloud-providers v1.29.15-vip
	k8s.io/metrics => github.com/k3s-vip/kubernetes/staging/src/k8s.io/metrics v1.29.15-vip
	k8s.io/mount-utils => github.com/k3s-vip/kubernetes/staging/src/k8s.io/mount-utils v1.29.15-vip
	k8s.io/node-api => github.com/k3s-vip/kubernetes/staging/src/k8s.io/node-api v1.29.15-vip
	k8s.io/pod-security-admission => github.com/k3s-vip/kubernetes/staging/src/k8s.io/pod-security-admission v1.29.15-vip
	k8s.io/sample-apiserver => github.com/k3s-vip/kubernetes/staging/src/k8s.io/sample-apiserver v1.29.15-vip
	k8s.io/sample-cli-plugin => github.com/k3s-vip/kubernetes/staging/src/k8s.io/sample-cli-plugin v1.29.15-vip
	k8s.io/sample-controller => github.com/k3s-vip/kubernetes/staging/src/k8s.io/sample-controller v1.29.15-vip
	sigs.k8s.io/structured-merge-diff/v4 => sigs.k8s.io/structured-merge-diff/v4 v4.4.1
)

require (
	github.com/Microsoft/hcsshim latest // replace
	github.com/Mirantis/cri-dockerd latest // replace
	github.com/blang/semver/v4 v4.0.0
	github.com/cloudnativelabs/kube-router/v2 latest // replace
	github.com/containerd/aufs latest // stable
	github.com/containerd/cgroups/v3 v3.1.0 // for runtime-spec v1.2.1
	github.com/containerd/containerd latest // replace
	github.com/containerd/fuse-overlayfs-snapshotter latest // stable
	github.com/containerd/stargz-snapshotter latest // replace
	github.com/containerd/zfs latest // stable
	github.com/coreos/go-iptables v0.8.0
	github.com/coreos/go-systemd/v22 v22.5.0
	github.com/docker/docker latest // replace
	github.com/erikdubbelboer/gspt v0.0.0-20190125194910-e68493906b83
	github.com/flannel-io/flannel latest // replace
	github.com/fsnotify/fsnotify v1.7.0
	github.com/go-logr/logr v1.4.2
	github.com/go-logr/stdr v1.2.3-0.20220714215716-96bad1d688c5
	github.com/go-test/deep v1.0.7
	github.com/golang/mock v1.6.0
	github.com/google/cadvisor v0.54.1 // replace
	github.com/google/go-containerregistry latest // stable
	github.com/google/uuid latest // stable
	github.com/gorilla/mux v1.8.1
	github.com/gorilla/websocket v1.5.3
	github.com/inetaf/tcpproxy latest // stable
	github.com/ipfs/go-ds-leveldb v0.5.0
	github.com/ipfs/go-log/v2 v2.5.1
	github.com/joho/godotenv latest // stable
	github.com/json-iterator/go v1.1.12
	github.com/k3s-io/helm-controller v0.15.17 // latest
	github.com/k3s-io/kine v0 // stable
	github.com/klauspost/compress v1.18.0
	github.com/klauspost/crc32 v1.3.0 // indirect
	github.com/kubernetes-sigs/cri-tools v1.29.0 // replace
	github.com/libp2p/go-libp2p v0.38.2
	github.com/minio/minio-go/v7 latest // stable
	github.com/mwitkow/go-http-dialer v0.0.0-20161116154839-378f744fb2b8
	github.com/natefinch/lumberjack v2.0.0+incompatible
	github.com/onsi/ginkgo/v2 v2.22.0
	github.com/onsi/gomega v1.36.0
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
	github.com/rancher/wrangler latest // replace
	github.com/robfig/cron/v3 v3.0.1
	github.com/rootless-containers/rootlesskit/v2 v2.3.6 // stable
	github.com/sirupsen/logrus v1.9.3
	github.com/spegel-org/spegel latest // replace
	github.com/spf13/pflag v1.0.6
	github.com/stretchr/objx v0.5.2 // indirect
	github.com/stretchr/testify v1.10.0
	github.com/urfave/cli v1.22.15
	github.com/vishvananda/netlink latest // stable
	github.com/yl2chen/cidranger v1.0.2
	go.etcd.io/etcd/api/v3 latest // replace
	go.etcd.io/etcd/client/pkg/v3 latest // replace
	go.etcd.io/etcd/client/v3 latest // replace
	go.etcd.io/etcd/etcdutl/v3 latest // replace
	go.etcd.io/etcd/server/v3 latest // replace
	golang.org/x/crypto v0.37.0
	golang.org/x/mod v0.22.0
	golang.org/x/net v0.39.0
	golang.org/x/sync v0.13.0
	golang.org/x/sys v0.32.0
	google.golang.org/grpc v1.71.1
	gopkg.in/yaml.v2 latest // stable
	k8s.io/api v0.29.15 // replace
	k8s.io/apimachinery v0.29.15 // replace
	k8s.io/apiserver v0.29.15 // replace
	k8s.io/cli-runtime v0.29.15 // replace
	k8s.io/client-go v0.29.15 // replace
	k8s.io/cloud-provider v0.29.15 // replace
	k8s.io/cluster-bootstrap v0.29.15 // replace
	k8s.io/component-base v0.29.15 // replace
	k8s.io/component-helpers v0.29.15 // replace
	k8s.io/cri-api v0.29.15 // replace
	k8s.io/klog/v2 latest // replace
	k8s.io/kube-proxy v0.29.15 // replace
	k8s.io/kubectl v1.29.15 // replace
	k8s.io/kubernetes v0.29.15 // replace
	k8s.io/utils latest // stable
	sigs.k8s.io/yaml latest // stable
)

require (
	github.com/knqyf263/go-plugin latest // stable
	github.com/philhofer/fwd latest // stable
	github.com/tetratelabs/wazero latest // stable
	github.com/tinylib/msgp latest // stable
)

require (
	cloud.google.com/go/auth latest // stable
	cloud.google.com/go/auth/oauth2adapt latest // stable
	cloud.google.com/go/compute/metadata v0.6.0 // indirect
	dario.cat/mergo latest // stable
	filippo.io/edwards25519 latest // stable
	github.com/AdaLogics/go-fuzz-headers v0.0.0-20230811130428-ced1acdcaa24 // indirect
	github.com/AdamKorcz/go-118-fuzz-build v0.0.0-20230306123547-8075edf89bb0 // indirect
	github.com/Azure/azure-sdk-for-go v68.0.0+incompatible // indirect
	github.com/Azure/go-ansiterm v0.0.0-20210617225240-d185dfc1b5a1 // indirect
	github.com/Azure/go-autorest v14.2.0+incompatible // indirect
	github.com/Azure/go-autorest/autorest v0.11.29 // indirect
	github.com/Azure/go-autorest/autorest/adal v0.9.23 // indirect
	github.com/Azure/go-autorest/autorest/date v0.3.0 // indirect
	github.com/Azure/go-autorest/autorest/mocks v0.4.2 // indirect
	github.com/Azure/go-autorest/autorest/to v0.4.0 // indirect
	github.com/Azure/go-autorest/autorest/validation v0.3.1 // indirect
	github.com/Azure/go-autorest/logger v0.2.1 // indirect
	github.com/Azure/go-autorest/tracing v0.6.0 // indirect
	github.com/GoogleCloudPlatform/k8s-cloud-provider v1.18.1-0.20220218231025-f11817397a1b // indirect
	github.com/JeffAshton/win_pdh v0.0.0-20161109143554-76bb4ee9f0ab // indirect
	github.com/MakeNowJust/heredoc v1.0.0 // indirect
	github.com/Masterminds/semver/v3 v3.4.0 // indirect
	github.com/Microsoft/go-winio v0.6.2 // indirect
	github.com/NYTimes/gziphandler v1.1.1 // indirect
	github.com/Rican7/retry v0.3.1 // indirect
	github.com/antithesishq/antithesis-sdk-go v0.5.0 // indirect
	github.com/antlr/antlr4/runtime/Go/antlr/v4 v4.0.0-20230305170008-8188dc5388df // indirect
	github.com/armon/circbuf v0.0.0-20150827004946-bbbad097214e // indirect
	github.com/avast/retry-go/v4 v4.6.1 // indirect
	github.com/benbjohnson/clock v1.3.5 // indirect
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/blang/semver v3.5.1+incompatible // indirect
	github.com/bronze1man/goStrongswanVici v0.0.0-20231128135937-211cef3b0b20 // indirect
	github.com/cenkalti/backoff/v5 latest // stable
	github.com/cespare/xxhash/v2 v2.3.0 // indirect
	github.com/chai2010/gettext-go v1.0.2 // indirect
	github.com/cilium/ebpf v0.16.0 // indirect
	github.com/container-storage-interface/spec v1.8.0 // indirect
	github.com/containerd/btrfs/v2 latest // stable
	github.com/containerd/cgroups latest // stable
	github.com/containerd/console latest // stable
	github.com/containerd/containerd/api latest // stable
	github.com/containerd/continuity latest // stable
	github.com/containerd/errdefs latest // stable
	github.com/containerd/errdefs/pkg latest // stable
	github.com/containerd/fifo latest // stable
	github.com/containerd/go-cni latest // stable
	github.com/containerd/go-runc latest // stable
	github.com/containerd/imgcrypt latest // stable
	github.com/containerd/log latest // stable
	github.com/containerd/nri latest // replace
	github.com/containerd/platforms latest // stable
	github.com/containerd/stargz-snapshotter/estargz v0.18.1 // latest
	github.com/containerd/ttrpc latest // stable
	github.com/containerd/typeurl latest // stable
	github.com/containerd/typeurl/v2 latest // stable
	github.com/containernetworking/cni latest // stable
	github.com/containernetworking/plugins latest // stable
	github.com/containers/ocicrypt latest // stable
	github.com/coreos/go-oidc v2.2.1+incompatible // indirect
	github.com/coreos/go-semver v0.3.1 // indirect
	github.com/cpuguy83/go-md2man/v2 v2.0.5 // indirect
	github.com/cyphar/filepath-securejoin v0.5.2 // stable
	github.com/danwinship/knftables v0.0.13 // indirect
	github.com/davecgh/go-spew v1.1.2-0.20180830191138-d8f796af33cc // indirect
	github.com/daviddengcn/go-colortext v1.0.0 // indirect
	github.com/davidlazar/go-crypto v0.0.0-20200604182044-b73af7476f6c // indirect
	github.com/decred/dcrd/dcrec/secp256k1/v4 v4.3.0 // indirect
	github.com/distribution/reference latest // replace
	github.com/docker/cli latest // stable
	github.com/docker/distribution latest // stable
	github.com/docker/docker-credential-helpers latest // stable
	github.com/docker/go-connections latest // stable
	github.com/docker/go-events latest // stable
	github.com/docker/go-metrics latest // stable
	github.com/docker/go-units latest // stable
	github.com/dustin/go-humanize v1.0.1 // indirect
	github.com/dunglas/httpsfv v1.1.0
	github.com/emicklei/go-restful v2.16.0+incompatible // indirect
	github.com/emicklei/go-restful/v3 v3.12.2 // indirect
	github.com/euank/go-kmsg-parser v2.0.0+incompatible // indirect
	github.com/evanphx/json-patch v5.6.0+incompatible // indirect
	github.com/exponent-io/jsonpath v0.0.0-20151013193312-d6023ce2651d // indirect
	github.com/expr-lang/expr latest // stable
	github.com/fatih/camelcase v1.0.0 // indirect
	github.com/felixge/httpsnoop v1.0.4 // indirect
	github.com/filecoin-project/go-clock v0.1.0 // indirect
	github.com/flynn/noise v1.1.0 // indirect
	github.com/fvbommel/sortorder v1.1.0 // indirect
	github.com/gammazero/deque v1.2.0 // indirect
	github.com/ghodss/yaml v1.0.0 // indirect
	github.com/go-errors/errors v1.4.2 // indirect
	github.com/go-ini/ini v1.66.6 // latest
	github.com/go-jose/go-jose/v4 latest // stable
	github.com/go-openapi/jsonpointer v0.21.1 // indirect
	github.com/go-openapi/jsonreference v0.21.0 // indirect
	github.com/go-openapi/swag v0.23.1 // indirect
	github.com/go-sql-driver/mysql latest // stable
	github.com/go-task/slim-sprig/v3 latest // stable
	github.com/godbus/dbus/v5 v5.1.0 // indirect
	github.com/gofrs/flock v0.8.1 // indirect
	github.com/gofrs/uuid v4.4.0+incompatible // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang-jwt/jwt/v4 latest // stable
	github.com/golang-jwt/jwt/v5 latest // stable
	github.com/golang/groupcache v0.0.0-20241129210726-2c02b8208cf8 // indirect
	github.com/golang/protobuf v1.5.4 // indirect
	github.com/golang/snappy v0.0.4 // indirect
	github.com/google/btree latest // stable
	github.com/google/cel-go latest // replace
	github.com/google/gnostic-models latest // stable
	github.com/google/go-cmp latest // stable
	github.com/google/go-tpm latest // stable
	github.com/google/gofuzz latest // stable
	github.com/google/gopacket latest // stable
	github.com/google/pprof latest // stable
	github.com/google/s2a-go latest // stable
	github.com/google/shlex latest // stable
	github.com/googleapis/enterprise-certificate-proxy latest // stable
	github.com/googleapis/gax-go/v2 v2.14.1 // indirect
	github.com/gregjones/httpcache v0.0.0-20180305231024-9cad4c3443a7 // indirect
	github.com/grpc-ecosystem/go-grpc-middleware latest // stable
	github.com/grpc-ecosystem/go-grpc-middleware/providers/prometheus latest // stable
	github.com/grpc-ecosystem/go-grpc-middleware/v2 latest // stable
	github.com/grpc-ecosystem/go-grpc-prometheus latest // stable
	github.com/grpc-ecosystem/grpc-gateway/v2 latest // stable
	github.com/guillaumemichel/reservedpool v0.3.0 // indirect
	github.com/hanwen/go-fuse/v2 v2.5.1 // indirect
	github.com/hashicorp/errwrap v1.1.0 // indirect
	github.com/hashicorp/go-cleanhttp v0.5.2 // indirect
	github.com/hashicorp/go-multierror v1.1.1 // indirect
	github.com/hashicorp/go-retryablehttp v0.7.7 // indirect
	github.com/hashicorp/go-version v1.7.0 // indirect
	github.com/hashicorp/golang-lru v1.0.2 // indirect
	github.com/hashicorp/golang-lru/arc/v2 v2.0.7 // indirect
	github.com/hashicorp/golang-lru/v2 v2.0.7 // indirect
	github.com/huin/goupnp v1.3.0 // indirect
	github.com/imdario/mergo v0.3.16 // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/intel/goresctrl v0.7.0 // indirect
	github.com/ipfs/boxo v0.27.0 // indirect
	github.com/ipfs/go-block-format v0.2.3 // indirect
	github.com/ipfs/go-cid v0.4.1 // indirect
	github.com/ipfs/go-datastore v0.6.0 // indirect
	github.com/ipfs/go-test v0.2.3 // indirect
	github.com/ipld/go-ipld-prime v0.21.0 // indirect
	github.com/jackc/pgerrcode v0.0.0-20240316143900-6e2875d9b438 // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20240606120523-5a60cdf6a761 // indirect
	github.com/jackc/pgx/v5 v5.7.4 // indirect
	github.com/jackc/puddle/v2 latest // stable
	github.com/jackpal/go-nat-pmp v1.0.2 // indirect
	github.com/jbenet/go-temp-err-catcher v0.1.0 // indirect
	github.com/jonboulle/clockwork v0.5.0 // indirect
	github.com/josharian/intern v1.0.0 // indirect
	github.com/josharian/native v1.1.0 // indirect
	github.com/karrick/godirwalk v1.17.0 // indirect
	github.com/klauspost/cpuid/v2 v2.2.9 // indirect
	github.com/koron/go-ssdp v0.0.4 // indirect
	github.com/kylelemons/godebug v1.1.0 // indirect
	github.com/libopenstorage/openstorage v1.0.0 // indirect
	github.com/libp2p/go-buffer-pool v0.1.0 // indirect
	github.com/libp2p/go-cidranger v1.1.0 // indirect
	github.com/libp2p/go-flow-metrics v0.2.0 // indirect
	github.com/libp2p/go-libp2p-asn-util v0.4.1 // indirect
	github.com/libp2p/go-libp2p-kad-dht v0.28.2 // indirect
	github.com/libp2p/go-libp2p-kbucket v0.6.4 // indirect
	github.com/libp2p/go-libp2p-record v0.2.0 // indirect
	github.com/libp2p/go-libp2p-routing-helpers v0.7.4 // indirect
	github.com/libp2p/go-msgio v0.3.0 // indirect
	github.com/libp2p/go-netroute v0.2.2 // indirect
	github.com/libp2p/go-reuseport v0.4.0 // indirect
	github.com/libp2p/go-yamux/v5 latest // stable
	github.com/liggitt/tabwriter v0.0.0-20181228230101-89fcab3d43de // indirect
	github.com/lithammer/dedent v1.1.0 // indirect
	github.com/mailru/easyjson v0.9.0 // indirect
	github.com/marten-seemann/tcp v0.0.0-20210406111302-dfbc87cc63fd // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/mattn/go-sqlite3 latest // stable
	github.com/mdlayher/genetlink v1.3.2 // indirect
	github.com/mdlayher/netlink v1.7.2 // indirect
	github.com/mdlayher/socket v0.5.1 // indirect
	github.com/miekg/dns v1.1.62 // indirect
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
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/mohae/deepcopy v0.0.0-20170929034955-c48cc78d4826 // indirect
	github.com/monochromegane/go-gitignore v0.0.0-20200626010858-205db1a8cc00 // indirect
	github.com/morikuni/aec v1.0.0 // indirect
	github.com/mr-tron/base58 v1.2.0 // indirect
	github.com/multiformats/go-base32 v0.1.0 // indirect
	github.com/multiformats/go-base36 v0.2.0 // indirect
	github.com/multiformats/go-multiaddr v0.14.0 // indirect
	github.com/multiformats/go-multiaddr-dns v0.4.1 // indirect
	github.com/multiformats/go-multiaddr-fmt v0.1.0 // indirect
	github.com/multiformats/go-multibase v0.2.0 // indirect
	github.com/multiformats/go-multicodec v0.9.0 // indirect
	github.com/multiformats/go-multihash v0.2.3 // indirect
	github.com/multiformats/go-multistream v0.6.0 // indirect
	github.com/multiformats/go-varint v0.0.7 // indirect
	github.com/munnerz/goautoneg v0.0.0-20191010083416-a7dc8b61c822 // indirect
	github.com/mxk/go-flowrate v0.0.0-20140419014527-cca7078d478f // indirect
	github.com/nats-io/jsm.go v0.0.31-0.20220317133147-fe318f464eee // indirect
	github.com/nats-io/jwt/v2 v2.7.3 // indirect
	github.com/nats-io/nats-server/v2 v2.10.27 // indirect
	github.com/nats-io/nats.go v1.39.1 // indirect
	github.com/nats-io/nkeys v0.4.10 // indirect
	github.com/nats-io/nuid v1.0.1 // indirect
	github.com/opencontainers/go-digest latest // stable
	github.com/opencontainers/image-spec latest // stable
	github.com/opencontainers/runtime-spec v1.2.1 // stable
	github.com/opencontainers/runtime-tools v0.9.1-0.20251111083745-e5b454202754 // for runtime-spec v1.2.1
	github.com/pbnjay/memory v0.0.0-20210728143218-7b4eea64cf58 // indirect
	github.com/pelletier/go-toml v1.9.5 // indirect
	github.com/pelletier/go-toml/v2 v2.2.3 // indirect
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
	github.com/quic-go/qpack v0.5.1 // indirect
	github.com/quic-go/quic-go v0.48.2 // indirect
	github.com/quic-go/webtransport-go v0.8.1-0.20241018022711-4ac2c9250e66 // indirect
	github.com/rs/xid v1.6.0 // indirect
	github.com/rubiojr/go-vhd v0.0.0-20200706105327-02e210299021 // indirect
	github.com/russross/blackfriday/v2 v2.1.0 // indirect
	github.com/sasha-s/go-deadlock latest // stable
	github.com/shengdoushi/base58 v1.0.0 // indirect
	github.com/smallstep/pkcs7 latest // stable
	github.com/soheilhy/cmux v0.1.5 // indirect
	github.com/spaolacci/murmur3 v1.1.0 // indirect
	github.com/spf13/cobra v1.8.1 // indirect
	github.com/stefanberger/go-pkcs11uri v0.0.0-20230803200340-78284954bff6 // indirect
	github.com/stoewer/go-strcase v1.3.0 // indirect
	github.com/syndtr/goleveldb v1.0.0 // indirect
	github.com/tchap/go-patricia/v2 v2.3.1 // indirect
	github.com/tidwall/btree v1.6.0 // indirect
	github.com/tmc/grpc-websocket-proxy v0.0.0-20220101234140-673ab2c3ae75 // indirect
	github.com/urfave/cli/v2 v2.27.6 // indirect
	github.com/vbatts/tar-split v0.11.5 // indirect
	github.com/vishvananda/netns latest // stable
	github.com/vmware/govmomi v0.30.6 // indirect
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
	go.opentelemetry.io/contrib/instrumentation/github.com/emicklei/go-restful/otelrestful v0.53.0 // indirect
	go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc v0.54.0 // indirect
	go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp v0.56.0 // indirect
	go.opentelemetry.io/otel v1.34.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlptrace v1.31.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc v1.31.0 // indirect
	go.opentelemetry.io/otel/metric v1.34.0 // indirect
	go.opentelemetry.io/otel/sdk v1.34.0 // indirect
	go.opentelemetry.io/otel/trace v1.34.0 // indirect
	go.opentelemetry.io/proto/otlp v1.3.1 // indirect
	go.starlark.net v0.0.0-20230525235612-a134d8f9ddca // indirect
	go.uber.org/automaxprocs latest // stable
	go.uber.org/dig latest // stable
	go.uber.org/fx latest // stable
	go.uber.org/mock latest // stable
	go.uber.org/multierr latest // stable
	go.uber.org/zap latest // stable
	go.yaml.in/yaml/v2 latest // stable
	go.yaml.in/yaml/v3 latest // stable
	golang.org/x/exp v0.0.0-20241217172543-b2144cdd0a67 // indirect
	golang.org/x/oauth2 v0.25.0 // indirect
	golang.org/x/telemetry master // latest
	golang.org/x/term v0.30.0 // indirect
	golang.org/x/text v0.23.0 // indirect
	golang.org/x/time v0.10.0 // indirect
	golang.org/x/tools v0.28.0 // indirect
	golang.zx2c4.com/wireguard v0.0.0-20231211153847-12269c276173 // indirect
	golang.zx2c4.com/wireguard/wgctrl v0.0.0-20241231184526-a9ab2273dd10 // indirect
	gonum.org/v1/gonum v0.15.0 // indirect
	google.golang.org/api v0.215.0 // indirect
	google.golang.org/genproto v0.0.0-20241118233622-e639e219e697 // indirect
	google.golang.org/genproto/googleapis/api latest // stable
	google.golang.org/genproto/googleapis/rpc latest // stable
	google.golang.org/protobuf v1.36.4 // indirect
	gopkg.in/gcfg.v1 latest // stable
	gopkg.in/inf.v0 latest // stable
	gopkg.in/natefinch/lumberjack.v2 latest // stable
	gopkg.in/square/go-jose.v2 latest // stable
	gopkg.in/warnings.v0 latest // stable
	gopkg.in/yaml.v3 latest // stable
	k8s.io/apiextensions-apiserver v0.29.15 // replace
	k8s.io/code-generator v0.29.15 // replace
	k8s.io/controller-manager v0.29.15 // replace
	k8s.io/csi-translation-lib v0.29.15 // replace
	k8s.io/dynamic-resource-allocation v0.29.15 // replace
	k8s.io/endpointslice v0.29.15 // replace
	k8s.io/gengo latest // stable
	k8s.io/kms v0.29.15 // replace
	k8s.io/kube-aggregator v0.29.15 // replace
	k8s.io/kube-controller-manager v0.29.15 // replace
	k8s.io/kube-openapi v0.0.0-20250710124328-f3f2b991d03b // replace
	k8s.io/kube-scheduler v0.29.15 // replace
	k8s.io/kubelet v0.29.15 // replace
	k8s.io/legacy-cloud-providers v0.29.15 // replace
	k8s.io/metrics v0.29.15 // replace
	k8s.io/mount-utils v0.29.15 // replace
	k8s.io/pod-security-admission v0.29.15 // replace
	lukechampine.com/blake3 v1.3.0 // indirect
	sigs.k8s.io/apiserver-network-proxy/konnectivity-client v0.31.4 // stable
	sigs.k8s.io/json latest // stable
	sigs.k8s.io/knftables latest // stable
	sigs.k8s.io/kustomize/api v0.13.5-0.20230601165947-6ce0bf390ce3 // indirect
	sigs.k8s.io/kustomize/kustomize/v5 v5.0.4-0.20230601165947-6ce0bf390ce3 // indirect
	sigs.k8s.io/kustomize/kyaml v0.14.3-0.20230601165947-6ce0bf390ce3 // indirect
	sigs.k8s.io/randfill latest // stable
	sigs.k8s.io/structured-merge-diff/v4 latest // replace
	tags.cncf.io/container-device-interface v0.8.1 // indirect
	tags.cncf.io/container-device-interface/specs-go v0.8.0 // indirect
)
