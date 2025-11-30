package __obf_030d39f7a8de96c6

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
var __obf_b01bba1694eaae6f = [utf8.RuneSelf]bool{
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
var __obf_a555af9a7c7d8eea = [utf8.RuneSelf]bool{
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

var __obf_09bb009dbfa0a3d4 = "0123456789abcdef"

// WriteStringWithHTMLEscaped write string to stream with html special characters escaped
func (__obf_8a2c51fe22775655 *Stream) WriteStringWithHTMLEscaped(__obf_1d7ed4e74380ec76 string) {
	__obf_51f269c3f31c8795 := len(__obf_1d7ed4e74380ec76)
	__obf_8a2c51fe22775655.__obf_0b1656d7ef4bd9df = append(__obf_8a2c51fe22775655.
		// write string, the fast path, without utf8 and escape support
		__obf_0b1656d7ef4bd9df, '"')
	__obf_82c6e05b9d226c58 := 0
	for ; __obf_82c6e05b9d226c58 < __obf_51f269c3f31c8795; __obf_82c6e05b9d226c58++ {
		__obf_24309b3b0ff9ba22 := __obf_1d7ed4e74380ec76[__obf_82c6e05b9d226c58]
		if __obf_24309b3b0ff9ba22 < utf8.RuneSelf && __obf_b01bba1694eaae6f[__obf_24309b3b0ff9ba22] {
			__obf_8a2c51fe22775655.__obf_0b1656d7ef4bd9df = append(__obf_8a2c51fe22775655.__obf_0b1656d7ef4bd9df, __obf_24309b3b0ff9ba22)
		} else {
			break
		}
	}
	if __obf_82c6e05b9d226c58 == __obf_51f269c3f31c8795 {
		__obf_8a2c51fe22775655.__obf_0b1656d7ef4bd9df = append(__obf_8a2c51fe22775655.__obf_0b1656d7ef4bd9df, '"')
		return
	}
	__obf_37db73f8c1f01c54(__obf_8a2c51fe22775655, __obf_82c6e05b9d226c58, __obf_1d7ed4e74380ec76, __obf_51f269c3f31c8795)
}

func __obf_37db73f8c1f01c54(__obf_8a2c51fe22775655 *Stream, __obf_82c6e05b9d226c58 int, __obf_1d7ed4e74380ec76 string, __obf_51f269c3f31c8795 int) {
	__obf_447a3399d55be0d6 := __obf_82c6e05b9d226c58
	// for the remaining parts, we process them char by char
	for __obf_82c6e05b9d226c58 < __obf_51f269c3f31c8795 {
		if __obf_ea107e11b66371dd := __obf_1d7ed4e74380ec76[__obf_82c6e05b9d226c58]; __obf_ea107e11b66371dd < utf8.RuneSelf {
			if __obf_b01bba1694eaae6f[__obf_ea107e11b66371dd] {
				__obf_82c6e05b9d226c58++
				continue
			}
			if __obf_447a3399d55be0d6 < __obf_82c6e05b9d226c58 {
				__obf_8a2c51fe22775655.
					WriteRaw(__obf_1d7ed4e74380ec76[__obf_447a3399d55be0d6:__obf_82c6e05b9d226c58])
			}
			switch __obf_ea107e11b66371dd {
			case '\\', '"':
				__obf_8a2c51fe22775655.__obf_3635bad429be5857('\\', __obf_ea107e11b66371dd)
			case '\n':
				__obf_8a2c51fe22775655.__obf_3635bad429be5857('\\', 'n')
			case '\r':
				__obf_8a2c51fe22775655.__obf_3635bad429be5857('\\', 'r')
			case '\t':
				__obf_8a2c51fe22775655.__obf_3635bad429be5857('\\', 't')
			default:
				__obf_8a2c51fe22775655.
					// This encodes bytes < 0x20 except for \t, \n and \r.
					// If escapeHTML is set, it also escapes <, >, and &
					// because they can lead to security holes when
					// user-controlled strings are rendered into JSON
					// and served to some browsers.
					WriteRaw(`\u00`)
				__obf_8a2c51fe22775655.__obf_3635bad429be5857(__obf_09bb009dbfa0a3d4[__obf_ea107e11b66371dd>>4], __obf_09bb009dbfa0a3d4[__obf_ea107e11b66371dd&0xF])
			}
			__obf_82c6e05b9d226c58++
			__obf_447a3399d55be0d6 = __obf_82c6e05b9d226c58
			continue
		}
		__obf_24309b3b0ff9ba22, __obf_452ca3698eb735a5 := utf8.DecodeRuneInString(__obf_1d7ed4e74380ec76[__obf_82c6e05b9d226c58:])
		if __obf_24309b3b0ff9ba22 == utf8.RuneError && __obf_452ca3698eb735a5 == 1 {
			if __obf_447a3399d55be0d6 < __obf_82c6e05b9d226c58 {
				__obf_8a2c51fe22775655.
					WriteRaw(__obf_1d7ed4e74380ec76[__obf_447a3399d55be0d6:__obf_82c6e05b9d226c58])
			}
			__obf_8a2c51fe22775655.
				WriteRaw(`\ufffd`)
			__obf_82c6e05b9d226c58++
			__obf_447a3399d55be0d6 = __obf_82c6e05b9d226c58
			continue
		}
		// U+2028 is LINE SEPARATOR.
		// U+2029 is PARAGRAPH SEPARATOR.
		// They are both technically valid characters in JSON strings,
		// but don't work in JSONP, which has to be evaluated as JavaScript,
		// and can lead to security holes there. It is valid JSON to
		// escape them, so we do so unconditionally.
		// See http://timelessrepo.com/json-isnt-a-javascript-subset for discussion.
		if __obf_24309b3b0ff9ba22 == '\u2028' || __obf_24309b3b0ff9ba22 == '\u2029' {
			if __obf_447a3399d55be0d6 < __obf_82c6e05b9d226c58 {
				__obf_8a2c51fe22775655.
					WriteRaw(__obf_1d7ed4e74380ec76[__obf_447a3399d55be0d6:__obf_82c6e05b9d226c58])
			}
			__obf_8a2c51fe22775655.
				WriteRaw(`\u202`)
			__obf_8a2c51fe22775655.__obf_41130daa346c0e53(__obf_09bb009dbfa0a3d4[__obf_24309b3b0ff9ba22&0xF])
			__obf_82c6e05b9d226c58 += __obf_452ca3698eb735a5
			__obf_447a3399d55be0d6 = __obf_82c6e05b9d226c58
			continue
		}
		__obf_82c6e05b9d226c58 += __obf_452ca3698eb735a5
	}
	if __obf_447a3399d55be0d6 < len(__obf_1d7ed4e74380ec76) {
		__obf_8a2c51fe22775655.
			WriteRaw(__obf_1d7ed4e74380ec76[__obf_447a3399d55be0d6:])
	}
	__obf_8a2c51fe22775655.__obf_41130daa346c0e53('"')
}

// WriteString write string to stream without html escape
func (__obf_8a2c51fe22775655 *Stream) WriteString(__obf_1d7ed4e74380ec76 string) {
	__obf_51f269c3f31c8795 := len(__obf_1d7ed4e74380ec76)
	__obf_8a2c51fe22775655.__obf_0b1656d7ef4bd9df = append(__obf_8a2c51fe22775655.
		// write string, the fast path, without utf8 and escape support
		__obf_0b1656d7ef4bd9df, '"')
	__obf_82c6e05b9d226c58 := 0
	for ; __obf_82c6e05b9d226c58 < __obf_51f269c3f31c8795; __obf_82c6e05b9d226c58++ {
		__obf_24309b3b0ff9ba22 := __obf_1d7ed4e74380ec76[__obf_82c6e05b9d226c58]
		if __obf_24309b3b0ff9ba22 > 31 && __obf_24309b3b0ff9ba22 != '"' && __obf_24309b3b0ff9ba22 != '\\' {
			__obf_8a2c51fe22775655.__obf_0b1656d7ef4bd9df = append(__obf_8a2c51fe22775655.__obf_0b1656d7ef4bd9df, __obf_24309b3b0ff9ba22)
		} else {
			break
		}
	}
	if __obf_82c6e05b9d226c58 == __obf_51f269c3f31c8795 {
		__obf_8a2c51fe22775655.__obf_0b1656d7ef4bd9df = append(__obf_8a2c51fe22775655.__obf_0b1656d7ef4bd9df, '"')
		return
	}
	__obf_d4d730ff20d6807b(__obf_8a2c51fe22775655, __obf_82c6e05b9d226c58, __obf_1d7ed4e74380ec76, __obf_51f269c3f31c8795)
}

func __obf_d4d730ff20d6807b(__obf_8a2c51fe22775655 *Stream, __obf_82c6e05b9d226c58 int, __obf_1d7ed4e74380ec76 string, __obf_51f269c3f31c8795 int) {
	__obf_447a3399d55be0d6 := __obf_82c6e05b9d226c58
	// for the remaining parts, we process them char by char
	for __obf_82c6e05b9d226c58 < __obf_51f269c3f31c8795 {
		if __obf_ea107e11b66371dd := __obf_1d7ed4e74380ec76[__obf_82c6e05b9d226c58]; __obf_ea107e11b66371dd < utf8.RuneSelf {
			if __obf_a555af9a7c7d8eea[__obf_ea107e11b66371dd] {
				__obf_82c6e05b9d226c58++
				continue
			}
			if __obf_447a3399d55be0d6 < __obf_82c6e05b9d226c58 {
				__obf_8a2c51fe22775655.
					WriteRaw(__obf_1d7ed4e74380ec76[__obf_447a3399d55be0d6:__obf_82c6e05b9d226c58])
			}
			switch __obf_ea107e11b66371dd {
			case '\\', '"':
				__obf_8a2c51fe22775655.__obf_3635bad429be5857('\\', __obf_ea107e11b66371dd)
			case '\n':
				__obf_8a2c51fe22775655.__obf_3635bad429be5857('\\', 'n')
			case '\r':
				__obf_8a2c51fe22775655.__obf_3635bad429be5857('\\', 'r')
			case '\t':
				__obf_8a2c51fe22775655.__obf_3635bad429be5857('\\', 't')
			default:
				__obf_8a2c51fe22775655.
					// This encodes bytes < 0x20 except for \t, \n and \r.
					// If escapeHTML is set, it also escapes <, >, and &
					// because they can lead to security holes when
					// user-controlled strings are rendered into JSON
					// and served to some browsers.
					WriteRaw(`\u00`)
				__obf_8a2c51fe22775655.__obf_3635bad429be5857(__obf_09bb009dbfa0a3d4[__obf_ea107e11b66371dd>>4], __obf_09bb009dbfa0a3d4[__obf_ea107e11b66371dd&0xF])
			}
			__obf_82c6e05b9d226c58++
			__obf_447a3399d55be0d6 = __obf_82c6e05b9d226c58
			continue
		}
		__obf_82c6e05b9d226c58++
		continue
	}
	if __obf_447a3399d55be0d6 < len(__obf_1d7ed4e74380ec76) {
		__obf_8a2c51fe22775655.
			WriteRaw(__obf_1d7ed4e74380ec76[__obf_447a3399d55be0d6:])
	}
	__obf_8a2c51fe22775655.__obf_41130daa346c0e53('"')
}
