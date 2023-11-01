FROM --platform=linux/amd64 golang:1.21.3-alpine3.18

RUN adduser -D go

USER go

WORKDIR /home/go/app

COPY . .

