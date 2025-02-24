FROM golang:1.23.3

ENV CGO_ENABLED=1

RUN apt update && apt install -y build-essential

RUN mkdir /work

WORKDIR /work
