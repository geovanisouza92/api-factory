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

	analyticsProviderRaw, err := protocolClient.Dispense("analytics")
	if err != nil {
		fmt.Println("Error:", err.Error())
		os.Exit(1)
	}

	_ = analyticsProviderRaw.(apifactory.AnalyticsProvider)
}
