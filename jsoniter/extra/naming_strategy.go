package __obf_789dc33d47f4ab2c

import (
	jsoniter "github.com/ArtisanHiram/go-pkg/jsoniter"
	"strings"
	"unicode"
)

// SetNamingStrategy rename struct fields uniformly
func SetNamingStrategy(__obf_9bb5e6c978141659 func(string) string) {
	jsoniter.RegisterExtension(&__obf_2b8c89a8d54c45fb{jsoniter.DummyExtension{}, __obf_9bb5e6c978141659})
}

type __obf_2b8c89a8d54c45fb struct {
	jsoniter.DummyExtension
	__obf_9bb5e6c978141659 func(string) string
}

func (__obf_369258fb13d82ac5 *__obf_2b8c89a8d54c45fb) UpdateStructDescriptor(__obf_0c6523f08d195dd4 *jsoniter.StructDescriptor) {
	for _, __obf_7b97f262d18b7c91 := range __obf_0c6523f08d195dd4.Fields {
		if unicode.IsLower(rune(__obf_7b97f262d18b7c91.Field.Name()[0])) || __obf_7b97f262d18b7c91.Field.Name()[0] == '_' {
			continue
		}
		__obf_ba151e7945153a22, __obf_9136b4f5c8f3c198 := __obf_7b97f262d18b7c91.Field.Tag().Lookup("json")
		if __obf_9136b4f5c8f3c198 {
			__obf_1d1c24a713f957c3 := strings.Split(__obf_ba151e7945153a22, ",")
			if __obf_1d1c24a713f957c3[0] == "-" {
				continue // hidden field
			}
			if __obf_1d1c24a713f957c3[0] != "" {
				continue // field explicitly named
			}
		}
		__obf_7b97f262d18b7c91.
			ToNames = []string{__obf_369258fb13d82ac5.__obf_9bb5e6c978141659(__obf_7b97f262d18b7c91.Field.Name())}
		__obf_7b97f262d18b7c91.
			FromNames = []string{__obf_369258fb13d82ac5.__obf_9bb5e6c978141659(__obf_7b97f262d18b7c91.Field.Name())}
	}
}

// LowerCaseWithUnderscores one strategy to SetNamingStrategy for. It will change HelloWorld to hello_world.
func LowerCaseWithUnderscores(__obf_e9274217ff0621b6 string) string {
	__obf_114d7264c618cd3c := []rune{}
	for __obf_c1c90a275d52d9e6, __obf_c39e01a0b80979b1 := range __obf_e9274217ff0621b6 {
		if __obf_c1c90a275d52d9e6 == 0 {
			__obf_114d7264c618cd3c = append(__obf_114d7264c618cd3c, unicode.ToLower(__obf_c39e01a0b80979b1))
		} else {
			if unicode.IsUpper(__obf_c39e01a0b80979b1) {
				__obf_114d7264c618cd3c = append(__obf_114d7264c618cd3c, '_')
				__obf_114d7264c618cd3c = append(__obf_114d7264c618cd3c, unicode.ToLower(__obf_c39e01a0b80979b1))
			} else {
				__obf_114d7264c618cd3c = append(__obf_114d7264c618cd3c, __obf_c39e01a0b80979b1)
			}
		}
	}
	return string(__obf_114d7264c618cd3c)
}
