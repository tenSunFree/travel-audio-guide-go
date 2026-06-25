# Stage 1: Build
FROM golang:1.25-alpine AS builder

WORKDIR /app

# First copy go.mod / go.sum, then use Docker layer cache
# Only re-download when dependencies change
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Compile static binary (disable CGO for easier placement in scratch/alpine image)
RUN CGO_ENABLED=0 GOOS=linux go build \
    -ldflags="-s -w" \
    -o /app/server \
    ./cmd/server

# Stage 2: Runtime
FROM alpine:3.20

WORKDIR /app

# ca-certificates: Required for HTTPS connections
# tzdata: Time zone information
RUN apk --no-cache add ca-certificates tzdata

COPY --from=builder /app/server ./server

EXPOSE 8080

CMD ["./server"]
