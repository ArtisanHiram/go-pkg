package __obf_c3cd04a15c56f32f

import (
	"fmt"
	"github.com/modern-go/reflect2"
	"io"
	"unsafe"
)

func __obf_2281a40bdf22829a(__obf_62435d89ebefd5aa *__obf_62435d89ebefd5aa, __obf_8667d4fc2a95b0d7 reflect2.Type) ValDecoder {
	__obf_b2fdbbc22827e2d8 := __obf_8667d4fc2a95b0d7.(*reflect2.UnsafeArrayType)
	__obf_924cc7ef585cdfb0 := __obf_eddb00a5736b0fe7(__obf_62435d89ebefd5aa.append("[arrayElem]"), __obf_b2fdbbc22827e2d8.Elem())
	return &__obf_ca6c591cf4cd95e2{__obf_b2fdbbc22827e2d8, __obf_924cc7ef585cdfb0}
}

func __obf_66f4ed70a70f0e6c(__obf_62435d89ebefd5aa *__obf_62435d89ebefd5aa, __obf_8667d4fc2a95b0d7 reflect2.Type) ValEncoder {
	__obf_b2fdbbc22827e2d8 := __obf_8667d4fc2a95b0d7.(*reflect2.UnsafeArrayType)
	if __obf_b2fdbbc22827e2d8.Len() == 0 {
		return __obf_c1cd16973db444c6{}
	}
	__obf_c85b895560db528f := __obf_dd4ab965a9fde81c(__obf_62435d89ebefd5aa.append("[arrayElem]"), __obf_b2fdbbc22827e2d8.Elem())
	return &__obf_06afadf28a5822f4{__obf_b2fdbbc22827e2d8, __obf_c85b895560db528f}
}

type __obf_c1cd16973db444c6 struct{}

func (__obf_c85b895560db528f __obf_c1cd16973db444c6) Encode(__obf_753ab3fb72cdd659 unsafe.Pointer, __obf_2361f5214e84c60f *Stream) {
	__obf_2361f5214e84c60f.
		WriteEmptyArray()
}

func (__obf_c85b895560db528f __obf_c1cd16973db444c6) IsEmpty(__obf_753ab3fb72cdd659 unsafe.Pointer) bool {
	return true
}

type __obf_06afadf28a5822f4 struct {
	__obf_b2fdbbc22827e2d8 *reflect2.UnsafeArrayType
	__obf_2b4ba80da80f8306 ValEncoder
}

func (__obf_c85b895560db528f *__obf_06afadf28a5822f4) Encode(__obf_753ab3fb72cdd659 unsafe.Pointer, __obf_2361f5214e84c60f *Stream) {
	__obf_2361f5214e84c60f.
		WriteArrayStart()
	__obf_2d94a9e4f3d309fd := unsafe.Pointer(__obf_753ab3fb72cdd659)
	__obf_c85b895560db528f.__obf_2b4ba80da80f8306.
		Encode(__obf_2d94a9e4f3d309fd, __obf_2361f5214e84c60f)
	for __obf_28d099df85f083a8 := 1; __obf_28d099df85f083a8 < __obf_c85b895560db528f.__obf_b2fdbbc22827e2d8.Len(); __obf_28d099df85f083a8++ {
		__obf_2361f5214e84c60f.
			WriteMore()
		__obf_2d94a9e4f3d309fd = __obf_c85b895560db528f.__obf_b2fdbbc22827e2d8.UnsafeGetIndex(__obf_753ab3fb72cdd659, __obf_28d099df85f083a8)
		__obf_c85b895560db528f.__obf_2b4ba80da80f8306.
			Encode(__obf_2d94a9e4f3d309fd, __obf_2361f5214e84c60f)
	}
	__obf_2361f5214e84c60f.
		WriteArrayEnd()
	if __obf_2361f5214e84c60f.Error != nil && __obf_2361f5214e84c60f.Error != io.EOF {
		__obf_2361f5214e84c60f.
			Error = fmt.Errorf("%v: %s", __obf_c85b895560db528f.__obf_b2fdbbc22827e2d8, __obf_2361f5214e84c60f.Error.Error())
	}
}

func (__obf_c85b895560db528f *__obf_06afadf28a5822f4) IsEmpty(__obf_753ab3fb72cdd659 unsafe.Pointer) bool {
	return false
}

type __obf_ca6c591cf4cd95e2 struct {
	__obf_b2fdbbc22827e2d8 *reflect2.UnsafeArrayType
	__obf_fb5db223a2c09df6 ValDecoder
}

func (__obf_924cc7ef585cdfb0 *__obf_ca6c591cf4cd95e2) Decode(__obf_753ab3fb72cdd659 unsafe.Pointer, __obf_edd9fe48d39445e4 *Iterator) {
	__obf_924cc7ef585cdfb0.__obf_8e9e333265155dbf(__obf_753ab3fb72cdd659, __obf_edd9fe48d39445e4)
	if __obf_edd9fe48d39445e4.Error != nil && __obf_edd9fe48d39445e4.Error != io.EOF {
		__obf_edd9fe48d39445e4.
			Error = fmt.Errorf("%v: %s", __obf_924cc7ef585cdfb0.__obf_b2fdbbc22827e2d8, __obf_edd9fe48d39445e4.Error.Error())
	}
}

func (__obf_924cc7ef585cdfb0 *__obf_ca6c591cf4cd95e2) __obf_8e9e333265155dbf(__obf_753ab3fb72cdd659 unsafe.Pointer, __obf_edd9fe48d39445e4 *Iterator) {
	__obf_0c1bc1e511a43120 := __obf_edd9fe48d39445e4.__obf_a26eeeac1d6f5a11()
	__obf_b2fdbbc22827e2d8 := __obf_924cc7ef585cdfb0.__obf_b2fdbbc22827e2d8
	if __obf_0c1bc1e511a43120 == 'n' {
		__obf_edd9fe48d39445e4.__obf_8bc1f4b8d62f5247('u', 'l', 'l')
		return
	}
	if __obf_0c1bc1e511a43120 != '[' {
		__obf_edd9fe48d39445e4.
			ReportError("decode array", "expect [ or n, but found "+string([]byte{__obf_0c1bc1e511a43120}))
		return
	}
	__obf_0c1bc1e511a43120 = __obf_edd9fe48d39445e4.__obf_a26eeeac1d6f5a11()
	if __obf_0c1bc1e511a43120 == ']' {
		return
	}
	__obf_edd9fe48d39445e4.__obf_a0622aded2d86ded()
	__obf_2d94a9e4f3d309fd := __obf_b2fdbbc22827e2d8.UnsafeGetIndex(__obf_753ab3fb72cdd659, 0)
	__obf_924cc7ef585cdfb0.__obf_fb5db223a2c09df6.
		Decode(__obf_2d94a9e4f3d309fd, __obf_edd9fe48d39445e4)
	__obf_c29c662da997f101 := 1
	for __obf_0c1bc1e511a43120 = __obf_edd9fe48d39445e4.__obf_a26eeeac1d6f5a11(); __obf_0c1bc1e511a43120 == ','; __obf_0c1bc1e511a43120 = __obf_edd9fe48d39445e4.__obf_a26eeeac1d6f5a11() {
		if __obf_c29c662da997f101 >= __obf_b2fdbbc22827e2d8.Len() {
			__obf_edd9fe48d39445e4.
				Skip()
			continue
		}
		__obf_ad034d1ef83c8230 := __obf_c29c662da997f101
		__obf_c29c662da997f101 += 1
		__obf_2d94a9e4f3d309fd = __obf_b2fdbbc22827e2d8.UnsafeGetIndex(__obf_753ab3fb72cdd659, __obf_ad034d1ef83c8230)
		__obf_924cc7ef585cdfb0.__obf_fb5db223a2c09df6.
			Decode(__obf_2d94a9e4f3d309fd, __obf_edd9fe48d39445e4)
	}
	if __obf_0c1bc1e511a43120 != ']' {
		__obf_edd9fe48d39445e4.
			ReportError("decode array", "expect ], but found "+string([]byte{__obf_0c1bc1e511a43120}))
		return
	}
}
