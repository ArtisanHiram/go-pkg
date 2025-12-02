package __obf_c7b79b12b33d8238

import "fmt"

type __obf_cb30187c64f44cee struct {
	__obf_3a44d9c5204c2e12
	__obf_ea0680f8b461a85b error
}

func __obf_5b62b1b9d47ae70b(__obf_5dde9fb4007294d4 []any) *__obf_cb30187c64f44cee {
	return &__obf_cb30187c64f44cee{__obf_3a44d9c5204c2e12{}, fmt.Errorf("%v not found", __obf_5dde9fb4007294d4)}
}

func (any *__obf_cb30187c64f44cee,) LastError() error {
	return any.__obf_ea0680f8b461a85b
}

func (any *__obf_cb30187c64f44cee,) ValueType() ValueType {
	return InvalidValue
}

func (any *__obf_cb30187c64f44cee,) MustBeValid() Any {
	panic(any.__obf_ea0680f8b461a85b)
}

func (any *__obf_cb30187c64f44cee,) ToBool() bool {
	return false
}

func (any *__obf_cb30187c64f44cee,) ToInt() int {
	return 0
}

func (any *__obf_cb30187c64f44cee,) ToInt32() int32 {
	return 0
}

func (any *__obf_cb30187c64f44cee,) ToInt64() int64 {
	return 0
}

func (any *__obf_cb30187c64f44cee,) ToUint() uint {
	return 0
}

func (any *__obf_cb30187c64f44cee,) ToUint32() uint32 {
	return 0
}

func (any *__obf_cb30187c64f44cee,) ToUint64() uint64 {
	return 0
}

func (any *__obf_cb30187c64f44cee,) ToFloat32() float32 {
	return 0
}

func (any *__obf_cb30187c64f44cee,) ToFloat64() float64 {
	return 0
}

func (any *__obf_cb30187c64f44cee,) ToString() string {
	return ""
}

func (any *__obf_cb30187c64f44cee,) WriteTo(__obf_d8f50bcfe87d8b47 *Stream) {
}

func (any *__obf_cb30187c64f44cee,) Get(__obf_5dde9fb4007294d4 ...any) Any {
	if any.__obf_ea0680f8b461a85b == nil {
		return &__obf_cb30187c64f44cee{__obf_3a44d9c5204c2e12{}, fmt.Errorf("get %v from invalid", __obf_5dde9fb4007294d4)}
	}
	return &__obf_cb30187c64f44cee{__obf_3a44d9c5204c2e12{}, fmt.Errorf("%v, get %v from invalid", any.__obf_ea0680f8b461a85b, __obf_5dde9fb4007294d4)}
}

func (any *__obf_cb30187c64f44cee,) Parse() *Iterator {
	return nil
}

func (any *__obf_cb30187c64f44cee,) GetInterface() any {
	return nil
}
