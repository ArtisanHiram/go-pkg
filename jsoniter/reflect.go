package __obf_030d39f7a8de96c6

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
	Decode(__obf_dbbf371b91136494 unsafe.Pointer, __obf_4ab56a99965952e7 *Iterator)
}

// ValEncoder is an internal type registered to cache as needed.
// Don't confuse jsoniter.ValEncoder with json.Encoder.
// For json.Encoder's adapter, refer to jsoniter.AdapterEncoder(todo godoc link).
type ValEncoder interface {
	IsEmpty(__obf_dbbf371b91136494 unsafe.Pointer) bool
	Encode(__obf_dbbf371b91136494 unsafe.Pointer, __obf_8a2c51fe22775655 *Stream)
}

type __obf_b5475e6763577f46 interface {
	IsEmpty(__obf_dbbf371b91136494 unsafe.Pointer) bool
}

type __obf_a565fbaffca944f0 struct {
	*__obf_81639538813814ff
	__obf_834e1679b8081f46 string
	__obf_2a91a6a342abfb1a map[reflect2.Type]ValEncoder
	__obf_660c95f82dedaffe map[reflect2.Type]ValDecoder
}

func (__obf_ea107e11b66371dd *__obf_a565fbaffca944f0) __obf_af13e3babb780a4e() bool {
	if __obf_ea107e11b66371dd.__obf_81639538813814ff == nil {
		// default is case-insensitive
		return false
	}
	return __obf_ea107e11b66371dd.__obf_81639538813814ff.__obf_af13e3babb780a4e
}

func (__obf_ea107e11b66371dd *__obf_a565fbaffca944f0) append(__obf_834e1679b8081f46 string) *__obf_a565fbaffca944f0 {
	return &__obf_a565fbaffca944f0{__obf_81639538813814ff: __obf_ea107e11b66371dd.__obf_81639538813814ff, __obf_834e1679b8081f46: __obf_ea107e11b66371dd.__obf_834e1679b8081f46 + " " + __obf_834e1679b8081f46, __obf_2a91a6a342abfb1a: __obf_ea107e11b66371dd.__obf_2a91a6a342abfb1a, __obf_660c95f82dedaffe: __obf_ea107e11b66371dd.__obf_660c95f82dedaffe}
}

// ReadVal copy the underlying JSON into go interface, same as json.Unmarshal
func (__obf_4ab56a99965952e7 *Iterator) ReadVal(__obf_eefa92392cc2442c any) {
	__obf_66f73b1f1e6b88e2 := __obf_4ab56a99965952e7.__obf_66f73b1f1e6b88e2
	__obf_5ddcbb0301e0cf46 := reflect2.RTypeOf(__obf_eefa92392cc2442c)
	__obf_11a3f28116164d09 := __obf_4ab56a99965952e7.__obf_679611dc56ff465b.__obf_dff673ee34435a0e(__obf_5ddcbb0301e0cf46)
	if __obf_11a3f28116164d09 == nil {
		__obf_8744d0b8c80ccdc1 := reflect2.TypeOf(__obf_eefa92392cc2442c)
		if __obf_8744d0b8c80ccdc1 == nil || __obf_8744d0b8c80ccdc1.Kind() != reflect.Ptr {
			__obf_4ab56a99965952e7.
				ReportError("ReadVal", "can only unmarshal into pointer")
			return
		}
		__obf_11a3f28116164d09 = __obf_4ab56a99965952e7.__obf_679611dc56ff465b.DecoderOf(__obf_8744d0b8c80ccdc1)
	}
	__obf_dbbf371b91136494 := reflect2.PtrOf(__obf_eefa92392cc2442c)
	if __obf_dbbf371b91136494 == nil {
		__obf_4ab56a99965952e7.
			ReportError("ReadVal", "can not read into nil pointer")
		return
	}
	__obf_11a3f28116164d09.
		Decode(__obf_dbbf371b91136494, __obf_4ab56a99965952e7)
	if __obf_4ab56a99965952e7.__obf_66f73b1f1e6b88e2 != __obf_66f73b1f1e6b88e2 {
		__obf_4ab56a99965952e7.
			ReportError("ReadVal", "unexpected mismatched nesting")
		return
	}
}

// WriteVal copy the go interface into underlying JSON, same as json.Marshal
func (__obf_8a2c51fe22775655 *Stream) WriteVal(__obf_efaf2719b44ad8ac any) {
	if nil == __obf_efaf2719b44ad8ac {
		__obf_8a2c51fe22775655.
			WriteNil()
		return
	}
	__obf_5ddcbb0301e0cf46 := reflect2.RTypeOf(__obf_efaf2719b44ad8ac)
	__obf_008f61596d7e9523 := __obf_8a2c51fe22775655.__obf_679611dc56ff465b.__obf_f3694d09db16d770(__obf_5ddcbb0301e0cf46)
	if __obf_008f61596d7e9523 == nil {
		__obf_8744d0b8c80ccdc1 := reflect2.TypeOf(__obf_efaf2719b44ad8ac)
		__obf_008f61596d7e9523 = __obf_8a2c51fe22775655.__obf_679611dc56ff465b.EncoderOf(__obf_8744d0b8c80ccdc1)
	}
	__obf_008f61596d7e9523.
		Encode(reflect2.PtrOf(__obf_efaf2719b44ad8ac), __obf_8a2c51fe22775655)
}

func (__obf_679611dc56ff465b *__obf_81639538813814ff) DecoderOf(__obf_8744d0b8c80ccdc1 reflect2.Type) ValDecoder {
	__obf_5ddcbb0301e0cf46 := __obf_8744d0b8c80ccdc1.RType()
	__obf_11a3f28116164d09 := __obf_679611dc56ff465b.__obf_dff673ee34435a0e(__obf_5ddcbb0301e0cf46)
	if __obf_11a3f28116164d09 != nil {
		return __obf_11a3f28116164d09
	}
	__obf_a565fbaffca944f0 := &__obf_a565fbaffca944f0{__obf_81639538813814ff: __obf_679611dc56ff465b, __obf_834e1679b8081f46: "", __obf_660c95f82dedaffe: map[reflect2.Type]ValDecoder{}, __obf_2a91a6a342abfb1a: map[reflect2.Type]ValEncoder{}}
	__obf_3a51157620f9910b := __obf_8744d0b8c80ccdc1.(*reflect2.UnsafePtrType)
	__obf_11a3f28116164d09 = __obf_729ef08c8b8e060e(__obf_a565fbaffca944f0, __obf_3a51157620f9910b.Elem())
	__obf_679611dc56ff465b.__obf_0599525836bd5d66(__obf_5ddcbb0301e0cf46, __obf_11a3f28116164d09)
	return __obf_11a3f28116164d09
}

func __obf_729ef08c8b8e060e(__obf_a565fbaffca944f0 *__obf_a565fbaffca944f0, __obf_8744d0b8c80ccdc1 reflect2.Type) ValDecoder {
	__obf_11a3f28116164d09 := __obf_8a762691a84adf6f(__obf_a565fbaffca944f0, __obf_8744d0b8c80ccdc1)
	if __obf_11a3f28116164d09 != nil {
		return __obf_11a3f28116164d09
	}
	__obf_11a3f28116164d09 = __obf_0c188ea02a9c0025(__obf_a565fbaffca944f0, __obf_8744d0b8c80ccdc1)
	for _, __obf_621544a57e52000f := range __obf_04498d7bbd7ab2ba {
		__obf_11a3f28116164d09 = __obf_621544a57e52000f.DecorateDecoder(__obf_8744d0b8c80ccdc1, __obf_11a3f28116164d09)
	}
	__obf_11a3f28116164d09 = __obf_a565fbaffca944f0.__obf_6313bfb9c913bd50.DecorateDecoder(__obf_8744d0b8c80ccdc1, __obf_11a3f28116164d09)
	for _, __obf_621544a57e52000f := range __obf_a565fbaffca944f0.__obf_6621255bc1f80c8e {
		__obf_11a3f28116164d09 = __obf_621544a57e52000f.DecorateDecoder(__obf_8744d0b8c80ccdc1, __obf_11a3f28116164d09)
	}
	return __obf_11a3f28116164d09
}

func __obf_0c188ea02a9c0025(__obf_a565fbaffca944f0 *__obf_a565fbaffca944f0, __obf_8744d0b8c80ccdc1 reflect2.Type) ValDecoder {
	__obf_11a3f28116164d09 := __obf_a565fbaffca944f0.__obf_660c95f82dedaffe[__obf_8744d0b8c80ccdc1]
	if __obf_11a3f28116164d09 != nil {
		return __obf_11a3f28116164d09
	}
	__obf_e4948ae8f6e37b72 := &__obf_733bd443cfa30455{}
	__obf_a565fbaffca944f0.__obf_660c95f82dedaffe[__obf_8744d0b8c80ccdc1] = __obf_e4948ae8f6e37b72
	__obf_11a3f28116164d09 = _createDecoderOfType(__obf_a565fbaffca944f0, __obf_8744d0b8c80ccdc1)
	__obf_e4948ae8f6e37b72.__obf_11a3f28116164d09 = __obf_11a3f28116164d09
	return __obf_11a3f28116164d09
}

func _createDecoderOfType(__obf_a565fbaffca944f0 *__obf_a565fbaffca944f0, __obf_8744d0b8c80ccdc1 reflect2.Type) ValDecoder {
	__obf_11a3f28116164d09 := __obf_3d1edc0e2c1c1e07(__obf_8744d0b8c80ccdc1)
	if __obf_11a3f28116164d09 != nil {
		return __obf_11a3f28116164d09
	}
	__obf_11a3f28116164d09 = __obf_18e6804367792db0(__obf_a565fbaffca944f0, __obf_8744d0b8c80ccdc1)
	if __obf_11a3f28116164d09 != nil {
		return __obf_11a3f28116164d09
	}
	__obf_11a3f28116164d09 = __obf_a0c12030f23db855(__obf_a565fbaffca944f0, __obf_8744d0b8c80ccdc1)
	if __obf_11a3f28116164d09 != nil {
		return __obf_11a3f28116164d09
	}
	__obf_11a3f28116164d09 = __obf_b457e3c1bf18f46a(__obf_a565fbaffca944f0, __obf_8744d0b8c80ccdc1)
	if __obf_11a3f28116164d09 != nil {
		return __obf_11a3f28116164d09
	}
	__obf_11a3f28116164d09 = __obf_506acfeb1c549d6a(__obf_a565fbaffca944f0, __obf_8744d0b8c80ccdc1)
	if __obf_11a3f28116164d09 != nil {
		return __obf_11a3f28116164d09
	}
	switch __obf_8744d0b8c80ccdc1.Kind() {
	case reflect.Interface:
		__obf_3c844a64621b06a1, __obf_0a552960b2a5212a := __obf_8744d0b8c80ccdc1.(*reflect2.UnsafeIFaceType)
		if __obf_0a552960b2a5212a {
			return &__obf_ee8a1f4d7afb7af9{__obf_a7610e23a63622fd: __obf_3c844a64621b06a1}
		}
		return &__obf_b4fd1ed1693bf959{}
	case reflect.Struct:
		return __obf_0891d5da085d6093(__obf_a565fbaffca944f0, __obf_8744d0b8c80ccdc1)
	case reflect.Array:
		return __obf_9dc19cae1724dfa5(__obf_a565fbaffca944f0, __obf_8744d0b8c80ccdc1)
	case reflect.Slice:
		return __obf_6f35214e10cf09dc(__obf_a565fbaffca944f0, __obf_8744d0b8c80ccdc1)
	case reflect.Map:
		return __obf_d63eb7ded243ae52(__obf_a565fbaffca944f0, __obf_8744d0b8c80ccdc1)
	case reflect.Ptr:
		return __obf_2ca1ae44e249759c(__obf_a565fbaffca944f0, __obf_8744d0b8c80ccdc1)
	default:
		return &__obf_195d4c78d647b89f{__obf_fcc907dd69879566: fmt.Errorf("%s%s is unsupported type", __obf_a565fbaffca944f0.__obf_834e1679b8081f46, __obf_8744d0b8c80ccdc1.String())}
	}
}

func (__obf_679611dc56ff465b *__obf_81639538813814ff) EncoderOf(__obf_8744d0b8c80ccdc1 reflect2.Type) ValEncoder {
	__obf_5ddcbb0301e0cf46 := __obf_8744d0b8c80ccdc1.RType()
	__obf_008f61596d7e9523 := __obf_679611dc56ff465b.__obf_f3694d09db16d770(__obf_5ddcbb0301e0cf46)
	if __obf_008f61596d7e9523 != nil {
		return __obf_008f61596d7e9523
	}
	__obf_a565fbaffca944f0 := &__obf_a565fbaffca944f0{__obf_81639538813814ff: __obf_679611dc56ff465b, __obf_834e1679b8081f46: "", __obf_660c95f82dedaffe: map[reflect2.Type]ValDecoder{}, __obf_2a91a6a342abfb1a: map[reflect2.Type]ValEncoder{}}
	__obf_008f61596d7e9523 = __obf_37ef471fa5e40404(__obf_a565fbaffca944f0, __obf_8744d0b8c80ccdc1)
	if __obf_8744d0b8c80ccdc1.LikePtr() {
		__obf_008f61596d7e9523 = &__obf_f719899ddc8f9bae{__obf_008f61596d7e9523}
	}
	__obf_679611dc56ff465b.__obf_ec0f7f63f3dc6639(__obf_5ddcbb0301e0cf46, __obf_008f61596d7e9523)
	return __obf_008f61596d7e9523
}

type __obf_f719899ddc8f9bae struct {
	__obf_008f61596d7e9523 ValEncoder
}

func (__obf_008f61596d7e9523 *__obf_f719899ddc8f9bae) IsEmpty(__obf_dbbf371b91136494 unsafe.Pointer) bool {
	return __obf_008f61596d7e9523.__obf_008f61596d7e9523.IsEmpty(unsafe.Pointer(&__obf_dbbf371b91136494))
}

func (__obf_008f61596d7e9523 *__obf_f719899ddc8f9bae) Encode(__obf_dbbf371b91136494 unsafe.Pointer, __obf_8a2c51fe22775655 *Stream) {
	__obf_008f61596d7e9523.__obf_008f61596d7e9523.
		Encode(unsafe.Pointer(&__obf_dbbf371b91136494), __obf_8a2c51fe22775655)
}

func __obf_37ef471fa5e40404(__obf_a565fbaffca944f0 *__obf_a565fbaffca944f0, __obf_8744d0b8c80ccdc1 reflect2.Type) ValEncoder {
	__obf_008f61596d7e9523 := __obf_4f7cd8fb242103cc(__obf_a565fbaffca944f0, __obf_8744d0b8c80ccdc1)
	if __obf_008f61596d7e9523 != nil {
		return __obf_008f61596d7e9523
	}
	__obf_008f61596d7e9523 = __obf_8dd9c351c8c3d090(__obf_a565fbaffca944f0, __obf_8744d0b8c80ccdc1)
	for _, __obf_621544a57e52000f := range __obf_04498d7bbd7ab2ba {
		__obf_008f61596d7e9523 = __obf_621544a57e52000f.DecorateEncoder(__obf_8744d0b8c80ccdc1, __obf_008f61596d7e9523)
	}
	__obf_008f61596d7e9523 = __obf_a565fbaffca944f0.__obf_7db1f0a55b319e45.DecorateEncoder(__obf_8744d0b8c80ccdc1, __obf_008f61596d7e9523)
	for _, __obf_621544a57e52000f := range __obf_a565fbaffca944f0.__obf_6621255bc1f80c8e {
		__obf_008f61596d7e9523 = __obf_621544a57e52000f.DecorateEncoder(__obf_8744d0b8c80ccdc1, __obf_008f61596d7e9523)
	}
	return __obf_008f61596d7e9523
}

func __obf_8dd9c351c8c3d090(__obf_a565fbaffca944f0 *__obf_a565fbaffca944f0, __obf_8744d0b8c80ccdc1 reflect2.Type) ValEncoder {
	__obf_008f61596d7e9523 := __obf_a565fbaffca944f0.__obf_2a91a6a342abfb1a[__obf_8744d0b8c80ccdc1]
	if __obf_008f61596d7e9523 != nil {
		return __obf_008f61596d7e9523
	}
	__obf_e4948ae8f6e37b72 := &__obf_54037d93d833eecf{}
	__obf_a565fbaffca944f0.__obf_2a91a6a342abfb1a[__obf_8744d0b8c80ccdc1] = __obf_e4948ae8f6e37b72
	__obf_008f61596d7e9523 = _createEncoderOfType(__obf_a565fbaffca944f0, __obf_8744d0b8c80ccdc1)
	__obf_e4948ae8f6e37b72.__obf_008f61596d7e9523 = __obf_008f61596d7e9523
	return __obf_008f61596d7e9523
}
func _createEncoderOfType(__obf_a565fbaffca944f0 *__obf_a565fbaffca944f0, __obf_8744d0b8c80ccdc1 reflect2.Type) ValEncoder {
	__obf_008f61596d7e9523 := __obf_8a2ee1f3ca9358c1(__obf_8744d0b8c80ccdc1)
	if __obf_008f61596d7e9523 != nil {
		return __obf_008f61596d7e9523
	}
	__obf_008f61596d7e9523 = __obf_21ace3604d05bc8e(__obf_a565fbaffca944f0, __obf_8744d0b8c80ccdc1)
	if __obf_008f61596d7e9523 != nil {
		return __obf_008f61596d7e9523
	}
	__obf_008f61596d7e9523 = __obf_35f11390081cc107(__obf_a565fbaffca944f0, __obf_8744d0b8c80ccdc1)
	if __obf_008f61596d7e9523 != nil {
		return __obf_008f61596d7e9523
	}
	__obf_008f61596d7e9523 = __obf_9ed4e77098f589a9(__obf_a565fbaffca944f0, __obf_8744d0b8c80ccdc1)
	if __obf_008f61596d7e9523 != nil {
		return __obf_008f61596d7e9523
	}
	__obf_008f61596d7e9523 = __obf_438eebf6331ef15b(__obf_a565fbaffca944f0, __obf_8744d0b8c80ccdc1)
	if __obf_008f61596d7e9523 != nil {
		return __obf_008f61596d7e9523
	}
	__obf_30e494a4f2832c2f := __obf_8744d0b8c80ccdc1.Kind()
	switch __obf_30e494a4f2832c2f {
	case reflect.Interface:
		return &__obf_3691828d7304f473{__obf_8744d0b8c80ccdc1}
	case reflect.Struct:
		return __obf_9070536cbef1ec38(__obf_a565fbaffca944f0, __obf_8744d0b8c80ccdc1)
	case reflect.Array:
		return __obf_dd1bf627e51d80bf(__obf_a565fbaffca944f0, __obf_8744d0b8c80ccdc1)
	case reflect.Slice:
		return __obf_f70c82e772cafe09(__obf_a565fbaffca944f0, __obf_8744d0b8c80ccdc1)
	case reflect.Map:
		return __obf_5cc5784cdf0ac33a(__obf_a565fbaffca944f0, __obf_8744d0b8c80ccdc1)
	case reflect.Ptr:
		return __obf_bb9550a27dc28be6(__obf_a565fbaffca944f0, __obf_8744d0b8c80ccdc1)
	default:
		return &__obf_f14cb8ddffae87b3{__obf_fcc907dd69879566: fmt.Errorf("%s%s is unsupported type", __obf_a565fbaffca944f0.__obf_834e1679b8081f46, __obf_8744d0b8c80ccdc1.String())}
	}
}

type __obf_195d4c78d647b89f struct {
	__obf_fcc907dd69879566 error
}

func (__obf_11a3f28116164d09 *__obf_195d4c78d647b89f) Decode(__obf_dbbf371b91136494 unsafe.Pointer, __obf_4ab56a99965952e7 *Iterator) {
	if __obf_4ab56a99965952e7.WhatIsNext() != NilValue {
		if __obf_4ab56a99965952e7.Error == nil {
			__obf_4ab56a99965952e7.
				Error = __obf_11a3f28116164d09.__obf_fcc907dd69879566
		}
	} else {
		__obf_4ab56a99965952e7.
			Skip()
	}
}

type __obf_f14cb8ddffae87b3 struct {
	__obf_fcc907dd69879566 error
}

func (__obf_008f61596d7e9523 *__obf_f14cb8ddffae87b3) Encode(__obf_dbbf371b91136494 unsafe.Pointer, __obf_8a2c51fe22775655 *Stream) {
	if __obf_dbbf371b91136494 == nil {
		__obf_8a2c51fe22775655.
			WriteNil()
	} else if __obf_8a2c51fe22775655.Error == nil {
		__obf_8a2c51fe22775655.
			Error = __obf_008f61596d7e9523.__obf_fcc907dd69879566
	}
}

func (__obf_008f61596d7e9523 *__obf_f14cb8ddffae87b3) IsEmpty(__obf_dbbf371b91136494 unsafe.Pointer) bool {
	return false
}

type __obf_733bd443cfa30455 struct {
	__obf_11a3f28116164d09 ValDecoder
}

func (__obf_11a3f28116164d09 *__obf_733bd443cfa30455) Decode(__obf_dbbf371b91136494 unsafe.Pointer, __obf_4ab56a99965952e7 *Iterator) {
	__obf_11a3f28116164d09.__obf_11a3f28116164d09.
		Decode(__obf_dbbf371b91136494, __obf_4ab56a99965952e7)
}

type __obf_54037d93d833eecf struct {
	__obf_008f61596d7e9523 ValEncoder
}

func (__obf_008f61596d7e9523 *__obf_54037d93d833eecf) Encode(__obf_dbbf371b91136494 unsafe.Pointer, __obf_8a2c51fe22775655 *Stream) {
	__obf_008f61596d7e9523.__obf_008f61596d7e9523.
		Encode(__obf_dbbf371b91136494, __obf_8a2c51fe22775655)
}

func (__obf_008f61596d7e9523 *__obf_54037d93d833eecf) IsEmpty(__obf_dbbf371b91136494 unsafe.Pointer) bool {
	return __obf_008f61596d7e9523.__obf_008f61596d7e9523.IsEmpty(__obf_dbbf371b91136494)
}
