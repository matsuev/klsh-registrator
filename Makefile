.PHONY: run
run:
	@go mod tidy
	@go run ./cmd/registrator

.PHONY: build
build:
	@go mod tidy
	go build -o ./build/registrator -v ./cmd/registrator

.PHONY: tidy
tidy:
	go mod tidy
	
.DEFAULT_GOAL := run
