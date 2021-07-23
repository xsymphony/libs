package slices_test

import (
	"testing"

	"github.com/xsymphony/libs/slices"
)

func TestMapTemplate(t *testing.T) {
	tests := []struct {
		name string
		fn   func(slices.Template) slices.Template
		in   []slices.Template
		out  []slices.Template
	}{
		{
			"double",
			func(t slices.Template) slices.Template { return t.(int) + t.(int) },
			[]slices.Template{1, 2, 3},
			[]slices.Template{2, 4, 6},
		},
		{
			"square",
			func(t slices.Template) slices.Template { return t.(int) * t.(int) },
			[]slices.Template{1, 2, 3},
			[]slices.Template{1, 4, 9},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := slices.MapTemplate(tt.fn, tt.in)
			if !equal(out, tt.out) {
				t.Fatalf("MapTemplate(%#v) -> %#v, but %#v", tt.in, tt.out, out)
			}
		})
	}
}

func TestFilterTemplate(t *testing.T) {
	tests := []struct {
		name string
		fn   func(slices.Template) bool
		in   []slices.Template
		out  []slices.Template
	}{
		{
			"odd",
			func(t slices.Template) bool { return t.(int)%2 == 1 },
			[]slices.Template{1, 2, 3, 4, 5},
			[]slices.Template{1, 3, 5},
		},
		{
			"more than 2",
			func(t slices.Template) bool { return len(t.(string)) >= 2 },
			[]slices.Template{"", "a", "aa", "bb", "c", "dddd"},
			[]slices.Template{"aa", "bb", "dddd"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := slices.FilterTemplate(tt.fn, tt.in)
			if !equal(out, tt.out) {
				t.Fatalf("FilterTemplate(%#v) -> %#v, but %#v", tt.in, tt.out, out)
			}
		})
	}
}

func TestReduceTemplate(t *testing.T) {
	tests := []struct {
		name string
		fn   func(slices.Template, slices.Template) slices.Template
		in   []slices.Template
		out  slices.Template
	}{
		{
			"sum",
			func(t1, t2 slices.Template) slices.Template { return t1.(int) + t2.(int) },
			[]slices.Template{1, 2, 3, 4, 5},
			15,
		},
		{
			"join",
			func(t1, t2 slices.Template) slices.Template { return t1.(string) + "," + t2.(string) },
			[]slices.Template{"how", "old", "are", "you"},
			"how,old,are,you",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := slices.ReduceTemplate(tt.fn, tt.in)
			if out != tt.out {
				t.Fatalf("ReduceTemplate(%#v) -> %v, but %v", tt.in, tt.out, out)
			}
		})
	}
}

func equal(a, b []slices.Template) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
