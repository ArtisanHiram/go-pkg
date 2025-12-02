package __obf_c7b79b12b33d8238

import (
	"fmt"
	"io"
	"reflect"
	"sort"
	"unsafe"

	"github.com/modern-go/reflect2"
)

func __obf_a3a1d2b70073655a(__obf_99ec45908bceabdb *__obf_99ec45908bceabdb, __obf_edcded11af6ebc4c reflect2.Type) ValDecoder {
	__obf_5d711cb2eb6457ca := __obf_edcded11af6ebc4c.(*reflect2.UnsafeMapType)
	__obf_022787125e88e30e := __obf_96930e171a6a4fd3(__obf_99ec45908bceabdb.append("[mapKey]"), __obf_5d711cb2eb6457ca.Key())
	__obf_5b1e51b09d209694 := __obf_bb14644cc3f030b3(__obf_99ec45908bceabdb.append("[mapElem]"), __obf_5d711cb2eb6457ca.Elem())
	return &__obf_157a9ea6b5aff297{__obf_5d711cb2eb6457ca: __obf_5d711cb2eb6457ca, __obf_dac100455708c87b: __obf_5d711cb2eb6457ca.Key(), __obf_f5ccfed08a3cf863: __obf_5d711cb2eb6457ca.Elem(), __obf_022787125e88e30e: __obf_022787125e88e30e, __obf_5b1e51b09d209694: __obf_5b1e51b09d209694}
}

func __obf_8be1ac6303cc573e(__obf_99ec45908bceabdb *__obf_99ec45908bceabdb, __obf_edcded11af6ebc4c reflect2.Type) ValEncoder {
	__obf_5d711cb2eb6457ca := __obf_edcded11af6ebc4c.(*reflect2.UnsafeMapType)
	if __obf_99ec45908bceabdb.__obf_8571acf590409abc {
		return &__obf_1489ce624db45bf2{__obf_5d711cb2eb6457ca: __obf_5d711cb2eb6457ca, __obf_b45e0b8dbf92317e: __obf_2ce3fff3039d320a(__obf_99ec45908bceabdb.append("[mapKey]"), __obf_5d711cb2eb6457ca.Key()), __obf_e6894379459b04e8: __obf_85f0e4c352da4678(__obf_99ec45908bceabdb.append("[mapElem]"), __obf_5d711cb2eb6457ca.Elem())}
	}
	return &__obf_a2ea064cccf9f52a{__obf_5d711cb2eb6457ca: __obf_5d711cb2eb6457ca, __obf_b45e0b8dbf92317e: __obf_2ce3fff3039d320a(__obf_99ec45908bceabdb.append("[mapKey]"), __obf_5d711cb2eb6457ca.Key()), __obf_e6894379459b04e8: __obf_85f0e4c352da4678(__obf_99ec45908bceabdb.append("[mapElem]"), __obf_5d711cb2eb6457ca.Elem())}
}

func __obf_96930e171a6a4fd3(__obf_99ec45908bceabdb *__obf_99ec45908bceabdb, __obf_edcded11af6ebc4c reflect2.Type) ValDecoder {
	__obf_801f19702638809c := __obf_99ec45908bceabdb.__obf_1b8392ccffa3bed5.CreateMapKeyDecoder(__obf_edcded11af6ebc4c)
	if __obf_801f19702638809c != nil {
		return __obf_801f19702638809c
	}
	for _, __obf_b8f2f726e65c4d89 := range __obf_99ec45908bceabdb.__obf_3c82f4beb30882eb {
		__obf_801f19702638809c := __obf_b8f2f726e65c4d89.CreateMapKeyDecoder(__obf_edcded11af6ebc4c)
		if __obf_801f19702638809c != nil {
			return __obf_801f19702638809c
		}
	}
	__obf_e2840a6d1d1a664b := reflect2.PtrTo(__obf_edcded11af6ebc4c)
	if __obf_e2840a6d1d1a664b.Implements(__obf_f5c4044aeb40e807) {
		return &__obf_488d5d64f097f8dc{
			&__obf_7829690424eb3cfa{__obf_e0ba73c3f13f4455: __obf_e2840a6d1d1a664b},
		}
	}
	if __obf_edcded11af6ebc4c.Implements(__obf_f5c4044aeb40e807) {
		return &__obf_7829690424eb3cfa{__obf_e0ba73c3f13f4455: __obf_edcded11af6ebc4c}
	}
	if __obf_e2840a6d1d1a664b.Implements(__obf_1233738a1f65c955) {
		return &__obf_488d5d64f097f8dc{
			&__obf_1ed943b6487d22ca{__obf_e0ba73c3f13f4455: __obf_e2840a6d1d1a664b},
		}
	}
	if __obf_edcded11af6ebc4c.Implements(__obf_1233738a1f65c955) {
		return &__obf_1ed943b6487d22ca{__obf_e0ba73c3f13f4455: __obf_edcded11af6ebc4c}
	}

	switch __obf_edcded11af6ebc4c.Kind() {
	case reflect.String:
		return __obf_bb14644cc3f030b3(__obf_99ec45908bceabdb, reflect2.DefaultTypeOfKind(reflect.String))
	case reflect.Bool,
		reflect.Uint8, reflect.Int8,
		reflect.Uint16, reflect.Int16,
		reflect.Uint32, reflect.Int32,
		reflect.Uint64, reflect.Int64,
		reflect.Uint, reflect.Int,
		reflect.Float32, reflect.Float64,
		reflect.Uintptr:
		__obf_edcded11af6ebc4c = reflect2.DefaultTypeOfKind(__obf_edcded11af6ebc4c.Kind())
		return &__obf_6cba33ef86edb78c{__obf_bb14644cc3f030b3(__obf_99ec45908bceabdb, __obf_edcded11af6ebc4c)}
	default:
		return &__obf_b7a21f1836187893{__obf_ea0680f8b461a85b: fmt.Errorf("unsupported map key type: %v", __obf_edcded11af6ebc4c)}
	}
}

func __obf_2ce3fff3039d320a(__obf_99ec45908bceabdb *__obf_99ec45908bceabdb, __obf_edcded11af6ebc4c reflect2.Type) ValEncoder {
	__obf_c091c27480fae550 := __obf_99ec45908bceabdb.__obf_3594368cedb76ac8.CreateMapKeyEncoder(__obf_edcded11af6ebc4c)
	if __obf_c091c27480fae550 != nil {
		return __obf_c091c27480fae550
	}
	for _, __obf_b8f2f726e65c4d89 := range __obf_99ec45908bceabdb.__obf_3c82f4beb30882eb {
		__obf_c091c27480fae550 := __obf_b8f2f726e65c4d89.CreateMapKeyEncoder(__obf_edcded11af6ebc4c)
		if __obf_c091c27480fae550 != nil {
			return __obf_c091c27480fae550
		}
	}

	if __obf_edcded11af6ebc4c.Kind() != reflect.String {
		if __obf_edcded11af6ebc4c == __obf_c3f942f0f1ba7a21 {
			return &__obf_848eb91cb8d800f4{__obf_7ae2c2c658aa1916: __obf_99ec45908bceabdb.EncoderOf(reflect2.TypeOf(""))}
		}
		if __obf_edcded11af6ebc4c.Implements(__obf_c3f942f0f1ba7a21) {
			return &__obf_c2c5222d98675825{__obf_e0ba73c3f13f4455: __obf_edcded11af6ebc4c, __obf_7ae2c2c658aa1916: __obf_99ec45908bceabdb.EncoderOf(reflect2.TypeOf(""))}
		}
	}

	switch __obf_edcded11af6ebc4c.Kind() {
	case reflect.String:
		return __obf_85f0e4c352da4678(__obf_99ec45908bceabdb, reflect2.DefaultTypeOfKind(reflect.String))
	case reflect.Bool,
		reflect.Uint8, reflect.Int8,
		reflect.Uint16, reflect.Int16,
		reflect.Uint32, reflect.Int32,
		reflect.Uint64, reflect.Int64,
		reflect.Uint, reflect.Int,
		reflect.Float32, reflect.Float64,
		reflect.Uintptr:
		__obf_edcded11af6ebc4c = reflect2.DefaultTypeOfKind(__obf_edcded11af6ebc4c.Kind())
		return &__obf_1ba66c7bfbdfe8b6{__obf_85f0e4c352da4678(__obf_99ec45908bceabdb, __obf_edcded11af6ebc4c)}
	default:
		if __obf_edcded11af6ebc4c.Kind() == reflect.Interface {
			return &__obf_bba82adc0aec156d{__obf_99ec45908bceabdb, __obf_edcded11af6ebc4c}
		}
		return &__obf_531b9739373c7313{__obf_ea0680f8b461a85b: fmt.Errorf("unsupported map key type: %v", __obf_edcded11af6ebc4c)}
	}
}

type __obf_157a9ea6b5aff297 struct {
	__obf_5d711cb2eb6457ca *reflect2.UnsafeMapType
	__obf_dac100455708c87b reflect2.Type
	__obf_f5ccfed08a3cf863 reflect2.Type
	__obf_022787125e88e30e ValDecoder
	__obf_5b1e51b09d209694 ValDecoder
}

func (__obf_801f19702638809c *__obf_157a9ea6b5aff297) Decode(__obf_575c04f2b9d91315 unsafe.Pointer, __obf_0da8c843dad13139 *Iterator) {
	__obf_5d711cb2eb6457ca := __obf_801f19702638809c.__obf_5d711cb2eb6457ca
	__obf_12577c03a12f0d2c := __obf_0da8c843dad13139.__obf_2b436fcb1c0ca36d()
	if __obf_12577c03a12f0d2c == 'n' {
		__obf_0da8c843dad13139.__obf_5da8d54e7a1e782c('u', 'l', 'l')
		*(*unsafe.Pointer)(__obf_575c04f2b9d91315) = nil
		__obf_5d711cb2eb6457ca.
			UnsafeSet(__obf_575c04f2b9d91315, __obf_5d711cb2eb6457ca.UnsafeNew())
		return
	}
	if __obf_5d711cb2eb6457ca.UnsafeIsNil(__obf_575c04f2b9d91315) {
		__obf_5d711cb2eb6457ca.
			UnsafeSet(__obf_575c04f2b9d91315, __obf_5d711cb2eb6457ca.UnsafeMakeMap(0))
	}
	if __obf_12577c03a12f0d2c != '{' {
		__obf_0da8c843dad13139.
			ReportError("ReadMapCB", `expect { or n, but found `+string([]byte{__obf_12577c03a12f0d2c}))
		return
	}
	__obf_12577c03a12f0d2c = __obf_0da8c843dad13139.__obf_2b436fcb1c0ca36d()
	if __obf_12577c03a12f0d2c == '}' {
		return
	}
	__obf_0da8c843dad13139.__obf_a675366c80290b83()
	__obf_efcbbccd38302ccb := __obf_801f19702638809c.__obf_dac100455708c87b.UnsafeNew()
	__obf_801f19702638809c.__obf_022787125e88e30e.
		Decode(__obf_efcbbccd38302ccb, __obf_0da8c843dad13139)
	__obf_12577c03a12f0d2c = __obf_0da8c843dad13139.__obf_2b436fcb1c0ca36d()
	if __obf_12577c03a12f0d2c != ':' {
		__obf_0da8c843dad13139.
			ReportError("ReadMapCB", "expect : after object field, but found "+string([]byte{__obf_12577c03a12f0d2c}))
		return
	}
	__obf_43cefbbf0648e3be := __obf_801f19702638809c.__obf_f5ccfed08a3cf863.UnsafeNew()
	__obf_801f19702638809c.__obf_5b1e51b09d209694.
		Decode(__obf_43cefbbf0648e3be, __obf_0da8c843dad13139)
	__obf_801f19702638809c.__obf_5d711cb2eb6457ca.
		UnsafeSetIndex(__obf_575c04f2b9d91315, __obf_efcbbccd38302ccb, __obf_43cefbbf0648e3be)
	for __obf_12577c03a12f0d2c = __obf_0da8c843dad13139.__obf_2b436fcb1c0ca36d(); __obf_12577c03a12f0d2c == ','; __obf_12577c03a12f0d2c = __obf_0da8c843dad13139.__obf_2b436fcb1c0ca36d() {
		__obf_efcbbccd38302ccb := __obf_801f19702638809c.__obf_dac100455708c87b.UnsafeNew()
		__obf_801f19702638809c.__obf_022787125e88e30e.
			Decode(__obf_efcbbccd38302ccb, __obf_0da8c843dad13139)
		__obf_12577c03a12f0d2c = __obf_0da8c843dad13139.__obf_2b436fcb1c0ca36d()
		if __obf_12577c03a12f0d2c != ':' {
			__obf_0da8c843dad13139.
				ReportError("ReadMapCB", "expect : after object field, but found "+string([]byte{__obf_12577c03a12f0d2c}))
			return
		}
		__obf_43cefbbf0648e3be := __obf_801f19702638809c.__obf_f5ccfed08a3cf863.UnsafeNew()
		__obf_801f19702638809c.__obf_5b1e51b09d209694.
			Decode(__obf_43cefbbf0648e3be, __obf_0da8c843dad13139)
		__obf_801f19702638809c.__obf_5d711cb2eb6457ca.
			UnsafeSetIndex(__obf_575c04f2b9d91315, __obf_efcbbccd38302ccb, __obf_43cefbbf0648e3be)
	}
	if __obf_12577c03a12f0d2c != '}' {
		__obf_0da8c843dad13139.
			ReportError("ReadMapCB", `expect }, but found `+string([]byte{__obf_12577c03a12f0d2c}))
	}
}

type __obf_6cba33ef86edb78c struct {
	__obf_801f19702638809c ValDecoder
}

func (__obf_801f19702638809c *__obf_6cba33ef86edb78c) Decode(__obf_575c04f2b9d91315 unsafe.Pointer, __obf_0da8c843dad13139 *Iterator) {
	__obf_12577c03a12f0d2c := __obf_0da8c843dad13139.__obf_2b436fcb1c0ca36d()
	if __obf_12577c03a12f0d2c != '"' {
		__obf_0da8c843dad13139.
			ReportError("ReadMapCB", `expect ", but found `+string([]byte{__obf_12577c03a12f0d2c}))
		return
	}
	__obf_801f19702638809c.__obf_801f19702638809c.
		Decode(__obf_575c04f2b9d91315, __obf_0da8c843dad13139)
	__obf_12577c03a12f0d2c = __obf_0da8c843dad13139.__obf_2b436fcb1c0ca36d()
	if __obf_12577c03a12f0d2c != '"' {
		__obf_0da8c843dad13139.
			ReportError("ReadMapCB", `expect ", but found `+string([]byte{__obf_12577c03a12f0d2c}))
		return
	}
}

type __obf_1ba66c7bfbdfe8b6 struct {
	__obf_c091c27480fae550 ValEncoder
}

func (__obf_c091c27480fae550 *__obf_1ba66c7bfbdfe8b6) Encode(__obf_575c04f2b9d91315 unsafe.Pointer, __obf_d8f50bcfe87d8b47 *Stream) {
	__obf_d8f50bcfe87d8b47.__obf_7340077d41996822('"')
	__obf_c091c27480fae550.__obf_c091c27480fae550.
		Encode(__obf_575c04f2b9d91315, __obf_d8f50bcfe87d8b47)
	__obf_d8f50bcfe87d8b47.__obf_7340077d41996822('"')
}

func (__obf_c091c27480fae550 *__obf_1ba66c7bfbdfe8b6) IsEmpty(__obf_575c04f2b9d91315 unsafe.Pointer) bool {
	return false
}

type __obf_bba82adc0aec156d struct {
	__obf_99ec45908bceabdb *__obf_99ec45908bceabdb
	__obf_e0ba73c3f13f4455 reflect2.Type
}

func (__obf_c091c27480fae550 *__obf_bba82adc0aec156d) Encode(__obf_575c04f2b9d91315 unsafe.Pointer, __obf_d8f50bcfe87d8b47 *Stream) {
	__obf_d6e2df4782353c64 := __obf_c091c27480fae550.__obf_e0ba73c3f13f4455.UnsafeIndirect(__obf_575c04f2b9d91315)
	__obf_2ce3fff3039d320a(__obf_c091c27480fae550.__obf_99ec45908bceabdb, reflect2.TypeOf(__obf_d6e2df4782353c64)).Encode(reflect2.PtrOf(__obf_d6e2df4782353c64), __obf_d8f50bcfe87d8b47)
}

func (__obf_c091c27480fae550 *__obf_bba82adc0aec156d) IsEmpty(__obf_575c04f2b9d91315 unsafe.Pointer) bool {
	__obf_d6e2df4782353c64 := __obf_c091c27480fae550.__obf_e0ba73c3f13f4455.UnsafeIndirect(__obf_575c04f2b9d91315)
	return __obf_2ce3fff3039d320a(__obf_c091c27480fae550.__obf_99ec45908bceabdb, reflect2.TypeOf(__obf_d6e2df4782353c64)).IsEmpty(reflect2.PtrOf(__obf_d6e2df4782353c64))
}

type __obf_a2ea064cccf9f52a struct {
	__obf_5d711cb2eb6457ca *reflect2.UnsafeMapType
	__obf_b45e0b8dbf92317e ValEncoder
	__obf_e6894379459b04e8 ValEncoder
}

func (__obf_c091c27480fae550 *__obf_a2ea064cccf9f52a) Encode(__obf_575c04f2b9d91315 unsafe.Pointer, __obf_d8f50bcfe87d8b47 *Stream) {
	if *(*unsafe.Pointer)(__obf_575c04f2b9d91315) == nil {
		__obf_d8f50bcfe87d8b47.
			WriteNil()
		return
	}
	__obf_d8f50bcfe87d8b47.
		WriteObjectStart()
	__obf_0da8c843dad13139 := __obf_c091c27480fae550.__obf_5d711cb2eb6457ca.UnsafeIterate(__obf_575c04f2b9d91315)
	for __obf_a051269af3a541bb := 0; __obf_0da8c843dad13139.HasNext(); __obf_a051269af3a541bb++ {
		if __obf_a051269af3a541bb != 0 {
			__obf_d8f50bcfe87d8b47.
				WriteMore()
		}
		__obf_efcbbccd38302ccb, __obf_43cefbbf0648e3be := __obf_0da8c843dad13139.UnsafeNext()
		__obf_c091c27480fae550.__obf_b45e0b8dbf92317e.
			Encode(__obf_efcbbccd38302ccb, __obf_d8f50bcfe87d8b47)
		if __obf_d8f50bcfe87d8b47.__obf_4f8e10a1467d205b > 0 {
			__obf_d8f50bcfe87d8b47.__obf_99147f1a9b9d5ffc(byte(':'), byte(' '))
		} else {
			__obf_d8f50bcfe87d8b47.__obf_7340077d41996822(':')
		}
		__obf_c091c27480fae550.__obf_e6894379459b04e8.
			Encode(__obf_43cefbbf0648e3be, __obf_d8f50bcfe87d8b47)
	}
	__obf_d8f50bcfe87d8b47.
		WriteObjectEnd()
}

func (__obf_c091c27480fae550 *__obf_a2ea064cccf9f52a) IsEmpty(__obf_575c04f2b9d91315 unsafe.Pointer) bool {
	__obf_0da8c843dad13139 := __obf_c091c27480fae550.__obf_5d711cb2eb6457ca.UnsafeIterate(__obf_575c04f2b9d91315)
	return !__obf_0da8c843dad13139.HasNext()
}

type __obf_1489ce624db45bf2 struct {
	__obf_5d711cb2eb6457ca *reflect2.UnsafeMapType
	__obf_b45e0b8dbf92317e ValEncoder
	__obf_e6894379459b04e8 ValEncoder
}

func (__obf_c091c27480fae550 *__obf_1489ce624db45bf2) Encode(__obf_575c04f2b9d91315 unsafe.Pointer, __obf_d8f50bcfe87d8b47 *Stream) {
	if *(*unsafe.Pointer)(__obf_575c04f2b9d91315) == nil {
		__obf_d8f50bcfe87d8b47.
			WriteNil()
		return
	}
	__obf_d8f50bcfe87d8b47.
		WriteObjectStart()
	__obf_4123c5f84ffa2f6f := __obf_c091c27480fae550.__obf_5d711cb2eb6457ca.UnsafeIterate(__obf_575c04f2b9d91315)
	__obf_809dbd7faddc40c6 := __obf_d8f50bcfe87d8b47.__obf_c52e0bcfad4b8b71.BorrowStream(nil)
	__obf_809dbd7faddc40c6.
		Attachment = __obf_d8f50bcfe87d8b47.Attachment
	__obf_990259e739e891ed := __obf_d8f50bcfe87d8b47.__obf_c52e0bcfad4b8b71.BorrowIterator(nil)
	__obf_4aa45569acc66879 := __obf_198eabc89da671aa{}
	for __obf_4123c5f84ffa2f6f.HasNext() {
		__obf_efcbbccd38302ccb, __obf_43cefbbf0648e3be := __obf_4123c5f84ffa2f6f.UnsafeNext()
		__obf_11b89918efd3f394 := __obf_809dbd7faddc40c6.Buffered()
		__obf_c091c27480fae550.__obf_b45e0b8dbf92317e.
			Encode(__obf_efcbbccd38302ccb, __obf_809dbd7faddc40c6)
		if __obf_809dbd7faddc40c6.Error != nil && __obf_809dbd7faddc40c6.Error != io.EOF && __obf_d8f50bcfe87d8b47.Error == nil {
			__obf_d8f50bcfe87d8b47.
				Error = __obf_809dbd7faddc40c6.Error
		}
		__obf_bf277d1a443f94a0 := __obf_809dbd7faddc40c6.Buffer()[__obf_11b89918efd3f394:]
		__obf_990259e739e891ed.
			ResetBytes(__obf_bf277d1a443f94a0)
		__obf_e2b45f7e56598800 := __obf_990259e739e891ed.ReadString()
		if __obf_d8f50bcfe87d8b47.__obf_4f8e10a1467d205b > 0 {
			__obf_809dbd7faddc40c6.__obf_99147f1a9b9d5ffc(byte(':'), byte(' '))
		} else {
			__obf_809dbd7faddc40c6.__obf_7340077d41996822(':')
		}
		__obf_c091c27480fae550.__obf_e6894379459b04e8.
			Encode(__obf_43cefbbf0648e3be, __obf_809dbd7faddc40c6)
		__obf_4aa45569acc66879 = append(__obf_4aa45569acc66879, __obf_307e277aac3dd82d{__obf_efcbbccd38302ccb: __obf_e2b45f7e56598800, __obf_254c6296a93c8c3e: __obf_809dbd7faddc40c6.Buffer()[__obf_11b89918efd3f394:]})
	}
	sort.Sort(__obf_4aa45569acc66879)
	for __obf_a051269af3a541bb, __obf_254c6296a93c8c3e := range __obf_4aa45569acc66879 {
		if __obf_a051269af3a541bb != 0 {
			__obf_d8f50bcfe87d8b47.
				WriteMore()
		}
		__obf_d8f50bcfe87d8b47.
			Write(__obf_254c6296a93c8c3e.__obf_254c6296a93c8c3e)
	}
	if __obf_809dbd7faddc40c6.Error != nil && __obf_d8f50bcfe87d8b47.Error == nil {
		__obf_d8f50bcfe87d8b47.
			Error = __obf_809dbd7faddc40c6.Error
	}
	__obf_d8f50bcfe87d8b47.
		WriteObjectEnd()
	__obf_d8f50bcfe87d8b47.__obf_c52e0bcfad4b8b71.
		ReturnStream(__obf_809dbd7faddc40c6)
	__obf_d8f50bcfe87d8b47.__obf_c52e0bcfad4b8b71.
		ReturnIterator(__obf_990259e739e891ed)
}

func (__obf_c091c27480fae550 *__obf_1489ce624db45bf2) IsEmpty(__obf_575c04f2b9d91315 unsafe.Pointer) bool {
	__obf_0da8c843dad13139 := __obf_c091c27480fae550.__obf_5d711cb2eb6457ca.UnsafeIterate(__obf_575c04f2b9d91315)
	return !__obf_0da8c843dad13139.HasNext()
}

type __obf_198eabc89da671aa []__obf_307e277aac3dd82d

type __obf_307e277aac3dd82d struct {
	__obf_efcbbccd38302ccb string
	__obf_254c6296a93c8c3e []byte
}

func (__obf_2ce6adba58f70a89 __obf_198eabc89da671aa) Len() int { return len(__obf_2ce6adba58f70a89) }
func (__obf_2ce6adba58f70a89 __obf_198eabc89da671aa) Swap(__obf_a051269af3a541bb, __obf_aac66f5ab0eab626 int) {
	__obf_2ce6adba58f70a89[__obf_a051269af3a541bb], __obf_2ce6adba58f70a89[__obf_aac66f5ab0eab626] = __obf_2ce6adba58f70a89[__obf_aac66f5ab0eab626], __obf_2ce6adba58f70a89[__obf_a051269af3a541bb]
}
func (__obf_2ce6adba58f70a89 __obf_198eabc89da671aa) Less(__obf_a051269af3a541bb, __obf_aac66f5ab0eab626 int) bool {
	return __obf_2ce6adba58f70a89[__obf_a051269af3a541bb].__obf_efcbbccd38302ccb < __obf_2ce6adba58f70a89[__obf_aac66f5ab0eab626].__obf_efcbbccd38302ccb
}
