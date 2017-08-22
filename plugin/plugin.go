package plugin

import (
	"net/rpc"

	goPlugin "github.com/hashicorp/go-plugin"
	"google.golang.org/grpc"

	"github.com/geovanisouza92/api-factory/proto"
)

var (
	HandshakeConfig = goPlugin.HandshakeConfig{
		ProtocolVersion:  1,
		MagicCookieKey:   "DATABASE_PLUGIN",
		MagicCookieValue: "database",
	}

	PluginMap = map[string]goPlugin.Plugin{
		"database": &DatabaseServicePlugin{},
	}
)

type DatabaseServicePlugin struct {
	Impl DatabaseService
}

func (p *DatabaseServicePlugin) Client(m *goPlugin.MuxBroker, c *rpc.Client) (interface{}, error) {
	return &rpcClient{client: c}, nil
}

func (p *DatabaseServicePlugin) Server(m *goPlugin.MuxBroker) (interface{}, error) {
	return &rpcServer{Impl: p.Impl}, nil
}

func (p *DatabaseServicePlugin) GRPCClient(c *grpc.ClientConn) (interface{}, error) {
	return &grpcClient{client: proto.NewDatabaseServiceClient(c)}, nil
}

func (p *DatabaseServicePlugin) GRPCServer(s *grpc.Server) error {
	proto.RegisterDatabaseServiceServer(s, &grpcServer{Impl: p.Impl})
	return nil
}

func Serve(plugins map[string]interface{}) {
	ps := map[string]goPlugin.Plugin{}
	for k, p := range plugins {
		ps[k] = p.(goPlugin.Plugin)
	}

	goPlugin.Serve(&goPlugin.ServeConfig{
		HandshakeConfig: HandshakeConfig,
		Plugins:         ps,
		GRPCServer:      goPlugin.DefaultGRPCServer,
	})
}
