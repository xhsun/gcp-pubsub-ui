package core

import (
	"context"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/xhsun/gcp-pubsub-ui/pubsub-ui-server/internal/config"
)

// IMessageStreamService provide streams of data from pubsub
type IMessageStreamService interface {
	// Stream passes the outstanding messages from the subscription to out channel.
	// If it didn't receive any message within PUBSUB_UI_SERVER_TIMEOUT, it will output an empty byte array to indicate no data.
	// It blocks until ctx is done, or the service returns a non-retryable error.
	//
	// It will close the out channel when ctx is done, or it encountered an error.
	//
	// The standard way to terminate a Stream is to cancel its context:
	//
	//	cctx, cancel := context.WithCancel(ctx)
	//	err := mss.Stream(cctx, projectID, topicName, out)
	//	// Call cancel to end Stream
	Stream(ctx context.Context, projectID string, topicName string, out chan<- []byte)
}

type MessageStreamService struct {
	defaultTimeout          int
	pubsubRepositoryBuilder IPubSubRepositoryBuilder
}

// NewMessageStreamService method creates a new MessageStreamService
func NewMessageStreamService(config *config.Config, pubsubRepositoryBuilder IPubSubRepositoryBuilder) *MessageStreamService {
	return &MessageStreamService{
		defaultTimeout:          config.Timeout,
		pubsubRepositoryBuilder: pubsubRepositoryBuilder,
	}
}

// Stream passes the outstanding messages from the subscription to out channel.
// If it didn't receive any message within PUBSUB_UI_SERVER_TIMEOUT, it will output an empty byte array to indicate no data.
// It blocks until ctx is done, or the service returns a non-retryable error.
//
// It will close the out channel when ctx is done, or it encountered an error.
//
// The standard way to terminate a Stream is to cancel its context:
//
//	cctx, cancel := context.WithCancel(ctx)
//	err := mss.Stream(cctx, projectID, topicName, out)
//	// Call cancel to end Stream
func (mss *MessageStreamService) Stream(ctx context.Context, projectID string, topicName string, out chan<- []byte) {
	logger := log.WithFields(log.Fields{"projectID": projectID, "topicName": topicName})

	client, err := mss.pubsubRepositoryBuilder.WithTopicName(topicName).Build(projectID)
	if err != nil {
		logger.WithError(err).Error("Failed to create PubSub subscription")
		close(out)
		return
	}

	data := make(chan []byte, 1)
	if err := client.Receive(ctx, topicName, data); err != nil {
		logger.WithError(err).Error("Failed to receive data from PubSub subscription")
		close(out)
		return
	}

	for {
		select {
		case <-ctx.Done():
			logger.Debug("Context cancelled, close the output channel")
			close(out)
			return
		case message, ok := <-data:
			if ok {
				out <- message
			} else {
				logger.Debug("data channel closed, no more data to pass on")
				close(out)
				return
			}
		case <-time.After(time.Duration(mss.defaultTimeout) * time.Second):
			out <- []byte{}
		}
	}
}
