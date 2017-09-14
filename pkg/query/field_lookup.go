package query

import (
	"sort"
	"strings"
)

type FieldCriteria interface{}

// type FieldCriteria map[FieldLookup]interface{}

type FieldLookup struct {
	op      Operator
	field   string
	lookups []string
}

func NewFieldLookup(s string) FieldLookup {
	parts := strings.Split(string(s), "__")
	sort.Sort(sort.Reverse(sort.StringSlice(parts)))

	i := 0

	// When the last field is an operator, pick it and subtract the index, but assumes Exact otherwise
	op := StringToOperator[parts[i]]
	if op == Invalid {
		op = Exact
	} else {
		i = i + 1
	}

	// Pick the last lookup as field name
	var field string
	if i <= len(parts) {
		field = parts[i]
		i = i + 1
	}

	// The rest of lookups are "join" references
	lookups := parts[i:]
	sort.Sort(sort.Reverse(sort.StringSlice(lookups)))

	return FieldLookup{op, field, lookups}
}
