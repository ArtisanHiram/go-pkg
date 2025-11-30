package __obf_c3cd04a15c56f32f

import (
	"strconv"
)

type __obf_e42a057d4a279be0 struct {
	__obf_86279f81acb381f3
	__obf_d59813f23ad740a8 int64
}

func (any *__obf_e42a057d4a279be0) LastError() error {
	return nil
}

func (any *__obf_e42a057d4a279be0) ValueType() ValueType {
	return NumberValue
}

func (any *__obf_e42a057d4a279be0) MustBeValid() Any {
	return any
}

func (any *__obf_e42a057d4a279be0) ToBool() bool {
	return any.__obf_d59813f23ad740a8 != 0
}

func (any *__obf_e42a057d4a279be0) ToInt() int {
	return int(any.__obf_d59813f23ad740a8)
}

func (any *__obf_e42a057d4a279be0) ToInt32() int32 {
	return int32(any.__obf_d59813f23ad740a8)
}

func (any *__obf_e42a057d4a279be0) ToInt64() int64 {
	return any.__obf_d59813f23ad740a8
}

func (any *__obf_e42a057d4a279be0) ToUint() uint {
	return uint(any.__obf_d59813f23ad740a8)
}

func (any *__obf_e42a057d4a279be0) ToUint32() uint32 {
	return uint32(any.__obf_d59813f23ad740a8)
}

func (any *__obf_e42a057d4a279be0) ToUint64() uint64 {
	return uint64(any.__obf_d59813f23ad740a8)
}

func (any *__obf_e42a057d4a279be0) ToFloat32() float32 {
	return float32(any.__obf_d59813f23ad740a8)
}

func (any *__obf_e42a057d4a279be0) ToFloat64() float64 {
	return float64(any.__obf_d59813f23ad740a8)
}

func (any *__obf_e42a057d4a279be0) ToString() string {
	return strconv.FormatInt(any.__obf_d59813f23ad740a8, 10)
}

func (any *__obf_e42a057d4a279be0) WriteTo(__obf_2361f5214e84c60f *Stream) {
	__obf_2361f5214e84c60f.
		WriteInt64(any.__obf_d59813f23ad740a8)
}

func (any *__obf_e42a057d4a279be0) Parse() *Iterator {
	return nil
}

func (any *__obf_e42a057d4a279be0) GetInterface() any {
	return any.__obf_d59813f23ad740a8
}
