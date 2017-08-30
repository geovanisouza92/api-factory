package plugin

import (
	"net/rpc"

	"github.com/hashicorp/go-plugin"
	"google.golang.org/grpc"

	"github.com/geovanisouza92/api-factory/apifactory"
	"github.com/geovanisouza92/api-factory/plugin/proto"
)

type TrafficControlProviderPlugin struct {
	F TrafficControlProviderFunc
}

func (*TrafficControlProviderPlugin) Client(*plugin.MuxBroker, *rpc.Client) (interface{}, error) {
	return nil, errNotImplemented
}

func (*TrafficControlProviderPlugin) Server(*plugin.MuxBroker) (interface{}, error) {
	return nil, errNotImplemented
}

func (p *TrafficControlProviderPlugin) GRPCClient(c *grpc.ClientConn) (interface{}, error) {
	return &TrafficControlProvider{client: proto.NewTrafficControlProviderClient(c)}, nil
}

func (p *TrafficControlProviderPlugin) GRPCServer(s *grpc.Server) error {
	proto.RegisterAuthenticationProviderServer(s, &TrafficControlProviderServer{impl: p.F()})
	return nil
}

type TrafficControlProvider struct {
	client proto.TrafficControlProviderClient
}

// TODO: TrafficControlProvider deve implementar apifactory.TrafficControlProvider

type TrafficControlProviderServer struct {
	impl apifactory.TrafficControlProvider
}

// TODO: TrafficControlProviderServer deve implementar apifactory.TrafficControlProvider
