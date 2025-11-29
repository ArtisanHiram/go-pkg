package __obf_91620b895eeff9ed

import (
	"reflect"
	"unsafe"
)

type __obf_8759b49e282d9de6 struct {
	__obf_58563642f42f4a04
	__obf_892632c77e859037 *__obf_972c2293804d141c
	__obf_184433571fa55237 []byte
	__obf_62967739bca1d11d error
}

func (any *__obf_8759b49e282d9de6) ValueType() ValueType {
	return ArrayValue
}

func (any *__obf_8759b49e282d9de6) MustBeValid() Any {
	return any
}

func (any *__obf_8759b49e282d9de6) LastError() error {
	return any.__obf_62967739bca1d11d
}

func (any *__obf_8759b49e282d9de6) ToBool() bool {
	__obf_1bb30e8a74ed8233 := any.__obf_892632c77e859037.BorrowIterator(any.__obf_184433571fa55237)
	defer any.__obf_892632c77e859037.ReturnIterator(__obf_1bb30e8a74ed8233)
	return __obf_1bb30e8a74ed8233.ReadArray()
}

func (any *__obf_8759b49e282d9de6) ToInt() int {
	if any.ToBool() {
		return 1
	}
	return 0
}

func (any *__obf_8759b49e282d9de6) ToInt32() int32 {
	if any.ToBool() {
		return 1
	}
	return 0
}

func (any *__obf_8759b49e282d9de6) ToInt64() int64 {
	if any.ToBool() {
		return 1
	}
	return 0
}

func (any *__obf_8759b49e282d9de6) ToUint() uint {
	if any.ToBool() {
		return 1
	}
	return 0
}

func (any *__obf_8759b49e282d9de6) ToUint32() uint32 {
	if any.ToBool() {
		return 1
	}
	return 0
}

func (any *__obf_8759b49e282d9de6) ToUint64() uint64 {
	if any.ToBool() {
		return 1
	}
	return 0
}

func (any *__obf_8759b49e282d9de6) ToFloat32() float32 {
	if any.ToBool() {
		return 1
	}
	return 0
}

func (any *__obf_8759b49e282d9de6) ToFloat64() float64 {
	if any.ToBool() {
		return 1
	}
	return 0
}

func (any *__obf_8759b49e282d9de6) ToString() string {
	return *(*string)(unsafe.Pointer(&any.__obf_184433571fa55237))
}

func (any *__obf_8759b49e282d9de6) ToVal(__obf_bbfd2ac8618a6f0c any) {
	__obf_1bb30e8a74ed8233 := any.__obf_892632c77e859037.BorrowIterator(any.__obf_184433571fa55237)
	defer any.__obf_892632c77e859037.ReturnIterator(__obf_1bb30e8a74ed8233)
	__obf_1bb30e8a74ed8233.
		ReadVal(__obf_bbfd2ac8618a6f0c)
}

func (any *__obf_8759b49e282d9de6) Get(__obf_82c50fcbfc9ab432 ...any) Any {
	if len(__obf_82c50fcbfc9ab432) == 0 {
		return any
	}
	switch __obf_c75099dc7388aeb3 := __obf_82c50fcbfc9ab432[0].(type) {
	case int:
		__obf_1bb30e8a74ed8233 := any.__obf_892632c77e859037.BorrowIterator(any.__obf_184433571fa55237)
		defer any.__obf_892632c77e859037.ReturnIterator(__obf_1bb30e8a74ed8233)
		__obf_930565f5477d9c6d := __obf_3f3216ff423a5303(__obf_1bb30e8a74ed8233, __obf_c75099dc7388aeb3)
		if __obf_930565f5477d9c6d == nil {
			return __obf_227499e666173bc3(__obf_82c50fcbfc9ab432)
		}
		__obf_1bb30e8a74ed8233.
			ResetBytes(__obf_930565f5477d9c6d)
		return __obf_67ce74ddb71f8159(__obf_1bb30e8a74ed8233, __obf_82c50fcbfc9ab432[1:])
	case int32:
		if '*' == __obf_c75099dc7388aeb3 {
			__obf_1bb30e8a74ed8233 := any.__obf_892632c77e859037.BorrowIterator(any.__obf_184433571fa55237)
			defer any.__obf_892632c77e859037.ReturnIterator(__obf_1bb30e8a74ed8233)
			__obf_6fc61a171ba497a3 := make([]Any, 0)
			__obf_1bb30e8a74ed8233.
				ReadArrayCB(func(__obf_1bb30e8a74ed8233 *Iterator) bool {
					__obf_93f189bffabaf5a4 := __obf_1bb30e8a74ed8233.__obf_6ec0844a84418f37().Get(__obf_82c50fcbfc9ab432[1:]...)
					if __obf_93f189bffabaf5a4.ValueType() != InvalidValue {
						__obf_6fc61a171ba497a3 = append(__obf_6fc61a171ba497a3, __obf_93f189bffabaf5a4)
					}
					return true
				})
			return __obf_856f321acac25752(__obf_6fc61a171ba497a3)
		}
		return __obf_227499e666173bc3(__obf_82c50fcbfc9ab432)
	default:
		return __obf_227499e666173bc3(__obf_82c50fcbfc9ab432)
	}
}

func (any *__obf_8759b49e282d9de6) Size() int {
	__obf_dc51447b63477324 := 0
	__obf_1bb30e8a74ed8233 := any.__obf_892632c77e859037.BorrowIterator(any.__obf_184433571fa55237)
	defer any.__obf_892632c77e859037.ReturnIterator(__obf_1bb30e8a74ed8233)
	__obf_1bb30e8a74ed8233.
		ReadArrayCB(func(__obf_1bb30e8a74ed8233 *Iterator) bool {
			__obf_dc51447b63477324++
			__obf_1bb30e8a74ed8233.
				Skip()
			return true
		})
	return __obf_dc51447b63477324
}

func (any *__obf_8759b49e282d9de6) WriteTo(__obf_850a7457bb739a32 *Stream) {
	__obf_850a7457bb739a32.
		Write(any.__obf_184433571fa55237)
}

func (any *__obf_8759b49e282d9de6) GetInterface() any {
	__obf_1bb30e8a74ed8233 := any.__obf_892632c77e859037.BorrowIterator(any.__obf_184433571fa55237)
	defer any.__obf_892632c77e859037.ReturnIterator(__obf_1bb30e8a74ed8233)
	return __obf_1bb30e8a74ed8233.Read()
}

type __obf_b38756b00a7b6b82 struct {
	__obf_58563642f42f4a04
	__obf_bbfd2ac8618a6f0c reflect.Value
}

func __obf_856f321acac25752(__obf_bbfd2ac8618a6f0c any) *__obf_b38756b00a7b6b82 {
	return &__obf_b38756b00a7b6b82{__obf_58563642f42f4a04{}, reflect.ValueOf(__obf_bbfd2ac8618a6f0c)}
}

func (any *__obf_b38756b00a7b6b82) ValueType() ValueType {
	return ArrayValue
}

func (any *__obf_b38756b00a7b6b82) MustBeValid() Any {
	return any
}

func (any *__obf_b38756b00a7b6b82) LastError() error {
	return nil
}

func (any *__obf_b38756b00a7b6b82) ToBool() bool {
	return any.__obf_bbfd2ac8618a6f0c.Len() != 0
}

func (any *__obf_b38756b00a7b6b82) ToInt() int {
	if any.__obf_bbfd2ac8618a6f0c.Len() == 0 {
		return 0
	}
	return 1
}

func (any *__obf_b38756b00a7b6b82) ToInt32() int32 {
	if any.__obf_bbfd2ac8618a6f0c.Len() == 0 {
		return 0
	}
	return 1
}

func (any *__obf_b38756b00a7b6b82) ToInt64() int64 {
	if any.__obf_bbfd2ac8618a6f0c.Len() == 0 {
		return 0
	}
	return 1
}

func (any *__obf_b38756b00a7b6b82) ToUint() uint {
	if any.__obf_bbfd2ac8618a6f0c.Len() == 0 {
		return 0
	}
	return 1
}

func (any *__obf_b38756b00a7b6b82) ToUint32() uint32 {
	if any.__obf_bbfd2ac8618a6f0c.Len() == 0 {
		return 0
	}
	return 1
}

func (any *__obf_b38756b00a7b6b82) ToUint64() uint64 {
	if any.__obf_bbfd2ac8618a6f0c.Len() == 0 {
		return 0
	}
	return 1
}

func (any *__obf_b38756b00a7b6b82) ToFloat32() float32 {
	if any.__obf_bbfd2ac8618a6f0c.Len() == 0 {
		return 0
	}
	return 1
}

func (any *__obf_b38756b00a7b6b82) ToFloat64() float64 {
	if any.__obf_bbfd2ac8618a6f0c.Len() == 0 {
		return 0
	}
	return 1
}

func (any *__obf_b38756b00a7b6b82) ToString() string {
	__obf_e91bd2feb751e4f1, _ := MarshalToString(any.__obf_bbfd2ac8618a6f0c.Interface())
	return __obf_e91bd2feb751e4f1
}

func (any *__obf_b38756b00a7b6b82) Get(__obf_82c50fcbfc9ab432 ...any) Any {
	if len(__obf_82c50fcbfc9ab432) == 0 {
		return any
	}
	switch __obf_c75099dc7388aeb3 := __obf_82c50fcbfc9ab432[0].(type) {
	case int:
		if __obf_c75099dc7388aeb3 < 0 || __obf_c75099dc7388aeb3 >= any.__obf_bbfd2ac8618a6f0c.Len() {
			return __obf_227499e666173bc3(__obf_82c50fcbfc9ab432)
		}
		return Wrap(any.__obf_bbfd2ac8618a6f0c.Index(__obf_c75099dc7388aeb3).Interface())
	case int32:
		if '*' == __obf_c75099dc7388aeb3 {
			__obf_8c73a46661d509f0 := make([]Any, 0)
			for __obf_5aa5c8829b97f182 := 0; __obf_5aa5c8829b97f182 < any.__obf_bbfd2ac8618a6f0c.Len(); __obf_5aa5c8829b97f182++ {
				__obf_d2a067019e17e3c8 := Wrap(any.__obf_bbfd2ac8618a6f0c.Index(__obf_5aa5c8829b97f182).Interface()).Get(__obf_82c50fcbfc9ab432[1:]...)
				if __obf_d2a067019e17e3c8.ValueType() != InvalidValue {
					__obf_8c73a46661d509f0 = append(__obf_8c73a46661d509f0, __obf_d2a067019e17e3c8)
				}
			}
			return __obf_856f321acac25752(__obf_8c73a46661d509f0)
		}
		return __obf_227499e666173bc3(__obf_82c50fcbfc9ab432)
	default:
		return __obf_227499e666173bc3(__obf_82c50fcbfc9ab432)
	}
}

func (any *__obf_b38756b00a7b6b82) Size() int {
	return any.__obf_bbfd2ac8618a6f0c.Len()
}

func (any *__obf_b38756b00a7b6b82) WriteTo(__obf_850a7457bb739a32 *Stream) {
	__obf_850a7457bb739a32.
		WriteVal(any.__obf_bbfd2ac8618a6f0c)
}

func (any *__obf_b38756b00a7b6b82) GetInterface() any {
	return any.__obf_bbfd2ac8618a6f0c.Interface()
}
