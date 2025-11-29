package __obf_91620b895eeff9ed

import (
	"reflect"
	"unsafe"

	"github.com/modern-go/reflect2"
)

type __obf_bcab05fd1729d9fa struct {
	__obf_5ea8ee327a6f7e0d reflect2.Type
}

func (__obf_96e65a4c8c4f2ce5 *__obf_bcab05fd1729d9fa) Encode(__obf_2a1474febb02279b unsafe.Pointer, __obf_850a7457bb739a32 *Stream) {
	__obf_35136ef2ac0f94e4 := __obf_96e65a4c8c4f2ce5.__obf_5ea8ee327a6f7e0d.UnsafeIndirect(__obf_2a1474febb02279b)
	__obf_850a7457bb739a32.
		WriteVal(__obf_35136ef2ac0f94e4)
}

func (__obf_96e65a4c8c4f2ce5 *__obf_bcab05fd1729d9fa) IsEmpty(__obf_2a1474febb02279b unsafe.Pointer) bool {
	return __obf_96e65a4c8c4f2ce5.__obf_5ea8ee327a6f7e0d.UnsafeIndirect(__obf_2a1474febb02279b) == nil
}

type __obf_5e2fcb326d03b9c1 struct {
}

func (__obf_6fd3f72eb9b5574c *__obf_5e2fcb326d03b9c1) Decode(__obf_2a1474febb02279b unsafe.Pointer, __obf_1bb30e8a74ed8233 *Iterator) {
	__obf_f46fc633d7e54b23 := (*any)(__obf_2a1474febb02279b)
	__obf_35136ef2ac0f94e4 := *__obf_f46fc633d7e54b23
	if __obf_35136ef2ac0f94e4 == nil {
		*__obf_f46fc633d7e54b23 = __obf_1bb30e8a74ed8233.Read()
		return
	}
	__obf_29ebd0f2c324f5ea := reflect2.TypeOf(__obf_35136ef2ac0f94e4)
	if __obf_29ebd0f2c324f5ea.Kind() != reflect.Ptr {
		*__obf_f46fc633d7e54b23 = __obf_1bb30e8a74ed8233.Read()
		return
	}
	__obf_f2fdafeb141957bd := __obf_29ebd0f2c324f5ea.(*reflect2.UnsafePtrType)
	__obf_d163fc3ca05fb54f := __obf_f2fdafeb141957bd.Elem()
	if __obf_1bb30e8a74ed8233.WhatIsNext() == NilValue {
		if __obf_d163fc3ca05fb54f.Kind() != reflect.Ptr {
			__obf_1bb30e8a74ed8233.__obf_94a8e39928c8440c('n', 'u', 'l', 'l')
			*__obf_f46fc633d7e54b23 = nil
			return
		}
	}
	if reflect2.IsNil(__obf_35136ef2ac0f94e4) {
		__obf_35136ef2ac0f94e4 := __obf_d163fc3ca05fb54f.New()
		__obf_1bb30e8a74ed8233.
			ReadVal(__obf_35136ef2ac0f94e4)
		*__obf_f46fc633d7e54b23 = __obf_35136ef2ac0f94e4
		return
	}
	__obf_1bb30e8a74ed8233.
		ReadVal(__obf_35136ef2ac0f94e4)
}

type __obf_36b06a4a164de621 struct {
	__obf_5ea8ee327a6f7e0d *reflect2.UnsafeIFaceType
}

func (__obf_6fd3f72eb9b5574c *__obf_36b06a4a164de621) Decode(__obf_2a1474febb02279b unsafe.Pointer, __obf_1bb30e8a74ed8233 *Iterator) {
	if __obf_1bb30e8a74ed8233.ReadNil() {
		__obf_6fd3f72eb9b5574c.__obf_5ea8ee327a6f7e0d.
			UnsafeSet(__obf_2a1474febb02279b, __obf_6fd3f72eb9b5574c.__obf_5ea8ee327a6f7e0d.UnsafeNew())
		return
	}
	__obf_35136ef2ac0f94e4 := __obf_6fd3f72eb9b5574c.__obf_5ea8ee327a6f7e0d.UnsafeIndirect(__obf_2a1474febb02279b)
	if reflect2.IsNil(__obf_35136ef2ac0f94e4) {
		__obf_1bb30e8a74ed8233.
			ReportError("decode non empty interface", "can not unmarshal into nil")
		return
	}
	__obf_1bb30e8a74ed8233.
		ReadVal(__obf_35136ef2ac0f94e4)
}
