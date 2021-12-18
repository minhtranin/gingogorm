FROM golang:1.16-alpine AS builder
LABEL MAINTAINER="minhtran.in@outlook.com"

RUN mkdir -p /go/src/work/gingo
WORKDIR /go/src/work/gingo

COPY go.mod .
COPY go.sum .
RUN GO111MODULE=on
# RUN chmod 777 ./go.sum
RUN go mod download all
# RUN go build ./main.go
COPY . .

EXPOSE 8888
RUN go get github.com/pilu/fresh
# ENTRYPOINT ["./main"]
ENTRYPOINT  fresh
#  docker run -it -p 8888:8888 --name gingo gingo
# GOPATH + module in go.mod