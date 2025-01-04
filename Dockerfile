# Choose whatever you want, version >= 1.16
FROM golang:1.23-alpine

WORKDIR /app

RUN go install github.com/air-verse/air@latest

COPY db-service/go.mod db-service/go.sum ./
RUN go mod download

CMD ["air"]