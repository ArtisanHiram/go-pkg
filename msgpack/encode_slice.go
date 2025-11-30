package __obf_de86cdc8ae98b45b

import (
	"math"
	"reflect"
)

var __obf_a9e74ed9c49454e3 = reflect.TypeOf(([]string)(nil))

func __obf_439390bb02ad86d8(__obf_7bae0b613da60dd3 *Encoder, __obf_17732590722140e7 reflect.Value) error {
	return __obf_7bae0b613da60dd3.EncodeString(__obf_17732590722140e7.String())
}

func __obf_efad3cfa905a24ec(__obf_7bae0b613da60dd3 *Encoder, __obf_17732590722140e7 reflect.Value) error {
	return __obf_7bae0b613da60dd3.EncodeBytes(__obf_17732590722140e7.Bytes())
}

func __obf_2c6376330bc8c07a(__obf_7bae0b613da60dd3 *Encoder, __obf_17732590722140e7 reflect.Value) error {
	if __obf_0feb3528b7b709ec := __obf_7bae0b613da60dd3.EncodeBytesLen(__obf_17732590722140e7.Len()); __obf_0feb3528b7b709ec != nil {
		return __obf_0feb3528b7b709ec
	}

	if __obf_17732590722140e7.CanAddr() {
		__obf_a7fd7acf2bd4435f := __obf_17732590722140e7.Slice(0, __obf_17732590722140e7.Len()).Bytes()
		return __obf_7bae0b613da60dd3.__obf_b0daaf664cd5cdfb(__obf_a7fd7acf2bd4435f)
	}
	__obf_7bae0b613da60dd3.__obf_9c8c6a5e4cc8822c = __obf_634cd331185f089d(__obf_7bae0b613da60dd3.__obf_9c8c6a5e4cc8822c, __obf_17732590722140e7.Len())
	reflect.Copy(reflect.ValueOf(__obf_7bae0b613da60dd3.__obf_9c8c6a5e4cc8822c), __obf_17732590722140e7)
	return __obf_7bae0b613da60dd3.__obf_b0daaf664cd5cdfb(__obf_7bae0b613da60dd3.__obf_9c8c6a5e4cc8822c)
}

func __obf_634cd331185f089d(__obf_a7fd7acf2bd4435f []byte, __obf_2b0247e819b1bf4a int) []byte {
	if cap(__obf_a7fd7acf2bd4435f) >= __obf_2b0247e819b1bf4a {
		return __obf_a7fd7acf2bd4435f[:__obf_2b0247e819b1bf4a]
	}
	__obf_a7fd7acf2bd4435f = __obf_a7fd7acf2bd4435f[:cap(__obf_a7fd7acf2bd4435f)]
	__obf_a7fd7acf2bd4435f = append(__obf_a7fd7acf2bd4435f, make([]byte, __obf_2b0247e819b1bf4a-len(__obf_a7fd7acf2bd4435f))...)
	return __obf_a7fd7acf2bd4435f
}

func (__obf_7bae0b613da60dd3 *Encoder) EncodeBytesLen(__obf_6dbc1eb636e18f94 int) error {
	if __obf_6dbc1eb636e18f94 < 256 {
		return __obf_7bae0b613da60dd3.__obf_fa2b4086fff3f28d(Bin8, uint8(__obf_6dbc1eb636e18f94))
	}
	if __obf_6dbc1eb636e18f94 <= math.MaxUint16 {
		return __obf_7bae0b613da60dd3.__obf_bd6e685495964ed5(Bin16, uint16(__obf_6dbc1eb636e18f94))
	}
	return __obf_7bae0b613da60dd3.__obf_a95f338c34ae207a(Bin32, uint32(__obf_6dbc1eb636e18f94))
}

func (__obf_7bae0b613da60dd3 *Encoder) __obf_a2dad3baa5ecee91(__obf_6dbc1eb636e18f94 int) error {
	if __obf_6dbc1eb636e18f94 < 32 {
		return __obf_7bae0b613da60dd3.__obf_47b86f684323fc33(FixedStrLow | byte(__obf_6dbc1eb636e18f94))
	}
	if __obf_6dbc1eb636e18f94 < 256 {
		return __obf_7bae0b613da60dd3.__obf_fa2b4086fff3f28d(Str8, uint8(__obf_6dbc1eb636e18f94))
	}
	if __obf_6dbc1eb636e18f94 <= math.MaxUint16 {
		return __obf_7bae0b613da60dd3.__obf_bd6e685495964ed5(Str16, uint16(__obf_6dbc1eb636e18f94))
	}
	return __obf_7bae0b613da60dd3.__obf_a95f338c34ae207a(Str32, uint32(__obf_6dbc1eb636e18f94))
}

func (__obf_7bae0b613da60dd3 *Encoder) EncodeString(__obf_17732590722140e7 string) error {
	if __obf_65f94b9eb2a3c80f := __obf_7bae0b613da60dd3.__obf_a15a642720e39c3e&__obf_1fcd090a6da01656 != 0; __obf_65f94b9eb2a3c80f || len(__obf_7bae0b613da60dd3.__obf_25128eea0d506b65) > 0 {
		return __obf_7bae0b613da60dd3.__obf_66adb399f41a0954(__obf_17732590722140e7, __obf_65f94b9eb2a3c80f)
	}
	return __obf_7bae0b613da60dd3.__obf_4920ad7ca3e462a3(__obf_17732590722140e7)
}

func (__obf_7bae0b613da60dd3 *Encoder) __obf_4920ad7ca3e462a3(__obf_17732590722140e7 string) error {
	if __obf_0feb3528b7b709ec := __obf_7bae0b613da60dd3.__obf_a2dad3baa5ecee91(len(__obf_17732590722140e7)); __obf_0feb3528b7b709ec != nil {
		return __obf_0feb3528b7b709ec
	}
	return __obf_7bae0b613da60dd3.__obf_ff0da04198a8b5e2(__obf_17732590722140e7)
}

func (__obf_7bae0b613da60dd3 *Encoder) EncodeBytes(__obf_17732590722140e7 []byte) error {
	if __obf_17732590722140e7 == nil {
		return __obf_7bae0b613da60dd3.EncodeNil()
	}
	if __obf_0feb3528b7b709ec := __obf_7bae0b613da60dd3.EncodeBytesLen(len(__obf_17732590722140e7)); __obf_0feb3528b7b709ec != nil {
		return __obf_0feb3528b7b709ec
	}
	return __obf_7bae0b613da60dd3.__obf_b0daaf664cd5cdfb(__obf_17732590722140e7)
}

func (__obf_7bae0b613da60dd3 *Encoder) EncodeArrayLen(__obf_6dbc1eb636e18f94 int) error {
	if __obf_6dbc1eb636e18f94 < 16 {
		return __obf_7bae0b613da60dd3.__obf_47b86f684323fc33(FixedArrayLow | byte(__obf_6dbc1eb636e18f94))
	}
	if __obf_6dbc1eb636e18f94 <= math.MaxUint16 {
		return __obf_7bae0b613da60dd3.__obf_bd6e685495964ed5(Array16, uint16(__obf_6dbc1eb636e18f94))
	}
	return __obf_7bae0b613da60dd3.__obf_a95f338c34ae207a(Array32, uint32(__obf_6dbc1eb636e18f94))
}

func __obf_c225cd0483ee2104(__obf_7bae0b613da60dd3 *Encoder, __obf_17732590722140e7 reflect.Value) error {
	__obf_02601a79721c2043 := __obf_17732590722140e7.Convert(__obf_a9e74ed9c49454e3).Interface().([]string)
	return __obf_7bae0b613da60dd3.__obf_9ac36acedd38be06(__obf_02601a79721c2043)
}

func (__obf_7bae0b613da60dd3 *Encoder) __obf_9ac36acedd38be06(__obf_a93d004caac96500 []string) error {
	if __obf_a93d004caac96500 == nil {
		return __obf_7bae0b613da60dd3.EncodeNil()
	}
	if __obf_0feb3528b7b709ec := __obf_7bae0b613da60dd3.EncodeArrayLen(len(__obf_a93d004caac96500)); __obf_0feb3528b7b709ec != nil {
		return __obf_0feb3528b7b709ec
	}
	for _, __obf_17732590722140e7 := range __obf_a93d004caac96500 {
		if __obf_0feb3528b7b709ec := __obf_7bae0b613da60dd3.EncodeString(__obf_17732590722140e7); __obf_0feb3528b7b709ec != nil {
			return __obf_0feb3528b7b709ec
		}
	}
	return nil
}

func __obf_0b25b51b78e0d671(__obf_7bae0b613da60dd3 *Encoder, __obf_17732590722140e7 reflect.Value) error {
	if __obf_17732590722140e7.IsNil() {
		return __obf_7bae0b613da60dd3.EncodeNil()
	}
	return __obf_743bdcb3d7776bf9(__obf_7bae0b613da60dd3, __obf_17732590722140e7)
}

func __obf_743bdcb3d7776bf9(__obf_7bae0b613da60dd3 *Encoder, __obf_17732590722140e7 reflect.Value) error {
	__obf_6dbc1eb636e18f94 := __obf_17732590722140e7.Len()
	if __obf_0feb3528b7b709ec := __obf_7bae0b613da60dd3.EncodeArrayLen(__obf_6dbc1eb636e18f94); __obf_0feb3528b7b709ec != nil {
		return __obf_0feb3528b7b709ec
	}
	for __obf_101eec84c8338296 := 0; __obf_101eec84c8338296 < __obf_6dbc1eb636e18f94; __obf_101eec84c8338296++ {
		if __obf_0feb3528b7b709ec := __obf_7bae0b613da60dd3.EncodeValue(__obf_17732590722140e7.Index(__obf_101eec84c8338296)); __obf_0feb3528b7b709ec != nil {
			return __obf_0feb3528b7b709ec
		}
	}
	return nil
}
