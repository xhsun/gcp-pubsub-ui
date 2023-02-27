package infra

import (
	"context"
	"errors"

	"cloud.google.com/go/pubsub"
	log "github.com/sirupsen/logrus"
	"github.com/xhsun/gcp-pubsub-ui/pubsub-ui-server/internal/config"
)

var (
	ErrSubscriberNotFound = errors.New("subscriber not found")
)

type GCPPubSubRepository struct {
	client                *pubsub.Client
	subscribers           map[string]*pubsub.Subscription
	defaultSubscriberName string
}

// NewGCPPubSubRepository method creates a new GCPPubSubRepository
func NewGCPPubSubRepository(config *config.Config, GCPProjectID string) (*GCPPubSubRepository, error) {
	client, err := pubsub.NewClient(context.Background(), GCPProjectID)
	if err != nil {
		log.WithField("GCPProjectID", GCPProjectID).WithError(err).Debug("Failed to create GCP PubSub Client")
		return nil, err
	}
	return &GCPPubSubRepository{
		client:                client,
		subscribers:           make(map[string]*pubsub.Subscription),
		defaultSubscriberName: config.TopicSubscriberName,
	}, nil
}

// CreateSubscriber will create a new subscriber to the given topic if there is no pre-existing subscriber for that topic
func (pr *GCPPubSubRepository) CreateSubscriber(topicName string) error {
	_, exists := pr.subscribers[topicName]
	if exists {
		return nil
	}
	logger := log.WithField("topicName", topicName)

	topic := pr.client.Topic(topicName)
	exists, err := topic.Exists(context.Background())
	if err != nil {
		logger.WithError(err).Debug("Failed to check if GCP PubSub topic exist or not")
		return err
	}
	if !exists {
		logger.Debug("Topic not found, creating it")
		if _, err = pr.client.CreateTopic(context.Background(), topicName); err != nil {
			logger.WithError(err).Debug("Failed to create GCP PubSub topic")
			return err
		}
	}

	subscriber := pr.client.Subscription(pr.defaultSubscriberName)
	exists, err = subscriber.Exists(context.Background())
	if err != nil {
		logger.WithError(err).Debug("Failed to check if GCP PubSub subscriber exist or not")
		return err
	}
	if !exists {
		if _, err = pr.client.CreateSubscription(context.Background(), pr.defaultSubscriberName, pubsub.SubscriptionConfig{Topic: topic}); err != nil {
			logger.WithError(err).Debug("Failed to create GCP PubSub topic subscriber")
			return err
		}
	}
	pr.subscribers[topicName] = subscriber
	return nil
}

// Receive passes the outstanding messages from the subscription to out channel.
// It returns ErrSubscriberNotFound if subscriber not found, call CreateSubscriber to create one.
//
// The standard way to terminate a Receive is to cancel its context:
//
//	cctx, cancel := context.WithCancel(ctx)
//	err := pr.Receive(cctx, topicName, out)
//	// Call cancel to end Receive
func (pr *GCPPubSubRepository) Receive(ctx context.Context, topicName string, out chan<- []byte) error {
	subscriber, exist := pr.subscribers[topicName]
	if !exist {
		return ErrSubscriberNotFound
	}

	go subscriber.Receive(ctx, func(c context.Context, message *pubsub.Message) {
		select {
		case <-c.Done():
			log.Debug("Context cancelled, close the output channel")
			close(out)
		case out <- message.Data:
			message.Ack()
		}
	})

	return nil
}
