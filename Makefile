# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=boomb
BINARY_UNIX=$(BINARY_NAME)_unix

all: test build
build:
	$(GOBUILD) -ldflags="-s -w" -o bin/$(BINARY_NAME) -v
test:
	$(GOTEST) -v ./...
clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_UNIX)

.PHONY: all test clean