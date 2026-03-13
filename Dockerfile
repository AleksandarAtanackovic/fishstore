# Build stage
FROM golang:1.25-alpine AS builder

WORKDIR /app

# Copy go mod files
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build binary
RUN go build -o server backend/cmd/server/main.go


# Run stage (small image)
FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/server .

EXPOSE 9090

CMD ["./server"]