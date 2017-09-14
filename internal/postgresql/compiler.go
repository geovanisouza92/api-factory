package postgresql

import (
	"github.com/geovanisouza92/api-factory/pkg/query"
)

type Compiler struct {
}

func (Compiler) Compile(q *query.QuerySet) string {
	return ""
}
