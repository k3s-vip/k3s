K3s - Lightweight Kubernetes
===============================================

Lightweight Kubernetes.  Production ready, easy to install, half the memory, all in a binary less than 100 MB.

![](https://github.com/Platane/snk/raw/refs/heads/output-svg-only/github-contribution-grid-snake.svg)

K3s bundles the following technologies together into a single cohesive distribution:

* [Containerd](https://github.com/containerd/containerd) & [runc](https://github.com/opencontainers/runc)
* [Flannel](https://github.com/flannel-io/flannel) for CNI
* [CoreDNS](https://github.com/coredns/coredns)
* [Metrics Server](https://github.com/kubernetes-sigs/metrics-server)
* [Traefik](https://github.com/traefik/traefik) for ingress
* [Klipper-lb](https://github.com/k3s-io/klipper-lb) as an embedded service load balancer provider
* [Kube-router](https://github.com/cloudnativelabs/kube-router) netpol controller for network policy
* [Helm-controller](https://github.com/k3s-io/helm-controller) to allow for CRD-driven deployment of helm manifests
* [Kine](https://github.com/k3s-io/kine) as a datastore shim that allows etcd to be replaced with other databases
* [Local-path-provisioner](https://github.com/rancher/local-path-provisioner) for provisioning volumes using local storage
* [Host utilities](https://github.com/k3s-io/k3s-root) such as iptables/nftables, ebtables, ethtool, & busybox

![](https://github.com/Platane/snk/raw/refs/heads/output-svg-only/github-contribution-grid-snake-dark.svg)
These technologies can be disabled or swapped out for technologies of your choice.

Additionally, K3s simplifies Kubernetes operations by maintaining functionality for:

* Managing the TLS certificates of Kubernetes components
* Managing the connection between worker and server nodes
* Auto-deploying Kubernetes resources from local manifests in realtime as they are changed.
* Managing an embedded etcd cluster
