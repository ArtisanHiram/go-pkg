package __obf_030d39f7a8de96c6

import (
	"strconv"
)

type __obf_99e4188e77cbabaa struct {
	__obf_7834a13a259af712
	__obf_efaf2719b44ad8ac uint32
}

func (any *__obf_99e4188e77cbabaa) LastError() error {
	return nil
}

func (any *__obf_99e4188e77cbabaa) ValueType() ValueType {
	return NumberValue
}

func (any *__obf_99e4188e77cbabaa) MustBeValid() Any {
	return any
}

func (any *__obf_99e4188e77cbabaa) ToBool() bool {
	return any.__obf_efaf2719b44ad8ac != 0
}

func (any *__obf_99e4188e77cbabaa) ToInt() int {
	return int(any.__obf_efaf2719b44ad8ac)
}

func (any *__obf_99e4188e77cbabaa) ToInt32() int32 {
	return int32(any.__obf_efaf2719b44ad8ac)
}

func (any *__obf_99e4188e77cbabaa) ToInt64() int64 {
	return int64(any.__obf_efaf2719b44ad8ac)
}

func (any *__obf_99e4188e77cbabaa) ToUint() uint {
	return uint(any.__obf_efaf2719b44ad8ac)
}

func (any *__obf_99e4188e77cbabaa) ToUint32() uint32 {
	return any.__obf_efaf2719b44ad8ac
}

func (any *__obf_99e4188e77cbabaa) ToUint64() uint64 {
	return uint64(any.__obf_efaf2719b44ad8ac)
}

func (any *__obf_99e4188e77cbabaa) ToFloat32() float32 {
	return float32(any.__obf_efaf2719b44ad8ac)
}

func (any *__obf_99e4188e77cbabaa) ToFloat64() float64 {
	return float64(any.__obf_efaf2719b44ad8ac)
}

func (any *__obf_99e4188e77cbabaa) ToString() string {
	return strconv.FormatInt(int64(any.__obf_efaf2719b44ad8ac), 10)
}

func (any *__obf_99e4188e77cbabaa) WriteTo(__obf_8a2c51fe22775655 *Stream) {
	__obf_8a2c51fe22775655.
		WriteUint32(any.__obf_efaf2719b44ad8ac)
}

func (any *__obf_99e4188e77cbabaa) Parse() *Iterator {
	return nil
}

func (any *__obf_99e4188e77cbabaa) GetInterface() any {
	return any.__obf_efaf2719b44ad8ac
}
