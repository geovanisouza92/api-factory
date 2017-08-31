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
		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *helper.ResourceData) (interface{}, error) {
	return nil, nil
}