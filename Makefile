# Makefile for Brev Cloud SDK
# A vendor-agnostic Go SDK for managing clusterable, GPU-accelerated compute

# Variables
BINARY_NAME=compute
MODULE_NAME=github.com/brevdev/cloud
BUILD_DIR=build
COVERAGE_DIR=coverage

# Load environment variables from .env file if it exists
ifneq (,$(wildcard .env))
    include .env
    export
endif

# Go related variables
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOMOD=$(GOCMD) mod
GOLINT=golangci-lint
GOVET=$(GOCMD) vet
GOFMT=gofumpt
GOINSTALL=$(GOCMD) install
GOSEC=gosec

# Third party tools
ARTIFACT_GOLINT=github.com/golangci/golangci-lint/cmd/golangci-lint@v1.64.8
ARTIFACT_GOSEC=github.com/securego/gosec/v2/cmd/gosec@v2.22.8

# Build flags
LDFLAGS=-ldflags "-X main.Version=$(shell git describe --tags --always --dirty) -X main.BuildTime=$(shell date -u '+%Y-%m-%d_%H:%M:%S')"

# Default target
.PHONY: all
all: clean build test lint

# Build the project
.PHONY: build
build:
	@echo "Building $(BINARY_NAME)..."
	@mkdir -p $(BUILD_DIR)
	$(GOBUILD) ./...

# Build for multiple platforms
.PHONY: build-all
build-all: build-linux build-darwin build-windows

.PHONY: build-linux
build-linux:
	@echo "Building for Linux..."
	@mkdir -p $(BUILD_DIR)
	GOOS=linux GOARCH=amd64 $(GOBUILD) ./...

.PHONY: build-darwin
build-darwin:
	@echo "Building for macOS..."
	@mkdir -p $(BUILD_DIR)
	GOOS=darwin GOARCH=amd64 $(GOBUILD) ./...

.PHONY: build-windows
build-windows:
	@echo "Building for Windows..."
	@mkdir -p $(BUILD_DIR)
	GOOS=windows GOARCH=amd64 $(GOBUILD) ./...

# Run tests
.PHONY: test
test:
	@echo "Running tests..."
	$(GOTEST) -v -short ./...

# Run validation tests
.PHONY: test-validation
test-validation:
	@echo "Running validation tests..."
	$(GOTEST) -v -short=false ./...

# Run all tests including validation
.PHONY: test-all
test-all:
	@echo "Running all tests..."
	$(GOTEST) -v ./...

# Run tests with coverage
.PHONY: test-coverage
test-coverage:
	@echo "Running tests with coverage..."
	@mkdir -p $(COVERAGE_DIR)
	$(GOTEST) -v -coverprofile=$(COVERAGE_DIR)/coverage.out ./...
	$(GOCMD) tool cover -html=$(COVERAGE_DIR)/coverage.out -o $(COVERAGE_DIR)/coverage.html
	@echo "Coverage report generated at $(COVERAGE_DIR)/coverage.html"

# Run tests with race detection
.PHONY: test-race
test-race:
	@echo "Running tests with race detection..."
	$(GOTEST) -race -v ./...

# Run benchmarks
.PHONY: bench
bench:
	@echo "Running benchmarks..."
	$(GOTEST) -bench=. -benchmem ./...

# Lint the code
.PHONY: lint
lint:
	@echo "Running linter..."
	@if ! command -v $(GOLINT) > /dev/null 2>&1; then \
		$(GOINSTALL) $(ARTIFACT_GOLINT); \
	fi
	$(GOLINT) run ./...

# Run go vet
.PHONY: vet
vet:
	@echo "Running go vet..."
	$(GOVET) ./...

# Format code
.PHONY: fmt
fmt:
	@echo "Formatting code..."
	$(GOFMT) -w .

# Check if code is formatted
.PHONY: fmt-check
fmt-check:
	@echo "Checking code formatting..."
	@if [ "$$($(GOFMT) -l . | wc -l)" -gt 0 ]; then \
	@files=$$($(GOFMT) -s -l .); \
	if [ -n "$$files" ]; then \
		echo "Code is not formatted. Run 'make fmt' to format."; \
		$(GOFMT) -l .; \
		echo "$$files"; \
		exit 1; \
	else \
		echo "Code is properly formatted."; \
	fi

# Security scan
.PHONY: security
security:
	@echo "Running security scan..."
	@if ! command -v $(GOSEC) > /dev/null 2>&1; then \
		echo "gosec not found. Installing..."; \
		$(GOINSTALL) $(ARTIFACT_GOSEC); \
	fi
	$(GOSEC) ./...

# Clean build artifacts
.PHONY: clean
clean:
	@echo "Cleaning build artifacts..."
	$(GOCLEAN)
	rm -rf $(BUILD_DIR)
	rm -rf $(COVERAGE_DIR)

# Install dependencies
.PHONY: deps
deps:
	@echo "Installing dependencies..."
	$(GOMOD) download
	$(GOMOD) tidy

# Update dependencies
.PHONY: deps-update
deps-update:
	@echo "Updating dependencies..."
	$(GOCMD) get -u ./...
	$(GOMOD) tidy

# Verify dependencies
.PHONY: deps-verify
deps-verify:
	@echo "Verifying dependencies..."
	$(GOMOD) verify

# Generate documentation
.PHONY: docs
docs:
	@echo "Generating documentation..."
	$(GOCMD) doc -all ./...

# Run all checks (lint, vet, fmt-check, test)
.PHONY: check
check: lint vet fmt-check test

# Install tools
.PHONY: install-tools
install-tools:
	@echo "Installing development tools..."
	$(GOINSTALL) $(ARTIFACT_GOLINT)
	$(GOINSTALL) $(ARTIFACT_GOSEC)

# Show help
.PHONY: help
help:
	@echo "Available targets:"
	@echo "  build          - Build the project"
	@echo "  build-all      - Build for Linux, macOS, and Windows"
	@echo "  test           - Run tests (with -short flag)"
	@echo "  test-validation - Run validation tests (without -short flag)"
	@echo "  test-all       - Run all tests including validation"
	@echo "  test-coverage  - Run tests with coverage report"
	@echo "  test-race      - Run tests with race detection"
	@echo "  bench          - Run benchmarks"
	@echo "  lint           - Run linter (golangci-lint)"
	@echo "  vet            - Run go vet"
	@echo "  fmt            - Format code"
	@echo "  fmt-check      - Check if code is formatted"
	@echo "  security       - Run security scan (gosec)"
	@echo "  clean          - Clean build artifacts"
	@echo "  deps           - Install dependencies"
	@echo "  deps-update    - Update dependencies"
	@echo "  deps-verify    - Verify dependencies"
	@echo "  docs           - Generate documentation"
	@echo "  check          - Run all checks (lint, vet, fmt-check, test)"
	@echo "  install-tools  - Install development tools"
	@echo "  help           - Show this help message"  