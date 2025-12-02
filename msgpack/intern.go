package __obf_a935eb7f548271a4

import (
	"fmt"
	"math"
	"reflect"
)

const (
	__obf_971340e496937bc3 = 3
	__obf_03790f8736d3c526 = math.MaxUint16
)

var __obf_4d4ad0d76fc43916 = int8(math.MinInt8)

func init() {
	__obf_335898695997d965[__obf_4d4ad0d76fc43916] = &__obf_9b64c962e55dc648{
		Type:    __obf_b128fe7c7d10a668,
		Decoder: __obf_4f40a74e1cbdb894,
	}
}

func __obf_4f40a74e1cbdb894(__obf_a21885da2425f2b2 *Decoder, __obf_6d570581f4b60dbc reflect.Value, __obf_f9e21d535a3562ea int) error {
	__obf_b8b71f09b098e0ae, __obf_4d327e1cd40c2e21 := __obf_a21885da2425f2b2.__obf_99d06338875dd380(__obf_f9e21d535a3562ea)
	if __obf_4d327e1cd40c2e21 != nil {
		return __obf_4d327e1cd40c2e21
	}
	__obf_b62c60fba2fd9788, __obf_4d327e1cd40c2e21 := __obf_a21885da2425f2b2.__obf_53754b3ca2dbbbab(__obf_b8b71f09b098e0ae)
	if __obf_4d327e1cd40c2e21 != nil {
		return __obf_4d327e1cd40c2e21
	}
	__obf_6d570581f4b60dbc.
		SetString(__obf_b62c60fba2fd9788)
	return nil
}

//------------------------------------------------------------------------------

func __obf_14ce8aee2699d8c1(__obf_ed37a34c347049f3 *Encoder, __obf_6d570581f4b60dbc reflect.Value) error {
	if __obf_6d570581f4b60dbc.IsNil() {
		return __obf_ed37a34c347049f3.EncodeNil()
	}
	__obf_6d570581f4b60dbc = __obf_6d570581f4b60dbc.Elem()
	if __obf_6d570581f4b60dbc.Kind() == reflect.String {
		return __obf_ed37a34c347049f3.__obf_f3306d3e9feeada5(__obf_6d570581f4b60dbc.String(), true)
	}
	return __obf_ed37a34c347049f3.EncodeValue(__obf_6d570581f4b60dbc)
}

func __obf_f1fd7fac8f632e34(__obf_ed37a34c347049f3 *Encoder, __obf_6d570581f4b60dbc reflect.Value) error {
	return __obf_ed37a34c347049f3.__obf_f3306d3e9feeada5(__obf_6d570581f4b60dbc.String(), true)
}

func (__obf_ed37a34c347049f3 *Encoder) __obf_f3306d3e9feeada5(__obf_b62c60fba2fd9788 string, __obf_65a759dbf3ace040 bool) error {
	// Interned string takes at least 3 bytes. Plain string 1 byte + string len.
	if __obf_b8b71f09b098e0ae, __obf_826ac2dbb957d7df := __obf_ed37a34c347049f3.__obf_2cdd7fad308aef1d[__obf_b62c60fba2fd9788]; __obf_826ac2dbb957d7df {
		return __obf_ed37a34c347049f3.__obf_c0d34985e47ef061(__obf_b8b71f09b098e0ae)
	}

	if __obf_65a759dbf3ace040 && len(__obf_b62c60fba2fd9788) >= __obf_971340e496937bc3 && len(__obf_ed37a34c347049f3.__obf_2cdd7fad308aef1d) < __obf_03790f8736d3c526 {
		if __obf_ed37a34c347049f3.__obf_2cdd7fad308aef1d == nil {
			__obf_ed37a34c347049f3.__obf_2cdd7fad308aef1d = make(map[string]int)
		}
		__obf_b8b71f09b098e0ae := len(__obf_ed37a34c347049f3.__obf_2cdd7fad308aef1d)
		__obf_ed37a34c347049f3.__obf_2cdd7fad308aef1d[__obf_b62c60fba2fd9788] = __obf_b8b71f09b098e0ae
	}

	return __obf_ed37a34c347049f3.__obf_dd784acb29c85421(__obf_b62c60fba2fd9788)
}

func (__obf_ed37a34c347049f3 *Encoder) __obf_c0d34985e47ef061(__obf_b8b71f09b098e0ae int) error {
	if __obf_b8b71f09b098e0ae <= math.MaxUint8 {
		if __obf_4d327e1cd40c2e21 := __obf_ed37a34c347049f3.__obf_2252dbfd2473b643(FixExt1); __obf_4d327e1cd40c2e21 != nil {
			return __obf_4d327e1cd40c2e21
		}
		return __obf_ed37a34c347049f3.__obf_327ee6f3f671bfd3(byte(__obf_4d4ad0d76fc43916), uint8(__obf_b8b71f09b098e0ae))
	}

	if __obf_b8b71f09b098e0ae <= math.MaxUint16 {
		if __obf_4d327e1cd40c2e21 := __obf_ed37a34c347049f3.__obf_2252dbfd2473b643(FixExt2); __obf_4d327e1cd40c2e21 != nil {
			return __obf_4d327e1cd40c2e21
		}
		return __obf_ed37a34c347049f3.__obf_dd92b251146e86fb(byte(__obf_4d4ad0d76fc43916), uint16(__obf_b8b71f09b098e0ae))
	}

	if uint64(__obf_b8b71f09b098e0ae) <= math.MaxUint32 {
		if __obf_4d327e1cd40c2e21 := __obf_ed37a34c347049f3.__obf_2252dbfd2473b643(FixExt4); __obf_4d327e1cd40c2e21 != nil {
			return __obf_4d327e1cd40c2e21
		}
		return __obf_ed37a34c347049f3.__obf_21945a55177e9fbd(byte(__obf_4d4ad0d76fc43916), uint32(__obf_b8b71f09b098e0ae))
	}

	return fmt.Errorf("msgpack: interned string index=%d is too large", __obf_b8b71f09b098e0ae)
}

//------------------------------------------------------------------------------

func __obf_9d0cdf74c5f70d62(__obf_a21885da2425f2b2 *Decoder, __obf_6d570581f4b60dbc reflect.Value) error {
	__obf_b62c60fba2fd9788, __obf_4d327e1cd40c2e21 := __obf_a21885da2425f2b2.__obf_f43916d12d5b1cf3(true)
	if __obf_4d327e1cd40c2e21 == nil {
		__obf_6d570581f4b60dbc.
			Set(reflect.ValueOf(__obf_b62c60fba2fd9788))
		return nil
	}
	if _, __obf_826ac2dbb957d7df := __obf_4d327e1cd40c2e21.(__obf_769655cd52f45f6d); !__obf_826ac2dbb957d7df {
		return __obf_4d327e1cd40c2e21
	}

	if __obf_4d327e1cd40c2e21 := __obf_a21885da2425f2b2.__obf_b62c60fba2fd9788.UnreadByte(); __obf_4d327e1cd40c2e21 != nil {
		return __obf_4d327e1cd40c2e21
	}
	return __obf_ceed3ef343f3463a(__obf_a21885da2425f2b2, __obf_6d570581f4b60dbc)
}

func __obf_f81a4aeb4252941f(__obf_a21885da2425f2b2 *Decoder, __obf_6d570581f4b60dbc reflect.Value) error {
	__obf_b62c60fba2fd9788, __obf_4d327e1cd40c2e21 := __obf_a21885da2425f2b2.__obf_f43916d12d5b1cf3(true)
	if __obf_4d327e1cd40c2e21 != nil {
		return __obf_4d327e1cd40c2e21
	}
	__obf_6d570581f4b60dbc.
		SetString(__obf_b62c60fba2fd9788)
	return nil
}

func (__obf_a21885da2425f2b2 *Decoder) __obf_f43916d12d5b1cf3(__obf_65a759dbf3ace040 bool) (string, error) {
	__obf_f5df560f4d67421b, __obf_4d327e1cd40c2e21 := __obf_a21885da2425f2b2.__obf_61e5879fcf7b35ce()
	if __obf_4d327e1cd40c2e21 != nil {
		return "", __obf_4d327e1cd40c2e21
	}

	if IsFixedString(__obf_f5df560f4d67421b) {
		__obf_326af9bd942662ac := int(__obf_f5df560f4d67421b & FixedStrMask)
		return __obf_a21885da2425f2b2.__obf_4bf50a57dd5d7421(__obf_326af9bd942662ac, __obf_65a759dbf3ace040)
	}

	switch __obf_f5df560f4d67421b {
	case Nil:
		return "", nil
	case FixExt1, FixExt2, FixExt4:
		__obf_d90f938afeddd764, __obf_f9e21d535a3562ea, __obf_4d327e1cd40c2e21 := __obf_a21885da2425f2b2.__obf_4dc04de1e73ebdac(__obf_f5df560f4d67421b)
		if __obf_4d327e1cd40c2e21 != nil {
			return "", __obf_4d327e1cd40c2e21
		}
		if __obf_d90f938afeddd764 != __obf_4d4ad0d76fc43916 {
			__obf_4d327e1cd40c2e21 := fmt.Errorf("msgpack: got ext type=%d, wanted %d", __obf_d90f938afeddd764, __obf_4d4ad0d76fc43916)
			return "", __obf_4d327e1cd40c2e21
		}
		__obf_b8b71f09b098e0ae, __obf_4d327e1cd40c2e21 := __obf_a21885da2425f2b2.__obf_99d06338875dd380(__obf_f9e21d535a3562ea)
		if __obf_4d327e1cd40c2e21 != nil {
			return "", __obf_4d327e1cd40c2e21
		}

		return __obf_a21885da2425f2b2.__obf_53754b3ca2dbbbab(__obf_b8b71f09b098e0ae)
	case Str8, Bin8:
		__obf_326af9bd942662ac, __obf_4d327e1cd40c2e21 := __obf_a21885da2425f2b2.uint8()
		if __obf_4d327e1cd40c2e21 != nil {
			return "", __obf_4d327e1cd40c2e21
		}
		return __obf_a21885da2425f2b2.__obf_4bf50a57dd5d7421(int(__obf_326af9bd942662ac), __obf_65a759dbf3ace040)
	case Str16, Bin16:
		__obf_326af9bd942662ac, __obf_4d327e1cd40c2e21 := __obf_a21885da2425f2b2.uint16()
		if __obf_4d327e1cd40c2e21 != nil {
			return "", __obf_4d327e1cd40c2e21
		}
		return __obf_a21885da2425f2b2.__obf_4bf50a57dd5d7421(int(__obf_326af9bd942662ac), __obf_65a759dbf3ace040)
	case Str32, Bin32:
		__obf_326af9bd942662ac, __obf_4d327e1cd40c2e21 := __obf_a21885da2425f2b2.uint32()
		if __obf_4d327e1cd40c2e21 != nil {
			return "", __obf_4d327e1cd40c2e21
		}
		return __obf_a21885da2425f2b2.__obf_4bf50a57dd5d7421(int(__obf_326af9bd942662ac), __obf_65a759dbf3ace040)
	}

	return "", __obf_769655cd52f45f6d{__obf_b983039ef2a336eb: __obf_f5df560f4d67421b, __obf_feeba0b68911eadc: "interned string"}
}

func (__obf_a21885da2425f2b2 *Decoder) __obf_99d06338875dd380(__obf_f9e21d535a3562ea int) (int, error) {
	switch __obf_f9e21d535a3562ea {
	case 1:
		__obf_326af9bd942662ac, __obf_4d327e1cd40c2e21 := __obf_a21885da2425f2b2.uint8()
		if __obf_4d327e1cd40c2e21 != nil {
			return 0, __obf_4d327e1cd40c2e21
		}
		return int(__obf_326af9bd942662ac), nil
	case 2:
		__obf_326af9bd942662ac, __obf_4d327e1cd40c2e21 := __obf_a21885da2425f2b2.uint16()
		if __obf_4d327e1cd40c2e21 != nil {
			return 0, __obf_4d327e1cd40c2e21
		}
		return int(__obf_326af9bd942662ac), nil
	case 4:
		__obf_326af9bd942662ac, __obf_4d327e1cd40c2e21 := __obf_a21885da2425f2b2.uint32()
		if __obf_4d327e1cd40c2e21 != nil {
			return 0, __obf_4d327e1cd40c2e21
		}
		return int(__obf_326af9bd942662ac), nil
	}
	__obf_4d327e1cd40c2e21 := fmt.Errorf("msgpack: unsupported ext len=%d decoding interned string", __obf_f9e21d535a3562ea)
	return 0, __obf_4d327e1cd40c2e21
}

func (__obf_a21885da2425f2b2 *Decoder) __obf_53754b3ca2dbbbab(__obf_b8b71f09b098e0ae int) (string, error) {
	if __obf_b8b71f09b098e0ae >= len(__obf_a21885da2425f2b2.__obf_2cdd7fad308aef1d) {
		__obf_4d327e1cd40c2e21 := fmt.Errorf("msgpack: interned string at index=%d does not exist", __obf_b8b71f09b098e0ae)
		return "", __obf_4d327e1cd40c2e21
	}
	return __obf_a21885da2425f2b2.__obf_2cdd7fad308aef1d[__obf_b8b71f09b098e0ae], nil
}

func (__obf_a21885da2425f2b2 *Decoder) __obf_4bf50a57dd5d7421(__obf_326af9bd942662ac int, __obf_65a759dbf3ace040 bool) (string, error) {
	if __obf_326af9bd942662ac <= 0 {
		return "", nil
	}
	__obf_b62c60fba2fd9788, __obf_4d327e1cd40c2e21 := __obf_a21885da2425f2b2.__obf_bb558e6f117ed933(__obf_326af9bd942662ac)
	if __obf_4d327e1cd40c2e21 != nil {
		return "", __obf_4d327e1cd40c2e21
	}

	if __obf_65a759dbf3ace040 && len(__obf_b62c60fba2fd9788) >= __obf_971340e496937bc3 && len(__obf_a21885da2425f2b2.__obf_2cdd7fad308aef1d) < __obf_03790f8736d3c526 {
		__obf_a21885da2425f2b2.__obf_2cdd7fad308aef1d = append(__obf_a21885da2425f2b2.__obf_2cdd7fad308aef1d, __obf_b62c60fba2fd9788)
	}

	return __obf_b62c60fba2fd9788, nil
}
