package __obf_030d39f7a8de96c6

import (
	"reflect"
	"unsafe"
)

type __obf_6861b65fa3b114d9 struct {
	__obf_7834a13a259af712
	__obf_679611dc56ff465b *__obf_81639538813814ff
	__obf_0b1656d7ef4bd9df []byte
	__obf_fcc907dd69879566 error
}

func (any *__obf_6861b65fa3b114d9) ValueType() ValueType {
	return ArrayValue
}

func (any *__obf_6861b65fa3b114d9) MustBeValid() Any {
	return any
}

func (any *__obf_6861b65fa3b114d9) LastError() error {
	return any.__obf_fcc907dd69879566
}

func (any *__obf_6861b65fa3b114d9) ToBool() bool {
	__obf_4ab56a99965952e7 := any.__obf_679611dc56ff465b.BorrowIterator(any.__obf_0b1656d7ef4bd9df)
	defer any.__obf_679611dc56ff465b.ReturnIterator(__obf_4ab56a99965952e7)
	return __obf_4ab56a99965952e7.ReadArray()
}

func (any *__obf_6861b65fa3b114d9) ToInt() int {
	if any.ToBool() {
		return 1
	}
	return 0
}

func (any *__obf_6861b65fa3b114d9) ToInt32() int32 {
	if any.ToBool() {
		return 1
	}
	return 0
}

func (any *__obf_6861b65fa3b114d9) ToInt64() int64 {
	if any.ToBool() {
		return 1
	}
	return 0
}

func (any *__obf_6861b65fa3b114d9) ToUint() uint {
	if any.ToBool() {
		return 1
	}
	return 0
}

func (any *__obf_6861b65fa3b114d9) ToUint32() uint32 {
	if any.ToBool() {
		return 1
	}
	return 0
}

func (any *__obf_6861b65fa3b114d9) ToUint64() uint64 {
	if any.ToBool() {
		return 1
	}
	return 0
}

func (any *__obf_6861b65fa3b114d9) ToFloat32() float32 {
	if any.ToBool() {
		return 1
	}
	return 0
}

func (any *__obf_6861b65fa3b114d9) ToFloat64() float64 {
	if any.ToBool() {
		return 1
	}
	return 0
}

func (any *__obf_6861b65fa3b114d9) ToString() string {
	return *(*string)(unsafe.Pointer(&any.__obf_0b1656d7ef4bd9df))
}

func (any *__obf_6861b65fa3b114d9) ToVal(__obf_efaf2719b44ad8ac any) {
	__obf_4ab56a99965952e7 := any.__obf_679611dc56ff465b.BorrowIterator(any.__obf_0b1656d7ef4bd9df)
	defer any.__obf_679611dc56ff465b.ReturnIterator(__obf_4ab56a99965952e7)
	__obf_4ab56a99965952e7.
		ReadVal(__obf_efaf2719b44ad8ac)
}

func (any *__obf_6861b65fa3b114d9) Get(__obf_f77a9507782b919d ...any) Any {
	if len(__obf_f77a9507782b919d) == 0 {
		return any
	}
	switch __obf_5f55b20fdc05864d := __obf_f77a9507782b919d[0].(type) {
	case int:
		__obf_4ab56a99965952e7 := any.__obf_679611dc56ff465b.BorrowIterator(any.__obf_0b1656d7ef4bd9df)
		defer any.__obf_679611dc56ff465b.ReturnIterator(__obf_4ab56a99965952e7)
		__obf_d2de1fa55acc5ff0 := __obf_6f7138d64db3c315(__obf_4ab56a99965952e7, __obf_5f55b20fdc05864d)
		if __obf_d2de1fa55acc5ff0 == nil {
			return __obf_0adb97f96f4a14bd(__obf_f77a9507782b919d)
		}
		__obf_4ab56a99965952e7.
			ResetBytes(__obf_d2de1fa55acc5ff0)
		return __obf_4cd15630eeb40546(__obf_4ab56a99965952e7, __obf_f77a9507782b919d[1:])
	case int32:
		if '*' == __obf_5f55b20fdc05864d {
			__obf_4ab56a99965952e7 := any.__obf_679611dc56ff465b.BorrowIterator(any.__obf_0b1656d7ef4bd9df)
			defer any.__obf_679611dc56ff465b.ReturnIterator(__obf_4ab56a99965952e7)
			__obf_612acb1933b33fd2 := make([]Any, 0)
			__obf_4ab56a99965952e7.
				ReadArrayCB(func(__obf_4ab56a99965952e7 *Iterator) bool {
					__obf_ccdc4ee8b9d2f52f := __obf_4ab56a99965952e7.__obf_d0ec77416ba1caaa().Get(__obf_f77a9507782b919d[1:]...)
					if __obf_ccdc4ee8b9d2f52f.ValueType() != InvalidValue {
						__obf_612acb1933b33fd2 = append(__obf_612acb1933b33fd2, __obf_ccdc4ee8b9d2f52f)
					}
					return true
				})
			return __obf_eb0024afd144815a(__obf_612acb1933b33fd2)
		}
		return __obf_0adb97f96f4a14bd(__obf_f77a9507782b919d)
	default:
		return __obf_0adb97f96f4a14bd(__obf_f77a9507782b919d)
	}
}

func (any *__obf_6861b65fa3b114d9) Size() int {
	__obf_452ca3698eb735a5 := 0
	__obf_4ab56a99965952e7 := any.__obf_679611dc56ff465b.BorrowIterator(any.__obf_0b1656d7ef4bd9df)
	defer any.__obf_679611dc56ff465b.ReturnIterator(__obf_4ab56a99965952e7)
	__obf_4ab56a99965952e7.
		ReadArrayCB(func(__obf_4ab56a99965952e7 *Iterator) bool {
			__obf_452ca3698eb735a5++
			__obf_4ab56a99965952e7.
				Skip()
			return true
		})
	return __obf_452ca3698eb735a5
}

func (any *__obf_6861b65fa3b114d9) WriteTo(__obf_8a2c51fe22775655 *Stream) {
	__obf_8a2c51fe22775655.
		Write(any.__obf_0b1656d7ef4bd9df)
}

func (any *__obf_6861b65fa3b114d9) GetInterface() any {
	__obf_4ab56a99965952e7 := any.__obf_679611dc56ff465b.BorrowIterator(any.__obf_0b1656d7ef4bd9df)
	defer any.__obf_679611dc56ff465b.ReturnIterator(__obf_4ab56a99965952e7)
	return __obf_4ab56a99965952e7.Read()
}

type __obf_da1b433c05d622d3 struct {
	__obf_7834a13a259af712
	__obf_efaf2719b44ad8ac reflect.Value
}

func __obf_eb0024afd144815a(__obf_efaf2719b44ad8ac any) *__obf_da1b433c05d622d3 {
	return &__obf_da1b433c05d622d3{__obf_7834a13a259af712{}, reflect.ValueOf(__obf_efaf2719b44ad8ac)}
}

func (any *__obf_da1b433c05d622d3) ValueType() ValueType {
	return ArrayValue
}

func (any *__obf_da1b433c05d622d3) MustBeValid() Any {
	return any
}

func (any *__obf_da1b433c05d622d3) LastError() error {
	return nil
}

func (any *__obf_da1b433c05d622d3) ToBool() bool {
	return any.__obf_efaf2719b44ad8ac.Len() != 0
}

func (any *__obf_da1b433c05d622d3) ToInt() int {
	if any.__obf_efaf2719b44ad8ac.Len() == 0 {
		return 0
	}
	return 1
}

func (any *__obf_da1b433c05d622d3) ToInt32() int32 {
	if any.__obf_efaf2719b44ad8ac.Len() == 0 {
		return 0
	}
	return 1
}

func (any *__obf_da1b433c05d622d3) ToInt64() int64 {
	if any.__obf_efaf2719b44ad8ac.Len() == 0 {
		return 0
	}
	return 1
}

func (any *__obf_da1b433c05d622d3) ToUint() uint {
	if any.__obf_efaf2719b44ad8ac.Len() == 0 {
		return 0
	}
	return 1
}

func (any *__obf_da1b433c05d622d3) ToUint32() uint32 {
	if any.__obf_efaf2719b44ad8ac.Len() == 0 {
		return 0
	}
	return 1
}

func (any *__obf_da1b433c05d622d3) ToUint64() uint64 {
	if any.__obf_efaf2719b44ad8ac.Len() == 0 {
		return 0
	}
	return 1
}

func (any *__obf_da1b433c05d622d3) ToFloat32() float32 {
	if any.__obf_efaf2719b44ad8ac.Len() == 0 {
		return 0
	}
	return 1
}

func (any *__obf_da1b433c05d622d3) ToFloat64() float64 {
	if any.__obf_efaf2719b44ad8ac.Len() == 0 {
		return 0
	}
	return 1
}

func (any *__obf_da1b433c05d622d3) ToString() string {
	__obf_428c3b4a95725c4a, _ := MarshalToString(any.__obf_efaf2719b44ad8ac.Interface())
	return __obf_428c3b4a95725c4a
}

func (any *__obf_da1b433c05d622d3) Get(__obf_f77a9507782b919d ...any) Any {
	if len(__obf_f77a9507782b919d) == 0 {
		return any
	}
	switch __obf_5f55b20fdc05864d := __obf_f77a9507782b919d[0].(type) {
	case int:
		if __obf_5f55b20fdc05864d < 0 || __obf_5f55b20fdc05864d >= any.__obf_efaf2719b44ad8ac.Len() {
			return __obf_0adb97f96f4a14bd(__obf_f77a9507782b919d)
		}
		return Wrap(any.__obf_efaf2719b44ad8ac.Index(__obf_5f55b20fdc05864d).Interface())
	case int32:
		if '*' == __obf_5f55b20fdc05864d {
			__obf_849b132b302d3a9c := make([]Any, 0)
			for __obf_82c6e05b9d226c58 := 0; __obf_82c6e05b9d226c58 < any.__obf_efaf2719b44ad8ac.Len(); __obf_82c6e05b9d226c58++ {
				__obf_499e17e4254c75a7 := Wrap(any.__obf_efaf2719b44ad8ac.Index(__obf_82c6e05b9d226c58).Interface()).Get(__obf_f77a9507782b919d[1:]...)
				if __obf_499e17e4254c75a7.ValueType() != InvalidValue {
					__obf_849b132b302d3a9c = append(__obf_849b132b302d3a9c, __obf_499e17e4254c75a7)
				}
			}
			return __obf_eb0024afd144815a(__obf_849b132b302d3a9c)
		}
		return __obf_0adb97f96f4a14bd(__obf_f77a9507782b919d)
	default:
		return __obf_0adb97f96f4a14bd(__obf_f77a9507782b919d)
	}
}

func (any *__obf_da1b433c05d622d3) Size() int {
	return any.__obf_efaf2719b44ad8ac.Len()
}

func (any *__obf_da1b433c05d622d3) WriteTo(__obf_8a2c51fe22775655 *Stream) {
	__obf_8a2c51fe22775655.
		WriteVal(any.__obf_efaf2719b44ad8ac)
}

func (any *__obf_da1b433c05d622d3) GetInterface() any {
	return any.__obf_efaf2719b44ad8ac.Interface()
}
