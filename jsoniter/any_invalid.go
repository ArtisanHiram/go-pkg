package __obf_703d23ba520c3295

import "fmt"

type __obf_5d650408c7f7f591 struct {
	__obf_b2a2e31336581ab8
	__obf_e56742d6e52f642e error
}

func __obf_b82811e6c38b9e10(__obf_267bdf615cb1b310 []any) *__obf_5d650408c7f7f591 {
	return &__obf_5d650408c7f7f591{__obf_b2a2e31336581ab8{}, fmt.Errorf("%v not found", __obf_267bdf615cb1b310)}
}

func (any *__obf_5d650408c7f7f591,) LastError() error {
	return any.__obf_e56742d6e52f642e
}

func (any *__obf_5d650408c7f7f591,) ValueType() ValueType {
	return InvalidValue
}

func (any *__obf_5d650408c7f7f591,) MustBeValid() Any {
	panic(any.__obf_e56742d6e52f642e)
}

func (any *__obf_5d650408c7f7f591,) ToBool() bool {
	return false
}

func (any *__obf_5d650408c7f7f591,) ToInt() int {
	return 0
}

func (any *__obf_5d650408c7f7f591,) ToInt32() int32 {
	return 0
}

func (any *__obf_5d650408c7f7f591,) ToInt64() int64 {
	return 0
}

func (any *__obf_5d650408c7f7f591,) ToUint() uint {
	return 0
}

func (any *__obf_5d650408c7f7f591,) ToUint32() uint32 {
	return 0
}

func (any *__obf_5d650408c7f7f591,) ToUint64() uint64 {
	return 0
}

func (any *__obf_5d650408c7f7f591,) ToFloat32() float32 {
	return 0
}

func (any *__obf_5d650408c7f7f591,) ToFloat64() float64 {
	return 0
}

func (any *__obf_5d650408c7f7f591,) ToString() string {
	return ""
}

func (any *__obf_5d650408c7f7f591,) WriteTo(__obf_9a34f62857fb1d1d *Stream) {
}

func (any *__obf_5d650408c7f7f591,) Get(__obf_267bdf615cb1b310 ...any) Any {
	if any.__obf_e56742d6e52f642e == nil {
		return &__obf_5d650408c7f7f591{__obf_b2a2e31336581ab8{}, fmt.Errorf("get %v from invalid", __obf_267bdf615cb1b310)}
	}
	return &__obf_5d650408c7f7f591{__obf_b2a2e31336581ab8{}, fmt.Errorf("%v, get %v from invalid", any.__obf_e56742d6e52f642e, __obf_267bdf615cb1b310)}
}

func (any *__obf_5d650408c7f7f591,) Parse() *Iterator {
	return nil
}

func (any *__obf_5d650408c7f7f591,) GetInterface() any {
	return nil
}
