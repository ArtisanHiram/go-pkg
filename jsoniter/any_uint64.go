package __obf_703d23ba520c3295

import (
	"strconv"
)

type __obf_c0a5fd87615531ba struct {
	__obf_b2a2e31336581ab8
	__obf_b924538fffe5fd64 uint64
}

func (any *__obf_c0a5fd87615531ba) LastError() error {
	return nil
}

func (any *__obf_c0a5fd87615531ba) ValueType() ValueType {
	return NumberValue
}

func (any *__obf_c0a5fd87615531ba) MustBeValid() Any {
	return any
}

func (any *__obf_c0a5fd87615531ba) ToBool() bool {
	return any.__obf_b924538fffe5fd64 != 0
}

func (any *__obf_c0a5fd87615531ba) ToInt() int {
	return int(any.__obf_b924538fffe5fd64)
}

func (any *__obf_c0a5fd87615531ba) ToInt32() int32 {
	return int32(any.__obf_b924538fffe5fd64)
}

func (any *__obf_c0a5fd87615531ba) ToInt64() int64 {
	return int64(any.__obf_b924538fffe5fd64)
}

func (any *__obf_c0a5fd87615531ba) ToUint() uint {
	return uint(any.__obf_b924538fffe5fd64)
}

func (any *__obf_c0a5fd87615531ba) ToUint32() uint32 {
	return uint32(any.__obf_b924538fffe5fd64)
}

func (any *__obf_c0a5fd87615531ba) ToUint64() uint64 {
	return any.__obf_b924538fffe5fd64
}

func (any *__obf_c0a5fd87615531ba) ToFloat32() float32 {
	return float32(any.__obf_b924538fffe5fd64)
}

func (any *__obf_c0a5fd87615531ba) ToFloat64() float64 {
	return float64(any.__obf_b924538fffe5fd64)
}

func (any *__obf_c0a5fd87615531ba) ToString() string {
	return strconv.FormatUint(any.__obf_b924538fffe5fd64, 10)
}

func (any *__obf_c0a5fd87615531ba) WriteTo(__obf_9a34f62857fb1d1d *Stream) {
	__obf_9a34f62857fb1d1d.
		WriteUint64(any.__obf_b924538fffe5fd64)
}

func (any *__obf_c0a5fd87615531ba) Parse() *Iterator {
	return nil
}

func (any *__obf_c0a5fd87615531ba) GetInterface() any {
	return any.__obf_b924538fffe5fd64
}
