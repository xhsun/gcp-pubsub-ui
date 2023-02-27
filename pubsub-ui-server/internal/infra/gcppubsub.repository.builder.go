package infra

import (
	"context"
	"fmt"

	"cloud.google.com/go/pubsub"
	log "github.com/sirupsen/logrus"
	"github.com/xhsun/gcp-pubsub-ui/pubsub-ui-server/internal/config"
	"github.com/xhsun/gcp-pubsub-ui/pubsub-ui-server/internal/core"
	"golang.org/x/exp/slices"
)

type GCPPubSubRepositoryBuilder struct {
	config         *config.Config
	clients        map[string]*pubsub.Client
	subscriberKeys []string
}

// NewGCPPubSubRepositoryBuilder method creates a new GCPPubSubRepositoryBuilder
func NewGCPPubSubRepositoryBuilder(config *config.Config) *GCPPubSubRepositoryBuilder {
	return &GCPPubSubRepositoryBuilder{
		config:         config,
		clients:        make(map[string]*pubsub.Client),
		subscriberKeys: []string{},
	}
}

// Build create a new client if there is no pre-existing client for that project
func (prb *GCPPubSubRepositoryBuilder) Build(gcpProjectID string, topicName string) (core.IPubSubRepository, error) {
	logger := log.WithField("GCPProjectID", gcpProjectID)
	client, exist := prb.clients[gcpProjectID]
	if !exist {
		temp, err := pubsub.NewClient(context.Background(), gcpProjectID)
		if err != nil {
			logger.WithError(err).Debug("Failed to create GCP PubSub Client")
			return nil, err
		}
		prb.clients[gcpProjectID] = temp
		client = temp
	}

	repo := NewGCPPubSubRepository(prb.config, client)
	key := fmt.Sprintf("%s%s", gcpProjectID, topicName)
	if !slices.Contains(prb.subscriberKeys, key) {
		if err := repo.CreateSubscriber(topicName); err != nil {
			logger.WithField("topicName", topicName).WithError(err).Debug("Failed to create GCP PubSub topic subscription")
			return nil, err
		}
		prb.subscriberKeys = append(prb.subscriberKeys, key)
	}

	return repo, nil
}
