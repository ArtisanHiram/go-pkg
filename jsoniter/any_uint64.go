package __obf_5b802ce8d9ba56d6

import (
	"strconv"
)

type __obf_1eb46b63f55e7ee7 struct {
	__obf_fb75d4e4562132ae
	__obf_5406b11e33b9d1d3 uint64
}

func (any *__obf_1eb46b63f55e7ee7) LastError() error {
	return nil
}

func (any *__obf_1eb46b63f55e7ee7) ValueType() ValueType {
	return NumberValue
}

func (any *__obf_1eb46b63f55e7ee7) MustBeValid() Any {
	return any
}

func (any *__obf_1eb46b63f55e7ee7) ToBool() bool {
	return any.__obf_5406b11e33b9d1d3 != 0
}

func (any *__obf_1eb46b63f55e7ee7) ToInt() int {
	return int(any.__obf_5406b11e33b9d1d3)
}

func (any *__obf_1eb46b63f55e7ee7) ToInt32() int32 {
	return int32(any.__obf_5406b11e33b9d1d3)
}

func (any *__obf_1eb46b63f55e7ee7) ToInt64() int64 {
	return int64(any.__obf_5406b11e33b9d1d3)
}

func (any *__obf_1eb46b63f55e7ee7) ToUint() uint {
	return uint(any.__obf_5406b11e33b9d1d3)
}

func (any *__obf_1eb46b63f55e7ee7) ToUint32() uint32 {
	return uint32(any.__obf_5406b11e33b9d1d3)
}

func (any *__obf_1eb46b63f55e7ee7) ToUint64() uint64 {
	return any.__obf_5406b11e33b9d1d3
}

func (any *__obf_1eb46b63f55e7ee7) ToFloat32() float32 {
	return float32(any.__obf_5406b11e33b9d1d3)
}

func (any *__obf_1eb46b63f55e7ee7) ToFloat64() float64 {
	return float64(any.__obf_5406b11e33b9d1d3)
}

func (any *__obf_1eb46b63f55e7ee7) ToString() string {
	return strconv.FormatUint(any.__obf_5406b11e33b9d1d3, 10)
}

func (any *__obf_1eb46b63f55e7ee7) WriteTo(__obf_00fc491caa967a74 *Stream) {
	__obf_00fc491caa967a74.
		WriteUint64(any.__obf_5406b11e33b9d1d3)
}

func (any *__obf_1eb46b63f55e7ee7) Parse() *Iterator {
	return nil
}

func (any *__obf_1eb46b63f55e7ee7) GetInterface() any {
	return any.__obf_5406b11e33b9d1d3
}
