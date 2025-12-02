package __obf_a935eb7f548271a4

import (
	"math"
	"reflect"
	"sort"
)

func __obf_a743877260781a87(__obf_ed37a34c347049f3 *Encoder, __obf_6d570581f4b60dbc reflect.Value) error {
	if __obf_6d570581f4b60dbc.IsNil() {
		return __obf_ed37a34c347049f3.EncodeNil()
	}

	if __obf_4d327e1cd40c2e21 := __obf_ed37a34c347049f3.EncodeMapLen(__obf_6d570581f4b60dbc.Len()); __obf_4d327e1cd40c2e21 != nil {
		return __obf_4d327e1cd40c2e21
	}
	__obf_f5332b3bc41330b8 := __obf_6d570581f4b60dbc.MapRange()
	for __obf_f5332b3bc41330b8.Next() {
		if __obf_4d327e1cd40c2e21 := __obf_ed37a34c347049f3.EncodeValue(__obf_f5332b3bc41330b8.Key()); __obf_4d327e1cd40c2e21 != nil {
			return __obf_4d327e1cd40c2e21
		}
		if __obf_4d327e1cd40c2e21 := __obf_ed37a34c347049f3.EncodeValue(__obf_f5332b3bc41330b8.Value()); __obf_4d327e1cd40c2e21 != nil {
			return __obf_4d327e1cd40c2e21
		}
	}

	return nil
}

func __obf_19525146631c9c00(__obf_ed37a34c347049f3 *Encoder, __obf_6d570581f4b60dbc reflect.Value) error {
	if __obf_6d570581f4b60dbc.IsNil() {
		return __obf_ed37a34c347049f3.EncodeNil()
	}

	if __obf_4d327e1cd40c2e21 := __obf_ed37a34c347049f3.EncodeMapLen(__obf_6d570581f4b60dbc.Len()); __obf_4d327e1cd40c2e21 != nil {
		return __obf_4d327e1cd40c2e21
	}
	__obf_26d12ef0df36a324 := __obf_6d570581f4b60dbc.Convert(__obf_9436e2f0a394e1f9).Interface().(map[string]bool)
	if __obf_ed37a34c347049f3.__obf_d617f3769ce16c47&__obf_51d2317856ab3853 != 0 {
		return __obf_ed37a34c347049f3.__obf_5a48a47ce344519b(__obf_26d12ef0df36a324)
	}

	for __obf_3b7905b85fce1fb1, __obf_b35b927f0bf61896 := range __obf_26d12ef0df36a324 {
		if __obf_4d327e1cd40c2e21 := __obf_ed37a34c347049f3.EncodeString(__obf_3b7905b85fce1fb1); __obf_4d327e1cd40c2e21 != nil {
			return __obf_4d327e1cd40c2e21
		}
		if __obf_4d327e1cd40c2e21 := __obf_ed37a34c347049f3.EncodeBool(__obf_b35b927f0bf61896); __obf_4d327e1cd40c2e21 != nil {
			return __obf_4d327e1cd40c2e21
		}
	}

	return nil
}

func __obf_e9d97c6d9e0e5b5f(__obf_ed37a34c347049f3 *Encoder, __obf_6d570581f4b60dbc reflect.Value) error {
	if __obf_6d570581f4b60dbc.IsNil() {
		return __obf_ed37a34c347049f3.EncodeNil()
	}

	if __obf_4d327e1cd40c2e21 := __obf_ed37a34c347049f3.EncodeMapLen(__obf_6d570581f4b60dbc.Len()); __obf_4d327e1cd40c2e21 != nil {
		return __obf_4d327e1cd40c2e21
	}
	__obf_26d12ef0df36a324 := __obf_6d570581f4b60dbc.Convert(__obf_eead1ffc8ffe8371).Interface().(map[string]string)
	if __obf_ed37a34c347049f3.__obf_d617f3769ce16c47&__obf_51d2317856ab3853 != 0 {
		return __obf_ed37a34c347049f3.__obf_efcc6860b9252a3a(__obf_26d12ef0df36a324)
	}

	for __obf_3b7905b85fce1fb1, __obf_b35b927f0bf61896 := range __obf_26d12ef0df36a324 {
		if __obf_4d327e1cd40c2e21 := __obf_ed37a34c347049f3.EncodeString(__obf_3b7905b85fce1fb1); __obf_4d327e1cd40c2e21 != nil {
			return __obf_4d327e1cd40c2e21
		}
		if __obf_4d327e1cd40c2e21 := __obf_ed37a34c347049f3.EncodeString(__obf_b35b927f0bf61896); __obf_4d327e1cd40c2e21 != nil {
			return __obf_4d327e1cd40c2e21
		}
	}

	return nil
}

func __obf_25ee786fca8b560b(__obf_ed37a34c347049f3 *Encoder, __obf_6d570581f4b60dbc reflect.Value) error {
	if __obf_6d570581f4b60dbc.IsNil() {
		return __obf_ed37a34c347049f3.EncodeNil()
	}
	__obf_26d12ef0df36a324 := __obf_6d570581f4b60dbc.Convert(__obf_6782064e44047a94).Interface().(map[string]any)
	if __obf_ed37a34c347049f3.__obf_d617f3769ce16c47&__obf_51d2317856ab3853 != 0 {
		return __obf_ed37a34c347049f3.EncodeMapSorted(__obf_26d12ef0df36a324)
	}
	return __obf_ed37a34c347049f3.EncodeMap(__obf_26d12ef0df36a324)
}

func (__obf_ed37a34c347049f3 *Encoder) EncodeMap(__obf_26d12ef0df36a324 map[string]any) error {
	if __obf_26d12ef0df36a324 == nil {
		return __obf_ed37a34c347049f3.EncodeNil()
	}
	if __obf_4d327e1cd40c2e21 := __obf_ed37a34c347049f3.EncodeMapLen(len(__obf_26d12ef0df36a324)); __obf_4d327e1cd40c2e21 != nil {
		return __obf_4d327e1cd40c2e21
	}
	for __obf_3b7905b85fce1fb1, __obf_b35b927f0bf61896 := range __obf_26d12ef0df36a324 {
		if __obf_4d327e1cd40c2e21 := __obf_ed37a34c347049f3.EncodeString(__obf_3b7905b85fce1fb1); __obf_4d327e1cd40c2e21 != nil {
			return __obf_4d327e1cd40c2e21
		}
		if __obf_4d327e1cd40c2e21 := __obf_ed37a34c347049f3.Encode(__obf_b35b927f0bf61896); __obf_4d327e1cd40c2e21 != nil {
			return __obf_4d327e1cd40c2e21
		}
	}
	return nil
}

func (__obf_ed37a34c347049f3 *Encoder) EncodeMapSorted(__obf_26d12ef0df36a324 map[string]any) error {
	if __obf_26d12ef0df36a324 == nil {
		return __obf_ed37a34c347049f3.EncodeNil()
	}
	if __obf_4d327e1cd40c2e21 := __obf_ed37a34c347049f3.EncodeMapLen(len(__obf_26d12ef0df36a324)); __obf_4d327e1cd40c2e21 != nil {
		return __obf_4d327e1cd40c2e21
	}
	__obf_b65b0c2f8480a025 := make([]string, 0, len(__obf_26d12ef0df36a324))

	for __obf_27cd90e99138f7bb := range __obf_26d12ef0df36a324 {
		__obf_b65b0c2f8480a025 = append(__obf_b65b0c2f8480a025, __obf_27cd90e99138f7bb)
	}

	sort.Strings(__obf_b65b0c2f8480a025)

	for _, __obf_27cd90e99138f7bb := range __obf_b65b0c2f8480a025 {
		if __obf_4d327e1cd40c2e21 := __obf_ed37a34c347049f3.EncodeString(__obf_27cd90e99138f7bb); __obf_4d327e1cd40c2e21 != nil {
			return __obf_4d327e1cd40c2e21
		}
		if __obf_4d327e1cd40c2e21 := __obf_ed37a34c347049f3.Encode(__obf_26d12ef0df36a324[__obf_27cd90e99138f7bb]); __obf_4d327e1cd40c2e21 != nil {
			return __obf_4d327e1cd40c2e21
		}
	}

	return nil
}

func (__obf_ed37a34c347049f3 *Encoder) __obf_5a48a47ce344519b(__obf_26d12ef0df36a324 map[string]bool) error {
	__obf_b65b0c2f8480a025 := make([]string, 0, len(__obf_26d12ef0df36a324))
	for __obf_27cd90e99138f7bb := range __obf_26d12ef0df36a324 {
		__obf_b65b0c2f8480a025 = append(__obf_b65b0c2f8480a025, __obf_27cd90e99138f7bb)
	}
	sort.Strings(__obf_b65b0c2f8480a025)

	for _, __obf_27cd90e99138f7bb := range __obf_b65b0c2f8480a025 {
		__obf_4d327e1cd40c2e21 := __obf_ed37a34c347049f3.EncodeString(__obf_27cd90e99138f7bb)
		if __obf_4d327e1cd40c2e21 != nil {
			return __obf_4d327e1cd40c2e21
		}
		if __obf_4d327e1cd40c2e21 = __obf_ed37a34c347049f3.EncodeBool(__obf_26d12ef0df36a324[__obf_27cd90e99138f7bb]); __obf_4d327e1cd40c2e21 != nil {
			return __obf_4d327e1cd40c2e21
		}
	}

	return nil
}

func (__obf_ed37a34c347049f3 *Encoder) __obf_efcc6860b9252a3a(__obf_26d12ef0df36a324 map[string]string) error {
	__obf_b65b0c2f8480a025 := make([]string, 0, len(__obf_26d12ef0df36a324))
	for __obf_27cd90e99138f7bb := range __obf_26d12ef0df36a324 {
		__obf_b65b0c2f8480a025 = append(__obf_b65b0c2f8480a025, __obf_27cd90e99138f7bb)
	}
	sort.Strings(__obf_b65b0c2f8480a025)

	for _, __obf_27cd90e99138f7bb := range __obf_b65b0c2f8480a025 {
		__obf_4d327e1cd40c2e21 := __obf_ed37a34c347049f3.EncodeString(__obf_27cd90e99138f7bb)
		if __obf_4d327e1cd40c2e21 != nil {
			return __obf_4d327e1cd40c2e21
		}
		if __obf_4d327e1cd40c2e21 = __obf_ed37a34c347049f3.EncodeString(__obf_26d12ef0df36a324[__obf_27cd90e99138f7bb]); __obf_4d327e1cd40c2e21 != nil {
			return __obf_4d327e1cd40c2e21
		}
	}

	return nil
}

func (__obf_ed37a34c347049f3 *Encoder) EncodeMapLen(__obf_3c213c92f0597e4d int) error {
	if __obf_3c213c92f0597e4d < 16 {
		return __obf_ed37a34c347049f3.__obf_2252dbfd2473b643(FixedMapLow | byte(__obf_3c213c92f0597e4d))
	}
	if __obf_3c213c92f0597e4d <= math.MaxUint16 {
		return __obf_ed37a34c347049f3.__obf_dd92b251146e86fb(Map16, uint16(__obf_3c213c92f0597e4d))
	}
	return __obf_ed37a34c347049f3.__obf_21945a55177e9fbd(Map32, uint32(__obf_3c213c92f0597e4d))
}

func __obf_1d06e69b22f89920(__obf_ed37a34c347049f3 *Encoder, __obf_5e97d46983ad5951 reflect.Value) error {
	__obf_bf5291b060bc4afb := __obf_b7d9de4837c900e0.Fields(__obf_5e97d46983ad5951.Type(), __obf_ed37a34c347049f3.__obf_90e41d275c699c20)
	if __obf_ed37a34c347049f3.__obf_d617f3769ce16c47&__obf_d6a65589ccd01600 != 0 || __obf_bf5291b060bc4afb.AsArray {
		return __obf_8d399e52493f0915(__obf_ed37a34c347049f3, __obf_5e97d46983ad5951, __obf_bf5291b060bc4afb.List)
	}
	__obf_02acefac40101465 := __obf_bf5291b060bc4afb.OmitEmpty(__obf_ed37a34c347049f3, __obf_5e97d46983ad5951)

	if __obf_4d327e1cd40c2e21 := __obf_ed37a34c347049f3.EncodeMapLen(len(__obf_02acefac40101465)); __obf_4d327e1cd40c2e21 != nil {
		return __obf_4d327e1cd40c2e21
	}

	for _, __obf_f6dd34b5068d19f3 := range __obf_02acefac40101465 {
		if __obf_4d327e1cd40c2e21 := __obf_ed37a34c347049f3.EncodeString(__obf_f6dd34b5068d19f3.__obf_e80b5f2d9b7252c4); __obf_4d327e1cd40c2e21 != nil {
			return __obf_4d327e1cd40c2e21
		}
		if __obf_4d327e1cd40c2e21 := __obf_f6dd34b5068d19f3.EncodeValue(__obf_ed37a34c347049f3, __obf_5e97d46983ad5951); __obf_4d327e1cd40c2e21 != nil {
			return __obf_4d327e1cd40c2e21
		}
	}

	return nil
}

func __obf_8d399e52493f0915(__obf_ed37a34c347049f3 *Encoder, __obf_5e97d46983ad5951 reflect.Value, __obf_02acefac40101465 []*__obf_3b4b6de10b0b5ac0) error {
	if __obf_4d327e1cd40c2e21 := __obf_ed37a34c347049f3.EncodeArrayLen(len(__obf_02acefac40101465)); __obf_4d327e1cd40c2e21 != nil {
		return __obf_4d327e1cd40c2e21
	}
	for _, __obf_f6dd34b5068d19f3 := range __obf_02acefac40101465 {
		if __obf_4d327e1cd40c2e21 := __obf_f6dd34b5068d19f3.EncodeValue(__obf_ed37a34c347049f3, __obf_5e97d46983ad5951); __obf_4d327e1cd40c2e21 != nil {
			return __obf_4d327e1cd40c2e21
		}
	}
	return nil
}
