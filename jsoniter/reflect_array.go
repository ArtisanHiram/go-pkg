package __obf_703d23ba520c3295

import (
	"fmt"
	"github.com/modern-go/reflect2"
	"io"
	"unsafe"
)

func __obf_71ca52bfb4b492aa(__obf_2aaf7367de3ff86d *__obf_2aaf7367de3ff86d, __obf_3b7f6abbae19451e reflect2.Type) ValDecoder {
	__obf_a699c6c1641cf09d := __obf_3b7f6abbae19451e.(*reflect2.UnsafeArrayType)
	__obf_c9719e68325f7d44 := __obf_b27fe2bc17d94621(__obf_2aaf7367de3ff86d.append("[arrayElem]"), __obf_a699c6c1641cf09d.Elem())
	return &__obf_a3c612b96bef8595{__obf_a699c6c1641cf09d, __obf_c9719e68325f7d44}
}

func __obf_55bee799e6e81597(__obf_2aaf7367de3ff86d *__obf_2aaf7367de3ff86d, __obf_3b7f6abbae19451e reflect2.Type) ValEncoder {
	__obf_a699c6c1641cf09d := __obf_3b7f6abbae19451e.(*reflect2.UnsafeArrayType)
	if __obf_a699c6c1641cf09d.Len() == 0 {
		return __obf_8068b642c9d99edd{}
	}
	__obf_cf840243a12ee308 := __obf_906496c658b349cf(__obf_2aaf7367de3ff86d.append("[arrayElem]"), __obf_a699c6c1641cf09d.Elem())
	return &__obf_dfdeb193f963ce62{__obf_a699c6c1641cf09d, __obf_cf840243a12ee308}
}

type __obf_8068b642c9d99edd struct{}

func (__obf_cf840243a12ee308 __obf_8068b642c9d99edd) Encode(__obf_47fa53f3d161f45c unsafe.Pointer, __obf_9a34f62857fb1d1d *Stream) {
	__obf_9a34f62857fb1d1d.
		WriteEmptyArray()
}

func (__obf_cf840243a12ee308 __obf_8068b642c9d99edd) IsEmpty(__obf_47fa53f3d161f45c unsafe.Pointer) bool {
	return true
}

type __obf_dfdeb193f963ce62 struct {
	__obf_a699c6c1641cf09d *reflect2.UnsafeArrayType
	__obf_8bb551383a0ba60c ValEncoder
}

func (__obf_cf840243a12ee308 *__obf_dfdeb193f963ce62) Encode(__obf_47fa53f3d161f45c unsafe.Pointer, __obf_9a34f62857fb1d1d *Stream) {
	__obf_9a34f62857fb1d1d.
		WriteArrayStart()
	__obf_0863d94fb2cc079b := unsafe.Pointer(__obf_47fa53f3d161f45c)
	__obf_cf840243a12ee308.__obf_8bb551383a0ba60c.
		Encode(__obf_0863d94fb2cc079b, __obf_9a34f62857fb1d1d)
	for __obf_b0a5d2bd48690f1d := 1; __obf_b0a5d2bd48690f1d < __obf_cf840243a12ee308.__obf_a699c6c1641cf09d.Len(); __obf_b0a5d2bd48690f1d++ {
		__obf_9a34f62857fb1d1d.
			WriteMore()
		__obf_0863d94fb2cc079b = __obf_cf840243a12ee308.__obf_a699c6c1641cf09d.UnsafeGetIndex(__obf_47fa53f3d161f45c, __obf_b0a5d2bd48690f1d)
		__obf_cf840243a12ee308.__obf_8bb551383a0ba60c.
			Encode(__obf_0863d94fb2cc079b, __obf_9a34f62857fb1d1d)
	}
	__obf_9a34f62857fb1d1d.
		WriteArrayEnd()
	if __obf_9a34f62857fb1d1d.Error != nil && __obf_9a34f62857fb1d1d.Error != io.EOF {
		__obf_9a34f62857fb1d1d.
			Error = fmt.Errorf("%v: %s", __obf_cf840243a12ee308.__obf_a699c6c1641cf09d, __obf_9a34f62857fb1d1d.Error.Error())
	}
}

func (__obf_cf840243a12ee308 *__obf_dfdeb193f963ce62) IsEmpty(__obf_47fa53f3d161f45c unsafe.Pointer) bool {
	return false
}

type __obf_a3c612b96bef8595 struct {
	__obf_a699c6c1641cf09d *reflect2.UnsafeArrayType
	__obf_0db873ef26515049 ValDecoder
}

func (__obf_c9719e68325f7d44 *__obf_a3c612b96bef8595) Decode(__obf_47fa53f3d161f45c unsafe.Pointer, __obf_47edab4c16a2d88d *Iterator) {
	__obf_c9719e68325f7d44.__obf_968f6753c844a49c(__obf_47fa53f3d161f45c, __obf_47edab4c16a2d88d)
	if __obf_47edab4c16a2d88d.Error != nil && __obf_47edab4c16a2d88d.Error != io.EOF {
		__obf_47edab4c16a2d88d.
			Error = fmt.Errorf("%v: %s", __obf_c9719e68325f7d44.__obf_a699c6c1641cf09d, __obf_47edab4c16a2d88d.Error.Error())
	}
}

func (__obf_c9719e68325f7d44 *__obf_a3c612b96bef8595) __obf_968f6753c844a49c(__obf_47fa53f3d161f45c unsafe.Pointer, __obf_47edab4c16a2d88d *Iterator) {
	__obf_bd08f5161fff294a := __obf_47edab4c16a2d88d.__obf_6f7a9d6f8e846eec()
	__obf_a699c6c1641cf09d := __obf_c9719e68325f7d44.__obf_a699c6c1641cf09d
	if __obf_bd08f5161fff294a == 'n' {
		__obf_47edab4c16a2d88d.__obf_0704d4be12be5e96('u', 'l', 'l')
		return
	}
	if __obf_bd08f5161fff294a != '[' {
		__obf_47edab4c16a2d88d.
			ReportError("decode array", "expect [ or n, but found "+string([]byte{__obf_bd08f5161fff294a}))
		return
	}
	__obf_bd08f5161fff294a = __obf_47edab4c16a2d88d.__obf_6f7a9d6f8e846eec()
	if __obf_bd08f5161fff294a == ']' {
		return
	}
	__obf_47edab4c16a2d88d.__obf_743b1a47c346c8a3()
	__obf_0863d94fb2cc079b := __obf_a699c6c1641cf09d.UnsafeGetIndex(__obf_47fa53f3d161f45c, 0)
	__obf_c9719e68325f7d44.__obf_0db873ef26515049.
		Decode(__obf_0863d94fb2cc079b, __obf_47edab4c16a2d88d)
	__obf_467174c7140a26b0 := 1
	for __obf_bd08f5161fff294a = __obf_47edab4c16a2d88d.__obf_6f7a9d6f8e846eec(); __obf_bd08f5161fff294a == ','; __obf_bd08f5161fff294a = __obf_47edab4c16a2d88d.__obf_6f7a9d6f8e846eec() {
		if __obf_467174c7140a26b0 >= __obf_a699c6c1641cf09d.Len() {
			__obf_47edab4c16a2d88d.
				Skip()
			continue
		}
		__obf_3e35e690d9043ba8 := __obf_467174c7140a26b0
		__obf_467174c7140a26b0 += 1
		__obf_0863d94fb2cc079b = __obf_a699c6c1641cf09d.UnsafeGetIndex(__obf_47fa53f3d161f45c, __obf_3e35e690d9043ba8)
		__obf_c9719e68325f7d44.__obf_0db873ef26515049.
			Decode(__obf_0863d94fb2cc079b, __obf_47edab4c16a2d88d)
	}
	if __obf_bd08f5161fff294a != ']' {
		__obf_47edab4c16a2d88d.
			ReportError("decode array", "expect ], but found "+string([]byte{__obf_bd08f5161fff294a}))
		return
	}
}
