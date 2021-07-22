package functional

import "github.com/cheekybits/genny/generic"

type Template generic.Type

//go:generate genny -in=$GOFILE -out=gen-$GOFILE gen "Template=string,int,int64,int32"
// IfTemplate是三目运算
// 例如，期望判断a,b的大小，a>=b的时候返回a，否则返回b
// 正常可写作
// ```
// if a >= b {
//     return a
// } else {
//     return b
// }
// ```
// 使用IfTemplate，可简化为
// ```
// IfTemplate(func() {return a >= b })(a)(b)
// ```
// 等价于JavaScript中`a>=b?a:b`
func IfTemplate(condition func() bool) func(Template) func(Template) Template {
	return func(trueValue Template) func(Template) Template {
		return func(falseValue Template) Template {
			if condition() {
				return trueValue
			}
			return falseValue
		}
	}
}
