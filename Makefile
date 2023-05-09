PROJECT_NAME := "go-rest-api"
PKG := "gitlab.com/devops-projects6943118/$(PROJECT_NAME)"
PKG_LIST := $(shell go list ${PKG}/... | grep -v /vendor/)
BINARY_NAME=main

.PHONY: deps build clean test lint

build:
	echo "Compiling for every OS and Platform"
	mkdir -p out/bin
	GOARCH=amd64 GOOS=darwin go build -o out/bin/${BINARY_NAME}-darwin main.go
	GOARCH=amd64 GOOS=linux go build -o out/bin/${BINARY_NAME}-linux main.go
	GOARCH=amd64 GOOS=windows go build -o out/bin/${BINARY_NAME}-windows main.go

lint: ## Lint the files
	@golint -set_exit_status ${PKG_LIST}

deps:
	@go get -v -d ./...
	@go get -u github.com/golang/lint/golint

test:
	@go test -short ./src/tests/...

clean: ## Clean the project:
	go clean
	rm -f out/bin/${BINARY_NAME}-darwin
	rm -f out/bin/${BINARY_NAME}-linux
	rm -f out/bin/${BINARY_NAME}-windows