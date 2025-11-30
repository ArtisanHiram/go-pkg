package __obf_9a397ef96534ad45

import (
	jsoniter "github.com/ArtisanHiram/go-pkg/jsoniter"
	"strings"
	"unicode"
)

// SupportPrivateFields include private fields when encoding/decoding
func SupportPrivateFields() {
	jsoniter.RegisterExtension(&__obf_0c6963813fdfb0d9{})
}

type __obf_0c6963813fdfb0d9 struct {
	jsoniter.DummyExtension
}

func (__obf_454cba947156c7ed *__obf_0c6963813fdfb0d9) UpdateStructDescriptor(__obf_44b5443a119ac92f *jsoniter.StructDescriptor) {
	for _, __obf_3c7c779fd3d51e1a := range __obf_44b5443a119ac92f.Fields {
		__obf_cd6cf7a381e55119 := unicode.IsLower(rune(__obf_3c7c779fd3d51e1a.Field.Name()[0]))
		if __obf_cd6cf7a381e55119 {
			__obf_3d974a4e3139d5d8, __obf_b09e0e0d5c507063 := __obf_3c7c779fd3d51e1a.Field.Tag().Lookup("json")
			if !__obf_b09e0e0d5c507063 {
				__obf_3c7c779fd3d51e1a.
					FromNames = []string{__obf_3c7c779fd3d51e1a.Field.Name()}
				__obf_3c7c779fd3d51e1a.
					ToNames = []string{__obf_3c7c779fd3d51e1a.Field.Name()}
				continue
			}
			__obf_9b9c76ccfbe36390 := strings.Split(__obf_3d974a4e3139d5d8, ",")
			__obf_b787d8eef1d5d957 := __obf_1ad0312523b7ae84(__obf_3c7c779fd3d51e1a.Field.Name(), __obf_9b9c76ccfbe36390[0], __obf_3d974a4e3139d5d8)
			__obf_3c7c779fd3d51e1a.
				FromNames = __obf_b787d8eef1d5d957
			__obf_3c7c779fd3d51e1a.
				ToNames = __obf_b787d8eef1d5d957
		}
	}
}

func __obf_1ad0312523b7ae84(__obf_635724196001dd2a string, __obf_628ac40cb26c672f string, __obf_76722ee82ba114cd string) []string {
	// ignore?
	if __obf_76722ee82ba114cd == "-" {
		return []string{}
	}
	// rename?
	var __obf_8e0e35b4a5d4bad6 []string
	if __obf_628ac40cb26c672f == "" {
		__obf_8e0e35b4a5d4bad6 = []string{__obf_635724196001dd2a}
	} else {
		__obf_8e0e35b4a5d4bad6 = []string{__obf_628ac40cb26c672f}
	}
	__obf_c63e51f7a514b9cf := // private?
		unicode.IsLower(rune(__obf_635724196001dd2a[0]))
	if __obf_c63e51f7a514b9cf {
		__obf_8e0e35b4a5d4bad6 = []string{}
	}
	return __obf_8e0e35b4a5d4bad6
}
