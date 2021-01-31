FROM golang:alpine

RUN apk update && apk add --no-cache \
  alpine-sdk \
  git \
  && go get github.com/pilu/fresh

ENV GOPATH /go
ENV SRCPATH $GOPATH/src/github.com/naoki/gacha

WORKDIR $SRCPATH

ENV GO111MODULE=on

ADD . $SRCPATH
RUN go get -v -d


# ENV CGO_ENABLED=0
# ENV GOOS=linux
# RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /tmp/build/gacha

# FROM scratchZ
# ENV GIN_MODE=release
# ENV PORT 8080
EXPOSE 8080

# COPY --from=builder /tmp/build/gacha main

# ENTRYPOINT [ "./main" ]


