package query

import (
	"testing"

	"github.com/smartystreets/assertions"
)

func TestFieldLookup1(t *testing.T) {
	expected := FieldLookup{
		op:      Exact,
		field:   "id",
		lookups: []string{},
	}
	actual := NewFieldLookup("id")

	if msg := assertions.ShouldResemble(actual, expected); msg != "" {
		t.Error(msg)
	}
}

func TestFieldLookup2(t *testing.T) {
	expected := FieldLookup{
		op:      Lt,
		field:   "id",
		lookups: []string{},
	}
	actual := NewFieldLookup("id__lt")

	if msg := assertions.ShouldResemble(actual, expected); msg != "" {
		t.Error(msg)
	}
}

func TestFieldLookup3(t *testing.T) {
	expected := FieldLookup{
		op:      StartsWith,
		field:   "name",
		lookups: []string{"author"},
	}
	actual := NewFieldLookup("author__name__startswith")

	if msg := assertions.ShouldResemble(actual, expected); msg != "" {
		t.Error(msg)
	}
}

func TestFieldLookup4(t *testing.T) {
	expected := FieldLookup{
		op:      Exact,
		field:   "name",
		lookups: []string{"blog", "author"},
	}
	actual := NewFieldLookup("blog__author__name")

	if msg := assertions.ShouldResemble(actual, expected); msg != "" {
		t.Error(msg)
	}
}
