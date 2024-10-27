SRCMAIN = ./cmd/main.go
IMAGE_NAME = number-search
IMAGE_TAG = latest

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

docker-build:
	docker build -t ${IMAGE_NAME}:${IMAGE_TAG} .
.PHONY: docker-build

docker-run:
	docker run -p 8080:8080 \
		-e PORT=":8080" \
		-e LOG_LEVEL="info" \
		-e VARIATION="10" \
		-e FILE_PATH="./input.txt" \
		${IMAGE_NAME}:${IMAGE_TAG}
.PHONY: docker-run

docker-up: ## Build and run docker container
	make docker-build
	make docker-run
.PHONY: docker-up