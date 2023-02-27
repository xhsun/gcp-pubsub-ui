GO := go

GO_LDFLAGS_STATIC:=-ldflags "-extldflags -static"
GO_SERVER_OUT:=pubsubui_server

build-server: ## Builds a static executable
	@echo "+ $@"
	$(GO) build -tags "static_build" ${GO_LDFLAGS_STATIC} -o $(GO_SERVER_OUT) ./cmd

clean:
	@echo "+ $@"
	$(RM) $(GO_SERVER_OUT)