package __obf_91620b895eeff9ed

import (
	"reflect"
	"unsafe"
)

type __obf_9460519df17f70ab struct {
	__obf_58563642f42f4a04
	__obf_892632c77e859037 *__obf_972c2293804d141c
	__obf_184433571fa55237 []byte
	__obf_62967739bca1d11d error
}

func (any *__obf_9460519df17f70ab) ValueType() ValueType {
	return ObjectValue
}

func (any *__obf_9460519df17f70ab) MustBeValid() Any {
	return any
}

func (any *__obf_9460519df17f70ab) LastError() error {
	return any.__obf_62967739bca1d11d
}

func (any *__obf_9460519df17f70ab) ToBool() bool {
	return true
}

func (any *__obf_9460519df17f70ab) ToInt() int {
	return 0
}

func (any *__obf_9460519df17f70ab) ToInt32() int32 {
	return 0
}

func (any *__obf_9460519df17f70ab) ToInt64() int64 {
	return 0
}

func (any *__obf_9460519df17f70ab) ToUint() uint {
	return 0
}

func (any *__obf_9460519df17f70ab) ToUint32() uint32 {
	return 0
}

func (any *__obf_9460519df17f70ab) ToUint64() uint64 {
	return 0
}

func (any *__obf_9460519df17f70ab) ToFloat32() float32 {
	return 0
}

func (any *__obf_9460519df17f70ab) ToFloat64() float64 {
	return 0
}

func (any *__obf_9460519df17f70ab) ToString() string {
	return *(*string)(unsafe.Pointer(&any.__obf_184433571fa55237))
}

func (any *__obf_9460519df17f70ab) ToVal(__obf_35136ef2ac0f94e4 any) {
	__obf_1bb30e8a74ed8233 := any.__obf_892632c77e859037.BorrowIterator(any.__obf_184433571fa55237)
	defer any.__obf_892632c77e859037.ReturnIterator(__obf_1bb30e8a74ed8233)
	__obf_1bb30e8a74ed8233.
		ReadVal(__obf_35136ef2ac0f94e4)
}

func (any *__obf_9460519df17f70ab) Get(__obf_82c50fcbfc9ab432 ...any) Any {
	if len(__obf_82c50fcbfc9ab432) == 0 {
		return any
	}
	switch __obf_c75099dc7388aeb3 := __obf_82c50fcbfc9ab432[0].(type) {
	case string:
		__obf_1bb30e8a74ed8233 := any.__obf_892632c77e859037.BorrowIterator(any.__obf_184433571fa55237)
		defer any.__obf_892632c77e859037.ReturnIterator(__obf_1bb30e8a74ed8233)
		__obf_930565f5477d9c6d := __obf_820a957cfe88153b(__obf_1bb30e8a74ed8233, __obf_c75099dc7388aeb3)
		if __obf_930565f5477d9c6d == nil {
			return __obf_227499e666173bc3(__obf_82c50fcbfc9ab432)
		}
		__obf_1bb30e8a74ed8233.
			ResetBytes(__obf_930565f5477d9c6d)
		return __obf_67ce74ddb71f8159(__obf_1bb30e8a74ed8233, __obf_82c50fcbfc9ab432[1:])
	case int32:
		if '*' == __obf_c75099dc7388aeb3 {
			__obf_8c73a46661d509f0 := map[string]Any{}
			__obf_1bb30e8a74ed8233 := any.__obf_892632c77e859037.BorrowIterator(any.__obf_184433571fa55237)
			defer any.__obf_892632c77e859037.ReturnIterator(__obf_1bb30e8a74ed8233)
			__obf_1bb30e8a74ed8233.
				ReadMapCB(func(__obf_1bb30e8a74ed8233 *Iterator, __obf_7e01b5b4651074d4 string) bool {
					__obf_d2a067019e17e3c8 := __obf_67ce74ddb71f8159(__obf_1bb30e8a74ed8233, __obf_82c50fcbfc9ab432[1:])
					if __obf_d2a067019e17e3c8.ValueType() != InvalidValue {
						__obf_8c73a46661d509f0[__obf_7e01b5b4651074d4] = __obf_d2a067019e17e3c8
					}
					return true
				})
			return __obf_200d6b32a72d70af(__obf_8c73a46661d509f0)
		}
		return __obf_227499e666173bc3(__obf_82c50fcbfc9ab432)
	default:
		return __obf_227499e666173bc3(__obf_82c50fcbfc9ab432)
	}
}

func (any *__obf_9460519df17f70ab) Keys() []string {
	__obf_51f7acc049544d5a := []string{}
	__obf_1bb30e8a74ed8233 := any.__obf_892632c77e859037.BorrowIterator(any.__obf_184433571fa55237)
	defer any.__obf_892632c77e859037.ReturnIterator(__obf_1bb30e8a74ed8233)
	__obf_1bb30e8a74ed8233.
		ReadMapCB(func(__obf_1bb30e8a74ed8233 *Iterator, __obf_7e01b5b4651074d4 string) bool {
			__obf_1bb30e8a74ed8233.
				Skip()
			__obf_51f7acc049544d5a = append(__obf_51f7acc049544d5a, __obf_7e01b5b4651074d4)
			return true
		})
	return __obf_51f7acc049544d5a
}

func (any *__obf_9460519df17f70ab) Size() int {
	__obf_dc51447b63477324 := 0
	__obf_1bb30e8a74ed8233 := any.__obf_892632c77e859037.BorrowIterator(any.__obf_184433571fa55237)
	defer any.__obf_892632c77e859037.ReturnIterator(__obf_1bb30e8a74ed8233)
	__obf_1bb30e8a74ed8233.
		ReadObjectCB(func(__obf_1bb30e8a74ed8233 *Iterator, __obf_7e01b5b4651074d4 string) bool {
			__obf_1bb30e8a74ed8233.
				Skip()
			__obf_dc51447b63477324++
			return true
		})
	return __obf_dc51447b63477324
}

func (any *__obf_9460519df17f70ab) WriteTo(__obf_850a7457bb739a32 *Stream) {
	__obf_850a7457bb739a32.
		Write(any.__obf_184433571fa55237)
}

func (any *__obf_9460519df17f70ab) GetInterface() any {
	__obf_1bb30e8a74ed8233 := any.__obf_892632c77e859037.BorrowIterator(any.__obf_184433571fa55237)
	defer any.__obf_892632c77e859037.ReturnIterator(__obf_1bb30e8a74ed8233)
	return __obf_1bb30e8a74ed8233.Read()
}

type __obf_2f60b320017f4217 struct {
	__obf_58563642f42f4a04
	__obf_62967739bca1d11d error
	__obf_bbfd2ac8618a6f0c reflect.Value
}

func __obf_1dcf4efa25cd71a9(__obf_bbfd2ac8618a6f0c any) *__obf_2f60b320017f4217 {
	return &__obf_2f60b320017f4217{__obf_58563642f42f4a04{}, nil, reflect.ValueOf(__obf_bbfd2ac8618a6f0c)}
}

func (any *__obf_2f60b320017f4217) ValueType() ValueType {
	return ObjectValue
}

func (any *__obf_2f60b320017f4217) MustBeValid() Any {
	return any
}

func (any *__obf_2f60b320017f4217) Parse() *Iterator {
	return nil
}

func (any *__obf_2f60b320017f4217) LastError() error {
	return any.__obf_62967739bca1d11d
}

func (any *__obf_2f60b320017f4217) ToBool() bool {
	return any.__obf_bbfd2ac8618a6f0c.NumField() != 0
}

func (any *__obf_2f60b320017f4217) ToInt() int {
	return 0
}

func (any *__obf_2f60b320017f4217) ToInt32() int32 {
	return 0
}

func (any *__obf_2f60b320017f4217) ToInt64() int64 {
	return 0
}

func (any *__obf_2f60b320017f4217) ToUint() uint {
	return 0
}

func (any *__obf_2f60b320017f4217) ToUint32() uint32 {
	return 0
}

func (any *__obf_2f60b320017f4217) ToUint64() uint64 {
	return 0
}

func (any *__obf_2f60b320017f4217) ToFloat32() float32 {
	return 0
}

func (any *__obf_2f60b320017f4217) ToFloat64() float64 {
	return 0
}

func (any *__obf_2f60b320017f4217) ToString() string {
	__obf_e91bd2feb751e4f1, __obf_62967739bca1d11d := MarshalToString(any.__obf_bbfd2ac8618a6f0c.Interface())
	any.__obf_62967739bca1d11d = __obf_62967739bca1d11d
	return __obf_e91bd2feb751e4f1
}

func (any *__obf_2f60b320017f4217) Get(__obf_82c50fcbfc9ab432 ...any) Any {
	if len(__obf_82c50fcbfc9ab432) == 0 {
		return any
	}
	switch __obf_c75099dc7388aeb3 := __obf_82c50fcbfc9ab432[0].(type) {
	case string:
		__obf_7e01b5b4651074d4 := any.__obf_bbfd2ac8618a6f0c.FieldByName(__obf_c75099dc7388aeb3)
		if !__obf_7e01b5b4651074d4.IsValid() {
			return __obf_227499e666173bc3(__obf_82c50fcbfc9ab432)
		}
		return Wrap(__obf_7e01b5b4651074d4.Interface())
	case int32:
		if '*' == __obf_c75099dc7388aeb3 {
			__obf_8c73a46661d509f0 := map[string]Any{}
			for __obf_5aa5c8829b97f182 := 0; __obf_5aa5c8829b97f182 < any.__obf_bbfd2ac8618a6f0c.NumField(); __obf_5aa5c8829b97f182++ {
				__obf_7e01b5b4651074d4 := any.__obf_bbfd2ac8618a6f0c.Field(__obf_5aa5c8829b97f182)
				if __obf_7e01b5b4651074d4.CanInterface() {
					__obf_d2a067019e17e3c8 := Wrap(__obf_7e01b5b4651074d4.Interface()).Get(__obf_82c50fcbfc9ab432[1:]...)
					if __obf_d2a067019e17e3c8.ValueType() != InvalidValue {
						__obf_8c73a46661d509f0[any.__obf_bbfd2ac8618a6f0c.Type().Field(__obf_5aa5c8829b97f182).Name] = __obf_d2a067019e17e3c8
					}
				}
			}
			return __obf_200d6b32a72d70af(__obf_8c73a46661d509f0)
		}
		return __obf_227499e666173bc3(__obf_82c50fcbfc9ab432)
	default:
		return __obf_227499e666173bc3(__obf_82c50fcbfc9ab432)
	}
}

func (any *__obf_2f60b320017f4217) Keys() []string {
	__obf_51f7acc049544d5a := make([]string, 0, any.__obf_bbfd2ac8618a6f0c.NumField())
	for __obf_5aa5c8829b97f182 := 0; __obf_5aa5c8829b97f182 < any.__obf_bbfd2ac8618a6f0c.NumField(); __obf_5aa5c8829b97f182++ {
		__obf_51f7acc049544d5a = append(__obf_51f7acc049544d5a, any.__obf_bbfd2ac8618a6f0c.Type().Field(__obf_5aa5c8829b97f182).Name)
	}
	return __obf_51f7acc049544d5a
}

func (any *__obf_2f60b320017f4217) Size() int {
	return any.__obf_bbfd2ac8618a6f0c.NumField()
}

func (any *__obf_2f60b320017f4217) WriteTo(__obf_850a7457bb739a32 *Stream) {
	__obf_850a7457bb739a32.
		WriteVal(any.__obf_bbfd2ac8618a6f0c)
}

func (any *__obf_2f60b320017f4217) GetInterface() any {
	return any.__obf_bbfd2ac8618a6f0c.Interface()
}

type __obf_354164febd1caa81 struct {
	__obf_58563642f42f4a04
	__obf_62967739bca1d11d error
	__obf_bbfd2ac8618a6f0c reflect.Value
}

func __obf_200d6b32a72d70af(__obf_bbfd2ac8618a6f0c any) *__obf_354164febd1caa81 {
	return &__obf_354164febd1caa81{__obf_58563642f42f4a04{}, nil, reflect.ValueOf(__obf_bbfd2ac8618a6f0c)}
}

func (any *__obf_354164febd1caa81) ValueType() ValueType {
	return ObjectValue
}

func (any *__obf_354164febd1caa81) MustBeValid() Any {
	return any
}

func (any *__obf_354164febd1caa81) Parse() *Iterator {
	return nil
}

func (any *__obf_354164febd1caa81) LastError() error {
	return any.__obf_62967739bca1d11d
}

func (any *__obf_354164febd1caa81) ToBool() bool {
	return true
}

func (any *__obf_354164febd1caa81) ToInt() int {
	return 0
}

func (any *__obf_354164febd1caa81) ToInt32() int32 {
	return 0
}

func (any *__obf_354164febd1caa81) ToInt64() int64 {
	return 0
}

func (any *__obf_354164febd1caa81) ToUint() uint {
	return 0
}

func (any *__obf_354164febd1caa81) ToUint32() uint32 {
	return 0
}

func (any *__obf_354164febd1caa81) ToUint64() uint64 {
	return 0
}

func (any *__obf_354164febd1caa81) ToFloat32() float32 {
	return 0
}

func (any *__obf_354164febd1caa81) ToFloat64() float64 {
	return 0
}

func (any *__obf_354164febd1caa81) ToString() string {
	__obf_e91bd2feb751e4f1, __obf_62967739bca1d11d := MarshalToString(any.__obf_bbfd2ac8618a6f0c.Interface())
	any.__obf_62967739bca1d11d = __obf_62967739bca1d11d
	return __obf_e91bd2feb751e4f1
}

func (any *__obf_354164febd1caa81) Get(__obf_82c50fcbfc9ab432 ...any) Any {
	if len(__obf_82c50fcbfc9ab432) == 0 {
		return any
	}
	switch __obf_c75099dc7388aeb3 := __obf_82c50fcbfc9ab432[0].(type) {
	case int32:
		if '*' == __obf_c75099dc7388aeb3 {
			__obf_8c73a46661d509f0 := map[string]Any{}
			for _, __obf_fc61abc90a2a7b45 := range any.__obf_bbfd2ac8618a6f0c.MapKeys() {
				__obf_6d89b8540e6802c9 := __obf_fc61abc90a2a7b45.String()
				__obf_03e4a7738c1c02b0 := Wrap(any.__obf_bbfd2ac8618a6f0c.MapIndex(__obf_fc61abc90a2a7b45).Interface())
				__obf_d2a067019e17e3c8 := __obf_03e4a7738c1c02b0.Get(__obf_82c50fcbfc9ab432[1:]...)
				if __obf_d2a067019e17e3c8.ValueType() != InvalidValue {
					__obf_8c73a46661d509f0[__obf_6d89b8540e6802c9] = __obf_d2a067019e17e3c8
				}
			}
			return __obf_200d6b32a72d70af(__obf_8c73a46661d509f0)
		}
		return __obf_227499e666173bc3(__obf_82c50fcbfc9ab432)
	default:
		__obf_4724d1b596d6a0c0 := any.__obf_bbfd2ac8618a6f0c.MapIndex(reflect.ValueOf(__obf_c75099dc7388aeb3))
		if !__obf_4724d1b596d6a0c0.IsValid() {
			return __obf_227499e666173bc3(__obf_82c50fcbfc9ab432)
		}
		return Wrap(__obf_4724d1b596d6a0c0.Interface())
	}
}

func (any *__obf_354164febd1caa81) Keys() []string {
	__obf_51f7acc049544d5a := make([]string, 0, any.__obf_bbfd2ac8618a6f0c.Len())
	for _, __obf_fc61abc90a2a7b45 := range any.__obf_bbfd2ac8618a6f0c.MapKeys() {
		__obf_51f7acc049544d5a = append(__obf_51f7acc049544d5a, __obf_fc61abc90a2a7b45.String())
	}
	return __obf_51f7acc049544d5a
}

func (any *__obf_354164febd1caa81) Size() int {
	return any.__obf_bbfd2ac8618a6f0c.Len()
}

func (any *__obf_354164febd1caa81) WriteTo(__obf_850a7457bb739a32 *Stream) {
	__obf_850a7457bb739a32.
		WriteVal(any.__obf_bbfd2ac8618a6f0c)
}

func (any *__obf_354164febd1caa81) GetInterface() any {
	return any.__obf_bbfd2ac8618a6f0c.Interface()
}
