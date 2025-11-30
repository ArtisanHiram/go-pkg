package __obf_030d39f7a8de96c6

import "fmt"

type __obf_aa2610a72837656e struct {
	__obf_7834a13a259af712
	__obf_fcc907dd69879566 error
}

func __obf_0adb97f96f4a14bd(__obf_f77a9507782b919d []any) *__obf_aa2610a72837656e {
	return &__obf_aa2610a72837656e{__obf_7834a13a259af712{}, fmt.Errorf("%v not found", __obf_f77a9507782b919d)}
}

func (any *__obf_aa2610a72837656e,) LastError() error {
	return any.__obf_fcc907dd69879566
}

func (any *__obf_aa2610a72837656e,) ValueType() ValueType {
	return InvalidValue
}

func (any *__obf_aa2610a72837656e,) MustBeValid() Any {
	panic(any.__obf_fcc907dd69879566)
}

func (any *__obf_aa2610a72837656e,) ToBool() bool {
	return false
}

func (any *__obf_aa2610a72837656e,) ToInt() int {
	return 0
}

func (any *__obf_aa2610a72837656e,) ToInt32() int32 {
	return 0
}

func (any *__obf_aa2610a72837656e,) ToInt64() int64 {
	return 0
}

func (any *__obf_aa2610a72837656e,) ToUint() uint {
	return 0
}

func (any *__obf_aa2610a72837656e,) ToUint32() uint32 {
	return 0
}

func (any *__obf_aa2610a72837656e,) ToUint64() uint64 {
	return 0
}

func (any *__obf_aa2610a72837656e,) ToFloat32() float32 {
	return 0
}

func (any *__obf_aa2610a72837656e,) ToFloat64() float64 {
	return 0
}

func (any *__obf_aa2610a72837656e,) ToString() string {
	return ""
}

func (any *__obf_aa2610a72837656e,) WriteTo(__obf_8a2c51fe22775655 *Stream) {
}

func (any *__obf_aa2610a72837656e,) Get(__obf_f77a9507782b919d ...any) Any {
	if any.__obf_fcc907dd69879566 == nil {
		return &__obf_aa2610a72837656e{__obf_7834a13a259af712{}, fmt.Errorf("get %v from invalid", __obf_f77a9507782b919d)}
	}
	return &__obf_aa2610a72837656e{__obf_7834a13a259af712{}, fmt.Errorf("%v, get %v from invalid", any.__obf_fcc907dd69879566, __obf_f77a9507782b919d)}
}

func (any *__obf_aa2610a72837656e,) Parse() *Iterator {
	return nil
}

func (any *__obf_aa2610a72837656e,) GetInterface() any {
	return nil
}
