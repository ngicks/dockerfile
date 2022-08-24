ARG UBUNTU_TAG=jammy-20220801
ARG RUST_VERSION=1.63.0
ARG DENO_VERSION=1.24.3

FROM krallin/ubuntu-tini:trusty AS tini_bin
FROM rust:${RUST_VERSION}-slim-buster AS rust_bin
FROM denoland/deno:bin-${DENO_VERSION} AS deno_bin

FROM ubuntu:${UBUNTU_TAG}

ARG LIBS_DEP="autoconf automake bzip2 dpkg-dev file g++ gcc imagemagick libbz2-dev libc6-dev\
    libcurl4-openssl-dev libdb-dev libevent-dev libffi-dev libgdbm-dev libglib2.0-dev libgmp-dev\
    libjpeg-dev libkrb5-dev liblzma-dev libmagickcore-dev libmagickwand-dev libmaxminddb-dev\
    libncurses5-dev libncursesw5-dev libpng-dev libpq-dev libreadline-dev libsqlite3-dev libssl-dev\
    libtool libwebp-dev libxml2-dev libxslt-dev libyaml-dev make patch unzip xz-utils zlib1g-dev\
    libmysqlclient-dev"
RUN apt-get update && apt-get install -y ${LIBS_DEP}

ARG NETBASE="ca-certificates curl netbase wget"
ARG DEV_DEP="git less"
ARG BINDGEN_DEP="llvm-dev libclang-dev clang"
RUN apt-get update && apt-get install -y\
    ${NETBASE}\
    ${DEV_DEP}\
    ${BINDGEN_DEP}

COPY --from=tini_bin /usr/bin/tini /usr/bin/tini

# rust
COPY --from=rust_bin /usr/local/cargo /usr/local/cargo
COPY --from=rust_bin /usr/local/rustup /usr/local/rustup

ENV RUSTUP_HOME=/usr/local/rustup\
    CARGO_HOME=/usr/local/cargo\
    PATH=/usr/local/cargo/bin:$PATH\
    RUST_VERSION=${RUST_VERSION}

VOLUME [ "/usr/local/cargo", "/usr/local/rustup" ]

# Deno
COPY --from=deno_bin /deno /usr/bin/deno
# see https://deno.land/manual/getting_started/setup_your_environment#environment-variables
ENV DENO_DIR=/deno/cache\
    DENO_INSTALL_ROOT=/deno/bin\
    PATH=${DENO_INSTALL_ROOT}:$PATH
VOLUME [ "/deno" ]

ENTRYPOINT [ "/usr/bin/tini", "--", "/bin/bash" ]
