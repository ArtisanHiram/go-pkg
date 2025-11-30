package __obf_030d39f7a8de96c6

import (
	"fmt"
	"github.com/modern-go/reflect2"
	"io"
	"unsafe"
)

func __obf_6f35214e10cf09dc(__obf_a565fbaffca944f0 *__obf_a565fbaffca944f0, __obf_8744d0b8c80ccdc1 reflect2.Type) ValDecoder {
	__obf_c44e8ae6d57faefd := __obf_8744d0b8c80ccdc1.(*reflect2.UnsafeSliceType)
	__obf_11a3f28116164d09 := __obf_729ef08c8b8e060e(__obf_a565fbaffca944f0.append("[sliceElem]"), __obf_c44e8ae6d57faefd.Elem())
	return &__obf_b122f1d695f8466c{__obf_c44e8ae6d57faefd, __obf_11a3f28116164d09}
}

func __obf_f70c82e772cafe09(__obf_a565fbaffca944f0 *__obf_a565fbaffca944f0, __obf_8744d0b8c80ccdc1 reflect2.Type) ValEncoder {
	__obf_c44e8ae6d57faefd := __obf_8744d0b8c80ccdc1.(*reflect2.UnsafeSliceType)
	__obf_008f61596d7e9523 := __obf_37ef471fa5e40404(__obf_a565fbaffca944f0.append("[sliceElem]"), __obf_c44e8ae6d57faefd.Elem())
	return &__obf_a69908ce40c0d051{__obf_c44e8ae6d57faefd, __obf_008f61596d7e9523}
}

type __obf_a69908ce40c0d051 struct {
	__obf_c44e8ae6d57faefd *reflect2.UnsafeSliceType
	__obf_73c590fa207e5962 ValEncoder
}

func (__obf_008f61596d7e9523 *__obf_a69908ce40c0d051) Encode(__obf_dbbf371b91136494 unsafe.Pointer, __obf_8a2c51fe22775655 *Stream) {
	if __obf_008f61596d7e9523.__obf_c44e8ae6d57faefd.UnsafeIsNil(__obf_dbbf371b91136494) {
		__obf_8a2c51fe22775655.
			WriteNil()
		return
	}
	__obf_9b9647f9a13cc8c1 := __obf_008f61596d7e9523.__obf_c44e8ae6d57faefd.UnsafeLengthOf(__obf_dbbf371b91136494)
	if __obf_9b9647f9a13cc8c1 == 0 {
		__obf_8a2c51fe22775655.
			WriteEmptyArray()
		return
	}
	__obf_8a2c51fe22775655.
		WriteArrayStart()
	__obf_008f61596d7e9523.__obf_73c590fa207e5962.
		Encode(__obf_008f61596d7e9523.__obf_c44e8ae6d57faefd.UnsafeGetIndex(__obf_dbbf371b91136494, 0), __obf_8a2c51fe22775655)
	for __obf_82c6e05b9d226c58 := 1; __obf_82c6e05b9d226c58 < __obf_9b9647f9a13cc8c1; __obf_82c6e05b9d226c58++ {
		__obf_8a2c51fe22775655.
			WriteMore()
		__obf_569a31123aaa281e := __obf_008f61596d7e9523.__obf_c44e8ae6d57faefd.UnsafeGetIndex(__obf_dbbf371b91136494, __obf_82c6e05b9d226c58)
		__obf_008f61596d7e9523.__obf_73c590fa207e5962.
			Encode(__obf_569a31123aaa281e, __obf_8a2c51fe22775655)
	}
	__obf_8a2c51fe22775655.
		WriteArrayEnd()
	if __obf_8a2c51fe22775655.Error != nil && __obf_8a2c51fe22775655.Error != io.EOF {
		__obf_8a2c51fe22775655.
			Error = fmt.Errorf("%v: %s", __obf_008f61596d7e9523.__obf_c44e8ae6d57faefd, __obf_8a2c51fe22775655.Error.Error())
	}
}

func (__obf_008f61596d7e9523 *__obf_a69908ce40c0d051) IsEmpty(__obf_dbbf371b91136494 unsafe.Pointer) bool {
	return __obf_008f61596d7e9523.__obf_c44e8ae6d57faefd.UnsafeLengthOf(__obf_dbbf371b91136494) == 0
}

type __obf_b122f1d695f8466c struct {
	__obf_c44e8ae6d57faefd *reflect2.UnsafeSliceType
	__obf_be23981cae4f81b9 ValDecoder
}

func (__obf_11a3f28116164d09 *__obf_b122f1d695f8466c) Decode(__obf_dbbf371b91136494 unsafe.Pointer, __obf_4ab56a99965952e7 *Iterator) {
	__obf_11a3f28116164d09.__obf_b2eef4fad6c777b1(__obf_dbbf371b91136494, __obf_4ab56a99965952e7)
	if __obf_4ab56a99965952e7.Error != nil && __obf_4ab56a99965952e7.Error != io.EOF {
		__obf_4ab56a99965952e7.
			Error = fmt.Errorf("%v: %s", __obf_11a3f28116164d09.__obf_c44e8ae6d57faefd, __obf_4ab56a99965952e7.Error.Error())
	}
}

func (__obf_11a3f28116164d09 *__obf_b122f1d695f8466c) __obf_b2eef4fad6c777b1(__obf_dbbf371b91136494 unsafe.Pointer, __obf_4ab56a99965952e7 *Iterator) {
	__obf_24309b3b0ff9ba22 := __obf_4ab56a99965952e7.__obf_61df301d1f67ad73()
	__obf_c44e8ae6d57faefd := __obf_11a3f28116164d09.__obf_c44e8ae6d57faefd
	if __obf_24309b3b0ff9ba22 == 'n' {
		__obf_4ab56a99965952e7.__obf_01038f7a5ba9ca56('u', 'l', 'l')
		__obf_c44e8ae6d57faefd.
			UnsafeSetNil(__obf_dbbf371b91136494)
		return
	}
	if __obf_24309b3b0ff9ba22 != '[' {
		__obf_4ab56a99965952e7.
			ReportError("decode slice", "expect [ or n, but found "+string([]byte{__obf_24309b3b0ff9ba22}))
		return
	}
	__obf_24309b3b0ff9ba22 = __obf_4ab56a99965952e7.__obf_61df301d1f67ad73()
	if __obf_24309b3b0ff9ba22 == ']' {
		__obf_c44e8ae6d57faefd.
			UnsafeSet(__obf_dbbf371b91136494, __obf_c44e8ae6d57faefd.UnsafeMakeSlice(0, 0))
		return
	}
	__obf_4ab56a99965952e7.__obf_ba6d0be9a7ab6ae4()
	__obf_c44e8ae6d57faefd.
		UnsafeGrow(__obf_dbbf371b91136494, 1)
	__obf_569a31123aaa281e := __obf_c44e8ae6d57faefd.UnsafeGetIndex(__obf_dbbf371b91136494, 0)
	__obf_11a3f28116164d09.__obf_be23981cae4f81b9.
		Decode(__obf_569a31123aaa281e, __obf_4ab56a99965952e7)
	__obf_9b9647f9a13cc8c1 := 1
	for __obf_24309b3b0ff9ba22 = __obf_4ab56a99965952e7.__obf_61df301d1f67ad73(); __obf_24309b3b0ff9ba22 == ','; __obf_24309b3b0ff9ba22 = __obf_4ab56a99965952e7.__obf_61df301d1f67ad73() {
		__obf_17119ea1bee712b3 := __obf_9b9647f9a13cc8c1
		__obf_9b9647f9a13cc8c1 += 1
		__obf_c44e8ae6d57faefd.
			UnsafeGrow(__obf_dbbf371b91136494, __obf_9b9647f9a13cc8c1)
		__obf_569a31123aaa281e = __obf_c44e8ae6d57faefd.UnsafeGetIndex(__obf_dbbf371b91136494, __obf_17119ea1bee712b3)
		__obf_11a3f28116164d09.__obf_be23981cae4f81b9.
			Decode(__obf_569a31123aaa281e, __obf_4ab56a99965952e7)
	}
	if __obf_24309b3b0ff9ba22 != ']' {
		__obf_4ab56a99965952e7.
			ReportError("decode slice", "expect ], but found "+string([]byte{__obf_24309b3b0ff9ba22}))
		return
	}
}
