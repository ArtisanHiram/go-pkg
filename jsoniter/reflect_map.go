package __obf_91620b895eeff9ed

import (
	"fmt"
	"io"
	"reflect"
	"sort"
	"unsafe"

	"github.com/modern-go/reflect2"
)

func __obf_04b67c87ffeaab64(__obf_2f9c5aed866cce75 *__obf_2f9c5aed866cce75, __obf_29ebd0f2c324f5ea reflect2.Type) ValDecoder {
	__obf_08a00a78c23c9481 := __obf_29ebd0f2c324f5ea.(*reflect2.UnsafeMapType)
	__obf_147a5351cadb41fb := __obf_968e73852121dd3f(__obf_2f9c5aed866cce75.append("[mapKey]"), __obf_08a00a78c23c9481.Key())
	__obf_0110b2c027051777 := __obf_0b44a7afc1523314(__obf_2f9c5aed866cce75.append("[mapElem]"), __obf_08a00a78c23c9481.Elem())
	return &__obf_6add9c8dc4379363{__obf_08a00a78c23c9481: __obf_08a00a78c23c9481, __obf_b0e7c0b32b50aefc: __obf_08a00a78c23c9481.Key(), __obf_eef27b4522cbbab1: __obf_08a00a78c23c9481.Elem(), __obf_147a5351cadb41fb: __obf_147a5351cadb41fb, __obf_0110b2c027051777: __obf_0110b2c027051777}
}

func __obf_00ff8b3760381bbe(__obf_2f9c5aed866cce75 *__obf_2f9c5aed866cce75, __obf_29ebd0f2c324f5ea reflect2.Type) ValEncoder {
	__obf_08a00a78c23c9481 := __obf_29ebd0f2c324f5ea.(*reflect2.UnsafeMapType)
	if __obf_2f9c5aed866cce75.__obf_e63708933643279b {
		return &__obf_3c655ef2fb5953e5{__obf_08a00a78c23c9481: __obf_08a00a78c23c9481, __obf_38ad2485e5b56490: __obf_11b7350fb8580af4(__obf_2f9c5aed866cce75.append("[mapKey]"), __obf_08a00a78c23c9481.Key()), __obf_11328395cf86586a: __obf_d1233db7a73cc96c(__obf_2f9c5aed866cce75.append("[mapElem]"), __obf_08a00a78c23c9481.Elem())}
	}
	return &__obf_ef22fb2a3a24383f{__obf_08a00a78c23c9481: __obf_08a00a78c23c9481, __obf_38ad2485e5b56490: __obf_11b7350fb8580af4(__obf_2f9c5aed866cce75.append("[mapKey]"), __obf_08a00a78c23c9481.Key()), __obf_11328395cf86586a: __obf_d1233db7a73cc96c(__obf_2f9c5aed866cce75.append("[mapElem]"), __obf_08a00a78c23c9481.Elem())}
}

func __obf_968e73852121dd3f(__obf_2f9c5aed866cce75 *__obf_2f9c5aed866cce75, __obf_29ebd0f2c324f5ea reflect2.Type) ValDecoder {
	__obf_6fd3f72eb9b5574c := __obf_2f9c5aed866cce75.__obf_920e2f1ddf47b5e1.CreateMapKeyDecoder(__obf_29ebd0f2c324f5ea)
	if __obf_6fd3f72eb9b5574c != nil {
		return __obf_6fd3f72eb9b5574c
	}
	for _, __obf_8c9d73a8f8319687 := range __obf_2f9c5aed866cce75.__obf_b4dfae156c7ac26c {
		__obf_6fd3f72eb9b5574c := __obf_8c9d73a8f8319687.CreateMapKeyDecoder(__obf_29ebd0f2c324f5ea)
		if __obf_6fd3f72eb9b5574c != nil {
			return __obf_6fd3f72eb9b5574c
		}
	}
	__obf_f2fdafeb141957bd := reflect2.PtrTo(__obf_29ebd0f2c324f5ea)
	if __obf_f2fdafeb141957bd.Implements(__obf_731bba38e4f2c30a) {
		return &__obf_a82dfb83d40f1b70{
			&__obf_c90773957546cc29{__obf_5ea8ee327a6f7e0d: __obf_f2fdafeb141957bd},
		}
	}
	if __obf_29ebd0f2c324f5ea.Implements(__obf_731bba38e4f2c30a) {
		return &__obf_c90773957546cc29{__obf_5ea8ee327a6f7e0d: __obf_29ebd0f2c324f5ea}
	}
	if __obf_f2fdafeb141957bd.Implements(__obf_39d8082c5a1ebfa8) {
		return &__obf_a82dfb83d40f1b70{
			&__obf_ebde06d9f0cb70ce{__obf_5ea8ee327a6f7e0d: __obf_f2fdafeb141957bd},
		}
	}
	if __obf_29ebd0f2c324f5ea.Implements(__obf_39d8082c5a1ebfa8) {
		return &__obf_ebde06d9f0cb70ce{__obf_5ea8ee327a6f7e0d: __obf_29ebd0f2c324f5ea}
	}

	switch __obf_29ebd0f2c324f5ea.Kind() {
	case reflect.String:
		return __obf_0b44a7afc1523314(__obf_2f9c5aed866cce75, reflect2.DefaultTypeOfKind(reflect.String))
	case reflect.Bool,
		reflect.Uint8, reflect.Int8,
		reflect.Uint16, reflect.Int16,
		reflect.Uint32, reflect.Int32,
		reflect.Uint64, reflect.Int64,
		reflect.Uint, reflect.Int,
		reflect.Float32, reflect.Float64,
		reflect.Uintptr:
		__obf_29ebd0f2c324f5ea = reflect2.DefaultTypeOfKind(__obf_29ebd0f2c324f5ea.Kind())
		return &__obf_92b5aaa0aa6b4949{__obf_0b44a7afc1523314(__obf_2f9c5aed866cce75, __obf_29ebd0f2c324f5ea)}
	default:
		return &__obf_b540734d2036f0d8{__obf_62967739bca1d11d: fmt.Errorf("unsupported map key type: %v", __obf_29ebd0f2c324f5ea)}
	}
}

func __obf_11b7350fb8580af4(__obf_2f9c5aed866cce75 *__obf_2f9c5aed866cce75, __obf_29ebd0f2c324f5ea reflect2.Type) ValEncoder {
	__obf_96e65a4c8c4f2ce5 := __obf_2f9c5aed866cce75.__obf_47929f7efe51b371.CreateMapKeyEncoder(__obf_29ebd0f2c324f5ea)
	if __obf_96e65a4c8c4f2ce5 != nil {
		return __obf_96e65a4c8c4f2ce5
	}
	for _, __obf_8c9d73a8f8319687 := range __obf_2f9c5aed866cce75.__obf_b4dfae156c7ac26c {
		__obf_96e65a4c8c4f2ce5 := __obf_8c9d73a8f8319687.CreateMapKeyEncoder(__obf_29ebd0f2c324f5ea)
		if __obf_96e65a4c8c4f2ce5 != nil {
			return __obf_96e65a4c8c4f2ce5
		}
	}

	if __obf_29ebd0f2c324f5ea.Kind() != reflect.String {
		if __obf_29ebd0f2c324f5ea == __obf_c32cd3ec60132ef8 {
			return &__obf_df238ef33130aad9{__obf_7ebb8230b2719231: __obf_2f9c5aed866cce75.EncoderOf(reflect2.TypeOf(""))}
		}
		if __obf_29ebd0f2c324f5ea.Implements(__obf_c32cd3ec60132ef8) {
			return &__obf_9f57b5893251d4b1{__obf_5ea8ee327a6f7e0d: __obf_29ebd0f2c324f5ea, __obf_7ebb8230b2719231: __obf_2f9c5aed866cce75.EncoderOf(reflect2.TypeOf(""))}
		}
	}

	switch __obf_29ebd0f2c324f5ea.Kind() {
	case reflect.String:
		return __obf_d1233db7a73cc96c(__obf_2f9c5aed866cce75, reflect2.DefaultTypeOfKind(reflect.String))
	case reflect.Bool,
		reflect.Uint8, reflect.Int8,
		reflect.Uint16, reflect.Int16,
		reflect.Uint32, reflect.Int32,
		reflect.Uint64, reflect.Int64,
		reflect.Uint, reflect.Int,
		reflect.Float32, reflect.Float64,
		reflect.Uintptr:
		__obf_29ebd0f2c324f5ea = reflect2.DefaultTypeOfKind(__obf_29ebd0f2c324f5ea.Kind())
		return &__obf_9d8e2fdc19ab0434{__obf_d1233db7a73cc96c(__obf_2f9c5aed866cce75, __obf_29ebd0f2c324f5ea)}
	default:
		if __obf_29ebd0f2c324f5ea.Kind() == reflect.Interface {
			return &__obf_653048ea55761927{__obf_2f9c5aed866cce75, __obf_29ebd0f2c324f5ea}
		}
		return &__obf_01cb1af95bf874f9{__obf_62967739bca1d11d: fmt.Errorf("unsupported map key type: %v", __obf_29ebd0f2c324f5ea)}
	}
}

type __obf_6add9c8dc4379363 struct {
	__obf_08a00a78c23c9481 *reflect2.UnsafeMapType
	__obf_b0e7c0b32b50aefc reflect2.Type
	__obf_eef27b4522cbbab1 reflect2.Type
	__obf_147a5351cadb41fb ValDecoder
	__obf_0110b2c027051777 ValDecoder
}

func (__obf_6fd3f72eb9b5574c *__obf_6add9c8dc4379363) Decode(__obf_2a1474febb02279b unsafe.Pointer, __obf_1bb30e8a74ed8233 *Iterator) {
	__obf_08a00a78c23c9481 := __obf_6fd3f72eb9b5574c.__obf_08a00a78c23c9481
	__obf_f16b4157911bc9af := __obf_1bb30e8a74ed8233.__obf_684faa48ae8c5049()
	if __obf_f16b4157911bc9af == 'n' {
		__obf_1bb30e8a74ed8233.__obf_3e1d2ad9a54f0d22('u', 'l', 'l')
		*(*unsafe.Pointer)(__obf_2a1474febb02279b) = nil
		__obf_08a00a78c23c9481.
			UnsafeSet(__obf_2a1474febb02279b, __obf_08a00a78c23c9481.UnsafeNew())
		return
	}
	if __obf_08a00a78c23c9481.UnsafeIsNil(__obf_2a1474febb02279b) {
		__obf_08a00a78c23c9481.
			UnsafeSet(__obf_2a1474febb02279b, __obf_08a00a78c23c9481.UnsafeMakeMap(0))
	}
	if __obf_f16b4157911bc9af != '{' {
		__obf_1bb30e8a74ed8233.
			ReportError("ReadMapCB", `expect { or n, but found `+string([]byte{__obf_f16b4157911bc9af}))
		return
	}
	__obf_f16b4157911bc9af = __obf_1bb30e8a74ed8233.__obf_684faa48ae8c5049()
	if __obf_f16b4157911bc9af == '}' {
		return
	}
	__obf_1bb30e8a74ed8233.__obf_a163df67f9bb1c4b()
	__obf_fc61abc90a2a7b45 := __obf_6fd3f72eb9b5574c.__obf_b0e7c0b32b50aefc.UnsafeNew()
	__obf_6fd3f72eb9b5574c.__obf_147a5351cadb41fb.
		Decode(__obf_fc61abc90a2a7b45, __obf_1bb30e8a74ed8233)
	__obf_f16b4157911bc9af = __obf_1bb30e8a74ed8233.__obf_684faa48ae8c5049()
	if __obf_f16b4157911bc9af != ':' {
		__obf_1bb30e8a74ed8233.
			ReportError("ReadMapCB", "expect : after object field, but found "+string([]byte{__obf_f16b4157911bc9af}))
		return
	}
	__obf_97d7079fd160f22c := __obf_6fd3f72eb9b5574c.__obf_eef27b4522cbbab1.UnsafeNew()
	__obf_6fd3f72eb9b5574c.__obf_0110b2c027051777.
		Decode(__obf_97d7079fd160f22c, __obf_1bb30e8a74ed8233)
	__obf_6fd3f72eb9b5574c.__obf_08a00a78c23c9481.
		UnsafeSetIndex(__obf_2a1474febb02279b, __obf_fc61abc90a2a7b45, __obf_97d7079fd160f22c)
	for __obf_f16b4157911bc9af = __obf_1bb30e8a74ed8233.__obf_684faa48ae8c5049(); __obf_f16b4157911bc9af == ','; __obf_f16b4157911bc9af = __obf_1bb30e8a74ed8233.__obf_684faa48ae8c5049() {
		__obf_fc61abc90a2a7b45 := __obf_6fd3f72eb9b5574c.__obf_b0e7c0b32b50aefc.UnsafeNew()
		__obf_6fd3f72eb9b5574c.__obf_147a5351cadb41fb.
			Decode(__obf_fc61abc90a2a7b45, __obf_1bb30e8a74ed8233)
		__obf_f16b4157911bc9af = __obf_1bb30e8a74ed8233.__obf_684faa48ae8c5049()
		if __obf_f16b4157911bc9af != ':' {
			__obf_1bb30e8a74ed8233.
				ReportError("ReadMapCB", "expect : after object field, but found "+string([]byte{__obf_f16b4157911bc9af}))
			return
		}
		__obf_97d7079fd160f22c := __obf_6fd3f72eb9b5574c.__obf_eef27b4522cbbab1.UnsafeNew()
		__obf_6fd3f72eb9b5574c.__obf_0110b2c027051777.
			Decode(__obf_97d7079fd160f22c, __obf_1bb30e8a74ed8233)
		__obf_6fd3f72eb9b5574c.__obf_08a00a78c23c9481.
			UnsafeSetIndex(__obf_2a1474febb02279b, __obf_fc61abc90a2a7b45, __obf_97d7079fd160f22c)
	}
	if __obf_f16b4157911bc9af != '}' {
		__obf_1bb30e8a74ed8233.
			ReportError("ReadMapCB", `expect }, but found `+string([]byte{__obf_f16b4157911bc9af}))
	}
}

type __obf_92b5aaa0aa6b4949 struct {
	__obf_6fd3f72eb9b5574c ValDecoder
}

func (__obf_6fd3f72eb9b5574c *__obf_92b5aaa0aa6b4949) Decode(__obf_2a1474febb02279b unsafe.Pointer, __obf_1bb30e8a74ed8233 *Iterator) {
	__obf_f16b4157911bc9af := __obf_1bb30e8a74ed8233.__obf_684faa48ae8c5049()
	if __obf_f16b4157911bc9af != '"' {
		__obf_1bb30e8a74ed8233.
			ReportError("ReadMapCB", `expect ", but found `+string([]byte{__obf_f16b4157911bc9af}))
		return
	}
	__obf_6fd3f72eb9b5574c.__obf_6fd3f72eb9b5574c.
		Decode(__obf_2a1474febb02279b, __obf_1bb30e8a74ed8233)
	__obf_f16b4157911bc9af = __obf_1bb30e8a74ed8233.__obf_684faa48ae8c5049()
	if __obf_f16b4157911bc9af != '"' {
		__obf_1bb30e8a74ed8233.
			ReportError("ReadMapCB", `expect ", but found `+string([]byte{__obf_f16b4157911bc9af}))
		return
	}
}

type __obf_9d8e2fdc19ab0434 struct {
	__obf_96e65a4c8c4f2ce5 ValEncoder
}

func (__obf_96e65a4c8c4f2ce5 *__obf_9d8e2fdc19ab0434) Encode(__obf_2a1474febb02279b unsafe.Pointer, __obf_850a7457bb739a32 *Stream) {
	__obf_850a7457bb739a32.__obf_72837f6fe737f078('"')
	__obf_96e65a4c8c4f2ce5.__obf_96e65a4c8c4f2ce5.
		Encode(__obf_2a1474febb02279b, __obf_850a7457bb739a32)
	__obf_850a7457bb739a32.__obf_72837f6fe737f078('"')
}

func (__obf_96e65a4c8c4f2ce5 *__obf_9d8e2fdc19ab0434) IsEmpty(__obf_2a1474febb02279b unsafe.Pointer) bool {
	return false
}

type __obf_653048ea55761927 struct {
	__obf_2f9c5aed866cce75 *__obf_2f9c5aed866cce75
	__obf_5ea8ee327a6f7e0d reflect2.Type
}

func (__obf_96e65a4c8c4f2ce5 *__obf_653048ea55761927) Encode(__obf_2a1474febb02279b unsafe.Pointer, __obf_850a7457bb739a32 *Stream) {
	__obf_35136ef2ac0f94e4 := __obf_96e65a4c8c4f2ce5.__obf_5ea8ee327a6f7e0d.UnsafeIndirect(__obf_2a1474febb02279b)
	__obf_11b7350fb8580af4(__obf_96e65a4c8c4f2ce5.__obf_2f9c5aed866cce75, reflect2.TypeOf(__obf_35136ef2ac0f94e4)).Encode(reflect2.PtrOf(__obf_35136ef2ac0f94e4), __obf_850a7457bb739a32)
}

func (__obf_96e65a4c8c4f2ce5 *__obf_653048ea55761927) IsEmpty(__obf_2a1474febb02279b unsafe.Pointer) bool {
	__obf_35136ef2ac0f94e4 := __obf_96e65a4c8c4f2ce5.__obf_5ea8ee327a6f7e0d.UnsafeIndirect(__obf_2a1474febb02279b)
	return __obf_11b7350fb8580af4(__obf_96e65a4c8c4f2ce5.__obf_2f9c5aed866cce75, reflect2.TypeOf(__obf_35136ef2ac0f94e4)).IsEmpty(reflect2.PtrOf(__obf_35136ef2ac0f94e4))
}

type __obf_ef22fb2a3a24383f struct {
	__obf_08a00a78c23c9481 *reflect2.UnsafeMapType
	__obf_38ad2485e5b56490 ValEncoder
	__obf_11328395cf86586a ValEncoder
}

func (__obf_96e65a4c8c4f2ce5 *__obf_ef22fb2a3a24383f) Encode(__obf_2a1474febb02279b unsafe.Pointer, __obf_850a7457bb739a32 *Stream) {
	if *(*unsafe.Pointer)(__obf_2a1474febb02279b) == nil {
		__obf_850a7457bb739a32.
			WriteNil()
		return
	}
	__obf_850a7457bb739a32.
		WriteObjectStart()
	__obf_1bb30e8a74ed8233 := __obf_96e65a4c8c4f2ce5.__obf_08a00a78c23c9481.UnsafeIterate(__obf_2a1474febb02279b)
	for __obf_5aa5c8829b97f182 := 0; __obf_1bb30e8a74ed8233.HasNext(); __obf_5aa5c8829b97f182++ {
		if __obf_5aa5c8829b97f182 != 0 {
			__obf_850a7457bb739a32.
				WriteMore()
		}
		__obf_fc61abc90a2a7b45, __obf_97d7079fd160f22c := __obf_1bb30e8a74ed8233.UnsafeNext()
		__obf_96e65a4c8c4f2ce5.__obf_38ad2485e5b56490.
			Encode(__obf_fc61abc90a2a7b45, __obf_850a7457bb739a32)
		if __obf_850a7457bb739a32.__obf_2dfa32334fb1aa21 > 0 {
			__obf_850a7457bb739a32.__obf_56747a2b5f26ca3c(byte(':'), byte(' '))
		} else {
			__obf_850a7457bb739a32.__obf_72837f6fe737f078(':')
		}
		__obf_96e65a4c8c4f2ce5.__obf_11328395cf86586a.
			Encode(__obf_97d7079fd160f22c, __obf_850a7457bb739a32)
	}
	__obf_850a7457bb739a32.
		WriteObjectEnd()
}

func (__obf_96e65a4c8c4f2ce5 *__obf_ef22fb2a3a24383f) IsEmpty(__obf_2a1474febb02279b unsafe.Pointer) bool {
	__obf_1bb30e8a74ed8233 := __obf_96e65a4c8c4f2ce5.__obf_08a00a78c23c9481.UnsafeIterate(__obf_2a1474febb02279b)
	return !__obf_1bb30e8a74ed8233.HasNext()
}

type __obf_3c655ef2fb5953e5 struct {
	__obf_08a00a78c23c9481 *reflect2.UnsafeMapType
	__obf_38ad2485e5b56490 ValEncoder
	__obf_11328395cf86586a ValEncoder
}

func (__obf_96e65a4c8c4f2ce5 *__obf_3c655ef2fb5953e5) Encode(__obf_2a1474febb02279b unsafe.Pointer, __obf_850a7457bb739a32 *Stream) {
	if *(*unsafe.Pointer)(__obf_2a1474febb02279b) == nil {
		__obf_850a7457bb739a32.
			WriteNil()
		return
	}
	__obf_850a7457bb739a32.
		WriteObjectStart()
	__obf_1caa4d865a8cf9b3 := __obf_96e65a4c8c4f2ce5.__obf_08a00a78c23c9481.UnsafeIterate(__obf_2a1474febb02279b)
	__obf_6170edf961c8d699 := __obf_850a7457bb739a32.__obf_892632c77e859037.BorrowStream(nil)
	__obf_6170edf961c8d699.
		Attachment = __obf_850a7457bb739a32.Attachment
	__obf_84296603e957b922 := __obf_850a7457bb739a32.__obf_892632c77e859037.BorrowIterator(nil)
	__obf_ffa4a9085b19c905 := __obf_ad9d3cf9389c93d1{}
	for __obf_1caa4d865a8cf9b3.HasNext() {
		__obf_fc61abc90a2a7b45, __obf_97d7079fd160f22c := __obf_1caa4d865a8cf9b3.UnsafeNext()
		__obf_c1331b415fb9908c := __obf_6170edf961c8d699.Buffered()
		__obf_96e65a4c8c4f2ce5.__obf_38ad2485e5b56490.
			Encode(__obf_fc61abc90a2a7b45, __obf_6170edf961c8d699)
		if __obf_6170edf961c8d699.Error != nil && __obf_6170edf961c8d699.Error != io.EOF && __obf_850a7457bb739a32.Error == nil {
			__obf_850a7457bb739a32.
				Error = __obf_6170edf961c8d699.Error
		}
		__obf_44238afd063cfd94 := __obf_6170edf961c8d699.Buffer()[__obf_c1331b415fb9908c:]
		__obf_84296603e957b922.
			ResetBytes(__obf_44238afd063cfd94)
		__obf_d555f030f951bd37 := __obf_84296603e957b922.ReadString()
		if __obf_850a7457bb739a32.__obf_2dfa32334fb1aa21 > 0 {
			__obf_6170edf961c8d699.__obf_56747a2b5f26ca3c(byte(':'), byte(' '))
		} else {
			__obf_6170edf961c8d699.__obf_72837f6fe737f078(':')
		}
		__obf_96e65a4c8c4f2ce5.__obf_11328395cf86586a.
			Encode(__obf_97d7079fd160f22c, __obf_6170edf961c8d699)
		__obf_ffa4a9085b19c905 = append(__obf_ffa4a9085b19c905, __obf_c7f698397bd82d50{__obf_fc61abc90a2a7b45: __obf_d555f030f951bd37, __obf_4d55aa768b40f2fb: __obf_6170edf961c8d699.Buffer()[__obf_c1331b415fb9908c:]})
	}
	sort.Sort(__obf_ffa4a9085b19c905)
	for __obf_5aa5c8829b97f182, __obf_4d55aa768b40f2fb := range __obf_ffa4a9085b19c905 {
		if __obf_5aa5c8829b97f182 != 0 {
			__obf_850a7457bb739a32.
				WriteMore()
		}
		__obf_850a7457bb739a32.
			Write(__obf_4d55aa768b40f2fb.__obf_4d55aa768b40f2fb)
	}
	if __obf_6170edf961c8d699.Error != nil && __obf_850a7457bb739a32.Error == nil {
		__obf_850a7457bb739a32.
			Error = __obf_6170edf961c8d699.Error
	}
	__obf_850a7457bb739a32.
		WriteObjectEnd()
	__obf_850a7457bb739a32.__obf_892632c77e859037.
		ReturnStream(__obf_6170edf961c8d699)
	__obf_850a7457bb739a32.__obf_892632c77e859037.
		ReturnIterator(__obf_84296603e957b922)
}

func (__obf_96e65a4c8c4f2ce5 *__obf_3c655ef2fb5953e5) IsEmpty(__obf_2a1474febb02279b unsafe.Pointer) bool {
	__obf_1bb30e8a74ed8233 := __obf_96e65a4c8c4f2ce5.__obf_08a00a78c23c9481.UnsafeIterate(__obf_2a1474febb02279b)
	return !__obf_1bb30e8a74ed8233.HasNext()
}

type __obf_ad9d3cf9389c93d1 []__obf_c7f698397bd82d50

type __obf_c7f698397bd82d50 struct {
	__obf_fc61abc90a2a7b45 string
	__obf_4d55aa768b40f2fb []byte
}

func (__obf_7fb91ef3074fde14 __obf_ad9d3cf9389c93d1) Len() int { return len(__obf_7fb91ef3074fde14) }
func (__obf_7fb91ef3074fde14 __obf_ad9d3cf9389c93d1) Swap(__obf_5aa5c8829b97f182, __obf_7b29092764c6c9cb int) {
	__obf_7fb91ef3074fde14[__obf_5aa5c8829b97f182], __obf_7fb91ef3074fde14[__obf_7b29092764c6c9cb] = __obf_7fb91ef3074fde14[__obf_7b29092764c6c9cb], __obf_7fb91ef3074fde14[__obf_5aa5c8829b97f182]
}
func (__obf_7fb91ef3074fde14 __obf_ad9d3cf9389c93d1) Less(__obf_5aa5c8829b97f182, __obf_7b29092764c6c9cb int) bool {
	return __obf_7fb91ef3074fde14[__obf_5aa5c8829b97f182].__obf_fc61abc90a2a7b45 < __obf_7fb91ef3074fde14[__obf_7b29092764c6c9cb].__obf_fc61abc90a2a7b45
}
