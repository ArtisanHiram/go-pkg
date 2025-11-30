package __obf_c3cd04a15c56f32f

import (
	"strconv"
)

type __obf_286b4b716339f0b1 struct {
	__obf_86279f81acb381f3
	__obf_d59813f23ad740a8 uint32
}

func (any *__obf_286b4b716339f0b1) LastError() error {
	return nil
}

func (any *__obf_286b4b716339f0b1) ValueType() ValueType {
	return NumberValue
}

func (any *__obf_286b4b716339f0b1) MustBeValid() Any {
	return any
}

func (any *__obf_286b4b716339f0b1) ToBool() bool {
	return any.__obf_d59813f23ad740a8 != 0
}

func (any *__obf_286b4b716339f0b1) ToInt() int {
	return int(any.__obf_d59813f23ad740a8)
}

func (any *__obf_286b4b716339f0b1) ToInt32() int32 {
	return int32(any.__obf_d59813f23ad740a8)
}

func (any *__obf_286b4b716339f0b1) ToInt64() int64 {
	return int64(any.__obf_d59813f23ad740a8)
}

func (any *__obf_286b4b716339f0b1) ToUint() uint {
	return uint(any.__obf_d59813f23ad740a8)
}

func (any *__obf_286b4b716339f0b1) ToUint32() uint32 {
	return any.__obf_d59813f23ad740a8
}

func (any *__obf_286b4b716339f0b1) ToUint64() uint64 {
	return uint64(any.__obf_d59813f23ad740a8)
}

func (any *__obf_286b4b716339f0b1) ToFloat32() float32 {
	return float32(any.__obf_d59813f23ad740a8)
}

func (any *__obf_286b4b716339f0b1) ToFloat64() float64 {
	return float64(any.__obf_d59813f23ad740a8)
}

func (any *__obf_286b4b716339f0b1) ToString() string {
	return strconv.FormatInt(int64(any.__obf_d59813f23ad740a8), 10)
}

func (any *__obf_286b4b716339f0b1) WriteTo(__obf_2361f5214e84c60f *Stream) {
	__obf_2361f5214e84c60f.
		WriteUint32(any.__obf_d59813f23ad740a8)
}

func (any *__obf_286b4b716339f0b1) Parse() *Iterator {
	return nil
}

func (any *__obf_286b4b716339f0b1) GetInterface() any {
	return any.__obf_d59813f23ad740a8
}
