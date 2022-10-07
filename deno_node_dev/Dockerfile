ARG UBUNTU_TAG=jammy-20220801
ARG NODE_VERSION=18.7.0
ARG DENO_VERSION=1.26.1

FROM krallin/ubuntu-tini:trusty AS tini_bin
FROM node:${NODE_VERSION}-bullseye AS node_bin
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
RUN apt-get update && apt-get install -y\
    ${NETBASE}\
    ${DEV_DEP}

COPY --from=tini_bin /usr/bin/tini /usr/bin/tini

# Node.js
COPY --from=node_bin /usr/local/include /usr/local/include
COPY --from=node_bin /usr/local/lib/node_modules /usr/local/lib/node_modules
COPY --from=node_bin /usr/local/share/man /usr/local/share/man
COPY --from=node_bin /usr/local/share/doc /usr/local/share/doc
COPY --from=node_bin /usr/local/share/systemtap /usr/local/share/systemtap

COPY --from=node_bin /usr/local/bin/node /usr/local/bin/node
RUN ln -s /usr/local/lib/node_modules/corepack/dist/corepack.js /usr/local/bin/corepack &&\
    ln -s /usr/local/bin/node /usr/local/bin/nodejs &&\
    ln -s /usr/local/lib/node_modules/npm/bin/npm-cli.js /usr/local/bin/npm &&\
    ln -s /usr/local/lib/node_modules/npm/bin/npx-cli.js /usr/local/bin/npx 
VOLUME [ "/root/.npm" ]

# Deno
COPY --from=deno_bin /deno /usr/bin/deno
# see https://deno.land/manual/getting_started/setup_your_environment#environment-variables
ENV DENO_DIR=/deno/cache\
    DENO_INSTALL_ROOT=/deno/bin\
    PATH=${DENO_INSTALL_ROOT}:$PATH
VOLUME [ "/deno" ]

ENTRYPOINT [ "/usr/bin/tini", "--", "/bin/bash" ]
