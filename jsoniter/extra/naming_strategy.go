package __obf_060749efdc6ad522

import (
	jsoniter "github.com/ArtisanHiram/go-pkg/jsoniter"
	"strings"
	"unicode"
)

// SetNamingStrategy rename struct fields uniformly
func SetNamingStrategy(__obf_230467da7ea96415 func(string) string) {
	jsoniter.RegisterExtension(&__obf_e80ef1c9cb107176{jsoniter.DummyExtension{}, __obf_230467da7ea96415})
}

type __obf_e80ef1c9cb107176 struct {
	jsoniter.DummyExtension
	__obf_230467da7ea96415 func(string) string
}

func (__obf_1ffe89e9c54d7879 *__obf_e80ef1c9cb107176) UpdateStructDescriptor(__obf_d7a80073d1dfd7d7 *jsoniter.StructDescriptor) {
	for _, __obf_88585cf978e0c17e := range __obf_d7a80073d1dfd7d7.Fields {
		if unicode.IsLower(rune(__obf_88585cf978e0c17e.Field.Name()[0])) || __obf_88585cf978e0c17e.Field.Name()[0] == '_' {
			continue
		}
		__obf_5c03b64bc5e37fc7, __obf_ca50ba45dceff983 := __obf_88585cf978e0c17e.Field.Tag().Lookup("json")
		if __obf_ca50ba45dceff983 {
			__obf_cc68dfcdabdda8bc := strings.Split(__obf_5c03b64bc5e37fc7, ",")
			if __obf_cc68dfcdabdda8bc[0] == "-" {
				continue // hidden field
			}
			if __obf_cc68dfcdabdda8bc[0] != "" {
				continue // field explicitly named
			}
		}
		__obf_88585cf978e0c17e.
			ToNames = []string{__obf_1ffe89e9c54d7879.__obf_230467da7ea96415(__obf_88585cf978e0c17e.Field.Name())}
		__obf_88585cf978e0c17e.
			FromNames = []string{__obf_1ffe89e9c54d7879.__obf_230467da7ea96415(__obf_88585cf978e0c17e.Field.Name())}
	}
}

// LowerCaseWithUnderscores one strategy to SetNamingStrategy for. It will change HelloWorld to hello_world.
func LowerCaseWithUnderscores(__obf_49ffc53ac93b13d0 string) string {
	__obf_7ca81e1200cb049d := []rune{}
	for __obf_9e607be55d8e3092, __obf_e5ca5e05b9822fee := range __obf_49ffc53ac93b13d0 {
		if __obf_9e607be55d8e3092 == 0 {
			__obf_7ca81e1200cb049d = append(__obf_7ca81e1200cb049d, unicode.ToLower(__obf_e5ca5e05b9822fee))
		} else {
			if unicode.IsUpper(__obf_e5ca5e05b9822fee) {
				__obf_7ca81e1200cb049d = append(__obf_7ca81e1200cb049d, '_')
				__obf_7ca81e1200cb049d = append(__obf_7ca81e1200cb049d, unicode.ToLower(__obf_e5ca5e05b9822fee))
			} else {
				__obf_7ca81e1200cb049d = append(__obf_7ca81e1200cb049d, __obf_e5ca5e05b9822fee)
			}
		}
	}
	return string(__obf_7ca81e1200cb049d)
}
