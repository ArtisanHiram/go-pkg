package __obf_fd770cb9675903b0

import (
	"math"
	"reflect"
	"sort"
)

func __obf_50d653704a4e67b8(__obf_e9038cda3b5cf711 *Encoder, __obf_f328a048f76b7256 reflect.Value) error {
	if __obf_f328a048f76b7256.IsNil() {
		return __obf_e9038cda3b5cf711.EncodeNil()
	}

	if __obf_45342a3a754d12f5 := __obf_e9038cda3b5cf711.EncodeMapLen(__obf_f328a048f76b7256.Len()); __obf_45342a3a754d12f5 != nil {
		return __obf_45342a3a754d12f5
	}
	__obf_952e96cf1beb18c1 := __obf_f328a048f76b7256.MapRange()
	for __obf_952e96cf1beb18c1.Next() {
		if __obf_45342a3a754d12f5 := __obf_e9038cda3b5cf711.EncodeValue(__obf_952e96cf1beb18c1.Key()); __obf_45342a3a754d12f5 != nil {
			return __obf_45342a3a754d12f5
		}
		if __obf_45342a3a754d12f5 := __obf_e9038cda3b5cf711.EncodeValue(__obf_952e96cf1beb18c1.Value()); __obf_45342a3a754d12f5 != nil {
			return __obf_45342a3a754d12f5
		}
	}

	return nil
}

func __obf_e92628670d4494a8(__obf_e9038cda3b5cf711 *Encoder, __obf_f328a048f76b7256 reflect.Value) error {
	if __obf_f328a048f76b7256.IsNil() {
		return __obf_e9038cda3b5cf711.EncodeNil()
	}

	if __obf_45342a3a754d12f5 := __obf_e9038cda3b5cf711.EncodeMapLen(__obf_f328a048f76b7256.Len()); __obf_45342a3a754d12f5 != nil {
		return __obf_45342a3a754d12f5
	}
	__obf_777489ec8e6b2044 := __obf_f328a048f76b7256.Convert(__obf_66e0ed467aaa55ff).Interface().(map[string]bool)
	if __obf_e9038cda3b5cf711.__obf_7185cb62de7af792&__obf_4010aa531727deba != 0 {
		return __obf_e9038cda3b5cf711.__obf_aa08922f873407c2(__obf_777489ec8e6b2044)
	}

	for __obf_a3ae43155694e198, __obf_2cca23d5c936f4ab := range __obf_777489ec8e6b2044 {
		if __obf_45342a3a754d12f5 := __obf_e9038cda3b5cf711.EncodeString(__obf_a3ae43155694e198); __obf_45342a3a754d12f5 != nil {
			return __obf_45342a3a754d12f5
		}
		if __obf_45342a3a754d12f5 := __obf_e9038cda3b5cf711.EncodeBool(__obf_2cca23d5c936f4ab); __obf_45342a3a754d12f5 != nil {
			return __obf_45342a3a754d12f5
		}
	}

	return nil
}

func __obf_617b6f6590f5807c(__obf_e9038cda3b5cf711 *Encoder, __obf_f328a048f76b7256 reflect.Value) error {
	if __obf_f328a048f76b7256.IsNil() {
		return __obf_e9038cda3b5cf711.EncodeNil()
	}

	if __obf_45342a3a754d12f5 := __obf_e9038cda3b5cf711.EncodeMapLen(__obf_f328a048f76b7256.Len()); __obf_45342a3a754d12f5 != nil {
		return __obf_45342a3a754d12f5
	}
	__obf_777489ec8e6b2044 := __obf_f328a048f76b7256.Convert(__obf_a25da01fd293fb59).Interface().(map[string]string)
	if __obf_e9038cda3b5cf711.__obf_7185cb62de7af792&__obf_4010aa531727deba != 0 {
		return __obf_e9038cda3b5cf711.__obf_0ae7478903c7794a(__obf_777489ec8e6b2044)
	}

	for __obf_a3ae43155694e198, __obf_2cca23d5c936f4ab := range __obf_777489ec8e6b2044 {
		if __obf_45342a3a754d12f5 := __obf_e9038cda3b5cf711.EncodeString(__obf_a3ae43155694e198); __obf_45342a3a754d12f5 != nil {
			return __obf_45342a3a754d12f5
		}
		if __obf_45342a3a754d12f5 := __obf_e9038cda3b5cf711.EncodeString(__obf_2cca23d5c936f4ab); __obf_45342a3a754d12f5 != nil {
			return __obf_45342a3a754d12f5
		}
	}

	return nil
}

func __obf_d3a88247514432a8(__obf_e9038cda3b5cf711 *Encoder, __obf_f328a048f76b7256 reflect.Value) error {
	if __obf_f328a048f76b7256.IsNil() {
		return __obf_e9038cda3b5cf711.EncodeNil()
	}
	__obf_777489ec8e6b2044 := __obf_f328a048f76b7256.Convert(__obf_3c15f58273d2641c).Interface().(map[string]any)
	if __obf_e9038cda3b5cf711.__obf_7185cb62de7af792&__obf_4010aa531727deba != 0 {
		return __obf_e9038cda3b5cf711.EncodeMapSorted(__obf_777489ec8e6b2044)
	}
	return __obf_e9038cda3b5cf711.EncodeMap(__obf_777489ec8e6b2044)
}

func (__obf_e9038cda3b5cf711 *Encoder) EncodeMap(__obf_777489ec8e6b2044 map[string]any) error {
	if __obf_777489ec8e6b2044 == nil {
		return __obf_e9038cda3b5cf711.EncodeNil()
	}
	if __obf_45342a3a754d12f5 := __obf_e9038cda3b5cf711.EncodeMapLen(len(__obf_777489ec8e6b2044)); __obf_45342a3a754d12f5 != nil {
		return __obf_45342a3a754d12f5
	}
	for __obf_a3ae43155694e198, __obf_2cca23d5c936f4ab := range __obf_777489ec8e6b2044 {
		if __obf_45342a3a754d12f5 := __obf_e9038cda3b5cf711.EncodeString(__obf_a3ae43155694e198); __obf_45342a3a754d12f5 != nil {
			return __obf_45342a3a754d12f5
		}
		if __obf_45342a3a754d12f5 := __obf_e9038cda3b5cf711.Encode(__obf_2cca23d5c936f4ab); __obf_45342a3a754d12f5 != nil {
			return __obf_45342a3a754d12f5
		}
	}
	return nil
}

func (__obf_e9038cda3b5cf711 *Encoder) EncodeMapSorted(__obf_777489ec8e6b2044 map[string]any) error {
	if __obf_777489ec8e6b2044 == nil {
		return __obf_e9038cda3b5cf711.EncodeNil()
	}
	if __obf_45342a3a754d12f5 := __obf_e9038cda3b5cf711.EncodeMapLen(len(__obf_777489ec8e6b2044)); __obf_45342a3a754d12f5 != nil {
		return __obf_45342a3a754d12f5
	}
	__obf_b8ee03f219c7dd0a := make([]string, 0, len(__obf_777489ec8e6b2044))

	for __obf_592b715dc5180198 := range __obf_777489ec8e6b2044 {
		__obf_b8ee03f219c7dd0a = append(__obf_b8ee03f219c7dd0a, __obf_592b715dc5180198)
	}

	sort.Strings(__obf_b8ee03f219c7dd0a)

	for _, __obf_592b715dc5180198 := range __obf_b8ee03f219c7dd0a {
		if __obf_45342a3a754d12f5 := __obf_e9038cda3b5cf711.EncodeString(__obf_592b715dc5180198); __obf_45342a3a754d12f5 != nil {
			return __obf_45342a3a754d12f5
		}
		if __obf_45342a3a754d12f5 := __obf_e9038cda3b5cf711.Encode(__obf_777489ec8e6b2044[__obf_592b715dc5180198]); __obf_45342a3a754d12f5 != nil {
			return __obf_45342a3a754d12f5
		}
	}

	return nil
}

func (__obf_e9038cda3b5cf711 *Encoder) __obf_aa08922f873407c2(__obf_777489ec8e6b2044 map[string]bool) error {
	__obf_b8ee03f219c7dd0a := make([]string, 0, len(__obf_777489ec8e6b2044))
	for __obf_592b715dc5180198 := range __obf_777489ec8e6b2044 {
		__obf_b8ee03f219c7dd0a = append(__obf_b8ee03f219c7dd0a, __obf_592b715dc5180198)
	}
	sort.Strings(__obf_b8ee03f219c7dd0a)

	for _, __obf_592b715dc5180198 := range __obf_b8ee03f219c7dd0a {
		__obf_45342a3a754d12f5 := __obf_e9038cda3b5cf711.EncodeString(__obf_592b715dc5180198)
		if __obf_45342a3a754d12f5 != nil {
			return __obf_45342a3a754d12f5
		}
		if __obf_45342a3a754d12f5 = __obf_e9038cda3b5cf711.EncodeBool(__obf_777489ec8e6b2044[__obf_592b715dc5180198]); __obf_45342a3a754d12f5 != nil {
			return __obf_45342a3a754d12f5
		}
	}

	return nil
}

func (__obf_e9038cda3b5cf711 *Encoder) __obf_0ae7478903c7794a(__obf_777489ec8e6b2044 map[string]string) error {
	__obf_b8ee03f219c7dd0a := make([]string, 0, len(__obf_777489ec8e6b2044))
	for __obf_592b715dc5180198 := range __obf_777489ec8e6b2044 {
		__obf_b8ee03f219c7dd0a = append(__obf_b8ee03f219c7dd0a, __obf_592b715dc5180198)
	}
	sort.Strings(__obf_b8ee03f219c7dd0a)

	for _, __obf_592b715dc5180198 := range __obf_b8ee03f219c7dd0a {
		__obf_45342a3a754d12f5 := __obf_e9038cda3b5cf711.EncodeString(__obf_592b715dc5180198)
		if __obf_45342a3a754d12f5 != nil {
			return __obf_45342a3a754d12f5
		}
		if __obf_45342a3a754d12f5 = __obf_e9038cda3b5cf711.EncodeString(__obf_777489ec8e6b2044[__obf_592b715dc5180198]); __obf_45342a3a754d12f5 != nil {
			return __obf_45342a3a754d12f5
		}
	}

	return nil
}

func (__obf_e9038cda3b5cf711 *Encoder) EncodeMapLen(__obf_bddea2836c583aa2 int) error {
	if __obf_bddea2836c583aa2 < 16 {
		return __obf_e9038cda3b5cf711.__obf_a72c49227d01f9dd(FixedMapLow | byte(__obf_bddea2836c583aa2))
	}
	if __obf_bddea2836c583aa2 <= math.MaxUint16 {
		return __obf_e9038cda3b5cf711.__obf_a98b5ca0c9f03534(Map16, uint16(__obf_bddea2836c583aa2))
	}
	return __obf_e9038cda3b5cf711.__obf_e0e3a3f6b942651e(Map32, uint32(__obf_bddea2836c583aa2))
}

func __obf_c52bd37d7024764c(__obf_e9038cda3b5cf711 *Encoder, __obf_475491b3031fb8b4 reflect.Value) error {
	__obf_7d081654985e1a56 := __obf_a93e575d0cc03783.Fields(__obf_475491b3031fb8b4.Type(), __obf_e9038cda3b5cf711.__obf_6621ba3773b8696a)
	if __obf_e9038cda3b5cf711.__obf_7185cb62de7af792&__obf_b1ce96c64c5697ef != 0 || __obf_7d081654985e1a56.AsArray {
		return __obf_69cc590da446ae20(__obf_e9038cda3b5cf711, __obf_475491b3031fb8b4, __obf_7d081654985e1a56.List)
	}
	__obf_f92dfca778102077 := __obf_7d081654985e1a56.OmitEmpty(__obf_e9038cda3b5cf711, __obf_475491b3031fb8b4)

	if __obf_45342a3a754d12f5 := __obf_e9038cda3b5cf711.EncodeMapLen(len(__obf_f92dfca778102077)); __obf_45342a3a754d12f5 != nil {
		return __obf_45342a3a754d12f5
	}

	for _, __obf_e25b9e550ea05af3 := range __obf_f92dfca778102077 {
		if __obf_45342a3a754d12f5 := __obf_e9038cda3b5cf711.EncodeString(__obf_e25b9e550ea05af3.__obf_2a5a18f2bc2906b2); __obf_45342a3a754d12f5 != nil {
			return __obf_45342a3a754d12f5
		}
		if __obf_45342a3a754d12f5 := __obf_e25b9e550ea05af3.EncodeValue(__obf_e9038cda3b5cf711, __obf_475491b3031fb8b4); __obf_45342a3a754d12f5 != nil {
			return __obf_45342a3a754d12f5
		}
	}

	return nil
}

func __obf_69cc590da446ae20(__obf_e9038cda3b5cf711 *Encoder, __obf_475491b3031fb8b4 reflect.Value, __obf_f92dfca778102077 []*__obf_9b29490b522018e0) error {
	if __obf_45342a3a754d12f5 := __obf_e9038cda3b5cf711.EncodeArrayLen(len(__obf_f92dfca778102077)); __obf_45342a3a754d12f5 != nil {
		return __obf_45342a3a754d12f5
	}
	for _, __obf_e25b9e550ea05af3 := range __obf_f92dfca778102077 {
		if __obf_45342a3a754d12f5 := __obf_e25b9e550ea05af3.EncodeValue(__obf_e9038cda3b5cf711, __obf_475491b3031fb8b4); __obf_45342a3a754d12f5 != nil {
			return __obf_45342a3a754d12f5
		}
	}
	return nil
}
