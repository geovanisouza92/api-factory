package plugin

import (
	"net/rpc"

	"github.com/hashicorp/go-plugin"
	"google.golang.org/grpc"

	"github.com/geovanisouza92/api-factory/apifactory"
	"github.com/geovanisouza92/api-factory/plugin/proto"
)

type AnalyticsProviderPlugin struct {
	// NOTE: Necessário apenas dentro dos plugins, que devem fornecer uma implementação válida
	F AnalyticsProviderFunc
}

func (*AnalyticsProviderPlugin) Client(*plugin.MuxBroker, *rpc.Client) (interface{}, error) {
	return nil, errNotImplemented
}

func (*AnalyticsProviderPlugin) Server(*plugin.MuxBroker) (interface{}, error) {
	return nil, errNotImplemented
}

func (p *AnalyticsProviderPlugin) GRPCClient(c *grpc.ClientConn) (interface{}, error) {
	return &AnalyticsProvider{client: proto.NewAnalyticsProviderClient(c)}, nil
}

func (p *AnalyticsProviderPlugin) GRPCServer(s *grpc.Server) error {
	proto.RegisterAnalyticsProviderServer(s, &AnalyticsProviderServer{impl: p.F()})
	return nil
}

type AnalyticsProvider struct {
	client proto.AnalyticsProviderClient
}

// TODO: AnalyticsProvider deve implementar apifactory.AnalyticsProvider

type AnalyticsProviderServer struct {
	impl apifactory.AnalyticsProvider
}

// TODO: AnalyticsProviderServer deve implementar apifactory.AnalyticsProvider
