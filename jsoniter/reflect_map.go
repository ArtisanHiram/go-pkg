package __obf_5b802ce8d9ba56d6

import (
	"fmt"
	"io"
	"reflect"
	"sort"
	"unsafe"

	"github.com/modern-go/reflect2"
)

func __obf_b13ed867b35030eb(__obf_08da24be66d0d13c *__obf_08da24be66d0d13c, __obf_5efc66d8c338c133 reflect2.Type) ValDecoder {
	__obf_fbe55896ffba69b4 := __obf_5efc66d8c338c133.(*reflect2.UnsafeMapType)
	__obf_ba6b21866e0588c7 := __obf_6ebaf33c79de08a9(__obf_08da24be66d0d13c.append("[mapKey]"), __obf_fbe55896ffba69b4.Key())
	__obf_45aaf5212411947b := __obf_c3a46fc9dd10c84e(__obf_08da24be66d0d13c.append("[mapElem]"), __obf_fbe55896ffba69b4.Elem())
	return &__obf_c321df75699018fd{__obf_fbe55896ffba69b4: __obf_fbe55896ffba69b4, __obf_a2884ca0ef7b7e56: __obf_fbe55896ffba69b4.Key(), __obf_0b19d2b7ea8e1be5: __obf_fbe55896ffba69b4.Elem(), __obf_ba6b21866e0588c7: __obf_ba6b21866e0588c7, __obf_45aaf5212411947b: __obf_45aaf5212411947b}
}

func __obf_8fcbff4be3a3cf2e(__obf_08da24be66d0d13c *__obf_08da24be66d0d13c, __obf_5efc66d8c338c133 reflect2.Type) ValEncoder {
	__obf_fbe55896ffba69b4 := __obf_5efc66d8c338c133.(*reflect2.UnsafeMapType)
	if __obf_08da24be66d0d13c.__obf_805c0ed9960be9bd {
		return &__obf_979124c0decc73e5{__obf_fbe55896ffba69b4: __obf_fbe55896ffba69b4, __obf_8cc36892af8499f0: __obf_1168a391e1dff270(__obf_08da24be66d0d13c.append("[mapKey]"), __obf_fbe55896ffba69b4.Key()), __obf_e57e229f9a2faf9a: __obf_3bb380f7abc03208(__obf_08da24be66d0d13c.append("[mapElem]"), __obf_fbe55896ffba69b4.Elem())}
	}
	return &__obf_a23df74fc49011e1{__obf_fbe55896ffba69b4: __obf_fbe55896ffba69b4, __obf_8cc36892af8499f0: __obf_1168a391e1dff270(__obf_08da24be66d0d13c.append("[mapKey]"), __obf_fbe55896ffba69b4.Key()), __obf_e57e229f9a2faf9a: __obf_3bb380f7abc03208(__obf_08da24be66d0d13c.append("[mapElem]"), __obf_fbe55896ffba69b4.Elem())}
}

func __obf_6ebaf33c79de08a9(__obf_08da24be66d0d13c *__obf_08da24be66d0d13c, __obf_5efc66d8c338c133 reflect2.Type) ValDecoder {
	__obf_18f746d7b5b440ee := __obf_08da24be66d0d13c.__obf_de1968c22ba7e047.CreateMapKeyDecoder(__obf_5efc66d8c338c133)
	if __obf_18f746d7b5b440ee != nil {
		return __obf_18f746d7b5b440ee
	}
	for _, __obf_6b17b29e9b6f5243 := range __obf_08da24be66d0d13c.__obf_251d7593467966e4 {
		__obf_18f746d7b5b440ee := __obf_6b17b29e9b6f5243.CreateMapKeyDecoder(__obf_5efc66d8c338c133)
		if __obf_18f746d7b5b440ee != nil {
			return __obf_18f746d7b5b440ee
		}
	}
	__obf_d0cac7bfcf0092ea := reflect2.PtrTo(__obf_5efc66d8c338c133)
	if __obf_d0cac7bfcf0092ea.Implements(__obf_738fff33396f272d) {
		return &__obf_46f8310df6424f3f{
			&__obf_1ce729a0ec8c5634{__obf_5e777ac7034ab220: __obf_d0cac7bfcf0092ea},
		}
	}
	if __obf_5efc66d8c338c133.Implements(__obf_738fff33396f272d) {
		return &__obf_1ce729a0ec8c5634{__obf_5e777ac7034ab220: __obf_5efc66d8c338c133}
	}
	if __obf_d0cac7bfcf0092ea.Implements(__obf_5cbbb40316a5c06a) {
		return &__obf_46f8310df6424f3f{
			&__obf_646b36ac5dd53924{__obf_5e777ac7034ab220: __obf_d0cac7bfcf0092ea},
		}
	}
	if __obf_5efc66d8c338c133.Implements(__obf_5cbbb40316a5c06a) {
		return &__obf_646b36ac5dd53924{__obf_5e777ac7034ab220: __obf_5efc66d8c338c133}
	}

	switch __obf_5efc66d8c338c133.Kind() {
	case reflect.String:
		return __obf_c3a46fc9dd10c84e(__obf_08da24be66d0d13c, reflect2.DefaultTypeOfKind(reflect.String))
	case reflect.Bool,
		reflect.Uint8, reflect.Int8,
		reflect.Uint16, reflect.Int16,
		reflect.Uint32, reflect.Int32,
		reflect.Uint64, reflect.Int64,
		reflect.Uint, reflect.Int,
		reflect.Float32, reflect.Float64,
		reflect.Uintptr:
		__obf_5efc66d8c338c133 = reflect2.DefaultTypeOfKind(__obf_5efc66d8c338c133.Kind())
		return &__obf_4b2c64dc9ee0d8cd{__obf_c3a46fc9dd10c84e(__obf_08da24be66d0d13c, __obf_5efc66d8c338c133)}
	default:
		return &__obf_a12de69c81d1c528{__obf_6d071d48f3b321ad: fmt.Errorf("unsupported map key type: %v", __obf_5efc66d8c338c133)}
	}
}

func __obf_1168a391e1dff270(__obf_08da24be66d0d13c *__obf_08da24be66d0d13c, __obf_5efc66d8c338c133 reflect2.Type) ValEncoder {
	__obf_29366c3ad76a8c51 := __obf_08da24be66d0d13c.__obf_f754da59a9c09bdc.CreateMapKeyEncoder(__obf_5efc66d8c338c133)
	if __obf_29366c3ad76a8c51 != nil {
		return __obf_29366c3ad76a8c51
	}
	for _, __obf_6b17b29e9b6f5243 := range __obf_08da24be66d0d13c.__obf_251d7593467966e4 {
		__obf_29366c3ad76a8c51 := __obf_6b17b29e9b6f5243.CreateMapKeyEncoder(__obf_5efc66d8c338c133)
		if __obf_29366c3ad76a8c51 != nil {
			return __obf_29366c3ad76a8c51
		}
	}

	if __obf_5efc66d8c338c133.Kind() != reflect.String {
		if __obf_5efc66d8c338c133 == __obf_9dca2b8aa916a9d0 {
			return &__obf_8201f27af3b81439{__obf_29006411c501f852: __obf_08da24be66d0d13c.EncoderOf(reflect2.TypeOf(""))}
		}
		if __obf_5efc66d8c338c133.Implements(__obf_9dca2b8aa916a9d0) {
			return &__obf_3290483a33b841e1{__obf_5e777ac7034ab220: __obf_5efc66d8c338c133, __obf_29006411c501f852: __obf_08da24be66d0d13c.EncoderOf(reflect2.TypeOf(""))}
		}
	}

	switch __obf_5efc66d8c338c133.Kind() {
	case reflect.String:
		return __obf_3bb380f7abc03208(__obf_08da24be66d0d13c, reflect2.DefaultTypeOfKind(reflect.String))
	case reflect.Bool,
		reflect.Uint8, reflect.Int8,
		reflect.Uint16, reflect.Int16,
		reflect.Uint32, reflect.Int32,
		reflect.Uint64, reflect.Int64,
		reflect.Uint, reflect.Int,
		reflect.Float32, reflect.Float64,
		reflect.Uintptr:
		__obf_5efc66d8c338c133 = reflect2.DefaultTypeOfKind(__obf_5efc66d8c338c133.Kind())
		return &__obf_fadeee7b4108e895{__obf_3bb380f7abc03208(__obf_08da24be66d0d13c, __obf_5efc66d8c338c133)}
	default:
		if __obf_5efc66d8c338c133.Kind() == reflect.Interface {
			return &__obf_0bd7ffe15ed7683a{__obf_08da24be66d0d13c, __obf_5efc66d8c338c133}
		}
		return &__obf_5d9ef7d0fc378907{__obf_6d071d48f3b321ad: fmt.Errorf("unsupported map key type: %v", __obf_5efc66d8c338c133)}
	}
}

type __obf_c321df75699018fd struct {
	__obf_fbe55896ffba69b4 *reflect2.UnsafeMapType
	__obf_a2884ca0ef7b7e56 reflect2.Type
	__obf_0b19d2b7ea8e1be5 reflect2.Type
	__obf_ba6b21866e0588c7 ValDecoder
	__obf_45aaf5212411947b ValDecoder
}

func (__obf_18f746d7b5b440ee *__obf_c321df75699018fd) Decode(__obf_d3c919547bf7e47a unsafe.Pointer, __obf_67008a6a9e5ba828 *Iterator) {
	__obf_fbe55896ffba69b4 := __obf_18f746d7b5b440ee.__obf_fbe55896ffba69b4
	__obf_dab9baaadfa7c8c2 := __obf_67008a6a9e5ba828.__obf_b781a59d5a0d2490()
	if __obf_dab9baaadfa7c8c2 == 'n' {
		__obf_67008a6a9e5ba828.__obf_4aeb767e0be7277a('u', 'l', 'l')
		*(*unsafe.Pointer)(__obf_d3c919547bf7e47a) = nil
		__obf_fbe55896ffba69b4.
			UnsafeSet(__obf_d3c919547bf7e47a, __obf_fbe55896ffba69b4.UnsafeNew())
		return
	}
	if __obf_fbe55896ffba69b4.UnsafeIsNil(__obf_d3c919547bf7e47a) {
		__obf_fbe55896ffba69b4.
			UnsafeSet(__obf_d3c919547bf7e47a, __obf_fbe55896ffba69b4.UnsafeMakeMap(0))
	}
	if __obf_dab9baaadfa7c8c2 != '{' {
		__obf_67008a6a9e5ba828.
			ReportError("ReadMapCB", `expect { or n, but found `+string([]byte{__obf_dab9baaadfa7c8c2}))
		return
	}
	__obf_dab9baaadfa7c8c2 = __obf_67008a6a9e5ba828.__obf_b781a59d5a0d2490()
	if __obf_dab9baaadfa7c8c2 == '}' {
		return
	}
	__obf_67008a6a9e5ba828.__obf_3284a1eaa2a0abb6()
	__obf_6998442f85330120 := __obf_18f746d7b5b440ee.__obf_a2884ca0ef7b7e56.UnsafeNew()
	__obf_18f746d7b5b440ee.__obf_ba6b21866e0588c7.
		Decode(__obf_6998442f85330120, __obf_67008a6a9e5ba828)
	__obf_dab9baaadfa7c8c2 = __obf_67008a6a9e5ba828.__obf_b781a59d5a0d2490()
	if __obf_dab9baaadfa7c8c2 != ':' {
		__obf_67008a6a9e5ba828.
			ReportError("ReadMapCB", "expect : after object field, but found "+string([]byte{__obf_dab9baaadfa7c8c2}))
		return
	}
	__obf_77670b1f1899744d := __obf_18f746d7b5b440ee.__obf_0b19d2b7ea8e1be5.UnsafeNew()
	__obf_18f746d7b5b440ee.__obf_45aaf5212411947b.
		Decode(__obf_77670b1f1899744d, __obf_67008a6a9e5ba828)
	__obf_18f746d7b5b440ee.__obf_fbe55896ffba69b4.
		UnsafeSetIndex(__obf_d3c919547bf7e47a, __obf_6998442f85330120, __obf_77670b1f1899744d)
	for __obf_dab9baaadfa7c8c2 = __obf_67008a6a9e5ba828.__obf_b781a59d5a0d2490(); __obf_dab9baaadfa7c8c2 == ','; __obf_dab9baaadfa7c8c2 = __obf_67008a6a9e5ba828.__obf_b781a59d5a0d2490() {
		__obf_6998442f85330120 := __obf_18f746d7b5b440ee.__obf_a2884ca0ef7b7e56.UnsafeNew()
		__obf_18f746d7b5b440ee.__obf_ba6b21866e0588c7.
			Decode(__obf_6998442f85330120, __obf_67008a6a9e5ba828)
		__obf_dab9baaadfa7c8c2 = __obf_67008a6a9e5ba828.__obf_b781a59d5a0d2490()
		if __obf_dab9baaadfa7c8c2 != ':' {
			__obf_67008a6a9e5ba828.
				ReportError("ReadMapCB", "expect : after object field, but found "+string([]byte{__obf_dab9baaadfa7c8c2}))
			return
		}
		__obf_77670b1f1899744d := __obf_18f746d7b5b440ee.__obf_0b19d2b7ea8e1be5.UnsafeNew()
		__obf_18f746d7b5b440ee.__obf_45aaf5212411947b.
			Decode(__obf_77670b1f1899744d, __obf_67008a6a9e5ba828)
		__obf_18f746d7b5b440ee.__obf_fbe55896ffba69b4.
			UnsafeSetIndex(__obf_d3c919547bf7e47a, __obf_6998442f85330120, __obf_77670b1f1899744d)
	}
	if __obf_dab9baaadfa7c8c2 != '}' {
		__obf_67008a6a9e5ba828.
			ReportError("ReadMapCB", `expect }, but found `+string([]byte{__obf_dab9baaadfa7c8c2}))
	}
}

type __obf_4b2c64dc9ee0d8cd struct {
	__obf_18f746d7b5b440ee ValDecoder
}

func (__obf_18f746d7b5b440ee *__obf_4b2c64dc9ee0d8cd) Decode(__obf_d3c919547bf7e47a unsafe.Pointer, __obf_67008a6a9e5ba828 *Iterator) {
	__obf_dab9baaadfa7c8c2 := __obf_67008a6a9e5ba828.__obf_b781a59d5a0d2490()
	if __obf_dab9baaadfa7c8c2 != '"' {
		__obf_67008a6a9e5ba828.
			ReportError("ReadMapCB", `expect ", but found `+string([]byte{__obf_dab9baaadfa7c8c2}))
		return
	}
	__obf_18f746d7b5b440ee.__obf_18f746d7b5b440ee.
		Decode(__obf_d3c919547bf7e47a, __obf_67008a6a9e5ba828)
	__obf_dab9baaadfa7c8c2 = __obf_67008a6a9e5ba828.__obf_b781a59d5a0d2490()
	if __obf_dab9baaadfa7c8c2 != '"' {
		__obf_67008a6a9e5ba828.
			ReportError("ReadMapCB", `expect ", but found `+string([]byte{__obf_dab9baaadfa7c8c2}))
		return
	}
}

type __obf_fadeee7b4108e895 struct {
	__obf_29366c3ad76a8c51 ValEncoder
}

func (__obf_29366c3ad76a8c51 *__obf_fadeee7b4108e895) Encode(__obf_d3c919547bf7e47a unsafe.Pointer, __obf_00fc491caa967a74 *Stream) {
	__obf_00fc491caa967a74.__obf_487ee8035c7dd8f6('"')
	__obf_29366c3ad76a8c51.__obf_29366c3ad76a8c51.
		Encode(__obf_d3c919547bf7e47a, __obf_00fc491caa967a74)
	__obf_00fc491caa967a74.__obf_487ee8035c7dd8f6('"')
}

func (__obf_29366c3ad76a8c51 *__obf_fadeee7b4108e895) IsEmpty(__obf_d3c919547bf7e47a unsafe.Pointer) bool {
	return false
}

type __obf_0bd7ffe15ed7683a struct {
	__obf_08da24be66d0d13c *__obf_08da24be66d0d13c
	__obf_5e777ac7034ab220 reflect2.Type
}

func (__obf_29366c3ad76a8c51 *__obf_0bd7ffe15ed7683a) Encode(__obf_d3c919547bf7e47a unsafe.Pointer, __obf_00fc491caa967a74 *Stream) {
	__obf_f9b1526531f3c024 := __obf_29366c3ad76a8c51.__obf_5e777ac7034ab220.UnsafeIndirect(__obf_d3c919547bf7e47a)
	__obf_1168a391e1dff270(__obf_29366c3ad76a8c51.__obf_08da24be66d0d13c, reflect2.TypeOf(__obf_f9b1526531f3c024)).Encode(reflect2.PtrOf(__obf_f9b1526531f3c024), __obf_00fc491caa967a74)
}

func (__obf_29366c3ad76a8c51 *__obf_0bd7ffe15ed7683a) IsEmpty(__obf_d3c919547bf7e47a unsafe.Pointer) bool {
	__obf_f9b1526531f3c024 := __obf_29366c3ad76a8c51.__obf_5e777ac7034ab220.UnsafeIndirect(__obf_d3c919547bf7e47a)
	return __obf_1168a391e1dff270(__obf_29366c3ad76a8c51.__obf_08da24be66d0d13c, reflect2.TypeOf(__obf_f9b1526531f3c024)).IsEmpty(reflect2.PtrOf(__obf_f9b1526531f3c024))
}

type __obf_a23df74fc49011e1 struct {
	__obf_fbe55896ffba69b4 *reflect2.UnsafeMapType
	__obf_8cc36892af8499f0 ValEncoder
	__obf_e57e229f9a2faf9a ValEncoder
}

func (__obf_29366c3ad76a8c51 *__obf_a23df74fc49011e1) Encode(__obf_d3c919547bf7e47a unsafe.Pointer, __obf_00fc491caa967a74 *Stream) {
	if *(*unsafe.Pointer)(__obf_d3c919547bf7e47a) == nil {
		__obf_00fc491caa967a74.
			WriteNil()
		return
	}
	__obf_00fc491caa967a74.
		WriteObjectStart()
	__obf_67008a6a9e5ba828 := __obf_29366c3ad76a8c51.__obf_fbe55896ffba69b4.UnsafeIterate(__obf_d3c919547bf7e47a)
	for __obf_2deec7c38ffb6ae3 := 0; __obf_67008a6a9e5ba828.HasNext(); __obf_2deec7c38ffb6ae3++ {
		if __obf_2deec7c38ffb6ae3 != 0 {
			__obf_00fc491caa967a74.
				WriteMore()
		}
		__obf_6998442f85330120, __obf_77670b1f1899744d := __obf_67008a6a9e5ba828.UnsafeNext()
		__obf_29366c3ad76a8c51.__obf_8cc36892af8499f0.
			Encode(__obf_6998442f85330120, __obf_00fc491caa967a74)
		if __obf_00fc491caa967a74.__obf_e68cefffff828d90 > 0 {
			__obf_00fc491caa967a74.__obf_dd59782cf69113a3(byte(':'), byte(' '))
		} else {
			__obf_00fc491caa967a74.__obf_487ee8035c7dd8f6(':')
		}
		__obf_29366c3ad76a8c51.__obf_e57e229f9a2faf9a.
			Encode(__obf_77670b1f1899744d, __obf_00fc491caa967a74)
	}
	__obf_00fc491caa967a74.
		WriteObjectEnd()
}

func (__obf_29366c3ad76a8c51 *__obf_a23df74fc49011e1) IsEmpty(__obf_d3c919547bf7e47a unsafe.Pointer) bool {
	__obf_67008a6a9e5ba828 := __obf_29366c3ad76a8c51.__obf_fbe55896ffba69b4.UnsafeIterate(__obf_d3c919547bf7e47a)
	return !__obf_67008a6a9e5ba828.HasNext()
}

type __obf_979124c0decc73e5 struct {
	__obf_fbe55896ffba69b4 *reflect2.UnsafeMapType
	__obf_8cc36892af8499f0 ValEncoder
	__obf_e57e229f9a2faf9a ValEncoder
}

func (__obf_29366c3ad76a8c51 *__obf_979124c0decc73e5) Encode(__obf_d3c919547bf7e47a unsafe.Pointer, __obf_00fc491caa967a74 *Stream) {
	if *(*unsafe.Pointer)(__obf_d3c919547bf7e47a) == nil {
		__obf_00fc491caa967a74.
			WriteNil()
		return
	}
	__obf_00fc491caa967a74.
		WriteObjectStart()
	__obf_f9eb7766c5644b54 := __obf_29366c3ad76a8c51.__obf_fbe55896ffba69b4.UnsafeIterate(__obf_d3c919547bf7e47a)
	__obf_4b295e47f4216f3f := __obf_00fc491caa967a74.__obf_dca106e1cdf85392.BorrowStream(nil)
	__obf_4b295e47f4216f3f.
		Attachment = __obf_00fc491caa967a74.Attachment
	__obf_50929c65e436c5e8 := __obf_00fc491caa967a74.__obf_dca106e1cdf85392.BorrowIterator(nil)
	__obf_27b74d1e79a5eca4 := __obf_363c0b331bbcc76c{}
	for __obf_f9eb7766c5644b54.HasNext() {
		__obf_6998442f85330120, __obf_77670b1f1899744d := __obf_f9eb7766c5644b54.UnsafeNext()
		__obf_aa119ebbfdd94885 := __obf_4b295e47f4216f3f.Buffered()
		__obf_29366c3ad76a8c51.__obf_8cc36892af8499f0.
			Encode(__obf_6998442f85330120, __obf_4b295e47f4216f3f)
		if __obf_4b295e47f4216f3f.Error != nil && __obf_4b295e47f4216f3f.Error != io.EOF && __obf_00fc491caa967a74.Error == nil {
			__obf_00fc491caa967a74.
				Error = __obf_4b295e47f4216f3f.Error
		}
		__obf_a539367a37ebbe53 := __obf_4b295e47f4216f3f.Buffer()[__obf_aa119ebbfdd94885:]
		__obf_50929c65e436c5e8.
			ResetBytes(__obf_a539367a37ebbe53)
		__obf_9f0b6145adaac566 := __obf_50929c65e436c5e8.ReadString()
		if __obf_00fc491caa967a74.__obf_e68cefffff828d90 > 0 {
			__obf_4b295e47f4216f3f.__obf_dd59782cf69113a3(byte(':'), byte(' '))
		} else {
			__obf_4b295e47f4216f3f.__obf_487ee8035c7dd8f6(':')
		}
		__obf_29366c3ad76a8c51.__obf_e57e229f9a2faf9a.
			Encode(__obf_77670b1f1899744d, __obf_4b295e47f4216f3f)
		__obf_27b74d1e79a5eca4 = append(__obf_27b74d1e79a5eca4, __obf_12b5c7b31c19fcd1{__obf_6998442f85330120: __obf_9f0b6145adaac566, __obf_23f4faed92c890f8: __obf_4b295e47f4216f3f.Buffer()[__obf_aa119ebbfdd94885:]})
	}
	sort.Sort(__obf_27b74d1e79a5eca4)
	for __obf_2deec7c38ffb6ae3, __obf_23f4faed92c890f8 := range __obf_27b74d1e79a5eca4 {
		if __obf_2deec7c38ffb6ae3 != 0 {
			__obf_00fc491caa967a74.
				WriteMore()
		}
		__obf_00fc491caa967a74.
			Write(__obf_23f4faed92c890f8.__obf_23f4faed92c890f8)
	}
	if __obf_4b295e47f4216f3f.Error != nil && __obf_00fc491caa967a74.Error == nil {
		__obf_00fc491caa967a74.
			Error = __obf_4b295e47f4216f3f.Error
	}
	__obf_00fc491caa967a74.
		WriteObjectEnd()
	__obf_00fc491caa967a74.__obf_dca106e1cdf85392.
		ReturnStream(__obf_4b295e47f4216f3f)
	__obf_00fc491caa967a74.__obf_dca106e1cdf85392.
		ReturnIterator(__obf_50929c65e436c5e8)
}

func (__obf_29366c3ad76a8c51 *__obf_979124c0decc73e5) IsEmpty(__obf_d3c919547bf7e47a unsafe.Pointer) bool {
	__obf_67008a6a9e5ba828 := __obf_29366c3ad76a8c51.__obf_fbe55896ffba69b4.UnsafeIterate(__obf_d3c919547bf7e47a)
	return !__obf_67008a6a9e5ba828.HasNext()
}

type __obf_363c0b331bbcc76c []__obf_12b5c7b31c19fcd1

type __obf_12b5c7b31c19fcd1 struct {
	__obf_6998442f85330120 string
	__obf_23f4faed92c890f8 []byte
}

func (__obf_16a45220fceac36d __obf_363c0b331bbcc76c) Len() int { return len(__obf_16a45220fceac36d) }
func (__obf_16a45220fceac36d __obf_363c0b331bbcc76c) Swap(__obf_2deec7c38ffb6ae3, __obf_973404809dee3093 int) {
	__obf_16a45220fceac36d[__obf_2deec7c38ffb6ae3], __obf_16a45220fceac36d[__obf_973404809dee3093] = __obf_16a45220fceac36d[__obf_973404809dee3093], __obf_16a45220fceac36d[__obf_2deec7c38ffb6ae3]
}
func (__obf_16a45220fceac36d __obf_363c0b331bbcc76c) Less(__obf_2deec7c38ffb6ae3, __obf_973404809dee3093 int) bool {
	return __obf_16a45220fceac36d[__obf_2deec7c38ffb6ae3].__obf_6998442f85330120 < __obf_16a45220fceac36d[__obf_973404809dee3093].__obf_6998442f85330120
}
