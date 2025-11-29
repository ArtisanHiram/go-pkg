package __obf_5b802ce8d9ba56d6

import (
	"strconv"
)

type __obf_dbd8d5f30edec5df struct {
	__obf_fb75d4e4562132ae
	__obf_5406b11e33b9d1d3 uint32
}

func (any *__obf_dbd8d5f30edec5df) LastError() error {
	return nil
}

func (any *__obf_dbd8d5f30edec5df) ValueType() ValueType {
	return NumberValue
}

func (any *__obf_dbd8d5f30edec5df) MustBeValid() Any {
	return any
}

func (any *__obf_dbd8d5f30edec5df) ToBool() bool {
	return any.__obf_5406b11e33b9d1d3 != 0
}

func (any *__obf_dbd8d5f30edec5df) ToInt() int {
	return int(any.__obf_5406b11e33b9d1d3)
}

func (any *__obf_dbd8d5f30edec5df) ToInt32() int32 {
	return int32(any.__obf_5406b11e33b9d1d3)
}

func (any *__obf_dbd8d5f30edec5df) ToInt64() int64 {
	return int64(any.__obf_5406b11e33b9d1d3)
}

func (any *__obf_dbd8d5f30edec5df) ToUint() uint {
	return uint(any.__obf_5406b11e33b9d1d3)
}

func (any *__obf_dbd8d5f30edec5df) ToUint32() uint32 {
	return any.__obf_5406b11e33b9d1d3
}

func (any *__obf_dbd8d5f30edec5df) ToUint64() uint64 {
	return uint64(any.__obf_5406b11e33b9d1d3)
}

func (any *__obf_dbd8d5f30edec5df) ToFloat32() float32 {
	return float32(any.__obf_5406b11e33b9d1d3)
}

func (any *__obf_dbd8d5f30edec5df) ToFloat64() float64 {
	return float64(any.__obf_5406b11e33b9d1d3)
}

func (any *__obf_dbd8d5f30edec5df) ToString() string {
	return strconv.FormatInt(int64(any.__obf_5406b11e33b9d1d3), 10)
}

func (any *__obf_dbd8d5f30edec5df) WriteTo(__obf_00fc491caa967a74 *Stream) {
	__obf_00fc491caa967a74.
		WriteUint32(any.__obf_5406b11e33b9d1d3)
}

func (any *__obf_dbd8d5f30edec5df) Parse() *Iterator {
	return nil
}

func (any *__obf_dbd8d5f30edec5df) GetInterface() any {
	return any.__obf_5406b11e33b9d1d3
}
