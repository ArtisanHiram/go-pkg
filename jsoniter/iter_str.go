package __obf_030d39f7a8de96c6

import (
	"fmt"
	"unicode/utf16"
)

// ReadString read string from iterator
func (__obf_4ab56a99965952e7 *Iterator) ReadString() (__obf_59c72400ec1a6110 string) {
	__obf_24309b3b0ff9ba22 := __obf_4ab56a99965952e7.__obf_61df301d1f67ad73()
	if __obf_24309b3b0ff9ba22 == '"' {
		for __obf_82c6e05b9d226c58 := __obf_4ab56a99965952e7.__obf_5e22d6270d27491f; __obf_82c6e05b9d226c58 < __obf_4ab56a99965952e7.__obf_7c17042d9b4e73cc; __obf_82c6e05b9d226c58++ {
			__obf_24309b3b0ff9ba22 := __obf_4ab56a99965952e7.__obf_0b1656d7ef4bd9df[__obf_82c6e05b9d226c58]
			if __obf_24309b3b0ff9ba22 == '"' {
				__obf_59c72400ec1a6110 = string(__obf_4ab56a99965952e7.__obf_0b1656d7ef4bd9df[__obf_4ab56a99965952e7.__obf_5e22d6270d27491f:__obf_82c6e05b9d226c58])
				__obf_4ab56a99965952e7.__obf_5e22d6270d27491f = __obf_82c6e05b9d226c58 + 1
				return __obf_59c72400ec1a6110
			} else if __obf_24309b3b0ff9ba22 == '\\' {
				break
			} else if __obf_24309b3b0ff9ba22 < ' ' {
				__obf_4ab56a99965952e7.
					ReportError("ReadString",
						fmt.Sprintf(`invalid control character found: %d`, __obf_24309b3b0ff9ba22))
				return
			}
		}
		return __obf_4ab56a99965952e7.__obf_506d8d9b7d95ca67()
	} else if __obf_24309b3b0ff9ba22 == 'n' {
		__obf_4ab56a99965952e7.__obf_01038f7a5ba9ca56('u', 'l', 'l')
		return ""
	}
	__obf_4ab56a99965952e7.
		ReportError("ReadString", `expects " or n, but found `+string([]byte{__obf_24309b3b0ff9ba22}))
	return
}

func (__obf_4ab56a99965952e7 *Iterator) __obf_506d8d9b7d95ca67() (__obf_59c72400ec1a6110 string) {
	var __obf_428c3b4a95725c4a []byte
	var __obf_24309b3b0ff9ba22 byte
	for __obf_4ab56a99965952e7.Error == nil {
		__obf_24309b3b0ff9ba22 = __obf_4ab56a99965952e7.__obf_2b6cf70e8dd32b68()
		if __obf_24309b3b0ff9ba22 == '"' {
			return string(__obf_428c3b4a95725c4a)
		}
		if __obf_24309b3b0ff9ba22 == '\\' {
			__obf_24309b3b0ff9ba22 = __obf_4ab56a99965952e7.__obf_2b6cf70e8dd32b68()
			__obf_428c3b4a95725c4a = __obf_4ab56a99965952e7.__obf_4728db252a465f75(__obf_24309b3b0ff9ba22, __obf_428c3b4a95725c4a)
		} else {
			__obf_428c3b4a95725c4a = append(__obf_428c3b4a95725c4a, __obf_24309b3b0ff9ba22)
		}
	}
	__obf_4ab56a99965952e7.
		ReportError("readStringSlowPath", "unexpected end of input")
	return
}

func (__obf_4ab56a99965952e7 *Iterator) __obf_4728db252a465f75(__obf_24309b3b0ff9ba22 byte, __obf_428c3b4a95725c4a []byte) []byte {
	switch __obf_24309b3b0ff9ba22 {
	case 'u':
		__obf_9aca26302d2c5887 := __obf_4ab56a99965952e7.__obf_4a65d6eac537eda8()
		if utf16.IsSurrogate(__obf_9aca26302d2c5887) {
			__obf_24309b3b0ff9ba22 = __obf_4ab56a99965952e7.__obf_2b6cf70e8dd32b68()
			if __obf_4ab56a99965952e7.Error != nil {
				return nil
			}
			if __obf_24309b3b0ff9ba22 != '\\' {
				__obf_4ab56a99965952e7.__obf_ba6d0be9a7ab6ae4()
				__obf_428c3b4a95725c4a = __obf_f5dbf537a0450c02(__obf_428c3b4a95725c4a, __obf_9aca26302d2c5887)
				return __obf_428c3b4a95725c4a
			}
			__obf_24309b3b0ff9ba22 = __obf_4ab56a99965952e7.__obf_2b6cf70e8dd32b68()
			if __obf_4ab56a99965952e7.Error != nil {
				return nil
			}
			if __obf_24309b3b0ff9ba22 != 'u' {
				__obf_428c3b4a95725c4a = __obf_f5dbf537a0450c02(__obf_428c3b4a95725c4a, __obf_9aca26302d2c5887)
				return __obf_4ab56a99965952e7.__obf_4728db252a465f75(__obf_24309b3b0ff9ba22, __obf_428c3b4a95725c4a)
			}
			__obf_12837937a8e321f5 := __obf_4ab56a99965952e7.__obf_4a65d6eac537eda8()
			if __obf_4ab56a99965952e7.Error != nil {
				return nil
			}
			__obf_fde97cbb580f6f76 := utf16.DecodeRune(__obf_9aca26302d2c5887, __obf_12837937a8e321f5)
			if __obf_fde97cbb580f6f76 == '\uFFFD' {
				__obf_428c3b4a95725c4a = __obf_f5dbf537a0450c02(__obf_428c3b4a95725c4a, __obf_9aca26302d2c5887)
				__obf_428c3b4a95725c4a = __obf_f5dbf537a0450c02(__obf_428c3b4a95725c4a, __obf_12837937a8e321f5)
			} else {
				__obf_428c3b4a95725c4a = __obf_f5dbf537a0450c02(__obf_428c3b4a95725c4a, __obf_fde97cbb580f6f76)
			}
		} else {
			__obf_428c3b4a95725c4a = __obf_f5dbf537a0450c02(__obf_428c3b4a95725c4a, __obf_9aca26302d2c5887)
		}
	case '"':
		__obf_428c3b4a95725c4a = append(__obf_428c3b4a95725c4a, '"')
	case '\\':
		__obf_428c3b4a95725c4a = append(__obf_428c3b4a95725c4a, '\\')
	case '/':
		__obf_428c3b4a95725c4a = append(__obf_428c3b4a95725c4a, '/')
	case 'b':
		__obf_428c3b4a95725c4a = append(__obf_428c3b4a95725c4a, '\b')
	case 'f':
		__obf_428c3b4a95725c4a = append(__obf_428c3b4a95725c4a, '\f')
	case 'n':
		__obf_428c3b4a95725c4a = append(__obf_428c3b4a95725c4a, '\n')
	case 'r':
		__obf_428c3b4a95725c4a = append(__obf_428c3b4a95725c4a, '\r')
	case 't':
		__obf_428c3b4a95725c4a = append(__obf_428c3b4a95725c4a, '\t')
	default:
		__obf_4ab56a99965952e7.
			ReportError("readEscapedChar",
				`invalid escape char after \`)
		return nil
	}
	return __obf_428c3b4a95725c4a
}

// ReadStringAsSlice read string from iterator without copying into string form.
// The []byte can not be kept, as it will change after next iterator call.
func (__obf_4ab56a99965952e7 *Iterator) ReadStringAsSlice() (__obf_59c72400ec1a6110 []byte) {
	__obf_24309b3b0ff9ba22 := __obf_4ab56a99965952e7.__obf_61df301d1f67ad73()
	if __obf_24309b3b0ff9ba22 == '"' {
		for __obf_82c6e05b9d226c58 := __obf_4ab56a99965952e7.__obf_5e22d6270d27491f;
		// require ascii string and no escape
		// for: field name, base64, number
		__obf_82c6e05b9d226c58 < __obf_4ab56a99965952e7.__obf_7c17042d9b4e73cc; __obf_82c6e05b9d226c58++ {

			if __obf_4ab56a99965952e7.__obf_0b1656d7ef4bd9df[__obf_82c6e05b9d226c58] == '"' {
				__obf_59c72400ec1a6110 = // fast path: reuse the underlying buffer
					__obf_4ab56a99965952e7.__obf_0b1656d7ef4bd9df[__obf_4ab56a99965952e7.__obf_5e22d6270d27491f:__obf_82c6e05b9d226c58]
				__obf_4ab56a99965952e7.__obf_5e22d6270d27491f = __obf_82c6e05b9d226c58 + 1
				return __obf_59c72400ec1a6110
			}
		}
		__obf_8324ba83052af51f := __obf_4ab56a99965952e7.__obf_7c17042d9b4e73cc - __obf_4ab56a99965952e7.__obf_5e22d6270d27491f
		__obf_390a561ab5257f8f := make([]byte, __obf_8324ba83052af51f, __obf_8324ba83052af51f*2)
		copy(__obf_390a561ab5257f8f, __obf_4ab56a99965952e7.__obf_0b1656d7ef4bd9df[__obf_4ab56a99965952e7.__obf_5e22d6270d27491f:__obf_4ab56a99965952e7.__obf_7c17042d9b4e73cc])
		__obf_4ab56a99965952e7.__obf_5e22d6270d27491f = __obf_4ab56a99965952e7.__obf_7c17042d9b4e73cc
		for __obf_4ab56a99965952e7.Error == nil {
			__obf_24309b3b0ff9ba22 := __obf_4ab56a99965952e7.__obf_2b6cf70e8dd32b68()
			if __obf_24309b3b0ff9ba22 == '"' {
				return __obf_390a561ab5257f8f
			}
			__obf_390a561ab5257f8f = append(__obf_390a561ab5257f8f, __obf_24309b3b0ff9ba22)
		}
		return __obf_390a561ab5257f8f
	}
	__obf_4ab56a99965952e7.
		ReportError("ReadStringAsSlice", `expects " or n, but found `+string([]byte{__obf_24309b3b0ff9ba22}))
	return
}

func (__obf_4ab56a99965952e7 *Iterator) __obf_4a65d6eac537eda8() (__obf_59c72400ec1a6110 rune) {
	for __obf_82c6e05b9d226c58 := 0; __obf_82c6e05b9d226c58 < 4; __obf_82c6e05b9d226c58++ {
		__obf_24309b3b0ff9ba22 := __obf_4ab56a99965952e7.__obf_2b6cf70e8dd32b68()
		if __obf_4ab56a99965952e7.Error != nil {
			return
		}
		if __obf_24309b3b0ff9ba22 >= '0' && __obf_24309b3b0ff9ba22 <= '9' {
			__obf_59c72400ec1a6110 = __obf_59c72400ec1a6110*16 + rune(__obf_24309b3b0ff9ba22-'0')
		} else if __obf_24309b3b0ff9ba22 >= 'a' && __obf_24309b3b0ff9ba22 <= 'f' {
			__obf_59c72400ec1a6110 = __obf_59c72400ec1a6110*16 + rune(__obf_24309b3b0ff9ba22-'a'+10)
		} else if __obf_24309b3b0ff9ba22 >= 'A' && __obf_24309b3b0ff9ba22 <= 'F' {
			__obf_59c72400ec1a6110 = __obf_59c72400ec1a6110*16 + rune(__obf_24309b3b0ff9ba22-'A'+10)
		} else {
			__obf_4ab56a99965952e7.
				ReportError("readU4", "expects 0~9 or a~f, but found "+string([]byte{__obf_24309b3b0ff9ba22}))
			return
		}
	}
	return __obf_59c72400ec1a6110
}

const (
	__obf_81f43c3c813b206f = 0x00
	__obf_0d0247a16a02da0b = // 0000 0000
	0x80
	__obf_018a11d730e8bc93 = // 1000 0000
	0xC0
	__obf_deecb1247ccb1b8a = // 1100 0000
	0xE0
	__obf_0d0f92458d170a7e = // 1110 0000
	0xF0
	__obf_a34e46079bd3b052 = // 1111 0000
	0xF8
	__obf_14f92024c8f90caa = // 1111 1000

	0x3F
	__obf_ea60cf3105384bb0 = // 0011 1111
	0x1F
	__obf_1bf25876f19a8f85 = // 0001 1111
	0x0F
	__obf_8591faf540b04df3 = // 0000 1111
	0x07
	__obf_5f0b81f102df35cd = // 0000 0111

	1<<7 - 1
	__obf_7ebc883351e09b33 = 1<<11 - 1
	__obf_1185685a6a537cc6 = 1<<16 - 1
	__obf_bef5b78babc69b6c = 0xD800
	__obf_ab55c6faf11600cb = 0xDFFF
	__obf_430bcade1f2409ef = '\U0010FFFF'
	__obf_8e8ee390f2f3d4e4 = // Maximum valid Unicode code point.
	'\uFFFD'                 // the "error" Rune or "Unicode replacement character"
)

func __obf_f5dbf537a0450c02(__obf_23ceea01a3d73f6a []byte, __obf_9aca26302d2c5887 rune) []byte {
	// Negative values are erroneous. Making it unsigned addresses the problem.
	switch __obf_82c6e05b9d226c58 := uint32(__obf_9aca26302d2c5887); {
	case __obf_82c6e05b9d226c58 <= __obf_5f0b81f102df35cd:
		__obf_23ceea01a3d73f6a = append(__obf_23ceea01a3d73f6a, byte(__obf_9aca26302d2c5887))
		return __obf_23ceea01a3d73f6a
	case __obf_82c6e05b9d226c58 <= __obf_7ebc883351e09b33:
		__obf_23ceea01a3d73f6a = append(__obf_23ceea01a3d73f6a, __obf_018a11d730e8bc93|byte(__obf_9aca26302d2c5887>>6))
		__obf_23ceea01a3d73f6a = append(__obf_23ceea01a3d73f6a, __obf_0d0247a16a02da0b|byte(__obf_9aca26302d2c5887)&__obf_14f92024c8f90caa)
		return __obf_23ceea01a3d73f6a
	case __obf_82c6e05b9d226c58 > __obf_430bcade1f2409ef, __obf_bef5b78babc69b6c <= __obf_82c6e05b9d226c58 && __obf_82c6e05b9d226c58 <= __obf_ab55c6faf11600cb:
		__obf_9aca26302d2c5887 = __obf_8e8ee390f2f3d4e4
		fallthrough
	case __obf_82c6e05b9d226c58 <= __obf_1185685a6a537cc6:
		__obf_23ceea01a3d73f6a = append(__obf_23ceea01a3d73f6a, __obf_deecb1247ccb1b8a|byte(__obf_9aca26302d2c5887>>12))
		__obf_23ceea01a3d73f6a = append(__obf_23ceea01a3d73f6a, __obf_0d0247a16a02da0b|byte(__obf_9aca26302d2c5887>>6)&__obf_14f92024c8f90caa)
		__obf_23ceea01a3d73f6a = append(__obf_23ceea01a3d73f6a, __obf_0d0247a16a02da0b|byte(__obf_9aca26302d2c5887)&__obf_14f92024c8f90caa)
		return __obf_23ceea01a3d73f6a
	default:
		__obf_23ceea01a3d73f6a = append(__obf_23ceea01a3d73f6a, __obf_0d0f92458d170a7e|byte(__obf_9aca26302d2c5887>>18))
		__obf_23ceea01a3d73f6a = append(__obf_23ceea01a3d73f6a, __obf_0d0247a16a02da0b|byte(__obf_9aca26302d2c5887>>12)&__obf_14f92024c8f90caa)
		__obf_23ceea01a3d73f6a = append(__obf_23ceea01a3d73f6a, __obf_0d0247a16a02da0b|byte(__obf_9aca26302d2c5887>>6)&__obf_14f92024c8f90caa)
		__obf_23ceea01a3d73f6a = append(__obf_23ceea01a3d73f6a, __obf_0d0247a16a02da0b|byte(__obf_9aca26302d2c5887)&__obf_14f92024c8f90caa)
		return __obf_23ceea01a3d73f6a
	}
}
