package __obf_703d23ba520c3295

import (
	"fmt"
	"unicode/utf16"
)

// ReadString read string from iterator
func (__obf_47edab4c16a2d88d *Iterator) ReadString() (__obf_551cbec38242ce0c string) {
	__obf_bd08f5161fff294a := __obf_47edab4c16a2d88d.__obf_6f7a9d6f8e846eec()
	if __obf_bd08f5161fff294a == '"' {
		for __obf_b0a5d2bd48690f1d := __obf_47edab4c16a2d88d.__obf_da6aa1cfbd772c11; __obf_b0a5d2bd48690f1d < __obf_47edab4c16a2d88d.__obf_687f71f0b8c00ef5; __obf_b0a5d2bd48690f1d++ {
			__obf_bd08f5161fff294a := __obf_47edab4c16a2d88d.__obf_a065f8e0da5f5952[__obf_b0a5d2bd48690f1d]
			if __obf_bd08f5161fff294a == '"' {
				__obf_551cbec38242ce0c = string(__obf_47edab4c16a2d88d.__obf_a065f8e0da5f5952[__obf_47edab4c16a2d88d.__obf_da6aa1cfbd772c11:__obf_b0a5d2bd48690f1d])
				__obf_47edab4c16a2d88d.__obf_da6aa1cfbd772c11 = __obf_b0a5d2bd48690f1d + 1
				return __obf_551cbec38242ce0c
			} else if __obf_bd08f5161fff294a == '\\' {
				break
			} else if __obf_bd08f5161fff294a < ' ' {
				__obf_47edab4c16a2d88d.
					ReportError("ReadString",
						fmt.Sprintf(`invalid control character found: %d`, __obf_bd08f5161fff294a))
				return
			}
		}
		return __obf_47edab4c16a2d88d.__obf_61de6bc49a932c21()
	} else if __obf_bd08f5161fff294a == 'n' {
		__obf_47edab4c16a2d88d.__obf_0704d4be12be5e96('u', 'l', 'l')
		return ""
	}
	__obf_47edab4c16a2d88d.
		ReportError("ReadString", `expects " or n, but found `+string([]byte{__obf_bd08f5161fff294a}))
	return
}

func (__obf_47edab4c16a2d88d *Iterator) __obf_61de6bc49a932c21() (__obf_551cbec38242ce0c string) {
	var __obf_44b48c26051f8033 []byte
	var __obf_bd08f5161fff294a byte
	for __obf_47edab4c16a2d88d.Error == nil {
		__obf_bd08f5161fff294a = __obf_47edab4c16a2d88d.__obf_2fa919905fa99cc3()
		if __obf_bd08f5161fff294a == '"' {
			return string(__obf_44b48c26051f8033)
		}
		if __obf_bd08f5161fff294a == '\\' {
			__obf_bd08f5161fff294a = __obf_47edab4c16a2d88d.__obf_2fa919905fa99cc3()
			__obf_44b48c26051f8033 = __obf_47edab4c16a2d88d.__obf_26a57e7c441237dd(__obf_bd08f5161fff294a, __obf_44b48c26051f8033)
		} else {
			__obf_44b48c26051f8033 = append(__obf_44b48c26051f8033, __obf_bd08f5161fff294a)
		}
	}
	__obf_47edab4c16a2d88d.
		ReportError("readStringSlowPath", "unexpected end of input")
	return
}

func (__obf_47edab4c16a2d88d *Iterator) __obf_26a57e7c441237dd(__obf_bd08f5161fff294a byte, __obf_44b48c26051f8033 []byte) []byte {
	switch __obf_bd08f5161fff294a {
	case 'u':
		__obf_a44636b770caa848 := __obf_47edab4c16a2d88d.__obf_8c66bb094410791e()
		if utf16.IsSurrogate(__obf_a44636b770caa848) {
			__obf_bd08f5161fff294a = __obf_47edab4c16a2d88d.__obf_2fa919905fa99cc3()
			if __obf_47edab4c16a2d88d.Error != nil {
				return nil
			}
			if __obf_bd08f5161fff294a != '\\' {
				__obf_47edab4c16a2d88d.__obf_743b1a47c346c8a3()
				__obf_44b48c26051f8033 = __obf_bae71db04b4b33ab(__obf_44b48c26051f8033, __obf_a44636b770caa848)
				return __obf_44b48c26051f8033
			}
			__obf_bd08f5161fff294a = __obf_47edab4c16a2d88d.__obf_2fa919905fa99cc3()
			if __obf_47edab4c16a2d88d.Error != nil {
				return nil
			}
			if __obf_bd08f5161fff294a != 'u' {
				__obf_44b48c26051f8033 = __obf_bae71db04b4b33ab(__obf_44b48c26051f8033, __obf_a44636b770caa848)
				return __obf_47edab4c16a2d88d.__obf_26a57e7c441237dd(__obf_bd08f5161fff294a, __obf_44b48c26051f8033)
			}
			__obf_f2b2cb07840ef812 := __obf_47edab4c16a2d88d.__obf_8c66bb094410791e()
			if __obf_47edab4c16a2d88d.Error != nil {
				return nil
			}
			__obf_fcaa8188c00cc090 := utf16.DecodeRune(__obf_a44636b770caa848, __obf_f2b2cb07840ef812)
			if __obf_fcaa8188c00cc090 == '\uFFFD' {
				__obf_44b48c26051f8033 = __obf_bae71db04b4b33ab(__obf_44b48c26051f8033, __obf_a44636b770caa848)
				__obf_44b48c26051f8033 = __obf_bae71db04b4b33ab(__obf_44b48c26051f8033, __obf_f2b2cb07840ef812)
			} else {
				__obf_44b48c26051f8033 = __obf_bae71db04b4b33ab(__obf_44b48c26051f8033, __obf_fcaa8188c00cc090)
			}
		} else {
			__obf_44b48c26051f8033 = __obf_bae71db04b4b33ab(__obf_44b48c26051f8033, __obf_a44636b770caa848)
		}
	case '"':
		__obf_44b48c26051f8033 = append(__obf_44b48c26051f8033, '"')
	case '\\':
		__obf_44b48c26051f8033 = append(__obf_44b48c26051f8033, '\\')
	case '/':
		__obf_44b48c26051f8033 = append(__obf_44b48c26051f8033, '/')
	case 'b':
		__obf_44b48c26051f8033 = append(__obf_44b48c26051f8033, '\b')
	case 'f':
		__obf_44b48c26051f8033 = append(__obf_44b48c26051f8033, '\f')
	case 'n':
		__obf_44b48c26051f8033 = append(__obf_44b48c26051f8033, '\n')
	case 'r':
		__obf_44b48c26051f8033 = append(__obf_44b48c26051f8033, '\r')
	case 't':
		__obf_44b48c26051f8033 = append(__obf_44b48c26051f8033, '\t')
	default:
		__obf_47edab4c16a2d88d.
			ReportError("readEscapedChar",
				`invalid escape char after \`)
		return nil
	}
	return __obf_44b48c26051f8033
}

// ReadStringAsSlice read string from iterator without copying into string form.
// The []byte can not be kept, as it will change after next iterator call.
func (__obf_47edab4c16a2d88d *Iterator) ReadStringAsSlice() (__obf_551cbec38242ce0c []byte) {
	__obf_bd08f5161fff294a := __obf_47edab4c16a2d88d.__obf_6f7a9d6f8e846eec()
	if __obf_bd08f5161fff294a == '"' {
		for __obf_b0a5d2bd48690f1d := __obf_47edab4c16a2d88d.__obf_da6aa1cfbd772c11;
		// require ascii string and no escape
		// for: field name, base64, number
		__obf_b0a5d2bd48690f1d < __obf_47edab4c16a2d88d.__obf_687f71f0b8c00ef5; __obf_b0a5d2bd48690f1d++ {

			if __obf_47edab4c16a2d88d.__obf_a065f8e0da5f5952[__obf_b0a5d2bd48690f1d] == '"' {
				__obf_551cbec38242ce0c = // fast path: reuse the underlying buffer
					__obf_47edab4c16a2d88d.__obf_a065f8e0da5f5952[__obf_47edab4c16a2d88d.__obf_da6aa1cfbd772c11:__obf_b0a5d2bd48690f1d]
				__obf_47edab4c16a2d88d.__obf_da6aa1cfbd772c11 = __obf_b0a5d2bd48690f1d + 1
				return __obf_551cbec38242ce0c
			}
		}
		__obf_b0144297144986c4 := __obf_47edab4c16a2d88d.__obf_687f71f0b8c00ef5 - __obf_47edab4c16a2d88d.__obf_da6aa1cfbd772c11
		__obf_6fd9ea911f2d54de := make([]byte, __obf_b0144297144986c4, __obf_b0144297144986c4*2)
		copy(__obf_6fd9ea911f2d54de, __obf_47edab4c16a2d88d.__obf_a065f8e0da5f5952[__obf_47edab4c16a2d88d.__obf_da6aa1cfbd772c11:__obf_47edab4c16a2d88d.__obf_687f71f0b8c00ef5])
		__obf_47edab4c16a2d88d.__obf_da6aa1cfbd772c11 = __obf_47edab4c16a2d88d.__obf_687f71f0b8c00ef5
		for __obf_47edab4c16a2d88d.Error == nil {
			__obf_bd08f5161fff294a := __obf_47edab4c16a2d88d.__obf_2fa919905fa99cc3()
			if __obf_bd08f5161fff294a == '"' {
				return __obf_6fd9ea911f2d54de
			}
			__obf_6fd9ea911f2d54de = append(__obf_6fd9ea911f2d54de, __obf_bd08f5161fff294a)
		}
		return __obf_6fd9ea911f2d54de
	}
	__obf_47edab4c16a2d88d.
		ReportError("ReadStringAsSlice", `expects " or n, but found `+string([]byte{__obf_bd08f5161fff294a}))
	return
}

func (__obf_47edab4c16a2d88d *Iterator) __obf_8c66bb094410791e() (__obf_551cbec38242ce0c rune) {
	for __obf_b0a5d2bd48690f1d := 0; __obf_b0a5d2bd48690f1d < 4; __obf_b0a5d2bd48690f1d++ {
		__obf_bd08f5161fff294a := __obf_47edab4c16a2d88d.__obf_2fa919905fa99cc3()
		if __obf_47edab4c16a2d88d.Error != nil {
			return
		}
		if __obf_bd08f5161fff294a >= '0' && __obf_bd08f5161fff294a <= '9' {
			__obf_551cbec38242ce0c = __obf_551cbec38242ce0c*16 + rune(__obf_bd08f5161fff294a-'0')
		} else if __obf_bd08f5161fff294a >= 'a' && __obf_bd08f5161fff294a <= 'f' {
			__obf_551cbec38242ce0c = __obf_551cbec38242ce0c*16 + rune(__obf_bd08f5161fff294a-'a'+10)
		} else if __obf_bd08f5161fff294a >= 'A' && __obf_bd08f5161fff294a <= 'F' {
			__obf_551cbec38242ce0c = __obf_551cbec38242ce0c*16 + rune(__obf_bd08f5161fff294a-'A'+10)
		} else {
			__obf_47edab4c16a2d88d.
				ReportError("readU4", "expects 0~9 or a~f, but found "+string([]byte{__obf_bd08f5161fff294a}))
			return
		}
	}
	return __obf_551cbec38242ce0c
}

const (
	__obf_2ffca0acdd6f48cb = 0x00
	__obf_04a1f63de4da8292 = // 0000 0000
	0x80
	__obf_f064f18789555652 = // 1000 0000
	0xC0
	__obf_319827e0d402ff5c = // 1100 0000
	0xE0
	__obf_445b2c4e109493b8 = // 1110 0000
	0xF0
	__obf_c43f1bbe41457db5 = // 1111 0000
	0xF8
	__obf_0f40016f272b2e0e = // 1111 1000

	0x3F
	__obf_2b1b9b59b7a3cc87 = // 0011 1111
	0x1F
	__obf_2d6e67db2b4d07a3 = // 0001 1111
	0x0F
	__obf_d5bafb593b356c01 = // 0000 1111
	0x07
	__obf_55cfddde33a36ccf = // 0000 0111

	1<<7 - 1
	__obf_660e8c2eedf27453 = 1<<11 - 1
	__obf_514ab62506bec9f8 = 1<<16 - 1
	__obf_7fa51ad04124416a = 0xD800
	__obf_7ec173b5f5bf907b = 0xDFFF
	__obf_1024c82656bfe07a = '\U0010FFFF'
	__obf_64886e7251f3524c = // Maximum valid Unicode code point.
	'\uFFFD'                 // the "error" Rune or "Unicode replacement character"
)

func __obf_bae71db04b4b33ab(__obf_f9ed235be8f52eaf []byte, __obf_a44636b770caa848 rune) []byte {
	// Negative values are erroneous. Making it unsigned addresses the problem.
	switch __obf_b0a5d2bd48690f1d := uint32(__obf_a44636b770caa848); {
	case __obf_b0a5d2bd48690f1d <= __obf_55cfddde33a36ccf:
		__obf_f9ed235be8f52eaf = append(__obf_f9ed235be8f52eaf, byte(__obf_a44636b770caa848))
		return __obf_f9ed235be8f52eaf
	case __obf_b0a5d2bd48690f1d <= __obf_660e8c2eedf27453:
		__obf_f9ed235be8f52eaf = append(__obf_f9ed235be8f52eaf, __obf_f064f18789555652|byte(__obf_a44636b770caa848>>6))
		__obf_f9ed235be8f52eaf = append(__obf_f9ed235be8f52eaf, __obf_04a1f63de4da8292|byte(__obf_a44636b770caa848)&__obf_0f40016f272b2e0e)
		return __obf_f9ed235be8f52eaf
	case __obf_b0a5d2bd48690f1d > __obf_1024c82656bfe07a, __obf_7fa51ad04124416a <= __obf_b0a5d2bd48690f1d && __obf_b0a5d2bd48690f1d <= __obf_7ec173b5f5bf907b:
		__obf_a44636b770caa848 = __obf_64886e7251f3524c
		fallthrough
	case __obf_b0a5d2bd48690f1d <= __obf_514ab62506bec9f8:
		__obf_f9ed235be8f52eaf = append(__obf_f9ed235be8f52eaf, __obf_319827e0d402ff5c|byte(__obf_a44636b770caa848>>12))
		__obf_f9ed235be8f52eaf = append(__obf_f9ed235be8f52eaf, __obf_04a1f63de4da8292|byte(__obf_a44636b770caa848>>6)&__obf_0f40016f272b2e0e)
		__obf_f9ed235be8f52eaf = append(__obf_f9ed235be8f52eaf, __obf_04a1f63de4da8292|byte(__obf_a44636b770caa848)&__obf_0f40016f272b2e0e)
		return __obf_f9ed235be8f52eaf
	default:
		__obf_f9ed235be8f52eaf = append(__obf_f9ed235be8f52eaf, __obf_445b2c4e109493b8|byte(__obf_a44636b770caa848>>18))
		__obf_f9ed235be8f52eaf = append(__obf_f9ed235be8f52eaf, __obf_04a1f63de4da8292|byte(__obf_a44636b770caa848>>12)&__obf_0f40016f272b2e0e)
		__obf_f9ed235be8f52eaf = append(__obf_f9ed235be8f52eaf, __obf_04a1f63de4da8292|byte(__obf_a44636b770caa848>>6)&__obf_0f40016f272b2e0e)
		__obf_f9ed235be8f52eaf = append(__obf_f9ed235be8f52eaf, __obf_04a1f63de4da8292|byte(__obf_a44636b770caa848)&__obf_0f40016f272b2e0e)
		return __obf_f9ed235be8f52eaf
	}
}
