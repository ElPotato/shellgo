.DEFAULT_GOAL := all
BIN_NAME := "shellgo"
LD_FLAGS := -ldflags='-s -w'

all: test build compress

test:
	@go test shell_test.go --count=1

build:
	@go build $(LD_FLAGS) -o $(BIN_NAME) *.go

compress: 
	@upx $(BIN_NAME)