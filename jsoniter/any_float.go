package __obf_030d39f7a8de96c6

import (
	"strconv"
)

type __obf_07c9a175e2e50f0d struct {
	__obf_7834a13a259af712
	__obf_efaf2719b44ad8ac float64
}

func (any *__obf_07c9a175e2e50f0d) Parse() *Iterator {
	return nil
}

func (any *__obf_07c9a175e2e50f0d) ValueType() ValueType {
	return NumberValue
}

func (any *__obf_07c9a175e2e50f0d) MustBeValid() Any {
	return any
}

func (any *__obf_07c9a175e2e50f0d) LastError() error {
	return nil
}

func (any *__obf_07c9a175e2e50f0d) ToBool() bool {
	return any.ToFloat64() != 0
}

func (any *__obf_07c9a175e2e50f0d) ToInt() int {
	return int(any.__obf_efaf2719b44ad8ac)
}

func (any *__obf_07c9a175e2e50f0d) ToInt32() int32 {
	return int32(any.__obf_efaf2719b44ad8ac)
}

func (any *__obf_07c9a175e2e50f0d) ToInt64() int64 {
	return int64(any.__obf_efaf2719b44ad8ac)
}

func (any *__obf_07c9a175e2e50f0d) ToUint() uint {
	if any.__obf_efaf2719b44ad8ac > 0 {
		return uint(any.__obf_efaf2719b44ad8ac)
	}
	return 0
}

func (any *__obf_07c9a175e2e50f0d) ToUint32() uint32 {
	if any.__obf_efaf2719b44ad8ac > 0 {
		return uint32(any.__obf_efaf2719b44ad8ac)
	}
	return 0
}

func (any *__obf_07c9a175e2e50f0d) ToUint64() uint64 {
	if any.__obf_efaf2719b44ad8ac > 0 {
		return uint64(any.__obf_efaf2719b44ad8ac)
	}
	return 0
}

func (any *__obf_07c9a175e2e50f0d) ToFloat32() float32 {
	return float32(any.__obf_efaf2719b44ad8ac)
}

func (any *__obf_07c9a175e2e50f0d) ToFloat64() float64 {
	return any.__obf_efaf2719b44ad8ac
}

func (any *__obf_07c9a175e2e50f0d) ToString() string {
	return strconv.FormatFloat(any.__obf_efaf2719b44ad8ac, 'E', -1, 64)
}

func (any *__obf_07c9a175e2e50f0d) WriteTo(__obf_8a2c51fe22775655 *Stream) {
	__obf_8a2c51fe22775655.
		WriteFloat64(any.__obf_efaf2719b44ad8ac)
}

func (any *__obf_07c9a175e2e50f0d) GetInterface() any {
	return any.__obf_efaf2719b44ad8ac
}
