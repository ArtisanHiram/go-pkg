package __obf_1f22c6b8dfc77bff

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
var __obf_565e8d57d2410ecc = [utf8.RuneSelf]bool{
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

var __obf_24c6d2ed2e7faadc = reflect2.TypeOfPtr((*[]byte)(nil)).Elem()

type BinaryAsStringExtension struct {
	jsoniter.DummyExtension
}

func (__obf_9382a03d04135598 *BinaryAsStringExtension) CreateEncoder(__obf_66f97a1172af63eb reflect2.Type) jsoniter.ValEncoder {
	if __obf_66f97a1172af63eb == __obf_24c6d2ed2e7faadc {
		return &__obf_ee90dc36ed738ee2{}
	}
	return nil
}

func (__obf_9382a03d04135598 *BinaryAsStringExtension) CreateDecoder(__obf_66f97a1172af63eb reflect2.Type) jsoniter.ValDecoder {
	if __obf_66f97a1172af63eb == __obf_24c6d2ed2e7faadc {
		return &__obf_ee90dc36ed738ee2{}
	}
	return nil
}

type __obf_ee90dc36ed738ee2 struct {
}

func (__obf_fd6e92bcd33a0284 *__obf_ee90dc36ed738ee2) Decode(__obf_a5271c65f4ae17af unsafe.Pointer, __obf_d021dab62946a708 *jsoniter.Iterator) {
	__obf_e66be3973a3fc36a := __obf_d021dab62946a708.ReadStringAsSlice()
	__obf_635420ed849f1fa1 := make([]byte, 0, len(__obf_e66be3973a3fc36a))
	for __obf_f52abe4fb4613d80 := 0; __obf_f52abe4fb4613d80 < len(__obf_e66be3973a3fc36a); __obf_f52abe4fb4613d80++ {
		__obf_bc5528e6d68c497c := __obf_e66be3973a3fc36a[__obf_f52abe4fb4613d80]
		if __obf_bc5528e6d68c497c == '\\' {
			__obf_8f3d534ee3a0b541 := __obf_e66be3973a3fc36a[__obf_f52abe4fb4613d80+1]
			if __obf_8f3d534ee3a0b541 != '\\' {
				__obf_d021dab62946a708.
					ReportError("decode binary as string", `\\x is only supported escape`)
				return
			}
			__obf_99dfc793cc21305d := __obf_e66be3973a3fc36a[__obf_f52abe4fb4613d80+2]
			if __obf_99dfc793cc21305d != 'x' {
				__obf_d021dab62946a708.
					ReportError("decode binary as string", `\\x is only supported escape`)
				return
			}
			__obf_bd106fbaa0f67d97 := __obf_e66be3973a3fc36a[__obf_f52abe4fb4613d80+3]
			__obf_08dda0e7fde14fc9 := __obf_e66be3973a3fc36a[__obf_f52abe4fb4613d80+4]
			__obf_f52abe4fb4613d80 += 4
			__obf_bc5528e6d68c497c = __obf_837ff3f0d09d44dc(__obf_d021dab62946a708, __obf_bd106fbaa0f67d97, __obf_08dda0e7fde14fc9)
		}
		__obf_635420ed849f1fa1 = append(__obf_635420ed849f1fa1, __obf_bc5528e6d68c497c)
	}
	*(*[]byte)(__obf_a5271c65f4ae17af) = __obf_635420ed849f1fa1
}
func (__obf_fd6e92bcd33a0284 *__obf_ee90dc36ed738ee2) IsEmpty(__obf_a5271c65f4ae17af unsafe.Pointer) bool {
	return len(*((*[]byte)(__obf_a5271c65f4ae17af))) == 0
}
func (__obf_fd6e92bcd33a0284 *__obf_ee90dc36ed738ee2) Encode(__obf_a5271c65f4ae17af unsafe.Pointer, __obf_d178d558227696d6 *jsoniter.Stream) {
	__obf_114e226830ea8f9c := __obf_5dbded1ce3d29daa(__obf_d178d558227696d6.Buffer(), *(*[]byte)(__obf_a5271c65f4ae17af))
	__obf_d178d558227696d6.
		SetBuffer(__obf_114e226830ea8f9c)
}

func __obf_837ff3f0d09d44dc(__obf_d021dab62946a708 *jsoniter.Iterator, __obf_b109797a860b3dc6, __obf_8f3d534ee3a0b541 byte) byte {
	var __obf_03a5ce02d309c1fd byte
	if __obf_b109797a860b3dc6 >= '0' && __obf_b109797a860b3dc6 <= '9' {
		__obf_03a5ce02d309c1fd = __obf_b109797a860b3dc6 - '0'
	} else if __obf_b109797a860b3dc6 >= 'a' && __obf_b109797a860b3dc6 <= 'f' {
		__obf_03a5ce02d309c1fd = __obf_b109797a860b3dc6 - 'a' + 10
	} else {
		__obf_d021dab62946a708.
			ReportError("read hex", "expects 0~9 or a~f, but found "+string([]byte{__obf_b109797a860b3dc6}))
		return 0
	}
	__obf_03a5ce02d309c1fd *= 16
	if __obf_8f3d534ee3a0b541 >= '0' && __obf_8f3d534ee3a0b541 <= '9' {
		__obf_03a5ce02d309c1fd += __obf_8f3d534ee3a0b541 - '0'
	} else if __obf_8f3d534ee3a0b541 >= 'a' && __obf_8f3d534ee3a0b541 <= 'f' {
		__obf_03a5ce02d309c1fd += __obf_8f3d534ee3a0b541 - 'a' + 10
	} else {
		__obf_d021dab62946a708.
			ReportError("read hex", "expects 0~9 or a~f, but found "+string([]byte{__obf_8f3d534ee3a0b541}))
		return 0
	}
	return __obf_03a5ce02d309c1fd
}

var __obf_87d49a9f55d1c0e7 = "0123456789abcdef"

func __obf_5dbded1ce3d29daa(__obf_82bb9369f940a1ff []byte, __obf_0d8a34228fa11a3f []byte) []byte {
	__obf_82bb9369f940a1ff = append(__obf_82bb9369f940a1ff, '"')
	// write string, the fast path, without utf8 and escape support
	var __obf_f52abe4fb4613d80 int
	var __obf_f32e4134794a18a0 byte
	for __obf_f52abe4fb4613d80, __obf_f32e4134794a18a0 = range __obf_0d8a34228fa11a3f {
		if __obf_f32e4134794a18a0 < utf8.RuneSelf && __obf_565e8d57d2410ecc[__obf_f32e4134794a18a0] {
			__obf_82bb9369f940a1ff = append(__obf_82bb9369f940a1ff, __obf_f32e4134794a18a0)
		} else {
			break
		}
	}
	if __obf_f52abe4fb4613d80 == len(__obf_0d8a34228fa11a3f)-1 {
		__obf_82bb9369f940a1ff = append(__obf_82bb9369f940a1ff, '"')
		return __obf_82bb9369f940a1ff
	}
	return __obf_02ac89197205cb99(__obf_82bb9369f940a1ff, __obf_0d8a34228fa11a3f[__obf_f52abe4fb4613d80:])
}

func __obf_02ac89197205cb99(__obf_82bb9369f940a1ff []byte, __obf_0d8a34228fa11a3f []byte) []byte {
	__obf_ca64af48b909a39e := 0
	// for the remaining parts, we process them char by char
	var __obf_f52abe4fb4613d80 int
	var __obf_bc5528e6d68c497c byte
	for __obf_f52abe4fb4613d80, __obf_bc5528e6d68c497c = range __obf_0d8a34228fa11a3f {
		if __obf_bc5528e6d68c497c >= utf8.RuneSelf {
			__obf_82bb9369f940a1ff = append(__obf_82bb9369f940a1ff, '\\', '\\', 'x', __obf_87d49a9f55d1c0e7[__obf_bc5528e6d68c497c>>4], __obf_87d49a9f55d1c0e7[__obf_bc5528e6d68c497c&0xF])
			__obf_ca64af48b909a39e = __obf_f52abe4fb4613d80 + 1
			continue
		}
		if __obf_565e8d57d2410ecc[__obf_bc5528e6d68c497c] {
			continue
		}
		if __obf_ca64af48b909a39e < __obf_f52abe4fb4613d80 {
			__obf_82bb9369f940a1ff = append(__obf_82bb9369f940a1ff, __obf_0d8a34228fa11a3f[__obf_ca64af48b909a39e:__obf_f52abe4fb4613d80]...)
		}
		__obf_82bb9369f940a1ff = append(__obf_82bb9369f940a1ff, '\\', '\\', 'x', __obf_87d49a9f55d1c0e7[__obf_bc5528e6d68c497c>>4], __obf_87d49a9f55d1c0e7[__obf_bc5528e6d68c497c&0xF])
		__obf_ca64af48b909a39e = __obf_f52abe4fb4613d80 + 1
	}
	if __obf_ca64af48b909a39e < len(__obf_0d8a34228fa11a3f) {
		__obf_82bb9369f940a1ff = append(__obf_82bb9369f940a1ff, __obf_0d8a34228fa11a3f[__obf_ca64af48b909a39e:]...)
	}
	return append(__obf_82bb9369f940a1ff, '"')
}
