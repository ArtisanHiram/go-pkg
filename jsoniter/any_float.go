package __obf_91620b895eeff9ed

import (
	"strconv"
)

type __obf_36d93f8a7e11e503 struct {
	__obf_58563642f42f4a04
	__obf_bbfd2ac8618a6f0c float64
}

func (any *__obf_36d93f8a7e11e503) Parse() *Iterator {
	return nil
}

func (any *__obf_36d93f8a7e11e503) ValueType() ValueType {
	return NumberValue
}

func (any *__obf_36d93f8a7e11e503) MustBeValid() Any {
	return any
}

func (any *__obf_36d93f8a7e11e503) LastError() error {
	return nil
}

func (any *__obf_36d93f8a7e11e503) ToBool() bool {
	return any.ToFloat64() != 0
}

func (any *__obf_36d93f8a7e11e503) ToInt() int {
	return int(any.__obf_bbfd2ac8618a6f0c)
}

func (any *__obf_36d93f8a7e11e503) ToInt32() int32 {
	return int32(any.__obf_bbfd2ac8618a6f0c)
}

func (any *__obf_36d93f8a7e11e503) ToInt64() int64 {
	return int64(any.__obf_bbfd2ac8618a6f0c)
}

func (any *__obf_36d93f8a7e11e503) ToUint() uint {
	if any.__obf_bbfd2ac8618a6f0c > 0 {
		return uint(any.__obf_bbfd2ac8618a6f0c)
	}
	return 0
}

func (any *__obf_36d93f8a7e11e503) ToUint32() uint32 {
	if any.__obf_bbfd2ac8618a6f0c > 0 {
		return uint32(any.__obf_bbfd2ac8618a6f0c)
	}
	return 0
}

func (any *__obf_36d93f8a7e11e503) ToUint64() uint64 {
	if any.__obf_bbfd2ac8618a6f0c > 0 {
		return uint64(any.__obf_bbfd2ac8618a6f0c)
	}
	return 0
}

func (any *__obf_36d93f8a7e11e503) ToFloat32() float32 {
	return float32(any.__obf_bbfd2ac8618a6f0c)
}

func (any *__obf_36d93f8a7e11e503) ToFloat64() float64 {
	return any.__obf_bbfd2ac8618a6f0c
}

func (any *__obf_36d93f8a7e11e503) ToString() string {
	return strconv.FormatFloat(any.__obf_bbfd2ac8618a6f0c, 'E', -1, 64)
}

func (any *__obf_36d93f8a7e11e503) WriteTo(__obf_850a7457bb739a32 *Stream) {
	__obf_850a7457bb739a32.
		WriteFloat64(any.__obf_bbfd2ac8618a6f0c)
}

func (any *__obf_36d93f8a7e11e503) GetInterface() any {
	return any.__obf_bbfd2ac8618a6f0c
}
