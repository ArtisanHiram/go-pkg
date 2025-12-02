package __obf_c7b79b12b33d8238

import (
	"reflect"
	"unsafe"
)

type __obf_293b886b2fccdeac struct {
	__obf_3a44d9c5204c2e12
	__obf_c52e0bcfad4b8b71 *__obf_30fe5c95cabd69c0
	__obf_684d738bc3ac851a []byte
	__obf_ea0680f8b461a85b error
}

func (any *__obf_293b886b2fccdeac) ValueType() ValueType {
	return ObjectValue
}

func (any *__obf_293b886b2fccdeac) MustBeValid() Any {
	return any
}

func (any *__obf_293b886b2fccdeac) LastError() error {
	return any.__obf_ea0680f8b461a85b
}

func (any *__obf_293b886b2fccdeac) ToBool() bool {
	return true
}

func (any *__obf_293b886b2fccdeac) ToInt() int {
	return 0
}

func (any *__obf_293b886b2fccdeac) ToInt32() int32 {
	return 0
}

func (any *__obf_293b886b2fccdeac) ToInt64() int64 {
	return 0
}

func (any *__obf_293b886b2fccdeac) ToUint() uint {
	return 0
}

func (any *__obf_293b886b2fccdeac) ToUint32() uint32 {
	return 0
}

func (any *__obf_293b886b2fccdeac) ToUint64() uint64 {
	return 0
}

func (any *__obf_293b886b2fccdeac) ToFloat32() float32 {
	return 0
}

func (any *__obf_293b886b2fccdeac) ToFloat64() float64 {
	return 0
}

func (any *__obf_293b886b2fccdeac) ToString() string {
	return *(*string)(unsafe.Pointer(&any.__obf_684d738bc3ac851a))
}

func (any *__obf_293b886b2fccdeac) ToVal(__obf_d6e2df4782353c64 any) {
	__obf_0da8c843dad13139 := any.__obf_c52e0bcfad4b8b71.BorrowIterator(any.__obf_684d738bc3ac851a)
	defer any.__obf_c52e0bcfad4b8b71.ReturnIterator(__obf_0da8c843dad13139)
	__obf_0da8c843dad13139.
		ReadVal(__obf_d6e2df4782353c64)
}

func (any *__obf_293b886b2fccdeac) Get(__obf_5dde9fb4007294d4 ...any) Any {
	if len(__obf_5dde9fb4007294d4) == 0 {
		return any
	}
	switch __obf_a9297d670cd1f39b := __obf_5dde9fb4007294d4[0].(type) {
	case string:
		__obf_0da8c843dad13139 := any.__obf_c52e0bcfad4b8b71.BorrowIterator(any.__obf_684d738bc3ac851a)
		defer any.__obf_c52e0bcfad4b8b71.ReturnIterator(__obf_0da8c843dad13139)
		__obf_39f3961bda254565 := __obf_17f2269292836639(__obf_0da8c843dad13139, __obf_a9297d670cd1f39b)
		if __obf_39f3961bda254565 == nil {
			return __obf_5b62b1b9d47ae70b(__obf_5dde9fb4007294d4)
		}
		__obf_0da8c843dad13139.
			ResetBytes(__obf_39f3961bda254565)
		return __obf_fedb6216061761ae(__obf_0da8c843dad13139, __obf_5dde9fb4007294d4[1:])
	case int32:
		if '*' == __obf_a9297d670cd1f39b {
			__obf_328e2515a5e98ad9 := map[string]Any{}
			__obf_0da8c843dad13139 := any.__obf_c52e0bcfad4b8b71.BorrowIterator(any.__obf_684d738bc3ac851a)
			defer any.__obf_c52e0bcfad4b8b71.ReturnIterator(__obf_0da8c843dad13139)
			__obf_0da8c843dad13139.
				ReadMapCB(func(__obf_0da8c843dad13139 *Iterator, __obf_ad34f8a6de2b01cb string) bool {
					__obf_cef0ddc31c052dab := __obf_fedb6216061761ae(__obf_0da8c843dad13139, __obf_5dde9fb4007294d4[1:])
					if __obf_cef0ddc31c052dab.ValueType() != InvalidValue {
						__obf_328e2515a5e98ad9[__obf_ad34f8a6de2b01cb] = __obf_cef0ddc31c052dab
					}
					return true
				})
			return __obf_2259c4eb0c25660e(__obf_328e2515a5e98ad9)
		}
		return __obf_5b62b1b9d47ae70b(__obf_5dde9fb4007294d4)
	default:
		return __obf_5b62b1b9d47ae70b(__obf_5dde9fb4007294d4)
	}
}

func (any *__obf_293b886b2fccdeac) Keys() []string {
	__obf_5308cb213e4a53a1 := []string{}
	__obf_0da8c843dad13139 := any.__obf_c52e0bcfad4b8b71.BorrowIterator(any.__obf_684d738bc3ac851a)
	defer any.__obf_c52e0bcfad4b8b71.ReturnIterator(__obf_0da8c843dad13139)
	__obf_0da8c843dad13139.
		ReadMapCB(func(__obf_0da8c843dad13139 *Iterator, __obf_ad34f8a6de2b01cb string) bool {
			__obf_0da8c843dad13139.
				Skip()
			__obf_5308cb213e4a53a1 = append(__obf_5308cb213e4a53a1, __obf_ad34f8a6de2b01cb)
			return true
		})
	return __obf_5308cb213e4a53a1
}

func (any *__obf_293b886b2fccdeac) Size() int {
	__obf_86901472dc2ab8e1 := 0
	__obf_0da8c843dad13139 := any.__obf_c52e0bcfad4b8b71.BorrowIterator(any.__obf_684d738bc3ac851a)
	defer any.__obf_c52e0bcfad4b8b71.ReturnIterator(__obf_0da8c843dad13139)
	__obf_0da8c843dad13139.
		ReadObjectCB(func(__obf_0da8c843dad13139 *Iterator, __obf_ad34f8a6de2b01cb string) bool {
			__obf_0da8c843dad13139.
				Skip()
			__obf_86901472dc2ab8e1++
			return true
		})
	return __obf_86901472dc2ab8e1
}

func (any *__obf_293b886b2fccdeac) WriteTo(__obf_d8f50bcfe87d8b47 *Stream) {
	__obf_d8f50bcfe87d8b47.
		Write(any.__obf_684d738bc3ac851a)
}

func (any *__obf_293b886b2fccdeac) GetInterface() any {
	__obf_0da8c843dad13139 := any.__obf_c52e0bcfad4b8b71.BorrowIterator(any.__obf_684d738bc3ac851a)
	defer any.__obf_c52e0bcfad4b8b71.ReturnIterator(__obf_0da8c843dad13139)
	return __obf_0da8c843dad13139.Read()
}

type __obf_cb32dfdd7061c2a0 struct {
	__obf_3a44d9c5204c2e12
	__obf_ea0680f8b461a85b error
	__obf_35accd096612ebe4 reflect.Value
}

func __obf_e7c64fed69971093(__obf_35accd096612ebe4 any) *__obf_cb32dfdd7061c2a0 {
	return &__obf_cb32dfdd7061c2a0{__obf_3a44d9c5204c2e12{}, nil, reflect.ValueOf(__obf_35accd096612ebe4)}
}

func (any *__obf_cb32dfdd7061c2a0) ValueType() ValueType {
	return ObjectValue
}

func (any *__obf_cb32dfdd7061c2a0) MustBeValid() Any {
	return any
}

func (any *__obf_cb32dfdd7061c2a0) Parse() *Iterator {
	return nil
}

func (any *__obf_cb32dfdd7061c2a0) LastError() error {
	return any.__obf_ea0680f8b461a85b
}

func (any *__obf_cb32dfdd7061c2a0) ToBool() bool {
	return any.__obf_35accd096612ebe4.NumField() != 0
}

func (any *__obf_cb32dfdd7061c2a0) ToInt() int {
	return 0
}

func (any *__obf_cb32dfdd7061c2a0) ToInt32() int32 {
	return 0
}

func (any *__obf_cb32dfdd7061c2a0) ToInt64() int64 {
	return 0
}

func (any *__obf_cb32dfdd7061c2a0) ToUint() uint {
	return 0
}

func (any *__obf_cb32dfdd7061c2a0) ToUint32() uint32 {
	return 0
}

func (any *__obf_cb32dfdd7061c2a0) ToUint64() uint64 {
	return 0
}

func (any *__obf_cb32dfdd7061c2a0) ToFloat32() float32 {
	return 0
}

func (any *__obf_cb32dfdd7061c2a0) ToFloat64() float64 {
	return 0
}

func (any *__obf_cb32dfdd7061c2a0) ToString() string {
	__obf_a3eaafc22faf1644, __obf_ea0680f8b461a85b := MarshalToString(any.__obf_35accd096612ebe4.Interface())
	any.__obf_ea0680f8b461a85b = __obf_ea0680f8b461a85b
	return __obf_a3eaafc22faf1644
}

func (any *__obf_cb32dfdd7061c2a0) Get(__obf_5dde9fb4007294d4 ...any) Any {
	if len(__obf_5dde9fb4007294d4) == 0 {
		return any
	}
	switch __obf_a9297d670cd1f39b := __obf_5dde9fb4007294d4[0].(type) {
	case string:
		__obf_ad34f8a6de2b01cb := any.__obf_35accd096612ebe4.FieldByName(__obf_a9297d670cd1f39b)
		if !__obf_ad34f8a6de2b01cb.IsValid() {
			return __obf_5b62b1b9d47ae70b(__obf_5dde9fb4007294d4)
		}
		return Wrap(__obf_ad34f8a6de2b01cb.Interface())
	case int32:
		if '*' == __obf_a9297d670cd1f39b {
			__obf_328e2515a5e98ad9 := map[string]Any{}
			for __obf_a051269af3a541bb := 0; __obf_a051269af3a541bb < any.__obf_35accd096612ebe4.NumField(); __obf_a051269af3a541bb++ {
				__obf_ad34f8a6de2b01cb := any.__obf_35accd096612ebe4.Field(__obf_a051269af3a541bb)
				if __obf_ad34f8a6de2b01cb.CanInterface() {
					__obf_cef0ddc31c052dab := Wrap(__obf_ad34f8a6de2b01cb.Interface()).Get(__obf_5dde9fb4007294d4[1:]...)
					if __obf_cef0ddc31c052dab.ValueType() != InvalidValue {
						__obf_328e2515a5e98ad9[any.__obf_35accd096612ebe4.Type().Field(__obf_a051269af3a541bb).Name] = __obf_cef0ddc31c052dab
					}
				}
			}
			return __obf_2259c4eb0c25660e(__obf_328e2515a5e98ad9)
		}
		return __obf_5b62b1b9d47ae70b(__obf_5dde9fb4007294d4)
	default:
		return __obf_5b62b1b9d47ae70b(__obf_5dde9fb4007294d4)
	}
}

func (any *__obf_cb32dfdd7061c2a0) Keys() []string {
	__obf_5308cb213e4a53a1 := make([]string, 0, any.__obf_35accd096612ebe4.NumField())
	for __obf_a051269af3a541bb := 0; __obf_a051269af3a541bb < any.__obf_35accd096612ebe4.NumField(); __obf_a051269af3a541bb++ {
		__obf_5308cb213e4a53a1 = append(__obf_5308cb213e4a53a1, any.__obf_35accd096612ebe4.Type().Field(__obf_a051269af3a541bb).Name)
	}
	return __obf_5308cb213e4a53a1
}

func (any *__obf_cb32dfdd7061c2a0) Size() int {
	return any.__obf_35accd096612ebe4.NumField()
}

func (any *__obf_cb32dfdd7061c2a0) WriteTo(__obf_d8f50bcfe87d8b47 *Stream) {
	__obf_d8f50bcfe87d8b47.
		WriteVal(any.__obf_35accd096612ebe4)
}

func (any *__obf_cb32dfdd7061c2a0) GetInterface() any {
	return any.__obf_35accd096612ebe4.Interface()
}

type __obf_25190d85d579f33a struct {
	__obf_3a44d9c5204c2e12
	__obf_ea0680f8b461a85b error
	__obf_35accd096612ebe4 reflect.Value
}

func __obf_2259c4eb0c25660e(__obf_35accd096612ebe4 any) *__obf_25190d85d579f33a {
	return &__obf_25190d85d579f33a{__obf_3a44d9c5204c2e12{}, nil, reflect.ValueOf(__obf_35accd096612ebe4)}
}

func (any *__obf_25190d85d579f33a) ValueType() ValueType {
	return ObjectValue
}

func (any *__obf_25190d85d579f33a) MustBeValid() Any {
	return any
}

func (any *__obf_25190d85d579f33a) Parse() *Iterator {
	return nil
}

func (any *__obf_25190d85d579f33a) LastError() error {
	return any.__obf_ea0680f8b461a85b
}

func (any *__obf_25190d85d579f33a) ToBool() bool {
	return true
}

func (any *__obf_25190d85d579f33a) ToInt() int {
	return 0
}

func (any *__obf_25190d85d579f33a) ToInt32() int32 {
	return 0
}

func (any *__obf_25190d85d579f33a) ToInt64() int64 {
	return 0
}

func (any *__obf_25190d85d579f33a) ToUint() uint {
	return 0
}

func (any *__obf_25190d85d579f33a) ToUint32() uint32 {
	return 0
}

func (any *__obf_25190d85d579f33a) ToUint64() uint64 {
	return 0
}

func (any *__obf_25190d85d579f33a) ToFloat32() float32 {
	return 0
}

func (any *__obf_25190d85d579f33a) ToFloat64() float64 {
	return 0
}

func (any *__obf_25190d85d579f33a) ToString() string {
	__obf_a3eaafc22faf1644, __obf_ea0680f8b461a85b := MarshalToString(any.__obf_35accd096612ebe4.Interface())
	any.__obf_ea0680f8b461a85b = __obf_ea0680f8b461a85b
	return __obf_a3eaafc22faf1644
}

func (any *__obf_25190d85d579f33a) Get(__obf_5dde9fb4007294d4 ...any) Any {
	if len(__obf_5dde9fb4007294d4) == 0 {
		return any
	}
	switch __obf_a9297d670cd1f39b := __obf_5dde9fb4007294d4[0].(type) {
	case int32:
		if '*' == __obf_a9297d670cd1f39b {
			__obf_328e2515a5e98ad9 := map[string]Any{}
			for _, __obf_efcbbccd38302ccb := range any.__obf_35accd096612ebe4.MapKeys() {
				__obf_3ad1ab12c6ade438 := __obf_efcbbccd38302ccb.String()
				__obf_472279c4f949a602 := Wrap(any.__obf_35accd096612ebe4.MapIndex(__obf_efcbbccd38302ccb).Interface())
				__obf_cef0ddc31c052dab := __obf_472279c4f949a602.Get(__obf_5dde9fb4007294d4[1:]...)
				if __obf_cef0ddc31c052dab.ValueType() != InvalidValue {
					__obf_328e2515a5e98ad9[__obf_3ad1ab12c6ade438] = __obf_cef0ddc31c052dab
				}
			}
			return __obf_2259c4eb0c25660e(__obf_328e2515a5e98ad9)
		}
		return __obf_5b62b1b9d47ae70b(__obf_5dde9fb4007294d4)
	default:
		__obf_52321dce0d8db989 := any.__obf_35accd096612ebe4.MapIndex(reflect.ValueOf(__obf_a9297d670cd1f39b))
		if !__obf_52321dce0d8db989.IsValid() {
			return __obf_5b62b1b9d47ae70b(__obf_5dde9fb4007294d4)
		}
		return Wrap(__obf_52321dce0d8db989.Interface())
	}
}

func (any *__obf_25190d85d579f33a) Keys() []string {
	__obf_5308cb213e4a53a1 := make([]string, 0, any.__obf_35accd096612ebe4.Len())
	for _, __obf_efcbbccd38302ccb := range any.__obf_35accd096612ebe4.MapKeys() {
		__obf_5308cb213e4a53a1 = append(__obf_5308cb213e4a53a1, __obf_efcbbccd38302ccb.String())
	}
	return __obf_5308cb213e4a53a1
}

func (any *__obf_25190d85d579f33a) Size() int {
	return any.__obf_35accd096612ebe4.Len()
}

func (any *__obf_25190d85d579f33a) WriteTo(__obf_d8f50bcfe87d8b47 *Stream) {
	__obf_d8f50bcfe87d8b47.
		WriteVal(any.__obf_35accd096612ebe4)
}

func (any *__obf_25190d85d579f33a) GetInterface() any {
	return any.__obf_35accd096612ebe4.Interface()
}
