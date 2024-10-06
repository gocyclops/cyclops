# Stage 1: Build the app
FROM golang:1.22-alpine3.20 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

COPY .env .env

RUN go build -o main .

# Stage 2: Run the Go App
FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/main .
COPY .env /root/

COPY --from=builder /app/.env .

EXPOSE 8080

CMD ["./main"]
