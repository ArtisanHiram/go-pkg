package __obf_5b802ce8d9ba56d6

import (
	"fmt"
	"github.com/modern-go/reflect2"
	"io"
	"unsafe"
)

func __obf_868f472bae9503f0(__obf_08da24be66d0d13c *__obf_08da24be66d0d13c, __obf_5efc66d8c338c133 reflect2.Type) ValDecoder {
	__obf_2aefb0cc09305308 := __obf_5efc66d8c338c133.(*reflect2.UnsafeSliceType)
	__obf_18f746d7b5b440ee := __obf_c3a46fc9dd10c84e(__obf_08da24be66d0d13c.append("[sliceElem]"), __obf_2aefb0cc09305308.Elem())
	return &__obf_fae02aecb00c25f9{__obf_2aefb0cc09305308, __obf_18f746d7b5b440ee}
}

func __obf_0dab538fd3e96275(__obf_08da24be66d0d13c *__obf_08da24be66d0d13c, __obf_5efc66d8c338c133 reflect2.Type) ValEncoder {
	__obf_2aefb0cc09305308 := __obf_5efc66d8c338c133.(*reflect2.UnsafeSliceType)
	__obf_29366c3ad76a8c51 := __obf_3bb380f7abc03208(__obf_08da24be66d0d13c.append("[sliceElem]"), __obf_2aefb0cc09305308.Elem())
	return &__obf_47f54f698749f41b{__obf_2aefb0cc09305308, __obf_29366c3ad76a8c51}
}

type __obf_47f54f698749f41b struct {
	__obf_2aefb0cc09305308 *reflect2.UnsafeSliceType
	__obf_e57e229f9a2faf9a ValEncoder
}

func (__obf_29366c3ad76a8c51 *__obf_47f54f698749f41b) Encode(__obf_d3c919547bf7e47a unsafe.Pointer, __obf_00fc491caa967a74 *Stream) {
	if __obf_29366c3ad76a8c51.__obf_2aefb0cc09305308.UnsafeIsNil(__obf_d3c919547bf7e47a) {
		__obf_00fc491caa967a74.
			WriteNil()
		return
	}
	__obf_b1248a9e4497b06b := __obf_29366c3ad76a8c51.__obf_2aefb0cc09305308.UnsafeLengthOf(__obf_d3c919547bf7e47a)
	if __obf_b1248a9e4497b06b == 0 {
		__obf_00fc491caa967a74.
			WriteEmptyArray()
		return
	}
	__obf_00fc491caa967a74.
		WriteArrayStart()
	__obf_29366c3ad76a8c51.__obf_e57e229f9a2faf9a.
		Encode(__obf_29366c3ad76a8c51.__obf_2aefb0cc09305308.UnsafeGetIndex(__obf_d3c919547bf7e47a, 0), __obf_00fc491caa967a74)
	for __obf_2deec7c38ffb6ae3 := 1; __obf_2deec7c38ffb6ae3 < __obf_b1248a9e4497b06b; __obf_2deec7c38ffb6ae3++ {
		__obf_00fc491caa967a74.
			WriteMore()
		__obf_7a37ccab267895e6 := __obf_29366c3ad76a8c51.__obf_2aefb0cc09305308.UnsafeGetIndex(__obf_d3c919547bf7e47a, __obf_2deec7c38ffb6ae3)
		__obf_29366c3ad76a8c51.__obf_e57e229f9a2faf9a.
			Encode(__obf_7a37ccab267895e6, __obf_00fc491caa967a74)
	}
	__obf_00fc491caa967a74.
		WriteArrayEnd()
	if __obf_00fc491caa967a74.Error != nil && __obf_00fc491caa967a74.Error != io.EOF {
		__obf_00fc491caa967a74.
			Error = fmt.Errorf("%v: %s", __obf_29366c3ad76a8c51.__obf_2aefb0cc09305308, __obf_00fc491caa967a74.Error.Error())
	}
}

func (__obf_29366c3ad76a8c51 *__obf_47f54f698749f41b) IsEmpty(__obf_d3c919547bf7e47a unsafe.Pointer) bool {
	return __obf_29366c3ad76a8c51.__obf_2aefb0cc09305308.UnsafeLengthOf(__obf_d3c919547bf7e47a) == 0
}

type __obf_fae02aecb00c25f9 struct {
	__obf_2aefb0cc09305308 *reflect2.UnsafeSliceType
	__obf_45aaf5212411947b ValDecoder
}

func (__obf_18f746d7b5b440ee *__obf_fae02aecb00c25f9) Decode(__obf_d3c919547bf7e47a unsafe.Pointer, __obf_67008a6a9e5ba828 *Iterator) {
	__obf_18f746d7b5b440ee.__obf_95ea2b5b23c8db5c(__obf_d3c919547bf7e47a, __obf_67008a6a9e5ba828)
	if __obf_67008a6a9e5ba828.Error != nil && __obf_67008a6a9e5ba828.Error != io.EOF {
		__obf_67008a6a9e5ba828.
			Error = fmt.Errorf("%v: %s", __obf_18f746d7b5b440ee.__obf_2aefb0cc09305308, __obf_67008a6a9e5ba828.Error.Error())
	}
}

func (__obf_18f746d7b5b440ee *__obf_fae02aecb00c25f9) __obf_95ea2b5b23c8db5c(__obf_d3c919547bf7e47a unsafe.Pointer, __obf_67008a6a9e5ba828 *Iterator) {
	__obf_dab9baaadfa7c8c2 := __obf_67008a6a9e5ba828.__obf_b781a59d5a0d2490()
	__obf_2aefb0cc09305308 := __obf_18f746d7b5b440ee.__obf_2aefb0cc09305308
	if __obf_dab9baaadfa7c8c2 == 'n' {
		__obf_67008a6a9e5ba828.__obf_4aeb767e0be7277a('u', 'l', 'l')
		__obf_2aefb0cc09305308.
			UnsafeSetNil(__obf_d3c919547bf7e47a)
		return
	}
	if __obf_dab9baaadfa7c8c2 != '[' {
		__obf_67008a6a9e5ba828.
			ReportError("decode slice", "expect [ or n, but found "+string([]byte{__obf_dab9baaadfa7c8c2}))
		return
	}
	__obf_dab9baaadfa7c8c2 = __obf_67008a6a9e5ba828.__obf_b781a59d5a0d2490()
	if __obf_dab9baaadfa7c8c2 == ']' {
		__obf_2aefb0cc09305308.
			UnsafeSet(__obf_d3c919547bf7e47a, __obf_2aefb0cc09305308.UnsafeMakeSlice(0, 0))
		return
	}
	__obf_67008a6a9e5ba828.__obf_3284a1eaa2a0abb6()
	__obf_2aefb0cc09305308.
		UnsafeGrow(__obf_d3c919547bf7e47a, 1)
	__obf_7a37ccab267895e6 := __obf_2aefb0cc09305308.UnsafeGetIndex(__obf_d3c919547bf7e47a, 0)
	__obf_18f746d7b5b440ee.__obf_45aaf5212411947b.
		Decode(__obf_7a37ccab267895e6, __obf_67008a6a9e5ba828)
	__obf_b1248a9e4497b06b := 1
	for __obf_dab9baaadfa7c8c2 = __obf_67008a6a9e5ba828.__obf_b781a59d5a0d2490(); __obf_dab9baaadfa7c8c2 == ','; __obf_dab9baaadfa7c8c2 = __obf_67008a6a9e5ba828.__obf_b781a59d5a0d2490() {
		__obf_e9657e567969f634 := __obf_b1248a9e4497b06b
		__obf_b1248a9e4497b06b += 1
		__obf_2aefb0cc09305308.
			UnsafeGrow(__obf_d3c919547bf7e47a, __obf_b1248a9e4497b06b)
		__obf_7a37ccab267895e6 = __obf_2aefb0cc09305308.UnsafeGetIndex(__obf_d3c919547bf7e47a, __obf_e9657e567969f634)
		__obf_18f746d7b5b440ee.__obf_45aaf5212411947b.
			Decode(__obf_7a37ccab267895e6, __obf_67008a6a9e5ba828)
	}
	if __obf_dab9baaadfa7c8c2 != ']' {
		__obf_67008a6a9e5ba828.
			ReportError("decode slice", "expect ], but found "+string([]byte{__obf_dab9baaadfa7c8c2}))
		return
	}
}
