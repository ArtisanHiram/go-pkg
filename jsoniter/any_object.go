package __obf_030d39f7a8de96c6

import (
	"reflect"
	"unsafe"
)

type __obf_93b2c1e9efd325de struct {
	__obf_7834a13a259af712
	__obf_679611dc56ff465b *__obf_81639538813814ff
	__obf_0b1656d7ef4bd9df []byte
	__obf_fcc907dd69879566 error
}

func (any *__obf_93b2c1e9efd325de) ValueType() ValueType {
	return ObjectValue
}

func (any *__obf_93b2c1e9efd325de) MustBeValid() Any {
	return any
}

func (any *__obf_93b2c1e9efd325de) LastError() error {
	return any.__obf_fcc907dd69879566
}

func (any *__obf_93b2c1e9efd325de) ToBool() bool {
	return true
}

func (any *__obf_93b2c1e9efd325de) ToInt() int {
	return 0
}

func (any *__obf_93b2c1e9efd325de) ToInt32() int32 {
	return 0
}

func (any *__obf_93b2c1e9efd325de) ToInt64() int64 {
	return 0
}

func (any *__obf_93b2c1e9efd325de) ToUint() uint {
	return 0
}

func (any *__obf_93b2c1e9efd325de) ToUint32() uint32 {
	return 0
}

func (any *__obf_93b2c1e9efd325de) ToUint64() uint64 {
	return 0
}

func (any *__obf_93b2c1e9efd325de) ToFloat32() float32 {
	return 0
}

func (any *__obf_93b2c1e9efd325de) ToFloat64() float64 {
	return 0
}

func (any *__obf_93b2c1e9efd325de) ToString() string {
	return *(*string)(unsafe.Pointer(&any.__obf_0b1656d7ef4bd9df))
}

func (any *__obf_93b2c1e9efd325de) ToVal(__obf_eefa92392cc2442c any) {
	__obf_4ab56a99965952e7 := any.__obf_679611dc56ff465b.BorrowIterator(any.__obf_0b1656d7ef4bd9df)
	defer any.__obf_679611dc56ff465b.ReturnIterator(__obf_4ab56a99965952e7)
	__obf_4ab56a99965952e7.
		ReadVal(__obf_eefa92392cc2442c)
}

func (any *__obf_93b2c1e9efd325de) Get(__obf_f77a9507782b919d ...any) Any {
	if len(__obf_f77a9507782b919d) == 0 {
		return any
	}
	switch __obf_5f55b20fdc05864d := __obf_f77a9507782b919d[0].(type) {
	case string:
		__obf_4ab56a99965952e7 := any.__obf_679611dc56ff465b.BorrowIterator(any.__obf_0b1656d7ef4bd9df)
		defer any.__obf_679611dc56ff465b.ReturnIterator(__obf_4ab56a99965952e7)
		__obf_d2de1fa55acc5ff0 := __obf_d17dcb541db3da34(__obf_4ab56a99965952e7, __obf_5f55b20fdc05864d)
		if __obf_d2de1fa55acc5ff0 == nil {
			return __obf_0adb97f96f4a14bd(__obf_f77a9507782b919d)
		}
		__obf_4ab56a99965952e7.
			ResetBytes(__obf_d2de1fa55acc5ff0)
		return __obf_4cd15630eeb40546(__obf_4ab56a99965952e7, __obf_f77a9507782b919d[1:])
	case int32:
		if '*' == __obf_5f55b20fdc05864d {
			__obf_849b132b302d3a9c := map[string]Any{}
			__obf_4ab56a99965952e7 := any.__obf_679611dc56ff465b.BorrowIterator(any.__obf_0b1656d7ef4bd9df)
			defer any.__obf_679611dc56ff465b.ReturnIterator(__obf_4ab56a99965952e7)
			__obf_4ab56a99965952e7.
				ReadMapCB(func(__obf_4ab56a99965952e7 *Iterator, __obf_cd4d02f264b18d55 string) bool {
					__obf_499e17e4254c75a7 := __obf_4cd15630eeb40546(__obf_4ab56a99965952e7, __obf_f77a9507782b919d[1:])
					if __obf_499e17e4254c75a7.ValueType() != InvalidValue {
						__obf_849b132b302d3a9c[__obf_cd4d02f264b18d55] = __obf_499e17e4254c75a7
					}
					return true
				})
			return __obf_1dc61d22b3defb62(__obf_849b132b302d3a9c)
		}
		return __obf_0adb97f96f4a14bd(__obf_f77a9507782b919d)
	default:
		return __obf_0adb97f96f4a14bd(__obf_f77a9507782b919d)
	}
}

func (any *__obf_93b2c1e9efd325de) Keys() []string {
	__obf_ba60b41c7ef0aae8 := []string{}
	__obf_4ab56a99965952e7 := any.__obf_679611dc56ff465b.BorrowIterator(any.__obf_0b1656d7ef4bd9df)
	defer any.__obf_679611dc56ff465b.ReturnIterator(__obf_4ab56a99965952e7)
	__obf_4ab56a99965952e7.
		ReadMapCB(func(__obf_4ab56a99965952e7 *Iterator, __obf_cd4d02f264b18d55 string) bool {
			__obf_4ab56a99965952e7.
				Skip()
			__obf_ba60b41c7ef0aae8 = append(__obf_ba60b41c7ef0aae8, __obf_cd4d02f264b18d55)
			return true
		})
	return __obf_ba60b41c7ef0aae8
}

func (any *__obf_93b2c1e9efd325de) Size() int {
	__obf_452ca3698eb735a5 := 0
	__obf_4ab56a99965952e7 := any.__obf_679611dc56ff465b.BorrowIterator(any.__obf_0b1656d7ef4bd9df)
	defer any.__obf_679611dc56ff465b.ReturnIterator(__obf_4ab56a99965952e7)
	__obf_4ab56a99965952e7.
		ReadObjectCB(func(__obf_4ab56a99965952e7 *Iterator, __obf_cd4d02f264b18d55 string) bool {
			__obf_4ab56a99965952e7.
				Skip()
			__obf_452ca3698eb735a5++
			return true
		})
	return __obf_452ca3698eb735a5
}

func (any *__obf_93b2c1e9efd325de) WriteTo(__obf_8a2c51fe22775655 *Stream) {
	__obf_8a2c51fe22775655.
		Write(any.__obf_0b1656d7ef4bd9df)
}

func (any *__obf_93b2c1e9efd325de) GetInterface() any {
	__obf_4ab56a99965952e7 := any.__obf_679611dc56ff465b.BorrowIterator(any.__obf_0b1656d7ef4bd9df)
	defer any.__obf_679611dc56ff465b.ReturnIterator(__obf_4ab56a99965952e7)
	return __obf_4ab56a99965952e7.Read()
}

type __obf_698fe270afe3739e struct {
	__obf_7834a13a259af712
	__obf_fcc907dd69879566 error
	__obf_efaf2719b44ad8ac reflect.Value
}

func __obf_0cb32cddab75abc0(__obf_efaf2719b44ad8ac any) *__obf_698fe270afe3739e {
	return &__obf_698fe270afe3739e{__obf_7834a13a259af712{}, nil, reflect.ValueOf(__obf_efaf2719b44ad8ac)}
}

func (any *__obf_698fe270afe3739e) ValueType() ValueType {
	return ObjectValue
}

func (any *__obf_698fe270afe3739e) MustBeValid() Any {
	return any
}

func (any *__obf_698fe270afe3739e) Parse() *Iterator {
	return nil
}

func (any *__obf_698fe270afe3739e) LastError() error {
	return any.__obf_fcc907dd69879566
}

func (any *__obf_698fe270afe3739e) ToBool() bool {
	return any.__obf_efaf2719b44ad8ac.NumField() != 0
}

func (any *__obf_698fe270afe3739e) ToInt() int {
	return 0
}

func (any *__obf_698fe270afe3739e) ToInt32() int32 {
	return 0
}

func (any *__obf_698fe270afe3739e) ToInt64() int64 {
	return 0
}

func (any *__obf_698fe270afe3739e) ToUint() uint {
	return 0
}

func (any *__obf_698fe270afe3739e) ToUint32() uint32 {
	return 0
}

func (any *__obf_698fe270afe3739e) ToUint64() uint64 {
	return 0
}

func (any *__obf_698fe270afe3739e) ToFloat32() float32 {
	return 0
}

func (any *__obf_698fe270afe3739e) ToFloat64() float64 {
	return 0
}

func (any *__obf_698fe270afe3739e) ToString() string {
	__obf_428c3b4a95725c4a, __obf_fcc907dd69879566 := MarshalToString(any.__obf_efaf2719b44ad8ac.Interface())
	any.__obf_fcc907dd69879566 = __obf_fcc907dd69879566
	return __obf_428c3b4a95725c4a
}

func (any *__obf_698fe270afe3739e) Get(__obf_f77a9507782b919d ...any) Any {
	if len(__obf_f77a9507782b919d) == 0 {
		return any
	}
	switch __obf_5f55b20fdc05864d := __obf_f77a9507782b919d[0].(type) {
	case string:
		__obf_cd4d02f264b18d55 := any.__obf_efaf2719b44ad8ac.FieldByName(__obf_5f55b20fdc05864d)
		if !__obf_cd4d02f264b18d55.IsValid() {
			return __obf_0adb97f96f4a14bd(__obf_f77a9507782b919d)
		}
		return Wrap(__obf_cd4d02f264b18d55.Interface())
	case int32:
		if '*' == __obf_5f55b20fdc05864d {
			__obf_849b132b302d3a9c := map[string]Any{}
			for __obf_82c6e05b9d226c58 := 0; __obf_82c6e05b9d226c58 < any.__obf_efaf2719b44ad8ac.NumField(); __obf_82c6e05b9d226c58++ {
				__obf_cd4d02f264b18d55 := any.__obf_efaf2719b44ad8ac.Field(__obf_82c6e05b9d226c58)
				if __obf_cd4d02f264b18d55.CanInterface() {
					__obf_499e17e4254c75a7 := Wrap(__obf_cd4d02f264b18d55.Interface()).Get(__obf_f77a9507782b919d[1:]...)
					if __obf_499e17e4254c75a7.ValueType() != InvalidValue {
						__obf_849b132b302d3a9c[any.__obf_efaf2719b44ad8ac.Type().Field(__obf_82c6e05b9d226c58).Name] = __obf_499e17e4254c75a7
					}
				}
			}
			return __obf_1dc61d22b3defb62(__obf_849b132b302d3a9c)
		}
		return __obf_0adb97f96f4a14bd(__obf_f77a9507782b919d)
	default:
		return __obf_0adb97f96f4a14bd(__obf_f77a9507782b919d)
	}
}

func (any *__obf_698fe270afe3739e) Keys() []string {
	__obf_ba60b41c7ef0aae8 := make([]string, 0, any.__obf_efaf2719b44ad8ac.NumField())
	for __obf_82c6e05b9d226c58 := 0; __obf_82c6e05b9d226c58 < any.__obf_efaf2719b44ad8ac.NumField(); __obf_82c6e05b9d226c58++ {
		__obf_ba60b41c7ef0aae8 = append(__obf_ba60b41c7ef0aae8, any.__obf_efaf2719b44ad8ac.Type().Field(__obf_82c6e05b9d226c58).Name)
	}
	return __obf_ba60b41c7ef0aae8
}

func (any *__obf_698fe270afe3739e) Size() int {
	return any.__obf_efaf2719b44ad8ac.NumField()
}

func (any *__obf_698fe270afe3739e) WriteTo(__obf_8a2c51fe22775655 *Stream) {
	__obf_8a2c51fe22775655.
		WriteVal(any.__obf_efaf2719b44ad8ac)
}

func (any *__obf_698fe270afe3739e) GetInterface() any {
	return any.__obf_efaf2719b44ad8ac.Interface()
}

type __obf_d52c3ff6008a9c10 struct {
	__obf_7834a13a259af712
	__obf_fcc907dd69879566 error
	__obf_efaf2719b44ad8ac reflect.Value
}

func __obf_1dc61d22b3defb62(__obf_efaf2719b44ad8ac any) *__obf_d52c3ff6008a9c10 {
	return &__obf_d52c3ff6008a9c10{__obf_7834a13a259af712{}, nil, reflect.ValueOf(__obf_efaf2719b44ad8ac)}
}

func (any *__obf_d52c3ff6008a9c10) ValueType() ValueType {
	return ObjectValue
}

func (any *__obf_d52c3ff6008a9c10) MustBeValid() Any {
	return any
}

func (any *__obf_d52c3ff6008a9c10) Parse() *Iterator {
	return nil
}

func (any *__obf_d52c3ff6008a9c10) LastError() error {
	return any.__obf_fcc907dd69879566
}

func (any *__obf_d52c3ff6008a9c10) ToBool() bool {
	return true
}

func (any *__obf_d52c3ff6008a9c10) ToInt() int {
	return 0
}

func (any *__obf_d52c3ff6008a9c10) ToInt32() int32 {
	return 0
}

func (any *__obf_d52c3ff6008a9c10) ToInt64() int64 {
	return 0
}

func (any *__obf_d52c3ff6008a9c10) ToUint() uint {
	return 0
}

func (any *__obf_d52c3ff6008a9c10) ToUint32() uint32 {
	return 0
}

func (any *__obf_d52c3ff6008a9c10) ToUint64() uint64 {
	return 0
}

func (any *__obf_d52c3ff6008a9c10) ToFloat32() float32 {
	return 0
}

func (any *__obf_d52c3ff6008a9c10) ToFloat64() float64 {
	return 0
}

func (any *__obf_d52c3ff6008a9c10) ToString() string {
	__obf_428c3b4a95725c4a, __obf_fcc907dd69879566 := MarshalToString(any.__obf_efaf2719b44ad8ac.Interface())
	any.__obf_fcc907dd69879566 = __obf_fcc907dd69879566
	return __obf_428c3b4a95725c4a
}

func (any *__obf_d52c3ff6008a9c10) Get(__obf_f77a9507782b919d ...any) Any {
	if len(__obf_f77a9507782b919d) == 0 {
		return any
	}
	switch __obf_5f55b20fdc05864d := __obf_f77a9507782b919d[0].(type) {
	case int32:
		if '*' == __obf_5f55b20fdc05864d {
			__obf_849b132b302d3a9c := map[string]Any{}
			for _, __obf_53d763816351c178 := range any.__obf_efaf2719b44ad8ac.MapKeys() {
				__obf_cfe04ed2246e45b7 := __obf_53d763816351c178.String()
				__obf_2cb4d0263b55edb1 := Wrap(any.__obf_efaf2719b44ad8ac.MapIndex(__obf_53d763816351c178).Interface())
				__obf_499e17e4254c75a7 := __obf_2cb4d0263b55edb1.Get(__obf_f77a9507782b919d[1:]...)
				if __obf_499e17e4254c75a7.ValueType() != InvalidValue {
					__obf_849b132b302d3a9c[__obf_cfe04ed2246e45b7] = __obf_499e17e4254c75a7
				}
			}
			return __obf_1dc61d22b3defb62(__obf_849b132b302d3a9c)
		}
		return __obf_0adb97f96f4a14bd(__obf_f77a9507782b919d)
	default:
		__obf_b80d5217e1232943 := any.__obf_efaf2719b44ad8ac.MapIndex(reflect.ValueOf(__obf_5f55b20fdc05864d))
		if !__obf_b80d5217e1232943.IsValid() {
			return __obf_0adb97f96f4a14bd(__obf_f77a9507782b919d)
		}
		return Wrap(__obf_b80d5217e1232943.Interface())
	}
}

func (any *__obf_d52c3ff6008a9c10) Keys() []string {
	__obf_ba60b41c7ef0aae8 := make([]string, 0, any.__obf_efaf2719b44ad8ac.Len())
	for _, __obf_53d763816351c178 := range any.__obf_efaf2719b44ad8ac.MapKeys() {
		__obf_ba60b41c7ef0aae8 = append(__obf_ba60b41c7ef0aae8, __obf_53d763816351c178.String())
	}
	return __obf_ba60b41c7ef0aae8
}

func (any *__obf_d52c3ff6008a9c10) Size() int {
	return any.__obf_efaf2719b44ad8ac.Len()
}

func (any *__obf_d52c3ff6008a9c10) WriteTo(__obf_8a2c51fe22775655 *Stream) {
	__obf_8a2c51fe22775655.
		WriteVal(any.__obf_efaf2719b44ad8ac)
}

func (any *__obf_d52c3ff6008a9c10) GetInterface() any {
	return any.__obf_efaf2719b44ad8ac.Interface()
}
