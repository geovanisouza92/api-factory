package plugin

import (
	"os/exec"

	"github.com/hashicorp/go-plugin"
)

func ClientConfig(cmd string) *plugin.ClientConfig {
	return &plugin.ClientConfig{
		Cmd:              exec.Command(cmd),
		HandshakeConfig:  HandshakeConfig,
		Managed:          true,
		Plugins:          PluginMap,
		AllowedProtocols: []plugin.Protocol{plugin.ProtocolGRPC},
	}
}
