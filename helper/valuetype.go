package helper

//go:generate stringer -type=ValueType

type ValueType int

const (
	TypeInvalid ValueType = iota
	TypeBool
	TypeInt
	TypeFloat
	TypeString
	TypeList
	TypeMap
	TypeURL
	// TypeTimestamp ??
)
