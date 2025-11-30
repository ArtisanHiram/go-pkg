package __obf_703d23ba520c3295

import (
	"reflect"
	"unsafe"
)

type __obf_c7e0be96be0cd81d struct {
	__obf_b2a2e31336581ab8
	__obf_25bd4f754a37b862 *__obf_ef74248d8ae9ba78
	__obf_a065f8e0da5f5952 []byte
	__obf_e56742d6e52f642e error
}

func (any *__obf_c7e0be96be0cd81d) ValueType() ValueType {
	return ObjectValue
}

func (any *__obf_c7e0be96be0cd81d) MustBeValid() Any {
	return any
}

func (any *__obf_c7e0be96be0cd81d) LastError() error {
	return any.__obf_e56742d6e52f642e
}

func (any *__obf_c7e0be96be0cd81d) ToBool() bool {
	return true
}

func (any *__obf_c7e0be96be0cd81d) ToInt() int {
	return 0
}

func (any *__obf_c7e0be96be0cd81d) ToInt32() int32 {
	return 0
}

func (any *__obf_c7e0be96be0cd81d) ToInt64() int64 {
	return 0
}

func (any *__obf_c7e0be96be0cd81d) ToUint() uint {
	return 0
}

func (any *__obf_c7e0be96be0cd81d) ToUint32() uint32 {
	return 0
}

func (any *__obf_c7e0be96be0cd81d) ToUint64() uint64 {
	return 0
}

func (any *__obf_c7e0be96be0cd81d) ToFloat32() float32 {
	return 0
}

func (any *__obf_c7e0be96be0cd81d) ToFloat64() float64 {
	return 0
}

func (any *__obf_c7e0be96be0cd81d) ToString() string {
	return *(*string)(unsafe.Pointer(&any.__obf_a065f8e0da5f5952))
}

func (any *__obf_c7e0be96be0cd81d) ToVal(__obf_02ba706f4f104d71 any) {
	__obf_47edab4c16a2d88d := any.__obf_25bd4f754a37b862.BorrowIterator(any.__obf_a065f8e0da5f5952)
	defer any.__obf_25bd4f754a37b862.ReturnIterator(__obf_47edab4c16a2d88d)
	__obf_47edab4c16a2d88d.
		ReadVal(__obf_02ba706f4f104d71)
}

func (any *__obf_c7e0be96be0cd81d) Get(__obf_267bdf615cb1b310 ...any) Any {
	if len(__obf_267bdf615cb1b310) == 0 {
		return any
	}
	switch __obf_1bdb78ff13702ddf := __obf_267bdf615cb1b310[0].(type) {
	case string:
		__obf_47edab4c16a2d88d := any.__obf_25bd4f754a37b862.BorrowIterator(any.__obf_a065f8e0da5f5952)
		defer any.__obf_25bd4f754a37b862.ReturnIterator(__obf_47edab4c16a2d88d)
		__obf_24e18de62be92d01 := __obf_29b231eb99e7ebaa(__obf_47edab4c16a2d88d, __obf_1bdb78ff13702ddf)
		if __obf_24e18de62be92d01 == nil {
			return __obf_b82811e6c38b9e10(__obf_267bdf615cb1b310)
		}
		__obf_47edab4c16a2d88d.
			ResetBytes(__obf_24e18de62be92d01)
		return __obf_da70fd4b038d543a(__obf_47edab4c16a2d88d, __obf_267bdf615cb1b310[1:])
	case int32:
		if '*' == __obf_1bdb78ff13702ddf {
			__obf_cc4c2d1b183679ba := map[string]Any{}
			__obf_47edab4c16a2d88d := any.__obf_25bd4f754a37b862.BorrowIterator(any.__obf_a065f8e0da5f5952)
			defer any.__obf_25bd4f754a37b862.ReturnIterator(__obf_47edab4c16a2d88d)
			__obf_47edab4c16a2d88d.
				ReadMapCB(func(__obf_47edab4c16a2d88d *Iterator, __obf_13f6440419533990 string) bool {
					__obf_f7372bba52102679 := __obf_da70fd4b038d543a(__obf_47edab4c16a2d88d, __obf_267bdf615cb1b310[1:])
					if __obf_f7372bba52102679.ValueType() != InvalidValue {
						__obf_cc4c2d1b183679ba[__obf_13f6440419533990] = __obf_f7372bba52102679
					}
					return true
				})
			return __obf_495ed748a54702d1(__obf_cc4c2d1b183679ba)
		}
		return __obf_b82811e6c38b9e10(__obf_267bdf615cb1b310)
	default:
		return __obf_b82811e6c38b9e10(__obf_267bdf615cb1b310)
	}
}

func (any *__obf_c7e0be96be0cd81d) Keys() []string {
	__obf_ca4e023a99d875e2 := []string{}
	__obf_47edab4c16a2d88d := any.__obf_25bd4f754a37b862.BorrowIterator(any.__obf_a065f8e0da5f5952)
	defer any.__obf_25bd4f754a37b862.ReturnIterator(__obf_47edab4c16a2d88d)
	__obf_47edab4c16a2d88d.
		ReadMapCB(func(__obf_47edab4c16a2d88d *Iterator, __obf_13f6440419533990 string) bool {
			__obf_47edab4c16a2d88d.
				Skip()
			__obf_ca4e023a99d875e2 = append(__obf_ca4e023a99d875e2, __obf_13f6440419533990)
			return true
		})
	return __obf_ca4e023a99d875e2
}

func (any *__obf_c7e0be96be0cd81d) Size() int {
	__obf_0126ec6b3c37befb := 0
	__obf_47edab4c16a2d88d := any.__obf_25bd4f754a37b862.BorrowIterator(any.__obf_a065f8e0da5f5952)
	defer any.__obf_25bd4f754a37b862.ReturnIterator(__obf_47edab4c16a2d88d)
	__obf_47edab4c16a2d88d.
		ReadObjectCB(func(__obf_47edab4c16a2d88d *Iterator, __obf_13f6440419533990 string) bool {
			__obf_47edab4c16a2d88d.
				Skip()
			__obf_0126ec6b3c37befb++
			return true
		})
	return __obf_0126ec6b3c37befb
}

func (any *__obf_c7e0be96be0cd81d) WriteTo(__obf_9a34f62857fb1d1d *Stream) {
	__obf_9a34f62857fb1d1d.
		Write(any.__obf_a065f8e0da5f5952)
}

func (any *__obf_c7e0be96be0cd81d) GetInterface() any {
	__obf_47edab4c16a2d88d := any.__obf_25bd4f754a37b862.BorrowIterator(any.__obf_a065f8e0da5f5952)
	defer any.__obf_25bd4f754a37b862.ReturnIterator(__obf_47edab4c16a2d88d)
	return __obf_47edab4c16a2d88d.Read()
}

type __obf_f73f34b3e3f3e932 struct {
	__obf_b2a2e31336581ab8
	__obf_e56742d6e52f642e error
	__obf_b924538fffe5fd64 reflect.Value
}

func __obf_70c50c635ea444d2(__obf_b924538fffe5fd64 any) *__obf_f73f34b3e3f3e932 {
	return &__obf_f73f34b3e3f3e932{__obf_b2a2e31336581ab8{}, nil, reflect.ValueOf(__obf_b924538fffe5fd64)}
}

func (any *__obf_f73f34b3e3f3e932) ValueType() ValueType {
	return ObjectValue
}

func (any *__obf_f73f34b3e3f3e932) MustBeValid() Any {
	return any
}

func (any *__obf_f73f34b3e3f3e932) Parse() *Iterator {
	return nil
}

func (any *__obf_f73f34b3e3f3e932) LastError() error {
	return any.__obf_e56742d6e52f642e
}

func (any *__obf_f73f34b3e3f3e932) ToBool() bool {
	return any.__obf_b924538fffe5fd64.NumField() != 0
}

func (any *__obf_f73f34b3e3f3e932) ToInt() int {
	return 0
}

func (any *__obf_f73f34b3e3f3e932) ToInt32() int32 {
	return 0
}

func (any *__obf_f73f34b3e3f3e932) ToInt64() int64 {
	return 0
}

func (any *__obf_f73f34b3e3f3e932) ToUint() uint {
	return 0
}

func (any *__obf_f73f34b3e3f3e932) ToUint32() uint32 {
	return 0
}

func (any *__obf_f73f34b3e3f3e932) ToUint64() uint64 {
	return 0
}

func (any *__obf_f73f34b3e3f3e932) ToFloat32() float32 {
	return 0
}

func (any *__obf_f73f34b3e3f3e932) ToFloat64() float64 {
	return 0
}

func (any *__obf_f73f34b3e3f3e932) ToString() string {
	__obf_44b48c26051f8033, __obf_e56742d6e52f642e := MarshalToString(any.__obf_b924538fffe5fd64.Interface())
	any.__obf_e56742d6e52f642e = __obf_e56742d6e52f642e
	return __obf_44b48c26051f8033
}

func (any *__obf_f73f34b3e3f3e932) Get(__obf_267bdf615cb1b310 ...any) Any {
	if len(__obf_267bdf615cb1b310) == 0 {
		return any
	}
	switch __obf_1bdb78ff13702ddf := __obf_267bdf615cb1b310[0].(type) {
	case string:
		__obf_13f6440419533990 := any.__obf_b924538fffe5fd64.FieldByName(__obf_1bdb78ff13702ddf)
		if !__obf_13f6440419533990.IsValid() {
			return __obf_b82811e6c38b9e10(__obf_267bdf615cb1b310)
		}
		return Wrap(__obf_13f6440419533990.Interface())
	case int32:
		if '*' == __obf_1bdb78ff13702ddf {
			__obf_cc4c2d1b183679ba := map[string]Any{}
			for __obf_b0a5d2bd48690f1d := 0; __obf_b0a5d2bd48690f1d < any.__obf_b924538fffe5fd64.NumField(); __obf_b0a5d2bd48690f1d++ {
				__obf_13f6440419533990 := any.__obf_b924538fffe5fd64.Field(__obf_b0a5d2bd48690f1d)
				if __obf_13f6440419533990.CanInterface() {
					__obf_f7372bba52102679 := Wrap(__obf_13f6440419533990.Interface()).Get(__obf_267bdf615cb1b310[1:]...)
					if __obf_f7372bba52102679.ValueType() != InvalidValue {
						__obf_cc4c2d1b183679ba[any.__obf_b924538fffe5fd64.Type().Field(__obf_b0a5d2bd48690f1d).Name] = __obf_f7372bba52102679
					}
				}
			}
			return __obf_495ed748a54702d1(__obf_cc4c2d1b183679ba)
		}
		return __obf_b82811e6c38b9e10(__obf_267bdf615cb1b310)
	default:
		return __obf_b82811e6c38b9e10(__obf_267bdf615cb1b310)
	}
}

func (any *__obf_f73f34b3e3f3e932) Keys() []string {
	__obf_ca4e023a99d875e2 := make([]string, 0, any.__obf_b924538fffe5fd64.NumField())
	for __obf_b0a5d2bd48690f1d := 0; __obf_b0a5d2bd48690f1d < any.__obf_b924538fffe5fd64.NumField(); __obf_b0a5d2bd48690f1d++ {
		__obf_ca4e023a99d875e2 = append(__obf_ca4e023a99d875e2, any.__obf_b924538fffe5fd64.Type().Field(__obf_b0a5d2bd48690f1d).Name)
	}
	return __obf_ca4e023a99d875e2
}

func (any *__obf_f73f34b3e3f3e932) Size() int {
	return any.__obf_b924538fffe5fd64.NumField()
}

func (any *__obf_f73f34b3e3f3e932) WriteTo(__obf_9a34f62857fb1d1d *Stream) {
	__obf_9a34f62857fb1d1d.
		WriteVal(any.__obf_b924538fffe5fd64)
}

func (any *__obf_f73f34b3e3f3e932) GetInterface() any {
	return any.__obf_b924538fffe5fd64.Interface()
}

type __obf_1e9102439d3f1d8e struct {
	__obf_b2a2e31336581ab8
	__obf_e56742d6e52f642e error
	__obf_b924538fffe5fd64 reflect.Value
}

func __obf_495ed748a54702d1(__obf_b924538fffe5fd64 any) *__obf_1e9102439d3f1d8e {
	return &__obf_1e9102439d3f1d8e{__obf_b2a2e31336581ab8{}, nil, reflect.ValueOf(__obf_b924538fffe5fd64)}
}

func (any *__obf_1e9102439d3f1d8e) ValueType() ValueType {
	return ObjectValue
}

func (any *__obf_1e9102439d3f1d8e) MustBeValid() Any {
	return any
}

func (any *__obf_1e9102439d3f1d8e) Parse() *Iterator {
	return nil
}

func (any *__obf_1e9102439d3f1d8e) LastError() error {
	return any.__obf_e56742d6e52f642e
}

func (any *__obf_1e9102439d3f1d8e) ToBool() bool {
	return true
}

func (any *__obf_1e9102439d3f1d8e) ToInt() int {
	return 0
}

func (any *__obf_1e9102439d3f1d8e) ToInt32() int32 {
	return 0
}

func (any *__obf_1e9102439d3f1d8e) ToInt64() int64 {
	return 0
}

func (any *__obf_1e9102439d3f1d8e) ToUint() uint {
	return 0
}

func (any *__obf_1e9102439d3f1d8e) ToUint32() uint32 {
	return 0
}

func (any *__obf_1e9102439d3f1d8e) ToUint64() uint64 {
	return 0
}

func (any *__obf_1e9102439d3f1d8e) ToFloat32() float32 {
	return 0
}

func (any *__obf_1e9102439d3f1d8e) ToFloat64() float64 {
	return 0
}

func (any *__obf_1e9102439d3f1d8e) ToString() string {
	__obf_44b48c26051f8033, __obf_e56742d6e52f642e := MarshalToString(any.__obf_b924538fffe5fd64.Interface())
	any.__obf_e56742d6e52f642e = __obf_e56742d6e52f642e
	return __obf_44b48c26051f8033
}

func (any *__obf_1e9102439d3f1d8e) Get(__obf_267bdf615cb1b310 ...any) Any {
	if len(__obf_267bdf615cb1b310) == 0 {
		return any
	}
	switch __obf_1bdb78ff13702ddf := __obf_267bdf615cb1b310[0].(type) {
	case int32:
		if '*' == __obf_1bdb78ff13702ddf {
			__obf_cc4c2d1b183679ba := map[string]Any{}
			for _, __obf_4580e6a9e5b1ee92 := range any.__obf_b924538fffe5fd64.MapKeys() {
				__obf_29b3d2ee7910f2ff := __obf_4580e6a9e5b1ee92.String()
				__obf_40727fc8bdb3249a := Wrap(any.__obf_b924538fffe5fd64.MapIndex(__obf_4580e6a9e5b1ee92).Interface())
				__obf_f7372bba52102679 := __obf_40727fc8bdb3249a.Get(__obf_267bdf615cb1b310[1:]...)
				if __obf_f7372bba52102679.ValueType() != InvalidValue {
					__obf_cc4c2d1b183679ba[__obf_29b3d2ee7910f2ff] = __obf_f7372bba52102679
				}
			}
			return __obf_495ed748a54702d1(__obf_cc4c2d1b183679ba)
		}
		return __obf_b82811e6c38b9e10(__obf_267bdf615cb1b310)
	default:
		__obf_ccb7d8cb07a5350f := any.__obf_b924538fffe5fd64.MapIndex(reflect.ValueOf(__obf_1bdb78ff13702ddf))
		if !__obf_ccb7d8cb07a5350f.IsValid() {
			return __obf_b82811e6c38b9e10(__obf_267bdf615cb1b310)
		}
		return Wrap(__obf_ccb7d8cb07a5350f.Interface())
	}
}

func (any *__obf_1e9102439d3f1d8e) Keys() []string {
	__obf_ca4e023a99d875e2 := make([]string, 0, any.__obf_b924538fffe5fd64.Len())
	for _, __obf_4580e6a9e5b1ee92 := range any.__obf_b924538fffe5fd64.MapKeys() {
		__obf_ca4e023a99d875e2 = append(__obf_ca4e023a99d875e2, __obf_4580e6a9e5b1ee92.String())
	}
	return __obf_ca4e023a99d875e2
}

func (any *__obf_1e9102439d3f1d8e) Size() int {
	return any.__obf_b924538fffe5fd64.Len()
}

func (any *__obf_1e9102439d3f1d8e) WriteTo(__obf_9a34f62857fb1d1d *Stream) {
	__obf_9a34f62857fb1d1d.
		WriteVal(any.__obf_b924538fffe5fd64)
}

func (any *__obf_1e9102439d3f1d8e) GetInterface() any {
	return any.__obf_b924538fffe5fd64.Interface()
}
