package __obf_c3cd04a15c56f32f

import (
	"reflect"
	"unsafe"
)

type __obf_1a2e99c79ce5a488 struct {
	__obf_86279f81acb381f3
	__obf_0c8065c219ccf0ab *__obf_3a25fb4c9a8342bb
	__obf_ace979f6622823f3 []byte
	__obf_5966ecc5edb9b17e error
}

func (any *__obf_1a2e99c79ce5a488) ValueType() ValueType {
	return ArrayValue
}

func (any *__obf_1a2e99c79ce5a488) MustBeValid() Any {
	return any
}

func (any *__obf_1a2e99c79ce5a488) LastError() error {
	return any.__obf_5966ecc5edb9b17e
}

func (any *__obf_1a2e99c79ce5a488) ToBool() bool {
	__obf_edd9fe48d39445e4 := any.__obf_0c8065c219ccf0ab.BorrowIterator(any.__obf_ace979f6622823f3)
	defer any.__obf_0c8065c219ccf0ab.ReturnIterator(__obf_edd9fe48d39445e4)
	return __obf_edd9fe48d39445e4.ReadArray()
}

func (any *__obf_1a2e99c79ce5a488) ToInt() int {
	if any.ToBool() {
		return 1
	}
	return 0
}

func (any *__obf_1a2e99c79ce5a488) ToInt32() int32 {
	if any.ToBool() {
		return 1
	}
	return 0
}

func (any *__obf_1a2e99c79ce5a488) ToInt64() int64 {
	if any.ToBool() {
		return 1
	}
	return 0
}

func (any *__obf_1a2e99c79ce5a488) ToUint() uint {
	if any.ToBool() {
		return 1
	}
	return 0
}

func (any *__obf_1a2e99c79ce5a488) ToUint32() uint32 {
	if any.ToBool() {
		return 1
	}
	return 0
}

func (any *__obf_1a2e99c79ce5a488) ToUint64() uint64 {
	if any.ToBool() {
		return 1
	}
	return 0
}

func (any *__obf_1a2e99c79ce5a488) ToFloat32() float32 {
	if any.ToBool() {
		return 1
	}
	return 0
}

func (any *__obf_1a2e99c79ce5a488) ToFloat64() float64 {
	if any.ToBool() {
		return 1
	}
	return 0
}

func (any *__obf_1a2e99c79ce5a488) ToString() string {
	return *(*string)(unsafe.Pointer(&any.__obf_ace979f6622823f3))
}

func (any *__obf_1a2e99c79ce5a488) ToVal(__obf_d59813f23ad740a8 any) {
	__obf_edd9fe48d39445e4 := any.__obf_0c8065c219ccf0ab.BorrowIterator(any.__obf_ace979f6622823f3)
	defer any.__obf_0c8065c219ccf0ab.ReturnIterator(__obf_edd9fe48d39445e4)
	__obf_edd9fe48d39445e4.
		ReadVal(__obf_d59813f23ad740a8)
}

func (any *__obf_1a2e99c79ce5a488) Get(__obf_b90e4ca332607787 ...any) Any {
	if len(__obf_b90e4ca332607787) == 0 {
		return any
	}
	switch __obf_0f79e2c35188222f := __obf_b90e4ca332607787[0].(type) {
	case int:
		__obf_edd9fe48d39445e4 := any.__obf_0c8065c219ccf0ab.BorrowIterator(any.__obf_ace979f6622823f3)
		defer any.__obf_0c8065c219ccf0ab.ReturnIterator(__obf_edd9fe48d39445e4)
		__obf_cd9b85b271fff465 := __obf_4c9d33490f37e941(__obf_edd9fe48d39445e4, __obf_0f79e2c35188222f)
		if __obf_cd9b85b271fff465 == nil {
			return __obf_5a6677fd9236ae8d(__obf_b90e4ca332607787)
		}
		__obf_edd9fe48d39445e4.
			ResetBytes(__obf_cd9b85b271fff465)
		return __obf_9979e7313b51aeca(__obf_edd9fe48d39445e4, __obf_b90e4ca332607787[1:])
	case int32:
		if '*' == __obf_0f79e2c35188222f {
			__obf_edd9fe48d39445e4 := any.__obf_0c8065c219ccf0ab.BorrowIterator(any.__obf_ace979f6622823f3)
			defer any.__obf_0c8065c219ccf0ab.ReturnIterator(__obf_edd9fe48d39445e4)
			__obf_003fd2c5d2c6b965 := make([]Any, 0)
			__obf_edd9fe48d39445e4.
				ReadArrayCB(func(__obf_edd9fe48d39445e4 *Iterator) bool {
					__obf_da9e40840d5dbfdf := __obf_edd9fe48d39445e4.__obf_1da8db8a1d168ded().Get(__obf_b90e4ca332607787[1:]...)
					if __obf_da9e40840d5dbfdf.ValueType() != InvalidValue {
						__obf_003fd2c5d2c6b965 = append(__obf_003fd2c5d2c6b965, __obf_da9e40840d5dbfdf)
					}
					return true
				})
			return __obf_90999f5d50faa290(__obf_003fd2c5d2c6b965)
		}
		return __obf_5a6677fd9236ae8d(__obf_b90e4ca332607787)
	default:
		return __obf_5a6677fd9236ae8d(__obf_b90e4ca332607787)
	}
}

func (any *__obf_1a2e99c79ce5a488) Size() int {
	__obf_ecf95be2d6e27166 := 0
	__obf_edd9fe48d39445e4 := any.__obf_0c8065c219ccf0ab.BorrowIterator(any.__obf_ace979f6622823f3)
	defer any.__obf_0c8065c219ccf0ab.ReturnIterator(__obf_edd9fe48d39445e4)
	__obf_edd9fe48d39445e4.
		ReadArrayCB(func(__obf_edd9fe48d39445e4 *Iterator) bool {
			__obf_ecf95be2d6e27166++
			__obf_edd9fe48d39445e4.
				Skip()
			return true
		})
	return __obf_ecf95be2d6e27166
}

func (any *__obf_1a2e99c79ce5a488) WriteTo(__obf_2361f5214e84c60f *Stream) {
	__obf_2361f5214e84c60f.
		Write(any.__obf_ace979f6622823f3)
}

func (any *__obf_1a2e99c79ce5a488) GetInterface() any {
	__obf_edd9fe48d39445e4 := any.__obf_0c8065c219ccf0ab.BorrowIterator(any.__obf_ace979f6622823f3)
	defer any.__obf_0c8065c219ccf0ab.ReturnIterator(__obf_edd9fe48d39445e4)
	return __obf_edd9fe48d39445e4.Read()
}

type __obf_f34b0c210608c484 struct {
	__obf_86279f81acb381f3
	__obf_d59813f23ad740a8 reflect.Value
}

func __obf_90999f5d50faa290(__obf_d59813f23ad740a8 any) *__obf_f34b0c210608c484 {
	return &__obf_f34b0c210608c484{__obf_86279f81acb381f3{}, reflect.ValueOf(__obf_d59813f23ad740a8)}
}

func (any *__obf_f34b0c210608c484) ValueType() ValueType {
	return ArrayValue
}

func (any *__obf_f34b0c210608c484) MustBeValid() Any {
	return any
}

func (any *__obf_f34b0c210608c484) LastError() error {
	return nil
}

func (any *__obf_f34b0c210608c484) ToBool() bool {
	return any.__obf_d59813f23ad740a8.Len() != 0
}

func (any *__obf_f34b0c210608c484) ToInt() int {
	if any.__obf_d59813f23ad740a8.Len() == 0 {
		return 0
	}
	return 1
}

func (any *__obf_f34b0c210608c484) ToInt32() int32 {
	if any.__obf_d59813f23ad740a8.Len() == 0 {
		return 0
	}
	return 1
}

func (any *__obf_f34b0c210608c484) ToInt64() int64 {
	if any.__obf_d59813f23ad740a8.Len() == 0 {
		return 0
	}
	return 1
}

func (any *__obf_f34b0c210608c484) ToUint() uint {
	if any.__obf_d59813f23ad740a8.Len() == 0 {
		return 0
	}
	return 1
}

func (any *__obf_f34b0c210608c484) ToUint32() uint32 {
	if any.__obf_d59813f23ad740a8.Len() == 0 {
		return 0
	}
	return 1
}

func (any *__obf_f34b0c210608c484) ToUint64() uint64 {
	if any.__obf_d59813f23ad740a8.Len() == 0 {
		return 0
	}
	return 1
}

func (any *__obf_f34b0c210608c484) ToFloat32() float32 {
	if any.__obf_d59813f23ad740a8.Len() == 0 {
		return 0
	}
	return 1
}

func (any *__obf_f34b0c210608c484) ToFloat64() float64 {
	if any.__obf_d59813f23ad740a8.Len() == 0 {
		return 0
	}
	return 1
}

func (any *__obf_f34b0c210608c484) ToString() string {
	__obf_2d944450d48e5810, _ := MarshalToString(any.__obf_d59813f23ad740a8.Interface())
	return __obf_2d944450d48e5810
}

func (any *__obf_f34b0c210608c484) Get(__obf_b90e4ca332607787 ...any) Any {
	if len(__obf_b90e4ca332607787) == 0 {
		return any
	}
	switch __obf_0f79e2c35188222f := __obf_b90e4ca332607787[0].(type) {
	case int:
		if __obf_0f79e2c35188222f < 0 || __obf_0f79e2c35188222f >= any.__obf_d59813f23ad740a8.Len() {
			return __obf_5a6677fd9236ae8d(__obf_b90e4ca332607787)
		}
		return Wrap(any.__obf_d59813f23ad740a8.Index(__obf_0f79e2c35188222f).Interface())
	case int32:
		if '*' == __obf_0f79e2c35188222f {
			__obf_db9539110ac34bf8 := make([]Any, 0)
			for __obf_28d099df85f083a8 := 0; __obf_28d099df85f083a8 < any.__obf_d59813f23ad740a8.Len(); __obf_28d099df85f083a8++ {
				__obf_7348ad94a5462ee9 := Wrap(any.__obf_d59813f23ad740a8.Index(__obf_28d099df85f083a8).Interface()).Get(__obf_b90e4ca332607787[1:]...)
				if __obf_7348ad94a5462ee9.ValueType() != InvalidValue {
					__obf_db9539110ac34bf8 = append(__obf_db9539110ac34bf8, __obf_7348ad94a5462ee9)
				}
			}
			return __obf_90999f5d50faa290(__obf_db9539110ac34bf8)
		}
		return __obf_5a6677fd9236ae8d(__obf_b90e4ca332607787)
	default:
		return __obf_5a6677fd9236ae8d(__obf_b90e4ca332607787)
	}
}

func (any *__obf_f34b0c210608c484) Size() int {
	return any.__obf_d59813f23ad740a8.Len()
}

func (any *__obf_f34b0c210608c484) WriteTo(__obf_2361f5214e84c60f *Stream) {
	__obf_2361f5214e84c60f.
		WriteVal(any.__obf_d59813f23ad740a8)
}

func (any *__obf_f34b0c210608c484) GetInterface() any {
	return any.__obf_d59813f23ad740a8.Interface()
}
