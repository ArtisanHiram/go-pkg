package __obf_060749efdc6ad522

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
var __obf_458bb987840c8d86 = [utf8.RuneSelf]bool{
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

var __obf_81335f7d913029d0 = reflect2.TypeOfPtr((*[]byte)(nil)).Elem()

type BinaryAsStringExtension struct {
	jsoniter.DummyExtension
}

func (__obf_1ffe89e9c54d7879 *BinaryAsStringExtension) CreateEncoder(__obf_036de3429d9233a8 reflect2.Type) jsoniter.ValEncoder {
	if __obf_036de3429d9233a8 == __obf_81335f7d913029d0 {
		return &__obf_6380fc5f49aa2bee{}
	}
	return nil
}

func (__obf_1ffe89e9c54d7879 *BinaryAsStringExtension) CreateDecoder(__obf_036de3429d9233a8 reflect2.Type) jsoniter.ValDecoder {
	if __obf_036de3429d9233a8 == __obf_81335f7d913029d0 {
		return &__obf_6380fc5f49aa2bee{}
	}
	return nil
}

type __obf_6380fc5f49aa2bee struct {
}

func (__obf_5b4176d65d17beeb *__obf_6380fc5f49aa2bee) Decode(__obf_deff95ef54922509 unsafe.Pointer, __obf_7eafecc31dc47100 *jsoniter.Iterator) {
	__obf_812f4ae8d1b54953 := __obf_7eafecc31dc47100.ReadStringAsSlice()
	__obf_2e09cc1a0ba5306e := make([]byte, 0, len(__obf_812f4ae8d1b54953))
	for __obf_9e607be55d8e3092 := 0; __obf_9e607be55d8e3092 < len(__obf_812f4ae8d1b54953); __obf_9e607be55d8e3092++ {
		__obf_4e9fe63935c10129 := __obf_812f4ae8d1b54953[__obf_9e607be55d8e3092]
		if __obf_4e9fe63935c10129 == '\\' {
			__obf_9cfea7136645448b := __obf_812f4ae8d1b54953[__obf_9e607be55d8e3092+1]
			if __obf_9cfea7136645448b != '\\' {
				__obf_7eafecc31dc47100.
					ReportError("decode binary as string", `\\x is only supported escape`)
				return
			}
			__obf_c343aebc4aaac55b := __obf_812f4ae8d1b54953[__obf_9e607be55d8e3092+2]
			if __obf_c343aebc4aaac55b != 'x' {
				__obf_7eafecc31dc47100.
					ReportError("decode binary as string", `\\x is only supported escape`)
				return
			}
			__obf_d791ad46f7b4dec4 := __obf_812f4ae8d1b54953[__obf_9e607be55d8e3092+3]
			__obf_6b1c8dd85cf2fe99 := __obf_812f4ae8d1b54953[__obf_9e607be55d8e3092+4]
			__obf_9e607be55d8e3092 += 4
			__obf_4e9fe63935c10129 = __obf_723c9dd124eadbb0(__obf_7eafecc31dc47100, __obf_d791ad46f7b4dec4, __obf_6b1c8dd85cf2fe99)
		}
		__obf_2e09cc1a0ba5306e = append(__obf_2e09cc1a0ba5306e, __obf_4e9fe63935c10129)
	}
	*(*[]byte)(__obf_deff95ef54922509) = __obf_2e09cc1a0ba5306e
}
func (__obf_5b4176d65d17beeb *__obf_6380fc5f49aa2bee) IsEmpty(__obf_deff95ef54922509 unsafe.Pointer) bool {
	return len(*((*[]byte)(__obf_deff95ef54922509))) == 0
}
func (__obf_5b4176d65d17beeb *__obf_6380fc5f49aa2bee) Encode(__obf_deff95ef54922509 unsafe.Pointer, __obf_d9afe3b0234a705f *jsoniter.Stream) {
	__obf_36e5f64c7d941c45 := __obf_7a27b73696e841d3(__obf_d9afe3b0234a705f.Buffer(), *(*[]byte)(__obf_deff95ef54922509))
	__obf_d9afe3b0234a705f.
		SetBuffer(__obf_36e5f64c7d941c45)
}

func __obf_723c9dd124eadbb0(__obf_7eafecc31dc47100 *jsoniter.Iterator, __obf_02eb07db493b4883, __obf_9cfea7136645448b byte) byte {
	var __obf_a55ddbc21711e6e6 byte
	if __obf_02eb07db493b4883 >= '0' && __obf_02eb07db493b4883 <= '9' {
		__obf_a55ddbc21711e6e6 = __obf_02eb07db493b4883 - '0'
	} else if __obf_02eb07db493b4883 >= 'a' && __obf_02eb07db493b4883 <= 'f' {
		__obf_a55ddbc21711e6e6 = __obf_02eb07db493b4883 - 'a' + 10
	} else {
		__obf_7eafecc31dc47100.
			ReportError("read hex", "expects 0~9 or a~f, but found "+string([]byte{__obf_02eb07db493b4883}))
		return 0
	}
	__obf_a55ddbc21711e6e6 *= 16
	if __obf_9cfea7136645448b >= '0' && __obf_9cfea7136645448b <= '9' {
		__obf_a55ddbc21711e6e6 += __obf_9cfea7136645448b - '0'
	} else if __obf_9cfea7136645448b >= 'a' && __obf_9cfea7136645448b <= 'f' {
		__obf_a55ddbc21711e6e6 += __obf_9cfea7136645448b - 'a' + 10
	} else {
		__obf_7eafecc31dc47100.
			ReportError("read hex", "expects 0~9 or a~f, but found "+string([]byte{__obf_9cfea7136645448b}))
		return 0
	}
	return __obf_a55ddbc21711e6e6
}

var __obf_591ee55a5d75d6ba = "0123456789abcdef"

func __obf_7a27b73696e841d3(__obf_b0ffe8e7c94a06dd []byte, __obf_335a8301c054fd1f []byte) []byte {
	__obf_b0ffe8e7c94a06dd = append(__obf_b0ffe8e7c94a06dd, '"')
	// write string, the fast path, without utf8 and escape support
	var __obf_9e607be55d8e3092 int
	var __obf_e5ca5e05b9822fee byte
	for __obf_9e607be55d8e3092, __obf_e5ca5e05b9822fee = range __obf_335a8301c054fd1f {
		if __obf_e5ca5e05b9822fee < utf8.RuneSelf && __obf_458bb987840c8d86[__obf_e5ca5e05b9822fee] {
			__obf_b0ffe8e7c94a06dd = append(__obf_b0ffe8e7c94a06dd, __obf_e5ca5e05b9822fee)
		} else {
			break
		}
	}
	if __obf_9e607be55d8e3092 == len(__obf_335a8301c054fd1f)-1 {
		__obf_b0ffe8e7c94a06dd = append(__obf_b0ffe8e7c94a06dd, '"')
		return __obf_b0ffe8e7c94a06dd
	}
	return __obf_ca3657c3fe783da7(__obf_b0ffe8e7c94a06dd, __obf_335a8301c054fd1f[__obf_9e607be55d8e3092:])
}

func __obf_ca3657c3fe783da7(__obf_b0ffe8e7c94a06dd []byte, __obf_335a8301c054fd1f []byte) []byte {
	__obf_f5388c04f667abf3 := 0
	// for the remaining parts, we process them char by char
	var __obf_9e607be55d8e3092 int
	var __obf_4e9fe63935c10129 byte
	for __obf_9e607be55d8e3092, __obf_4e9fe63935c10129 = range __obf_335a8301c054fd1f {
		if __obf_4e9fe63935c10129 >= utf8.RuneSelf {
			__obf_b0ffe8e7c94a06dd = append(__obf_b0ffe8e7c94a06dd, '\\', '\\', 'x', __obf_591ee55a5d75d6ba[__obf_4e9fe63935c10129>>4], __obf_591ee55a5d75d6ba[__obf_4e9fe63935c10129&0xF])
			__obf_f5388c04f667abf3 = __obf_9e607be55d8e3092 + 1
			continue
		}
		if __obf_458bb987840c8d86[__obf_4e9fe63935c10129] {
			continue
		}
		if __obf_f5388c04f667abf3 < __obf_9e607be55d8e3092 {
			__obf_b0ffe8e7c94a06dd = append(__obf_b0ffe8e7c94a06dd, __obf_335a8301c054fd1f[__obf_f5388c04f667abf3:__obf_9e607be55d8e3092]...)
		}
		__obf_b0ffe8e7c94a06dd = append(__obf_b0ffe8e7c94a06dd, '\\', '\\', 'x', __obf_591ee55a5d75d6ba[__obf_4e9fe63935c10129>>4], __obf_591ee55a5d75d6ba[__obf_4e9fe63935c10129&0xF])
		__obf_f5388c04f667abf3 = __obf_9e607be55d8e3092 + 1
	}
	if __obf_f5388c04f667abf3 < len(__obf_335a8301c054fd1f) {
		__obf_b0ffe8e7c94a06dd = append(__obf_b0ffe8e7c94a06dd, __obf_335a8301c054fd1f[__obf_f5388c04f667abf3:]...)
	}
	return append(__obf_b0ffe8e7c94a06dd, '"')
}
