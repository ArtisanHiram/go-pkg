package __obf_5b802ce8d9ba56d6

import (
	"strconv"
)

type __obf_3073bf0ff9855ed4 struct {
	__obf_fb75d4e4562132ae
	__obf_5406b11e33b9d1d3 float64
}

func (any *__obf_3073bf0ff9855ed4) Parse() *Iterator {
	return nil
}

func (any *__obf_3073bf0ff9855ed4) ValueType() ValueType {
	return NumberValue
}

func (any *__obf_3073bf0ff9855ed4) MustBeValid() Any {
	return any
}

func (any *__obf_3073bf0ff9855ed4) LastError() error {
	return nil
}

func (any *__obf_3073bf0ff9855ed4) ToBool() bool {
	return any.ToFloat64() != 0
}

func (any *__obf_3073bf0ff9855ed4) ToInt() int {
	return int(any.__obf_5406b11e33b9d1d3)
}

func (any *__obf_3073bf0ff9855ed4) ToInt32() int32 {
	return int32(any.__obf_5406b11e33b9d1d3)
}

func (any *__obf_3073bf0ff9855ed4) ToInt64() int64 {
	return int64(any.__obf_5406b11e33b9d1d3)
}

func (any *__obf_3073bf0ff9855ed4) ToUint() uint {
	if any.__obf_5406b11e33b9d1d3 > 0 {
		return uint(any.__obf_5406b11e33b9d1d3)
	}
	return 0
}

func (any *__obf_3073bf0ff9855ed4) ToUint32() uint32 {
	if any.__obf_5406b11e33b9d1d3 > 0 {
		return uint32(any.__obf_5406b11e33b9d1d3)
	}
	return 0
}

func (any *__obf_3073bf0ff9855ed4) ToUint64() uint64 {
	if any.__obf_5406b11e33b9d1d3 > 0 {
		return uint64(any.__obf_5406b11e33b9d1d3)
	}
	return 0
}

func (any *__obf_3073bf0ff9855ed4) ToFloat32() float32 {
	return float32(any.__obf_5406b11e33b9d1d3)
}

func (any *__obf_3073bf0ff9855ed4) ToFloat64() float64 {
	return any.__obf_5406b11e33b9d1d3
}

func (any *__obf_3073bf0ff9855ed4) ToString() string {
	return strconv.FormatFloat(any.__obf_5406b11e33b9d1d3, 'E', -1, 64)
}

func (any *__obf_3073bf0ff9855ed4) WriteTo(__obf_00fc491caa967a74 *Stream) {
	__obf_00fc491caa967a74.
		WriteFloat64(any.__obf_5406b11e33b9d1d3)
}

func (any *__obf_3073bf0ff9855ed4) GetInterface() any {
	return any.__obf_5406b11e33b9d1d3
}
