# KSUI - Go Web Application

A simple Go web application that demonstrates HTMX integration with a clean, modern interface.

## Features

- **HTMX Integration**: Dynamic content loading without JavaScript
- **Go Backend**: Fast, efficient HTTP server
- **Template Rendering**: Server-side HTML templates
- **RESTful API**: Clean API endpoints for dynamic content

## Quick Start

### Prerequisites

- Go 1.23 or later
- Git

### Local Development

1. Clone the repository:
   ```bash
   git clone https://github.com/yourusername/ksui.git
   cd ksui
   ```

2. Install dependencies:
   ```bash
   go mod download
   ```

3. Run the application:
   ```bash
   go run main.go
   ```

4. Open your browser and navigate to `http://localhost:8080`

### Testing

Run the test suite:
```bash
go test -v
```

Run tests with coverage:
```bash
go test -v -cover
```

### Building

Build the application:
```bash
go build -o ksui main.go
```

## Docker

### Build and Run with Docker

```bash
# Build the Docker image
docker build -t ksui .

# Run the container
docker run -p 8080:8080 ksui
```

## CI/CD Pipeline

This project uses GitHub Actions for continuous integration and deployment. The pipeline includes:

### Workflow Jobs

1. **Test and Build** (`test`)
   - Runs on every push and pull request
   - Installs dependencies
   - Runs code linting with golangci-lint
   - Executes tests with race detection and coverage
   - Uploads coverage reports to Codecov
   - Builds the application binary
   - Uploads build artifacts

2. **Security Scan** (`security`)
   - Runs Trivy vulnerability scanner
   - Uploads results to GitHub Security tab
   - Scans for known vulnerabilities in dependencies

3. **Docker Build** (`docker`)
   - Only runs on main branch
   - Builds Docker image with caching
   - Uses multi-stage build for optimization

### Pipeline Features

- **Code Quality**: Automated linting with golangci-lint
- **Test Coverage**: Comprehensive test suite with coverage reporting
- **Security**: Vulnerability scanning with Trivy
- **Artifacts**: Build artifacts are preserved for 30 days
- **Caching**: Go module and Docker layer caching for faster builds
- **Parallel Jobs**: Security scanning runs in parallel with testing

### Local Development with CI Tools

Install golangci-lint locally:
```bash
go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
```

Run linting:
```bash
golangci-lint run
```

## Project Structure

```
ksui/
├── main.go              # Main application entry point
├── main_test.go         # Test suite
├── Dockerfile           # Docker configuration
├── .dockerignore        # Docker build exclusions
├── .golangci.yml        # Linting configuration
├── go.mod               # Go module definition
├── go.sum               # Go module checksums
├── templates/
│   └── index.html       # HTML template
└── .github/
    └── workflows/
        └── go.yml       # GitHub Actions workflow
```

## API Endpoints

- `GET /` - Main page with HTMX interface
- `GET /api/message` - Returns HTML snippet for dynamic loading

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests for new functionality
5. Ensure all tests pass
6. Submit a pull request

## License

This project is licensed under the MIT License. 