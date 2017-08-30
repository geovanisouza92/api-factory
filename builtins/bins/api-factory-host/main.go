package main

import (
	"fmt"
	"os"

	goPlugin "github.com/hashicorp/go-plugin"

	"github.com/geovanisouza92/api-factory/apifactory"
	"github.com/geovanisouza92/api-factory/plugin"
)

func main() {
	pluginClient := goPlugin.NewClient(plugin.ClientConfig("foo bar"))

	protocolClient, err := pluginClient.Client()
	if err != nil {
		fmt.Println("Error:", err.Error())
		os.Exit(1)
	}

	serviceProviderRaw, err := protocolClient.Dispense("service")
	if err != nil {
		fmt.Println("Error:", err.Error())
		os.Exit(1)
	}

	_ = serviceProviderRaw.(apifactory.ServiceProvider)
}
