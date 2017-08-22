package main

import (
	"github.com/geovanisouza92/api-factory/plugin"
)

type DatabaseService struct{}

func (s *DatabaseService) List() ([]*plugin.Resource, error) {
	return []*plugin.Resource{}, nil
}

func main() {
	plugin.Serve(map[string]interface{}{
		"database": &plugin.DatabaseServicePlugin{Impl: &DatabaseService{}},
	})
}
