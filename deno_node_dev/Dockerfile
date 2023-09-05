# syntax = docker/dockerfile:1.4
ARG UBUNTU_TAG=jammy-20230624
ARG NODE_VERSION=20.5.0
ARG DENO_VERSION=1.36.0
FROM node:${NODE_VERSION}-bookworm AS node_bin
FROM denoland/deno:bin-${DENO_VERSION} AS deno_bin

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
ARG DEV_DEP="git less"
RUN --mount=type=secret,id=certs,target=/etc/ssl/certs/ca-certificates.crt \
    --mount=type=cache,target=/var/cache/apt,sharing=locked \
    --mount=type=cache,target=/var/lib/apt,sharing=locked \
    rm -f /etc/apt/apt.conf.d/docker-clean \
    ; echo 'Binary::apt::APT::Keep-Downloaded-Packages "true";' > /etc/apt/apt.conf.d/keep-cache &&\
    apt-get update && apt-get install --no-install-recommends -y\
    ${NETBASE}\
    ${DEV_DEP}

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
