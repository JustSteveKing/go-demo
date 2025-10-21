# Go Demo

A simple HTTP server built with Go demonstrating graceful shutdown, middleware, and clean architecture.

## Features

- ✨ Clean, organized code structure
- 🔄 Graceful shutdown handling
- 📝 Request logging middleware
- 🛡️ Panic recovery middleware
- 🏥 Health check endpoint
- 📦 Version information endpoint

## Getting Started

### Prerequisites

- Go 1.21 or higher

### Installation

```bash
git clone <repository-url>
cd go-demo
go mod download
```

### Running the Server

```bash
go run .
```

Or build and run:

```bash
go build
./go-demo
```

The server will start on `http://localhost:8000`

## API Endpoints

### `GET /`
Returns a simple greeting message.

```bash
curl http://localhost:8000/
```

### `GET /health`
Health check endpoint returning JSON status.

```bash
curl http://localhost:8000/health
```

Response:
```json
{"status":"ok"}
```

### `GET /version`
Returns build version information.

```bash
curl http://localhost:8000/version
```

Response:
```json
{"version":"dev","commit":"","built":""}
```

## Building with Version Info

You can inject build information at compile time:

```bash
go build -ldflags "-X main.buildVersion=1.0.0 -X main.buildCommit=$(git rev-parse HEAD) -X main.buildTime=$(date -u +%Y-%m-%dT%H:%M:%SZ)"
```

## Project Structure

```
.
├── main.go        # Server setup and configuration
├── handlers.go    # HTTP request handlers
├── middleware.go  # HTTP middleware functions
├── types.go       # Type definitions and build info
└── go.mod         # Go module definition
```

## Configuration

Configuration is done via constants in `main.go`:

- `serverPort` - Server listen address (default: `:8000`)
- `shutdownTimeout` - Graceful shutdown timeout (default: `10s`)
- `readTimeout` - HTTP read timeout (default: `5s`)
- `writeTimeout` - HTTP write timeout (default: `10s`)
- `idleTimeout` - HTTP idle timeout (default: `60s`)

## Graceful Shutdown

The server handles `SIGINT` and `SIGTERM` signals for graceful shutdown, allowing in-flight requests to complete before stopping.

## License

MIT
