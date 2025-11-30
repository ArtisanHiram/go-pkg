package __obf_c3cd04a15c56f32f

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
	Decode(__obf_753ab3fb72cdd659 unsafe.Pointer, __obf_edd9fe48d39445e4 *Iterator)
}

// ValEncoder is an internal type registered to cache as needed.
// Don't confuse jsoniter.ValEncoder with json.Encoder.
// For json.Encoder's adapter, refer to jsoniter.AdapterEncoder(todo godoc link).
type ValEncoder interface {
	IsEmpty(__obf_753ab3fb72cdd659 unsafe.Pointer) bool
	Encode(__obf_753ab3fb72cdd659 unsafe.Pointer, __obf_2361f5214e84c60f *Stream)
}

type __obf_7cb1bad9ea4a8234 interface {
	IsEmpty(__obf_753ab3fb72cdd659 unsafe.Pointer) bool
}

type __obf_62435d89ebefd5aa struct {
	*__obf_3a25fb4c9a8342bb
	__obf_f517983aa5897f07 string
	__obf_cdd1406522f9a3ab map[reflect2.Type]ValEncoder
	__obf_8237dc8cee160fdb map[reflect2.Type]ValDecoder
}

func (__obf_902d7026e8a09dd2 *__obf_62435d89ebefd5aa) __obf_b600271a247f99b8() bool {
	if __obf_902d7026e8a09dd2.__obf_3a25fb4c9a8342bb == nil {
		// default is case-insensitive
		return false
	}
	return __obf_902d7026e8a09dd2.__obf_3a25fb4c9a8342bb.__obf_b600271a247f99b8
}

func (__obf_902d7026e8a09dd2 *__obf_62435d89ebefd5aa) append(__obf_f517983aa5897f07 string) *__obf_62435d89ebefd5aa {
	return &__obf_62435d89ebefd5aa{__obf_3a25fb4c9a8342bb: __obf_902d7026e8a09dd2.__obf_3a25fb4c9a8342bb, __obf_f517983aa5897f07: __obf_902d7026e8a09dd2.__obf_f517983aa5897f07 + " " + __obf_f517983aa5897f07, __obf_cdd1406522f9a3ab: __obf_902d7026e8a09dd2.__obf_cdd1406522f9a3ab, __obf_8237dc8cee160fdb: __obf_902d7026e8a09dd2.__obf_8237dc8cee160fdb}
}

// ReadVal copy the underlying JSON into go interface, same as json.Unmarshal
func (__obf_edd9fe48d39445e4 *Iterator) ReadVal(__obf_50acae8c0a871ba1 any) {
	__obf_2da3205c67480a39 := __obf_edd9fe48d39445e4.__obf_2da3205c67480a39
	__obf_cb552b6dac1d7edc := reflect2.RTypeOf(__obf_50acae8c0a871ba1)
	__obf_924cc7ef585cdfb0 := __obf_edd9fe48d39445e4.__obf_0c8065c219ccf0ab.__obf_70f630ceb9086844(__obf_cb552b6dac1d7edc)
	if __obf_924cc7ef585cdfb0 == nil {
		__obf_8667d4fc2a95b0d7 := reflect2.TypeOf(__obf_50acae8c0a871ba1)
		if __obf_8667d4fc2a95b0d7 == nil || __obf_8667d4fc2a95b0d7.Kind() != reflect.Ptr {
			__obf_edd9fe48d39445e4.
				ReportError("ReadVal", "can only unmarshal into pointer")
			return
		}
		__obf_924cc7ef585cdfb0 = __obf_edd9fe48d39445e4.__obf_0c8065c219ccf0ab.DecoderOf(__obf_8667d4fc2a95b0d7)
	}
	__obf_753ab3fb72cdd659 := reflect2.PtrOf(__obf_50acae8c0a871ba1)
	if __obf_753ab3fb72cdd659 == nil {
		__obf_edd9fe48d39445e4.
			ReportError("ReadVal", "can not read into nil pointer")
		return
	}
	__obf_924cc7ef585cdfb0.
		Decode(__obf_753ab3fb72cdd659, __obf_edd9fe48d39445e4)
	if __obf_edd9fe48d39445e4.__obf_2da3205c67480a39 != __obf_2da3205c67480a39 {
		__obf_edd9fe48d39445e4.
			ReportError("ReadVal", "unexpected mismatched nesting")
		return
	}
}

// WriteVal copy the go interface into underlying JSON, same as json.Marshal
func (__obf_2361f5214e84c60f *Stream) WriteVal(__obf_d59813f23ad740a8 any) {
	if nil == __obf_d59813f23ad740a8 {
		__obf_2361f5214e84c60f.
			WriteNil()
		return
	}
	__obf_cb552b6dac1d7edc := reflect2.RTypeOf(__obf_d59813f23ad740a8)
	__obf_c85b895560db528f := __obf_2361f5214e84c60f.__obf_0c8065c219ccf0ab.__obf_ca7fd3d0f377cbfd(__obf_cb552b6dac1d7edc)
	if __obf_c85b895560db528f == nil {
		__obf_8667d4fc2a95b0d7 := reflect2.TypeOf(__obf_d59813f23ad740a8)
		__obf_c85b895560db528f = __obf_2361f5214e84c60f.__obf_0c8065c219ccf0ab.EncoderOf(__obf_8667d4fc2a95b0d7)
	}
	__obf_c85b895560db528f.
		Encode(reflect2.PtrOf(__obf_d59813f23ad740a8), __obf_2361f5214e84c60f)
}

func (__obf_0c8065c219ccf0ab *__obf_3a25fb4c9a8342bb) DecoderOf(__obf_8667d4fc2a95b0d7 reflect2.Type) ValDecoder {
	__obf_cb552b6dac1d7edc := __obf_8667d4fc2a95b0d7.RType()
	__obf_924cc7ef585cdfb0 := __obf_0c8065c219ccf0ab.__obf_70f630ceb9086844(__obf_cb552b6dac1d7edc)
	if __obf_924cc7ef585cdfb0 != nil {
		return __obf_924cc7ef585cdfb0
	}
	__obf_62435d89ebefd5aa := &__obf_62435d89ebefd5aa{__obf_3a25fb4c9a8342bb: __obf_0c8065c219ccf0ab, __obf_f517983aa5897f07: "", __obf_8237dc8cee160fdb: map[reflect2.Type]ValDecoder{}, __obf_cdd1406522f9a3ab: map[reflect2.Type]ValEncoder{}}
	__obf_9fd3068ebc25d7f9 := __obf_8667d4fc2a95b0d7.(*reflect2.UnsafePtrType)
	__obf_924cc7ef585cdfb0 = __obf_eddb00a5736b0fe7(__obf_62435d89ebefd5aa, __obf_9fd3068ebc25d7f9.Elem())
	__obf_0c8065c219ccf0ab.__obf_eb4fd71a722727bf(__obf_cb552b6dac1d7edc, __obf_924cc7ef585cdfb0)
	return __obf_924cc7ef585cdfb0
}

func __obf_eddb00a5736b0fe7(__obf_62435d89ebefd5aa *__obf_62435d89ebefd5aa, __obf_8667d4fc2a95b0d7 reflect2.Type) ValDecoder {
	__obf_924cc7ef585cdfb0 := __obf_e17667896a54b4c5(__obf_62435d89ebefd5aa, __obf_8667d4fc2a95b0d7)
	if __obf_924cc7ef585cdfb0 != nil {
		return __obf_924cc7ef585cdfb0
	}
	__obf_924cc7ef585cdfb0 = __obf_a2f920895161517a(__obf_62435d89ebefd5aa, __obf_8667d4fc2a95b0d7)
	for _, __obf_e4a614f491c1bb0c := range __obf_572808a33f9754ff {
		__obf_924cc7ef585cdfb0 = __obf_e4a614f491c1bb0c.DecorateDecoder(__obf_8667d4fc2a95b0d7, __obf_924cc7ef585cdfb0)
	}
	__obf_924cc7ef585cdfb0 = __obf_62435d89ebefd5aa.__obf_d86968bf26bed402.DecorateDecoder(__obf_8667d4fc2a95b0d7, __obf_924cc7ef585cdfb0)
	for _, __obf_e4a614f491c1bb0c := range __obf_62435d89ebefd5aa.__obf_a1eac24bcc56f0a6 {
		__obf_924cc7ef585cdfb0 = __obf_e4a614f491c1bb0c.DecorateDecoder(__obf_8667d4fc2a95b0d7, __obf_924cc7ef585cdfb0)
	}
	return __obf_924cc7ef585cdfb0
}

func __obf_a2f920895161517a(__obf_62435d89ebefd5aa *__obf_62435d89ebefd5aa, __obf_8667d4fc2a95b0d7 reflect2.Type) ValDecoder {
	__obf_924cc7ef585cdfb0 := __obf_62435d89ebefd5aa.__obf_8237dc8cee160fdb[__obf_8667d4fc2a95b0d7]
	if __obf_924cc7ef585cdfb0 != nil {
		return __obf_924cc7ef585cdfb0
	}
	__obf_bf791dcb299dea37 := &__obf_44f231ad268e82e4{}
	__obf_62435d89ebefd5aa.__obf_8237dc8cee160fdb[__obf_8667d4fc2a95b0d7] = __obf_bf791dcb299dea37
	__obf_924cc7ef585cdfb0 = _createDecoderOfType(__obf_62435d89ebefd5aa, __obf_8667d4fc2a95b0d7)
	__obf_bf791dcb299dea37.__obf_924cc7ef585cdfb0 = __obf_924cc7ef585cdfb0
	return __obf_924cc7ef585cdfb0
}

func _createDecoderOfType(__obf_62435d89ebefd5aa *__obf_62435d89ebefd5aa, __obf_8667d4fc2a95b0d7 reflect2.Type) ValDecoder {
	__obf_924cc7ef585cdfb0 := __obf_93af885996aef0fb(__obf_8667d4fc2a95b0d7)
	if __obf_924cc7ef585cdfb0 != nil {
		return __obf_924cc7ef585cdfb0
	}
	__obf_924cc7ef585cdfb0 = __obf_2ffedf0111a1f93e(__obf_62435d89ebefd5aa, __obf_8667d4fc2a95b0d7)
	if __obf_924cc7ef585cdfb0 != nil {
		return __obf_924cc7ef585cdfb0
	}
	__obf_924cc7ef585cdfb0 = __obf_2b6c587dc5a76788(__obf_62435d89ebefd5aa, __obf_8667d4fc2a95b0d7)
	if __obf_924cc7ef585cdfb0 != nil {
		return __obf_924cc7ef585cdfb0
	}
	__obf_924cc7ef585cdfb0 = __obf_79cadeaf3aef9cd1(__obf_62435d89ebefd5aa, __obf_8667d4fc2a95b0d7)
	if __obf_924cc7ef585cdfb0 != nil {
		return __obf_924cc7ef585cdfb0
	}
	__obf_924cc7ef585cdfb0 = __obf_985b053eae9c9342(__obf_62435d89ebefd5aa, __obf_8667d4fc2a95b0d7)
	if __obf_924cc7ef585cdfb0 != nil {
		return __obf_924cc7ef585cdfb0
	}
	switch __obf_8667d4fc2a95b0d7.Kind() {
	case reflect.Interface:
		__obf_bf9c464b0d370e81, __obf_5543542a87e64fe1 := __obf_8667d4fc2a95b0d7.(*reflect2.UnsafeIFaceType)
		if __obf_5543542a87e64fe1 {
			return &__obf_59cd180ea7ef10b5{__obf_797d4fe23b3556f8: __obf_bf9c464b0d370e81}
		}
		return &__obf_a3ca22e5223e9e8a{}
	case reflect.Struct:
		return __obf_a2a209dd7fe7e0d3(__obf_62435d89ebefd5aa, __obf_8667d4fc2a95b0d7)
	case reflect.Array:
		return __obf_2281a40bdf22829a(__obf_62435d89ebefd5aa, __obf_8667d4fc2a95b0d7)
	case reflect.Slice:
		return __obf_f7de3e7720df8df1(__obf_62435d89ebefd5aa, __obf_8667d4fc2a95b0d7)
	case reflect.Map:
		return __obf_96d7d01ea15ea41d(__obf_62435d89ebefd5aa, __obf_8667d4fc2a95b0d7)
	case reflect.Ptr:
		return __obf_272a3d13803d4851(__obf_62435d89ebefd5aa, __obf_8667d4fc2a95b0d7)
	default:
		return &__obf_5ba850d645a49aec{__obf_5966ecc5edb9b17e: fmt.Errorf("%s%s is unsupported type", __obf_62435d89ebefd5aa.__obf_f517983aa5897f07, __obf_8667d4fc2a95b0d7.String())}
	}
}

func (__obf_0c8065c219ccf0ab *__obf_3a25fb4c9a8342bb) EncoderOf(__obf_8667d4fc2a95b0d7 reflect2.Type) ValEncoder {
	__obf_cb552b6dac1d7edc := __obf_8667d4fc2a95b0d7.RType()
	__obf_c85b895560db528f := __obf_0c8065c219ccf0ab.__obf_ca7fd3d0f377cbfd(__obf_cb552b6dac1d7edc)
	if __obf_c85b895560db528f != nil {
		return __obf_c85b895560db528f
	}
	__obf_62435d89ebefd5aa := &__obf_62435d89ebefd5aa{__obf_3a25fb4c9a8342bb: __obf_0c8065c219ccf0ab, __obf_f517983aa5897f07: "", __obf_8237dc8cee160fdb: map[reflect2.Type]ValDecoder{}, __obf_cdd1406522f9a3ab: map[reflect2.Type]ValEncoder{}}
	__obf_c85b895560db528f = __obf_dd4ab965a9fde81c(__obf_62435d89ebefd5aa, __obf_8667d4fc2a95b0d7)
	if __obf_8667d4fc2a95b0d7.LikePtr() {
		__obf_c85b895560db528f = &__obf_e9fdf8807fd4de0f{__obf_c85b895560db528f}
	}
	__obf_0c8065c219ccf0ab.__obf_7ac7db4cbd198195(__obf_cb552b6dac1d7edc, __obf_c85b895560db528f)
	return __obf_c85b895560db528f
}

type __obf_e9fdf8807fd4de0f struct {
	__obf_c85b895560db528f ValEncoder
}

func (__obf_c85b895560db528f *__obf_e9fdf8807fd4de0f) IsEmpty(__obf_753ab3fb72cdd659 unsafe.Pointer) bool {
	return __obf_c85b895560db528f.__obf_c85b895560db528f.IsEmpty(unsafe.Pointer(&__obf_753ab3fb72cdd659))
}

func (__obf_c85b895560db528f *__obf_e9fdf8807fd4de0f) Encode(__obf_753ab3fb72cdd659 unsafe.Pointer, __obf_2361f5214e84c60f *Stream) {
	__obf_c85b895560db528f.__obf_c85b895560db528f.
		Encode(unsafe.Pointer(&__obf_753ab3fb72cdd659), __obf_2361f5214e84c60f)
}

func __obf_dd4ab965a9fde81c(__obf_62435d89ebefd5aa *__obf_62435d89ebefd5aa, __obf_8667d4fc2a95b0d7 reflect2.Type) ValEncoder {
	__obf_c85b895560db528f := __obf_31b5d333a775e58a(__obf_62435d89ebefd5aa, __obf_8667d4fc2a95b0d7)
	if __obf_c85b895560db528f != nil {
		return __obf_c85b895560db528f
	}
	__obf_c85b895560db528f = __obf_b5904c6dc99194f8(__obf_62435d89ebefd5aa, __obf_8667d4fc2a95b0d7)
	for _, __obf_e4a614f491c1bb0c := range __obf_572808a33f9754ff {
		__obf_c85b895560db528f = __obf_e4a614f491c1bb0c.DecorateEncoder(__obf_8667d4fc2a95b0d7, __obf_c85b895560db528f)
	}
	__obf_c85b895560db528f = __obf_62435d89ebefd5aa.__obf_912608e689e59f9b.DecorateEncoder(__obf_8667d4fc2a95b0d7, __obf_c85b895560db528f)
	for _, __obf_e4a614f491c1bb0c := range __obf_62435d89ebefd5aa.__obf_a1eac24bcc56f0a6 {
		__obf_c85b895560db528f = __obf_e4a614f491c1bb0c.DecorateEncoder(__obf_8667d4fc2a95b0d7, __obf_c85b895560db528f)
	}
	return __obf_c85b895560db528f
}

func __obf_b5904c6dc99194f8(__obf_62435d89ebefd5aa *__obf_62435d89ebefd5aa, __obf_8667d4fc2a95b0d7 reflect2.Type) ValEncoder {
	__obf_c85b895560db528f := __obf_62435d89ebefd5aa.__obf_cdd1406522f9a3ab[__obf_8667d4fc2a95b0d7]
	if __obf_c85b895560db528f != nil {
		return __obf_c85b895560db528f
	}
	__obf_bf791dcb299dea37 := &__obf_16e24315cd385130{}
	__obf_62435d89ebefd5aa.__obf_cdd1406522f9a3ab[__obf_8667d4fc2a95b0d7] = __obf_bf791dcb299dea37
	__obf_c85b895560db528f = _createEncoderOfType(__obf_62435d89ebefd5aa, __obf_8667d4fc2a95b0d7)
	__obf_bf791dcb299dea37.__obf_c85b895560db528f = __obf_c85b895560db528f
	return __obf_c85b895560db528f
}
func _createEncoderOfType(__obf_62435d89ebefd5aa *__obf_62435d89ebefd5aa, __obf_8667d4fc2a95b0d7 reflect2.Type) ValEncoder {
	__obf_c85b895560db528f := __obf_5939b4177daf663d(__obf_8667d4fc2a95b0d7)
	if __obf_c85b895560db528f != nil {
		return __obf_c85b895560db528f
	}
	__obf_c85b895560db528f = __obf_28c9f319ac0c0137(__obf_62435d89ebefd5aa, __obf_8667d4fc2a95b0d7)
	if __obf_c85b895560db528f != nil {
		return __obf_c85b895560db528f
	}
	__obf_c85b895560db528f = __obf_34128745798a55d1(__obf_62435d89ebefd5aa, __obf_8667d4fc2a95b0d7)
	if __obf_c85b895560db528f != nil {
		return __obf_c85b895560db528f
	}
	__obf_c85b895560db528f = __obf_257945e5ab511848(__obf_62435d89ebefd5aa, __obf_8667d4fc2a95b0d7)
	if __obf_c85b895560db528f != nil {
		return __obf_c85b895560db528f
	}
	__obf_c85b895560db528f = __obf_3828217e60017848(__obf_62435d89ebefd5aa, __obf_8667d4fc2a95b0d7)
	if __obf_c85b895560db528f != nil {
		return __obf_c85b895560db528f
	}
	__obf_0e6540508f0e8fa6 := __obf_8667d4fc2a95b0d7.Kind()
	switch __obf_0e6540508f0e8fa6 {
	case reflect.Interface:
		return &__obf_4fb4b91e56916c0b{__obf_8667d4fc2a95b0d7}
	case reflect.Struct:
		return __obf_fca27e6ceeb592ed(__obf_62435d89ebefd5aa, __obf_8667d4fc2a95b0d7)
	case reflect.Array:
		return __obf_66f4ed70a70f0e6c(__obf_62435d89ebefd5aa, __obf_8667d4fc2a95b0d7)
	case reflect.Slice:
		return __obf_b1da7c0e6535593d(__obf_62435d89ebefd5aa, __obf_8667d4fc2a95b0d7)
	case reflect.Map:
		return __obf_f36df00362e2ff3a(__obf_62435d89ebefd5aa, __obf_8667d4fc2a95b0d7)
	case reflect.Ptr:
		return __obf_d2cf57edf2ac20e2(__obf_62435d89ebefd5aa, __obf_8667d4fc2a95b0d7)
	default:
		return &__obf_a580dbbcc400f543{__obf_5966ecc5edb9b17e: fmt.Errorf("%s%s is unsupported type", __obf_62435d89ebefd5aa.__obf_f517983aa5897f07, __obf_8667d4fc2a95b0d7.String())}
	}
}

type __obf_5ba850d645a49aec struct {
	__obf_5966ecc5edb9b17e error
}

func (__obf_924cc7ef585cdfb0 *__obf_5ba850d645a49aec) Decode(__obf_753ab3fb72cdd659 unsafe.Pointer, __obf_edd9fe48d39445e4 *Iterator) {
	if __obf_edd9fe48d39445e4.WhatIsNext() != NilValue {
		if __obf_edd9fe48d39445e4.Error == nil {
			__obf_edd9fe48d39445e4.
				Error = __obf_924cc7ef585cdfb0.__obf_5966ecc5edb9b17e
		}
	} else {
		__obf_edd9fe48d39445e4.
			Skip()
	}
}

type __obf_a580dbbcc400f543 struct {
	__obf_5966ecc5edb9b17e error
}

func (__obf_c85b895560db528f *__obf_a580dbbcc400f543) Encode(__obf_753ab3fb72cdd659 unsafe.Pointer, __obf_2361f5214e84c60f *Stream) {
	if __obf_753ab3fb72cdd659 == nil {
		__obf_2361f5214e84c60f.
			WriteNil()
	} else if __obf_2361f5214e84c60f.Error == nil {
		__obf_2361f5214e84c60f.
			Error = __obf_c85b895560db528f.__obf_5966ecc5edb9b17e
	}
}

func (__obf_c85b895560db528f *__obf_a580dbbcc400f543) IsEmpty(__obf_753ab3fb72cdd659 unsafe.Pointer) bool {
	return false
}

type __obf_44f231ad268e82e4 struct {
	__obf_924cc7ef585cdfb0 ValDecoder
}

func (__obf_924cc7ef585cdfb0 *__obf_44f231ad268e82e4) Decode(__obf_753ab3fb72cdd659 unsafe.Pointer, __obf_edd9fe48d39445e4 *Iterator) {
	__obf_924cc7ef585cdfb0.__obf_924cc7ef585cdfb0.
		Decode(__obf_753ab3fb72cdd659, __obf_edd9fe48d39445e4)
}

type __obf_16e24315cd385130 struct {
	__obf_c85b895560db528f ValEncoder
}

func (__obf_c85b895560db528f *__obf_16e24315cd385130) Encode(__obf_753ab3fb72cdd659 unsafe.Pointer, __obf_2361f5214e84c60f *Stream) {
	__obf_c85b895560db528f.__obf_c85b895560db528f.
		Encode(__obf_753ab3fb72cdd659, __obf_2361f5214e84c60f)
}

func (__obf_c85b895560db528f *__obf_16e24315cd385130) IsEmpty(__obf_753ab3fb72cdd659 unsafe.Pointer) bool {
	return __obf_c85b895560db528f.__obf_c85b895560db528f.IsEmpty(__obf_753ab3fb72cdd659)
}
