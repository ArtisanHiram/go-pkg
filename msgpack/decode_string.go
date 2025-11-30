package __obf_de86cdc8ae98b45b

import (
	"fmt"
	"reflect"
)

func (__obf_dcaad1165bb07af9 *Decoder) __obf_869c18dc6d397629(__obf_ecf6d2d3253a058d byte) (int, error) {
	if __obf_ecf6d2d3253a058d == Nil {
		return -1, nil
	}

	if IsFixedString(__obf_ecf6d2d3253a058d) {
		return int(__obf_ecf6d2d3253a058d & FixedStrMask), nil
	}

	switch __obf_ecf6d2d3253a058d {
	case Str8, Bin8:
		__obf_2b0247e819b1bf4a, __obf_0feb3528b7b709ec := __obf_dcaad1165bb07af9.uint8()
		return int(__obf_2b0247e819b1bf4a), __obf_0feb3528b7b709ec
	case Str16, Bin16:
		__obf_2b0247e819b1bf4a, __obf_0feb3528b7b709ec := __obf_dcaad1165bb07af9.uint16()
		return int(__obf_2b0247e819b1bf4a), __obf_0feb3528b7b709ec
	case Str32, Bin32:
		__obf_2b0247e819b1bf4a, __obf_0feb3528b7b709ec := __obf_dcaad1165bb07af9.uint32()
		return int(__obf_2b0247e819b1bf4a), __obf_0feb3528b7b709ec
	}

	return 0, fmt.Errorf("msgpack: invalid code=%x decoding string/bytes length", __obf_ecf6d2d3253a058d)
}

func (__obf_dcaad1165bb07af9 *Decoder) DecodeString() (string, error) {
	if __obf_65f94b9eb2a3c80f := __obf_dcaad1165bb07af9.__obf_a15a642720e39c3e&__obf_1fcd090a6da01656 != 0; __obf_65f94b9eb2a3c80f || len(__obf_dcaad1165bb07af9.__obf_25128eea0d506b65) > 0 {
		return __obf_dcaad1165bb07af9.__obf_76b76bd32a6c2a62(__obf_65f94b9eb2a3c80f)
	}
	__obf_ecf6d2d3253a058d, __obf_0feb3528b7b709ec := __obf_dcaad1165bb07af9.__obf_74ac05d9a9da01f4()
	if __obf_0feb3528b7b709ec != nil {
		return "", __obf_0feb3528b7b709ec
	}
	return __obf_dcaad1165bb07af9.string(__obf_ecf6d2d3253a058d)
}

func (__obf_dcaad1165bb07af9 *Decoder) string(__obf_ecf6d2d3253a058d byte) (string, error) {
	__obf_2b0247e819b1bf4a, __obf_0feb3528b7b709ec := __obf_dcaad1165bb07af9.__obf_869c18dc6d397629(__obf_ecf6d2d3253a058d)
	if __obf_0feb3528b7b709ec != nil {
		return "", __obf_0feb3528b7b709ec
	}
	return __obf_dcaad1165bb07af9.__obf_281f9be3791bf742(__obf_2b0247e819b1bf4a)
}

func (__obf_dcaad1165bb07af9 *Decoder) __obf_281f9be3791bf742(__obf_2b0247e819b1bf4a int) (string, error) {
	if __obf_2b0247e819b1bf4a <= 0 {
		return "", nil
	}
	__obf_a7fd7acf2bd4435f, __obf_0feb3528b7b709ec := __obf_dcaad1165bb07af9.__obf_8f5a813341411779(__obf_2b0247e819b1bf4a)
	return string(__obf_a7fd7acf2bd4435f), __obf_0feb3528b7b709ec
}

func __obf_47189e506d9624a7(__obf_dcaad1165bb07af9 *Decoder, __obf_17732590722140e7 reflect.Value) error {
	__obf_a93d004caac96500, __obf_0feb3528b7b709ec := __obf_dcaad1165bb07af9.DecodeString()
	if __obf_0feb3528b7b709ec != nil {
		return __obf_0feb3528b7b709ec
	}
	__obf_17732590722140e7.
		SetString(__obf_a93d004caac96500)
	return nil
}

func (__obf_dcaad1165bb07af9 *Decoder) __obf_0605127b8a406c17() (string, error) {
	if __obf_65f94b9eb2a3c80f := __obf_dcaad1165bb07af9.__obf_a15a642720e39c3e&__obf_1fcd090a6da01656 != 0; __obf_65f94b9eb2a3c80f || len(__obf_dcaad1165bb07af9.__obf_25128eea0d506b65) > 0 {
		return __obf_dcaad1165bb07af9.__obf_76b76bd32a6c2a62(__obf_65f94b9eb2a3c80f)
	}
	__obf_ecf6d2d3253a058d, __obf_0feb3528b7b709ec := __obf_dcaad1165bb07af9.__obf_74ac05d9a9da01f4()
	if __obf_0feb3528b7b709ec != nil {
		return "", __obf_0feb3528b7b709ec
	}
	__obf_2b0247e819b1bf4a, __obf_0feb3528b7b709ec := __obf_dcaad1165bb07af9.__obf_869c18dc6d397629(__obf_ecf6d2d3253a058d)
	if __obf_0feb3528b7b709ec != nil {
		return "", __obf_0feb3528b7b709ec
	}
	if __obf_2b0247e819b1bf4a == -1 {
		return "", nil
	}
	__obf_a7fd7acf2bd4435f, __obf_0feb3528b7b709ec := __obf_dcaad1165bb07af9.__obf_8f5a813341411779(__obf_2b0247e819b1bf4a)
	if __obf_0feb3528b7b709ec != nil {
		return "", __obf_0feb3528b7b709ec
	}

	return __obf_16417a69f4ce2476(__obf_a7fd7acf2bd4435f), nil
}

func (__obf_dcaad1165bb07af9 *Decoder) __obf_a7cf3730c2f8bbd8(__obf_01f0154f1fce8e5a *[]byte) error {
	__obf_ecf6d2d3253a058d, __obf_0feb3528b7b709ec := __obf_dcaad1165bb07af9.__obf_74ac05d9a9da01f4()
	if __obf_0feb3528b7b709ec != nil {
		return __obf_0feb3528b7b709ec
	}
	return __obf_dcaad1165bb07af9.__obf_8dc07bf67c6e9356(__obf_ecf6d2d3253a058d, __obf_01f0154f1fce8e5a)
}

func (__obf_dcaad1165bb07af9 *Decoder) __obf_8dc07bf67c6e9356(__obf_ecf6d2d3253a058d byte, __obf_01f0154f1fce8e5a *[]byte) error {
	__obf_2b0247e819b1bf4a, __obf_0feb3528b7b709ec := __obf_dcaad1165bb07af9.__obf_869c18dc6d397629(__obf_ecf6d2d3253a058d)
	if __obf_0feb3528b7b709ec != nil {
		return __obf_0feb3528b7b709ec
	}
	if __obf_2b0247e819b1bf4a == -1 {
		*__obf_01f0154f1fce8e5a = nil
		return nil
	}

	*__obf_01f0154f1fce8e5a, __obf_0feb3528b7b709ec = __obf_8f5a813341411779(__obf_dcaad1165bb07af9.__obf_4609cadeaf50917c, *__obf_01f0154f1fce8e5a, __obf_2b0247e819b1bf4a)
	return __obf_0feb3528b7b709ec
}

func (__obf_dcaad1165bb07af9 *Decoder) __obf_d22198f61bd105ee(__obf_ecf6d2d3253a058d byte) error {
	__obf_2b0247e819b1bf4a, __obf_0feb3528b7b709ec := __obf_dcaad1165bb07af9.__obf_869c18dc6d397629(__obf_ecf6d2d3253a058d)
	if __obf_0feb3528b7b709ec != nil {
		return __obf_0feb3528b7b709ec
	}
	if __obf_2b0247e819b1bf4a <= 0 {
		return nil
	}
	return __obf_dcaad1165bb07af9.__obf_6f540045c1336cac(__obf_2b0247e819b1bf4a)
}

func __obf_9dd8ad0f86a4d0ca(__obf_dcaad1165bb07af9 *Decoder, __obf_17732590722140e7 reflect.Value) error {
	__obf_ecf6d2d3253a058d, __obf_0feb3528b7b709ec := __obf_dcaad1165bb07af9.__obf_74ac05d9a9da01f4()
	if __obf_0feb3528b7b709ec != nil {
		return __obf_0feb3528b7b709ec
	}
	__obf_2b0247e819b1bf4a, __obf_0feb3528b7b709ec := __obf_dcaad1165bb07af9.__obf_869c18dc6d397629(__obf_ecf6d2d3253a058d)
	if __obf_0feb3528b7b709ec != nil {
		return __obf_0feb3528b7b709ec
	}
	if __obf_2b0247e819b1bf4a == -1 {
		return nil
	}
	if __obf_2b0247e819b1bf4a > __obf_17732590722140e7.Len() {
		return fmt.Errorf("%s len is %d, but msgpack has %d elements", __obf_17732590722140e7.Type(), __obf_17732590722140e7.Len(), __obf_2b0247e819b1bf4a)
	}
	__obf_a7fd7acf2bd4435f := __obf_17732590722140e7.Slice(0, __obf_2b0247e819b1bf4a).Bytes()
	return __obf_dcaad1165bb07af9.__obf_73059ce4dea16cf3(__obf_a7fd7acf2bd4435f)
}
