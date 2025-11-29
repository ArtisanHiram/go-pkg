package __obf_91620b895eeff9ed

import (
	"fmt"
	"github.com/modern-go/reflect2"
	"io"
	"unsafe"
)

func __obf_c1de99114b629ad2(__obf_2f9c5aed866cce75 *__obf_2f9c5aed866cce75, __obf_29ebd0f2c324f5ea reflect2.Type) ValDecoder {
	__obf_7fc2ead51ea6d78f := __obf_29ebd0f2c324f5ea.(*reflect2.UnsafeArrayType)
	__obf_6fd3f72eb9b5574c := __obf_0b44a7afc1523314(__obf_2f9c5aed866cce75.append("[arrayElem]"), __obf_7fc2ead51ea6d78f.Elem())
	return &__obf_05f0f29146ccfa6b{__obf_7fc2ead51ea6d78f, __obf_6fd3f72eb9b5574c}
}

func __obf_77d3a743e2d92123(__obf_2f9c5aed866cce75 *__obf_2f9c5aed866cce75, __obf_29ebd0f2c324f5ea reflect2.Type) ValEncoder {
	__obf_7fc2ead51ea6d78f := __obf_29ebd0f2c324f5ea.(*reflect2.UnsafeArrayType)
	if __obf_7fc2ead51ea6d78f.Len() == 0 {
		return __obf_62549f97568c8f48{}
	}
	__obf_96e65a4c8c4f2ce5 := __obf_d1233db7a73cc96c(__obf_2f9c5aed866cce75.append("[arrayElem]"), __obf_7fc2ead51ea6d78f.Elem())
	return &__obf_8dfe2d1964f42840{__obf_7fc2ead51ea6d78f, __obf_96e65a4c8c4f2ce5}
}

type __obf_62549f97568c8f48 struct{}

func (__obf_96e65a4c8c4f2ce5 __obf_62549f97568c8f48) Encode(__obf_2a1474febb02279b unsafe.Pointer, __obf_850a7457bb739a32 *Stream) {
	__obf_850a7457bb739a32.
		WriteEmptyArray()
}

func (__obf_96e65a4c8c4f2ce5 __obf_62549f97568c8f48) IsEmpty(__obf_2a1474febb02279b unsafe.Pointer) bool {
	return true
}

type __obf_8dfe2d1964f42840 struct {
	__obf_7fc2ead51ea6d78f *reflect2.UnsafeArrayType
	__obf_11328395cf86586a ValEncoder
}

func (__obf_96e65a4c8c4f2ce5 *__obf_8dfe2d1964f42840) Encode(__obf_2a1474febb02279b unsafe.Pointer, __obf_850a7457bb739a32 *Stream) {
	__obf_850a7457bb739a32.
		WriteArrayStart()
	__obf_2b0afef38cdec72a := unsafe.Pointer(__obf_2a1474febb02279b)
	__obf_96e65a4c8c4f2ce5.__obf_11328395cf86586a.
		Encode(__obf_2b0afef38cdec72a, __obf_850a7457bb739a32)
	for __obf_5aa5c8829b97f182 := 1; __obf_5aa5c8829b97f182 < __obf_96e65a4c8c4f2ce5.__obf_7fc2ead51ea6d78f.Len(); __obf_5aa5c8829b97f182++ {
		__obf_850a7457bb739a32.
			WriteMore()
		__obf_2b0afef38cdec72a = __obf_96e65a4c8c4f2ce5.__obf_7fc2ead51ea6d78f.UnsafeGetIndex(__obf_2a1474febb02279b, __obf_5aa5c8829b97f182)
		__obf_96e65a4c8c4f2ce5.__obf_11328395cf86586a.
			Encode(__obf_2b0afef38cdec72a, __obf_850a7457bb739a32)
	}
	__obf_850a7457bb739a32.
		WriteArrayEnd()
	if __obf_850a7457bb739a32.Error != nil && __obf_850a7457bb739a32.Error != io.EOF {
		__obf_850a7457bb739a32.
			Error = fmt.Errorf("%v: %s", __obf_96e65a4c8c4f2ce5.__obf_7fc2ead51ea6d78f, __obf_850a7457bb739a32.Error.Error())
	}
}

func (__obf_96e65a4c8c4f2ce5 *__obf_8dfe2d1964f42840) IsEmpty(__obf_2a1474febb02279b unsafe.Pointer) bool {
	return false
}

type __obf_05f0f29146ccfa6b struct {
	__obf_7fc2ead51ea6d78f *reflect2.UnsafeArrayType
	__obf_0110b2c027051777 ValDecoder
}

func (__obf_6fd3f72eb9b5574c *__obf_05f0f29146ccfa6b) Decode(__obf_2a1474febb02279b unsafe.Pointer, __obf_1bb30e8a74ed8233 *Iterator) {
	__obf_6fd3f72eb9b5574c.__obf_f5d689d3cfbfa27f(__obf_2a1474febb02279b, __obf_1bb30e8a74ed8233)
	if __obf_1bb30e8a74ed8233.Error != nil && __obf_1bb30e8a74ed8233.Error != io.EOF {
		__obf_1bb30e8a74ed8233.
			Error = fmt.Errorf("%v: %s", __obf_6fd3f72eb9b5574c.__obf_7fc2ead51ea6d78f, __obf_1bb30e8a74ed8233.Error.Error())
	}
}

func (__obf_6fd3f72eb9b5574c *__obf_05f0f29146ccfa6b) __obf_f5d689d3cfbfa27f(__obf_2a1474febb02279b unsafe.Pointer, __obf_1bb30e8a74ed8233 *Iterator) {
	__obf_f16b4157911bc9af := __obf_1bb30e8a74ed8233.__obf_684faa48ae8c5049()
	__obf_7fc2ead51ea6d78f := __obf_6fd3f72eb9b5574c.__obf_7fc2ead51ea6d78f
	if __obf_f16b4157911bc9af == 'n' {
		__obf_1bb30e8a74ed8233.__obf_3e1d2ad9a54f0d22('u', 'l', 'l')
		return
	}
	if __obf_f16b4157911bc9af != '[' {
		__obf_1bb30e8a74ed8233.
			ReportError("decode array", "expect [ or n, but found "+string([]byte{__obf_f16b4157911bc9af}))
		return
	}
	__obf_f16b4157911bc9af = __obf_1bb30e8a74ed8233.__obf_684faa48ae8c5049()
	if __obf_f16b4157911bc9af == ']' {
		return
	}
	__obf_1bb30e8a74ed8233.__obf_a163df67f9bb1c4b()
	__obf_2b0afef38cdec72a := __obf_7fc2ead51ea6d78f.UnsafeGetIndex(__obf_2a1474febb02279b, 0)
	__obf_6fd3f72eb9b5574c.__obf_0110b2c027051777.
		Decode(__obf_2b0afef38cdec72a, __obf_1bb30e8a74ed8233)
	__obf_e3fec96653416447 := 1
	for __obf_f16b4157911bc9af = __obf_1bb30e8a74ed8233.__obf_684faa48ae8c5049(); __obf_f16b4157911bc9af == ','; __obf_f16b4157911bc9af = __obf_1bb30e8a74ed8233.__obf_684faa48ae8c5049() {
		if __obf_e3fec96653416447 >= __obf_7fc2ead51ea6d78f.Len() {
			__obf_1bb30e8a74ed8233.
				Skip()
			continue
		}
		__obf_b0acd231dbd22334 := __obf_e3fec96653416447
		__obf_e3fec96653416447 += 1
		__obf_2b0afef38cdec72a = __obf_7fc2ead51ea6d78f.UnsafeGetIndex(__obf_2a1474febb02279b, __obf_b0acd231dbd22334)
		__obf_6fd3f72eb9b5574c.__obf_0110b2c027051777.
			Decode(__obf_2b0afef38cdec72a, __obf_1bb30e8a74ed8233)
	}
	if __obf_f16b4157911bc9af != ']' {
		__obf_1bb30e8a74ed8233.
			ReportError("decode array", "expect ], but found "+string([]byte{__obf_f16b4157911bc9af}))
		return
	}
}
