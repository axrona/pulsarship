BINARY := pulsarship
BUILD_DIR := build
VERSION := $(shell git describe --tags --always --dirty)
#LDFLAGS := -s -w -X main.version=$(VERSION)

GOCMD := go
GOBUILD := $(GOCMD) build
GOTEST := $(GOCMD) test
GOVET := $(GOCMD) vet
GOFMT := gofmt -s -w
GOFILES := $(shell find . -type f -name '*.go' -not -path "./vendor/*")

.PHONY: all
all: fmt vet test build install

.PHONY: build
build:
#	$(GOBUILD) -ldflags "$(LDFLAGS)" -o $(BUILD_DIR)/$(BINARY) .
	$(GOBUILD) -o $(BUILD_DIR)/$(BINARY) .

.PHONY: install
install:
	sudo install -Dm755 $(BUILD_DIR)/$(BINARY) /usr/bin/$(BINARY)

.PHONY: fmt
fmt:
	$(GOFMT) $(GOFILES)

.PHONY: vet
vet:
	$(GOVET) ./...

.PHONY: test
test:
	$(GOTEST) ./...

.PHONY: clean
clean:
	rm -rf $(BUILD_DIR)

.PHONY: tidy
tidy:
	$(GOCMD) mod tidy

.PHONY: run
run:
	$(GOCMD) run .
