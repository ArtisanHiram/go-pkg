package __obf_703d23ba520c3295

import (
	"reflect"
	"unsafe"
)

type __obf_909c8add24236c9b struct {
	__obf_b2a2e31336581ab8
	__obf_25bd4f754a37b862 *__obf_ef74248d8ae9ba78
	__obf_a065f8e0da5f5952 []byte
	__obf_e56742d6e52f642e error
}

func (any *__obf_909c8add24236c9b) ValueType() ValueType {
	return ArrayValue
}

func (any *__obf_909c8add24236c9b) MustBeValid() Any {
	return any
}

func (any *__obf_909c8add24236c9b) LastError() error {
	return any.__obf_e56742d6e52f642e
}

func (any *__obf_909c8add24236c9b) ToBool() bool {
	__obf_47edab4c16a2d88d := any.__obf_25bd4f754a37b862.BorrowIterator(any.__obf_a065f8e0da5f5952)
	defer any.__obf_25bd4f754a37b862.ReturnIterator(__obf_47edab4c16a2d88d)
	return __obf_47edab4c16a2d88d.ReadArray()
}

func (any *__obf_909c8add24236c9b) ToInt() int {
	if any.ToBool() {
		return 1
	}
	return 0
}

func (any *__obf_909c8add24236c9b) ToInt32() int32 {
	if any.ToBool() {
		return 1
	}
	return 0
}

func (any *__obf_909c8add24236c9b) ToInt64() int64 {
	if any.ToBool() {
		return 1
	}
	return 0
}

func (any *__obf_909c8add24236c9b) ToUint() uint {
	if any.ToBool() {
		return 1
	}
	return 0
}

func (any *__obf_909c8add24236c9b) ToUint32() uint32 {
	if any.ToBool() {
		return 1
	}
	return 0
}

func (any *__obf_909c8add24236c9b) ToUint64() uint64 {
	if any.ToBool() {
		return 1
	}
	return 0
}

func (any *__obf_909c8add24236c9b) ToFloat32() float32 {
	if any.ToBool() {
		return 1
	}
	return 0
}

func (any *__obf_909c8add24236c9b) ToFloat64() float64 {
	if any.ToBool() {
		return 1
	}
	return 0
}

func (any *__obf_909c8add24236c9b) ToString() string {
	return *(*string)(unsafe.Pointer(&any.__obf_a065f8e0da5f5952))
}

func (any *__obf_909c8add24236c9b) ToVal(__obf_b924538fffe5fd64 any) {
	__obf_47edab4c16a2d88d := any.__obf_25bd4f754a37b862.BorrowIterator(any.__obf_a065f8e0da5f5952)
	defer any.__obf_25bd4f754a37b862.ReturnIterator(__obf_47edab4c16a2d88d)
	__obf_47edab4c16a2d88d.
		ReadVal(__obf_b924538fffe5fd64)
}

func (any *__obf_909c8add24236c9b) Get(__obf_267bdf615cb1b310 ...any) Any {
	if len(__obf_267bdf615cb1b310) == 0 {
		return any
	}
	switch __obf_1bdb78ff13702ddf := __obf_267bdf615cb1b310[0].(type) {
	case int:
		__obf_47edab4c16a2d88d := any.__obf_25bd4f754a37b862.BorrowIterator(any.__obf_a065f8e0da5f5952)
		defer any.__obf_25bd4f754a37b862.ReturnIterator(__obf_47edab4c16a2d88d)
		__obf_24e18de62be92d01 := __obf_275bd1a9c11e5d3d(__obf_47edab4c16a2d88d, __obf_1bdb78ff13702ddf)
		if __obf_24e18de62be92d01 == nil {
			return __obf_b82811e6c38b9e10(__obf_267bdf615cb1b310)
		}
		__obf_47edab4c16a2d88d.
			ResetBytes(__obf_24e18de62be92d01)
		return __obf_da70fd4b038d543a(__obf_47edab4c16a2d88d, __obf_267bdf615cb1b310[1:])
	case int32:
		if '*' == __obf_1bdb78ff13702ddf {
			__obf_47edab4c16a2d88d := any.__obf_25bd4f754a37b862.BorrowIterator(any.__obf_a065f8e0da5f5952)
			defer any.__obf_25bd4f754a37b862.ReturnIterator(__obf_47edab4c16a2d88d)
			__obf_f21bafabcc6ccf82 := make([]Any, 0)
			__obf_47edab4c16a2d88d.
				ReadArrayCB(func(__obf_47edab4c16a2d88d *Iterator) bool {
					__obf_87a836f73d31233b := __obf_47edab4c16a2d88d.__obf_721b69294316727a().Get(__obf_267bdf615cb1b310[1:]...)
					if __obf_87a836f73d31233b.ValueType() != InvalidValue {
						__obf_f21bafabcc6ccf82 = append(__obf_f21bafabcc6ccf82, __obf_87a836f73d31233b)
					}
					return true
				})
			return __obf_7484f9ce272651b2(__obf_f21bafabcc6ccf82)
		}
		return __obf_b82811e6c38b9e10(__obf_267bdf615cb1b310)
	default:
		return __obf_b82811e6c38b9e10(__obf_267bdf615cb1b310)
	}
}

func (any *__obf_909c8add24236c9b) Size() int {
	__obf_0126ec6b3c37befb := 0
	__obf_47edab4c16a2d88d := any.__obf_25bd4f754a37b862.BorrowIterator(any.__obf_a065f8e0da5f5952)
	defer any.__obf_25bd4f754a37b862.ReturnIterator(__obf_47edab4c16a2d88d)
	__obf_47edab4c16a2d88d.
		ReadArrayCB(func(__obf_47edab4c16a2d88d *Iterator) bool {
			__obf_0126ec6b3c37befb++
			__obf_47edab4c16a2d88d.
				Skip()
			return true
		})
	return __obf_0126ec6b3c37befb
}

func (any *__obf_909c8add24236c9b) WriteTo(__obf_9a34f62857fb1d1d *Stream) {
	__obf_9a34f62857fb1d1d.
		Write(any.__obf_a065f8e0da5f5952)
}

func (any *__obf_909c8add24236c9b) GetInterface() any {
	__obf_47edab4c16a2d88d := any.__obf_25bd4f754a37b862.BorrowIterator(any.__obf_a065f8e0da5f5952)
	defer any.__obf_25bd4f754a37b862.ReturnIterator(__obf_47edab4c16a2d88d)
	return __obf_47edab4c16a2d88d.Read()
}

type __obf_345b4c12577be4bb struct {
	__obf_b2a2e31336581ab8
	__obf_b924538fffe5fd64 reflect.Value
}

func __obf_7484f9ce272651b2(__obf_b924538fffe5fd64 any) *__obf_345b4c12577be4bb {
	return &__obf_345b4c12577be4bb{__obf_b2a2e31336581ab8{}, reflect.ValueOf(__obf_b924538fffe5fd64)}
}

func (any *__obf_345b4c12577be4bb) ValueType() ValueType {
	return ArrayValue
}

func (any *__obf_345b4c12577be4bb) MustBeValid() Any {
	return any
}

func (any *__obf_345b4c12577be4bb) LastError() error {
	return nil
}

func (any *__obf_345b4c12577be4bb) ToBool() bool {
	return any.__obf_b924538fffe5fd64.Len() != 0
}

func (any *__obf_345b4c12577be4bb) ToInt() int {
	if any.__obf_b924538fffe5fd64.Len() == 0 {
		return 0
	}
	return 1
}

func (any *__obf_345b4c12577be4bb) ToInt32() int32 {
	if any.__obf_b924538fffe5fd64.Len() == 0 {
		return 0
	}
	return 1
}

func (any *__obf_345b4c12577be4bb) ToInt64() int64 {
	if any.__obf_b924538fffe5fd64.Len() == 0 {
		return 0
	}
	return 1
}

func (any *__obf_345b4c12577be4bb) ToUint() uint {
	if any.__obf_b924538fffe5fd64.Len() == 0 {
		return 0
	}
	return 1
}

func (any *__obf_345b4c12577be4bb) ToUint32() uint32 {
	if any.__obf_b924538fffe5fd64.Len() == 0 {
		return 0
	}
	return 1
}

func (any *__obf_345b4c12577be4bb) ToUint64() uint64 {
	if any.__obf_b924538fffe5fd64.Len() == 0 {
		return 0
	}
	return 1
}

func (any *__obf_345b4c12577be4bb) ToFloat32() float32 {
	if any.__obf_b924538fffe5fd64.Len() == 0 {
		return 0
	}
	return 1
}

func (any *__obf_345b4c12577be4bb) ToFloat64() float64 {
	if any.__obf_b924538fffe5fd64.Len() == 0 {
		return 0
	}
	return 1
}

func (any *__obf_345b4c12577be4bb) ToString() string {
	__obf_44b48c26051f8033, _ := MarshalToString(any.__obf_b924538fffe5fd64.Interface())
	return __obf_44b48c26051f8033
}

func (any *__obf_345b4c12577be4bb) Get(__obf_267bdf615cb1b310 ...any) Any {
	if len(__obf_267bdf615cb1b310) == 0 {
		return any
	}
	switch __obf_1bdb78ff13702ddf := __obf_267bdf615cb1b310[0].(type) {
	case int:
		if __obf_1bdb78ff13702ddf < 0 || __obf_1bdb78ff13702ddf >= any.__obf_b924538fffe5fd64.Len() {
			return __obf_b82811e6c38b9e10(__obf_267bdf615cb1b310)
		}
		return Wrap(any.__obf_b924538fffe5fd64.Index(__obf_1bdb78ff13702ddf).Interface())
	case int32:
		if '*' == __obf_1bdb78ff13702ddf {
			__obf_cc4c2d1b183679ba := make([]Any, 0)
			for __obf_b0a5d2bd48690f1d := 0; __obf_b0a5d2bd48690f1d < any.__obf_b924538fffe5fd64.Len(); __obf_b0a5d2bd48690f1d++ {
				__obf_f7372bba52102679 := Wrap(any.__obf_b924538fffe5fd64.Index(__obf_b0a5d2bd48690f1d).Interface()).Get(__obf_267bdf615cb1b310[1:]...)
				if __obf_f7372bba52102679.ValueType() != InvalidValue {
					__obf_cc4c2d1b183679ba = append(__obf_cc4c2d1b183679ba, __obf_f7372bba52102679)
				}
			}
			return __obf_7484f9ce272651b2(__obf_cc4c2d1b183679ba)
		}
		return __obf_b82811e6c38b9e10(__obf_267bdf615cb1b310)
	default:
		return __obf_b82811e6c38b9e10(__obf_267bdf615cb1b310)
	}
}

func (any *__obf_345b4c12577be4bb) Size() int {
	return any.__obf_b924538fffe5fd64.Len()
}

func (any *__obf_345b4c12577be4bb) WriteTo(__obf_9a34f62857fb1d1d *Stream) {
	__obf_9a34f62857fb1d1d.
		WriteVal(any.__obf_b924538fffe5fd64)
}

func (any *__obf_345b4c12577be4bb) GetInterface() any {
	return any.__obf_b924538fffe5fd64.Interface()
}
