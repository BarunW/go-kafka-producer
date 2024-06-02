# Builder stage
FROM golang:1.22-bookworm AS builder

# Set the working directory
WORKDIR /app

# Copy module files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Install build dependencies 
RUN apt-get update && apt-get install -y --no-install-recommends \
    build-essential \
    librdkafka-dev 

# Install the Confluent Kafka Go client library (confluent-kafka-go)
RUN go get github.com/confluentinc/confluent-kafka-go/kafka

# Copy the application code
COPY . .

# Build the Go binary with CGO enabled
RUN CGO_ENABLED=1 go build -o producer

# Final stage (for a smaller production image)
FROM debian:bookworm-slim

# Set the working directory in the final image
WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder /app/producer .

# Copy librdkafka from the builder stage
COPY --from=builder /usr/lib/x86_64-linux-gnu/librdkafka.so.1 /usr/lib/x86_64-linux-gnu/

# Expose the port your application uses (if applicable)
EXPOSE 8080

# Command to run the application
CMD ["./producer"]

