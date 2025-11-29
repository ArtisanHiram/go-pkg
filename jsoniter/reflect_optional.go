package __obf_5b802ce8d9ba56d6

import (
	"github.com/modern-go/reflect2"
	"unsafe"
)

func __obf_cb0144ac2d5afb8d(__obf_08da24be66d0d13c *__obf_08da24be66d0d13c, __obf_5efc66d8c338c133 reflect2.Type) ValDecoder {
	__obf_d0cac7bfcf0092ea := __obf_5efc66d8c338c133.(*reflect2.UnsafePtrType)
	__obf_0b19d2b7ea8e1be5 := __obf_d0cac7bfcf0092ea.Elem()
	__obf_18f746d7b5b440ee := __obf_c3a46fc9dd10c84e(__obf_08da24be66d0d13c, __obf_0b19d2b7ea8e1be5)
	return &OptionalDecoder{__obf_0b19d2b7ea8e1be5, __obf_18f746d7b5b440ee}
}

func __obf_7894a79375435eef(__obf_08da24be66d0d13c *__obf_08da24be66d0d13c, __obf_5efc66d8c338c133 reflect2.Type) ValEncoder {
	__obf_d0cac7bfcf0092ea := __obf_5efc66d8c338c133.(*reflect2.UnsafePtrType)
	__obf_0b19d2b7ea8e1be5 := __obf_d0cac7bfcf0092ea.Elem()
	__obf_e57e229f9a2faf9a := __obf_3bb380f7abc03208(__obf_08da24be66d0d13c, __obf_0b19d2b7ea8e1be5)
	__obf_29366c3ad76a8c51 := &OptionalEncoder{__obf_e57e229f9a2faf9a}
	return __obf_29366c3ad76a8c51
}

type OptionalDecoder struct {
	ValueType    reflect2.Type
	ValueDecoder ValDecoder
}

func (__obf_18f746d7b5b440ee *OptionalDecoder) Decode(__obf_d3c919547bf7e47a unsafe.Pointer, __obf_67008a6a9e5ba828 *Iterator) {
	if __obf_67008a6a9e5ba828.ReadNil() {
		*((*unsafe.Pointer)(__obf_d3c919547bf7e47a)) = nil
	} else {
		if *((*unsafe.Pointer)(__obf_d3c919547bf7e47a)) == nil {
			__obf_74d5f8c22a26960b := //pointer to null, we have to allocate memory to hold the value
				__obf_18f746d7b5b440ee.ValueType.UnsafeNew()
			__obf_18f746d7b5b440ee.
				ValueDecoder.Decode(__obf_74d5f8c22a26960b, __obf_67008a6a9e5ba828)
			*((*unsafe.Pointer)(__obf_d3c919547bf7e47a)) = __obf_74d5f8c22a26960b
		} else {
			__obf_18f746d7b5b440ee.
				//reuse existing instance
				ValueDecoder.Decode(*((*unsafe.Pointer)(__obf_d3c919547bf7e47a)), __obf_67008a6a9e5ba828)
		}
	}
}

type __obf_7ad15af757dc337b struct {
	__obf_9ad2e389a2a6fdb8 reflect2. // only to deference a pointer
				Type
	__obf_49b56423e80a9797 ValDecoder
}

func (__obf_18f746d7b5b440ee *__obf_7ad15af757dc337b) Decode(__obf_d3c919547bf7e47a unsafe.Pointer, __obf_67008a6a9e5ba828 *Iterator) {
	if *((*unsafe.Pointer)(__obf_d3c919547bf7e47a)) == nil {
		__obf_74d5f8c22a26960b := //pointer to null, we have to allocate memory to hold the value
			__obf_18f746d7b5b440ee.__obf_9ad2e389a2a6fdb8.UnsafeNew()
		__obf_18f746d7b5b440ee.__obf_49b56423e80a9797.
			Decode(__obf_74d5f8c22a26960b, __obf_67008a6a9e5ba828)
		*((*unsafe.Pointer)(__obf_d3c919547bf7e47a)) = __obf_74d5f8c22a26960b
	} else {
		__obf_18f746d7b5b440ee.
			//reuse existing instance
			__obf_49b56423e80a9797.
			Decode(*((*unsafe.Pointer)(__obf_d3c919547bf7e47a)), __obf_67008a6a9e5ba828)
	}
}

type OptionalEncoder struct {
	ValueEncoder ValEncoder
}

func (__obf_29366c3ad76a8c51 *OptionalEncoder) Encode(__obf_d3c919547bf7e47a unsafe.Pointer, __obf_00fc491caa967a74 *Stream) {
	if *((*unsafe.Pointer)(__obf_d3c919547bf7e47a)) == nil {
		__obf_00fc491caa967a74.
			WriteNil()
	} else {
		__obf_29366c3ad76a8c51.
			ValueEncoder.Encode(*((*unsafe.Pointer)(__obf_d3c919547bf7e47a)), __obf_00fc491caa967a74)
	}
}

func (__obf_29366c3ad76a8c51 *OptionalEncoder) IsEmpty(__obf_d3c919547bf7e47a unsafe.Pointer) bool {
	return *((*unsafe.Pointer)(__obf_d3c919547bf7e47a)) == nil
}

type __obf_b7df4ea38882225e struct {
	ValueEncoder ValEncoder
}

func (__obf_29366c3ad76a8c51 *__obf_b7df4ea38882225e) Encode(__obf_d3c919547bf7e47a unsafe.Pointer, __obf_00fc491caa967a74 *Stream) {
	if *((*unsafe.Pointer)(__obf_d3c919547bf7e47a)) == nil {
		__obf_00fc491caa967a74.
			WriteNil()
	} else {
		__obf_29366c3ad76a8c51.
			ValueEncoder.Encode(*((*unsafe.Pointer)(__obf_d3c919547bf7e47a)), __obf_00fc491caa967a74)
	}
}

func (__obf_29366c3ad76a8c51 *__obf_b7df4ea38882225e) IsEmpty(__obf_d3c919547bf7e47a unsafe.Pointer) bool {
	__obf_43dffb9252574488 := *((*unsafe.Pointer)(__obf_d3c919547bf7e47a))
	if __obf_43dffb9252574488 == nil {
		return true
	}
	return __obf_29366c3ad76a8c51.ValueEncoder.IsEmpty(__obf_43dffb9252574488)
}

func (__obf_29366c3ad76a8c51 *__obf_b7df4ea38882225e) IsEmbeddedPtrNil(__obf_d3c919547bf7e47a unsafe.Pointer) bool {
	__obf_4080b42e3a615d8e := *((*unsafe.Pointer)(__obf_d3c919547bf7e47a))
	if __obf_4080b42e3a615d8e == nil {
		return true
	}
	__obf_8fcef6e2878720af, __obf_ecd3f2241015639e := __obf_29366c3ad76a8c51.ValueEncoder.(IsEmbeddedPtrNil)
	if !__obf_ecd3f2241015639e {
		return false
	}
	__obf_a6502f11c1c12aab := unsafe.Pointer(__obf_4080b42e3a615d8e)
	return __obf_8fcef6e2878720af.IsEmbeddedPtrNil(__obf_a6502f11c1c12aab)
}

type __obf_f668640819fd46fc struct {
	__obf_29366c3ad76a8c51 ValEncoder
}

func (__obf_29366c3ad76a8c51 *__obf_f668640819fd46fc) Encode(__obf_d3c919547bf7e47a unsafe.Pointer, __obf_00fc491caa967a74 *Stream) {
	__obf_29366c3ad76a8c51.__obf_29366c3ad76a8c51.
		Encode(unsafe.Pointer(&__obf_d3c919547bf7e47a), __obf_00fc491caa967a74)
}

func (__obf_29366c3ad76a8c51 *__obf_f668640819fd46fc) IsEmpty(__obf_d3c919547bf7e47a unsafe.Pointer) bool {
	return __obf_29366c3ad76a8c51.__obf_29366c3ad76a8c51.IsEmpty(unsafe.Pointer(&__obf_d3c919547bf7e47a))
}

type __obf_46f8310df6424f3f struct {
	__obf_18f746d7b5b440ee ValDecoder
}

func (__obf_18f746d7b5b440ee *__obf_46f8310df6424f3f) Decode(__obf_d3c919547bf7e47a unsafe.Pointer, __obf_67008a6a9e5ba828 *Iterator) {
	__obf_18f746d7b5b440ee.__obf_18f746d7b5b440ee.
		Decode(unsafe.Pointer(&__obf_d3c919547bf7e47a), __obf_67008a6a9e5ba828)
}
