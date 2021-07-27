package template

//go:generate genny -in=$GOFILE -out=../$GOFILE -pkg=slices gen "Template=BUILTINS"
// InsertTemplate在切片指定位置插入元素
// usage:
// 在列表倒数第一位之前插入元素: `InsertTemplate(s, len(s)-1, element)`
// 也可以简化为: `InsertTemplate(s, -1, element)`
func InsertTemplate(s []Template, index int, elements ...Template) []Template {
	length := len(s)
	if length == 0 {
		return s
	}

	// 传入负数时，计算绝对位置
	if index < 0 {
		index = length + index
	}

	// 绝对位置大于s的长度时，代表直接在s的尾端插入elements
	if index >= length {
		index = length
	}

	// 绝对位置小于0，代表需要在s之前插入elements
	if index <= 0 {
		index = 0
	}

	t := make([]Template, len(s)+len(elements))
	copy(t[:index], s[:index])
	copy(t[index:index+len(elements)], elements)
	copy(t[index+len(elements):], s[index:])

	return t
}
