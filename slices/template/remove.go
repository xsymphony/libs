package template

//go:generate genny -in=$GOFILE -out=../$GOFILE -pkg=slices gen "Template=BUILTINS"
// RemoveTemplate在切片移除指定的元素，返回一个新的切片
// usage:
// 移除元素: `RemoveTemplate(s, element)`
func RemoveTemplate(s []Template, element Template) []Template {
	t := make([]Template, 0, len(s))
	for _, i := range s {
		if i == element {
			continue
		}
		t = append(t, i)
	}
	return t
}
