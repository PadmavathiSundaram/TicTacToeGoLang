GOCMD=go
GOBUILD=$(GOCMD) build
GOTEST=$(GOCMD) test
CMD_DIR=$(shell pwd)/cmd/game

local: all run
all: lint build test

build:
	go build --race -v ./...
test:
	rm -rf reports;
	mkdir reports;
	go test -mod=vendor --race -v -covermode=atomic -coverprofile=reports/coverage.out ./... >> reports/test_Result
	go tool cover -html=reports/coverage.out
lint:
	golangci-lint  run
run:
	cd ${CMD_DIR}; \
	go run ./...