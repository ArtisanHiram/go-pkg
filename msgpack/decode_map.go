package __obf_3edaa49e853afa16

import (
	"errors"
	"fmt"
	"reflect"
)

var __obf_607b50c18f29ff3f = errors.New("msgpack: number of fields in array-encoded struct has changed")

var (
	__obf_4f4bcf58d6ad6cea = reflect.TypeOf((*map[string]string)(nil))
	__obf_3e2ae0e33841c7b2 = __obf_4f4bcf58d6ad6cea.
				Elem()
	__obf_c721376e99d352d0 = reflect.TypeOf((*map[string]bool)(nil))
	__obf_faa0fb9daafd869d = __obf_c721376e99d352d0.
				Elem()
)

var (
	__obf_4b12002d82e4bb80 = reflect.TypeOf((*map[string]any)(nil))
	__obf_739f33e1b9d59004 = __obf_4b12002d82e4bb80.
				Elem()
)

func __obf_943cb3215d99d848(__obf_015afbee33862a01 *Decoder, __obf_85d270343a0dfe14 reflect.Value) error {
	__obf_56127cd370854f0b, __obf_3aa78ad28f50ed46 := __obf_015afbee33862a01.DecodeMapLen()
	if __obf_3aa78ad28f50ed46 != nil {
		return __obf_3aa78ad28f50ed46
	}
	__obf_af250e9784b6a92c := __obf_85d270343a0dfe14.Type()
	if __obf_56127cd370854f0b == -1 {
		__obf_85d270343a0dfe14.
			Set(reflect.Zero(__obf_af250e9784b6a92c))
		return nil
	}

	if __obf_85d270343a0dfe14.IsNil() {
		__obf_a3b64f93c4732565 := __obf_56127cd370854f0b
		if __obf_015afbee33862a01.__obf_6aa0efd537415192&__obf_9d8a15eb3a073456 == 0 {
			__obf_a3b64f93c4732565 = min(__obf_a3b64f93c4732565, __obf_5d6aaa093307f6f6)
		}
		__obf_85d270343a0dfe14.
			Set(reflect.MakeMapWithSize(__obf_af250e9784b6a92c, __obf_a3b64f93c4732565))
	}
	if __obf_56127cd370854f0b == 0 {
		return nil
	}

	return __obf_015afbee33862a01.__obf_1f0ded3b3d232ea2(__obf_85d270343a0dfe14, __obf_56127cd370854f0b)
}

func (__obf_015afbee33862a01 *Decoder) __obf_987a7825732f4656() (any, error) {
	if __obf_015afbee33862a01.__obf_e63e8a485c325197 != nil {
		return __obf_015afbee33862a01.__obf_e63e8a485c325197(__obf_015afbee33862a01)
	}
	return __obf_015afbee33862a01.DecodeMap()
}

// DecodeMapLen decodes map length. Length is -1 when map is nil.
func (__obf_015afbee33862a01 *Decoder) DecodeMapLen() (int, error) {
	__obf_145c7a7d25eea2bd, __obf_3aa78ad28f50ed46 := __obf_015afbee33862a01.__obf_4c5f8d154c9027ea()
	if __obf_3aa78ad28f50ed46 != nil {
		return 0, __obf_3aa78ad28f50ed46
	}

	if IsExt(__obf_145c7a7d25eea2bd) {
		if __obf_3aa78ad28f50ed46 = __obf_015afbee33862a01.__obf_c7f8d446fe3a2767(__obf_145c7a7d25eea2bd); __obf_3aa78ad28f50ed46 != nil {
			return 0, __obf_3aa78ad28f50ed46
		}
		__obf_145c7a7d25eea2bd, __obf_3aa78ad28f50ed46 = __obf_015afbee33862a01.__obf_4c5f8d154c9027ea()
		if __obf_3aa78ad28f50ed46 != nil {
			return 0, __obf_3aa78ad28f50ed46
		}
	}
	return __obf_015afbee33862a01.__obf_9cb09f4e20a89b72(__obf_145c7a7d25eea2bd)
}

func (__obf_015afbee33862a01 *Decoder) __obf_9cb09f4e20a89b72(__obf_145c7a7d25eea2bd byte) (int, error) {
	if __obf_145c7a7d25eea2bd == Nil {
		return -1, nil
	}
	if __obf_145c7a7d25eea2bd >= FixedMapLow && __obf_145c7a7d25eea2bd <= FixedMapHigh {
		return int(__obf_145c7a7d25eea2bd & FixedMapMask), nil
	}
	if __obf_145c7a7d25eea2bd == Map16 {
		__obf_5cbee6b126b9acb5, __obf_3aa78ad28f50ed46 := __obf_015afbee33862a01.uint16()
		return int(__obf_5cbee6b126b9acb5), __obf_3aa78ad28f50ed46
	}
	if __obf_145c7a7d25eea2bd == Map32 {
		__obf_5cbee6b126b9acb5, __obf_3aa78ad28f50ed46 := __obf_015afbee33862a01.uint32()
		return int(__obf_5cbee6b126b9acb5), __obf_3aa78ad28f50ed46
	}
	return 0, __obf_6b688bac335393d7{__obf_508e2d6ff53655c0: __obf_145c7a7d25eea2bd, __obf_6e0dac6e0c9d5f66: "map length"}
}

func __obf_107f1a504832f4ec(__obf_015afbee33862a01 *Decoder, __obf_85d270343a0dfe14 reflect.Value) error {
	__obf_b56b4968c8cc4763 := __obf_85d270343a0dfe14.Addr().Convert(__obf_4f4bcf58d6ad6cea).Interface().(*map[string]string)
	return __obf_015afbee33862a01.__obf_7e971a691672ab09(__obf_b56b4968c8cc4763)
}

func (__obf_015afbee33862a01 *Decoder) __obf_7e971a691672ab09(__obf_8f0c71619c0382f1 *map[string]string) error {
	__obf_5cbee6b126b9acb5, __obf_3aa78ad28f50ed46 := __obf_015afbee33862a01.DecodeMapLen()
	if __obf_3aa78ad28f50ed46 != nil {
		return __obf_3aa78ad28f50ed46
	}
	if __obf_5cbee6b126b9acb5 == -1 {
		*__obf_8f0c71619c0382f1 = nil
		return nil
	}
	__obf_c26f5adfb3152475 := *__obf_8f0c71619c0382f1
	if __obf_c26f5adfb3152475 == nil {
		__obf_a3b64f93c4732565 := __obf_5cbee6b126b9acb5
		if __obf_015afbee33862a01.__obf_6aa0efd537415192&__obf_9d8a15eb3a073456 == 0 {
			__obf_a3b64f93c4732565 = min(__obf_5cbee6b126b9acb5, __obf_5d6aaa093307f6f6)
		}
		*__obf_8f0c71619c0382f1 = make(map[string]string, __obf_a3b64f93c4732565)
		__obf_c26f5adfb3152475 = *__obf_8f0c71619c0382f1
	}

	for range __obf_5cbee6b126b9acb5 {
		__obf_8f7121221504784f, __obf_3aa78ad28f50ed46 := __obf_015afbee33862a01.DecodeString()
		if __obf_3aa78ad28f50ed46 != nil {
			return __obf_3aa78ad28f50ed46
		}
		__obf_a03d608795b584a3, __obf_3aa78ad28f50ed46 := __obf_015afbee33862a01.DecodeString()
		if __obf_3aa78ad28f50ed46 != nil {
			return __obf_3aa78ad28f50ed46
		}
		__obf_c26f5adfb3152475[__obf_8f7121221504784f] = __obf_a03d608795b584a3
	}

	return nil
}

func __obf_aa82e0459275f272(__obf_015afbee33862a01 *Decoder, __obf_85d270343a0dfe14 reflect.Value) error {
	__obf_8f0c71619c0382f1 := __obf_85d270343a0dfe14.Addr().Convert(__obf_4b12002d82e4bb80).Interface().(*map[string]any)
	return __obf_015afbee33862a01.__obf_c5ba7e7660e774e8(__obf_8f0c71619c0382f1)
}

func (__obf_015afbee33862a01 *Decoder) __obf_c5ba7e7660e774e8(__obf_8f0c71619c0382f1 *map[string]any) error {
	__obf_c26f5adfb3152475, __obf_3aa78ad28f50ed46 := __obf_015afbee33862a01.DecodeMap()
	if __obf_3aa78ad28f50ed46 != nil {
		return __obf_3aa78ad28f50ed46
	}
	*__obf_8f0c71619c0382f1 = __obf_c26f5adfb3152475
	return nil
}

func (__obf_015afbee33862a01 *Decoder) DecodeMap() (map[string]any, error) {
	__obf_56127cd370854f0b, __obf_3aa78ad28f50ed46 := __obf_015afbee33862a01.DecodeMapLen()
	if __obf_3aa78ad28f50ed46 != nil {
		return nil, __obf_3aa78ad28f50ed46
	}

	if __obf_56127cd370854f0b == -1 {
		return nil, nil
	}
	__obf_c26f5adfb3152475 := make(map[string]any, __obf_56127cd370854f0b)

	for range __obf_56127cd370854f0b {
		__obf_8f7121221504784f, __obf_3aa78ad28f50ed46 := __obf_015afbee33862a01.DecodeString()
		if __obf_3aa78ad28f50ed46 != nil {
			return nil, __obf_3aa78ad28f50ed46
		}
		__obf_a03d608795b584a3, __obf_3aa78ad28f50ed46 := __obf_015afbee33862a01.__obf_ed59dd67f1d39cb1()
		if __obf_3aa78ad28f50ed46 != nil {
			return nil, __obf_3aa78ad28f50ed46
		}
		__obf_c26f5adfb3152475[__obf_8f7121221504784f] = __obf_a03d608795b584a3
	}

	return __obf_c26f5adfb3152475, nil
}

func (__obf_015afbee33862a01 *Decoder) DecodeUntypedMap() (map[any]any, error) {
	__obf_56127cd370854f0b, __obf_3aa78ad28f50ed46 := __obf_015afbee33862a01.DecodeMapLen()
	if __obf_3aa78ad28f50ed46 != nil {
		return nil, __obf_3aa78ad28f50ed46
	}

	if __obf_56127cd370854f0b == -1 {
		return nil, nil
	}
	__obf_c26f5adfb3152475 := make(map[any]any, __obf_56127cd370854f0b)

	for range __obf_56127cd370854f0b {
		__obf_8f7121221504784f, __obf_3aa78ad28f50ed46 := __obf_015afbee33862a01.__obf_ed59dd67f1d39cb1()
		if __obf_3aa78ad28f50ed46 != nil {
			return nil, __obf_3aa78ad28f50ed46
		}
		__obf_a03d608795b584a3, __obf_3aa78ad28f50ed46 := __obf_015afbee33862a01.__obf_ed59dd67f1d39cb1()
		if __obf_3aa78ad28f50ed46 != nil {
			return nil, __obf_3aa78ad28f50ed46
		}
		__obf_c26f5adfb3152475[__obf_8f7121221504784f] = __obf_a03d608795b584a3
	}

	return __obf_c26f5adfb3152475, nil
}

// DecodeTypedMap decodes a typed map. Typed map is a map that has a fixed type for keys and values.
// Key and value types may be different.
func (__obf_015afbee33862a01 *Decoder) DecodeTypedMap() (any, error) {
	__obf_56127cd370854f0b, __obf_3aa78ad28f50ed46 := __obf_015afbee33862a01.DecodeMapLen()
	if __obf_3aa78ad28f50ed46 != nil {
		return nil, __obf_3aa78ad28f50ed46
	}
	if __obf_56127cd370854f0b <= 0 {
		return nil, nil
	}
	__obf_813ce87902031fb7, __obf_3aa78ad28f50ed46 := __obf_015afbee33862a01.__obf_ed59dd67f1d39cb1()
	if __obf_3aa78ad28f50ed46 != nil {
		return nil, __obf_3aa78ad28f50ed46
	}
	__obf_e6992bb5202a647c, __obf_3aa78ad28f50ed46 := __obf_015afbee33862a01.__obf_ed59dd67f1d39cb1()
	if __obf_3aa78ad28f50ed46 != nil {
		return nil, __obf_3aa78ad28f50ed46
	}
	__obf_659b0034deae1ed3 := reflect.TypeOf(__obf_813ce87902031fb7)
	__obf_f45b2e9db50eeed5 := reflect.TypeOf(__obf_e6992bb5202a647c)

	if !__obf_659b0034deae1ed3.Comparable() {
		return nil, fmt.Errorf("msgpack: unsupported map key: %s", __obf_659b0034deae1ed3.String())
	}
	__obf_ef271ec4cfd40c08 := reflect.MapOf(__obf_659b0034deae1ed3, __obf_f45b2e9db50eeed5)
	__obf_a3b64f93c4732565 := __obf_56127cd370854f0b
	if __obf_015afbee33862a01.__obf_6aa0efd537415192&__obf_9d8a15eb3a073456 == 0 {
		__obf_a3b64f93c4732565 = min(__obf_a3b64f93c4732565, __obf_5d6aaa093307f6f6)
	}
	__obf_635ef3f7c728ba18 := reflect.MakeMapWithSize(__obf_ef271ec4cfd40c08, __obf_a3b64f93c4732565)
	__obf_635ef3f7c728ba18.
		SetMapIndex(reflect.ValueOf(__obf_813ce87902031fb7), reflect.ValueOf(__obf_e6992bb5202a647c))
	__obf_56127cd370854f0b--
	if __obf_3aa78ad28f50ed46 := __obf_015afbee33862a01.__obf_1f0ded3b3d232ea2(__obf_635ef3f7c728ba18, __obf_56127cd370854f0b); __obf_3aa78ad28f50ed46 != nil {
		return nil, __obf_3aa78ad28f50ed46
	}

	return __obf_635ef3f7c728ba18.Interface(), nil
}

func (__obf_015afbee33862a01 *Decoder) __obf_1f0ded3b3d232ea2(__obf_85d270343a0dfe14 reflect.Value, __obf_56127cd370854f0b int) error {
	var (
		__obf_af250e9784b6a92c = __obf_85d270343a0dfe14.
					Type()
		__obf_659b0034deae1ed3 = __obf_af250e9784b6a92c.
					Key()
		__obf_f45b2e9db50eeed5 = __obf_af250e9784b6a92c.
					Elem()
	)
	for range __obf_56127cd370854f0b {
		__obf_8f7121221504784f := __obf_015afbee33862a01.__obf_4c87e5dcce642cc7(__obf_659b0034deae1ed3).Elem()
		if __obf_3aa78ad28f50ed46 := __obf_015afbee33862a01.DecodeValue(__obf_8f7121221504784f); __obf_3aa78ad28f50ed46 != nil {
			return __obf_3aa78ad28f50ed46
		}
		__obf_a03d608795b584a3 := __obf_015afbee33862a01.__obf_4c87e5dcce642cc7(__obf_f45b2e9db50eeed5).Elem()
		if __obf_3aa78ad28f50ed46 := __obf_015afbee33862a01.DecodeValue(__obf_a03d608795b584a3); __obf_3aa78ad28f50ed46 != nil {
			return __obf_3aa78ad28f50ed46
		}
		__obf_85d270343a0dfe14.
			SetMapIndex(__obf_8f7121221504784f, __obf_a03d608795b584a3)
	}

	return nil
}

func (__obf_015afbee33862a01 *Decoder) __obf_609d01762403129e(__obf_145c7a7d25eea2bd byte) error {
	__obf_56127cd370854f0b, __obf_3aa78ad28f50ed46 := __obf_015afbee33862a01.__obf_9cb09f4e20a89b72(__obf_145c7a7d25eea2bd)
	if __obf_3aa78ad28f50ed46 != nil {
		return __obf_3aa78ad28f50ed46
	}
	for range __obf_56127cd370854f0b {
		if __obf_3aa78ad28f50ed46 := __obf_015afbee33862a01.Skip(); __obf_3aa78ad28f50ed46 != nil {
			return __obf_3aa78ad28f50ed46
		}
		if __obf_3aa78ad28f50ed46 := __obf_015afbee33862a01.Skip(); __obf_3aa78ad28f50ed46 != nil {
			return __obf_3aa78ad28f50ed46
		}
	}
	return nil
}

func __obf_4f268d405fba2c41(__obf_015afbee33862a01 *Decoder, __obf_85d270343a0dfe14 reflect.Value) error {
	__obf_145c7a7d25eea2bd, __obf_3aa78ad28f50ed46 := __obf_015afbee33862a01.__obf_4c5f8d154c9027ea()
	if __obf_3aa78ad28f50ed46 != nil {
		return __obf_3aa78ad28f50ed46
	}
	__obf_56127cd370854f0b, __obf_3aa78ad28f50ed46 := __obf_015afbee33862a01.__obf_9cb09f4e20a89b72(__obf_145c7a7d25eea2bd)
	if __obf_3aa78ad28f50ed46 == nil {
		return __obf_015afbee33862a01.__obf_fae4b6ade41d23dd(__obf_85d270343a0dfe14, __obf_56127cd370854f0b)
	}

	var __obf_45fce952dcac91a7 error
	__obf_56127cd370854f0b, __obf_45fce952dcac91a7 = __obf_015afbee33862a01.__obf_5102e56f561052c2(__obf_145c7a7d25eea2bd)
	if __obf_45fce952dcac91a7 != nil {
		return __obf_3aa78ad28f50ed46
	}

	if __obf_56127cd370854f0b <= 0 {
		__obf_85d270343a0dfe14.
			Set(reflect.Zero(__obf_85d270343a0dfe14.Type()))
		return nil
	}
	__obf_fdeb38537c10bbcb := __obf_be0ec04e52f483d9.Fields(__obf_85d270343a0dfe14.Type(), __obf_015afbee33862a01.__obf_ec5d9a4ce8dd41be)
	if __obf_56127cd370854f0b != len(__obf_fdeb38537c10bbcb.List) {
		return __obf_607b50c18f29ff3f
	}

	for _, __obf_40a54345a49c3cd5 := range __obf_fdeb38537c10bbcb.List {
		if __obf_3aa78ad28f50ed46 := __obf_40a54345a49c3cd5.DecodeValue(__obf_015afbee33862a01, __obf_85d270343a0dfe14); __obf_3aa78ad28f50ed46 != nil {
			return __obf_3aa78ad28f50ed46
		}
	}

	return nil
}

func (__obf_015afbee33862a01 *Decoder) __obf_fae4b6ade41d23dd(__obf_85d270343a0dfe14 reflect.Value, __obf_56127cd370854f0b int) error {
	if __obf_56127cd370854f0b == -1 {
		__obf_85d270343a0dfe14.
			Set(reflect.Zero(__obf_85d270343a0dfe14.Type()))
		return nil
	}
	__obf_fdeb38537c10bbcb := __obf_be0ec04e52f483d9.Fields(__obf_85d270343a0dfe14.Type(), __obf_015afbee33862a01.__obf_ec5d9a4ce8dd41be)
	for range __obf_56127cd370854f0b {
		__obf_90a8fc4880887fcb, __obf_3aa78ad28f50ed46 := __obf_015afbee33862a01.__obf_aeaabfb6cb02451a()
		if __obf_3aa78ad28f50ed46 != nil {
			return __obf_3aa78ad28f50ed46
		}

		if __obf_40a54345a49c3cd5 := __obf_fdeb38537c10bbcb.Map[__obf_90a8fc4880887fcb]; __obf_40a54345a49c3cd5 != nil {
			if __obf_3aa78ad28f50ed46 := __obf_40a54345a49c3cd5.DecodeValue(__obf_015afbee33862a01, __obf_85d270343a0dfe14); __obf_3aa78ad28f50ed46 != nil {
				return __obf_3aa78ad28f50ed46
			}
			continue
		}

		if __obf_015afbee33862a01.__obf_6aa0efd537415192&__obf_2b5f039e8d42091e != 0 {
			return fmt.Errorf("msgpack: unknown field %q", __obf_90a8fc4880887fcb)
		}
		if __obf_3aa78ad28f50ed46 := __obf_015afbee33862a01.Skip(); __obf_3aa78ad28f50ed46 != nil {
			return __obf_3aa78ad28f50ed46
		}
	}

	return nil
}
