# Dockerfile for service2
FROM golang:latest

WORKDIR /app

COPY . .
RUN go mod download

COPY . .

RUN go build -o service2

ENV SERVICE_TYPE=service2

EXPOSE 8080

CMD ["./service2"]
