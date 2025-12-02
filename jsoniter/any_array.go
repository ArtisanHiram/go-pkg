package __obf_c7b79b12b33d8238

import (
	"reflect"
	"unsafe"
)

type __obf_94d970ccdadcfb91 struct {
	__obf_3a44d9c5204c2e12
	__obf_c52e0bcfad4b8b71 *__obf_30fe5c95cabd69c0
	__obf_684d738bc3ac851a []byte
	__obf_ea0680f8b461a85b error
}

func (any *__obf_94d970ccdadcfb91) ValueType() ValueType {
	return ArrayValue
}

func (any *__obf_94d970ccdadcfb91) MustBeValid() Any {
	return any
}

func (any *__obf_94d970ccdadcfb91) LastError() error {
	return any.__obf_ea0680f8b461a85b
}

func (any *__obf_94d970ccdadcfb91) ToBool() bool {
	__obf_0da8c843dad13139 := any.__obf_c52e0bcfad4b8b71.BorrowIterator(any.__obf_684d738bc3ac851a)
	defer any.__obf_c52e0bcfad4b8b71.ReturnIterator(__obf_0da8c843dad13139)
	return __obf_0da8c843dad13139.ReadArray()
}

func (any *__obf_94d970ccdadcfb91) ToInt() int {
	if any.ToBool() {
		return 1
	}
	return 0
}

func (any *__obf_94d970ccdadcfb91) ToInt32() int32 {
	if any.ToBool() {
		return 1
	}
	return 0
}

func (any *__obf_94d970ccdadcfb91) ToInt64() int64 {
	if any.ToBool() {
		return 1
	}
	return 0
}

func (any *__obf_94d970ccdadcfb91) ToUint() uint {
	if any.ToBool() {
		return 1
	}
	return 0
}

func (any *__obf_94d970ccdadcfb91) ToUint32() uint32 {
	if any.ToBool() {
		return 1
	}
	return 0
}

func (any *__obf_94d970ccdadcfb91) ToUint64() uint64 {
	if any.ToBool() {
		return 1
	}
	return 0
}

func (any *__obf_94d970ccdadcfb91) ToFloat32() float32 {
	if any.ToBool() {
		return 1
	}
	return 0
}

func (any *__obf_94d970ccdadcfb91) ToFloat64() float64 {
	if any.ToBool() {
		return 1
	}
	return 0
}

func (any *__obf_94d970ccdadcfb91) ToString() string {
	return *(*string)(unsafe.Pointer(&any.__obf_684d738bc3ac851a))
}

func (any *__obf_94d970ccdadcfb91) ToVal(__obf_35accd096612ebe4 any) {
	__obf_0da8c843dad13139 := any.__obf_c52e0bcfad4b8b71.BorrowIterator(any.__obf_684d738bc3ac851a)
	defer any.__obf_c52e0bcfad4b8b71.ReturnIterator(__obf_0da8c843dad13139)
	__obf_0da8c843dad13139.
		ReadVal(__obf_35accd096612ebe4)
}

func (any *__obf_94d970ccdadcfb91) Get(__obf_5dde9fb4007294d4 ...any) Any {
	if len(__obf_5dde9fb4007294d4) == 0 {
		return any
	}
	switch __obf_a9297d670cd1f39b := __obf_5dde9fb4007294d4[0].(type) {
	case int:
		__obf_0da8c843dad13139 := any.__obf_c52e0bcfad4b8b71.BorrowIterator(any.__obf_684d738bc3ac851a)
		defer any.__obf_c52e0bcfad4b8b71.ReturnIterator(__obf_0da8c843dad13139)
		__obf_39f3961bda254565 := __obf_7a195008f4d0d118(__obf_0da8c843dad13139, __obf_a9297d670cd1f39b)
		if __obf_39f3961bda254565 == nil {
			return __obf_5b62b1b9d47ae70b(__obf_5dde9fb4007294d4)
		}
		__obf_0da8c843dad13139.
			ResetBytes(__obf_39f3961bda254565)
		return __obf_fedb6216061761ae(__obf_0da8c843dad13139, __obf_5dde9fb4007294d4[1:])
	case int32:
		if '*' == __obf_a9297d670cd1f39b {
			__obf_0da8c843dad13139 := any.__obf_c52e0bcfad4b8b71.BorrowIterator(any.__obf_684d738bc3ac851a)
			defer any.__obf_c52e0bcfad4b8b71.ReturnIterator(__obf_0da8c843dad13139)
			__obf_3e7cc2c62bf39a94 := make([]Any, 0)
			__obf_0da8c843dad13139.
				ReadArrayCB(func(__obf_0da8c843dad13139 *Iterator) bool {
					__obf_22405d874667f998 := __obf_0da8c843dad13139.__obf_772f435663149392().Get(__obf_5dde9fb4007294d4[1:]...)
					if __obf_22405d874667f998.ValueType() != InvalidValue {
						__obf_3e7cc2c62bf39a94 = append(__obf_3e7cc2c62bf39a94, __obf_22405d874667f998)
					}
					return true
				})
			return __obf_a040b37b3279e567(__obf_3e7cc2c62bf39a94)
		}
		return __obf_5b62b1b9d47ae70b(__obf_5dde9fb4007294d4)
	default:
		return __obf_5b62b1b9d47ae70b(__obf_5dde9fb4007294d4)
	}
}

func (any *__obf_94d970ccdadcfb91) Size() int {
	__obf_86901472dc2ab8e1 := 0
	__obf_0da8c843dad13139 := any.__obf_c52e0bcfad4b8b71.BorrowIterator(any.__obf_684d738bc3ac851a)
	defer any.__obf_c52e0bcfad4b8b71.ReturnIterator(__obf_0da8c843dad13139)
	__obf_0da8c843dad13139.
		ReadArrayCB(func(__obf_0da8c843dad13139 *Iterator) bool {
			__obf_86901472dc2ab8e1++
			__obf_0da8c843dad13139.
				Skip()
			return true
		})
	return __obf_86901472dc2ab8e1
}

func (any *__obf_94d970ccdadcfb91) WriteTo(__obf_d8f50bcfe87d8b47 *Stream) {
	__obf_d8f50bcfe87d8b47.
		Write(any.__obf_684d738bc3ac851a)
}

func (any *__obf_94d970ccdadcfb91) GetInterface() any {
	__obf_0da8c843dad13139 := any.__obf_c52e0bcfad4b8b71.BorrowIterator(any.__obf_684d738bc3ac851a)
	defer any.__obf_c52e0bcfad4b8b71.ReturnIterator(__obf_0da8c843dad13139)
	return __obf_0da8c843dad13139.Read()
}

type __obf_fa8a1919e95ff1f8 struct {
	__obf_3a44d9c5204c2e12
	__obf_35accd096612ebe4 reflect.Value
}

func __obf_a040b37b3279e567(__obf_35accd096612ebe4 any) *__obf_fa8a1919e95ff1f8 {
	return &__obf_fa8a1919e95ff1f8{__obf_3a44d9c5204c2e12{}, reflect.ValueOf(__obf_35accd096612ebe4)}
}

func (any *__obf_fa8a1919e95ff1f8) ValueType() ValueType {
	return ArrayValue
}

func (any *__obf_fa8a1919e95ff1f8) MustBeValid() Any {
	return any
}

func (any *__obf_fa8a1919e95ff1f8) LastError() error {
	return nil
}

func (any *__obf_fa8a1919e95ff1f8) ToBool() bool {
	return any.__obf_35accd096612ebe4.Len() != 0
}

func (any *__obf_fa8a1919e95ff1f8) ToInt() int {
	if any.__obf_35accd096612ebe4.Len() == 0 {
		return 0
	}
	return 1
}

func (any *__obf_fa8a1919e95ff1f8) ToInt32() int32 {
	if any.__obf_35accd096612ebe4.Len() == 0 {
		return 0
	}
	return 1
}

func (any *__obf_fa8a1919e95ff1f8) ToInt64() int64 {
	if any.__obf_35accd096612ebe4.Len() == 0 {
		return 0
	}
	return 1
}

func (any *__obf_fa8a1919e95ff1f8) ToUint() uint {
	if any.__obf_35accd096612ebe4.Len() == 0 {
		return 0
	}
	return 1
}

func (any *__obf_fa8a1919e95ff1f8) ToUint32() uint32 {
	if any.__obf_35accd096612ebe4.Len() == 0 {
		return 0
	}
	return 1
}

func (any *__obf_fa8a1919e95ff1f8) ToUint64() uint64 {
	if any.__obf_35accd096612ebe4.Len() == 0 {
		return 0
	}
	return 1
}

func (any *__obf_fa8a1919e95ff1f8) ToFloat32() float32 {
	if any.__obf_35accd096612ebe4.Len() == 0 {
		return 0
	}
	return 1
}

func (any *__obf_fa8a1919e95ff1f8) ToFloat64() float64 {
	if any.__obf_35accd096612ebe4.Len() == 0 {
		return 0
	}
	return 1
}

func (any *__obf_fa8a1919e95ff1f8) ToString() string {
	__obf_a3eaafc22faf1644, _ := MarshalToString(any.__obf_35accd096612ebe4.Interface())
	return __obf_a3eaafc22faf1644
}

func (any *__obf_fa8a1919e95ff1f8) Get(__obf_5dde9fb4007294d4 ...any) Any {
	if len(__obf_5dde9fb4007294d4) == 0 {
		return any
	}
	switch __obf_a9297d670cd1f39b := __obf_5dde9fb4007294d4[0].(type) {
	case int:
		if __obf_a9297d670cd1f39b < 0 || __obf_a9297d670cd1f39b >= any.__obf_35accd096612ebe4.Len() {
			return __obf_5b62b1b9d47ae70b(__obf_5dde9fb4007294d4)
		}
		return Wrap(any.__obf_35accd096612ebe4.Index(__obf_a9297d670cd1f39b).Interface())
	case int32:
		if '*' == __obf_a9297d670cd1f39b {
			__obf_328e2515a5e98ad9 := make([]Any, 0)
			for __obf_a051269af3a541bb := 0; __obf_a051269af3a541bb < any.__obf_35accd096612ebe4.Len(); __obf_a051269af3a541bb++ {
				__obf_cef0ddc31c052dab := Wrap(any.__obf_35accd096612ebe4.Index(__obf_a051269af3a541bb).Interface()).Get(__obf_5dde9fb4007294d4[1:]...)
				if __obf_cef0ddc31c052dab.ValueType() != InvalidValue {
					__obf_328e2515a5e98ad9 = append(__obf_328e2515a5e98ad9, __obf_cef0ddc31c052dab)
				}
			}
			return __obf_a040b37b3279e567(__obf_328e2515a5e98ad9)
		}
		return __obf_5b62b1b9d47ae70b(__obf_5dde9fb4007294d4)
	default:
		return __obf_5b62b1b9d47ae70b(__obf_5dde9fb4007294d4)
	}
}

func (any *__obf_fa8a1919e95ff1f8) Size() int {
	return any.__obf_35accd096612ebe4.Len()
}

func (any *__obf_fa8a1919e95ff1f8) WriteTo(__obf_d8f50bcfe87d8b47 *Stream) {
	__obf_d8f50bcfe87d8b47.
		WriteVal(any.__obf_35accd096612ebe4)
}

func (any *__obf_fa8a1919e95ff1f8) GetInterface() any {
	return any.__obf_35accd096612ebe4.Interface()
}
