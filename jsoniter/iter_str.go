package __obf_c3cd04a15c56f32f

import (
	"fmt"
	"unicode/utf16"
)

// ReadString read string from iterator
func (__obf_edd9fe48d39445e4 *Iterator) ReadString() (__obf_316af79472661247 string) {
	__obf_0c1bc1e511a43120 := __obf_edd9fe48d39445e4.__obf_a26eeeac1d6f5a11()
	if __obf_0c1bc1e511a43120 == '"' {
		for __obf_28d099df85f083a8 := __obf_edd9fe48d39445e4.__obf_edd3c3885447229b; __obf_28d099df85f083a8 < __obf_edd9fe48d39445e4.__obf_3a475d1c1ae2cd57; __obf_28d099df85f083a8++ {
			__obf_0c1bc1e511a43120 := __obf_edd9fe48d39445e4.__obf_ace979f6622823f3[__obf_28d099df85f083a8]
			if __obf_0c1bc1e511a43120 == '"' {
				__obf_316af79472661247 = string(__obf_edd9fe48d39445e4.__obf_ace979f6622823f3[__obf_edd9fe48d39445e4.__obf_edd3c3885447229b:__obf_28d099df85f083a8])
				__obf_edd9fe48d39445e4.__obf_edd3c3885447229b = __obf_28d099df85f083a8 + 1
				return __obf_316af79472661247
			} else if __obf_0c1bc1e511a43120 == '\\' {
				break
			} else if __obf_0c1bc1e511a43120 < ' ' {
				__obf_edd9fe48d39445e4.
					ReportError("ReadString",
						fmt.Sprintf(`invalid control character found: %d`, __obf_0c1bc1e511a43120))
				return
			}
		}
		return __obf_edd9fe48d39445e4.__obf_561396e75273c9da()
	} else if __obf_0c1bc1e511a43120 == 'n' {
		__obf_edd9fe48d39445e4.__obf_8bc1f4b8d62f5247('u', 'l', 'l')
		return ""
	}
	__obf_edd9fe48d39445e4.
		ReportError("ReadString", `expects " or n, but found `+string([]byte{__obf_0c1bc1e511a43120}))
	return
}

func (__obf_edd9fe48d39445e4 *Iterator) __obf_561396e75273c9da() (__obf_316af79472661247 string) {
	var __obf_2d944450d48e5810 []byte
	var __obf_0c1bc1e511a43120 byte
	for __obf_edd9fe48d39445e4.Error == nil {
		__obf_0c1bc1e511a43120 = __obf_edd9fe48d39445e4.__obf_70c673c91de4f883()
		if __obf_0c1bc1e511a43120 == '"' {
			return string(__obf_2d944450d48e5810)
		}
		if __obf_0c1bc1e511a43120 == '\\' {
			__obf_0c1bc1e511a43120 = __obf_edd9fe48d39445e4.__obf_70c673c91de4f883()
			__obf_2d944450d48e5810 = __obf_edd9fe48d39445e4.__obf_641c7defdc4ce200(__obf_0c1bc1e511a43120, __obf_2d944450d48e5810)
		} else {
			__obf_2d944450d48e5810 = append(__obf_2d944450d48e5810, __obf_0c1bc1e511a43120)
		}
	}
	__obf_edd9fe48d39445e4.
		ReportError("readStringSlowPath", "unexpected end of input")
	return
}

func (__obf_edd9fe48d39445e4 *Iterator) __obf_641c7defdc4ce200(__obf_0c1bc1e511a43120 byte, __obf_2d944450d48e5810 []byte) []byte {
	switch __obf_0c1bc1e511a43120 {
	case 'u':
		__obf_c9c2210cc516a267 := __obf_edd9fe48d39445e4.__obf_573176d14d38234c()
		if utf16.IsSurrogate(__obf_c9c2210cc516a267) {
			__obf_0c1bc1e511a43120 = __obf_edd9fe48d39445e4.__obf_70c673c91de4f883()
			if __obf_edd9fe48d39445e4.Error != nil {
				return nil
			}
			if __obf_0c1bc1e511a43120 != '\\' {
				__obf_edd9fe48d39445e4.__obf_a0622aded2d86ded()
				__obf_2d944450d48e5810 = __obf_4dc4e0bee3d6e5b0(__obf_2d944450d48e5810, __obf_c9c2210cc516a267)
				return __obf_2d944450d48e5810
			}
			__obf_0c1bc1e511a43120 = __obf_edd9fe48d39445e4.__obf_70c673c91de4f883()
			if __obf_edd9fe48d39445e4.Error != nil {
				return nil
			}
			if __obf_0c1bc1e511a43120 != 'u' {
				__obf_2d944450d48e5810 = __obf_4dc4e0bee3d6e5b0(__obf_2d944450d48e5810, __obf_c9c2210cc516a267)
				return __obf_edd9fe48d39445e4.__obf_641c7defdc4ce200(__obf_0c1bc1e511a43120, __obf_2d944450d48e5810)
			}
			__obf_695eff5f2dce5e40 := __obf_edd9fe48d39445e4.__obf_573176d14d38234c()
			if __obf_edd9fe48d39445e4.Error != nil {
				return nil
			}
			__obf_510a540cf25c88d2 := utf16.DecodeRune(__obf_c9c2210cc516a267, __obf_695eff5f2dce5e40)
			if __obf_510a540cf25c88d2 == '\uFFFD' {
				__obf_2d944450d48e5810 = __obf_4dc4e0bee3d6e5b0(__obf_2d944450d48e5810, __obf_c9c2210cc516a267)
				__obf_2d944450d48e5810 = __obf_4dc4e0bee3d6e5b0(__obf_2d944450d48e5810, __obf_695eff5f2dce5e40)
			} else {
				__obf_2d944450d48e5810 = __obf_4dc4e0bee3d6e5b0(__obf_2d944450d48e5810, __obf_510a540cf25c88d2)
			}
		} else {
			__obf_2d944450d48e5810 = __obf_4dc4e0bee3d6e5b0(__obf_2d944450d48e5810, __obf_c9c2210cc516a267)
		}
	case '"':
		__obf_2d944450d48e5810 = append(__obf_2d944450d48e5810, '"')
	case '\\':
		__obf_2d944450d48e5810 = append(__obf_2d944450d48e5810, '\\')
	case '/':
		__obf_2d944450d48e5810 = append(__obf_2d944450d48e5810, '/')
	case 'b':
		__obf_2d944450d48e5810 = append(__obf_2d944450d48e5810, '\b')
	case 'f':
		__obf_2d944450d48e5810 = append(__obf_2d944450d48e5810, '\f')
	case 'n':
		__obf_2d944450d48e5810 = append(__obf_2d944450d48e5810, '\n')
	case 'r':
		__obf_2d944450d48e5810 = append(__obf_2d944450d48e5810, '\r')
	case 't':
		__obf_2d944450d48e5810 = append(__obf_2d944450d48e5810, '\t')
	default:
		__obf_edd9fe48d39445e4.
			ReportError("readEscapedChar",
				`invalid escape char after \`)
		return nil
	}
	return __obf_2d944450d48e5810
}

// ReadStringAsSlice read string from iterator without copying into string form.
// The []byte can not be kept, as it will change after next iterator call.
func (__obf_edd9fe48d39445e4 *Iterator) ReadStringAsSlice() (__obf_316af79472661247 []byte) {
	__obf_0c1bc1e511a43120 := __obf_edd9fe48d39445e4.__obf_a26eeeac1d6f5a11()
	if __obf_0c1bc1e511a43120 == '"' {
		for __obf_28d099df85f083a8 := __obf_edd9fe48d39445e4.__obf_edd3c3885447229b;
		// require ascii string and no escape
		// for: field name, base64, number
		__obf_28d099df85f083a8 < __obf_edd9fe48d39445e4.__obf_3a475d1c1ae2cd57; __obf_28d099df85f083a8++ {

			if __obf_edd9fe48d39445e4.__obf_ace979f6622823f3[__obf_28d099df85f083a8] == '"' {
				__obf_316af79472661247 = // fast path: reuse the underlying buffer
					__obf_edd9fe48d39445e4.__obf_ace979f6622823f3[__obf_edd9fe48d39445e4.__obf_edd3c3885447229b:__obf_28d099df85f083a8]
				__obf_edd9fe48d39445e4.__obf_edd3c3885447229b = __obf_28d099df85f083a8 + 1
				return __obf_316af79472661247
			}
		}
		__obf_50f30ca9f423e440 := __obf_edd9fe48d39445e4.__obf_3a475d1c1ae2cd57 - __obf_edd9fe48d39445e4.__obf_edd3c3885447229b
		__obf_4cb2b46ddea01d73 := make([]byte, __obf_50f30ca9f423e440, __obf_50f30ca9f423e440*2)
		copy(__obf_4cb2b46ddea01d73, __obf_edd9fe48d39445e4.__obf_ace979f6622823f3[__obf_edd9fe48d39445e4.__obf_edd3c3885447229b:__obf_edd9fe48d39445e4.__obf_3a475d1c1ae2cd57])
		__obf_edd9fe48d39445e4.__obf_edd3c3885447229b = __obf_edd9fe48d39445e4.__obf_3a475d1c1ae2cd57
		for __obf_edd9fe48d39445e4.Error == nil {
			__obf_0c1bc1e511a43120 := __obf_edd9fe48d39445e4.__obf_70c673c91de4f883()
			if __obf_0c1bc1e511a43120 == '"' {
				return __obf_4cb2b46ddea01d73
			}
			__obf_4cb2b46ddea01d73 = append(__obf_4cb2b46ddea01d73, __obf_0c1bc1e511a43120)
		}
		return __obf_4cb2b46ddea01d73
	}
	__obf_edd9fe48d39445e4.
		ReportError("ReadStringAsSlice", `expects " or n, but found `+string([]byte{__obf_0c1bc1e511a43120}))
	return
}

func (__obf_edd9fe48d39445e4 *Iterator) __obf_573176d14d38234c() (__obf_316af79472661247 rune) {
	for __obf_28d099df85f083a8 := 0; __obf_28d099df85f083a8 < 4; __obf_28d099df85f083a8++ {
		__obf_0c1bc1e511a43120 := __obf_edd9fe48d39445e4.__obf_70c673c91de4f883()
		if __obf_edd9fe48d39445e4.Error != nil {
			return
		}
		if __obf_0c1bc1e511a43120 >= '0' && __obf_0c1bc1e511a43120 <= '9' {
			__obf_316af79472661247 = __obf_316af79472661247*16 + rune(__obf_0c1bc1e511a43120-'0')
		} else if __obf_0c1bc1e511a43120 >= 'a' && __obf_0c1bc1e511a43120 <= 'f' {
			__obf_316af79472661247 = __obf_316af79472661247*16 + rune(__obf_0c1bc1e511a43120-'a'+10)
		} else if __obf_0c1bc1e511a43120 >= 'A' && __obf_0c1bc1e511a43120 <= 'F' {
			__obf_316af79472661247 = __obf_316af79472661247*16 + rune(__obf_0c1bc1e511a43120-'A'+10)
		} else {
			__obf_edd9fe48d39445e4.
				ReportError("readU4", "expects 0~9 or a~f, but found "+string([]byte{__obf_0c1bc1e511a43120}))
			return
		}
	}
	return __obf_316af79472661247
}

const (
	__obf_bd5846cb5f7add36 = 0x00
	__obf_972beef8d8fbc858 = // 0000 0000
	0x80
	__obf_cde254558e35efdc = // 1000 0000
	0xC0
	__obf_a50bbdfcdef483f4 = // 1100 0000
	0xE0
	__obf_18c257a2fe547b82 = // 1110 0000
	0xF0
	__obf_12c994522e3f9f82 = // 1111 0000
	0xF8
	__obf_807c4555c7c9e94a = // 1111 1000

	0x3F
	__obf_bc5cb31fa0747c02 = // 0011 1111
	0x1F
	__obf_ca489d32ada69b6d = // 0001 1111
	0x0F
	__obf_3ee3d96f92082fb2 = // 0000 1111
	0x07
	__obf_aacaf6d0380b657a = // 0000 0111

	1<<7 - 1
	__obf_d5512c623921f247 = 1<<11 - 1
	__obf_892d2387149d6a03 = 1<<16 - 1
	__obf_d8b46204a6034e44 = 0xD800
	__obf_b108d5a9dfa66fd2 = 0xDFFF
	__obf_d4b9599da3bb84b2 = '\U0010FFFF'
	__obf_7dfab770381dc16a = // Maximum valid Unicode code point.
	'\uFFFD'                 // the "error" Rune or "Unicode replacement character"
)

func __obf_4dc4e0bee3d6e5b0(__obf_420b1bf30275e403 []byte, __obf_c9c2210cc516a267 rune) []byte {
	// Negative values are erroneous. Making it unsigned addresses the problem.
	switch __obf_28d099df85f083a8 := uint32(__obf_c9c2210cc516a267); {
	case __obf_28d099df85f083a8 <= __obf_aacaf6d0380b657a:
		__obf_420b1bf30275e403 = append(__obf_420b1bf30275e403, byte(__obf_c9c2210cc516a267))
		return __obf_420b1bf30275e403
	case __obf_28d099df85f083a8 <= __obf_d5512c623921f247:
		__obf_420b1bf30275e403 = append(__obf_420b1bf30275e403, __obf_cde254558e35efdc|byte(__obf_c9c2210cc516a267>>6))
		__obf_420b1bf30275e403 = append(__obf_420b1bf30275e403, __obf_972beef8d8fbc858|byte(__obf_c9c2210cc516a267)&__obf_807c4555c7c9e94a)
		return __obf_420b1bf30275e403
	case __obf_28d099df85f083a8 > __obf_d4b9599da3bb84b2, __obf_d8b46204a6034e44 <= __obf_28d099df85f083a8 && __obf_28d099df85f083a8 <= __obf_b108d5a9dfa66fd2:
		__obf_c9c2210cc516a267 = __obf_7dfab770381dc16a
		fallthrough
	case __obf_28d099df85f083a8 <= __obf_892d2387149d6a03:
		__obf_420b1bf30275e403 = append(__obf_420b1bf30275e403, __obf_a50bbdfcdef483f4|byte(__obf_c9c2210cc516a267>>12))
		__obf_420b1bf30275e403 = append(__obf_420b1bf30275e403, __obf_972beef8d8fbc858|byte(__obf_c9c2210cc516a267>>6)&__obf_807c4555c7c9e94a)
		__obf_420b1bf30275e403 = append(__obf_420b1bf30275e403, __obf_972beef8d8fbc858|byte(__obf_c9c2210cc516a267)&__obf_807c4555c7c9e94a)
		return __obf_420b1bf30275e403
	default:
		__obf_420b1bf30275e403 = append(__obf_420b1bf30275e403, __obf_18c257a2fe547b82|byte(__obf_c9c2210cc516a267>>18))
		__obf_420b1bf30275e403 = append(__obf_420b1bf30275e403, __obf_972beef8d8fbc858|byte(__obf_c9c2210cc516a267>>12)&__obf_807c4555c7c9e94a)
		__obf_420b1bf30275e403 = append(__obf_420b1bf30275e403, __obf_972beef8d8fbc858|byte(__obf_c9c2210cc516a267>>6)&__obf_807c4555c7c9e94a)
		__obf_420b1bf30275e403 = append(__obf_420b1bf30275e403, __obf_972beef8d8fbc858|byte(__obf_c9c2210cc516a267)&__obf_807c4555c7c9e94a)
		return __obf_420b1bf30275e403
	}
}
