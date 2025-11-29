package __obf_5b802ce8d9ba56d6

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
	Decode(__obf_d3c919547bf7e47a unsafe.Pointer, __obf_67008a6a9e5ba828 *Iterator)
}

// ValEncoder is an internal type registered to cache as needed.
// Don't confuse jsoniter.ValEncoder with json.Encoder.
// For json.Encoder's adapter, refer to jsoniter.AdapterEncoder(todo godoc link).
type ValEncoder interface {
	IsEmpty(__obf_d3c919547bf7e47a unsafe.Pointer) bool
	Encode(__obf_d3c919547bf7e47a unsafe.Pointer, __obf_00fc491caa967a74 *Stream)
}

type __obf_9cd9bb2e92b7d8c2 interface {
	IsEmpty(__obf_d3c919547bf7e47a unsafe.Pointer) bool
}

type __obf_08da24be66d0d13c struct {
	*__obf_5d13d7f3bd06c6cf
	__obf_f3dd783c02964acf string
	__obf_80c7056dc65b6b61 map[reflect2.Type]ValEncoder
	__obf_1aedff1a5fb21bbf map[reflect2.Type]ValDecoder
}

func (__obf_1c7157183bc604f5 *__obf_08da24be66d0d13c) __obf_f48e3198f571baa9() bool {
	if __obf_1c7157183bc604f5.__obf_5d13d7f3bd06c6cf == nil {
		// default is case-insensitive
		return false
	}
	return __obf_1c7157183bc604f5.__obf_5d13d7f3bd06c6cf.__obf_f48e3198f571baa9
}

func (__obf_1c7157183bc604f5 *__obf_08da24be66d0d13c) append(__obf_f3dd783c02964acf string) *__obf_08da24be66d0d13c {
	return &__obf_08da24be66d0d13c{__obf_5d13d7f3bd06c6cf: __obf_1c7157183bc604f5.__obf_5d13d7f3bd06c6cf, __obf_f3dd783c02964acf: __obf_1c7157183bc604f5.__obf_f3dd783c02964acf + " " + __obf_f3dd783c02964acf, __obf_80c7056dc65b6b61: __obf_1c7157183bc604f5.__obf_80c7056dc65b6b61, __obf_1aedff1a5fb21bbf: __obf_1c7157183bc604f5.__obf_1aedff1a5fb21bbf}
}

// ReadVal copy the underlying JSON into go interface, same as json.Unmarshal
func (__obf_67008a6a9e5ba828 *Iterator) ReadVal(__obf_f9b1526531f3c024 any) {
	__obf_e0b51eaba31d0671 := __obf_67008a6a9e5ba828.__obf_e0b51eaba31d0671
	__obf_760ddf980dcc31a5 := reflect2.RTypeOf(__obf_f9b1526531f3c024)
	__obf_18f746d7b5b440ee := __obf_67008a6a9e5ba828.__obf_dca106e1cdf85392.__obf_0b82b8c45af29ba7(__obf_760ddf980dcc31a5)
	if __obf_18f746d7b5b440ee == nil {
		__obf_5efc66d8c338c133 := reflect2.TypeOf(__obf_f9b1526531f3c024)
		if __obf_5efc66d8c338c133 == nil || __obf_5efc66d8c338c133.Kind() != reflect.Ptr {
			__obf_67008a6a9e5ba828.
				ReportError("ReadVal", "can only unmarshal into pointer")
			return
		}
		__obf_18f746d7b5b440ee = __obf_67008a6a9e5ba828.__obf_dca106e1cdf85392.DecoderOf(__obf_5efc66d8c338c133)
	}
	__obf_d3c919547bf7e47a := reflect2.PtrOf(__obf_f9b1526531f3c024)
	if __obf_d3c919547bf7e47a == nil {
		__obf_67008a6a9e5ba828.
			ReportError("ReadVal", "can not read into nil pointer")
		return
	}
	__obf_18f746d7b5b440ee.
		Decode(__obf_d3c919547bf7e47a, __obf_67008a6a9e5ba828)
	if __obf_67008a6a9e5ba828.__obf_e0b51eaba31d0671 != __obf_e0b51eaba31d0671 {
		__obf_67008a6a9e5ba828.
			ReportError("ReadVal", "unexpected mismatched nesting")
		return
	}
}

// WriteVal copy the go interface into underlying JSON, same as json.Marshal
func (__obf_00fc491caa967a74 *Stream) WriteVal(__obf_5406b11e33b9d1d3 any) {
	if nil == __obf_5406b11e33b9d1d3 {
		__obf_00fc491caa967a74.
			WriteNil()
		return
	}
	__obf_760ddf980dcc31a5 := reflect2.RTypeOf(__obf_5406b11e33b9d1d3)
	__obf_29366c3ad76a8c51 := __obf_00fc491caa967a74.__obf_dca106e1cdf85392.__obf_2ca4a0a018fe3cfc(__obf_760ddf980dcc31a5)
	if __obf_29366c3ad76a8c51 == nil {
		__obf_5efc66d8c338c133 := reflect2.TypeOf(__obf_5406b11e33b9d1d3)
		__obf_29366c3ad76a8c51 = __obf_00fc491caa967a74.__obf_dca106e1cdf85392.EncoderOf(__obf_5efc66d8c338c133)
	}
	__obf_29366c3ad76a8c51.
		Encode(reflect2.PtrOf(__obf_5406b11e33b9d1d3), __obf_00fc491caa967a74)
}

func (__obf_dca106e1cdf85392 *__obf_5d13d7f3bd06c6cf) DecoderOf(__obf_5efc66d8c338c133 reflect2.Type) ValDecoder {
	__obf_760ddf980dcc31a5 := __obf_5efc66d8c338c133.RType()
	__obf_18f746d7b5b440ee := __obf_dca106e1cdf85392.__obf_0b82b8c45af29ba7(__obf_760ddf980dcc31a5)
	if __obf_18f746d7b5b440ee != nil {
		return __obf_18f746d7b5b440ee
	}
	__obf_08da24be66d0d13c := &__obf_08da24be66d0d13c{__obf_5d13d7f3bd06c6cf: __obf_dca106e1cdf85392, __obf_f3dd783c02964acf: "", __obf_1aedff1a5fb21bbf: map[reflect2.Type]ValDecoder{}, __obf_80c7056dc65b6b61: map[reflect2.Type]ValEncoder{}}
	__obf_d0cac7bfcf0092ea := __obf_5efc66d8c338c133.(*reflect2.UnsafePtrType)
	__obf_18f746d7b5b440ee = __obf_c3a46fc9dd10c84e(__obf_08da24be66d0d13c, __obf_d0cac7bfcf0092ea.Elem())
	__obf_dca106e1cdf85392.__obf_88c5e2d37a5eaf26(__obf_760ddf980dcc31a5, __obf_18f746d7b5b440ee)
	return __obf_18f746d7b5b440ee
}

func __obf_c3a46fc9dd10c84e(__obf_08da24be66d0d13c *__obf_08da24be66d0d13c, __obf_5efc66d8c338c133 reflect2.Type) ValDecoder {
	__obf_18f746d7b5b440ee := __obf_8b38c87279d9c5d7(__obf_08da24be66d0d13c, __obf_5efc66d8c338c133)
	if __obf_18f746d7b5b440ee != nil {
		return __obf_18f746d7b5b440ee
	}
	__obf_18f746d7b5b440ee = __obf_e5cc0c9ea6c53bd5(__obf_08da24be66d0d13c, __obf_5efc66d8c338c133)
	for _, __obf_6b17b29e9b6f5243 := range __obf_d12055ebaa226ca1 {
		__obf_18f746d7b5b440ee = __obf_6b17b29e9b6f5243.DecorateDecoder(__obf_5efc66d8c338c133, __obf_18f746d7b5b440ee)
	}
	__obf_18f746d7b5b440ee = __obf_08da24be66d0d13c.__obf_de1968c22ba7e047.DecorateDecoder(__obf_5efc66d8c338c133, __obf_18f746d7b5b440ee)
	for _, __obf_6b17b29e9b6f5243 := range __obf_08da24be66d0d13c.__obf_251d7593467966e4 {
		__obf_18f746d7b5b440ee = __obf_6b17b29e9b6f5243.DecorateDecoder(__obf_5efc66d8c338c133, __obf_18f746d7b5b440ee)
	}
	return __obf_18f746d7b5b440ee
}

func __obf_e5cc0c9ea6c53bd5(__obf_08da24be66d0d13c *__obf_08da24be66d0d13c, __obf_5efc66d8c338c133 reflect2.Type) ValDecoder {
	__obf_18f746d7b5b440ee := __obf_08da24be66d0d13c.__obf_1aedff1a5fb21bbf[__obf_5efc66d8c338c133]
	if __obf_18f746d7b5b440ee != nil {
		return __obf_18f746d7b5b440ee
	}
	__obf_a732a16069395004 := &__obf_600a982bed70642e{}
	__obf_08da24be66d0d13c.__obf_1aedff1a5fb21bbf[__obf_5efc66d8c338c133] = __obf_a732a16069395004
	__obf_18f746d7b5b440ee = _createDecoderOfType(__obf_08da24be66d0d13c, __obf_5efc66d8c338c133)
	__obf_a732a16069395004.__obf_18f746d7b5b440ee = __obf_18f746d7b5b440ee
	return __obf_18f746d7b5b440ee
}

func _createDecoderOfType(__obf_08da24be66d0d13c *__obf_08da24be66d0d13c, __obf_5efc66d8c338c133 reflect2.Type) ValDecoder {
	__obf_18f746d7b5b440ee := __obf_e31e45c2077f2473(__obf_5efc66d8c338c133)
	if __obf_18f746d7b5b440ee != nil {
		return __obf_18f746d7b5b440ee
	}
	__obf_18f746d7b5b440ee = __obf_3fe7eb60741bd404(__obf_08da24be66d0d13c, __obf_5efc66d8c338c133)
	if __obf_18f746d7b5b440ee != nil {
		return __obf_18f746d7b5b440ee
	}
	__obf_18f746d7b5b440ee = __obf_308321b3ba24eb8c(__obf_08da24be66d0d13c, __obf_5efc66d8c338c133)
	if __obf_18f746d7b5b440ee != nil {
		return __obf_18f746d7b5b440ee
	}
	__obf_18f746d7b5b440ee = __obf_d7f8199541699318(__obf_08da24be66d0d13c, __obf_5efc66d8c338c133)
	if __obf_18f746d7b5b440ee != nil {
		return __obf_18f746d7b5b440ee
	}
	__obf_18f746d7b5b440ee = __obf_145e07780e2db77f(__obf_08da24be66d0d13c, __obf_5efc66d8c338c133)
	if __obf_18f746d7b5b440ee != nil {
		return __obf_18f746d7b5b440ee
	}
	switch __obf_5efc66d8c338c133.Kind() {
	case reflect.Interface:
		__obf_3838254586b2b190, __obf_ba4f3e0033f885b6 := __obf_5efc66d8c338c133.(*reflect2.UnsafeIFaceType)
		if __obf_ba4f3e0033f885b6 {
			return &__obf_acb33c4a1597a4a2{__obf_5e777ac7034ab220: __obf_3838254586b2b190}
		}
		return &__obf_7560826a5f4553cb{}
	case reflect.Struct:
		return __obf_f6a69b5d2ae02314(__obf_08da24be66d0d13c, __obf_5efc66d8c338c133)
	case reflect.Array:
		return __obf_ee5a34160cacf7d6(__obf_08da24be66d0d13c, __obf_5efc66d8c338c133)
	case reflect.Slice:
		return __obf_868f472bae9503f0(__obf_08da24be66d0d13c, __obf_5efc66d8c338c133)
	case reflect.Map:
		return __obf_b13ed867b35030eb(__obf_08da24be66d0d13c, __obf_5efc66d8c338c133)
	case reflect.Ptr:
		return __obf_cb0144ac2d5afb8d(__obf_08da24be66d0d13c, __obf_5efc66d8c338c133)
	default:
		return &__obf_a12de69c81d1c528{__obf_6d071d48f3b321ad: fmt.Errorf("%s%s is unsupported type", __obf_08da24be66d0d13c.__obf_f3dd783c02964acf, __obf_5efc66d8c338c133.String())}
	}
}

func (__obf_dca106e1cdf85392 *__obf_5d13d7f3bd06c6cf) EncoderOf(__obf_5efc66d8c338c133 reflect2.Type) ValEncoder {
	__obf_760ddf980dcc31a5 := __obf_5efc66d8c338c133.RType()
	__obf_29366c3ad76a8c51 := __obf_dca106e1cdf85392.__obf_2ca4a0a018fe3cfc(__obf_760ddf980dcc31a5)
	if __obf_29366c3ad76a8c51 != nil {
		return __obf_29366c3ad76a8c51
	}
	__obf_08da24be66d0d13c := &__obf_08da24be66d0d13c{__obf_5d13d7f3bd06c6cf: __obf_dca106e1cdf85392, __obf_f3dd783c02964acf: "", __obf_1aedff1a5fb21bbf: map[reflect2.Type]ValDecoder{}, __obf_80c7056dc65b6b61: map[reflect2.Type]ValEncoder{}}
	__obf_29366c3ad76a8c51 = __obf_3bb380f7abc03208(__obf_08da24be66d0d13c, __obf_5efc66d8c338c133)
	if __obf_5efc66d8c338c133.LikePtr() {
		__obf_29366c3ad76a8c51 = &__obf_351ba5a1e19cebc1{__obf_29366c3ad76a8c51}
	}
	__obf_dca106e1cdf85392.__obf_7662f98a8ed819d0(__obf_760ddf980dcc31a5, __obf_29366c3ad76a8c51)
	return __obf_29366c3ad76a8c51
}

type __obf_351ba5a1e19cebc1 struct {
	__obf_29366c3ad76a8c51 ValEncoder
}

func (__obf_29366c3ad76a8c51 *__obf_351ba5a1e19cebc1) IsEmpty(__obf_d3c919547bf7e47a unsafe.Pointer) bool {
	return __obf_29366c3ad76a8c51.__obf_29366c3ad76a8c51.IsEmpty(unsafe.Pointer(&__obf_d3c919547bf7e47a))
}

func (__obf_29366c3ad76a8c51 *__obf_351ba5a1e19cebc1) Encode(__obf_d3c919547bf7e47a unsafe.Pointer, __obf_00fc491caa967a74 *Stream) {
	__obf_29366c3ad76a8c51.__obf_29366c3ad76a8c51.
		Encode(unsafe.Pointer(&__obf_d3c919547bf7e47a), __obf_00fc491caa967a74)
}

func __obf_3bb380f7abc03208(__obf_08da24be66d0d13c *__obf_08da24be66d0d13c, __obf_5efc66d8c338c133 reflect2.Type) ValEncoder {
	__obf_29366c3ad76a8c51 := __obf_4563eaa1c4c5e82b(__obf_08da24be66d0d13c, __obf_5efc66d8c338c133)
	if __obf_29366c3ad76a8c51 != nil {
		return __obf_29366c3ad76a8c51
	}
	__obf_29366c3ad76a8c51 = __obf_b7bc4a82851e2c07(__obf_08da24be66d0d13c, __obf_5efc66d8c338c133)
	for _, __obf_6b17b29e9b6f5243 := range __obf_d12055ebaa226ca1 {
		__obf_29366c3ad76a8c51 = __obf_6b17b29e9b6f5243.DecorateEncoder(__obf_5efc66d8c338c133, __obf_29366c3ad76a8c51)
	}
	__obf_29366c3ad76a8c51 = __obf_08da24be66d0d13c.__obf_f754da59a9c09bdc.DecorateEncoder(__obf_5efc66d8c338c133, __obf_29366c3ad76a8c51)
	for _, __obf_6b17b29e9b6f5243 := range __obf_08da24be66d0d13c.__obf_251d7593467966e4 {
		__obf_29366c3ad76a8c51 = __obf_6b17b29e9b6f5243.DecorateEncoder(__obf_5efc66d8c338c133, __obf_29366c3ad76a8c51)
	}
	return __obf_29366c3ad76a8c51
}

func __obf_b7bc4a82851e2c07(__obf_08da24be66d0d13c *__obf_08da24be66d0d13c, __obf_5efc66d8c338c133 reflect2.Type) ValEncoder {
	__obf_29366c3ad76a8c51 := __obf_08da24be66d0d13c.__obf_80c7056dc65b6b61[__obf_5efc66d8c338c133]
	if __obf_29366c3ad76a8c51 != nil {
		return __obf_29366c3ad76a8c51
	}
	__obf_a732a16069395004 := &__obf_f57780403ff790e6{}
	__obf_08da24be66d0d13c.__obf_80c7056dc65b6b61[__obf_5efc66d8c338c133] = __obf_a732a16069395004
	__obf_29366c3ad76a8c51 = _createEncoderOfType(__obf_08da24be66d0d13c, __obf_5efc66d8c338c133)
	__obf_a732a16069395004.__obf_29366c3ad76a8c51 = __obf_29366c3ad76a8c51
	return __obf_29366c3ad76a8c51
}
func _createEncoderOfType(__obf_08da24be66d0d13c *__obf_08da24be66d0d13c, __obf_5efc66d8c338c133 reflect2.Type) ValEncoder {
	__obf_29366c3ad76a8c51 := __obf_248bb80e8824b081(__obf_5efc66d8c338c133)
	if __obf_29366c3ad76a8c51 != nil {
		return __obf_29366c3ad76a8c51
	}
	__obf_29366c3ad76a8c51 = __obf_4e38aa1ee94cd7cf(__obf_08da24be66d0d13c, __obf_5efc66d8c338c133)
	if __obf_29366c3ad76a8c51 != nil {
		return __obf_29366c3ad76a8c51
	}
	__obf_29366c3ad76a8c51 = __obf_238d567082abf8bc(__obf_08da24be66d0d13c, __obf_5efc66d8c338c133)
	if __obf_29366c3ad76a8c51 != nil {
		return __obf_29366c3ad76a8c51
	}
	__obf_29366c3ad76a8c51 = __obf_3f8d707915d65625(__obf_08da24be66d0d13c, __obf_5efc66d8c338c133)
	if __obf_29366c3ad76a8c51 != nil {
		return __obf_29366c3ad76a8c51
	}
	__obf_29366c3ad76a8c51 = __obf_062c5f9f03d60bc9(__obf_08da24be66d0d13c, __obf_5efc66d8c338c133)
	if __obf_29366c3ad76a8c51 != nil {
		return __obf_29366c3ad76a8c51
	}
	__obf_f1764122678433dc := __obf_5efc66d8c338c133.Kind()
	switch __obf_f1764122678433dc {
	case reflect.Interface:
		return &__obf_cb5eb67b8b3d51d0{__obf_5efc66d8c338c133}
	case reflect.Struct:
		return __obf_04c6c702681f689e(__obf_08da24be66d0d13c, __obf_5efc66d8c338c133)
	case reflect.Array:
		return __obf_d2984c9ef31dac9d(__obf_08da24be66d0d13c, __obf_5efc66d8c338c133)
	case reflect.Slice:
		return __obf_0dab538fd3e96275(__obf_08da24be66d0d13c, __obf_5efc66d8c338c133)
	case reflect.Map:
		return __obf_8fcbff4be3a3cf2e(__obf_08da24be66d0d13c, __obf_5efc66d8c338c133)
	case reflect.Ptr:
		return __obf_7894a79375435eef(__obf_08da24be66d0d13c, __obf_5efc66d8c338c133)
	default:
		return &__obf_5d9ef7d0fc378907{__obf_6d071d48f3b321ad: fmt.Errorf("%s%s is unsupported type", __obf_08da24be66d0d13c.__obf_f3dd783c02964acf, __obf_5efc66d8c338c133.String())}
	}
}

type __obf_a12de69c81d1c528 struct {
	__obf_6d071d48f3b321ad error
}

func (__obf_18f746d7b5b440ee *__obf_a12de69c81d1c528) Decode(__obf_d3c919547bf7e47a unsafe.Pointer, __obf_67008a6a9e5ba828 *Iterator) {
	if __obf_67008a6a9e5ba828.WhatIsNext() != NilValue {
		if __obf_67008a6a9e5ba828.Error == nil {
			__obf_67008a6a9e5ba828.
				Error = __obf_18f746d7b5b440ee.__obf_6d071d48f3b321ad
		}
	} else {
		__obf_67008a6a9e5ba828.
			Skip()
	}
}

type __obf_5d9ef7d0fc378907 struct {
	__obf_6d071d48f3b321ad error
}

func (__obf_29366c3ad76a8c51 *__obf_5d9ef7d0fc378907) Encode(__obf_d3c919547bf7e47a unsafe.Pointer, __obf_00fc491caa967a74 *Stream) {
	if __obf_d3c919547bf7e47a == nil {
		__obf_00fc491caa967a74.
			WriteNil()
	} else if __obf_00fc491caa967a74.Error == nil {
		__obf_00fc491caa967a74.
			Error = __obf_29366c3ad76a8c51.__obf_6d071d48f3b321ad
	}
}

func (__obf_29366c3ad76a8c51 *__obf_5d9ef7d0fc378907) IsEmpty(__obf_d3c919547bf7e47a unsafe.Pointer) bool {
	return false
}

type __obf_600a982bed70642e struct {
	__obf_18f746d7b5b440ee ValDecoder
}

func (__obf_18f746d7b5b440ee *__obf_600a982bed70642e) Decode(__obf_d3c919547bf7e47a unsafe.Pointer, __obf_67008a6a9e5ba828 *Iterator) {
	__obf_18f746d7b5b440ee.__obf_18f746d7b5b440ee.
		Decode(__obf_d3c919547bf7e47a, __obf_67008a6a9e5ba828)
}

type __obf_f57780403ff790e6 struct {
	__obf_29366c3ad76a8c51 ValEncoder
}

func (__obf_29366c3ad76a8c51 *__obf_f57780403ff790e6) Encode(__obf_d3c919547bf7e47a unsafe.Pointer, __obf_00fc491caa967a74 *Stream) {
	__obf_29366c3ad76a8c51.__obf_29366c3ad76a8c51.
		Encode(__obf_d3c919547bf7e47a, __obf_00fc491caa967a74)
}

func (__obf_29366c3ad76a8c51 *__obf_f57780403ff790e6) IsEmpty(__obf_d3c919547bf7e47a unsafe.Pointer) bool {
	return __obf_29366c3ad76a8c51.__obf_29366c3ad76a8c51.IsEmpty(__obf_d3c919547bf7e47a)
}
