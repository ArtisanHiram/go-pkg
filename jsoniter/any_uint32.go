package __obf_91620b895eeff9ed

import (
	"strconv"
)

type __obf_935849676e112271 struct {
	__obf_58563642f42f4a04
	__obf_bbfd2ac8618a6f0c uint32
}

func (any *__obf_935849676e112271) LastError() error {
	return nil
}

func (any *__obf_935849676e112271) ValueType() ValueType {
	return NumberValue
}

func (any *__obf_935849676e112271) MustBeValid() Any {
	return any
}

func (any *__obf_935849676e112271) ToBool() bool {
	return any.__obf_bbfd2ac8618a6f0c != 0
}

func (any *__obf_935849676e112271) ToInt() int {
	return int(any.__obf_bbfd2ac8618a6f0c)
}

func (any *__obf_935849676e112271) ToInt32() int32 {
	return int32(any.__obf_bbfd2ac8618a6f0c)
}

func (any *__obf_935849676e112271) ToInt64() int64 {
	return int64(any.__obf_bbfd2ac8618a6f0c)
}

func (any *__obf_935849676e112271) ToUint() uint {
	return uint(any.__obf_bbfd2ac8618a6f0c)
}

func (any *__obf_935849676e112271) ToUint32() uint32 {
	return any.__obf_bbfd2ac8618a6f0c
}

func (any *__obf_935849676e112271) ToUint64() uint64 {
	return uint64(any.__obf_bbfd2ac8618a6f0c)
}

func (any *__obf_935849676e112271) ToFloat32() float32 {
	return float32(any.__obf_bbfd2ac8618a6f0c)
}

func (any *__obf_935849676e112271) ToFloat64() float64 {
	return float64(any.__obf_bbfd2ac8618a6f0c)
}

func (any *__obf_935849676e112271) ToString() string {
	return strconv.FormatInt(int64(any.__obf_bbfd2ac8618a6f0c), 10)
}

func (any *__obf_935849676e112271) WriteTo(__obf_850a7457bb739a32 *Stream) {
	__obf_850a7457bb739a32.
		WriteUint32(any.__obf_bbfd2ac8618a6f0c)
}

func (any *__obf_935849676e112271) Parse() *Iterator {
	return nil
}

func (any *__obf_935849676e112271) GetInterface() any {
	return any.__obf_bbfd2ac8618a6f0c
}
