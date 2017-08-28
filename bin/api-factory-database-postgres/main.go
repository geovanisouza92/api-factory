package main

import (
	"github.com/geovanisouza92/api-factory/apifactory"
	"github.com/geovanisouza92/api-factory/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		AnalyticsProviderFunc: func() apifactory.AnalyticsProvider {
			return nil
		},
	})
}
