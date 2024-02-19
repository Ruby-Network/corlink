PATH := $(PATH):$(PWD)/bin
SHELL := env PATH=$(PATH) bash
.PHONY: config

config:
	@echo "Generating config files..."
	@./bin/pkl-gen-go.bin config/schema.pkl --base-path github.com/ruby-network/corlink/licensing
