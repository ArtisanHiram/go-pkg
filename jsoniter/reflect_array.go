package __obf_5b802ce8d9ba56d6

import (
	"fmt"
	"github.com/modern-go/reflect2"
	"io"
	"unsafe"
)

func __obf_ee5a34160cacf7d6(__obf_08da24be66d0d13c *__obf_08da24be66d0d13c, __obf_5efc66d8c338c133 reflect2.Type) ValDecoder {
	__obf_1060c952158f31ac := __obf_5efc66d8c338c133.(*reflect2.UnsafeArrayType)
	__obf_18f746d7b5b440ee := __obf_c3a46fc9dd10c84e(__obf_08da24be66d0d13c.append("[arrayElem]"), __obf_1060c952158f31ac.Elem())
	return &__obf_eab64d10d33e7ed7{__obf_1060c952158f31ac, __obf_18f746d7b5b440ee}
}

func __obf_d2984c9ef31dac9d(__obf_08da24be66d0d13c *__obf_08da24be66d0d13c, __obf_5efc66d8c338c133 reflect2.Type) ValEncoder {
	__obf_1060c952158f31ac := __obf_5efc66d8c338c133.(*reflect2.UnsafeArrayType)
	if __obf_1060c952158f31ac.Len() == 0 {
		return __obf_f805be59b9f4911f{}
	}
	__obf_29366c3ad76a8c51 := __obf_3bb380f7abc03208(__obf_08da24be66d0d13c.append("[arrayElem]"), __obf_1060c952158f31ac.Elem())
	return &__obf_55abf1f4a08610a6{__obf_1060c952158f31ac, __obf_29366c3ad76a8c51}
}

type __obf_f805be59b9f4911f struct{}

func (__obf_29366c3ad76a8c51 __obf_f805be59b9f4911f) Encode(__obf_d3c919547bf7e47a unsafe.Pointer, __obf_00fc491caa967a74 *Stream) {
	__obf_00fc491caa967a74.
		WriteEmptyArray()
}

func (__obf_29366c3ad76a8c51 __obf_f805be59b9f4911f) IsEmpty(__obf_d3c919547bf7e47a unsafe.Pointer) bool {
	return true
}

type __obf_55abf1f4a08610a6 struct {
	__obf_1060c952158f31ac *reflect2.UnsafeArrayType
	__obf_e57e229f9a2faf9a ValEncoder
}

func (__obf_29366c3ad76a8c51 *__obf_55abf1f4a08610a6) Encode(__obf_d3c919547bf7e47a unsafe.Pointer, __obf_00fc491caa967a74 *Stream) {
	__obf_00fc491caa967a74.
		WriteArrayStart()
	__obf_7a37ccab267895e6 := unsafe.Pointer(__obf_d3c919547bf7e47a)
	__obf_29366c3ad76a8c51.__obf_e57e229f9a2faf9a.
		Encode(__obf_7a37ccab267895e6, __obf_00fc491caa967a74)
	for __obf_2deec7c38ffb6ae3 := 1; __obf_2deec7c38ffb6ae3 < __obf_29366c3ad76a8c51.__obf_1060c952158f31ac.Len(); __obf_2deec7c38ffb6ae3++ {
		__obf_00fc491caa967a74.
			WriteMore()
		__obf_7a37ccab267895e6 = __obf_29366c3ad76a8c51.__obf_1060c952158f31ac.UnsafeGetIndex(__obf_d3c919547bf7e47a, __obf_2deec7c38ffb6ae3)
		__obf_29366c3ad76a8c51.__obf_e57e229f9a2faf9a.
			Encode(__obf_7a37ccab267895e6, __obf_00fc491caa967a74)
	}
	__obf_00fc491caa967a74.
		WriteArrayEnd()
	if __obf_00fc491caa967a74.Error != nil && __obf_00fc491caa967a74.Error != io.EOF {
		__obf_00fc491caa967a74.
			Error = fmt.Errorf("%v: %s", __obf_29366c3ad76a8c51.__obf_1060c952158f31ac, __obf_00fc491caa967a74.Error.Error())
	}
}

func (__obf_29366c3ad76a8c51 *__obf_55abf1f4a08610a6) IsEmpty(__obf_d3c919547bf7e47a unsafe.Pointer) bool {
	return false
}

type __obf_eab64d10d33e7ed7 struct {
	__obf_1060c952158f31ac *reflect2.UnsafeArrayType
	__obf_45aaf5212411947b ValDecoder
}

func (__obf_18f746d7b5b440ee *__obf_eab64d10d33e7ed7) Decode(__obf_d3c919547bf7e47a unsafe.Pointer, __obf_67008a6a9e5ba828 *Iterator) {
	__obf_18f746d7b5b440ee.__obf_95ea2b5b23c8db5c(__obf_d3c919547bf7e47a, __obf_67008a6a9e5ba828)
	if __obf_67008a6a9e5ba828.Error != nil && __obf_67008a6a9e5ba828.Error != io.EOF {
		__obf_67008a6a9e5ba828.
			Error = fmt.Errorf("%v: %s", __obf_18f746d7b5b440ee.__obf_1060c952158f31ac, __obf_67008a6a9e5ba828.Error.Error())
	}
}

func (__obf_18f746d7b5b440ee *__obf_eab64d10d33e7ed7) __obf_95ea2b5b23c8db5c(__obf_d3c919547bf7e47a unsafe.Pointer, __obf_67008a6a9e5ba828 *Iterator) {
	__obf_dab9baaadfa7c8c2 := __obf_67008a6a9e5ba828.__obf_b781a59d5a0d2490()
	__obf_1060c952158f31ac := __obf_18f746d7b5b440ee.__obf_1060c952158f31ac
	if __obf_dab9baaadfa7c8c2 == 'n' {
		__obf_67008a6a9e5ba828.__obf_4aeb767e0be7277a('u', 'l', 'l')
		return
	}
	if __obf_dab9baaadfa7c8c2 != '[' {
		__obf_67008a6a9e5ba828.
			ReportError("decode array", "expect [ or n, but found "+string([]byte{__obf_dab9baaadfa7c8c2}))
		return
	}
	__obf_dab9baaadfa7c8c2 = __obf_67008a6a9e5ba828.__obf_b781a59d5a0d2490()
	if __obf_dab9baaadfa7c8c2 == ']' {
		return
	}
	__obf_67008a6a9e5ba828.__obf_3284a1eaa2a0abb6()
	__obf_7a37ccab267895e6 := __obf_1060c952158f31ac.UnsafeGetIndex(__obf_d3c919547bf7e47a, 0)
	__obf_18f746d7b5b440ee.__obf_45aaf5212411947b.
		Decode(__obf_7a37ccab267895e6, __obf_67008a6a9e5ba828)
	__obf_b1248a9e4497b06b := 1
	for __obf_dab9baaadfa7c8c2 = __obf_67008a6a9e5ba828.__obf_b781a59d5a0d2490(); __obf_dab9baaadfa7c8c2 == ','; __obf_dab9baaadfa7c8c2 = __obf_67008a6a9e5ba828.__obf_b781a59d5a0d2490() {
		if __obf_b1248a9e4497b06b >= __obf_1060c952158f31ac.Len() {
			__obf_67008a6a9e5ba828.
				Skip()
			continue
		}
		__obf_e9657e567969f634 := __obf_b1248a9e4497b06b
		__obf_b1248a9e4497b06b += 1
		__obf_7a37ccab267895e6 = __obf_1060c952158f31ac.UnsafeGetIndex(__obf_d3c919547bf7e47a, __obf_e9657e567969f634)
		__obf_18f746d7b5b440ee.__obf_45aaf5212411947b.
			Decode(__obf_7a37ccab267895e6, __obf_67008a6a9e5ba828)
	}
	if __obf_dab9baaadfa7c8c2 != ']' {
		__obf_67008a6a9e5ba828.
			ReportError("decode array", "expect ], but found "+string([]byte{__obf_dab9baaadfa7c8c2}))
		return
	}
}
