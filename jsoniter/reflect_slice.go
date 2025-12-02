package __obf_c7b79b12b33d8238

import (
	"fmt"
	"github.com/modern-go/reflect2"
	"io"
	"unsafe"
)

func __obf_978533a1047109a9(__obf_99ec45908bceabdb *__obf_99ec45908bceabdb, __obf_edcded11af6ebc4c reflect2.Type) ValDecoder {
	__obf_782666043ff6c7ac := __obf_edcded11af6ebc4c.(*reflect2.UnsafeSliceType)
	__obf_801f19702638809c := __obf_bb14644cc3f030b3(__obf_99ec45908bceabdb.append("[sliceElem]"), __obf_782666043ff6c7ac.Elem())
	return &__obf_06c2733d7b043ec8{__obf_782666043ff6c7ac, __obf_801f19702638809c}
}

func __obf_1a6fc40f1f862856(__obf_99ec45908bceabdb *__obf_99ec45908bceabdb, __obf_edcded11af6ebc4c reflect2.Type) ValEncoder {
	__obf_782666043ff6c7ac := __obf_edcded11af6ebc4c.(*reflect2.UnsafeSliceType)
	__obf_c091c27480fae550 := __obf_85f0e4c352da4678(__obf_99ec45908bceabdb.append("[sliceElem]"), __obf_782666043ff6c7ac.Elem())
	return &__obf_577577703e34ae45{__obf_782666043ff6c7ac, __obf_c091c27480fae550}
}

type __obf_577577703e34ae45 struct {
	__obf_782666043ff6c7ac *reflect2.UnsafeSliceType
	__obf_e6894379459b04e8 ValEncoder
}

func (__obf_c091c27480fae550 *__obf_577577703e34ae45) Encode(__obf_575c04f2b9d91315 unsafe.Pointer, __obf_d8f50bcfe87d8b47 *Stream) {
	if __obf_c091c27480fae550.__obf_782666043ff6c7ac.UnsafeIsNil(__obf_575c04f2b9d91315) {
		__obf_d8f50bcfe87d8b47.
			WriteNil()
		return
	}
	__obf_55238f3512972ca4 := __obf_c091c27480fae550.__obf_782666043ff6c7ac.UnsafeLengthOf(__obf_575c04f2b9d91315)
	if __obf_55238f3512972ca4 == 0 {
		__obf_d8f50bcfe87d8b47.
			WriteEmptyArray()
		return
	}
	__obf_d8f50bcfe87d8b47.
		WriteArrayStart()
	__obf_c091c27480fae550.__obf_e6894379459b04e8.
		Encode(__obf_c091c27480fae550.__obf_782666043ff6c7ac.UnsafeGetIndex(__obf_575c04f2b9d91315, 0), __obf_d8f50bcfe87d8b47)
	for __obf_a051269af3a541bb := 1; __obf_a051269af3a541bb < __obf_55238f3512972ca4; __obf_a051269af3a541bb++ {
		__obf_d8f50bcfe87d8b47.
			WriteMore()
		__obf_0c67f6b9b84c15fa := __obf_c091c27480fae550.__obf_782666043ff6c7ac.UnsafeGetIndex(__obf_575c04f2b9d91315, __obf_a051269af3a541bb)
		__obf_c091c27480fae550.__obf_e6894379459b04e8.
			Encode(__obf_0c67f6b9b84c15fa, __obf_d8f50bcfe87d8b47)
	}
	__obf_d8f50bcfe87d8b47.
		WriteArrayEnd()
	if __obf_d8f50bcfe87d8b47.Error != nil && __obf_d8f50bcfe87d8b47.Error != io.EOF {
		__obf_d8f50bcfe87d8b47.
			Error = fmt.Errorf("%v: %s", __obf_c091c27480fae550.__obf_782666043ff6c7ac, __obf_d8f50bcfe87d8b47.Error.Error())
	}
}

func (__obf_c091c27480fae550 *__obf_577577703e34ae45) IsEmpty(__obf_575c04f2b9d91315 unsafe.Pointer) bool {
	return __obf_c091c27480fae550.__obf_782666043ff6c7ac.UnsafeLengthOf(__obf_575c04f2b9d91315) == 0
}

type __obf_06c2733d7b043ec8 struct {
	__obf_782666043ff6c7ac *reflect2.UnsafeSliceType
	__obf_5b1e51b09d209694 ValDecoder
}

func (__obf_801f19702638809c *__obf_06c2733d7b043ec8) Decode(__obf_575c04f2b9d91315 unsafe.Pointer, __obf_0da8c843dad13139 *Iterator) {
	__obf_801f19702638809c.__obf_346fc39501146782(__obf_575c04f2b9d91315, __obf_0da8c843dad13139)
	if __obf_0da8c843dad13139.Error != nil && __obf_0da8c843dad13139.Error != io.EOF {
		__obf_0da8c843dad13139.
			Error = fmt.Errorf("%v: %s", __obf_801f19702638809c.__obf_782666043ff6c7ac, __obf_0da8c843dad13139.Error.Error())
	}
}

func (__obf_801f19702638809c *__obf_06c2733d7b043ec8) __obf_346fc39501146782(__obf_575c04f2b9d91315 unsafe.Pointer, __obf_0da8c843dad13139 *Iterator) {
	__obf_12577c03a12f0d2c := __obf_0da8c843dad13139.__obf_2b436fcb1c0ca36d()
	__obf_782666043ff6c7ac := __obf_801f19702638809c.__obf_782666043ff6c7ac
	if __obf_12577c03a12f0d2c == 'n' {
		__obf_0da8c843dad13139.__obf_5da8d54e7a1e782c('u', 'l', 'l')
		__obf_782666043ff6c7ac.
			UnsafeSetNil(__obf_575c04f2b9d91315)
		return
	}
	if __obf_12577c03a12f0d2c != '[' {
		__obf_0da8c843dad13139.
			ReportError("decode slice", "expect [ or n, but found "+string([]byte{__obf_12577c03a12f0d2c}))
		return
	}
	__obf_12577c03a12f0d2c = __obf_0da8c843dad13139.__obf_2b436fcb1c0ca36d()
	if __obf_12577c03a12f0d2c == ']' {
		__obf_782666043ff6c7ac.
			UnsafeSet(__obf_575c04f2b9d91315, __obf_782666043ff6c7ac.UnsafeMakeSlice(0, 0))
		return
	}
	__obf_0da8c843dad13139.__obf_a675366c80290b83()
	__obf_782666043ff6c7ac.
		UnsafeGrow(__obf_575c04f2b9d91315, 1)
	__obf_0c67f6b9b84c15fa := __obf_782666043ff6c7ac.UnsafeGetIndex(__obf_575c04f2b9d91315, 0)
	__obf_801f19702638809c.__obf_5b1e51b09d209694.
		Decode(__obf_0c67f6b9b84c15fa, __obf_0da8c843dad13139)
	__obf_55238f3512972ca4 := 1
	for __obf_12577c03a12f0d2c = __obf_0da8c843dad13139.__obf_2b436fcb1c0ca36d(); __obf_12577c03a12f0d2c == ','; __obf_12577c03a12f0d2c = __obf_0da8c843dad13139.__obf_2b436fcb1c0ca36d() {
		__obf_359f6e4a6d05c6da := __obf_55238f3512972ca4
		__obf_55238f3512972ca4 += 1
		__obf_782666043ff6c7ac.
			UnsafeGrow(__obf_575c04f2b9d91315, __obf_55238f3512972ca4)
		__obf_0c67f6b9b84c15fa = __obf_782666043ff6c7ac.UnsafeGetIndex(__obf_575c04f2b9d91315, __obf_359f6e4a6d05c6da)
		__obf_801f19702638809c.__obf_5b1e51b09d209694.
			Decode(__obf_0c67f6b9b84c15fa, __obf_0da8c843dad13139)
	}
	if __obf_12577c03a12f0d2c != ']' {
		__obf_0da8c843dad13139.
			ReportError("decode slice", "expect ], but found "+string([]byte{__obf_12577c03a12f0d2c}))
		return
	}
}
