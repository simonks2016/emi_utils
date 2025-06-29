package AnyPro

import "strconv"

type NumberType interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~float32 | ~float64
}

func Any2Number[T NumberType](d interface{}) T {
	var zero T
	if d == nil {
		return zero
	}

	switch v := d.(type) {
	case string:
		// 尝试转 float64，再强制转换为 T
		if f, err := strconv.ParseFloat(v, 64); err == nil {
			return castNumber[T](f)
		}
	case float64:
		return castNumber[T](v)
	case float32:
		return castNumber[T](float64(v))
	case int:
		return castNumber[T](float64(v))
	case int8:
		return castNumber[T](float64(v))
	case int16:
		return castNumber[T](float64(v))
	case int32:
		return castNumber[T](float64(v))
	case int64:
		return castNumber[T](float64(v))
	case uint:
		return castNumber[T](float64(v))
	case uint8:
		return castNumber[T](float64(v))
	case uint16:
		return castNumber[T](float64(v))
	case uint32:
		return castNumber[T](float64(v))
	case uint64:
		return castNumber[T](float64(v))
	default:
		// 不支持的类型
		return zero
	}
	return zero
}

// 帮助函数：将 float64 转换为目标类型
func castNumber[T NumberType](f float64) T {
	var t any

	switch any(*new(T)).(type) {
	case int:
		t = int(f)
	case int8:
		t = int8(f)
	case int16:
		t = int16(f)
	case int32:
		t = int32(f)
	case int64:
		t = int64(f)
	case uint:
		t = uint(f)
	case uint8:
		t = uint8(f)
	case uint16:
		t = uint16(f)
	case uint32:
		t = uint32(f)
	case uint64:
		t = uint64(f)
	case float32:
		t = float32(f)
	case float64:
		t = float64(f)
	default:
		t = *new(T) // zero value fallback
	}
	return t.(T)
}
