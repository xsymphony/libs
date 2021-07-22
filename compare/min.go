package compare

//go:generate genny -in=$GOFILE -out=gen-$GOFILE gen "Template=int,int64,int32,float64,float32"
func MinTemplate(arg Template, others ...Template) Template {
	min := arg
	for _, other := range others {
		if other < min {
			min = other
		}
	}
	return min
}
