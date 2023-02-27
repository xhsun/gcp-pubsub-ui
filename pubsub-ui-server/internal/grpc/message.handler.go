package grpc

import (
	"context"

	log "github.com/sirupsen/logrus"
	"github.com/xhsun/gcp-pubsub-ui/pubsub-ui-server/internal/core"
	"github.com/xhsun/gcp-pubsub-ui/pubsub-ui-server/internal/pubsubui"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type MessageHandler struct {
	pubsubui.UnimplementedPubSubUIServer
	streamer core.IMessageStreamService
}

// NewMessageHandler method creates a new MessageHandler
func NewMessageHandler(streamer core.IMessageStreamService) *MessageHandler {
	return &MessageHandler{
		streamer: streamer,
	}
}

// Fetch retrieves messages from PubSub and pass it to the caller
func (mh *MessageHandler) Fetch(topic *pubsubui.TopicSubscription, stream pubsubui.PubSubUI_FetchServer) error {
	if topic == nil || topic.GcpProjectId == "" || topic.PubsubTopicName == "" {
		log.Error("Topic information cannot be empty")
		return status.Error(codes.InvalidArgument, "Please provide valid topic information")
	}

	logger := log.WithField("projectID", topic.GcpProjectId).WithField("topicName", topic.PubsubTopicName)
	data := make(chan []byte, 1)
	cctx, cancel := context.WithCancel(stream.Context())
	go mh.streamer.Stream(cctx, topic.GcpProjectId, topic.PubsubTopicName, data)

	logger.Debug("Start retrieving PubSub messages")
	for {
		select {
		case <-stream.Context().Done():
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
				return status.Error(codes.Unavailable, "Encountered unexpected error, please try again")
			}
		}
	}
}
