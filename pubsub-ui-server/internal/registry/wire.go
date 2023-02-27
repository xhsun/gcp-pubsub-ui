//go:build wireinject
// +build wireinject

package registry

import (
	"github.com/google/wire"
	"github.com/xhsun/gcp-pubsub-ui/pubsub-ui-server/internal/config"
	"github.com/xhsun/gcp-pubsub-ui/pubsub-ui-server/internal/core"
	"github.com/xhsun/gcp-pubsub-ui/pubsub-ui-server/internal/grpc"
	"github.com/xhsun/gcp-pubsub-ui/pubsub-ui-server/internal/infra"
)

var ServiceBuilderSet = wire.NewSet(
	infra.NewGCPPubSubRepository,
	wire.Bind(new(core.IPubSubRepository), new(*infra.GCPPubSubRepository)),
	infra.NewGCPPubSubRepositoryBuilder,
	wire.Bind(new(core.IPubSubRepositoryBuilder), new(*infra.GCPPubSubRepositoryBuilder)),
	core.NewMessageStreamService,
	wire.Bind(new(core.IMessageStreamService), new(*core.MessageStreamService)),
)

func InitializeServer(config *config.Config) (*grpc.Server, error) {
	wire.Build(ServiceBuilderSet, grpc.NewMessageHandler, grpc.NewHealthCheckHandler, grpc.NewServer)
	return &grpc.Server{}, nil
}
