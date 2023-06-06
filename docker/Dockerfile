# Initial stage: download modules
FROM golang:1.17-alpine as builder

ENV config=docker

WORKDIR /usr/local/go/src/app

COPY ./ /usr/local/go/src/app

RUN go mod download


# Intermediate stage: Build the binary
FROM golang:1.17-alpine as runner

COPY --from=builder ./usr/local/go/src/ /app

RUN go get github.com/githubnemo/CompileDaemon

WORKDIR /app
ENV config=docker

EXPOSE 5005
EXPOSE 5555
EXPOSE 7070

ENTRYPOINT CompileDaemon --build="go build cmd/api/main.go" --command=./main
