package helper

import (
	"github.com/geovanisouza92/api-factory/apifactory"
)

type ConfigureFunc func(*apifactory.ProviderConfig) (interface{}, error)
