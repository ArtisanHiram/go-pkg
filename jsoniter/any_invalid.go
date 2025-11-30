package __obf_c3cd04a15c56f32f

import "fmt"

type __obf_5183f6b11c973faf struct {
	__obf_86279f81acb381f3
	__obf_5966ecc5edb9b17e error
}

func __obf_5a6677fd9236ae8d(__obf_b90e4ca332607787 []any) *__obf_5183f6b11c973faf {
	return &__obf_5183f6b11c973faf{__obf_86279f81acb381f3{}, fmt.Errorf("%v not found", __obf_b90e4ca332607787)}
}

func (any *__obf_5183f6b11c973faf,) LastError() error {
	return any.__obf_5966ecc5edb9b17e
}

func (any *__obf_5183f6b11c973faf,) ValueType() ValueType {
	return InvalidValue
}

func (any *__obf_5183f6b11c973faf,) MustBeValid() Any {
	panic(any.__obf_5966ecc5edb9b17e)
}

func (any *__obf_5183f6b11c973faf,) ToBool() bool {
	return false
}

func (any *__obf_5183f6b11c973faf,) ToInt() int {
	return 0
}

func (any *__obf_5183f6b11c973faf,) ToInt32() int32 {
	return 0
}

func (any *__obf_5183f6b11c973faf,) ToInt64() int64 {
	return 0
}

func (any *__obf_5183f6b11c973faf,) ToUint() uint {
	return 0
}

func (any *__obf_5183f6b11c973faf,) ToUint32() uint32 {
	return 0
}

func (any *__obf_5183f6b11c973faf,) ToUint64() uint64 {
	return 0
}

func (any *__obf_5183f6b11c973faf,) ToFloat32() float32 {
	return 0
}

func (any *__obf_5183f6b11c973faf,) ToFloat64() float64 {
	return 0
}

func (any *__obf_5183f6b11c973faf,) ToString() string {
	return ""
}

func (any *__obf_5183f6b11c973faf,) WriteTo(__obf_2361f5214e84c60f *Stream) {
}

func (any *__obf_5183f6b11c973faf,) Get(__obf_b90e4ca332607787 ...any) Any {
	if any.__obf_5966ecc5edb9b17e == nil {
		return &__obf_5183f6b11c973faf{__obf_86279f81acb381f3{}, fmt.Errorf("get %v from invalid", __obf_b90e4ca332607787)}
	}
	return &__obf_5183f6b11c973faf{__obf_86279f81acb381f3{}, fmt.Errorf("%v, get %v from invalid", any.__obf_5966ecc5edb9b17e, __obf_b90e4ca332607787)}
}

func (any *__obf_5183f6b11c973faf,) Parse() *Iterator {
	return nil
}

func (any *__obf_5183f6b11c973faf,) GetInterface() any {
	return nil
}
