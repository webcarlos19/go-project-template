# Go Project Template Makefile
.PHONY: help build run test clean docker-build docker-run docker-compose-up docker-compose-down lint format deps

# Variables
APP_NAME := go-project-template
BUILD_DIR := ./build
CMD_DIR := ./cmd/$(APP_NAME)
DOCKER_IMAGE := $(APP_NAME):latest
GO_VERSION := 1.21

# Default target
help: ## Show this help message
	@echo 'Usage: make [target]'
	@echo ''
	@echo 'Targets:'
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "  %-15s %s\n", $$1, $$2}' $(MAKEFILE_LIST)

# Development
run: ## Run the application locally
	@echo "Running $(APP_NAME)..."
	go run $(CMD_DIR)/main.go

run-prod: ## Run the application in production mode
	@echo "Running $(APP_NAME) in production mode..."
	APP_ENV=production go run $(CMD_DIR)/main.go

run-test: ## Run the application in test mode
	@echo "Running $(APP_NAME) in test mode..."
	APP_ENV=test go run $(CMD_DIR)/main.go

# Build
build: ## Build the application
	@echo "Building $(APP_NAME)..."
	mkdir -p $(BUILD_DIR)
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o $(BUILD_DIR)/$(APP_NAME) $(CMD_DIR)/main.go

build-windows: ## Build for Windows
	@echo "Building $(APP_NAME) for Windows..."
	mkdir -p $(BUILD_DIR)
	CGO_ENABLED=0 GOOS=windows go build -a -installsuffix cgo -o $(BUILD_DIR)/$(APP_NAME).exe $(CMD_DIR)/main.go

build-macos: ## Build for macOS
	@echo "Building $(APP_NAME) for macOS..."
	mkdir -p $(BUILD_DIR)
	CGO_ENABLED=0 GOOS=darwin go build -a -installsuffix cgo -o $(BUILD_DIR)/$(APP_NAME)-macos $(CMD_DIR)/main.go

# Dependencies
deps: ## Download dependencies
	@echo "Downloading dependencies..."
	go mod download
	go mod tidy

deps-update: ## Update dependencies
	@echo "Updating dependencies..."
	go get -u ./...
	go mod tidy

# Testing
test: ## Run tests
	@echo "Running tests..."
	go test -v ./...

test-coverage: ## Run tests with coverage
	@echo "Running tests with coverage..."
	go test -v -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html

test-race: ## Run tests with race detection
	@echo "Running tests with race detection..."
	go test -v -race ./...

benchmark: ## Run benchmarks
	@echo "Running benchmarks..."
	go test -bench=. -benchmem ./...

# Code quality
lint: ## Run linter
	@echo "Running linter..."
	golangci-lint run

format: ## Format code
	@echo "Formatting code..."
	go fmt ./...
	goimports -w .

vet: ## Run go vet
	@echo "Running go vet..."
	go vet ./...

# Docker
docker-build: ## Build Docker image
	@echo "Building Docker image..."
	docker build -t $(DOCKER_IMAGE) .

docker-run: ## Run Docker container
	@echo "Running Docker container..."
	docker run -p 8080:8080 --env-file .env $(DOCKER_IMAGE)

docker-compose-up: ## Start services with docker-compose
	@echo "Starting services with docker-compose..."
	docker-compose up -d

docker-compose-down: ## Stop services with docker-compose
	@echo "Stopping services with docker-compose..."
	docker-compose down

docker-compose-logs: ## View docker-compose logs
	@echo "Viewing docker-compose logs..."
	docker-compose logs -f

# Database
db-migrate: ## Run database migrations (placeholder)
	@echo "Running database migrations..."
	# Add your migration tool command here

db-seed: ## Seed database with test data (placeholder)
	@echo "Seeding database..."
	# Add your seed command here

# Cleanup
clean: ## Clean build artifacts
	@echo "Cleaning up..."
	rm -rf $(BUILD_DIR)
	rm -f coverage.out coverage.html
	go clean

clean-docker: ## Clean Docker images and containers
	@echo "Cleaning Docker artifacts..."
	docker system prune -f

# Security
security-scan: ## Run security scan
	@echo "Running security scan..."
	gosec ./...

# Release
release: clean test build ## Clean, test, and build for release
	@echo "Release build completed"

# Development setup
dev-setup: deps ## Setup development environment
	@echo "Setting up development environment..."
	@echo "Installing development tools..."
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	go install golang.org/x/tools/cmd/goimports@latest
	go install github.com/securecodewarrior/gosec/v2/cmd/gosec@latest
	@echo "Development environment setup complete!"

# Generate
generate: ## Run go generate
	@echo "Running go generate..."
	go generate ./...

# Module verification
mod-verify: ## Verify module dependencies
	@echo "Verifying module dependencies..."
	go mod verify

mod-why: ## Explain why packages are needed
	@echo "Explaining module dependencies..."
	go mod why -m all
