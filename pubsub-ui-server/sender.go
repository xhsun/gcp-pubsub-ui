package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"cloud.google.com/go/pubsub"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"syreclabs.com/go/faker"
)

const (
	projectID = "development"
	topicName = "example2"
)

type Example struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func main() {
	client, err := pubsub.NewClient(context.Background(), projectID)
	if err != nil {
		log.Fatal(err)
	}

	topic, err := client.CreateTopic(context.Background(), topicName)
	if err != nil {
		if status.Convert(err).Code() != codes.AlreadyExists {
			log.Fatal(err)
		}
		topic = client.Topic(topicName)
	}

	b, _ := json.Marshal(Example{
		ID:   uuid.NewString(),
		Name: faker.Name().Name(),
	})

	fmt.Printf("%s ::%s\n", topicName, b)
	if _, err := topic.Publish(context.Background(), &pubsub.Message{
		Data: b,
	}).Get(context.Background()); err != nil {
		fmt.Printf("Could not publish message: %v", err)
	}
}
