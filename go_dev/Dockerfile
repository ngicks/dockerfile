# syntax = docker/dockerfile:1.4
ARG UBUNTU_TAG=jammy-20230624
ARG GO_VERSION=1.21.0
ARG DENO_VERSION=1.36.0
FROM golang:${GO_VERSION}-bookworm AS go_bin
FROM denoland/deno:bin-${DENO_VERSION} AS deno_bin

FROM ubuntu:${UBUNTU_TAG}

ARG NETBASE="curl netbase wget"
ARG DEV_DEP="git less"
ARG GO_DEP="g++ gcc libc6-dev make pkg-config"
RUN --mount=type=secret,id=certs,target=/etc/ssl/certs/ca-certificates.crt \
    --mount=type=cache,target=/var/cache/apt,sharing=locked \
    --mount=type=cache,target=/var/lib/apt,sharing=locked \
    rm -f /etc/apt/apt.conf.d/docker-clean \
    ; echo 'Binary::apt::APT::Keep-Downloaded-Packages "true";' > /etc/apt/apt.conf.d/keep-cache &&\
    apt-get update && apt-get install --no-install-recommends -y\
    ${NETBASE}\
    ${DEV_DEP}\
    ${GO_DEP}

# Go
COPY --from=go_bin /usr/local/go /usr/local/go
ENV GOPATH=/go\
    PATH=/go/bin:/usr/local/go/bin:$PATH
RUN mkdir -p "$GOPATH/src" "$GOPATH/bin" && chmod -R 777 "$GOPATH"    
VOLUME [ "/go" ]

# Deno
ENV DENO_CERT=/etc/ssl/certs/ca-certificates.crt
COPY --from=deno_bin /deno /usr/bin/deno
# see https://deno.land/manual/getting_started/setup_your_environment#environment-variables
ENV DENO_DIR=/deno/cache\
    DENO_INSTALL_ROOT=/deno\
    PATH=/deno/bin:$PATH
VOLUME [ "/deno" ]

ENTRYPOINT [ "/bin/bash" ]
