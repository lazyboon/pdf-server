FROM ubuntu:22.10

ENV LANG C.UTF-8

WORKDIR /home

COPY ./server /home/server

RUN apt-get update  \
    && apt-get install -y wget \
    \
    && wget http://archive.ubuntu.com/ubuntu/pool/main/o/openssl/libssl1.1_1.1.1-1ubuntu2.1\~18.04.20_amd64.deb \
    && dpkg -i libssl1.1_1.1.1-1ubuntu2.1~18.04.20_amd64.deb \
    && rm -f libssl1.1_1.1.1-1ubuntu2.1~18.04.20_amd64.deb \
    \
    && wget https://github.com/wkhtmltopdf/packaging/releases/download/0.12.6-1/wkhtmltox_0.12.6-1.focal_amd64.deb \
    && apt --fix-broken install -y ./wkhtmltox_0.12.6-1.focal_amd64.deb \
    && rm -f wkhtmltox_0.12.6-1.focal_amd64.deb \
    && apt-get install -y ttf-wqy-microhei

CMD ["/home/server", "&"]
