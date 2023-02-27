ARG GO_VERSION=1-bullseye
FROM golang:$GO_VERSION AS build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -tags "static_build" -ldflags "-extldflags -static" -o pubsubui_server ./cmd

FROM golang:$GO_VERSION
WORKDIR  /server

ARG PUBSUB_EMULATOR_HOST=cloudpubsub:8086
ARG PUBSUB_UI_SERVER_VERBOSE=false
ARG PUBSUB_UI_SERVER_TIMEOUT=3
ENV PUBSUB_UI_SERVER_PORT=50051
EXPOSE 50051

COPY --from=build /app/pubsubui_server /server
CMD ["/server/pubsubui_server"]