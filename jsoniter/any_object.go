package __obf_5b802ce8d9ba56d6

import (
	"reflect"
	"unsafe"
)

type __obf_92721e3ae32bb50a struct {
	__obf_fb75d4e4562132ae
	__obf_dca106e1cdf85392 *__obf_5d13d7f3bd06c6cf
	__obf_9fc06d9180f0daca []byte
	__obf_6d071d48f3b321ad error
}

func (any *__obf_92721e3ae32bb50a) ValueType() ValueType {
	return ObjectValue
}

func (any *__obf_92721e3ae32bb50a) MustBeValid() Any {
	return any
}

func (any *__obf_92721e3ae32bb50a) LastError() error {
	return any.__obf_6d071d48f3b321ad
}

func (any *__obf_92721e3ae32bb50a) ToBool() bool {
	return true
}

func (any *__obf_92721e3ae32bb50a) ToInt() int {
	return 0
}

func (any *__obf_92721e3ae32bb50a) ToInt32() int32 {
	return 0
}

func (any *__obf_92721e3ae32bb50a) ToInt64() int64 {
	return 0
}

func (any *__obf_92721e3ae32bb50a) ToUint() uint {
	return 0
}

func (any *__obf_92721e3ae32bb50a) ToUint32() uint32 {
	return 0
}

func (any *__obf_92721e3ae32bb50a) ToUint64() uint64 {
	return 0
}

func (any *__obf_92721e3ae32bb50a) ToFloat32() float32 {
	return 0
}

func (any *__obf_92721e3ae32bb50a) ToFloat64() float64 {
	return 0
}

func (any *__obf_92721e3ae32bb50a) ToString() string {
	return *(*string)(unsafe.Pointer(&any.__obf_9fc06d9180f0daca))
}

func (any *__obf_92721e3ae32bb50a) ToVal(__obf_f9b1526531f3c024 any) {
	__obf_67008a6a9e5ba828 := any.__obf_dca106e1cdf85392.BorrowIterator(any.__obf_9fc06d9180f0daca)
	defer any.__obf_dca106e1cdf85392.ReturnIterator(__obf_67008a6a9e5ba828)
	__obf_67008a6a9e5ba828.
		ReadVal(__obf_f9b1526531f3c024)
}

func (any *__obf_92721e3ae32bb50a) Get(__obf_c5d71353f4393b3c ...any) Any {
	if len(__obf_c5d71353f4393b3c) == 0 {
		return any
	}
	switch __obf_837762e0f8c45322 := __obf_c5d71353f4393b3c[0].(type) {
	case string:
		__obf_67008a6a9e5ba828 := any.__obf_dca106e1cdf85392.BorrowIterator(any.__obf_9fc06d9180f0daca)
		defer any.__obf_dca106e1cdf85392.ReturnIterator(__obf_67008a6a9e5ba828)
		__obf_611f49298438d349 := __obf_51a7c2c2318325e4(__obf_67008a6a9e5ba828, __obf_837762e0f8c45322)
		if __obf_611f49298438d349 == nil {
			return __obf_e990980207770dd2(__obf_c5d71353f4393b3c)
		}
		__obf_67008a6a9e5ba828.
			ResetBytes(__obf_611f49298438d349)
		return __obf_9c645421e8d71b4a(__obf_67008a6a9e5ba828, __obf_c5d71353f4393b3c[1:])
	case int32:
		if '*' == __obf_837762e0f8c45322 {
			__obf_4615d688422b7c63 := map[string]Any{}
			__obf_67008a6a9e5ba828 := any.__obf_dca106e1cdf85392.BorrowIterator(any.__obf_9fc06d9180f0daca)
			defer any.__obf_dca106e1cdf85392.ReturnIterator(__obf_67008a6a9e5ba828)
			__obf_67008a6a9e5ba828.
				ReadMapCB(func(__obf_67008a6a9e5ba828 *Iterator, __obf_61998fb083387c3c string) bool {
					__obf_6b02486c8eab4fbc := __obf_9c645421e8d71b4a(__obf_67008a6a9e5ba828, __obf_c5d71353f4393b3c[1:])
					if __obf_6b02486c8eab4fbc.ValueType() != InvalidValue {
						__obf_4615d688422b7c63[__obf_61998fb083387c3c] = __obf_6b02486c8eab4fbc
					}
					return true
				})
			return __obf_03d594f9555cce37(__obf_4615d688422b7c63)
		}
		return __obf_e990980207770dd2(__obf_c5d71353f4393b3c)
	default:
		return __obf_e990980207770dd2(__obf_c5d71353f4393b3c)
	}
}

func (any *__obf_92721e3ae32bb50a) Keys() []string {
	__obf_9dcce33a678d336b := []string{}
	__obf_67008a6a9e5ba828 := any.__obf_dca106e1cdf85392.BorrowIterator(any.__obf_9fc06d9180f0daca)
	defer any.__obf_dca106e1cdf85392.ReturnIterator(__obf_67008a6a9e5ba828)
	__obf_67008a6a9e5ba828.
		ReadMapCB(func(__obf_67008a6a9e5ba828 *Iterator, __obf_61998fb083387c3c string) bool {
			__obf_67008a6a9e5ba828.
				Skip()
			__obf_9dcce33a678d336b = append(__obf_9dcce33a678d336b, __obf_61998fb083387c3c)
			return true
		})
	return __obf_9dcce33a678d336b
}

func (any *__obf_92721e3ae32bb50a) Size() int {
	__obf_aedccbde4481459c := 0
	__obf_67008a6a9e5ba828 := any.__obf_dca106e1cdf85392.BorrowIterator(any.__obf_9fc06d9180f0daca)
	defer any.__obf_dca106e1cdf85392.ReturnIterator(__obf_67008a6a9e5ba828)
	__obf_67008a6a9e5ba828.
		ReadObjectCB(func(__obf_67008a6a9e5ba828 *Iterator, __obf_61998fb083387c3c string) bool {
			__obf_67008a6a9e5ba828.
				Skip()
			__obf_aedccbde4481459c++
			return true
		})
	return __obf_aedccbde4481459c
}

func (any *__obf_92721e3ae32bb50a) WriteTo(__obf_00fc491caa967a74 *Stream) {
	__obf_00fc491caa967a74.
		Write(any.__obf_9fc06d9180f0daca)
}

func (any *__obf_92721e3ae32bb50a) GetInterface() any {
	__obf_67008a6a9e5ba828 := any.__obf_dca106e1cdf85392.BorrowIterator(any.__obf_9fc06d9180f0daca)
	defer any.__obf_dca106e1cdf85392.ReturnIterator(__obf_67008a6a9e5ba828)
	return __obf_67008a6a9e5ba828.Read()
}

type __obf_f69d5a3c92902506 struct {
	__obf_fb75d4e4562132ae
	__obf_6d071d48f3b321ad error
	__obf_5406b11e33b9d1d3 reflect.Value
}

func __obf_4b9ae4bb2bbc6708(__obf_5406b11e33b9d1d3 any) *__obf_f69d5a3c92902506 {
	return &__obf_f69d5a3c92902506{__obf_fb75d4e4562132ae{}, nil, reflect.ValueOf(__obf_5406b11e33b9d1d3)}
}

func (any *__obf_f69d5a3c92902506) ValueType() ValueType {
	return ObjectValue
}

func (any *__obf_f69d5a3c92902506) MustBeValid() Any {
	return any
}

func (any *__obf_f69d5a3c92902506) Parse() *Iterator {
	return nil
}

func (any *__obf_f69d5a3c92902506) LastError() error {
	return any.__obf_6d071d48f3b321ad
}

func (any *__obf_f69d5a3c92902506) ToBool() bool {
	return any.__obf_5406b11e33b9d1d3.NumField() != 0
}

func (any *__obf_f69d5a3c92902506) ToInt() int {
	return 0
}

func (any *__obf_f69d5a3c92902506) ToInt32() int32 {
	return 0
}

func (any *__obf_f69d5a3c92902506) ToInt64() int64 {
	return 0
}

func (any *__obf_f69d5a3c92902506) ToUint() uint {
	return 0
}

func (any *__obf_f69d5a3c92902506) ToUint32() uint32 {
	return 0
}

func (any *__obf_f69d5a3c92902506) ToUint64() uint64 {
	return 0
}

func (any *__obf_f69d5a3c92902506) ToFloat32() float32 {
	return 0
}

func (any *__obf_f69d5a3c92902506) ToFloat64() float64 {
	return 0
}

func (any *__obf_f69d5a3c92902506) ToString() string {
	__obf_12c21b79fa86dcba, __obf_6d071d48f3b321ad := MarshalToString(any.__obf_5406b11e33b9d1d3.Interface())
	any.__obf_6d071d48f3b321ad = __obf_6d071d48f3b321ad
	return __obf_12c21b79fa86dcba
}

func (any *__obf_f69d5a3c92902506) Get(__obf_c5d71353f4393b3c ...any) Any {
	if len(__obf_c5d71353f4393b3c) == 0 {
		return any
	}
	switch __obf_837762e0f8c45322 := __obf_c5d71353f4393b3c[0].(type) {
	case string:
		__obf_61998fb083387c3c := any.__obf_5406b11e33b9d1d3.FieldByName(__obf_837762e0f8c45322)
		if !__obf_61998fb083387c3c.IsValid() {
			return __obf_e990980207770dd2(__obf_c5d71353f4393b3c)
		}
		return Wrap(__obf_61998fb083387c3c.Interface())
	case int32:
		if '*' == __obf_837762e0f8c45322 {
			__obf_4615d688422b7c63 := map[string]Any{}
			for __obf_2deec7c38ffb6ae3 := 0; __obf_2deec7c38ffb6ae3 < any.__obf_5406b11e33b9d1d3.NumField(); __obf_2deec7c38ffb6ae3++ {
				__obf_61998fb083387c3c := any.__obf_5406b11e33b9d1d3.Field(__obf_2deec7c38ffb6ae3)
				if __obf_61998fb083387c3c.CanInterface() {
					__obf_6b02486c8eab4fbc := Wrap(__obf_61998fb083387c3c.Interface()).Get(__obf_c5d71353f4393b3c[1:]...)
					if __obf_6b02486c8eab4fbc.ValueType() != InvalidValue {
						__obf_4615d688422b7c63[any.__obf_5406b11e33b9d1d3.Type().Field(__obf_2deec7c38ffb6ae3).Name] = __obf_6b02486c8eab4fbc
					}
				}
			}
			return __obf_03d594f9555cce37(__obf_4615d688422b7c63)
		}
		return __obf_e990980207770dd2(__obf_c5d71353f4393b3c)
	default:
		return __obf_e990980207770dd2(__obf_c5d71353f4393b3c)
	}
}

func (any *__obf_f69d5a3c92902506) Keys() []string {
	__obf_9dcce33a678d336b := make([]string, 0, any.__obf_5406b11e33b9d1d3.NumField())
	for __obf_2deec7c38ffb6ae3 := 0; __obf_2deec7c38ffb6ae3 < any.__obf_5406b11e33b9d1d3.NumField(); __obf_2deec7c38ffb6ae3++ {
		__obf_9dcce33a678d336b = append(__obf_9dcce33a678d336b, any.__obf_5406b11e33b9d1d3.Type().Field(__obf_2deec7c38ffb6ae3).Name)
	}
	return __obf_9dcce33a678d336b
}

func (any *__obf_f69d5a3c92902506) Size() int {
	return any.__obf_5406b11e33b9d1d3.NumField()
}

func (any *__obf_f69d5a3c92902506) WriteTo(__obf_00fc491caa967a74 *Stream) {
	__obf_00fc491caa967a74.
		WriteVal(any.__obf_5406b11e33b9d1d3)
}

func (any *__obf_f69d5a3c92902506) GetInterface() any {
	return any.__obf_5406b11e33b9d1d3.Interface()
}

type __obf_bc1094b03f59c52b struct {
	__obf_fb75d4e4562132ae
	__obf_6d071d48f3b321ad error
	__obf_5406b11e33b9d1d3 reflect.Value
}

func __obf_03d594f9555cce37(__obf_5406b11e33b9d1d3 any) *__obf_bc1094b03f59c52b {
	return &__obf_bc1094b03f59c52b{__obf_fb75d4e4562132ae{}, nil, reflect.ValueOf(__obf_5406b11e33b9d1d3)}
}

func (any *__obf_bc1094b03f59c52b) ValueType() ValueType {
	return ObjectValue
}

func (any *__obf_bc1094b03f59c52b) MustBeValid() Any {
	return any
}

func (any *__obf_bc1094b03f59c52b) Parse() *Iterator {
	return nil
}

func (any *__obf_bc1094b03f59c52b) LastError() error {
	return any.__obf_6d071d48f3b321ad
}

func (any *__obf_bc1094b03f59c52b) ToBool() bool {
	return true
}

func (any *__obf_bc1094b03f59c52b) ToInt() int {
	return 0
}

func (any *__obf_bc1094b03f59c52b) ToInt32() int32 {
	return 0
}

func (any *__obf_bc1094b03f59c52b) ToInt64() int64 {
	return 0
}

func (any *__obf_bc1094b03f59c52b) ToUint() uint {
	return 0
}

func (any *__obf_bc1094b03f59c52b) ToUint32() uint32 {
	return 0
}

func (any *__obf_bc1094b03f59c52b) ToUint64() uint64 {
	return 0
}

func (any *__obf_bc1094b03f59c52b) ToFloat32() float32 {
	return 0
}

func (any *__obf_bc1094b03f59c52b) ToFloat64() float64 {
	return 0
}

func (any *__obf_bc1094b03f59c52b) ToString() string {
	__obf_12c21b79fa86dcba, __obf_6d071d48f3b321ad := MarshalToString(any.__obf_5406b11e33b9d1d3.Interface())
	any.__obf_6d071d48f3b321ad = __obf_6d071d48f3b321ad
	return __obf_12c21b79fa86dcba
}

func (any *__obf_bc1094b03f59c52b) Get(__obf_c5d71353f4393b3c ...any) Any {
	if len(__obf_c5d71353f4393b3c) == 0 {
		return any
	}
	switch __obf_837762e0f8c45322 := __obf_c5d71353f4393b3c[0].(type) {
	case int32:
		if '*' == __obf_837762e0f8c45322 {
			__obf_4615d688422b7c63 := map[string]Any{}
			for _, __obf_6998442f85330120 := range any.__obf_5406b11e33b9d1d3.MapKeys() {
				__obf_91a752381e0ce24a := __obf_6998442f85330120.String()
				__obf_d35a0de47a035484 := Wrap(any.__obf_5406b11e33b9d1d3.MapIndex(__obf_6998442f85330120).Interface())
				__obf_6b02486c8eab4fbc := __obf_d35a0de47a035484.Get(__obf_c5d71353f4393b3c[1:]...)
				if __obf_6b02486c8eab4fbc.ValueType() != InvalidValue {
					__obf_4615d688422b7c63[__obf_91a752381e0ce24a] = __obf_6b02486c8eab4fbc
				}
			}
			return __obf_03d594f9555cce37(__obf_4615d688422b7c63)
		}
		return __obf_e990980207770dd2(__obf_c5d71353f4393b3c)
	default:
		__obf_c949477a4af2082d := any.__obf_5406b11e33b9d1d3.MapIndex(reflect.ValueOf(__obf_837762e0f8c45322))
		if !__obf_c949477a4af2082d.IsValid() {
			return __obf_e990980207770dd2(__obf_c5d71353f4393b3c)
		}
		return Wrap(__obf_c949477a4af2082d.Interface())
	}
}

func (any *__obf_bc1094b03f59c52b) Keys() []string {
	__obf_9dcce33a678d336b := make([]string, 0, any.__obf_5406b11e33b9d1d3.Len())
	for _, __obf_6998442f85330120 := range any.__obf_5406b11e33b9d1d3.MapKeys() {
		__obf_9dcce33a678d336b = append(__obf_9dcce33a678d336b, __obf_6998442f85330120.String())
	}
	return __obf_9dcce33a678d336b
}

func (any *__obf_bc1094b03f59c52b) Size() int {
	return any.__obf_5406b11e33b9d1d3.Len()
}

func (any *__obf_bc1094b03f59c52b) WriteTo(__obf_00fc491caa967a74 *Stream) {
	__obf_00fc491caa967a74.
		WriteVal(any.__obf_5406b11e33b9d1d3)
}

func (any *__obf_bc1094b03f59c52b) GetInterface() any {
	return any.__obf_5406b11e33b9d1d3.Interface()
}
