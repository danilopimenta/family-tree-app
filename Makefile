export GO111MODULE=on

-include .env
export

# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=go-echo
BINARY_UNIX=$(BINARY_NAME)_unix

test:
	$(GOTEST) -v ./...

generate:
	$(GOCMD) generate ./...

tidy:
	$(GOCMD) mod tidy -v

fmt:
	$(GOCMD) fmt ./...

goimports:
	$(GOCMD) install golang.org/x/tools/cmd/goimports@latest
	goimports -local github.com/danilopimenta/family-tree-app -w .
