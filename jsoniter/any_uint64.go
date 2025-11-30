package __obf_c3cd04a15c56f32f

import (
	"strconv"
)

type __obf_226272cc8e4cb56e struct {
	__obf_86279f81acb381f3
	__obf_d59813f23ad740a8 uint64
}

func (any *__obf_226272cc8e4cb56e) LastError() error {
	return nil
}

func (any *__obf_226272cc8e4cb56e) ValueType() ValueType {
	return NumberValue
}

func (any *__obf_226272cc8e4cb56e) MustBeValid() Any {
	return any
}

func (any *__obf_226272cc8e4cb56e) ToBool() bool {
	return any.__obf_d59813f23ad740a8 != 0
}

func (any *__obf_226272cc8e4cb56e) ToInt() int {
	return int(any.__obf_d59813f23ad740a8)
}

func (any *__obf_226272cc8e4cb56e) ToInt32() int32 {
	return int32(any.__obf_d59813f23ad740a8)
}

func (any *__obf_226272cc8e4cb56e) ToInt64() int64 {
	return int64(any.__obf_d59813f23ad740a8)
}

func (any *__obf_226272cc8e4cb56e) ToUint() uint {
	return uint(any.__obf_d59813f23ad740a8)
}

func (any *__obf_226272cc8e4cb56e) ToUint32() uint32 {
	return uint32(any.__obf_d59813f23ad740a8)
}

func (any *__obf_226272cc8e4cb56e) ToUint64() uint64 {
	return any.__obf_d59813f23ad740a8
}

func (any *__obf_226272cc8e4cb56e) ToFloat32() float32 {
	return float32(any.__obf_d59813f23ad740a8)
}

func (any *__obf_226272cc8e4cb56e) ToFloat64() float64 {
	return float64(any.__obf_d59813f23ad740a8)
}

func (any *__obf_226272cc8e4cb56e) ToString() string {
	return strconv.FormatUint(any.__obf_d59813f23ad740a8, 10)
}

func (any *__obf_226272cc8e4cb56e) WriteTo(__obf_2361f5214e84c60f *Stream) {
	__obf_2361f5214e84c60f.
		WriteUint64(any.__obf_d59813f23ad740a8)
}

func (any *__obf_226272cc8e4cb56e) Parse() *Iterator {
	return nil
}

func (any *__obf_226272cc8e4cb56e) GetInterface() any {
	return any.__obf_d59813f23ad740a8
}
