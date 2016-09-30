FROM golang:1.7.1-alpine

ARG component

COPY dist/$component bin/$component

ENV component $component

ENTRYPOINT /go/bin/$component
