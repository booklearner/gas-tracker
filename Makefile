BINARY_NAME=gas-tracker

default: build
.PHONY: deps run server clean fmt

deps:
	go mod download all

build:
	go build -o ${BINARY_NAME} cmd/cli.go

run:
	go run cmd/cli.go

server:
	go run cmd/cli.go server

fmt:
	go fmt ./...

clean:
	go clean
	rm ${BINARY_NAME}

