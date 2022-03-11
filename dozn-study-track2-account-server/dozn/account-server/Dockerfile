FROM golang:alpine AS builder
MAINTAINER hjkim

WORKDIR /src/
COPY . .

RUN go env -w GO111MODULE=auto
RUN go build main.go

FROM alpine

COPY --from=builder /src/main /src/main

CMD ["./src/main"]