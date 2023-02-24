package config

type Config struct {
	VerboseLog          bool   `json:"verbose" env:"PUBSUB_UI_SERVER_VERBOSE" env-default:"false"`
	Port                uint16 `json:"port" env:"PUBSUB_UI_SERVER_PORT" env-default:"50051"`
	Timeout             int    `json:"timeout" env:"PUBSUB_UI_SERVER_TIMEOUT" env-default:"3"`
	TopicSubscriberName string `json:"topicSubscriberName" env:"PUBSUB_UI_SERVER_SUBSCRIBER_NAME" env-default:"pubsubuisubscriber"`
}
