package __obf_de86cdc8ae98b45b

import (
	"errors"
	"fmt"
	"reflect"
)

var __obf_4c7bbc253fc40d15 = errors.New("msgpack: number of fields in array-encoded struct has changed")

var (
	__obf_89b1bbe408c24ecc = reflect.TypeOf((*map[string]string)(nil))
	__obf_34d6b0138b97b155 = __obf_89b1bbe408c24ecc.
				Elem()
	__obf_d8ad9b13ddaf6b80 = reflect.TypeOf((*map[string]bool)(nil))
	__obf_9ea30858b9b6e940 = __obf_d8ad9b13ddaf6b80.
				Elem()
)

var (
	__obf_fad714615854b426 = reflect.TypeOf((*map[string]any)(nil))
	__obf_e8e62bb83070a108 = __obf_fad714615854b426.
				Elem()
)

func __obf_fc3c99dfcc57d122(__obf_dcaad1165bb07af9 *Decoder, __obf_17732590722140e7 reflect.Value) error {
	__obf_2b0247e819b1bf4a, __obf_0feb3528b7b709ec := __obf_dcaad1165bb07af9.DecodeMapLen()
	if __obf_0feb3528b7b709ec != nil {
		return __obf_0feb3528b7b709ec
	}
	__obf_2dea2b2773c8afdf := __obf_17732590722140e7.Type()
	if __obf_2b0247e819b1bf4a == -1 {
		__obf_17732590722140e7.
			Set(reflect.Zero(__obf_2dea2b2773c8afdf))
		return nil
	}

	if __obf_17732590722140e7.IsNil() {
		__obf_89ea3eb7b8e3cc7f := __obf_2b0247e819b1bf4a
		if __obf_dcaad1165bb07af9.__obf_a15a642720e39c3e&__obf_c646a26ee823abc3 == 0 {
			__obf_89ea3eb7b8e3cc7f = min(__obf_89ea3eb7b8e3cc7f, __obf_5e8134fa61b47535)
		}
		__obf_17732590722140e7.
			Set(reflect.MakeMapWithSize(__obf_2dea2b2773c8afdf, __obf_89ea3eb7b8e3cc7f))
	}
	if __obf_2b0247e819b1bf4a == 0 {
		return nil
	}

	return __obf_dcaad1165bb07af9.__obf_3307ce2fba98dd22(__obf_17732590722140e7, __obf_2b0247e819b1bf4a)
}

func (__obf_dcaad1165bb07af9 *Decoder) __obf_8468f219671f8d5a() (any, error) {
	if __obf_dcaad1165bb07af9.__obf_c1135a4c578558d5 != nil {
		return __obf_dcaad1165bb07af9.__obf_c1135a4c578558d5(__obf_dcaad1165bb07af9)
	}
	return __obf_dcaad1165bb07af9.DecodeMap()
}

// DecodeMapLen decodes map length. Length is -1 when map is nil.
func (__obf_dcaad1165bb07af9 *Decoder) DecodeMapLen() (int, error) {
	__obf_ecf6d2d3253a058d, __obf_0feb3528b7b709ec := __obf_dcaad1165bb07af9.__obf_74ac05d9a9da01f4()
	if __obf_0feb3528b7b709ec != nil {
		return 0, __obf_0feb3528b7b709ec
	}

	if IsExt(__obf_ecf6d2d3253a058d) {
		if __obf_0feb3528b7b709ec = __obf_dcaad1165bb07af9.__obf_9019231c4f64c0e5(__obf_ecf6d2d3253a058d); __obf_0feb3528b7b709ec != nil {
			return 0, __obf_0feb3528b7b709ec
		}
		__obf_ecf6d2d3253a058d, __obf_0feb3528b7b709ec = __obf_dcaad1165bb07af9.__obf_74ac05d9a9da01f4()
		if __obf_0feb3528b7b709ec != nil {
			return 0, __obf_0feb3528b7b709ec
		}
	}
	return __obf_dcaad1165bb07af9.__obf_0317c5198258183c(__obf_ecf6d2d3253a058d)
}

func (__obf_dcaad1165bb07af9 *Decoder) __obf_0317c5198258183c(__obf_ecf6d2d3253a058d byte) (int, error) {
	if __obf_ecf6d2d3253a058d == Nil {
		return -1, nil
	}
	if __obf_ecf6d2d3253a058d >= FixedMapLow && __obf_ecf6d2d3253a058d <= FixedMapHigh {
		return int(__obf_ecf6d2d3253a058d & FixedMapMask), nil
	}
	if __obf_ecf6d2d3253a058d == Map16 {
		__obf_4d6037e256b6f76d, __obf_0feb3528b7b709ec := __obf_dcaad1165bb07af9.uint16()
		return int(__obf_4d6037e256b6f76d), __obf_0feb3528b7b709ec
	}
	if __obf_ecf6d2d3253a058d == Map32 {
		__obf_4d6037e256b6f76d, __obf_0feb3528b7b709ec := __obf_dcaad1165bb07af9.uint32()
		return int(__obf_4d6037e256b6f76d), __obf_0feb3528b7b709ec
	}
	return 0, __obf_150b3812dfd829ef{__obf_34e3ba264a6bb541: __obf_ecf6d2d3253a058d, __obf_ee374540697f97c0: "map length"}
}

func __obf_62de82f75c0986e4(__obf_dcaad1165bb07af9 *Decoder, __obf_17732590722140e7 reflect.Value) error {
	__obf_49790e9e8244dd55 := __obf_17732590722140e7.Addr().Convert(__obf_89b1bbe408c24ecc).Interface().(*map[string]string)
	return __obf_dcaad1165bb07af9.__obf_234bb66c0ec0a03d(__obf_49790e9e8244dd55)
}

func (__obf_dcaad1165bb07af9 *Decoder) __obf_234bb66c0ec0a03d(__obf_01f0154f1fce8e5a *map[string]string) error {
	__obf_4d6037e256b6f76d, __obf_0feb3528b7b709ec := __obf_dcaad1165bb07af9.DecodeMapLen()
	if __obf_0feb3528b7b709ec != nil {
		return __obf_0feb3528b7b709ec
	}
	if __obf_4d6037e256b6f76d == -1 {
		*__obf_01f0154f1fce8e5a = nil
		return nil
	}
	__obf_2a26e7104b4c4373 := *__obf_01f0154f1fce8e5a
	if __obf_2a26e7104b4c4373 == nil {
		__obf_89ea3eb7b8e3cc7f := __obf_4d6037e256b6f76d
		if __obf_dcaad1165bb07af9.__obf_a15a642720e39c3e&__obf_c646a26ee823abc3 == 0 {
			__obf_89ea3eb7b8e3cc7f = min(__obf_4d6037e256b6f76d, __obf_5e8134fa61b47535)
		}
		*__obf_01f0154f1fce8e5a = make(map[string]string, __obf_89ea3eb7b8e3cc7f)
		__obf_2a26e7104b4c4373 = *__obf_01f0154f1fce8e5a
	}

	for range __obf_4d6037e256b6f76d {
		__obf_420732d362059ffa, __obf_0feb3528b7b709ec := __obf_dcaad1165bb07af9.DecodeString()
		if __obf_0feb3528b7b709ec != nil {
			return __obf_0feb3528b7b709ec
		}
		__obf_a6cb485687dffc12, __obf_0feb3528b7b709ec := __obf_dcaad1165bb07af9.DecodeString()
		if __obf_0feb3528b7b709ec != nil {
			return __obf_0feb3528b7b709ec
		}
		__obf_2a26e7104b4c4373[__obf_420732d362059ffa] = __obf_a6cb485687dffc12
	}

	return nil
}

func __obf_b673e9c8a63629b4(__obf_dcaad1165bb07af9 *Decoder, __obf_17732590722140e7 reflect.Value) error {
	__obf_01f0154f1fce8e5a := __obf_17732590722140e7.Addr().Convert(__obf_fad714615854b426).Interface().(*map[string]any)
	return __obf_dcaad1165bb07af9.__obf_0886d84e85a12f76(__obf_01f0154f1fce8e5a)
}

func (__obf_dcaad1165bb07af9 *Decoder) __obf_0886d84e85a12f76(__obf_01f0154f1fce8e5a *map[string]any) error {
	__obf_2a26e7104b4c4373, __obf_0feb3528b7b709ec := __obf_dcaad1165bb07af9.DecodeMap()
	if __obf_0feb3528b7b709ec != nil {
		return __obf_0feb3528b7b709ec
	}
	*__obf_01f0154f1fce8e5a = __obf_2a26e7104b4c4373
	return nil
}

func (__obf_dcaad1165bb07af9 *Decoder) DecodeMap() (map[string]any, error) {
	__obf_2b0247e819b1bf4a, __obf_0feb3528b7b709ec := __obf_dcaad1165bb07af9.DecodeMapLen()
	if __obf_0feb3528b7b709ec != nil {
		return nil, __obf_0feb3528b7b709ec
	}

	if __obf_2b0247e819b1bf4a == -1 {
		return nil, nil
	}
	__obf_2a26e7104b4c4373 := make(map[string]any, __obf_2b0247e819b1bf4a)

	for range __obf_2b0247e819b1bf4a {
		__obf_420732d362059ffa, __obf_0feb3528b7b709ec := __obf_dcaad1165bb07af9.DecodeString()
		if __obf_0feb3528b7b709ec != nil {
			return nil, __obf_0feb3528b7b709ec
		}
		__obf_a6cb485687dffc12, __obf_0feb3528b7b709ec := __obf_dcaad1165bb07af9.__obf_7d3129f334945713()
		if __obf_0feb3528b7b709ec != nil {
			return nil, __obf_0feb3528b7b709ec
		}
		__obf_2a26e7104b4c4373[__obf_420732d362059ffa] = __obf_a6cb485687dffc12
	}

	return __obf_2a26e7104b4c4373, nil
}

func (__obf_dcaad1165bb07af9 *Decoder) DecodeUntypedMap() (map[any]any, error) {
	__obf_2b0247e819b1bf4a, __obf_0feb3528b7b709ec := __obf_dcaad1165bb07af9.DecodeMapLen()
	if __obf_0feb3528b7b709ec != nil {
		return nil, __obf_0feb3528b7b709ec
	}

	if __obf_2b0247e819b1bf4a == -1 {
		return nil, nil
	}
	__obf_2a26e7104b4c4373 := make(map[any]any, __obf_2b0247e819b1bf4a)

	for range __obf_2b0247e819b1bf4a {
		__obf_420732d362059ffa, __obf_0feb3528b7b709ec := __obf_dcaad1165bb07af9.__obf_7d3129f334945713()
		if __obf_0feb3528b7b709ec != nil {
			return nil, __obf_0feb3528b7b709ec
		}
		__obf_a6cb485687dffc12, __obf_0feb3528b7b709ec := __obf_dcaad1165bb07af9.__obf_7d3129f334945713()
		if __obf_0feb3528b7b709ec != nil {
			return nil, __obf_0feb3528b7b709ec
		}
		__obf_2a26e7104b4c4373[__obf_420732d362059ffa] = __obf_a6cb485687dffc12
	}

	return __obf_2a26e7104b4c4373, nil
}

// DecodeTypedMap decodes a typed map. Typed map is a map that has a fixed type for keys and values.
// Key and value types may be different.
func (__obf_dcaad1165bb07af9 *Decoder) DecodeTypedMap() (any, error) {
	__obf_2b0247e819b1bf4a, __obf_0feb3528b7b709ec := __obf_dcaad1165bb07af9.DecodeMapLen()
	if __obf_0feb3528b7b709ec != nil {
		return nil, __obf_0feb3528b7b709ec
	}
	if __obf_2b0247e819b1bf4a <= 0 {
		return nil, nil
	}
	__obf_9cab2959bb95a876, __obf_0feb3528b7b709ec := __obf_dcaad1165bb07af9.__obf_7d3129f334945713()
	if __obf_0feb3528b7b709ec != nil {
		return nil, __obf_0feb3528b7b709ec
	}
	__obf_1926a1e373f9e647, __obf_0feb3528b7b709ec := __obf_dcaad1165bb07af9.__obf_7d3129f334945713()
	if __obf_0feb3528b7b709ec != nil {
		return nil, __obf_0feb3528b7b709ec
	}
	__obf_bad9c9cc3e01eaf6 := reflect.TypeOf(__obf_9cab2959bb95a876)
	__obf_3b6e1488559a31f6 := reflect.TypeOf(__obf_1926a1e373f9e647)

	if !__obf_bad9c9cc3e01eaf6.Comparable() {
		return nil, fmt.Errorf("msgpack: unsupported map key: %s", __obf_bad9c9cc3e01eaf6.String())
	}
	__obf_e6043d225e420b8b := reflect.MapOf(__obf_bad9c9cc3e01eaf6, __obf_3b6e1488559a31f6)
	__obf_89ea3eb7b8e3cc7f := __obf_2b0247e819b1bf4a
	if __obf_dcaad1165bb07af9.__obf_a15a642720e39c3e&__obf_c646a26ee823abc3 == 0 {
		__obf_89ea3eb7b8e3cc7f = min(__obf_89ea3eb7b8e3cc7f, __obf_5e8134fa61b47535)
	}
	__obf_9f8ff94f6953c947 := reflect.MakeMapWithSize(__obf_e6043d225e420b8b, __obf_89ea3eb7b8e3cc7f)
	__obf_9f8ff94f6953c947.
		SetMapIndex(reflect.ValueOf(__obf_9cab2959bb95a876), reflect.ValueOf(__obf_1926a1e373f9e647))
	__obf_2b0247e819b1bf4a--
	if __obf_0feb3528b7b709ec := __obf_dcaad1165bb07af9.__obf_3307ce2fba98dd22(__obf_9f8ff94f6953c947, __obf_2b0247e819b1bf4a); __obf_0feb3528b7b709ec != nil {
		return nil, __obf_0feb3528b7b709ec
	}

	return __obf_9f8ff94f6953c947.Interface(), nil
}

func (__obf_dcaad1165bb07af9 *Decoder) __obf_3307ce2fba98dd22(__obf_17732590722140e7 reflect.Value, __obf_2b0247e819b1bf4a int) error {
	var (
		__obf_2dea2b2773c8afdf = __obf_17732590722140e7.
					Type()
		__obf_bad9c9cc3e01eaf6 = __obf_2dea2b2773c8afdf.
					Key()
		__obf_3b6e1488559a31f6 = __obf_2dea2b2773c8afdf.
					Elem()
	)
	for range __obf_2b0247e819b1bf4a {
		__obf_420732d362059ffa := __obf_dcaad1165bb07af9.__obf_598c5a4b42040fb1(__obf_bad9c9cc3e01eaf6).Elem()
		if __obf_0feb3528b7b709ec := __obf_dcaad1165bb07af9.DecodeValue(__obf_420732d362059ffa); __obf_0feb3528b7b709ec != nil {
			return __obf_0feb3528b7b709ec
		}
		__obf_a6cb485687dffc12 := __obf_dcaad1165bb07af9.__obf_598c5a4b42040fb1(__obf_3b6e1488559a31f6).Elem()
		if __obf_0feb3528b7b709ec := __obf_dcaad1165bb07af9.DecodeValue(__obf_a6cb485687dffc12); __obf_0feb3528b7b709ec != nil {
			return __obf_0feb3528b7b709ec
		}
		__obf_17732590722140e7.
			SetMapIndex(__obf_420732d362059ffa, __obf_a6cb485687dffc12)
	}

	return nil
}

func (__obf_dcaad1165bb07af9 *Decoder) __obf_293bf0315964dccb(__obf_ecf6d2d3253a058d byte) error {
	__obf_2b0247e819b1bf4a, __obf_0feb3528b7b709ec := __obf_dcaad1165bb07af9.__obf_0317c5198258183c(__obf_ecf6d2d3253a058d)
	if __obf_0feb3528b7b709ec != nil {
		return __obf_0feb3528b7b709ec
	}
	for range __obf_2b0247e819b1bf4a {
		if __obf_0feb3528b7b709ec := __obf_dcaad1165bb07af9.Skip(); __obf_0feb3528b7b709ec != nil {
			return __obf_0feb3528b7b709ec
		}
		if __obf_0feb3528b7b709ec := __obf_dcaad1165bb07af9.Skip(); __obf_0feb3528b7b709ec != nil {
			return __obf_0feb3528b7b709ec
		}
	}
	return nil
}

func __obf_d6cbe6ba2b1343be(__obf_dcaad1165bb07af9 *Decoder, __obf_17732590722140e7 reflect.Value) error {
	__obf_ecf6d2d3253a058d, __obf_0feb3528b7b709ec := __obf_dcaad1165bb07af9.__obf_74ac05d9a9da01f4()
	if __obf_0feb3528b7b709ec != nil {
		return __obf_0feb3528b7b709ec
	}
	__obf_2b0247e819b1bf4a, __obf_0feb3528b7b709ec := __obf_dcaad1165bb07af9.__obf_0317c5198258183c(__obf_ecf6d2d3253a058d)
	if __obf_0feb3528b7b709ec == nil {
		return __obf_dcaad1165bb07af9.__obf_7f10921311c5c64d(__obf_17732590722140e7, __obf_2b0247e819b1bf4a)
	}

	var __obf_b98e12e36d4a1089 error
	__obf_2b0247e819b1bf4a, __obf_b98e12e36d4a1089 = __obf_dcaad1165bb07af9.__obf_5f964cf181cfcbf3(__obf_ecf6d2d3253a058d)
	if __obf_b98e12e36d4a1089 != nil {
		return __obf_0feb3528b7b709ec
	}

	if __obf_2b0247e819b1bf4a <= 0 {
		__obf_17732590722140e7.
			Set(reflect.Zero(__obf_17732590722140e7.Type()))
		return nil
	}
	__obf_7553cd24ebfd7da5 := __obf_4642342e308ea07b.Fields(__obf_17732590722140e7.Type(), __obf_dcaad1165bb07af9.__obf_603dff1ad8d49539)
	if __obf_2b0247e819b1bf4a != len(__obf_7553cd24ebfd7da5.List) {
		return __obf_4c7bbc253fc40d15
	}

	for _, __obf_cf8245c3f9570af5 := range __obf_7553cd24ebfd7da5.List {
		if __obf_0feb3528b7b709ec := __obf_cf8245c3f9570af5.DecodeValue(__obf_dcaad1165bb07af9, __obf_17732590722140e7); __obf_0feb3528b7b709ec != nil {
			return __obf_0feb3528b7b709ec
		}
	}

	return nil
}

func (__obf_dcaad1165bb07af9 *Decoder) __obf_7f10921311c5c64d(__obf_17732590722140e7 reflect.Value, __obf_2b0247e819b1bf4a int) error {
	if __obf_2b0247e819b1bf4a == -1 {
		__obf_17732590722140e7.
			Set(reflect.Zero(__obf_17732590722140e7.Type()))
		return nil
	}
	__obf_7553cd24ebfd7da5 := __obf_4642342e308ea07b.Fields(__obf_17732590722140e7.Type(), __obf_dcaad1165bb07af9.__obf_603dff1ad8d49539)
	for range __obf_2b0247e819b1bf4a {
		__obf_d7fb5470f7183af5, __obf_0feb3528b7b709ec := __obf_dcaad1165bb07af9.__obf_0605127b8a406c17()
		if __obf_0feb3528b7b709ec != nil {
			return __obf_0feb3528b7b709ec
		}

		if __obf_cf8245c3f9570af5 := __obf_7553cd24ebfd7da5.Map[__obf_d7fb5470f7183af5]; __obf_cf8245c3f9570af5 != nil {
			if __obf_0feb3528b7b709ec := __obf_cf8245c3f9570af5.DecodeValue(__obf_dcaad1165bb07af9, __obf_17732590722140e7); __obf_0feb3528b7b709ec != nil {
				return __obf_0feb3528b7b709ec
			}
			continue
		}

		if __obf_dcaad1165bb07af9.__obf_a15a642720e39c3e&__obf_64c53d062f7fdfd5 != 0 {
			return fmt.Errorf("msgpack: unknown field %q", __obf_d7fb5470f7183af5)
		}
		if __obf_0feb3528b7b709ec := __obf_dcaad1165bb07af9.Skip(); __obf_0feb3528b7b709ec != nil {
			return __obf_0feb3528b7b709ec
		}
	}

	return nil
}
