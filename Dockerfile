ARG GOLANG=golang:1.23-alpine
FROM ${GOLANG} AS build


# Install necessary packages
RUN apk -U --no-cache add \
    bash git gcc musl-dev docker vim less file curl wget ca-certificates jq linux-headers \
    zlib-dev tar zip squashfs-tools npm coreutils python3 py3-pip openssl-dev libffi-dev libseccomp \
    libseccomp-dev libseccomp-static make libuv-static sqlite-dev sqlite-static libselinux \
    libselinux-dev zlib-dev zlib-static zstd pigz alpine-sdk binutils-gold btrfs-progs-dev \
    btrfs-progs-static gawk yq pipx

# Install AWS CLI
RUN PIPX_BIN_DIR=/usr/local/bin pipx install awscli

# Cleanup
RUN rm -rf /go/src /go/pkg


# Set SELINUX environment variable
ARG SELINUX=true
ENV STATIC_BUILD=true GODEBUG=gotypesalias=0 DOCKER_BUILDKIT=1 SKIP_IMAGE=1 SKIP_AIRGAP=1

WORKDIR /go/src/github.com/k3s-io/k3s
COPY . .
ARG OEM
RUN bash scripts/ci;cp -av dist/artifacts/k3s* /tmp/k3s
FROM busybox:1.36.1
COPY --from=build /tmp/k3s /
