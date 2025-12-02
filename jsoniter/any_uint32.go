package __obf_c7b79b12b33d8238

import (
	"strconv"
)

type __obf_e407dbf923c34658 struct {
	__obf_3a44d9c5204c2e12
	__obf_35accd096612ebe4 uint32
}

func (any *__obf_e407dbf923c34658) LastError() error {
	return nil
}

func (any *__obf_e407dbf923c34658) ValueType() ValueType {
	return NumberValue
}

func (any *__obf_e407dbf923c34658) MustBeValid() Any {
	return any
}

func (any *__obf_e407dbf923c34658) ToBool() bool {
	return any.__obf_35accd096612ebe4 != 0
}

func (any *__obf_e407dbf923c34658) ToInt() int {
	return int(any.__obf_35accd096612ebe4)
}

func (any *__obf_e407dbf923c34658) ToInt32() int32 {
	return int32(any.__obf_35accd096612ebe4)
}

func (any *__obf_e407dbf923c34658) ToInt64() int64 {
	return int64(any.__obf_35accd096612ebe4)
}

func (any *__obf_e407dbf923c34658) ToUint() uint {
	return uint(any.__obf_35accd096612ebe4)
}

func (any *__obf_e407dbf923c34658) ToUint32() uint32 {
	return any.__obf_35accd096612ebe4
}

func (any *__obf_e407dbf923c34658) ToUint64() uint64 {
	return uint64(any.__obf_35accd096612ebe4)
}

func (any *__obf_e407dbf923c34658) ToFloat32() float32 {
	return float32(any.__obf_35accd096612ebe4)
}

func (any *__obf_e407dbf923c34658) ToFloat64() float64 {
	return float64(any.__obf_35accd096612ebe4)
}

func (any *__obf_e407dbf923c34658) ToString() string {
	return strconv.FormatInt(int64(any.__obf_35accd096612ebe4), 10)
}

func (any *__obf_e407dbf923c34658) WriteTo(__obf_d8f50bcfe87d8b47 *Stream) {
	__obf_d8f50bcfe87d8b47.
		WriteUint32(any.__obf_35accd096612ebe4)
}

func (any *__obf_e407dbf923c34658) Parse() *Iterator {
	return nil
}

func (any *__obf_e407dbf923c34658) GetInterface() any {
	return any.__obf_35accd096612ebe4
}
