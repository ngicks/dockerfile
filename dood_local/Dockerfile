# syntax = docker/dockerfile:1.4
ARG UBUNTU_TAG=jammy-20240911.1
ARG GO_VERSION=1.24.0
ARG NODE_VERSION=22.9.0
ARG DENO_VERSION=1.46.3
FROM golang:${GO_VERSION}-bookworm AS go_bin
FROM node:${NODE_VERSION}-bookworm AS node_bin
FROM denoland/deno:bin-${DENO_VERSION} AS deno_bin
FROM docker:24.0.5 AS docker_bin
FROM docker/compose-bin:v2.20.2 AS docker_compose_bin
FROM docker/buildx-bin:0.11.2 AS docker_buildx_bin

FROM ubuntu:${UBUNTU_TAG}

RUN --mount=type=cache,target=/var/cache/apt,sharing=locked \
    --mount=type=cache,target=/var/lib/apt,sharing=locked \
    rm -f /etc/apt/apt.conf.d/docker-clean \
    ; echo 'Binary::apt::APT::Keep-Downloaded-Packages "true";' > /etc/apt/apt.conf.d/keep-cache &&\
    apt-get update && apt-get install --no-install-recommends -y ca-certificates

ARG LIBS_DEP="autoconf automake bzip2 dpkg-dev file g++ gcc imagemagick libbz2-dev libc6-dev\
    libcurl4-openssl-dev libdb-dev libevent-dev libffi-dev libgdbm-dev libglib2.0-dev libgmp-dev\
    libjpeg-dev libkrb5-dev liblzma-dev libmagickcore-dev libmagickwand-dev libmaxminddb-dev\
    libncurses5-dev libncursesw5-dev libpng-dev libpq-dev libreadline-dev libsqlite3-dev libssl-dev\
    libtool libwebp-dev libxml2-dev libxslt-dev libyaml-dev make patch unzip xz-utils zlib1g-dev\
    libmysqlclient-dev"
RUN --mount=type=secret,id=certs,target=/etc/ssl/certs/ca-certificates.crt \
    --mount=type=cache,target=/var/cache/apt,sharing=locked \
    --mount=type=cache,target=/var/lib/apt,sharing=locked \
    rm -f /etc/apt/apt.conf.d/docker-clean \
    ; echo 'Binary::apt::APT::Keep-Downloaded-Packages "true";' > /etc/apt/apt.conf.d/keep-cache &&\
    apt-get update && apt-get install --no-install-recommends -y ${LIBS_DEP}

ARG NETBASE="curl netbase wget"
ARG DEV_DEP="git less gpg"
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

# docker cli
COPY --from=docker_bin /usr/local/bin/docker /usr/local/bin/docker
COPY --from=docker_compose_bin /docker-compose /usr/libexec/docker/cli-plugins/docker-compose
COPY --from=docker_buildx_bin /buildx /usr/libexec/docker/cli-plugins/docker-buildx

# Go
COPY --from=go_bin /usr/local/go /usr/local/go
ENV GOPATH=/go\
    PATH=/go/bin:/usr/local/go/bin:$PATH
RUN mkdir -p "$GOPATH/src" "$GOPATH/bin" && chmod -R 777 "$GOPATH"    
VOLUME [ "/go" ]

# Node.js
COPY --from=node_bin /usr/local/include /usr/local/include
COPY --from=node_bin /usr/local/lib/node_modules /usr/local/lib/node_modules
COPY --from=node_bin /usr/local/share/man /usr/local/share/man
COPY --from=node_bin /usr/local/share/doc /usr/local/share/doc

COPY --from=node_bin /usr/local/bin/node /usr/local/bin/node
RUN ln -s /usr/local/lib/node_modules/corepack/dist/corepack.js /usr/local/bin/corepack &&\
    ln -s /usr/local/bin/node /usr/local/bin/nodejs &&\
    ln -s /usr/local/lib/node_modules/npm/bin/npm-cli.js /usr/local/bin/npm &&\
    ln -s /usr/local/lib/node_modules/npm/bin/npx-cli.js /usr/local/bin/npx 
VOLUME [ "/root/.npm" ]

# Deno
ENV DENO_CERT=/etc/ssl/certs/ca-certificates.crt
COPY --from=deno_bin /deno /usr/bin/deno
# see https://deno.land/manual/getting_started/setup_your_environment#environment-variables
ENV DENO_DIR=/deno/cache\
    DENO_INSTALL_ROOT=/deno\
    PATH=/deno/bin:$PATH
VOLUME [ "/deno" ]

ENTRYPOINT [ "/bin/bash" ]
