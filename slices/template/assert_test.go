package template

import "testing"

func TestAnyTemplate(t *testing.T) {
	tests := []struct {
		name   string
		in     []Template
		assert func(Template) bool
		out    bool
	}{
		{
			"some one is odd",
			[]Template{2, 4, 5, 6, 10},
			func(t Template) bool { return t.(int)%2 == 1 },
			true,
		},
		{
			"some one is more than 20",
			[]Template{2, 4, 5, 6, 10},
			func(t Template) bool { return t.(int) >= 20 },
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := AnyTemplate(tt.in, tt.assert)
			if out != tt.out {
				t.Fatalf("AnyTemplate(%#v) -> %v, but %v", tt.in, tt.out, out)
			}
		})
	}
}

func TestAllTemplate(t *testing.T) {
	tests := []struct {
		name   string
		in     []Template
		assert func(Template) bool
		out    bool
	}{
		{
			"all is odd",
			[]Template{1, 3, 5, 7, 9},
			func(t Template) bool { return t.(int)%2 == 1 },
			true,
		},
		{
			"all more than 20",
			[]Template{21, 22, 23, 24, 10},
			func(t Template) bool { return t.(int) >= 20 },
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := AllTemplate(tt.in, tt.assert)
			if out != tt.out {
				t.Fatalf("AllTemplate(%#v) -> %v, but %v", tt.in, tt.out, out)
			}
		})
	}
}
