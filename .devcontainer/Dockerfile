# [Choice] Go version (use -bullseye variants on local arm64/Apple Silicon): 1, 1.19, 1.18, 1-bullseye, 1.19-bullseye, 1.18-bullseye, 1-buster, 1.19-buster, 1.18-buster
ARG VARIANT=1-bullseye
FROM mcr.microsoft.com/vscode/devcontainers/go:0-${VARIANT}

# [Optional] Uncomment this section to install additional OS packages.
RUN apt-get update && export DEBIAN_FRONTEND=noninteractive \
    && apt-get -y install --no-install-recommends protobuf-compiler libprotobuf-dev

# [Optional] Uncomment the next lines to use go get to install anything else you need
USER vscode
RUN go env -w GOPROXY=https://proxy.golang.org,direct
RUN go install github.com/yoheimuta/protolint/cmd/protolint@latest
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
RUN go install github.com/google/wire/cmd/wire@latest

# Install protoc-gen-grpc-web protoc plugin
ARG GRPC_WEB_VERSION=1.4.2
ARG GRPC_WEB_ARCH=x86_64
RUN mkdir -p ~/.local/bin
RUN curl -LO https://github.com/grpc/grpc-web/releases/download/${GRPC_WEB_VERSION}/protoc-gen-grpc-web-${GRPC_WEB_VERSION}-linux-${GRPC_WEB_ARCH}
RUN mv protoc-gen-grpc-web-${GRPC_WEB_VERSION}-linux-${GRPC_WEB_ARCH} ~/.local/bin/protoc-gen-grpc-web
RUN chmod +x  ~/.local/bin/protoc-gen-grpc-web