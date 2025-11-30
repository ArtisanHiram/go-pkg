package __obf_c3cd04a15c56f32f

import (
	"fmt"
	"io"
	"reflect"
	"sort"
	"unsafe"

	"github.com/modern-go/reflect2"
)

func __obf_96d7d01ea15ea41d(__obf_62435d89ebefd5aa *__obf_62435d89ebefd5aa, __obf_8667d4fc2a95b0d7 reflect2.Type) ValDecoder {
	__obf_3915d58c01112789 := __obf_8667d4fc2a95b0d7.(*reflect2.UnsafeMapType)
	__obf_258fab27b416dbfd := __obf_22de96dea5400836(__obf_62435d89ebefd5aa.append("[mapKey]"), __obf_3915d58c01112789.Key())
	__obf_fb5db223a2c09df6 := __obf_eddb00a5736b0fe7(__obf_62435d89ebefd5aa.append("[mapElem]"), __obf_3915d58c01112789.Elem())
	return &__obf_11d4688ddb117a11{__obf_3915d58c01112789: __obf_3915d58c01112789, __obf_1774f10425ffe935: __obf_3915d58c01112789.Key(), __obf_78b97e7cf83971dd: __obf_3915d58c01112789.Elem(), __obf_258fab27b416dbfd: __obf_258fab27b416dbfd, __obf_fb5db223a2c09df6: __obf_fb5db223a2c09df6}
}

func __obf_f36df00362e2ff3a(__obf_62435d89ebefd5aa *__obf_62435d89ebefd5aa, __obf_8667d4fc2a95b0d7 reflect2.Type) ValEncoder {
	__obf_3915d58c01112789 := __obf_8667d4fc2a95b0d7.(*reflect2.UnsafeMapType)
	if __obf_62435d89ebefd5aa.__obf_70c2e3a03299f745 {
		return &__obf_e06fd245f5efd38c{__obf_3915d58c01112789: __obf_3915d58c01112789, __obf_4f90b2990512879d: __obf_afebd67e4fbc0eb1(__obf_62435d89ebefd5aa.append("[mapKey]"), __obf_3915d58c01112789.Key()), __obf_2b4ba80da80f8306: __obf_dd4ab965a9fde81c(__obf_62435d89ebefd5aa.append("[mapElem]"), __obf_3915d58c01112789.Elem())}
	}
	return &__obf_b4c25ee21eb89b02{__obf_3915d58c01112789: __obf_3915d58c01112789, __obf_4f90b2990512879d: __obf_afebd67e4fbc0eb1(__obf_62435d89ebefd5aa.append("[mapKey]"), __obf_3915d58c01112789.Key()), __obf_2b4ba80da80f8306: __obf_dd4ab965a9fde81c(__obf_62435d89ebefd5aa.append("[mapElem]"), __obf_3915d58c01112789.Elem())}
}

func __obf_22de96dea5400836(__obf_62435d89ebefd5aa *__obf_62435d89ebefd5aa, __obf_8667d4fc2a95b0d7 reflect2.Type) ValDecoder {
	__obf_924cc7ef585cdfb0 := __obf_62435d89ebefd5aa.__obf_d86968bf26bed402.CreateMapKeyDecoder(__obf_8667d4fc2a95b0d7)
	if __obf_924cc7ef585cdfb0 != nil {
		return __obf_924cc7ef585cdfb0
	}
	for _, __obf_e4a614f491c1bb0c := range __obf_62435d89ebefd5aa.__obf_a1eac24bcc56f0a6 {
		__obf_924cc7ef585cdfb0 := __obf_e4a614f491c1bb0c.CreateMapKeyDecoder(__obf_8667d4fc2a95b0d7)
		if __obf_924cc7ef585cdfb0 != nil {
			return __obf_924cc7ef585cdfb0
		}
	}
	__obf_9fd3068ebc25d7f9 := reflect2.PtrTo(__obf_8667d4fc2a95b0d7)
	if __obf_9fd3068ebc25d7f9.Implements(__obf_31f96b637210e6eb) {
		return &__obf_cec111c41697b8e9{
			&__obf_646b7564744762af{__obf_797d4fe23b3556f8: __obf_9fd3068ebc25d7f9},
		}
	}
	if __obf_8667d4fc2a95b0d7.Implements(__obf_31f96b637210e6eb) {
		return &__obf_646b7564744762af{__obf_797d4fe23b3556f8: __obf_8667d4fc2a95b0d7}
	}
	if __obf_9fd3068ebc25d7f9.Implements(__obf_b7c5ad5d21139f9f) {
		return &__obf_cec111c41697b8e9{
			&__obf_d73f3e91d310f340{__obf_797d4fe23b3556f8: __obf_9fd3068ebc25d7f9},
		}
	}
	if __obf_8667d4fc2a95b0d7.Implements(__obf_b7c5ad5d21139f9f) {
		return &__obf_d73f3e91d310f340{__obf_797d4fe23b3556f8: __obf_8667d4fc2a95b0d7}
	}

	switch __obf_8667d4fc2a95b0d7.Kind() {
	case reflect.String:
		return __obf_eddb00a5736b0fe7(__obf_62435d89ebefd5aa, reflect2.DefaultTypeOfKind(reflect.String))
	case reflect.Bool,
		reflect.Uint8, reflect.Int8,
		reflect.Uint16, reflect.Int16,
		reflect.Uint32, reflect.Int32,
		reflect.Uint64, reflect.Int64,
		reflect.Uint, reflect.Int,
		reflect.Float32, reflect.Float64,
		reflect.Uintptr:
		__obf_8667d4fc2a95b0d7 = reflect2.DefaultTypeOfKind(__obf_8667d4fc2a95b0d7.Kind())
		return &__obf_360986424625997b{__obf_eddb00a5736b0fe7(__obf_62435d89ebefd5aa, __obf_8667d4fc2a95b0d7)}
	default:
		return &__obf_5ba850d645a49aec{__obf_5966ecc5edb9b17e: fmt.Errorf("unsupported map key type: %v", __obf_8667d4fc2a95b0d7)}
	}
}

func __obf_afebd67e4fbc0eb1(__obf_62435d89ebefd5aa *__obf_62435d89ebefd5aa, __obf_8667d4fc2a95b0d7 reflect2.Type) ValEncoder {
	__obf_c85b895560db528f := __obf_62435d89ebefd5aa.__obf_912608e689e59f9b.CreateMapKeyEncoder(__obf_8667d4fc2a95b0d7)
	if __obf_c85b895560db528f != nil {
		return __obf_c85b895560db528f
	}
	for _, __obf_e4a614f491c1bb0c := range __obf_62435d89ebefd5aa.__obf_a1eac24bcc56f0a6 {
		__obf_c85b895560db528f := __obf_e4a614f491c1bb0c.CreateMapKeyEncoder(__obf_8667d4fc2a95b0d7)
		if __obf_c85b895560db528f != nil {
			return __obf_c85b895560db528f
		}
	}

	if __obf_8667d4fc2a95b0d7.Kind() != reflect.String {
		if __obf_8667d4fc2a95b0d7 == __obf_118e2643fe9fc774 {
			return &__obf_956493b86d70ce45{__obf_de47f56a3aec06da: __obf_62435d89ebefd5aa.EncoderOf(reflect2.TypeOf(""))}
		}
		if __obf_8667d4fc2a95b0d7.Implements(__obf_118e2643fe9fc774) {
			return &__obf_d7eae9a85745e3c8{__obf_797d4fe23b3556f8: __obf_8667d4fc2a95b0d7, __obf_de47f56a3aec06da: __obf_62435d89ebefd5aa.EncoderOf(reflect2.TypeOf(""))}
		}
	}

	switch __obf_8667d4fc2a95b0d7.Kind() {
	case reflect.String:
		return __obf_dd4ab965a9fde81c(__obf_62435d89ebefd5aa, reflect2.DefaultTypeOfKind(reflect.String))
	case reflect.Bool,
		reflect.Uint8, reflect.Int8,
		reflect.Uint16, reflect.Int16,
		reflect.Uint32, reflect.Int32,
		reflect.Uint64, reflect.Int64,
		reflect.Uint, reflect.Int,
		reflect.Float32, reflect.Float64,
		reflect.Uintptr:
		__obf_8667d4fc2a95b0d7 = reflect2.DefaultTypeOfKind(__obf_8667d4fc2a95b0d7.Kind())
		return &__obf_9cb2d594f1f201ad{__obf_dd4ab965a9fde81c(__obf_62435d89ebefd5aa, __obf_8667d4fc2a95b0d7)}
	default:
		if __obf_8667d4fc2a95b0d7.Kind() == reflect.Interface {
			return &__obf_fe7c69e3b915406f{__obf_62435d89ebefd5aa, __obf_8667d4fc2a95b0d7}
		}
		return &__obf_a580dbbcc400f543{__obf_5966ecc5edb9b17e: fmt.Errorf("unsupported map key type: %v", __obf_8667d4fc2a95b0d7)}
	}
}

type __obf_11d4688ddb117a11 struct {
	__obf_3915d58c01112789 *reflect2.UnsafeMapType
	__obf_1774f10425ffe935 reflect2.Type
	__obf_78b97e7cf83971dd reflect2.Type
	__obf_258fab27b416dbfd ValDecoder
	__obf_fb5db223a2c09df6 ValDecoder
}

func (__obf_924cc7ef585cdfb0 *__obf_11d4688ddb117a11) Decode(__obf_753ab3fb72cdd659 unsafe.Pointer, __obf_edd9fe48d39445e4 *Iterator) {
	__obf_3915d58c01112789 := __obf_924cc7ef585cdfb0.__obf_3915d58c01112789
	__obf_0c1bc1e511a43120 := __obf_edd9fe48d39445e4.__obf_a26eeeac1d6f5a11()
	if __obf_0c1bc1e511a43120 == 'n' {
		__obf_edd9fe48d39445e4.__obf_8bc1f4b8d62f5247('u', 'l', 'l')
		*(*unsafe.Pointer)(__obf_753ab3fb72cdd659) = nil
		__obf_3915d58c01112789.
			UnsafeSet(__obf_753ab3fb72cdd659, __obf_3915d58c01112789.UnsafeNew())
		return
	}
	if __obf_3915d58c01112789.UnsafeIsNil(__obf_753ab3fb72cdd659) {
		__obf_3915d58c01112789.
			UnsafeSet(__obf_753ab3fb72cdd659, __obf_3915d58c01112789.UnsafeMakeMap(0))
	}
	if __obf_0c1bc1e511a43120 != '{' {
		__obf_edd9fe48d39445e4.
			ReportError("ReadMapCB", `expect { or n, but found `+string([]byte{__obf_0c1bc1e511a43120}))
		return
	}
	__obf_0c1bc1e511a43120 = __obf_edd9fe48d39445e4.__obf_a26eeeac1d6f5a11()
	if __obf_0c1bc1e511a43120 == '}' {
		return
	}
	__obf_edd9fe48d39445e4.__obf_a0622aded2d86ded()
	__obf_09e2475e3bcc63c4 := __obf_924cc7ef585cdfb0.__obf_1774f10425ffe935.UnsafeNew()
	__obf_924cc7ef585cdfb0.__obf_258fab27b416dbfd.
		Decode(__obf_09e2475e3bcc63c4, __obf_edd9fe48d39445e4)
	__obf_0c1bc1e511a43120 = __obf_edd9fe48d39445e4.__obf_a26eeeac1d6f5a11()
	if __obf_0c1bc1e511a43120 != ':' {
		__obf_edd9fe48d39445e4.
			ReportError("ReadMapCB", "expect : after object field, but found "+string([]byte{__obf_0c1bc1e511a43120}))
		return
	}
	__obf_991e715ce3758c78 := __obf_924cc7ef585cdfb0.__obf_78b97e7cf83971dd.UnsafeNew()
	__obf_924cc7ef585cdfb0.__obf_fb5db223a2c09df6.
		Decode(__obf_991e715ce3758c78, __obf_edd9fe48d39445e4)
	__obf_924cc7ef585cdfb0.__obf_3915d58c01112789.
		UnsafeSetIndex(__obf_753ab3fb72cdd659, __obf_09e2475e3bcc63c4, __obf_991e715ce3758c78)
	for __obf_0c1bc1e511a43120 = __obf_edd9fe48d39445e4.__obf_a26eeeac1d6f5a11(); __obf_0c1bc1e511a43120 == ','; __obf_0c1bc1e511a43120 = __obf_edd9fe48d39445e4.__obf_a26eeeac1d6f5a11() {
		__obf_09e2475e3bcc63c4 := __obf_924cc7ef585cdfb0.__obf_1774f10425ffe935.UnsafeNew()
		__obf_924cc7ef585cdfb0.__obf_258fab27b416dbfd.
			Decode(__obf_09e2475e3bcc63c4, __obf_edd9fe48d39445e4)
		__obf_0c1bc1e511a43120 = __obf_edd9fe48d39445e4.__obf_a26eeeac1d6f5a11()
		if __obf_0c1bc1e511a43120 != ':' {
			__obf_edd9fe48d39445e4.
				ReportError("ReadMapCB", "expect : after object field, but found "+string([]byte{__obf_0c1bc1e511a43120}))
			return
		}
		__obf_991e715ce3758c78 := __obf_924cc7ef585cdfb0.__obf_78b97e7cf83971dd.UnsafeNew()
		__obf_924cc7ef585cdfb0.__obf_fb5db223a2c09df6.
			Decode(__obf_991e715ce3758c78, __obf_edd9fe48d39445e4)
		__obf_924cc7ef585cdfb0.__obf_3915d58c01112789.
			UnsafeSetIndex(__obf_753ab3fb72cdd659, __obf_09e2475e3bcc63c4, __obf_991e715ce3758c78)
	}
	if __obf_0c1bc1e511a43120 != '}' {
		__obf_edd9fe48d39445e4.
			ReportError("ReadMapCB", `expect }, but found `+string([]byte{__obf_0c1bc1e511a43120}))
	}
}

type __obf_360986424625997b struct {
	__obf_924cc7ef585cdfb0 ValDecoder
}

func (__obf_924cc7ef585cdfb0 *__obf_360986424625997b) Decode(__obf_753ab3fb72cdd659 unsafe.Pointer, __obf_edd9fe48d39445e4 *Iterator) {
	__obf_0c1bc1e511a43120 := __obf_edd9fe48d39445e4.__obf_a26eeeac1d6f5a11()
	if __obf_0c1bc1e511a43120 != '"' {
		__obf_edd9fe48d39445e4.
			ReportError("ReadMapCB", `expect ", but found `+string([]byte{__obf_0c1bc1e511a43120}))
		return
	}
	__obf_924cc7ef585cdfb0.__obf_924cc7ef585cdfb0.
		Decode(__obf_753ab3fb72cdd659, __obf_edd9fe48d39445e4)
	__obf_0c1bc1e511a43120 = __obf_edd9fe48d39445e4.__obf_a26eeeac1d6f5a11()
	if __obf_0c1bc1e511a43120 != '"' {
		__obf_edd9fe48d39445e4.
			ReportError("ReadMapCB", `expect ", but found `+string([]byte{__obf_0c1bc1e511a43120}))
		return
	}
}

type __obf_9cb2d594f1f201ad struct {
	__obf_c85b895560db528f ValEncoder
}

func (__obf_c85b895560db528f *__obf_9cb2d594f1f201ad) Encode(__obf_753ab3fb72cdd659 unsafe.Pointer, __obf_2361f5214e84c60f *Stream) {
	__obf_2361f5214e84c60f.__obf_c4fec0edfb3875ad('"')
	__obf_c85b895560db528f.__obf_c85b895560db528f.
		Encode(__obf_753ab3fb72cdd659, __obf_2361f5214e84c60f)
	__obf_2361f5214e84c60f.__obf_c4fec0edfb3875ad('"')
}

func (__obf_c85b895560db528f *__obf_9cb2d594f1f201ad) IsEmpty(__obf_753ab3fb72cdd659 unsafe.Pointer) bool {
	return false
}

type __obf_fe7c69e3b915406f struct {
	__obf_62435d89ebefd5aa *__obf_62435d89ebefd5aa
	__obf_797d4fe23b3556f8 reflect2.Type
}

func (__obf_c85b895560db528f *__obf_fe7c69e3b915406f) Encode(__obf_753ab3fb72cdd659 unsafe.Pointer, __obf_2361f5214e84c60f *Stream) {
	__obf_50acae8c0a871ba1 := __obf_c85b895560db528f.__obf_797d4fe23b3556f8.UnsafeIndirect(__obf_753ab3fb72cdd659)
	__obf_afebd67e4fbc0eb1(__obf_c85b895560db528f.__obf_62435d89ebefd5aa, reflect2.TypeOf(__obf_50acae8c0a871ba1)).Encode(reflect2.PtrOf(__obf_50acae8c0a871ba1), __obf_2361f5214e84c60f)
}

func (__obf_c85b895560db528f *__obf_fe7c69e3b915406f) IsEmpty(__obf_753ab3fb72cdd659 unsafe.Pointer) bool {
	__obf_50acae8c0a871ba1 := __obf_c85b895560db528f.__obf_797d4fe23b3556f8.UnsafeIndirect(__obf_753ab3fb72cdd659)
	return __obf_afebd67e4fbc0eb1(__obf_c85b895560db528f.__obf_62435d89ebefd5aa, reflect2.TypeOf(__obf_50acae8c0a871ba1)).IsEmpty(reflect2.PtrOf(__obf_50acae8c0a871ba1))
}

type __obf_b4c25ee21eb89b02 struct {
	__obf_3915d58c01112789 *reflect2.UnsafeMapType
	__obf_4f90b2990512879d ValEncoder
	__obf_2b4ba80da80f8306 ValEncoder
}

func (__obf_c85b895560db528f *__obf_b4c25ee21eb89b02) Encode(__obf_753ab3fb72cdd659 unsafe.Pointer, __obf_2361f5214e84c60f *Stream) {
	if *(*unsafe.Pointer)(__obf_753ab3fb72cdd659) == nil {
		__obf_2361f5214e84c60f.
			WriteNil()
		return
	}
	__obf_2361f5214e84c60f.
		WriteObjectStart()
	__obf_edd9fe48d39445e4 := __obf_c85b895560db528f.__obf_3915d58c01112789.UnsafeIterate(__obf_753ab3fb72cdd659)
	for __obf_28d099df85f083a8 := 0; __obf_edd9fe48d39445e4.HasNext(); __obf_28d099df85f083a8++ {
		if __obf_28d099df85f083a8 != 0 {
			__obf_2361f5214e84c60f.
				WriteMore()
		}
		__obf_09e2475e3bcc63c4, __obf_991e715ce3758c78 := __obf_edd9fe48d39445e4.UnsafeNext()
		__obf_c85b895560db528f.__obf_4f90b2990512879d.
			Encode(__obf_09e2475e3bcc63c4, __obf_2361f5214e84c60f)
		if __obf_2361f5214e84c60f.__obf_7ccde3f02c81cd81 > 0 {
			__obf_2361f5214e84c60f.__obf_5e728551f00598e5(byte(':'), byte(' '))
		} else {
			__obf_2361f5214e84c60f.__obf_c4fec0edfb3875ad(':')
		}
		__obf_c85b895560db528f.__obf_2b4ba80da80f8306.
			Encode(__obf_991e715ce3758c78, __obf_2361f5214e84c60f)
	}
	__obf_2361f5214e84c60f.
		WriteObjectEnd()
}

func (__obf_c85b895560db528f *__obf_b4c25ee21eb89b02) IsEmpty(__obf_753ab3fb72cdd659 unsafe.Pointer) bool {
	__obf_edd9fe48d39445e4 := __obf_c85b895560db528f.__obf_3915d58c01112789.UnsafeIterate(__obf_753ab3fb72cdd659)
	return !__obf_edd9fe48d39445e4.HasNext()
}

type __obf_e06fd245f5efd38c struct {
	__obf_3915d58c01112789 *reflect2.UnsafeMapType
	__obf_4f90b2990512879d ValEncoder
	__obf_2b4ba80da80f8306 ValEncoder
}

func (__obf_c85b895560db528f *__obf_e06fd245f5efd38c) Encode(__obf_753ab3fb72cdd659 unsafe.Pointer, __obf_2361f5214e84c60f *Stream) {
	if *(*unsafe.Pointer)(__obf_753ab3fb72cdd659) == nil {
		__obf_2361f5214e84c60f.
			WriteNil()
		return
	}
	__obf_2361f5214e84c60f.
		WriteObjectStart()
	__obf_f19a793f1cf72d9e := __obf_c85b895560db528f.__obf_3915d58c01112789.UnsafeIterate(__obf_753ab3fb72cdd659)
	__obf_69f9feb9e74f95e8 := __obf_2361f5214e84c60f.__obf_0c8065c219ccf0ab.BorrowStream(nil)
	__obf_69f9feb9e74f95e8.
		Attachment = __obf_2361f5214e84c60f.Attachment
	__obf_e348a31983adaf66 := __obf_2361f5214e84c60f.__obf_0c8065c219ccf0ab.BorrowIterator(nil)
	__obf_ca4dfe9c97d70114 := __obf_1f1f8389aeed6181{}
	for __obf_f19a793f1cf72d9e.HasNext() {
		__obf_09e2475e3bcc63c4, __obf_991e715ce3758c78 := __obf_f19a793f1cf72d9e.UnsafeNext()
		__obf_8bd02256cde98ca6 := __obf_69f9feb9e74f95e8.Buffered()
		__obf_c85b895560db528f.__obf_4f90b2990512879d.
			Encode(__obf_09e2475e3bcc63c4, __obf_69f9feb9e74f95e8)
		if __obf_69f9feb9e74f95e8.Error != nil && __obf_69f9feb9e74f95e8.Error != io.EOF && __obf_2361f5214e84c60f.Error == nil {
			__obf_2361f5214e84c60f.
				Error = __obf_69f9feb9e74f95e8.Error
		}
		__obf_14c0ff3025a221b9 := __obf_69f9feb9e74f95e8.Buffer()[__obf_8bd02256cde98ca6:]
		__obf_e348a31983adaf66.
			ResetBytes(__obf_14c0ff3025a221b9)
		__obf_d4c9f69716253614 := __obf_e348a31983adaf66.ReadString()
		if __obf_2361f5214e84c60f.__obf_7ccde3f02c81cd81 > 0 {
			__obf_69f9feb9e74f95e8.__obf_5e728551f00598e5(byte(':'), byte(' '))
		} else {
			__obf_69f9feb9e74f95e8.__obf_c4fec0edfb3875ad(':')
		}
		__obf_c85b895560db528f.__obf_2b4ba80da80f8306.
			Encode(__obf_991e715ce3758c78, __obf_69f9feb9e74f95e8)
		__obf_ca4dfe9c97d70114 = append(__obf_ca4dfe9c97d70114, __obf_1e422500e8a12c19{__obf_09e2475e3bcc63c4: __obf_d4c9f69716253614, __obf_fa23b47bdefcc639: __obf_69f9feb9e74f95e8.Buffer()[__obf_8bd02256cde98ca6:]})
	}
	sort.Sort(__obf_ca4dfe9c97d70114)
	for __obf_28d099df85f083a8, __obf_fa23b47bdefcc639 := range __obf_ca4dfe9c97d70114 {
		if __obf_28d099df85f083a8 != 0 {
			__obf_2361f5214e84c60f.
				WriteMore()
		}
		__obf_2361f5214e84c60f.
			Write(__obf_fa23b47bdefcc639.__obf_fa23b47bdefcc639)
	}
	if __obf_69f9feb9e74f95e8.Error != nil && __obf_2361f5214e84c60f.Error == nil {
		__obf_2361f5214e84c60f.
			Error = __obf_69f9feb9e74f95e8.Error
	}
	__obf_2361f5214e84c60f.
		WriteObjectEnd()
	__obf_2361f5214e84c60f.__obf_0c8065c219ccf0ab.
		ReturnStream(__obf_69f9feb9e74f95e8)
	__obf_2361f5214e84c60f.__obf_0c8065c219ccf0ab.
		ReturnIterator(__obf_e348a31983adaf66)
}

func (__obf_c85b895560db528f *__obf_e06fd245f5efd38c) IsEmpty(__obf_753ab3fb72cdd659 unsafe.Pointer) bool {
	__obf_edd9fe48d39445e4 := __obf_c85b895560db528f.__obf_3915d58c01112789.UnsafeIterate(__obf_753ab3fb72cdd659)
	return !__obf_edd9fe48d39445e4.HasNext()
}

type __obf_1f1f8389aeed6181 []__obf_1e422500e8a12c19

type __obf_1e422500e8a12c19 struct {
	__obf_09e2475e3bcc63c4 string
	__obf_fa23b47bdefcc639 []byte
}

func (__obf_0a2b808592c93074 __obf_1f1f8389aeed6181) Len() int { return len(__obf_0a2b808592c93074) }
func (__obf_0a2b808592c93074 __obf_1f1f8389aeed6181) Swap(__obf_28d099df85f083a8, __obf_57cd9c621e7db075 int) {
	__obf_0a2b808592c93074[__obf_28d099df85f083a8], __obf_0a2b808592c93074[__obf_57cd9c621e7db075] = __obf_0a2b808592c93074[__obf_57cd9c621e7db075], __obf_0a2b808592c93074[__obf_28d099df85f083a8]
}
func (__obf_0a2b808592c93074 __obf_1f1f8389aeed6181) Less(__obf_28d099df85f083a8, __obf_57cd9c621e7db075 int) bool {
	return __obf_0a2b808592c93074[__obf_28d099df85f083a8].__obf_09e2475e3bcc63c4 < __obf_0a2b808592c93074[__obf_57cd9c621e7db075].__obf_09e2475e3bcc63c4
}
