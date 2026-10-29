K3s - Lightweight Kubernetes
===============================================
[![FOSSA Status](https://app.fossa.com/api/projects/custom%2B25850%2Fgithub.com%2Fk3s-io%2Fk3s.svg?type=shield)](https://app.fossa.com/projects/custom%2B25850%2Fgithub.com%2Fk3s-io%2Fk3s?ref=badge_shield)
[![Nightly CI](https://github.com/k3s-io/k3s/actions/workflows/nightly-install.yaml/badge.svg)](https://github.com/k3s-io/k3s/actions/workflows/nightly-install.yaml)
[![Build Status](https://drone-publish.k3s.io/api/badges/k3s-io/k3s/status.svg)](https://drone-publish.k3s.io/k3s-io/k3s)
[![Integration Test Coverage](https://github.com/k3s-io/k3s/actions/workflows/integration.yaml/badge.svg)](https://github.com/k3s-io/k3s/actions/workflows/integration.yaml)
[![Unit Test Coverage](https://github.com/k3s-io/k3s/actions/workflows/unitcoverage.yaml/badge.svg)](https://github.com/k3s-io/k3s/actions/workflows/unitcoverage.yaml)
[![OpenSSF Best Practices](https://www.bestpractices.dev/projects/6835/badge)](https://www.bestpractices.dev/projects/6835)
[![OpenSSF Scorecard](https://api.scorecard.dev/projects/github.com/k3s-io/k3s/badge)](https://scorecard.dev/viewer/?uri=github.com/k3s-io/k3s)
[![Releases](https://img.shields.io/github/downloads/k3s-io/k3s/total.svg)](https://github.com/k3s-io/k3s/tags?label=Downloads)
[![CLOMonitor](https://img.shields.io/endpoint?url=https://clomonitor.io/api/projects/cncf/k3s/badge)](https://clomonitor.io/projects/cncf/k3s)

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
