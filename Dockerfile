FROM golang:1.19-alpine

WORKDIR /app

COPY ./ ./

CMD go build -o go_listener .; ./go_listener
