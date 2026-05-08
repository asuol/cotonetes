FROM golang:1.25.9-trixie

ENV CGO_ENABLED=1

RUN apt update && apt install -y build-essential

RUN mkdir /work

WORKDIR /work
