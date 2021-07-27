package template

import "testing"

func TestIfTemplate(t *testing.T) {
	tests := []struct {
		name                  string
		cond                  func() bool
		trueValue, falseValue Template
		out                   Template
	}{
		{
			"(1) must true",
			func() bool { return 2 > 0 },
			"true",
			"false",
			"true",
		},
		{
			"(2) must false",
			func() bool { return len("aa") > len("aaa") },
			"true",
			"false",
			"false",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := IfTemplate(tt.cond)(tt.trueValue)(tt.falseValue)
			if out != tt.out {
				t.Fatalf("IfTemplate(%v)(%v)(%v) -> %v, but %v", tt.cond(), tt.trueValue, tt.falseValue, tt.out, out)
			}
		})
	}
}
