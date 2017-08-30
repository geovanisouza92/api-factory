package query

//go:generate stringer -type=Operator

type Operator byte

const (
	Invalid Operator = iota
	Exact
	IExact
	Contains
	IContains
	In
	Gt
	Gte
	Lt
	Lte
	StartsWith
	IStartsWith
	EndsWith
	IEndsWith
	Range

	// Date
	// Year
	// Month
	// Day
	// WeekDay
	// Hour
	// Minute
	// Second

	IsNull
	Regex
	IRegex
)

var OperatorToString = map[Operator]string{
	Exact:       "exact",
	IExact:      "iexact",
	Contains:    "contains",
	IContains:   "icontains",
	In:          "in",
	Gt:          "gt",
	Gte:         "gte",
	Lt:          "lt",
	Lte:         "lte",
	StartsWith:  "startswith",
	IStartsWith: "istartswith",
	EndsWith:    "endswith",
	IEndsWith:   "iendswith",
	Range:       "range",

	// Date
	// Year
	// Month
	// Day
	// WeekDay
	// Hour
	// Minute
	// Second

	IsNull: "isnull",
	Regex:  "regex",
	IRegex: "iregex",
}

var StringToOperator = map[string]Operator{
	"exact":       Exact,
	"iexact":      IExact,
	"contains":    Contains,
	"icontains":   IContains,
	"in":          In,
	"gt":          Gt,
	"gte":         Gte,
	"lt":          Lt,
	"lte":         Lte,
	"startswith":  StartsWith,
	"istartswith": IStartsWith,
	"endswith":    EndsWith,
	"iendswith":   IEndsWith,
	"range":       Range,

	// Date
	// Year
	// Month
	// Day
	// WeekDay
	// Hour
	// Minute
	// Second

	"isnull": IsNull,
	"regex":  Regex,
	"iregex": IRegex,
}
