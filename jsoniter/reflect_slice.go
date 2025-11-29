package __obf_91620b895eeff9ed

import (
	"fmt"
	"github.com/modern-go/reflect2"
	"io"
	"unsafe"
)

func __obf_5f6f4a91bd6a3b9b(__obf_2f9c5aed866cce75 *__obf_2f9c5aed866cce75, __obf_29ebd0f2c324f5ea reflect2.Type) ValDecoder {
	__obf_56bad1827144ee34 := __obf_29ebd0f2c324f5ea.(*reflect2.UnsafeSliceType)
	__obf_6fd3f72eb9b5574c := __obf_0b44a7afc1523314(__obf_2f9c5aed866cce75.append("[sliceElem]"), __obf_56bad1827144ee34.Elem())
	return &__obf_786f7abbe25edc1b{__obf_56bad1827144ee34, __obf_6fd3f72eb9b5574c}
}

func __obf_9df0b29f9ca5de5b(__obf_2f9c5aed866cce75 *__obf_2f9c5aed866cce75, __obf_29ebd0f2c324f5ea reflect2.Type) ValEncoder {
	__obf_56bad1827144ee34 := __obf_29ebd0f2c324f5ea.(*reflect2.UnsafeSliceType)
	__obf_96e65a4c8c4f2ce5 := __obf_d1233db7a73cc96c(__obf_2f9c5aed866cce75.append("[sliceElem]"), __obf_56bad1827144ee34.Elem())
	return &__obf_a70647718c3a51db{__obf_56bad1827144ee34, __obf_96e65a4c8c4f2ce5}
}

type __obf_a70647718c3a51db struct {
	__obf_56bad1827144ee34 *reflect2.UnsafeSliceType
	__obf_11328395cf86586a ValEncoder
}

func (__obf_96e65a4c8c4f2ce5 *__obf_a70647718c3a51db) Encode(__obf_2a1474febb02279b unsafe.Pointer, __obf_850a7457bb739a32 *Stream) {
	if __obf_96e65a4c8c4f2ce5.__obf_56bad1827144ee34.UnsafeIsNil(__obf_2a1474febb02279b) {
		__obf_850a7457bb739a32.
			WriteNil()
		return
	}
	__obf_e3fec96653416447 := __obf_96e65a4c8c4f2ce5.__obf_56bad1827144ee34.UnsafeLengthOf(__obf_2a1474febb02279b)
	if __obf_e3fec96653416447 == 0 {
		__obf_850a7457bb739a32.
			WriteEmptyArray()
		return
	}
	__obf_850a7457bb739a32.
		WriteArrayStart()
	__obf_96e65a4c8c4f2ce5.__obf_11328395cf86586a.
		Encode(__obf_96e65a4c8c4f2ce5.__obf_56bad1827144ee34.UnsafeGetIndex(__obf_2a1474febb02279b, 0), __obf_850a7457bb739a32)
	for __obf_5aa5c8829b97f182 := 1; __obf_5aa5c8829b97f182 < __obf_e3fec96653416447; __obf_5aa5c8829b97f182++ {
		__obf_850a7457bb739a32.
			WriteMore()
		__obf_2b0afef38cdec72a := __obf_96e65a4c8c4f2ce5.__obf_56bad1827144ee34.UnsafeGetIndex(__obf_2a1474febb02279b, __obf_5aa5c8829b97f182)
		__obf_96e65a4c8c4f2ce5.__obf_11328395cf86586a.
			Encode(__obf_2b0afef38cdec72a, __obf_850a7457bb739a32)
	}
	__obf_850a7457bb739a32.
		WriteArrayEnd()
	if __obf_850a7457bb739a32.Error != nil && __obf_850a7457bb739a32.Error != io.EOF {
		__obf_850a7457bb739a32.
			Error = fmt.Errorf("%v: %s", __obf_96e65a4c8c4f2ce5.__obf_56bad1827144ee34, __obf_850a7457bb739a32.Error.Error())
	}
}

func (__obf_96e65a4c8c4f2ce5 *__obf_a70647718c3a51db) IsEmpty(__obf_2a1474febb02279b unsafe.Pointer) bool {
	return __obf_96e65a4c8c4f2ce5.__obf_56bad1827144ee34.UnsafeLengthOf(__obf_2a1474febb02279b) == 0
}

type __obf_786f7abbe25edc1b struct {
	__obf_56bad1827144ee34 *reflect2.UnsafeSliceType
	__obf_0110b2c027051777 ValDecoder
}

func (__obf_6fd3f72eb9b5574c *__obf_786f7abbe25edc1b) Decode(__obf_2a1474febb02279b unsafe.Pointer, __obf_1bb30e8a74ed8233 *Iterator) {
	__obf_6fd3f72eb9b5574c.__obf_f5d689d3cfbfa27f(__obf_2a1474febb02279b, __obf_1bb30e8a74ed8233)
	if __obf_1bb30e8a74ed8233.Error != nil && __obf_1bb30e8a74ed8233.Error != io.EOF {
		__obf_1bb30e8a74ed8233.
			Error = fmt.Errorf("%v: %s", __obf_6fd3f72eb9b5574c.__obf_56bad1827144ee34, __obf_1bb30e8a74ed8233.Error.Error())
	}
}

func (__obf_6fd3f72eb9b5574c *__obf_786f7abbe25edc1b) __obf_f5d689d3cfbfa27f(__obf_2a1474febb02279b unsafe.Pointer, __obf_1bb30e8a74ed8233 *Iterator) {
	__obf_f16b4157911bc9af := __obf_1bb30e8a74ed8233.__obf_684faa48ae8c5049()
	__obf_56bad1827144ee34 := __obf_6fd3f72eb9b5574c.__obf_56bad1827144ee34
	if __obf_f16b4157911bc9af == 'n' {
		__obf_1bb30e8a74ed8233.__obf_3e1d2ad9a54f0d22('u', 'l', 'l')
		__obf_56bad1827144ee34.
			UnsafeSetNil(__obf_2a1474febb02279b)
		return
	}
	if __obf_f16b4157911bc9af != '[' {
		__obf_1bb30e8a74ed8233.
			ReportError("decode slice", "expect [ or n, but found "+string([]byte{__obf_f16b4157911bc9af}))
		return
	}
	__obf_f16b4157911bc9af = __obf_1bb30e8a74ed8233.__obf_684faa48ae8c5049()
	if __obf_f16b4157911bc9af == ']' {
		__obf_56bad1827144ee34.
			UnsafeSet(__obf_2a1474febb02279b, __obf_56bad1827144ee34.UnsafeMakeSlice(0, 0))
		return
	}
	__obf_1bb30e8a74ed8233.__obf_a163df67f9bb1c4b()
	__obf_56bad1827144ee34.
		UnsafeGrow(__obf_2a1474febb02279b, 1)
	__obf_2b0afef38cdec72a := __obf_56bad1827144ee34.UnsafeGetIndex(__obf_2a1474febb02279b, 0)
	__obf_6fd3f72eb9b5574c.__obf_0110b2c027051777.
		Decode(__obf_2b0afef38cdec72a, __obf_1bb30e8a74ed8233)
	__obf_e3fec96653416447 := 1
	for __obf_f16b4157911bc9af = __obf_1bb30e8a74ed8233.__obf_684faa48ae8c5049(); __obf_f16b4157911bc9af == ','; __obf_f16b4157911bc9af = __obf_1bb30e8a74ed8233.__obf_684faa48ae8c5049() {
		__obf_b0acd231dbd22334 := __obf_e3fec96653416447
		__obf_e3fec96653416447 += 1
		__obf_56bad1827144ee34.
			UnsafeGrow(__obf_2a1474febb02279b, __obf_e3fec96653416447)
		__obf_2b0afef38cdec72a = __obf_56bad1827144ee34.UnsafeGetIndex(__obf_2a1474febb02279b, __obf_b0acd231dbd22334)
		__obf_6fd3f72eb9b5574c.__obf_0110b2c027051777.
			Decode(__obf_2b0afef38cdec72a, __obf_1bb30e8a74ed8233)
	}
	if __obf_f16b4157911bc9af != ']' {
		__obf_1bb30e8a74ed8233.
			ReportError("decode slice", "expect ], but found "+string([]byte{__obf_f16b4157911bc9af}))
		return
	}
}
