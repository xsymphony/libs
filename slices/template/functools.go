package template

import "github.com/cheekybits/genny/generic"

//go:generate genny -in=$GOFILE -out=../$GOFILE -pkg=slices gen "Template=BUILTINS"
type Template generic.Type

func MapTemplate(f func(Template) Template, s []Template) []Template {
	n := make([]Template, 0, len(s))
	for i := 0; i < len(s); i++ {
		n = append(n, f(s[i]))
	}
	return n
}

func FilterTemplate(f func(Template) bool, s []Template) []Template {
	n := make([]Template, 0, len(s))
	for i := 0; i < len(s); i++ {
		if f(s[i]) {
			n = append(n, s[i])
		}
	}
	return n
}

func ReduceTemplate(f func(Template, Template) Template, s []Template) Template {
	t := s[0]
	for i := 1; i < len(s); i++ {
		t = f(t, s[i])
	}
	return t
}
