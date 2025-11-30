package __obf_3e0c303610a19bd4

import (
	"fmt"
	"math"
	"reflect"
)

const (
	__obf_7d6f8399c1a04ea0 = 3
	__obf_dfd218286236277a = math.MaxUint16
)

var __obf_5eb26c8ee897c63c = int8(math.MinInt8)

func init() {
	__obf_d7c9956a45f4419c[__obf_5eb26c8ee897c63c] = &__obf_6be15d2128e93f39{
		Type:    __obf_24e46306ee5355ff,
		Decoder: __obf_6a5629e1b3a86b03,
	}
}

func __obf_6a5629e1b3a86b03(__obf_dc35117108ba8439 *Decoder, __obf_63bbcee86d44fdde reflect.Value, __obf_7e99dd6aa706c57b int) error {
	__obf_4040153f7700644e, __obf_8882f6cf6e378ded := __obf_dc35117108ba8439.__obf_48883bece586e37e(__obf_7e99dd6aa706c57b)
	if __obf_8882f6cf6e378ded != nil {
		return __obf_8882f6cf6e378ded
	}
	__obf_61027e0491b6dd3d, __obf_8882f6cf6e378ded := __obf_dc35117108ba8439.__obf_a8c76388b4390765(__obf_4040153f7700644e)
	if __obf_8882f6cf6e378ded != nil {
		return __obf_8882f6cf6e378ded
	}
	__obf_63bbcee86d44fdde.
		SetString(__obf_61027e0491b6dd3d)
	return nil
}

//------------------------------------------------------------------------------

func __obf_864478adcb603602(__obf_77240dc7776b60b4 *Encoder, __obf_63bbcee86d44fdde reflect.Value) error {
	if __obf_63bbcee86d44fdde.IsNil() {
		return __obf_77240dc7776b60b4.EncodeNil()
	}
	__obf_63bbcee86d44fdde = __obf_63bbcee86d44fdde.Elem()
	if __obf_63bbcee86d44fdde.Kind() == reflect.String {
		return __obf_77240dc7776b60b4.__obf_553232b05cafcb2b(__obf_63bbcee86d44fdde.String(), true)
	}
	return __obf_77240dc7776b60b4.EncodeValue(__obf_63bbcee86d44fdde)
}

func __obf_51341bc888154b6f(__obf_77240dc7776b60b4 *Encoder, __obf_63bbcee86d44fdde reflect.Value) error {
	return __obf_77240dc7776b60b4.__obf_553232b05cafcb2b(__obf_63bbcee86d44fdde.String(), true)
}

func (__obf_77240dc7776b60b4 *Encoder) __obf_553232b05cafcb2b(__obf_61027e0491b6dd3d string, __obf_14e539ac760532b2 bool) error {
	// Interned string takes at least 3 bytes. Plain string 1 byte + string len.
	if __obf_4040153f7700644e, __obf_ce8bef141eff8aab := __obf_77240dc7776b60b4.__obf_a22a31b815544cae[__obf_61027e0491b6dd3d]; __obf_ce8bef141eff8aab {
		return __obf_77240dc7776b60b4.__obf_77548058c113d558(__obf_4040153f7700644e)
	}

	if __obf_14e539ac760532b2 && len(__obf_61027e0491b6dd3d) >= __obf_7d6f8399c1a04ea0 && len(__obf_77240dc7776b60b4.__obf_a22a31b815544cae) < __obf_dfd218286236277a {
		if __obf_77240dc7776b60b4.__obf_a22a31b815544cae == nil {
			__obf_77240dc7776b60b4.__obf_a22a31b815544cae = make(map[string]int)
		}
		__obf_4040153f7700644e := len(__obf_77240dc7776b60b4.__obf_a22a31b815544cae)
		__obf_77240dc7776b60b4.__obf_a22a31b815544cae[__obf_61027e0491b6dd3d] = __obf_4040153f7700644e
	}

	return __obf_77240dc7776b60b4.__obf_a30f8acd02da8779(__obf_61027e0491b6dd3d)
}

func (__obf_77240dc7776b60b4 *Encoder) __obf_77548058c113d558(__obf_4040153f7700644e int) error {
	if __obf_4040153f7700644e <= math.MaxUint8 {
		if __obf_8882f6cf6e378ded := __obf_77240dc7776b60b4.__obf_e8d9aa5d4f5eb974(FixExt1); __obf_8882f6cf6e378ded != nil {
			return __obf_8882f6cf6e378ded
		}
		return __obf_77240dc7776b60b4.__obf_a5ea8b1a318bf9eb(byte(__obf_5eb26c8ee897c63c), uint8(__obf_4040153f7700644e))
	}

	if __obf_4040153f7700644e <= math.MaxUint16 {
		if __obf_8882f6cf6e378ded := __obf_77240dc7776b60b4.__obf_e8d9aa5d4f5eb974(FixExt2); __obf_8882f6cf6e378ded != nil {
			return __obf_8882f6cf6e378ded
		}
		return __obf_77240dc7776b60b4.__obf_9ecc220d5defc1d9(byte(__obf_5eb26c8ee897c63c), uint16(__obf_4040153f7700644e))
	}

	if uint64(__obf_4040153f7700644e) <= math.MaxUint32 {
		if __obf_8882f6cf6e378ded := __obf_77240dc7776b60b4.__obf_e8d9aa5d4f5eb974(FixExt4); __obf_8882f6cf6e378ded != nil {
			return __obf_8882f6cf6e378ded
		}
		return __obf_77240dc7776b60b4.__obf_70cbe6e3e07b6347(byte(__obf_5eb26c8ee897c63c), uint32(__obf_4040153f7700644e))
	}

	return fmt.Errorf("msgpack: interned string index=%d is too large", __obf_4040153f7700644e)
}

//------------------------------------------------------------------------------

func __obf_68c203f3636fd603(__obf_dc35117108ba8439 *Decoder, __obf_63bbcee86d44fdde reflect.Value) error {
	__obf_61027e0491b6dd3d, __obf_8882f6cf6e378ded := __obf_dc35117108ba8439.__obf_7f3d352cfbc0638e(true)
	if __obf_8882f6cf6e378ded == nil {
		__obf_63bbcee86d44fdde.
			Set(reflect.ValueOf(__obf_61027e0491b6dd3d))
		return nil
	}
	if _, __obf_ce8bef141eff8aab := __obf_8882f6cf6e378ded.(__obf_dc8806ddbc8f0bfa); !__obf_ce8bef141eff8aab {
		return __obf_8882f6cf6e378ded
	}

	if __obf_8882f6cf6e378ded := __obf_dc35117108ba8439.__obf_61027e0491b6dd3d.UnreadByte(); __obf_8882f6cf6e378ded != nil {
		return __obf_8882f6cf6e378ded
	}
	return __obf_c0b16016f640c84c(__obf_dc35117108ba8439, __obf_63bbcee86d44fdde)
}

func __obf_113b700fc6a29ee1(__obf_dc35117108ba8439 *Decoder, __obf_63bbcee86d44fdde reflect.Value) error {
	__obf_61027e0491b6dd3d, __obf_8882f6cf6e378ded := __obf_dc35117108ba8439.__obf_7f3d352cfbc0638e(true)
	if __obf_8882f6cf6e378ded != nil {
		return __obf_8882f6cf6e378ded
	}
	__obf_63bbcee86d44fdde.
		SetString(__obf_61027e0491b6dd3d)
	return nil
}

func (__obf_dc35117108ba8439 *Decoder) __obf_7f3d352cfbc0638e(__obf_14e539ac760532b2 bool) (string, error) {
	__obf_e46289218af709bf, __obf_8882f6cf6e378ded := __obf_dc35117108ba8439.__obf_5b7d85f9093902c3()
	if __obf_8882f6cf6e378ded != nil {
		return "", __obf_8882f6cf6e378ded
	}

	if IsFixedString(__obf_e46289218af709bf) {
		__obf_4909ae60ffbb8e53 := int(__obf_e46289218af709bf & FixedStrMask)
		return __obf_dc35117108ba8439.__obf_4b9fe77516ccbc85(__obf_4909ae60ffbb8e53, __obf_14e539ac760532b2)
	}

	switch __obf_e46289218af709bf {
	case Nil:
		return "", nil
	case FixExt1, FixExt2, FixExt4:
		__obf_bcf14844c70bfa3b, __obf_7e99dd6aa706c57b, __obf_8882f6cf6e378ded := __obf_dc35117108ba8439.__obf_3ac0957b06c080b5(__obf_e46289218af709bf)
		if __obf_8882f6cf6e378ded != nil {
			return "", __obf_8882f6cf6e378ded
		}
		if __obf_bcf14844c70bfa3b != __obf_5eb26c8ee897c63c {
			__obf_8882f6cf6e378ded := fmt.Errorf("msgpack: got ext type=%d, wanted %d", __obf_bcf14844c70bfa3b, __obf_5eb26c8ee897c63c)
			return "", __obf_8882f6cf6e378ded
		}
		__obf_4040153f7700644e, __obf_8882f6cf6e378ded := __obf_dc35117108ba8439.__obf_48883bece586e37e(__obf_7e99dd6aa706c57b)
		if __obf_8882f6cf6e378ded != nil {
			return "", __obf_8882f6cf6e378ded
		}

		return __obf_dc35117108ba8439.__obf_a8c76388b4390765(__obf_4040153f7700644e)
	case Str8, Bin8:
		__obf_4909ae60ffbb8e53, __obf_8882f6cf6e378ded := __obf_dc35117108ba8439.uint8()
		if __obf_8882f6cf6e378ded != nil {
			return "", __obf_8882f6cf6e378ded
		}
		return __obf_dc35117108ba8439.__obf_4b9fe77516ccbc85(int(__obf_4909ae60ffbb8e53), __obf_14e539ac760532b2)
	case Str16, Bin16:
		__obf_4909ae60ffbb8e53, __obf_8882f6cf6e378ded := __obf_dc35117108ba8439.uint16()
		if __obf_8882f6cf6e378ded != nil {
			return "", __obf_8882f6cf6e378ded
		}
		return __obf_dc35117108ba8439.__obf_4b9fe77516ccbc85(int(__obf_4909ae60ffbb8e53), __obf_14e539ac760532b2)
	case Str32, Bin32:
		__obf_4909ae60ffbb8e53, __obf_8882f6cf6e378ded := __obf_dc35117108ba8439.uint32()
		if __obf_8882f6cf6e378ded != nil {
			return "", __obf_8882f6cf6e378ded
		}
		return __obf_dc35117108ba8439.__obf_4b9fe77516ccbc85(int(__obf_4909ae60ffbb8e53), __obf_14e539ac760532b2)
	}

	return "", __obf_dc8806ddbc8f0bfa{__obf_545372fefbb733e5: __obf_e46289218af709bf, __obf_41a3520de571ca37: "interned string"}
}

func (__obf_dc35117108ba8439 *Decoder) __obf_48883bece586e37e(__obf_7e99dd6aa706c57b int) (int, error) {
	switch __obf_7e99dd6aa706c57b {
	case 1:
		__obf_4909ae60ffbb8e53, __obf_8882f6cf6e378ded := __obf_dc35117108ba8439.uint8()
		if __obf_8882f6cf6e378ded != nil {
			return 0, __obf_8882f6cf6e378ded
		}
		return int(__obf_4909ae60ffbb8e53), nil
	case 2:
		__obf_4909ae60ffbb8e53, __obf_8882f6cf6e378ded := __obf_dc35117108ba8439.uint16()
		if __obf_8882f6cf6e378ded != nil {
			return 0, __obf_8882f6cf6e378ded
		}
		return int(__obf_4909ae60ffbb8e53), nil
	case 4:
		__obf_4909ae60ffbb8e53, __obf_8882f6cf6e378ded := __obf_dc35117108ba8439.uint32()
		if __obf_8882f6cf6e378ded != nil {
			return 0, __obf_8882f6cf6e378ded
		}
		return int(__obf_4909ae60ffbb8e53), nil
	}
	__obf_8882f6cf6e378ded := fmt.Errorf("msgpack: unsupported ext len=%d decoding interned string", __obf_7e99dd6aa706c57b)
	return 0, __obf_8882f6cf6e378ded
}

func (__obf_dc35117108ba8439 *Decoder) __obf_a8c76388b4390765(__obf_4040153f7700644e int) (string, error) {
	if __obf_4040153f7700644e >= len(__obf_dc35117108ba8439.__obf_a22a31b815544cae) {
		__obf_8882f6cf6e378ded := fmt.Errorf("msgpack: interned string at index=%d does not exist", __obf_4040153f7700644e)
		return "", __obf_8882f6cf6e378ded
	}
	return __obf_dc35117108ba8439.__obf_a22a31b815544cae[__obf_4040153f7700644e], nil
}

func (__obf_dc35117108ba8439 *Decoder) __obf_4b9fe77516ccbc85(__obf_4909ae60ffbb8e53 int, __obf_14e539ac760532b2 bool) (string, error) {
	if __obf_4909ae60ffbb8e53 <= 0 {
		return "", nil
	}
	__obf_61027e0491b6dd3d, __obf_8882f6cf6e378ded := __obf_dc35117108ba8439.__obf_9198a6182aacf2fb(__obf_4909ae60ffbb8e53)
	if __obf_8882f6cf6e378ded != nil {
		return "", __obf_8882f6cf6e378ded
	}

	if __obf_14e539ac760532b2 && len(__obf_61027e0491b6dd3d) >= __obf_7d6f8399c1a04ea0 && len(__obf_dc35117108ba8439.__obf_a22a31b815544cae) < __obf_dfd218286236277a {
		__obf_dc35117108ba8439.__obf_a22a31b815544cae = append(__obf_dc35117108ba8439.__obf_a22a31b815544cae, __obf_61027e0491b6dd3d)
	}

	return __obf_61027e0491b6dd3d, nil
}
