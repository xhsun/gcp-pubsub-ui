package core

// IPubSubRepositoryBuilder used to build IPubSubRepository
type IPubSubRepositoryBuilder interface {
	// WithTopicName provide a topic name to the builder, which allow it to create a topic while building the client
	WithTopicName(topicName string) IPubSubRepositoryBuilder
	// Build create a new client if there is no pre-existing client for that project
	Build(gcpProjectID string) (IPubSubRepository, error)
}
