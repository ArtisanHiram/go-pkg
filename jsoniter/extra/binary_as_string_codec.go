package __obf_38f1d2f091ad74e0

import (
	jsoniter "github.com/ArtisanHiram/go-pkg/jsoniter"
	"github.com/modern-go/reflect2"
	"unicode/utf8"
	"unsafe"
)

// safeSet holds the value true if the ASCII character with the given array
// position can be represented inside a JSON string without any further
// escaping.
//
// All values are true except for the ASCII control characters (0-31), the
// double quote ("), and the backslash character ("\").
var __obf_a65e817db81ad76a = [utf8.RuneSelf]bool{
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

var __obf_63990f46a4264b7c = reflect2.TypeOfPtr((*[]byte)(nil)).Elem()

type BinaryAsStringExtension struct {
	jsoniter.DummyExtension
}

func (__obf_9c280f4c530aab7f *BinaryAsStringExtension) CreateEncoder(__obf_3c1b235923760ab5 reflect2.Type) jsoniter.ValEncoder {
	if __obf_3c1b235923760ab5 == __obf_63990f46a4264b7c {
		return &__obf_33bc4247e0e4a397{}
	}
	return nil
}

func (__obf_9c280f4c530aab7f *BinaryAsStringExtension) CreateDecoder(__obf_3c1b235923760ab5 reflect2.Type) jsoniter.ValDecoder {
	if __obf_3c1b235923760ab5 == __obf_63990f46a4264b7c {
		return &__obf_33bc4247e0e4a397{}
	}
	return nil
}

type __obf_33bc4247e0e4a397 struct {
}

func (__obf_af0a31f95254d8e7 *__obf_33bc4247e0e4a397) Decode(__obf_35567cf7daf6e12d unsafe.Pointer, __obf_113f80f39cc94185 *jsoniter.Iterator) {
	__obf_37b3fb015c18f5fb := __obf_113f80f39cc94185.ReadStringAsSlice()
	__obf_f7731faf22eea84c := make([]byte, 0, len(__obf_37b3fb015c18f5fb))
	for __obf_41047c070beaa4ab := 0; __obf_41047c070beaa4ab < len(__obf_37b3fb015c18f5fb); __obf_41047c070beaa4ab++ {
		__obf_fd3bcfa6ea6d8c13 := __obf_37b3fb015c18f5fb[__obf_41047c070beaa4ab]
		if __obf_fd3bcfa6ea6d8c13 == '\\' {
			__obf_5a5920d6a4bf4246 := __obf_37b3fb015c18f5fb[__obf_41047c070beaa4ab+1]
			if __obf_5a5920d6a4bf4246 != '\\' {
				__obf_113f80f39cc94185.
					ReportError("decode binary as string", `\\x is only supported escape`)
				return
			}
			__obf_033b0be8dbe7532c := __obf_37b3fb015c18f5fb[__obf_41047c070beaa4ab+2]
			if __obf_033b0be8dbe7532c != 'x' {
				__obf_113f80f39cc94185.
					ReportError("decode binary as string", `\\x is only supported escape`)
				return
			}
			__obf_18e93bc62ab96553 := __obf_37b3fb015c18f5fb[__obf_41047c070beaa4ab+3]
			__obf_4b0050bb13162e55 := __obf_37b3fb015c18f5fb[__obf_41047c070beaa4ab+4]
			__obf_41047c070beaa4ab += 4
			__obf_fd3bcfa6ea6d8c13 = __obf_9fb3097b2a27643e(__obf_113f80f39cc94185, __obf_18e93bc62ab96553, __obf_4b0050bb13162e55)
		}
		__obf_f7731faf22eea84c = append(__obf_f7731faf22eea84c, __obf_fd3bcfa6ea6d8c13)
	}
	*(*[]byte)(__obf_35567cf7daf6e12d) = __obf_f7731faf22eea84c
}
func (__obf_af0a31f95254d8e7 *__obf_33bc4247e0e4a397) IsEmpty(__obf_35567cf7daf6e12d unsafe.Pointer) bool {
	return len(*((*[]byte)(__obf_35567cf7daf6e12d))) == 0
}
func (__obf_af0a31f95254d8e7 *__obf_33bc4247e0e4a397) Encode(__obf_35567cf7daf6e12d unsafe.Pointer, __obf_27f0100519cc4eeb *jsoniter.Stream) {
	__obf_3d26aea54211d5a8 := __obf_5a86e4659b46a409(__obf_27f0100519cc4eeb.Buffer(), *(*[]byte)(__obf_35567cf7daf6e12d))
	__obf_27f0100519cc4eeb.
		SetBuffer(__obf_3d26aea54211d5a8)
}

func __obf_9fb3097b2a27643e(__obf_113f80f39cc94185 *jsoniter.Iterator, __obf_d64068806bf1ae95, __obf_5a5920d6a4bf4246 byte) byte {
	var __obf_39ba29f32e852854 byte
	if __obf_d64068806bf1ae95 >= '0' && __obf_d64068806bf1ae95 <= '9' {
		__obf_39ba29f32e852854 = __obf_d64068806bf1ae95 - '0'
	} else if __obf_d64068806bf1ae95 >= 'a' && __obf_d64068806bf1ae95 <= 'f' {
		__obf_39ba29f32e852854 = __obf_d64068806bf1ae95 - 'a' + 10
	} else {
		__obf_113f80f39cc94185.
			ReportError("read hex", "expects 0~9 or a~f, but found "+string([]byte{__obf_d64068806bf1ae95}))
		return 0
	}
	__obf_39ba29f32e852854 *= 16
	if __obf_5a5920d6a4bf4246 >= '0' && __obf_5a5920d6a4bf4246 <= '9' {
		__obf_39ba29f32e852854 += __obf_5a5920d6a4bf4246 - '0'
	} else if __obf_5a5920d6a4bf4246 >= 'a' && __obf_5a5920d6a4bf4246 <= 'f' {
		__obf_39ba29f32e852854 += __obf_5a5920d6a4bf4246 - 'a' + 10
	} else {
		__obf_113f80f39cc94185.
			ReportError("read hex", "expects 0~9 or a~f, but found "+string([]byte{__obf_5a5920d6a4bf4246}))
		return 0
	}
	return __obf_39ba29f32e852854
}

var __obf_806bafc73e053b98 = "0123456789abcdef"

func __obf_5a86e4659b46a409(__obf_67c7391ef0be6acc []byte, __obf_cf57c4db34e4cb93 []byte) []byte {
	__obf_67c7391ef0be6acc = append(__obf_67c7391ef0be6acc, '"')
	// write string, the fast path, without utf8 and escape support
	var __obf_41047c070beaa4ab int
	var __obf_412744d53cf48990 byte
	for __obf_41047c070beaa4ab, __obf_412744d53cf48990 = range __obf_cf57c4db34e4cb93 {
		if __obf_412744d53cf48990 < utf8.RuneSelf && __obf_a65e817db81ad76a[__obf_412744d53cf48990] {
			__obf_67c7391ef0be6acc = append(__obf_67c7391ef0be6acc, __obf_412744d53cf48990)
		} else {
			break
		}
	}
	if __obf_41047c070beaa4ab == len(__obf_cf57c4db34e4cb93)-1 {
		__obf_67c7391ef0be6acc = append(__obf_67c7391ef0be6acc, '"')
		return __obf_67c7391ef0be6acc
	}
	return __obf_1ba08cd61096ff96(__obf_67c7391ef0be6acc, __obf_cf57c4db34e4cb93[__obf_41047c070beaa4ab:])
}

func __obf_1ba08cd61096ff96(__obf_67c7391ef0be6acc []byte, __obf_cf57c4db34e4cb93 []byte) []byte {
	__obf_7e92dc336515a8b8 := 0
	// for the remaining parts, we process them char by char
	var __obf_41047c070beaa4ab int
	var __obf_fd3bcfa6ea6d8c13 byte
	for __obf_41047c070beaa4ab, __obf_fd3bcfa6ea6d8c13 = range __obf_cf57c4db34e4cb93 {
		if __obf_fd3bcfa6ea6d8c13 >= utf8.RuneSelf {
			__obf_67c7391ef0be6acc = append(__obf_67c7391ef0be6acc, '\\', '\\', 'x', __obf_806bafc73e053b98[__obf_fd3bcfa6ea6d8c13>>4], __obf_806bafc73e053b98[__obf_fd3bcfa6ea6d8c13&0xF])
			__obf_7e92dc336515a8b8 = __obf_41047c070beaa4ab + 1
			continue
		}
		if __obf_a65e817db81ad76a[__obf_fd3bcfa6ea6d8c13] {
			continue
		}
		if __obf_7e92dc336515a8b8 < __obf_41047c070beaa4ab {
			__obf_67c7391ef0be6acc = append(__obf_67c7391ef0be6acc, __obf_cf57c4db34e4cb93[__obf_7e92dc336515a8b8:__obf_41047c070beaa4ab]...)
		}
		__obf_67c7391ef0be6acc = append(__obf_67c7391ef0be6acc, '\\', '\\', 'x', __obf_806bafc73e053b98[__obf_fd3bcfa6ea6d8c13>>4], __obf_806bafc73e053b98[__obf_fd3bcfa6ea6d8c13&0xF])
		__obf_7e92dc336515a8b8 = __obf_41047c070beaa4ab + 1
	}
	if __obf_7e92dc336515a8b8 < len(__obf_cf57c4db34e4cb93) {
		__obf_67c7391ef0be6acc = append(__obf_67c7391ef0be6acc, __obf_cf57c4db34e4cb93[__obf_7e92dc336515a8b8:]...)
	}
	return append(__obf_67c7391ef0be6acc, '"')
}
