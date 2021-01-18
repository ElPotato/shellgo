.DEFAULT_GOAL := bci
BIN_NAME := "shellgo"
VERSION := $(shell git describe --tags --always --dirty="")
LD_FLAGS := -ldflags='-X "main.BuildVersion=$(VERSION)" -s -w'

bci: build compress install

test:
	@go test -v shell_test.go --count=1

build:
	@go build $(LD_FLAGS) -o $(BIN_NAME) cmd/cli.go

compress: 
	@upx $(BIN_NAME)

install:
	@cp -f $(BIN_NAME) $(GOPATH)/bin/

remove:
	@rm -f $(GOPATH)/bin/$(BIN_NAME)