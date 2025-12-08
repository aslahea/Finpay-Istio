# Build stage
FROM golang:1.22-alpine AS builder
WORKDIR /app

COPY go.mod ./
RUN go mod tidy

COPY . .
RUN go build -o service .

# Run stage
FROM alpine:3.19
WORKDIR /app
COPY --from=builder /app/service /app/service

EXPOSE 8080
CMD ["/app/service"]
