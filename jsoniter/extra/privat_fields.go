package __obf_060749efdc6ad522

import (
	jsoniter "github.com/ArtisanHiram/go-pkg/jsoniter"
	"strings"
	"unicode"
)

// SupportPrivateFields include private fields when encoding/decoding
func SupportPrivateFields() {
	jsoniter.RegisterExtension(&__obf_e285c56e9f9e8dfa{})
}

type __obf_e285c56e9f9e8dfa struct {
	jsoniter.DummyExtension
}

func (__obf_1ffe89e9c54d7879 *__obf_e285c56e9f9e8dfa) UpdateStructDescriptor(__obf_d7a80073d1dfd7d7 *jsoniter.StructDescriptor) {
	for _, __obf_88585cf978e0c17e := range __obf_d7a80073d1dfd7d7.Fields {
		__obf_ba72d0834cfde195 := unicode.IsLower(rune(__obf_88585cf978e0c17e.Field.Name()[0]))
		if __obf_ba72d0834cfde195 {
			__obf_5c03b64bc5e37fc7, __obf_ca50ba45dceff983 := __obf_88585cf978e0c17e.Field.Tag().Lookup("json")
			if !__obf_ca50ba45dceff983 {
				__obf_88585cf978e0c17e.
					FromNames = []string{__obf_88585cf978e0c17e.Field.Name()}
				__obf_88585cf978e0c17e.
					ToNames = []string{__obf_88585cf978e0c17e.Field.Name()}
				continue
			}
			__obf_cc68dfcdabdda8bc := strings.Split(__obf_5c03b64bc5e37fc7, ",")
			__obf_eb7baabb5b71d1d0 := __obf_4e7c4f6fede743a6(__obf_88585cf978e0c17e.Field.Name(), __obf_cc68dfcdabdda8bc[0], __obf_5c03b64bc5e37fc7)
			__obf_88585cf978e0c17e.
				FromNames = __obf_eb7baabb5b71d1d0
			__obf_88585cf978e0c17e.
				ToNames = __obf_eb7baabb5b71d1d0
		}
	}
}

func __obf_4e7c4f6fede743a6(__obf_095e834ba26620a1 string, __obf_cdbb8d3ed65466e2 string, __obf_fee925dda7b66007 string) []string {
	// ignore?
	if __obf_fee925dda7b66007 == "-" {
		return []string{}
	}
	// rename?
	var __obf_f1e5c5a48e90b312 []string
	if __obf_cdbb8d3ed65466e2 == "" {
		__obf_f1e5c5a48e90b312 = []string{__obf_095e834ba26620a1}
	} else {
		__obf_f1e5c5a48e90b312 = []string{__obf_cdbb8d3ed65466e2}
	}
	__obf_10bea95324e25b5d := // private?
		unicode.IsLower(rune(__obf_095e834ba26620a1[0]))
	if __obf_10bea95324e25b5d {
		__obf_f1e5c5a48e90b312 = []string{}
	}
	return __obf_f1e5c5a48e90b312
}
