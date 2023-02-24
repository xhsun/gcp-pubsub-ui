package grpc

import (
	"fmt"
	"net"

	log "github.com/sirupsen/logrus"
	"github.com/xhsun/gcp-pubsub-ui/pubsub-ui-server/internal/config"
	"github.com/xhsun/gcp-pubsub-ui/pubsub-ui-server/internal/pubsubui"
	"google.golang.org/grpc"
)

// Server - The gRPC server
type Server struct {
	config         *config.Config
	messageHandler *MessageHandler
}

// NewServer method creates a new gRPC server
func NewServer(config *config.Config, messageHandler *MessageHandler) *Server {
	return &Server{
		config:         config,
		messageHandler: messageHandler,
	}
}

// Start starts the gRPC server
func (s *Server) Start() error {
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", s.config.Port))
	if err != nil {
		log.WithError(err).Fatalf("failed to listen to port %d", s.config.Port)
	}

	grpcServer := grpc.NewServer([]grpc.ServerOption{}...)
	pubsubui.RegisterPubSubUIServer(grpcServer, s.messageHandler)
	log.Info("Started PubSub UI server")
	return grpcServer.Serve(lis)
}
