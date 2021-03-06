// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/cheekybits/genny

package slices

// InsertBool在切片指定位置插入元素
// usage:
// 在列表倒数第一位之前插入元素: `InsertBool(s, len(s)-1, element)`
// 也可以简化为: `InsertBool(s, -1, element)`
func InsertBool(s []bool, index int, elements ...bool) []bool {
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

	t := make([]bool, len(s)+len(elements))
	copy(t[:index], s[:index])
	copy(t[index:index+len(elements)], elements)
	copy(t[index+len(elements):], s[index:])

	return t
}

// InsertByte在切片指定位置插入元素
// usage:
// 在列表倒数第一位之前插入元素: `InsertByte(s, len(s)-1, element)`
// 也可以简化为: `InsertByte(s, -1, element)`
func InsertByte(s []byte, index int, elements ...byte) []byte {
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

	t := make([]byte, len(s)+len(elements))
	copy(t[:index], s[:index])
	copy(t[index:index+len(elements)], elements)
	copy(t[index+len(elements):], s[index:])

	return t
}

// InsertComplex128在切片指定位置插入元素
// usage:
// 在列表倒数第一位之前插入元素: `InsertComplex128(s, len(s)-1, element)`
// 也可以简化为: `InsertComplex128(s, -1, element)`
func InsertComplex128(s []complex128, index int, elements ...complex128) []complex128 {
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

	t := make([]complex128, len(s)+len(elements))
	copy(t[:index], s[:index])
	copy(t[index:index+len(elements)], elements)
	copy(t[index+len(elements):], s[index:])

	return t
}

// InsertComplex64在切片指定位置插入元素
// usage:
// 在列表倒数第一位之前插入元素: `InsertComplex64(s, len(s)-1, element)`
// 也可以简化为: `InsertComplex64(s, -1, element)`
func InsertComplex64(s []complex64, index int, elements ...complex64) []complex64 {
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

	t := make([]complex64, len(s)+len(elements))
	copy(t[:index], s[:index])
	copy(t[index:index+len(elements)], elements)
	copy(t[index+len(elements):], s[index:])

	return t
}

// InsertError在切片指定位置插入元素
// usage:
// 在列表倒数第一位之前插入元素: `InsertError(s, len(s)-1, element)`
// 也可以简化为: `InsertError(s, -1, element)`
func InsertError(s []error, index int, elements ...error) []error {
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

	t := make([]error, len(s)+len(elements))
	copy(t[:index], s[:index])
	copy(t[index:index+len(elements)], elements)
	copy(t[index+len(elements):], s[index:])

	return t
}

// InsertFloat32在切片指定位置插入元素
// usage:
// 在列表倒数第一位之前插入元素: `InsertFloat32(s, len(s)-1, element)`
// 也可以简化为: `InsertFloat32(s, -1, element)`
func InsertFloat32(s []float32, index int, elements ...float32) []float32 {
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

	t := make([]float32, len(s)+len(elements))
	copy(t[:index], s[:index])
	copy(t[index:index+len(elements)], elements)
	copy(t[index+len(elements):], s[index:])

	return t
}

// InsertFloat64在切片指定位置插入元素
// usage:
// 在列表倒数第一位之前插入元素: `InsertFloat64(s, len(s)-1, element)`
// 也可以简化为: `InsertFloat64(s, -1, element)`
func InsertFloat64(s []float64, index int, elements ...float64) []float64 {
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

	t := make([]float64, len(s)+len(elements))
	copy(t[:index], s[:index])
	copy(t[index:index+len(elements)], elements)
	copy(t[index+len(elements):], s[index:])

	return t
}

// InsertInt在切片指定位置插入元素
// usage:
// 在列表倒数第一位之前插入元素: `InsertInt(s, len(s)-1, element)`
// 也可以简化为: `InsertInt(s, -1, element)`
func InsertInt(s []int, index int, elements ...int) []int {
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

	t := make([]int, len(s)+len(elements))
	copy(t[:index], s[:index])
	copy(t[index:index+len(elements)], elements)
	copy(t[index+len(elements):], s[index:])

	return t
}

// InsertInt16在切片指定位置插入元素
// usage:
// 在列表倒数第一位之前插入元素: `InsertInt16(s, len(s)-1, element)`
// 也可以简化为: `InsertInt16(s, -1, element)`
func InsertInt16(s []int16, index int, elements ...int16) []int16 {
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

	t := make([]int16, len(s)+len(elements))
	copy(t[:index], s[:index])
	copy(t[index:index+len(elements)], elements)
	copy(t[index+len(elements):], s[index:])

	return t
}

// InsertInt32在切片指定位置插入元素
// usage:
// 在列表倒数第一位之前插入元素: `InsertInt32(s, len(s)-1, element)`
// 也可以简化为: `InsertInt32(s, -1, element)`
func InsertInt32(s []int32, index int, elements ...int32) []int32 {
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

	t := make([]int32, len(s)+len(elements))
	copy(t[:index], s[:index])
	copy(t[index:index+len(elements)], elements)
	copy(t[index+len(elements):], s[index:])

	return t
}

// InsertInt64在切片指定位置插入元素
// usage:
// 在列表倒数第一位之前插入元素: `InsertInt64(s, len(s)-1, element)`
// 也可以简化为: `InsertInt64(s, -1, element)`
func InsertInt64(s []int64, index int, elements ...int64) []int64 {
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

	t := make([]int64, len(s)+len(elements))
	copy(t[:index], s[:index])
	copy(t[index:index+len(elements)], elements)
	copy(t[index+len(elements):], s[index:])

	return t
}

// InsertInt8在切片指定位置插入元素
// usage:
// 在列表倒数第一位之前插入元素: `InsertInt8(s, len(s)-1, element)`
// 也可以简化为: `InsertInt8(s, -1, element)`
func InsertInt8(s []int8, index int, elements ...int8) []int8 {
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

	t := make([]int8, len(s)+len(elements))
	copy(t[:index], s[:index])
	copy(t[index:index+len(elements)], elements)
	copy(t[index+len(elements):], s[index:])

	return t
}

// InsertRune在切片指定位置插入元素
// usage:
// 在列表倒数第一位之前插入元素: `InsertRune(s, len(s)-1, element)`
// 也可以简化为: `InsertRune(s, -1, element)`
func InsertRune(s []rune, index int, elements ...rune) []rune {
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

	t := make([]rune, len(s)+len(elements))
	copy(t[:index], s[:index])
	copy(t[index:index+len(elements)], elements)
	copy(t[index+len(elements):], s[index:])

	return t
}

// InsertString在切片指定位置插入元素
// usage:
// 在列表倒数第一位之前插入元素: `InsertString(s, len(s)-1, element)`
// 也可以简化为: `InsertString(s, -1, element)`
func InsertString(s []string, index int, elements ...string) []string {
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

	t := make([]string, len(s)+len(elements))
	copy(t[:index], s[:index])
	copy(t[index:index+len(elements)], elements)
	copy(t[index+len(elements):], s[index:])

	return t
}

// InsertUint在切片指定位置插入元素
// usage:
// 在列表倒数第一位之前插入元素: `InsertUint(s, len(s)-1, element)`
// 也可以简化为: `InsertUint(s, -1, element)`
func InsertUint(s []uint, index int, elements ...uint) []uint {
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

	t := make([]uint, len(s)+len(elements))
	copy(t[:index], s[:index])
	copy(t[index:index+len(elements)], elements)
	copy(t[index+len(elements):], s[index:])

	return t
}

// InsertUint16在切片指定位置插入元素
// usage:
// 在列表倒数第一位之前插入元素: `InsertUint16(s, len(s)-1, element)`
// 也可以简化为: `InsertUint16(s, -1, element)`
func InsertUint16(s []uint16, index int, elements ...uint16) []uint16 {
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

	t := make([]uint16, len(s)+len(elements))
	copy(t[:index], s[:index])
	copy(t[index:index+len(elements)], elements)
	copy(t[index+len(elements):], s[index:])

	return t
}

// InsertUint32在切片指定位置插入元素
// usage:
// 在列表倒数第一位之前插入元素: `InsertUint32(s, len(s)-1, element)`
// 也可以简化为: `InsertUint32(s, -1, element)`
func InsertUint32(s []uint32, index int, elements ...uint32) []uint32 {
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

	t := make([]uint32, len(s)+len(elements))
	copy(t[:index], s[:index])
	copy(t[index:index+len(elements)], elements)
	copy(t[index+len(elements):], s[index:])

	return t
}

// InsertUint64在切片指定位置插入元素
// usage:
// 在列表倒数第一位之前插入元素: `InsertUint64(s, len(s)-1, element)`
// 也可以简化为: `InsertUint64(s, -1, element)`
func InsertUint64(s []uint64, index int, elements ...uint64) []uint64 {
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

	t := make([]uint64, len(s)+len(elements))
	copy(t[:index], s[:index])
	copy(t[index:index+len(elements)], elements)
	copy(t[index+len(elements):], s[index:])

	return t
}

// InsertUint8在切片指定位置插入元素
// usage:
// 在列表倒数第一位之前插入元素: `InsertUint8(s, len(s)-1, element)`
// 也可以简化为: `InsertUint8(s, -1, element)`
func InsertUint8(s []uint8, index int, elements ...uint8) []uint8 {
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

	t := make([]uint8, len(s)+len(elements))
	copy(t[:index], s[:index])
	copy(t[index:index+len(elements)], elements)
	copy(t[index+len(elements):], s[index:])

	return t
}

// InsertUintptr在切片指定位置插入元素
// usage:
// 在列表倒数第一位之前插入元素: `InsertUintptr(s, len(s)-1, element)`
// 也可以简化为: `InsertUintptr(s, -1, element)`
func InsertUintptr(s []uintptr, index int, elements ...uintptr) []uintptr {
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

	t := make([]uintptr, len(s)+len(elements))
	copy(t[:index], s[:index])
	copy(t[index:index+len(elements)], elements)
	copy(t[index+len(elements):], s[index:])

	return t
}
