.PHONY: help build run dev clean lint fmt install-tools

# Variables
BINARY_NAME=pharmacy
BINARY_PATH=./tmp/$(BINARY_NAME)
MAIN_FILE=./cmd/pharmacy/main.go
GO=go

# Default target
help:
	@echo "Available targets:"
	@echo "  make build          - Build the project"
	@echo "  make run            - Run the compiled binary"
	@echo "  make dev            - Run with air (hot reload)"
	@echo "  make clean          - Clean build artifacts and tmp files"
	@echo "  make fmt            - Format code with gofmt"
	@echo "  make lint           - Run linter (requires golangci-lint)"
	@echo "  make install-deps   - Install Go dependencies"
	@echo "  make install-tools  - Install development tools (air, golangci-lint)"

# Build the project
build:
	@echo "Building $(BINARY_NAME)..."
	@mkdir -p tmp
	$(GO) build -o $(BINARY_PATH) $(MAIN_FILE)
	@echo "Build complete: $(BINARY_PATH)"

# Run the compiled binary
run: build
	@echo "Running $(BINARY_NAME)..."
	$(BINARY_PATH)

# Run with air for hot reload
dev:
	@echo "Starting development server with air..."
	@if ! command -v air &> /dev/null; then \
		echo "air not found. Installing..."; \
		$(GO) install github.com/cosmtrek/air@latest; \
	fi
	air

# Format code
fmt:
	@echo "Formatting code..."
	$(GO) fmt ./...
	@echo "Code formatted"

# Run linter
lint:
	@echo "Running linter..."
	@if ! command -v golangci-lint &> /dev/null; then \
		echo "golangci-lint not found. Installing..."; \
		$(GO) install github.com/golangci/golangci-lint/cmd/golangci-lint@latest; \
	fi
	golangci-lint run ./...

# Install dependencies
install-deps:
	@echo "Installing dependencies..."
	$(GO) mod download
	$(GO) mod tidy
	@echo "Dependencies installed"

# Install development tools
install-tools: install-deps
	@echo "Installing development tools..."
	$(GO) install github.com/cosmtrek/air@latest
	$(GO) install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	@echo "Tools installed"

# Clean build artifacts and tmp files
clean:
	@echo "Cleaning..."
	@rm -rf tmp/
	$(GO) clean
	@echo "Clean complete"
