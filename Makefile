GOPATH := $(shell go env GOPATH)

build:
	go build -o bin/gendiff ./cmd/gendiff

lint:
	$(GOPATH)/bin/golangci-lint run ./...

lint-fix:
	$(GOPATH)/bin/golangci-lint run --fix ./...
