package __obf_de86cdc8ae98b45b

import (
	"math"
	"reflect"
	"sort"
)

func __obf_90655a0691bba685(__obf_7bae0b613da60dd3 *Encoder, __obf_17732590722140e7 reflect.Value) error {
	if __obf_17732590722140e7.IsNil() {
		return __obf_7bae0b613da60dd3.EncodeNil()
	}

	if __obf_0feb3528b7b709ec := __obf_7bae0b613da60dd3.EncodeMapLen(__obf_17732590722140e7.Len()); __obf_0feb3528b7b709ec != nil {
		return __obf_0feb3528b7b709ec
	}
	__obf_6faef04689465bf5 := __obf_17732590722140e7.MapRange()
	for __obf_6faef04689465bf5.Next() {
		if __obf_0feb3528b7b709ec := __obf_7bae0b613da60dd3.EncodeValue(__obf_6faef04689465bf5.Key()); __obf_0feb3528b7b709ec != nil {
			return __obf_0feb3528b7b709ec
		}
		if __obf_0feb3528b7b709ec := __obf_7bae0b613da60dd3.EncodeValue(__obf_6faef04689465bf5.Value()); __obf_0feb3528b7b709ec != nil {
			return __obf_0feb3528b7b709ec
		}
	}

	return nil
}

func __obf_59e60c775a39bdfb(__obf_7bae0b613da60dd3 *Encoder, __obf_17732590722140e7 reflect.Value) error {
	if __obf_17732590722140e7.IsNil() {
		return __obf_7bae0b613da60dd3.EncodeNil()
	}

	if __obf_0feb3528b7b709ec := __obf_7bae0b613da60dd3.EncodeMapLen(__obf_17732590722140e7.Len()); __obf_0feb3528b7b709ec != nil {
		return __obf_0feb3528b7b709ec
	}
	__obf_2a26e7104b4c4373 := __obf_17732590722140e7.Convert(__obf_9ea30858b9b6e940).Interface().(map[string]bool)
	if __obf_7bae0b613da60dd3.__obf_a15a642720e39c3e&__obf_3da2e5262e6dbb1b != 0 {
		return __obf_7bae0b613da60dd3.__obf_70140837eb28ac8c(__obf_2a26e7104b4c4373)
	}

	for __obf_420732d362059ffa, __obf_a6cb485687dffc12 := range __obf_2a26e7104b4c4373 {
		if __obf_0feb3528b7b709ec := __obf_7bae0b613da60dd3.EncodeString(__obf_420732d362059ffa); __obf_0feb3528b7b709ec != nil {
			return __obf_0feb3528b7b709ec
		}
		if __obf_0feb3528b7b709ec := __obf_7bae0b613da60dd3.EncodeBool(__obf_a6cb485687dffc12); __obf_0feb3528b7b709ec != nil {
			return __obf_0feb3528b7b709ec
		}
	}

	return nil
}

func __obf_9805b50aa1981a7e(__obf_7bae0b613da60dd3 *Encoder, __obf_17732590722140e7 reflect.Value) error {
	if __obf_17732590722140e7.IsNil() {
		return __obf_7bae0b613da60dd3.EncodeNil()
	}

	if __obf_0feb3528b7b709ec := __obf_7bae0b613da60dd3.EncodeMapLen(__obf_17732590722140e7.Len()); __obf_0feb3528b7b709ec != nil {
		return __obf_0feb3528b7b709ec
	}
	__obf_2a26e7104b4c4373 := __obf_17732590722140e7.Convert(__obf_34d6b0138b97b155).Interface().(map[string]string)
	if __obf_7bae0b613da60dd3.__obf_a15a642720e39c3e&__obf_3da2e5262e6dbb1b != 0 {
		return __obf_7bae0b613da60dd3.__obf_ea8d119701730aa3(__obf_2a26e7104b4c4373)
	}

	for __obf_420732d362059ffa, __obf_a6cb485687dffc12 := range __obf_2a26e7104b4c4373 {
		if __obf_0feb3528b7b709ec := __obf_7bae0b613da60dd3.EncodeString(__obf_420732d362059ffa); __obf_0feb3528b7b709ec != nil {
			return __obf_0feb3528b7b709ec
		}
		if __obf_0feb3528b7b709ec := __obf_7bae0b613da60dd3.EncodeString(__obf_a6cb485687dffc12); __obf_0feb3528b7b709ec != nil {
			return __obf_0feb3528b7b709ec
		}
	}

	return nil
}

func __obf_62ccef293845516e(__obf_7bae0b613da60dd3 *Encoder, __obf_17732590722140e7 reflect.Value) error {
	if __obf_17732590722140e7.IsNil() {
		return __obf_7bae0b613da60dd3.EncodeNil()
	}
	__obf_2a26e7104b4c4373 := __obf_17732590722140e7.Convert(__obf_e8e62bb83070a108).Interface().(map[string]any)
	if __obf_7bae0b613da60dd3.__obf_a15a642720e39c3e&__obf_3da2e5262e6dbb1b != 0 {
		return __obf_7bae0b613da60dd3.EncodeMapSorted(__obf_2a26e7104b4c4373)
	}
	return __obf_7bae0b613da60dd3.EncodeMap(__obf_2a26e7104b4c4373)
}

func (__obf_7bae0b613da60dd3 *Encoder) EncodeMap(__obf_2a26e7104b4c4373 map[string]any) error {
	if __obf_2a26e7104b4c4373 == nil {
		return __obf_7bae0b613da60dd3.EncodeNil()
	}
	if __obf_0feb3528b7b709ec := __obf_7bae0b613da60dd3.EncodeMapLen(len(__obf_2a26e7104b4c4373)); __obf_0feb3528b7b709ec != nil {
		return __obf_0feb3528b7b709ec
	}
	for __obf_420732d362059ffa, __obf_a6cb485687dffc12 := range __obf_2a26e7104b4c4373 {
		if __obf_0feb3528b7b709ec := __obf_7bae0b613da60dd3.EncodeString(__obf_420732d362059ffa); __obf_0feb3528b7b709ec != nil {
			return __obf_0feb3528b7b709ec
		}
		if __obf_0feb3528b7b709ec := __obf_7bae0b613da60dd3.Encode(__obf_a6cb485687dffc12); __obf_0feb3528b7b709ec != nil {
			return __obf_0feb3528b7b709ec
		}
	}
	return nil
}

func (__obf_7bae0b613da60dd3 *Encoder) EncodeMapSorted(__obf_2a26e7104b4c4373 map[string]any) error {
	if __obf_2a26e7104b4c4373 == nil {
		return __obf_7bae0b613da60dd3.EncodeNil()
	}
	if __obf_0feb3528b7b709ec := __obf_7bae0b613da60dd3.EncodeMapLen(len(__obf_2a26e7104b4c4373)); __obf_0feb3528b7b709ec != nil {
		return __obf_0feb3528b7b709ec
	}
	__obf_c629eadb7f08c1bb := make([]string, 0, len(__obf_2a26e7104b4c4373))

	for __obf_6cc7432073215be0 := range __obf_2a26e7104b4c4373 {
		__obf_c629eadb7f08c1bb = append(__obf_c629eadb7f08c1bb, __obf_6cc7432073215be0)
	}

	sort.Strings(__obf_c629eadb7f08c1bb)

	for _, __obf_6cc7432073215be0 := range __obf_c629eadb7f08c1bb {
		if __obf_0feb3528b7b709ec := __obf_7bae0b613da60dd3.EncodeString(__obf_6cc7432073215be0); __obf_0feb3528b7b709ec != nil {
			return __obf_0feb3528b7b709ec
		}
		if __obf_0feb3528b7b709ec := __obf_7bae0b613da60dd3.Encode(__obf_2a26e7104b4c4373[__obf_6cc7432073215be0]); __obf_0feb3528b7b709ec != nil {
			return __obf_0feb3528b7b709ec
		}
	}

	return nil
}

func (__obf_7bae0b613da60dd3 *Encoder) __obf_70140837eb28ac8c(__obf_2a26e7104b4c4373 map[string]bool) error {
	__obf_c629eadb7f08c1bb := make([]string, 0, len(__obf_2a26e7104b4c4373))
	for __obf_6cc7432073215be0 := range __obf_2a26e7104b4c4373 {
		__obf_c629eadb7f08c1bb = append(__obf_c629eadb7f08c1bb, __obf_6cc7432073215be0)
	}
	sort.Strings(__obf_c629eadb7f08c1bb)

	for _, __obf_6cc7432073215be0 := range __obf_c629eadb7f08c1bb {
		__obf_0feb3528b7b709ec := __obf_7bae0b613da60dd3.EncodeString(__obf_6cc7432073215be0)
		if __obf_0feb3528b7b709ec != nil {
			return __obf_0feb3528b7b709ec
		}
		if __obf_0feb3528b7b709ec = __obf_7bae0b613da60dd3.EncodeBool(__obf_2a26e7104b4c4373[__obf_6cc7432073215be0]); __obf_0feb3528b7b709ec != nil {
			return __obf_0feb3528b7b709ec
		}
	}

	return nil
}

func (__obf_7bae0b613da60dd3 *Encoder) __obf_ea8d119701730aa3(__obf_2a26e7104b4c4373 map[string]string) error {
	__obf_c629eadb7f08c1bb := make([]string, 0, len(__obf_2a26e7104b4c4373))
	for __obf_6cc7432073215be0 := range __obf_2a26e7104b4c4373 {
		__obf_c629eadb7f08c1bb = append(__obf_c629eadb7f08c1bb, __obf_6cc7432073215be0)
	}
	sort.Strings(__obf_c629eadb7f08c1bb)

	for _, __obf_6cc7432073215be0 := range __obf_c629eadb7f08c1bb {
		__obf_0feb3528b7b709ec := __obf_7bae0b613da60dd3.EncodeString(__obf_6cc7432073215be0)
		if __obf_0feb3528b7b709ec != nil {
			return __obf_0feb3528b7b709ec
		}
		if __obf_0feb3528b7b709ec = __obf_7bae0b613da60dd3.EncodeString(__obf_2a26e7104b4c4373[__obf_6cc7432073215be0]); __obf_0feb3528b7b709ec != nil {
			return __obf_0feb3528b7b709ec
		}
	}

	return nil
}

func (__obf_7bae0b613da60dd3 *Encoder) EncodeMapLen(__obf_6dbc1eb636e18f94 int) error {
	if __obf_6dbc1eb636e18f94 < 16 {
		return __obf_7bae0b613da60dd3.__obf_47b86f684323fc33(FixedMapLow | byte(__obf_6dbc1eb636e18f94))
	}
	if __obf_6dbc1eb636e18f94 <= math.MaxUint16 {
		return __obf_7bae0b613da60dd3.__obf_bd6e685495964ed5(Map16, uint16(__obf_6dbc1eb636e18f94))
	}
	return __obf_7bae0b613da60dd3.__obf_a95f338c34ae207a(Map32, uint32(__obf_6dbc1eb636e18f94))
}

func __obf_95de04fb2a8b0af3(__obf_7bae0b613da60dd3 *Encoder, __obf_84ee659cd19e5239 reflect.Value) error {
	__obf_fa8592fc07f7c881 := __obf_4642342e308ea07b.Fields(__obf_84ee659cd19e5239.Type(), __obf_7bae0b613da60dd3.__obf_603dff1ad8d49539)
	if __obf_7bae0b613da60dd3.__obf_a15a642720e39c3e&__obf_1724e39b1a63df3b != 0 || __obf_fa8592fc07f7c881.AsArray {
		return __obf_9cb93bef79a30d58(__obf_7bae0b613da60dd3, __obf_84ee659cd19e5239, __obf_fa8592fc07f7c881.List)
	}
	__obf_7553cd24ebfd7da5 := __obf_fa8592fc07f7c881.OmitEmpty(__obf_7bae0b613da60dd3, __obf_84ee659cd19e5239)

	if __obf_0feb3528b7b709ec := __obf_7bae0b613da60dd3.EncodeMapLen(len(__obf_7553cd24ebfd7da5)); __obf_0feb3528b7b709ec != nil {
		return __obf_0feb3528b7b709ec
	}

	for _, __obf_cf8245c3f9570af5 := range __obf_7553cd24ebfd7da5 {
		if __obf_0feb3528b7b709ec := __obf_7bae0b613da60dd3.EncodeString(__obf_cf8245c3f9570af5.__obf_d7fb5470f7183af5); __obf_0feb3528b7b709ec != nil {
			return __obf_0feb3528b7b709ec
		}
		if __obf_0feb3528b7b709ec := __obf_cf8245c3f9570af5.EncodeValue(__obf_7bae0b613da60dd3, __obf_84ee659cd19e5239); __obf_0feb3528b7b709ec != nil {
			return __obf_0feb3528b7b709ec
		}
	}

	return nil
}

func __obf_9cb93bef79a30d58(__obf_7bae0b613da60dd3 *Encoder, __obf_84ee659cd19e5239 reflect.Value, __obf_7553cd24ebfd7da5 []*__obf_4ee6afc37d3486aa) error {
	if __obf_0feb3528b7b709ec := __obf_7bae0b613da60dd3.EncodeArrayLen(len(__obf_7553cd24ebfd7da5)); __obf_0feb3528b7b709ec != nil {
		return __obf_0feb3528b7b709ec
	}
	for _, __obf_cf8245c3f9570af5 := range __obf_7553cd24ebfd7da5 {
		if __obf_0feb3528b7b709ec := __obf_cf8245c3f9570af5.EncodeValue(__obf_7bae0b613da60dd3, __obf_84ee659cd19e5239); __obf_0feb3528b7b709ec != nil {
			return __obf_0feb3528b7b709ec
		}
	}
	return nil
}
