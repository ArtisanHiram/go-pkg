package __obf_c7b79b12b33d8238

import (
	"github.com/modern-go/reflect2"
	"unsafe"
)

func __obf_472d93c11781e39d(__obf_99ec45908bceabdb *__obf_99ec45908bceabdb, __obf_edcded11af6ebc4c reflect2.Type) ValDecoder {
	__obf_e2840a6d1d1a664b := __obf_edcded11af6ebc4c.(*reflect2.UnsafePtrType)
	__obf_f5ccfed08a3cf863 := __obf_e2840a6d1d1a664b.Elem()
	__obf_801f19702638809c := __obf_bb14644cc3f030b3(__obf_99ec45908bceabdb, __obf_f5ccfed08a3cf863)
	return &OptionalDecoder{__obf_f5ccfed08a3cf863, __obf_801f19702638809c}
}

func __obf_172fdbd74d552cd4(__obf_99ec45908bceabdb *__obf_99ec45908bceabdb, __obf_edcded11af6ebc4c reflect2.Type) ValEncoder {
	__obf_e2840a6d1d1a664b := __obf_edcded11af6ebc4c.(*reflect2.UnsafePtrType)
	__obf_f5ccfed08a3cf863 := __obf_e2840a6d1d1a664b.Elem()
	__obf_e6894379459b04e8 := __obf_85f0e4c352da4678(__obf_99ec45908bceabdb, __obf_f5ccfed08a3cf863)
	__obf_c091c27480fae550 := &OptionalEncoder{__obf_e6894379459b04e8}
	return __obf_c091c27480fae550
}

type OptionalDecoder struct {
	ValueType    reflect2.Type
	ValueDecoder ValDecoder
}

func (__obf_801f19702638809c *OptionalDecoder) Decode(__obf_575c04f2b9d91315 unsafe.Pointer, __obf_0da8c843dad13139 *Iterator) {
	if __obf_0da8c843dad13139.ReadNil() {
		*((*unsafe.Pointer)(__obf_575c04f2b9d91315)) = nil
	} else {
		if *((*unsafe.Pointer)(__obf_575c04f2b9d91315)) == nil {
			__obf_c7e531fc9b564031 := //pointer to null, we have to allocate memory to hold the value
				__obf_801f19702638809c.ValueType.UnsafeNew()
			__obf_801f19702638809c.
				ValueDecoder.Decode(__obf_c7e531fc9b564031, __obf_0da8c843dad13139)
			*((*unsafe.Pointer)(__obf_575c04f2b9d91315)) = __obf_c7e531fc9b564031
		} else {
			__obf_801f19702638809c.
				//reuse existing instance
				ValueDecoder.Decode(*((*unsafe.Pointer)(__obf_575c04f2b9d91315)), __obf_0da8c843dad13139)
		}
	}
}

type __obf_af9e91113c760212 struct {
	__obf_2cd064ad0a3f5c35 reflect2. // only to deference a pointer
				Type
	__obf_80388639278c34a2 ValDecoder
}

func (__obf_801f19702638809c *__obf_af9e91113c760212) Decode(__obf_575c04f2b9d91315 unsafe.Pointer, __obf_0da8c843dad13139 *Iterator) {
	if *((*unsafe.Pointer)(__obf_575c04f2b9d91315)) == nil {
		__obf_c7e531fc9b564031 := //pointer to null, we have to allocate memory to hold the value
			__obf_801f19702638809c.__obf_2cd064ad0a3f5c35.UnsafeNew()
		__obf_801f19702638809c.__obf_80388639278c34a2.
			Decode(__obf_c7e531fc9b564031, __obf_0da8c843dad13139)
		*((*unsafe.Pointer)(__obf_575c04f2b9d91315)) = __obf_c7e531fc9b564031
	} else {
		__obf_801f19702638809c.
			//reuse existing instance
			__obf_80388639278c34a2.
			Decode(*((*unsafe.Pointer)(__obf_575c04f2b9d91315)), __obf_0da8c843dad13139)
	}
}

type OptionalEncoder struct {
	ValueEncoder ValEncoder
}

func (__obf_c091c27480fae550 *OptionalEncoder) Encode(__obf_575c04f2b9d91315 unsafe.Pointer, __obf_d8f50bcfe87d8b47 *Stream) {
	if *((*unsafe.Pointer)(__obf_575c04f2b9d91315)) == nil {
		__obf_d8f50bcfe87d8b47.
			WriteNil()
	} else {
		__obf_c091c27480fae550.
			ValueEncoder.Encode(*((*unsafe.Pointer)(__obf_575c04f2b9d91315)), __obf_d8f50bcfe87d8b47)
	}
}

func (__obf_c091c27480fae550 *OptionalEncoder) IsEmpty(__obf_575c04f2b9d91315 unsafe.Pointer) bool {
	return *((*unsafe.Pointer)(__obf_575c04f2b9d91315)) == nil
}

type __obf_a76744e81695632d struct {
	ValueEncoder ValEncoder
}

func (__obf_c091c27480fae550 *__obf_a76744e81695632d) Encode(__obf_575c04f2b9d91315 unsafe.Pointer, __obf_d8f50bcfe87d8b47 *Stream) {
	if *((*unsafe.Pointer)(__obf_575c04f2b9d91315)) == nil {
		__obf_d8f50bcfe87d8b47.
			WriteNil()
	} else {
		__obf_c091c27480fae550.
			ValueEncoder.Encode(*((*unsafe.Pointer)(__obf_575c04f2b9d91315)), __obf_d8f50bcfe87d8b47)
	}
}

func (__obf_c091c27480fae550 *__obf_a76744e81695632d) IsEmpty(__obf_575c04f2b9d91315 unsafe.Pointer) bool {
	__obf_c192e74422df4ed6 := *((*unsafe.Pointer)(__obf_575c04f2b9d91315))
	if __obf_c192e74422df4ed6 == nil {
		return true
	}
	return __obf_c091c27480fae550.ValueEncoder.IsEmpty(__obf_c192e74422df4ed6)
}

func (__obf_c091c27480fae550 *__obf_a76744e81695632d) IsEmbeddedPtrNil(__obf_575c04f2b9d91315 unsafe.Pointer) bool {
	__obf_a722610daca6a84a := *((*unsafe.Pointer)(__obf_575c04f2b9d91315))
	if __obf_a722610daca6a84a == nil {
		return true
	}
	__obf_bdbbea1ea376ff9b, __obf_a7ca508d9e901782 := __obf_c091c27480fae550.ValueEncoder.(IsEmbeddedPtrNil)
	if !__obf_a7ca508d9e901782 {
		return false
	}
	__obf_c34985751cfe0592 := unsafe.Pointer(__obf_a722610daca6a84a)
	return __obf_bdbbea1ea376ff9b.IsEmbeddedPtrNil(__obf_c34985751cfe0592)
}

type __obf_f992d744271af46a struct {
	__obf_c091c27480fae550 ValEncoder
}

func (__obf_c091c27480fae550 *__obf_f992d744271af46a) Encode(__obf_575c04f2b9d91315 unsafe.Pointer, __obf_d8f50bcfe87d8b47 *Stream) {
	__obf_c091c27480fae550.__obf_c091c27480fae550.
		Encode(unsafe.Pointer(&__obf_575c04f2b9d91315), __obf_d8f50bcfe87d8b47)
}

func (__obf_c091c27480fae550 *__obf_f992d744271af46a) IsEmpty(__obf_575c04f2b9d91315 unsafe.Pointer) bool {
	return __obf_c091c27480fae550.__obf_c091c27480fae550.IsEmpty(unsafe.Pointer(&__obf_575c04f2b9d91315))
}

type __obf_488d5d64f097f8dc struct {
	__obf_801f19702638809c ValDecoder
}

func (__obf_801f19702638809c *__obf_488d5d64f097f8dc) Decode(__obf_575c04f2b9d91315 unsafe.Pointer, __obf_0da8c843dad13139 *Iterator) {
	__obf_801f19702638809c.__obf_801f19702638809c.
		Decode(unsafe.Pointer(&__obf_575c04f2b9d91315), __obf_0da8c843dad13139)
}
