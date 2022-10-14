ARG UBUNTU_TAG=jammy-20220801

FROM krallin/ubuntu-tini:trusty AS tini_bin

FROM ubuntu:${UBUNTU_TAG}

WORKDIR /tmp

RUN apt-get update &&\
    apt-get install -y wget &&\
    wget https://dl.google.com/linux/direct/google-chrome-stable_current_amd64.deb &&\
    apt-get install -y ./google-chrome-stable_current_amd64.deb &&\
    rm -rf ./*

RUN apt-get update &&\
    apt-get install -y fonts-noto

WORKDIR /root

COPY --from=tini_bin /usr/bin/tini /usr/bin/tini

ENV DISPLAY=host.docker.internal:0.0

ENTRYPOINT [ "/usr/bin/tini", "--" ]
CMD ["google-chrome", "--no-sandbox"]
