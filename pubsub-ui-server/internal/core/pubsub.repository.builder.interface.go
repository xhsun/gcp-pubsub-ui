package core

// IPubSubRepositoryBuilder used to build IPubSubRepository
type IPubSubRepositoryBuilder interface {
	// Build create a new client if there is no pre-existing client for that project
	Build(gcpProjectID string, topicName string) (IPubSubRepository, error)
}
