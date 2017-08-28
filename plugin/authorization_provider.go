package plugin

import (
	"net/rpc"

	"github.com/hashicorp/go-plugin"
	"google.golang.org/grpc"

	"github.com/geovanisouza92/api-factory/apifactory"
	"github.com/geovanisouza92/api-factory/apifactory/proto"
)

type AuthorizationProviderPlugin struct {
	F AuthorizationProviderFunc
}

func (*AuthorizationProviderPlugin) Client(*plugin.MuxBroker, *rpc.Client) (interface{}, error) {
	return nil, errNotImplemented
}

func (*AuthorizationProviderPlugin) Server(*plugin.MuxBroker) (interface{}, error) {
	return nil, errNotImplemented
}

func (p *AuthorizationProviderPlugin) GRPCClient(c *grpc.ClientConn) (interface{}, error) {
	return &AuthorizationProvider{client: proto.NewAuthorizationProviderClient(c)}, nil
}

func (p *AuthorizationProviderPlugin) GRPCServer(s *grpc.Server) error {
	proto.RegisterAuthenticationProviderServer(s, &AuthorizationProviderServer{impl: p.F()})
	return nil
}

type AuthorizationProvider struct {
	client proto.AuthorizationProviderClient
}

// TODO: AuthorizationProvider deve implementar apifactory.AuthorizationProvider

type AuthorizationProviderServer struct {
	impl apifactory.AuthorizationProvider
}

// TODO: AuthorizationProviderServer deve implementar apifactory.AuthorizationProvider
