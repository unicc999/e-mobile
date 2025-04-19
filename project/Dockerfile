FROM golang:1.23.2 AS builder

WORKDIR /app

COPY ./go.mod ./go.sum /

RUN go mod download

COPY . ./

RUN go build -o peopleAPI ./cmd/main.go

FROM alpine:latest

RUN apk --no-cache 