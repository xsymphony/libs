package template

import "testing"

func TestInsertTemplate(t *testing.T) {
	tests := []struct {
		name     string
		in       []Template
		index    int
		elements []Template
		out      []Template
	}{
		{
			"nil `in`",
			nil,
			-1,
			[]Template{1, 2, 3},
			[]Template{1, 2, 3},
		},
		{
			"nil `elements`",
			[]Template{1, 2, 3},
			-1,
			nil,
			[]Template{1, 2, 3},
		},
		{
			"`index` out of range",
			[]Template{1, 2, 3},
			100,
			[]Template{4, 5},
			[]Template{1, 2, 3, 4, 5},
		},
		{
			"`index` out of range",
			[]Template{1, 2, 3},
			-100,
			[]Template{4, 5},
			[]Template{4, 5, 1, 2, 3},
		},
		{
			"`index` out of range",
			[]Template{1, 2, 3},
			100,
			[]Template{4, 5},
			[]Template{1, 2, 3, 4, 5},
		},
		{
			"`index` head",
			[]Template{1, 2, 3},
			0,
			[]Template{4, 5},
			[]Template{4, 5, 1, 2, 3},
		},
		{
			"`index` tail",
			[]Template{1, 2, 3},
			-1,
			[]Template{4, 5},
			[]Template{1, 2, 4, 5, 3},
		},
		{
			"`index` mid",
			[]Template{1, 2, 3},
			1,
			[]Template{4, 5},
			[]Template{1, 4, 5, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := InsertTemplate(tt.in, tt.index, tt.elements...)
			if !equal(out, tt.out) {
				t.Fatalf("InsertTemplate(%#v, %d, %#v) -> %#v, but %#v", tt.in, tt.index, tt.elements, tt.out, out)
			}
		})
	}
}
