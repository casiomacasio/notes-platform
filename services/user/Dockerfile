# syntax=docker/dockerfile:1.4
FROM golang:1.21 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o main ./cmd

# -- Final image
FROM gcr.io/distroless/base-debian11

WORKDIR /app

COPY --from=builder /app/main .
COPY .env .env

CMD ["./main"]
