package plugin

import (
	"github.com/hashicorp/go-plugin"

	"github.com/geovanisouza92/api-factory/pkg/apifactory"
)

var HandshakeConfig = plugin.HandshakeConfig{
	ProtocolVersion:  1,
	MagicCookieKey:   "API_FACTORY_PLUGIN",
	MagicCookieValue: "Pnt9RTftm3K7tJnjsLqL4whVxwmPC7n9wzrpNPCdxbsjNdqX9TXNt7Tc7wkHtWJv",
}

type ServeOpts struct {
	TracingProviderFunc        TracingProviderFunc
	AuthenticationProviderFunc AuthenticationProviderFunc
	AuthorizationProviderFunc  AuthorizationProviderFunc
	ScriptingProviderFunc      ScriptingProviderFunc
	ServiceProviderFunc        ServiceProviderFunc
	TrafficControlProviderFunc TrafficControlProviderFunc
}

type TracingProviderFunc func() apifactory.TracingProvider
type AuthenticationProviderFunc func() apifactory.AuthenticationProvider
type AuthorizationProviderFunc func() apifactory.AuthorizationProvider
type ScriptingProviderFunc func() apifactory.ScriptingProvider
type ServiceProviderFunc func() apifactory.ServiceProvider
type TrafficControlProviderFunc func() apifactory.TrafficControlProvider

func Serve(opts *ServeOpts) {
	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: HandshakeConfig,
		// TODO: TLS
		Plugins: makePluginMapFrom(opts),
	})
}

func makePluginMapFrom(opts *ServeOpts) map[string]plugin.Plugin {
	return map[string]plugin.Plugin{
		"authentication":  &AuthenticationProviderPlugin{F: opts.AuthenticationProviderFunc},
		"authorization":   &AuthorizationProviderPlugin{F: opts.AuthorizationProviderFunc},
		"scripting":       &ScriptingProviderPlugin{F: opts.ScriptingProviderFunc},
		"service":         &ServiceProviderPlugin{F: opts.ServiceProviderFunc},
		"tracing":         &TracingProviderPlugin{F: opts.TracingProviderFunc},
		"traffic_control": &TrafficControlProviderPlugin{F: opts.TrafficControlProviderFunc},
	}
}
