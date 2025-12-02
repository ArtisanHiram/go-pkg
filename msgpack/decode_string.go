package __obf_a935eb7f548271a4

import (
	"fmt"
	"reflect"
)

func (__obf_a21885da2425f2b2 *Decoder) __obf_c0a95532c414ce37(__obf_f5df560f4d67421b byte) (int, error) {
	if __obf_f5df560f4d67421b == Nil {
		return -1, nil
	}

	if IsFixedString(__obf_f5df560f4d67421b) {
		return int(__obf_f5df560f4d67421b & FixedStrMask), nil
	}

	switch __obf_f5df560f4d67421b {
	case Str8, Bin8:
		__obf_326af9bd942662ac, __obf_4d327e1cd40c2e21 := __obf_a21885da2425f2b2.uint8()
		return int(__obf_326af9bd942662ac), __obf_4d327e1cd40c2e21
	case Str16, Bin16:
		__obf_326af9bd942662ac, __obf_4d327e1cd40c2e21 := __obf_a21885da2425f2b2.uint16()
		return int(__obf_326af9bd942662ac), __obf_4d327e1cd40c2e21
	case Str32, Bin32:
		__obf_326af9bd942662ac, __obf_4d327e1cd40c2e21 := __obf_a21885da2425f2b2.uint32()
		return int(__obf_326af9bd942662ac), __obf_4d327e1cd40c2e21
	}

	return 0, fmt.Errorf("msgpack: invalid code=%x decoding string/bytes length", __obf_f5df560f4d67421b)
}

func (__obf_a21885da2425f2b2 *Decoder) DecodeString() (string, error) {
	if __obf_65a759dbf3ace040 := __obf_a21885da2425f2b2.__obf_d617f3769ce16c47&__obf_bf24fd2d64b91255 != 0; __obf_65a759dbf3ace040 || len(__obf_a21885da2425f2b2.__obf_2cdd7fad308aef1d) > 0 {
		return __obf_a21885da2425f2b2.__obf_f43916d12d5b1cf3(__obf_65a759dbf3ace040)
	}
	__obf_f5df560f4d67421b, __obf_4d327e1cd40c2e21 := __obf_a21885da2425f2b2.__obf_61e5879fcf7b35ce()
	if __obf_4d327e1cd40c2e21 != nil {
		return "", __obf_4d327e1cd40c2e21
	}
	return __obf_a21885da2425f2b2.string(__obf_f5df560f4d67421b)
}

func (__obf_a21885da2425f2b2 *Decoder) string(__obf_f5df560f4d67421b byte) (string, error) {
	__obf_326af9bd942662ac, __obf_4d327e1cd40c2e21 := __obf_a21885da2425f2b2.__obf_c0a95532c414ce37(__obf_f5df560f4d67421b)
	if __obf_4d327e1cd40c2e21 != nil {
		return "", __obf_4d327e1cd40c2e21
	}
	return __obf_a21885da2425f2b2.__obf_bb558e6f117ed933(__obf_326af9bd942662ac)
}

func (__obf_a21885da2425f2b2 *Decoder) __obf_bb558e6f117ed933(__obf_326af9bd942662ac int) (string, error) {
	if __obf_326af9bd942662ac <= 0 {
		return "", nil
	}
	__obf_f2ca794293605b73, __obf_4d327e1cd40c2e21 := __obf_a21885da2425f2b2.__obf_595cbf51b6653ebf(__obf_326af9bd942662ac)
	return string(__obf_f2ca794293605b73), __obf_4d327e1cd40c2e21
}

func __obf_d48bbb6e52137026(__obf_a21885da2425f2b2 *Decoder, __obf_6d570581f4b60dbc reflect.Value) error {
	__obf_b62c60fba2fd9788, __obf_4d327e1cd40c2e21 := __obf_a21885da2425f2b2.DecodeString()
	if __obf_4d327e1cd40c2e21 != nil {
		return __obf_4d327e1cd40c2e21
	}
	__obf_6d570581f4b60dbc.
		SetString(__obf_b62c60fba2fd9788)
	return nil
}

func (__obf_a21885da2425f2b2 *Decoder) __obf_4e3696e2f7b41e6e() (string, error) {
	if __obf_65a759dbf3ace040 := __obf_a21885da2425f2b2.__obf_d617f3769ce16c47&__obf_bf24fd2d64b91255 != 0; __obf_65a759dbf3ace040 || len(__obf_a21885da2425f2b2.__obf_2cdd7fad308aef1d) > 0 {
		return __obf_a21885da2425f2b2.__obf_f43916d12d5b1cf3(__obf_65a759dbf3ace040)
	}
	__obf_f5df560f4d67421b, __obf_4d327e1cd40c2e21 := __obf_a21885da2425f2b2.__obf_61e5879fcf7b35ce()
	if __obf_4d327e1cd40c2e21 != nil {
		return "", __obf_4d327e1cd40c2e21
	}
	__obf_326af9bd942662ac, __obf_4d327e1cd40c2e21 := __obf_a21885da2425f2b2.__obf_c0a95532c414ce37(__obf_f5df560f4d67421b)
	if __obf_4d327e1cd40c2e21 != nil {
		return "", __obf_4d327e1cd40c2e21
	}
	if __obf_326af9bd942662ac == -1 {
		return "", nil
	}
	__obf_f2ca794293605b73, __obf_4d327e1cd40c2e21 := __obf_a21885da2425f2b2.__obf_595cbf51b6653ebf(__obf_326af9bd942662ac)
	if __obf_4d327e1cd40c2e21 != nil {
		return "", __obf_4d327e1cd40c2e21
	}

	return __obf_ae16b98e7d90edc4(__obf_f2ca794293605b73), nil
}

func (__obf_a21885da2425f2b2 *Decoder) __obf_24d3dffd03973291(__obf_0d8a994785cda6df *[]byte) error {
	__obf_f5df560f4d67421b, __obf_4d327e1cd40c2e21 := __obf_a21885da2425f2b2.__obf_61e5879fcf7b35ce()
	if __obf_4d327e1cd40c2e21 != nil {
		return __obf_4d327e1cd40c2e21
	}
	return __obf_a21885da2425f2b2.__obf_0ca5adbd48337aba(__obf_f5df560f4d67421b, __obf_0d8a994785cda6df)
}

func (__obf_a21885da2425f2b2 *Decoder) __obf_0ca5adbd48337aba(__obf_f5df560f4d67421b byte, __obf_0d8a994785cda6df *[]byte) error {
	__obf_326af9bd942662ac, __obf_4d327e1cd40c2e21 := __obf_a21885da2425f2b2.__obf_c0a95532c414ce37(__obf_f5df560f4d67421b)
	if __obf_4d327e1cd40c2e21 != nil {
		return __obf_4d327e1cd40c2e21
	}
	if __obf_326af9bd942662ac == -1 {
		*__obf_0d8a994785cda6df = nil
		return nil
	}

	*__obf_0d8a994785cda6df, __obf_4d327e1cd40c2e21 = __obf_595cbf51b6653ebf(__obf_a21885da2425f2b2.__obf_4f94e5d98c7e4bb3, *__obf_0d8a994785cda6df, __obf_326af9bd942662ac)
	return __obf_4d327e1cd40c2e21
}

func (__obf_a21885da2425f2b2 *Decoder) __obf_4b7348bc9d246008(__obf_f5df560f4d67421b byte) error {
	__obf_326af9bd942662ac, __obf_4d327e1cd40c2e21 := __obf_a21885da2425f2b2.__obf_c0a95532c414ce37(__obf_f5df560f4d67421b)
	if __obf_4d327e1cd40c2e21 != nil {
		return __obf_4d327e1cd40c2e21
	}
	if __obf_326af9bd942662ac <= 0 {
		return nil
	}
	return __obf_a21885da2425f2b2.__obf_b2541a0cb78c8e1f(__obf_326af9bd942662ac)
}

func __obf_499a52ef00987055(__obf_a21885da2425f2b2 *Decoder, __obf_6d570581f4b60dbc reflect.Value) error {
	__obf_f5df560f4d67421b, __obf_4d327e1cd40c2e21 := __obf_a21885da2425f2b2.__obf_61e5879fcf7b35ce()
	if __obf_4d327e1cd40c2e21 != nil {
		return __obf_4d327e1cd40c2e21
	}
	__obf_326af9bd942662ac, __obf_4d327e1cd40c2e21 := __obf_a21885da2425f2b2.__obf_c0a95532c414ce37(__obf_f5df560f4d67421b)
	if __obf_4d327e1cd40c2e21 != nil {
		return __obf_4d327e1cd40c2e21
	}
	if __obf_326af9bd942662ac == -1 {
		return nil
	}
	if __obf_326af9bd942662ac > __obf_6d570581f4b60dbc.Len() {
		return fmt.Errorf("%s len is %d, but msgpack has %d elements", __obf_6d570581f4b60dbc.Type(), __obf_6d570581f4b60dbc.Len(), __obf_326af9bd942662ac)
	}
	__obf_f2ca794293605b73 := __obf_6d570581f4b60dbc.Slice(0, __obf_326af9bd942662ac).Bytes()
	return __obf_a21885da2425f2b2.__obf_fd322376412b2809(__obf_f2ca794293605b73)
}
