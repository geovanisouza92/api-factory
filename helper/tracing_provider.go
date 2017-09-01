package helper

import (
	"github.com/geovanisouza92/api-factory/apifactory"
)

type TracingProvider struct {
	Schema        map[string]*Schema
	ConfigureFunc ConfigureFunc
	SendFunc      SendFunc

	meta interface{}
}

type SendFunc func(apifactory.TracingEvent, interface{})

func (p *TracingProvider) Configure(c *apifactory.ProviderConfig) error {
	if p.ConfigureFunc == nil {
		return nil
	}

	meta, err := p.ConfigureFunc(c)
	if err != nil {
		return err
	}

	p.meta = meta
	return nil
}

func (p *TracingProvider) Send(ev apifactory.TracingEvent) {
	p.SendFunc(ev, nil)
}
