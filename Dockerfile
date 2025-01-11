# Choose whatever you want, version >= 1.16
FROM golang:1.23-alpine

WORKDIR /app

RUN go install github.com/air-verse/air@latest

COPY pc-service/go.mod ./
RUN go mod download

CMD ["air"]