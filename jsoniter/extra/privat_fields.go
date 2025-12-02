package __obf_eed9c5a643743c33

import (
	jsoniter "github.com/ArtisanHiram/go-pkg/jsoniter"
	"strings"
	"unicode"
)

// SupportPrivateFields include private fields when encoding/decoding
func SupportPrivateFields() {
	jsoniter.RegisterExtension(&__obf_e29fe8a5dfe68243{})
}

type __obf_e29fe8a5dfe68243 struct {
	jsoniter.DummyExtension
}

func (__obf_7ab69ccfb0d084de *__obf_e29fe8a5dfe68243) UpdateStructDescriptor(__obf_329fcbf95fe89472 *jsoniter.StructDescriptor) {
	for _, __obf_7ab709a6862eec48 := range __obf_329fcbf95fe89472.Fields {
		__obf_2611f3af652cc264 := unicode.IsLower(rune(__obf_7ab709a6862eec48.Field.Name()[0]))
		if __obf_2611f3af652cc264 {
			__obf_2b0695717d2ca1a3, __obf_e2ef35e6eb4b8b83 := __obf_7ab709a6862eec48.Field.Tag().Lookup("json")
			if !__obf_e2ef35e6eb4b8b83 {
				__obf_7ab709a6862eec48.
					FromNames = []string{__obf_7ab709a6862eec48.Field.Name()}
				__obf_7ab709a6862eec48.
					ToNames = []string{__obf_7ab709a6862eec48.Field.Name()}
				continue
			}
			__obf_843475ca8e35a946 := strings.Split(__obf_2b0695717d2ca1a3, ",")
			__obf_a8feb41f36b0c166 := __obf_9283e5d76cc46717(__obf_7ab709a6862eec48.Field.Name(), __obf_843475ca8e35a946[0], __obf_2b0695717d2ca1a3)
			__obf_7ab709a6862eec48.
				FromNames = __obf_a8feb41f36b0c166
			__obf_7ab709a6862eec48.
				ToNames = __obf_a8feb41f36b0c166
		}
	}
}

func __obf_9283e5d76cc46717(__obf_dca459323daad3c5 string, __obf_a885d8a803ede2fb string, __obf_bdebc4b1a3ae526d string) []string {
	// ignore?
	if __obf_bdebc4b1a3ae526d == "-" {
		return []string{}
	}
	// rename?
	var __obf_9fda920569c2f7b3 []string
	if __obf_a885d8a803ede2fb == "" {
		__obf_9fda920569c2f7b3 = []string{__obf_dca459323daad3c5}
	} else {
		__obf_9fda920569c2f7b3 = []string{__obf_a885d8a803ede2fb}
	}
	__obf_dbd177e186d76a12 := // private?
		unicode.IsLower(rune(__obf_dca459323daad3c5[0]))
	if __obf_dbd177e186d76a12 {
		__obf_9fda920569c2f7b3 = []string{}
	}
	return __obf_9fda920569c2f7b3
}
