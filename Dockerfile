# Builder
FROM golang:1.21-alpine3.18 AS builder

WORKDIR /app

COPY . .

RUN go build -o /app/affiliate-forwarder main.go

# Runner
FROM alpine:3.18

WORKDIR /app

# Copy binary from builder
COPY --from=builder /app/affiliate-forwarder /app/affiliate-forwarder

# Run the binary
CMD ["/bin/sh", "-c", "/app/affiliate-forwarder"]