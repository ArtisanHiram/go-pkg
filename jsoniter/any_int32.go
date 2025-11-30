package __obf_c3cd04a15c56f32f

import (
	"strconv"
)

type __obf_c335eb1f76d110d7 struct {
	__obf_86279f81acb381f3
	__obf_d59813f23ad740a8 int32
}

func (any *__obf_c335eb1f76d110d7) LastError() error {
	return nil
}

func (any *__obf_c335eb1f76d110d7) ValueType() ValueType {
	return NumberValue
}

func (any *__obf_c335eb1f76d110d7) MustBeValid() Any {
	return any
}

func (any *__obf_c335eb1f76d110d7) ToBool() bool {
	return any.__obf_d59813f23ad740a8 != 0
}

func (any *__obf_c335eb1f76d110d7) ToInt() int {
	return int(any.__obf_d59813f23ad740a8)
}

func (any *__obf_c335eb1f76d110d7) ToInt32() int32 {
	return any.__obf_d59813f23ad740a8
}

func (any *__obf_c335eb1f76d110d7) ToInt64() int64 {
	return int64(any.__obf_d59813f23ad740a8)
}

func (any *__obf_c335eb1f76d110d7) ToUint() uint {
	return uint(any.__obf_d59813f23ad740a8)
}

func (any *__obf_c335eb1f76d110d7) ToUint32() uint32 {
	return uint32(any.__obf_d59813f23ad740a8)
}

func (any *__obf_c335eb1f76d110d7) ToUint64() uint64 {
	return uint64(any.__obf_d59813f23ad740a8)
}

func (any *__obf_c335eb1f76d110d7) ToFloat32() float32 {
	return float32(any.__obf_d59813f23ad740a8)
}

func (any *__obf_c335eb1f76d110d7) ToFloat64() float64 {
	return float64(any.__obf_d59813f23ad740a8)
}

func (any *__obf_c335eb1f76d110d7) ToString() string {
	return strconv.FormatInt(int64(any.__obf_d59813f23ad740a8), 10)
}

func (any *__obf_c335eb1f76d110d7) WriteTo(__obf_2361f5214e84c60f *Stream) {
	__obf_2361f5214e84c60f.
		WriteInt32(any.__obf_d59813f23ad740a8)
}

func (any *__obf_c335eb1f76d110d7) Parse() *Iterator {
	return nil
}

func (any *__obf_c335eb1f76d110d7) GetInterface() any {
	return any.__obf_d59813f23ad740a8
}
