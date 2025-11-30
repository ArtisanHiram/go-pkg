package __obf_030d39f7a8de96c6

import (
	"fmt"
	"io"
	"reflect"
	"sort"
	"unsafe"

	"github.com/modern-go/reflect2"
)

func __obf_d63eb7ded243ae52(__obf_a565fbaffca944f0 *__obf_a565fbaffca944f0, __obf_8744d0b8c80ccdc1 reflect2.Type) ValDecoder {
	__obf_2153a234428f3398 := __obf_8744d0b8c80ccdc1.(*reflect2.UnsafeMapType)
	__obf_1bd3e17936aaaae3 := __obf_61610649fc8afc4b(__obf_a565fbaffca944f0.append("[mapKey]"), __obf_2153a234428f3398.Key())
	__obf_be23981cae4f81b9 := __obf_729ef08c8b8e060e(__obf_a565fbaffca944f0.append("[mapElem]"), __obf_2153a234428f3398.Elem())
	return &__obf_7b221609a941b0a1{__obf_2153a234428f3398: __obf_2153a234428f3398, __obf_f80ccf0a95ddf294: __obf_2153a234428f3398.Key(), __obf_d4fd7b4aa577e66f: __obf_2153a234428f3398.Elem(), __obf_1bd3e17936aaaae3: __obf_1bd3e17936aaaae3, __obf_be23981cae4f81b9: __obf_be23981cae4f81b9}
}

func __obf_5cc5784cdf0ac33a(__obf_a565fbaffca944f0 *__obf_a565fbaffca944f0, __obf_8744d0b8c80ccdc1 reflect2.Type) ValEncoder {
	__obf_2153a234428f3398 := __obf_8744d0b8c80ccdc1.(*reflect2.UnsafeMapType)
	if __obf_a565fbaffca944f0.__obf_e01b84a9a1b0f266 {
		return &__obf_dfe61d8c4987b512{__obf_2153a234428f3398: __obf_2153a234428f3398, __obf_bf24e60df05dbd7a: __obf_0c4f0e0fbb717718(__obf_a565fbaffca944f0.append("[mapKey]"), __obf_2153a234428f3398.Key()), __obf_73c590fa207e5962: __obf_37ef471fa5e40404(__obf_a565fbaffca944f0.append("[mapElem]"), __obf_2153a234428f3398.Elem())}
	}
	return &__obf_a18951b953f76dd9{__obf_2153a234428f3398: __obf_2153a234428f3398, __obf_bf24e60df05dbd7a: __obf_0c4f0e0fbb717718(__obf_a565fbaffca944f0.append("[mapKey]"), __obf_2153a234428f3398.Key()), __obf_73c590fa207e5962: __obf_37ef471fa5e40404(__obf_a565fbaffca944f0.append("[mapElem]"), __obf_2153a234428f3398.Elem())}
}

func __obf_61610649fc8afc4b(__obf_a565fbaffca944f0 *__obf_a565fbaffca944f0, __obf_8744d0b8c80ccdc1 reflect2.Type) ValDecoder {
	__obf_11a3f28116164d09 := __obf_a565fbaffca944f0.__obf_6313bfb9c913bd50.CreateMapKeyDecoder(__obf_8744d0b8c80ccdc1)
	if __obf_11a3f28116164d09 != nil {
		return __obf_11a3f28116164d09
	}
	for _, __obf_621544a57e52000f := range __obf_a565fbaffca944f0.__obf_6621255bc1f80c8e {
		__obf_11a3f28116164d09 := __obf_621544a57e52000f.CreateMapKeyDecoder(__obf_8744d0b8c80ccdc1)
		if __obf_11a3f28116164d09 != nil {
			return __obf_11a3f28116164d09
		}
	}
	__obf_3a51157620f9910b := reflect2.PtrTo(__obf_8744d0b8c80ccdc1)
	if __obf_3a51157620f9910b.Implements(__obf_b7ddfd40cfe62a98) {
		return &__obf_2e9c42d03f0d8cd5{
			&__obf_cd063ac27c40d5c6{__obf_a7610e23a63622fd: __obf_3a51157620f9910b},
		}
	}
	if __obf_8744d0b8c80ccdc1.Implements(__obf_b7ddfd40cfe62a98) {
		return &__obf_cd063ac27c40d5c6{__obf_a7610e23a63622fd: __obf_8744d0b8c80ccdc1}
	}
	if __obf_3a51157620f9910b.Implements(__obf_f58653b68e3e7233) {
		return &__obf_2e9c42d03f0d8cd5{
			&__obf_e0e26cbc14bdbad8{__obf_a7610e23a63622fd: __obf_3a51157620f9910b},
		}
	}
	if __obf_8744d0b8c80ccdc1.Implements(__obf_f58653b68e3e7233) {
		return &__obf_e0e26cbc14bdbad8{__obf_a7610e23a63622fd: __obf_8744d0b8c80ccdc1}
	}

	switch __obf_8744d0b8c80ccdc1.Kind() {
	case reflect.String:
		return __obf_729ef08c8b8e060e(__obf_a565fbaffca944f0, reflect2.DefaultTypeOfKind(reflect.String))
	case reflect.Bool,
		reflect.Uint8, reflect.Int8,
		reflect.Uint16, reflect.Int16,
		reflect.Uint32, reflect.Int32,
		reflect.Uint64, reflect.Int64,
		reflect.Uint, reflect.Int,
		reflect.Float32, reflect.Float64,
		reflect.Uintptr:
		__obf_8744d0b8c80ccdc1 = reflect2.DefaultTypeOfKind(__obf_8744d0b8c80ccdc1.Kind())
		return &__obf_27c2e8b9fed7f4ba{__obf_729ef08c8b8e060e(__obf_a565fbaffca944f0, __obf_8744d0b8c80ccdc1)}
	default:
		return &__obf_195d4c78d647b89f{__obf_fcc907dd69879566: fmt.Errorf("unsupported map key type: %v", __obf_8744d0b8c80ccdc1)}
	}
}

func __obf_0c4f0e0fbb717718(__obf_a565fbaffca944f0 *__obf_a565fbaffca944f0, __obf_8744d0b8c80ccdc1 reflect2.Type) ValEncoder {
	__obf_008f61596d7e9523 := __obf_a565fbaffca944f0.__obf_7db1f0a55b319e45.CreateMapKeyEncoder(__obf_8744d0b8c80ccdc1)
	if __obf_008f61596d7e9523 != nil {
		return __obf_008f61596d7e9523
	}
	for _, __obf_621544a57e52000f := range __obf_a565fbaffca944f0.__obf_6621255bc1f80c8e {
		__obf_008f61596d7e9523 := __obf_621544a57e52000f.CreateMapKeyEncoder(__obf_8744d0b8c80ccdc1)
		if __obf_008f61596d7e9523 != nil {
			return __obf_008f61596d7e9523
		}
	}

	if __obf_8744d0b8c80ccdc1.Kind() != reflect.String {
		if __obf_8744d0b8c80ccdc1 == __obf_5353ad9fd2397e25 {
			return &__obf_be22d3d9cea5924c{__obf_6e4d12bda4dd7e33: __obf_a565fbaffca944f0.EncoderOf(reflect2.TypeOf(""))}
		}
		if __obf_8744d0b8c80ccdc1.Implements(__obf_5353ad9fd2397e25) {
			return &__obf_1632fdd756ee8530{__obf_a7610e23a63622fd: __obf_8744d0b8c80ccdc1, __obf_6e4d12bda4dd7e33: __obf_a565fbaffca944f0.EncoderOf(reflect2.TypeOf(""))}
		}
	}

	switch __obf_8744d0b8c80ccdc1.Kind() {
	case reflect.String:
		return __obf_37ef471fa5e40404(__obf_a565fbaffca944f0, reflect2.DefaultTypeOfKind(reflect.String))
	case reflect.Bool,
		reflect.Uint8, reflect.Int8,
		reflect.Uint16, reflect.Int16,
		reflect.Uint32, reflect.Int32,
		reflect.Uint64, reflect.Int64,
		reflect.Uint, reflect.Int,
		reflect.Float32, reflect.Float64,
		reflect.Uintptr:
		__obf_8744d0b8c80ccdc1 = reflect2.DefaultTypeOfKind(__obf_8744d0b8c80ccdc1.Kind())
		return &__obf_250d39cf19eeed62{__obf_37ef471fa5e40404(__obf_a565fbaffca944f0, __obf_8744d0b8c80ccdc1)}
	default:
		if __obf_8744d0b8c80ccdc1.Kind() == reflect.Interface {
			return &__obf_611b016ea7f35155{__obf_a565fbaffca944f0, __obf_8744d0b8c80ccdc1}
		}
		return &__obf_f14cb8ddffae87b3{__obf_fcc907dd69879566: fmt.Errorf("unsupported map key type: %v", __obf_8744d0b8c80ccdc1)}
	}
}

type __obf_7b221609a941b0a1 struct {
	__obf_2153a234428f3398 *reflect2.UnsafeMapType
	__obf_f80ccf0a95ddf294 reflect2.Type
	__obf_d4fd7b4aa577e66f reflect2.Type
	__obf_1bd3e17936aaaae3 ValDecoder
	__obf_be23981cae4f81b9 ValDecoder
}

func (__obf_11a3f28116164d09 *__obf_7b221609a941b0a1) Decode(__obf_dbbf371b91136494 unsafe.Pointer, __obf_4ab56a99965952e7 *Iterator) {
	__obf_2153a234428f3398 := __obf_11a3f28116164d09.__obf_2153a234428f3398
	__obf_24309b3b0ff9ba22 := __obf_4ab56a99965952e7.__obf_61df301d1f67ad73()
	if __obf_24309b3b0ff9ba22 == 'n' {
		__obf_4ab56a99965952e7.__obf_01038f7a5ba9ca56('u', 'l', 'l')
		*(*unsafe.Pointer)(__obf_dbbf371b91136494) = nil
		__obf_2153a234428f3398.
			UnsafeSet(__obf_dbbf371b91136494, __obf_2153a234428f3398.UnsafeNew())
		return
	}
	if __obf_2153a234428f3398.UnsafeIsNil(__obf_dbbf371b91136494) {
		__obf_2153a234428f3398.
			UnsafeSet(__obf_dbbf371b91136494, __obf_2153a234428f3398.UnsafeMakeMap(0))
	}
	if __obf_24309b3b0ff9ba22 != '{' {
		__obf_4ab56a99965952e7.
			ReportError("ReadMapCB", `expect { or n, but found `+string([]byte{__obf_24309b3b0ff9ba22}))
		return
	}
	__obf_24309b3b0ff9ba22 = __obf_4ab56a99965952e7.__obf_61df301d1f67ad73()
	if __obf_24309b3b0ff9ba22 == '}' {
		return
	}
	__obf_4ab56a99965952e7.__obf_ba6d0be9a7ab6ae4()
	__obf_53d763816351c178 := __obf_11a3f28116164d09.__obf_f80ccf0a95ddf294.UnsafeNew()
	__obf_11a3f28116164d09.__obf_1bd3e17936aaaae3.
		Decode(__obf_53d763816351c178, __obf_4ab56a99965952e7)
	__obf_24309b3b0ff9ba22 = __obf_4ab56a99965952e7.__obf_61df301d1f67ad73()
	if __obf_24309b3b0ff9ba22 != ':' {
		__obf_4ab56a99965952e7.
			ReportError("ReadMapCB", "expect : after object field, but found "+string([]byte{__obf_24309b3b0ff9ba22}))
		return
	}
	__obf_e4d99e200ff06def := __obf_11a3f28116164d09.__obf_d4fd7b4aa577e66f.UnsafeNew()
	__obf_11a3f28116164d09.__obf_be23981cae4f81b9.
		Decode(__obf_e4d99e200ff06def, __obf_4ab56a99965952e7)
	__obf_11a3f28116164d09.__obf_2153a234428f3398.
		UnsafeSetIndex(__obf_dbbf371b91136494, __obf_53d763816351c178, __obf_e4d99e200ff06def)
	for __obf_24309b3b0ff9ba22 = __obf_4ab56a99965952e7.__obf_61df301d1f67ad73(); __obf_24309b3b0ff9ba22 == ','; __obf_24309b3b0ff9ba22 = __obf_4ab56a99965952e7.__obf_61df301d1f67ad73() {
		__obf_53d763816351c178 := __obf_11a3f28116164d09.__obf_f80ccf0a95ddf294.UnsafeNew()
		__obf_11a3f28116164d09.__obf_1bd3e17936aaaae3.
			Decode(__obf_53d763816351c178, __obf_4ab56a99965952e7)
		__obf_24309b3b0ff9ba22 = __obf_4ab56a99965952e7.__obf_61df301d1f67ad73()
		if __obf_24309b3b0ff9ba22 != ':' {
			__obf_4ab56a99965952e7.
				ReportError("ReadMapCB", "expect : after object field, but found "+string([]byte{__obf_24309b3b0ff9ba22}))
			return
		}
		__obf_e4d99e200ff06def := __obf_11a3f28116164d09.__obf_d4fd7b4aa577e66f.UnsafeNew()
		__obf_11a3f28116164d09.__obf_be23981cae4f81b9.
			Decode(__obf_e4d99e200ff06def, __obf_4ab56a99965952e7)
		__obf_11a3f28116164d09.__obf_2153a234428f3398.
			UnsafeSetIndex(__obf_dbbf371b91136494, __obf_53d763816351c178, __obf_e4d99e200ff06def)
	}
	if __obf_24309b3b0ff9ba22 != '}' {
		__obf_4ab56a99965952e7.
			ReportError("ReadMapCB", `expect }, but found `+string([]byte{__obf_24309b3b0ff9ba22}))
	}
}

type __obf_27c2e8b9fed7f4ba struct {
	__obf_11a3f28116164d09 ValDecoder
}

func (__obf_11a3f28116164d09 *__obf_27c2e8b9fed7f4ba) Decode(__obf_dbbf371b91136494 unsafe.Pointer, __obf_4ab56a99965952e7 *Iterator) {
	__obf_24309b3b0ff9ba22 := __obf_4ab56a99965952e7.__obf_61df301d1f67ad73()
	if __obf_24309b3b0ff9ba22 != '"' {
		__obf_4ab56a99965952e7.
			ReportError("ReadMapCB", `expect ", but found `+string([]byte{__obf_24309b3b0ff9ba22}))
		return
	}
	__obf_11a3f28116164d09.__obf_11a3f28116164d09.
		Decode(__obf_dbbf371b91136494, __obf_4ab56a99965952e7)
	__obf_24309b3b0ff9ba22 = __obf_4ab56a99965952e7.__obf_61df301d1f67ad73()
	if __obf_24309b3b0ff9ba22 != '"' {
		__obf_4ab56a99965952e7.
			ReportError("ReadMapCB", `expect ", but found `+string([]byte{__obf_24309b3b0ff9ba22}))
		return
	}
}

type __obf_250d39cf19eeed62 struct {
	__obf_008f61596d7e9523 ValEncoder
}

func (__obf_008f61596d7e9523 *__obf_250d39cf19eeed62) Encode(__obf_dbbf371b91136494 unsafe.Pointer, __obf_8a2c51fe22775655 *Stream) {
	__obf_8a2c51fe22775655.__obf_41130daa346c0e53('"')
	__obf_008f61596d7e9523.__obf_008f61596d7e9523.
		Encode(__obf_dbbf371b91136494, __obf_8a2c51fe22775655)
	__obf_8a2c51fe22775655.__obf_41130daa346c0e53('"')
}

func (__obf_008f61596d7e9523 *__obf_250d39cf19eeed62) IsEmpty(__obf_dbbf371b91136494 unsafe.Pointer) bool {
	return false
}

type __obf_611b016ea7f35155 struct {
	__obf_a565fbaffca944f0 *__obf_a565fbaffca944f0
	__obf_a7610e23a63622fd reflect2.Type
}

func (__obf_008f61596d7e9523 *__obf_611b016ea7f35155) Encode(__obf_dbbf371b91136494 unsafe.Pointer, __obf_8a2c51fe22775655 *Stream) {
	__obf_eefa92392cc2442c := __obf_008f61596d7e9523.__obf_a7610e23a63622fd.UnsafeIndirect(__obf_dbbf371b91136494)
	__obf_0c4f0e0fbb717718(__obf_008f61596d7e9523.__obf_a565fbaffca944f0, reflect2.TypeOf(__obf_eefa92392cc2442c)).Encode(reflect2.PtrOf(__obf_eefa92392cc2442c), __obf_8a2c51fe22775655)
}

func (__obf_008f61596d7e9523 *__obf_611b016ea7f35155) IsEmpty(__obf_dbbf371b91136494 unsafe.Pointer) bool {
	__obf_eefa92392cc2442c := __obf_008f61596d7e9523.__obf_a7610e23a63622fd.UnsafeIndirect(__obf_dbbf371b91136494)
	return __obf_0c4f0e0fbb717718(__obf_008f61596d7e9523.__obf_a565fbaffca944f0, reflect2.TypeOf(__obf_eefa92392cc2442c)).IsEmpty(reflect2.PtrOf(__obf_eefa92392cc2442c))
}

type __obf_a18951b953f76dd9 struct {
	__obf_2153a234428f3398 *reflect2.UnsafeMapType
	__obf_bf24e60df05dbd7a ValEncoder
	__obf_73c590fa207e5962 ValEncoder
}

func (__obf_008f61596d7e9523 *__obf_a18951b953f76dd9) Encode(__obf_dbbf371b91136494 unsafe.Pointer, __obf_8a2c51fe22775655 *Stream) {
	if *(*unsafe.Pointer)(__obf_dbbf371b91136494) == nil {
		__obf_8a2c51fe22775655.
			WriteNil()
		return
	}
	__obf_8a2c51fe22775655.
		WriteObjectStart()
	__obf_4ab56a99965952e7 := __obf_008f61596d7e9523.__obf_2153a234428f3398.UnsafeIterate(__obf_dbbf371b91136494)
	for __obf_82c6e05b9d226c58 := 0; __obf_4ab56a99965952e7.HasNext(); __obf_82c6e05b9d226c58++ {
		if __obf_82c6e05b9d226c58 != 0 {
			__obf_8a2c51fe22775655.
				WriteMore()
		}
		__obf_53d763816351c178, __obf_e4d99e200ff06def := __obf_4ab56a99965952e7.UnsafeNext()
		__obf_008f61596d7e9523.__obf_bf24e60df05dbd7a.
			Encode(__obf_53d763816351c178, __obf_8a2c51fe22775655)
		if __obf_8a2c51fe22775655.__obf_508619ef3f08fbb3 > 0 {
			__obf_8a2c51fe22775655.__obf_3635bad429be5857(byte(':'), byte(' '))
		} else {
			__obf_8a2c51fe22775655.__obf_41130daa346c0e53(':')
		}
		__obf_008f61596d7e9523.__obf_73c590fa207e5962.
			Encode(__obf_e4d99e200ff06def, __obf_8a2c51fe22775655)
	}
	__obf_8a2c51fe22775655.
		WriteObjectEnd()
}

func (__obf_008f61596d7e9523 *__obf_a18951b953f76dd9) IsEmpty(__obf_dbbf371b91136494 unsafe.Pointer) bool {
	__obf_4ab56a99965952e7 := __obf_008f61596d7e9523.__obf_2153a234428f3398.UnsafeIterate(__obf_dbbf371b91136494)
	return !__obf_4ab56a99965952e7.HasNext()
}

type __obf_dfe61d8c4987b512 struct {
	__obf_2153a234428f3398 *reflect2.UnsafeMapType
	__obf_bf24e60df05dbd7a ValEncoder
	__obf_73c590fa207e5962 ValEncoder
}

func (__obf_008f61596d7e9523 *__obf_dfe61d8c4987b512) Encode(__obf_dbbf371b91136494 unsafe.Pointer, __obf_8a2c51fe22775655 *Stream) {
	if *(*unsafe.Pointer)(__obf_dbbf371b91136494) == nil {
		__obf_8a2c51fe22775655.
			WriteNil()
		return
	}
	__obf_8a2c51fe22775655.
		WriteObjectStart()
	__obf_6b532ffbb7374d2f := __obf_008f61596d7e9523.__obf_2153a234428f3398.UnsafeIterate(__obf_dbbf371b91136494)
	__obf_1b90c9809fcfcb01 := __obf_8a2c51fe22775655.__obf_679611dc56ff465b.BorrowStream(nil)
	__obf_1b90c9809fcfcb01.
		Attachment = __obf_8a2c51fe22775655.Attachment
	__obf_4e82475946772fdc := __obf_8a2c51fe22775655.__obf_679611dc56ff465b.BorrowIterator(nil)
	__obf_bce7ea49343beec7 := __obf_56653a7ff233df96{}
	for __obf_6b532ffbb7374d2f.HasNext() {
		__obf_53d763816351c178, __obf_e4d99e200ff06def := __obf_6b532ffbb7374d2f.UnsafeNext()
		__obf_783e2aacd53ceb12 := __obf_1b90c9809fcfcb01.Buffered()
		__obf_008f61596d7e9523.__obf_bf24e60df05dbd7a.
			Encode(__obf_53d763816351c178, __obf_1b90c9809fcfcb01)
		if __obf_1b90c9809fcfcb01.Error != nil && __obf_1b90c9809fcfcb01.Error != io.EOF && __obf_8a2c51fe22775655.Error == nil {
			__obf_8a2c51fe22775655.
				Error = __obf_1b90c9809fcfcb01.Error
		}
		__obf_b3a60355d7e31a25 := __obf_1b90c9809fcfcb01.Buffer()[__obf_783e2aacd53ceb12:]
		__obf_4e82475946772fdc.
			ResetBytes(__obf_b3a60355d7e31a25)
		__obf_a82df6db6b1bf288 := __obf_4e82475946772fdc.ReadString()
		if __obf_8a2c51fe22775655.__obf_508619ef3f08fbb3 > 0 {
			__obf_1b90c9809fcfcb01.__obf_3635bad429be5857(byte(':'), byte(' '))
		} else {
			__obf_1b90c9809fcfcb01.__obf_41130daa346c0e53(':')
		}
		__obf_008f61596d7e9523.__obf_73c590fa207e5962.
			Encode(__obf_e4d99e200ff06def, __obf_1b90c9809fcfcb01)
		__obf_bce7ea49343beec7 = append(__obf_bce7ea49343beec7, __obf_0a5ef920593d7c5c{__obf_53d763816351c178: __obf_a82df6db6b1bf288, __obf_37ffb80dd19e4d37: __obf_1b90c9809fcfcb01.Buffer()[__obf_783e2aacd53ceb12:]})
	}
	sort.Sort(__obf_bce7ea49343beec7)
	for __obf_82c6e05b9d226c58, __obf_37ffb80dd19e4d37 := range __obf_bce7ea49343beec7 {
		if __obf_82c6e05b9d226c58 != 0 {
			__obf_8a2c51fe22775655.
				WriteMore()
		}
		__obf_8a2c51fe22775655.
			Write(__obf_37ffb80dd19e4d37.__obf_37ffb80dd19e4d37)
	}
	if __obf_1b90c9809fcfcb01.Error != nil && __obf_8a2c51fe22775655.Error == nil {
		__obf_8a2c51fe22775655.
			Error = __obf_1b90c9809fcfcb01.Error
	}
	__obf_8a2c51fe22775655.
		WriteObjectEnd()
	__obf_8a2c51fe22775655.__obf_679611dc56ff465b.
		ReturnStream(__obf_1b90c9809fcfcb01)
	__obf_8a2c51fe22775655.__obf_679611dc56ff465b.
		ReturnIterator(__obf_4e82475946772fdc)
}

func (__obf_008f61596d7e9523 *__obf_dfe61d8c4987b512) IsEmpty(__obf_dbbf371b91136494 unsafe.Pointer) bool {
	__obf_4ab56a99965952e7 := __obf_008f61596d7e9523.__obf_2153a234428f3398.UnsafeIterate(__obf_dbbf371b91136494)
	return !__obf_4ab56a99965952e7.HasNext()
}

type __obf_56653a7ff233df96 []__obf_0a5ef920593d7c5c

type __obf_0a5ef920593d7c5c struct {
	__obf_53d763816351c178 string
	__obf_37ffb80dd19e4d37 []byte
}

func (__obf_5c5e2789837e58e7 __obf_56653a7ff233df96) Len() int { return len(__obf_5c5e2789837e58e7) }
func (__obf_5c5e2789837e58e7 __obf_56653a7ff233df96) Swap(__obf_82c6e05b9d226c58, __obf_354f1fc1cab1d938 int) {
	__obf_5c5e2789837e58e7[__obf_82c6e05b9d226c58], __obf_5c5e2789837e58e7[__obf_354f1fc1cab1d938] = __obf_5c5e2789837e58e7[__obf_354f1fc1cab1d938], __obf_5c5e2789837e58e7[__obf_82c6e05b9d226c58]
}
func (__obf_5c5e2789837e58e7 __obf_56653a7ff233df96) Less(__obf_82c6e05b9d226c58, __obf_354f1fc1cab1d938 int) bool {
	return __obf_5c5e2789837e58e7[__obf_82c6e05b9d226c58].__obf_53d763816351c178 < __obf_5c5e2789837e58e7[__obf_354f1fc1cab1d938].__obf_53d763816351c178
}
