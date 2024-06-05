# Dockerfile for service1
FROM golang:latest

WORKDIR /app

COPY . .
RUN go mod download

COPY . .

RUN go build -o service1

ENV SERVICE_TYPE=service1

EXPOSE 8080

CMD ["./service1"]
