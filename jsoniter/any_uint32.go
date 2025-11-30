package __obf_703d23ba520c3295

import (
	"strconv"
)

type __obf_71f4ab36a81c97f3 struct {
	__obf_b2a2e31336581ab8
	__obf_b924538fffe5fd64 uint32
}

func (any *__obf_71f4ab36a81c97f3) LastError() error {
	return nil
}

func (any *__obf_71f4ab36a81c97f3) ValueType() ValueType {
	return NumberValue
}

func (any *__obf_71f4ab36a81c97f3) MustBeValid() Any {
	return any
}

func (any *__obf_71f4ab36a81c97f3) ToBool() bool {
	return any.__obf_b924538fffe5fd64 != 0
}

func (any *__obf_71f4ab36a81c97f3) ToInt() int {
	return int(any.__obf_b924538fffe5fd64)
}

func (any *__obf_71f4ab36a81c97f3) ToInt32() int32 {
	return int32(any.__obf_b924538fffe5fd64)
}

func (any *__obf_71f4ab36a81c97f3) ToInt64() int64 {
	return int64(any.__obf_b924538fffe5fd64)
}

func (any *__obf_71f4ab36a81c97f3) ToUint() uint {
	return uint(any.__obf_b924538fffe5fd64)
}

func (any *__obf_71f4ab36a81c97f3) ToUint32() uint32 {
	return any.__obf_b924538fffe5fd64
}

func (any *__obf_71f4ab36a81c97f3) ToUint64() uint64 {
	return uint64(any.__obf_b924538fffe5fd64)
}

func (any *__obf_71f4ab36a81c97f3) ToFloat32() float32 {
	return float32(any.__obf_b924538fffe5fd64)
}

func (any *__obf_71f4ab36a81c97f3) ToFloat64() float64 {
	return float64(any.__obf_b924538fffe5fd64)
}

func (any *__obf_71f4ab36a81c97f3) ToString() string {
	return strconv.FormatInt(int64(any.__obf_b924538fffe5fd64), 10)
}

func (any *__obf_71f4ab36a81c97f3) WriteTo(__obf_9a34f62857fb1d1d *Stream) {
	__obf_9a34f62857fb1d1d.
		WriteUint32(any.__obf_b924538fffe5fd64)
}

func (any *__obf_71f4ab36a81c97f3) Parse() *Iterator {
	return nil
}

func (any *__obf_71f4ab36a81c97f3) GetInterface() any {
	return any.__obf_b924538fffe5fd64
}
