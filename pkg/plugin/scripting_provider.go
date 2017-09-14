package plugin

import (
	"net/rpc"

	"github.com/hashicorp/go-plugin"
	"google.golang.org/grpc"

	"github.com/geovanisouza92/api-factory/internal/proto"
	"github.com/geovanisouza92/api-factory/pkg/apifactory"
)

type ScriptingProviderPlugin struct {
	F ScriptingProviderFunc
}

func (*ScriptingProviderPlugin) Client(*plugin.MuxBroker, *rpc.Client) (interface{}, error) {
	return nil, errNotImplemented
}

func (*ScriptingProviderPlugin) Server(*plugin.MuxBroker) (interface{}, error) {
	return nil, errNotImplemented
}

func (p *ScriptingProviderPlugin) GRPCClient(c *grpc.ClientConn) (interface{}, error) {
	return &ScriptingProvider{client: proto.NewScriptingProviderClient(c)}, nil
}

func (p *ScriptingProviderPlugin) GRPCServer(s *grpc.Server) error {
	proto.RegisterAuthenticationProviderServer(s, &ScriptingProviderServer{impl: p.F()})
	return nil
}

type ScriptingProvider struct {
	client proto.ScriptingProviderClient
}

// TODO: ScriptingProvider deve implementar apifactory.ScriptingProvider

type ScriptingProviderServer struct {
	impl apifactory.ScriptingProvider
}

// TODO: ScriptingProviderServer deve implementar apifactory.ScriptingProvider
