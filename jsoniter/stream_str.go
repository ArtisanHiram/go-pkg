package __obf_c7b79b12b33d8238

import (
	"unicode/utf8"
)

// htmlSafeSet holds the value true if the ASCII character with the given
// array position can be safely represented inside a JSON string, embedded
// inside of HTML <script> tags, without any additional escaping.
//
// All values are true except for the ASCII control characters (0-31), the
// double quote ("), the backslash character ("\"), HTML opening and closing
// tags ("<" and ">"), and the ampersand ("&").
var __obf_96e8abc9e8fb6ef8 = [utf8.RuneSelf]bool{
	' ':      true,
	'!':      true,
	'"':      false,
	'#':      true,
	'$':      true,
	'%':      true,
	'&':      false,
	'\'':     true,
	'(':      true,
	')':      true,
	'*':      true,
	'+':      true,
	',':      true,
	'-':      true,
	'.':      true,
	'/':      true,
	'0':      true,
	'1':      true,
	'2':      true,
	'3':      true,
	'4':      true,
	'5':      true,
	'6':      true,
	'7':      true,
	'8':      true,
	'9':      true,
	':':      true,
	';':      true,
	'<':      false,
	'=':      true,
	'>':      false,
	'?':      true,
	'@':      true,
	'A':      true,
	'B':      true,
	'C':      true,
	'D':      true,
	'E':      true,
	'F':      true,
	'G':      true,
	'H':      true,
	'I':      true,
	'J':      true,
	'K':      true,
	'L':      true,
	'M':      true,
	'N':      true,
	'O':      true,
	'P':      true,
	'Q':      true,
	'R':      true,
	'S':      true,
	'T':      true,
	'U':      true,
	'V':      true,
	'W':      true,
	'X':      true,
	'Y':      true,
	'Z':      true,
	'[':      true,
	'\\':     false,
	']':      true,
	'^':      true,
	'_':      true,
	'`':      true,
	'a':      true,
	'b':      true,
	'c':      true,
	'd':      true,
	'e':      true,
	'f':      true,
	'g':      true,
	'h':      true,
	'i':      true,
	'j':      true,
	'k':      true,
	'l':      true,
	'm':      true,
	'n':      true,
	'o':      true,
	'p':      true,
	'q':      true,
	'r':      true,
	's':      true,
	't':      true,
	'u':      true,
	'v':      true,
	'w':      true,
	'x':      true,
	'y':      true,
	'z':      true,
	'{':      true,
	'|':      true,
	'}':      true,
	'~':      true,
	'\u007f': true,
}

// safeSet holds the value true if the ASCII character with the given array
// position can be represented inside a JSON string without any further
// escaping.
//
// All values are true except for the ASCII control characters (0-31), the
// double quote ("), and the backslash character ("\").
var __obf_b7d1921dad40e65e = [utf8.RuneSelf]bool{
	' ':      true,
	'!':      true,
	'"':      false,
	'#':      true,
	'$':      true,
	'%':      true,
	'&':      true,
	'\'':     true,
	'(':      true,
	')':      true,
	'*':      true,
	'+':      true,
	',':      true,
	'-':      true,
	'.':      true,
	'/':      true,
	'0':      true,
	'1':      true,
	'2':      true,
	'3':      true,
	'4':      true,
	'5':      true,
	'6':      true,
	'7':      true,
	'8':      true,
	'9':      true,
	':':      true,
	';':      true,
	'<':      true,
	'=':      true,
	'>':      true,
	'?':      true,
	'@':      true,
	'A':      true,
	'B':      true,
	'C':      true,
	'D':      true,
	'E':      true,
	'F':      true,
	'G':      true,
	'H':      true,
	'I':      true,
	'J':      true,
	'K':      true,
	'L':      true,
	'M':      true,
	'N':      true,
	'O':      true,
	'P':      true,
	'Q':      true,
	'R':      true,
	'S':      true,
	'T':      true,
	'U':      true,
	'V':      true,
	'W':      true,
	'X':      true,
	'Y':      true,
	'Z':      true,
	'[':      true,
	'\\':     false,
	']':      true,
	'^':      true,
	'_':      true,
	'`':      true,
	'a':      true,
	'b':      true,
	'c':      true,
	'd':      true,
	'e':      true,
	'f':      true,
	'g':      true,
	'h':      true,
	'i':      true,
	'j':      true,
	'k':      true,
	'l':      true,
	'm':      true,
	'n':      true,
	'o':      true,
	'p':      true,
	'q':      true,
	'r':      true,
	's':      true,
	't':      true,
	'u':      true,
	'v':      true,
	'w':      true,
	'x':      true,
	'y':      true,
	'z':      true,
	'{':      true,
	'|':      true,
	'}':      true,
	'~':      true,
	'\u007f': true,
}

var __obf_0295be282ee06228 = "0123456789abcdef"

// WriteStringWithHTMLEscaped write string to stream with html special characters escaped
func (__obf_d8f50bcfe87d8b47 *Stream) WriteStringWithHTMLEscaped(__obf_057e9641d10a1246 string) {
	__obf_31ccee095dd054a8 := len(__obf_057e9641d10a1246)
	__obf_d8f50bcfe87d8b47.__obf_684d738bc3ac851a = append(__obf_d8f50bcfe87d8b47.
		// write string, the fast path, without utf8 and escape support
		__obf_684d738bc3ac851a, '"')
	__obf_a051269af3a541bb := 0
	for ; __obf_a051269af3a541bb < __obf_31ccee095dd054a8; __obf_a051269af3a541bb++ {
		__obf_12577c03a12f0d2c := __obf_057e9641d10a1246[__obf_a051269af3a541bb]
		if __obf_12577c03a12f0d2c < utf8.RuneSelf && __obf_96e8abc9e8fb6ef8[__obf_12577c03a12f0d2c] {
			__obf_d8f50bcfe87d8b47.__obf_684d738bc3ac851a = append(__obf_d8f50bcfe87d8b47.__obf_684d738bc3ac851a, __obf_12577c03a12f0d2c)
		} else {
			break
		}
	}
	if __obf_a051269af3a541bb == __obf_31ccee095dd054a8 {
		__obf_d8f50bcfe87d8b47.__obf_684d738bc3ac851a = append(__obf_d8f50bcfe87d8b47.__obf_684d738bc3ac851a, '"')
		return
	}
	__obf_70fbb381f8712f64(__obf_d8f50bcfe87d8b47, __obf_a051269af3a541bb, __obf_057e9641d10a1246, __obf_31ccee095dd054a8)
}

func __obf_70fbb381f8712f64(__obf_d8f50bcfe87d8b47 *Stream, __obf_a051269af3a541bb int, __obf_057e9641d10a1246 string, __obf_31ccee095dd054a8 int) {
	__obf_28e7d8a603535072 := __obf_a051269af3a541bb
	// for the remaining parts, we process them char by char
	for __obf_a051269af3a541bb < __obf_31ccee095dd054a8 {
		if __obf_fa8a3db302183a92 := __obf_057e9641d10a1246[__obf_a051269af3a541bb]; __obf_fa8a3db302183a92 < utf8.RuneSelf {
			if __obf_96e8abc9e8fb6ef8[__obf_fa8a3db302183a92] {
				__obf_a051269af3a541bb++
				continue
			}
			if __obf_28e7d8a603535072 < __obf_a051269af3a541bb {
				__obf_d8f50bcfe87d8b47.
					WriteRaw(__obf_057e9641d10a1246[__obf_28e7d8a603535072:__obf_a051269af3a541bb])
			}
			switch __obf_fa8a3db302183a92 {
			case '\\', '"':
				__obf_d8f50bcfe87d8b47.__obf_99147f1a9b9d5ffc('\\', __obf_fa8a3db302183a92)
			case '\n':
				__obf_d8f50bcfe87d8b47.__obf_99147f1a9b9d5ffc('\\', 'n')
			case '\r':
				__obf_d8f50bcfe87d8b47.__obf_99147f1a9b9d5ffc('\\', 'r')
			case '\t':
				__obf_d8f50bcfe87d8b47.__obf_99147f1a9b9d5ffc('\\', 't')
			default:
				__obf_d8f50bcfe87d8b47.
					// This encodes bytes < 0x20 except for \t, \n and \r.
					// If escapeHTML is set, it also escapes <, >, and &
					// because they can lead to security holes when
					// user-controlled strings are rendered into JSON
					// and served to some browsers.
					WriteRaw(`\u00`)
				__obf_d8f50bcfe87d8b47.__obf_99147f1a9b9d5ffc(__obf_0295be282ee06228[__obf_fa8a3db302183a92>>4], __obf_0295be282ee06228[__obf_fa8a3db302183a92&0xF])
			}
			__obf_a051269af3a541bb++
			__obf_28e7d8a603535072 = __obf_a051269af3a541bb
			continue
		}
		__obf_12577c03a12f0d2c, __obf_86901472dc2ab8e1 := utf8.DecodeRuneInString(__obf_057e9641d10a1246[__obf_a051269af3a541bb:])
		if __obf_12577c03a12f0d2c == utf8.RuneError && __obf_86901472dc2ab8e1 == 1 {
			if __obf_28e7d8a603535072 < __obf_a051269af3a541bb {
				__obf_d8f50bcfe87d8b47.
					WriteRaw(__obf_057e9641d10a1246[__obf_28e7d8a603535072:__obf_a051269af3a541bb])
			}
			__obf_d8f50bcfe87d8b47.
				WriteRaw(`\ufffd`)
			__obf_a051269af3a541bb++
			__obf_28e7d8a603535072 = __obf_a051269af3a541bb
			continue
		}
		// U+2028 is LINE SEPARATOR.
		// U+2029 is PARAGRAPH SEPARATOR.
		// They are both technically valid characters in JSON strings,
		// but don't work in JSONP, which has to be evaluated as JavaScript,
		// and can lead to security holes there. It is valid JSON to
		// escape them, so we do so unconditionally.
		// See http://timelessrepo.com/json-isnt-a-javascript-subset for discussion.
		if __obf_12577c03a12f0d2c == '\u2028' || __obf_12577c03a12f0d2c == '\u2029' {
			if __obf_28e7d8a603535072 < __obf_a051269af3a541bb {
				__obf_d8f50bcfe87d8b47.
					WriteRaw(__obf_057e9641d10a1246[__obf_28e7d8a603535072:__obf_a051269af3a541bb])
			}
			__obf_d8f50bcfe87d8b47.
				WriteRaw(`\u202`)
			__obf_d8f50bcfe87d8b47.__obf_7340077d41996822(__obf_0295be282ee06228[__obf_12577c03a12f0d2c&0xF])
			__obf_a051269af3a541bb += __obf_86901472dc2ab8e1
			__obf_28e7d8a603535072 = __obf_a051269af3a541bb
			continue
		}
		__obf_a051269af3a541bb += __obf_86901472dc2ab8e1
	}
	if __obf_28e7d8a603535072 < len(__obf_057e9641d10a1246) {
		__obf_d8f50bcfe87d8b47.
			WriteRaw(__obf_057e9641d10a1246[__obf_28e7d8a603535072:])
	}
	__obf_d8f50bcfe87d8b47.__obf_7340077d41996822('"')
}

// WriteString write string to stream without html escape
func (__obf_d8f50bcfe87d8b47 *Stream) WriteString(__obf_057e9641d10a1246 string) {
	__obf_31ccee095dd054a8 := len(__obf_057e9641d10a1246)
	__obf_d8f50bcfe87d8b47.__obf_684d738bc3ac851a = append(__obf_d8f50bcfe87d8b47.
		// write string, the fast path, without utf8 and escape support
		__obf_684d738bc3ac851a, '"')
	__obf_a051269af3a541bb := 0
	for ; __obf_a051269af3a541bb < __obf_31ccee095dd054a8; __obf_a051269af3a541bb++ {
		__obf_12577c03a12f0d2c := __obf_057e9641d10a1246[__obf_a051269af3a541bb]
		if __obf_12577c03a12f0d2c > 31 && __obf_12577c03a12f0d2c != '"' && __obf_12577c03a12f0d2c != '\\' {
			__obf_d8f50bcfe87d8b47.__obf_684d738bc3ac851a = append(__obf_d8f50bcfe87d8b47.__obf_684d738bc3ac851a, __obf_12577c03a12f0d2c)
		} else {
			break
		}
	}
	if __obf_a051269af3a541bb == __obf_31ccee095dd054a8 {
		__obf_d8f50bcfe87d8b47.__obf_684d738bc3ac851a = append(__obf_d8f50bcfe87d8b47.__obf_684d738bc3ac851a, '"')
		return
	}
	__obf_52035bf0e0d0c96a(__obf_d8f50bcfe87d8b47, __obf_a051269af3a541bb, __obf_057e9641d10a1246, __obf_31ccee095dd054a8)
}

func __obf_52035bf0e0d0c96a(__obf_d8f50bcfe87d8b47 *Stream, __obf_a051269af3a541bb int, __obf_057e9641d10a1246 string, __obf_31ccee095dd054a8 int) {
	__obf_28e7d8a603535072 := __obf_a051269af3a541bb
	// for the remaining parts, we process them char by char
	for __obf_a051269af3a541bb < __obf_31ccee095dd054a8 {
		if __obf_fa8a3db302183a92 := __obf_057e9641d10a1246[__obf_a051269af3a541bb]; __obf_fa8a3db302183a92 < utf8.RuneSelf {
			if __obf_b7d1921dad40e65e[__obf_fa8a3db302183a92] {
				__obf_a051269af3a541bb++
				continue
			}
			if __obf_28e7d8a603535072 < __obf_a051269af3a541bb {
				__obf_d8f50bcfe87d8b47.
					WriteRaw(__obf_057e9641d10a1246[__obf_28e7d8a603535072:__obf_a051269af3a541bb])
			}
			switch __obf_fa8a3db302183a92 {
			case '\\', '"':
				__obf_d8f50bcfe87d8b47.__obf_99147f1a9b9d5ffc('\\', __obf_fa8a3db302183a92)
			case '\n':
				__obf_d8f50bcfe87d8b47.__obf_99147f1a9b9d5ffc('\\', 'n')
			case '\r':
				__obf_d8f50bcfe87d8b47.__obf_99147f1a9b9d5ffc('\\', 'r')
			case '\t':
				__obf_d8f50bcfe87d8b47.__obf_99147f1a9b9d5ffc('\\', 't')
			default:
				__obf_d8f50bcfe87d8b47.
					// This encodes bytes < 0x20 except for \t, \n and \r.
					// If escapeHTML is set, it also escapes <, >, and &
					// because they can lead to security holes when
					// user-controlled strings are rendered into JSON
					// and served to some browsers.
					WriteRaw(`\u00`)
				__obf_d8f50bcfe87d8b47.__obf_99147f1a9b9d5ffc(__obf_0295be282ee06228[__obf_fa8a3db302183a92>>4], __obf_0295be282ee06228[__obf_fa8a3db302183a92&0xF])
			}
			__obf_a051269af3a541bb++
			__obf_28e7d8a603535072 = __obf_a051269af3a541bb
			continue
		}
		__obf_a051269af3a541bb++
		continue
	}
	if __obf_28e7d8a603535072 < len(__obf_057e9641d10a1246) {
		__obf_d8f50bcfe87d8b47.
			WriteRaw(__obf_057e9641d10a1246[__obf_28e7d8a603535072:])
	}
	__obf_d8f50bcfe87d8b47.__obf_7340077d41996822('"')
}
