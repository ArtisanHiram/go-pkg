package __obf_91620b895eeff9ed

import (
	"fmt"
	"reflect"
	"unsafe"

	"github.com/modern-go/reflect2"
)

// ValDecoder is an internal type registered to cache as needed.
// Don't confuse jsoniter.ValDecoder with json.Decoder.
// For json.Decoder's adapter, refer to jsoniter.AdapterDecoder(todo link).
//
// Reflection on type to create decoders, which is then cached
// Reflection on value is avoided as we can, as the reflect.Value itself will allocate, with following exceptions
// 1. create instance of new value, for example *int will need a int to be allocated
// 2. append to slice, if the existing cap is not enough, allocate will be done using Reflect.New
// 3. assignment to map, both key and value will be reflect.Value
// For a simple struct binding, it will be reflect.Value free and allocation free
type ValDecoder interface {
	Decode(__obf_2a1474febb02279b unsafe.Pointer, __obf_1bb30e8a74ed8233 *Iterator)
}

// ValEncoder is an internal type registered to cache as needed.
// Don't confuse jsoniter.ValEncoder with json.Encoder.
// For json.Encoder's adapter, refer to jsoniter.AdapterEncoder(todo godoc link).
type ValEncoder interface {
	IsEmpty(__obf_2a1474febb02279b unsafe.Pointer) bool
	Encode(__obf_2a1474febb02279b unsafe.Pointer, __obf_850a7457bb739a32 *Stream)
}

type __obf_7a8d7d05140d0aaa interface {
	IsEmpty(__obf_2a1474febb02279b unsafe.Pointer) bool
}

type __obf_2f9c5aed866cce75 struct {
	*__obf_972c2293804d141c
	__obf_6e20c91e70903582 string
	__obf_f6b2feb1a20c21df map[reflect2.Type]ValEncoder
	__obf_77e81c6c93a73f92 map[reflect2.Type]ValDecoder
}

func (__obf_1b9dc2dc0f9af749 *__obf_2f9c5aed866cce75) __obf_f087a7a47617f72a() bool {
	if __obf_1b9dc2dc0f9af749.__obf_972c2293804d141c == nil {
		// default is case-insensitive
		return false
	}
	return __obf_1b9dc2dc0f9af749.__obf_972c2293804d141c.__obf_f087a7a47617f72a
}

func (__obf_1b9dc2dc0f9af749 *__obf_2f9c5aed866cce75) append(__obf_6e20c91e70903582 string) *__obf_2f9c5aed866cce75 {
	return &__obf_2f9c5aed866cce75{__obf_972c2293804d141c: __obf_1b9dc2dc0f9af749.__obf_972c2293804d141c, __obf_6e20c91e70903582: __obf_1b9dc2dc0f9af749.__obf_6e20c91e70903582 + " " + __obf_6e20c91e70903582, __obf_f6b2feb1a20c21df: __obf_1b9dc2dc0f9af749.__obf_f6b2feb1a20c21df, __obf_77e81c6c93a73f92: __obf_1b9dc2dc0f9af749.__obf_77e81c6c93a73f92}
}

// ReadVal copy the underlying JSON into go interface, same as json.Unmarshal
func (__obf_1bb30e8a74ed8233 *Iterator) ReadVal(__obf_35136ef2ac0f94e4 any) {
	__obf_1802821f4ca77d0e := __obf_1bb30e8a74ed8233.__obf_1802821f4ca77d0e
	__obf_f5a66ce7c84d060d := reflect2.RTypeOf(__obf_35136ef2ac0f94e4)
	__obf_6fd3f72eb9b5574c := __obf_1bb30e8a74ed8233.__obf_892632c77e859037.__obf_49679f79c9853606(__obf_f5a66ce7c84d060d)
	if __obf_6fd3f72eb9b5574c == nil {
		__obf_29ebd0f2c324f5ea := reflect2.TypeOf(__obf_35136ef2ac0f94e4)
		if __obf_29ebd0f2c324f5ea == nil || __obf_29ebd0f2c324f5ea.Kind() != reflect.Ptr {
			__obf_1bb30e8a74ed8233.
				ReportError("ReadVal", "can only unmarshal into pointer")
			return
		}
		__obf_6fd3f72eb9b5574c = __obf_1bb30e8a74ed8233.__obf_892632c77e859037.DecoderOf(__obf_29ebd0f2c324f5ea)
	}
	__obf_2a1474febb02279b := reflect2.PtrOf(__obf_35136ef2ac0f94e4)
	if __obf_2a1474febb02279b == nil {
		__obf_1bb30e8a74ed8233.
			ReportError("ReadVal", "can not read into nil pointer")
		return
	}
	__obf_6fd3f72eb9b5574c.
		Decode(__obf_2a1474febb02279b, __obf_1bb30e8a74ed8233)
	if __obf_1bb30e8a74ed8233.__obf_1802821f4ca77d0e != __obf_1802821f4ca77d0e {
		__obf_1bb30e8a74ed8233.
			ReportError("ReadVal", "unexpected mismatched nesting")
		return
	}
}

// WriteVal copy the go interface into underlying JSON, same as json.Marshal
func (__obf_850a7457bb739a32 *Stream) WriteVal(__obf_bbfd2ac8618a6f0c any) {
	if nil == __obf_bbfd2ac8618a6f0c {
		__obf_850a7457bb739a32.
			WriteNil()
		return
	}
	__obf_f5a66ce7c84d060d := reflect2.RTypeOf(__obf_bbfd2ac8618a6f0c)
	__obf_96e65a4c8c4f2ce5 := __obf_850a7457bb739a32.__obf_892632c77e859037.__obf_7892b430524f2484(__obf_f5a66ce7c84d060d)
	if __obf_96e65a4c8c4f2ce5 == nil {
		__obf_29ebd0f2c324f5ea := reflect2.TypeOf(__obf_bbfd2ac8618a6f0c)
		__obf_96e65a4c8c4f2ce5 = __obf_850a7457bb739a32.__obf_892632c77e859037.EncoderOf(__obf_29ebd0f2c324f5ea)
	}
	__obf_96e65a4c8c4f2ce5.
		Encode(reflect2.PtrOf(__obf_bbfd2ac8618a6f0c), __obf_850a7457bb739a32)
}

func (__obf_892632c77e859037 *__obf_972c2293804d141c) DecoderOf(__obf_29ebd0f2c324f5ea reflect2.Type) ValDecoder {
	__obf_f5a66ce7c84d060d := __obf_29ebd0f2c324f5ea.RType()
	__obf_6fd3f72eb9b5574c := __obf_892632c77e859037.__obf_49679f79c9853606(__obf_f5a66ce7c84d060d)
	if __obf_6fd3f72eb9b5574c != nil {
		return __obf_6fd3f72eb9b5574c
	}
	__obf_2f9c5aed866cce75 := &__obf_2f9c5aed866cce75{__obf_972c2293804d141c: __obf_892632c77e859037, __obf_6e20c91e70903582: "", __obf_77e81c6c93a73f92: map[reflect2.Type]ValDecoder{}, __obf_f6b2feb1a20c21df: map[reflect2.Type]ValEncoder{}}
	__obf_f2fdafeb141957bd := __obf_29ebd0f2c324f5ea.(*reflect2.UnsafePtrType)
	__obf_6fd3f72eb9b5574c = __obf_0b44a7afc1523314(__obf_2f9c5aed866cce75, __obf_f2fdafeb141957bd.Elem())
	__obf_892632c77e859037.__obf_70fb224d773d428b(__obf_f5a66ce7c84d060d, __obf_6fd3f72eb9b5574c)
	return __obf_6fd3f72eb9b5574c
}

func __obf_0b44a7afc1523314(__obf_2f9c5aed866cce75 *__obf_2f9c5aed866cce75, __obf_29ebd0f2c324f5ea reflect2.Type) ValDecoder {
	__obf_6fd3f72eb9b5574c := __obf_deb887de9e48895a(__obf_2f9c5aed866cce75, __obf_29ebd0f2c324f5ea)
	if __obf_6fd3f72eb9b5574c != nil {
		return __obf_6fd3f72eb9b5574c
	}
	__obf_6fd3f72eb9b5574c = __obf_b64b0813f814034f(__obf_2f9c5aed866cce75, __obf_29ebd0f2c324f5ea)
	for _, __obf_8c9d73a8f8319687 := range __obf_bb7ace2aaa9423b9 {
		__obf_6fd3f72eb9b5574c = __obf_8c9d73a8f8319687.DecorateDecoder(__obf_29ebd0f2c324f5ea, __obf_6fd3f72eb9b5574c)
	}
	__obf_6fd3f72eb9b5574c = __obf_2f9c5aed866cce75.__obf_920e2f1ddf47b5e1.DecorateDecoder(__obf_29ebd0f2c324f5ea, __obf_6fd3f72eb9b5574c)
	for _, __obf_8c9d73a8f8319687 := range __obf_2f9c5aed866cce75.__obf_b4dfae156c7ac26c {
		__obf_6fd3f72eb9b5574c = __obf_8c9d73a8f8319687.DecorateDecoder(__obf_29ebd0f2c324f5ea, __obf_6fd3f72eb9b5574c)
	}
	return __obf_6fd3f72eb9b5574c
}

func __obf_b64b0813f814034f(__obf_2f9c5aed866cce75 *__obf_2f9c5aed866cce75, __obf_29ebd0f2c324f5ea reflect2.Type) ValDecoder {
	__obf_6fd3f72eb9b5574c := __obf_2f9c5aed866cce75.__obf_77e81c6c93a73f92[__obf_29ebd0f2c324f5ea]
	if __obf_6fd3f72eb9b5574c != nil {
		return __obf_6fd3f72eb9b5574c
	}
	__obf_a414ce00908cfb47 := &__obf_adc0a0c5c132c296{}
	__obf_2f9c5aed866cce75.__obf_77e81c6c93a73f92[__obf_29ebd0f2c324f5ea] = __obf_a414ce00908cfb47
	__obf_6fd3f72eb9b5574c = _createDecoderOfType(__obf_2f9c5aed866cce75, __obf_29ebd0f2c324f5ea)
	__obf_a414ce00908cfb47.__obf_6fd3f72eb9b5574c = __obf_6fd3f72eb9b5574c
	return __obf_6fd3f72eb9b5574c
}

func _createDecoderOfType(__obf_2f9c5aed866cce75 *__obf_2f9c5aed866cce75, __obf_29ebd0f2c324f5ea reflect2.Type) ValDecoder {
	__obf_6fd3f72eb9b5574c := __obf_f5b325d98e33bfe7(__obf_29ebd0f2c324f5ea)
	if __obf_6fd3f72eb9b5574c != nil {
		return __obf_6fd3f72eb9b5574c
	}
	__obf_6fd3f72eb9b5574c = __obf_5a852a950afc440b(__obf_2f9c5aed866cce75, __obf_29ebd0f2c324f5ea)
	if __obf_6fd3f72eb9b5574c != nil {
		return __obf_6fd3f72eb9b5574c
	}
	__obf_6fd3f72eb9b5574c = __obf_a0300a1b15f3fca6(__obf_2f9c5aed866cce75, __obf_29ebd0f2c324f5ea)
	if __obf_6fd3f72eb9b5574c != nil {
		return __obf_6fd3f72eb9b5574c
	}
	__obf_6fd3f72eb9b5574c = __obf_9aa737014fad08d9(__obf_2f9c5aed866cce75, __obf_29ebd0f2c324f5ea)
	if __obf_6fd3f72eb9b5574c != nil {
		return __obf_6fd3f72eb9b5574c
	}
	__obf_6fd3f72eb9b5574c = __obf_f1b7f12340262e83(__obf_2f9c5aed866cce75, __obf_29ebd0f2c324f5ea)
	if __obf_6fd3f72eb9b5574c != nil {
		return __obf_6fd3f72eb9b5574c
	}
	switch __obf_29ebd0f2c324f5ea.Kind() {
	case reflect.Interface:
		__obf_4c24859f59053937, __obf_06869ed2256bec63 := __obf_29ebd0f2c324f5ea.(*reflect2.UnsafeIFaceType)
		if __obf_06869ed2256bec63 {
			return &__obf_36b06a4a164de621{__obf_5ea8ee327a6f7e0d: __obf_4c24859f59053937}
		}
		return &__obf_5e2fcb326d03b9c1{}
	case reflect.Struct:
		return __obf_5e7b461732099c43(__obf_2f9c5aed866cce75, __obf_29ebd0f2c324f5ea)
	case reflect.Array:
		return __obf_c1de99114b629ad2(__obf_2f9c5aed866cce75, __obf_29ebd0f2c324f5ea)
	case reflect.Slice:
		return __obf_5f6f4a91bd6a3b9b(__obf_2f9c5aed866cce75, __obf_29ebd0f2c324f5ea)
	case reflect.Map:
		return __obf_04b67c87ffeaab64(__obf_2f9c5aed866cce75, __obf_29ebd0f2c324f5ea)
	case reflect.Ptr:
		return __obf_06a295cf16fb16e3(__obf_2f9c5aed866cce75, __obf_29ebd0f2c324f5ea)
	default:
		return &__obf_b540734d2036f0d8{__obf_62967739bca1d11d: fmt.Errorf("%s%s is unsupported type", __obf_2f9c5aed866cce75.__obf_6e20c91e70903582, __obf_29ebd0f2c324f5ea.String())}
	}
}

func (__obf_892632c77e859037 *__obf_972c2293804d141c) EncoderOf(__obf_29ebd0f2c324f5ea reflect2.Type) ValEncoder {
	__obf_f5a66ce7c84d060d := __obf_29ebd0f2c324f5ea.RType()
	__obf_96e65a4c8c4f2ce5 := __obf_892632c77e859037.__obf_7892b430524f2484(__obf_f5a66ce7c84d060d)
	if __obf_96e65a4c8c4f2ce5 != nil {
		return __obf_96e65a4c8c4f2ce5
	}
	__obf_2f9c5aed866cce75 := &__obf_2f9c5aed866cce75{__obf_972c2293804d141c: __obf_892632c77e859037, __obf_6e20c91e70903582: "", __obf_77e81c6c93a73f92: map[reflect2.Type]ValDecoder{}, __obf_f6b2feb1a20c21df: map[reflect2.Type]ValEncoder{}}
	__obf_96e65a4c8c4f2ce5 = __obf_d1233db7a73cc96c(__obf_2f9c5aed866cce75, __obf_29ebd0f2c324f5ea)
	if __obf_29ebd0f2c324f5ea.LikePtr() {
		__obf_96e65a4c8c4f2ce5 = &__obf_4af16ee9f391fb6f{__obf_96e65a4c8c4f2ce5}
	}
	__obf_892632c77e859037.__obf_56d16fae33267fc7(__obf_f5a66ce7c84d060d, __obf_96e65a4c8c4f2ce5)
	return __obf_96e65a4c8c4f2ce5
}

type __obf_4af16ee9f391fb6f struct {
	__obf_96e65a4c8c4f2ce5 ValEncoder
}

func (__obf_96e65a4c8c4f2ce5 *__obf_4af16ee9f391fb6f) IsEmpty(__obf_2a1474febb02279b unsafe.Pointer) bool {
	return __obf_96e65a4c8c4f2ce5.__obf_96e65a4c8c4f2ce5.IsEmpty(unsafe.Pointer(&__obf_2a1474febb02279b))
}

func (__obf_96e65a4c8c4f2ce5 *__obf_4af16ee9f391fb6f) Encode(__obf_2a1474febb02279b unsafe.Pointer, __obf_850a7457bb739a32 *Stream) {
	__obf_96e65a4c8c4f2ce5.__obf_96e65a4c8c4f2ce5.
		Encode(unsafe.Pointer(&__obf_2a1474febb02279b), __obf_850a7457bb739a32)
}

func __obf_d1233db7a73cc96c(__obf_2f9c5aed866cce75 *__obf_2f9c5aed866cce75, __obf_29ebd0f2c324f5ea reflect2.Type) ValEncoder {
	__obf_96e65a4c8c4f2ce5 := __obf_ee643627fa1cfd77(__obf_2f9c5aed866cce75, __obf_29ebd0f2c324f5ea)
	if __obf_96e65a4c8c4f2ce5 != nil {
		return __obf_96e65a4c8c4f2ce5
	}
	__obf_96e65a4c8c4f2ce5 = __obf_bce6f9eae7da4ad0(__obf_2f9c5aed866cce75, __obf_29ebd0f2c324f5ea)
	for _, __obf_8c9d73a8f8319687 := range __obf_bb7ace2aaa9423b9 {
		__obf_96e65a4c8c4f2ce5 = __obf_8c9d73a8f8319687.DecorateEncoder(__obf_29ebd0f2c324f5ea, __obf_96e65a4c8c4f2ce5)
	}
	__obf_96e65a4c8c4f2ce5 = __obf_2f9c5aed866cce75.__obf_47929f7efe51b371.DecorateEncoder(__obf_29ebd0f2c324f5ea, __obf_96e65a4c8c4f2ce5)
	for _, __obf_8c9d73a8f8319687 := range __obf_2f9c5aed866cce75.__obf_b4dfae156c7ac26c {
		__obf_96e65a4c8c4f2ce5 = __obf_8c9d73a8f8319687.DecorateEncoder(__obf_29ebd0f2c324f5ea, __obf_96e65a4c8c4f2ce5)
	}
	return __obf_96e65a4c8c4f2ce5
}

func __obf_bce6f9eae7da4ad0(__obf_2f9c5aed866cce75 *__obf_2f9c5aed866cce75, __obf_29ebd0f2c324f5ea reflect2.Type) ValEncoder {
	__obf_96e65a4c8c4f2ce5 := __obf_2f9c5aed866cce75.__obf_f6b2feb1a20c21df[__obf_29ebd0f2c324f5ea]
	if __obf_96e65a4c8c4f2ce5 != nil {
		return __obf_96e65a4c8c4f2ce5
	}
	__obf_a414ce00908cfb47 := &__obf_df35623308757bf4{}
	__obf_2f9c5aed866cce75.__obf_f6b2feb1a20c21df[__obf_29ebd0f2c324f5ea] = __obf_a414ce00908cfb47
	__obf_96e65a4c8c4f2ce5 = _createEncoderOfType(__obf_2f9c5aed866cce75, __obf_29ebd0f2c324f5ea)
	__obf_a414ce00908cfb47.__obf_96e65a4c8c4f2ce5 = __obf_96e65a4c8c4f2ce5
	return __obf_96e65a4c8c4f2ce5
}
func _createEncoderOfType(__obf_2f9c5aed866cce75 *__obf_2f9c5aed866cce75, __obf_29ebd0f2c324f5ea reflect2.Type) ValEncoder {
	__obf_96e65a4c8c4f2ce5 := __obf_70759df1070bb0f4(__obf_29ebd0f2c324f5ea)
	if __obf_96e65a4c8c4f2ce5 != nil {
		return __obf_96e65a4c8c4f2ce5
	}
	__obf_96e65a4c8c4f2ce5 = __obf_2d74fdc730e873e9(__obf_2f9c5aed866cce75, __obf_29ebd0f2c324f5ea)
	if __obf_96e65a4c8c4f2ce5 != nil {
		return __obf_96e65a4c8c4f2ce5
	}
	__obf_96e65a4c8c4f2ce5 = __obf_656f25d36a6603f4(__obf_2f9c5aed866cce75, __obf_29ebd0f2c324f5ea)
	if __obf_96e65a4c8c4f2ce5 != nil {
		return __obf_96e65a4c8c4f2ce5
	}
	__obf_96e65a4c8c4f2ce5 = __obf_d3e2baf3cccfde91(__obf_2f9c5aed866cce75, __obf_29ebd0f2c324f5ea)
	if __obf_96e65a4c8c4f2ce5 != nil {
		return __obf_96e65a4c8c4f2ce5
	}
	__obf_96e65a4c8c4f2ce5 = __obf_6e137b9e634e5072(__obf_2f9c5aed866cce75, __obf_29ebd0f2c324f5ea)
	if __obf_96e65a4c8c4f2ce5 != nil {
		return __obf_96e65a4c8c4f2ce5
	}
	__obf_c2f7526bd7f1cf55 := __obf_29ebd0f2c324f5ea.Kind()
	switch __obf_c2f7526bd7f1cf55 {
	case reflect.Interface:
		return &__obf_bcab05fd1729d9fa{__obf_29ebd0f2c324f5ea}
	case reflect.Struct:
		return __obf_64b71f72ded515d8(__obf_2f9c5aed866cce75, __obf_29ebd0f2c324f5ea)
	case reflect.Array:
		return __obf_77d3a743e2d92123(__obf_2f9c5aed866cce75, __obf_29ebd0f2c324f5ea)
	case reflect.Slice:
		return __obf_9df0b29f9ca5de5b(__obf_2f9c5aed866cce75, __obf_29ebd0f2c324f5ea)
	case reflect.Map:
		return __obf_00ff8b3760381bbe(__obf_2f9c5aed866cce75, __obf_29ebd0f2c324f5ea)
	case reflect.Ptr:
		return __obf_0c129db9075e37b1(__obf_2f9c5aed866cce75, __obf_29ebd0f2c324f5ea)
	default:
		return &__obf_01cb1af95bf874f9{__obf_62967739bca1d11d: fmt.Errorf("%s%s is unsupported type", __obf_2f9c5aed866cce75.__obf_6e20c91e70903582, __obf_29ebd0f2c324f5ea.String())}
	}
}

type __obf_b540734d2036f0d8 struct {
	__obf_62967739bca1d11d error
}

func (__obf_6fd3f72eb9b5574c *__obf_b540734d2036f0d8) Decode(__obf_2a1474febb02279b unsafe.Pointer, __obf_1bb30e8a74ed8233 *Iterator) {
	if __obf_1bb30e8a74ed8233.WhatIsNext() != NilValue {
		if __obf_1bb30e8a74ed8233.Error == nil {
			__obf_1bb30e8a74ed8233.
				Error = __obf_6fd3f72eb9b5574c.__obf_62967739bca1d11d
		}
	} else {
		__obf_1bb30e8a74ed8233.
			Skip()
	}
}

type __obf_01cb1af95bf874f9 struct {
	__obf_62967739bca1d11d error
}

func (__obf_96e65a4c8c4f2ce5 *__obf_01cb1af95bf874f9) Encode(__obf_2a1474febb02279b unsafe.Pointer, __obf_850a7457bb739a32 *Stream) {
	if __obf_2a1474febb02279b == nil {
		__obf_850a7457bb739a32.
			WriteNil()
	} else if __obf_850a7457bb739a32.Error == nil {
		__obf_850a7457bb739a32.
			Error = __obf_96e65a4c8c4f2ce5.__obf_62967739bca1d11d
	}
}

func (__obf_96e65a4c8c4f2ce5 *__obf_01cb1af95bf874f9) IsEmpty(__obf_2a1474febb02279b unsafe.Pointer) bool {
	return false
}

type __obf_adc0a0c5c132c296 struct {
	__obf_6fd3f72eb9b5574c ValDecoder
}

func (__obf_6fd3f72eb9b5574c *__obf_adc0a0c5c132c296) Decode(__obf_2a1474febb02279b unsafe.Pointer, __obf_1bb30e8a74ed8233 *Iterator) {
	__obf_6fd3f72eb9b5574c.__obf_6fd3f72eb9b5574c.
		Decode(__obf_2a1474febb02279b, __obf_1bb30e8a74ed8233)
}

type __obf_df35623308757bf4 struct {
	__obf_96e65a4c8c4f2ce5 ValEncoder
}

func (__obf_96e65a4c8c4f2ce5 *__obf_df35623308757bf4) Encode(__obf_2a1474febb02279b unsafe.Pointer, __obf_850a7457bb739a32 *Stream) {
	__obf_96e65a4c8c4f2ce5.__obf_96e65a4c8c4f2ce5.
		Encode(__obf_2a1474febb02279b, __obf_850a7457bb739a32)
}

func (__obf_96e65a4c8c4f2ce5 *__obf_df35623308757bf4) IsEmpty(__obf_2a1474febb02279b unsafe.Pointer) bool {
	return __obf_96e65a4c8c4f2ce5.__obf_96e65a4c8c4f2ce5.IsEmpty(__obf_2a1474febb02279b)
}
