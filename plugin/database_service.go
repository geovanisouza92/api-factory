package plugin

import (
	"errors"
	"net/rpc"

	"github.com/hashicorp/go-plugin"
	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"github.com/geovanisouza92/api-factory/database"
	"github.com/geovanisouza92/api-factory/proto"
)

var errNotImplemented = errors.New("Not implemented")

type DatabaseServicePlugin struct {
	F DatabaseFunc
}

func (*DatabaseServicePlugin) Client(*plugin.MuxBroker, *rpc.Client) (interface{}, error) {
	return nil, errNotImplemented
}

func (*DatabaseServicePlugin) Server(*plugin.MuxBroker) (interface{}, error) {
	return nil, errNotImplemented
}

func (*DatabaseServicePlugin) GRPCClient(c *grpc.ClientConn) (interface{}, error) {
	return &databaseService{client: proto.NewDatabaseClient(c)}, nil
}

func (self *DatabaseServicePlugin) GRPCServer(s *grpc.Server) error {
	proto.RegisterDatabaseServer(s, &databaseServiceServer{impl: self.F()})
	return nil
}

type databaseService struct {
	client proto.DatabaseClient
}

func (self *databaseService) List() ([]*database.Resource, error) {
	a := []*database.Resource{}

	res, err := self.client.List(context.Background(), &proto.ListRequest{})
	if err != nil {
		return a, err
	}

	for _, resource := range res.GetResources() {
		r := &database.Resource{Name: resource.GetName()}
		if resource.GetKind() == proto.Resource_TABLE {
			r.Kind = database.Table
		} else if resource.GetKind() == proto.Resource_VIEW {
			r.Kind = database.View
		} else if resource.GetKind() == proto.Resource_STORED_PROC {
			r.Kind = database.StoredProc
		}
		a = append(a, r)
	}

	return a, nil
}

type databaseServiceServer struct {
	impl database.Service
}

func (self *databaseServiceServer) List(ctx context.Context, req *proto.ListRequest) (*proto.ListResponse, error) {
	rs, err := self.impl.List()
	if err != nil {
		return &proto.ListResponse{}, err
	}

	res := []*proto.Resource{}
	for _, resource := range rs {
		r := &proto.Resource{Name: resource.Name}
		if resource.Kind == database.Table {
			r.Kind = proto.Resource_TABLE
		} else if resource.Kind == database.View {
			r.Kind = proto.Resource_VIEW
		} else if resource.Kind == database.StoredProc {
			r.Kind = proto.Resource_STORED_PROC
		}
	}

	return &proto.ListResponse{Resources: res}, nil
}
