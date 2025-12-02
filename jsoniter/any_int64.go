package __obf_c7b79b12b33d8238

import (
	"strconv"
)

type __obf_064ca4e3fc05241a struct {
	__obf_3a44d9c5204c2e12
	__obf_35accd096612ebe4 int64
}

func (any *__obf_064ca4e3fc05241a) LastError() error {
	return nil
}

func (any *__obf_064ca4e3fc05241a) ValueType() ValueType {
	return NumberValue
}

func (any *__obf_064ca4e3fc05241a) MustBeValid() Any {
	return any
}

func (any *__obf_064ca4e3fc05241a) ToBool() bool {
	return any.__obf_35accd096612ebe4 != 0
}

func (any *__obf_064ca4e3fc05241a) ToInt() int {
	return int(any.__obf_35accd096612ebe4)
}

func (any *__obf_064ca4e3fc05241a) ToInt32() int32 {
	return int32(any.__obf_35accd096612ebe4)
}

func (any *__obf_064ca4e3fc05241a) ToInt64() int64 {
	return any.__obf_35accd096612ebe4
}

func (any *__obf_064ca4e3fc05241a) ToUint() uint {
	return uint(any.__obf_35accd096612ebe4)
}

func (any *__obf_064ca4e3fc05241a) ToUint32() uint32 {
	return uint32(any.__obf_35accd096612ebe4)
}

func (any *__obf_064ca4e3fc05241a) ToUint64() uint64 {
	return uint64(any.__obf_35accd096612ebe4)
}

func (any *__obf_064ca4e3fc05241a) ToFloat32() float32 {
	return float32(any.__obf_35accd096612ebe4)
}

func (any *__obf_064ca4e3fc05241a) ToFloat64() float64 {
	return float64(any.__obf_35accd096612ebe4)
}

func (any *__obf_064ca4e3fc05241a) ToString() string {
	return strconv.FormatInt(any.__obf_35accd096612ebe4, 10)
}

func (any *__obf_064ca4e3fc05241a) WriteTo(__obf_d8f50bcfe87d8b47 *Stream) {
	__obf_d8f50bcfe87d8b47.
		WriteInt64(any.__obf_35accd096612ebe4)
}

func (any *__obf_064ca4e3fc05241a) Parse() *Iterator {
	return nil
}

func (any *__obf_064ca4e3fc05241a) GetInterface() any {
	return any.__obf_35accd096612ebe4
}
