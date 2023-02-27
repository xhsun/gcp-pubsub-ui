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

	client, err := mss.pubsubRepositoryBuilder.Build(projectID, topicName)
	if err != nil {
		logger.WithError(err).Error("Failed to create PubSub subscription")
		close(out)
		return
	}
	client.Receive(ctx, topicName, out)

	timeout := time.Duration(mss.defaultTimeout) * time.Second
	timer := time.NewTimer(timeout)
	defer timer.Stop()
	for {
		select {
		case <-ctx.Done():
			return
		case <-timer.C:
			out <- []byte{}
			timer.Reset(timeout)
		default:
			if len(out) > 0 {
				if !timer.Stop() {
					<-timer.C
				}
				timer.Reset(timeout)
			}
		}
	}
}
