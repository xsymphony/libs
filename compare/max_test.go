package compare_test

import (
	"testing"

	"github.com/xsymphony/libs/compare"
)

func TestMaxTemplate(t *testing.T) {
	tests := []struct {
		name string
		in   []compare.Template
		out  compare.Template
	}{
		{
			"int",
			[]compare.Template{1, 2, 3},
			3,
		},
		{
			"float",
			[]compare.Template{1.1, 2.1, 3.1},
			3.1,
		},
		{
			"minus",
			[]compare.Template{-1, -2, -3},
			-1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := compare.MaxTemplate(tt.in[0], tt.in[1:]...)
			if out != tt.out {
				t.Logf("MaxTemplate(%f, %#v) -> %f, but %f", tt.in[0], tt.in[1:], tt.out, out)
			}
		})
	}
}
