package __obf_91620b895eeff9ed

import (
	"strconv"
)

type __obf_d8d434196b35220b struct {
	__obf_58563642f42f4a04
	__obf_bbfd2ac8618a6f0c uint64
}

func (any *__obf_d8d434196b35220b) LastError() error {
	return nil
}

func (any *__obf_d8d434196b35220b) ValueType() ValueType {
	return NumberValue
}

func (any *__obf_d8d434196b35220b) MustBeValid() Any {
	return any
}

func (any *__obf_d8d434196b35220b) ToBool() bool {
	return any.__obf_bbfd2ac8618a6f0c != 0
}

func (any *__obf_d8d434196b35220b) ToInt() int {
	return int(any.__obf_bbfd2ac8618a6f0c)
}

func (any *__obf_d8d434196b35220b) ToInt32() int32 {
	return int32(any.__obf_bbfd2ac8618a6f0c)
}

func (any *__obf_d8d434196b35220b) ToInt64() int64 {
	return int64(any.__obf_bbfd2ac8618a6f0c)
}

func (any *__obf_d8d434196b35220b) ToUint() uint {
	return uint(any.__obf_bbfd2ac8618a6f0c)
}

func (any *__obf_d8d434196b35220b) ToUint32() uint32 {
	return uint32(any.__obf_bbfd2ac8618a6f0c)
}

func (any *__obf_d8d434196b35220b) ToUint64() uint64 {
	return any.__obf_bbfd2ac8618a6f0c
}

func (any *__obf_d8d434196b35220b) ToFloat32() float32 {
	return float32(any.__obf_bbfd2ac8618a6f0c)
}

func (any *__obf_d8d434196b35220b) ToFloat64() float64 {
	return float64(any.__obf_bbfd2ac8618a6f0c)
}

func (any *__obf_d8d434196b35220b) ToString() string {
	return strconv.FormatUint(any.__obf_bbfd2ac8618a6f0c, 10)
}

func (any *__obf_d8d434196b35220b) WriteTo(__obf_850a7457bb739a32 *Stream) {
	__obf_850a7457bb739a32.
		WriteUint64(any.__obf_bbfd2ac8618a6f0c)
}

func (any *__obf_d8d434196b35220b) Parse() *Iterator {
	return nil
}

func (any *__obf_d8d434196b35220b) GetInterface() any {
	return any.__obf_bbfd2ac8618a6f0c
}
