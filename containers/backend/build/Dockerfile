FROM golang:1.13-alpine

RUN mkdir -p /go/src/api
ADD . /go/src/api

WORKDIR /go/src/api
RUN  go install ./cmd/server
RUN cat prod.env | grep -v ^# | xargs -I {} -n 1 sh -c "echo export {} >> .envrc"


CMD ["sh", "-c", "source .envrc && server"]
