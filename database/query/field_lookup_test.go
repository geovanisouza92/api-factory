package query

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestFieldLookup(t *testing.T) {
	Convey("Field lookups translate as expected", t, func() {
		Convey("example 1", func() {
			expected := FieldLookup{
				op:      Exact,
				field:   "id",
				lookups: []string{},
			}
			actual := NewFieldLookup("id")
			So(actual, ShouldResemble, expected)
		})

		Convey("example 2", func() {
			expected := FieldLookup{
				op:      Lt,
				field:   "id",
				lookups: []string{},
			}
			actual := NewFieldLookup("id__lt")
			So(actual, ShouldResemble, expected)
		})

		Convey("example 3", func() {
			expected := FieldLookup{
				op:      StartsWith,
				field:   "name",
				lookups: []string{"author"},
			}
			actual := NewFieldLookup("author__name__startswith")
			So(actual, ShouldResemble, expected)
		})

		Convey("example 4", func() {
			expected := FieldLookup{
				op:      Exact,
				field:   "name",
				lookups: []string{"blog", "author"},
			}
			actual := NewFieldLookup("blog__author__name")
			So(actual, ShouldResemble, expected)
		})
	})
}
