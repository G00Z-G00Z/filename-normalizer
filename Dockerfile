FROM golang:1.18 as dev

WORKDIR /app

RUN go install -v golang.org/x/tools/cmd/godoc@latest

