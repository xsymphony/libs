package collection

import "github.com/cheekybits/genny/generic"

type Template generic.Type

//go:generate genny -in=$GOFILE -out=gen-$GOFILE gen "Template=string,int,int64,int32"
type SetTemplate struct {
	values map[Template]struct{}
}

func NewSetTemplate(elements ...Template) *SetTemplate {
	set := &SetTemplate{
		values: make(map[Template]struct{}),
	}
	for _, element := range elements {
		set.add(element)
	}

	return set
}

func (set *SetTemplate) has(elem Template) bool {
	_, exist := set.values[elem]
	return exist
}

func (set *SetTemplate) add(elem Template) {
	if set.values == nil {
		set.values = make(map[Template]struct{})
	}
	set.values[elem] = struct{}{}
}

func (set *SetTemplate) Add(elem Template) bool {
	if set.has(elem) {
		return true
	}
	set.add(elem)
	return false
}

func (set *SetTemplate) Update(elements ...Template) {
	for _, element := range elements {
		set.add(element)
	}
}

func (set *SetTemplate) Remove(elem Template) bool {
	if !set.has(elem) {
		return false
	}
	delete(set.values, elem)
	return true
}

func (set *SetTemplate) Has(elem Template) bool {
	return set.has(elem)
}

// 遍历集合
func (set *SetTemplate) Range(f func(element Template)) {
	for value := range set.values {
		f(value)
	}
}

// copy一个新的集合
func (set *SetTemplate) Copy() *SetTemplate {
	copyed := NewSetTemplate()
	set.Range(func(element Template) {
		copyed.Add(element)
	})
	return copyed
}

// 求并集
func (set *SetTemplate) Union(another *SetTemplate) *SetTemplate {
	union := set.Copy()
	another.Range(func(element Template) {
		union.Add(element)
	})
	return union
}

// 求交集
func (set *SetTemplate) Intersect(another *SetTemplate) *SetTemplate {
	intersect := NewSetTemplate()
	set.Range(func(element Template) {
		if another.Has(element) {
			intersect.Add(element)
		}
	})
	return intersect
}

// 求差集
func (set *SetTemplate) Except(another *SetTemplate) *SetTemplate {
	except := NewSetTemplate()
	set.Range(func(element Template) {
		if !another.Has(element) {
			except.Add(element)
		}
	})
	return except
}

// 集合长度
func (set *SetTemplate) Length() int {
	return len(set.values)
}

// 集合转化为slice结构
func (set *SetTemplate) TemplateList() []Template {
	lst := make([]Template, 0, set.Length())
	set.Range(func(element Template) {
		lst = append(lst, element)
	})

	return lst
}
