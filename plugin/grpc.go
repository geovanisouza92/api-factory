package plugin

import (
	"golang.org/x/net/context"

	"github.com/geovanisouza92/api-factory/proto"
)

type grpcClient struct {
	client proto.DatabaseServiceClient
}

func (c *grpcClient) List() ([]*Resource, error) {
	a := []*Resource{}

	res, err := c.client.List(context.Background(), &proto.ListRequest{})
	if err != nil {
		return a, err
	}

	for _, resource := range res.GetResources() {
		r := &Resource{Name: resource.GetName()}
		if resource.GetKind() == proto.Resource_TABLE {
			r.Kind = Table
		} else if resource.GetKind() == proto.Resource_VIEW {
			r.Kind = View
		} else if resource.GetKind() == proto.Resource_STORED_PROC {
			r.Kind = StoredProc
		}
		a = append(a, r)
	}

	return a, nil
}

type grpcServer struct {
	Impl DatabaseService
}

func (s *grpcServer) List(ctx context.Context, req *proto.ListRequest) (*proto.ListResponse, error) {
	rs, err := s.Impl.List()
	if err != nil {
		return &proto.ListResponse{}, err
	}

	res := []*proto.Resource{}
	for _, resource := range rs {
		r := &proto.Resource{Name: resource.Name}
		if resource.Kind == Table {
			r.Kind = proto.Resource_TABLE
		} else if resource.Kind == View {
			r.Kind = proto.Resource_VIEW
		} else if resource.Kind == StoredProc {
			r.Kind = proto.Resource_STORED_PROC
		}
	}

	return &proto.ListResponse{Resources: res}, nil
}
