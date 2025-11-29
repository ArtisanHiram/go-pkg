package __obf_91620b895eeff9ed

import (
	"strconv"
)

type __obf_2a9d95e27a8da146 struct {
	__obf_58563642f42f4a04
	__obf_bbfd2ac8618a6f0c int64
}

func (any *__obf_2a9d95e27a8da146) LastError() error {
	return nil
}

func (any *__obf_2a9d95e27a8da146) ValueType() ValueType {
	return NumberValue
}

func (any *__obf_2a9d95e27a8da146) MustBeValid() Any {
	return any
}

func (any *__obf_2a9d95e27a8da146) ToBool() bool {
	return any.__obf_bbfd2ac8618a6f0c != 0
}

func (any *__obf_2a9d95e27a8da146) ToInt() int {
	return int(any.__obf_bbfd2ac8618a6f0c)
}

func (any *__obf_2a9d95e27a8da146) ToInt32() int32 {
	return int32(any.__obf_bbfd2ac8618a6f0c)
}

func (any *__obf_2a9d95e27a8da146) ToInt64() int64 {
	return any.__obf_bbfd2ac8618a6f0c
}

func (any *__obf_2a9d95e27a8da146) ToUint() uint {
	return uint(any.__obf_bbfd2ac8618a6f0c)
}

func (any *__obf_2a9d95e27a8da146) ToUint32() uint32 {
	return uint32(any.__obf_bbfd2ac8618a6f0c)
}

func (any *__obf_2a9d95e27a8da146) ToUint64() uint64 {
	return uint64(any.__obf_bbfd2ac8618a6f0c)
}

func (any *__obf_2a9d95e27a8da146) ToFloat32() float32 {
	return float32(any.__obf_bbfd2ac8618a6f0c)
}

func (any *__obf_2a9d95e27a8da146) ToFloat64() float64 {
	return float64(any.__obf_bbfd2ac8618a6f0c)
}

func (any *__obf_2a9d95e27a8da146) ToString() string {
	return strconv.FormatInt(any.__obf_bbfd2ac8618a6f0c, 10)
}

func (any *__obf_2a9d95e27a8da146) WriteTo(__obf_850a7457bb739a32 *Stream) {
	__obf_850a7457bb739a32.
		WriteInt64(any.__obf_bbfd2ac8618a6f0c)
}

func (any *__obf_2a9d95e27a8da146) Parse() *Iterator {
	return nil
}

func (any *__obf_2a9d95e27a8da146) GetInterface() any {
	return any.__obf_bbfd2ac8618a6f0c
}
