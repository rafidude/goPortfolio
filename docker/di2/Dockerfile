# syntax=docker/dockerfile:1

# Build

FROM golang:1.18 AS build

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY *.go .

RUN go build -o /hello

# Deploy

FROM debian:latest AS deploy

WORKDIR /

COPY --from=build /hello /usr/local/bin/hello

EXPOSE 8080

ENTRYPOINT ["/usr/local/bin/hello"]