package __obf_a935eb7f548271a4

import (
	"errors"
	"fmt"
	"reflect"
)

var __obf_33a82116e47fa106 = errors.New("msgpack: number of fields in array-encoded struct has changed")

var (
	__obf_b2b51d791ee37b07 = reflect.TypeOf((*map[string]string)(nil))
	__obf_eead1ffc8ffe8371 = __obf_b2b51d791ee37b07.
				Elem()
	__obf_9282ea7b5a7f8ff4 = reflect.TypeOf((*map[string]bool)(nil))
	__obf_9436e2f0a394e1f9 = __obf_9282ea7b5a7f8ff4.
				Elem()
)

var (
	__obf_b419e839a55aa236 = reflect.TypeOf((*map[string]any)(nil))
	__obf_6782064e44047a94 = __obf_b419e839a55aa236.
				Elem()
)

func __obf_d50c45158b108e4f(__obf_a21885da2425f2b2 *Decoder, __obf_6d570581f4b60dbc reflect.Value) error {
	__obf_326af9bd942662ac, __obf_4d327e1cd40c2e21 := __obf_a21885da2425f2b2.DecodeMapLen()
	if __obf_4d327e1cd40c2e21 != nil {
		return __obf_4d327e1cd40c2e21
	}
	__obf_7bd64e8df8357cab := __obf_6d570581f4b60dbc.Type()
	if __obf_326af9bd942662ac == -1 {
		__obf_6d570581f4b60dbc.
			Set(reflect.Zero(__obf_7bd64e8df8357cab))
		return nil
	}

	if __obf_6d570581f4b60dbc.IsNil() {
		__obf_a3e41621ce2ae57b := __obf_326af9bd942662ac
		if __obf_a21885da2425f2b2.__obf_d617f3769ce16c47&__obf_20f89e0d9adcdf29 == 0 {
			__obf_a3e41621ce2ae57b = min(__obf_a3e41621ce2ae57b, __obf_c02b114384c3720b)
		}
		__obf_6d570581f4b60dbc.
			Set(reflect.MakeMapWithSize(__obf_7bd64e8df8357cab, __obf_a3e41621ce2ae57b))
	}
	if __obf_326af9bd942662ac == 0 {
		return nil
	}

	return __obf_a21885da2425f2b2.__obf_e8eb71d4abd07c0e(__obf_6d570581f4b60dbc, __obf_326af9bd942662ac)
}

func (__obf_a21885da2425f2b2 *Decoder) __obf_1d0c0ae9a1e78fcd() (any, error) {
	if __obf_a21885da2425f2b2.__obf_6ef8e6548ebf1dd5 != nil {
		return __obf_a21885da2425f2b2.__obf_6ef8e6548ebf1dd5(__obf_a21885da2425f2b2)
	}
	return __obf_a21885da2425f2b2.DecodeMap()
}

// DecodeMapLen decodes map length. Length is -1 when map is nil.
func (__obf_a21885da2425f2b2 *Decoder) DecodeMapLen() (int, error) {
	__obf_f5df560f4d67421b, __obf_4d327e1cd40c2e21 := __obf_a21885da2425f2b2.__obf_61e5879fcf7b35ce()
	if __obf_4d327e1cd40c2e21 != nil {
		return 0, __obf_4d327e1cd40c2e21
	}

	if IsExt(__obf_f5df560f4d67421b) {
		if __obf_4d327e1cd40c2e21 = __obf_a21885da2425f2b2.__obf_b5db8a2c243ae15f(__obf_f5df560f4d67421b); __obf_4d327e1cd40c2e21 != nil {
			return 0, __obf_4d327e1cd40c2e21
		}
		__obf_f5df560f4d67421b, __obf_4d327e1cd40c2e21 = __obf_a21885da2425f2b2.__obf_61e5879fcf7b35ce()
		if __obf_4d327e1cd40c2e21 != nil {
			return 0, __obf_4d327e1cd40c2e21
		}
	}
	return __obf_a21885da2425f2b2.__obf_ee4d07636640bafb(__obf_f5df560f4d67421b)
}

func (__obf_a21885da2425f2b2 *Decoder) __obf_ee4d07636640bafb(__obf_f5df560f4d67421b byte) (int, error) {
	if __obf_f5df560f4d67421b == Nil {
		return -1, nil
	}
	if __obf_f5df560f4d67421b >= FixedMapLow && __obf_f5df560f4d67421b <= FixedMapHigh {
		return int(__obf_f5df560f4d67421b & FixedMapMask), nil
	}
	if __obf_f5df560f4d67421b == Map16 {
		__obf_2b6e1001534e539e, __obf_4d327e1cd40c2e21 := __obf_a21885da2425f2b2.uint16()
		return int(__obf_2b6e1001534e539e), __obf_4d327e1cd40c2e21
	}
	if __obf_f5df560f4d67421b == Map32 {
		__obf_2b6e1001534e539e, __obf_4d327e1cd40c2e21 := __obf_a21885da2425f2b2.uint32()
		return int(__obf_2b6e1001534e539e), __obf_4d327e1cd40c2e21
	}
	return 0, __obf_769655cd52f45f6d{__obf_b983039ef2a336eb: __obf_f5df560f4d67421b, __obf_feeba0b68911eadc: "map length"}
}

func __obf_aa35887bffb65a6c(__obf_a21885da2425f2b2 *Decoder, __obf_6d570581f4b60dbc reflect.Value) error {
	__obf_0b9b75168991fde0 := __obf_6d570581f4b60dbc.Addr().Convert(__obf_b2b51d791ee37b07).Interface().(*map[string]string)
	return __obf_a21885da2425f2b2.__obf_17bf0c8c3e4bc6de(__obf_0b9b75168991fde0)
}

func (__obf_a21885da2425f2b2 *Decoder) __obf_17bf0c8c3e4bc6de(__obf_0d8a994785cda6df *map[string]string) error {
	__obf_2b6e1001534e539e, __obf_4d327e1cd40c2e21 := __obf_a21885da2425f2b2.DecodeMapLen()
	if __obf_4d327e1cd40c2e21 != nil {
		return __obf_4d327e1cd40c2e21
	}
	if __obf_2b6e1001534e539e == -1 {
		*__obf_0d8a994785cda6df = nil
		return nil
	}
	__obf_26d12ef0df36a324 := *__obf_0d8a994785cda6df
	if __obf_26d12ef0df36a324 == nil {
		__obf_a3e41621ce2ae57b := __obf_2b6e1001534e539e
		if __obf_a21885da2425f2b2.__obf_d617f3769ce16c47&__obf_20f89e0d9adcdf29 == 0 {
			__obf_a3e41621ce2ae57b = min(__obf_2b6e1001534e539e, __obf_c02b114384c3720b)
		}
		*__obf_0d8a994785cda6df = make(map[string]string, __obf_a3e41621ce2ae57b)
		__obf_26d12ef0df36a324 = *__obf_0d8a994785cda6df
	}

	for range __obf_2b6e1001534e539e {
		__obf_3b7905b85fce1fb1, __obf_4d327e1cd40c2e21 := __obf_a21885da2425f2b2.DecodeString()
		if __obf_4d327e1cd40c2e21 != nil {
			return __obf_4d327e1cd40c2e21
		}
		__obf_b35b927f0bf61896, __obf_4d327e1cd40c2e21 := __obf_a21885da2425f2b2.DecodeString()
		if __obf_4d327e1cd40c2e21 != nil {
			return __obf_4d327e1cd40c2e21
		}
		__obf_26d12ef0df36a324[__obf_3b7905b85fce1fb1] = __obf_b35b927f0bf61896
	}

	return nil
}

func __obf_b269ad4ecc53ab26(__obf_a21885da2425f2b2 *Decoder, __obf_6d570581f4b60dbc reflect.Value) error {
	__obf_0d8a994785cda6df := __obf_6d570581f4b60dbc.Addr().Convert(__obf_b419e839a55aa236).Interface().(*map[string]any)
	return __obf_a21885da2425f2b2.__obf_7b03c6c7ec65752f(__obf_0d8a994785cda6df)
}

func (__obf_a21885da2425f2b2 *Decoder) __obf_7b03c6c7ec65752f(__obf_0d8a994785cda6df *map[string]any) error {
	__obf_26d12ef0df36a324, __obf_4d327e1cd40c2e21 := __obf_a21885da2425f2b2.DecodeMap()
	if __obf_4d327e1cd40c2e21 != nil {
		return __obf_4d327e1cd40c2e21
	}
	*__obf_0d8a994785cda6df = __obf_26d12ef0df36a324
	return nil
}

func (__obf_a21885da2425f2b2 *Decoder) DecodeMap() (map[string]any, error) {
	__obf_326af9bd942662ac, __obf_4d327e1cd40c2e21 := __obf_a21885da2425f2b2.DecodeMapLen()
	if __obf_4d327e1cd40c2e21 != nil {
		return nil, __obf_4d327e1cd40c2e21
	}

	if __obf_326af9bd942662ac == -1 {
		return nil, nil
	}
	__obf_26d12ef0df36a324 := make(map[string]any, __obf_326af9bd942662ac)

	for range __obf_326af9bd942662ac {
		__obf_3b7905b85fce1fb1, __obf_4d327e1cd40c2e21 := __obf_a21885da2425f2b2.DecodeString()
		if __obf_4d327e1cd40c2e21 != nil {
			return nil, __obf_4d327e1cd40c2e21
		}
		__obf_b35b927f0bf61896, __obf_4d327e1cd40c2e21 := __obf_a21885da2425f2b2.__obf_c4227998727b27c6()
		if __obf_4d327e1cd40c2e21 != nil {
			return nil, __obf_4d327e1cd40c2e21
		}
		__obf_26d12ef0df36a324[__obf_3b7905b85fce1fb1] = __obf_b35b927f0bf61896
	}

	return __obf_26d12ef0df36a324, nil
}

func (__obf_a21885da2425f2b2 *Decoder) DecodeUntypedMap() (map[any]any, error) {
	__obf_326af9bd942662ac, __obf_4d327e1cd40c2e21 := __obf_a21885da2425f2b2.DecodeMapLen()
	if __obf_4d327e1cd40c2e21 != nil {
		return nil, __obf_4d327e1cd40c2e21
	}

	if __obf_326af9bd942662ac == -1 {
		return nil, nil
	}
	__obf_26d12ef0df36a324 := make(map[any]any, __obf_326af9bd942662ac)

	for range __obf_326af9bd942662ac {
		__obf_3b7905b85fce1fb1, __obf_4d327e1cd40c2e21 := __obf_a21885da2425f2b2.__obf_c4227998727b27c6()
		if __obf_4d327e1cd40c2e21 != nil {
			return nil, __obf_4d327e1cd40c2e21
		}
		__obf_b35b927f0bf61896, __obf_4d327e1cd40c2e21 := __obf_a21885da2425f2b2.__obf_c4227998727b27c6()
		if __obf_4d327e1cd40c2e21 != nil {
			return nil, __obf_4d327e1cd40c2e21
		}
		__obf_26d12ef0df36a324[__obf_3b7905b85fce1fb1] = __obf_b35b927f0bf61896
	}

	return __obf_26d12ef0df36a324, nil
}

// DecodeTypedMap decodes a typed map. Typed map is a map that has a fixed type for keys and values.
// Key and value types may be different.
func (__obf_a21885da2425f2b2 *Decoder) DecodeTypedMap() (any, error) {
	__obf_326af9bd942662ac, __obf_4d327e1cd40c2e21 := __obf_a21885da2425f2b2.DecodeMapLen()
	if __obf_4d327e1cd40c2e21 != nil {
		return nil, __obf_4d327e1cd40c2e21
	}
	if __obf_326af9bd942662ac <= 0 {
		return nil, nil
	}
	__obf_5603ee7e7690c5f5, __obf_4d327e1cd40c2e21 := __obf_a21885da2425f2b2.__obf_c4227998727b27c6()
	if __obf_4d327e1cd40c2e21 != nil {
		return nil, __obf_4d327e1cd40c2e21
	}
	__obf_205ee5d820f118f1, __obf_4d327e1cd40c2e21 := __obf_a21885da2425f2b2.__obf_c4227998727b27c6()
	if __obf_4d327e1cd40c2e21 != nil {
		return nil, __obf_4d327e1cd40c2e21
	}
	__obf_c98c06f55b7d5f42 := reflect.TypeOf(__obf_5603ee7e7690c5f5)
	__obf_be822a1292db6343 := reflect.TypeOf(__obf_205ee5d820f118f1)

	if !__obf_c98c06f55b7d5f42.Comparable() {
		return nil, fmt.Errorf("msgpack: unsupported map key: %s", __obf_c98c06f55b7d5f42.String())
	}
	__obf_73234a7ac3b083bf := reflect.MapOf(__obf_c98c06f55b7d5f42, __obf_be822a1292db6343)
	__obf_a3e41621ce2ae57b := __obf_326af9bd942662ac
	if __obf_a21885da2425f2b2.__obf_d617f3769ce16c47&__obf_20f89e0d9adcdf29 == 0 {
		__obf_a3e41621ce2ae57b = min(__obf_a3e41621ce2ae57b, __obf_c02b114384c3720b)
	}
	__obf_c5e1c1e12c411c85 := reflect.MakeMapWithSize(__obf_73234a7ac3b083bf, __obf_a3e41621ce2ae57b)
	__obf_c5e1c1e12c411c85.
		SetMapIndex(reflect.ValueOf(__obf_5603ee7e7690c5f5), reflect.ValueOf(__obf_205ee5d820f118f1))
	__obf_326af9bd942662ac--
	if __obf_4d327e1cd40c2e21 := __obf_a21885da2425f2b2.__obf_e8eb71d4abd07c0e(__obf_c5e1c1e12c411c85, __obf_326af9bd942662ac); __obf_4d327e1cd40c2e21 != nil {
		return nil, __obf_4d327e1cd40c2e21
	}

	return __obf_c5e1c1e12c411c85.Interface(), nil
}

func (__obf_a21885da2425f2b2 *Decoder) __obf_e8eb71d4abd07c0e(__obf_6d570581f4b60dbc reflect.Value, __obf_326af9bd942662ac int) error {
	var (
		__obf_7bd64e8df8357cab = __obf_6d570581f4b60dbc.
					Type()
		__obf_c98c06f55b7d5f42 = __obf_7bd64e8df8357cab.
					Key()
		__obf_be822a1292db6343 = __obf_7bd64e8df8357cab.
					Elem()
	)
	for range __obf_326af9bd942662ac {
		__obf_3b7905b85fce1fb1 := __obf_a21885da2425f2b2.__obf_1fc1190bc702b3d7(__obf_c98c06f55b7d5f42).Elem()
		if __obf_4d327e1cd40c2e21 := __obf_a21885da2425f2b2.DecodeValue(__obf_3b7905b85fce1fb1); __obf_4d327e1cd40c2e21 != nil {
			return __obf_4d327e1cd40c2e21
		}
		__obf_b35b927f0bf61896 := __obf_a21885da2425f2b2.__obf_1fc1190bc702b3d7(__obf_be822a1292db6343).Elem()
		if __obf_4d327e1cd40c2e21 := __obf_a21885da2425f2b2.DecodeValue(__obf_b35b927f0bf61896); __obf_4d327e1cd40c2e21 != nil {
			return __obf_4d327e1cd40c2e21
		}
		__obf_6d570581f4b60dbc.
			SetMapIndex(__obf_3b7905b85fce1fb1, __obf_b35b927f0bf61896)
	}

	return nil
}

func (__obf_a21885da2425f2b2 *Decoder) __obf_ad6b3dbd72d54464(__obf_f5df560f4d67421b byte) error {
	__obf_326af9bd942662ac, __obf_4d327e1cd40c2e21 := __obf_a21885da2425f2b2.__obf_ee4d07636640bafb(__obf_f5df560f4d67421b)
	if __obf_4d327e1cd40c2e21 != nil {
		return __obf_4d327e1cd40c2e21
	}
	for range __obf_326af9bd942662ac {
		if __obf_4d327e1cd40c2e21 := __obf_a21885da2425f2b2.Skip(); __obf_4d327e1cd40c2e21 != nil {
			return __obf_4d327e1cd40c2e21
		}
		if __obf_4d327e1cd40c2e21 := __obf_a21885da2425f2b2.Skip(); __obf_4d327e1cd40c2e21 != nil {
			return __obf_4d327e1cd40c2e21
		}
	}
	return nil
}

func __obf_9651070b33634b16(__obf_a21885da2425f2b2 *Decoder, __obf_6d570581f4b60dbc reflect.Value) error {
	__obf_f5df560f4d67421b, __obf_4d327e1cd40c2e21 := __obf_a21885da2425f2b2.__obf_61e5879fcf7b35ce()
	if __obf_4d327e1cd40c2e21 != nil {
		return __obf_4d327e1cd40c2e21
	}
	__obf_326af9bd942662ac, __obf_4d327e1cd40c2e21 := __obf_a21885da2425f2b2.__obf_ee4d07636640bafb(__obf_f5df560f4d67421b)
	if __obf_4d327e1cd40c2e21 == nil {
		return __obf_a21885da2425f2b2.__obf_6ff623855786a4ae(__obf_6d570581f4b60dbc, __obf_326af9bd942662ac)
	}

	var __obf_48db166600b89531 error
	__obf_326af9bd942662ac, __obf_48db166600b89531 = __obf_a21885da2425f2b2.__obf_2aa510435f9160c6(__obf_f5df560f4d67421b)
	if __obf_48db166600b89531 != nil {
		return __obf_4d327e1cd40c2e21
	}

	if __obf_326af9bd942662ac <= 0 {
		__obf_6d570581f4b60dbc.
			Set(reflect.Zero(__obf_6d570581f4b60dbc.Type()))
		return nil
	}
	__obf_02acefac40101465 := __obf_b7d9de4837c900e0.Fields(__obf_6d570581f4b60dbc.Type(), __obf_a21885da2425f2b2.__obf_90e41d275c699c20)
	if __obf_326af9bd942662ac != len(__obf_02acefac40101465.List) {
		return __obf_33a82116e47fa106
	}

	for _, __obf_f6dd34b5068d19f3 := range __obf_02acefac40101465.List {
		if __obf_4d327e1cd40c2e21 := __obf_f6dd34b5068d19f3.DecodeValue(__obf_a21885da2425f2b2, __obf_6d570581f4b60dbc); __obf_4d327e1cd40c2e21 != nil {
			return __obf_4d327e1cd40c2e21
		}
	}

	return nil
}

func (__obf_a21885da2425f2b2 *Decoder) __obf_6ff623855786a4ae(__obf_6d570581f4b60dbc reflect.Value, __obf_326af9bd942662ac int) error {
	if __obf_326af9bd942662ac == -1 {
		__obf_6d570581f4b60dbc.
			Set(reflect.Zero(__obf_6d570581f4b60dbc.Type()))
		return nil
	}
	__obf_02acefac40101465 := __obf_b7d9de4837c900e0.Fields(__obf_6d570581f4b60dbc.Type(), __obf_a21885da2425f2b2.__obf_90e41d275c699c20)
	for range __obf_326af9bd942662ac {
		__obf_e80b5f2d9b7252c4, __obf_4d327e1cd40c2e21 := __obf_a21885da2425f2b2.__obf_4e3696e2f7b41e6e()
		if __obf_4d327e1cd40c2e21 != nil {
			return __obf_4d327e1cd40c2e21
		}

		if __obf_f6dd34b5068d19f3 := __obf_02acefac40101465.Map[__obf_e80b5f2d9b7252c4]; __obf_f6dd34b5068d19f3 != nil {
			if __obf_4d327e1cd40c2e21 := __obf_f6dd34b5068d19f3.DecodeValue(__obf_a21885da2425f2b2, __obf_6d570581f4b60dbc); __obf_4d327e1cd40c2e21 != nil {
				return __obf_4d327e1cd40c2e21
			}
			continue
		}

		if __obf_a21885da2425f2b2.__obf_d617f3769ce16c47&__obf_5c467062833ca44f != 0 {
			return fmt.Errorf("msgpack: unknown field %q", __obf_e80b5f2d9b7252c4)
		}
		if __obf_4d327e1cd40c2e21 := __obf_a21885da2425f2b2.Skip(); __obf_4d327e1cd40c2e21 != nil {
			return __obf_4d327e1cd40c2e21
		}
	}

	return nil
}
