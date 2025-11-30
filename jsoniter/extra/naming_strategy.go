package __obf_1f22c6b8dfc77bff

import (
	jsoniter "github.com/ArtisanHiram/go-pkg/jsoniter"
	"strings"
	"unicode"
)

// SetNamingStrategy rename struct fields uniformly
func SetNamingStrategy(__obf_e2b6bf0cec6020a4 func(string) string) {
	jsoniter.RegisterExtension(&__obf_4749c3e7b81b324c{jsoniter.DummyExtension{}, __obf_e2b6bf0cec6020a4})
}

type __obf_4749c3e7b81b324c struct {
	jsoniter.DummyExtension
	__obf_e2b6bf0cec6020a4 func(string) string
}

func (__obf_9382a03d04135598 *__obf_4749c3e7b81b324c) UpdateStructDescriptor(__obf_098e889630cd7d27 *jsoniter.StructDescriptor) {
	for _, __obf_423918f677e1b1d0 := range __obf_098e889630cd7d27.Fields {
		if unicode.IsLower(rune(__obf_423918f677e1b1d0.Field.Name()[0])) || __obf_423918f677e1b1d0.Field.Name()[0] == '_' {
			continue
		}
		__obf_44f78316a232f68f, __obf_69245d9902c316c9 := __obf_423918f677e1b1d0.Field.Tag().Lookup("json")
		if __obf_69245d9902c316c9 {
			__obf_1e7ae97a62113f1e := strings.Split(__obf_44f78316a232f68f, ",")
			if __obf_1e7ae97a62113f1e[0] == "-" {
				continue // hidden field
			}
			if __obf_1e7ae97a62113f1e[0] != "" {
				continue // field explicitly named
			}
		}
		__obf_423918f677e1b1d0.
			ToNames = []string{__obf_9382a03d04135598.__obf_e2b6bf0cec6020a4(__obf_423918f677e1b1d0.Field.Name())}
		__obf_423918f677e1b1d0.
			FromNames = []string{__obf_9382a03d04135598.__obf_e2b6bf0cec6020a4(__obf_423918f677e1b1d0.Field.Name())}
	}
}

// LowerCaseWithUnderscores one strategy to SetNamingStrategy for. It will change HelloWorld to hello_world.
func LowerCaseWithUnderscores(__obf_d0acfe041546e537 string) string {
	__obf_e1a7dc2cb1c8cf76 := []rune{}
	for __obf_f52abe4fb4613d80, __obf_f32e4134794a18a0 := range __obf_d0acfe041546e537 {
		if __obf_f52abe4fb4613d80 == 0 {
			__obf_e1a7dc2cb1c8cf76 = append(__obf_e1a7dc2cb1c8cf76, unicode.ToLower(__obf_f32e4134794a18a0))
		} else {
			if unicode.IsUpper(__obf_f32e4134794a18a0) {
				__obf_e1a7dc2cb1c8cf76 = append(__obf_e1a7dc2cb1c8cf76, '_')
				__obf_e1a7dc2cb1c8cf76 = append(__obf_e1a7dc2cb1c8cf76, unicode.ToLower(__obf_f32e4134794a18a0))
			} else {
				__obf_e1a7dc2cb1c8cf76 = append(__obf_e1a7dc2cb1c8cf76, __obf_f32e4134794a18a0)
			}
		}
	}
	return string(__obf_e1a7dc2cb1c8cf76)
}
