package infra

import (
	log "github.com/sirupsen/logrus"
	"github.com/xhsun/gcp-pubsub-ui/pubsub-ui-server/internal/config"
	"github.com/xhsun/gcp-pubsub-ui/pubsub-ui-server/internal/core"
)

type GCPPubSubRepositoryBuilder struct {
	config    *config.Config
	clients   map[string]*GCPPubSubRepository
	topicName string
}

// NewGCPPubSubRepositoryBuilder method creates a new GCPPubSubRepositoryBuilder
func NewGCPPubSubRepositoryBuilder(config *config.Config) *GCPPubSubRepositoryBuilder {
	return &GCPPubSubRepositoryBuilder{
		config:    config,
		clients:   make(map[string]*GCPPubSubRepository),
		topicName: "",
	}
}

// WithTopicName provide a topic name to the builder, which allow it to create a topic while building the client
func (prb *GCPPubSubRepositoryBuilder) WithTopicName(topicName string) core.IPubSubRepositoryBuilder {
	return &GCPPubSubRepositoryBuilder{
		config:    prb.config,
		clients:   prb.clients,
		topicName: topicName,
	}
}

// Build create a new client if there is no pre-existing client for that project
func (prb *GCPPubSubRepositoryBuilder) Build(gcpProjectID string) (core.IPubSubRepository, error) {
	client, exist := prb.clients[gcpProjectID]
	if !exist {
		client, err := NewGCPPubSubRepository(prb.config, gcpProjectID)
		if err != nil {
			log.WithField("gcpProjectID", gcpProjectID).WithError(err).Debug("Failed to create pubsub repository")
			return nil, err
		}
		prb.clients[gcpProjectID] = client
	}

	if prb.topicName != "" {
		return client, client.CreateSubscriber(prb.topicName)
	}

	return client, nil
}
