# STAGE 1: Build the binary
FROM golang:1.24-alpine AS builder

# Install build tools for SQLite (CGO)
RUN apk add --no-cache gcc musl-dev

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .

# CGO_ENABLED=1 is required for the SQLite driver
RUN CGO_ENABLED=1 GOOS=linux go build -o main .

# STAGE 2: Run the binary
FROM alpine:latest
RUN apk add --no-cache ca-certificates

WORKDIR /root/
# Copy the binary from the builder stage
COPY --from=builder /app/main .

# Expose the port your app runs on
EXPOSE 8080

# Command to run the executable
CMD ["./main"]