package plugin

import (
	"os/exec"

	"github.com/hashicorp/go-plugin"
)

func ClientConfig(cmd string) *plugin.ClientConfig {
	return &plugin.ClientConfig{
		HandshakeConfig: HandshakeConfig,
		Plugins:         PluginMap,
		Cmd:             exec.Command(cmd),
		// TODO: ReattachConfig ou SecureConfig ?
		// TODO: TLS
		Managed:          true,
		AllowedProtocols: []plugin.Protocol{plugin.ProtocolGRPC},
		// TODO: Logger
	}
}
