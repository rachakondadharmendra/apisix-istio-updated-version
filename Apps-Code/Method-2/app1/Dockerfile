
FROM golang:1.22-alpine

WORKDIR /app

COPY . .
RUN go mod download

COPY . .

RUN go build -o service1

ENV SERVICE2_ADDR=http://service2:8080

EXPOSE 8080

CMD ["./service1"]
