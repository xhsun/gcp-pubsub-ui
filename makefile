SHELL := /bin/bash
PREFIX?=$(shell pwd)
GO := go
PROTOC := protoc

GO_LDFLAGS_STATIC:=-ldflags "-extldflags -static"
GO_SERVER_PATH:=$(PREFIX)/pubsub-ui-server
GO_SERVER_OUT:=$(PREFIX)/pubsubui_server

PROTO_DIR:= $(PREFIX)/pubsub-ui
PROTO_FILE := pubsub_ui.proto
PROTO_GO_OUTPUT:= $(GO_SERVER_PATH)/internal/pubsubui
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

build-server: ## Builds a static executable
	@echo "+ $@"
	cd $(GO_SERVER_PATH)/ && $(GO) build -tags "static_build" ${GO_LDFLAGS_STATIC} -o $(GO_SERVER_OUT) ./cmd

run-server: ## Run server
	@echo "+ $@"
	cd $(GO_SERVER_PATH)/ && $(GO) run cmd/main.go

clean:
	@echo "+ $@"
	$(RM) $(GO_SERVER_OUT)