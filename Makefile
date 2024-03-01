SHELL := bash
.PHONY: all server docs
.DEFAULT_GOAL := all

all: server bot

 server:
	@echo "Building server..."
	@cd licensing/ && make
	@echo "Server built successfully!"

docs:
	@echo "Building docs..."
	@cd docs/ && npm run build
	@echo "Docs built succesfully!"

bot:
	@echo "Building bot..."
	@cd bot/ && bundle install
	@echo "Bot built successfully!"

botStart:
	@echo "Starting bot..."
	@cd bot/ && bundle exec ruby main.rb &
	@echo "Bot started successfully!"

serverStart:
	@echo "Starting server..."
	@cd licensing/ && ./bin/corlink-server start &
	@echo "Server started successfully!"

start: serverStart botStart
