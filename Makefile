#vars

BINARY_NAME=challengefile
BINARY_PATH=/usr/bin/$(BINARY_NAME)

#build
.PHONY: build
build:
	@go build -o $(BINARY_NAME) -v

#test
.PHONY: test
test:
	@go test -v

#install dependencies
install:
	@go mod tidy

