package __obf_c3cd04a15c56f32f

import (
	"reflect"
	"unsafe"
)

type __obf_0caad688333e6f4f struct {
	__obf_86279f81acb381f3
	__obf_0c8065c219ccf0ab *__obf_3a25fb4c9a8342bb
	__obf_ace979f6622823f3 []byte
	__obf_5966ecc5edb9b17e error
}

func (any *__obf_0caad688333e6f4f) ValueType() ValueType {
	return ObjectValue
}

func (any *__obf_0caad688333e6f4f) MustBeValid() Any {
	return any
}

func (any *__obf_0caad688333e6f4f) LastError() error {
	return any.__obf_5966ecc5edb9b17e
}

func (any *__obf_0caad688333e6f4f) ToBool() bool {
	return true
}

func (any *__obf_0caad688333e6f4f) ToInt() int {
	return 0
}

func (any *__obf_0caad688333e6f4f) ToInt32() int32 {
	return 0
}

func (any *__obf_0caad688333e6f4f) ToInt64() int64 {
	return 0
}

func (any *__obf_0caad688333e6f4f) ToUint() uint {
	return 0
}

func (any *__obf_0caad688333e6f4f) ToUint32() uint32 {
	return 0
}

func (any *__obf_0caad688333e6f4f) ToUint64() uint64 {
	return 0
}

func (any *__obf_0caad688333e6f4f) ToFloat32() float32 {
	return 0
}

func (any *__obf_0caad688333e6f4f) ToFloat64() float64 {
	return 0
}

func (any *__obf_0caad688333e6f4f) ToString() string {
	return *(*string)(unsafe.Pointer(&any.__obf_ace979f6622823f3))
}

func (any *__obf_0caad688333e6f4f) ToVal(__obf_50acae8c0a871ba1 any) {
	__obf_edd9fe48d39445e4 := any.__obf_0c8065c219ccf0ab.BorrowIterator(any.__obf_ace979f6622823f3)
	defer any.__obf_0c8065c219ccf0ab.ReturnIterator(__obf_edd9fe48d39445e4)
	__obf_edd9fe48d39445e4.
		ReadVal(__obf_50acae8c0a871ba1)
}

func (any *__obf_0caad688333e6f4f) Get(__obf_b90e4ca332607787 ...any) Any {
	if len(__obf_b90e4ca332607787) == 0 {
		return any
	}
	switch __obf_0f79e2c35188222f := __obf_b90e4ca332607787[0].(type) {
	case string:
		__obf_edd9fe48d39445e4 := any.__obf_0c8065c219ccf0ab.BorrowIterator(any.__obf_ace979f6622823f3)
		defer any.__obf_0c8065c219ccf0ab.ReturnIterator(__obf_edd9fe48d39445e4)
		__obf_cd9b85b271fff465 := __obf_931a5a6194f47dee(__obf_edd9fe48d39445e4, __obf_0f79e2c35188222f)
		if __obf_cd9b85b271fff465 == nil {
			return __obf_5a6677fd9236ae8d(__obf_b90e4ca332607787)
		}
		__obf_edd9fe48d39445e4.
			ResetBytes(__obf_cd9b85b271fff465)
		return __obf_9979e7313b51aeca(__obf_edd9fe48d39445e4, __obf_b90e4ca332607787[1:])
	case int32:
		if '*' == __obf_0f79e2c35188222f {
			__obf_db9539110ac34bf8 := map[string]Any{}
			__obf_edd9fe48d39445e4 := any.__obf_0c8065c219ccf0ab.BorrowIterator(any.__obf_ace979f6622823f3)
			defer any.__obf_0c8065c219ccf0ab.ReturnIterator(__obf_edd9fe48d39445e4)
			__obf_edd9fe48d39445e4.
				ReadMapCB(func(__obf_edd9fe48d39445e4 *Iterator, __obf_f992c91036745ca0 string) bool {
					__obf_7348ad94a5462ee9 := __obf_9979e7313b51aeca(__obf_edd9fe48d39445e4, __obf_b90e4ca332607787[1:])
					if __obf_7348ad94a5462ee9.ValueType() != InvalidValue {
						__obf_db9539110ac34bf8[__obf_f992c91036745ca0] = __obf_7348ad94a5462ee9
					}
					return true
				})
			return __obf_af999c25758b0ff1(__obf_db9539110ac34bf8)
		}
		return __obf_5a6677fd9236ae8d(__obf_b90e4ca332607787)
	default:
		return __obf_5a6677fd9236ae8d(__obf_b90e4ca332607787)
	}
}

func (any *__obf_0caad688333e6f4f) Keys() []string {
	__obf_870ded78ba396f54 := []string{}
	__obf_edd9fe48d39445e4 := any.__obf_0c8065c219ccf0ab.BorrowIterator(any.__obf_ace979f6622823f3)
	defer any.__obf_0c8065c219ccf0ab.ReturnIterator(__obf_edd9fe48d39445e4)
	__obf_edd9fe48d39445e4.
		ReadMapCB(func(__obf_edd9fe48d39445e4 *Iterator, __obf_f992c91036745ca0 string) bool {
			__obf_edd9fe48d39445e4.
				Skip()
			__obf_870ded78ba396f54 = append(__obf_870ded78ba396f54, __obf_f992c91036745ca0)
			return true
		})
	return __obf_870ded78ba396f54
}

func (any *__obf_0caad688333e6f4f) Size() int {
	__obf_ecf95be2d6e27166 := 0
	__obf_edd9fe48d39445e4 := any.__obf_0c8065c219ccf0ab.BorrowIterator(any.__obf_ace979f6622823f3)
	defer any.__obf_0c8065c219ccf0ab.ReturnIterator(__obf_edd9fe48d39445e4)
	__obf_edd9fe48d39445e4.
		ReadObjectCB(func(__obf_edd9fe48d39445e4 *Iterator, __obf_f992c91036745ca0 string) bool {
			__obf_edd9fe48d39445e4.
				Skip()
			__obf_ecf95be2d6e27166++
			return true
		})
	return __obf_ecf95be2d6e27166
}

func (any *__obf_0caad688333e6f4f) WriteTo(__obf_2361f5214e84c60f *Stream) {
	__obf_2361f5214e84c60f.
		Write(any.__obf_ace979f6622823f3)
}

func (any *__obf_0caad688333e6f4f) GetInterface() any {
	__obf_edd9fe48d39445e4 := any.__obf_0c8065c219ccf0ab.BorrowIterator(any.__obf_ace979f6622823f3)
	defer any.__obf_0c8065c219ccf0ab.ReturnIterator(__obf_edd9fe48d39445e4)
	return __obf_edd9fe48d39445e4.Read()
}

type __obf_d2e8e028a5c60fc6 struct {
	__obf_86279f81acb381f3
	__obf_5966ecc5edb9b17e error
	__obf_d59813f23ad740a8 reflect.Value
}

func __obf_5ed91bdc7ddeb0d6(__obf_d59813f23ad740a8 any) *__obf_d2e8e028a5c60fc6 {
	return &__obf_d2e8e028a5c60fc6{__obf_86279f81acb381f3{}, nil, reflect.ValueOf(__obf_d59813f23ad740a8)}
}

func (any *__obf_d2e8e028a5c60fc6) ValueType() ValueType {
	return ObjectValue
}

func (any *__obf_d2e8e028a5c60fc6) MustBeValid() Any {
	return any
}

func (any *__obf_d2e8e028a5c60fc6) Parse() *Iterator {
	return nil
}

func (any *__obf_d2e8e028a5c60fc6) LastError() error {
	return any.__obf_5966ecc5edb9b17e
}

func (any *__obf_d2e8e028a5c60fc6) ToBool() bool {
	return any.__obf_d59813f23ad740a8.NumField() != 0
}

func (any *__obf_d2e8e028a5c60fc6) ToInt() int {
	return 0
}

func (any *__obf_d2e8e028a5c60fc6) ToInt32() int32 {
	return 0
}

func (any *__obf_d2e8e028a5c60fc6) ToInt64() int64 {
	return 0
}

func (any *__obf_d2e8e028a5c60fc6) ToUint() uint {
	return 0
}

func (any *__obf_d2e8e028a5c60fc6) ToUint32() uint32 {
	return 0
}

func (any *__obf_d2e8e028a5c60fc6) ToUint64() uint64 {
	return 0
}

func (any *__obf_d2e8e028a5c60fc6) ToFloat32() float32 {
	return 0
}

func (any *__obf_d2e8e028a5c60fc6) ToFloat64() float64 {
	return 0
}

func (any *__obf_d2e8e028a5c60fc6) ToString() string {
	__obf_2d944450d48e5810, __obf_5966ecc5edb9b17e := MarshalToString(any.__obf_d59813f23ad740a8.Interface())
	any.__obf_5966ecc5edb9b17e = __obf_5966ecc5edb9b17e
	return __obf_2d944450d48e5810
}

func (any *__obf_d2e8e028a5c60fc6) Get(__obf_b90e4ca332607787 ...any) Any {
	if len(__obf_b90e4ca332607787) == 0 {
		return any
	}
	switch __obf_0f79e2c35188222f := __obf_b90e4ca332607787[0].(type) {
	case string:
		__obf_f992c91036745ca0 := any.__obf_d59813f23ad740a8.FieldByName(__obf_0f79e2c35188222f)
		if !__obf_f992c91036745ca0.IsValid() {
			return __obf_5a6677fd9236ae8d(__obf_b90e4ca332607787)
		}
		return Wrap(__obf_f992c91036745ca0.Interface())
	case int32:
		if '*' == __obf_0f79e2c35188222f {
			__obf_db9539110ac34bf8 := map[string]Any{}
			for __obf_28d099df85f083a8 := 0; __obf_28d099df85f083a8 < any.__obf_d59813f23ad740a8.NumField(); __obf_28d099df85f083a8++ {
				__obf_f992c91036745ca0 := any.__obf_d59813f23ad740a8.Field(__obf_28d099df85f083a8)
				if __obf_f992c91036745ca0.CanInterface() {
					__obf_7348ad94a5462ee9 := Wrap(__obf_f992c91036745ca0.Interface()).Get(__obf_b90e4ca332607787[1:]...)
					if __obf_7348ad94a5462ee9.ValueType() != InvalidValue {
						__obf_db9539110ac34bf8[any.__obf_d59813f23ad740a8.Type().Field(__obf_28d099df85f083a8).Name] = __obf_7348ad94a5462ee9
					}
				}
			}
			return __obf_af999c25758b0ff1(__obf_db9539110ac34bf8)
		}
		return __obf_5a6677fd9236ae8d(__obf_b90e4ca332607787)
	default:
		return __obf_5a6677fd9236ae8d(__obf_b90e4ca332607787)
	}
}

func (any *__obf_d2e8e028a5c60fc6) Keys() []string {
	__obf_870ded78ba396f54 := make([]string, 0, any.__obf_d59813f23ad740a8.NumField())
	for __obf_28d099df85f083a8 := 0; __obf_28d099df85f083a8 < any.__obf_d59813f23ad740a8.NumField(); __obf_28d099df85f083a8++ {
		__obf_870ded78ba396f54 = append(__obf_870ded78ba396f54, any.__obf_d59813f23ad740a8.Type().Field(__obf_28d099df85f083a8).Name)
	}
	return __obf_870ded78ba396f54
}

func (any *__obf_d2e8e028a5c60fc6) Size() int {
	return any.__obf_d59813f23ad740a8.NumField()
}

func (any *__obf_d2e8e028a5c60fc6) WriteTo(__obf_2361f5214e84c60f *Stream) {
	__obf_2361f5214e84c60f.
		WriteVal(any.__obf_d59813f23ad740a8)
}

func (any *__obf_d2e8e028a5c60fc6) GetInterface() any {
	return any.__obf_d59813f23ad740a8.Interface()
}

type __obf_27302ea006b70fcd struct {
	__obf_86279f81acb381f3
	__obf_5966ecc5edb9b17e error
	__obf_d59813f23ad740a8 reflect.Value
}

func __obf_af999c25758b0ff1(__obf_d59813f23ad740a8 any) *__obf_27302ea006b70fcd {
	return &__obf_27302ea006b70fcd{__obf_86279f81acb381f3{}, nil, reflect.ValueOf(__obf_d59813f23ad740a8)}
}

func (any *__obf_27302ea006b70fcd) ValueType() ValueType {
	return ObjectValue
}

func (any *__obf_27302ea006b70fcd) MustBeValid() Any {
	return any
}

func (any *__obf_27302ea006b70fcd) Parse() *Iterator {
	return nil
}

func (any *__obf_27302ea006b70fcd) LastError() error {
	return any.__obf_5966ecc5edb9b17e
}

func (any *__obf_27302ea006b70fcd) ToBool() bool {
	return true
}

func (any *__obf_27302ea006b70fcd) ToInt() int {
	return 0
}

func (any *__obf_27302ea006b70fcd) ToInt32() int32 {
	return 0
}

func (any *__obf_27302ea006b70fcd) ToInt64() int64 {
	return 0
}

func (any *__obf_27302ea006b70fcd) ToUint() uint {
	return 0
}

func (any *__obf_27302ea006b70fcd) ToUint32() uint32 {
	return 0
}

func (any *__obf_27302ea006b70fcd) ToUint64() uint64 {
	return 0
}

func (any *__obf_27302ea006b70fcd) ToFloat32() float32 {
	return 0
}

func (any *__obf_27302ea006b70fcd) ToFloat64() float64 {
	return 0
}

func (any *__obf_27302ea006b70fcd) ToString() string {
	__obf_2d944450d48e5810, __obf_5966ecc5edb9b17e := MarshalToString(any.__obf_d59813f23ad740a8.Interface())
	any.__obf_5966ecc5edb9b17e = __obf_5966ecc5edb9b17e
	return __obf_2d944450d48e5810
}

func (any *__obf_27302ea006b70fcd) Get(__obf_b90e4ca332607787 ...any) Any {
	if len(__obf_b90e4ca332607787) == 0 {
		return any
	}
	switch __obf_0f79e2c35188222f := __obf_b90e4ca332607787[0].(type) {
	case int32:
		if '*' == __obf_0f79e2c35188222f {
			__obf_db9539110ac34bf8 := map[string]Any{}
			for _, __obf_09e2475e3bcc63c4 := range any.__obf_d59813f23ad740a8.MapKeys() {
				__obf_bd452debe72a9d3c := __obf_09e2475e3bcc63c4.String()
				__obf_bc320e3215fd14b8 := Wrap(any.__obf_d59813f23ad740a8.MapIndex(__obf_09e2475e3bcc63c4).Interface())
				__obf_7348ad94a5462ee9 := __obf_bc320e3215fd14b8.Get(__obf_b90e4ca332607787[1:]...)
				if __obf_7348ad94a5462ee9.ValueType() != InvalidValue {
					__obf_db9539110ac34bf8[__obf_bd452debe72a9d3c] = __obf_7348ad94a5462ee9
				}
			}
			return __obf_af999c25758b0ff1(__obf_db9539110ac34bf8)
		}
		return __obf_5a6677fd9236ae8d(__obf_b90e4ca332607787)
	default:
		__obf_1a534fe904baaf4a := any.__obf_d59813f23ad740a8.MapIndex(reflect.ValueOf(__obf_0f79e2c35188222f))
		if !__obf_1a534fe904baaf4a.IsValid() {
			return __obf_5a6677fd9236ae8d(__obf_b90e4ca332607787)
		}
		return Wrap(__obf_1a534fe904baaf4a.Interface())
	}
}

func (any *__obf_27302ea006b70fcd) Keys() []string {
	__obf_870ded78ba396f54 := make([]string, 0, any.__obf_d59813f23ad740a8.Len())
	for _, __obf_09e2475e3bcc63c4 := range any.__obf_d59813f23ad740a8.MapKeys() {
		__obf_870ded78ba396f54 = append(__obf_870ded78ba396f54, __obf_09e2475e3bcc63c4.String())
	}
	return __obf_870ded78ba396f54
}

func (any *__obf_27302ea006b70fcd) Size() int {
	return any.__obf_d59813f23ad740a8.Len()
}

func (any *__obf_27302ea006b70fcd) WriteTo(__obf_2361f5214e84c60f *Stream) {
	__obf_2361f5214e84c60f.
		WriteVal(any.__obf_d59813f23ad740a8)
}

func (any *__obf_27302ea006b70fcd) GetInterface() any {
	return any.__obf_d59813f23ad740a8.Interface()
}
