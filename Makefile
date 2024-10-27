SRCMAIN = ./cmd/main.go

run:
	go run ${SRCMAIN}
.PHONY: run

build:
	go build -o bin/number-search ${SRCMAIN}
.PHONY: build

fmt:
	go fmt ./...
.PHONY: fmt

lint:
	golangci-lint run ./...
.PHONY: lint

test:
	go test ./...
.PHONY: test

generate:
	go generate ./...
.PHONY: generate