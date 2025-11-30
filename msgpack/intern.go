package __obf_de86cdc8ae98b45b

import (
	"fmt"
	"math"
	"reflect"
)

const (
	__obf_70ae198308a0c7fd = 3
	__obf_b4c1069eab73d093 = math.MaxUint16
)

var __obf_492e8397e6ec6eac = int8(math.MinInt8)

func init() {
	__obf_11268e7b8aa6c195[__obf_492e8397e6ec6eac] = &__obf_55911355a3eaccfd{
		Type:    __obf_620340e4a4ed2e24,
		Decoder: __obf_aefb4e9179a89374,
	}
}

func __obf_aefb4e9179a89374(__obf_dcaad1165bb07af9 *Decoder, __obf_17732590722140e7 reflect.Value, __obf_eef44dfac5af7de1 int) error {
	__obf_a6d03241f2c8bded, __obf_0feb3528b7b709ec := __obf_dcaad1165bb07af9.__obf_c93cc81cece3b79b(__obf_eef44dfac5af7de1)
	if __obf_0feb3528b7b709ec != nil {
		return __obf_0feb3528b7b709ec
	}
	__obf_a93d004caac96500, __obf_0feb3528b7b709ec := __obf_dcaad1165bb07af9.__obf_a60a0b612562fe30(__obf_a6d03241f2c8bded)
	if __obf_0feb3528b7b709ec != nil {
		return __obf_0feb3528b7b709ec
	}
	__obf_17732590722140e7.
		SetString(__obf_a93d004caac96500)
	return nil
}

//------------------------------------------------------------------------------

func __obf_c4de622e12ad86a8(__obf_7bae0b613da60dd3 *Encoder, __obf_17732590722140e7 reflect.Value) error {
	if __obf_17732590722140e7.IsNil() {
		return __obf_7bae0b613da60dd3.EncodeNil()
	}
	__obf_17732590722140e7 = __obf_17732590722140e7.Elem()
	if __obf_17732590722140e7.Kind() == reflect.String {
		return __obf_7bae0b613da60dd3.__obf_66adb399f41a0954(__obf_17732590722140e7.String(), true)
	}
	return __obf_7bae0b613da60dd3.EncodeValue(__obf_17732590722140e7)
}

func __obf_f85540ca8b34fdcf(__obf_7bae0b613da60dd3 *Encoder, __obf_17732590722140e7 reflect.Value) error {
	return __obf_7bae0b613da60dd3.__obf_66adb399f41a0954(__obf_17732590722140e7.String(), true)
}

func (__obf_7bae0b613da60dd3 *Encoder) __obf_66adb399f41a0954(__obf_a93d004caac96500 string, __obf_65f94b9eb2a3c80f bool) error {
	// Interned string takes at least 3 bytes. Plain string 1 byte + string len.
	if __obf_a6d03241f2c8bded, __obf_77cfff95beb3cc99 := __obf_7bae0b613da60dd3.__obf_25128eea0d506b65[__obf_a93d004caac96500]; __obf_77cfff95beb3cc99 {
		return __obf_7bae0b613da60dd3.__obf_773b609f6ddb4c5e(__obf_a6d03241f2c8bded)
	}

	if __obf_65f94b9eb2a3c80f && len(__obf_a93d004caac96500) >= __obf_70ae198308a0c7fd && len(__obf_7bae0b613da60dd3.__obf_25128eea0d506b65) < __obf_b4c1069eab73d093 {
		if __obf_7bae0b613da60dd3.__obf_25128eea0d506b65 == nil {
			__obf_7bae0b613da60dd3.__obf_25128eea0d506b65 = make(map[string]int)
		}
		__obf_a6d03241f2c8bded := len(__obf_7bae0b613da60dd3.__obf_25128eea0d506b65)
		__obf_7bae0b613da60dd3.__obf_25128eea0d506b65[__obf_a93d004caac96500] = __obf_a6d03241f2c8bded
	}

	return __obf_7bae0b613da60dd3.__obf_4920ad7ca3e462a3(__obf_a93d004caac96500)
}

func (__obf_7bae0b613da60dd3 *Encoder) __obf_773b609f6ddb4c5e(__obf_a6d03241f2c8bded int) error {
	if __obf_a6d03241f2c8bded <= math.MaxUint8 {
		if __obf_0feb3528b7b709ec := __obf_7bae0b613da60dd3.__obf_47b86f684323fc33(FixExt1); __obf_0feb3528b7b709ec != nil {
			return __obf_0feb3528b7b709ec
		}
		return __obf_7bae0b613da60dd3.__obf_fa2b4086fff3f28d(byte(__obf_492e8397e6ec6eac), uint8(__obf_a6d03241f2c8bded))
	}

	if __obf_a6d03241f2c8bded <= math.MaxUint16 {
		if __obf_0feb3528b7b709ec := __obf_7bae0b613da60dd3.__obf_47b86f684323fc33(FixExt2); __obf_0feb3528b7b709ec != nil {
			return __obf_0feb3528b7b709ec
		}
		return __obf_7bae0b613da60dd3.__obf_bd6e685495964ed5(byte(__obf_492e8397e6ec6eac), uint16(__obf_a6d03241f2c8bded))
	}

	if uint64(__obf_a6d03241f2c8bded) <= math.MaxUint32 {
		if __obf_0feb3528b7b709ec := __obf_7bae0b613da60dd3.__obf_47b86f684323fc33(FixExt4); __obf_0feb3528b7b709ec != nil {
			return __obf_0feb3528b7b709ec
		}
		return __obf_7bae0b613da60dd3.__obf_a95f338c34ae207a(byte(__obf_492e8397e6ec6eac), uint32(__obf_a6d03241f2c8bded))
	}

	return fmt.Errorf("msgpack: interned string index=%d is too large", __obf_a6d03241f2c8bded)
}

//------------------------------------------------------------------------------

func __obf_06166b79d4faa976(__obf_dcaad1165bb07af9 *Decoder, __obf_17732590722140e7 reflect.Value) error {
	__obf_a93d004caac96500, __obf_0feb3528b7b709ec := __obf_dcaad1165bb07af9.__obf_76b76bd32a6c2a62(true)
	if __obf_0feb3528b7b709ec == nil {
		__obf_17732590722140e7.
			Set(reflect.ValueOf(__obf_a93d004caac96500))
		return nil
	}
	if _, __obf_77cfff95beb3cc99 := __obf_0feb3528b7b709ec.(__obf_150b3812dfd829ef); !__obf_77cfff95beb3cc99 {
		return __obf_0feb3528b7b709ec
	}

	if __obf_0feb3528b7b709ec := __obf_dcaad1165bb07af9.__obf_a93d004caac96500.UnreadByte(); __obf_0feb3528b7b709ec != nil {
		return __obf_0feb3528b7b709ec
	}
	return __obf_19f61f7db3d5aa0f(__obf_dcaad1165bb07af9, __obf_17732590722140e7)
}

func __obf_0e872b4337eae851(__obf_dcaad1165bb07af9 *Decoder, __obf_17732590722140e7 reflect.Value) error {
	__obf_a93d004caac96500, __obf_0feb3528b7b709ec := __obf_dcaad1165bb07af9.__obf_76b76bd32a6c2a62(true)
	if __obf_0feb3528b7b709ec != nil {
		return __obf_0feb3528b7b709ec
	}
	__obf_17732590722140e7.
		SetString(__obf_a93d004caac96500)
	return nil
}

func (__obf_dcaad1165bb07af9 *Decoder) __obf_76b76bd32a6c2a62(__obf_65f94b9eb2a3c80f bool) (string, error) {
	__obf_ecf6d2d3253a058d, __obf_0feb3528b7b709ec := __obf_dcaad1165bb07af9.__obf_74ac05d9a9da01f4()
	if __obf_0feb3528b7b709ec != nil {
		return "", __obf_0feb3528b7b709ec
	}

	if IsFixedString(__obf_ecf6d2d3253a058d) {
		__obf_2b0247e819b1bf4a := int(__obf_ecf6d2d3253a058d & FixedStrMask)
		return __obf_dcaad1165bb07af9.__obf_b3eec65a1e4f0c6a(__obf_2b0247e819b1bf4a, __obf_65f94b9eb2a3c80f)
	}

	switch __obf_ecf6d2d3253a058d {
	case Nil:
		return "", nil
	case FixExt1, FixExt2, FixExt4:
		__obf_203b7fcdb6d67b3e, __obf_eef44dfac5af7de1, __obf_0feb3528b7b709ec := __obf_dcaad1165bb07af9.__obf_73f4503729af0402(__obf_ecf6d2d3253a058d)
		if __obf_0feb3528b7b709ec != nil {
			return "", __obf_0feb3528b7b709ec
		}
		if __obf_203b7fcdb6d67b3e != __obf_492e8397e6ec6eac {
			__obf_0feb3528b7b709ec := fmt.Errorf("msgpack: got ext type=%d, wanted %d", __obf_203b7fcdb6d67b3e, __obf_492e8397e6ec6eac)
			return "", __obf_0feb3528b7b709ec
		}
		__obf_a6d03241f2c8bded, __obf_0feb3528b7b709ec := __obf_dcaad1165bb07af9.__obf_c93cc81cece3b79b(__obf_eef44dfac5af7de1)
		if __obf_0feb3528b7b709ec != nil {
			return "", __obf_0feb3528b7b709ec
		}

		return __obf_dcaad1165bb07af9.__obf_a60a0b612562fe30(__obf_a6d03241f2c8bded)
	case Str8, Bin8:
		__obf_2b0247e819b1bf4a, __obf_0feb3528b7b709ec := __obf_dcaad1165bb07af9.uint8()
		if __obf_0feb3528b7b709ec != nil {
			return "", __obf_0feb3528b7b709ec
		}
		return __obf_dcaad1165bb07af9.__obf_b3eec65a1e4f0c6a(int(__obf_2b0247e819b1bf4a), __obf_65f94b9eb2a3c80f)
	case Str16, Bin16:
		__obf_2b0247e819b1bf4a, __obf_0feb3528b7b709ec := __obf_dcaad1165bb07af9.uint16()
		if __obf_0feb3528b7b709ec != nil {
			return "", __obf_0feb3528b7b709ec
		}
		return __obf_dcaad1165bb07af9.__obf_b3eec65a1e4f0c6a(int(__obf_2b0247e819b1bf4a), __obf_65f94b9eb2a3c80f)
	case Str32, Bin32:
		__obf_2b0247e819b1bf4a, __obf_0feb3528b7b709ec := __obf_dcaad1165bb07af9.uint32()
		if __obf_0feb3528b7b709ec != nil {
			return "", __obf_0feb3528b7b709ec
		}
		return __obf_dcaad1165bb07af9.__obf_b3eec65a1e4f0c6a(int(__obf_2b0247e819b1bf4a), __obf_65f94b9eb2a3c80f)
	}

	return "", __obf_150b3812dfd829ef{__obf_34e3ba264a6bb541: __obf_ecf6d2d3253a058d, __obf_ee374540697f97c0: "interned string"}
}

func (__obf_dcaad1165bb07af9 *Decoder) __obf_c93cc81cece3b79b(__obf_eef44dfac5af7de1 int) (int, error) {
	switch __obf_eef44dfac5af7de1 {
	case 1:
		__obf_2b0247e819b1bf4a, __obf_0feb3528b7b709ec := __obf_dcaad1165bb07af9.uint8()
		if __obf_0feb3528b7b709ec != nil {
			return 0, __obf_0feb3528b7b709ec
		}
		return int(__obf_2b0247e819b1bf4a), nil
	case 2:
		__obf_2b0247e819b1bf4a, __obf_0feb3528b7b709ec := __obf_dcaad1165bb07af9.uint16()
		if __obf_0feb3528b7b709ec != nil {
			return 0, __obf_0feb3528b7b709ec
		}
		return int(__obf_2b0247e819b1bf4a), nil
	case 4:
		__obf_2b0247e819b1bf4a, __obf_0feb3528b7b709ec := __obf_dcaad1165bb07af9.uint32()
		if __obf_0feb3528b7b709ec != nil {
			return 0, __obf_0feb3528b7b709ec
		}
		return int(__obf_2b0247e819b1bf4a), nil
	}
	__obf_0feb3528b7b709ec := fmt.Errorf("msgpack: unsupported ext len=%d decoding interned string", __obf_eef44dfac5af7de1)
	return 0, __obf_0feb3528b7b709ec
}

func (__obf_dcaad1165bb07af9 *Decoder) __obf_a60a0b612562fe30(__obf_a6d03241f2c8bded int) (string, error) {
	if __obf_a6d03241f2c8bded >= len(__obf_dcaad1165bb07af9.__obf_25128eea0d506b65) {
		__obf_0feb3528b7b709ec := fmt.Errorf("msgpack: interned string at index=%d does not exist", __obf_a6d03241f2c8bded)
		return "", __obf_0feb3528b7b709ec
	}
	return __obf_dcaad1165bb07af9.__obf_25128eea0d506b65[__obf_a6d03241f2c8bded], nil
}

func (__obf_dcaad1165bb07af9 *Decoder) __obf_b3eec65a1e4f0c6a(__obf_2b0247e819b1bf4a int, __obf_65f94b9eb2a3c80f bool) (string, error) {
	if __obf_2b0247e819b1bf4a <= 0 {
		return "", nil
	}
	__obf_a93d004caac96500, __obf_0feb3528b7b709ec := __obf_dcaad1165bb07af9.__obf_281f9be3791bf742(__obf_2b0247e819b1bf4a)
	if __obf_0feb3528b7b709ec != nil {
		return "", __obf_0feb3528b7b709ec
	}

	if __obf_65f94b9eb2a3c80f && len(__obf_a93d004caac96500) >= __obf_70ae198308a0c7fd && len(__obf_dcaad1165bb07af9.__obf_25128eea0d506b65) < __obf_b4c1069eab73d093 {
		__obf_dcaad1165bb07af9.__obf_25128eea0d506b65 = append(__obf_dcaad1165bb07af9.__obf_25128eea0d506b65, __obf_a93d004caac96500)
	}

	return __obf_a93d004caac96500, nil
}
