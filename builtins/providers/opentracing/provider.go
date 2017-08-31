package opentracing

import (
	"github.com/geovanisouza92/api-factory/apifactory"
	"github.com/geovanisouza92/api-factory/helper"
)

func Provider() apifactory.TracingProvider {
	return &helper.TracingProvider{
		Schema:        map[string]*helper.Schema{},
		ConfigureFunc: configureTracing,
		SendFunc:      providerSend,
	}
}

func configureTracing(d *helper.ResourceData) (interface{}, error) {
	// TODO: Return `c` as an opentracing client/instance
	return nil, nil
}

func providerSend(ev apifactory.TracingEvent, c interface{}) {
	// TODO: Use `c` to send `ev`
}
