FROM golang:1.26-alpine AS builder

WORKDIR /app

# Copy dependency files
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build the application
# CGO_ENABLED=0 creates a statically linked binary
# -ldflags="-w -s" reduces binary size by stripping debug info
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o tasm-api .

# Final stage
FROM alpine:latest

RUN apk add --no-cache tzdata

WORKDIR /app

# Copy the binary from the builder
COPY --from=builder /app/tasm-api .

# Expose port 8080
EXPOSE 8080

# Command to run the application
CMD ["./tasm-api"]
