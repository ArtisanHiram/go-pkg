package __obf_a4aad98aaf3178f4

import (
	"errors"
	"fmt"
	"reflect"
)

var __obf_7b6975bcf1390025 = errors.New("msgpack: number of fields in array-encoded struct has changed")

var (
	__obf_69759c415a929047 = reflect.TypeOf((*map[string]string)(nil))
	__obf_23cc385d7847ade3 = __obf_69759c415a929047.
				Elem()
	__obf_24b0e5a73747ecf6 = reflect.TypeOf((*map[string]bool)(nil))
	__obf_6742791176d1c4ba = __obf_24b0e5a73747ecf6.
				Elem()
)

var (
	__obf_8b7de4e099520c07 = reflect.TypeOf((*map[string]any)(nil))
	__obf_ef257d4bd242e017 = __obf_8b7de4e099520c07.
				Elem()
)

func __obf_c97ee5b3b53fb2c5(__obf_613397fefdec0ed0 *Decoder, __obf_a1f43267eeb48a1a reflect.Value) error {
	__obf_99a74e41c9c640c0, __obf_4061ca0039150c39 := __obf_613397fefdec0ed0.DecodeMapLen()
	if __obf_4061ca0039150c39 != nil {
		return __obf_4061ca0039150c39
	}
	__obf_bbfd30fcc08dc1bc := __obf_a1f43267eeb48a1a.Type()
	if __obf_99a74e41c9c640c0 == -1 {
		__obf_a1f43267eeb48a1a.
			Set(reflect.Zero(__obf_bbfd30fcc08dc1bc))
		return nil
	}

	if __obf_a1f43267eeb48a1a.IsNil() {
		__obf_2a56557d09674668 := __obf_99a74e41c9c640c0
		if __obf_613397fefdec0ed0.__obf_a8400e70a712e9fa&__obf_beecb71ff037aa1b == 0 {
			__obf_2a56557d09674668 = min(__obf_2a56557d09674668, __obf_0dcaecf62133ed63)
		}
		__obf_a1f43267eeb48a1a.
			Set(reflect.MakeMapWithSize(__obf_bbfd30fcc08dc1bc, __obf_2a56557d09674668))
	}
	if __obf_99a74e41c9c640c0 == 0 {
		return nil
	}

	return __obf_613397fefdec0ed0.__obf_fcd041d31e24cfc5(__obf_a1f43267eeb48a1a, __obf_99a74e41c9c640c0)
}

func (__obf_613397fefdec0ed0 *Decoder) __obf_01c9a38c34ed0945() (any, error) {
	if __obf_613397fefdec0ed0.__obf_d20effb6daf18186 != nil {
		return __obf_613397fefdec0ed0.__obf_d20effb6daf18186(__obf_613397fefdec0ed0)
	}
	return __obf_613397fefdec0ed0.DecodeMap()
}

// DecodeMapLen decodes map length. Length is -1 when map is nil.
func (__obf_613397fefdec0ed0 *Decoder) DecodeMapLen() (int, error) {
	__obf_6dbe86b3d9d9b037, __obf_4061ca0039150c39 := __obf_613397fefdec0ed0.__obf_46b95b57ed5617b5()
	if __obf_4061ca0039150c39 != nil {
		return 0, __obf_4061ca0039150c39
	}

	if IsExt(__obf_6dbe86b3d9d9b037) {
		if __obf_4061ca0039150c39 = __obf_613397fefdec0ed0.__obf_0428b2f11094e5f5(__obf_6dbe86b3d9d9b037); __obf_4061ca0039150c39 != nil {
			return 0, __obf_4061ca0039150c39
		}
		__obf_6dbe86b3d9d9b037, __obf_4061ca0039150c39 = __obf_613397fefdec0ed0.__obf_46b95b57ed5617b5()
		if __obf_4061ca0039150c39 != nil {
			return 0, __obf_4061ca0039150c39
		}
	}
	return __obf_613397fefdec0ed0.__obf_3f2de7171e23c38e(__obf_6dbe86b3d9d9b037)
}

func (__obf_613397fefdec0ed0 *Decoder) __obf_3f2de7171e23c38e(__obf_6dbe86b3d9d9b037 byte) (int, error) {
	if __obf_6dbe86b3d9d9b037 == Nil {
		return -1, nil
	}
	if __obf_6dbe86b3d9d9b037 >= FixedMapLow && __obf_6dbe86b3d9d9b037 <= FixedMapHigh {
		return int(__obf_6dbe86b3d9d9b037 & FixedMapMask), nil
	}
	if __obf_6dbe86b3d9d9b037 == Map16 {
		__obf_73d124c794e3a6a2, __obf_4061ca0039150c39 := __obf_613397fefdec0ed0.uint16()
		return int(__obf_73d124c794e3a6a2), __obf_4061ca0039150c39
	}
	if __obf_6dbe86b3d9d9b037 == Map32 {
		__obf_73d124c794e3a6a2, __obf_4061ca0039150c39 := __obf_613397fefdec0ed0.uint32()
		return int(__obf_73d124c794e3a6a2), __obf_4061ca0039150c39
	}
	return 0, __obf_31d9eceb380e57ed{__obf_987b95dd4dcfc308: __obf_6dbe86b3d9d9b037, __obf_1ec1fd8709fc0afd: "map length"}
}

func __obf_cc88c63d6bb28c3c(__obf_613397fefdec0ed0 *Decoder, __obf_a1f43267eeb48a1a reflect.Value) error {
	__obf_8f294218dfea15c7 := __obf_a1f43267eeb48a1a.Addr().Convert(__obf_69759c415a929047).Interface().(*map[string]string)
	return __obf_613397fefdec0ed0.__obf_b01b056f781c9c97(__obf_8f294218dfea15c7)
}

func (__obf_613397fefdec0ed0 *Decoder) __obf_b01b056f781c9c97(__obf_cf3077802722ecc2 *map[string]string) error {
	__obf_73d124c794e3a6a2, __obf_4061ca0039150c39 := __obf_613397fefdec0ed0.DecodeMapLen()
	if __obf_4061ca0039150c39 != nil {
		return __obf_4061ca0039150c39
	}
	if __obf_73d124c794e3a6a2 == -1 {
		*__obf_cf3077802722ecc2 = nil
		return nil
	}
	__obf_66cc4b26e3c9cdbb := *__obf_cf3077802722ecc2
	if __obf_66cc4b26e3c9cdbb == nil {
		__obf_2a56557d09674668 := __obf_73d124c794e3a6a2
		if __obf_613397fefdec0ed0.__obf_a8400e70a712e9fa&__obf_beecb71ff037aa1b == 0 {
			__obf_2a56557d09674668 = min(__obf_73d124c794e3a6a2, __obf_0dcaecf62133ed63)
		}
		*__obf_cf3077802722ecc2 = make(map[string]string, __obf_2a56557d09674668)
		__obf_66cc4b26e3c9cdbb = *__obf_cf3077802722ecc2
	}

	for range __obf_73d124c794e3a6a2 {
		__obf_25e3964b56bdb154, __obf_4061ca0039150c39 := __obf_613397fefdec0ed0.DecodeString()
		if __obf_4061ca0039150c39 != nil {
			return __obf_4061ca0039150c39
		}
		__obf_e23729dae69c9f0b, __obf_4061ca0039150c39 := __obf_613397fefdec0ed0.DecodeString()
		if __obf_4061ca0039150c39 != nil {
			return __obf_4061ca0039150c39
		}
		__obf_66cc4b26e3c9cdbb[__obf_25e3964b56bdb154] = __obf_e23729dae69c9f0b
	}

	return nil
}

func __obf_a5e827593d54b583(__obf_613397fefdec0ed0 *Decoder, __obf_a1f43267eeb48a1a reflect.Value) error {
	__obf_cf3077802722ecc2 := __obf_a1f43267eeb48a1a.Addr().Convert(__obf_8b7de4e099520c07).Interface().(*map[string]any)
	return __obf_613397fefdec0ed0.__obf_27d72527258884d5(__obf_cf3077802722ecc2)
}

func (__obf_613397fefdec0ed0 *Decoder) __obf_27d72527258884d5(__obf_cf3077802722ecc2 *map[string]any) error {
	__obf_66cc4b26e3c9cdbb, __obf_4061ca0039150c39 := __obf_613397fefdec0ed0.DecodeMap()
	if __obf_4061ca0039150c39 != nil {
		return __obf_4061ca0039150c39
	}
	*__obf_cf3077802722ecc2 = __obf_66cc4b26e3c9cdbb
	return nil
}

func (__obf_613397fefdec0ed0 *Decoder) DecodeMap() (map[string]any, error) {
	__obf_99a74e41c9c640c0, __obf_4061ca0039150c39 := __obf_613397fefdec0ed0.DecodeMapLen()
	if __obf_4061ca0039150c39 != nil {
		return nil, __obf_4061ca0039150c39
	}

	if __obf_99a74e41c9c640c0 == -1 {
		return nil, nil
	}
	__obf_66cc4b26e3c9cdbb := make(map[string]any, __obf_99a74e41c9c640c0)

	for range __obf_99a74e41c9c640c0 {
		__obf_25e3964b56bdb154, __obf_4061ca0039150c39 := __obf_613397fefdec0ed0.DecodeString()
		if __obf_4061ca0039150c39 != nil {
			return nil, __obf_4061ca0039150c39
		}
		__obf_e23729dae69c9f0b, __obf_4061ca0039150c39 := __obf_613397fefdec0ed0.__obf_ec86210fb048ddc6()
		if __obf_4061ca0039150c39 != nil {
			return nil, __obf_4061ca0039150c39
		}
		__obf_66cc4b26e3c9cdbb[__obf_25e3964b56bdb154] = __obf_e23729dae69c9f0b
	}

	return __obf_66cc4b26e3c9cdbb, nil
}

func (__obf_613397fefdec0ed0 *Decoder) DecodeUntypedMap() (map[any]any, error) {
	__obf_99a74e41c9c640c0, __obf_4061ca0039150c39 := __obf_613397fefdec0ed0.DecodeMapLen()
	if __obf_4061ca0039150c39 != nil {
		return nil, __obf_4061ca0039150c39
	}

	if __obf_99a74e41c9c640c0 == -1 {
		return nil, nil
	}
	__obf_66cc4b26e3c9cdbb := make(map[any]any, __obf_99a74e41c9c640c0)

	for range __obf_99a74e41c9c640c0 {
		__obf_25e3964b56bdb154, __obf_4061ca0039150c39 := __obf_613397fefdec0ed0.__obf_ec86210fb048ddc6()
		if __obf_4061ca0039150c39 != nil {
			return nil, __obf_4061ca0039150c39
		}
		__obf_e23729dae69c9f0b, __obf_4061ca0039150c39 := __obf_613397fefdec0ed0.__obf_ec86210fb048ddc6()
		if __obf_4061ca0039150c39 != nil {
			return nil, __obf_4061ca0039150c39
		}
		__obf_66cc4b26e3c9cdbb[__obf_25e3964b56bdb154] = __obf_e23729dae69c9f0b
	}

	return __obf_66cc4b26e3c9cdbb, nil
}

// DecodeTypedMap decodes a typed map. Typed map is a map that has a fixed type for keys and values.
// Key and value types may be different.
func (__obf_613397fefdec0ed0 *Decoder) DecodeTypedMap() (any, error) {
	__obf_99a74e41c9c640c0, __obf_4061ca0039150c39 := __obf_613397fefdec0ed0.DecodeMapLen()
	if __obf_4061ca0039150c39 != nil {
		return nil, __obf_4061ca0039150c39
	}
	if __obf_99a74e41c9c640c0 <= 0 {
		return nil, nil
	}
	__obf_992cb53d7b9cb024, __obf_4061ca0039150c39 := __obf_613397fefdec0ed0.__obf_ec86210fb048ddc6()
	if __obf_4061ca0039150c39 != nil {
		return nil, __obf_4061ca0039150c39
	}
	__obf_9055e1481edcd436, __obf_4061ca0039150c39 := __obf_613397fefdec0ed0.__obf_ec86210fb048ddc6()
	if __obf_4061ca0039150c39 != nil {
		return nil, __obf_4061ca0039150c39
	}
	__obf_8ed73698802e3906 := reflect.TypeOf(__obf_992cb53d7b9cb024)
	__obf_fea988b4f5ef4c3c := reflect.TypeOf(__obf_9055e1481edcd436)

	if !__obf_8ed73698802e3906.Comparable() {
		return nil, fmt.Errorf("msgpack: unsupported map key: %s", __obf_8ed73698802e3906.String())
	}
	__obf_560c892fccb23944 := reflect.MapOf(__obf_8ed73698802e3906, __obf_fea988b4f5ef4c3c)
	__obf_2a56557d09674668 := __obf_99a74e41c9c640c0
	if __obf_613397fefdec0ed0.__obf_a8400e70a712e9fa&__obf_beecb71ff037aa1b == 0 {
		__obf_2a56557d09674668 = min(__obf_2a56557d09674668, __obf_0dcaecf62133ed63)
	}
	__obf_2c6058099d421374 := reflect.MakeMapWithSize(__obf_560c892fccb23944, __obf_2a56557d09674668)
	__obf_2c6058099d421374.
		SetMapIndex(reflect.ValueOf(__obf_992cb53d7b9cb024), reflect.ValueOf(__obf_9055e1481edcd436))
	__obf_99a74e41c9c640c0--
	if __obf_4061ca0039150c39 := __obf_613397fefdec0ed0.__obf_fcd041d31e24cfc5(__obf_2c6058099d421374, __obf_99a74e41c9c640c0); __obf_4061ca0039150c39 != nil {
		return nil, __obf_4061ca0039150c39
	}

	return __obf_2c6058099d421374.Interface(), nil
}

func (__obf_613397fefdec0ed0 *Decoder) __obf_fcd041d31e24cfc5(__obf_a1f43267eeb48a1a reflect.Value, __obf_99a74e41c9c640c0 int) error {
	var (
		__obf_bbfd30fcc08dc1bc = __obf_a1f43267eeb48a1a.
					Type()
		__obf_8ed73698802e3906 = __obf_bbfd30fcc08dc1bc.
					Key()
		__obf_fea988b4f5ef4c3c = __obf_bbfd30fcc08dc1bc.
					Elem()
	)
	for range __obf_99a74e41c9c640c0 {
		__obf_25e3964b56bdb154 := __obf_613397fefdec0ed0.__obf_7a63f73d6b3d6972(__obf_8ed73698802e3906).Elem()
		if __obf_4061ca0039150c39 := __obf_613397fefdec0ed0.DecodeValue(__obf_25e3964b56bdb154); __obf_4061ca0039150c39 != nil {
			return __obf_4061ca0039150c39
		}
		__obf_e23729dae69c9f0b := __obf_613397fefdec0ed0.__obf_7a63f73d6b3d6972(__obf_fea988b4f5ef4c3c).Elem()
		if __obf_4061ca0039150c39 := __obf_613397fefdec0ed0.DecodeValue(__obf_e23729dae69c9f0b); __obf_4061ca0039150c39 != nil {
			return __obf_4061ca0039150c39
		}
		__obf_a1f43267eeb48a1a.
			SetMapIndex(__obf_25e3964b56bdb154, __obf_e23729dae69c9f0b)
	}

	return nil
}

func (__obf_613397fefdec0ed0 *Decoder) __obf_a25930d7ee975bcc(__obf_6dbe86b3d9d9b037 byte) error {
	__obf_99a74e41c9c640c0, __obf_4061ca0039150c39 := __obf_613397fefdec0ed0.__obf_3f2de7171e23c38e(__obf_6dbe86b3d9d9b037)
	if __obf_4061ca0039150c39 != nil {
		return __obf_4061ca0039150c39
	}
	for range __obf_99a74e41c9c640c0 {
		if __obf_4061ca0039150c39 := __obf_613397fefdec0ed0.Skip(); __obf_4061ca0039150c39 != nil {
			return __obf_4061ca0039150c39
		}
		if __obf_4061ca0039150c39 := __obf_613397fefdec0ed0.Skip(); __obf_4061ca0039150c39 != nil {
			return __obf_4061ca0039150c39
		}
	}
	return nil
}

func __obf_cf4307adf926d203(__obf_613397fefdec0ed0 *Decoder, __obf_a1f43267eeb48a1a reflect.Value) error {
	__obf_6dbe86b3d9d9b037, __obf_4061ca0039150c39 := __obf_613397fefdec0ed0.__obf_46b95b57ed5617b5()
	if __obf_4061ca0039150c39 != nil {
		return __obf_4061ca0039150c39
	}
	__obf_99a74e41c9c640c0, __obf_4061ca0039150c39 := __obf_613397fefdec0ed0.__obf_3f2de7171e23c38e(__obf_6dbe86b3d9d9b037)
	if __obf_4061ca0039150c39 == nil {
		return __obf_613397fefdec0ed0.__obf_b333f2a241a3898e(__obf_a1f43267eeb48a1a, __obf_99a74e41c9c640c0)
	}

	var __obf_3007455580ef212f error
	__obf_99a74e41c9c640c0, __obf_3007455580ef212f = __obf_613397fefdec0ed0.__obf_0b139571f0c923ec(__obf_6dbe86b3d9d9b037)
	if __obf_3007455580ef212f != nil {
		return __obf_4061ca0039150c39
	}

	if __obf_99a74e41c9c640c0 <= 0 {
		__obf_a1f43267eeb48a1a.
			Set(reflect.Zero(__obf_a1f43267eeb48a1a.Type()))
		return nil
	}
	__obf_0cf794214f02df4d := __obf_bd827009b4ede0e2.Fields(__obf_a1f43267eeb48a1a.Type(), __obf_613397fefdec0ed0.__obf_72ba705ed504d210)
	if __obf_99a74e41c9c640c0 != len(__obf_0cf794214f02df4d.List) {
		return __obf_7b6975bcf1390025
	}

	for _, __obf_da4a2296298d318f := range __obf_0cf794214f02df4d.List {
		if __obf_4061ca0039150c39 := __obf_da4a2296298d318f.DecodeValue(__obf_613397fefdec0ed0, __obf_a1f43267eeb48a1a); __obf_4061ca0039150c39 != nil {
			return __obf_4061ca0039150c39
		}
	}

	return nil
}

func (__obf_613397fefdec0ed0 *Decoder) __obf_b333f2a241a3898e(__obf_a1f43267eeb48a1a reflect.Value, __obf_99a74e41c9c640c0 int) error {
	if __obf_99a74e41c9c640c0 == -1 {
		__obf_a1f43267eeb48a1a.
			Set(reflect.Zero(__obf_a1f43267eeb48a1a.Type()))
		return nil
	}
	__obf_0cf794214f02df4d := __obf_bd827009b4ede0e2.Fields(__obf_a1f43267eeb48a1a.Type(), __obf_613397fefdec0ed0.__obf_72ba705ed504d210)
	for range __obf_99a74e41c9c640c0 {
		__obf_071b55c16b16fe35, __obf_4061ca0039150c39 := __obf_613397fefdec0ed0.__obf_dc853c0361533caa()
		if __obf_4061ca0039150c39 != nil {
			return __obf_4061ca0039150c39
		}

		if __obf_da4a2296298d318f := __obf_0cf794214f02df4d.Map[__obf_071b55c16b16fe35]; __obf_da4a2296298d318f != nil {
			if __obf_4061ca0039150c39 := __obf_da4a2296298d318f.DecodeValue(__obf_613397fefdec0ed0, __obf_a1f43267eeb48a1a); __obf_4061ca0039150c39 != nil {
				return __obf_4061ca0039150c39
			}
			continue
		}

		if __obf_613397fefdec0ed0.__obf_a8400e70a712e9fa&__obf_00d6d684a301507c != 0 {
			return fmt.Errorf("msgpack: unknown field %q", __obf_071b55c16b16fe35)
		}
		if __obf_4061ca0039150c39 := __obf_613397fefdec0ed0.Skip(); __obf_4061ca0039150c39 != nil {
			return __obf_4061ca0039150c39
		}
	}

	return nil
}
