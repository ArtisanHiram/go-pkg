package __obf_030d39f7a8de96c6

import (
	"github.com/modern-go/reflect2"
	"unsafe"
)

func __obf_2ca1ae44e249759c(__obf_a565fbaffca944f0 *__obf_a565fbaffca944f0, __obf_8744d0b8c80ccdc1 reflect2.Type) ValDecoder {
	__obf_3a51157620f9910b := __obf_8744d0b8c80ccdc1.(*reflect2.UnsafePtrType)
	__obf_d4fd7b4aa577e66f := __obf_3a51157620f9910b.Elem()
	__obf_11a3f28116164d09 := __obf_729ef08c8b8e060e(__obf_a565fbaffca944f0, __obf_d4fd7b4aa577e66f)
	return &OptionalDecoder{__obf_d4fd7b4aa577e66f, __obf_11a3f28116164d09}
}

func __obf_bb9550a27dc28be6(__obf_a565fbaffca944f0 *__obf_a565fbaffca944f0, __obf_8744d0b8c80ccdc1 reflect2.Type) ValEncoder {
	__obf_3a51157620f9910b := __obf_8744d0b8c80ccdc1.(*reflect2.UnsafePtrType)
	__obf_d4fd7b4aa577e66f := __obf_3a51157620f9910b.Elem()
	__obf_73c590fa207e5962 := __obf_37ef471fa5e40404(__obf_a565fbaffca944f0, __obf_d4fd7b4aa577e66f)
	__obf_008f61596d7e9523 := &OptionalEncoder{__obf_73c590fa207e5962}
	return __obf_008f61596d7e9523
}

type OptionalDecoder struct {
	ValueType    reflect2.Type
	ValueDecoder ValDecoder
}

func (__obf_11a3f28116164d09 *OptionalDecoder) Decode(__obf_dbbf371b91136494 unsafe.Pointer, __obf_4ab56a99965952e7 *Iterator) {
	if __obf_4ab56a99965952e7.ReadNil() {
		*((*unsafe.Pointer)(__obf_dbbf371b91136494)) = nil
	} else {
		if *((*unsafe.Pointer)(__obf_dbbf371b91136494)) == nil {
			__obf_1218ea139ffa7953 := //pointer to null, we have to allocate memory to hold the value
				__obf_11a3f28116164d09.ValueType.UnsafeNew()
			__obf_11a3f28116164d09.
				ValueDecoder.Decode(__obf_1218ea139ffa7953, __obf_4ab56a99965952e7)
			*((*unsafe.Pointer)(__obf_dbbf371b91136494)) = __obf_1218ea139ffa7953
		} else {
			__obf_11a3f28116164d09.
				//reuse existing instance
				ValueDecoder.Decode(*((*unsafe.Pointer)(__obf_dbbf371b91136494)), __obf_4ab56a99965952e7)
		}
	}
}

type __obf_174b577e33a12aa4 struct {
	__obf_8afc98a56c05c14e reflect2. // only to deference a pointer
				Type
	__obf_d11a33e2f1240793 ValDecoder
}

func (__obf_11a3f28116164d09 *__obf_174b577e33a12aa4) Decode(__obf_dbbf371b91136494 unsafe.Pointer, __obf_4ab56a99965952e7 *Iterator) {
	if *((*unsafe.Pointer)(__obf_dbbf371b91136494)) == nil {
		__obf_1218ea139ffa7953 := //pointer to null, we have to allocate memory to hold the value
			__obf_11a3f28116164d09.__obf_8afc98a56c05c14e.UnsafeNew()
		__obf_11a3f28116164d09.__obf_d11a33e2f1240793.
			Decode(__obf_1218ea139ffa7953, __obf_4ab56a99965952e7)
		*((*unsafe.Pointer)(__obf_dbbf371b91136494)) = __obf_1218ea139ffa7953
	} else {
		__obf_11a3f28116164d09.
			//reuse existing instance
			__obf_d11a33e2f1240793.
			Decode(*((*unsafe.Pointer)(__obf_dbbf371b91136494)), __obf_4ab56a99965952e7)
	}
}

type OptionalEncoder struct {
	ValueEncoder ValEncoder
}

func (__obf_008f61596d7e9523 *OptionalEncoder) Encode(__obf_dbbf371b91136494 unsafe.Pointer, __obf_8a2c51fe22775655 *Stream) {
	if *((*unsafe.Pointer)(__obf_dbbf371b91136494)) == nil {
		__obf_8a2c51fe22775655.
			WriteNil()
	} else {
		__obf_008f61596d7e9523.
			ValueEncoder.Encode(*((*unsafe.Pointer)(__obf_dbbf371b91136494)), __obf_8a2c51fe22775655)
	}
}

func (__obf_008f61596d7e9523 *OptionalEncoder) IsEmpty(__obf_dbbf371b91136494 unsafe.Pointer) bool {
	return *((*unsafe.Pointer)(__obf_dbbf371b91136494)) == nil
}

type __obf_e3ea39e92054d646 struct {
	ValueEncoder ValEncoder
}

func (__obf_008f61596d7e9523 *__obf_e3ea39e92054d646) Encode(__obf_dbbf371b91136494 unsafe.Pointer, __obf_8a2c51fe22775655 *Stream) {
	if *((*unsafe.Pointer)(__obf_dbbf371b91136494)) == nil {
		__obf_8a2c51fe22775655.
			WriteNil()
	} else {
		__obf_008f61596d7e9523.
			ValueEncoder.Encode(*((*unsafe.Pointer)(__obf_dbbf371b91136494)), __obf_8a2c51fe22775655)
	}
}

func (__obf_008f61596d7e9523 *__obf_e3ea39e92054d646) IsEmpty(__obf_dbbf371b91136494 unsafe.Pointer) bool {
	__obf_43056da345a9146e := *((*unsafe.Pointer)(__obf_dbbf371b91136494))
	if __obf_43056da345a9146e == nil {
		return true
	}
	return __obf_008f61596d7e9523.ValueEncoder.IsEmpty(__obf_43056da345a9146e)
}

func (__obf_008f61596d7e9523 *__obf_e3ea39e92054d646) IsEmbeddedPtrNil(__obf_dbbf371b91136494 unsafe.Pointer) bool {
	__obf_85e89b6c0ad7dab9 := *((*unsafe.Pointer)(__obf_dbbf371b91136494))
	if __obf_85e89b6c0ad7dab9 == nil {
		return true
	}
	__obf_f22b71bcff862810, __obf_ae94b7511a43a3e7 := __obf_008f61596d7e9523.ValueEncoder.(IsEmbeddedPtrNil)
	if !__obf_ae94b7511a43a3e7 {
		return false
	}
	__obf_ba13c65f6e3a796c := unsafe.Pointer(__obf_85e89b6c0ad7dab9)
	return __obf_f22b71bcff862810.IsEmbeddedPtrNil(__obf_ba13c65f6e3a796c)
}

type __obf_d086959fc9847188 struct {
	__obf_008f61596d7e9523 ValEncoder
}

func (__obf_008f61596d7e9523 *__obf_d086959fc9847188) Encode(__obf_dbbf371b91136494 unsafe.Pointer, __obf_8a2c51fe22775655 *Stream) {
	__obf_008f61596d7e9523.__obf_008f61596d7e9523.
		Encode(unsafe.Pointer(&__obf_dbbf371b91136494), __obf_8a2c51fe22775655)
}

func (__obf_008f61596d7e9523 *__obf_d086959fc9847188) IsEmpty(__obf_dbbf371b91136494 unsafe.Pointer) bool {
	return __obf_008f61596d7e9523.__obf_008f61596d7e9523.IsEmpty(unsafe.Pointer(&__obf_dbbf371b91136494))
}

type __obf_2e9c42d03f0d8cd5 struct {
	__obf_11a3f28116164d09 ValDecoder
}

func (__obf_11a3f28116164d09 *__obf_2e9c42d03f0d8cd5) Decode(__obf_dbbf371b91136494 unsafe.Pointer, __obf_4ab56a99965952e7 *Iterator) {
	__obf_11a3f28116164d09.__obf_11a3f28116164d09.
		Decode(unsafe.Pointer(&__obf_dbbf371b91136494), __obf_4ab56a99965952e7)
}
