FROM golang:1.14.4-alpine3.12 AS build

WORKDIR /go
RUN mkdir -p /go/modules
RUN mkdir -p /go/sampleapp
ENV GOPATH=/go/modules

RUN go get -v github.com/stamblerre/gocode &&\
    go get -v github.com/uudashr/gopkgs/v2/cmd/gopkgs

RUN apk update && \
    apk add --no-cache git

CMD ["go", "run", "main.go"]