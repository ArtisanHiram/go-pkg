package __obf_c3cd04a15c56f32f

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
var __obf_0e44dbc035010740 = [utf8.RuneSelf]bool{
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
var __obf_0d124a3a00bcd99e = [utf8.RuneSelf]bool{
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

var __obf_2fe6db6e4b4829cc = "0123456789abcdef"

// WriteStringWithHTMLEscaped write string to stream with html special characters escaped
func (__obf_2361f5214e84c60f *Stream) WriteStringWithHTMLEscaped(__obf_20e65aaba6bfc813 string) {
	__obf_dec599590af4a21c := len(__obf_20e65aaba6bfc813)
	__obf_2361f5214e84c60f.__obf_ace979f6622823f3 = append(__obf_2361f5214e84c60f.
		// write string, the fast path, without utf8 and escape support
		__obf_ace979f6622823f3, '"')
	__obf_28d099df85f083a8 := 0
	for ; __obf_28d099df85f083a8 < __obf_dec599590af4a21c; __obf_28d099df85f083a8++ {
		__obf_0c1bc1e511a43120 := __obf_20e65aaba6bfc813[__obf_28d099df85f083a8]
		if __obf_0c1bc1e511a43120 < utf8.RuneSelf && __obf_0e44dbc035010740[__obf_0c1bc1e511a43120] {
			__obf_2361f5214e84c60f.__obf_ace979f6622823f3 = append(__obf_2361f5214e84c60f.__obf_ace979f6622823f3, __obf_0c1bc1e511a43120)
		} else {
			break
		}
	}
	if __obf_28d099df85f083a8 == __obf_dec599590af4a21c {
		__obf_2361f5214e84c60f.__obf_ace979f6622823f3 = append(__obf_2361f5214e84c60f.__obf_ace979f6622823f3, '"')
		return
	}
	__obf_878cb4feba50d089(__obf_2361f5214e84c60f, __obf_28d099df85f083a8, __obf_20e65aaba6bfc813, __obf_dec599590af4a21c)
}

func __obf_878cb4feba50d089(__obf_2361f5214e84c60f *Stream, __obf_28d099df85f083a8 int, __obf_20e65aaba6bfc813 string, __obf_dec599590af4a21c int) {
	__obf_91c0b88ab32aec51 := __obf_28d099df85f083a8
	// for the remaining parts, we process them char by char
	for __obf_28d099df85f083a8 < __obf_dec599590af4a21c {
		if __obf_902d7026e8a09dd2 := __obf_20e65aaba6bfc813[__obf_28d099df85f083a8]; __obf_902d7026e8a09dd2 < utf8.RuneSelf {
			if __obf_0e44dbc035010740[__obf_902d7026e8a09dd2] {
				__obf_28d099df85f083a8++
				continue
			}
			if __obf_91c0b88ab32aec51 < __obf_28d099df85f083a8 {
				__obf_2361f5214e84c60f.
					WriteRaw(__obf_20e65aaba6bfc813[__obf_91c0b88ab32aec51:__obf_28d099df85f083a8])
			}
			switch __obf_902d7026e8a09dd2 {
			case '\\', '"':
				__obf_2361f5214e84c60f.__obf_5e728551f00598e5('\\', __obf_902d7026e8a09dd2)
			case '\n':
				__obf_2361f5214e84c60f.__obf_5e728551f00598e5('\\', 'n')
			case '\r':
				__obf_2361f5214e84c60f.__obf_5e728551f00598e5('\\', 'r')
			case '\t':
				__obf_2361f5214e84c60f.__obf_5e728551f00598e5('\\', 't')
			default:
				__obf_2361f5214e84c60f.
					// This encodes bytes < 0x20 except for \t, \n and \r.
					// If escapeHTML is set, it also escapes <, >, and &
					// because they can lead to security holes when
					// user-controlled strings are rendered into JSON
					// and served to some browsers.
					WriteRaw(`\u00`)
				__obf_2361f5214e84c60f.__obf_5e728551f00598e5(__obf_2fe6db6e4b4829cc[__obf_902d7026e8a09dd2>>4], __obf_2fe6db6e4b4829cc[__obf_902d7026e8a09dd2&0xF])
			}
			__obf_28d099df85f083a8++
			__obf_91c0b88ab32aec51 = __obf_28d099df85f083a8
			continue
		}
		__obf_0c1bc1e511a43120, __obf_ecf95be2d6e27166 := utf8.DecodeRuneInString(__obf_20e65aaba6bfc813[__obf_28d099df85f083a8:])
		if __obf_0c1bc1e511a43120 == utf8.RuneError && __obf_ecf95be2d6e27166 == 1 {
			if __obf_91c0b88ab32aec51 < __obf_28d099df85f083a8 {
				__obf_2361f5214e84c60f.
					WriteRaw(__obf_20e65aaba6bfc813[__obf_91c0b88ab32aec51:__obf_28d099df85f083a8])
			}
			__obf_2361f5214e84c60f.
				WriteRaw(`\ufffd`)
			__obf_28d099df85f083a8++
			__obf_91c0b88ab32aec51 = __obf_28d099df85f083a8
			continue
		}
		// U+2028 is LINE SEPARATOR.
		// U+2029 is PARAGRAPH SEPARATOR.
		// They are both technically valid characters in JSON strings,
		// but don't work in JSONP, which has to be evaluated as JavaScript,
		// and can lead to security holes there. It is valid JSON to
		// escape them, so we do so unconditionally.
		// See http://timelessrepo.com/json-isnt-a-javascript-subset for discussion.
		if __obf_0c1bc1e511a43120 == '\u2028' || __obf_0c1bc1e511a43120 == '\u2029' {
			if __obf_91c0b88ab32aec51 < __obf_28d099df85f083a8 {
				__obf_2361f5214e84c60f.
					WriteRaw(__obf_20e65aaba6bfc813[__obf_91c0b88ab32aec51:__obf_28d099df85f083a8])
			}
			__obf_2361f5214e84c60f.
				WriteRaw(`\u202`)
			__obf_2361f5214e84c60f.__obf_c4fec0edfb3875ad(__obf_2fe6db6e4b4829cc[__obf_0c1bc1e511a43120&0xF])
			__obf_28d099df85f083a8 += __obf_ecf95be2d6e27166
			__obf_91c0b88ab32aec51 = __obf_28d099df85f083a8
			continue
		}
		__obf_28d099df85f083a8 += __obf_ecf95be2d6e27166
	}
	if __obf_91c0b88ab32aec51 < len(__obf_20e65aaba6bfc813) {
		__obf_2361f5214e84c60f.
			WriteRaw(__obf_20e65aaba6bfc813[__obf_91c0b88ab32aec51:])
	}
	__obf_2361f5214e84c60f.__obf_c4fec0edfb3875ad('"')
}

// WriteString write string to stream without html escape
func (__obf_2361f5214e84c60f *Stream) WriteString(__obf_20e65aaba6bfc813 string) {
	__obf_dec599590af4a21c := len(__obf_20e65aaba6bfc813)
	__obf_2361f5214e84c60f.__obf_ace979f6622823f3 = append(__obf_2361f5214e84c60f.
		// write string, the fast path, without utf8 and escape support
		__obf_ace979f6622823f3, '"')
	__obf_28d099df85f083a8 := 0
	for ; __obf_28d099df85f083a8 < __obf_dec599590af4a21c; __obf_28d099df85f083a8++ {
		__obf_0c1bc1e511a43120 := __obf_20e65aaba6bfc813[__obf_28d099df85f083a8]
		if __obf_0c1bc1e511a43120 > 31 && __obf_0c1bc1e511a43120 != '"' && __obf_0c1bc1e511a43120 != '\\' {
			__obf_2361f5214e84c60f.__obf_ace979f6622823f3 = append(__obf_2361f5214e84c60f.__obf_ace979f6622823f3, __obf_0c1bc1e511a43120)
		} else {
			break
		}
	}
	if __obf_28d099df85f083a8 == __obf_dec599590af4a21c {
		__obf_2361f5214e84c60f.__obf_ace979f6622823f3 = append(__obf_2361f5214e84c60f.__obf_ace979f6622823f3, '"')
		return
	}
	__obf_85c4b0ee02a95710(__obf_2361f5214e84c60f, __obf_28d099df85f083a8, __obf_20e65aaba6bfc813, __obf_dec599590af4a21c)
}

func __obf_85c4b0ee02a95710(__obf_2361f5214e84c60f *Stream, __obf_28d099df85f083a8 int, __obf_20e65aaba6bfc813 string, __obf_dec599590af4a21c int) {
	__obf_91c0b88ab32aec51 := __obf_28d099df85f083a8
	// for the remaining parts, we process them char by char
	for __obf_28d099df85f083a8 < __obf_dec599590af4a21c {
		if __obf_902d7026e8a09dd2 := __obf_20e65aaba6bfc813[__obf_28d099df85f083a8]; __obf_902d7026e8a09dd2 < utf8.RuneSelf {
			if __obf_0d124a3a00bcd99e[__obf_902d7026e8a09dd2] {
				__obf_28d099df85f083a8++
				continue
			}
			if __obf_91c0b88ab32aec51 < __obf_28d099df85f083a8 {
				__obf_2361f5214e84c60f.
					WriteRaw(__obf_20e65aaba6bfc813[__obf_91c0b88ab32aec51:__obf_28d099df85f083a8])
			}
			switch __obf_902d7026e8a09dd2 {
			case '\\', '"':
				__obf_2361f5214e84c60f.__obf_5e728551f00598e5('\\', __obf_902d7026e8a09dd2)
			case '\n':
				__obf_2361f5214e84c60f.__obf_5e728551f00598e5('\\', 'n')
			case '\r':
				__obf_2361f5214e84c60f.__obf_5e728551f00598e5('\\', 'r')
			case '\t':
				__obf_2361f5214e84c60f.__obf_5e728551f00598e5('\\', 't')
			default:
				__obf_2361f5214e84c60f.
					// This encodes bytes < 0x20 except for \t, \n and \r.
					// If escapeHTML is set, it also escapes <, >, and &
					// because they can lead to security holes when
					// user-controlled strings are rendered into JSON
					// and served to some browsers.
					WriteRaw(`\u00`)
				__obf_2361f5214e84c60f.__obf_5e728551f00598e5(__obf_2fe6db6e4b4829cc[__obf_902d7026e8a09dd2>>4], __obf_2fe6db6e4b4829cc[__obf_902d7026e8a09dd2&0xF])
			}
			__obf_28d099df85f083a8++
			__obf_91c0b88ab32aec51 = __obf_28d099df85f083a8
			continue
		}
		__obf_28d099df85f083a8++
		continue
	}
	if __obf_91c0b88ab32aec51 < len(__obf_20e65aaba6bfc813) {
		__obf_2361f5214e84c60f.
			WriteRaw(__obf_20e65aaba6bfc813[__obf_91c0b88ab32aec51:])
	}
	__obf_2361f5214e84c60f.__obf_c4fec0edfb3875ad('"')
}
