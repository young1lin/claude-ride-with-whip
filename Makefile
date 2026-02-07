# Makefile for claude-ride-with-whip

BINARY_NAME=statusline
GOEXE=$(shell go env GOEXE)
VERSION=$(shell cat VERSION 2>/dev/null || echo "0.1.0")
BIN_DIR=bin

.PHONY: build clean test release help

## build: Build the binary for the current platform
build:
	@echo "Building $(BINARY_NAME)..."
	@mkdir -p $(BIN_DIR)
	go build -o $(BIN_DIR)/$(BINARY_NAME)$(GOEXE) ./cmd/statusline
	@echo "Built $(BIN_DIR)/$(BINARY_NAME)$(GOEXE)"

## clean: Remove build artifacts
clean:
	@echo "Cleaning..."
	rm -f $(BIN_DIR)/$(BINARY_NAME)$(GOEXE)
	rm -f dist/*
	@echo "Cleaned"

## test: Run tests
test:
	@echo "Running tests..."
	go test -v -race -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html
	@echo "Tests completed"

## lint: Run linter
lint:
	@echo "Running linters..."
	golangci-lint run ./...

## fmt: Format code
fmt:
	@echo "Formatting code..."
	go fmt ./...
	gofmt -w -s .

## release: Create a new release (requires gh CLI)
release:
	@echo "Current version: $(VERSION)"
	@read -p "Enter new version (e.g., 0.2.0): " NEW_VERSION; \
	echo $$NEW_VERSION > VERSION; \
	git add VERSION; \
	git commit -m "chore: release v$$NEW_VERSION"; \
	git tag "v$$NEW_VERSION"; \
	git push origin main; \
	git push origin "v$$NEW_VERSION"

## build-all: Build for all platforms
build-all:
	@echo "Building for all platforms..."
	@mkdir -p dist
	@GOOS=windows GOARCH=amd64 go build -o dist/$(BINARY_NAME)_windows_amd64.exe ./cmd/statusline
	@GOOS=darwin GOARCH=amd64 go build -o dist/$(BINARY_NAME)_darwin_amd64 ./cmd/statusline
	@GOOS=darwin GOARCH=arm64 go build -o dist/$(BINARY_NAME)_darwin_arm64 ./cmd/statusline
	@GOOS=linux GOARCH=amd64 go build -o dist/$(BINARY_NAME)_linux_amd64 ./cmd/statusline
	@echo "Built binaries in dist/"

## help: Show this help message
help:
	@echo "claude-ride-with-whip Makefile"
	@echo ""
	@echo "Usage: make [target]"
	@echo ""
	@grep -E '^## ' $(MAKEFILE_LIST) | sed 's/## /  /'
