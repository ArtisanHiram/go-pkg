package __obf_5b802ce8d9ba56d6

import (
	"strconv"
)

type __obf_7d29c083a07482f6 struct {
	__obf_fb75d4e4562132ae
	__obf_5406b11e33b9d1d3 int64
}

func (any *__obf_7d29c083a07482f6) LastError() error {
	return nil
}

func (any *__obf_7d29c083a07482f6) ValueType() ValueType {
	return NumberValue
}

func (any *__obf_7d29c083a07482f6) MustBeValid() Any {
	return any
}

func (any *__obf_7d29c083a07482f6) ToBool() bool {
	return any.__obf_5406b11e33b9d1d3 != 0
}

func (any *__obf_7d29c083a07482f6) ToInt() int {
	return int(any.__obf_5406b11e33b9d1d3)
}

func (any *__obf_7d29c083a07482f6) ToInt32() int32 {
	return int32(any.__obf_5406b11e33b9d1d3)
}

func (any *__obf_7d29c083a07482f6) ToInt64() int64 {
	return any.__obf_5406b11e33b9d1d3
}

func (any *__obf_7d29c083a07482f6) ToUint() uint {
	return uint(any.__obf_5406b11e33b9d1d3)
}

func (any *__obf_7d29c083a07482f6) ToUint32() uint32 {
	return uint32(any.__obf_5406b11e33b9d1d3)
}

func (any *__obf_7d29c083a07482f6) ToUint64() uint64 {
	return uint64(any.__obf_5406b11e33b9d1d3)
}

func (any *__obf_7d29c083a07482f6) ToFloat32() float32 {
	return float32(any.__obf_5406b11e33b9d1d3)
}

func (any *__obf_7d29c083a07482f6) ToFloat64() float64 {
	return float64(any.__obf_5406b11e33b9d1d3)
}

func (any *__obf_7d29c083a07482f6) ToString() string {
	return strconv.FormatInt(any.__obf_5406b11e33b9d1d3, 10)
}

func (any *__obf_7d29c083a07482f6) WriteTo(__obf_00fc491caa967a74 *Stream) {
	__obf_00fc491caa967a74.
		WriteInt64(any.__obf_5406b11e33b9d1d3)
}

func (any *__obf_7d29c083a07482f6) Parse() *Iterator {
	return nil
}

func (any *__obf_7d29c083a07482f6) GetInterface() any {
	return any.__obf_5406b11e33b9d1d3
}
