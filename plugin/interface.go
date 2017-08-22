package plugin

type Kind byte

const (
	UnknownKind Kind = iota
	Table
	View
	StoredProc
)

type Resource struct {
	Name string
	Kind Kind
}

type DatabaseService interface {
	List() ([]*Resource, error)
}
