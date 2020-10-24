.EXPORT_ALL_VARIABLES:
NAME := rides
PKG := github.com/ajbosco/rides/cmd/rides
BUILD_DIR := $(shell pwd)/build
TARGET := ${BUILD_DIR}/${NAME}
VERSION := $(shell cat VERSION.txt)
LDFLAGS ?= -X github.com/ajbosco/rides/version.VERSION=${VERSION}

.PHONY: fmt
fmt: ## Verifies all files have been `gofmt`ed.
	@gofmt -s -l . | grep -v vendor | tee /dev/stderr

.PHONY: lint
lint: ## Verifies `golint` passes.
	@golint ./... | grep -v vendor | tee /dev/stderr

.PHONY: test
test: ## Runs the go tests.
	@go test -cover -race $(shell go list ./... | grep -v vendor)

.PHONY: vet
vet: ## Verifies `go vet` passes.
	@go vet $(shell go list ./... | grep -v vendor) | tee /dev/stderr

.PHONY: build
build: ## Run go build for current OS
	@go build -ldflags "$(LDFLAGS)" -o "${TARGET}" ${PKG}

.PHONY: clean
clean: ## Cleanup any build binaries or packages.
	@$(RM) -r $(BUILD_DIR)

.PHONY: help
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
