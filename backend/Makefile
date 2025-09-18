.PHONY: build run dev migrate seed setup full-setup clean test

# Build the application
build:
	@echo "Building application..."
	@go build -o bin/server cmd/main.go

# Run the application
run: build
	@echo "Starting server..."
	@./bin/server

# Run in development mode
dev:
	@echo "Starting development server..."
	@go run cmd/main.go

# Run database migrations
migrate:
	@echo "Running database migrations..."
	@go run cmd/main.go -migrate

# Seed database with sample data
seed:
	@echo "Seeding database with sample data..."
	@go run cmd/main.go -seed

# Setup database (migrate + seed)
setup: migrate seed
	@echo "Database setup completed!"

# Full setup (build + migrate + seed)
full-setup: build migrate seed
	@echo "Full setup completed!"

# Clean build artifacts
clean:
	@echo "Cleaning build artifacts..."
	@rm -rf bin/

# Run tests
test:
	@echo "Running tests..."
	@go test ./...

# Install dependencies
deps:
	@echo "Installing dependencies..."
	@go mod tidy
	@go mod download

# Help
help:
	@echo "Available commands:"
	@echo "  build       - Build the application"
	@echo "  run         - Build and run the application"
	@echo "  dev         - Run in development mode"
	@echo "  migrate     - Run database migrations"
	@echo "  seed        - Seed database with sample data"
	@echo "  setup       - Run migrations and seed data"
	@echo "  full-setup  - Build, migrate, and seed"
	@echo "  clean       - Clean build artifacts"
	@echo "  test        - Run tests"
	@echo "  deps        - Install dependencies"
	@echo "  help        - Show this help message"
