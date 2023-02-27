package core

import "context"

// IPubSubRepository provide functionalities to subscribe and receive message from a topic
type IPubSubRepository interface {
	// CreateSubscriber will create a new subscriber to the given topic if there is no pre-existing subscriber for that topic
	CreateSubscriber(topicName string) error
	// Receive passes the outstanding messages from the subscription to out channel.
	//
	// The standard way to terminate a Receive is to cancel its context:
	//
	//	cctx, cancel := context.WithCancel(ctx)
	//	pr.Receive(cctx, topicName, out)
	//	// Call cancel to end Receive
	Receive(ctx context.Context, topicName string, out chan<- []byte)
}
