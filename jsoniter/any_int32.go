package __obf_5b802ce8d9ba56d6

import (
	"strconv"
)

type __obf_3b63a102cb2fa2bb struct {
	__obf_fb75d4e4562132ae
	__obf_5406b11e33b9d1d3 int32
}

func (any *__obf_3b63a102cb2fa2bb) LastError() error {
	return nil
}

func (any *__obf_3b63a102cb2fa2bb) ValueType() ValueType {
	return NumberValue
}

func (any *__obf_3b63a102cb2fa2bb) MustBeValid() Any {
	return any
}

func (any *__obf_3b63a102cb2fa2bb) ToBool() bool {
	return any.__obf_5406b11e33b9d1d3 != 0
}

func (any *__obf_3b63a102cb2fa2bb) ToInt() int {
	return int(any.__obf_5406b11e33b9d1d3)
}

func (any *__obf_3b63a102cb2fa2bb) ToInt32() int32 {
	return any.__obf_5406b11e33b9d1d3
}

func (any *__obf_3b63a102cb2fa2bb) ToInt64() int64 {
	return int64(any.__obf_5406b11e33b9d1d3)
}

func (any *__obf_3b63a102cb2fa2bb) ToUint() uint {
	return uint(any.__obf_5406b11e33b9d1d3)
}

func (any *__obf_3b63a102cb2fa2bb) ToUint32() uint32 {
	return uint32(any.__obf_5406b11e33b9d1d3)
}

func (any *__obf_3b63a102cb2fa2bb) ToUint64() uint64 {
	return uint64(any.__obf_5406b11e33b9d1d3)
}

func (any *__obf_3b63a102cb2fa2bb) ToFloat32() float32 {
	return float32(any.__obf_5406b11e33b9d1d3)
}

func (any *__obf_3b63a102cb2fa2bb) ToFloat64() float64 {
	return float64(any.__obf_5406b11e33b9d1d3)
}

func (any *__obf_3b63a102cb2fa2bb) ToString() string {
	return strconv.FormatInt(int64(any.__obf_5406b11e33b9d1d3), 10)
}

func (any *__obf_3b63a102cb2fa2bb) WriteTo(__obf_00fc491caa967a74 *Stream) {
	__obf_00fc491caa967a74.
		WriteInt32(any.__obf_5406b11e33b9d1d3)
}

func (any *__obf_3b63a102cb2fa2bb) Parse() *Iterator {
	return nil
}

func (any *__obf_3b63a102cb2fa2bb) GetInterface() any {
	return any.__obf_5406b11e33b9d1d3
}
