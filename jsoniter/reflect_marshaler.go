package __obf_c3cd04a15c56f32f

import (
	"encoding"
	"encoding/json"
	"unsafe"

	"github.com/modern-go/reflect2"
)

var __obf_b16b276d11f0ab2f = reflect2.TypeOfPtr((*json.Marshaler)(nil)).Elem()
var __obf_31f96b637210e6eb = reflect2.TypeOfPtr((*json.Unmarshaler)(nil)).Elem()
var __obf_118e2643fe9fc774 = reflect2.TypeOfPtr((*encoding.TextMarshaler)(nil)).Elem()
var __obf_b7c5ad5d21139f9f = reflect2.TypeOfPtr((*encoding.TextUnmarshaler)(nil)).Elem()

func __obf_2b6c587dc5a76788(__obf_62435d89ebefd5aa *__obf_62435d89ebefd5aa, __obf_8667d4fc2a95b0d7 reflect2.Type) ValDecoder {
	__obf_9fd3068ebc25d7f9 := reflect2.PtrTo(__obf_8667d4fc2a95b0d7)
	if __obf_9fd3068ebc25d7f9.Implements(__obf_31f96b637210e6eb) {
		return &__obf_cec111c41697b8e9{
			&__obf_646b7564744762af{__obf_9fd3068ebc25d7f9},
		}
	}
	if __obf_9fd3068ebc25d7f9.Implements(__obf_b7c5ad5d21139f9f) {
		return &__obf_cec111c41697b8e9{
			&__obf_d73f3e91d310f340{__obf_9fd3068ebc25d7f9},
		}
	}
	return nil
}

func __obf_34128745798a55d1(__obf_62435d89ebefd5aa *__obf_62435d89ebefd5aa, __obf_8667d4fc2a95b0d7 reflect2.Type) ValEncoder {
	if __obf_8667d4fc2a95b0d7 == __obf_b16b276d11f0ab2f {
		__obf_7cb1bad9ea4a8234 := __obf_6ef591ca3f1a443c(__obf_62435d89ebefd5aa, __obf_8667d4fc2a95b0d7)
		var __obf_c85b895560db528f ValEncoder = &__obf_322b8d5218335b8c{__obf_7cb1bad9ea4a8234: __obf_7cb1bad9ea4a8234}
		return __obf_c85b895560db528f
	}
	if __obf_8667d4fc2a95b0d7.Implements(__obf_b16b276d11f0ab2f) {
		__obf_7cb1bad9ea4a8234 := __obf_6ef591ca3f1a443c(__obf_62435d89ebefd5aa, __obf_8667d4fc2a95b0d7)
		var __obf_c85b895560db528f ValEncoder = &__obf_e8d8a3c523a99c1f{__obf_797d4fe23b3556f8: __obf_8667d4fc2a95b0d7, __obf_7cb1bad9ea4a8234: __obf_7cb1bad9ea4a8234}
		return __obf_c85b895560db528f
	}
	__obf_9fd3068ebc25d7f9 := reflect2.PtrTo(__obf_8667d4fc2a95b0d7)
	if __obf_62435d89ebefd5aa.__obf_f517983aa5897f07 != "" && __obf_9fd3068ebc25d7f9.Implements(__obf_b16b276d11f0ab2f) {
		__obf_7cb1bad9ea4a8234 := __obf_6ef591ca3f1a443c(__obf_62435d89ebefd5aa, __obf_9fd3068ebc25d7f9)
		var __obf_c85b895560db528f ValEncoder = &__obf_e8d8a3c523a99c1f{__obf_797d4fe23b3556f8: __obf_9fd3068ebc25d7f9, __obf_7cb1bad9ea4a8234: __obf_7cb1bad9ea4a8234}
		return &__obf_19c8951c4dfee8a0{__obf_c85b895560db528f}
	}
	if __obf_8667d4fc2a95b0d7 == __obf_118e2643fe9fc774 {
		__obf_7cb1bad9ea4a8234 := __obf_6ef591ca3f1a443c(__obf_62435d89ebefd5aa, __obf_8667d4fc2a95b0d7)
		var __obf_c85b895560db528f ValEncoder = &__obf_956493b86d70ce45{__obf_7cb1bad9ea4a8234: __obf_7cb1bad9ea4a8234, __obf_de47f56a3aec06da: __obf_62435d89ebefd5aa.EncoderOf(reflect2.TypeOf(""))}
		return __obf_c85b895560db528f
	}
	if __obf_8667d4fc2a95b0d7.Implements(__obf_118e2643fe9fc774) {
		__obf_7cb1bad9ea4a8234 := __obf_6ef591ca3f1a443c(__obf_62435d89ebefd5aa, __obf_8667d4fc2a95b0d7)
		var __obf_c85b895560db528f ValEncoder = &__obf_d7eae9a85745e3c8{__obf_797d4fe23b3556f8: __obf_8667d4fc2a95b0d7, __obf_de47f56a3aec06da: __obf_62435d89ebefd5aa.EncoderOf(reflect2.TypeOf("")), __obf_7cb1bad9ea4a8234: __obf_7cb1bad9ea4a8234}
		return __obf_c85b895560db528f
	}
	// if prefix is empty, the type is the root type
	if __obf_62435d89ebefd5aa.__obf_f517983aa5897f07 != "" && __obf_9fd3068ebc25d7f9.Implements(__obf_118e2643fe9fc774) {
		__obf_7cb1bad9ea4a8234 := __obf_6ef591ca3f1a443c(__obf_62435d89ebefd5aa, __obf_9fd3068ebc25d7f9)
		var __obf_c85b895560db528f ValEncoder = &__obf_d7eae9a85745e3c8{__obf_797d4fe23b3556f8: __obf_9fd3068ebc25d7f9, __obf_de47f56a3aec06da: __obf_62435d89ebefd5aa.EncoderOf(reflect2.TypeOf("")), __obf_7cb1bad9ea4a8234: __obf_7cb1bad9ea4a8234}
		return &__obf_19c8951c4dfee8a0{__obf_c85b895560db528f}
	}
	return nil
}

type __obf_e8d8a3c523a99c1f struct {
	__obf_7cb1bad9ea4a8234 __obf_7cb1bad9ea4a8234
	__obf_797d4fe23b3556f8 reflect2.Type
}

func (__obf_c85b895560db528f *__obf_e8d8a3c523a99c1f) Encode(__obf_753ab3fb72cdd659 unsafe.Pointer, __obf_2361f5214e84c60f *Stream) {
	__obf_50acae8c0a871ba1 := __obf_c85b895560db528f.__obf_797d4fe23b3556f8.UnsafeIndirect(__obf_753ab3fb72cdd659)
	if __obf_c85b895560db528f.__obf_797d4fe23b3556f8.IsNullable() && reflect2.IsNil(__obf_50acae8c0a871ba1) {
		__obf_2361f5214e84c60f.
			WriteNil()
		return
	}
	__obf_86a83a0f40546556 := __obf_50acae8c0a871ba1.(json.Marshaler)
	__obf_d6d6ee2e4eba7d3f, __obf_5966ecc5edb9b17e := __obf_86a83a0f40546556.MarshalJSON()
	if __obf_5966ecc5edb9b17e != nil {
		__obf_2361f5214e84c60f.
			Error = __obf_5966ecc5edb9b17e
	} else {
		__obf_28656aba21654308 := // html escape was already done by jsoniter
			// but the extra '\n' should be trimed
			len(__obf_d6d6ee2e4eba7d3f)
		if __obf_28656aba21654308 > 0 && __obf_d6d6ee2e4eba7d3f[__obf_28656aba21654308-1] == '\n' {
			__obf_d6d6ee2e4eba7d3f = __obf_d6d6ee2e4eba7d3f[:__obf_28656aba21654308-1]
		}
		__obf_2361f5214e84c60f.
			Write(__obf_d6d6ee2e4eba7d3f)
	}
}

func (__obf_c85b895560db528f *__obf_e8d8a3c523a99c1f) IsEmpty(__obf_753ab3fb72cdd659 unsafe.Pointer) bool {
	return __obf_c85b895560db528f.__obf_7cb1bad9ea4a8234.IsEmpty(__obf_753ab3fb72cdd659)
}

type __obf_322b8d5218335b8c struct {
	__obf_7cb1bad9ea4a8234 __obf_7cb1bad9ea4a8234
}

func (__obf_c85b895560db528f *__obf_322b8d5218335b8c) Encode(__obf_753ab3fb72cdd659 unsafe.Pointer, __obf_2361f5214e84c60f *Stream) {
	__obf_86a83a0f40546556 := *(*json.Marshaler)(__obf_753ab3fb72cdd659)
	if __obf_86a83a0f40546556 == nil {
		__obf_2361f5214e84c60f.
			WriteNil()
		return
	}
	__obf_d6d6ee2e4eba7d3f, __obf_5966ecc5edb9b17e := __obf_86a83a0f40546556.MarshalJSON()
	if __obf_5966ecc5edb9b17e != nil {
		__obf_2361f5214e84c60f.
			Error = __obf_5966ecc5edb9b17e
	} else {
		__obf_2361f5214e84c60f.
			Write(__obf_d6d6ee2e4eba7d3f)
	}
}

func (__obf_c85b895560db528f *__obf_322b8d5218335b8c) IsEmpty(__obf_753ab3fb72cdd659 unsafe.Pointer) bool {
	return __obf_c85b895560db528f.__obf_7cb1bad9ea4a8234.IsEmpty(__obf_753ab3fb72cdd659)
}

type __obf_d7eae9a85745e3c8 struct {
	__obf_797d4fe23b3556f8 reflect2.Type
	__obf_de47f56a3aec06da ValEncoder
	__obf_7cb1bad9ea4a8234 __obf_7cb1bad9ea4a8234
}

func (__obf_c85b895560db528f *__obf_d7eae9a85745e3c8) Encode(__obf_753ab3fb72cdd659 unsafe.Pointer, __obf_2361f5214e84c60f *Stream) {
	__obf_50acae8c0a871ba1 := __obf_c85b895560db528f.__obf_797d4fe23b3556f8.UnsafeIndirect(__obf_753ab3fb72cdd659)
	if __obf_c85b895560db528f.__obf_797d4fe23b3556f8.IsNullable() && reflect2.IsNil(__obf_50acae8c0a871ba1) {
		__obf_2361f5214e84c60f.
			WriteNil()
		return
	}
	__obf_86a83a0f40546556 := (__obf_50acae8c0a871ba1).(encoding.TextMarshaler)
	__obf_d6d6ee2e4eba7d3f, __obf_5966ecc5edb9b17e := __obf_86a83a0f40546556.MarshalText()
	if __obf_5966ecc5edb9b17e != nil {
		__obf_2361f5214e84c60f.
			Error = __obf_5966ecc5edb9b17e
	} else {
		__obf_2d944450d48e5810 := string(__obf_d6d6ee2e4eba7d3f)
		__obf_c85b895560db528f.__obf_de47f56a3aec06da.
			Encode(unsafe.Pointer(&__obf_2d944450d48e5810), __obf_2361f5214e84c60f)
	}
}

func (__obf_c85b895560db528f *__obf_d7eae9a85745e3c8) IsEmpty(__obf_753ab3fb72cdd659 unsafe.Pointer) bool {
	return __obf_c85b895560db528f.__obf_7cb1bad9ea4a8234.IsEmpty(__obf_753ab3fb72cdd659)
}

type __obf_956493b86d70ce45 struct {
	__obf_de47f56a3aec06da ValEncoder
	__obf_7cb1bad9ea4a8234 __obf_7cb1bad9ea4a8234
}

func (__obf_c85b895560db528f *__obf_956493b86d70ce45) Encode(__obf_753ab3fb72cdd659 unsafe.Pointer, __obf_2361f5214e84c60f *Stream) {
	__obf_86a83a0f40546556 := *(*encoding.TextMarshaler)(__obf_753ab3fb72cdd659)
	if __obf_86a83a0f40546556 == nil {
		__obf_2361f5214e84c60f.
			WriteNil()
		return
	}
	__obf_d6d6ee2e4eba7d3f, __obf_5966ecc5edb9b17e := __obf_86a83a0f40546556.MarshalText()
	if __obf_5966ecc5edb9b17e != nil {
		__obf_2361f5214e84c60f.
			Error = __obf_5966ecc5edb9b17e
	} else {
		__obf_2d944450d48e5810 := string(__obf_d6d6ee2e4eba7d3f)
		__obf_c85b895560db528f.__obf_de47f56a3aec06da.
			Encode(unsafe.Pointer(&__obf_2d944450d48e5810), __obf_2361f5214e84c60f)
	}
}

func (__obf_c85b895560db528f *__obf_956493b86d70ce45) IsEmpty(__obf_753ab3fb72cdd659 unsafe.Pointer) bool {
	return __obf_c85b895560db528f.__obf_7cb1bad9ea4a8234.IsEmpty(__obf_753ab3fb72cdd659)
}

type __obf_646b7564744762af struct {
	__obf_797d4fe23b3556f8 reflect2.Type
}

func (__obf_924cc7ef585cdfb0 *__obf_646b7564744762af) Decode(__obf_753ab3fb72cdd659 unsafe.Pointer, __obf_edd9fe48d39445e4 *Iterator) {
	__obf_797d4fe23b3556f8 := __obf_924cc7ef585cdfb0.__obf_797d4fe23b3556f8
	__obf_50acae8c0a871ba1 := __obf_797d4fe23b3556f8.UnsafeIndirect(__obf_753ab3fb72cdd659)
	__obf_588894717c4179ce := __obf_50acae8c0a871ba1.(json.Unmarshaler)
	__obf_edd9fe48d39445e4.__obf_a26eeeac1d6f5a11()
	__obf_edd9fe48d39445e4.
		// skip spaces
		__obf_a0622aded2d86ded()
	__obf_d6d6ee2e4eba7d3f := __obf_edd9fe48d39445e4.SkipAndReturnBytes()
	__obf_5966ecc5edb9b17e := __obf_588894717c4179ce.UnmarshalJSON(__obf_d6d6ee2e4eba7d3f)
	if __obf_5966ecc5edb9b17e != nil {
		__obf_edd9fe48d39445e4.
			ReportError("unmarshalerDecoder", __obf_5966ecc5edb9b17e.Error())
	}
}

type __obf_d73f3e91d310f340 struct {
	__obf_797d4fe23b3556f8 reflect2.Type
}

func (__obf_924cc7ef585cdfb0 *__obf_d73f3e91d310f340) Decode(__obf_753ab3fb72cdd659 unsafe.Pointer, __obf_edd9fe48d39445e4 *Iterator) {
	__obf_797d4fe23b3556f8 := __obf_924cc7ef585cdfb0.__obf_797d4fe23b3556f8
	__obf_50acae8c0a871ba1 := __obf_797d4fe23b3556f8.UnsafeIndirect(__obf_753ab3fb72cdd659)
	if reflect2.IsNil(__obf_50acae8c0a871ba1) {
		__obf_9fd3068ebc25d7f9 := __obf_797d4fe23b3556f8.(*reflect2.UnsafePtrType)
		__obf_78b97e7cf83971dd := __obf_9fd3068ebc25d7f9.Elem()
		__obf_991e715ce3758c78 := __obf_78b97e7cf83971dd.UnsafeNew()
		__obf_9fd3068ebc25d7f9.
			UnsafeSet(__obf_753ab3fb72cdd659, unsafe.Pointer(&__obf_991e715ce3758c78))
		__obf_50acae8c0a871ba1 = __obf_797d4fe23b3556f8.UnsafeIndirect(__obf_753ab3fb72cdd659)
	}
	__obf_588894717c4179ce := (__obf_50acae8c0a871ba1).(encoding.TextUnmarshaler)
	__obf_2d944450d48e5810 := __obf_edd9fe48d39445e4.ReadString()
	__obf_5966ecc5edb9b17e := __obf_588894717c4179ce.UnmarshalText([]byte(__obf_2d944450d48e5810))
	if __obf_5966ecc5edb9b17e != nil {
		__obf_edd9fe48d39445e4.
			ReportError("textUnmarshalerDecoder", __obf_5966ecc5edb9b17e.Error())
	}
}
