package template

import "testing"

func TestMapTemplate(t *testing.T) {
	tests := []struct {
		name string
		fn   func(Template) Template
		in   []Template
		out  []Template
	}{
		{
			"double",
			func(t Template) Template { return t.(int) + t.(int) },
			[]Template{1, 2, 3},
			[]Template{2, 4, 6},
		},
		{
			"square",
			func(t Template) Template { return t.(int) * t.(int) },
			[]Template{1, 2, 3},
			[]Template{1, 4, 9},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := MapTemplate(tt.fn, tt.in)
			if !equal(out, tt.out) {
				t.Fatalf("MapTemplate(%#v) -> %#v, but %#v", tt.in, tt.out, out)
			}
		})
	}
}

func TestFilterTemplate(t *testing.T) {
	tests := []struct {
		name string
		fn   func(Template) bool
		in   []Template
		out  []Template
	}{
		{
			"odd",
			func(t Template) bool { return t.(int)%2 == 1 },
			[]Template{1, 2, 3, 4, 5},
			[]Template{1, 3, 5},
		},
		{
			"more than 2",
			func(t Template) bool { return len(t.(string)) >= 2 },
			[]Template{"", "a", "aa", "bb", "c", "dddd"},
			[]Template{"aa", "bb", "dddd"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := FilterTemplate(tt.fn, tt.in)
			if !equal(out, tt.out) {
				t.Fatalf("FilterTemplate(%#v) -> %#v, but %#v", tt.in, tt.out, out)
			}
		})
	}
}

func TestReduceTemplate(t *testing.T) {
	tests := []struct {
		name string
		fn   func(Template, Template) Template
		in   []Template
		out  Template
	}{
		{
			"sum",
			func(t1, t2 Template) Template { return t1.(int) + t2.(int) },
			[]Template{1, 2, 3, 4, 5},
			15,
		},
		{
			"join",
			func(t1, t2 Template) Template { return t1.(string) + "," + t2.(string) },
			[]Template{"how", "old", "are", "you"},
			"how,old,are,you",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := ReduceTemplate(tt.fn, tt.in)
			if out != tt.out {
				t.Fatalf("ReduceTemplate(%#v) -> %v, but %v", tt.in, tt.out, out)
			}
		})
	}
}

func equal(a, b []Template) bool {
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
