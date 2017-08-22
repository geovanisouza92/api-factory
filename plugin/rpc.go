package plugin

import (
	"net/rpc"
)

type rpcClient struct {
	client *rpc.Client
}

func (c *rpcClient) List() ([]*Resource, error) {
	resp := []*Resource{}
	err := c.client.Call("/proto.DatabaseService/List", map[string]interface{}{}, &resp)
	return resp, err
}

type rpcServer struct {
	Impl DatabaseService
}

func (s *rpcServer) List(args map[string]interface{}, resp *interface{}) error {
	res, err := s.Impl.List()
	if err != nil {
		return err
	}
	*resp = res
	return nil
}
