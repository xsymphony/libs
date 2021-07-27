package template

//go:generate genny -in=$GOFILE -out=../$GOFILE -pkg=slices gen "Template=string,int,int32,int64,uint32,uint64,bool"
func AnyTemplate(s []Template, assert func(Template) bool) bool {
	for i := 0; i < len(s); i++ {
		if assert(s[i]) {
			return true
		}
	}
	return false
}

func AllTemplate(s []Template, assert func(Template) bool) bool {
	for i := 0; i < len(s); i++ {
		if !assert(s[i]) {
			return false
		}
	}
	return true
}
