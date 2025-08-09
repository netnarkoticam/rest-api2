GOFILES = $(shell find . -type f -name '*.go')
LOCAL_BIN:=$(CURDIR)/bin
APP_NAME=order-service

ifeq ($(OS),Windows_NT)
	GOLANGCI_BIN:=$(LOCAL_BIN)/golangci-lint.exe
	GOTESTSUM_BIN:=$(LOCAL_BIN)/gotestsum.exe
else
	GOLANGCI_BIN:=$(LOCAL_BIN)/golangci-lint
	GOTESTSUM_BIN:=$(LOCAL_BIN)/gotestsum
endif

install-lint:
ifeq ($(wildcard $(GOLANGCI_BIN)),)
	$(info Downloading golangci-lint latest)
	GOBIN=$(LOCAL_BIN) go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
endif

install-gotestsum:
ifeq ($(wildcard $(GOTESTSUM_BIN)),)
	$(info Downloading gotestsum latest)
	GOBIN=$(LOCAL_BIN) go install gotest.tools/gotestsum@latest
endif

fmt: # Format code
	$(info Formatting...)
	@gofmt -s -w ${GOFILES}

lint: install-lint # Run lint
	$(info Running lint...)
	$(GOLANGCI_BIN) run --config=.golangci.yaml --fix ./...

.PHONY: build
build: # Build app
	$(info Building app locally...)
	go build -ldflags="-X main.version=$(shell git rev-parse --short HEAD)" -o "$(LOCAL_BIN)/$(APP_NAME)" ./cmd/app

test: install-gotestsum # Run all tests
	$(info Cleaning test cache...)
	go clean -testcache
	$(info Running tests...)
	$(GOTESTSUM_BIN) ./...

install-pre-push-hook: # Install pre push git hook
	$(info Installing pre-push hook...)
	@mkdir -p .git/hooks
	@cp -r scripts/pre-push .git/hooks/pre-push
	@chmod +x .git/hooks/pre-push
	$(info Pre-push hook installed.)

install-hooks: install-pre-push-hook # Install git hooks

clean-binaries: # Remove all binaries from bin and build folders
	rm -rf ${LOCAL_BIN}/*
	rm -rf $(CURDIR)/build/*

test-coverage: # Create and open HTML test coverage report via browser
	$(info Preparing HTML tests coverage profile...)
	go test ./... -coverprofile c.out
	go tool cover -html=c.out
	rm c.out

add-migration: # Add migration
	cd migrations/scheme; \
	goose create $(name) sql

help:
	@grep -E '^[a-zA-Z0-9 -]+:.*#'  Makefile | sort | while read -r l; do printf "\033[1;32m$$(echo $$l | cut -f 1 -d':')\033[00m:$$(echo $$l | cut -f 2- -d'#')\n"; done