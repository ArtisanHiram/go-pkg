package __obf_c7b79b12b33d8238

import (
	"strconv"
)

type __obf_aeafd238348bf351 struct {
	__obf_3a44d9c5204c2e12
	__obf_35accd096612ebe4 uint64
}

func (any *__obf_aeafd238348bf351) LastError() error {
	return nil
}

func (any *__obf_aeafd238348bf351) ValueType() ValueType {
	return NumberValue
}

func (any *__obf_aeafd238348bf351) MustBeValid() Any {
	return any
}

func (any *__obf_aeafd238348bf351) ToBool() bool {
	return any.__obf_35accd096612ebe4 != 0
}

func (any *__obf_aeafd238348bf351) ToInt() int {
	return int(any.__obf_35accd096612ebe4)
}

func (any *__obf_aeafd238348bf351) ToInt32() int32 {
	return int32(any.__obf_35accd096612ebe4)
}

func (any *__obf_aeafd238348bf351) ToInt64() int64 {
	return int64(any.__obf_35accd096612ebe4)
}

func (any *__obf_aeafd238348bf351) ToUint() uint {
	return uint(any.__obf_35accd096612ebe4)
}

func (any *__obf_aeafd238348bf351) ToUint32() uint32 {
	return uint32(any.__obf_35accd096612ebe4)
}

func (any *__obf_aeafd238348bf351) ToUint64() uint64 {
	return any.__obf_35accd096612ebe4
}

func (any *__obf_aeafd238348bf351) ToFloat32() float32 {
	return float32(any.__obf_35accd096612ebe4)
}

func (any *__obf_aeafd238348bf351) ToFloat64() float64 {
	return float64(any.__obf_35accd096612ebe4)
}

func (any *__obf_aeafd238348bf351) ToString() string {
	return strconv.FormatUint(any.__obf_35accd096612ebe4, 10)
}

func (any *__obf_aeafd238348bf351) WriteTo(__obf_d8f50bcfe87d8b47 *Stream) {
	__obf_d8f50bcfe87d8b47.
		WriteUint64(any.__obf_35accd096612ebe4)
}

func (any *__obf_aeafd238348bf351) Parse() *Iterator {
	return nil
}

func (any *__obf_aeafd238348bf351) GetInterface() any {
	return any.__obf_35accd096612ebe4
}
