SHELL := /bin/bash
PREFIX?=$(shell pwd)
GO := go
NPM := npm

GO_SERVER_PATH:=$(PREFIX)/pubsub-ui-server
ANGULAR_CLIENT_PATH:=$(PREFIX)/pubsub-ui-client

run-server: ## Run server
	@echo "+ $@"
	cd $(GO_SERVER_PATH)/ && $(GO) run cmd/main.go

run-client: ## Run client
	@echo "+ $@"
	cd $(ANGULAR_CLIENT_PATH)/ && $(NPM) run start
