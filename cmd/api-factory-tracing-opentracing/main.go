package main

import (
	"github.com/geovanisouza92/api-factory/builtins/providers/opentracing"
	"github.com/geovanisouza92/api-factory/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		TracingProviderFunc: opentracing.Provider,
	})
}
