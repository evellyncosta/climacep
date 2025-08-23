FROM golang:1.21-alpine AS builder

WORKDIR /app

# Copy go.mod file
COPY go.mod ./

# Download dependencies and generate go.sum
RUN go mod download && go mod tidy

# Copy the source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o climacep ./cmd/api

# Use a small alpine image for the final image
FROM alpine:latest AS final

WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder /app/climacep .

# Expose port
EXPOSE 8080

# Run the application
CMD ["./climacep"]