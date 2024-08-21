package grpc

import (
	"fmt"
	"net/http"

	"connectrpc.com/grpchealth"
	"connectrpc.com/grpcreflect"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"

	connectcors "connectrpc.com/cors"
	"github.com/rs/cors"
	"github.com/xhsun/gcp-pubsub-ui/pubsub-ui-server/internal/config"
	"github.com/xhsun/gcp-pubsub-ui/pubsub-ui-server/internal/pubsubui/pubsubuiconnect"
)

// Server - The gRPC server
type Server struct {
	config             *config.Config
	messageHandler     *MessageHandler
	healthCheckHandler *HealthCheckHandler
}

// NewServer method creates a new gRPC server
func NewServer(config *config.Config, messageHandler *MessageHandler, healthCheckHandler *HealthCheckHandler) *Server {
	return &Server{
		config:             config,
		messageHandler:     messageHandler,
		healthCheckHandler: healthCheckHandler,
	}
}

// Start starts the gRPC server
func (s *Server) Start() error {
	mux := http.NewServeMux()
	// protovalidate interceptor: https://github.com/connectrpc/validate-go
	// otel interceptor:
	//  - https://github.com/connectrpc/otelconnect-go
	//  - https://connectrpc.com/docs/go/observability
	mux.Handle(pubsubuiconnect.NewPubSubUIHandler(s.messageHandler))
	// gRPC health API support
	mux.Handle(grpchealth.NewHandler(grpchealth.NewStaticChecker(
		pubsubuiconnect.PubSubUIName,
	)))
	// gRPC reflect support
	reflector := grpcreflect.NewStaticReflector(
		pubsubuiconnect.PubSubUIName,
	)
	mux.Handle(grpcreflect.NewHandlerV1(reflector))
	mux.Handle(grpcreflect.NewHandlerV1Alpha(reflector)) // Support backwards compatibility
	return http.ListenAndServe(
		fmt.Sprintf("localhost:%d", s.config.Port),
		h2c.NewHandler(withCORS(mux), &http2.Server{}),
	)
}

// withCORS adds CORS support to a Connect HTTP handler.
func withCORS(h http.Handler) http.Handler {
	middleware := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: connectcors.AllowedMethods(),
		AllowedHeaders: connectcors.AllowedHeaders(),
		ExposedHeaders: connectcors.ExposedHeaders(),
	})
	return middleware.Handler(h)
}
