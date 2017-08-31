package helper

import (
	"github.com/geovanisouza92/api-factory/apifactory"
)

type TracingProvider struct {
	Schema        map[string]*Schema
	ConfigureFunc ConfigureFunc
	SendFunc      SendFunc
}

type SendFunc func(apifactory.TracingEvent, interface{})

func (p *TracingProvider) Send(ev apifactory.TracingEvent) {
	p.SendFunc(ev)
}
