package main

import (
	"github.com/geovanisouza92/api-factory/builtins/providers/postgresql"
	"github.com/geovanisouza92/api-factory/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ServiceProviderFunc: postgresql.Provider,
	})
}
