FROM golang:1.19-alpine

WORKDIR /app

COPY ./ ./

CMD go build -o go_consumer .; ./go_consumer
