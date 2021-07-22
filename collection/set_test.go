package collection_test

import (
	"testing"

	"github.com/xsymphony/libs/collection"
)

func TestSet(t *testing.T) {
	c := collection.NewSetTemplate(1, 2, 3)

	s := c.Copy()
	s.Add(4)
	s.Update(3, 5, 6)
	s.Remove(2)

	if s.Length() != 5 {
		t.Fatal("length not equal 5")
	}

	for _, i := range []int{1, 3, 4, 5, 6} {
		if !s.Has(i) {
			t.Fatalf("set not has %v", i)
		}
	}

	if len(s.TemplateList()) != s.Length() {
		t.Fatal("list not equal set")
	}

	another := collection.NewSetTemplate(2, 3, 4)
	union := c.Union(another)
	if union.Length() != 4 {
		t.Fatal("length not equal 6")
	}

	intersect := c.Intersect(another)
	if intersect.Length() != 2 {
		t.Fatal("length not equal 2")
	}

	except := c.Except(another)
	if except.Length() != 1 {
		t.Fatal("length not equal 1")
	}
}
