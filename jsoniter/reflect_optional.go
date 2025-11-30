package __obf_c3cd04a15c56f32f

import (
	"github.com/modern-go/reflect2"
	"unsafe"
)

func __obf_272a3d13803d4851(__obf_62435d89ebefd5aa *__obf_62435d89ebefd5aa, __obf_8667d4fc2a95b0d7 reflect2.Type) ValDecoder {
	__obf_9fd3068ebc25d7f9 := __obf_8667d4fc2a95b0d7.(*reflect2.UnsafePtrType)
	__obf_78b97e7cf83971dd := __obf_9fd3068ebc25d7f9.Elem()
	__obf_924cc7ef585cdfb0 := __obf_eddb00a5736b0fe7(__obf_62435d89ebefd5aa, __obf_78b97e7cf83971dd)
	return &OptionalDecoder{__obf_78b97e7cf83971dd, __obf_924cc7ef585cdfb0}
}

func __obf_d2cf57edf2ac20e2(__obf_62435d89ebefd5aa *__obf_62435d89ebefd5aa, __obf_8667d4fc2a95b0d7 reflect2.Type) ValEncoder {
	__obf_9fd3068ebc25d7f9 := __obf_8667d4fc2a95b0d7.(*reflect2.UnsafePtrType)
	__obf_78b97e7cf83971dd := __obf_9fd3068ebc25d7f9.Elem()
	__obf_2b4ba80da80f8306 := __obf_dd4ab965a9fde81c(__obf_62435d89ebefd5aa, __obf_78b97e7cf83971dd)
	__obf_c85b895560db528f := &OptionalEncoder{__obf_2b4ba80da80f8306}
	return __obf_c85b895560db528f
}

type OptionalDecoder struct {
	ValueType    reflect2.Type
	ValueDecoder ValDecoder
}

func (__obf_924cc7ef585cdfb0 *OptionalDecoder) Decode(__obf_753ab3fb72cdd659 unsafe.Pointer, __obf_edd9fe48d39445e4 *Iterator) {
	if __obf_edd9fe48d39445e4.ReadNil() {
		*((*unsafe.Pointer)(__obf_753ab3fb72cdd659)) = nil
	} else {
		if *((*unsafe.Pointer)(__obf_753ab3fb72cdd659)) == nil {
			__obf_25685aaa7e50c763 := //pointer to null, we have to allocate memory to hold the value
				__obf_924cc7ef585cdfb0.ValueType.UnsafeNew()
			__obf_924cc7ef585cdfb0.
				ValueDecoder.Decode(__obf_25685aaa7e50c763, __obf_edd9fe48d39445e4)
			*((*unsafe.Pointer)(__obf_753ab3fb72cdd659)) = __obf_25685aaa7e50c763
		} else {
			__obf_924cc7ef585cdfb0.
				//reuse existing instance
				ValueDecoder.Decode(*((*unsafe.Pointer)(__obf_753ab3fb72cdd659)), __obf_edd9fe48d39445e4)
		}
	}
}

type __obf_ff8360cd52b76a62 struct {
	__obf_64fa966d17d74b1c reflect2. // only to deference a pointer
				Type
	__obf_67383beae8dd522e ValDecoder
}

func (__obf_924cc7ef585cdfb0 *__obf_ff8360cd52b76a62) Decode(__obf_753ab3fb72cdd659 unsafe.Pointer, __obf_edd9fe48d39445e4 *Iterator) {
	if *((*unsafe.Pointer)(__obf_753ab3fb72cdd659)) == nil {
		__obf_25685aaa7e50c763 := //pointer to null, we have to allocate memory to hold the value
			__obf_924cc7ef585cdfb0.__obf_64fa966d17d74b1c.UnsafeNew()
		__obf_924cc7ef585cdfb0.__obf_67383beae8dd522e.
			Decode(__obf_25685aaa7e50c763, __obf_edd9fe48d39445e4)
		*((*unsafe.Pointer)(__obf_753ab3fb72cdd659)) = __obf_25685aaa7e50c763
	} else {
		__obf_924cc7ef585cdfb0.
			//reuse existing instance
			__obf_67383beae8dd522e.
			Decode(*((*unsafe.Pointer)(__obf_753ab3fb72cdd659)), __obf_edd9fe48d39445e4)
	}
}

type OptionalEncoder struct {
	ValueEncoder ValEncoder
}

func (__obf_c85b895560db528f *OptionalEncoder) Encode(__obf_753ab3fb72cdd659 unsafe.Pointer, __obf_2361f5214e84c60f *Stream) {
	if *((*unsafe.Pointer)(__obf_753ab3fb72cdd659)) == nil {
		__obf_2361f5214e84c60f.
			WriteNil()
	} else {
		__obf_c85b895560db528f.
			ValueEncoder.Encode(*((*unsafe.Pointer)(__obf_753ab3fb72cdd659)), __obf_2361f5214e84c60f)
	}
}

func (__obf_c85b895560db528f *OptionalEncoder) IsEmpty(__obf_753ab3fb72cdd659 unsafe.Pointer) bool {
	return *((*unsafe.Pointer)(__obf_753ab3fb72cdd659)) == nil
}

type __obf_09858d643dd5d97b struct {
	ValueEncoder ValEncoder
}

func (__obf_c85b895560db528f *__obf_09858d643dd5d97b) Encode(__obf_753ab3fb72cdd659 unsafe.Pointer, __obf_2361f5214e84c60f *Stream) {
	if *((*unsafe.Pointer)(__obf_753ab3fb72cdd659)) == nil {
		__obf_2361f5214e84c60f.
			WriteNil()
	} else {
		__obf_c85b895560db528f.
			ValueEncoder.Encode(*((*unsafe.Pointer)(__obf_753ab3fb72cdd659)), __obf_2361f5214e84c60f)
	}
}

func (__obf_c85b895560db528f *__obf_09858d643dd5d97b) IsEmpty(__obf_753ab3fb72cdd659 unsafe.Pointer) bool {
	__obf_35d0a8402b564ad6 := *((*unsafe.Pointer)(__obf_753ab3fb72cdd659))
	if __obf_35d0a8402b564ad6 == nil {
		return true
	}
	return __obf_c85b895560db528f.ValueEncoder.IsEmpty(__obf_35d0a8402b564ad6)
}

func (__obf_c85b895560db528f *__obf_09858d643dd5d97b) IsEmbeddedPtrNil(__obf_753ab3fb72cdd659 unsafe.Pointer) bool {
	__obf_fa11a68748b91bc9 := *((*unsafe.Pointer)(__obf_753ab3fb72cdd659))
	if __obf_fa11a68748b91bc9 == nil {
		return true
	}
	__obf_defaafd546a4fefb, __obf_315a503b16d9c8a0 := __obf_c85b895560db528f.ValueEncoder.(IsEmbeddedPtrNil)
	if !__obf_315a503b16d9c8a0 {
		return false
	}
	__obf_f4916a65ae1a5a5a := unsafe.Pointer(__obf_fa11a68748b91bc9)
	return __obf_defaafd546a4fefb.IsEmbeddedPtrNil(__obf_f4916a65ae1a5a5a)
}

type __obf_19c8951c4dfee8a0 struct {
	__obf_c85b895560db528f ValEncoder
}

func (__obf_c85b895560db528f *__obf_19c8951c4dfee8a0) Encode(__obf_753ab3fb72cdd659 unsafe.Pointer, __obf_2361f5214e84c60f *Stream) {
	__obf_c85b895560db528f.__obf_c85b895560db528f.
		Encode(unsafe.Pointer(&__obf_753ab3fb72cdd659), __obf_2361f5214e84c60f)
}

func (__obf_c85b895560db528f *__obf_19c8951c4dfee8a0) IsEmpty(__obf_753ab3fb72cdd659 unsafe.Pointer) bool {
	return __obf_c85b895560db528f.__obf_c85b895560db528f.IsEmpty(unsafe.Pointer(&__obf_753ab3fb72cdd659))
}

type __obf_cec111c41697b8e9 struct {
	__obf_924cc7ef585cdfb0 ValDecoder
}

func (__obf_924cc7ef585cdfb0 *__obf_cec111c41697b8e9) Decode(__obf_753ab3fb72cdd659 unsafe.Pointer, __obf_edd9fe48d39445e4 *Iterator) {
	__obf_924cc7ef585cdfb0.__obf_924cc7ef585cdfb0.
		Decode(unsafe.Pointer(&__obf_753ab3fb72cdd659), __obf_edd9fe48d39445e4)
}
