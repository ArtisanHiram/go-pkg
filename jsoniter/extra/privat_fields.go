package __obf_38f1d2f091ad74e0

import (
	jsoniter "github.com/ArtisanHiram/go-pkg/jsoniter"
	"strings"
	"unicode"
)

// SupportPrivateFields include private fields when encoding/decoding
func SupportPrivateFields() {
	jsoniter.RegisterExtension(&__obf_cd6b67693812e1cb{})
}

type __obf_cd6b67693812e1cb struct {
	jsoniter.DummyExtension
}

func (__obf_9c280f4c530aab7f *__obf_cd6b67693812e1cb) UpdateStructDescriptor(__obf_cf80c64718a9d630 *jsoniter.StructDescriptor) {
	for _, __obf_5ac789aeaf7879e1 := range __obf_cf80c64718a9d630.Fields {
		__obf_3b28129690e3ff31 := unicode.IsLower(rune(__obf_5ac789aeaf7879e1.Field.Name()[0]))
		if __obf_3b28129690e3ff31 {
			__obf_8f18cfca726ca2b1, __obf_691758d45d138fc7 := __obf_5ac789aeaf7879e1.Field.Tag().Lookup("json")
			if !__obf_691758d45d138fc7 {
				__obf_5ac789aeaf7879e1.
					FromNames = []string{__obf_5ac789aeaf7879e1.Field.Name()}
				__obf_5ac789aeaf7879e1.
					ToNames = []string{__obf_5ac789aeaf7879e1.Field.Name()}
				continue
			}
			__obf_86cb161c59473cad := strings.Split(__obf_8f18cfca726ca2b1, ",")
			__obf_218caee415bee8ff := __obf_b75177ede013c874(__obf_5ac789aeaf7879e1.Field.Name(), __obf_86cb161c59473cad[0], __obf_8f18cfca726ca2b1)
			__obf_5ac789aeaf7879e1.
				FromNames = __obf_218caee415bee8ff
			__obf_5ac789aeaf7879e1.
				ToNames = __obf_218caee415bee8ff
		}
	}
}

func __obf_b75177ede013c874(__obf_185e5ff1fc2038f8 string, __obf_34e8a5e6af542b28 string, __obf_1a53c39df833133c string) []string {
	// ignore?
	if __obf_1a53c39df833133c == "-" {
		return []string{}
	}
	// rename?
	var __obf_f9ed816a19e02856 []string
	if __obf_34e8a5e6af542b28 == "" {
		__obf_f9ed816a19e02856 = []string{__obf_185e5ff1fc2038f8}
	} else {
		__obf_f9ed816a19e02856 = []string{__obf_34e8a5e6af542b28}
	}
	__obf_3bd272a3c35bbc8f := // private?
		unicode.IsLower(rune(__obf_185e5ff1fc2038f8[0]))
	if __obf_3bd272a3c35bbc8f {
		__obf_f9ed816a19e02856 = []string{}
	}
	return __obf_f9ed816a19e02856
}
