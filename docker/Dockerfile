FROM golang:1.13-alpine as base

ARG GO111MODULE

ENV GO111MODULE on

ENV HOME /home/app

RUN apk update && \
  apk upgrade && \
  apk add --no-cache bash git

WORKDIR $HOME

COPY . $HOME

FROM base as dev

RUN go get -u github.com/githubnemo/CompileDaemon

ENTRYPOINT CompileDaemon -directory="$HOME/src" -build="go build -o $HOME/bin/main $HOME/src/main.go" -command="$HOME/bin/main"
