ARG UBUNTU_TAG=jammy-20220801
ARG RUST_VERSION=1.63.0
# deno as script runtime.
ARG DENO_VERSION=1.24.3

FROM krallin/ubuntu-tini:trusty AS tini_bin
FROM rust:${RUST_VERSION}-slim-buster AS rust_bin
FROM denoland/deno:bin-${DENO_VERSION} AS deno_bin

FROM ubuntu:${UBUNTU_TAG}

ARG NETBASE="ca-certificates curl netbase wget"
ARG RUST_DEP="gcc libc6-dev"
ARG BINDGEN_DEP="llvm-dev libclang-dev clang"
ARG DEV_DEP="git less"
RUN apt-get update && apt-get install -y\
    ${NETBASE}\
    ${RUST_DEP}\
    ${BINDGEN_DEP}\
    ${DEV_DEP}

COPY --from=tini_bin /usr/bin/tini /usr/bin/tini

COPY --from=rust_bin /usr/local/cargo /usr/local/cargo
COPY --from=rust_bin /usr/local/rustup /usr/local/rustup

ENV RUSTUP_HOME=/usr/local/rustup\
    CARGO_HOME=/usr/local/cargo\
    PATH=/usr/local/cargo/bin:$PATH\
    RUST_VERSION=${RUST_VERSION}

VOLUME [ "/usr/local/cargo", "/usr/local/rustup" ]

COPY --from=deno_bin /deno /usr/bin/deno
# see https://deno.land/manual/getting_started/setup_your_environment#environment-variables
ENV DENO_DIR=/deno/cache\
    DENO_INSTALL_ROOT=/deno/bin\
    PATH=${DENO_INSTALL_ROOT}:$PATH

VOLUME [ "/deno" ]

ENTRYPOINT [ "/usr/bin/tini", "--", "/bin/bash" ]