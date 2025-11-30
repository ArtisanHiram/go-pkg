package __obf_fd770cb9675903b0

import (
	"errors"
	"fmt"
	"reflect"
)

var __obf_86399171bbf77573 = errors.New("msgpack: number of fields in array-encoded struct has changed")

var (
	__obf_bc0036d6832eedeb = reflect.TypeOf((*map[string]string)(nil))
	__obf_a25da01fd293fb59 = __obf_bc0036d6832eedeb.
				Elem()
	__obf_0a48c023045a8303 = reflect.TypeOf((*map[string]bool)(nil))
	__obf_66e0ed467aaa55ff = __obf_0a48c023045a8303.
				Elem()
)

var (
	__obf_4cd48a477958f66c = reflect.TypeOf((*map[string]any)(nil))
	__obf_3c15f58273d2641c = __obf_4cd48a477958f66c.
				Elem()
)

func __obf_274fc8cba465d312(__obf_5d389b660eefb08c *Decoder, __obf_f328a048f76b7256 reflect.Value) error {
	__obf_d774e4844c74bfe9, __obf_45342a3a754d12f5 := __obf_5d389b660eefb08c.DecodeMapLen()
	if __obf_45342a3a754d12f5 != nil {
		return __obf_45342a3a754d12f5
	}
	__obf_8733059f76088ffc := __obf_f328a048f76b7256.Type()
	if __obf_d774e4844c74bfe9 == -1 {
		__obf_f328a048f76b7256.
			Set(reflect.Zero(__obf_8733059f76088ffc))
		return nil
	}

	if __obf_f328a048f76b7256.IsNil() {
		__obf_209aa71f611c3c00 := __obf_d774e4844c74bfe9
		if __obf_5d389b660eefb08c.__obf_7185cb62de7af792&__obf_d0b80752c6a68c61 == 0 {
			__obf_209aa71f611c3c00 = min(__obf_209aa71f611c3c00, __obf_d6d48ab082cf77ea)
		}
		__obf_f328a048f76b7256.
			Set(reflect.MakeMapWithSize(__obf_8733059f76088ffc, __obf_209aa71f611c3c00))
	}
	if __obf_d774e4844c74bfe9 == 0 {
		return nil
	}

	return __obf_5d389b660eefb08c.__obf_a6baa5a2cb5419e2(__obf_f328a048f76b7256, __obf_d774e4844c74bfe9)
}

func (__obf_5d389b660eefb08c *Decoder) __obf_ee0b22a7a221cee0() (any, error) {
	if __obf_5d389b660eefb08c.__obf_a212983b973aa252 != nil {
		return __obf_5d389b660eefb08c.__obf_a212983b973aa252(__obf_5d389b660eefb08c)
	}
	return __obf_5d389b660eefb08c.DecodeMap()
}

// DecodeMapLen decodes map length. Length is -1 when map is nil.
func (__obf_5d389b660eefb08c *Decoder) DecodeMapLen() (int, error) {
	__obf_4148125b350dfea2, __obf_45342a3a754d12f5 := __obf_5d389b660eefb08c.__obf_4786096a37bd32ce()
	if __obf_45342a3a754d12f5 != nil {
		return 0, __obf_45342a3a754d12f5
	}

	if IsExt(__obf_4148125b350dfea2) {
		if __obf_45342a3a754d12f5 = __obf_5d389b660eefb08c.__obf_57d324efa1b98131(__obf_4148125b350dfea2); __obf_45342a3a754d12f5 != nil {
			return 0, __obf_45342a3a754d12f5
		}
		__obf_4148125b350dfea2, __obf_45342a3a754d12f5 = __obf_5d389b660eefb08c.__obf_4786096a37bd32ce()
		if __obf_45342a3a754d12f5 != nil {
			return 0, __obf_45342a3a754d12f5
		}
	}
	return __obf_5d389b660eefb08c.__obf_2ec728407227b54d(__obf_4148125b350dfea2)
}

func (__obf_5d389b660eefb08c *Decoder) __obf_2ec728407227b54d(__obf_4148125b350dfea2 byte) (int, error) {
	if __obf_4148125b350dfea2 == Nil {
		return -1, nil
	}
	if __obf_4148125b350dfea2 >= FixedMapLow && __obf_4148125b350dfea2 <= FixedMapHigh {
		return int(__obf_4148125b350dfea2 & FixedMapMask), nil
	}
	if __obf_4148125b350dfea2 == Map16 {
		__obf_977dae2e8ffd32f8, __obf_45342a3a754d12f5 := __obf_5d389b660eefb08c.uint16()
		return int(__obf_977dae2e8ffd32f8), __obf_45342a3a754d12f5
	}
	if __obf_4148125b350dfea2 == Map32 {
		__obf_977dae2e8ffd32f8, __obf_45342a3a754d12f5 := __obf_5d389b660eefb08c.uint32()
		return int(__obf_977dae2e8ffd32f8), __obf_45342a3a754d12f5
	}
	return 0, __obf_b1f4502463d63ed9{__obf_cde82889ba8e4822: __obf_4148125b350dfea2, __obf_d6c122c727b9e88f: "map length"}
}

func __obf_9a941cbc62083f7f(__obf_5d389b660eefb08c *Decoder, __obf_f328a048f76b7256 reflect.Value) error {
	__obf_58e6da7ea3b944e4 := __obf_f328a048f76b7256.Addr().Convert(__obf_bc0036d6832eedeb).Interface().(*map[string]string)
	return __obf_5d389b660eefb08c.__obf_4acef8852f5f3f0b(__obf_58e6da7ea3b944e4)
}

func (__obf_5d389b660eefb08c *Decoder) __obf_4acef8852f5f3f0b(__obf_2c49f9d2007cfaf6 *map[string]string) error {
	__obf_977dae2e8ffd32f8, __obf_45342a3a754d12f5 := __obf_5d389b660eefb08c.DecodeMapLen()
	if __obf_45342a3a754d12f5 != nil {
		return __obf_45342a3a754d12f5
	}
	if __obf_977dae2e8ffd32f8 == -1 {
		*__obf_2c49f9d2007cfaf6 = nil
		return nil
	}
	__obf_777489ec8e6b2044 := *__obf_2c49f9d2007cfaf6
	if __obf_777489ec8e6b2044 == nil {
		__obf_209aa71f611c3c00 := __obf_977dae2e8ffd32f8
		if __obf_5d389b660eefb08c.__obf_7185cb62de7af792&__obf_d0b80752c6a68c61 == 0 {
			__obf_209aa71f611c3c00 = min(__obf_977dae2e8ffd32f8, __obf_d6d48ab082cf77ea)
		}
		*__obf_2c49f9d2007cfaf6 = make(map[string]string, __obf_209aa71f611c3c00)
		__obf_777489ec8e6b2044 = *__obf_2c49f9d2007cfaf6
	}

	for range __obf_977dae2e8ffd32f8 {
		__obf_a3ae43155694e198, __obf_45342a3a754d12f5 := __obf_5d389b660eefb08c.DecodeString()
		if __obf_45342a3a754d12f5 != nil {
			return __obf_45342a3a754d12f5
		}
		__obf_2cca23d5c936f4ab, __obf_45342a3a754d12f5 := __obf_5d389b660eefb08c.DecodeString()
		if __obf_45342a3a754d12f5 != nil {
			return __obf_45342a3a754d12f5
		}
		__obf_777489ec8e6b2044[__obf_a3ae43155694e198] = __obf_2cca23d5c936f4ab
	}

	return nil
}

func __obf_80486e14ddfacbaf(__obf_5d389b660eefb08c *Decoder, __obf_f328a048f76b7256 reflect.Value) error {
	__obf_2c49f9d2007cfaf6 := __obf_f328a048f76b7256.Addr().Convert(__obf_4cd48a477958f66c).Interface().(*map[string]any)
	return __obf_5d389b660eefb08c.__obf_1bc77a41576c2d43(__obf_2c49f9d2007cfaf6)
}

func (__obf_5d389b660eefb08c *Decoder) __obf_1bc77a41576c2d43(__obf_2c49f9d2007cfaf6 *map[string]any) error {
	__obf_777489ec8e6b2044, __obf_45342a3a754d12f5 := __obf_5d389b660eefb08c.DecodeMap()
	if __obf_45342a3a754d12f5 != nil {
		return __obf_45342a3a754d12f5
	}
	*__obf_2c49f9d2007cfaf6 = __obf_777489ec8e6b2044
	return nil
}

func (__obf_5d389b660eefb08c *Decoder) DecodeMap() (map[string]any, error) {
	__obf_d774e4844c74bfe9, __obf_45342a3a754d12f5 := __obf_5d389b660eefb08c.DecodeMapLen()
	if __obf_45342a3a754d12f5 != nil {
		return nil, __obf_45342a3a754d12f5
	}

	if __obf_d774e4844c74bfe9 == -1 {
		return nil, nil
	}
	__obf_777489ec8e6b2044 := make(map[string]any, __obf_d774e4844c74bfe9)

	for range __obf_d774e4844c74bfe9 {
		__obf_a3ae43155694e198, __obf_45342a3a754d12f5 := __obf_5d389b660eefb08c.DecodeString()
		if __obf_45342a3a754d12f5 != nil {
			return nil, __obf_45342a3a754d12f5
		}
		__obf_2cca23d5c936f4ab, __obf_45342a3a754d12f5 := __obf_5d389b660eefb08c.__obf_d8bd92f8abc7c84f()
		if __obf_45342a3a754d12f5 != nil {
			return nil, __obf_45342a3a754d12f5
		}
		__obf_777489ec8e6b2044[__obf_a3ae43155694e198] = __obf_2cca23d5c936f4ab
	}

	return __obf_777489ec8e6b2044, nil
}

func (__obf_5d389b660eefb08c *Decoder) DecodeUntypedMap() (map[any]any, error) {
	__obf_d774e4844c74bfe9, __obf_45342a3a754d12f5 := __obf_5d389b660eefb08c.DecodeMapLen()
	if __obf_45342a3a754d12f5 != nil {
		return nil, __obf_45342a3a754d12f5
	}

	if __obf_d774e4844c74bfe9 == -1 {
		return nil, nil
	}
	__obf_777489ec8e6b2044 := make(map[any]any, __obf_d774e4844c74bfe9)

	for range __obf_d774e4844c74bfe9 {
		__obf_a3ae43155694e198, __obf_45342a3a754d12f5 := __obf_5d389b660eefb08c.__obf_d8bd92f8abc7c84f()
		if __obf_45342a3a754d12f5 != nil {
			return nil, __obf_45342a3a754d12f5
		}
		__obf_2cca23d5c936f4ab, __obf_45342a3a754d12f5 := __obf_5d389b660eefb08c.__obf_d8bd92f8abc7c84f()
		if __obf_45342a3a754d12f5 != nil {
			return nil, __obf_45342a3a754d12f5
		}
		__obf_777489ec8e6b2044[__obf_a3ae43155694e198] = __obf_2cca23d5c936f4ab
	}

	return __obf_777489ec8e6b2044, nil
}

// DecodeTypedMap decodes a typed map. Typed map is a map that has a fixed type for keys and values.
// Key and value types may be different.
func (__obf_5d389b660eefb08c *Decoder) DecodeTypedMap() (any, error) {
	__obf_d774e4844c74bfe9, __obf_45342a3a754d12f5 := __obf_5d389b660eefb08c.DecodeMapLen()
	if __obf_45342a3a754d12f5 != nil {
		return nil, __obf_45342a3a754d12f5
	}
	if __obf_d774e4844c74bfe9 <= 0 {
		return nil, nil
	}
	__obf_e155b5a95e20e61d, __obf_45342a3a754d12f5 := __obf_5d389b660eefb08c.__obf_d8bd92f8abc7c84f()
	if __obf_45342a3a754d12f5 != nil {
		return nil, __obf_45342a3a754d12f5
	}
	__obf_28cbfc96ff0a56b6, __obf_45342a3a754d12f5 := __obf_5d389b660eefb08c.__obf_d8bd92f8abc7c84f()
	if __obf_45342a3a754d12f5 != nil {
		return nil, __obf_45342a3a754d12f5
	}
	__obf_6e9ffa0694f822e0 := reflect.TypeOf(__obf_e155b5a95e20e61d)
	__obf_ee5b9c8e748434dc := reflect.TypeOf(__obf_28cbfc96ff0a56b6)

	if !__obf_6e9ffa0694f822e0.Comparable() {
		return nil, fmt.Errorf("msgpack: unsupported map key: %s", __obf_6e9ffa0694f822e0.String())
	}
	__obf_ccc5efd6264390d4 := reflect.MapOf(__obf_6e9ffa0694f822e0, __obf_ee5b9c8e748434dc)
	__obf_209aa71f611c3c00 := __obf_d774e4844c74bfe9
	if __obf_5d389b660eefb08c.__obf_7185cb62de7af792&__obf_d0b80752c6a68c61 == 0 {
		__obf_209aa71f611c3c00 = min(__obf_209aa71f611c3c00, __obf_d6d48ab082cf77ea)
	}
	__obf_50a36b9f2cb55f5b := reflect.MakeMapWithSize(__obf_ccc5efd6264390d4, __obf_209aa71f611c3c00)
	__obf_50a36b9f2cb55f5b.
		SetMapIndex(reflect.ValueOf(__obf_e155b5a95e20e61d), reflect.ValueOf(__obf_28cbfc96ff0a56b6))
	__obf_d774e4844c74bfe9--
	if __obf_45342a3a754d12f5 := __obf_5d389b660eefb08c.__obf_a6baa5a2cb5419e2(__obf_50a36b9f2cb55f5b, __obf_d774e4844c74bfe9); __obf_45342a3a754d12f5 != nil {
		return nil, __obf_45342a3a754d12f5
	}

	return __obf_50a36b9f2cb55f5b.Interface(), nil
}

func (__obf_5d389b660eefb08c *Decoder) __obf_a6baa5a2cb5419e2(__obf_f328a048f76b7256 reflect.Value, __obf_d774e4844c74bfe9 int) error {
	var (
		__obf_8733059f76088ffc = __obf_f328a048f76b7256.
					Type()
		__obf_6e9ffa0694f822e0 = __obf_8733059f76088ffc.
					Key()
		__obf_ee5b9c8e748434dc = __obf_8733059f76088ffc.
					Elem()
	)
	for range __obf_d774e4844c74bfe9 {
		__obf_a3ae43155694e198 := __obf_5d389b660eefb08c.__obf_9eb3ea28cca5995b(__obf_6e9ffa0694f822e0).Elem()
		if __obf_45342a3a754d12f5 := __obf_5d389b660eefb08c.DecodeValue(__obf_a3ae43155694e198); __obf_45342a3a754d12f5 != nil {
			return __obf_45342a3a754d12f5
		}
		__obf_2cca23d5c936f4ab := __obf_5d389b660eefb08c.__obf_9eb3ea28cca5995b(__obf_ee5b9c8e748434dc).Elem()
		if __obf_45342a3a754d12f5 := __obf_5d389b660eefb08c.DecodeValue(__obf_2cca23d5c936f4ab); __obf_45342a3a754d12f5 != nil {
			return __obf_45342a3a754d12f5
		}
		__obf_f328a048f76b7256.
			SetMapIndex(__obf_a3ae43155694e198, __obf_2cca23d5c936f4ab)
	}

	return nil
}

func (__obf_5d389b660eefb08c *Decoder) __obf_f977b6e0dc35fc2e(__obf_4148125b350dfea2 byte) error {
	__obf_d774e4844c74bfe9, __obf_45342a3a754d12f5 := __obf_5d389b660eefb08c.__obf_2ec728407227b54d(__obf_4148125b350dfea2)
	if __obf_45342a3a754d12f5 != nil {
		return __obf_45342a3a754d12f5
	}
	for range __obf_d774e4844c74bfe9 {
		if __obf_45342a3a754d12f5 := __obf_5d389b660eefb08c.Skip(); __obf_45342a3a754d12f5 != nil {
			return __obf_45342a3a754d12f5
		}
		if __obf_45342a3a754d12f5 := __obf_5d389b660eefb08c.Skip(); __obf_45342a3a754d12f5 != nil {
			return __obf_45342a3a754d12f5
		}
	}
	return nil
}

func __obf_9fde65ded02d475f(__obf_5d389b660eefb08c *Decoder, __obf_f328a048f76b7256 reflect.Value) error {
	__obf_4148125b350dfea2, __obf_45342a3a754d12f5 := __obf_5d389b660eefb08c.__obf_4786096a37bd32ce()
	if __obf_45342a3a754d12f5 != nil {
		return __obf_45342a3a754d12f5
	}
	__obf_d774e4844c74bfe9, __obf_45342a3a754d12f5 := __obf_5d389b660eefb08c.__obf_2ec728407227b54d(__obf_4148125b350dfea2)
	if __obf_45342a3a754d12f5 == nil {
		return __obf_5d389b660eefb08c.__obf_65f40b078c0a6bdf(__obf_f328a048f76b7256, __obf_d774e4844c74bfe9)
	}

	var __obf_30dddc44fc6b7c5f error
	__obf_d774e4844c74bfe9, __obf_30dddc44fc6b7c5f = __obf_5d389b660eefb08c.__obf_a07a0c18e5d282be(__obf_4148125b350dfea2)
	if __obf_30dddc44fc6b7c5f != nil {
		return __obf_45342a3a754d12f5
	}

	if __obf_d774e4844c74bfe9 <= 0 {
		__obf_f328a048f76b7256.
			Set(reflect.Zero(__obf_f328a048f76b7256.Type()))
		return nil
	}
	__obf_f92dfca778102077 := __obf_a93e575d0cc03783.Fields(__obf_f328a048f76b7256.Type(), __obf_5d389b660eefb08c.__obf_6621ba3773b8696a)
	if __obf_d774e4844c74bfe9 != len(__obf_f92dfca778102077.List) {
		return __obf_86399171bbf77573
	}

	for _, __obf_e25b9e550ea05af3 := range __obf_f92dfca778102077.List {
		if __obf_45342a3a754d12f5 := __obf_e25b9e550ea05af3.DecodeValue(__obf_5d389b660eefb08c, __obf_f328a048f76b7256); __obf_45342a3a754d12f5 != nil {
			return __obf_45342a3a754d12f5
		}
	}

	return nil
}

func (__obf_5d389b660eefb08c *Decoder) __obf_65f40b078c0a6bdf(__obf_f328a048f76b7256 reflect.Value, __obf_d774e4844c74bfe9 int) error {
	if __obf_d774e4844c74bfe9 == -1 {
		__obf_f328a048f76b7256.
			Set(reflect.Zero(__obf_f328a048f76b7256.Type()))
		return nil
	}
	__obf_f92dfca778102077 := __obf_a93e575d0cc03783.Fields(__obf_f328a048f76b7256.Type(), __obf_5d389b660eefb08c.__obf_6621ba3773b8696a)
	for range __obf_d774e4844c74bfe9 {
		__obf_2a5a18f2bc2906b2, __obf_45342a3a754d12f5 := __obf_5d389b660eefb08c.__obf_479ae9aab7b05d68()
		if __obf_45342a3a754d12f5 != nil {
			return __obf_45342a3a754d12f5
		}

		if __obf_e25b9e550ea05af3 := __obf_f92dfca778102077.Map[__obf_2a5a18f2bc2906b2]; __obf_e25b9e550ea05af3 != nil {
			if __obf_45342a3a754d12f5 := __obf_e25b9e550ea05af3.DecodeValue(__obf_5d389b660eefb08c, __obf_f328a048f76b7256); __obf_45342a3a754d12f5 != nil {
				return __obf_45342a3a754d12f5
			}
			continue
		}

		if __obf_5d389b660eefb08c.__obf_7185cb62de7af792&__obf_5726d2683a07c9c6 != 0 {
			return fmt.Errorf("msgpack: unknown field %q", __obf_2a5a18f2bc2906b2)
		}
		if __obf_45342a3a754d12f5 := __obf_5d389b660eefb08c.Skip(); __obf_45342a3a754d12f5 != nil {
			return __obf_45342a3a754d12f5
		}
	}

	return nil
}
