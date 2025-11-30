package __obf_38f1d2f091ad74e0

import (
	jsoniter "github.com/ArtisanHiram/go-pkg/jsoniter"
	"strings"
	"unicode"
)

// SetNamingStrategy rename struct fields uniformly
func SetNamingStrategy(__obf_b68f4dcdf654054f func(string) string) {
	jsoniter.RegisterExtension(&__obf_b5bbf868a272dcc8{jsoniter.DummyExtension{}, __obf_b68f4dcdf654054f})
}

type __obf_b5bbf868a272dcc8 struct {
	jsoniter.DummyExtension
	__obf_b68f4dcdf654054f func(string) string
}

func (__obf_9c280f4c530aab7f *__obf_b5bbf868a272dcc8) UpdateStructDescriptor(__obf_cf80c64718a9d630 *jsoniter.StructDescriptor) {
	for _, __obf_5ac789aeaf7879e1 := range __obf_cf80c64718a9d630.Fields {
		if unicode.IsLower(rune(__obf_5ac789aeaf7879e1.Field.Name()[0])) || __obf_5ac789aeaf7879e1.Field.Name()[0] == '_' {
			continue
		}
		__obf_8f18cfca726ca2b1, __obf_691758d45d138fc7 := __obf_5ac789aeaf7879e1.Field.Tag().Lookup("json")
		if __obf_691758d45d138fc7 {
			__obf_86cb161c59473cad := strings.Split(__obf_8f18cfca726ca2b1, ",")
			if __obf_86cb161c59473cad[0] == "-" {
				continue // hidden field
			}
			if __obf_86cb161c59473cad[0] != "" {
				continue // field explicitly named
			}
		}
		__obf_5ac789aeaf7879e1.
			ToNames = []string{__obf_9c280f4c530aab7f.__obf_b68f4dcdf654054f(__obf_5ac789aeaf7879e1.Field.Name())}
		__obf_5ac789aeaf7879e1.
			FromNames = []string{__obf_9c280f4c530aab7f.__obf_b68f4dcdf654054f(__obf_5ac789aeaf7879e1.Field.Name())}
	}
}

// LowerCaseWithUnderscores one strategy to SetNamingStrategy for. It will change HelloWorld to hello_world.
func LowerCaseWithUnderscores(__obf_e3c0131bb9072d3b string) string {
	__obf_6a6b3734dcfb5f64 := []rune{}
	for __obf_41047c070beaa4ab, __obf_412744d53cf48990 := range __obf_e3c0131bb9072d3b {
		if __obf_41047c070beaa4ab == 0 {
			__obf_6a6b3734dcfb5f64 = append(__obf_6a6b3734dcfb5f64, unicode.ToLower(__obf_412744d53cf48990))
		} else {
			if unicode.IsUpper(__obf_412744d53cf48990) {
				__obf_6a6b3734dcfb5f64 = append(__obf_6a6b3734dcfb5f64, '_')
				__obf_6a6b3734dcfb5f64 = append(__obf_6a6b3734dcfb5f64, unicode.ToLower(__obf_412744d53cf48990))
			} else {
				__obf_6a6b3734dcfb5f64 = append(__obf_6a6b3734dcfb5f64, __obf_412744d53cf48990)
			}
		}
	}
	return string(__obf_6a6b3734dcfb5f64)
}
