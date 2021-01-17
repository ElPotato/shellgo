.DEFAULT_GOAL := all
BIN_NAME := "shellgo"
LD_FLAGS := -ldflags='-s -w'

all: build compress

build:
	@go build $(LD_FLAGS) -o $(BIN_NAME) *.go

compress: 
	@upx $(BIN_NAME)