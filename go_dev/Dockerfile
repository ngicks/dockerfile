ARG UBUNTU_TAG=jammy-20220801
ARG GO_VERSION=1.19.0
ARG DENO_VERSION=1.24.3

FROM krallin/ubuntu-tini:trusty AS tini_bin
FROM golang:${GO_VERSION}-bullseye AS go_bin
FROM denoland/deno:bin-${DENO_VERSION} AS deno_bin

FROM ubuntu:${UBUNTU_TAG}

ARG NETBASE="ca-certificates curl netbase wget"
ARG DEV_DEP="git less"
ARG GO_DEP="g++ gcc libc6-dev make pkg-config"
RUN apt-get update && apt-get install -y\
    ${NETBASE}\
    ${DEV_DEP}\
    ${GO_DEP}

COPY --from=tini_bin /usr/bin/tini /usr/bin/tini

# Go
COPY --from=go_bin /usr/local/go /usr/local/go
ENV GOPATH=/go\
    PATH=${GOPATH}/bin:/usr/local/go/bin:$PATH
RUN mkdir -p "$GOPATH/src" "$GOPATH/bin" && chmod -R 777 "$GOPATH"    
VOLUME [ "/go" ]

# Deno
COPY --from=deno_bin /deno /usr/bin/deno
# see https://deno.land/manual/getting_started/setup_your_environment#environment-variables
ENV DENO_DIR=/deno/cache\
    DENO_INSTALL_ROOT=/deno/bin\
    PATH=${DENO_INSTALL_ROOT}:$PATH
VOLUME [ "/deno" ]

ENTRYPOINT [ "/usr/bin/tini", "--", "/bin/bash" ]
