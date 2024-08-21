SHELL := /bin/bash
PREFIX?=$(shell pwd)
GO := go
NPM := npm

GO_SERVER_PATH:=$(PREFIX)/pubsub-ui-server
ANGULAR_CLIENT_PATH:=$(PREFIX)/pubsub-ui-client
REACT_CLIENT_PATH:=$(PREFIX)/openapi/test-app

run-server: ## Run server
	@echo "+ $@"
	cd $(GO_SERVER_PATH)/ && $(GO) run cmd/main.go

run-angular-client: ## Run client
	@echo "+ $@"
	cd $(ANGULAR_CLIENT_PATH)/ && $(NPM) run start

run-react-client: ## Run client
	@echo "+ $@"
	cd $(REACT_CLIENT_PATH)/ && $(NPM) run start

build-rtk:
	cd openapi/client && \
	npm run spec:bundle && \
	npm run spec:bundle:json && \
	npm run codegen && \
	npm run build