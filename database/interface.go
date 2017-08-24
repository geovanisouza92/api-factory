package database

type Service interface {
	List() ([]*Resource, error)
}

type Resource struct {
	Name string
	Kind Kind
}

type Kind byte

const (
	UnknownKind Kind = iota
	Table
	View
	StoredProc
)
