package postgresql

import (
	"github.com/geovanisouza92/api-factory/helper"
)

func resourceTable() *helper.Resource {
	return &helper.Resource{
		Schema: map[string]*helper.Schema{
			"namespace": &helper.Schema{
				Type:        helper.TypeString,
				Required:    true,
				Description: "",
			},
			"table_name": &helper.Schema{
				Type:        helper.TypeString,
				Required:    true,
				Description: "",
			},
			"id": &helper.Schema{
				Type:        helper.TypeString,
				Description: "",
			},
		},
	}
}
