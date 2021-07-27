package template

import "github.com/cheekybits/genny/generic"

type Template generic.Number

//go:generate genny -in=$GOFILE -out=../$GOFILE -pkg=compare gen "Template=int,int64,int32,float64,float32"
func MaxTemplate(arg Template, others ...Template) Template {
	max := arg
	for _, other := range others {
		if other > max {
			max = other
		}
	}
	return max
}
