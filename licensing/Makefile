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

npm-linux: linux
	@echo "Moving files to npm folders..."
	@mv bin/corlink-server npm/@rubynetwork/linux-x64/corlink-server 
	@mv bin/corlink-server-arm npm/@rubynetwork/linux-arm/corlink-server 
	@mv bin/corlink-server-arm64 npm/@rubynetwork/linux-arm64/corlink-server 
	@mv bin/corlink-server-386 npm/@rubynetwork/linux-ia32/corlink-server 

npm-windows: windows
	@echo "Moving files to npm folders..."
	@mv bin/corlink-server.exe npm/@rubynetwork/win32-x64/corlink-server.exe 
	@mv bin/corlink-server-386.exe npm/@rubynetwork/win32-ia32/corlink-server.exe

npm-mac: mac
	@echo "Moving files to npm folders..."
	@mv bin/corlink-server-mac npm/@rubynetwork/darwin-x64/corlink-server 
	@mv bin/corlink-server-mac-arm64 npm/@rubynetwork/darwin-arm64/corlink-server

npm-publish:
	@echo "Publishing to npm..."
	@cd npm/@rubynetwork/linux-x64 && npm publish --access public
	@cd npm/@rubynetwork/linux-arm && npm publish --access public
	@cd npm/@rubynetwork/linux-arm64 && npm publish --access public
	@cd npm/@rubynetwork/linux-ia32 && npm publish --access public
	@cd npm/@rubynetwork/win32-x64 && npm publish --access public
	@cd npm/@rubynetwork/win32-ia32 && npm publish --access public
	@cd npm/@rubynetwork/darwin-x64 && npm publish --access public
	@cd npm/@rubynetwork/darwin-arm64 && npm publish --access public
	@cd npm/corlink-server && npm publish --access public

npm: npm-linux npm-windows npm-mac
