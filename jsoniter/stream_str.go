package __obf_5b802ce8d9ba56d6

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
var __obf_cac39019abde8d68 = [utf8.RuneSelf]bool{
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
var __obf_d99476afc0910d4f = [utf8.RuneSelf]bool{
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

var __obf_5b60e46c496b0f98 = "0123456789abcdef"

// WriteStringWithHTMLEscaped write string to stream with html special characters escaped
func (__obf_00fc491caa967a74 *Stream) WriteStringWithHTMLEscaped(__obf_5ba76a0156a3a826 string) {
	__obf_324a62039bf9903d := len(__obf_5ba76a0156a3a826)
	__obf_00fc491caa967a74.__obf_9fc06d9180f0daca = append(__obf_00fc491caa967a74.
		// write string, the fast path, without utf8 and escape support
		__obf_9fc06d9180f0daca, '"')
	__obf_2deec7c38ffb6ae3 := 0
	for ; __obf_2deec7c38ffb6ae3 < __obf_324a62039bf9903d; __obf_2deec7c38ffb6ae3++ {
		__obf_dab9baaadfa7c8c2 := __obf_5ba76a0156a3a826[__obf_2deec7c38ffb6ae3]
		if __obf_dab9baaadfa7c8c2 < utf8.RuneSelf && __obf_cac39019abde8d68[__obf_dab9baaadfa7c8c2] {
			__obf_00fc491caa967a74.__obf_9fc06d9180f0daca = append(__obf_00fc491caa967a74.__obf_9fc06d9180f0daca, __obf_dab9baaadfa7c8c2)
		} else {
			break
		}
	}
	if __obf_2deec7c38ffb6ae3 == __obf_324a62039bf9903d {
		__obf_00fc491caa967a74.__obf_9fc06d9180f0daca = append(__obf_00fc491caa967a74.__obf_9fc06d9180f0daca, '"')
		return
	}
	__obf_8a04e5a5b073fae3(__obf_00fc491caa967a74, __obf_2deec7c38ffb6ae3, __obf_5ba76a0156a3a826, __obf_324a62039bf9903d)
}

func __obf_8a04e5a5b073fae3(__obf_00fc491caa967a74 *Stream, __obf_2deec7c38ffb6ae3 int, __obf_5ba76a0156a3a826 string, __obf_324a62039bf9903d int) {
	__obf_4773198b4b8f9524 := __obf_2deec7c38ffb6ae3
	// for the remaining parts, we process them char by char
	for __obf_2deec7c38ffb6ae3 < __obf_324a62039bf9903d {
		if __obf_1c7157183bc604f5 := __obf_5ba76a0156a3a826[__obf_2deec7c38ffb6ae3]; __obf_1c7157183bc604f5 < utf8.RuneSelf {
			if __obf_cac39019abde8d68[__obf_1c7157183bc604f5] {
				__obf_2deec7c38ffb6ae3++
				continue
			}
			if __obf_4773198b4b8f9524 < __obf_2deec7c38ffb6ae3 {
				__obf_00fc491caa967a74.
					WriteRaw(__obf_5ba76a0156a3a826[__obf_4773198b4b8f9524:__obf_2deec7c38ffb6ae3])
			}
			switch __obf_1c7157183bc604f5 {
			case '\\', '"':
				__obf_00fc491caa967a74.__obf_dd59782cf69113a3('\\', __obf_1c7157183bc604f5)
			case '\n':
				__obf_00fc491caa967a74.__obf_dd59782cf69113a3('\\', 'n')
			case '\r':
				__obf_00fc491caa967a74.__obf_dd59782cf69113a3('\\', 'r')
			case '\t':
				__obf_00fc491caa967a74.__obf_dd59782cf69113a3('\\', 't')
			default:
				__obf_00fc491caa967a74.
					// This encodes bytes < 0x20 except for \t, \n and \r.
					// If escapeHTML is set, it also escapes <, >, and &
					// because they can lead to security holes when
					// user-controlled strings are rendered into JSON
					// and served to some browsers.
					WriteRaw(`\u00`)
				__obf_00fc491caa967a74.__obf_dd59782cf69113a3(__obf_5b60e46c496b0f98[__obf_1c7157183bc604f5>>4], __obf_5b60e46c496b0f98[__obf_1c7157183bc604f5&0xF])
			}
			__obf_2deec7c38ffb6ae3++
			__obf_4773198b4b8f9524 = __obf_2deec7c38ffb6ae3
			continue
		}
		__obf_dab9baaadfa7c8c2, __obf_aedccbde4481459c := utf8.DecodeRuneInString(__obf_5ba76a0156a3a826[__obf_2deec7c38ffb6ae3:])
		if __obf_dab9baaadfa7c8c2 == utf8.RuneError && __obf_aedccbde4481459c == 1 {
			if __obf_4773198b4b8f9524 < __obf_2deec7c38ffb6ae3 {
				__obf_00fc491caa967a74.
					WriteRaw(__obf_5ba76a0156a3a826[__obf_4773198b4b8f9524:__obf_2deec7c38ffb6ae3])
			}
			__obf_00fc491caa967a74.
				WriteRaw(`\ufffd`)
			__obf_2deec7c38ffb6ae3++
			__obf_4773198b4b8f9524 = __obf_2deec7c38ffb6ae3
			continue
		}
		// U+2028 is LINE SEPARATOR.
		// U+2029 is PARAGRAPH SEPARATOR.
		// They are both technically valid characters in JSON strings,
		// but don't work in JSONP, which has to be evaluated as JavaScript,
		// and can lead to security holes there. It is valid JSON to
		// escape them, so we do so unconditionally.
		// See http://timelessrepo.com/json-isnt-a-javascript-subset for discussion.
		if __obf_dab9baaadfa7c8c2 == '\u2028' || __obf_dab9baaadfa7c8c2 == '\u2029' {
			if __obf_4773198b4b8f9524 < __obf_2deec7c38ffb6ae3 {
				__obf_00fc491caa967a74.
					WriteRaw(__obf_5ba76a0156a3a826[__obf_4773198b4b8f9524:__obf_2deec7c38ffb6ae3])
			}
			__obf_00fc491caa967a74.
				WriteRaw(`\u202`)
			__obf_00fc491caa967a74.__obf_487ee8035c7dd8f6(__obf_5b60e46c496b0f98[__obf_dab9baaadfa7c8c2&0xF])
			__obf_2deec7c38ffb6ae3 += __obf_aedccbde4481459c
			__obf_4773198b4b8f9524 = __obf_2deec7c38ffb6ae3
			continue
		}
		__obf_2deec7c38ffb6ae3 += __obf_aedccbde4481459c
	}
	if __obf_4773198b4b8f9524 < len(__obf_5ba76a0156a3a826) {
		__obf_00fc491caa967a74.
			WriteRaw(__obf_5ba76a0156a3a826[__obf_4773198b4b8f9524:])
	}
	__obf_00fc491caa967a74.__obf_487ee8035c7dd8f6('"')
}

// WriteString write string to stream without html escape
func (__obf_00fc491caa967a74 *Stream) WriteString(__obf_5ba76a0156a3a826 string) {
	__obf_324a62039bf9903d := len(__obf_5ba76a0156a3a826)
	__obf_00fc491caa967a74.__obf_9fc06d9180f0daca = append(__obf_00fc491caa967a74.
		// write string, the fast path, without utf8 and escape support
		__obf_9fc06d9180f0daca, '"')
	__obf_2deec7c38ffb6ae3 := 0
	for ; __obf_2deec7c38ffb6ae3 < __obf_324a62039bf9903d; __obf_2deec7c38ffb6ae3++ {
		__obf_dab9baaadfa7c8c2 := __obf_5ba76a0156a3a826[__obf_2deec7c38ffb6ae3]
		if __obf_dab9baaadfa7c8c2 > 31 && __obf_dab9baaadfa7c8c2 != '"' && __obf_dab9baaadfa7c8c2 != '\\' {
			__obf_00fc491caa967a74.__obf_9fc06d9180f0daca = append(__obf_00fc491caa967a74.__obf_9fc06d9180f0daca, __obf_dab9baaadfa7c8c2)
		} else {
			break
		}
	}
	if __obf_2deec7c38ffb6ae3 == __obf_324a62039bf9903d {
		__obf_00fc491caa967a74.__obf_9fc06d9180f0daca = append(__obf_00fc491caa967a74.__obf_9fc06d9180f0daca, '"')
		return
	}
	__obf_10917a340a1e4a93(__obf_00fc491caa967a74, __obf_2deec7c38ffb6ae3, __obf_5ba76a0156a3a826, __obf_324a62039bf9903d)
}

func __obf_10917a340a1e4a93(__obf_00fc491caa967a74 *Stream, __obf_2deec7c38ffb6ae3 int, __obf_5ba76a0156a3a826 string, __obf_324a62039bf9903d int) {
	__obf_4773198b4b8f9524 := __obf_2deec7c38ffb6ae3
	// for the remaining parts, we process them char by char
	for __obf_2deec7c38ffb6ae3 < __obf_324a62039bf9903d {
		if __obf_1c7157183bc604f5 := __obf_5ba76a0156a3a826[__obf_2deec7c38ffb6ae3]; __obf_1c7157183bc604f5 < utf8.RuneSelf {
			if __obf_d99476afc0910d4f[__obf_1c7157183bc604f5] {
				__obf_2deec7c38ffb6ae3++
				continue
			}
			if __obf_4773198b4b8f9524 < __obf_2deec7c38ffb6ae3 {
				__obf_00fc491caa967a74.
					WriteRaw(__obf_5ba76a0156a3a826[__obf_4773198b4b8f9524:__obf_2deec7c38ffb6ae3])
			}
			switch __obf_1c7157183bc604f5 {
			case '\\', '"':
				__obf_00fc491caa967a74.__obf_dd59782cf69113a3('\\', __obf_1c7157183bc604f5)
			case '\n':
				__obf_00fc491caa967a74.__obf_dd59782cf69113a3('\\', 'n')
			case '\r':
				__obf_00fc491caa967a74.__obf_dd59782cf69113a3('\\', 'r')
			case '\t':
				__obf_00fc491caa967a74.__obf_dd59782cf69113a3('\\', 't')
			default:
				__obf_00fc491caa967a74.
					// This encodes bytes < 0x20 except for \t, \n and \r.
					// If escapeHTML is set, it also escapes <, >, and &
					// because they can lead to security holes when
					// user-controlled strings are rendered into JSON
					// and served to some browsers.
					WriteRaw(`\u00`)
				__obf_00fc491caa967a74.__obf_dd59782cf69113a3(__obf_5b60e46c496b0f98[__obf_1c7157183bc604f5>>4], __obf_5b60e46c496b0f98[__obf_1c7157183bc604f5&0xF])
			}
			__obf_2deec7c38ffb6ae3++
			__obf_4773198b4b8f9524 = __obf_2deec7c38ffb6ae3
			continue
		}
		__obf_2deec7c38ffb6ae3++
		continue
	}
	if __obf_4773198b4b8f9524 < len(__obf_5ba76a0156a3a826) {
		__obf_00fc491caa967a74.
			WriteRaw(__obf_5ba76a0156a3a826[__obf_4773198b4b8f9524:])
	}
	__obf_00fc491caa967a74.__obf_487ee8035c7dd8f6('"')
}
