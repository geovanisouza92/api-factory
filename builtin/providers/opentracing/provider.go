package opentracing

import (
	"github.com/geovanisouza92/api-factory/apifactory"
	"github.com/geovanisouza92/api-factory/helper"
)

func Provider() apifactory.TracingProvider {
	return &helper.TracingProvider{
		Schema:        map[string]*helper.Schema{},
		ConfigureFunc: configureFunc,
		SendFunc:      sendFunc,
	}
}

func configureFunc(c *apifactory.ProviderConfig) (interface{}, error) {
	// TODO: Return `meta` as an opentracing client/instance
	return nil, nil
}

func sendFunc(ev apifactory.TracingEvent, meta interface{}) {
	// TODO: Use `meta` to send `ev`
}
