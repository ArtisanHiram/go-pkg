package __obf_91620b895eeff9ed

import (
	"github.com/modern-go/reflect2"
	"unsafe"
)

func __obf_06a295cf16fb16e3(__obf_2f9c5aed866cce75 *__obf_2f9c5aed866cce75, __obf_29ebd0f2c324f5ea reflect2.Type) ValDecoder {
	__obf_f2fdafeb141957bd := __obf_29ebd0f2c324f5ea.(*reflect2.UnsafePtrType)
	__obf_eef27b4522cbbab1 := __obf_f2fdafeb141957bd.Elem()
	__obf_6fd3f72eb9b5574c := __obf_0b44a7afc1523314(__obf_2f9c5aed866cce75, __obf_eef27b4522cbbab1)
	return &OptionalDecoder{__obf_eef27b4522cbbab1, __obf_6fd3f72eb9b5574c}
}

func __obf_0c129db9075e37b1(__obf_2f9c5aed866cce75 *__obf_2f9c5aed866cce75, __obf_29ebd0f2c324f5ea reflect2.Type) ValEncoder {
	__obf_f2fdafeb141957bd := __obf_29ebd0f2c324f5ea.(*reflect2.UnsafePtrType)
	__obf_eef27b4522cbbab1 := __obf_f2fdafeb141957bd.Elem()
	__obf_11328395cf86586a := __obf_d1233db7a73cc96c(__obf_2f9c5aed866cce75, __obf_eef27b4522cbbab1)
	__obf_96e65a4c8c4f2ce5 := &OptionalEncoder{__obf_11328395cf86586a}
	return __obf_96e65a4c8c4f2ce5
}

type OptionalDecoder struct {
	ValueType    reflect2.Type
	ValueDecoder ValDecoder
}

func (__obf_6fd3f72eb9b5574c *OptionalDecoder) Decode(__obf_2a1474febb02279b unsafe.Pointer, __obf_1bb30e8a74ed8233 *Iterator) {
	if __obf_1bb30e8a74ed8233.ReadNil() {
		*((*unsafe.Pointer)(__obf_2a1474febb02279b)) = nil
	} else {
		if *((*unsafe.Pointer)(__obf_2a1474febb02279b)) == nil {
			__obf_81e4d61da9ac740f := //pointer to null, we have to allocate memory to hold the value
				__obf_6fd3f72eb9b5574c.ValueType.UnsafeNew()
			__obf_6fd3f72eb9b5574c.
				ValueDecoder.Decode(__obf_81e4d61da9ac740f, __obf_1bb30e8a74ed8233)
			*((*unsafe.Pointer)(__obf_2a1474febb02279b)) = __obf_81e4d61da9ac740f
		} else {
			__obf_6fd3f72eb9b5574c.
				//reuse existing instance
				ValueDecoder.Decode(*((*unsafe.Pointer)(__obf_2a1474febb02279b)), __obf_1bb30e8a74ed8233)
		}
	}
}

type __obf_7a31d659954d3208 struct {
	__obf_f47701e28c5effac reflect2. // only to deference a pointer
				Type
	__obf_4f68c814fce49229 ValDecoder
}

func (__obf_6fd3f72eb9b5574c *__obf_7a31d659954d3208) Decode(__obf_2a1474febb02279b unsafe.Pointer, __obf_1bb30e8a74ed8233 *Iterator) {
	if *((*unsafe.Pointer)(__obf_2a1474febb02279b)) == nil {
		__obf_81e4d61da9ac740f := //pointer to null, we have to allocate memory to hold the value
			__obf_6fd3f72eb9b5574c.__obf_f47701e28c5effac.UnsafeNew()
		__obf_6fd3f72eb9b5574c.__obf_4f68c814fce49229.
			Decode(__obf_81e4d61da9ac740f, __obf_1bb30e8a74ed8233)
		*((*unsafe.Pointer)(__obf_2a1474febb02279b)) = __obf_81e4d61da9ac740f
	} else {
		__obf_6fd3f72eb9b5574c.
			//reuse existing instance
			__obf_4f68c814fce49229.
			Decode(*((*unsafe.Pointer)(__obf_2a1474febb02279b)), __obf_1bb30e8a74ed8233)
	}
}

type OptionalEncoder struct {
	ValueEncoder ValEncoder
}

func (__obf_96e65a4c8c4f2ce5 *OptionalEncoder) Encode(__obf_2a1474febb02279b unsafe.Pointer, __obf_850a7457bb739a32 *Stream) {
	if *((*unsafe.Pointer)(__obf_2a1474febb02279b)) == nil {
		__obf_850a7457bb739a32.
			WriteNil()
	} else {
		__obf_96e65a4c8c4f2ce5.
			ValueEncoder.Encode(*((*unsafe.Pointer)(__obf_2a1474febb02279b)), __obf_850a7457bb739a32)
	}
}

func (__obf_96e65a4c8c4f2ce5 *OptionalEncoder) IsEmpty(__obf_2a1474febb02279b unsafe.Pointer) bool {
	return *((*unsafe.Pointer)(__obf_2a1474febb02279b)) == nil
}

type __obf_9146060c74844107 struct {
	ValueEncoder ValEncoder
}

func (__obf_96e65a4c8c4f2ce5 *__obf_9146060c74844107) Encode(__obf_2a1474febb02279b unsafe.Pointer, __obf_850a7457bb739a32 *Stream) {
	if *((*unsafe.Pointer)(__obf_2a1474febb02279b)) == nil {
		__obf_850a7457bb739a32.
			WriteNil()
	} else {
		__obf_96e65a4c8c4f2ce5.
			ValueEncoder.Encode(*((*unsafe.Pointer)(__obf_2a1474febb02279b)), __obf_850a7457bb739a32)
	}
}

func (__obf_96e65a4c8c4f2ce5 *__obf_9146060c74844107) IsEmpty(__obf_2a1474febb02279b unsafe.Pointer) bool {
	__obf_cdc115f93c6164b0 := *((*unsafe.Pointer)(__obf_2a1474febb02279b))
	if __obf_cdc115f93c6164b0 == nil {
		return true
	}
	return __obf_96e65a4c8c4f2ce5.ValueEncoder.IsEmpty(__obf_cdc115f93c6164b0)
}

func (__obf_96e65a4c8c4f2ce5 *__obf_9146060c74844107) IsEmbeddedPtrNil(__obf_2a1474febb02279b unsafe.Pointer) bool {
	__obf_8e4248b0a8cb631c := *((*unsafe.Pointer)(__obf_2a1474febb02279b))
	if __obf_8e4248b0a8cb631c == nil {
		return true
	}
	__obf_c96270edda65ae80, __obf_9fb2af2f6f03e773 := __obf_96e65a4c8c4f2ce5.ValueEncoder.(IsEmbeddedPtrNil)
	if !__obf_9fb2af2f6f03e773 {
		return false
	}
	__obf_43dac9f05df8fcc6 := unsafe.Pointer(__obf_8e4248b0a8cb631c)
	return __obf_c96270edda65ae80.IsEmbeddedPtrNil(__obf_43dac9f05df8fcc6)
}

type __obf_9d9e26186edb63a5 struct {
	__obf_96e65a4c8c4f2ce5 ValEncoder
}

func (__obf_96e65a4c8c4f2ce5 *__obf_9d9e26186edb63a5) Encode(__obf_2a1474febb02279b unsafe.Pointer, __obf_850a7457bb739a32 *Stream) {
	__obf_96e65a4c8c4f2ce5.__obf_96e65a4c8c4f2ce5.
		Encode(unsafe.Pointer(&__obf_2a1474febb02279b), __obf_850a7457bb739a32)
}

func (__obf_96e65a4c8c4f2ce5 *__obf_9d9e26186edb63a5) IsEmpty(__obf_2a1474febb02279b unsafe.Pointer) bool {
	return __obf_96e65a4c8c4f2ce5.__obf_96e65a4c8c4f2ce5.IsEmpty(unsafe.Pointer(&__obf_2a1474febb02279b))
}

type __obf_a82dfb83d40f1b70 struct {
	__obf_6fd3f72eb9b5574c ValDecoder
}

func (__obf_6fd3f72eb9b5574c *__obf_a82dfb83d40f1b70) Decode(__obf_2a1474febb02279b unsafe.Pointer, __obf_1bb30e8a74ed8233 *Iterator) {
	__obf_6fd3f72eb9b5574c.__obf_6fd3f72eb9b5574c.
		Decode(unsafe.Pointer(&__obf_2a1474febb02279b), __obf_1bb30e8a74ed8233)
}
