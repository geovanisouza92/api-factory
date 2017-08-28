package plugin

import (
	"github.com/hashicorp/go-plugin"

	"github.com/geovanisouza92/api-factory/apifactory"
)

var HandshakeConfig = plugin.HandshakeConfig{
	ProtocolVersion:  1,
	MagicCookieKey:   "API_FACTORY_PLUGIN",
	MagicCookieValue: "Pnt9RTftm3K7tJnjsLqL4whVxwmPC7n9wzrpNPCdxbsjNdqX9TXNt7Tc7wkHtWJv",
}

type ServeOpts struct {
	AnalyticsProviderFunc      AnalyticsProviderFunc
	AuthenticationProviderFunc AuthenticationProviderFunc
	AuthorizationProviderFunc  AuthorizationProviderFunc
	LoggingProviderFunc        LoggingProviderFunc
	ScriptingProviderFunc      ScriptingProviderFunc
	ServiceProviderFunc        ServiceProviderFunc
	TrafficControlProviderFunc TrafficControlProviderFunc
}

type AnalyticsProviderFunc func() apifactory.AnalyticsProvider
type AuthenticationProviderFunc func() apifactory.AuthenticationProvider
type AuthorizationProviderFunc func() apifactory.AuthorizationProvider
type LoggingProviderFunc func() apifactory.LoggingProvider
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
		"analytics":       &AnalyticsProviderPlugin{F: opts.AnalyticsProviderFunc},
		"authentication":  &AuthenticationProviderPlugin{F: opts.AuthenticationProviderFunc},
		"authorization":   &AuthorizationProviderPlugin{F: opts.AuthorizationProviderFunc},
		"logging":         &LoggingProviderPlugin{F: opts.LoggingProviderFunc},
		"scripting":       &ScriptingProviderPlugin{F: opts.ScriptingProviderFunc},
		"service":         &ServiceProviderPlugin{F: opts.ServiceProviderFunc},
		"traffic_control": &TrafficControlProviderPlugin{F: opts.TrafficControlProviderFunc},
	}
}
