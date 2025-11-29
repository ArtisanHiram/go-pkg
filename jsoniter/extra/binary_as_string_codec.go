package __obf_789dc33d47f4ab2c

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
var __obf_3bd25e1bf21b7feb = [utf8.RuneSelf]bool{
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

var __obf_1dfdec16d9ee05e6 = reflect2.TypeOfPtr((*[]byte)(nil)).Elem()

type BinaryAsStringExtension struct {
	jsoniter.DummyExtension
}

func (__obf_369258fb13d82ac5 *BinaryAsStringExtension) CreateEncoder(__obf_487781debab9b36c reflect2.Type) jsoniter.ValEncoder {
	if __obf_487781debab9b36c == __obf_1dfdec16d9ee05e6 {
		return &__obf_1021b76cf8d4e45c{}
	}
	return nil
}

func (__obf_369258fb13d82ac5 *BinaryAsStringExtension) CreateDecoder(__obf_487781debab9b36c reflect2.Type) jsoniter.ValDecoder {
	if __obf_487781debab9b36c == __obf_1dfdec16d9ee05e6 {
		return &__obf_1021b76cf8d4e45c{}
	}
	return nil
}

type __obf_1021b76cf8d4e45c struct {
}

func (__obf_d02baddfcf2572b4 *__obf_1021b76cf8d4e45c) Decode(__obf_4280d1a4ac42e464 unsafe.Pointer, __obf_e97e2ea2e3d1b0a2 *jsoniter.Iterator) {
	__obf_23cc5c440aff96b7 := __obf_e97e2ea2e3d1b0a2.ReadStringAsSlice()
	__obf_5dba6019524f725e := make([]byte, 0, len(__obf_23cc5c440aff96b7))
	for __obf_c1c90a275d52d9e6 := 0; __obf_c1c90a275d52d9e6 < len(__obf_23cc5c440aff96b7); __obf_c1c90a275d52d9e6++ {
		__obf_2402aae457c079d1 := __obf_23cc5c440aff96b7[__obf_c1c90a275d52d9e6]
		if __obf_2402aae457c079d1 == '\\' {
			__obf_c232999a3905136c := __obf_23cc5c440aff96b7[__obf_c1c90a275d52d9e6+1]
			if __obf_c232999a3905136c != '\\' {
				__obf_e97e2ea2e3d1b0a2.
					ReportError("decode binary as string", `\\x is only supported escape`)
				return
			}
			__obf_4428dd66ab560623 := __obf_23cc5c440aff96b7[__obf_c1c90a275d52d9e6+2]
			if __obf_4428dd66ab560623 != 'x' {
				__obf_e97e2ea2e3d1b0a2.
					ReportError("decode binary as string", `\\x is only supported escape`)
				return
			}
			__obf_c69d98e4cb8da9a8 := __obf_23cc5c440aff96b7[__obf_c1c90a275d52d9e6+3]
			__obf_4c97ad69dfbb29aa := __obf_23cc5c440aff96b7[__obf_c1c90a275d52d9e6+4]
			__obf_c1c90a275d52d9e6 += 4
			__obf_2402aae457c079d1 = __obf_ca6c51d9c7042c69(__obf_e97e2ea2e3d1b0a2, __obf_c69d98e4cb8da9a8, __obf_4c97ad69dfbb29aa)
		}
		__obf_5dba6019524f725e = append(__obf_5dba6019524f725e, __obf_2402aae457c079d1)
	}
	*(*[]byte)(__obf_4280d1a4ac42e464) = __obf_5dba6019524f725e
}
func (__obf_d02baddfcf2572b4 *__obf_1021b76cf8d4e45c) IsEmpty(__obf_4280d1a4ac42e464 unsafe.Pointer) bool {
	return len(*((*[]byte)(__obf_4280d1a4ac42e464))) == 0
}
func (__obf_d02baddfcf2572b4 *__obf_1021b76cf8d4e45c) Encode(__obf_4280d1a4ac42e464 unsafe.Pointer, __obf_fe6de34d139b90e2 *jsoniter.Stream) {
	__obf_7c33ba66dbc211bc := __obf_d884f3d261f1244c(__obf_fe6de34d139b90e2.Buffer(), *(*[]byte)(__obf_4280d1a4ac42e464))
	__obf_fe6de34d139b90e2.
		SetBuffer(__obf_7c33ba66dbc211bc)
}

func __obf_ca6c51d9c7042c69(__obf_e97e2ea2e3d1b0a2 *jsoniter.Iterator, __obf_783d70441558c8f4, __obf_c232999a3905136c byte) byte {
	var __obf_99a010aa00b760bc byte
	if __obf_783d70441558c8f4 >= '0' && __obf_783d70441558c8f4 <= '9' {
		__obf_99a010aa00b760bc = __obf_783d70441558c8f4 - '0'
	} else if __obf_783d70441558c8f4 >= 'a' && __obf_783d70441558c8f4 <= 'f' {
		__obf_99a010aa00b760bc = __obf_783d70441558c8f4 - 'a' + 10
	} else {
		__obf_e97e2ea2e3d1b0a2.
			ReportError("read hex", "expects 0~9 or a~f, but found "+string([]byte{__obf_783d70441558c8f4}))
		return 0
	}
	__obf_99a010aa00b760bc *= 16
	if __obf_c232999a3905136c >= '0' && __obf_c232999a3905136c <= '9' {
		__obf_99a010aa00b760bc += __obf_c232999a3905136c - '0'
	} else if __obf_c232999a3905136c >= 'a' && __obf_c232999a3905136c <= 'f' {
		__obf_99a010aa00b760bc += __obf_c232999a3905136c - 'a' + 10
	} else {
		__obf_e97e2ea2e3d1b0a2.
			ReportError("read hex", "expects 0~9 or a~f, but found "+string([]byte{__obf_c232999a3905136c}))
		return 0
	}
	return __obf_99a010aa00b760bc
}

var __obf_cc4bcfa902100f9d = "0123456789abcdef"

func __obf_d884f3d261f1244c(__obf_2b86e14b7d787b5f []byte, __obf_4a9c70e1ac3f841e []byte) []byte {
	__obf_2b86e14b7d787b5f = append(__obf_2b86e14b7d787b5f, '"')
	// write string, the fast path, without utf8 and escape support
	var __obf_c1c90a275d52d9e6 int
	var __obf_c39e01a0b80979b1 byte
	for __obf_c1c90a275d52d9e6, __obf_c39e01a0b80979b1 = range __obf_4a9c70e1ac3f841e {
		if __obf_c39e01a0b80979b1 < utf8.RuneSelf && __obf_3bd25e1bf21b7feb[__obf_c39e01a0b80979b1] {
			__obf_2b86e14b7d787b5f = append(__obf_2b86e14b7d787b5f, __obf_c39e01a0b80979b1)
		} else {
			break
		}
	}
	if __obf_c1c90a275d52d9e6 == len(__obf_4a9c70e1ac3f841e)-1 {
		__obf_2b86e14b7d787b5f = append(__obf_2b86e14b7d787b5f, '"')
		return __obf_2b86e14b7d787b5f
	}
	return __obf_41b79aa4aeaa54e4(__obf_2b86e14b7d787b5f, __obf_4a9c70e1ac3f841e[__obf_c1c90a275d52d9e6:])
}

func __obf_41b79aa4aeaa54e4(__obf_2b86e14b7d787b5f []byte, __obf_4a9c70e1ac3f841e []byte) []byte {
	__obf_755e3f990a78402f := 0
	// for the remaining parts, we process them char by char
	var __obf_c1c90a275d52d9e6 int
	var __obf_2402aae457c079d1 byte
	for __obf_c1c90a275d52d9e6, __obf_2402aae457c079d1 = range __obf_4a9c70e1ac3f841e {
		if __obf_2402aae457c079d1 >= utf8.RuneSelf {
			__obf_2b86e14b7d787b5f = append(__obf_2b86e14b7d787b5f, '\\', '\\', 'x', __obf_cc4bcfa902100f9d[__obf_2402aae457c079d1>>4], __obf_cc4bcfa902100f9d[__obf_2402aae457c079d1&0xF])
			__obf_755e3f990a78402f = __obf_c1c90a275d52d9e6 + 1
			continue
		}
		if __obf_3bd25e1bf21b7feb[__obf_2402aae457c079d1] {
			continue
		}
		if __obf_755e3f990a78402f < __obf_c1c90a275d52d9e6 {
			__obf_2b86e14b7d787b5f = append(__obf_2b86e14b7d787b5f, __obf_4a9c70e1ac3f841e[__obf_755e3f990a78402f:__obf_c1c90a275d52d9e6]...)
		}
		__obf_2b86e14b7d787b5f = append(__obf_2b86e14b7d787b5f, '\\', '\\', 'x', __obf_cc4bcfa902100f9d[__obf_2402aae457c079d1>>4], __obf_cc4bcfa902100f9d[__obf_2402aae457c079d1&0xF])
		__obf_755e3f990a78402f = __obf_c1c90a275d52d9e6 + 1
	}
	if __obf_755e3f990a78402f < len(__obf_4a9c70e1ac3f841e) {
		__obf_2b86e14b7d787b5f = append(__obf_2b86e14b7d787b5f, __obf_4a9c70e1ac3f841e[__obf_755e3f990a78402f:]...)
	}
	return append(__obf_2b86e14b7d787b5f, '"')
}
