FROM daocloud.io/library/golang:1.7.0-wheezy

ARG component

COPY dist/$component bin/$component

CMD bin/$component
