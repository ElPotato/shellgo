.DEFAULT_GOAL := bci
BIN_NAME := "shellgo"
LD_FLAGS := -ldflags='-s -w'

bci: build compress install

test:
	@go test shell_test.go --count=1

build:
	@go build $(LD_FLAGS) -o $(BIN_NAME) *.go

compress: 
	@upx $(BIN_NAME)

install:
	@cp -f $(BIN_NAME) $(GOPATH)/bin/

remove:
	@rm -f $(GOPATH)/bin/$(BIN_NAME)