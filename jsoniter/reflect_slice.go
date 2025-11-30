package __obf_c3cd04a15c56f32f

import (
	"fmt"
	"github.com/modern-go/reflect2"
	"io"
	"unsafe"
)

func __obf_f7de3e7720df8df1(__obf_62435d89ebefd5aa *__obf_62435d89ebefd5aa, __obf_8667d4fc2a95b0d7 reflect2.Type) ValDecoder {
	__obf_ada810583e2cffde := __obf_8667d4fc2a95b0d7.(*reflect2.UnsafeSliceType)
	__obf_924cc7ef585cdfb0 := __obf_eddb00a5736b0fe7(__obf_62435d89ebefd5aa.append("[sliceElem]"), __obf_ada810583e2cffde.Elem())
	return &__obf_c58a03309b11d437{__obf_ada810583e2cffde, __obf_924cc7ef585cdfb0}
}

func __obf_b1da7c0e6535593d(__obf_62435d89ebefd5aa *__obf_62435d89ebefd5aa, __obf_8667d4fc2a95b0d7 reflect2.Type) ValEncoder {
	__obf_ada810583e2cffde := __obf_8667d4fc2a95b0d7.(*reflect2.UnsafeSliceType)
	__obf_c85b895560db528f := __obf_dd4ab965a9fde81c(__obf_62435d89ebefd5aa.append("[sliceElem]"), __obf_ada810583e2cffde.Elem())
	return &__obf_9a75119645c258dc{__obf_ada810583e2cffde, __obf_c85b895560db528f}
}

type __obf_9a75119645c258dc struct {
	__obf_ada810583e2cffde *reflect2.UnsafeSliceType
	__obf_2b4ba80da80f8306 ValEncoder
}

func (__obf_c85b895560db528f *__obf_9a75119645c258dc) Encode(__obf_753ab3fb72cdd659 unsafe.Pointer, __obf_2361f5214e84c60f *Stream) {
	if __obf_c85b895560db528f.__obf_ada810583e2cffde.UnsafeIsNil(__obf_753ab3fb72cdd659) {
		__obf_2361f5214e84c60f.
			WriteNil()
		return
	}
	__obf_c29c662da997f101 := __obf_c85b895560db528f.__obf_ada810583e2cffde.UnsafeLengthOf(__obf_753ab3fb72cdd659)
	if __obf_c29c662da997f101 == 0 {
		__obf_2361f5214e84c60f.
			WriteEmptyArray()
		return
	}
	__obf_2361f5214e84c60f.
		WriteArrayStart()
	__obf_c85b895560db528f.__obf_2b4ba80da80f8306.
		Encode(__obf_c85b895560db528f.__obf_ada810583e2cffde.UnsafeGetIndex(__obf_753ab3fb72cdd659, 0), __obf_2361f5214e84c60f)
	for __obf_28d099df85f083a8 := 1; __obf_28d099df85f083a8 < __obf_c29c662da997f101; __obf_28d099df85f083a8++ {
		__obf_2361f5214e84c60f.
			WriteMore()
		__obf_2d94a9e4f3d309fd := __obf_c85b895560db528f.__obf_ada810583e2cffde.UnsafeGetIndex(__obf_753ab3fb72cdd659, __obf_28d099df85f083a8)
		__obf_c85b895560db528f.__obf_2b4ba80da80f8306.
			Encode(__obf_2d94a9e4f3d309fd, __obf_2361f5214e84c60f)
	}
	__obf_2361f5214e84c60f.
		WriteArrayEnd()
	if __obf_2361f5214e84c60f.Error != nil && __obf_2361f5214e84c60f.Error != io.EOF {
		__obf_2361f5214e84c60f.
			Error = fmt.Errorf("%v: %s", __obf_c85b895560db528f.__obf_ada810583e2cffde, __obf_2361f5214e84c60f.Error.Error())
	}
}

func (__obf_c85b895560db528f *__obf_9a75119645c258dc) IsEmpty(__obf_753ab3fb72cdd659 unsafe.Pointer) bool {
	return __obf_c85b895560db528f.__obf_ada810583e2cffde.UnsafeLengthOf(__obf_753ab3fb72cdd659) == 0
}

type __obf_c58a03309b11d437 struct {
	__obf_ada810583e2cffde *reflect2.UnsafeSliceType
	__obf_fb5db223a2c09df6 ValDecoder
}

func (__obf_924cc7ef585cdfb0 *__obf_c58a03309b11d437) Decode(__obf_753ab3fb72cdd659 unsafe.Pointer, __obf_edd9fe48d39445e4 *Iterator) {
	__obf_924cc7ef585cdfb0.__obf_8e9e333265155dbf(__obf_753ab3fb72cdd659, __obf_edd9fe48d39445e4)
	if __obf_edd9fe48d39445e4.Error != nil && __obf_edd9fe48d39445e4.Error != io.EOF {
		__obf_edd9fe48d39445e4.
			Error = fmt.Errorf("%v: %s", __obf_924cc7ef585cdfb0.__obf_ada810583e2cffde, __obf_edd9fe48d39445e4.Error.Error())
	}
}

func (__obf_924cc7ef585cdfb0 *__obf_c58a03309b11d437) __obf_8e9e333265155dbf(__obf_753ab3fb72cdd659 unsafe.Pointer, __obf_edd9fe48d39445e4 *Iterator) {
	__obf_0c1bc1e511a43120 := __obf_edd9fe48d39445e4.__obf_a26eeeac1d6f5a11()
	__obf_ada810583e2cffde := __obf_924cc7ef585cdfb0.__obf_ada810583e2cffde
	if __obf_0c1bc1e511a43120 == 'n' {
		__obf_edd9fe48d39445e4.__obf_8bc1f4b8d62f5247('u', 'l', 'l')
		__obf_ada810583e2cffde.
			UnsafeSetNil(__obf_753ab3fb72cdd659)
		return
	}
	if __obf_0c1bc1e511a43120 != '[' {
		__obf_edd9fe48d39445e4.
			ReportError("decode slice", "expect [ or n, but found "+string([]byte{__obf_0c1bc1e511a43120}))
		return
	}
	__obf_0c1bc1e511a43120 = __obf_edd9fe48d39445e4.__obf_a26eeeac1d6f5a11()
	if __obf_0c1bc1e511a43120 == ']' {
		__obf_ada810583e2cffde.
			UnsafeSet(__obf_753ab3fb72cdd659, __obf_ada810583e2cffde.UnsafeMakeSlice(0, 0))
		return
	}
	__obf_edd9fe48d39445e4.__obf_a0622aded2d86ded()
	__obf_ada810583e2cffde.
		UnsafeGrow(__obf_753ab3fb72cdd659, 1)
	__obf_2d94a9e4f3d309fd := __obf_ada810583e2cffde.UnsafeGetIndex(__obf_753ab3fb72cdd659, 0)
	__obf_924cc7ef585cdfb0.__obf_fb5db223a2c09df6.
		Decode(__obf_2d94a9e4f3d309fd, __obf_edd9fe48d39445e4)
	__obf_c29c662da997f101 := 1
	for __obf_0c1bc1e511a43120 = __obf_edd9fe48d39445e4.__obf_a26eeeac1d6f5a11(); __obf_0c1bc1e511a43120 == ','; __obf_0c1bc1e511a43120 = __obf_edd9fe48d39445e4.__obf_a26eeeac1d6f5a11() {
		__obf_ad034d1ef83c8230 := __obf_c29c662da997f101
		__obf_c29c662da997f101 += 1
		__obf_ada810583e2cffde.
			UnsafeGrow(__obf_753ab3fb72cdd659, __obf_c29c662da997f101)
		__obf_2d94a9e4f3d309fd = __obf_ada810583e2cffde.UnsafeGetIndex(__obf_753ab3fb72cdd659, __obf_ad034d1ef83c8230)
		__obf_924cc7ef585cdfb0.__obf_fb5db223a2c09df6.
			Decode(__obf_2d94a9e4f3d309fd, __obf_edd9fe48d39445e4)
	}
	if __obf_0c1bc1e511a43120 != ']' {
		__obf_edd9fe48d39445e4.
			ReportError("decode slice", "expect ], but found "+string([]byte{__obf_0c1bc1e511a43120}))
		return
	}
}
