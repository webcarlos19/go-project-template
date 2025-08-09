#!/bin/bash

# Setup script for development environment
set -e

echo "🚀 Setting up Go Project Template development environment..."

# Check if Go is installed
if ! command -v go &> /dev/null; then
    echo "❌ Go is not installed. Please install Go 1.21 or later."
    exit 1
fi

# Check Go version
GO_VERSION=$(go version | cut -d' ' -f3 | cut -d'o' -f2)
echo "✅ Go version: $GO_VERSION"

# Create .env file if it doesn't exist
if [ ! -f .env ]; then
    echo "📋 Creating .env file from .env.example..."
    cp .env.example .env
else
    echo "✅ .env file already exists"
fi

# Download dependencies
echo "📦 Downloading Go dependencies..."
go mod download
go mod tidy

# Install development tools
echo "🔧 Installing development tools..."
go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
go install golang.org/x/tools/cmd/goimports@latest
go install github.com/securecodewarrior/gosec/v2/cmd/gosec@latest

# Create build directory
mkdir -p build

echo "✅ Development environment setup complete!"
echo ""
echo "📖 Next steps:"
echo "   1. Review and update .env file with your configuration"
echo "   2. Run 'make run' to start the development server"
echo "   3. Visit http://localhost:8080/health to verify the application is running"
echo "   4. Check out the README.md for more information"
