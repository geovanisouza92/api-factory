package plugin

import (
	"net/rpc"

	"github.com/hashicorp/go-plugin"
	"google.golang.org/grpc"

	"github.com/geovanisouza92/api-factory/apifactory"
	"github.com/geovanisouza92/api-factory/apifactory/proto"
)

type ServiceProviderPlugin struct {
	F ServiceProviderFunc
}

func (*ServiceProviderPlugin) Client(*plugin.MuxBroker, *rpc.Client) (interface{}, error) {
	return nil, errNotImplemented
}

func (*ServiceProviderPlugin) Server(*plugin.MuxBroker) (interface{}, error) {
	return nil, errNotImplemented
}

func (p *ServiceProviderPlugin) GRPCClient(c *grpc.ClientConn) (interface{}, error) {
	return &ServiceProvider{client: proto.NewServiceProviderClient(c)}, nil
}

func (p *ServiceProviderPlugin) GRPCServer(s *grpc.Server) error {
	proto.RegisterAuthenticationProviderServer(s, &ServiceProviderServer{impl: p.F()})
	return nil
}

type ServiceProvider struct {
	client proto.ServiceProviderClient
}

// TODO: ServiceProvider deve implementar apifactory.ServiceProvider

type ServiceProviderServer struct {
	impl apifactory.ServiceProvider
}

// TODO: ServiceProviderServer deve implementar apifactory.ServiceProvider
