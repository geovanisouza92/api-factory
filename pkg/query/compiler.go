package query

type Compiler interface {
	Compile(QuerySet) string
}
