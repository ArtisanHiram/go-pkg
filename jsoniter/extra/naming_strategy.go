package __obf_eed9c5a643743c33

import (
	jsoniter "github.com/ArtisanHiram/go-pkg/jsoniter"
	"strings"
	"unicode"
)

// SetNamingStrategy rename struct fields uniformly
func SetNamingStrategy(__obf_3c4096e0431a0c5e func(string) string) {
	jsoniter.RegisterExtension(&__obf_83e0d91bc178ed68{jsoniter.DummyExtension{}, __obf_3c4096e0431a0c5e})
}

type __obf_83e0d91bc178ed68 struct {
	jsoniter.DummyExtension
	__obf_3c4096e0431a0c5e func(string) string
}

func (__obf_7ab69ccfb0d084de *__obf_83e0d91bc178ed68) UpdateStructDescriptor(__obf_329fcbf95fe89472 *jsoniter.StructDescriptor) {
	for _, __obf_7ab709a6862eec48 := range __obf_329fcbf95fe89472.Fields {
		if unicode.IsLower(rune(__obf_7ab709a6862eec48.Field.Name()[0])) || __obf_7ab709a6862eec48.Field.Name()[0] == '_' {
			continue
		}
		__obf_2b0695717d2ca1a3, __obf_e2ef35e6eb4b8b83 := __obf_7ab709a6862eec48.Field.Tag().Lookup("json")
		if __obf_e2ef35e6eb4b8b83 {
			__obf_843475ca8e35a946 := strings.Split(__obf_2b0695717d2ca1a3, ",")
			if __obf_843475ca8e35a946[0] == "-" {
				continue // hidden field
			}
			if __obf_843475ca8e35a946[0] != "" {
				continue // field explicitly named
			}
		}
		__obf_7ab709a6862eec48.
			ToNames = []string{__obf_7ab69ccfb0d084de.__obf_3c4096e0431a0c5e(__obf_7ab709a6862eec48.Field.Name())}
		__obf_7ab709a6862eec48.
			FromNames = []string{__obf_7ab69ccfb0d084de.__obf_3c4096e0431a0c5e(__obf_7ab709a6862eec48.Field.Name())}
	}
}

// LowerCaseWithUnderscores one strategy to SetNamingStrategy for. It will change HelloWorld to hello_world.
func LowerCaseWithUnderscores(__obf_84fe317de61750b1 string) string {
	__obf_f8b6d7773b8c42a1 := []rune{}
	for __obf_6db94c729874d6be, __obf_d609c92bf973ccdd := range __obf_84fe317de61750b1 {
		if __obf_6db94c729874d6be == 0 {
			__obf_f8b6d7773b8c42a1 = append(__obf_f8b6d7773b8c42a1, unicode.ToLower(__obf_d609c92bf973ccdd))
		} else {
			if unicode.IsUpper(__obf_d609c92bf973ccdd) {
				__obf_f8b6d7773b8c42a1 = append(__obf_f8b6d7773b8c42a1, '_')
				__obf_f8b6d7773b8c42a1 = append(__obf_f8b6d7773b8c42a1, unicode.ToLower(__obf_d609c92bf973ccdd))
			} else {
				__obf_f8b6d7773b8c42a1 = append(__obf_f8b6d7773b8c42a1, __obf_d609c92bf973ccdd)
			}
		}
	}
	return string(__obf_f8b6d7773b8c42a1)
}
