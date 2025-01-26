#vars

BINARY_NAME=challengefile
BINARY_PATH=/usr/bin/$(BINARY_NAME)
BINARY_TEST_PATH=./binary/$(BINARY_NAME)

.PHONY: build
build:
	@go build -o $(BINARY_TEST_PATH) -v

#test
.PHONY: test
test:
	@go test -v

#install dependencies
.PHONY: install-deps
install-deps:
	@go mod tidy

#install binary
.PHONY: install
install:
	@sudo go build -o $(BINARY_PATH) -v
