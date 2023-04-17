.DEFAULT_GOAL := build
VERSION ?= $(shell git describe --tags --always --match=v* 2> /dev/null || echo v0)
COMMIT = $(shell git rev-parse HEAD)
BUILD_TS = $(shell date +%FT%T%:z)
MODULEPATH := $(shell go mod edit -json 2> /dev/null | jq -r '.Module.Path')

PLATFORM=local

# Runs go lint
.PHONY: lint
lint:
	@echo "Linting..."
	@golangci-lint run

# Runs go clean
.PHONY: clean
clean:
	@echo "Cleaning..."
	@go clean

# Runs go fmt
.PHONY: fmt
fmt:
	@echo "Formatting..."
	@go fmt ./...

# Runs go build
.PHONY: build
build: clean fmt lint
	@echo "Building k8s-inventory..."
	CGO_ENABLED=0 go build \
		-v \
		-a \
		-tags release \
		-ldflags '-s -w -X main.version=$(VERSION) -X main.commit=$(COMMIT) -X main.date=$(BUILD_TS)'


# Runs go build
.PHONY: build2
build2: clean fmt
	@echo "Building k8s-inventory"
	CGO_ENABLED=0 go build \
		-v \
		-tags release \
		-ldflags '-s -w -X main.version=$(VERSION) -X main.commit=$(COMMIT) -X main.date=$(BUILD_TS)'
