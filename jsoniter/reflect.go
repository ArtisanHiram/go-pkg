package __obf_c7b79b12b33d8238

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
	Decode(__obf_575c04f2b9d91315 unsafe.Pointer, __obf_0da8c843dad13139 *Iterator)
}

// ValEncoder is an internal type registered to cache as needed.
// Don't confuse jsoniter.ValEncoder with json.Encoder.
// For json.Encoder's adapter, refer to jsoniter.AdapterEncoder(todo godoc link).
type ValEncoder interface {
	IsEmpty(__obf_575c04f2b9d91315 unsafe.Pointer) bool
	Encode(__obf_575c04f2b9d91315 unsafe.Pointer, __obf_d8f50bcfe87d8b47 *Stream)
}

type __obf_07e25a5cafedd0df interface {
	IsEmpty(__obf_575c04f2b9d91315 unsafe.Pointer) bool
}

type __obf_99ec45908bceabdb struct {
	*__obf_30fe5c95cabd69c0
	__obf_5de9ece8fa3a16e3 string
	__obf_2f5d685c6dbf1e22 map[reflect2.Type]ValEncoder
	__obf_ecab71bd3ecfd43c map[reflect2.Type]ValDecoder
}

func (__obf_fa8a3db302183a92 *__obf_99ec45908bceabdb) __obf_3704f04b7ac67609() bool {
	if __obf_fa8a3db302183a92.__obf_30fe5c95cabd69c0 == nil {
		// default is case-insensitive
		return false
	}
	return __obf_fa8a3db302183a92.__obf_30fe5c95cabd69c0.__obf_3704f04b7ac67609
}

func (__obf_fa8a3db302183a92 *__obf_99ec45908bceabdb) append(__obf_5de9ece8fa3a16e3 string) *__obf_99ec45908bceabdb {
	return &__obf_99ec45908bceabdb{__obf_30fe5c95cabd69c0: __obf_fa8a3db302183a92.__obf_30fe5c95cabd69c0, __obf_5de9ece8fa3a16e3: __obf_fa8a3db302183a92.__obf_5de9ece8fa3a16e3 + " " + __obf_5de9ece8fa3a16e3, __obf_2f5d685c6dbf1e22: __obf_fa8a3db302183a92.__obf_2f5d685c6dbf1e22, __obf_ecab71bd3ecfd43c: __obf_fa8a3db302183a92.__obf_ecab71bd3ecfd43c}
}

// ReadVal copy the underlying JSON into go interface, same as json.Unmarshal
func (__obf_0da8c843dad13139 *Iterator) ReadVal(__obf_d6e2df4782353c64 any) {
	__obf_5e89e9b75fd360f4 := __obf_0da8c843dad13139.__obf_5e89e9b75fd360f4
	__obf_cb0de621f162c1ef := reflect2.RTypeOf(__obf_d6e2df4782353c64)
	__obf_801f19702638809c := __obf_0da8c843dad13139.__obf_c52e0bcfad4b8b71.__obf_73853dd89dff1885(__obf_cb0de621f162c1ef)
	if __obf_801f19702638809c == nil {
		__obf_edcded11af6ebc4c := reflect2.TypeOf(__obf_d6e2df4782353c64)
		if __obf_edcded11af6ebc4c == nil || __obf_edcded11af6ebc4c.Kind() != reflect.Ptr {
			__obf_0da8c843dad13139.
				ReportError("ReadVal", "can only unmarshal into pointer")
			return
		}
		__obf_801f19702638809c = __obf_0da8c843dad13139.__obf_c52e0bcfad4b8b71.DecoderOf(__obf_edcded11af6ebc4c)
	}
	__obf_575c04f2b9d91315 := reflect2.PtrOf(__obf_d6e2df4782353c64)
	if __obf_575c04f2b9d91315 == nil {
		__obf_0da8c843dad13139.
			ReportError("ReadVal", "can not read into nil pointer")
		return
	}
	__obf_801f19702638809c.
		Decode(__obf_575c04f2b9d91315, __obf_0da8c843dad13139)
	if __obf_0da8c843dad13139.__obf_5e89e9b75fd360f4 != __obf_5e89e9b75fd360f4 {
		__obf_0da8c843dad13139.
			ReportError("ReadVal", "unexpected mismatched nesting")
		return
	}
}

// WriteVal copy the go interface into underlying JSON, same as json.Marshal
func (__obf_d8f50bcfe87d8b47 *Stream) WriteVal(__obf_35accd096612ebe4 any) {
	if nil == __obf_35accd096612ebe4 {
		__obf_d8f50bcfe87d8b47.
			WriteNil()
		return
	}
	__obf_cb0de621f162c1ef := reflect2.RTypeOf(__obf_35accd096612ebe4)
	__obf_c091c27480fae550 := __obf_d8f50bcfe87d8b47.__obf_c52e0bcfad4b8b71.__obf_627bd91264fcb8f2(__obf_cb0de621f162c1ef)
	if __obf_c091c27480fae550 == nil {
		__obf_edcded11af6ebc4c := reflect2.TypeOf(__obf_35accd096612ebe4)
		__obf_c091c27480fae550 = __obf_d8f50bcfe87d8b47.__obf_c52e0bcfad4b8b71.EncoderOf(__obf_edcded11af6ebc4c)
	}
	__obf_c091c27480fae550.
		Encode(reflect2.PtrOf(__obf_35accd096612ebe4), __obf_d8f50bcfe87d8b47)
}

func (__obf_c52e0bcfad4b8b71 *__obf_30fe5c95cabd69c0) DecoderOf(__obf_edcded11af6ebc4c reflect2.Type) ValDecoder {
	__obf_cb0de621f162c1ef := __obf_edcded11af6ebc4c.RType()
	__obf_801f19702638809c := __obf_c52e0bcfad4b8b71.__obf_73853dd89dff1885(__obf_cb0de621f162c1ef)
	if __obf_801f19702638809c != nil {
		return __obf_801f19702638809c
	}
	__obf_99ec45908bceabdb := &__obf_99ec45908bceabdb{__obf_30fe5c95cabd69c0: __obf_c52e0bcfad4b8b71, __obf_5de9ece8fa3a16e3: "", __obf_ecab71bd3ecfd43c: map[reflect2.Type]ValDecoder{}, __obf_2f5d685c6dbf1e22: map[reflect2.Type]ValEncoder{}}
	__obf_e2840a6d1d1a664b := __obf_edcded11af6ebc4c.(*reflect2.UnsafePtrType)
	__obf_801f19702638809c = __obf_bb14644cc3f030b3(__obf_99ec45908bceabdb, __obf_e2840a6d1d1a664b.Elem())
	__obf_c52e0bcfad4b8b71.__obf_84b44689116bd691(__obf_cb0de621f162c1ef, __obf_801f19702638809c)
	return __obf_801f19702638809c
}

func __obf_bb14644cc3f030b3(__obf_99ec45908bceabdb *__obf_99ec45908bceabdb, __obf_edcded11af6ebc4c reflect2.Type) ValDecoder {
	__obf_801f19702638809c := __obf_62b401f4386d713a(__obf_99ec45908bceabdb, __obf_edcded11af6ebc4c)
	if __obf_801f19702638809c != nil {
		return __obf_801f19702638809c
	}
	__obf_801f19702638809c = __obf_8748f326144f5165(__obf_99ec45908bceabdb, __obf_edcded11af6ebc4c)
	for _, __obf_b8f2f726e65c4d89 := range __obf_0303af627a3db690 {
		__obf_801f19702638809c = __obf_b8f2f726e65c4d89.DecorateDecoder(__obf_edcded11af6ebc4c, __obf_801f19702638809c)
	}
	__obf_801f19702638809c = __obf_99ec45908bceabdb.__obf_1b8392ccffa3bed5.DecorateDecoder(__obf_edcded11af6ebc4c, __obf_801f19702638809c)
	for _, __obf_b8f2f726e65c4d89 := range __obf_99ec45908bceabdb.__obf_3c82f4beb30882eb {
		__obf_801f19702638809c = __obf_b8f2f726e65c4d89.DecorateDecoder(__obf_edcded11af6ebc4c, __obf_801f19702638809c)
	}
	return __obf_801f19702638809c
}

func __obf_8748f326144f5165(__obf_99ec45908bceabdb *__obf_99ec45908bceabdb, __obf_edcded11af6ebc4c reflect2.Type) ValDecoder {
	__obf_801f19702638809c := __obf_99ec45908bceabdb.__obf_ecab71bd3ecfd43c[__obf_edcded11af6ebc4c]
	if __obf_801f19702638809c != nil {
		return __obf_801f19702638809c
	}
	__obf_9cfd5795603f4278 := &__obf_bdefb6d099ae0972{}
	__obf_99ec45908bceabdb.__obf_ecab71bd3ecfd43c[__obf_edcded11af6ebc4c] = __obf_9cfd5795603f4278
	__obf_801f19702638809c = _createDecoderOfType(__obf_99ec45908bceabdb, __obf_edcded11af6ebc4c)
	__obf_9cfd5795603f4278.__obf_801f19702638809c = __obf_801f19702638809c
	return __obf_801f19702638809c
}

func _createDecoderOfType(__obf_99ec45908bceabdb *__obf_99ec45908bceabdb, __obf_edcded11af6ebc4c reflect2.Type) ValDecoder {
	__obf_801f19702638809c := __obf_4360c3949f7a6e97(__obf_edcded11af6ebc4c)
	if __obf_801f19702638809c != nil {
		return __obf_801f19702638809c
	}
	__obf_801f19702638809c = __obf_6bf3f557fcfe51c2(__obf_99ec45908bceabdb, __obf_edcded11af6ebc4c)
	if __obf_801f19702638809c != nil {
		return __obf_801f19702638809c
	}
	__obf_801f19702638809c = __obf_2c53588571656d99(__obf_99ec45908bceabdb, __obf_edcded11af6ebc4c)
	if __obf_801f19702638809c != nil {
		return __obf_801f19702638809c
	}
	__obf_801f19702638809c = __obf_1a6dc69a4bb1ccaa(__obf_99ec45908bceabdb, __obf_edcded11af6ebc4c)
	if __obf_801f19702638809c != nil {
		return __obf_801f19702638809c
	}
	__obf_801f19702638809c = __obf_8cd5deb18d9f5f28(__obf_99ec45908bceabdb, __obf_edcded11af6ebc4c)
	if __obf_801f19702638809c != nil {
		return __obf_801f19702638809c
	}
	switch __obf_edcded11af6ebc4c.Kind() {
	case reflect.Interface:
		__obf_de14529cdcf7529a, __obf_27a5845dd4ee2a4e := __obf_edcded11af6ebc4c.(*reflect2.UnsafeIFaceType)
		if __obf_27a5845dd4ee2a4e {
			return &__obf_8936e3e768812864{__obf_e0ba73c3f13f4455: __obf_de14529cdcf7529a}
		}
		return &__obf_756d9e02226ec8bc{}
	case reflect.Struct:
		return __obf_00d9b9fc4fc9dd14(__obf_99ec45908bceabdb, __obf_edcded11af6ebc4c)
	case reflect.Array:
		return __obf_5597a9555feba901(__obf_99ec45908bceabdb, __obf_edcded11af6ebc4c)
	case reflect.Slice:
		return __obf_978533a1047109a9(__obf_99ec45908bceabdb, __obf_edcded11af6ebc4c)
	case reflect.Map:
		return __obf_a3a1d2b70073655a(__obf_99ec45908bceabdb, __obf_edcded11af6ebc4c)
	case reflect.Ptr:
		return __obf_472d93c11781e39d(__obf_99ec45908bceabdb, __obf_edcded11af6ebc4c)
	default:
		return &__obf_b7a21f1836187893{__obf_ea0680f8b461a85b: fmt.Errorf("%s%s is unsupported type", __obf_99ec45908bceabdb.__obf_5de9ece8fa3a16e3, __obf_edcded11af6ebc4c.String())}
	}
}

func (__obf_c52e0bcfad4b8b71 *__obf_30fe5c95cabd69c0) EncoderOf(__obf_edcded11af6ebc4c reflect2.Type) ValEncoder {
	__obf_cb0de621f162c1ef := __obf_edcded11af6ebc4c.RType()
	__obf_c091c27480fae550 := __obf_c52e0bcfad4b8b71.__obf_627bd91264fcb8f2(__obf_cb0de621f162c1ef)
	if __obf_c091c27480fae550 != nil {
		return __obf_c091c27480fae550
	}
	__obf_99ec45908bceabdb := &__obf_99ec45908bceabdb{__obf_30fe5c95cabd69c0: __obf_c52e0bcfad4b8b71, __obf_5de9ece8fa3a16e3: "", __obf_ecab71bd3ecfd43c: map[reflect2.Type]ValDecoder{}, __obf_2f5d685c6dbf1e22: map[reflect2.Type]ValEncoder{}}
	__obf_c091c27480fae550 = __obf_85f0e4c352da4678(__obf_99ec45908bceabdb, __obf_edcded11af6ebc4c)
	if __obf_edcded11af6ebc4c.LikePtr() {
		__obf_c091c27480fae550 = &__obf_ca1b3b4a6d845336{__obf_c091c27480fae550}
	}
	__obf_c52e0bcfad4b8b71.__obf_fca11eaf15932cf2(__obf_cb0de621f162c1ef, __obf_c091c27480fae550)
	return __obf_c091c27480fae550
}

type __obf_ca1b3b4a6d845336 struct {
	__obf_c091c27480fae550 ValEncoder
}

func (__obf_c091c27480fae550 *__obf_ca1b3b4a6d845336) IsEmpty(__obf_575c04f2b9d91315 unsafe.Pointer) bool {
	return __obf_c091c27480fae550.__obf_c091c27480fae550.IsEmpty(unsafe.Pointer(&__obf_575c04f2b9d91315))
}

func (__obf_c091c27480fae550 *__obf_ca1b3b4a6d845336) Encode(__obf_575c04f2b9d91315 unsafe.Pointer, __obf_d8f50bcfe87d8b47 *Stream) {
	__obf_c091c27480fae550.__obf_c091c27480fae550.
		Encode(unsafe.Pointer(&__obf_575c04f2b9d91315), __obf_d8f50bcfe87d8b47)
}

func __obf_85f0e4c352da4678(__obf_99ec45908bceabdb *__obf_99ec45908bceabdb, __obf_edcded11af6ebc4c reflect2.Type) ValEncoder {
	__obf_c091c27480fae550 := __obf_558310ac59c2b1f5(__obf_99ec45908bceabdb, __obf_edcded11af6ebc4c)
	if __obf_c091c27480fae550 != nil {
		return __obf_c091c27480fae550
	}
	__obf_c091c27480fae550 = __obf_e3dad93963a78f83(__obf_99ec45908bceabdb, __obf_edcded11af6ebc4c)
	for _, __obf_b8f2f726e65c4d89 := range __obf_0303af627a3db690 {
		__obf_c091c27480fae550 = __obf_b8f2f726e65c4d89.DecorateEncoder(__obf_edcded11af6ebc4c, __obf_c091c27480fae550)
	}
	__obf_c091c27480fae550 = __obf_99ec45908bceabdb.__obf_3594368cedb76ac8.DecorateEncoder(__obf_edcded11af6ebc4c, __obf_c091c27480fae550)
	for _, __obf_b8f2f726e65c4d89 := range __obf_99ec45908bceabdb.__obf_3c82f4beb30882eb {
		__obf_c091c27480fae550 = __obf_b8f2f726e65c4d89.DecorateEncoder(__obf_edcded11af6ebc4c, __obf_c091c27480fae550)
	}
	return __obf_c091c27480fae550
}

func __obf_e3dad93963a78f83(__obf_99ec45908bceabdb *__obf_99ec45908bceabdb, __obf_edcded11af6ebc4c reflect2.Type) ValEncoder {
	__obf_c091c27480fae550 := __obf_99ec45908bceabdb.__obf_2f5d685c6dbf1e22[__obf_edcded11af6ebc4c]
	if __obf_c091c27480fae550 != nil {
		return __obf_c091c27480fae550
	}
	__obf_9cfd5795603f4278 := &__obf_f87453715ae2882d{}
	__obf_99ec45908bceabdb.__obf_2f5d685c6dbf1e22[__obf_edcded11af6ebc4c] = __obf_9cfd5795603f4278
	__obf_c091c27480fae550 = _createEncoderOfType(__obf_99ec45908bceabdb, __obf_edcded11af6ebc4c)
	__obf_9cfd5795603f4278.__obf_c091c27480fae550 = __obf_c091c27480fae550
	return __obf_c091c27480fae550
}
func _createEncoderOfType(__obf_99ec45908bceabdb *__obf_99ec45908bceabdb, __obf_edcded11af6ebc4c reflect2.Type) ValEncoder {
	__obf_c091c27480fae550 := __obf_f49908b292adac72(__obf_edcded11af6ebc4c)
	if __obf_c091c27480fae550 != nil {
		return __obf_c091c27480fae550
	}
	__obf_c091c27480fae550 = __obf_84a8ea4f190dbe8e(__obf_99ec45908bceabdb, __obf_edcded11af6ebc4c)
	if __obf_c091c27480fae550 != nil {
		return __obf_c091c27480fae550
	}
	__obf_c091c27480fae550 = __obf_5ddc38919c4ef5f0(__obf_99ec45908bceabdb, __obf_edcded11af6ebc4c)
	if __obf_c091c27480fae550 != nil {
		return __obf_c091c27480fae550
	}
	__obf_c091c27480fae550 = __obf_eab6c90d596283e9(__obf_99ec45908bceabdb, __obf_edcded11af6ebc4c)
	if __obf_c091c27480fae550 != nil {
		return __obf_c091c27480fae550
	}
	__obf_c091c27480fae550 = __obf_1f22c50e61ed6284(__obf_99ec45908bceabdb, __obf_edcded11af6ebc4c)
	if __obf_c091c27480fae550 != nil {
		return __obf_c091c27480fae550
	}
	__obf_105d1cd1a59219cd := __obf_edcded11af6ebc4c.Kind()
	switch __obf_105d1cd1a59219cd {
	case reflect.Interface:
		return &__obf_16a4e7713c271876{__obf_edcded11af6ebc4c}
	case reflect.Struct:
		return __obf_5191b4476a0b2b2f(__obf_99ec45908bceabdb, __obf_edcded11af6ebc4c)
	case reflect.Array:
		return __obf_947e009e73982acf(__obf_99ec45908bceabdb, __obf_edcded11af6ebc4c)
	case reflect.Slice:
		return __obf_1a6fc40f1f862856(__obf_99ec45908bceabdb, __obf_edcded11af6ebc4c)
	case reflect.Map:
		return __obf_8be1ac6303cc573e(__obf_99ec45908bceabdb, __obf_edcded11af6ebc4c)
	case reflect.Ptr:
		return __obf_172fdbd74d552cd4(__obf_99ec45908bceabdb, __obf_edcded11af6ebc4c)
	default:
		return &__obf_531b9739373c7313{__obf_ea0680f8b461a85b: fmt.Errorf("%s%s is unsupported type", __obf_99ec45908bceabdb.__obf_5de9ece8fa3a16e3, __obf_edcded11af6ebc4c.String())}
	}
}

type __obf_b7a21f1836187893 struct {
	__obf_ea0680f8b461a85b error
}

func (__obf_801f19702638809c *__obf_b7a21f1836187893) Decode(__obf_575c04f2b9d91315 unsafe.Pointer, __obf_0da8c843dad13139 *Iterator) {
	if __obf_0da8c843dad13139.WhatIsNext() != NilValue {
		if __obf_0da8c843dad13139.Error == nil {
			__obf_0da8c843dad13139.
				Error = __obf_801f19702638809c.__obf_ea0680f8b461a85b
		}
	} else {
		__obf_0da8c843dad13139.
			Skip()
	}
}

type __obf_531b9739373c7313 struct {
	__obf_ea0680f8b461a85b error
}

func (__obf_c091c27480fae550 *__obf_531b9739373c7313) Encode(__obf_575c04f2b9d91315 unsafe.Pointer, __obf_d8f50bcfe87d8b47 *Stream) {
	if __obf_575c04f2b9d91315 == nil {
		__obf_d8f50bcfe87d8b47.
			WriteNil()
	} else if __obf_d8f50bcfe87d8b47.Error == nil {
		__obf_d8f50bcfe87d8b47.
			Error = __obf_c091c27480fae550.__obf_ea0680f8b461a85b
	}
}

func (__obf_c091c27480fae550 *__obf_531b9739373c7313) IsEmpty(__obf_575c04f2b9d91315 unsafe.Pointer) bool {
	return false
}

type __obf_bdefb6d099ae0972 struct {
	__obf_801f19702638809c ValDecoder
}

func (__obf_801f19702638809c *__obf_bdefb6d099ae0972) Decode(__obf_575c04f2b9d91315 unsafe.Pointer, __obf_0da8c843dad13139 *Iterator) {
	__obf_801f19702638809c.__obf_801f19702638809c.
		Decode(__obf_575c04f2b9d91315, __obf_0da8c843dad13139)
}

type __obf_f87453715ae2882d struct {
	__obf_c091c27480fae550 ValEncoder
}

func (__obf_c091c27480fae550 *__obf_f87453715ae2882d) Encode(__obf_575c04f2b9d91315 unsafe.Pointer, __obf_d8f50bcfe87d8b47 *Stream) {
	__obf_c091c27480fae550.__obf_c091c27480fae550.
		Encode(__obf_575c04f2b9d91315, __obf_d8f50bcfe87d8b47)
}

func (__obf_c091c27480fae550 *__obf_f87453715ae2882d) IsEmpty(__obf_575c04f2b9d91315 unsafe.Pointer) bool {
	return __obf_c091c27480fae550.__obf_c091c27480fae550.IsEmpty(__obf_575c04f2b9d91315)
}
