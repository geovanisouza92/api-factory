package plugin

import (
	"net/rpc"

	"github.com/hashicorp/go-plugin"
	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"github.com/geovanisouza92/api-factory/apifactory"
	"github.com/geovanisouza92/api-factory/plugin/proto"
)

type TracingProviderPlugin struct {
	// NOTE: Necessário apenas dentro dos plugins, que devem fornecer uma implementação válida
	F TracingProviderFunc
}

func (*TracingProviderPlugin) Client(*plugin.MuxBroker, *rpc.Client) (interface{}, error) {
	return nil, errNotImplemented
}

func (*TracingProviderPlugin) Server(*plugin.MuxBroker) (interface{}, error) {
	return nil, errNotImplemented
}

func (p *TracingProviderPlugin) GRPCClient(c *grpc.ClientConn) (interface{}, error) {
	return &TracingProvider{client: proto.NewTracingProviderClient(c)}, nil
}

func (p *TracingProviderPlugin) GRPCServer(s *grpc.Server) error {
	proto.RegisterTracingProviderServer(s, &TracingProviderServer{impl: p.F()})
	return nil
}

type TracingProvider struct {
	client proto.TracingProviderClient
}

// TODO: TracingProvider deve implementar apifactory.TracingProvider

func (c *TracingProvider) Configure(conf *apifactory.ProviderConfig) error {
	config := &proto.ProviderConfig{}
	_, err := c.client.Configure(context.Background(), config)
	if err != nil {
		return err
	}
	return nil
}

func (c *TracingProvider) Send(ev apifactory.TracingEvent) {
	event := &proto.TracingEvent{
		RequestId:  ev.RequestID,
		Scheme:     ev.Scheme,
		Host:       ev.Host,
		Method:     ev.Method,
		RequestUri: ev.RequestURI,
		Proto:      ev.Proto,
		RemoteAddr: ev.RemoteAddr,
		Status:     int32(ev.Status),
		Bytes:      int32(ev.Bytes),
		Duration:   ev.Duration,
	}
	_, err := c.client.Send(context.Background(), event)
	if err != nil {
		panic(err)
	}
}

type TracingProviderServer struct {
	impl apifactory.TracingProvider
}

// TODO: TracingProviderServer deve implementar proto.TracingProviderServer

func (s *TracingProviderServer) Configure(ctx context.Context, c *proto.ProviderConfig) (*proto.Error, error) {
	config := &apifactory.ProviderConfig{}
	err := s.impl.Configure(config)
	if err != nil {
		return &proto.Error{}, err
	}
	return &proto.Error{}, nil
}

func (s *TracingProviderServer) Send(ctx context.Context, ev *proto.TracingEvent) (*proto.Empty, error) {
	event := apifactory.TracingEvent{
		RequestID:  ev.RequestId,
		Scheme:     ev.Scheme,
		Host:       ev.Host,
		Method:     ev.Method,
		RequestURI: ev.RequestUri,
		Proto:      ev.Proto,
		RemoteAddr: ev.RemoteAddr,
		Status:     int(ev.Status),
		Bytes:      int(ev.Bytes),
		Duration:   ev.Duration,
	}
	s.impl.Send(event)
	return &proto.Empty{}, nil
}
