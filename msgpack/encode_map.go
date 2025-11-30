package __obf_3e0c303610a19bd4

import (
	"math"
	"reflect"
	"sort"
)

func __obf_10a19aa600413a97(__obf_77240dc7776b60b4 *Encoder, __obf_63bbcee86d44fdde reflect.Value) error {
	if __obf_63bbcee86d44fdde.IsNil() {
		return __obf_77240dc7776b60b4.EncodeNil()
	}

	if __obf_8882f6cf6e378ded := __obf_77240dc7776b60b4.EncodeMapLen(__obf_63bbcee86d44fdde.Len()); __obf_8882f6cf6e378ded != nil {
		return __obf_8882f6cf6e378ded
	}
	__obf_dc466dd663d7625d := __obf_63bbcee86d44fdde.MapRange()
	for __obf_dc466dd663d7625d.Next() {
		if __obf_8882f6cf6e378ded := __obf_77240dc7776b60b4.EncodeValue(__obf_dc466dd663d7625d.Key()); __obf_8882f6cf6e378ded != nil {
			return __obf_8882f6cf6e378ded
		}
		if __obf_8882f6cf6e378ded := __obf_77240dc7776b60b4.EncodeValue(__obf_dc466dd663d7625d.Value()); __obf_8882f6cf6e378ded != nil {
			return __obf_8882f6cf6e378ded
		}
	}

	return nil
}

func __obf_6666d9a2247b6097(__obf_77240dc7776b60b4 *Encoder, __obf_63bbcee86d44fdde reflect.Value) error {
	if __obf_63bbcee86d44fdde.IsNil() {
		return __obf_77240dc7776b60b4.EncodeNil()
	}

	if __obf_8882f6cf6e378ded := __obf_77240dc7776b60b4.EncodeMapLen(__obf_63bbcee86d44fdde.Len()); __obf_8882f6cf6e378ded != nil {
		return __obf_8882f6cf6e378ded
	}
	__obf_b416b6a4555be5c8 := __obf_63bbcee86d44fdde.Convert(__obf_2f37cd23e2c01e1a).Interface().(map[string]bool)
	if __obf_77240dc7776b60b4.__obf_3cf0882fa5a4cafb&__obf_8dd54bc25ec67c4c != 0 {
		return __obf_77240dc7776b60b4.__obf_7f98e493e52dfa5b(__obf_b416b6a4555be5c8)
	}

	for __obf_e4440ef7e7f689d0, __obf_782ca3783399c281 := range __obf_b416b6a4555be5c8 {
		if __obf_8882f6cf6e378ded := __obf_77240dc7776b60b4.EncodeString(__obf_e4440ef7e7f689d0); __obf_8882f6cf6e378ded != nil {
			return __obf_8882f6cf6e378ded
		}
		if __obf_8882f6cf6e378ded := __obf_77240dc7776b60b4.EncodeBool(__obf_782ca3783399c281); __obf_8882f6cf6e378ded != nil {
			return __obf_8882f6cf6e378ded
		}
	}

	return nil
}

func __obf_af298ee09c591531(__obf_77240dc7776b60b4 *Encoder, __obf_63bbcee86d44fdde reflect.Value) error {
	if __obf_63bbcee86d44fdde.IsNil() {
		return __obf_77240dc7776b60b4.EncodeNil()
	}

	if __obf_8882f6cf6e378ded := __obf_77240dc7776b60b4.EncodeMapLen(__obf_63bbcee86d44fdde.Len()); __obf_8882f6cf6e378ded != nil {
		return __obf_8882f6cf6e378ded
	}
	__obf_b416b6a4555be5c8 := __obf_63bbcee86d44fdde.Convert(__obf_76d48a5285b334b8).Interface().(map[string]string)
	if __obf_77240dc7776b60b4.__obf_3cf0882fa5a4cafb&__obf_8dd54bc25ec67c4c != 0 {
		return __obf_77240dc7776b60b4.__obf_03d6994566a31ecc(__obf_b416b6a4555be5c8)
	}

	for __obf_e4440ef7e7f689d0, __obf_782ca3783399c281 := range __obf_b416b6a4555be5c8 {
		if __obf_8882f6cf6e378ded := __obf_77240dc7776b60b4.EncodeString(__obf_e4440ef7e7f689d0); __obf_8882f6cf6e378ded != nil {
			return __obf_8882f6cf6e378ded
		}
		if __obf_8882f6cf6e378ded := __obf_77240dc7776b60b4.EncodeString(__obf_782ca3783399c281); __obf_8882f6cf6e378ded != nil {
			return __obf_8882f6cf6e378ded
		}
	}

	return nil
}

func __obf_026b0bc834cd7f40(__obf_77240dc7776b60b4 *Encoder, __obf_63bbcee86d44fdde reflect.Value) error {
	if __obf_63bbcee86d44fdde.IsNil() {
		return __obf_77240dc7776b60b4.EncodeNil()
	}
	__obf_b416b6a4555be5c8 := __obf_63bbcee86d44fdde.Convert(__obf_46f915c67d10ce94).Interface().(map[string]any)
	if __obf_77240dc7776b60b4.__obf_3cf0882fa5a4cafb&__obf_8dd54bc25ec67c4c != 0 {
		return __obf_77240dc7776b60b4.EncodeMapSorted(__obf_b416b6a4555be5c8)
	}
	return __obf_77240dc7776b60b4.EncodeMap(__obf_b416b6a4555be5c8)
}

func (__obf_77240dc7776b60b4 *Encoder) EncodeMap(__obf_b416b6a4555be5c8 map[string]any) error {
	if __obf_b416b6a4555be5c8 == nil {
		return __obf_77240dc7776b60b4.EncodeNil()
	}
	if __obf_8882f6cf6e378ded := __obf_77240dc7776b60b4.EncodeMapLen(len(__obf_b416b6a4555be5c8)); __obf_8882f6cf6e378ded != nil {
		return __obf_8882f6cf6e378ded
	}
	for __obf_e4440ef7e7f689d0, __obf_782ca3783399c281 := range __obf_b416b6a4555be5c8 {
		if __obf_8882f6cf6e378ded := __obf_77240dc7776b60b4.EncodeString(__obf_e4440ef7e7f689d0); __obf_8882f6cf6e378ded != nil {
			return __obf_8882f6cf6e378ded
		}
		if __obf_8882f6cf6e378ded := __obf_77240dc7776b60b4.Encode(__obf_782ca3783399c281); __obf_8882f6cf6e378ded != nil {
			return __obf_8882f6cf6e378ded
		}
	}
	return nil
}

func (__obf_77240dc7776b60b4 *Encoder) EncodeMapSorted(__obf_b416b6a4555be5c8 map[string]any) error {
	if __obf_b416b6a4555be5c8 == nil {
		return __obf_77240dc7776b60b4.EncodeNil()
	}
	if __obf_8882f6cf6e378ded := __obf_77240dc7776b60b4.EncodeMapLen(len(__obf_b416b6a4555be5c8)); __obf_8882f6cf6e378ded != nil {
		return __obf_8882f6cf6e378ded
	}
	__obf_a97796a39b27a4b0 := make([]string, 0, len(__obf_b416b6a4555be5c8))

	for __obf_659eb681649d5d0a := range __obf_b416b6a4555be5c8 {
		__obf_a97796a39b27a4b0 = append(__obf_a97796a39b27a4b0, __obf_659eb681649d5d0a)
	}

	sort.Strings(__obf_a97796a39b27a4b0)

	for _, __obf_659eb681649d5d0a := range __obf_a97796a39b27a4b0 {
		if __obf_8882f6cf6e378ded := __obf_77240dc7776b60b4.EncodeString(__obf_659eb681649d5d0a); __obf_8882f6cf6e378ded != nil {
			return __obf_8882f6cf6e378ded
		}
		if __obf_8882f6cf6e378ded := __obf_77240dc7776b60b4.Encode(__obf_b416b6a4555be5c8[__obf_659eb681649d5d0a]); __obf_8882f6cf6e378ded != nil {
			return __obf_8882f6cf6e378ded
		}
	}

	return nil
}

func (__obf_77240dc7776b60b4 *Encoder) __obf_7f98e493e52dfa5b(__obf_b416b6a4555be5c8 map[string]bool) error {
	__obf_a97796a39b27a4b0 := make([]string, 0, len(__obf_b416b6a4555be5c8))
	for __obf_659eb681649d5d0a := range __obf_b416b6a4555be5c8 {
		__obf_a97796a39b27a4b0 = append(__obf_a97796a39b27a4b0, __obf_659eb681649d5d0a)
	}
	sort.Strings(__obf_a97796a39b27a4b0)

	for _, __obf_659eb681649d5d0a := range __obf_a97796a39b27a4b0 {
		__obf_8882f6cf6e378ded := __obf_77240dc7776b60b4.EncodeString(__obf_659eb681649d5d0a)
		if __obf_8882f6cf6e378ded != nil {
			return __obf_8882f6cf6e378ded
		}
		if __obf_8882f6cf6e378ded = __obf_77240dc7776b60b4.EncodeBool(__obf_b416b6a4555be5c8[__obf_659eb681649d5d0a]); __obf_8882f6cf6e378ded != nil {
			return __obf_8882f6cf6e378ded
		}
	}

	return nil
}

func (__obf_77240dc7776b60b4 *Encoder) __obf_03d6994566a31ecc(__obf_b416b6a4555be5c8 map[string]string) error {
	__obf_a97796a39b27a4b0 := make([]string, 0, len(__obf_b416b6a4555be5c8))
	for __obf_659eb681649d5d0a := range __obf_b416b6a4555be5c8 {
		__obf_a97796a39b27a4b0 = append(__obf_a97796a39b27a4b0, __obf_659eb681649d5d0a)
	}
	sort.Strings(__obf_a97796a39b27a4b0)

	for _, __obf_659eb681649d5d0a := range __obf_a97796a39b27a4b0 {
		__obf_8882f6cf6e378ded := __obf_77240dc7776b60b4.EncodeString(__obf_659eb681649d5d0a)
		if __obf_8882f6cf6e378ded != nil {
			return __obf_8882f6cf6e378ded
		}
		if __obf_8882f6cf6e378ded = __obf_77240dc7776b60b4.EncodeString(__obf_b416b6a4555be5c8[__obf_659eb681649d5d0a]); __obf_8882f6cf6e378ded != nil {
			return __obf_8882f6cf6e378ded
		}
	}

	return nil
}

func (__obf_77240dc7776b60b4 *Encoder) EncodeMapLen(__obf_d62f82ed950927da int) error {
	if __obf_d62f82ed950927da < 16 {
		return __obf_77240dc7776b60b4.__obf_e8d9aa5d4f5eb974(FixedMapLow | byte(__obf_d62f82ed950927da))
	}
	if __obf_d62f82ed950927da <= math.MaxUint16 {
		return __obf_77240dc7776b60b4.__obf_9ecc220d5defc1d9(Map16, uint16(__obf_d62f82ed950927da))
	}
	return __obf_77240dc7776b60b4.__obf_70cbe6e3e07b6347(Map32, uint32(__obf_d62f82ed950927da))
}

func __obf_cbc3a9db722bf76c(__obf_77240dc7776b60b4 *Encoder, __obf_b8f9ff0ad69ee384 reflect.Value) error {
	__obf_ab0626d3c29fac63 := __obf_3a8fae1624b5849e.Fields(__obf_b8f9ff0ad69ee384.Type(), __obf_77240dc7776b60b4.__obf_6d0a0d8a79287f26)
	if __obf_77240dc7776b60b4.__obf_3cf0882fa5a4cafb&__obf_7680454fe6a3ccd7 != 0 || __obf_ab0626d3c29fac63.AsArray {
		return __obf_630c97e8b5c4e4ec(__obf_77240dc7776b60b4, __obf_b8f9ff0ad69ee384, __obf_ab0626d3c29fac63.List)
	}
	__obf_ce19c38e443cefa8 := __obf_ab0626d3c29fac63.OmitEmpty(__obf_77240dc7776b60b4, __obf_b8f9ff0ad69ee384)

	if __obf_8882f6cf6e378ded := __obf_77240dc7776b60b4.EncodeMapLen(len(__obf_ce19c38e443cefa8)); __obf_8882f6cf6e378ded != nil {
		return __obf_8882f6cf6e378ded
	}

	for _, __obf_4add40b0f5dc6b95 := range __obf_ce19c38e443cefa8 {
		if __obf_8882f6cf6e378ded := __obf_77240dc7776b60b4.EncodeString(__obf_4add40b0f5dc6b95.__obf_754359097736e0d5); __obf_8882f6cf6e378ded != nil {
			return __obf_8882f6cf6e378ded
		}
		if __obf_8882f6cf6e378ded := __obf_4add40b0f5dc6b95.EncodeValue(__obf_77240dc7776b60b4, __obf_b8f9ff0ad69ee384); __obf_8882f6cf6e378ded != nil {
			return __obf_8882f6cf6e378ded
		}
	}

	return nil
}

func __obf_630c97e8b5c4e4ec(__obf_77240dc7776b60b4 *Encoder, __obf_b8f9ff0ad69ee384 reflect.Value, __obf_ce19c38e443cefa8 []*__obf_a5cb7d1c497d43ce) error {
	if __obf_8882f6cf6e378ded := __obf_77240dc7776b60b4.EncodeArrayLen(len(__obf_ce19c38e443cefa8)); __obf_8882f6cf6e378ded != nil {
		return __obf_8882f6cf6e378ded
	}
	for _, __obf_4add40b0f5dc6b95 := range __obf_ce19c38e443cefa8 {
		if __obf_8882f6cf6e378ded := __obf_4add40b0f5dc6b95.EncodeValue(__obf_77240dc7776b60b4, __obf_b8f9ff0ad69ee384); __obf_8882f6cf6e378ded != nil {
			return __obf_8882f6cf6e378ded
		}
	}
	return nil
}
