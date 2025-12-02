package __obf_c7b79b12b33d8238

import (
	"reflect"
	"unsafe"

	"github.com/modern-go/reflect2"
)

type __obf_16a4e7713c271876 struct {
	__obf_e0ba73c3f13f4455 reflect2.Type
}

func (__obf_c091c27480fae550 *__obf_16a4e7713c271876) Encode(__obf_575c04f2b9d91315 unsafe.Pointer, __obf_d8f50bcfe87d8b47 *Stream) {
	__obf_d6e2df4782353c64 := __obf_c091c27480fae550.__obf_e0ba73c3f13f4455.UnsafeIndirect(__obf_575c04f2b9d91315)
	__obf_d8f50bcfe87d8b47.
		WriteVal(__obf_d6e2df4782353c64)
}

func (__obf_c091c27480fae550 *__obf_16a4e7713c271876) IsEmpty(__obf_575c04f2b9d91315 unsafe.Pointer) bool {
	return __obf_c091c27480fae550.__obf_e0ba73c3f13f4455.UnsafeIndirect(__obf_575c04f2b9d91315) == nil
}

type __obf_756d9e02226ec8bc struct {
}

func (__obf_801f19702638809c *__obf_756d9e02226ec8bc) Decode(__obf_575c04f2b9d91315 unsafe.Pointer, __obf_0da8c843dad13139 *Iterator) {
	__obf_673edb05fb6950d1 := (*any)(__obf_575c04f2b9d91315)
	__obf_d6e2df4782353c64 := *__obf_673edb05fb6950d1
	if __obf_d6e2df4782353c64 == nil {
		*__obf_673edb05fb6950d1 = __obf_0da8c843dad13139.Read()
		return
	}
	__obf_edcded11af6ebc4c := reflect2.TypeOf(__obf_d6e2df4782353c64)
	if __obf_edcded11af6ebc4c.Kind() != reflect.Ptr {
		*__obf_673edb05fb6950d1 = __obf_0da8c843dad13139.Read()
		return
	}
	__obf_e2840a6d1d1a664b := __obf_edcded11af6ebc4c.(*reflect2.UnsafePtrType)
	__obf_12e01cd53fba7222 := __obf_e2840a6d1d1a664b.Elem()
	if __obf_0da8c843dad13139.WhatIsNext() == NilValue {
		if __obf_12e01cd53fba7222.Kind() != reflect.Ptr {
			__obf_0da8c843dad13139.__obf_62908d9424a8486b('n', 'u', 'l', 'l')
			*__obf_673edb05fb6950d1 = nil
			return
		}
	}
	if reflect2.IsNil(__obf_d6e2df4782353c64) {
		__obf_d6e2df4782353c64 := __obf_12e01cd53fba7222.New()
		__obf_0da8c843dad13139.
			ReadVal(__obf_d6e2df4782353c64)
		*__obf_673edb05fb6950d1 = __obf_d6e2df4782353c64
		return
	}
	__obf_0da8c843dad13139.
		ReadVal(__obf_d6e2df4782353c64)
}

type __obf_8936e3e768812864 struct {
	__obf_e0ba73c3f13f4455 *reflect2.UnsafeIFaceType
}

func (__obf_801f19702638809c *__obf_8936e3e768812864) Decode(__obf_575c04f2b9d91315 unsafe.Pointer, __obf_0da8c843dad13139 *Iterator) {
	if __obf_0da8c843dad13139.ReadNil() {
		__obf_801f19702638809c.__obf_e0ba73c3f13f4455.
			UnsafeSet(__obf_575c04f2b9d91315, __obf_801f19702638809c.__obf_e0ba73c3f13f4455.UnsafeNew())
		return
	}
	__obf_d6e2df4782353c64 := __obf_801f19702638809c.__obf_e0ba73c3f13f4455.UnsafeIndirect(__obf_575c04f2b9d91315)
	if reflect2.IsNil(__obf_d6e2df4782353c64) {
		__obf_0da8c843dad13139.
			ReportError("decode non empty interface", "can not unmarshal into nil")
		return
	}
	__obf_0da8c843dad13139.
		ReadVal(__obf_d6e2df4782353c64)
}
