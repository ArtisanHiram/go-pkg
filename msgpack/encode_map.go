package __obf_a4aad98aaf3178f4

import (
	"math"
	"reflect"
	"sort"
)

func __obf_c6d71f18960117fe(__obf_2c8e97779108ab17 *Encoder, __obf_a1f43267eeb48a1a reflect.Value) error {
	if __obf_a1f43267eeb48a1a.IsNil() {
		return __obf_2c8e97779108ab17.EncodeNil()
	}

	if __obf_4061ca0039150c39 := __obf_2c8e97779108ab17.EncodeMapLen(__obf_a1f43267eeb48a1a.Len()); __obf_4061ca0039150c39 != nil {
		return __obf_4061ca0039150c39
	}
	__obf_a932f1a9a32315d8 := __obf_a1f43267eeb48a1a.MapRange()
	for __obf_a932f1a9a32315d8.Next() {
		if __obf_4061ca0039150c39 := __obf_2c8e97779108ab17.EncodeValue(__obf_a932f1a9a32315d8.Key()); __obf_4061ca0039150c39 != nil {
			return __obf_4061ca0039150c39
		}
		if __obf_4061ca0039150c39 := __obf_2c8e97779108ab17.EncodeValue(__obf_a932f1a9a32315d8.Value()); __obf_4061ca0039150c39 != nil {
			return __obf_4061ca0039150c39
		}
	}

	return nil
}

func __obf_a8b5c3d4b26ca8c6(__obf_2c8e97779108ab17 *Encoder, __obf_a1f43267eeb48a1a reflect.Value) error {
	if __obf_a1f43267eeb48a1a.IsNil() {
		return __obf_2c8e97779108ab17.EncodeNil()
	}

	if __obf_4061ca0039150c39 := __obf_2c8e97779108ab17.EncodeMapLen(__obf_a1f43267eeb48a1a.Len()); __obf_4061ca0039150c39 != nil {
		return __obf_4061ca0039150c39
	}
	__obf_66cc4b26e3c9cdbb := __obf_a1f43267eeb48a1a.Convert(__obf_6742791176d1c4ba).Interface().(map[string]bool)
	if __obf_2c8e97779108ab17.__obf_a8400e70a712e9fa&__obf_1f7890468d217f4e != 0 {
		return __obf_2c8e97779108ab17.__obf_268e1476c4c0a178(__obf_66cc4b26e3c9cdbb)
	}

	for __obf_25e3964b56bdb154, __obf_e23729dae69c9f0b := range __obf_66cc4b26e3c9cdbb {
		if __obf_4061ca0039150c39 := __obf_2c8e97779108ab17.EncodeString(__obf_25e3964b56bdb154); __obf_4061ca0039150c39 != nil {
			return __obf_4061ca0039150c39
		}
		if __obf_4061ca0039150c39 := __obf_2c8e97779108ab17.EncodeBool(__obf_e23729dae69c9f0b); __obf_4061ca0039150c39 != nil {
			return __obf_4061ca0039150c39
		}
	}

	return nil
}

func __obf_af0bc339c396d17f(__obf_2c8e97779108ab17 *Encoder, __obf_a1f43267eeb48a1a reflect.Value) error {
	if __obf_a1f43267eeb48a1a.IsNil() {
		return __obf_2c8e97779108ab17.EncodeNil()
	}

	if __obf_4061ca0039150c39 := __obf_2c8e97779108ab17.EncodeMapLen(__obf_a1f43267eeb48a1a.Len()); __obf_4061ca0039150c39 != nil {
		return __obf_4061ca0039150c39
	}
	__obf_66cc4b26e3c9cdbb := __obf_a1f43267eeb48a1a.Convert(__obf_23cc385d7847ade3).Interface().(map[string]string)
	if __obf_2c8e97779108ab17.__obf_a8400e70a712e9fa&__obf_1f7890468d217f4e != 0 {
		return __obf_2c8e97779108ab17.__obf_959e7eaf3c1e489b(__obf_66cc4b26e3c9cdbb)
	}

	for __obf_25e3964b56bdb154, __obf_e23729dae69c9f0b := range __obf_66cc4b26e3c9cdbb {
		if __obf_4061ca0039150c39 := __obf_2c8e97779108ab17.EncodeString(__obf_25e3964b56bdb154); __obf_4061ca0039150c39 != nil {
			return __obf_4061ca0039150c39
		}
		if __obf_4061ca0039150c39 := __obf_2c8e97779108ab17.EncodeString(__obf_e23729dae69c9f0b); __obf_4061ca0039150c39 != nil {
			return __obf_4061ca0039150c39
		}
	}

	return nil
}

func __obf_76c5385b22000a50(__obf_2c8e97779108ab17 *Encoder, __obf_a1f43267eeb48a1a reflect.Value) error {
	if __obf_a1f43267eeb48a1a.IsNil() {
		return __obf_2c8e97779108ab17.EncodeNil()
	}
	__obf_66cc4b26e3c9cdbb := __obf_a1f43267eeb48a1a.Convert(__obf_ef257d4bd242e017).Interface().(map[string]any)
	if __obf_2c8e97779108ab17.__obf_a8400e70a712e9fa&__obf_1f7890468d217f4e != 0 {
		return __obf_2c8e97779108ab17.EncodeMapSorted(__obf_66cc4b26e3c9cdbb)
	}
	return __obf_2c8e97779108ab17.EncodeMap(__obf_66cc4b26e3c9cdbb)
}

func (__obf_2c8e97779108ab17 *Encoder) EncodeMap(__obf_66cc4b26e3c9cdbb map[string]any) error {
	if __obf_66cc4b26e3c9cdbb == nil {
		return __obf_2c8e97779108ab17.EncodeNil()
	}
	if __obf_4061ca0039150c39 := __obf_2c8e97779108ab17.EncodeMapLen(len(__obf_66cc4b26e3c9cdbb)); __obf_4061ca0039150c39 != nil {
		return __obf_4061ca0039150c39
	}
	for __obf_25e3964b56bdb154, __obf_e23729dae69c9f0b := range __obf_66cc4b26e3c9cdbb {
		if __obf_4061ca0039150c39 := __obf_2c8e97779108ab17.EncodeString(__obf_25e3964b56bdb154); __obf_4061ca0039150c39 != nil {
			return __obf_4061ca0039150c39
		}
		if __obf_4061ca0039150c39 := __obf_2c8e97779108ab17.Encode(__obf_e23729dae69c9f0b); __obf_4061ca0039150c39 != nil {
			return __obf_4061ca0039150c39
		}
	}
	return nil
}

func (__obf_2c8e97779108ab17 *Encoder) EncodeMapSorted(__obf_66cc4b26e3c9cdbb map[string]any) error {
	if __obf_66cc4b26e3c9cdbb == nil {
		return __obf_2c8e97779108ab17.EncodeNil()
	}
	if __obf_4061ca0039150c39 := __obf_2c8e97779108ab17.EncodeMapLen(len(__obf_66cc4b26e3c9cdbb)); __obf_4061ca0039150c39 != nil {
		return __obf_4061ca0039150c39
	}
	__obf_52d9f898077dc27a := make([]string, 0, len(__obf_66cc4b26e3c9cdbb))

	for __obf_e8b7c7defe7ee5ec := range __obf_66cc4b26e3c9cdbb {
		__obf_52d9f898077dc27a = append(__obf_52d9f898077dc27a, __obf_e8b7c7defe7ee5ec)
	}

	sort.Strings(__obf_52d9f898077dc27a)

	for _, __obf_e8b7c7defe7ee5ec := range __obf_52d9f898077dc27a {
		if __obf_4061ca0039150c39 := __obf_2c8e97779108ab17.EncodeString(__obf_e8b7c7defe7ee5ec); __obf_4061ca0039150c39 != nil {
			return __obf_4061ca0039150c39
		}
		if __obf_4061ca0039150c39 := __obf_2c8e97779108ab17.Encode(__obf_66cc4b26e3c9cdbb[__obf_e8b7c7defe7ee5ec]); __obf_4061ca0039150c39 != nil {
			return __obf_4061ca0039150c39
		}
	}

	return nil
}

func (__obf_2c8e97779108ab17 *Encoder) __obf_268e1476c4c0a178(__obf_66cc4b26e3c9cdbb map[string]bool) error {
	__obf_52d9f898077dc27a := make([]string, 0, len(__obf_66cc4b26e3c9cdbb))
	for __obf_e8b7c7defe7ee5ec := range __obf_66cc4b26e3c9cdbb {
		__obf_52d9f898077dc27a = append(__obf_52d9f898077dc27a, __obf_e8b7c7defe7ee5ec)
	}
	sort.Strings(__obf_52d9f898077dc27a)

	for _, __obf_e8b7c7defe7ee5ec := range __obf_52d9f898077dc27a {
		__obf_4061ca0039150c39 := __obf_2c8e97779108ab17.EncodeString(__obf_e8b7c7defe7ee5ec)
		if __obf_4061ca0039150c39 != nil {
			return __obf_4061ca0039150c39
		}
		if __obf_4061ca0039150c39 = __obf_2c8e97779108ab17.EncodeBool(__obf_66cc4b26e3c9cdbb[__obf_e8b7c7defe7ee5ec]); __obf_4061ca0039150c39 != nil {
			return __obf_4061ca0039150c39
		}
	}

	return nil
}

func (__obf_2c8e97779108ab17 *Encoder) __obf_959e7eaf3c1e489b(__obf_66cc4b26e3c9cdbb map[string]string) error {
	__obf_52d9f898077dc27a := make([]string, 0, len(__obf_66cc4b26e3c9cdbb))
	for __obf_e8b7c7defe7ee5ec := range __obf_66cc4b26e3c9cdbb {
		__obf_52d9f898077dc27a = append(__obf_52d9f898077dc27a, __obf_e8b7c7defe7ee5ec)
	}
	sort.Strings(__obf_52d9f898077dc27a)

	for _, __obf_e8b7c7defe7ee5ec := range __obf_52d9f898077dc27a {
		__obf_4061ca0039150c39 := __obf_2c8e97779108ab17.EncodeString(__obf_e8b7c7defe7ee5ec)
		if __obf_4061ca0039150c39 != nil {
			return __obf_4061ca0039150c39
		}
		if __obf_4061ca0039150c39 = __obf_2c8e97779108ab17.EncodeString(__obf_66cc4b26e3c9cdbb[__obf_e8b7c7defe7ee5ec]); __obf_4061ca0039150c39 != nil {
			return __obf_4061ca0039150c39
		}
	}

	return nil
}

func (__obf_2c8e97779108ab17 *Encoder) EncodeMapLen(__obf_4431cbe3c2f63394 int) error {
	if __obf_4431cbe3c2f63394 < 16 {
		return __obf_2c8e97779108ab17.__obf_c087409347698442(FixedMapLow | byte(__obf_4431cbe3c2f63394))
	}
	if __obf_4431cbe3c2f63394 <= math.MaxUint16 {
		return __obf_2c8e97779108ab17.__obf_71321d3206264f3b(Map16, uint16(__obf_4431cbe3c2f63394))
	}
	return __obf_2c8e97779108ab17.__obf_a40df49dbc3bfc37(Map32, uint32(__obf_4431cbe3c2f63394))
}

func __obf_8fcfe77f763e1de9(__obf_2c8e97779108ab17 *Encoder, __obf_7500ff0c7bba82eb reflect.Value) error {
	__obf_eb37ef6b7fc3d8f3 := __obf_bd827009b4ede0e2.Fields(__obf_7500ff0c7bba82eb.Type(), __obf_2c8e97779108ab17.__obf_72ba705ed504d210)
	if __obf_2c8e97779108ab17.__obf_a8400e70a712e9fa&__obf_3e711425f784a6a6 != 0 || __obf_eb37ef6b7fc3d8f3.AsArray {
		return __obf_e970405927cf1e5a(__obf_2c8e97779108ab17, __obf_7500ff0c7bba82eb, __obf_eb37ef6b7fc3d8f3.List)
	}
	__obf_0cf794214f02df4d := __obf_eb37ef6b7fc3d8f3.OmitEmpty(__obf_2c8e97779108ab17, __obf_7500ff0c7bba82eb)

	if __obf_4061ca0039150c39 := __obf_2c8e97779108ab17.EncodeMapLen(len(__obf_0cf794214f02df4d)); __obf_4061ca0039150c39 != nil {
		return __obf_4061ca0039150c39
	}

	for _, __obf_da4a2296298d318f := range __obf_0cf794214f02df4d {
		if __obf_4061ca0039150c39 := __obf_2c8e97779108ab17.EncodeString(__obf_da4a2296298d318f.__obf_071b55c16b16fe35); __obf_4061ca0039150c39 != nil {
			return __obf_4061ca0039150c39
		}
		if __obf_4061ca0039150c39 := __obf_da4a2296298d318f.EncodeValue(__obf_2c8e97779108ab17, __obf_7500ff0c7bba82eb); __obf_4061ca0039150c39 != nil {
			return __obf_4061ca0039150c39
		}
	}

	return nil
}

func __obf_e970405927cf1e5a(__obf_2c8e97779108ab17 *Encoder, __obf_7500ff0c7bba82eb reflect.Value, __obf_0cf794214f02df4d []*__obf_efecfc090178574b) error {
	if __obf_4061ca0039150c39 := __obf_2c8e97779108ab17.EncodeArrayLen(len(__obf_0cf794214f02df4d)); __obf_4061ca0039150c39 != nil {
		return __obf_4061ca0039150c39
	}
	for _, __obf_da4a2296298d318f := range __obf_0cf794214f02df4d {
		if __obf_4061ca0039150c39 := __obf_da4a2296298d318f.EncodeValue(__obf_2c8e97779108ab17, __obf_7500ff0c7bba82eb); __obf_4061ca0039150c39 != nil {
			return __obf_4061ca0039150c39
		}
	}
	return nil
}
