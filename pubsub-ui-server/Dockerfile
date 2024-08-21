ARG GO_VERSION=1-bullseye
FROM golang:$GO_VERSION AS build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -tags "static_build" -ldflags "-extldflags -static" -o pubsubui_server ./cmd

FROM alpine

WORKDIR /server

ARG PUBSUB_HOST=cloudpubsub:8086
ARG VERBOSE=false
ARG TIMEOUT=3
ENV PUBSUB_UI_SERVER_PORT=50051
ENV PUBSUB_EMULATOR_HOST=${PUBSUB_HOST}
ENV PUBSUB_UI_SERVER_VERBOSE=${VERBOSE}
ENV PUBSUB_UI_SERVER_TIMEOUT=${TIMEOUT}
EXPOSE 50051

COPY --from=build /app/pubsubui_server /server
CMD ["/server/pubsubui_server"]