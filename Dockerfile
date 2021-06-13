FROM golang:1.16-alpine AS builder
LABEL MAINTAINER="minhtran.in@outlook.com"

RUN mkdir -p /go/src/work/gingo
WORKDIR /go/src/work/gingo

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
RUN go build ./main.go

EXPOSE 8888

ENTRYPOINT ["./main"]
#  docker run -it -p 8888:8888 --name gingo gingo
# GOPATH + module in go.mod