FROM --platform=linux/amd64 golang:1.21.3-alpine3.18

RUN adduser -D go

USER go

WORKDIR /home/go/app

COPY go.mod go.sum ./

RUN go mod download && go mod verify

RUN go install github.com/cosmtrek/air@latest

