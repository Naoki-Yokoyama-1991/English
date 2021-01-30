FROM golang:alpine

RUN apk update && apk --no-cache add git

RUN mkdir -p $GOPATH/src/github.com/naoki/gacha

WORKDIR $GOPATH/src/github.com/naoki/gacha

ADD . $GOPATH/src/github.com/naoki/gacha

RUN go get