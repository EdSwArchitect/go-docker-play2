#
# This Dockerfile builds a recent curl with HTTP/2 client support, using
# a recnent nghttp2 build.
#
# See the Makefile for how to tag it. If Docker and that image is found, the
# Go tests use this curl binary for integration tests.
#
FROM ubuntu:18.04 AS EDPROXY

RUN mkdir -p /opt/playground
RUN mkdir -p /data

RUN ls -l /

WORKDIR /opt/playground

COPY go-docker-play2 ed-proxy

EXPOSE 8080

VOLUME /data

ENTRYPOINT ["/opt/playground/ed-proxy"]
