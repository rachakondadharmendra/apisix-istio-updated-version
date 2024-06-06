
FROM golang:1.22-alpine

WORKDIR /app

COPY . .
RUN go mod download

COPY . .

RUN go build -o service2

EXPOSE 8080

CMD ["./service2"]
