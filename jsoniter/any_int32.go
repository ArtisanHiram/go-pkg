package __obf_703d23ba520c3295

import (
	"strconv"
)

type __obf_592e58127a384a3b struct {
	__obf_b2a2e31336581ab8
	__obf_b924538fffe5fd64 int32
}

func (any *__obf_592e58127a384a3b) LastError() error {
	return nil
}

func (any *__obf_592e58127a384a3b) ValueType() ValueType {
	return NumberValue
}

func (any *__obf_592e58127a384a3b) MustBeValid() Any {
	return any
}

func (any *__obf_592e58127a384a3b) ToBool() bool {
	return any.__obf_b924538fffe5fd64 != 0
}

func (any *__obf_592e58127a384a3b) ToInt() int {
	return int(any.__obf_b924538fffe5fd64)
}

func (any *__obf_592e58127a384a3b) ToInt32() int32 {
	return any.__obf_b924538fffe5fd64
}

func (any *__obf_592e58127a384a3b) ToInt64() int64 {
	return int64(any.__obf_b924538fffe5fd64)
}

func (any *__obf_592e58127a384a3b) ToUint() uint {
	return uint(any.__obf_b924538fffe5fd64)
}

func (any *__obf_592e58127a384a3b) ToUint32() uint32 {
	return uint32(any.__obf_b924538fffe5fd64)
}

func (any *__obf_592e58127a384a3b) ToUint64() uint64 {
	return uint64(any.__obf_b924538fffe5fd64)
}

func (any *__obf_592e58127a384a3b) ToFloat32() float32 {
	return float32(any.__obf_b924538fffe5fd64)
}

func (any *__obf_592e58127a384a3b) ToFloat64() float64 {
	return float64(any.__obf_b924538fffe5fd64)
}

func (any *__obf_592e58127a384a3b) ToString() string {
	return strconv.FormatInt(int64(any.__obf_b924538fffe5fd64), 10)
}

func (any *__obf_592e58127a384a3b) WriteTo(__obf_9a34f62857fb1d1d *Stream) {
	__obf_9a34f62857fb1d1d.
		WriteInt32(any.__obf_b924538fffe5fd64)
}

func (any *__obf_592e58127a384a3b) Parse() *Iterator {
	return nil
}

func (any *__obf_592e58127a384a3b) GetInterface() any {
	return any.__obf_b924538fffe5fd64
}
