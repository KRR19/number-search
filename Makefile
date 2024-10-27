SRCMAIN = ./cmd/main.go

run:
	PORT=":8080" LOG_LEVEL=info VARIATION=10 FILE_PATH="./input.txt" go run ${SRCMAIN}
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