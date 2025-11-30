package __obf_030d39f7a8de96c6

import (
	"encoding"
	"encoding/json"
	"unsafe"

	"github.com/modern-go/reflect2"
)

var __obf_613216a45fd77451 = reflect2.TypeOfPtr((*json.Marshaler)(nil)).Elem()
var __obf_b7ddfd40cfe62a98 = reflect2.TypeOfPtr((*json.Unmarshaler)(nil)).Elem()
var __obf_5353ad9fd2397e25 = reflect2.TypeOfPtr((*encoding.TextMarshaler)(nil)).Elem()
var __obf_f58653b68e3e7233 = reflect2.TypeOfPtr((*encoding.TextUnmarshaler)(nil)).Elem()

func __obf_a0c12030f23db855(__obf_a565fbaffca944f0 *__obf_a565fbaffca944f0, __obf_8744d0b8c80ccdc1 reflect2.Type) ValDecoder {
	__obf_3a51157620f9910b := reflect2.PtrTo(__obf_8744d0b8c80ccdc1)
	if __obf_3a51157620f9910b.Implements(__obf_b7ddfd40cfe62a98) {
		return &__obf_2e9c42d03f0d8cd5{
			&__obf_cd063ac27c40d5c6{__obf_3a51157620f9910b},
		}
	}
	if __obf_3a51157620f9910b.Implements(__obf_f58653b68e3e7233) {
		return &__obf_2e9c42d03f0d8cd5{
			&__obf_e0e26cbc14bdbad8{__obf_3a51157620f9910b},
		}
	}
	return nil
}

func __obf_35f11390081cc107(__obf_a565fbaffca944f0 *__obf_a565fbaffca944f0, __obf_8744d0b8c80ccdc1 reflect2.Type) ValEncoder {
	if __obf_8744d0b8c80ccdc1 == __obf_613216a45fd77451 {
		__obf_b5475e6763577f46 := __obf_adf48ba0507caf28(__obf_a565fbaffca944f0, __obf_8744d0b8c80ccdc1)
		var __obf_008f61596d7e9523 ValEncoder = &__obf_b8430fac36f72789{__obf_b5475e6763577f46: __obf_b5475e6763577f46}
		return __obf_008f61596d7e9523
	}
	if __obf_8744d0b8c80ccdc1.Implements(__obf_613216a45fd77451) {
		__obf_b5475e6763577f46 := __obf_adf48ba0507caf28(__obf_a565fbaffca944f0, __obf_8744d0b8c80ccdc1)
		var __obf_008f61596d7e9523 ValEncoder = &__obf_1b29218d02f270d6{__obf_a7610e23a63622fd: __obf_8744d0b8c80ccdc1, __obf_b5475e6763577f46: __obf_b5475e6763577f46}
		return __obf_008f61596d7e9523
	}
	__obf_3a51157620f9910b := reflect2.PtrTo(__obf_8744d0b8c80ccdc1)
	if __obf_a565fbaffca944f0.__obf_834e1679b8081f46 != "" && __obf_3a51157620f9910b.Implements(__obf_613216a45fd77451) {
		__obf_b5475e6763577f46 := __obf_adf48ba0507caf28(__obf_a565fbaffca944f0, __obf_3a51157620f9910b)
		var __obf_008f61596d7e9523 ValEncoder = &__obf_1b29218d02f270d6{__obf_a7610e23a63622fd: __obf_3a51157620f9910b, __obf_b5475e6763577f46: __obf_b5475e6763577f46}
		return &__obf_d086959fc9847188{__obf_008f61596d7e9523}
	}
	if __obf_8744d0b8c80ccdc1 == __obf_5353ad9fd2397e25 {
		__obf_b5475e6763577f46 := __obf_adf48ba0507caf28(__obf_a565fbaffca944f0, __obf_8744d0b8c80ccdc1)
		var __obf_008f61596d7e9523 ValEncoder = &__obf_be22d3d9cea5924c{__obf_b5475e6763577f46: __obf_b5475e6763577f46, __obf_6e4d12bda4dd7e33: __obf_a565fbaffca944f0.EncoderOf(reflect2.TypeOf(""))}
		return __obf_008f61596d7e9523
	}
	if __obf_8744d0b8c80ccdc1.Implements(__obf_5353ad9fd2397e25) {
		__obf_b5475e6763577f46 := __obf_adf48ba0507caf28(__obf_a565fbaffca944f0, __obf_8744d0b8c80ccdc1)
		var __obf_008f61596d7e9523 ValEncoder = &__obf_1632fdd756ee8530{__obf_a7610e23a63622fd: __obf_8744d0b8c80ccdc1, __obf_6e4d12bda4dd7e33: __obf_a565fbaffca944f0.EncoderOf(reflect2.TypeOf("")), __obf_b5475e6763577f46: __obf_b5475e6763577f46}
		return __obf_008f61596d7e9523
	}
	// if prefix is empty, the type is the root type
	if __obf_a565fbaffca944f0.__obf_834e1679b8081f46 != "" && __obf_3a51157620f9910b.Implements(__obf_5353ad9fd2397e25) {
		__obf_b5475e6763577f46 := __obf_adf48ba0507caf28(__obf_a565fbaffca944f0, __obf_3a51157620f9910b)
		var __obf_008f61596d7e9523 ValEncoder = &__obf_1632fdd756ee8530{__obf_a7610e23a63622fd: __obf_3a51157620f9910b, __obf_6e4d12bda4dd7e33: __obf_a565fbaffca944f0.EncoderOf(reflect2.TypeOf("")), __obf_b5475e6763577f46: __obf_b5475e6763577f46}
		return &__obf_d086959fc9847188{__obf_008f61596d7e9523}
	}
	return nil
}

type __obf_1b29218d02f270d6 struct {
	__obf_b5475e6763577f46 __obf_b5475e6763577f46
	__obf_a7610e23a63622fd reflect2.Type
}

func (__obf_008f61596d7e9523 *__obf_1b29218d02f270d6) Encode(__obf_dbbf371b91136494 unsafe.Pointer, __obf_8a2c51fe22775655 *Stream) {
	__obf_eefa92392cc2442c := __obf_008f61596d7e9523.__obf_a7610e23a63622fd.UnsafeIndirect(__obf_dbbf371b91136494)
	if __obf_008f61596d7e9523.__obf_a7610e23a63622fd.IsNullable() && reflect2.IsNil(__obf_eefa92392cc2442c) {
		__obf_8a2c51fe22775655.
			WriteNil()
		return
	}
	__obf_69740b27cd33e616 := __obf_eefa92392cc2442c.(json.Marshaler)
	__obf_3e6f7e026dccbbaf, __obf_fcc907dd69879566 := __obf_69740b27cd33e616.MarshalJSON()
	if __obf_fcc907dd69879566 != nil {
		__obf_8a2c51fe22775655.
			Error = __obf_fcc907dd69879566
	} else {
		__obf_5fddcb0c561b8bd1 := // html escape was already done by jsoniter
			// but the extra '\n' should be trimed
			len(__obf_3e6f7e026dccbbaf)
		if __obf_5fddcb0c561b8bd1 > 0 && __obf_3e6f7e026dccbbaf[__obf_5fddcb0c561b8bd1-1] == '\n' {
			__obf_3e6f7e026dccbbaf = __obf_3e6f7e026dccbbaf[:__obf_5fddcb0c561b8bd1-1]
		}
		__obf_8a2c51fe22775655.
			Write(__obf_3e6f7e026dccbbaf)
	}
}

func (__obf_008f61596d7e9523 *__obf_1b29218d02f270d6) IsEmpty(__obf_dbbf371b91136494 unsafe.Pointer) bool {
	return __obf_008f61596d7e9523.__obf_b5475e6763577f46.IsEmpty(__obf_dbbf371b91136494)
}

type __obf_b8430fac36f72789 struct {
	__obf_b5475e6763577f46 __obf_b5475e6763577f46
}

func (__obf_008f61596d7e9523 *__obf_b8430fac36f72789) Encode(__obf_dbbf371b91136494 unsafe.Pointer, __obf_8a2c51fe22775655 *Stream) {
	__obf_69740b27cd33e616 := *(*json.Marshaler)(__obf_dbbf371b91136494)
	if __obf_69740b27cd33e616 == nil {
		__obf_8a2c51fe22775655.
			WriteNil()
		return
	}
	__obf_3e6f7e026dccbbaf, __obf_fcc907dd69879566 := __obf_69740b27cd33e616.MarshalJSON()
	if __obf_fcc907dd69879566 != nil {
		__obf_8a2c51fe22775655.
			Error = __obf_fcc907dd69879566
	} else {
		__obf_8a2c51fe22775655.
			Write(__obf_3e6f7e026dccbbaf)
	}
}

func (__obf_008f61596d7e9523 *__obf_b8430fac36f72789) IsEmpty(__obf_dbbf371b91136494 unsafe.Pointer) bool {
	return __obf_008f61596d7e9523.__obf_b5475e6763577f46.IsEmpty(__obf_dbbf371b91136494)
}

type __obf_1632fdd756ee8530 struct {
	__obf_a7610e23a63622fd reflect2.Type
	__obf_6e4d12bda4dd7e33 ValEncoder
	__obf_b5475e6763577f46 __obf_b5475e6763577f46
}

func (__obf_008f61596d7e9523 *__obf_1632fdd756ee8530) Encode(__obf_dbbf371b91136494 unsafe.Pointer, __obf_8a2c51fe22775655 *Stream) {
	__obf_eefa92392cc2442c := __obf_008f61596d7e9523.__obf_a7610e23a63622fd.UnsafeIndirect(__obf_dbbf371b91136494)
	if __obf_008f61596d7e9523.__obf_a7610e23a63622fd.IsNullable() && reflect2.IsNil(__obf_eefa92392cc2442c) {
		__obf_8a2c51fe22775655.
			WriteNil()
		return
	}
	__obf_69740b27cd33e616 := (__obf_eefa92392cc2442c).(encoding.TextMarshaler)
	__obf_3e6f7e026dccbbaf, __obf_fcc907dd69879566 := __obf_69740b27cd33e616.MarshalText()
	if __obf_fcc907dd69879566 != nil {
		__obf_8a2c51fe22775655.
			Error = __obf_fcc907dd69879566
	} else {
		__obf_428c3b4a95725c4a := string(__obf_3e6f7e026dccbbaf)
		__obf_008f61596d7e9523.__obf_6e4d12bda4dd7e33.
			Encode(unsafe.Pointer(&__obf_428c3b4a95725c4a), __obf_8a2c51fe22775655)
	}
}

func (__obf_008f61596d7e9523 *__obf_1632fdd756ee8530) IsEmpty(__obf_dbbf371b91136494 unsafe.Pointer) bool {
	return __obf_008f61596d7e9523.__obf_b5475e6763577f46.IsEmpty(__obf_dbbf371b91136494)
}

type __obf_be22d3d9cea5924c struct {
	__obf_6e4d12bda4dd7e33 ValEncoder
	__obf_b5475e6763577f46 __obf_b5475e6763577f46
}

func (__obf_008f61596d7e9523 *__obf_be22d3d9cea5924c) Encode(__obf_dbbf371b91136494 unsafe.Pointer, __obf_8a2c51fe22775655 *Stream) {
	__obf_69740b27cd33e616 := *(*encoding.TextMarshaler)(__obf_dbbf371b91136494)
	if __obf_69740b27cd33e616 == nil {
		__obf_8a2c51fe22775655.
			WriteNil()
		return
	}
	__obf_3e6f7e026dccbbaf, __obf_fcc907dd69879566 := __obf_69740b27cd33e616.MarshalText()
	if __obf_fcc907dd69879566 != nil {
		__obf_8a2c51fe22775655.
			Error = __obf_fcc907dd69879566
	} else {
		__obf_428c3b4a95725c4a := string(__obf_3e6f7e026dccbbaf)
		__obf_008f61596d7e9523.__obf_6e4d12bda4dd7e33.
			Encode(unsafe.Pointer(&__obf_428c3b4a95725c4a), __obf_8a2c51fe22775655)
	}
}

func (__obf_008f61596d7e9523 *__obf_be22d3d9cea5924c) IsEmpty(__obf_dbbf371b91136494 unsafe.Pointer) bool {
	return __obf_008f61596d7e9523.__obf_b5475e6763577f46.IsEmpty(__obf_dbbf371b91136494)
}

type __obf_cd063ac27c40d5c6 struct {
	__obf_a7610e23a63622fd reflect2.Type
}

func (__obf_11a3f28116164d09 *__obf_cd063ac27c40d5c6) Decode(__obf_dbbf371b91136494 unsafe.Pointer, __obf_4ab56a99965952e7 *Iterator) {
	__obf_a7610e23a63622fd := __obf_11a3f28116164d09.__obf_a7610e23a63622fd
	__obf_eefa92392cc2442c := __obf_a7610e23a63622fd.UnsafeIndirect(__obf_dbbf371b91136494)
	__obf_d32d7612a787ce59 := __obf_eefa92392cc2442c.(json.Unmarshaler)
	__obf_4ab56a99965952e7.__obf_61df301d1f67ad73()
	__obf_4ab56a99965952e7.
		// skip spaces
		__obf_ba6d0be9a7ab6ae4()
	__obf_3e6f7e026dccbbaf := __obf_4ab56a99965952e7.SkipAndReturnBytes()
	__obf_fcc907dd69879566 := __obf_d32d7612a787ce59.UnmarshalJSON(__obf_3e6f7e026dccbbaf)
	if __obf_fcc907dd69879566 != nil {
		__obf_4ab56a99965952e7.
			ReportError("unmarshalerDecoder", __obf_fcc907dd69879566.Error())
	}
}

type __obf_e0e26cbc14bdbad8 struct {
	__obf_a7610e23a63622fd reflect2.Type
}

func (__obf_11a3f28116164d09 *__obf_e0e26cbc14bdbad8) Decode(__obf_dbbf371b91136494 unsafe.Pointer, __obf_4ab56a99965952e7 *Iterator) {
	__obf_a7610e23a63622fd := __obf_11a3f28116164d09.__obf_a7610e23a63622fd
	__obf_eefa92392cc2442c := __obf_a7610e23a63622fd.UnsafeIndirect(__obf_dbbf371b91136494)
	if reflect2.IsNil(__obf_eefa92392cc2442c) {
		__obf_3a51157620f9910b := __obf_a7610e23a63622fd.(*reflect2.UnsafePtrType)
		__obf_d4fd7b4aa577e66f := __obf_3a51157620f9910b.Elem()
		__obf_e4d99e200ff06def := __obf_d4fd7b4aa577e66f.UnsafeNew()
		__obf_3a51157620f9910b.
			UnsafeSet(__obf_dbbf371b91136494, unsafe.Pointer(&__obf_e4d99e200ff06def))
		__obf_eefa92392cc2442c = __obf_a7610e23a63622fd.UnsafeIndirect(__obf_dbbf371b91136494)
	}
	__obf_d32d7612a787ce59 := (__obf_eefa92392cc2442c).(encoding.TextUnmarshaler)
	__obf_428c3b4a95725c4a := __obf_4ab56a99965952e7.ReadString()
	__obf_fcc907dd69879566 := __obf_d32d7612a787ce59.UnmarshalText([]byte(__obf_428c3b4a95725c4a))
	if __obf_fcc907dd69879566 != nil {
		__obf_4ab56a99965952e7.
			ReportError("textUnmarshalerDecoder", __obf_fcc907dd69879566.Error())
	}
}
