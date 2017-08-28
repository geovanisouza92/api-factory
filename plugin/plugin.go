package plugin

import (
	"errors"

	"github.com/hashicorp/go-plugin"
)

var PluginMap = map[string]plugin.Plugin{
	"analytics":       &AnalyticsProviderPlugin{},
	"authentication":  &AuthenticationProviderPlugin{},
	"authorization":   &AuthorizationProviderPlugin{},
	"logging":         &LoggingProviderPlugin{},
	"scripting":       &ScriptingProviderPlugin{},
	"service":         &ServiceProviderPlugin{},
	"traffic_control": &TrafficControlProviderPlugin{},
}

var errNotImplemented = errors.New("Not implemented")
