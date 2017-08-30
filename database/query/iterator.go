package query

type Iterable interface {
	Iter() Iterator
}

type Iterator interface {
	Next() (interface{}, bool)
}
