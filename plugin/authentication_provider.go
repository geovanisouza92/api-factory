package plugin

import (
	"net/rpc"

	"github.com/hashicorp/go-plugin"
	"google.golang.org/grpc"

	"github.com/geovanisouza92/api-factory/apifactory"
	"github.com/geovanisouza92/api-factory/plugin/proto"
)

type AuthenticationProviderPlugin struct {
	// NOTE: Necessário apenas dentro dos plugins, que devem fornecer uma implementação válida
	F AuthenticationProviderFunc
}

func (*AuthenticationProviderPlugin) Client(*plugin.MuxBroker, *rpc.Client) (interface{}, error) {
	return nil, errNotImplemented
}

func (*AuthenticationProviderPlugin) Server(*plugin.MuxBroker) (interface{}, error) {
	return nil, errNotImplemented
}

func (p *AuthenticationProviderPlugin) GRPCClient(c *grpc.ClientConn) (interface{}, error) {
	return &AuthenticationProvider{client: proto.NewAuthenticationProviderClient(c)}, nil
}

func (p *AuthenticationProviderPlugin) GRPCServer(s *grpc.Server) error {
	proto.RegisterAuthenticationProviderServer(s, &AuthenticationProviderServer{Impl: p.F()})
	return nil
}

type AuthenticationProvider struct {
	client proto.AuthenticationProviderClient
}

// TODO: AuthenticationProvider deve implementar apifactory.AuthenticationProvider

type AuthenticationProviderServer struct {
	Impl apifactory.AuthenticationProvider
}

// TODO: AuthenticationProviderServer deve implementar apifactory.AuthenticationProvider
