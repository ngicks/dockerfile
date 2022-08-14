ARG UBUNTU_TAG=jammy-20220801
# go as code generator
ARG GO_VERSION=1.19.0
# deno as script runtime.
ARG DENO_VERSION=1.24.3

FROM krallin/ubuntu-tini:trusty AS tini_bin
FROM golang:${GO_VERSION}-bullseye AS go_bin
FROM denoland/deno:bin-${DENO_VERSION} AS deno_bin

FROM ubuntu:${UBUNTU_TAG}

ARG NETBASE="ca-certificates curl netbase wget"
ARG GO_DEP="g++ gcc libc6-dev make pkg-config"
ARG DEV_DEP="git less"
RUN apt-get update && apt-get install -y\
    ${NETBASE}\
    ${GO_DEP}\
    ${DEV_DEP}

RUN apt-get update && apt-get install -y\
    ca-certificates curl gnupg lsb-release &&\
    curl -fsSL https://download.docker.com/linux/ubuntu/gpg | gpg --dearmor -o /usr/share/keyrings/docker-archive-keyring.gpg &&\
    echo "deb [arch=$(dpkg --print-architecture) signed-by=/usr/share/keyrings/docker-archive-keyring.gpg] https://download.docker.com/linux/ubuntu $(lsb_release -cs) stable"\
    | tee /etc/apt/sources.list.d/docker.list > /dev/null &&\
    apt-get update && apt-get install -y\
    docker-ce docker-ce-cli containerd.io

COPY --from=tini_bin /usr/bin/tini /usr/bin/tini

COPY --from=go_bin /usr/local/go /usr/local/go

ENV GOPATH=/go\
    PATH=${GOPATH}/bin:$PATH

RUN mkdir -p "$GOPATH/src" "$GOPATH/bin" && chmod -R 777 "$GOPATH"    

VOLUME [ "/go" ]

COPY --from=deno_bin /deno /usr/bin/deno
# see https://deno.land/manual/getting_started/setup_your_environment#environment-variables
ENV DENO_DIR=/deno/cache\
    DENO_INSTALL_ROOT=/deno/bin\
    PATH=${DENO_INSTALL_ROOT}:$PATH

VOLUME [ "/deno" ]

ENTRYPOINT [ "/usr/bin/tini", "--", "/bin/bash" ]