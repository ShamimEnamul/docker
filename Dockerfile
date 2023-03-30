FROM golang:alpine

RUN mkdir /build
WORKDIR /build

RUN export GO111MODULE=on
