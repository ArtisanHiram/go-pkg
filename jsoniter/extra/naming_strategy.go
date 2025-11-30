package __obf_9a397ef96534ad45

import (
	jsoniter "github.com/ArtisanHiram/go-pkg/jsoniter"
	"strings"
	"unicode"
)

// SetNamingStrategy rename struct fields uniformly
func SetNamingStrategy(__obf_bef596748bbd1441 func(string) string) {
	jsoniter.RegisterExtension(&__obf_2a424a87187ff2a9{jsoniter.DummyExtension{}, __obf_bef596748bbd1441})
}

type __obf_2a424a87187ff2a9 struct {
	jsoniter.DummyExtension
	__obf_bef596748bbd1441 func(string) string
}

func (__obf_454cba947156c7ed *__obf_2a424a87187ff2a9) UpdateStructDescriptor(__obf_44b5443a119ac92f *jsoniter.StructDescriptor) {
	for _, __obf_3c7c779fd3d51e1a := range __obf_44b5443a119ac92f.Fields {
		if unicode.IsLower(rune(__obf_3c7c779fd3d51e1a.Field.Name()[0])) || __obf_3c7c779fd3d51e1a.Field.Name()[0] == '_' {
			continue
		}
		__obf_3d974a4e3139d5d8, __obf_b09e0e0d5c507063 := __obf_3c7c779fd3d51e1a.Field.Tag().Lookup("json")
		if __obf_b09e0e0d5c507063 {
			__obf_9b9c76ccfbe36390 := strings.Split(__obf_3d974a4e3139d5d8, ",")
			if __obf_9b9c76ccfbe36390[0] == "-" {
				continue // hidden field
			}
			if __obf_9b9c76ccfbe36390[0] != "" {
				continue // field explicitly named
			}
		}
		__obf_3c7c779fd3d51e1a.
			ToNames = []string{__obf_454cba947156c7ed.__obf_bef596748bbd1441(__obf_3c7c779fd3d51e1a.Field.Name())}
		__obf_3c7c779fd3d51e1a.
			FromNames = []string{__obf_454cba947156c7ed.__obf_bef596748bbd1441(__obf_3c7c779fd3d51e1a.Field.Name())}
	}
}

// LowerCaseWithUnderscores one strategy to SetNamingStrategy for. It will change HelloWorld to hello_world.
func LowerCaseWithUnderscores(__obf_e393dc6c8c711a30 string) string {
	__obf_310656f5bca34cb1 := []rune{}
	for __obf_47b4fcf90774658f, __obf_08774a099cab6503 := range __obf_e393dc6c8c711a30 {
		if __obf_47b4fcf90774658f == 0 {
			__obf_310656f5bca34cb1 = append(__obf_310656f5bca34cb1, unicode.ToLower(__obf_08774a099cab6503))
		} else {
			if unicode.IsUpper(__obf_08774a099cab6503) {
				__obf_310656f5bca34cb1 = append(__obf_310656f5bca34cb1, '_')
				__obf_310656f5bca34cb1 = append(__obf_310656f5bca34cb1, unicode.ToLower(__obf_08774a099cab6503))
			} else {
				__obf_310656f5bca34cb1 = append(__obf_310656f5bca34cb1, __obf_08774a099cab6503)
			}
		}
	}
	return string(__obf_310656f5bca34cb1)
}
