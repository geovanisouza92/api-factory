package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"

	goPlugin "github.com/hashicorp/go-plugin"

	"github.com/geovanisouza92/api-factory/plugin"
)

func main() {
	pluginProcess := goPlugin.NewClient(&goPlugin.ClientConfig{
		HandshakeConfig:  plugin.HandshakeConfig,
		Plugins:          plugin.PluginMap,
		Cmd:              exec.Command("sh", "-c", os.Getenv("DATABASE_PLUGIN")),
		AllowedProtocols: []goPlugin.Protocol{goPlugin.ProtocolGRPC},
	})
	defer pluginProcess.Kill()

	client, err := pluginProcess.Client()
	if err != nil {
		fmt.Println("Error:", err.Error())
		os.Exit(1)
	}

	rawService, err := client.Dispense("database")
	if err != nil {
		fmt.Println("Error:", err.Error())
		os.Exit(1)
	}

	databaseService := rawService.(plugin.DatabaseService)
	res, err := databaseService.List()
	if err != nil {
		fmt.Println("Error:", err.Error())
		os.Exit(1)
	}

	_, err = json.Marshal(res)
	if err != nil {
		fmt.Println("Error:", err.Error())
		os.Exit(1)
	}
}
