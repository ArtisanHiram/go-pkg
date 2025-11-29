package __obf_5b802ce8d9ba56d6

import "fmt"

type __obf_00698f76feb93563 struct {
	__obf_fb75d4e4562132ae
	__obf_6d071d48f3b321ad error
}

func __obf_e990980207770dd2(__obf_c5d71353f4393b3c []any) *__obf_00698f76feb93563 {
	return &__obf_00698f76feb93563{__obf_fb75d4e4562132ae{}, fmt.Errorf("%v not found", __obf_c5d71353f4393b3c)}
}

func (any *__obf_00698f76feb93563,) LastError() error {
	return any.__obf_6d071d48f3b321ad
}

func (any *__obf_00698f76feb93563,) ValueType() ValueType {
	return InvalidValue
}

func (any *__obf_00698f76feb93563,) MustBeValid() Any {
	panic(any.__obf_6d071d48f3b321ad)
}

func (any *__obf_00698f76feb93563,) ToBool() bool {
	return false
}

func (any *__obf_00698f76feb93563,) ToInt() int {
	return 0
}

func (any *__obf_00698f76feb93563,) ToInt32() int32 {
	return 0
}

func (any *__obf_00698f76feb93563,) ToInt64() int64 {
	return 0
}

func (any *__obf_00698f76feb93563,) ToUint() uint {
	return 0
}

func (any *__obf_00698f76feb93563,) ToUint32() uint32 {
	return 0
}

func (any *__obf_00698f76feb93563,) ToUint64() uint64 {
	return 0
}

func (any *__obf_00698f76feb93563,) ToFloat32() float32 {
	return 0
}

func (any *__obf_00698f76feb93563,) ToFloat64() float64 {
	return 0
}

func (any *__obf_00698f76feb93563,) ToString() string {
	return ""
}

func (any *__obf_00698f76feb93563,) WriteTo(__obf_00fc491caa967a74 *Stream) {
}

func (any *__obf_00698f76feb93563,) Get(__obf_c5d71353f4393b3c ...any) Any {
	if any.__obf_6d071d48f3b321ad == nil {
		return &__obf_00698f76feb93563{__obf_fb75d4e4562132ae{}, fmt.Errorf("get %v from invalid", __obf_c5d71353f4393b3c)}
	}
	return &__obf_00698f76feb93563{__obf_fb75d4e4562132ae{}, fmt.Errorf("%v, get %v from invalid", any.__obf_6d071d48f3b321ad, __obf_c5d71353f4393b3c)}
}

func (any *__obf_00698f76feb93563,) Parse() *Iterator {
	return nil
}

func (any *__obf_00698f76feb93563,) GetInterface() any {
	return nil
}
