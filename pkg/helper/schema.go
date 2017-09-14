package helper

type Schema struct {
	Type ValueType

	Required bool
	Unique   bool

	Default     interface{}
	DefaultFunc func() (interface{}, error)

	Description string

	// Type == TypeList
	Elem     interface{}
	MinItems int
	MaxItems int

	// Type == TypeMap
	Schema map[string]Schema

	ValidateFunc func(interface{}, string) ([]string, []error)

	Deprecated string
}
