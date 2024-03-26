.PHONY: default run run-with-docs build test docs clean

APP_NAME=go-hire

default: run-with-docs

run:
	@go run cmd/main.go

run-with-docs:
	@swag init
	@go run cmd/main.go

build:
	@go build -o $(APP_NAME) cmd/main.go

test:
	@go test ./...

docs:
	@swag init

clean:
	@rm -f $(APP_NAME)
	@rm -rf docs