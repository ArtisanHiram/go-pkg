package __obf_703d23ba520c3295

import (
	"strconv"
)

type __obf_ba7fcab5df6006b8 struct {
	__obf_b2a2e31336581ab8
	__obf_b924538fffe5fd64 float64
}

func (any *__obf_ba7fcab5df6006b8) Parse() *Iterator {
	return nil
}

func (any *__obf_ba7fcab5df6006b8) ValueType() ValueType {
	return NumberValue
}

func (any *__obf_ba7fcab5df6006b8) MustBeValid() Any {
	return any
}

func (any *__obf_ba7fcab5df6006b8) LastError() error {
	return nil
}

func (any *__obf_ba7fcab5df6006b8) ToBool() bool {
	return any.ToFloat64() != 0
}

func (any *__obf_ba7fcab5df6006b8) ToInt() int {
	return int(any.__obf_b924538fffe5fd64)
}

func (any *__obf_ba7fcab5df6006b8) ToInt32() int32 {
	return int32(any.__obf_b924538fffe5fd64)
}

func (any *__obf_ba7fcab5df6006b8) ToInt64() int64 {
	return int64(any.__obf_b924538fffe5fd64)
}

func (any *__obf_ba7fcab5df6006b8) ToUint() uint {
	if any.__obf_b924538fffe5fd64 > 0 {
		return uint(any.__obf_b924538fffe5fd64)
	}
	return 0
}

func (any *__obf_ba7fcab5df6006b8) ToUint32() uint32 {
	if any.__obf_b924538fffe5fd64 > 0 {
		return uint32(any.__obf_b924538fffe5fd64)
	}
	return 0
}

func (any *__obf_ba7fcab5df6006b8) ToUint64() uint64 {
	if any.__obf_b924538fffe5fd64 > 0 {
		return uint64(any.__obf_b924538fffe5fd64)
	}
	return 0
}

func (any *__obf_ba7fcab5df6006b8) ToFloat32() float32 {
	return float32(any.__obf_b924538fffe5fd64)
}

func (any *__obf_ba7fcab5df6006b8) ToFloat64() float64 {
	return any.__obf_b924538fffe5fd64
}

func (any *__obf_ba7fcab5df6006b8) ToString() string {
	return strconv.FormatFloat(any.__obf_b924538fffe5fd64, 'E', -1, 64)
}

func (any *__obf_ba7fcab5df6006b8) WriteTo(__obf_9a34f62857fb1d1d *Stream) {
	__obf_9a34f62857fb1d1d.
		WriteFloat64(any.__obf_b924538fffe5fd64)
}

func (any *__obf_ba7fcab5df6006b8) GetInterface() any {
	return any.__obf_b924538fffe5fd64
}
