package __obf_3e0c303610a19bd4

import (
	"errors"
	"fmt"
	"reflect"
)

var __obf_71bf00f1e589d6d0 = errors.New("msgpack: number of fields in array-encoded struct has changed")

var (
	__obf_fe78dd9dd7a09942 = reflect.TypeOf((*map[string]string)(nil))
	__obf_76d48a5285b334b8 = __obf_fe78dd9dd7a09942.
				Elem()
	__obf_8b95b4f2234ae092 = reflect.TypeOf((*map[string]bool)(nil))
	__obf_2f37cd23e2c01e1a = __obf_8b95b4f2234ae092.
				Elem()
)

var (
	__obf_4b0627d9f1049b38 = reflect.TypeOf((*map[string]any)(nil))
	__obf_46f915c67d10ce94 = __obf_4b0627d9f1049b38.
				Elem()
)

func __obf_c02d6c3c3de4ca55(__obf_dc35117108ba8439 *Decoder, __obf_63bbcee86d44fdde reflect.Value) error {
	__obf_4909ae60ffbb8e53, __obf_8882f6cf6e378ded := __obf_dc35117108ba8439.DecodeMapLen()
	if __obf_8882f6cf6e378ded != nil {
		return __obf_8882f6cf6e378ded
	}
	__obf_616f98efc80197c6 := __obf_63bbcee86d44fdde.Type()
	if __obf_4909ae60ffbb8e53 == -1 {
		__obf_63bbcee86d44fdde.
			Set(reflect.Zero(__obf_616f98efc80197c6))
		return nil
	}

	if __obf_63bbcee86d44fdde.IsNil() {
		__obf_eeff06a204a0ff0d := __obf_4909ae60ffbb8e53
		if __obf_dc35117108ba8439.__obf_3cf0882fa5a4cafb&__obf_4b9eaa01d630fce4 == 0 {
			__obf_eeff06a204a0ff0d = min(__obf_eeff06a204a0ff0d, __obf_94b1e1f52a0a7c90)
		}
		__obf_63bbcee86d44fdde.
			Set(reflect.MakeMapWithSize(__obf_616f98efc80197c6, __obf_eeff06a204a0ff0d))
	}
	if __obf_4909ae60ffbb8e53 == 0 {
		return nil
	}

	return __obf_dc35117108ba8439.__obf_170642f58c47fb61(__obf_63bbcee86d44fdde, __obf_4909ae60ffbb8e53)
}

func (__obf_dc35117108ba8439 *Decoder) __obf_a2dd6e1a3453d2c3() (any, error) {
	if __obf_dc35117108ba8439.__obf_7162a9ba642f3ac5 != nil {
		return __obf_dc35117108ba8439.__obf_7162a9ba642f3ac5(__obf_dc35117108ba8439)
	}
	return __obf_dc35117108ba8439.DecodeMap()
}

// DecodeMapLen decodes map length. Length is -1 when map is nil.
func (__obf_dc35117108ba8439 *Decoder) DecodeMapLen() (int, error) {
	__obf_e46289218af709bf, __obf_8882f6cf6e378ded := __obf_dc35117108ba8439.__obf_5b7d85f9093902c3()
	if __obf_8882f6cf6e378ded != nil {
		return 0, __obf_8882f6cf6e378ded
	}

	if IsExt(__obf_e46289218af709bf) {
		if __obf_8882f6cf6e378ded = __obf_dc35117108ba8439.__obf_2ae9d023bf150409(__obf_e46289218af709bf); __obf_8882f6cf6e378ded != nil {
			return 0, __obf_8882f6cf6e378ded
		}
		__obf_e46289218af709bf, __obf_8882f6cf6e378ded = __obf_dc35117108ba8439.__obf_5b7d85f9093902c3()
		if __obf_8882f6cf6e378ded != nil {
			return 0, __obf_8882f6cf6e378ded
		}
	}
	return __obf_dc35117108ba8439.__obf_309efaee64cdf602(__obf_e46289218af709bf)
}

func (__obf_dc35117108ba8439 *Decoder) __obf_309efaee64cdf602(__obf_e46289218af709bf byte) (int, error) {
	if __obf_e46289218af709bf == Nil {
		return -1, nil
	}
	if __obf_e46289218af709bf >= FixedMapLow && __obf_e46289218af709bf <= FixedMapHigh {
		return int(__obf_e46289218af709bf & FixedMapMask), nil
	}
	if __obf_e46289218af709bf == Map16 {
		__obf_7ba4fcf68c1ec623, __obf_8882f6cf6e378ded := __obf_dc35117108ba8439.uint16()
		return int(__obf_7ba4fcf68c1ec623), __obf_8882f6cf6e378ded
	}
	if __obf_e46289218af709bf == Map32 {
		__obf_7ba4fcf68c1ec623, __obf_8882f6cf6e378ded := __obf_dc35117108ba8439.uint32()
		return int(__obf_7ba4fcf68c1ec623), __obf_8882f6cf6e378ded
	}
	return 0, __obf_dc8806ddbc8f0bfa{__obf_545372fefbb733e5: __obf_e46289218af709bf, __obf_41a3520de571ca37: "map length"}
}

func __obf_3a3b656fd33bb5aa(__obf_dc35117108ba8439 *Decoder, __obf_63bbcee86d44fdde reflect.Value) error {
	__obf_9a26563d46df476d := __obf_63bbcee86d44fdde.Addr().Convert(__obf_fe78dd9dd7a09942).Interface().(*map[string]string)
	return __obf_dc35117108ba8439.__obf_58af590e339dfa38(__obf_9a26563d46df476d)
}

func (__obf_dc35117108ba8439 *Decoder) __obf_58af590e339dfa38(__obf_b5a4664807537c0d *map[string]string) error {
	__obf_7ba4fcf68c1ec623, __obf_8882f6cf6e378ded := __obf_dc35117108ba8439.DecodeMapLen()
	if __obf_8882f6cf6e378ded != nil {
		return __obf_8882f6cf6e378ded
	}
	if __obf_7ba4fcf68c1ec623 == -1 {
		*__obf_b5a4664807537c0d = nil
		return nil
	}
	__obf_b416b6a4555be5c8 := *__obf_b5a4664807537c0d
	if __obf_b416b6a4555be5c8 == nil {
		__obf_eeff06a204a0ff0d := __obf_7ba4fcf68c1ec623
		if __obf_dc35117108ba8439.__obf_3cf0882fa5a4cafb&__obf_4b9eaa01d630fce4 == 0 {
			__obf_eeff06a204a0ff0d = min(__obf_7ba4fcf68c1ec623, __obf_94b1e1f52a0a7c90)
		}
		*__obf_b5a4664807537c0d = make(map[string]string, __obf_eeff06a204a0ff0d)
		__obf_b416b6a4555be5c8 = *__obf_b5a4664807537c0d
	}

	for range __obf_7ba4fcf68c1ec623 {
		__obf_e4440ef7e7f689d0, __obf_8882f6cf6e378ded := __obf_dc35117108ba8439.DecodeString()
		if __obf_8882f6cf6e378ded != nil {
			return __obf_8882f6cf6e378ded
		}
		__obf_782ca3783399c281, __obf_8882f6cf6e378ded := __obf_dc35117108ba8439.DecodeString()
		if __obf_8882f6cf6e378ded != nil {
			return __obf_8882f6cf6e378ded
		}
		__obf_b416b6a4555be5c8[__obf_e4440ef7e7f689d0] = __obf_782ca3783399c281
	}

	return nil
}

func __obf_ab027dc2684d6371(__obf_dc35117108ba8439 *Decoder, __obf_63bbcee86d44fdde reflect.Value) error {
	__obf_b5a4664807537c0d := __obf_63bbcee86d44fdde.Addr().Convert(__obf_4b0627d9f1049b38).Interface().(*map[string]any)
	return __obf_dc35117108ba8439.__obf_85426e39f3ed0878(__obf_b5a4664807537c0d)
}

func (__obf_dc35117108ba8439 *Decoder) __obf_85426e39f3ed0878(__obf_b5a4664807537c0d *map[string]any) error {
	__obf_b416b6a4555be5c8, __obf_8882f6cf6e378ded := __obf_dc35117108ba8439.DecodeMap()
	if __obf_8882f6cf6e378ded != nil {
		return __obf_8882f6cf6e378ded
	}
	*__obf_b5a4664807537c0d = __obf_b416b6a4555be5c8
	return nil
}

func (__obf_dc35117108ba8439 *Decoder) DecodeMap() (map[string]any, error) {
	__obf_4909ae60ffbb8e53, __obf_8882f6cf6e378ded := __obf_dc35117108ba8439.DecodeMapLen()
	if __obf_8882f6cf6e378ded != nil {
		return nil, __obf_8882f6cf6e378ded
	}

	if __obf_4909ae60ffbb8e53 == -1 {
		return nil, nil
	}
	__obf_b416b6a4555be5c8 := make(map[string]any, __obf_4909ae60ffbb8e53)

	for range __obf_4909ae60ffbb8e53 {
		__obf_e4440ef7e7f689d0, __obf_8882f6cf6e378ded := __obf_dc35117108ba8439.DecodeString()
		if __obf_8882f6cf6e378ded != nil {
			return nil, __obf_8882f6cf6e378ded
		}
		__obf_782ca3783399c281, __obf_8882f6cf6e378ded := __obf_dc35117108ba8439.__obf_93ff057c8a5c795c()
		if __obf_8882f6cf6e378ded != nil {
			return nil, __obf_8882f6cf6e378ded
		}
		__obf_b416b6a4555be5c8[__obf_e4440ef7e7f689d0] = __obf_782ca3783399c281
	}

	return __obf_b416b6a4555be5c8, nil
}

func (__obf_dc35117108ba8439 *Decoder) DecodeUntypedMap() (map[any]any, error) {
	__obf_4909ae60ffbb8e53, __obf_8882f6cf6e378ded := __obf_dc35117108ba8439.DecodeMapLen()
	if __obf_8882f6cf6e378ded != nil {
		return nil, __obf_8882f6cf6e378ded
	}

	if __obf_4909ae60ffbb8e53 == -1 {
		return nil, nil
	}
	__obf_b416b6a4555be5c8 := make(map[any]any, __obf_4909ae60ffbb8e53)

	for range __obf_4909ae60ffbb8e53 {
		__obf_e4440ef7e7f689d0, __obf_8882f6cf6e378ded := __obf_dc35117108ba8439.__obf_93ff057c8a5c795c()
		if __obf_8882f6cf6e378ded != nil {
			return nil, __obf_8882f6cf6e378ded
		}
		__obf_782ca3783399c281, __obf_8882f6cf6e378ded := __obf_dc35117108ba8439.__obf_93ff057c8a5c795c()
		if __obf_8882f6cf6e378ded != nil {
			return nil, __obf_8882f6cf6e378ded
		}
		__obf_b416b6a4555be5c8[__obf_e4440ef7e7f689d0] = __obf_782ca3783399c281
	}

	return __obf_b416b6a4555be5c8, nil
}

// DecodeTypedMap decodes a typed map. Typed map is a map that has a fixed type for keys and values.
// Key and value types may be different.
func (__obf_dc35117108ba8439 *Decoder) DecodeTypedMap() (any, error) {
	__obf_4909ae60ffbb8e53, __obf_8882f6cf6e378ded := __obf_dc35117108ba8439.DecodeMapLen()
	if __obf_8882f6cf6e378ded != nil {
		return nil, __obf_8882f6cf6e378ded
	}
	if __obf_4909ae60ffbb8e53 <= 0 {
		return nil, nil
	}
	__obf_ef8207c019b4cdb3, __obf_8882f6cf6e378ded := __obf_dc35117108ba8439.__obf_93ff057c8a5c795c()
	if __obf_8882f6cf6e378ded != nil {
		return nil, __obf_8882f6cf6e378ded
	}
	__obf_752fc3666f4041f7, __obf_8882f6cf6e378ded := __obf_dc35117108ba8439.__obf_93ff057c8a5c795c()
	if __obf_8882f6cf6e378ded != nil {
		return nil, __obf_8882f6cf6e378ded
	}
	__obf_404151d80b62a2be := reflect.TypeOf(__obf_ef8207c019b4cdb3)
	__obf_012e089811c37fe0 := reflect.TypeOf(__obf_752fc3666f4041f7)

	if !__obf_404151d80b62a2be.Comparable() {
		return nil, fmt.Errorf("msgpack: unsupported map key: %s", __obf_404151d80b62a2be.String())
	}
	__obf_b159e46ab1f048cd := reflect.MapOf(__obf_404151d80b62a2be, __obf_012e089811c37fe0)
	__obf_eeff06a204a0ff0d := __obf_4909ae60ffbb8e53
	if __obf_dc35117108ba8439.__obf_3cf0882fa5a4cafb&__obf_4b9eaa01d630fce4 == 0 {
		__obf_eeff06a204a0ff0d = min(__obf_eeff06a204a0ff0d, __obf_94b1e1f52a0a7c90)
	}
	__obf_381be812ca6eb2e6 := reflect.MakeMapWithSize(__obf_b159e46ab1f048cd, __obf_eeff06a204a0ff0d)
	__obf_381be812ca6eb2e6.
		SetMapIndex(reflect.ValueOf(__obf_ef8207c019b4cdb3), reflect.ValueOf(__obf_752fc3666f4041f7))
	__obf_4909ae60ffbb8e53--
	if __obf_8882f6cf6e378ded := __obf_dc35117108ba8439.__obf_170642f58c47fb61(__obf_381be812ca6eb2e6, __obf_4909ae60ffbb8e53); __obf_8882f6cf6e378ded != nil {
		return nil, __obf_8882f6cf6e378ded
	}

	return __obf_381be812ca6eb2e6.Interface(), nil
}

func (__obf_dc35117108ba8439 *Decoder) __obf_170642f58c47fb61(__obf_63bbcee86d44fdde reflect.Value, __obf_4909ae60ffbb8e53 int) error {
	var (
		__obf_616f98efc80197c6 = __obf_63bbcee86d44fdde.
					Type()
		__obf_404151d80b62a2be = __obf_616f98efc80197c6.
					Key()
		__obf_012e089811c37fe0 = __obf_616f98efc80197c6.
					Elem()
	)
	for range __obf_4909ae60ffbb8e53 {
		__obf_e4440ef7e7f689d0 := __obf_dc35117108ba8439.__obf_94f521ecbd0b4afa(__obf_404151d80b62a2be).Elem()
		if __obf_8882f6cf6e378ded := __obf_dc35117108ba8439.DecodeValue(__obf_e4440ef7e7f689d0); __obf_8882f6cf6e378ded != nil {
			return __obf_8882f6cf6e378ded
		}
		__obf_782ca3783399c281 := __obf_dc35117108ba8439.__obf_94f521ecbd0b4afa(__obf_012e089811c37fe0).Elem()
		if __obf_8882f6cf6e378ded := __obf_dc35117108ba8439.DecodeValue(__obf_782ca3783399c281); __obf_8882f6cf6e378ded != nil {
			return __obf_8882f6cf6e378ded
		}
		__obf_63bbcee86d44fdde.
			SetMapIndex(__obf_e4440ef7e7f689d0, __obf_782ca3783399c281)
	}

	return nil
}

func (__obf_dc35117108ba8439 *Decoder) __obf_ab578c2c2c260fb6(__obf_e46289218af709bf byte) error {
	__obf_4909ae60ffbb8e53, __obf_8882f6cf6e378ded := __obf_dc35117108ba8439.__obf_309efaee64cdf602(__obf_e46289218af709bf)
	if __obf_8882f6cf6e378ded != nil {
		return __obf_8882f6cf6e378ded
	}
	for range __obf_4909ae60ffbb8e53 {
		if __obf_8882f6cf6e378ded := __obf_dc35117108ba8439.Skip(); __obf_8882f6cf6e378ded != nil {
			return __obf_8882f6cf6e378ded
		}
		if __obf_8882f6cf6e378ded := __obf_dc35117108ba8439.Skip(); __obf_8882f6cf6e378ded != nil {
			return __obf_8882f6cf6e378ded
		}
	}
	return nil
}

func __obf_58fb4ae094054808(__obf_dc35117108ba8439 *Decoder, __obf_63bbcee86d44fdde reflect.Value) error {
	__obf_e46289218af709bf, __obf_8882f6cf6e378ded := __obf_dc35117108ba8439.__obf_5b7d85f9093902c3()
	if __obf_8882f6cf6e378ded != nil {
		return __obf_8882f6cf6e378ded
	}
	__obf_4909ae60ffbb8e53, __obf_8882f6cf6e378ded := __obf_dc35117108ba8439.__obf_309efaee64cdf602(__obf_e46289218af709bf)
	if __obf_8882f6cf6e378ded == nil {
		return __obf_dc35117108ba8439.__obf_d953238596e1307b(__obf_63bbcee86d44fdde, __obf_4909ae60ffbb8e53)
	}

	var __obf_3fe19b8bcd4a2a0b error
	__obf_4909ae60ffbb8e53, __obf_3fe19b8bcd4a2a0b = __obf_dc35117108ba8439.__obf_2a74eb3d00d3189b(__obf_e46289218af709bf)
	if __obf_3fe19b8bcd4a2a0b != nil {
		return __obf_8882f6cf6e378ded
	}

	if __obf_4909ae60ffbb8e53 <= 0 {
		__obf_63bbcee86d44fdde.
			Set(reflect.Zero(__obf_63bbcee86d44fdde.Type()))
		return nil
	}
	__obf_ce19c38e443cefa8 := __obf_3a8fae1624b5849e.Fields(__obf_63bbcee86d44fdde.Type(), __obf_dc35117108ba8439.__obf_6d0a0d8a79287f26)
	if __obf_4909ae60ffbb8e53 != len(__obf_ce19c38e443cefa8.List) {
		return __obf_71bf00f1e589d6d0
	}

	for _, __obf_4add40b0f5dc6b95 := range __obf_ce19c38e443cefa8.List {
		if __obf_8882f6cf6e378ded := __obf_4add40b0f5dc6b95.DecodeValue(__obf_dc35117108ba8439, __obf_63bbcee86d44fdde); __obf_8882f6cf6e378ded != nil {
			return __obf_8882f6cf6e378ded
		}
	}

	return nil
}

func (__obf_dc35117108ba8439 *Decoder) __obf_d953238596e1307b(__obf_63bbcee86d44fdde reflect.Value, __obf_4909ae60ffbb8e53 int) error {
	if __obf_4909ae60ffbb8e53 == -1 {
		__obf_63bbcee86d44fdde.
			Set(reflect.Zero(__obf_63bbcee86d44fdde.Type()))
		return nil
	}
	__obf_ce19c38e443cefa8 := __obf_3a8fae1624b5849e.Fields(__obf_63bbcee86d44fdde.Type(), __obf_dc35117108ba8439.__obf_6d0a0d8a79287f26)
	for range __obf_4909ae60ffbb8e53 {
		__obf_754359097736e0d5, __obf_8882f6cf6e378ded := __obf_dc35117108ba8439.__obf_5484baeee52d4c8a()
		if __obf_8882f6cf6e378ded != nil {
			return __obf_8882f6cf6e378ded
		}

		if __obf_4add40b0f5dc6b95 := __obf_ce19c38e443cefa8.Map[__obf_754359097736e0d5]; __obf_4add40b0f5dc6b95 != nil {
			if __obf_8882f6cf6e378ded := __obf_4add40b0f5dc6b95.DecodeValue(__obf_dc35117108ba8439, __obf_63bbcee86d44fdde); __obf_8882f6cf6e378ded != nil {
				return __obf_8882f6cf6e378ded
			}
			continue
		}

		if __obf_dc35117108ba8439.__obf_3cf0882fa5a4cafb&__obf_8807906d7618037e != 0 {
			return fmt.Errorf("msgpack: unknown field %q", __obf_754359097736e0d5)
		}
		if __obf_8882f6cf6e378ded := __obf_dc35117108ba8439.Skip(); __obf_8882f6cf6e378ded != nil {
			return __obf_8882f6cf6e378ded
		}
	}

	return nil
}
