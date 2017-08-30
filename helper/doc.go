package helper

type Doc struct {
	Type    DocType
	Content string
}

type DocType byte

const (
	DocInvalid DocType = iota
	DocBlueprint
	DocSwagger
)
