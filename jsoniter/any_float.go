package __obf_c3cd04a15c56f32f

import (
	"strconv"
)

type __obf_299c6b904e6b8eea struct {
	__obf_86279f81acb381f3
	__obf_d59813f23ad740a8 float64
}

func (any *__obf_299c6b904e6b8eea) Parse() *Iterator {
	return nil
}

func (any *__obf_299c6b904e6b8eea) ValueType() ValueType {
	return NumberValue
}

func (any *__obf_299c6b904e6b8eea) MustBeValid() Any {
	return any
}

func (any *__obf_299c6b904e6b8eea) LastError() error {
	return nil
}

func (any *__obf_299c6b904e6b8eea) ToBool() bool {
	return any.ToFloat64() != 0
}

func (any *__obf_299c6b904e6b8eea) ToInt() int {
	return int(any.__obf_d59813f23ad740a8)
}

func (any *__obf_299c6b904e6b8eea) ToInt32() int32 {
	return int32(any.__obf_d59813f23ad740a8)
}

func (any *__obf_299c6b904e6b8eea) ToInt64() int64 {
	return int64(any.__obf_d59813f23ad740a8)
}

func (any *__obf_299c6b904e6b8eea) ToUint() uint {
	if any.__obf_d59813f23ad740a8 > 0 {
		return uint(any.__obf_d59813f23ad740a8)
	}
	return 0
}

func (any *__obf_299c6b904e6b8eea) ToUint32() uint32 {
	if any.__obf_d59813f23ad740a8 > 0 {
		return uint32(any.__obf_d59813f23ad740a8)
	}
	return 0
}

func (any *__obf_299c6b904e6b8eea) ToUint64() uint64 {
	if any.__obf_d59813f23ad740a8 > 0 {
		return uint64(any.__obf_d59813f23ad740a8)
	}
	return 0
}

func (any *__obf_299c6b904e6b8eea) ToFloat32() float32 {
	return float32(any.__obf_d59813f23ad740a8)
}

func (any *__obf_299c6b904e6b8eea) ToFloat64() float64 {
	return any.__obf_d59813f23ad740a8
}

func (any *__obf_299c6b904e6b8eea) ToString() string {
	return strconv.FormatFloat(any.__obf_d59813f23ad740a8, 'E', -1, 64)
}

func (any *__obf_299c6b904e6b8eea) WriteTo(__obf_2361f5214e84c60f *Stream) {
	__obf_2361f5214e84c60f.
		WriteFloat64(any.__obf_d59813f23ad740a8)
}

func (any *__obf_299c6b904e6b8eea) GetInterface() any {
	return any.__obf_d59813f23ad740a8
}
