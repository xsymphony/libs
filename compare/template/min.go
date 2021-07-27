package template

//go:generate genny -in=$GOFILE -out=../$GOFILE -pkg=compare gen "Template=int,int64,int32,float64,float32"
func MinTemplate(arg Template, others ...Template) Template {
	min := arg
	for _, other := range others {
		if other < min {
			min = other
		}
	}
	return min
}
