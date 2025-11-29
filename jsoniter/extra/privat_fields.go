package __obf_789dc33d47f4ab2c

import (
	jsoniter "github.com/ArtisanHiram/go-pkg/jsoniter"
	"strings"
	"unicode"
)

// SupportPrivateFields include private fields when encoding/decoding
func SupportPrivateFields() {
	jsoniter.RegisterExtension(&__obf_3293b7ef8520d7c7{})
}

type __obf_3293b7ef8520d7c7 struct {
	jsoniter.DummyExtension
}

func (__obf_369258fb13d82ac5 *__obf_3293b7ef8520d7c7) UpdateStructDescriptor(__obf_0c6523f08d195dd4 *jsoniter.StructDescriptor) {
	for _, __obf_7b97f262d18b7c91 := range __obf_0c6523f08d195dd4.Fields {
		__obf_008a9f3cf4e19a8d := unicode.IsLower(rune(__obf_7b97f262d18b7c91.Field.Name()[0]))
		if __obf_008a9f3cf4e19a8d {
			__obf_ba151e7945153a22, __obf_9136b4f5c8f3c198 := __obf_7b97f262d18b7c91.Field.Tag().Lookup("json")
			if !__obf_9136b4f5c8f3c198 {
				__obf_7b97f262d18b7c91.
					FromNames = []string{__obf_7b97f262d18b7c91.Field.Name()}
				__obf_7b97f262d18b7c91.
					ToNames = []string{__obf_7b97f262d18b7c91.Field.Name()}
				continue
			}
			__obf_1d1c24a713f957c3 := strings.Split(__obf_ba151e7945153a22, ",")
			__obf_49684e759ac4d04b := __obf_c1bc14728875aeba(__obf_7b97f262d18b7c91.Field.Name(), __obf_1d1c24a713f957c3[0], __obf_ba151e7945153a22)
			__obf_7b97f262d18b7c91.
				FromNames = __obf_49684e759ac4d04b
			__obf_7b97f262d18b7c91.
				ToNames = __obf_49684e759ac4d04b
		}
	}
}

func __obf_c1bc14728875aeba(__obf_1e7667f97b5a53e4 string, __obf_007eedc28282801d string, __obf_7de1e1434ed30479 string) []string {
	// ignore?
	if __obf_7de1e1434ed30479 == "-" {
		return []string{}
	}
	// rename?
	var __obf_c2f82ea20f030066 []string
	if __obf_007eedc28282801d == "" {
		__obf_c2f82ea20f030066 = []string{__obf_1e7667f97b5a53e4}
	} else {
		__obf_c2f82ea20f030066 = []string{__obf_007eedc28282801d}
	}
	__obf_33c6fe270d3d2d45 := // private?
		unicode.IsLower(rune(__obf_1e7667f97b5a53e4[0]))
	if __obf_33c6fe270d3d2d45 {
		__obf_c2f82ea20f030066 = []string{}
	}
	return __obf_c2f82ea20f030066
}
