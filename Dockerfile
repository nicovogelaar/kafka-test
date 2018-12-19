FROM golang:alpine3.8

ARG LIBRESSL_VERSION=2.7
ARG LIBRDKAFKA_VERSION=0.11.6-r1
ARG ALPINE_VERSION=edge
ARG LIBCRYPTO_VERSION=1.1.1a-r1
ARG LIBSSL_VERSION=1.1.1a-r1

RUN apk update && \
    apk add libcrypto1.1=${LIBCRYPTO_VERSION} --update-cache --repository http://nl.alpinelinux.org/alpine/${ALPINE_VERSION}/main && \
    apk add libssl1.1=${LIBSSL_VERSION} --update-cache --repository http://nl.alpinelinux.org/alpine/${ALPINE_VERSION}/main && \
    apk add libressl${LIBRESSL_VERSION}-libcrypto libressl${LIBRESSL_VERSION}-libssl --update-cache --repository http://nl.alpinelinux.org/alpine/${ALPINE_VERSION}/main && \
    apk add librdkafka=${LIBRDKAFKA_VERSION} --update-cache --repository http://nl.alpinelinux.org/alpine/${ALPINE_VERSION}/community && \
    apk add librdkafka-dev=${LIBRDKAFKA_VERSION} --update-cache --repository http://nl.alpinelinux.org/alpine/${ALPINE_VERSION}/community && \
    apk add git openssh openssl yajl-dev zlib-dev cyrus-sasl-dev openssl-dev build-base coreutils && \
    rm -f /var/cache/apk/*
