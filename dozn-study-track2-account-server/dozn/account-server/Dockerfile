FROM golang:alpine AS builder
MAINTAINER hjkim

WORKDIR /src/
COPY . /src/

RUN go env -w GO111MODULE=auto
RUN go build main.go

FROM alpine

COPY --from=builder /go/main /go/main

CMD ["./main"]