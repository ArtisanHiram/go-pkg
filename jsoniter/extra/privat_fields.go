package __obf_1f22c6b8dfc77bff

import (
	jsoniter "github.com/ArtisanHiram/go-pkg/jsoniter"
	"strings"
	"unicode"
)

// SupportPrivateFields include private fields when encoding/decoding
func SupportPrivateFields() {
	jsoniter.RegisterExtension(&__obf_9b22217a5850829a{})
}

type __obf_9b22217a5850829a struct {
	jsoniter.DummyExtension
}

func (__obf_9382a03d04135598 *__obf_9b22217a5850829a) UpdateStructDescriptor(__obf_098e889630cd7d27 *jsoniter.StructDescriptor) {
	for _, __obf_423918f677e1b1d0 := range __obf_098e889630cd7d27.Fields {
		__obf_7a5302f0b83c9239 := unicode.IsLower(rune(__obf_423918f677e1b1d0.Field.Name()[0]))
		if __obf_7a5302f0b83c9239 {
			__obf_44f78316a232f68f, __obf_69245d9902c316c9 := __obf_423918f677e1b1d0.Field.Tag().Lookup("json")
			if !__obf_69245d9902c316c9 {
				__obf_423918f677e1b1d0.
					FromNames = []string{__obf_423918f677e1b1d0.Field.Name()}
				__obf_423918f677e1b1d0.
					ToNames = []string{__obf_423918f677e1b1d0.Field.Name()}
				continue
			}
			__obf_1e7ae97a62113f1e := strings.Split(__obf_44f78316a232f68f, ",")
			__obf_22b508d31c9e3f95 := __obf_3736d1208549a51d(__obf_423918f677e1b1d0.Field.Name(), __obf_1e7ae97a62113f1e[0], __obf_44f78316a232f68f)
			__obf_423918f677e1b1d0.
				FromNames = __obf_22b508d31c9e3f95
			__obf_423918f677e1b1d0.
				ToNames = __obf_22b508d31c9e3f95
		}
	}
}

func __obf_3736d1208549a51d(__obf_06294b16b409c1c4 string, __obf_9d91433b71213e06 string, __obf_5dff0486ad795b1e string) []string {
	// ignore?
	if __obf_5dff0486ad795b1e == "-" {
		return []string{}
	}
	// rename?
	var __obf_8d8f443afefb3b2b []string
	if __obf_9d91433b71213e06 == "" {
		__obf_8d8f443afefb3b2b = []string{__obf_06294b16b409c1c4}
	} else {
		__obf_8d8f443afefb3b2b = []string{__obf_9d91433b71213e06}
	}
	__obf_9ad8456f97ff4da6 := // private?
		unicode.IsLower(rune(__obf_06294b16b409c1c4[0]))
	if __obf_9ad8456f97ff4da6 {
		__obf_8d8f443afefb3b2b = []string{}
	}
	return __obf_8d8f443afefb3b2b
}
