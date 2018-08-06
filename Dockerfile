FROM golang:alpine as build

ENV GOPATH /go

COPY main.go /go/src/github.com/kamsz/certmon/main.go

WORKDIR /go/src/github.com/kamsz/certmon/

RUN apk add --no-cache git && \
    go get && \
    go build && \
    chmod +x /go/src/github.com/kamsz/certmon/certmon

FROM alpine:3.8

RUN apk add --no-cache ca-certificates

COPY --from=build /go/src/github.com/kamsz/certmon/certmon /certmon

CMD ["/certmon"]
