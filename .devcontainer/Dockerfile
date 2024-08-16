# [Choice] Go version (use -bullseye variants on local arm64/Apple Silicon): 1, 1.19, 1.18, 1-bullseye, 1.19-bullseye, 1.18-bullseye, 1-buster, 1.19-buster, 1.18-buster
ARG VARIANT=1-bullseye
FROM mcr.microsoft.com/vscode/devcontainers/go:0-${VARIANT}

# [Optional] Uncomment this section to install additional OS packages.
RUN apt-get update && export DEBIAN_FRONTEND=noninteractive \
    && apt-get -y install --no-install-recommends protobuf-compiler libprotobuf-dev

# [Optional] Uncomment the next lines to use go get to install anything else you need
USER vscode
RUN go env -w GOPROXY=https://proxy.golang.org,direct
RUN go install github.com/bufbuild/buf/cmd/buf@latest
RUN go install github.com/google/wire/cmd/wire@latest