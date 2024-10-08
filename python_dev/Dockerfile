# syntax = docker/dockerfile:1.4
ARG UBUNTU_TAG=jammy-20240911.1
ARG DENO_VERSION=1.46.3
ARG PYTHON_VERSION=3.12.7
FROM denoland/deno:bin-${DENO_VERSION} AS deno_bin
FROM python:${PYTHON_VERSION}-slim-bookworm as python_bin

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
ENV TZ=Asia/Tokyo
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone
ARG PYTHON_DEPS="gnupg libbluetooth-dev libexpat1-dev tk-dev uuid-dev"
RUN --mount=type=secret,id=certs,target=/etc/ssl/certs/ca-certificates.crt \
    --mount=type=cache,target=/var/cache/apt,sharing=locked \
    --mount=type=cache,target=/var/lib/apt,sharing=locked \
    rm -f /etc/apt/apt.conf.d/docker-clean \
    ; echo 'Binary::apt::APT::Keep-Downloaded-Packages "true";' > /etc/apt/apt.conf.d/keep-cache &&\
    apt-get update && apt-get install --no-install-recommends -y\
    ${NETBASE}\
    ${DEV_DEP}\
    ${PYTHON_DEPS}

# Deno
ENV DENO_CERT=/etc/ssl/certs/ca-certificates.crt
COPY --from=deno_bin /deno /usr/bin/deno
# see https://deno.land/manual/getting_started/setup_your_environment#environment-variables
ENV DENO_DIR=/deno/cache\
    DENO_INSTALL_ROOT=/deno\
    PATH=/deno/bin:$PATH
VOLUME [ "/deno" ]
# python
COPY --from=python_bin  /usr/local/bin/2to3-3.12 \ 
                        /usr/local/bin/idle3.12 \ 
                        /usr/local/bin/pip \ 
                        /usr/local/bin/pip3 \ 
                        /usr/local/bin/pip3.12 \ 
                        /usr/local/bin/pydoc3.12 \ 
                        /usr/local/bin/python3.12 \ 
                        /usr/local/bin/python3.12-config \ 
                        /usr/local/bin/
COPY --from=python_bin  /usr/local/lib/libpython3.12.so.1.0\
                        /usr/local/lib/libpython3.so\
                        /usr/local/lib/
RUN ln -s 2to3-3.12 /usr/local/bin/2to3 && \
    ln -s idle3  /usr/local/bin/idle && \
    ln -s idle3.12  /usr/local/bin/idle3 && \
    ln -s pydoc3  /usr/local/bin/pydoc && \
    ln -s pydoc3.12  /usr/local/bin/pydoc3 && \
    ln -s python3  /usr/local/bin/python && \
    ln -s python3.12  /usr/local/bin/python3 && \
    ln -s python3-config  /usr/local/bin/python-config && \
    ln -s python3.12-config  /usr/local/bin/python3-config && \
    # lib
    ln -s libpython3.12.so.1.0 /usr/local/lib/libpython3.12.so &&\
    # lib
    ln -s /usr/local/lib/libpython3.12.so.1.0 /lib/x86_64-linux-gnu/libpython3.12.so.1.0 &&\
    ln -s /usr/local/lib/libpython3.so /lib/x86_64-linux-gnu/libpython3.so 
COPY --from=python_bin /usr/local/lib/python3.12 /usr/local/lib/python3.12
ENV PYTHON_VERSION=3.12.7
RUN --mount=type=secret,id=certs,target=/etc/ssl/certs/ca-certificates.crt \
    curl -cafile /etc/ssl/certs/ca-certificates.crt -sSL https://install.python-poetry.org | python3 -
ENV PATH="/root/.local/bin:$PATH"\
    PYTHONDONTWRITEBYTECODE=1
VOLUME [ "/root/.cache/pip" ]

ENTRYPOINT [ "/bin/bash" ]
