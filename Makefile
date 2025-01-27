# Variables
BINARY_NAME := challengefile
BINARY_PATH := /usr/bin/$(BINARY_NAME)
BINARY_TEST_PATH := ./binary/$(BINARY_NAME)

.PHONY: all
all: build test

.PHONY: build
build:
	@echo "Building binary at $(BINARY_TEST_PATH)..."
	@go build -o $(BINARY_TEST_PATH) -v

# Run tests
.PHONY: test
test:
	@echo "Running tests..."
	@go test ./... -v

# Install dependencies
.PHONY: install-deps
install-deps:
	@echo "Tidying up dependencies..."
	@go mod tidy

.PHONY: install
install:
	@echo "Installing binary to $(BINARY_PATH)..."
	@go build -o $(BINARY_PATH) -v

.PHONY: clean
clean:
	@echo "Cleaning..."
	@rm -rf ./binary/$(BINARY_NAME)