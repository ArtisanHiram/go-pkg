package __obf_703d23ba520c3295

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
var __obf_f67877e5b4856b01 = [utf8.RuneSelf]bool{
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
var __obf_efc99c40077fe289 = [utf8.RuneSelf]bool{
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

var __obf_008d344daec45853 = "0123456789abcdef"

// WriteStringWithHTMLEscaped write string to stream with html special characters escaped
func (__obf_9a34f62857fb1d1d *Stream) WriteStringWithHTMLEscaped(__obf_61541a8d4689037b string) {
	__obf_19d1c5dc941d7f4b := len(__obf_61541a8d4689037b)
	__obf_9a34f62857fb1d1d.__obf_a065f8e0da5f5952 = append(__obf_9a34f62857fb1d1d.
		// write string, the fast path, without utf8 and escape support
		__obf_a065f8e0da5f5952, '"')
	__obf_b0a5d2bd48690f1d := 0
	for ; __obf_b0a5d2bd48690f1d < __obf_19d1c5dc941d7f4b; __obf_b0a5d2bd48690f1d++ {
		__obf_bd08f5161fff294a := __obf_61541a8d4689037b[__obf_b0a5d2bd48690f1d]
		if __obf_bd08f5161fff294a < utf8.RuneSelf && __obf_f67877e5b4856b01[__obf_bd08f5161fff294a] {
			__obf_9a34f62857fb1d1d.__obf_a065f8e0da5f5952 = append(__obf_9a34f62857fb1d1d.__obf_a065f8e0da5f5952, __obf_bd08f5161fff294a)
		} else {
			break
		}
	}
	if __obf_b0a5d2bd48690f1d == __obf_19d1c5dc941d7f4b {
		__obf_9a34f62857fb1d1d.__obf_a065f8e0da5f5952 = append(__obf_9a34f62857fb1d1d.__obf_a065f8e0da5f5952, '"')
		return
	}
	__obf_20c7028b9d2686d5(__obf_9a34f62857fb1d1d, __obf_b0a5d2bd48690f1d, __obf_61541a8d4689037b, __obf_19d1c5dc941d7f4b)
}

func __obf_20c7028b9d2686d5(__obf_9a34f62857fb1d1d *Stream, __obf_b0a5d2bd48690f1d int, __obf_61541a8d4689037b string, __obf_19d1c5dc941d7f4b int) {
	__obf_bb31006da87bf9e2 := __obf_b0a5d2bd48690f1d
	// for the remaining parts, we process them char by char
	for __obf_b0a5d2bd48690f1d < __obf_19d1c5dc941d7f4b {
		if __obf_85a417ca3a5d43c2 := __obf_61541a8d4689037b[__obf_b0a5d2bd48690f1d]; __obf_85a417ca3a5d43c2 < utf8.RuneSelf {
			if __obf_f67877e5b4856b01[__obf_85a417ca3a5d43c2] {
				__obf_b0a5d2bd48690f1d++
				continue
			}
			if __obf_bb31006da87bf9e2 < __obf_b0a5d2bd48690f1d {
				__obf_9a34f62857fb1d1d.
					WriteRaw(__obf_61541a8d4689037b[__obf_bb31006da87bf9e2:__obf_b0a5d2bd48690f1d])
			}
			switch __obf_85a417ca3a5d43c2 {
			case '\\', '"':
				__obf_9a34f62857fb1d1d.__obf_e3076a8aafc9c7f0('\\', __obf_85a417ca3a5d43c2)
			case '\n':
				__obf_9a34f62857fb1d1d.__obf_e3076a8aafc9c7f0('\\', 'n')
			case '\r':
				__obf_9a34f62857fb1d1d.__obf_e3076a8aafc9c7f0('\\', 'r')
			case '\t':
				__obf_9a34f62857fb1d1d.__obf_e3076a8aafc9c7f0('\\', 't')
			default:
				__obf_9a34f62857fb1d1d.
					// This encodes bytes < 0x20 except for \t, \n and \r.
					// If escapeHTML is set, it also escapes <, >, and &
					// because they can lead to security holes when
					// user-controlled strings are rendered into JSON
					// and served to some browsers.
					WriteRaw(`\u00`)
				__obf_9a34f62857fb1d1d.__obf_e3076a8aafc9c7f0(__obf_008d344daec45853[__obf_85a417ca3a5d43c2>>4], __obf_008d344daec45853[__obf_85a417ca3a5d43c2&0xF])
			}
			__obf_b0a5d2bd48690f1d++
			__obf_bb31006da87bf9e2 = __obf_b0a5d2bd48690f1d
			continue
		}
		__obf_bd08f5161fff294a, __obf_0126ec6b3c37befb := utf8.DecodeRuneInString(__obf_61541a8d4689037b[__obf_b0a5d2bd48690f1d:])
		if __obf_bd08f5161fff294a == utf8.RuneError && __obf_0126ec6b3c37befb == 1 {
			if __obf_bb31006da87bf9e2 < __obf_b0a5d2bd48690f1d {
				__obf_9a34f62857fb1d1d.
					WriteRaw(__obf_61541a8d4689037b[__obf_bb31006da87bf9e2:__obf_b0a5d2bd48690f1d])
			}
			__obf_9a34f62857fb1d1d.
				WriteRaw(`\ufffd`)
			__obf_b0a5d2bd48690f1d++
			__obf_bb31006da87bf9e2 = __obf_b0a5d2bd48690f1d
			continue
		}
		// U+2028 is LINE SEPARATOR.
		// U+2029 is PARAGRAPH SEPARATOR.
		// They are both technically valid characters in JSON strings,
		// but don't work in JSONP, which has to be evaluated as JavaScript,
		// and can lead to security holes there. It is valid JSON to
		// escape them, so we do so unconditionally.
		// See http://timelessrepo.com/json-isnt-a-javascript-subset for discussion.
		if __obf_bd08f5161fff294a == '\u2028' || __obf_bd08f5161fff294a == '\u2029' {
			if __obf_bb31006da87bf9e2 < __obf_b0a5d2bd48690f1d {
				__obf_9a34f62857fb1d1d.
					WriteRaw(__obf_61541a8d4689037b[__obf_bb31006da87bf9e2:__obf_b0a5d2bd48690f1d])
			}
			__obf_9a34f62857fb1d1d.
				WriteRaw(`\u202`)
			__obf_9a34f62857fb1d1d.__obf_f7df2a5224462d19(__obf_008d344daec45853[__obf_bd08f5161fff294a&0xF])
			__obf_b0a5d2bd48690f1d += __obf_0126ec6b3c37befb
			__obf_bb31006da87bf9e2 = __obf_b0a5d2bd48690f1d
			continue
		}
		__obf_b0a5d2bd48690f1d += __obf_0126ec6b3c37befb
	}
	if __obf_bb31006da87bf9e2 < len(__obf_61541a8d4689037b) {
		__obf_9a34f62857fb1d1d.
			WriteRaw(__obf_61541a8d4689037b[__obf_bb31006da87bf9e2:])
	}
	__obf_9a34f62857fb1d1d.__obf_f7df2a5224462d19('"')
}

// WriteString write string to stream without html escape
func (__obf_9a34f62857fb1d1d *Stream) WriteString(__obf_61541a8d4689037b string) {
	__obf_19d1c5dc941d7f4b := len(__obf_61541a8d4689037b)
	__obf_9a34f62857fb1d1d.__obf_a065f8e0da5f5952 = append(__obf_9a34f62857fb1d1d.
		// write string, the fast path, without utf8 and escape support
		__obf_a065f8e0da5f5952, '"')
	__obf_b0a5d2bd48690f1d := 0
	for ; __obf_b0a5d2bd48690f1d < __obf_19d1c5dc941d7f4b; __obf_b0a5d2bd48690f1d++ {
		__obf_bd08f5161fff294a := __obf_61541a8d4689037b[__obf_b0a5d2bd48690f1d]
		if __obf_bd08f5161fff294a > 31 && __obf_bd08f5161fff294a != '"' && __obf_bd08f5161fff294a != '\\' {
			__obf_9a34f62857fb1d1d.__obf_a065f8e0da5f5952 = append(__obf_9a34f62857fb1d1d.__obf_a065f8e0da5f5952, __obf_bd08f5161fff294a)
		} else {
			break
		}
	}
	if __obf_b0a5d2bd48690f1d == __obf_19d1c5dc941d7f4b {
		__obf_9a34f62857fb1d1d.__obf_a065f8e0da5f5952 = append(__obf_9a34f62857fb1d1d.__obf_a065f8e0da5f5952, '"')
		return
	}
	__obf_3b77d8b5b784abbe(__obf_9a34f62857fb1d1d, __obf_b0a5d2bd48690f1d, __obf_61541a8d4689037b, __obf_19d1c5dc941d7f4b)
}

func __obf_3b77d8b5b784abbe(__obf_9a34f62857fb1d1d *Stream, __obf_b0a5d2bd48690f1d int, __obf_61541a8d4689037b string, __obf_19d1c5dc941d7f4b int) {
	__obf_bb31006da87bf9e2 := __obf_b0a5d2bd48690f1d
	// for the remaining parts, we process them char by char
	for __obf_b0a5d2bd48690f1d < __obf_19d1c5dc941d7f4b {
		if __obf_85a417ca3a5d43c2 := __obf_61541a8d4689037b[__obf_b0a5d2bd48690f1d]; __obf_85a417ca3a5d43c2 < utf8.RuneSelf {
			if __obf_efc99c40077fe289[__obf_85a417ca3a5d43c2] {
				__obf_b0a5d2bd48690f1d++
				continue
			}
			if __obf_bb31006da87bf9e2 < __obf_b0a5d2bd48690f1d {
				__obf_9a34f62857fb1d1d.
					WriteRaw(__obf_61541a8d4689037b[__obf_bb31006da87bf9e2:__obf_b0a5d2bd48690f1d])
			}
			switch __obf_85a417ca3a5d43c2 {
			case '\\', '"':
				__obf_9a34f62857fb1d1d.__obf_e3076a8aafc9c7f0('\\', __obf_85a417ca3a5d43c2)
			case '\n':
				__obf_9a34f62857fb1d1d.__obf_e3076a8aafc9c7f0('\\', 'n')
			case '\r':
				__obf_9a34f62857fb1d1d.__obf_e3076a8aafc9c7f0('\\', 'r')
			case '\t':
				__obf_9a34f62857fb1d1d.__obf_e3076a8aafc9c7f0('\\', 't')
			default:
				__obf_9a34f62857fb1d1d.
					// This encodes bytes < 0x20 except for \t, \n and \r.
					// If escapeHTML is set, it also escapes <, >, and &
					// because they can lead to security holes when
					// user-controlled strings are rendered into JSON
					// and served to some browsers.
					WriteRaw(`\u00`)
				__obf_9a34f62857fb1d1d.__obf_e3076a8aafc9c7f0(__obf_008d344daec45853[__obf_85a417ca3a5d43c2>>4], __obf_008d344daec45853[__obf_85a417ca3a5d43c2&0xF])
			}
			__obf_b0a5d2bd48690f1d++
			__obf_bb31006da87bf9e2 = __obf_b0a5d2bd48690f1d
			continue
		}
		__obf_b0a5d2bd48690f1d++
		continue
	}
	if __obf_bb31006da87bf9e2 < len(__obf_61541a8d4689037b) {
		__obf_9a34f62857fb1d1d.
			WriteRaw(__obf_61541a8d4689037b[__obf_bb31006da87bf9e2:])
	}
	__obf_9a34f62857fb1d1d.__obf_f7df2a5224462d19('"')
}
