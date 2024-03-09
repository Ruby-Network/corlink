SHELL := bash
.PHONY: build clean download run linux windows mac default compress compress-output compress-linux compress-windows compress-mac compress-all format npm-linux npm-windows npm-mac npm
.DEFAULT_GOAL := default

clean:
	@echo "Cleaning up..."
	@rm -rf build

download:
	@echo "Downloading dependencies..."
	@go get

run:
	@echo "Running..."
	@go run main.go

linux: download
	@echo "Building for Linux..."
	@GOOS=linux GOARCH=amd64 go build -o bin/corlink-server -ldflags "-s -w" main.go
	@GOOS=linux GOARCH=arm go build -o bin/corlink-server-arm -ldflags "-s -w" main.go
	@GOOS=linux GOARCH=arm64 go build -o bin/corlink-server-arm64 -ldflags "-s -w" main.go
	@GOOS=linux GOARCH=386 go build -o bin/corlink-server-386 -ldflags "-s -w" main.go

windows: download
	@echo "Building for Windows..."
	@GOOS=windows GOARCH=amd64 go build -o bin/corlink-server.exe -ldflags "-s -w" main.go
	@GOOS=windows GOARCH=386 go build -o bin/corlink-server-386.exe -ldflags "-s -w" main.go

mac: download 
	@echo "Building for Mac..."
	@GOOS=darwin GOARCH=amd64 go build -o bin/corlink-server-mac -ldflags "-s -w" main.go
	@GOOS=darwin GOARCH=arm64 go build -o bin/corlink-server-mac-arm64 -ldflags "-s -w" main.go

default: clean download 
	@echo "Building for current OS..."
	@go build -o bin/corlink-server -ldflags "-s -w" main.go

compress:
	@echo "Compressing with no output file..."
	@upx --brute bin/corlink-server 

compress-output:
	@echo "Compressing with output file..."
	@upx --brute -o bin/corlink-server-compressed bin/corlink-server 

compress-linux:
	@echo "Compressing for Linux..."
	@upx --brute bin/corlink-server --output=bin/corlink-server-compressed
	@upx --brute bin/corlink-server-arm --output=bin/corlink-server-arm-compressed
	@upx --brute bin/corlink-server-arm64 --output=bin/corlink-server-arm64-compressed
	@upx --brute bin/corlink-server-386 --output=bin/corlink-server-386-compressed

compress-windows:
	@echo "Compressing for Windows..."
	@upx --brute bin/corlink-server.exe --output=bin/corlink-server-compressed.exe
	@upx --brute bin/corlink-server-386.exe --output=bin/corlink-server-386-compressed.exe

compress-mac:
	@echo "Compressing for Mac..."
	@upx --brute bin/corlink-server-mac --output=bin/corlink-server-mac-compressed
	@upx --brute bin/corlink-server-mac-arm64 --output=bin/corlink-server-mac-arm64-compressed

compress-all: compress-linux compress-windows compress-mac
	@echo "Compressing all files..."

format:
	@echo "Formatting..."
	@go fmt

all: clean download linux windows mac compress-all
