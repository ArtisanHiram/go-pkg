package __obf_9a397ef96534ad45

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
var __obf_58eb06631e97639d = [utf8.RuneSelf]bool{
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

var __obf_ed2a2840c6cd62ab = reflect2.TypeOfPtr((*[]byte)(nil)).Elem()

type BinaryAsStringExtension struct {
	jsoniter.DummyExtension
}

func (__obf_454cba947156c7ed *BinaryAsStringExtension) CreateEncoder(__obf_36cc3da433275e85 reflect2.Type) jsoniter.ValEncoder {
	if __obf_36cc3da433275e85 == __obf_ed2a2840c6cd62ab {
		return &__obf_c9eac5c11bd33652{}
	}
	return nil
}

func (__obf_454cba947156c7ed *BinaryAsStringExtension) CreateDecoder(__obf_36cc3da433275e85 reflect2.Type) jsoniter.ValDecoder {
	if __obf_36cc3da433275e85 == __obf_ed2a2840c6cd62ab {
		return &__obf_c9eac5c11bd33652{}
	}
	return nil
}

type __obf_c9eac5c11bd33652 struct {
}

func (__obf_d92fa10e8df80d33 *__obf_c9eac5c11bd33652) Decode(__obf_77e2593d34cc3a6a unsafe.Pointer, __obf_f2099f19d22a8175 *jsoniter.Iterator) {
	__obf_b7707f86f1b244a2 := __obf_f2099f19d22a8175.ReadStringAsSlice()
	__obf_73a7bbc094b11f03 := make([]byte, 0, len(__obf_b7707f86f1b244a2))
	for __obf_47b4fcf90774658f := 0; __obf_47b4fcf90774658f < len(__obf_b7707f86f1b244a2); __obf_47b4fcf90774658f++ {
		__obf_9f1019ba5c637808 := __obf_b7707f86f1b244a2[__obf_47b4fcf90774658f]
		if __obf_9f1019ba5c637808 == '\\' {
			__obf_a57242c44ace71f8 := __obf_b7707f86f1b244a2[__obf_47b4fcf90774658f+1]
			if __obf_a57242c44ace71f8 != '\\' {
				__obf_f2099f19d22a8175.
					ReportError("decode binary as string", `\\x is only supported escape`)
				return
			}
			__obf_cdf19b84dbf2691e := __obf_b7707f86f1b244a2[__obf_47b4fcf90774658f+2]
			if __obf_cdf19b84dbf2691e != 'x' {
				__obf_f2099f19d22a8175.
					ReportError("decode binary as string", `\\x is only supported escape`)
				return
			}
			__obf_c1d2bace8617c4de := __obf_b7707f86f1b244a2[__obf_47b4fcf90774658f+3]
			__obf_7a16490fedca451b := __obf_b7707f86f1b244a2[__obf_47b4fcf90774658f+4]
			__obf_47b4fcf90774658f += 4
			__obf_9f1019ba5c637808 = __obf_2383d0e3f5968196(__obf_f2099f19d22a8175, __obf_c1d2bace8617c4de, __obf_7a16490fedca451b)
		}
		__obf_73a7bbc094b11f03 = append(__obf_73a7bbc094b11f03, __obf_9f1019ba5c637808)
	}
	*(*[]byte)(__obf_77e2593d34cc3a6a) = __obf_73a7bbc094b11f03
}
func (__obf_d92fa10e8df80d33 *__obf_c9eac5c11bd33652) IsEmpty(__obf_77e2593d34cc3a6a unsafe.Pointer) bool {
	return len(*((*[]byte)(__obf_77e2593d34cc3a6a))) == 0
}
func (__obf_d92fa10e8df80d33 *__obf_c9eac5c11bd33652) Encode(__obf_77e2593d34cc3a6a unsafe.Pointer, __obf_3b68bc41f53ae503 *jsoniter.Stream) {
	__obf_7ecdc55c15da3dbf := __obf_804cf123d1338001(__obf_3b68bc41f53ae503.Buffer(), *(*[]byte)(__obf_77e2593d34cc3a6a))
	__obf_3b68bc41f53ae503.
		SetBuffer(__obf_7ecdc55c15da3dbf)
}

func __obf_2383d0e3f5968196(__obf_f2099f19d22a8175 *jsoniter.Iterator, __obf_baf4d3921d40b5f8, __obf_a57242c44ace71f8 byte) byte {
	var __obf_5d6932874351bc07 byte
	if __obf_baf4d3921d40b5f8 >= '0' && __obf_baf4d3921d40b5f8 <= '9' {
		__obf_5d6932874351bc07 = __obf_baf4d3921d40b5f8 - '0'
	} else if __obf_baf4d3921d40b5f8 >= 'a' && __obf_baf4d3921d40b5f8 <= 'f' {
		__obf_5d6932874351bc07 = __obf_baf4d3921d40b5f8 - 'a' + 10
	} else {
		__obf_f2099f19d22a8175.
			ReportError("read hex", "expects 0~9 or a~f, but found "+string([]byte{__obf_baf4d3921d40b5f8}))
		return 0
	}
	__obf_5d6932874351bc07 *= 16
	if __obf_a57242c44ace71f8 >= '0' && __obf_a57242c44ace71f8 <= '9' {
		__obf_5d6932874351bc07 += __obf_a57242c44ace71f8 - '0'
	} else if __obf_a57242c44ace71f8 >= 'a' && __obf_a57242c44ace71f8 <= 'f' {
		__obf_5d6932874351bc07 += __obf_a57242c44ace71f8 - 'a' + 10
	} else {
		__obf_f2099f19d22a8175.
			ReportError("read hex", "expects 0~9 or a~f, but found "+string([]byte{__obf_a57242c44ace71f8}))
		return 0
	}
	return __obf_5d6932874351bc07
}

var __obf_0d393f050940deba = "0123456789abcdef"

func __obf_804cf123d1338001(__obf_6861139d86b12566 []byte, __obf_3c2a0b9a02fa76c3 []byte) []byte {
	__obf_6861139d86b12566 = append(__obf_6861139d86b12566, '"')
	// write string, the fast path, without utf8 and escape support
	var __obf_47b4fcf90774658f int
	var __obf_08774a099cab6503 byte
	for __obf_47b4fcf90774658f, __obf_08774a099cab6503 = range __obf_3c2a0b9a02fa76c3 {
		if __obf_08774a099cab6503 < utf8.RuneSelf && __obf_58eb06631e97639d[__obf_08774a099cab6503] {
			__obf_6861139d86b12566 = append(__obf_6861139d86b12566, __obf_08774a099cab6503)
		} else {
			break
		}
	}
	if __obf_47b4fcf90774658f == len(__obf_3c2a0b9a02fa76c3)-1 {
		__obf_6861139d86b12566 = append(__obf_6861139d86b12566, '"')
		return __obf_6861139d86b12566
	}
	return __obf_90135c9cfc0bcfc1(__obf_6861139d86b12566, __obf_3c2a0b9a02fa76c3[__obf_47b4fcf90774658f:])
}

func __obf_90135c9cfc0bcfc1(__obf_6861139d86b12566 []byte, __obf_3c2a0b9a02fa76c3 []byte) []byte {
	__obf_78d081b33808be6b := 0
	// for the remaining parts, we process them char by char
	var __obf_47b4fcf90774658f int
	var __obf_9f1019ba5c637808 byte
	for __obf_47b4fcf90774658f, __obf_9f1019ba5c637808 = range __obf_3c2a0b9a02fa76c3 {
		if __obf_9f1019ba5c637808 >= utf8.RuneSelf {
			__obf_6861139d86b12566 = append(__obf_6861139d86b12566, '\\', '\\', 'x', __obf_0d393f050940deba[__obf_9f1019ba5c637808>>4], __obf_0d393f050940deba[__obf_9f1019ba5c637808&0xF])
			__obf_78d081b33808be6b = __obf_47b4fcf90774658f + 1
			continue
		}
		if __obf_58eb06631e97639d[__obf_9f1019ba5c637808] {
			continue
		}
		if __obf_78d081b33808be6b < __obf_47b4fcf90774658f {
			__obf_6861139d86b12566 = append(__obf_6861139d86b12566, __obf_3c2a0b9a02fa76c3[__obf_78d081b33808be6b:__obf_47b4fcf90774658f]...)
		}
		__obf_6861139d86b12566 = append(__obf_6861139d86b12566, '\\', '\\', 'x', __obf_0d393f050940deba[__obf_9f1019ba5c637808>>4], __obf_0d393f050940deba[__obf_9f1019ba5c637808&0xF])
		__obf_78d081b33808be6b = __obf_47b4fcf90774658f + 1
	}
	if __obf_78d081b33808be6b < len(__obf_3c2a0b9a02fa76c3) {
		__obf_6861139d86b12566 = append(__obf_6861139d86b12566, __obf_3c2a0b9a02fa76c3[__obf_78d081b33808be6b:]...)
	}
	return append(__obf_6861139d86b12566, '"')
}
