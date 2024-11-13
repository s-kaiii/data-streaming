# Real-Time Data Streaming API

A Go-based real-time data streaming API that integrates with Redpanda (Kafka).

## How to Run
1. Clone this repository.
2. Install dependencies: `go mod tidy`
3. Run the API server: `go run cmd/api/main.go`

## Endpoints
- **POST /stream/start**: Start a new data stream.
- **POST /stream/{stream_id}/send**: Send data to the stream.
- **GET /stream/{stream_id}/results**: Fetch processed results.

## Configuration
Modify `configs/config.yaml` to update server and Kafka settings.
