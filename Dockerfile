## Use the official Golang image for building
#FROM golang:1.20-alpine AS builder
#
## Install necessary dependencies
#RUN apk add --no-cache librdkafka-dev
#
## Set working directory
#WORKDIR /app
#
## Copy go.mod and go.sum and run go mod tidy
#COPY go.mod go.sum ./
#RUN go mod tidy && go mod download
#RUN cat go.mod && cat go.sum  # Debugging step: show contents of go.mod and go.sum
#
#
## Copy source code and build the binary
#COPY . .
#RUN CGO_ENABLED=1 GOOS=linux go build -o api-server ./cmd/api
#
## Final minimal runtime stage
#FROM alpine:3.17
#
## Install librdkafka runtime library for Redpanda/Kafka compatibility
#RUN apk add --no-cache librdkafka
#
## Set up working directory
#WORKDIR /app
#
## Copy the built binary from the builder stage
#COPY --from=builder /app/api-server .
#
## Copy configuration files, if any
#COPY configs /app/configs
#
## Expose necessary ports (e.g., 8080 for HTTP and WebSocket/SSE)
#EXPOSE 8080
#
## Run the binary
#CMD ["./api-server"]
