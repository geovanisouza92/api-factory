package query

type QuerySet struct{}

func (q *QuerySet) Iter() Iterator {
	return nil
}

func (q *QuerySet) Slice(low, high int) *QuerySet {
	return q
}

func (q *QuerySet) String() string {
	return ""
}

func (q *QuerySet) Len() int {
	return 0
}

func (q *QuerySet) Exists() bool {
	return false
}

func (q *QuerySet) Filter(c FieldCriteria) *QuerySet {
	return q
}

func (q *QuerySet) Exclude(c FieldCriteria) *QuerySet {
	return q
}

func (q *QuerySet) OrderBy(fls ...FieldLookup) *QuerySet {
	return q
}

func (q *QuerySet) Reverse() *QuerySet {
	return q
}

func (q *QuerySet) Distinct(fls ...FieldLookup) *QuerySet {
	return q
}

func (q *QuerySet) Values(fls ...FieldLookup) map[string]interface{} {
	return map[string]interface{}{}
}

func (q *QuerySet) All() {}
