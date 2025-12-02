package __obf_c7b79b12b33d8238

import (
	"strconv"
)

type __obf_fc7eb09dc621f051 struct {
	__obf_3a44d9c5204c2e12
	__obf_35accd096612ebe4 float64
}

func (any *__obf_fc7eb09dc621f051) Parse() *Iterator {
	return nil
}

func (any *__obf_fc7eb09dc621f051) ValueType() ValueType {
	return NumberValue
}

func (any *__obf_fc7eb09dc621f051) MustBeValid() Any {
	return any
}

func (any *__obf_fc7eb09dc621f051) LastError() error {
	return nil
}

func (any *__obf_fc7eb09dc621f051) ToBool() bool {
	return any.ToFloat64() != 0
}

func (any *__obf_fc7eb09dc621f051) ToInt() int {
	return int(any.__obf_35accd096612ebe4)
}

func (any *__obf_fc7eb09dc621f051) ToInt32() int32 {
	return int32(any.__obf_35accd096612ebe4)
}

func (any *__obf_fc7eb09dc621f051) ToInt64() int64 {
	return int64(any.__obf_35accd096612ebe4)
}

func (any *__obf_fc7eb09dc621f051) ToUint() uint {
	if any.__obf_35accd096612ebe4 > 0 {
		return uint(any.__obf_35accd096612ebe4)
	}
	return 0
}

func (any *__obf_fc7eb09dc621f051) ToUint32() uint32 {
	if any.__obf_35accd096612ebe4 > 0 {
		return uint32(any.__obf_35accd096612ebe4)
	}
	return 0
}

func (any *__obf_fc7eb09dc621f051) ToUint64() uint64 {
	if any.__obf_35accd096612ebe4 > 0 {
		return uint64(any.__obf_35accd096612ebe4)
	}
	return 0
}

func (any *__obf_fc7eb09dc621f051) ToFloat32() float32 {
	return float32(any.__obf_35accd096612ebe4)
}

func (any *__obf_fc7eb09dc621f051) ToFloat64() float64 {
	return any.__obf_35accd096612ebe4
}

func (any *__obf_fc7eb09dc621f051) ToString() string {
	return strconv.FormatFloat(any.__obf_35accd096612ebe4, 'E', -1, 64)
}

func (any *__obf_fc7eb09dc621f051) WriteTo(__obf_d8f50bcfe87d8b47 *Stream) {
	__obf_d8f50bcfe87d8b47.
		WriteFloat64(any.__obf_35accd096612ebe4)
}

func (any *__obf_fc7eb09dc621f051) GetInterface() any {
	return any.__obf_35accd096612ebe4
}
