package __obf_3edaa49e853afa16

import (
	"math"
	"reflect"
	"sort"
)

func __obf_c2aec17963f8c911(__obf_84d0d31a8288f191 *Encoder, __obf_85d270343a0dfe14 reflect.Value) error {
	if __obf_85d270343a0dfe14.IsNil() {
		return __obf_84d0d31a8288f191.EncodeNil()
	}

	if __obf_3aa78ad28f50ed46 := __obf_84d0d31a8288f191.EncodeMapLen(__obf_85d270343a0dfe14.Len()); __obf_3aa78ad28f50ed46 != nil {
		return __obf_3aa78ad28f50ed46
	}
	__obf_843dd74003fa7854 := __obf_85d270343a0dfe14.MapRange()
	for __obf_843dd74003fa7854.Next() {
		if __obf_3aa78ad28f50ed46 := __obf_84d0d31a8288f191.EncodeValue(__obf_843dd74003fa7854.Key()); __obf_3aa78ad28f50ed46 != nil {
			return __obf_3aa78ad28f50ed46
		}
		if __obf_3aa78ad28f50ed46 := __obf_84d0d31a8288f191.EncodeValue(__obf_843dd74003fa7854.Value()); __obf_3aa78ad28f50ed46 != nil {
			return __obf_3aa78ad28f50ed46
		}
	}

	return nil
}

func __obf_a421402c0bc3fb7f(__obf_84d0d31a8288f191 *Encoder, __obf_85d270343a0dfe14 reflect.Value) error {
	if __obf_85d270343a0dfe14.IsNil() {
		return __obf_84d0d31a8288f191.EncodeNil()
	}

	if __obf_3aa78ad28f50ed46 := __obf_84d0d31a8288f191.EncodeMapLen(__obf_85d270343a0dfe14.Len()); __obf_3aa78ad28f50ed46 != nil {
		return __obf_3aa78ad28f50ed46
	}
	__obf_c26f5adfb3152475 := __obf_85d270343a0dfe14.Convert(__obf_faa0fb9daafd869d).Interface().(map[string]bool)
	if __obf_84d0d31a8288f191.__obf_6aa0efd537415192&__obf_67f6be0f2d2d4aa4 != 0 {
		return __obf_84d0d31a8288f191.__obf_fd6039710e723b44(__obf_c26f5adfb3152475)
	}

	for __obf_8f7121221504784f, __obf_a03d608795b584a3 := range __obf_c26f5adfb3152475 {
		if __obf_3aa78ad28f50ed46 := __obf_84d0d31a8288f191.EncodeString(__obf_8f7121221504784f); __obf_3aa78ad28f50ed46 != nil {
			return __obf_3aa78ad28f50ed46
		}
		if __obf_3aa78ad28f50ed46 := __obf_84d0d31a8288f191.EncodeBool(__obf_a03d608795b584a3); __obf_3aa78ad28f50ed46 != nil {
			return __obf_3aa78ad28f50ed46
		}
	}

	return nil
}

func __obf_be42c64a29c7a474(__obf_84d0d31a8288f191 *Encoder, __obf_85d270343a0dfe14 reflect.Value) error {
	if __obf_85d270343a0dfe14.IsNil() {
		return __obf_84d0d31a8288f191.EncodeNil()
	}

	if __obf_3aa78ad28f50ed46 := __obf_84d0d31a8288f191.EncodeMapLen(__obf_85d270343a0dfe14.Len()); __obf_3aa78ad28f50ed46 != nil {
		return __obf_3aa78ad28f50ed46
	}
	__obf_c26f5adfb3152475 := __obf_85d270343a0dfe14.Convert(__obf_3e2ae0e33841c7b2).Interface().(map[string]string)
	if __obf_84d0d31a8288f191.__obf_6aa0efd537415192&__obf_67f6be0f2d2d4aa4 != 0 {
		return __obf_84d0d31a8288f191.__obf_39e93fec82dd6558(__obf_c26f5adfb3152475)
	}

	for __obf_8f7121221504784f, __obf_a03d608795b584a3 := range __obf_c26f5adfb3152475 {
		if __obf_3aa78ad28f50ed46 := __obf_84d0d31a8288f191.EncodeString(__obf_8f7121221504784f); __obf_3aa78ad28f50ed46 != nil {
			return __obf_3aa78ad28f50ed46
		}
		if __obf_3aa78ad28f50ed46 := __obf_84d0d31a8288f191.EncodeString(__obf_a03d608795b584a3); __obf_3aa78ad28f50ed46 != nil {
			return __obf_3aa78ad28f50ed46
		}
	}

	return nil
}

func __obf_f18ed5a38c2a6b1a(__obf_84d0d31a8288f191 *Encoder, __obf_85d270343a0dfe14 reflect.Value) error {
	if __obf_85d270343a0dfe14.IsNil() {
		return __obf_84d0d31a8288f191.EncodeNil()
	}
	__obf_c26f5adfb3152475 := __obf_85d270343a0dfe14.Convert(__obf_739f33e1b9d59004).Interface().(map[string]any)
	if __obf_84d0d31a8288f191.__obf_6aa0efd537415192&__obf_67f6be0f2d2d4aa4 != 0 {
		return __obf_84d0d31a8288f191.EncodeMapSorted(__obf_c26f5adfb3152475)
	}
	return __obf_84d0d31a8288f191.EncodeMap(__obf_c26f5adfb3152475)
}

func (__obf_84d0d31a8288f191 *Encoder) EncodeMap(__obf_c26f5adfb3152475 map[string]any) error {
	if __obf_c26f5adfb3152475 == nil {
		return __obf_84d0d31a8288f191.EncodeNil()
	}
	if __obf_3aa78ad28f50ed46 := __obf_84d0d31a8288f191.EncodeMapLen(len(__obf_c26f5adfb3152475)); __obf_3aa78ad28f50ed46 != nil {
		return __obf_3aa78ad28f50ed46
	}
	for __obf_8f7121221504784f, __obf_a03d608795b584a3 := range __obf_c26f5adfb3152475 {
		if __obf_3aa78ad28f50ed46 := __obf_84d0d31a8288f191.EncodeString(__obf_8f7121221504784f); __obf_3aa78ad28f50ed46 != nil {
			return __obf_3aa78ad28f50ed46
		}
		if __obf_3aa78ad28f50ed46 := __obf_84d0d31a8288f191.Encode(__obf_a03d608795b584a3); __obf_3aa78ad28f50ed46 != nil {
			return __obf_3aa78ad28f50ed46
		}
	}
	return nil
}

func (__obf_84d0d31a8288f191 *Encoder) EncodeMapSorted(__obf_c26f5adfb3152475 map[string]any) error {
	if __obf_c26f5adfb3152475 == nil {
		return __obf_84d0d31a8288f191.EncodeNil()
	}
	if __obf_3aa78ad28f50ed46 := __obf_84d0d31a8288f191.EncodeMapLen(len(__obf_c26f5adfb3152475)); __obf_3aa78ad28f50ed46 != nil {
		return __obf_3aa78ad28f50ed46
	}
	__obf_82030dcad96358f1 := make([]string, 0, len(__obf_c26f5adfb3152475))

	for __obf_a86a6197574e9220 := range __obf_c26f5adfb3152475 {
		__obf_82030dcad96358f1 = append(__obf_82030dcad96358f1, __obf_a86a6197574e9220)
	}

	sort.Strings(__obf_82030dcad96358f1)

	for _, __obf_a86a6197574e9220 := range __obf_82030dcad96358f1 {
		if __obf_3aa78ad28f50ed46 := __obf_84d0d31a8288f191.EncodeString(__obf_a86a6197574e9220); __obf_3aa78ad28f50ed46 != nil {
			return __obf_3aa78ad28f50ed46
		}
		if __obf_3aa78ad28f50ed46 := __obf_84d0d31a8288f191.Encode(__obf_c26f5adfb3152475[__obf_a86a6197574e9220]); __obf_3aa78ad28f50ed46 != nil {
			return __obf_3aa78ad28f50ed46
		}
	}

	return nil
}

func (__obf_84d0d31a8288f191 *Encoder) __obf_fd6039710e723b44(__obf_c26f5adfb3152475 map[string]bool) error {
	__obf_82030dcad96358f1 := make([]string, 0, len(__obf_c26f5adfb3152475))
	for __obf_a86a6197574e9220 := range __obf_c26f5adfb3152475 {
		__obf_82030dcad96358f1 = append(__obf_82030dcad96358f1, __obf_a86a6197574e9220)
	}
	sort.Strings(__obf_82030dcad96358f1)

	for _, __obf_a86a6197574e9220 := range __obf_82030dcad96358f1 {
		__obf_3aa78ad28f50ed46 := __obf_84d0d31a8288f191.EncodeString(__obf_a86a6197574e9220)
		if __obf_3aa78ad28f50ed46 != nil {
			return __obf_3aa78ad28f50ed46
		}
		if __obf_3aa78ad28f50ed46 = __obf_84d0d31a8288f191.EncodeBool(__obf_c26f5adfb3152475[__obf_a86a6197574e9220]); __obf_3aa78ad28f50ed46 != nil {
			return __obf_3aa78ad28f50ed46
		}
	}

	return nil
}

func (__obf_84d0d31a8288f191 *Encoder) __obf_39e93fec82dd6558(__obf_c26f5adfb3152475 map[string]string) error {
	__obf_82030dcad96358f1 := make([]string, 0, len(__obf_c26f5adfb3152475))
	for __obf_a86a6197574e9220 := range __obf_c26f5adfb3152475 {
		__obf_82030dcad96358f1 = append(__obf_82030dcad96358f1, __obf_a86a6197574e9220)
	}
	sort.Strings(__obf_82030dcad96358f1)

	for _, __obf_a86a6197574e9220 := range __obf_82030dcad96358f1 {
		__obf_3aa78ad28f50ed46 := __obf_84d0d31a8288f191.EncodeString(__obf_a86a6197574e9220)
		if __obf_3aa78ad28f50ed46 != nil {
			return __obf_3aa78ad28f50ed46
		}
		if __obf_3aa78ad28f50ed46 = __obf_84d0d31a8288f191.EncodeString(__obf_c26f5adfb3152475[__obf_a86a6197574e9220]); __obf_3aa78ad28f50ed46 != nil {
			return __obf_3aa78ad28f50ed46
		}
	}

	return nil
}

func (__obf_84d0d31a8288f191 *Encoder) EncodeMapLen(__obf_48b3b71f73d35432 int) error {
	if __obf_48b3b71f73d35432 < 16 {
		return __obf_84d0d31a8288f191.__obf_4268b610398f5a51(FixedMapLow | byte(__obf_48b3b71f73d35432))
	}
	if __obf_48b3b71f73d35432 <= math.MaxUint16 {
		return __obf_84d0d31a8288f191.__obf_941ea548b75e478e(Map16, uint16(__obf_48b3b71f73d35432))
	}
	return __obf_84d0d31a8288f191.__obf_7636556325987942(Map32, uint32(__obf_48b3b71f73d35432))
}

func __obf_bb1447402cee3139(__obf_84d0d31a8288f191 *Encoder, __obf_14a1657cb1cdf60b reflect.Value) error {
	__obf_94f9ea99a9fa147b := __obf_be0ec04e52f483d9.Fields(__obf_14a1657cb1cdf60b.Type(), __obf_84d0d31a8288f191.__obf_ec5d9a4ce8dd41be)
	if __obf_84d0d31a8288f191.__obf_6aa0efd537415192&__obf_a4be32a461c6088d != 0 || __obf_94f9ea99a9fa147b.AsArray {
		return __obf_c33a46b2f3f92a4c(__obf_84d0d31a8288f191, __obf_14a1657cb1cdf60b, __obf_94f9ea99a9fa147b.List)
	}
	__obf_fdeb38537c10bbcb := __obf_94f9ea99a9fa147b.OmitEmpty(__obf_84d0d31a8288f191, __obf_14a1657cb1cdf60b)

	if __obf_3aa78ad28f50ed46 := __obf_84d0d31a8288f191.EncodeMapLen(len(__obf_fdeb38537c10bbcb)); __obf_3aa78ad28f50ed46 != nil {
		return __obf_3aa78ad28f50ed46
	}

	for _, __obf_40a54345a49c3cd5 := range __obf_fdeb38537c10bbcb {
		if __obf_3aa78ad28f50ed46 := __obf_84d0d31a8288f191.EncodeString(__obf_40a54345a49c3cd5.__obf_90a8fc4880887fcb); __obf_3aa78ad28f50ed46 != nil {
			return __obf_3aa78ad28f50ed46
		}
		if __obf_3aa78ad28f50ed46 := __obf_40a54345a49c3cd5.EncodeValue(__obf_84d0d31a8288f191, __obf_14a1657cb1cdf60b); __obf_3aa78ad28f50ed46 != nil {
			return __obf_3aa78ad28f50ed46
		}
	}

	return nil
}

func __obf_c33a46b2f3f92a4c(__obf_84d0d31a8288f191 *Encoder, __obf_14a1657cb1cdf60b reflect.Value, __obf_fdeb38537c10bbcb []*__obf_90e4cc2c82c33de0) error {
	if __obf_3aa78ad28f50ed46 := __obf_84d0d31a8288f191.EncodeArrayLen(len(__obf_fdeb38537c10bbcb)); __obf_3aa78ad28f50ed46 != nil {
		return __obf_3aa78ad28f50ed46
	}
	for _, __obf_40a54345a49c3cd5 := range __obf_fdeb38537c10bbcb {
		if __obf_3aa78ad28f50ed46 := __obf_40a54345a49c3cd5.EncodeValue(__obf_84d0d31a8288f191, __obf_14a1657cb1cdf60b); __obf_3aa78ad28f50ed46 != nil {
			return __obf_3aa78ad28f50ed46
		}
	}
	return nil
}
