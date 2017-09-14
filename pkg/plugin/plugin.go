package plugin

import (
	"errors"

	"github.com/hashicorp/go-plugin"
)

var PluginMap = map[string]plugin.Plugin{
	"authentication":  &AuthenticationProviderPlugin{},
	"authorization":   &AuthorizationProviderPlugin{},
	"scripting":       &ScriptingProviderPlugin{},
	"service":         &ServiceProviderPlugin{},
	"tracing":         &TracingProviderPlugin{},
	"traffic_control": &TrafficControlProviderPlugin{},
}

var errNotImplemented = errors.New(`
This is not implemented.

Plugins on API Factory are meant to communicate over GRPC instead of plain
net/rpc, allowing a broader technology options when someone is trying to
implement new plugins.
`)
