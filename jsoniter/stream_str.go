package __obf_91620b895eeff9ed

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
var __obf_eb32a2c54983008b = [utf8.RuneSelf]bool{
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
var __obf_041351d57e164463 = [utf8.RuneSelf]bool{
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

var __obf_76ce97e7af7d5659 = "0123456789abcdef"

// WriteStringWithHTMLEscaped write string to stream with html special characters escaped
func (__obf_850a7457bb739a32 *Stream) WriteStringWithHTMLEscaped(__obf_6f3495a136f456ab string) {
	__obf_ed1950d3b33cfd95 := len(__obf_6f3495a136f456ab)
	__obf_850a7457bb739a32.__obf_184433571fa55237 = append(__obf_850a7457bb739a32.
		// write string, the fast path, without utf8 and escape support
		__obf_184433571fa55237, '"')
	__obf_5aa5c8829b97f182 := 0
	for ; __obf_5aa5c8829b97f182 < __obf_ed1950d3b33cfd95; __obf_5aa5c8829b97f182++ {
		__obf_f16b4157911bc9af := __obf_6f3495a136f456ab[__obf_5aa5c8829b97f182]
		if __obf_f16b4157911bc9af < utf8.RuneSelf && __obf_eb32a2c54983008b[__obf_f16b4157911bc9af] {
			__obf_850a7457bb739a32.__obf_184433571fa55237 = append(__obf_850a7457bb739a32.__obf_184433571fa55237, __obf_f16b4157911bc9af)
		} else {
			break
		}
	}
	if __obf_5aa5c8829b97f182 == __obf_ed1950d3b33cfd95 {
		__obf_850a7457bb739a32.__obf_184433571fa55237 = append(__obf_850a7457bb739a32.__obf_184433571fa55237, '"')
		return
	}
	__obf_7e1f068906eabb52(__obf_850a7457bb739a32, __obf_5aa5c8829b97f182, __obf_6f3495a136f456ab, __obf_ed1950d3b33cfd95)
}

func __obf_7e1f068906eabb52(__obf_850a7457bb739a32 *Stream, __obf_5aa5c8829b97f182 int, __obf_6f3495a136f456ab string, __obf_ed1950d3b33cfd95 int) {
	__obf_924c0b8ced3b6bc7 := __obf_5aa5c8829b97f182
	// for the remaining parts, we process them char by char
	for __obf_5aa5c8829b97f182 < __obf_ed1950d3b33cfd95 {
		if __obf_1b9dc2dc0f9af749 := __obf_6f3495a136f456ab[__obf_5aa5c8829b97f182]; __obf_1b9dc2dc0f9af749 < utf8.RuneSelf {
			if __obf_eb32a2c54983008b[__obf_1b9dc2dc0f9af749] {
				__obf_5aa5c8829b97f182++
				continue
			}
			if __obf_924c0b8ced3b6bc7 < __obf_5aa5c8829b97f182 {
				__obf_850a7457bb739a32.
					WriteRaw(__obf_6f3495a136f456ab[__obf_924c0b8ced3b6bc7:__obf_5aa5c8829b97f182])
			}
			switch __obf_1b9dc2dc0f9af749 {
			case '\\', '"':
				__obf_850a7457bb739a32.__obf_56747a2b5f26ca3c('\\', __obf_1b9dc2dc0f9af749)
			case '\n':
				__obf_850a7457bb739a32.__obf_56747a2b5f26ca3c('\\', 'n')
			case '\r':
				__obf_850a7457bb739a32.__obf_56747a2b5f26ca3c('\\', 'r')
			case '\t':
				__obf_850a7457bb739a32.__obf_56747a2b5f26ca3c('\\', 't')
			default:
				__obf_850a7457bb739a32.
					// This encodes bytes < 0x20 except for \t, \n and \r.
					// If escapeHTML is set, it also escapes <, >, and &
					// because they can lead to security holes when
					// user-controlled strings are rendered into JSON
					// and served to some browsers.
					WriteRaw(`\u00`)
				__obf_850a7457bb739a32.__obf_56747a2b5f26ca3c(__obf_76ce97e7af7d5659[__obf_1b9dc2dc0f9af749>>4], __obf_76ce97e7af7d5659[__obf_1b9dc2dc0f9af749&0xF])
			}
			__obf_5aa5c8829b97f182++
			__obf_924c0b8ced3b6bc7 = __obf_5aa5c8829b97f182
			continue
		}
		__obf_f16b4157911bc9af, __obf_dc51447b63477324 := utf8.DecodeRuneInString(__obf_6f3495a136f456ab[__obf_5aa5c8829b97f182:])
		if __obf_f16b4157911bc9af == utf8.RuneError && __obf_dc51447b63477324 == 1 {
			if __obf_924c0b8ced3b6bc7 < __obf_5aa5c8829b97f182 {
				__obf_850a7457bb739a32.
					WriteRaw(__obf_6f3495a136f456ab[__obf_924c0b8ced3b6bc7:__obf_5aa5c8829b97f182])
			}
			__obf_850a7457bb739a32.
				WriteRaw(`\ufffd`)
			__obf_5aa5c8829b97f182++
			__obf_924c0b8ced3b6bc7 = __obf_5aa5c8829b97f182
			continue
		}
		// U+2028 is LINE SEPARATOR.
		// U+2029 is PARAGRAPH SEPARATOR.
		// They are both technically valid characters in JSON strings,
		// but don't work in JSONP, which has to be evaluated as JavaScript,
		// and can lead to security holes there. It is valid JSON to
		// escape them, so we do so unconditionally.
		// See http://timelessrepo.com/json-isnt-a-javascript-subset for discussion.
		if __obf_f16b4157911bc9af == '\u2028' || __obf_f16b4157911bc9af == '\u2029' {
			if __obf_924c0b8ced3b6bc7 < __obf_5aa5c8829b97f182 {
				__obf_850a7457bb739a32.
					WriteRaw(__obf_6f3495a136f456ab[__obf_924c0b8ced3b6bc7:__obf_5aa5c8829b97f182])
			}
			__obf_850a7457bb739a32.
				WriteRaw(`\u202`)
			__obf_850a7457bb739a32.__obf_72837f6fe737f078(__obf_76ce97e7af7d5659[__obf_f16b4157911bc9af&0xF])
			__obf_5aa5c8829b97f182 += __obf_dc51447b63477324
			__obf_924c0b8ced3b6bc7 = __obf_5aa5c8829b97f182
			continue
		}
		__obf_5aa5c8829b97f182 += __obf_dc51447b63477324
	}
	if __obf_924c0b8ced3b6bc7 < len(__obf_6f3495a136f456ab) {
		__obf_850a7457bb739a32.
			WriteRaw(__obf_6f3495a136f456ab[__obf_924c0b8ced3b6bc7:])
	}
	__obf_850a7457bb739a32.__obf_72837f6fe737f078('"')
}

// WriteString write string to stream without html escape
func (__obf_850a7457bb739a32 *Stream) WriteString(__obf_6f3495a136f456ab string) {
	__obf_ed1950d3b33cfd95 := len(__obf_6f3495a136f456ab)
	__obf_850a7457bb739a32.__obf_184433571fa55237 = append(__obf_850a7457bb739a32.
		// write string, the fast path, without utf8 and escape support
		__obf_184433571fa55237, '"')
	__obf_5aa5c8829b97f182 := 0
	for ; __obf_5aa5c8829b97f182 < __obf_ed1950d3b33cfd95; __obf_5aa5c8829b97f182++ {
		__obf_f16b4157911bc9af := __obf_6f3495a136f456ab[__obf_5aa5c8829b97f182]
		if __obf_f16b4157911bc9af > 31 && __obf_f16b4157911bc9af != '"' && __obf_f16b4157911bc9af != '\\' {
			__obf_850a7457bb739a32.__obf_184433571fa55237 = append(__obf_850a7457bb739a32.__obf_184433571fa55237, __obf_f16b4157911bc9af)
		} else {
			break
		}
	}
	if __obf_5aa5c8829b97f182 == __obf_ed1950d3b33cfd95 {
		__obf_850a7457bb739a32.__obf_184433571fa55237 = append(__obf_850a7457bb739a32.__obf_184433571fa55237, '"')
		return
	}
	__obf_20f3f99253c01e0d(__obf_850a7457bb739a32, __obf_5aa5c8829b97f182, __obf_6f3495a136f456ab, __obf_ed1950d3b33cfd95)
}

func __obf_20f3f99253c01e0d(__obf_850a7457bb739a32 *Stream, __obf_5aa5c8829b97f182 int, __obf_6f3495a136f456ab string, __obf_ed1950d3b33cfd95 int) {
	__obf_924c0b8ced3b6bc7 := __obf_5aa5c8829b97f182
	// for the remaining parts, we process them char by char
	for __obf_5aa5c8829b97f182 < __obf_ed1950d3b33cfd95 {
		if __obf_1b9dc2dc0f9af749 := __obf_6f3495a136f456ab[__obf_5aa5c8829b97f182]; __obf_1b9dc2dc0f9af749 < utf8.RuneSelf {
			if __obf_041351d57e164463[__obf_1b9dc2dc0f9af749] {
				__obf_5aa5c8829b97f182++
				continue
			}
			if __obf_924c0b8ced3b6bc7 < __obf_5aa5c8829b97f182 {
				__obf_850a7457bb739a32.
					WriteRaw(__obf_6f3495a136f456ab[__obf_924c0b8ced3b6bc7:__obf_5aa5c8829b97f182])
			}
			switch __obf_1b9dc2dc0f9af749 {
			case '\\', '"':
				__obf_850a7457bb739a32.__obf_56747a2b5f26ca3c('\\', __obf_1b9dc2dc0f9af749)
			case '\n':
				__obf_850a7457bb739a32.__obf_56747a2b5f26ca3c('\\', 'n')
			case '\r':
				__obf_850a7457bb739a32.__obf_56747a2b5f26ca3c('\\', 'r')
			case '\t':
				__obf_850a7457bb739a32.__obf_56747a2b5f26ca3c('\\', 't')
			default:
				__obf_850a7457bb739a32.
					// This encodes bytes < 0x20 except for \t, \n and \r.
					// If escapeHTML is set, it also escapes <, >, and &
					// because they can lead to security holes when
					// user-controlled strings are rendered into JSON
					// and served to some browsers.
					WriteRaw(`\u00`)
				__obf_850a7457bb739a32.__obf_56747a2b5f26ca3c(__obf_76ce97e7af7d5659[__obf_1b9dc2dc0f9af749>>4], __obf_76ce97e7af7d5659[__obf_1b9dc2dc0f9af749&0xF])
			}
			__obf_5aa5c8829b97f182++
			__obf_924c0b8ced3b6bc7 = __obf_5aa5c8829b97f182
			continue
		}
		__obf_5aa5c8829b97f182++
		continue
	}
	if __obf_924c0b8ced3b6bc7 < len(__obf_6f3495a136f456ab) {
		__obf_850a7457bb739a32.
			WriteRaw(__obf_6f3495a136f456ab[__obf_924c0b8ced3b6bc7:])
	}
	__obf_850a7457bb739a32.__obf_72837f6fe737f078('"')
}
