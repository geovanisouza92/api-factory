package main

import (
	"github.com/geovanisouza92/api-factory/database"
	"github.com/geovanisouza92/api-factory/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		DatabaseFunc: func() database.Service {
			return &PostgreSQLDatabaseService{}
		},
	})
}

type PostgreSQLDatabaseService struct{}

func (s *PostgreSQLDatabaseService) List() ([]*database.Resource, error) {
	return []*database.Resource{}, nil
}
