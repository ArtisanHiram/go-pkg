package __obf_5b802ce8d9ba56d6

import (
	"fmt"
	"unicode/utf16"
)

// ReadString read string from iterator
func (__obf_67008a6a9e5ba828 *Iterator) ReadString() (__obf_5dabcdfee5097ed6 string) {
	__obf_dab9baaadfa7c8c2 := __obf_67008a6a9e5ba828.__obf_b781a59d5a0d2490()
	if __obf_dab9baaadfa7c8c2 == '"' {
		for __obf_2deec7c38ffb6ae3 := __obf_67008a6a9e5ba828.__obf_14babd6f9a55bd36; __obf_2deec7c38ffb6ae3 < __obf_67008a6a9e5ba828.__obf_3a36550914545c79; __obf_2deec7c38ffb6ae3++ {
			__obf_dab9baaadfa7c8c2 := __obf_67008a6a9e5ba828.__obf_9fc06d9180f0daca[__obf_2deec7c38ffb6ae3]
			if __obf_dab9baaadfa7c8c2 == '"' {
				__obf_5dabcdfee5097ed6 = string(__obf_67008a6a9e5ba828.__obf_9fc06d9180f0daca[__obf_67008a6a9e5ba828.__obf_14babd6f9a55bd36:__obf_2deec7c38ffb6ae3])
				__obf_67008a6a9e5ba828.__obf_14babd6f9a55bd36 = __obf_2deec7c38ffb6ae3 + 1
				return __obf_5dabcdfee5097ed6
			} else if __obf_dab9baaadfa7c8c2 == '\\' {
				break
			} else if __obf_dab9baaadfa7c8c2 < ' ' {
				__obf_67008a6a9e5ba828.
					ReportError("ReadString",
						fmt.Sprintf(`invalid control character found: %d`, __obf_dab9baaadfa7c8c2))
				return
			}
		}
		return __obf_67008a6a9e5ba828.__obf_ddc6a694dea48ef4()
	} else if __obf_dab9baaadfa7c8c2 == 'n' {
		__obf_67008a6a9e5ba828.__obf_4aeb767e0be7277a('u', 'l', 'l')
		return ""
	}
	__obf_67008a6a9e5ba828.
		ReportError("ReadString", `expects " or n, but found `+string([]byte{__obf_dab9baaadfa7c8c2}))
	return
}

func (__obf_67008a6a9e5ba828 *Iterator) __obf_ddc6a694dea48ef4() (__obf_5dabcdfee5097ed6 string) {
	var __obf_12c21b79fa86dcba []byte
	var __obf_dab9baaadfa7c8c2 byte
	for __obf_67008a6a9e5ba828.Error == nil {
		__obf_dab9baaadfa7c8c2 = __obf_67008a6a9e5ba828.__obf_ea3ebd5c6789bccb()
		if __obf_dab9baaadfa7c8c2 == '"' {
			return string(__obf_12c21b79fa86dcba)
		}
		if __obf_dab9baaadfa7c8c2 == '\\' {
			__obf_dab9baaadfa7c8c2 = __obf_67008a6a9e5ba828.__obf_ea3ebd5c6789bccb()
			__obf_12c21b79fa86dcba = __obf_67008a6a9e5ba828.__obf_57340dc7a0cce10d(__obf_dab9baaadfa7c8c2, __obf_12c21b79fa86dcba)
		} else {
			__obf_12c21b79fa86dcba = append(__obf_12c21b79fa86dcba, __obf_dab9baaadfa7c8c2)
		}
	}
	__obf_67008a6a9e5ba828.
		ReportError("readStringSlowPath", "unexpected end of input")
	return
}

func (__obf_67008a6a9e5ba828 *Iterator) __obf_57340dc7a0cce10d(__obf_dab9baaadfa7c8c2 byte, __obf_12c21b79fa86dcba []byte) []byte {
	switch __obf_dab9baaadfa7c8c2 {
	case 'u':
		__obf_14c0fc639db9b91c := __obf_67008a6a9e5ba828.__obf_5d1e2ea19783b872()
		if utf16.IsSurrogate(__obf_14c0fc639db9b91c) {
			__obf_dab9baaadfa7c8c2 = __obf_67008a6a9e5ba828.__obf_ea3ebd5c6789bccb()
			if __obf_67008a6a9e5ba828.Error != nil {
				return nil
			}
			if __obf_dab9baaadfa7c8c2 != '\\' {
				__obf_67008a6a9e5ba828.__obf_3284a1eaa2a0abb6()
				__obf_12c21b79fa86dcba = __obf_9717752b60f28ee0(__obf_12c21b79fa86dcba, __obf_14c0fc639db9b91c)
				return __obf_12c21b79fa86dcba
			}
			__obf_dab9baaadfa7c8c2 = __obf_67008a6a9e5ba828.__obf_ea3ebd5c6789bccb()
			if __obf_67008a6a9e5ba828.Error != nil {
				return nil
			}
			if __obf_dab9baaadfa7c8c2 != 'u' {
				__obf_12c21b79fa86dcba = __obf_9717752b60f28ee0(__obf_12c21b79fa86dcba, __obf_14c0fc639db9b91c)
				return __obf_67008a6a9e5ba828.__obf_57340dc7a0cce10d(__obf_dab9baaadfa7c8c2, __obf_12c21b79fa86dcba)
			}
			__obf_63efaa8efa743d97 := __obf_67008a6a9e5ba828.__obf_5d1e2ea19783b872()
			if __obf_67008a6a9e5ba828.Error != nil {
				return nil
			}
			__obf_c301db7626d23a96 := utf16.DecodeRune(__obf_14c0fc639db9b91c, __obf_63efaa8efa743d97)
			if __obf_c301db7626d23a96 == '\uFFFD' {
				__obf_12c21b79fa86dcba = __obf_9717752b60f28ee0(__obf_12c21b79fa86dcba, __obf_14c0fc639db9b91c)
				__obf_12c21b79fa86dcba = __obf_9717752b60f28ee0(__obf_12c21b79fa86dcba, __obf_63efaa8efa743d97)
			} else {
				__obf_12c21b79fa86dcba = __obf_9717752b60f28ee0(__obf_12c21b79fa86dcba, __obf_c301db7626d23a96)
			}
		} else {
			__obf_12c21b79fa86dcba = __obf_9717752b60f28ee0(__obf_12c21b79fa86dcba, __obf_14c0fc639db9b91c)
		}
	case '"':
		__obf_12c21b79fa86dcba = append(__obf_12c21b79fa86dcba, '"')
	case '\\':
		__obf_12c21b79fa86dcba = append(__obf_12c21b79fa86dcba, '\\')
	case '/':
		__obf_12c21b79fa86dcba = append(__obf_12c21b79fa86dcba, '/')
	case 'b':
		__obf_12c21b79fa86dcba = append(__obf_12c21b79fa86dcba, '\b')
	case 'f':
		__obf_12c21b79fa86dcba = append(__obf_12c21b79fa86dcba, '\f')
	case 'n':
		__obf_12c21b79fa86dcba = append(__obf_12c21b79fa86dcba, '\n')
	case 'r':
		__obf_12c21b79fa86dcba = append(__obf_12c21b79fa86dcba, '\r')
	case 't':
		__obf_12c21b79fa86dcba = append(__obf_12c21b79fa86dcba, '\t')
	default:
		__obf_67008a6a9e5ba828.
			ReportError("readEscapedChar",
				`invalid escape char after \`)
		return nil
	}
	return __obf_12c21b79fa86dcba
}

// ReadStringAsSlice read string from iterator without copying into string form.
// The []byte can not be kept, as it will change after next iterator call.
func (__obf_67008a6a9e5ba828 *Iterator) ReadStringAsSlice() (__obf_5dabcdfee5097ed6 []byte) {
	__obf_dab9baaadfa7c8c2 := __obf_67008a6a9e5ba828.__obf_b781a59d5a0d2490()
	if __obf_dab9baaadfa7c8c2 == '"' {
		for __obf_2deec7c38ffb6ae3 := __obf_67008a6a9e5ba828.__obf_14babd6f9a55bd36;
		// require ascii string and no escape
		// for: field name, base64, number
		__obf_2deec7c38ffb6ae3 < __obf_67008a6a9e5ba828.__obf_3a36550914545c79; __obf_2deec7c38ffb6ae3++ {

			if __obf_67008a6a9e5ba828.__obf_9fc06d9180f0daca[__obf_2deec7c38ffb6ae3] == '"' {
				__obf_5dabcdfee5097ed6 = // fast path: reuse the underlying buffer
					__obf_67008a6a9e5ba828.__obf_9fc06d9180f0daca[__obf_67008a6a9e5ba828.__obf_14babd6f9a55bd36:__obf_2deec7c38ffb6ae3]
				__obf_67008a6a9e5ba828.__obf_14babd6f9a55bd36 = __obf_2deec7c38ffb6ae3 + 1
				return __obf_5dabcdfee5097ed6
			}
		}
		__obf_ea6d371e3e0b2e07 := __obf_67008a6a9e5ba828.__obf_3a36550914545c79 - __obf_67008a6a9e5ba828.__obf_14babd6f9a55bd36
		__obf_cd8f569527424a26 := make([]byte, __obf_ea6d371e3e0b2e07, __obf_ea6d371e3e0b2e07*2)
		copy(__obf_cd8f569527424a26, __obf_67008a6a9e5ba828.__obf_9fc06d9180f0daca[__obf_67008a6a9e5ba828.__obf_14babd6f9a55bd36:__obf_67008a6a9e5ba828.__obf_3a36550914545c79])
		__obf_67008a6a9e5ba828.__obf_14babd6f9a55bd36 = __obf_67008a6a9e5ba828.__obf_3a36550914545c79
		for __obf_67008a6a9e5ba828.Error == nil {
			__obf_dab9baaadfa7c8c2 := __obf_67008a6a9e5ba828.__obf_ea3ebd5c6789bccb()
			if __obf_dab9baaadfa7c8c2 == '"' {
				return __obf_cd8f569527424a26
			}
			__obf_cd8f569527424a26 = append(__obf_cd8f569527424a26, __obf_dab9baaadfa7c8c2)
		}
		return __obf_cd8f569527424a26
	}
	__obf_67008a6a9e5ba828.
		ReportError("ReadStringAsSlice", `expects " or n, but found `+string([]byte{__obf_dab9baaadfa7c8c2}))
	return
}

func (__obf_67008a6a9e5ba828 *Iterator) __obf_5d1e2ea19783b872() (__obf_5dabcdfee5097ed6 rune) {
	for __obf_2deec7c38ffb6ae3 := 0; __obf_2deec7c38ffb6ae3 < 4; __obf_2deec7c38ffb6ae3++ {
		__obf_dab9baaadfa7c8c2 := __obf_67008a6a9e5ba828.__obf_ea3ebd5c6789bccb()
		if __obf_67008a6a9e5ba828.Error != nil {
			return
		}
		if __obf_dab9baaadfa7c8c2 >= '0' && __obf_dab9baaadfa7c8c2 <= '9' {
			__obf_5dabcdfee5097ed6 = __obf_5dabcdfee5097ed6*16 + rune(__obf_dab9baaadfa7c8c2-'0')
		} else if __obf_dab9baaadfa7c8c2 >= 'a' && __obf_dab9baaadfa7c8c2 <= 'f' {
			__obf_5dabcdfee5097ed6 = __obf_5dabcdfee5097ed6*16 + rune(__obf_dab9baaadfa7c8c2-'a'+10)
		} else if __obf_dab9baaadfa7c8c2 >= 'A' && __obf_dab9baaadfa7c8c2 <= 'F' {
			__obf_5dabcdfee5097ed6 = __obf_5dabcdfee5097ed6*16 + rune(__obf_dab9baaadfa7c8c2-'A'+10)
		} else {
			__obf_67008a6a9e5ba828.
				ReportError("readU4", "expects 0~9 or a~f, but found "+string([]byte{__obf_dab9baaadfa7c8c2}))
			return
		}
	}
	return __obf_5dabcdfee5097ed6
}

const (
	__obf_a9782c4a1bf8b50a = 0x00
	__obf_82cd435fbdf484f3 = // 0000 0000
	0x80
	__obf_75ac641cf429ef03 = // 1000 0000
	0xC0
	__obf_d83083f812432ec2 = // 1100 0000
	0xE0
	__obf_5c2242c59936e838 = // 1110 0000
	0xF0
	__obf_57c9d4421143801a = // 1111 0000
	0xF8
	__obf_237c261169ffa74c = // 1111 1000

	0x3F
	__obf_e29700c671609e81 = // 0011 1111
	0x1F
	__obf_703e3520b49146bd = // 0001 1111
	0x0F
	__obf_42db783fe8b3715c = // 0000 1111
	0x07
	__obf_d4704ef3e983b5c3 = // 0000 0111

	1<<7 - 1
	__obf_b7a94ae2a15e20ec = 1<<11 - 1
	__obf_e1dcb874310b5552 = 1<<16 - 1
	__obf_8700e3ed5e7d5d19 = 0xD800
	__obf_39f410a7d6a5db43 = 0xDFFF
	__obf_34a85f1f8455029c = '\U0010FFFF'
	__obf_3f85db3fcf6c0fb2 = // Maximum valid Unicode code point.
	'\uFFFD'                 // the "error" Rune or "Unicode replacement character"
)

func __obf_9717752b60f28ee0(__obf_82ddba9040ff8d8c []byte, __obf_14c0fc639db9b91c rune) []byte {
	// Negative values are erroneous. Making it unsigned addresses the problem.
	switch __obf_2deec7c38ffb6ae3 := uint32(__obf_14c0fc639db9b91c); {
	case __obf_2deec7c38ffb6ae3 <= __obf_d4704ef3e983b5c3:
		__obf_82ddba9040ff8d8c = append(__obf_82ddba9040ff8d8c, byte(__obf_14c0fc639db9b91c))
		return __obf_82ddba9040ff8d8c
	case __obf_2deec7c38ffb6ae3 <= __obf_b7a94ae2a15e20ec:
		__obf_82ddba9040ff8d8c = append(__obf_82ddba9040ff8d8c, __obf_75ac641cf429ef03|byte(__obf_14c0fc639db9b91c>>6))
		__obf_82ddba9040ff8d8c = append(__obf_82ddba9040ff8d8c, __obf_82cd435fbdf484f3|byte(__obf_14c0fc639db9b91c)&__obf_237c261169ffa74c)
		return __obf_82ddba9040ff8d8c
	case __obf_2deec7c38ffb6ae3 > __obf_34a85f1f8455029c, __obf_8700e3ed5e7d5d19 <= __obf_2deec7c38ffb6ae3 && __obf_2deec7c38ffb6ae3 <= __obf_39f410a7d6a5db43:
		__obf_14c0fc639db9b91c = __obf_3f85db3fcf6c0fb2
		fallthrough
	case __obf_2deec7c38ffb6ae3 <= __obf_e1dcb874310b5552:
		__obf_82ddba9040ff8d8c = append(__obf_82ddba9040ff8d8c, __obf_d83083f812432ec2|byte(__obf_14c0fc639db9b91c>>12))
		__obf_82ddba9040ff8d8c = append(__obf_82ddba9040ff8d8c, __obf_82cd435fbdf484f3|byte(__obf_14c0fc639db9b91c>>6)&__obf_237c261169ffa74c)
		__obf_82ddba9040ff8d8c = append(__obf_82ddba9040ff8d8c, __obf_82cd435fbdf484f3|byte(__obf_14c0fc639db9b91c)&__obf_237c261169ffa74c)
		return __obf_82ddba9040ff8d8c
	default:
		__obf_82ddba9040ff8d8c = append(__obf_82ddba9040ff8d8c, __obf_5c2242c59936e838|byte(__obf_14c0fc639db9b91c>>18))
		__obf_82ddba9040ff8d8c = append(__obf_82ddba9040ff8d8c, __obf_82cd435fbdf484f3|byte(__obf_14c0fc639db9b91c>>12)&__obf_237c261169ffa74c)
		__obf_82ddba9040ff8d8c = append(__obf_82ddba9040ff8d8c, __obf_82cd435fbdf484f3|byte(__obf_14c0fc639db9b91c>>6)&__obf_237c261169ffa74c)
		__obf_82ddba9040ff8d8c = append(__obf_82ddba9040ff8d8c, __obf_82cd435fbdf484f3|byte(__obf_14c0fc639db9b91c)&__obf_237c261169ffa74c)
		return __obf_82ddba9040ff8d8c
	}
}
