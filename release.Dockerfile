FROM golang:1.24.4-alpine

WORKDIR /app

RUN go install github.com/pressly/goose/v3/cmd/goose@latest

COPY src ./src
COPY migrations ./migrations
