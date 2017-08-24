package main

import (
	"encoding/json"
	"fmt"
	"os"

	goPlugin "github.com/hashicorp/go-plugin"

	"github.com/geovanisouza92/api-factory/database"
	"github.com/geovanisouza92/api-factory/plugin"
)

func main() {
	pluginClient := goPlugin.NewClient(plugin.ClientConfig("foo bar"))

	protocolClient, err := pluginClient.Client()
	if err != nil {
		fmt.Println("Error:", err.Error())
		os.Exit(1)
	}

	databaseServicePlugin, err := protocolClient.Dispense("database")
	if err != nil {
		fmt.Println("Error:", err.Error())
		os.Exit(1)
	}

	databaseService := databaseServicePlugin.(database.Service)
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
