package postgresql

import (
	"github.com/geovanisouza92/api-factory/apifactory"
	"github.com/geovanisouza92/api-factory/helper"
)

func Provider() apifactory.ServiceProvider {
	return &helper.ServiceProvider{
		Schema: map[string]*helper.Schema{
			"connection_string": &helper.Schema{
				Type:        helper.TypeString,
				Required:    true,
				Description: "",
			},
		},
		ResourceMap: map[string]*helper.Resource{
			"_table": resourceTable(),
			// "_view": resourceView(),
			// "_stored_proc": resourceStoredProc(),
		},
		ConfigureFunc: configureFunc,
	}
}

func configureFunc(c *apifactory.ProviderConfig) (interface{}, error) {
	return nil, nil
}
