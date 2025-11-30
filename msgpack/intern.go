package __obf_fd770cb9675903b0

import (
	"fmt"
	"math"
	"reflect"
)

const (
	__obf_98834cf68179cc6c = 3
	__obf_f305034286cabc10 = math.MaxUint16
)

var __obf_cd59f0eb0a2583f4 = int8(math.MinInt8)

func init() {
	__obf_8adb5c6b5b689f26[__obf_cd59f0eb0a2583f4] = &__obf_ae91b3f1af57dd05{
		Type:    __obf_e7498c925721ef3a,
		Decoder: __obf_717a84f71bdd880b,
	}
}

func __obf_717a84f71bdd880b(__obf_5d389b660eefb08c *Decoder, __obf_f328a048f76b7256 reflect.Value, __obf_803bccbf05dad29c int) error {
	__obf_261dffcd6ce7f08f, __obf_45342a3a754d12f5 := __obf_5d389b660eefb08c.__obf_9013b391da819303(__obf_803bccbf05dad29c)
	if __obf_45342a3a754d12f5 != nil {
		return __obf_45342a3a754d12f5
	}
	__obf_fe1ace7a2eb8bf9f, __obf_45342a3a754d12f5 := __obf_5d389b660eefb08c.__obf_7b2d4bb21ae842cc(__obf_261dffcd6ce7f08f)
	if __obf_45342a3a754d12f5 != nil {
		return __obf_45342a3a754d12f5
	}
	__obf_f328a048f76b7256.
		SetString(__obf_fe1ace7a2eb8bf9f)
	return nil
}

//------------------------------------------------------------------------------

func __obf_db4fdb7a2953f589(__obf_e9038cda3b5cf711 *Encoder, __obf_f328a048f76b7256 reflect.Value) error {
	if __obf_f328a048f76b7256.IsNil() {
		return __obf_e9038cda3b5cf711.EncodeNil()
	}
	__obf_f328a048f76b7256 = __obf_f328a048f76b7256.Elem()
	if __obf_f328a048f76b7256.Kind() == reflect.String {
		return __obf_e9038cda3b5cf711.__obf_27c3beed6bd703be(__obf_f328a048f76b7256.String(), true)
	}
	return __obf_e9038cda3b5cf711.EncodeValue(__obf_f328a048f76b7256)
}

func __obf_b1eb723829a575f6(__obf_e9038cda3b5cf711 *Encoder, __obf_f328a048f76b7256 reflect.Value) error {
	return __obf_e9038cda3b5cf711.__obf_27c3beed6bd703be(__obf_f328a048f76b7256.String(), true)
}

func (__obf_e9038cda3b5cf711 *Encoder) __obf_27c3beed6bd703be(__obf_fe1ace7a2eb8bf9f string, __obf_4968ec805e7c5822 bool) error {
	// Interned string takes at least 3 bytes. Plain string 1 byte + string len.
	if __obf_261dffcd6ce7f08f, __obf_b00b3c6a10f90467 := __obf_e9038cda3b5cf711.__obf_ff96db22e12b6842[__obf_fe1ace7a2eb8bf9f]; __obf_b00b3c6a10f90467 {
		return __obf_e9038cda3b5cf711.__obf_123942e252b03eb6(__obf_261dffcd6ce7f08f)
	}

	if __obf_4968ec805e7c5822 && len(__obf_fe1ace7a2eb8bf9f) >= __obf_98834cf68179cc6c && len(__obf_e9038cda3b5cf711.__obf_ff96db22e12b6842) < __obf_f305034286cabc10 {
		if __obf_e9038cda3b5cf711.__obf_ff96db22e12b6842 == nil {
			__obf_e9038cda3b5cf711.__obf_ff96db22e12b6842 = make(map[string]int)
		}
		__obf_261dffcd6ce7f08f := len(__obf_e9038cda3b5cf711.__obf_ff96db22e12b6842)
		__obf_e9038cda3b5cf711.__obf_ff96db22e12b6842[__obf_fe1ace7a2eb8bf9f] = __obf_261dffcd6ce7f08f
	}

	return __obf_e9038cda3b5cf711.__obf_8d015d0182bfd5d0(__obf_fe1ace7a2eb8bf9f)
}

func (__obf_e9038cda3b5cf711 *Encoder) __obf_123942e252b03eb6(__obf_261dffcd6ce7f08f int) error {
	if __obf_261dffcd6ce7f08f <= math.MaxUint8 {
		if __obf_45342a3a754d12f5 := __obf_e9038cda3b5cf711.__obf_a72c49227d01f9dd(FixExt1); __obf_45342a3a754d12f5 != nil {
			return __obf_45342a3a754d12f5
		}
		return __obf_e9038cda3b5cf711.__obf_c58302e6ea7043c4(byte(__obf_cd59f0eb0a2583f4), uint8(__obf_261dffcd6ce7f08f))
	}

	if __obf_261dffcd6ce7f08f <= math.MaxUint16 {
		if __obf_45342a3a754d12f5 := __obf_e9038cda3b5cf711.__obf_a72c49227d01f9dd(FixExt2); __obf_45342a3a754d12f5 != nil {
			return __obf_45342a3a754d12f5
		}
		return __obf_e9038cda3b5cf711.__obf_a98b5ca0c9f03534(byte(__obf_cd59f0eb0a2583f4), uint16(__obf_261dffcd6ce7f08f))
	}

	if uint64(__obf_261dffcd6ce7f08f) <= math.MaxUint32 {
		if __obf_45342a3a754d12f5 := __obf_e9038cda3b5cf711.__obf_a72c49227d01f9dd(FixExt4); __obf_45342a3a754d12f5 != nil {
			return __obf_45342a3a754d12f5
		}
		return __obf_e9038cda3b5cf711.__obf_e0e3a3f6b942651e(byte(__obf_cd59f0eb0a2583f4), uint32(__obf_261dffcd6ce7f08f))
	}

	return fmt.Errorf("msgpack: interned string index=%d is too large", __obf_261dffcd6ce7f08f)
}

//------------------------------------------------------------------------------

func __obf_2ae309e0ccdf900e(__obf_5d389b660eefb08c *Decoder, __obf_f328a048f76b7256 reflect.Value) error {
	__obf_fe1ace7a2eb8bf9f, __obf_45342a3a754d12f5 := __obf_5d389b660eefb08c.__obf_06f8691b61014235(true)
	if __obf_45342a3a754d12f5 == nil {
		__obf_f328a048f76b7256.
			Set(reflect.ValueOf(__obf_fe1ace7a2eb8bf9f))
		return nil
	}
	if _, __obf_b00b3c6a10f90467 := __obf_45342a3a754d12f5.(__obf_b1f4502463d63ed9); !__obf_b00b3c6a10f90467 {
		return __obf_45342a3a754d12f5
	}

	if __obf_45342a3a754d12f5 := __obf_5d389b660eefb08c.__obf_fe1ace7a2eb8bf9f.UnreadByte(); __obf_45342a3a754d12f5 != nil {
		return __obf_45342a3a754d12f5
	}
	return __obf_604515803d5e5b1f(__obf_5d389b660eefb08c, __obf_f328a048f76b7256)
}

func __obf_7d869d41e2fef94d(__obf_5d389b660eefb08c *Decoder, __obf_f328a048f76b7256 reflect.Value) error {
	__obf_fe1ace7a2eb8bf9f, __obf_45342a3a754d12f5 := __obf_5d389b660eefb08c.__obf_06f8691b61014235(true)
	if __obf_45342a3a754d12f5 != nil {
		return __obf_45342a3a754d12f5
	}
	__obf_f328a048f76b7256.
		SetString(__obf_fe1ace7a2eb8bf9f)
	return nil
}

func (__obf_5d389b660eefb08c *Decoder) __obf_06f8691b61014235(__obf_4968ec805e7c5822 bool) (string, error) {
	__obf_4148125b350dfea2, __obf_45342a3a754d12f5 := __obf_5d389b660eefb08c.__obf_4786096a37bd32ce()
	if __obf_45342a3a754d12f5 != nil {
		return "", __obf_45342a3a754d12f5
	}

	if IsFixedString(__obf_4148125b350dfea2) {
		__obf_d774e4844c74bfe9 := int(__obf_4148125b350dfea2 & FixedStrMask)
		return __obf_5d389b660eefb08c.__obf_437100740be0d32b(__obf_d774e4844c74bfe9, __obf_4968ec805e7c5822)
	}

	switch __obf_4148125b350dfea2 {
	case Nil:
		return "", nil
	case FixExt1, FixExt2, FixExt4:
		__obf_e18f55f32a4d960d, __obf_803bccbf05dad29c, __obf_45342a3a754d12f5 := __obf_5d389b660eefb08c.__obf_051bbe645d4e1aa5(__obf_4148125b350dfea2)
		if __obf_45342a3a754d12f5 != nil {
			return "", __obf_45342a3a754d12f5
		}
		if __obf_e18f55f32a4d960d != __obf_cd59f0eb0a2583f4 {
			__obf_45342a3a754d12f5 := fmt.Errorf("msgpack: got ext type=%d, wanted %d", __obf_e18f55f32a4d960d, __obf_cd59f0eb0a2583f4)
			return "", __obf_45342a3a754d12f5
		}
		__obf_261dffcd6ce7f08f, __obf_45342a3a754d12f5 := __obf_5d389b660eefb08c.__obf_9013b391da819303(__obf_803bccbf05dad29c)
		if __obf_45342a3a754d12f5 != nil {
			return "", __obf_45342a3a754d12f5
		}

		return __obf_5d389b660eefb08c.__obf_7b2d4bb21ae842cc(__obf_261dffcd6ce7f08f)
	case Str8, Bin8:
		__obf_d774e4844c74bfe9, __obf_45342a3a754d12f5 := __obf_5d389b660eefb08c.uint8()
		if __obf_45342a3a754d12f5 != nil {
			return "", __obf_45342a3a754d12f5
		}
		return __obf_5d389b660eefb08c.__obf_437100740be0d32b(int(__obf_d774e4844c74bfe9), __obf_4968ec805e7c5822)
	case Str16, Bin16:
		__obf_d774e4844c74bfe9, __obf_45342a3a754d12f5 := __obf_5d389b660eefb08c.uint16()
		if __obf_45342a3a754d12f5 != nil {
			return "", __obf_45342a3a754d12f5
		}
		return __obf_5d389b660eefb08c.__obf_437100740be0d32b(int(__obf_d774e4844c74bfe9), __obf_4968ec805e7c5822)
	case Str32, Bin32:
		__obf_d774e4844c74bfe9, __obf_45342a3a754d12f5 := __obf_5d389b660eefb08c.uint32()
		if __obf_45342a3a754d12f5 != nil {
			return "", __obf_45342a3a754d12f5
		}
		return __obf_5d389b660eefb08c.__obf_437100740be0d32b(int(__obf_d774e4844c74bfe9), __obf_4968ec805e7c5822)
	}

	return "", __obf_b1f4502463d63ed9{__obf_cde82889ba8e4822: __obf_4148125b350dfea2, __obf_d6c122c727b9e88f: "interned string"}
}

func (__obf_5d389b660eefb08c *Decoder) __obf_9013b391da819303(__obf_803bccbf05dad29c int) (int, error) {
	switch __obf_803bccbf05dad29c {
	case 1:
		__obf_d774e4844c74bfe9, __obf_45342a3a754d12f5 := __obf_5d389b660eefb08c.uint8()
		if __obf_45342a3a754d12f5 != nil {
			return 0, __obf_45342a3a754d12f5
		}
		return int(__obf_d774e4844c74bfe9), nil
	case 2:
		__obf_d774e4844c74bfe9, __obf_45342a3a754d12f5 := __obf_5d389b660eefb08c.uint16()
		if __obf_45342a3a754d12f5 != nil {
			return 0, __obf_45342a3a754d12f5
		}
		return int(__obf_d774e4844c74bfe9), nil
	case 4:
		__obf_d774e4844c74bfe9, __obf_45342a3a754d12f5 := __obf_5d389b660eefb08c.uint32()
		if __obf_45342a3a754d12f5 != nil {
			return 0, __obf_45342a3a754d12f5
		}
		return int(__obf_d774e4844c74bfe9), nil
	}
	__obf_45342a3a754d12f5 := fmt.Errorf("msgpack: unsupported ext len=%d decoding interned string", __obf_803bccbf05dad29c)
	return 0, __obf_45342a3a754d12f5
}

func (__obf_5d389b660eefb08c *Decoder) __obf_7b2d4bb21ae842cc(__obf_261dffcd6ce7f08f int) (string, error) {
	if __obf_261dffcd6ce7f08f >= len(__obf_5d389b660eefb08c.__obf_ff96db22e12b6842) {
		__obf_45342a3a754d12f5 := fmt.Errorf("msgpack: interned string at index=%d does not exist", __obf_261dffcd6ce7f08f)
		return "", __obf_45342a3a754d12f5
	}
	return __obf_5d389b660eefb08c.__obf_ff96db22e12b6842[__obf_261dffcd6ce7f08f], nil
}

func (__obf_5d389b660eefb08c *Decoder) __obf_437100740be0d32b(__obf_d774e4844c74bfe9 int, __obf_4968ec805e7c5822 bool) (string, error) {
	if __obf_d774e4844c74bfe9 <= 0 {
		return "", nil
	}
	__obf_fe1ace7a2eb8bf9f, __obf_45342a3a754d12f5 := __obf_5d389b660eefb08c.__obf_6b53a4200bd06df8(__obf_d774e4844c74bfe9)
	if __obf_45342a3a754d12f5 != nil {
		return "", __obf_45342a3a754d12f5
	}

	if __obf_4968ec805e7c5822 && len(__obf_fe1ace7a2eb8bf9f) >= __obf_98834cf68179cc6c && len(__obf_5d389b660eefb08c.__obf_ff96db22e12b6842) < __obf_f305034286cabc10 {
		__obf_5d389b660eefb08c.__obf_ff96db22e12b6842 = append(__obf_5d389b660eefb08c.__obf_ff96db22e12b6842, __obf_fe1ace7a2eb8bf9f)
	}

	return __obf_fe1ace7a2eb8bf9f, nil
}
