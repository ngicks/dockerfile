# syntax = docker/dockerfile:1.4
ARG UBUNTU_TAG=jammy-20230624
ARG RUST_VERSION=1.71.1
ARG DENO_VERSION=1.36.0
FROM rust:${RUST_VERSION}-slim-buster AS rust_bin
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
ARG DEV_DEP="git less gpg"
ARG BINDGEN_DEP="llvm-dev libclang-dev clang"
RUN --mount=type=secret,id=certs,target=/etc/ssl/certs/ca-certificates.crt \
    --mount=type=cache,target=/var/cache/apt,sharing=locked \
    --mount=type=cache,target=/var/lib/apt,sharing=locked \
    rm -f /etc/apt/apt.conf.d/docker-clean \
    ; echo 'Binary::apt::APT::Keep-Downloaded-Packages "true";' > /etc/apt/apt.conf.d/keep-cache &&\
    apt-get update && apt-get install --no-install-recommends -y\
    ${NETBASE}\
    ${DEV_DEP}\
    ${BINDGEN_DEP}

# rust
## We need to avoid COPYing some data to dir since it prevents cargo update.
## below VOLUME instruction overwrites contents of ./cargo and ./rustup on everytime the container launch.
COPY --from=rust_bin /usr/local/cargo /usr/local/cargo_initial
COPY --from=rust_bin /usr/local/rustup/toolchains /usr/local/rustup/toolchains
COPY --from=rust_bin /usr/local/rustup/update-hashes /usr/local/rustup/update-hashes

ENV RUSTUP_HOME=/usr/local/rustup\
    CARGO_HOME=/usr/local/cargo\
    PATH=/usr/local/cargo/bin:/usr/local/cargo_initial/bin:$PATH\
    RUST_VERSION=${RUST_VERSION}

VOLUME [ "/usr/local/cargo", "/usr/local/rustup" ]

# Deno
ENV DENO_CERT=/etc/ssl/certs/ca-certificates.crt
COPY --from=deno_bin /deno /usr/bin/deno
# see https://deno.land/manual/getting_started/setup_your_environment#environment-variables
ENV DENO_DIR=/deno/cache\
    DENO_INSTALL_ROOT=/deno\
    PATH=/deno/bin:$PATH
VOLUME [ "/deno" ]

ENTRYPOINT [ "/bin/bash" ]
