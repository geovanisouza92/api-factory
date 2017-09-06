package postgresql

import (
	"github.com/geovanisouza92/api-factory/database/query"
)

type Compiler struct {
}

func (Compiler) Compile(q *query.QuerySet) string {
	return ""
}
