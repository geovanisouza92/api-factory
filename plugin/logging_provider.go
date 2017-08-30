package plugin

import (
	"net/rpc"

	"github.com/hashicorp/go-plugin"
	"google.golang.org/grpc"

	"github.com/geovanisouza92/api-factory/apifactory"
	"github.com/geovanisouza92/api-factory/plugin/proto"
)

type LoggingProviderPlugin struct {
	F LoggingProviderFunc
}

func (*LoggingProviderPlugin) Client(*plugin.MuxBroker, *rpc.Client) (interface{}, error) {
	return nil, errNotImplemented
}

func (*LoggingProviderPlugin) Server(*plugin.MuxBroker) (interface{}, error) {
	return nil, errNotImplemented
}

func (p *LoggingProviderPlugin) GRPCClient(c *grpc.ClientConn) (interface{}, error) {
	return &LoggingProvider{client: proto.NewLoggingProviderClient(c)}, nil
}

func (p *LoggingProviderPlugin) GRPCServer(s *grpc.Server) error {
	proto.RegisterAuthenticationProviderServer(s, &LoggingProviderServer{impl: p.F()})
	return nil
}

type LoggingProvider struct {
	client proto.LoggingProviderClient
}

// TODO: LoggingProvider deve implementar apifactory.LoggingProvider

type LoggingProviderServer struct {
	impl apifactory.LoggingProvider
}

// TODO: LoggingProviderServer deve implementar apifactory.LoggingProvider
