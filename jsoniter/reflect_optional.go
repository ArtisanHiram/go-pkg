package __obf_703d23ba520c3295

import (
	"github.com/modern-go/reflect2"
	"unsafe"
)

func __obf_490e149853cb1478(__obf_2aaf7367de3ff86d *__obf_2aaf7367de3ff86d, __obf_3b7f6abbae19451e reflect2.Type) ValDecoder {
	__obf_95093efb9321684e := __obf_3b7f6abbae19451e.(*reflect2.UnsafePtrType)
	__obf_8b28c1d23e9d3881 := __obf_95093efb9321684e.Elem()
	__obf_c9719e68325f7d44 := __obf_b27fe2bc17d94621(__obf_2aaf7367de3ff86d, __obf_8b28c1d23e9d3881)
	return &OptionalDecoder{__obf_8b28c1d23e9d3881, __obf_c9719e68325f7d44}
}

func __obf_9a820bfca585fe9d(__obf_2aaf7367de3ff86d *__obf_2aaf7367de3ff86d, __obf_3b7f6abbae19451e reflect2.Type) ValEncoder {
	__obf_95093efb9321684e := __obf_3b7f6abbae19451e.(*reflect2.UnsafePtrType)
	__obf_8b28c1d23e9d3881 := __obf_95093efb9321684e.Elem()
	__obf_8bb551383a0ba60c := __obf_906496c658b349cf(__obf_2aaf7367de3ff86d, __obf_8b28c1d23e9d3881)
	__obf_cf840243a12ee308 := &OptionalEncoder{__obf_8bb551383a0ba60c}
	return __obf_cf840243a12ee308
}

type OptionalDecoder struct {
	ValueType    reflect2.Type
	ValueDecoder ValDecoder
}

func (__obf_c9719e68325f7d44 *OptionalDecoder) Decode(__obf_47fa53f3d161f45c unsafe.Pointer, __obf_47edab4c16a2d88d *Iterator) {
	if __obf_47edab4c16a2d88d.ReadNil() {
		*((*unsafe.Pointer)(__obf_47fa53f3d161f45c)) = nil
	} else {
		if *((*unsafe.Pointer)(__obf_47fa53f3d161f45c)) == nil {
			__obf_69a4174ebbf7fc7a := //pointer to null, we have to allocate memory to hold the value
				__obf_c9719e68325f7d44.ValueType.UnsafeNew()
			__obf_c9719e68325f7d44.
				ValueDecoder.Decode(__obf_69a4174ebbf7fc7a, __obf_47edab4c16a2d88d)
			*((*unsafe.Pointer)(__obf_47fa53f3d161f45c)) = __obf_69a4174ebbf7fc7a
		} else {
			__obf_c9719e68325f7d44.
				//reuse existing instance
				ValueDecoder.Decode(*((*unsafe.Pointer)(__obf_47fa53f3d161f45c)), __obf_47edab4c16a2d88d)
		}
	}
}

type __obf_af91af019702e067 struct {
	__obf_835102d74dbf01c6 reflect2. // only to deference a pointer
				Type
	__obf_dd9a1c6e4d240b40 ValDecoder
}

func (__obf_c9719e68325f7d44 *__obf_af91af019702e067) Decode(__obf_47fa53f3d161f45c unsafe.Pointer, __obf_47edab4c16a2d88d *Iterator) {
	if *((*unsafe.Pointer)(__obf_47fa53f3d161f45c)) == nil {
		__obf_69a4174ebbf7fc7a := //pointer to null, we have to allocate memory to hold the value
			__obf_c9719e68325f7d44.__obf_835102d74dbf01c6.UnsafeNew()
		__obf_c9719e68325f7d44.__obf_dd9a1c6e4d240b40.
			Decode(__obf_69a4174ebbf7fc7a, __obf_47edab4c16a2d88d)
		*((*unsafe.Pointer)(__obf_47fa53f3d161f45c)) = __obf_69a4174ebbf7fc7a
	} else {
		__obf_c9719e68325f7d44.
			//reuse existing instance
			__obf_dd9a1c6e4d240b40.
			Decode(*((*unsafe.Pointer)(__obf_47fa53f3d161f45c)), __obf_47edab4c16a2d88d)
	}
}

type OptionalEncoder struct {
	ValueEncoder ValEncoder
}

func (__obf_cf840243a12ee308 *OptionalEncoder) Encode(__obf_47fa53f3d161f45c unsafe.Pointer, __obf_9a34f62857fb1d1d *Stream) {
	if *((*unsafe.Pointer)(__obf_47fa53f3d161f45c)) == nil {
		__obf_9a34f62857fb1d1d.
			WriteNil()
	} else {
		__obf_cf840243a12ee308.
			ValueEncoder.Encode(*((*unsafe.Pointer)(__obf_47fa53f3d161f45c)), __obf_9a34f62857fb1d1d)
	}
}

func (__obf_cf840243a12ee308 *OptionalEncoder) IsEmpty(__obf_47fa53f3d161f45c unsafe.Pointer) bool {
	return *((*unsafe.Pointer)(__obf_47fa53f3d161f45c)) == nil
}

type __obf_d6bbd0845c4e0abb struct {
	ValueEncoder ValEncoder
}

func (__obf_cf840243a12ee308 *__obf_d6bbd0845c4e0abb) Encode(__obf_47fa53f3d161f45c unsafe.Pointer, __obf_9a34f62857fb1d1d *Stream) {
	if *((*unsafe.Pointer)(__obf_47fa53f3d161f45c)) == nil {
		__obf_9a34f62857fb1d1d.
			WriteNil()
	} else {
		__obf_cf840243a12ee308.
			ValueEncoder.Encode(*((*unsafe.Pointer)(__obf_47fa53f3d161f45c)), __obf_9a34f62857fb1d1d)
	}
}

func (__obf_cf840243a12ee308 *__obf_d6bbd0845c4e0abb) IsEmpty(__obf_47fa53f3d161f45c unsafe.Pointer) bool {
	__obf_2627aca83d15d7dc := *((*unsafe.Pointer)(__obf_47fa53f3d161f45c))
	if __obf_2627aca83d15d7dc == nil {
		return true
	}
	return __obf_cf840243a12ee308.ValueEncoder.IsEmpty(__obf_2627aca83d15d7dc)
}

func (__obf_cf840243a12ee308 *__obf_d6bbd0845c4e0abb) IsEmbeddedPtrNil(__obf_47fa53f3d161f45c unsafe.Pointer) bool {
	__obf_86ca9ddfc9ef02b6 := *((*unsafe.Pointer)(__obf_47fa53f3d161f45c))
	if __obf_86ca9ddfc9ef02b6 == nil {
		return true
	}
	__obf_0619e0ea83732b1b, __obf_22ffff425cd0177f := __obf_cf840243a12ee308.ValueEncoder.(IsEmbeddedPtrNil)
	if !__obf_22ffff425cd0177f {
		return false
	}
	__obf_94be1fdf2ec0d598 := unsafe.Pointer(__obf_86ca9ddfc9ef02b6)
	return __obf_0619e0ea83732b1b.IsEmbeddedPtrNil(__obf_94be1fdf2ec0d598)
}

type __obf_1de20804cb7c71ec struct {
	__obf_cf840243a12ee308 ValEncoder
}

func (__obf_cf840243a12ee308 *__obf_1de20804cb7c71ec) Encode(__obf_47fa53f3d161f45c unsafe.Pointer, __obf_9a34f62857fb1d1d *Stream) {
	__obf_cf840243a12ee308.__obf_cf840243a12ee308.
		Encode(unsafe.Pointer(&__obf_47fa53f3d161f45c), __obf_9a34f62857fb1d1d)
}

func (__obf_cf840243a12ee308 *__obf_1de20804cb7c71ec) IsEmpty(__obf_47fa53f3d161f45c unsafe.Pointer) bool {
	return __obf_cf840243a12ee308.__obf_cf840243a12ee308.IsEmpty(unsafe.Pointer(&__obf_47fa53f3d161f45c))
}

type __obf_8eed859ba47bd05c struct {
	__obf_c9719e68325f7d44 ValDecoder
}

func (__obf_c9719e68325f7d44 *__obf_8eed859ba47bd05c) Decode(__obf_47fa53f3d161f45c unsafe.Pointer, __obf_47edab4c16a2d88d *Iterator) {
	__obf_c9719e68325f7d44.__obf_c9719e68325f7d44.
		Decode(unsafe.Pointer(&__obf_47fa53f3d161f45c), __obf_47edab4c16a2d88d)
}
