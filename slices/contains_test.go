package slices_test

import (
	"testing"

	"github.com/xsymphony/libs/slices"
)

func TestContainsString(t *testing.T) {
	tests := []struct {
		name string
		arg1 []string
		arg2 string
		out  bool
	}{
		{
			"must in",
			[]string{"a", "aaa", "aaaa"},
			"aaa",
			true,
		},
		{
			"must out",
			[]string{"a", "aaa", "aaaa"},
			"aa",
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := slices.ContainsString(tt.arg1, tt.arg2)
			if out != tt.out {
				t.Fatalf("ContainsString(%#v, %s) -> %v, but %v", tt.arg1, tt.arg2, tt.out, out)
			}
		})
	}
}

func TestContainsAnyString(t *testing.T) {
	tests := []struct {
		name string
		arg1 []string
		arg2 []string
		out  bool
	}{
		{
			"1.must in",
			[]string{"a", "aaa", "aaaa"},
			[]string{"aaa"},
			true,
		},
		{
			"2.must in",
			[]string{"a", "aaa", "aaaa"},
			[]string{"b", "aaa"},
			true,
		},
		{
			"3.must in",
			[]string{"a", "aaa", "aaaa"},
			[]string{"a", "aaa", "bbb"},
			true,
		},
		{
			"4.must out",
			[]string{"a", "aaa", "aaaa"},
			[]string{"aa"},
			false,
		},
		{
			"5.must out",
			[]string{"a", "aaa", "aaaa"},
			[]string{"aa", "bb"},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := slices.ContainsAnyString(tt.arg1, tt.arg2...)
			if out != tt.out {
				t.Fatalf("ContainsAnyString(%#v, %#v) -> %v, but %v", tt.arg1, tt.arg2, tt.out, out)
			}
		})
	}
}
