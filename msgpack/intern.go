package __obf_a4aad98aaf3178f4

import (
	"fmt"
	"math"
	"reflect"
)

const (
	__obf_fa31df83bfa41769 = 3
	__obf_7b33188291e2bf2f = math.MaxUint16
)

var __obf_24fe567d301ad58a = int8(math.MinInt8)

func init() {
	__obf_b394b4262313f390[__obf_24fe567d301ad58a] = &__obf_3346ed3406077a0b{
		Type:    __obf_60ac326f01b0b63a,
		Decoder: __obf_07ae933bb5308aa1,
	}
}

func __obf_07ae933bb5308aa1(__obf_613397fefdec0ed0 *Decoder, __obf_a1f43267eeb48a1a reflect.Value, __obf_f975a2b18cb9c21c int) error {
	__obf_82aee0190d456a88, __obf_4061ca0039150c39 := __obf_613397fefdec0ed0.__obf_404e8bd8820334d0(__obf_f975a2b18cb9c21c)
	if __obf_4061ca0039150c39 != nil {
		return __obf_4061ca0039150c39
	}
	__obf_759f458bfa19abba, __obf_4061ca0039150c39 := __obf_613397fefdec0ed0.__obf_0818908d2fd05af4(__obf_82aee0190d456a88)
	if __obf_4061ca0039150c39 != nil {
		return __obf_4061ca0039150c39
	}
	__obf_a1f43267eeb48a1a.
		SetString(__obf_759f458bfa19abba)
	return nil
}

//------------------------------------------------------------------------------

func __obf_39ec8e83b1a4666e(__obf_2c8e97779108ab17 *Encoder, __obf_a1f43267eeb48a1a reflect.Value) error {
	if __obf_a1f43267eeb48a1a.IsNil() {
		return __obf_2c8e97779108ab17.EncodeNil()
	}
	__obf_a1f43267eeb48a1a = __obf_a1f43267eeb48a1a.Elem()
	if __obf_a1f43267eeb48a1a.Kind() == reflect.String {
		return __obf_2c8e97779108ab17.__obf_c00f3659b2fb4845(__obf_a1f43267eeb48a1a.String(), true)
	}
	return __obf_2c8e97779108ab17.EncodeValue(__obf_a1f43267eeb48a1a)
}

func __obf_825879324202941a(__obf_2c8e97779108ab17 *Encoder, __obf_a1f43267eeb48a1a reflect.Value) error {
	return __obf_2c8e97779108ab17.__obf_c00f3659b2fb4845(__obf_a1f43267eeb48a1a.String(), true)
}

func (__obf_2c8e97779108ab17 *Encoder) __obf_c00f3659b2fb4845(__obf_759f458bfa19abba string, __obf_0a0fa8c9dd348fed bool) error {
	// Interned string takes at least 3 bytes. Plain string 1 byte + string len.
	if __obf_82aee0190d456a88, __obf_81b19f2a6159ab89 := __obf_2c8e97779108ab17.__obf_57edbb6a6615f57c[__obf_759f458bfa19abba]; __obf_81b19f2a6159ab89 {
		return __obf_2c8e97779108ab17.__obf_fc4948a61270596f(__obf_82aee0190d456a88)
	}

	if __obf_0a0fa8c9dd348fed && len(__obf_759f458bfa19abba) >= __obf_fa31df83bfa41769 && len(__obf_2c8e97779108ab17.__obf_57edbb6a6615f57c) < __obf_7b33188291e2bf2f {
		if __obf_2c8e97779108ab17.__obf_57edbb6a6615f57c == nil {
			__obf_2c8e97779108ab17.__obf_57edbb6a6615f57c = make(map[string]int)
		}
		__obf_82aee0190d456a88 := len(__obf_2c8e97779108ab17.__obf_57edbb6a6615f57c)
		__obf_2c8e97779108ab17.__obf_57edbb6a6615f57c[__obf_759f458bfa19abba] = __obf_82aee0190d456a88
	}

	return __obf_2c8e97779108ab17.__obf_5d84555f6bd20e3c(__obf_759f458bfa19abba)
}

func (__obf_2c8e97779108ab17 *Encoder) __obf_fc4948a61270596f(__obf_82aee0190d456a88 int) error {
	if __obf_82aee0190d456a88 <= math.MaxUint8 {
		if __obf_4061ca0039150c39 := __obf_2c8e97779108ab17.__obf_c087409347698442(FixExt1); __obf_4061ca0039150c39 != nil {
			return __obf_4061ca0039150c39
		}
		return __obf_2c8e97779108ab17.__obf_2ebb4be6da23dcc7(byte(__obf_24fe567d301ad58a), uint8(__obf_82aee0190d456a88))
	}

	if __obf_82aee0190d456a88 <= math.MaxUint16 {
		if __obf_4061ca0039150c39 := __obf_2c8e97779108ab17.__obf_c087409347698442(FixExt2); __obf_4061ca0039150c39 != nil {
			return __obf_4061ca0039150c39
		}
		return __obf_2c8e97779108ab17.__obf_71321d3206264f3b(byte(__obf_24fe567d301ad58a), uint16(__obf_82aee0190d456a88))
	}

	if uint64(__obf_82aee0190d456a88) <= math.MaxUint32 {
		if __obf_4061ca0039150c39 := __obf_2c8e97779108ab17.__obf_c087409347698442(FixExt4); __obf_4061ca0039150c39 != nil {
			return __obf_4061ca0039150c39
		}
		return __obf_2c8e97779108ab17.__obf_a40df49dbc3bfc37(byte(__obf_24fe567d301ad58a), uint32(__obf_82aee0190d456a88))
	}

	return fmt.Errorf("msgpack: interned string index=%d is too large", __obf_82aee0190d456a88)
}

//------------------------------------------------------------------------------

func __obf_17a6f68227288f35(__obf_613397fefdec0ed0 *Decoder, __obf_a1f43267eeb48a1a reflect.Value) error {
	__obf_759f458bfa19abba, __obf_4061ca0039150c39 := __obf_613397fefdec0ed0.__obf_2fcfd6a9223bea0f(true)
	if __obf_4061ca0039150c39 == nil {
		__obf_a1f43267eeb48a1a.
			Set(reflect.ValueOf(__obf_759f458bfa19abba))
		return nil
	}
	if _, __obf_81b19f2a6159ab89 := __obf_4061ca0039150c39.(__obf_31d9eceb380e57ed); !__obf_81b19f2a6159ab89 {
		return __obf_4061ca0039150c39
	}

	if __obf_4061ca0039150c39 := __obf_613397fefdec0ed0.__obf_759f458bfa19abba.UnreadByte(); __obf_4061ca0039150c39 != nil {
		return __obf_4061ca0039150c39
	}
	return __obf_3b5214be8509686e(__obf_613397fefdec0ed0, __obf_a1f43267eeb48a1a)
}

func __obf_28e12bbdc2b0bf28(__obf_613397fefdec0ed0 *Decoder, __obf_a1f43267eeb48a1a reflect.Value) error {
	__obf_759f458bfa19abba, __obf_4061ca0039150c39 := __obf_613397fefdec0ed0.__obf_2fcfd6a9223bea0f(true)
	if __obf_4061ca0039150c39 != nil {
		return __obf_4061ca0039150c39
	}
	__obf_a1f43267eeb48a1a.
		SetString(__obf_759f458bfa19abba)
	return nil
}

func (__obf_613397fefdec0ed0 *Decoder) __obf_2fcfd6a9223bea0f(__obf_0a0fa8c9dd348fed bool) (string, error) {
	__obf_6dbe86b3d9d9b037, __obf_4061ca0039150c39 := __obf_613397fefdec0ed0.__obf_46b95b57ed5617b5()
	if __obf_4061ca0039150c39 != nil {
		return "", __obf_4061ca0039150c39
	}

	if IsFixedString(__obf_6dbe86b3d9d9b037) {
		__obf_99a74e41c9c640c0 := int(__obf_6dbe86b3d9d9b037 & FixedStrMask)
		return __obf_613397fefdec0ed0.__obf_c4d410fee2999cc5(__obf_99a74e41c9c640c0, __obf_0a0fa8c9dd348fed)
	}

	switch __obf_6dbe86b3d9d9b037 {
	case Nil:
		return "", nil
	case FixExt1, FixExt2, FixExt4:
		__obf_56090487efb2a9c7, __obf_f975a2b18cb9c21c, __obf_4061ca0039150c39 := __obf_613397fefdec0ed0.__obf_2227ad80d0435996(__obf_6dbe86b3d9d9b037)
		if __obf_4061ca0039150c39 != nil {
			return "", __obf_4061ca0039150c39
		}
		if __obf_56090487efb2a9c7 != __obf_24fe567d301ad58a {
			__obf_4061ca0039150c39 := fmt.Errorf("msgpack: got ext type=%d, wanted %d", __obf_56090487efb2a9c7, __obf_24fe567d301ad58a)
			return "", __obf_4061ca0039150c39
		}
		__obf_82aee0190d456a88, __obf_4061ca0039150c39 := __obf_613397fefdec0ed0.__obf_404e8bd8820334d0(__obf_f975a2b18cb9c21c)
		if __obf_4061ca0039150c39 != nil {
			return "", __obf_4061ca0039150c39
		}

		return __obf_613397fefdec0ed0.__obf_0818908d2fd05af4(__obf_82aee0190d456a88)
	case Str8, Bin8:
		__obf_99a74e41c9c640c0, __obf_4061ca0039150c39 := __obf_613397fefdec0ed0.uint8()
		if __obf_4061ca0039150c39 != nil {
			return "", __obf_4061ca0039150c39
		}
		return __obf_613397fefdec0ed0.__obf_c4d410fee2999cc5(int(__obf_99a74e41c9c640c0), __obf_0a0fa8c9dd348fed)
	case Str16, Bin16:
		__obf_99a74e41c9c640c0, __obf_4061ca0039150c39 := __obf_613397fefdec0ed0.uint16()
		if __obf_4061ca0039150c39 != nil {
			return "", __obf_4061ca0039150c39
		}
		return __obf_613397fefdec0ed0.__obf_c4d410fee2999cc5(int(__obf_99a74e41c9c640c0), __obf_0a0fa8c9dd348fed)
	case Str32, Bin32:
		__obf_99a74e41c9c640c0, __obf_4061ca0039150c39 := __obf_613397fefdec0ed0.uint32()
		if __obf_4061ca0039150c39 != nil {
			return "", __obf_4061ca0039150c39
		}
		return __obf_613397fefdec0ed0.__obf_c4d410fee2999cc5(int(__obf_99a74e41c9c640c0), __obf_0a0fa8c9dd348fed)
	}

	return "", __obf_31d9eceb380e57ed{__obf_987b95dd4dcfc308: __obf_6dbe86b3d9d9b037, __obf_1ec1fd8709fc0afd: "interned string"}
}

func (__obf_613397fefdec0ed0 *Decoder) __obf_404e8bd8820334d0(__obf_f975a2b18cb9c21c int) (int, error) {
	switch __obf_f975a2b18cb9c21c {
	case 1:
		__obf_99a74e41c9c640c0, __obf_4061ca0039150c39 := __obf_613397fefdec0ed0.uint8()
		if __obf_4061ca0039150c39 != nil {
			return 0, __obf_4061ca0039150c39
		}
		return int(__obf_99a74e41c9c640c0), nil
	case 2:
		__obf_99a74e41c9c640c0, __obf_4061ca0039150c39 := __obf_613397fefdec0ed0.uint16()
		if __obf_4061ca0039150c39 != nil {
			return 0, __obf_4061ca0039150c39
		}
		return int(__obf_99a74e41c9c640c0), nil
	case 4:
		__obf_99a74e41c9c640c0, __obf_4061ca0039150c39 := __obf_613397fefdec0ed0.uint32()
		if __obf_4061ca0039150c39 != nil {
			return 0, __obf_4061ca0039150c39
		}
		return int(__obf_99a74e41c9c640c0), nil
	}
	__obf_4061ca0039150c39 := fmt.Errorf("msgpack: unsupported ext len=%d decoding interned string", __obf_f975a2b18cb9c21c)
	return 0, __obf_4061ca0039150c39
}

func (__obf_613397fefdec0ed0 *Decoder) __obf_0818908d2fd05af4(__obf_82aee0190d456a88 int) (string, error) {
	if __obf_82aee0190d456a88 >= len(__obf_613397fefdec0ed0.__obf_57edbb6a6615f57c) {
		__obf_4061ca0039150c39 := fmt.Errorf("msgpack: interned string at index=%d does not exist", __obf_82aee0190d456a88)
		return "", __obf_4061ca0039150c39
	}
	return __obf_613397fefdec0ed0.__obf_57edbb6a6615f57c[__obf_82aee0190d456a88], nil
}

func (__obf_613397fefdec0ed0 *Decoder) __obf_c4d410fee2999cc5(__obf_99a74e41c9c640c0 int, __obf_0a0fa8c9dd348fed bool) (string, error) {
	if __obf_99a74e41c9c640c0 <= 0 {
		return "", nil
	}
	__obf_759f458bfa19abba, __obf_4061ca0039150c39 := __obf_613397fefdec0ed0.__obf_d6534b35621d5ef7(__obf_99a74e41c9c640c0)
	if __obf_4061ca0039150c39 != nil {
		return "", __obf_4061ca0039150c39
	}

	if __obf_0a0fa8c9dd348fed && len(__obf_759f458bfa19abba) >= __obf_fa31df83bfa41769 && len(__obf_613397fefdec0ed0.__obf_57edbb6a6615f57c) < __obf_7b33188291e2bf2f {
		__obf_613397fefdec0ed0.__obf_57edbb6a6615f57c = append(__obf_613397fefdec0ed0.__obf_57edbb6a6615f57c, __obf_759f458bfa19abba)
	}

	return __obf_759f458bfa19abba, nil
}
