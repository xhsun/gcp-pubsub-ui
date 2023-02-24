SHELL := /bin/bash
PREFIX?=$(shell pwd)
GO := go
PROTOC := protoc

PROTO_DIR:= $(PREFIX)/pubsub-ui
PROTO_FILE := pubsub_ui.proto
PROTO_GO_OUTPUT:= $(PREFIX)/pubsub-ui-server/internal/pubsubui
PROTO_ANGULAR_OUTPUT:=$(PREFIX)/pubsub-ui-client/src/app/core/pubsubui

gen-grpc: gen-go-grpc gen-web-grpc

gen-go-grpc: ## Generate gRPC Golang server and client stub
	@echo "+ $@"
	mkdir -p $(PROTO_GO_OUTPUT)
	$(PROTOC) -I=$(PROTO_DIR) $(PROTO_FILE) --go_out=$(PROTO_GO_OUTPUT) --go_opt=paths=source_relative --go-grpc_out=$(PROTO_GO_OUTPUT) --go-grpc_opt=paths=source_relative

gen-web-grpc: ## Generate gRPC web client stub
	@echo "+ $@"
	mkdir -p $(PROTO_ANGULAR_OUTPUT)
	$(PROTOC) -I=$(PROTO_DIR) $(PROTO_FILE) --js_out=import_style=commonjs,binary:$(PROTO_ANGULAR_OUTPUT) --grpc-web_out=import_style=typescript,mode=grpcweb:$(PROTO_ANGULAR_OUTPUT)
