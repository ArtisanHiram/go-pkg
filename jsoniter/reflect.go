package __obf_703d23ba520c3295

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
	Decode(__obf_47fa53f3d161f45c unsafe.Pointer, __obf_47edab4c16a2d88d *Iterator)
}

// ValEncoder is an internal type registered to cache as needed.
// Don't confuse jsoniter.ValEncoder with json.Encoder.
// For json.Encoder's adapter, refer to jsoniter.AdapterEncoder(todo godoc link).
type ValEncoder interface {
	IsEmpty(__obf_47fa53f3d161f45c unsafe.Pointer) bool
	Encode(__obf_47fa53f3d161f45c unsafe.Pointer, __obf_9a34f62857fb1d1d *Stream)
}

type __obf_8daa57d6b3a720c3 interface {
	IsEmpty(__obf_47fa53f3d161f45c unsafe.Pointer) bool
}

type __obf_2aaf7367de3ff86d struct {
	*__obf_ef74248d8ae9ba78
	__obf_c5fd316f3c4457de string
	__obf_7b863da151f705ed map[reflect2.Type]ValEncoder
	__obf_7a20f878b9108b4d map[reflect2.Type]ValDecoder
}

func (__obf_85a417ca3a5d43c2 *__obf_2aaf7367de3ff86d) __obf_9332366e75c2e023() bool {
	if __obf_85a417ca3a5d43c2.__obf_ef74248d8ae9ba78 == nil {
		// default is case-insensitive
		return false
	}
	return __obf_85a417ca3a5d43c2.__obf_ef74248d8ae9ba78.__obf_9332366e75c2e023
}

func (__obf_85a417ca3a5d43c2 *__obf_2aaf7367de3ff86d) append(__obf_c5fd316f3c4457de string) *__obf_2aaf7367de3ff86d {
	return &__obf_2aaf7367de3ff86d{__obf_ef74248d8ae9ba78: __obf_85a417ca3a5d43c2.__obf_ef74248d8ae9ba78, __obf_c5fd316f3c4457de: __obf_85a417ca3a5d43c2.__obf_c5fd316f3c4457de + " " + __obf_c5fd316f3c4457de, __obf_7b863da151f705ed: __obf_85a417ca3a5d43c2.__obf_7b863da151f705ed, __obf_7a20f878b9108b4d: __obf_85a417ca3a5d43c2.__obf_7a20f878b9108b4d}
}

// ReadVal copy the underlying JSON into go interface, same as json.Unmarshal
func (__obf_47edab4c16a2d88d *Iterator) ReadVal(__obf_02ba706f4f104d71 any) {
	__obf_5f5fdf79d12ab854 := __obf_47edab4c16a2d88d.__obf_5f5fdf79d12ab854
	__obf_5106f8b01c71afb8 := reflect2.RTypeOf(__obf_02ba706f4f104d71)
	__obf_c9719e68325f7d44 := __obf_47edab4c16a2d88d.__obf_25bd4f754a37b862.__obf_9d061f9912e56c90(__obf_5106f8b01c71afb8)
	if __obf_c9719e68325f7d44 == nil {
		__obf_3b7f6abbae19451e := reflect2.TypeOf(__obf_02ba706f4f104d71)
		if __obf_3b7f6abbae19451e == nil || __obf_3b7f6abbae19451e.Kind() != reflect.Ptr {
			__obf_47edab4c16a2d88d.
				ReportError("ReadVal", "can only unmarshal into pointer")
			return
		}
		__obf_c9719e68325f7d44 = __obf_47edab4c16a2d88d.__obf_25bd4f754a37b862.DecoderOf(__obf_3b7f6abbae19451e)
	}
	__obf_47fa53f3d161f45c := reflect2.PtrOf(__obf_02ba706f4f104d71)
	if __obf_47fa53f3d161f45c == nil {
		__obf_47edab4c16a2d88d.
			ReportError("ReadVal", "can not read into nil pointer")
		return
	}
	__obf_c9719e68325f7d44.
		Decode(__obf_47fa53f3d161f45c, __obf_47edab4c16a2d88d)
	if __obf_47edab4c16a2d88d.__obf_5f5fdf79d12ab854 != __obf_5f5fdf79d12ab854 {
		__obf_47edab4c16a2d88d.
			ReportError("ReadVal", "unexpected mismatched nesting")
		return
	}
}

// WriteVal copy the go interface into underlying JSON, same as json.Marshal
func (__obf_9a34f62857fb1d1d *Stream) WriteVal(__obf_b924538fffe5fd64 any) {
	if nil == __obf_b924538fffe5fd64 {
		__obf_9a34f62857fb1d1d.
			WriteNil()
		return
	}
	__obf_5106f8b01c71afb8 := reflect2.RTypeOf(__obf_b924538fffe5fd64)
	__obf_cf840243a12ee308 := __obf_9a34f62857fb1d1d.__obf_25bd4f754a37b862.__obf_4295547e93f1b46e(__obf_5106f8b01c71afb8)
	if __obf_cf840243a12ee308 == nil {
		__obf_3b7f6abbae19451e := reflect2.TypeOf(__obf_b924538fffe5fd64)
		__obf_cf840243a12ee308 = __obf_9a34f62857fb1d1d.__obf_25bd4f754a37b862.EncoderOf(__obf_3b7f6abbae19451e)
	}
	__obf_cf840243a12ee308.
		Encode(reflect2.PtrOf(__obf_b924538fffe5fd64), __obf_9a34f62857fb1d1d)
}

func (__obf_25bd4f754a37b862 *__obf_ef74248d8ae9ba78) DecoderOf(__obf_3b7f6abbae19451e reflect2.Type) ValDecoder {
	__obf_5106f8b01c71afb8 := __obf_3b7f6abbae19451e.RType()
	__obf_c9719e68325f7d44 := __obf_25bd4f754a37b862.__obf_9d061f9912e56c90(__obf_5106f8b01c71afb8)
	if __obf_c9719e68325f7d44 != nil {
		return __obf_c9719e68325f7d44
	}
	__obf_2aaf7367de3ff86d := &__obf_2aaf7367de3ff86d{__obf_ef74248d8ae9ba78: __obf_25bd4f754a37b862, __obf_c5fd316f3c4457de: "", __obf_7a20f878b9108b4d: map[reflect2.Type]ValDecoder{}, __obf_7b863da151f705ed: map[reflect2.Type]ValEncoder{}}
	__obf_95093efb9321684e := __obf_3b7f6abbae19451e.(*reflect2.UnsafePtrType)
	__obf_c9719e68325f7d44 = __obf_b27fe2bc17d94621(__obf_2aaf7367de3ff86d, __obf_95093efb9321684e.Elem())
	__obf_25bd4f754a37b862.__obf_4bf6eb1a561eb9a1(__obf_5106f8b01c71afb8, __obf_c9719e68325f7d44)
	return __obf_c9719e68325f7d44
}

func __obf_b27fe2bc17d94621(__obf_2aaf7367de3ff86d *__obf_2aaf7367de3ff86d, __obf_3b7f6abbae19451e reflect2.Type) ValDecoder {
	__obf_c9719e68325f7d44 := __obf_3c75139ba7ab27e4(__obf_2aaf7367de3ff86d, __obf_3b7f6abbae19451e)
	if __obf_c9719e68325f7d44 != nil {
		return __obf_c9719e68325f7d44
	}
	__obf_c9719e68325f7d44 = __obf_45dbbcbb35b6ebc6(__obf_2aaf7367de3ff86d, __obf_3b7f6abbae19451e)
	for _, __obf_dfc6ab1ee0bfd79e := range __obf_3e7b8f77e79afd5c {
		__obf_c9719e68325f7d44 = __obf_dfc6ab1ee0bfd79e.DecorateDecoder(__obf_3b7f6abbae19451e, __obf_c9719e68325f7d44)
	}
	__obf_c9719e68325f7d44 = __obf_2aaf7367de3ff86d.__obf_7c9972bb545a908d.DecorateDecoder(__obf_3b7f6abbae19451e, __obf_c9719e68325f7d44)
	for _, __obf_dfc6ab1ee0bfd79e := range __obf_2aaf7367de3ff86d.__obf_324ee50b8c0b2f3b {
		__obf_c9719e68325f7d44 = __obf_dfc6ab1ee0bfd79e.DecorateDecoder(__obf_3b7f6abbae19451e, __obf_c9719e68325f7d44)
	}
	return __obf_c9719e68325f7d44
}

func __obf_45dbbcbb35b6ebc6(__obf_2aaf7367de3ff86d *__obf_2aaf7367de3ff86d, __obf_3b7f6abbae19451e reflect2.Type) ValDecoder {
	__obf_c9719e68325f7d44 := __obf_2aaf7367de3ff86d.__obf_7a20f878b9108b4d[__obf_3b7f6abbae19451e]
	if __obf_c9719e68325f7d44 != nil {
		return __obf_c9719e68325f7d44
	}
	__obf_9296adb04ad51bf5 := &__obf_7cf8be0cfcbe2096{}
	__obf_2aaf7367de3ff86d.__obf_7a20f878b9108b4d[__obf_3b7f6abbae19451e] = __obf_9296adb04ad51bf5
	__obf_c9719e68325f7d44 = _createDecoderOfType(__obf_2aaf7367de3ff86d, __obf_3b7f6abbae19451e)
	__obf_9296adb04ad51bf5.__obf_c9719e68325f7d44 = __obf_c9719e68325f7d44
	return __obf_c9719e68325f7d44
}

func _createDecoderOfType(__obf_2aaf7367de3ff86d *__obf_2aaf7367de3ff86d, __obf_3b7f6abbae19451e reflect2.Type) ValDecoder {
	__obf_c9719e68325f7d44 := __obf_8dbf41d7a8e2a43b(__obf_3b7f6abbae19451e)
	if __obf_c9719e68325f7d44 != nil {
		return __obf_c9719e68325f7d44
	}
	__obf_c9719e68325f7d44 = __obf_1cd9fe2978e7054f(__obf_2aaf7367de3ff86d, __obf_3b7f6abbae19451e)
	if __obf_c9719e68325f7d44 != nil {
		return __obf_c9719e68325f7d44
	}
	__obf_c9719e68325f7d44 = __obf_8983286d3baa6ceb(__obf_2aaf7367de3ff86d, __obf_3b7f6abbae19451e)
	if __obf_c9719e68325f7d44 != nil {
		return __obf_c9719e68325f7d44
	}
	__obf_c9719e68325f7d44 = __obf_9e92fe737ca75ec6(__obf_2aaf7367de3ff86d, __obf_3b7f6abbae19451e)
	if __obf_c9719e68325f7d44 != nil {
		return __obf_c9719e68325f7d44
	}
	__obf_c9719e68325f7d44 = __obf_d310ff63d30070e1(__obf_2aaf7367de3ff86d, __obf_3b7f6abbae19451e)
	if __obf_c9719e68325f7d44 != nil {
		return __obf_c9719e68325f7d44
	}
	switch __obf_3b7f6abbae19451e.Kind() {
	case reflect.Interface:
		__obf_fa5bf8cb8f34498f, __obf_c0ab4528cf48f779 := __obf_3b7f6abbae19451e.(*reflect2.UnsafeIFaceType)
		if __obf_c0ab4528cf48f779 {
			return &__obf_a2ad861d69c8a6ea{__obf_15da62f311934a45: __obf_fa5bf8cb8f34498f}
		}
		return &__obf_a2ef1129a194d05a{}
	case reflect.Struct:
		return __obf_be37b1a24a3a47e8(__obf_2aaf7367de3ff86d, __obf_3b7f6abbae19451e)
	case reflect.Array:
		return __obf_71ca52bfb4b492aa(__obf_2aaf7367de3ff86d, __obf_3b7f6abbae19451e)
	case reflect.Slice:
		return __obf_aeff82b222c8ff5b(__obf_2aaf7367de3ff86d, __obf_3b7f6abbae19451e)
	case reflect.Map:
		return __obf_ef246e5fc788f28b(__obf_2aaf7367de3ff86d, __obf_3b7f6abbae19451e)
	case reflect.Ptr:
		return __obf_490e149853cb1478(__obf_2aaf7367de3ff86d, __obf_3b7f6abbae19451e)
	default:
		return &__obf_63b2b5a85529d72d{__obf_e56742d6e52f642e: fmt.Errorf("%s%s is unsupported type", __obf_2aaf7367de3ff86d.__obf_c5fd316f3c4457de, __obf_3b7f6abbae19451e.String())}
	}
}

func (__obf_25bd4f754a37b862 *__obf_ef74248d8ae9ba78) EncoderOf(__obf_3b7f6abbae19451e reflect2.Type) ValEncoder {
	__obf_5106f8b01c71afb8 := __obf_3b7f6abbae19451e.RType()
	__obf_cf840243a12ee308 := __obf_25bd4f754a37b862.__obf_4295547e93f1b46e(__obf_5106f8b01c71afb8)
	if __obf_cf840243a12ee308 != nil {
		return __obf_cf840243a12ee308
	}
	__obf_2aaf7367de3ff86d := &__obf_2aaf7367de3ff86d{__obf_ef74248d8ae9ba78: __obf_25bd4f754a37b862, __obf_c5fd316f3c4457de: "", __obf_7a20f878b9108b4d: map[reflect2.Type]ValDecoder{}, __obf_7b863da151f705ed: map[reflect2.Type]ValEncoder{}}
	__obf_cf840243a12ee308 = __obf_906496c658b349cf(__obf_2aaf7367de3ff86d, __obf_3b7f6abbae19451e)
	if __obf_3b7f6abbae19451e.LikePtr() {
		__obf_cf840243a12ee308 = &__obf_a2f77682b8377346{__obf_cf840243a12ee308}
	}
	__obf_25bd4f754a37b862.__obf_722868f5d55df1b8(__obf_5106f8b01c71afb8, __obf_cf840243a12ee308)
	return __obf_cf840243a12ee308
}

type __obf_a2f77682b8377346 struct {
	__obf_cf840243a12ee308 ValEncoder
}

func (__obf_cf840243a12ee308 *__obf_a2f77682b8377346) IsEmpty(__obf_47fa53f3d161f45c unsafe.Pointer) bool {
	return __obf_cf840243a12ee308.__obf_cf840243a12ee308.IsEmpty(unsafe.Pointer(&__obf_47fa53f3d161f45c))
}

func (__obf_cf840243a12ee308 *__obf_a2f77682b8377346) Encode(__obf_47fa53f3d161f45c unsafe.Pointer, __obf_9a34f62857fb1d1d *Stream) {
	__obf_cf840243a12ee308.__obf_cf840243a12ee308.
		Encode(unsafe.Pointer(&__obf_47fa53f3d161f45c), __obf_9a34f62857fb1d1d)
}

func __obf_906496c658b349cf(__obf_2aaf7367de3ff86d *__obf_2aaf7367de3ff86d, __obf_3b7f6abbae19451e reflect2.Type) ValEncoder {
	__obf_cf840243a12ee308 := __obf_6e923984d5a5d1d7(__obf_2aaf7367de3ff86d, __obf_3b7f6abbae19451e)
	if __obf_cf840243a12ee308 != nil {
		return __obf_cf840243a12ee308
	}
	__obf_cf840243a12ee308 = __obf_6b20c96ea553b73d(__obf_2aaf7367de3ff86d, __obf_3b7f6abbae19451e)
	for _, __obf_dfc6ab1ee0bfd79e := range __obf_3e7b8f77e79afd5c {
		__obf_cf840243a12ee308 = __obf_dfc6ab1ee0bfd79e.DecorateEncoder(__obf_3b7f6abbae19451e, __obf_cf840243a12ee308)
	}
	__obf_cf840243a12ee308 = __obf_2aaf7367de3ff86d.__obf_9289f2028bb085f2.DecorateEncoder(__obf_3b7f6abbae19451e, __obf_cf840243a12ee308)
	for _, __obf_dfc6ab1ee0bfd79e := range __obf_2aaf7367de3ff86d.__obf_324ee50b8c0b2f3b {
		__obf_cf840243a12ee308 = __obf_dfc6ab1ee0bfd79e.DecorateEncoder(__obf_3b7f6abbae19451e, __obf_cf840243a12ee308)
	}
	return __obf_cf840243a12ee308
}

func __obf_6b20c96ea553b73d(__obf_2aaf7367de3ff86d *__obf_2aaf7367de3ff86d, __obf_3b7f6abbae19451e reflect2.Type) ValEncoder {
	__obf_cf840243a12ee308 := __obf_2aaf7367de3ff86d.__obf_7b863da151f705ed[__obf_3b7f6abbae19451e]
	if __obf_cf840243a12ee308 != nil {
		return __obf_cf840243a12ee308
	}
	__obf_9296adb04ad51bf5 := &__obf_a74191d20e2f2013{}
	__obf_2aaf7367de3ff86d.__obf_7b863da151f705ed[__obf_3b7f6abbae19451e] = __obf_9296adb04ad51bf5
	__obf_cf840243a12ee308 = _createEncoderOfType(__obf_2aaf7367de3ff86d, __obf_3b7f6abbae19451e)
	__obf_9296adb04ad51bf5.__obf_cf840243a12ee308 = __obf_cf840243a12ee308
	return __obf_cf840243a12ee308
}
func _createEncoderOfType(__obf_2aaf7367de3ff86d *__obf_2aaf7367de3ff86d, __obf_3b7f6abbae19451e reflect2.Type) ValEncoder {
	__obf_cf840243a12ee308 := __obf_f3e98e16ceefbfa6(__obf_3b7f6abbae19451e)
	if __obf_cf840243a12ee308 != nil {
		return __obf_cf840243a12ee308
	}
	__obf_cf840243a12ee308 = __obf_0c0040a8d86c5f59(__obf_2aaf7367de3ff86d, __obf_3b7f6abbae19451e)
	if __obf_cf840243a12ee308 != nil {
		return __obf_cf840243a12ee308
	}
	__obf_cf840243a12ee308 = __obf_888be1be2b861d32(__obf_2aaf7367de3ff86d, __obf_3b7f6abbae19451e)
	if __obf_cf840243a12ee308 != nil {
		return __obf_cf840243a12ee308
	}
	__obf_cf840243a12ee308 = __obf_733be021e2b2eaf7(__obf_2aaf7367de3ff86d, __obf_3b7f6abbae19451e)
	if __obf_cf840243a12ee308 != nil {
		return __obf_cf840243a12ee308
	}
	__obf_cf840243a12ee308 = __obf_3774fbf920f93937(__obf_2aaf7367de3ff86d, __obf_3b7f6abbae19451e)
	if __obf_cf840243a12ee308 != nil {
		return __obf_cf840243a12ee308
	}
	__obf_2bb4ea546eb4165c := __obf_3b7f6abbae19451e.Kind()
	switch __obf_2bb4ea546eb4165c {
	case reflect.Interface:
		return &__obf_7f31cdbc4d682d50{__obf_3b7f6abbae19451e}
	case reflect.Struct:
		return __obf_76e40d1501c65ef3(__obf_2aaf7367de3ff86d, __obf_3b7f6abbae19451e)
	case reflect.Array:
		return __obf_55bee799e6e81597(__obf_2aaf7367de3ff86d, __obf_3b7f6abbae19451e)
	case reflect.Slice:
		return __obf_2aeca70011a01b7b(__obf_2aaf7367de3ff86d, __obf_3b7f6abbae19451e)
	case reflect.Map:
		return __obf_e70914e98bbb608c(__obf_2aaf7367de3ff86d, __obf_3b7f6abbae19451e)
	case reflect.Ptr:
		return __obf_9a820bfca585fe9d(__obf_2aaf7367de3ff86d, __obf_3b7f6abbae19451e)
	default:
		return &__obf_f80a54e26786ebf2{__obf_e56742d6e52f642e: fmt.Errorf("%s%s is unsupported type", __obf_2aaf7367de3ff86d.__obf_c5fd316f3c4457de, __obf_3b7f6abbae19451e.String())}
	}
}

type __obf_63b2b5a85529d72d struct {
	__obf_e56742d6e52f642e error
}

func (__obf_c9719e68325f7d44 *__obf_63b2b5a85529d72d) Decode(__obf_47fa53f3d161f45c unsafe.Pointer, __obf_47edab4c16a2d88d *Iterator) {
	if __obf_47edab4c16a2d88d.WhatIsNext() != NilValue {
		if __obf_47edab4c16a2d88d.Error == nil {
			__obf_47edab4c16a2d88d.
				Error = __obf_c9719e68325f7d44.__obf_e56742d6e52f642e
		}
	} else {
		__obf_47edab4c16a2d88d.
			Skip()
	}
}

type __obf_f80a54e26786ebf2 struct {
	__obf_e56742d6e52f642e error
}

func (__obf_cf840243a12ee308 *__obf_f80a54e26786ebf2) Encode(__obf_47fa53f3d161f45c unsafe.Pointer, __obf_9a34f62857fb1d1d *Stream) {
	if __obf_47fa53f3d161f45c == nil {
		__obf_9a34f62857fb1d1d.
			WriteNil()
	} else if __obf_9a34f62857fb1d1d.Error == nil {
		__obf_9a34f62857fb1d1d.
			Error = __obf_cf840243a12ee308.__obf_e56742d6e52f642e
	}
}

func (__obf_cf840243a12ee308 *__obf_f80a54e26786ebf2) IsEmpty(__obf_47fa53f3d161f45c unsafe.Pointer) bool {
	return false
}

type __obf_7cf8be0cfcbe2096 struct {
	__obf_c9719e68325f7d44 ValDecoder
}

func (__obf_c9719e68325f7d44 *__obf_7cf8be0cfcbe2096) Decode(__obf_47fa53f3d161f45c unsafe.Pointer, __obf_47edab4c16a2d88d *Iterator) {
	__obf_c9719e68325f7d44.__obf_c9719e68325f7d44.
		Decode(__obf_47fa53f3d161f45c, __obf_47edab4c16a2d88d)
}

type __obf_a74191d20e2f2013 struct {
	__obf_cf840243a12ee308 ValEncoder
}

func (__obf_cf840243a12ee308 *__obf_a74191d20e2f2013) Encode(__obf_47fa53f3d161f45c unsafe.Pointer, __obf_9a34f62857fb1d1d *Stream) {
	__obf_cf840243a12ee308.__obf_cf840243a12ee308.
		Encode(__obf_47fa53f3d161f45c, __obf_9a34f62857fb1d1d)
}

func (__obf_cf840243a12ee308 *__obf_a74191d20e2f2013) IsEmpty(__obf_47fa53f3d161f45c unsafe.Pointer) bool {
	return __obf_cf840243a12ee308.__obf_cf840243a12ee308.IsEmpty(__obf_47fa53f3d161f45c)
}
