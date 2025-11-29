package __obf_5b802ce8d9ba56d6

import (
	"reflect"
	"unsafe"
)

type __obf_41614ecb47fbcff9 struct {
	__obf_fb75d4e4562132ae
	__obf_dca106e1cdf85392 *__obf_5d13d7f3bd06c6cf
	__obf_9fc06d9180f0daca []byte
	__obf_6d071d48f3b321ad error
}

func (any *__obf_41614ecb47fbcff9) ValueType() ValueType {
	return ArrayValue
}

func (any *__obf_41614ecb47fbcff9) MustBeValid() Any {
	return any
}

func (any *__obf_41614ecb47fbcff9) LastError() error {
	return any.__obf_6d071d48f3b321ad
}

func (any *__obf_41614ecb47fbcff9) ToBool() bool {
	__obf_67008a6a9e5ba828 := any.__obf_dca106e1cdf85392.BorrowIterator(any.__obf_9fc06d9180f0daca)
	defer any.__obf_dca106e1cdf85392.ReturnIterator(__obf_67008a6a9e5ba828)
	return __obf_67008a6a9e5ba828.ReadArray()
}

func (any *__obf_41614ecb47fbcff9) ToInt() int {
	if any.ToBool() {
		return 1
	}
	return 0
}

func (any *__obf_41614ecb47fbcff9) ToInt32() int32 {
	if any.ToBool() {
		return 1
	}
	return 0
}

func (any *__obf_41614ecb47fbcff9) ToInt64() int64 {
	if any.ToBool() {
		return 1
	}
	return 0
}

func (any *__obf_41614ecb47fbcff9) ToUint() uint {
	if any.ToBool() {
		return 1
	}
	return 0
}

func (any *__obf_41614ecb47fbcff9) ToUint32() uint32 {
	if any.ToBool() {
		return 1
	}
	return 0
}

func (any *__obf_41614ecb47fbcff9) ToUint64() uint64 {
	if any.ToBool() {
		return 1
	}
	return 0
}

func (any *__obf_41614ecb47fbcff9) ToFloat32() float32 {
	if any.ToBool() {
		return 1
	}
	return 0
}

func (any *__obf_41614ecb47fbcff9) ToFloat64() float64 {
	if any.ToBool() {
		return 1
	}
	return 0
}

func (any *__obf_41614ecb47fbcff9) ToString() string {
	return *(*string)(unsafe.Pointer(&any.__obf_9fc06d9180f0daca))
}

func (any *__obf_41614ecb47fbcff9) ToVal(__obf_5406b11e33b9d1d3 any) {
	__obf_67008a6a9e5ba828 := any.__obf_dca106e1cdf85392.BorrowIterator(any.__obf_9fc06d9180f0daca)
	defer any.__obf_dca106e1cdf85392.ReturnIterator(__obf_67008a6a9e5ba828)
	__obf_67008a6a9e5ba828.
		ReadVal(__obf_5406b11e33b9d1d3)
}

func (any *__obf_41614ecb47fbcff9) Get(__obf_c5d71353f4393b3c ...any) Any {
	if len(__obf_c5d71353f4393b3c) == 0 {
		return any
	}
	switch __obf_837762e0f8c45322 := __obf_c5d71353f4393b3c[0].(type) {
	case int:
		__obf_67008a6a9e5ba828 := any.__obf_dca106e1cdf85392.BorrowIterator(any.__obf_9fc06d9180f0daca)
		defer any.__obf_dca106e1cdf85392.ReturnIterator(__obf_67008a6a9e5ba828)
		__obf_611f49298438d349 := __obf_e8d51e45c7dc6fc5(__obf_67008a6a9e5ba828, __obf_837762e0f8c45322)
		if __obf_611f49298438d349 == nil {
			return __obf_e990980207770dd2(__obf_c5d71353f4393b3c)
		}
		__obf_67008a6a9e5ba828.
			ResetBytes(__obf_611f49298438d349)
		return __obf_9c645421e8d71b4a(__obf_67008a6a9e5ba828, __obf_c5d71353f4393b3c[1:])
	case int32:
		if '*' == __obf_837762e0f8c45322 {
			__obf_67008a6a9e5ba828 := any.__obf_dca106e1cdf85392.BorrowIterator(any.__obf_9fc06d9180f0daca)
			defer any.__obf_dca106e1cdf85392.ReturnIterator(__obf_67008a6a9e5ba828)
			__obf_0ac6bb318bf53d42 := make([]Any, 0)
			__obf_67008a6a9e5ba828.
				ReadArrayCB(func(__obf_67008a6a9e5ba828 *Iterator) bool {
					__obf_98a3eefc4c187186 := __obf_67008a6a9e5ba828.__obf_f0a5bf7546ffe762().Get(__obf_c5d71353f4393b3c[1:]...)
					if __obf_98a3eefc4c187186.ValueType() != InvalidValue {
						__obf_0ac6bb318bf53d42 = append(__obf_0ac6bb318bf53d42, __obf_98a3eefc4c187186)
					}
					return true
				})
			return __obf_a6d72607e427694a(__obf_0ac6bb318bf53d42)
		}
		return __obf_e990980207770dd2(__obf_c5d71353f4393b3c)
	default:
		return __obf_e990980207770dd2(__obf_c5d71353f4393b3c)
	}
}

func (any *__obf_41614ecb47fbcff9) Size() int {
	__obf_aedccbde4481459c := 0
	__obf_67008a6a9e5ba828 := any.__obf_dca106e1cdf85392.BorrowIterator(any.__obf_9fc06d9180f0daca)
	defer any.__obf_dca106e1cdf85392.ReturnIterator(__obf_67008a6a9e5ba828)
	__obf_67008a6a9e5ba828.
		ReadArrayCB(func(__obf_67008a6a9e5ba828 *Iterator) bool {
			__obf_aedccbde4481459c++
			__obf_67008a6a9e5ba828.
				Skip()
			return true
		})
	return __obf_aedccbde4481459c
}

func (any *__obf_41614ecb47fbcff9) WriteTo(__obf_00fc491caa967a74 *Stream) {
	__obf_00fc491caa967a74.
		Write(any.__obf_9fc06d9180f0daca)
}

func (any *__obf_41614ecb47fbcff9) GetInterface() any {
	__obf_67008a6a9e5ba828 := any.__obf_dca106e1cdf85392.BorrowIterator(any.__obf_9fc06d9180f0daca)
	defer any.__obf_dca106e1cdf85392.ReturnIterator(__obf_67008a6a9e5ba828)
	return __obf_67008a6a9e5ba828.Read()
}

type __obf_db252f52bd20d21d struct {
	__obf_fb75d4e4562132ae
	__obf_5406b11e33b9d1d3 reflect.Value
}

func __obf_a6d72607e427694a(__obf_5406b11e33b9d1d3 any) *__obf_db252f52bd20d21d {
	return &__obf_db252f52bd20d21d{__obf_fb75d4e4562132ae{}, reflect.ValueOf(__obf_5406b11e33b9d1d3)}
}

func (any *__obf_db252f52bd20d21d) ValueType() ValueType {
	return ArrayValue
}

func (any *__obf_db252f52bd20d21d) MustBeValid() Any {
	return any
}

func (any *__obf_db252f52bd20d21d) LastError() error {
	return nil
}

func (any *__obf_db252f52bd20d21d) ToBool() bool {
	return any.__obf_5406b11e33b9d1d3.Len() != 0
}

func (any *__obf_db252f52bd20d21d) ToInt() int {
	if any.__obf_5406b11e33b9d1d3.Len() == 0 {
		return 0
	}
	return 1
}

func (any *__obf_db252f52bd20d21d) ToInt32() int32 {
	if any.__obf_5406b11e33b9d1d3.Len() == 0 {
		return 0
	}
	return 1
}

func (any *__obf_db252f52bd20d21d) ToInt64() int64 {
	if any.__obf_5406b11e33b9d1d3.Len() == 0 {
		return 0
	}
	return 1
}

func (any *__obf_db252f52bd20d21d) ToUint() uint {
	if any.__obf_5406b11e33b9d1d3.Len() == 0 {
		return 0
	}
	return 1
}

func (any *__obf_db252f52bd20d21d) ToUint32() uint32 {
	if any.__obf_5406b11e33b9d1d3.Len() == 0 {
		return 0
	}
	return 1
}

func (any *__obf_db252f52bd20d21d) ToUint64() uint64 {
	if any.__obf_5406b11e33b9d1d3.Len() == 0 {
		return 0
	}
	return 1
}

func (any *__obf_db252f52bd20d21d) ToFloat32() float32 {
	if any.__obf_5406b11e33b9d1d3.Len() == 0 {
		return 0
	}
	return 1
}

func (any *__obf_db252f52bd20d21d) ToFloat64() float64 {
	if any.__obf_5406b11e33b9d1d3.Len() == 0 {
		return 0
	}
	return 1
}

func (any *__obf_db252f52bd20d21d) ToString() string {
	__obf_12c21b79fa86dcba, _ := MarshalToString(any.__obf_5406b11e33b9d1d3.Interface())
	return __obf_12c21b79fa86dcba
}

func (any *__obf_db252f52bd20d21d) Get(__obf_c5d71353f4393b3c ...any) Any {
	if len(__obf_c5d71353f4393b3c) == 0 {
		return any
	}
	switch __obf_837762e0f8c45322 := __obf_c5d71353f4393b3c[0].(type) {
	case int:
		if __obf_837762e0f8c45322 < 0 || __obf_837762e0f8c45322 >= any.__obf_5406b11e33b9d1d3.Len() {
			return __obf_e990980207770dd2(__obf_c5d71353f4393b3c)
		}
		return Wrap(any.__obf_5406b11e33b9d1d3.Index(__obf_837762e0f8c45322).Interface())
	case int32:
		if '*' == __obf_837762e0f8c45322 {
			__obf_4615d688422b7c63 := make([]Any, 0)
			for __obf_2deec7c38ffb6ae3 := 0; __obf_2deec7c38ffb6ae3 < any.__obf_5406b11e33b9d1d3.Len(); __obf_2deec7c38ffb6ae3++ {
				__obf_6b02486c8eab4fbc := Wrap(any.__obf_5406b11e33b9d1d3.Index(__obf_2deec7c38ffb6ae3).Interface()).Get(__obf_c5d71353f4393b3c[1:]...)
				if __obf_6b02486c8eab4fbc.ValueType() != InvalidValue {
					__obf_4615d688422b7c63 = append(__obf_4615d688422b7c63, __obf_6b02486c8eab4fbc)
				}
			}
			return __obf_a6d72607e427694a(__obf_4615d688422b7c63)
		}
		return __obf_e990980207770dd2(__obf_c5d71353f4393b3c)
	default:
		return __obf_e990980207770dd2(__obf_c5d71353f4393b3c)
	}
}

func (any *__obf_db252f52bd20d21d) Size() int {
	return any.__obf_5406b11e33b9d1d3.Len()
}

func (any *__obf_db252f52bd20d21d) WriteTo(__obf_00fc491caa967a74 *Stream) {
	__obf_00fc491caa967a74.
		WriteVal(any.__obf_5406b11e33b9d1d3)
}

func (any *__obf_db252f52bd20d21d) GetInterface() any {
	return any.__obf_5406b11e33b9d1d3.Interface()
}
