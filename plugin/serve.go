package plugin

import (
	"github.com/geovanisouza92/api-factory/database"
	plugin "github.com/hashicorp/go-plugin"
)

var HandshakeConfig = plugin.HandshakeConfig{
	ProtocolVersion:  1,
	MagicCookieKey:   "API_FACTORY_PLUGIN",
	MagicCookieValue: "Pnt9RTftm3K7tJnjsLqL4whVxwmPC7n9wzrpNPCdxbsjNdqX9TXNt7Tc7wkHtWJv",
}

func Serve(opts *ServeOpts) {
	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: HandshakeConfig,
		Plugins:         makePluginMapFrom(opts),
	})
}

type ServeOpts struct {
	DatabaseFunc DatabaseFunc
}

type DatabaseFunc func() database.Service

func makePluginMapFrom(opts *ServeOpts) map[string]plugin.Plugin {
	return map[string]plugin.Plugin{
		"database": &DatabaseServicePlugin{F: opts.DatabaseFunc},
	}
}
