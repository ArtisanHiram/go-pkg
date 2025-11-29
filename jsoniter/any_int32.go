package __obf_91620b895eeff9ed

import (
	"strconv"
)

type __obf_aad087210be33f6a struct {
	__obf_58563642f42f4a04
	__obf_bbfd2ac8618a6f0c int32
}

func (any *__obf_aad087210be33f6a) LastError() error {
	return nil
}

func (any *__obf_aad087210be33f6a) ValueType() ValueType {
	return NumberValue
}

func (any *__obf_aad087210be33f6a) MustBeValid() Any {
	return any
}

func (any *__obf_aad087210be33f6a) ToBool() bool {
	return any.__obf_bbfd2ac8618a6f0c != 0
}

func (any *__obf_aad087210be33f6a) ToInt() int {
	return int(any.__obf_bbfd2ac8618a6f0c)
}

func (any *__obf_aad087210be33f6a) ToInt32() int32 {
	return any.__obf_bbfd2ac8618a6f0c
}

func (any *__obf_aad087210be33f6a) ToInt64() int64 {
	return int64(any.__obf_bbfd2ac8618a6f0c)
}

func (any *__obf_aad087210be33f6a) ToUint() uint {
	return uint(any.__obf_bbfd2ac8618a6f0c)
}

func (any *__obf_aad087210be33f6a) ToUint32() uint32 {
	return uint32(any.__obf_bbfd2ac8618a6f0c)
}

func (any *__obf_aad087210be33f6a) ToUint64() uint64 {
	return uint64(any.__obf_bbfd2ac8618a6f0c)
}

func (any *__obf_aad087210be33f6a) ToFloat32() float32 {
	return float32(any.__obf_bbfd2ac8618a6f0c)
}

func (any *__obf_aad087210be33f6a) ToFloat64() float64 {
	return float64(any.__obf_bbfd2ac8618a6f0c)
}

func (any *__obf_aad087210be33f6a) ToString() string {
	return strconv.FormatInt(int64(any.__obf_bbfd2ac8618a6f0c), 10)
}

func (any *__obf_aad087210be33f6a) WriteTo(__obf_850a7457bb739a32 *Stream) {
	__obf_850a7457bb739a32.
		WriteInt32(any.__obf_bbfd2ac8618a6f0c)
}

func (any *__obf_aad087210be33f6a) Parse() *Iterator {
	return nil
}

func (any *__obf_aad087210be33f6a) GetInterface() any {
	return any.__obf_bbfd2ac8618a6f0c
}
