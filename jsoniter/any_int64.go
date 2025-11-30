package __obf_703d23ba520c3295

import (
	"strconv"
)

type __obf_1d6d72beae0ea33d struct {
	__obf_b2a2e31336581ab8
	__obf_b924538fffe5fd64 int64
}

func (any *__obf_1d6d72beae0ea33d) LastError() error {
	return nil
}

func (any *__obf_1d6d72beae0ea33d) ValueType() ValueType {
	return NumberValue
}

func (any *__obf_1d6d72beae0ea33d) MustBeValid() Any {
	return any
}

func (any *__obf_1d6d72beae0ea33d) ToBool() bool {
	return any.__obf_b924538fffe5fd64 != 0
}

func (any *__obf_1d6d72beae0ea33d) ToInt() int {
	return int(any.__obf_b924538fffe5fd64)
}

func (any *__obf_1d6d72beae0ea33d) ToInt32() int32 {
	return int32(any.__obf_b924538fffe5fd64)
}

func (any *__obf_1d6d72beae0ea33d) ToInt64() int64 {
	return any.__obf_b924538fffe5fd64
}

func (any *__obf_1d6d72beae0ea33d) ToUint() uint {
	return uint(any.__obf_b924538fffe5fd64)
}

func (any *__obf_1d6d72beae0ea33d) ToUint32() uint32 {
	return uint32(any.__obf_b924538fffe5fd64)
}

func (any *__obf_1d6d72beae0ea33d) ToUint64() uint64 {
	return uint64(any.__obf_b924538fffe5fd64)
}

func (any *__obf_1d6d72beae0ea33d) ToFloat32() float32 {
	return float32(any.__obf_b924538fffe5fd64)
}

func (any *__obf_1d6d72beae0ea33d) ToFloat64() float64 {
	return float64(any.__obf_b924538fffe5fd64)
}

func (any *__obf_1d6d72beae0ea33d) ToString() string {
	return strconv.FormatInt(any.__obf_b924538fffe5fd64, 10)
}

func (any *__obf_1d6d72beae0ea33d) WriteTo(__obf_9a34f62857fb1d1d *Stream) {
	__obf_9a34f62857fb1d1d.
		WriteInt64(any.__obf_b924538fffe5fd64)
}

func (any *__obf_1d6d72beae0ea33d) Parse() *Iterator {
	return nil
}

func (any *__obf_1d6d72beae0ea33d) GetInterface() any {
	return any.__obf_b924538fffe5fd64
}
