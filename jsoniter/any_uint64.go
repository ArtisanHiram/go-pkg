package __obf_030d39f7a8de96c6

import (
	"strconv"
)

type __obf_4bd84d2e4678d344 struct {
	__obf_7834a13a259af712
	__obf_efaf2719b44ad8ac uint64
}

func (any *__obf_4bd84d2e4678d344) LastError() error {
	return nil
}

func (any *__obf_4bd84d2e4678d344) ValueType() ValueType {
	return NumberValue
}

func (any *__obf_4bd84d2e4678d344) MustBeValid() Any {
	return any
}

func (any *__obf_4bd84d2e4678d344) ToBool() bool {
	return any.__obf_efaf2719b44ad8ac != 0
}

func (any *__obf_4bd84d2e4678d344) ToInt() int {
	return int(any.__obf_efaf2719b44ad8ac)
}

func (any *__obf_4bd84d2e4678d344) ToInt32() int32 {
	return int32(any.__obf_efaf2719b44ad8ac)
}

func (any *__obf_4bd84d2e4678d344) ToInt64() int64 {
	return int64(any.__obf_efaf2719b44ad8ac)
}

func (any *__obf_4bd84d2e4678d344) ToUint() uint {
	return uint(any.__obf_efaf2719b44ad8ac)
}

func (any *__obf_4bd84d2e4678d344) ToUint32() uint32 {
	return uint32(any.__obf_efaf2719b44ad8ac)
}

func (any *__obf_4bd84d2e4678d344) ToUint64() uint64 {
	return any.__obf_efaf2719b44ad8ac
}

func (any *__obf_4bd84d2e4678d344) ToFloat32() float32 {
	return float32(any.__obf_efaf2719b44ad8ac)
}

func (any *__obf_4bd84d2e4678d344) ToFloat64() float64 {
	return float64(any.__obf_efaf2719b44ad8ac)
}

func (any *__obf_4bd84d2e4678d344) ToString() string {
	return strconv.FormatUint(any.__obf_efaf2719b44ad8ac, 10)
}

func (any *__obf_4bd84d2e4678d344) WriteTo(__obf_8a2c51fe22775655 *Stream) {
	__obf_8a2c51fe22775655.
		WriteUint64(any.__obf_efaf2719b44ad8ac)
}

func (any *__obf_4bd84d2e4678d344) Parse() *Iterator {
	return nil
}

func (any *__obf_4bd84d2e4678d344) GetInterface() any {
	return any.__obf_efaf2719b44ad8ac
}
