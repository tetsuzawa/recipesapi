FROM golang:1.13-alpine

RUN mkdir -p /go/src/api
ADD . /go/src/api

WORKDIR /go/src/api
RUN  go install ./cmd/server

CMD ["server"]
