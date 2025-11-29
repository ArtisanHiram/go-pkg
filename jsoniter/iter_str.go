package __obf_91620b895eeff9ed

import (
	"fmt"
	"unicode/utf16"
)

// ReadString read string from iterator
func (__obf_1bb30e8a74ed8233 *Iterator) ReadString() (__obf_e46f5fe3db5036fe string) {
	__obf_f16b4157911bc9af := __obf_1bb30e8a74ed8233.__obf_684faa48ae8c5049()
	if __obf_f16b4157911bc9af == '"' {
		for __obf_5aa5c8829b97f182 := __obf_1bb30e8a74ed8233.__obf_a657fb48fcb34e21; __obf_5aa5c8829b97f182 < __obf_1bb30e8a74ed8233.__obf_15d837671d2809ae; __obf_5aa5c8829b97f182++ {
			__obf_f16b4157911bc9af := __obf_1bb30e8a74ed8233.__obf_184433571fa55237[__obf_5aa5c8829b97f182]
			if __obf_f16b4157911bc9af == '"' {
				__obf_e46f5fe3db5036fe = string(__obf_1bb30e8a74ed8233.__obf_184433571fa55237[__obf_1bb30e8a74ed8233.__obf_a657fb48fcb34e21:__obf_5aa5c8829b97f182])
				__obf_1bb30e8a74ed8233.__obf_a657fb48fcb34e21 = __obf_5aa5c8829b97f182 + 1
				return __obf_e46f5fe3db5036fe
			} else if __obf_f16b4157911bc9af == '\\' {
				break
			} else if __obf_f16b4157911bc9af < ' ' {
				__obf_1bb30e8a74ed8233.
					ReportError("ReadString",
						fmt.Sprintf(`invalid control character found: %d`, __obf_f16b4157911bc9af))
				return
			}
		}
		return __obf_1bb30e8a74ed8233.__obf_e58c02da1e6cd884()
	} else if __obf_f16b4157911bc9af == 'n' {
		__obf_1bb30e8a74ed8233.__obf_3e1d2ad9a54f0d22('u', 'l', 'l')
		return ""
	}
	__obf_1bb30e8a74ed8233.
		ReportError("ReadString", `expects " or n, but found `+string([]byte{__obf_f16b4157911bc9af}))
	return
}

func (__obf_1bb30e8a74ed8233 *Iterator) __obf_e58c02da1e6cd884() (__obf_e46f5fe3db5036fe string) {
	var __obf_e91bd2feb751e4f1 []byte
	var __obf_f16b4157911bc9af byte
	for __obf_1bb30e8a74ed8233.Error == nil {
		__obf_f16b4157911bc9af = __obf_1bb30e8a74ed8233.__obf_9617ab9cc89bcddc()
		if __obf_f16b4157911bc9af == '"' {
			return string(__obf_e91bd2feb751e4f1)
		}
		if __obf_f16b4157911bc9af == '\\' {
			__obf_f16b4157911bc9af = __obf_1bb30e8a74ed8233.__obf_9617ab9cc89bcddc()
			__obf_e91bd2feb751e4f1 = __obf_1bb30e8a74ed8233.__obf_315d3526429ed7ec(__obf_f16b4157911bc9af, __obf_e91bd2feb751e4f1)
		} else {
			__obf_e91bd2feb751e4f1 = append(__obf_e91bd2feb751e4f1, __obf_f16b4157911bc9af)
		}
	}
	__obf_1bb30e8a74ed8233.
		ReportError("readStringSlowPath", "unexpected end of input")
	return
}

func (__obf_1bb30e8a74ed8233 *Iterator) __obf_315d3526429ed7ec(__obf_f16b4157911bc9af byte, __obf_e91bd2feb751e4f1 []byte) []byte {
	switch __obf_f16b4157911bc9af {
	case 'u':
		__obf_44946063bfe54da7 := __obf_1bb30e8a74ed8233.__obf_20f478da6e89426a()
		if utf16.IsSurrogate(__obf_44946063bfe54da7) {
			__obf_f16b4157911bc9af = __obf_1bb30e8a74ed8233.__obf_9617ab9cc89bcddc()
			if __obf_1bb30e8a74ed8233.Error != nil {
				return nil
			}
			if __obf_f16b4157911bc9af != '\\' {
				__obf_1bb30e8a74ed8233.__obf_a163df67f9bb1c4b()
				__obf_e91bd2feb751e4f1 = __obf_6c9d1b470209366b(__obf_e91bd2feb751e4f1, __obf_44946063bfe54da7)
				return __obf_e91bd2feb751e4f1
			}
			__obf_f16b4157911bc9af = __obf_1bb30e8a74ed8233.__obf_9617ab9cc89bcddc()
			if __obf_1bb30e8a74ed8233.Error != nil {
				return nil
			}
			if __obf_f16b4157911bc9af != 'u' {
				__obf_e91bd2feb751e4f1 = __obf_6c9d1b470209366b(__obf_e91bd2feb751e4f1, __obf_44946063bfe54da7)
				return __obf_1bb30e8a74ed8233.__obf_315d3526429ed7ec(__obf_f16b4157911bc9af, __obf_e91bd2feb751e4f1)
			}
			__obf_c3f7781dafa5b032 := __obf_1bb30e8a74ed8233.__obf_20f478da6e89426a()
			if __obf_1bb30e8a74ed8233.Error != nil {
				return nil
			}
			__obf_61d01ce508840a45 := utf16.DecodeRune(__obf_44946063bfe54da7, __obf_c3f7781dafa5b032)
			if __obf_61d01ce508840a45 == '\uFFFD' {
				__obf_e91bd2feb751e4f1 = __obf_6c9d1b470209366b(__obf_e91bd2feb751e4f1, __obf_44946063bfe54da7)
				__obf_e91bd2feb751e4f1 = __obf_6c9d1b470209366b(__obf_e91bd2feb751e4f1, __obf_c3f7781dafa5b032)
			} else {
				__obf_e91bd2feb751e4f1 = __obf_6c9d1b470209366b(__obf_e91bd2feb751e4f1, __obf_61d01ce508840a45)
			}
		} else {
			__obf_e91bd2feb751e4f1 = __obf_6c9d1b470209366b(__obf_e91bd2feb751e4f1, __obf_44946063bfe54da7)
		}
	case '"':
		__obf_e91bd2feb751e4f1 = append(__obf_e91bd2feb751e4f1, '"')
	case '\\':
		__obf_e91bd2feb751e4f1 = append(__obf_e91bd2feb751e4f1, '\\')
	case '/':
		__obf_e91bd2feb751e4f1 = append(__obf_e91bd2feb751e4f1, '/')
	case 'b':
		__obf_e91bd2feb751e4f1 = append(__obf_e91bd2feb751e4f1, '\b')
	case 'f':
		__obf_e91bd2feb751e4f1 = append(__obf_e91bd2feb751e4f1, '\f')
	case 'n':
		__obf_e91bd2feb751e4f1 = append(__obf_e91bd2feb751e4f1, '\n')
	case 'r':
		__obf_e91bd2feb751e4f1 = append(__obf_e91bd2feb751e4f1, '\r')
	case 't':
		__obf_e91bd2feb751e4f1 = append(__obf_e91bd2feb751e4f1, '\t')
	default:
		__obf_1bb30e8a74ed8233.
			ReportError("readEscapedChar",
				`invalid escape char after \`)
		return nil
	}
	return __obf_e91bd2feb751e4f1
}

// ReadStringAsSlice read string from iterator without copying into string form.
// The []byte can not be kept, as it will change after next iterator call.
func (__obf_1bb30e8a74ed8233 *Iterator) ReadStringAsSlice() (__obf_e46f5fe3db5036fe []byte) {
	__obf_f16b4157911bc9af := __obf_1bb30e8a74ed8233.__obf_684faa48ae8c5049()
	if __obf_f16b4157911bc9af == '"' {
		for __obf_5aa5c8829b97f182 := __obf_1bb30e8a74ed8233.__obf_a657fb48fcb34e21;
		// require ascii string and no escape
		// for: field name, base64, number
		__obf_5aa5c8829b97f182 < __obf_1bb30e8a74ed8233.__obf_15d837671d2809ae; __obf_5aa5c8829b97f182++ {

			if __obf_1bb30e8a74ed8233.__obf_184433571fa55237[__obf_5aa5c8829b97f182] == '"' {
				__obf_e46f5fe3db5036fe = // fast path: reuse the underlying buffer
					__obf_1bb30e8a74ed8233.__obf_184433571fa55237[__obf_1bb30e8a74ed8233.__obf_a657fb48fcb34e21:__obf_5aa5c8829b97f182]
				__obf_1bb30e8a74ed8233.__obf_a657fb48fcb34e21 = __obf_5aa5c8829b97f182 + 1
				return __obf_e46f5fe3db5036fe
			}
		}
		__obf_d81c22bc85ef5db2 := __obf_1bb30e8a74ed8233.__obf_15d837671d2809ae - __obf_1bb30e8a74ed8233.__obf_a657fb48fcb34e21
		__obf_fce0adcefc66d2ad := make([]byte, __obf_d81c22bc85ef5db2, __obf_d81c22bc85ef5db2*2)
		copy(__obf_fce0adcefc66d2ad, __obf_1bb30e8a74ed8233.__obf_184433571fa55237[__obf_1bb30e8a74ed8233.__obf_a657fb48fcb34e21:__obf_1bb30e8a74ed8233.__obf_15d837671d2809ae])
		__obf_1bb30e8a74ed8233.__obf_a657fb48fcb34e21 = __obf_1bb30e8a74ed8233.__obf_15d837671d2809ae
		for __obf_1bb30e8a74ed8233.Error == nil {
			__obf_f16b4157911bc9af := __obf_1bb30e8a74ed8233.__obf_9617ab9cc89bcddc()
			if __obf_f16b4157911bc9af == '"' {
				return __obf_fce0adcefc66d2ad
			}
			__obf_fce0adcefc66d2ad = append(__obf_fce0adcefc66d2ad, __obf_f16b4157911bc9af)
		}
		return __obf_fce0adcefc66d2ad
	}
	__obf_1bb30e8a74ed8233.
		ReportError("ReadStringAsSlice", `expects " or n, but found `+string([]byte{__obf_f16b4157911bc9af}))
	return
}

func (__obf_1bb30e8a74ed8233 *Iterator) __obf_20f478da6e89426a() (__obf_e46f5fe3db5036fe rune) {
	for __obf_5aa5c8829b97f182 := 0; __obf_5aa5c8829b97f182 < 4; __obf_5aa5c8829b97f182++ {
		__obf_f16b4157911bc9af := __obf_1bb30e8a74ed8233.__obf_9617ab9cc89bcddc()
		if __obf_1bb30e8a74ed8233.Error != nil {
			return
		}
		if __obf_f16b4157911bc9af >= '0' && __obf_f16b4157911bc9af <= '9' {
			__obf_e46f5fe3db5036fe = __obf_e46f5fe3db5036fe*16 + rune(__obf_f16b4157911bc9af-'0')
		} else if __obf_f16b4157911bc9af >= 'a' && __obf_f16b4157911bc9af <= 'f' {
			__obf_e46f5fe3db5036fe = __obf_e46f5fe3db5036fe*16 + rune(__obf_f16b4157911bc9af-'a'+10)
		} else if __obf_f16b4157911bc9af >= 'A' && __obf_f16b4157911bc9af <= 'F' {
			__obf_e46f5fe3db5036fe = __obf_e46f5fe3db5036fe*16 + rune(__obf_f16b4157911bc9af-'A'+10)
		} else {
			__obf_1bb30e8a74ed8233.
				ReportError("readU4", "expects 0~9 or a~f, but found "+string([]byte{__obf_f16b4157911bc9af}))
			return
		}
	}
	return __obf_e46f5fe3db5036fe
}

const (
	__obf_468c7c882a4b394d = 0x00
	__obf_369b5d087d80c67c = // 0000 0000
	0x80
	__obf_0618a11c70728488 = // 1000 0000
	0xC0
	__obf_c55bfaf16abf690d = // 1100 0000
	0xE0
	__obf_5ecb4cea72bf1f6e = // 1110 0000
	0xF0
	__obf_5f57b26dbecbd7a5 = // 1111 0000
	0xF8
	__obf_0ee71d2d658a37a4 = // 1111 1000

	0x3F
	__obf_e4623dea85b93a04 = // 0011 1111
	0x1F
	__obf_6af99dddee2cb3e8 = // 0001 1111
	0x0F
	__obf_a6e39fb6f5f628ca = // 0000 1111
	0x07
	__obf_fe5b632ca679df07 = // 0000 0111

	1<<7 - 1
	__obf_7a75d4c67a2d4cd1 = 1<<11 - 1
	__obf_3919a204519f7baf = 1<<16 - 1
	__obf_f4c37afdb1ea11c8 = 0xD800
	__obf_09046f928cd55921 = 0xDFFF
	__obf_7219879dd9af962a = '\U0010FFFF'
	__obf_9b1a6d61184cefd7 = // Maximum valid Unicode code point.
	'\uFFFD'                 // the "error" Rune or "Unicode replacement character"
)

func __obf_6c9d1b470209366b(__obf_bcd7bec668625a79 []byte, __obf_44946063bfe54da7 rune) []byte {
	// Negative values are erroneous. Making it unsigned addresses the problem.
	switch __obf_5aa5c8829b97f182 := uint32(__obf_44946063bfe54da7); {
	case __obf_5aa5c8829b97f182 <= __obf_fe5b632ca679df07:
		__obf_bcd7bec668625a79 = append(__obf_bcd7bec668625a79, byte(__obf_44946063bfe54da7))
		return __obf_bcd7bec668625a79
	case __obf_5aa5c8829b97f182 <= __obf_7a75d4c67a2d4cd1:
		__obf_bcd7bec668625a79 = append(__obf_bcd7bec668625a79, __obf_0618a11c70728488|byte(__obf_44946063bfe54da7>>6))
		__obf_bcd7bec668625a79 = append(__obf_bcd7bec668625a79, __obf_369b5d087d80c67c|byte(__obf_44946063bfe54da7)&__obf_0ee71d2d658a37a4)
		return __obf_bcd7bec668625a79
	case __obf_5aa5c8829b97f182 > __obf_7219879dd9af962a, __obf_f4c37afdb1ea11c8 <= __obf_5aa5c8829b97f182 && __obf_5aa5c8829b97f182 <= __obf_09046f928cd55921:
		__obf_44946063bfe54da7 = __obf_9b1a6d61184cefd7
		fallthrough
	case __obf_5aa5c8829b97f182 <= __obf_3919a204519f7baf:
		__obf_bcd7bec668625a79 = append(__obf_bcd7bec668625a79, __obf_c55bfaf16abf690d|byte(__obf_44946063bfe54da7>>12))
		__obf_bcd7bec668625a79 = append(__obf_bcd7bec668625a79, __obf_369b5d087d80c67c|byte(__obf_44946063bfe54da7>>6)&__obf_0ee71d2d658a37a4)
		__obf_bcd7bec668625a79 = append(__obf_bcd7bec668625a79, __obf_369b5d087d80c67c|byte(__obf_44946063bfe54da7)&__obf_0ee71d2d658a37a4)
		return __obf_bcd7bec668625a79
	default:
		__obf_bcd7bec668625a79 = append(__obf_bcd7bec668625a79, __obf_5ecb4cea72bf1f6e|byte(__obf_44946063bfe54da7>>18))
		__obf_bcd7bec668625a79 = append(__obf_bcd7bec668625a79, __obf_369b5d087d80c67c|byte(__obf_44946063bfe54da7>>12)&__obf_0ee71d2d658a37a4)
		__obf_bcd7bec668625a79 = append(__obf_bcd7bec668625a79, __obf_369b5d087d80c67c|byte(__obf_44946063bfe54da7>>6)&__obf_0ee71d2d658a37a4)
		__obf_bcd7bec668625a79 = append(__obf_bcd7bec668625a79, __obf_369b5d087d80c67c|byte(__obf_44946063bfe54da7)&__obf_0ee71d2d658a37a4)
		return __obf_bcd7bec668625a79
	}
}
