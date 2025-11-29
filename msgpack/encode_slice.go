package __obf_a4aad98aaf3178f4

import (
	"math"
	"reflect"
)

var __obf_e346e0bc14858511 = reflect.TypeOf(([]string)(nil))

func __obf_52eb49ecc239146d(__obf_2c8e97779108ab17 *Encoder, __obf_a1f43267eeb48a1a reflect.Value) error {
	return __obf_2c8e97779108ab17.EncodeString(__obf_a1f43267eeb48a1a.String())
}

func __obf_7b11a1a1aaa56165(__obf_2c8e97779108ab17 *Encoder, __obf_a1f43267eeb48a1a reflect.Value) error {
	return __obf_2c8e97779108ab17.EncodeBytes(__obf_a1f43267eeb48a1a.Bytes())
}

func __obf_7f7b053e6838c23d(__obf_2c8e97779108ab17 *Encoder, __obf_a1f43267eeb48a1a reflect.Value) error {
	if __obf_4061ca0039150c39 := __obf_2c8e97779108ab17.EncodeBytesLen(__obf_a1f43267eeb48a1a.Len()); __obf_4061ca0039150c39 != nil {
		return __obf_4061ca0039150c39
	}

	if __obf_a1f43267eeb48a1a.CanAddr() {
		__obf_f57209cfda8e17cf := __obf_a1f43267eeb48a1a.Slice(0, __obf_a1f43267eeb48a1a.Len()).Bytes()
		return __obf_2c8e97779108ab17.__obf_ab9f7a04111e0451(__obf_f57209cfda8e17cf)
	}
	__obf_2c8e97779108ab17.__obf_f5d391a424aa5a3a = __obf_7f5e839d6e2d001d(__obf_2c8e97779108ab17.__obf_f5d391a424aa5a3a, __obf_a1f43267eeb48a1a.Len())
	reflect.Copy(reflect.ValueOf(__obf_2c8e97779108ab17.__obf_f5d391a424aa5a3a), __obf_a1f43267eeb48a1a)
	return __obf_2c8e97779108ab17.__obf_ab9f7a04111e0451(__obf_2c8e97779108ab17.__obf_f5d391a424aa5a3a)
}

func __obf_7f5e839d6e2d001d(__obf_f57209cfda8e17cf []byte, __obf_99a74e41c9c640c0 int) []byte {
	if cap(__obf_f57209cfda8e17cf) >= __obf_99a74e41c9c640c0 {
		return __obf_f57209cfda8e17cf[:__obf_99a74e41c9c640c0]
	}
	__obf_f57209cfda8e17cf = __obf_f57209cfda8e17cf[:cap(__obf_f57209cfda8e17cf)]
	__obf_f57209cfda8e17cf = append(__obf_f57209cfda8e17cf, make([]byte, __obf_99a74e41c9c640c0-len(__obf_f57209cfda8e17cf))...)
	return __obf_f57209cfda8e17cf
}

func (__obf_2c8e97779108ab17 *Encoder) EncodeBytesLen(__obf_4431cbe3c2f63394 int) error {
	if __obf_4431cbe3c2f63394 < 256 {
		return __obf_2c8e97779108ab17.__obf_2ebb4be6da23dcc7(Bin8, uint8(__obf_4431cbe3c2f63394))
	}
	if __obf_4431cbe3c2f63394 <= math.MaxUint16 {
		return __obf_2c8e97779108ab17.__obf_71321d3206264f3b(Bin16, uint16(__obf_4431cbe3c2f63394))
	}
	return __obf_2c8e97779108ab17.__obf_a40df49dbc3bfc37(Bin32, uint32(__obf_4431cbe3c2f63394))
}

func (__obf_2c8e97779108ab17 *Encoder) __obf_f53cf6b2bc94676e(__obf_4431cbe3c2f63394 int) error {
	if __obf_4431cbe3c2f63394 < 32 {
		return __obf_2c8e97779108ab17.__obf_c087409347698442(FixedStrLow | byte(__obf_4431cbe3c2f63394))
	}
	if __obf_4431cbe3c2f63394 < 256 {
		return __obf_2c8e97779108ab17.__obf_2ebb4be6da23dcc7(Str8, uint8(__obf_4431cbe3c2f63394))
	}
	if __obf_4431cbe3c2f63394 <= math.MaxUint16 {
		return __obf_2c8e97779108ab17.__obf_71321d3206264f3b(Str16, uint16(__obf_4431cbe3c2f63394))
	}
	return __obf_2c8e97779108ab17.__obf_a40df49dbc3bfc37(Str32, uint32(__obf_4431cbe3c2f63394))
}

func (__obf_2c8e97779108ab17 *Encoder) EncodeString(__obf_a1f43267eeb48a1a string) error {
	if __obf_0a0fa8c9dd348fed := __obf_2c8e97779108ab17.__obf_a8400e70a712e9fa&__obf_ca379a8882c3af10 != 0; __obf_0a0fa8c9dd348fed || len(__obf_2c8e97779108ab17.__obf_57edbb6a6615f57c) > 0 {
		return __obf_2c8e97779108ab17.__obf_c00f3659b2fb4845(__obf_a1f43267eeb48a1a, __obf_0a0fa8c9dd348fed)
	}
	return __obf_2c8e97779108ab17.__obf_5d84555f6bd20e3c(__obf_a1f43267eeb48a1a)
}

func (__obf_2c8e97779108ab17 *Encoder) __obf_5d84555f6bd20e3c(__obf_a1f43267eeb48a1a string) error {
	if __obf_4061ca0039150c39 := __obf_2c8e97779108ab17.__obf_f53cf6b2bc94676e(len(__obf_a1f43267eeb48a1a)); __obf_4061ca0039150c39 != nil {
		return __obf_4061ca0039150c39
	}
	return __obf_2c8e97779108ab17.__obf_176357c847e61809(__obf_a1f43267eeb48a1a)
}

func (__obf_2c8e97779108ab17 *Encoder) EncodeBytes(__obf_a1f43267eeb48a1a []byte) error {
	if __obf_a1f43267eeb48a1a == nil {
		return __obf_2c8e97779108ab17.EncodeNil()
	}
	if __obf_4061ca0039150c39 := __obf_2c8e97779108ab17.EncodeBytesLen(len(__obf_a1f43267eeb48a1a)); __obf_4061ca0039150c39 != nil {
		return __obf_4061ca0039150c39
	}
	return __obf_2c8e97779108ab17.__obf_ab9f7a04111e0451(__obf_a1f43267eeb48a1a)
}

func (__obf_2c8e97779108ab17 *Encoder) EncodeArrayLen(__obf_4431cbe3c2f63394 int) error {
	if __obf_4431cbe3c2f63394 < 16 {
		return __obf_2c8e97779108ab17.__obf_c087409347698442(FixedArrayLow | byte(__obf_4431cbe3c2f63394))
	}
	if __obf_4431cbe3c2f63394 <= math.MaxUint16 {
		return __obf_2c8e97779108ab17.__obf_71321d3206264f3b(Array16, uint16(__obf_4431cbe3c2f63394))
	}
	return __obf_2c8e97779108ab17.__obf_a40df49dbc3bfc37(Array32, uint32(__obf_4431cbe3c2f63394))
}

func __obf_7c02e0223b928d18(__obf_2c8e97779108ab17 *Encoder, __obf_a1f43267eeb48a1a reflect.Value) error {
	__obf_a7ace2ee7c6566ec := __obf_a1f43267eeb48a1a.Convert(__obf_e346e0bc14858511).Interface().([]string)
	return __obf_2c8e97779108ab17.__obf_c18f6fc1aa28e1ac(__obf_a7ace2ee7c6566ec)
}

func (__obf_2c8e97779108ab17 *Encoder) __obf_c18f6fc1aa28e1ac(__obf_759f458bfa19abba []string) error {
	if __obf_759f458bfa19abba == nil {
		return __obf_2c8e97779108ab17.EncodeNil()
	}
	if __obf_4061ca0039150c39 := __obf_2c8e97779108ab17.EncodeArrayLen(len(__obf_759f458bfa19abba)); __obf_4061ca0039150c39 != nil {
		return __obf_4061ca0039150c39
	}
	for _, __obf_a1f43267eeb48a1a := range __obf_759f458bfa19abba {
		if __obf_4061ca0039150c39 := __obf_2c8e97779108ab17.EncodeString(__obf_a1f43267eeb48a1a); __obf_4061ca0039150c39 != nil {
			return __obf_4061ca0039150c39
		}
	}
	return nil
}

func __obf_6abbdd3481da14fa(__obf_2c8e97779108ab17 *Encoder, __obf_a1f43267eeb48a1a reflect.Value) error {
	if __obf_a1f43267eeb48a1a.IsNil() {
		return __obf_2c8e97779108ab17.EncodeNil()
	}
	return __obf_5711cd74f1e4cfdc(__obf_2c8e97779108ab17, __obf_a1f43267eeb48a1a)
}

func __obf_5711cd74f1e4cfdc(__obf_2c8e97779108ab17 *Encoder, __obf_a1f43267eeb48a1a reflect.Value) error {
	__obf_4431cbe3c2f63394 := __obf_a1f43267eeb48a1a.Len()
	if __obf_4061ca0039150c39 := __obf_2c8e97779108ab17.EncodeArrayLen(__obf_4431cbe3c2f63394); __obf_4061ca0039150c39 != nil {
		return __obf_4061ca0039150c39
	}
	for __obf_097d8b6061c9592b := 0; __obf_097d8b6061c9592b < __obf_4431cbe3c2f63394; __obf_097d8b6061c9592b++ {
		if __obf_4061ca0039150c39 := __obf_2c8e97779108ab17.EncodeValue(__obf_a1f43267eeb48a1a.Index(__obf_097d8b6061c9592b)); __obf_4061ca0039150c39 != nil {
			return __obf_4061ca0039150c39
		}
	}
	return nil
}
