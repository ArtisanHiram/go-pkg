package __obf_c7b79b12b33d8238

import (
	"strconv"
)

type __obf_67b34d104eef23fb struct {
	__obf_3a44d9c5204c2e12
	__obf_35accd096612ebe4 int32
}

func (any *__obf_67b34d104eef23fb) LastError() error {
	return nil
}

func (any *__obf_67b34d104eef23fb) ValueType() ValueType {
	return NumberValue
}

func (any *__obf_67b34d104eef23fb) MustBeValid() Any {
	return any
}

func (any *__obf_67b34d104eef23fb) ToBool() bool {
	return any.__obf_35accd096612ebe4 != 0
}

func (any *__obf_67b34d104eef23fb) ToInt() int {
	return int(any.__obf_35accd096612ebe4)
}

func (any *__obf_67b34d104eef23fb) ToInt32() int32 {
	return any.__obf_35accd096612ebe4
}

func (any *__obf_67b34d104eef23fb) ToInt64() int64 {
	return int64(any.__obf_35accd096612ebe4)
}

func (any *__obf_67b34d104eef23fb) ToUint() uint {
	return uint(any.__obf_35accd096612ebe4)
}

func (any *__obf_67b34d104eef23fb) ToUint32() uint32 {
	return uint32(any.__obf_35accd096612ebe4)
}

func (any *__obf_67b34d104eef23fb) ToUint64() uint64 {
	return uint64(any.__obf_35accd096612ebe4)
}

func (any *__obf_67b34d104eef23fb) ToFloat32() float32 {
	return float32(any.__obf_35accd096612ebe4)
}

func (any *__obf_67b34d104eef23fb) ToFloat64() float64 {
	return float64(any.__obf_35accd096612ebe4)
}

func (any *__obf_67b34d104eef23fb) ToString() string {
	return strconv.FormatInt(int64(any.__obf_35accd096612ebe4), 10)
}

func (any *__obf_67b34d104eef23fb) WriteTo(__obf_d8f50bcfe87d8b47 *Stream) {
	__obf_d8f50bcfe87d8b47.
		WriteInt32(any.__obf_35accd096612ebe4)
}

func (any *__obf_67b34d104eef23fb) Parse() *Iterator {
	return nil
}

func (any *__obf_67b34d104eef23fb) GetInterface() any {
	return any.__obf_35accd096612ebe4
}
