package grpc

import (
	"context"
	"errors"

	"connectrpc.com/connect"
	log "github.com/sirupsen/logrus"
	"github.com/xhsun/gcp-pubsub-ui/pubsub-ui-server/internal/core"
	"github.com/xhsun/gcp-pubsub-ui/pubsub-ui-server/internal/pubsubui"
	"github.com/xhsun/gcp-pubsub-ui/pubsub-ui-server/internal/pubsubui/pubsubuiconnect"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type MessageHandler struct {
	pubsubuiconnect.UnimplementedPubSubUIHandler
	streamer core.IMessageStreamService
}

// NewMessageHandler method creates a new MessageHandler
func NewMessageHandler(streamer core.IMessageStreamService) *MessageHandler {
	return &MessageHandler{
		streamer: streamer,
	}
}

// Fetch retrieves messages from PubSub and pass it to the caller
func (mh *MessageHandler) Fetch(ctx context.Context, request *connect.Request[pubsubui.TopicSubscription], stream *connect.ServerStream[pubsubui.Message]) error {
	if request == nil || request.Msg == nil || request.Msg.GcpProjectId == "" || request.Msg.PubsubTopicName == "" {
		log.Error("Topic information cannot be empty")
		return connect.NewWireError(connect.CodeInvalidArgument, errors.New("please provide valid topic information"))
	}

	projectID := request.Msg.GcpProjectId
	topicName := request.Msg.PubsubTopicName
	logger := log.WithField("projectID", projectID).WithField("topicName", topicName)
	data := make(chan []byte, 1)
	cctx, cancel := context.WithCancel(ctx)
	go mh.streamer.Stream(cctx, projectID, topicName, data)

	logger.Debug("Start retrieving PubSub messages")
	for {
		select {
		case <-ctx.Done():
			logger.Debug("Context cancelled, exit")
			cancel()
			return nil
		case message, ok := <-data:
			if ok {
				err := stream.Send(&pubsubui.Message{Data: message, Timestamp: timestamppb.Now()})
				if err != nil {
					logger.WithError(err).Error("Failed to send message to client")
					cancel()
					return err
				}
			} else {
				logger.Debug("data channel closed, no more data")
				cancel()
				return connect.NewWireError(connect.CodeUnavailable, errors.New("encountered unexpected error, please try again"))
			}
		}
	}
}

func (mh *MessageHandler) Echo(ctx context.Context, request *connect.Request[pubsubui.TopicSubscription]) (*connect.Response[pubsubui.TopicSubscription], error) {
	if request == nil || request.Msg == nil || request.Msg.GcpProjectId == "" || request.Msg.PubsubTopicName == "" {
		log.Error("Topic information cannot be empty")
		return nil, connect.NewWireError(connect.CodeInvalidArgument, errors.New("please provide valid topic information"))
	}
	projectID := request.Msg.GcpProjectId
	topicName := request.Msg.PubsubTopicName
	logger := log.WithField("projectID", projectID).WithField("topicName", topicName)
	logger.Debug("Start echoing")
	return connect.NewResponse(request.Msg), nil
}
