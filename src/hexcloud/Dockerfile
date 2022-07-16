FROM golang:1.18-alpine AS build

RUN apk update && apk add --no-cache git make build-base

WORKDIR $GOPATH/src/github.com/3vilM33pl3/hexcloud

COPY . ./
RUN CGO_ENABLED=1 GOOS=linux go build -o /go/bin/hexcloud ./cmd/hexcloud

FROM alpine:latest

RUN apk update && apk add --no-cache curl ca-certificates
RUN rm -rf /var/cache/apk/*

COPY --from=build --chown=65534:0 /go/bin/hexcloud /go/bin/hexcloud

USER 65534

ENTRYPOINT ["/go/bin/hexcloud"]