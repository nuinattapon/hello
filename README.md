# Hello - Go Web Application

This is a Go web application that demonstrates various HTTP handlers and features.

## Features
- Home page with server information
- JSON API endpoint
- HTML template rendering
- Prometheus metrics endpoint
- Ping endpoint for health checks
- Version endpoint
- Fibonacci calculator (optimized iterative implementation)

## Endpoints
- `GET /` - Home page with server information
- `GET /json` - Returns user data in JSON format
- `GET /template` - Renders HTML template with user data
- `GET /ping` - Returns "pong" for health checks
- `GET /version` - Returns application version
- `GET /fibo/{n}` - Calculates the nth Fibonacci number (0 ≤ n ≤ 45)
- `GET /metrics` - Prometheus metrics endpoint

## Configuration
The application can be configured using environment variables:
- `PORT` - Port to listen on (default: 80)

## Building and Running

### Using Go
```bash
go run .
```

### Using Docker
```bash
docker build -t hello .
docker run -p 8080:8080 hello
```

## Building for Different Architectures
```bash
# For Linux AMD64
GOOS=linux GOARCH=amd64 go build -o hello .

# For Linux ARM64
GOOS=linux GOARCH=arm64 go build -o hello .

# For Windows
GOOS=windows GOARCH=amd64 go build -o hello.exe .
```

## Improvements Made
1. **Performance**: Optimized Fibonacci calculation from O(2^n) recursive to O(n) iterative
2. **Error Handling**: Proper error handling with appropriate HTTP status codes
3. **Code Structure**: Organized code into packages (config, handlers, middleware, fibonacci)
4. **Logging**: Improved logging with structured format
5. **Configuration**: Made port configurable via environment variables
6. **Security**: Dockerfile updated to use non-root user
7. **Template**: Improved HTML template with better styling
