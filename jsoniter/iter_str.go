package __obf_c7b79b12b33d8238

import (
	"fmt"
	"unicode/utf16"
)

// ReadString read string from iterator
func (__obf_0da8c843dad13139 *Iterator) ReadString() (__obf_9a8638dac9e1c083 string) {
	__obf_12577c03a12f0d2c := __obf_0da8c843dad13139.__obf_2b436fcb1c0ca36d()
	if __obf_12577c03a12f0d2c == '"' {
		for __obf_a051269af3a541bb := __obf_0da8c843dad13139.__obf_6a63c9ac34fe84e2; __obf_a051269af3a541bb < __obf_0da8c843dad13139.__obf_840246080559848c; __obf_a051269af3a541bb++ {
			__obf_12577c03a12f0d2c := __obf_0da8c843dad13139.__obf_684d738bc3ac851a[__obf_a051269af3a541bb]
			if __obf_12577c03a12f0d2c == '"' {
				__obf_9a8638dac9e1c083 = string(__obf_0da8c843dad13139.__obf_684d738bc3ac851a[__obf_0da8c843dad13139.__obf_6a63c9ac34fe84e2:__obf_a051269af3a541bb])
				__obf_0da8c843dad13139.__obf_6a63c9ac34fe84e2 = __obf_a051269af3a541bb + 1
				return __obf_9a8638dac9e1c083
			} else if __obf_12577c03a12f0d2c == '\\' {
				break
			} else if __obf_12577c03a12f0d2c < ' ' {
				__obf_0da8c843dad13139.
					ReportError("ReadString",
						fmt.Sprintf(`invalid control character found: %d`, __obf_12577c03a12f0d2c))
				return
			}
		}
		return __obf_0da8c843dad13139.__obf_22b06ef3645a8e41()
	} else if __obf_12577c03a12f0d2c == 'n' {
		__obf_0da8c843dad13139.__obf_5da8d54e7a1e782c('u', 'l', 'l')
		return ""
	}
	__obf_0da8c843dad13139.
		ReportError("ReadString", `expects " or n, but found `+string([]byte{__obf_12577c03a12f0d2c}))
	return
}

func (__obf_0da8c843dad13139 *Iterator) __obf_22b06ef3645a8e41() (__obf_9a8638dac9e1c083 string) {
	var __obf_a3eaafc22faf1644 []byte
	var __obf_12577c03a12f0d2c byte
	for __obf_0da8c843dad13139.Error == nil {
		__obf_12577c03a12f0d2c = __obf_0da8c843dad13139.__obf_fe749dd9dadf32f8()
		if __obf_12577c03a12f0d2c == '"' {
			return string(__obf_a3eaafc22faf1644)
		}
		if __obf_12577c03a12f0d2c == '\\' {
			__obf_12577c03a12f0d2c = __obf_0da8c843dad13139.__obf_fe749dd9dadf32f8()
			__obf_a3eaafc22faf1644 = __obf_0da8c843dad13139.__obf_f951a1f38f8af71c(__obf_12577c03a12f0d2c, __obf_a3eaafc22faf1644)
		} else {
			__obf_a3eaafc22faf1644 = append(__obf_a3eaafc22faf1644, __obf_12577c03a12f0d2c)
		}
	}
	__obf_0da8c843dad13139.
		ReportError("readStringSlowPath", "unexpected end of input")
	return
}

func (__obf_0da8c843dad13139 *Iterator) __obf_f951a1f38f8af71c(__obf_12577c03a12f0d2c byte, __obf_a3eaafc22faf1644 []byte) []byte {
	switch __obf_12577c03a12f0d2c {
	case 'u':
		__obf_464eaebd868bf0fe := __obf_0da8c843dad13139.__obf_feed2b8accd8926b()
		if utf16.IsSurrogate(__obf_464eaebd868bf0fe) {
			__obf_12577c03a12f0d2c = __obf_0da8c843dad13139.__obf_fe749dd9dadf32f8()
			if __obf_0da8c843dad13139.Error != nil {
				return nil
			}
			if __obf_12577c03a12f0d2c != '\\' {
				__obf_0da8c843dad13139.__obf_a675366c80290b83()
				__obf_a3eaafc22faf1644 = __obf_9bca19f1764e1c19(__obf_a3eaafc22faf1644, __obf_464eaebd868bf0fe)
				return __obf_a3eaafc22faf1644
			}
			__obf_12577c03a12f0d2c = __obf_0da8c843dad13139.__obf_fe749dd9dadf32f8()
			if __obf_0da8c843dad13139.Error != nil {
				return nil
			}
			if __obf_12577c03a12f0d2c != 'u' {
				__obf_a3eaafc22faf1644 = __obf_9bca19f1764e1c19(__obf_a3eaafc22faf1644, __obf_464eaebd868bf0fe)
				return __obf_0da8c843dad13139.__obf_f951a1f38f8af71c(__obf_12577c03a12f0d2c, __obf_a3eaafc22faf1644)
			}
			__obf_a818cce3414b2179 := __obf_0da8c843dad13139.__obf_feed2b8accd8926b()
			if __obf_0da8c843dad13139.Error != nil {
				return nil
			}
			__obf_a672e99b6aa86630 := utf16.DecodeRune(__obf_464eaebd868bf0fe, __obf_a818cce3414b2179)
			if __obf_a672e99b6aa86630 == '\uFFFD' {
				__obf_a3eaafc22faf1644 = __obf_9bca19f1764e1c19(__obf_a3eaafc22faf1644, __obf_464eaebd868bf0fe)
				__obf_a3eaafc22faf1644 = __obf_9bca19f1764e1c19(__obf_a3eaafc22faf1644, __obf_a818cce3414b2179)
			} else {
				__obf_a3eaafc22faf1644 = __obf_9bca19f1764e1c19(__obf_a3eaafc22faf1644, __obf_a672e99b6aa86630)
			}
		} else {
			__obf_a3eaafc22faf1644 = __obf_9bca19f1764e1c19(__obf_a3eaafc22faf1644, __obf_464eaebd868bf0fe)
		}
	case '"':
		__obf_a3eaafc22faf1644 = append(__obf_a3eaafc22faf1644, '"')
	case '\\':
		__obf_a3eaafc22faf1644 = append(__obf_a3eaafc22faf1644, '\\')
	case '/':
		__obf_a3eaafc22faf1644 = append(__obf_a3eaafc22faf1644, '/')
	case 'b':
		__obf_a3eaafc22faf1644 = append(__obf_a3eaafc22faf1644, '\b')
	case 'f':
		__obf_a3eaafc22faf1644 = append(__obf_a3eaafc22faf1644, '\f')
	case 'n':
		__obf_a3eaafc22faf1644 = append(__obf_a3eaafc22faf1644, '\n')
	case 'r':
		__obf_a3eaafc22faf1644 = append(__obf_a3eaafc22faf1644, '\r')
	case 't':
		__obf_a3eaafc22faf1644 = append(__obf_a3eaafc22faf1644, '\t')
	default:
		__obf_0da8c843dad13139.
			ReportError("readEscapedChar",
				`invalid escape char after \`)
		return nil
	}
	return __obf_a3eaafc22faf1644
}

// ReadStringAsSlice read string from iterator without copying into string form.
// The []byte can not be kept, as it will change after next iterator call.
func (__obf_0da8c843dad13139 *Iterator) ReadStringAsSlice() (__obf_9a8638dac9e1c083 []byte) {
	__obf_12577c03a12f0d2c := __obf_0da8c843dad13139.__obf_2b436fcb1c0ca36d()
	if __obf_12577c03a12f0d2c == '"' {
		for __obf_a051269af3a541bb := __obf_0da8c843dad13139.__obf_6a63c9ac34fe84e2;
		// require ascii string and no escape
		// for: field name, base64, number
		__obf_a051269af3a541bb < __obf_0da8c843dad13139.__obf_840246080559848c; __obf_a051269af3a541bb++ {

			if __obf_0da8c843dad13139.__obf_684d738bc3ac851a[__obf_a051269af3a541bb] == '"' {
				__obf_9a8638dac9e1c083 = // fast path: reuse the underlying buffer
					__obf_0da8c843dad13139.__obf_684d738bc3ac851a[__obf_0da8c843dad13139.__obf_6a63c9ac34fe84e2:__obf_a051269af3a541bb]
				__obf_0da8c843dad13139.__obf_6a63c9ac34fe84e2 = __obf_a051269af3a541bb + 1
				return __obf_9a8638dac9e1c083
			}
		}
		__obf_29e70b787e35d945 := __obf_0da8c843dad13139.__obf_840246080559848c - __obf_0da8c843dad13139.__obf_6a63c9ac34fe84e2
		__obf_a3c6c71d32db2eb1 := make([]byte, __obf_29e70b787e35d945, __obf_29e70b787e35d945*2)
		copy(__obf_a3c6c71d32db2eb1, __obf_0da8c843dad13139.__obf_684d738bc3ac851a[__obf_0da8c843dad13139.__obf_6a63c9ac34fe84e2:__obf_0da8c843dad13139.__obf_840246080559848c])
		__obf_0da8c843dad13139.__obf_6a63c9ac34fe84e2 = __obf_0da8c843dad13139.__obf_840246080559848c
		for __obf_0da8c843dad13139.Error == nil {
			__obf_12577c03a12f0d2c := __obf_0da8c843dad13139.__obf_fe749dd9dadf32f8()
			if __obf_12577c03a12f0d2c == '"' {
				return __obf_a3c6c71d32db2eb1
			}
			__obf_a3c6c71d32db2eb1 = append(__obf_a3c6c71d32db2eb1, __obf_12577c03a12f0d2c)
		}
		return __obf_a3c6c71d32db2eb1
	}
	__obf_0da8c843dad13139.
		ReportError("ReadStringAsSlice", `expects " or n, but found `+string([]byte{__obf_12577c03a12f0d2c}))
	return
}

func (__obf_0da8c843dad13139 *Iterator) __obf_feed2b8accd8926b() (__obf_9a8638dac9e1c083 rune) {
	for __obf_a051269af3a541bb := 0; __obf_a051269af3a541bb < 4; __obf_a051269af3a541bb++ {
		__obf_12577c03a12f0d2c := __obf_0da8c843dad13139.__obf_fe749dd9dadf32f8()
		if __obf_0da8c843dad13139.Error != nil {
			return
		}
		if __obf_12577c03a12f0d2c >= '0' && __obf_12577c03a12f0d2c <= '9' {
			__obf_9a8638dac9e1c083 = __obf_9a8638dac9e1c083*16 + rune(__obf_12577c03a12f0d2c-'0')
		} else if __obf_12577c03a12f0d2c >= 'a' && __obf_12577c03a12f0d2c <= 'f' {
			__obf_9a8638dac9e1c083 = __obf_9a8638dac9e1c083*16 + rune(__obf_12577c03a12f0d2c-'a'+10)
		} else if __obf_12577c03a12f0d2c >= 'A' && __obf_12577c03a12f0d2c <= 'F' {
			__obf_9a8638dac9e1c083 = __obf_9a8638dac9e1c083*16 + rune(__obf_12577c03a12f0d2c-'A'+10)
		} else {
			__obf_0da8c843dad13139.
				ReportError("readU4", "expects 0~9 or a~f, but found "+string([]byte{__obf_12577c03a12f0d2c}))
			return
		}
	}
	return __obf_9a8638dac9e1c083
}

const (
	__obf_bd94ee5b39180582 = 0x00
	__obf_c80d46f677b920cc = // 0000 0000
	0x80
	__obf_6a7093331d3c1ead = // 1000 0000
	0xC0
	__obf_6ae418366c514de3 = // 1100 0000
	0xE0
	__obf_a3e07c5094670ee7 = // 1110 0000
	0xF0
	__obf_d41446154cb4c77d = // 1111 0000
	0xF8
	__obf_39db8b683ce0ed16 = // 1111 1000

	0x3F
	__obf_53a88abbe79362e9 = // 0011 1111
	0x1F
	__obf_227486b942a45954 = // 0001 1111
	0x0F
	__obf_339dddd1c85f3166 = // 0000 1111
	0x07
	__obf_fec31f5058f7df5d = // 0000 0111

	1<<7 - 1
	__obf_d14b321e9a2b2243 = 1<<11 - 1
	__obf_1eb40c5ecd81a2a0 = 1<<16 - 1
	__obf_e429a7eaa6ce06a3 = 0xD800
	__obf_f76a17621d63e3f2 = 0xDFFF
	__obf_9d6764936962d565 = '\U0010FFFF'
	__obf_dfdf034f7ead814b = // Maximum valid Unicode code point.
	'\uFFFD'                 // the "error" Rune or "Unicode replacement character"
)

func __obf_9bca19f1764e1c19(__obf_c03fbc322c1a9109 []byte, __obf_464eaebd868bf0fe rune) []byte {
	// Negative values are erroneous. Making it unsigned addresses the problem.
	switch __obf_a051269af3a541bb := uint32(__obf_464eaebd868bf0fe); {
	case __obf_a051269af3a541bb <= __obf_fec31f5058f7df5d:
		__obf_c03fbc322c1a9109 = append(__obf_c03fbc322c1a9109, byte(__obf_464eaebd868bf0fe))
		return __obf_c03fbc322c1a9109
	case __obf_a051269af3a541bb <= __obf_d14b321e9a2b2243:
		__obf_c03fbc322c1a9109 = append(__obf_c03fbc322c1a9109, __obf_6a7093331d3c1ead|byte(__obf_464eaebd868bf0fe>>6))
		__obf_c03fbc322c1a9109 = append(__obf_c03fbc322c1a9109, __obf_c80d46f677b920cc|byte(__obf_464eaebd868bf0fe)&__obf_39db8b683ce0ed16)
		return __obf_c03fbc322c1a9109
	case __obf_a051269af3a541bb > __obf_9d6764936962d565, __obf_e429a7eaa6ce06a3 <= __obf_a051269af3a541bb && __obf_a051269af3a541bb <= __obf_f76a17621d63e3f2:
		__obf_464eaebd868bf0fe = __obf_dfdf034f7ead814b
		fallthrough
	case __obf_a051269af3a541bb <= __obf_1eb40c5ecd81a2a0:
		__obf_c03fbc322c1a9109 = append(__obf_c03fbc322c1a9109, __obf_6ae418366c514de3|byte(__obf_464eaebd868bf0fe>>12))
		__obf_c03fbc322c1a9109 = append(__obf_c03fbc322c1a9109, __obf_c80d46f677b920cc|byte(__obf_464eaebd868bf0fe>>6)&__obf_39db8b683ce0ed16)
		__obf_c03fbc322c1a9109 = append(__obf_c03fbc322c1a9109, __obf_c80d46f677b920cc|byte(__obf_464eaebd868bf0fe)&__obf_39db8b683ce0ed16)
		return __obf_c03fbc322c1a9109
	default:
		__obf_c03fbc322c1a9109 = append(__obf_c03fbc322c1a9109, __obf_a3e07c5094670ee7|byte(__obf_464eaebd868bf0fe>>18))
		__obf_c03fbc322c1a9109 = append(__obf_c03fbc322c1a9109, __obf_c80d46f677b920cc|byte(__obf_464eaebd868bf0fe>>12)&__obf_39db8b683ce0ed16)
		__obf_c03fbc322c1a9109 = append(__obf_c03fbc322c1a9109, __obf_c80d46f677b920cc|byte(__obf_464eaebd868bf0fe>>6)&__obf_39db8b683ce0ed16)
		__obf_c03fbc322c1a9109 = append(__obf_c03fbc322c1a9109, __obf_c80d46f677b920cc|byte(__obf_464eaebd868bf0fe)&__obf_39db8b683ce0ed16)
		return __obf_c03fbc322c1a9109
	}
}
