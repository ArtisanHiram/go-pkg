package __obf_5b802ce8d9ba56d6

import (
	"encoding"
	"encoding/json"
	"unsafe"

	"github.com/modern-go/reflect2"
)

var __obf_a92e19d4213e7336 = reflect2.TypeOfPtr((*json.Marshaler)(nil)).Elem()
var __obf_738fff33396f272d = reflect2.TypeOfPtr((*json.Unmarshaler)(nil)).Elem()
var __obf_9dca2b8aa916a9d0 = reflect2.TypeOfPtr((*encoding.TextMarshaler)(nil)).Elem()
var __obf_5cbbb40316a5c06a = reflect2.TypeOfPtr((*encoding.TextUnmarshaler)(nil)).Elem()

func __obf_308321b3ba24eb8c(__obf_08da24be66d0d13c *__obf_08da24be66d0d13c, __obf_5efc66d8c338c133 reflect2.Type) ValDecoder {
	__obf_d0cac7bfcf0092ea := reflect2.PtrTo(__obf_5efc66d8c338c133)
	if __obf_d0cac7bfcf0092ea.Implements(__obf_738fff33396f272d) {
		return &__obf_46f8310df6424f3f{
			&__obf_1ce729a0ec8c5634{__obf_d0cac7bfcf0092ea},
		}
	}
	if __obf_d0cac7bfcf0092ea.Implements(__obf_5cbbb40316a5c06a) {
		return &__obf_46f8310df6424f3f{
			&__obf_646b36ac5dd53924{__obf_d0cac7bfcf0092ea},
		}
	}
	return nil
}

func __obf_238d567082abf8bc(__obf_08da24be66d0d13c *__obf_08da24be66d0d13c, __obf_5efc66d8c338c133 reflect2.Type) ValEncoder {
	if __obf_5efc66d8c338c133 == __obf_a92e19d4213e7336 {
		__obf_9cd9bb2e92b7d8c2 := __obf_35a60ba99630ce31(__obf_08da24be66d0d13c, __obf_5efc66d8c338c133)
		var __obf_29366c3ad76a8c51 ValEncoder = &__obf_cdbe4393cc3f8836{__obf_9cd9bb2e92b7d8c2: __obf_9cd9bb2e92b7d8c2}
		return __obf_29366c3ad76a8c51
	}
	if __obf_5efc66d8c338c133.Implements(__obf_a92e19d4213e7336) {
		__obf_9cd9bb2e92b7d8c2 := __obf_35a60ba99630ce31(__obf_08da24be66d0d13c, __obf_5efc66d8c338c133)
		var __obf_29366c3ad76a8c51 ValEncoder = &__obf_4dcb8b151485413d{__obf_5e777ac7034ab220: __obf_5efc66d8c338c133, __obf_9cd9bb2e92b7d8c2: __obf_9cd9bb2e92b7d8c2}
		return __obf_29366c3ad76a8c51
	}
	__obf_d0cac7bfcf0092ea := reflect2.PtrTo(__obf_5efc66d8c338c133)
	if __obf_08da24be66d0d13c.__obf_f3dd783c02964acf != "" && __obf_d0cac7bfcf0092ea.Implements(__obf_a92e19d4213e7336) {
		__obf_9cd9bb2e92b7d8c2 := __obf_35a60ba99630ce31(__obf_08da24be66d0d13c, __obf_d0cac7bfcf0092ea)
		var __obf_29366c3ad76a8c51 ValEncoder = &__obf_4dcb8b151485413d{__obf_5e777ac7034ab220: __obf_d0cac7bfcf0092ea, __obf_9cd9bb2e92b7d8c2: __obf_9cd9bb2e92b7d8c2}
		return &__obf_f668640819fd46fc{__obf_29366c3ad76a8c51}
	}
	if __obf_5efc66d8c338c133 == __obf_9dca2b8aa916a9d0 {
		__obf_9cd9bb2e92b7d8c2 := __obf_35a60ba99630ce31(__obf_08da24be66d0d13c, __obf_5efc66d8c338c133)
		var __obf_29366c3ad76a8c51 ValEncoder = &__obf_8201f27af3b81439{__obf_9cd9bb2e92b7d8c2: __obf_9cd9bb2e92b7d8c2, __obf_29006411c501f852: __obf_08da24be66d0d13c.EncoderOf(reflect2.TypeOf(""))}
		return __obf_29366c3ad76a8c51
	}
	if __obf_5efc66d8c338c133.Implements(__obf_9dca2b8aa916a9d0) {
		__obf_9cd9bb2e92b7d8c2 := __obf_35a60ba99630ce31(__obf_08da24be66d0d13c, __obf_5efc66d8c338c133)
		var __obf_29366c3ad76a8c51 ValEncoder = &__obf_3290483a33b841e1{__obf_5e777ac7034ab220: __obf_5efc66d8c338c133, __obf_29006411c501f852: __obf_08da24be66d0d13c.EncoderOf(reflect2.TypeOf("")), __obf_9cd9bb2e92b7d8c2: __obf_9cd9bb2e92b7d8c2}
		return __obf_29366c3ad76a8c51
	}
	// if prefix is empty, the type is the root type
	if __obf_08da24be66d0d13c.__obf_f3dd783c02964acf != "" && __obf_d0cac7bfcf0092ea.Implements(__obf_9dca2b8aa916a9d0) {
		__obf_9cd9bb2e92b7d8c2 := __obf_35a60ba99630ce31(__obf_08da24be66d0d13c, __obf_d0cac7bfcf0092ea)
		var __obf_29366c3ad76a8c51 ValEncoder = &__obf_3290483a33b841e1{__obf_5e777ac7034ab220: __obf_d0cac7bfcf0092ea, __obf_29006411c501f852: __obf_08da24be66d0d13c.EncoderOf(reflect2.TypeOf("")), __obf_9cd9bb2e92b7d8c2: __obf_9cd9bb2e92b7d8c2}
		return &__obf_f668640819fd46fc{__obf_29366c3ad76a8c51}
	}
	return nil
}

type __obf_4dcb8b151485413d struct {
	__obf_9cd9bb2e92b7d8c2 __obf_9cd9bb2e92b7d8c2
	__obf_5e777ac7034ab220 reflect2.Type
}

func (__obf_29366c3ad76a8c51 *__obf_4dcb8b151485413d) Encode(__obf_d3c919547bf7e47a unsafe.Pointer, __obf_00fc491caa967a74 *Stream) {
	__obf_f9b1526531f3c024 := __obf_29366c3ad76a8c51.__obf_5e777ac7034ab220.UnsafeIndirect(__obf_d3c919547bf7e47a)
	if __obf_29366c3ad76a8c51.__obf_5e777ac7034ab220.IsNullable() && reflect2.IsNil(__obf_f9b1526531f3c024) {
		__obf_00fc491caa967a74.
			WriteNil()
		return
	}
	__obf_8e41caaad555a905 := __obf_f9b1526531f3c024.(json.Marshaler)
	__obf_791123f235962b4b, __obf_6d071d48f3b321ad := __obf_8e41caaad555a905.MarshalJSON()
	if __obf_6d071d48f3b321ad != nil {
		__obf_00fc491caa967a74.
			Error = __obf_6d071d48f3b321ad
	} else {
		__obf_7fdd64b474fd171a := // html escape was already done by jsoniter
			// but the extra '\n' should be trimed
			len(__obf_791123f235962b4b)
		if __obf_7fdd64b474fd171a > 0 && __obf_791123f235962b4b[__obf_7fdd64b474fd171a-1] == '\n' {
			__obf_791123f235962b4b = __obf_791123f235962b4b[:__obf_7fdd64b474fd171a-1]
		}
		__obf_00fc491caa967a74.
			Write(__obf_791123f235962b4b)
	}
}

func (__obf_29366c3ad76a8c51 *__obf_4dcb8b151485413d) IsEmpty(__obf_d3c919547bf7e47a unsafe.Pointer) bool {
	return __obf_29366c3ad76a8c51.__obf_9cd9bb2e92b7d8c2.IsEmpty(__obf_d3c919547bf7e47a)
}

type __obf_cdbe4393cc3f8836 struct {
	__obf_9cd9bb2e92b7d8c2 __obf_9cd9bb2e92b7d8c2
}

func (__obf_29366c3ad76a8c51 *__obf_cdbe4393cc3f8836) Encode(__obf_d3c919547bf7e47a unsafe.Pointer, __obf_00fc491caa967a74 *Stream) {
	__obf_8e41caaad555a905 := *(*json.Marshaler)(__obf_d3c919547bf7e47a)
	if __obf_8e41caaad555a905 == nil {
		__obf_00fc491caa967a74.
			WriteNil()
		return
	}
	__obf_791123f235962b4b, __obf_6d071d48f3b321ad := __obf_8e41caaad555a905.MarshalJSON()
	if __obf_6d071d48f3b321ad != nil {
		__obf_00fc491caa967a74.
			Error = __obf_6d071d48f3b321ad
	} else {
		__obf_00fc491caa967a74.
			Write(__obf_791123f235962b4b)
	}
}

func (__obf_29366c3ad76a8c51 *__obf_cdbe4393cc3f8836) IsEmpty(__obf_d3c919547bf7e47a unsafe.Pointer) bool {
	return __obf_29366c3ad76a8c51.__obf_9cd9bb2e92b7d8c2.IsEmpty(__obf_d3c919547bf7e47a)
}

type __obf_3290483a33b841e1 struct {
	__obf_5e777ac7034ab220 reflect2.Type
	__obf_29006411c501f852 ValEncoder
	__obf_9cd9bb2e92b7d8c2 __obf_9cd9bb2e92b7d8c2
}

func (__obf_29366c3ad76a8c51 *__obf_3290483a33b841e1) Encode(__obf_d3c919547bf7e47a unsafe.Pointer, __obf_00fc491caa967a74 *Stream) {
	__obf_f9b1526531f3c024 := __obf_29366c3ad76a8c51.__obf_5e777ac7034ab220.UnsafeIndirect(__obf_d3c919547bf7e47a)
	if __obf_29366c3ad76a8c51.__obf_5e777ac7034ab220.IsNullable() && reflect2.IsNil(__obf_f9b1526531f3c024) {
		__obf_00fc491caa967a74.
			WriteNil()
		return
	}
	__obf_8e41caaad555a905 := (__obf_f9b1526531f3c024).(encoding.TextMarshaler)
	__obf_791123f235962b4b, __obf_6d071d48f3b321ad := __obf_8e41caaad555a905.MarshalText()
	if __obf_6d071d48f3b321ad != nil {
		__obf_00fc491caa967a74.
			Error = __obf_6d071d48f3b321ad
	} else {
		__obf_12c21b79fa86dcba := string(__obf_791123f235962b4b)
		__obf_29366c3ad76a8c51.__obf_29006411c501f852.
			Encode(unsafe.Pointer(&__obf_12c21b79fa86dcba), __obf_00fc491caa967a74)
	}
}

func (__obf_29366c3ad76a8c51 *__obf_3290483a33b841e1) IsEmpty(__obf_d3c919547bf7e47a unsafe.Pointer) bool {
	return __obf_29366c3ad76a8c51.__obf_9cd9bb2e92b7d8c2.IsEmpty(__obf_d3c919547bf7e47a)
}

type __obf_8201f27af3b81439 struct {
	__obf_29006411c501f852 ValEncoder
	__obf_9cd9bb2e92b7d8c2 __obf_9cd9bb2e92b7d8c2
}

func (__obf_29366c3ad76a8c51 *__obf_8201f27af3b81439) Encode(__obf_d3c919547bf7e47a unsafe.Pointer, __obf_00fc491caa967a74 *Stream) {
	__obf_8e41caaad555a905 := *(*encoding.TextMarshaler)(__obf_d3c919547bf7e47a)
	if __obf_8e41caaad555a905 == nil {
		__obf_00fc491caa967a74.
			WriteNil()
		return
	}
	__obf_791123f235962b4b, __obf_6d071d48f3b321ad := __obf_8e41caaad555a905.MarshalText()
	if __obf_6d071d48f3b321ad != nil {
		__obf_00fc491caa967a74.
			Error = __obf_6d071d48f3b321ad
	} else {
		__obf_12c21b79fa86dcba := string(__obf_791123f235962b4b)
		__obf_29366c3ad76a8c51.__obf_29006411c501f852.
			Encode(unsafe.Pointer(&__obf_12c21b79fa86dcba), __obf_00fc491caa967a74)
	}
}

func (__obf_29366c3ad76a8c51 *__obf_8201f27af3b81439) IsEmpty(__obf_d3c919547bf7e47a unsafe.Pointer) bool {
	return __obf_29366c3ad76a8c51.__obf_9cd9bb2e92b7d8c2.IsEmpty(__obf_d3c919547bf7e47a)
}

type __obf_1ce729a0ec8c5634 struct {
	__obf_5e777ac7034ab220 reflect2.Type
}

func (__obf_18f746d7b5b440ee *__obf_1ce729a0ec8c5634) Decode(__obf_d3c919547bf7e47a unsafe.Pointer, __obf_67008a6a9e5ba828 *Iterator) {
	__obf_5e777ac7034ab220 := __obf_18f746d7b5b440ee.__obf_5e777ac7034ab220
	__obf_f9b1526531f3c024 := __obf_5e777ac7034ab220.UnsafeIndirect(__obf_d3c919547bf7e47a)
	__obf_1afabbb7bd9b92a8 := __obf_f9b1526531f3c024.(json.Unmarshaler)
	__obf_67008a6a9e5ba828.__obf_b781a59d5a0d2490()
	__obf_67008a6a9e5ba828.
		// skip spaces
		__obf_3284a1eaa2a0abb6()
	__obf_791123f235962b4b := __obf_67008a6a9e5ba828.SkipAndReturnBytes()
	__obf_6d071d48f3b321ad := __obf_1afabbb7bd9b92a8.UnmarshalJSON(__obf_791123f235962b4b)
	if __obf_6d071d48f3b321ad != nil {
		__obf_67008a6a9e5ba828.
			ReportError("unmarshalerDecoder", __obf_6d071d48f3b321ad.Error())
	}
}

type __obf_646b36ac5dd53924 struct {
	__obf_5e777ac7034ab220 reflect2.Type
}

func (__obf_18f746d7b5b440ee *__obf_646b36ac5dd53924) Decode(__obf_d3c919547bf7e47a unsafe.Pointer, __obf_67008a6a9e5ba828 *Iterator) {
	__obf_5e777ac7034ab220 := __obf_18f746d7b5b440ee.__obf_5e777ac7034ab220
	__obf_f9b1526531f3c024 := __obf_5e777ac7034ab220.UnsafeIndirect(__obf_d3c919547bf7e47a)
	if reflect2.IsNil(__obf_f9b1526531f3c024) {
		__obf_d0cac7bfcf0092ea := __obf_5e777ac7034ab220.(*reflect2.UnsafePtrType)
		__obf_0b19d2b7ea8e1be5 := __obf_d0cac7bfcf0092ea.Elem()
		__obf_77670b1f1899744d := __obf_0b19d2b7ea8e1be5.UnsafeNew()
		__obf_d0cac7bfcf0092ea.
			UnsafeSet(__obf_d3c919547bf7e47a, unsafe.Pointer(&__obf_77670b1f1899744d))
		__obf_f9b1526531f3c024 = __obf_5e777ac7034ab220.UnsafeIndirect(__obf_d3c919547bf7e47a)
	}
	__obf_1afabbb7bd9b92a8 := (__obf_f9b1526531f3c024).(encoding.TextUnmarshaler)
	__obf_12c21b79fa86dcba := __obf_67008a6a9e5ba828.ReadString()
	__obf_6d071d48f3b321ad := __obf_1afabbb7bd9b92a8.UnmarshalText([]byte(__obf_12c21b79fa86dcba))
	if __obf_6d071d48f3b321ad != nil {
		__obf_67008a6a9e5ba828.
			ReportError("textUnmarshalerDecoder", __obf_6d071d48f3b321ad.Error())
	}
}
