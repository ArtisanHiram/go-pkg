package __obf_eed9c5a643743c33

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
var __obf_0eab2aa051a8dd12 = [utf8.RuneSelf]bool{
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

var __obf_cda44835ddd97604 = reflect2.TypeOfPtr((*[]byte)(nil)).Elem()

type BinaryAsStringExtension struct {
	jsoniter.DummyExtension
}

func (__obf_7ab69ccfb0d084de *BinaryAsStringExtension) CreateEncoder(__obf_4db0904e845b7200 reflect2.Type) jsoniter.ValEncoder {
	if __obf_4db0904e845b7200 == __obf_cda44835ddd97604 {
		return &__obf_9bf279fe700a7152{}
	}
	return nil
}

func (__obf_7ab69ccfb0d084de *BinaryAsStringExtension) CreateDecoder(__obf_4db0904e845b7200 reflect2.Type) jsoniter.ValDecoder {
	if __obf_4db0904e845b7200 == __obf_cda44835ddd97604 {
		return &__obf_9bf279fe700a7152{}
	}
	return nil
}

type __obf_9bf279fe700a7152 struct {
}

func (__obf_ccd1a335201fc2ad *__obf_9bf279fe700a7152) Decode(__obf_e88eec03e62b77ec unsafe.Pointer, __obf_338eca60ccc15d82 *jsoniter.Iterator) {
	__obf_a86e50dd7e44b5ab := __obf_338eca60ccc15d82.ReadStringAsSlice()
	__obf_d1915e73b125d016 := make([]byte, 0, len(__obf_a86e50dd7e44b5ab))
	for __obf_6db94c729874d6be := 0; __obf_6db94c729874d6be < len(__obf_a86e50dd7e44b5ab); __obf_6db94c729874d6be++ {
		__obf_62966ecce30cc903 := __obf_a86e50dd7e44b5ab[__obf_6db94c729874d6be]
		if __obf_62966ecce30cc903 == '\\' {
			__obf_bcbf560b54c31236 := __obf_a86e50dd7e44b5ab[__obf_6db94c729874d6be+1]
			if __obf_bcbf560b54c31236 != '\\' {
				__obf_338eca60ccc15d82.
					ReportError("decode binary as string", `\\x is only supported escape`)
				return
			}
			__obf_19eb7236a968298b := __obf_a86e50dd7e44b5ab[__obf_6db94c729874d6be+2]
			if __obf_19eb7236a968298b != 'x' {
				__obf_338eca60ccc15d82.
					ReportError("decode binary as string", `\\x is only supported escape`)
				return
			}
			__obf_bcaa1044ca193d6d := __obf_a86e50dd7e44b5ab[__obf_6db94c729874d6be+3]
			__obf_65feac19926b842b := __obf_a86e50dd7e44b5ab[__obf_6db94c729874d6be+4]
			__obf_6db94c729874d6be += 4
			__obf_62966ecce30cc903 = __obf_a027d487017539c2(__obf_338eca60ccc15d82, __obf_bcaa1044ca193d6d, __obf_65feac19926b842b)
		}
		__obf_d1915e73b125d016 = append(__obf_d1915e73b125d016, __obf_62966ecce30cc903)
	}
	*(*[]byte)(__obf_e88eec03e62b77ec) = __obf_d1915e73b125d016
}
func (__obf_ccd1a335201fc2ad *__obf_9bf279fe700a7152) IsEmpty(__obf_e88eec03e62b77ec unsafe.Pointer) bool {
	return len(*((*[]byte)(__obf_e88eec03e62b77ec))) == 0
}
func (__obf_ccd1a335201fc2ad *__obf_9bf279fe700a7152) Encode(__obf_e88eec03e62b77ec unsafe.Pointer, __obf_837be8d71267e11f *jsoniter.Stream) {
	__obf_fcfc67ed7b00a86f := __obf_11fd5d7520265863(__obf_837be8d71267e11f.Buffer(), *(*[]byte)(__obf_e88eec03e62b77ec))
	__obf_837be8d71267e11f.
		SetBuffer(__obf_fcfc67ed7b00a86f)
}

func __obf_a027d487017539c2(__obf_338eca60ccc15d82 *jsoniter.Iterator, __obf_7667ab75b6a91bb1, __obf_bcbf560b54c31236 byte) byte {
	var __obf_c32acc8b9240981b byte
	if __obf_7667ab75b6a91bb1 >= '0' && __obf_7667ab75b6a91bb1 <= '9' {
		__obf_c32acc8b9240981b = __obf_7667ab75b6a91bb1 - '0'
	} else if __obf_7667ab75b6a91bb1 >= 'a' && __obf_7667ab75b6a91bb1 <= 'f' {
		__obf_c32acc8b9240981b = __obf_7667ab75b6a91bb1 - 'a' + 10
	} else {
		__obf_338eca60ccc15d82.
			ReportError("read hex", "expects 0~9 or a~f, but found "+string([]byte{__obf_7667ab75b6a91bb1}))
		return 0
	}
	__obf_c32acc8b9240981b *= 16
	if __obf_bcbf560b54c31236 >= '0' && __obf_bcbf560b54c31236 <= '9' {
		__obf_c32acc8b9240981b += __obf_bcbf560b54c31236 - '0'
	} else if __obf_bcbf560b54c31236 >= 'a' && __obf_bcbf560b54c31236 <= 'f' {
		__obf_c32acc8b9240981b += __obf_bcbf560b54c31236 - 'a' + 10
	} else {
		__obf_338eca60ccc15d82.
			ReportError("read hex", "expects 0~9 or a~f, but found "+string([]byte{__obf_bcbf560b54c31236}))
		return 0
	}
	return __obf_c32acc8b9240981b
}

var __obf_61ea456740bf5ed3 = "0123456789abcdef"

func __obf_11fd5d7520265863(__obf_e7f5496ce1c3e6b2 []byte, __obf_f9431224f9a648b3 []byte) []byte {
	__obf_e7f5496ce1c3e6b2 = append(__obf_e7f5496ce1c3e6b2, '"')
	// write string, the fast path, without utf8 and escape support
	var __obf_6db94c729874d6be int
	var __obf_d609c92bf973ccdd byte
	for __obf_6db94c729874d6be, __obf_d609c92bf973ccdd = range __obf_f9431224f9a648b3 {
		if __obf_d609c92bf973ccdd < utf8.RuneSelf && __obf_0eab2aa051a8dd12[__obf_d609c92bf973ccdd] {
			__obf_e7f5496ce1c3e6b2 = append(__obf_e7f5496ce1c3e6b2, __obf_d609c92bf973ccdd)
		} else {
			break
		}
	}
	if __obf_6db94c729874d6be == len(__obf_f9431224f9a648b3)-1 {
		__obf_e7f5496ce1c3e6b2 = append(__obf_e7f5496ce1c3e6b2, '"')
		return __obf_e7f5496ce1c3e6b2
	}
	return __obf_b9c0c74a5bfbedb5(__obf_e7f5496ce1c3e6b2, __obf_f9431224f9a648b3[__obf_6db94c729874d6be:])
}

func __obf_b9c0c74a5bfbedb5(__obf_e7f5496ce1c3e6b2 []byte, __obf_f9431224f9a648b3 []byte) []byte {
	__obf_6ba320e83b6b02fb := 0
	// for the remaining parts, we process them char by char
	var __obf_6db94c729874d6be int
	var __obf_62966ecce30cc903 byte
	for __obf_6db94c729874d6be, __obf_62966ecce30cc903 = range __obf_f9431224f9a648b3 {
		if __obf_62966ecce30cc903 >= utf8.RuneSelf {
			__obf_e7f5496ce1c3e6b2 = append(__obf_e7f5496ce1c3e6b2, '\\', '\\', 'x', __obf_61ea456740bf5ed3[__obf_62966ecce30cc903>>4], __obf_61ea456740bf5ed3[__obf_62966ecce30cc903&0xF])
			__obf_6ba320e83b6b02fb = __obf_6db94c729874d6be + 1
			continue
		}
		if __obf_0eab2aa051a8dd12[__obf_62966ecce30cc903] {
			continue
		}
		if __obf_6ba320e83b6b02fb < __obf_6db94c729874d6be {
			__obf_e7f5496ce1c3e6b2 = append(__obf_e7f5496ce1c3e6b2, __obf_f9431224f9a648b3[__obf_6ba320e83b6b02fb:__obf_6db94c729874d6be]...)
		}
		__obf_e7f5496ce1c3e6b2 = append(__obf_e7f5496ce1c3e6b2, '\\', '\\', 'x', __obf_61ea456740bf5ed3[__obf_62966ecce30cc903>>4], __obf_61ea456740bf5ed3[__obf_62966ecce30cc903&0xF])
		__obf_6ba320e83b6b02fb = __obf_6db94c729874d6be + 1
	}
	if __obf_6ba320e83b6b02fb < len(__obf_f9431224f9a648b3) {
		__obf_e7f5496ce1c3e6b2 = append(__obf_e7f5496ce1c3e6b2, __obf_f9431224f9a648b3[__obf_6ba320e83b6b02fb:]...)
	}
	return append(__obf_e7f5496ce1c3e6b2, '"')
}
