package __obf_030d39f7a8de96c6

import (
	"fmt"
	"github.com/modern-go/reflect2"
	"io"
	"unsafe"
)

func __obf_9dc19cae1724dfa5(__obf_a565fbaffca944f0 *__obf_a565fbaffca944f0, __obf_8744d0b8c80ccdc1 reflect2.Type) ValDecoder {
	__obf_ab7510837e40b9df := __obf_8744d0b8c80ccdc1.(*reflect2.UnsafeArrayType)
	__obf_11a3f28116164d09 := __obf_729ef08c8b8e060e(__obf_a565fbaffca944f0.append("[arrayElem]"), __obf_ab7510837e40b9df.Elem())
	return &__obf_2687f9efadbadc9a{__obf_ab7510837e40b9df, __obf_11a3f28116164d09}
}

func __obf_dd1bf627e51d80bf(__obf_a565fbaffca944f0 *__obf_a565fbaffca944f0, __obf_8744d0b8c80ccdc1 reflect2.Type) ValEncoder {
	__obf_ab7510837e40b9df := __obf_8744d0b8c80ccdc1.(*reflect2.UnsafeArrayType)
	if __obf_ab7510837e40b9df.Len() == 0 {
		return __obf_9a6e970409e0bc3c{}
	}
	__obf_008f61596d7e9523 := __obf_37ef471fa5e40404(__obf_a565fbaffca944f0.append("[arrayElem]"), __obf_ab7510837e40b9df.Elem())
	return &__obf_31e9b5ebcd4b6405{__obf_ab7510837e40b9df, __obf_008f61596d7e9523}
}

type __obf_9a6e970409e0bc3c struct{}

func (__obf_008f61596d7e9523 __obf_9a6e970409e0bc3c) Encode(__obf_dbbf371b91136494 unsafe.Pointer, __obf_8a2c51fe22775655 *Stream) {
	__obf_8a2c51fe22775655.
		WriteEmptyArray()
}

func (__obf_008f61596d7e9523 __obf_9a6e970409e0bc3c) IsEmpty(__obf_dbbf371b91136494 unsafe.Pointer) bool {
	return true
}

type __obf_31e9b5ebcd4b6405 struct {
	__obf_ab7510837e40b9df *reflect2.UnsafeArrayType
	__obf_73c590fa207e5962 ValEncoder
}

func (__obf_008f61596d7e9523 *__obf_31e9b5ebcd4b6405) Encode(__obf_dbbf371b91136494 unsafe.Pointer, __obf_8a2c51fe22775655 *Stream) {
	__obf_8a2c51fe22775655.
		WriteArrayStart()
	__obf_569a31123aaa281e := unsafe.Pointer(__obf_dbbf371b91136494)
	__obf_008f61596d7e9523.__obf_73c590fa207e5962.
		Encode(__obf_569a31123aaa281e, __obf_8a2c51fe22775655)
	for __obf_82c6e05b9d226c58 := 1; __obf_82c6e05b9d226c58 < __obf_008f61596d7e9523.__obf_ab7510837e40b9df.Len(); __obf_82c6e05b9d226c58++ {
		__obf_8a2c51fe22775655.
			WriteMore()
		__obf_569a31123aaa281e = __obf_008f61596d7e9523.__obf_ab7510837e40b9df.UnsafeGetIndex(__obf_dbbf371b91136494, __obf_82c6e05b9d226c58)
		__obf_008f61596d7e9523.__obf_73c590fa207e5962.
			Encode(__obf_569a31123aaa281e, __obf_8a2c51fe22775655)
	}
	__obf_8a2c51fe22775655.
		WriteArrayEnd()
	if __obf_8a2c51fe22775655.Error != nil && __obf_8a2c51fe22775655.Error != io.EOF {
		__obf_8a2c51fe22775655.
			Error = fmt.Errorf("%v: %s", __obf_008f61596d7e9523.__obf_ab7510837e40b9df, __obf_8a2c51fe22775655.Error.Error())
	}
}

func (__obf_008f61596d7e9523 *__obf_31e9b5ebcd4b6405) IsEmpty(__obf_dbbf371b91136494 unsafe.Pointer) bool {
	return false
}

type __obf_2687f9efadbadc9a struct {
	__obf_ab7510837e40b9df *reflect2.UnsafeArrayType
	__obf_be23981cae4f81b9 ValDecoder
}

func (__obf_11a3f28116164d09 *__obf_2687f9efadbadc9a) Decode(__obf_dbbf371b91136494 unsafe.Pointer, __obf_4ab56a99965952e7 *Iterator) {
	__obf_11a3f28116164d09.__obf_b2eef4fad6c777b1(__obf_dbbf371b91136494, __obf_4ab56a99965952e7)
	if __obf_4ab56a99965952e7.Error != nil && __obf_4ab56a99965952e7.Error != io.EOF {
		__obf_4ab56a99965952e7.
			Error = fmt.Errorf("%v: %s", __obf_11a3f28116164d09.__obf_ab7510837e40b9df, __obf_4ab56a99965952e7.Error.Error())
	}
}

func (__obf_11a3f28116164d09 *__obf_2687f9efadbadc9a) __obf_b2eef4fad6c777b1(__obf_dbbf371b91136494 unsafe.Pointer, __obf_4ab56a99965952e7 *Iterator) {
	__obf_24309b3b0ff9ba22 := __obf_4ab56a99965952e7.__obf_61df301d1f67ad73()
	__obf_ab7510837e40b9df := __obf_11a3f28116164d09.__obf_ab7510837e40b9df
	if __obf_24309b3b0ff9ba22 == 'n' {
		__obf_4ab56a99965952e7.__obf_01038f7a5ba9ca56('u', 'l', 'l')
		return
	}
	if __obf_24309b3b0ff9ba22 != '[' {
		__obf_4ab56a99965952e7.
			ReportError("decode array", "expect [ or n, but found "+string([]byte{__obf_24309b3b0ff9ba22}))
		return
	}
	__obf_24309b3b0ff9ba22 = __obf_4ab56a99965952e7.__obf_61df301d1f67ad73()
	if __obf_24309b3b0ff9ba22 == ']' {
		return
	}
	__obf_4ab56a99965952e7.__obf_ba6d0be9a7ab6ae4()
	__obf_569a31123aaa281e := __obf_ab7510837e40b9df.UnsafeGetIndex(__obf_dbbf371b91136494, 0)
	__obf_11a3f28116164d09.__obf_be23981cae4f81b9.
		Decode(__obf_569a31123aaa281e, __obf_4ab56a99965952e7)
	__obf_9b9647f9a13cc8c1 := 1
	for __obf_24309b3b0ff9ba22 = __obf_4ab56a99965952e7.__obf_61df301d1f67ad73(); __obf_24309b3b0ff9ba22 == ','; __obf_24309b3b0ff9ba22 = __obf_4ab56a99965952e7.__obf_61df301d1f67ad73() {
		if __obf_9b9647f9a13cc8c1 >= __obf_ab7510837e40b9df.Len() {
			__obf_4ab56a99965952e7.
				Skip()
			continue
		}
		__obf_17119ea1bee712b3 := __obf_9b9647f9a13cc8c1
		__obf_9b9647f9a13cc8c1 += 1
		__obf_569a31123aaa281e = __obf_ab7510837e40b9df.UnsafeGetIndex(__obf_dbbf371b91136494, __obf_17119ea1bee712b3)
		__obf_11a3f28116164d09.__obf_be23981cae4f81b9.
			Decode(__obf_569a31123aaa281e, __obf_4ab56a99965952e7)
	}
	if __obf_24309b3b0ff9ba22 != ']' {
		__obf_4ab56a99965952e7.
			ReportError("decode array", "expect ], but found "+string([]byte{__obf_24309b3b0ff9ba22}))
		return
	}
}
