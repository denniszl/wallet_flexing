VERSION:=$(shell git rev-list --count HEAD)-$(shell git rev-parse --short HEAD)
DATE:=$(shell date -u '+%Y-%m-%d-%H%M UTC')
PATH:=$(GOPATH)/bin:$(PATH)

.PHONY: help
help:
	@echo 'Available commands:'
	@echo '* help                   - Show this message'
	@echo '* lint                   - Lint code'
	@echo '* test                   - Run tests'
	@echo '* build                  - Build binaries'
	@echo '* clean                  - Clean built binaries'

.PHONY: build
build: bin/api

bin/api:
	@echo "building api..."
	go build -o bin/api ./cmd/api

.PHONY: lint
lint:
	golangci-lint run --deadline=90s ./...

.PHONY: test
test:
	@go test -v -race $$(go list ./... | grep -v /vendor/)

.PHONY: clean
clean:
	rm -rf bin