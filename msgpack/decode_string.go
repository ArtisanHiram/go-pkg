package __obf_3edaa49e853afa16

import (
	"fmt"
	"reflect"
)

func (__obf_015afbee33862a01 *Decoder) __obf_196dd6b18db329d3(__obf_145c7a7d25eea2bd byte) (int, error) {
	if __obf_145c7a7d25eea2bd == Nil {
		return -1, nil
	}

	if IsFixedString(__obf_145c7a7d25eea2bd) {
		return int(__obf_145c7a7d25eea2bd & FixedStrMask), nil
	}

	switch __obf_145c7a7d25eea2bd {
	case Str8, Bin8:
		__obf_56127cd370854f0b, __obf_3aa78ad28f50ed46 := __obf_015afbee33862a01.uint8()
		return int(__obf_56127cd370854f0b), __obf_3aa78ad28f50ed46
	case Str16, Bin16:
		__obf_56127cd370854f0b, __obf_3aa78ad28f50ed46 := __obf_015afbee33862a01.uint16()
		return int(__obf_56127cd370854f0b), __obf_3aa78ad28f50ed46
	case Str32, Bin32:
		__obf_56127cd370854f0b, __obf_3aa78ad28f50ed46 := __obf_015afbee33862a01.uint32()
		return int(__obf_56127cd370854f0b), __obf_3aa78ad28f50ed46
	}

	return 0, fmt.Errorf("msgpack: invalid code=%x decoding string/bytes length", __obf_145c7a7d25eea2bd)
}

func (__obf_015afbee33862a01 *Decoder) DecodeString() (string, error) {
	if __obf_60ddbad9c6f41c91 := __obf_015afbee33862a01.__obf_6aa0efd537415192&__obf_48e006ba02b70b78 != 0; __obf_60ddbad9c6f41c91 || len(__obf_015afbee33862a01.__obf_b014355f64826d7b) > 0 {
		return __obf_015afbee33862a01.__obf_bc64b9136455b38b(__obf_60ddbad9c6f41c91)
	}
	__obf_145c7a7d25eea2bd, __obf_3aa78ad28f50ed46 := __obf_015afbee33862a01.__obf_4c5f8d154c9027ea()
	if __obf_3aa78ad28f50ed46 != nil {
		return "", __obf_3aa78ad28f50ed46
	}
	return __obf_015afbee33862a01.string(__obf_145c7a7d25eea2bd)
}

func (__obf_015afbee33862a01 *Decoder) string(__obf_145c7a7d25eea2bd byte) (string, error) {
	__obf_56127cd370854f0b, __obf_3aa78ad28f50ed46 := __obf_015afbee33862a01.__obf_196dd6b18db329d3(__obf_145c7a7d25eea2bd)
	if __obf_3aa78ad28f50ed46 != nil {
		return "", __obf_3aa78ad28f50ed46
	}
	return __obf_015afbee33862a01.__obf_b85f6957ac631eeb(__obf_56127cd370854f0b)
}

func (__obf_015afbee33862a01 *Decoder) __obf_b85f6957ac631eeb(__obf_56127cd370854f0b int) (string, error) {
	if __obf_56127cd370854f0b <= 0 {
		return "", nil
	}
	__obf_9b4a5a04bdad2347, __obf_3aa78ad28f50ed46 := __obf_015afbee33862a01.__obf_2c3d4aa0a72cea84(__obf_56127cd370854f0b)
	return string(__obf_9b4a5a04bdad2347), __obf_3aa78ad28f50ed46
}

func __obf_e018477bba540f48(__obf_015afbee33862a01 *Decoder, __obf_85d270343a0dfe14 reflect.Value) error {
	__obf_6827ff1b59d7ccec, __obf_3aa78ad28f50ed46 := __obf_015afbee33862a01.DecodeString()
	if __obf_3aa78ad28f50ed46 != nil {
		return __obf_3aa78ad28f50ed46
	}
	__obf_85d270343a0dfe14.
		SetString(__obf_6827ff1b59d7ccec)
	return nil
}

func (__obf_015afbee33862a01 *Decoder) __obf_aeaabfb6cb02451a() (string, error) {
	if __obf_60ddbad9c6f41c91 := __obf_015afbee33862a01.__obf_6aa0efd537415192&__obf_48e006ba02b70b78 != 0; __obf_60ddbad9c6f41c91 || len(__obf_015afbee33862a01.__obf_b014355f64826d7b) > 0 {
		return __obf_015afbee33862a01.__obf_bc64b9136455b38b(__obf_60ddbad9c6f41c91)
	}
	__obf_145c7a7d25eea2bd, __obf_3aa78ad28f50ed46 := __obf_015afbee33862a01.__obf_4c5f8d154c9027ea()
	if __obf_3aa78ad28f50ed46 != nil {
		return "", __obf_3aa78ad28f50ed46
	}
	__obf_56127cd370854f0b, __obf_3aa78ad28f50ed46 := __obf_015afbee33862a01.__obf_196dd6b18db329d3(__obf_145c7a7d25eea2bd)
	if __obf_3aa78ad28f50ed46 != nil {
		return "", __obf_3aa78ad28f50ed46
	}
	if __obf_56127cd370854f0b == -1 {
		return "", nil
	}
	__obf_9b4a5a04bdad2347, __obf_3aa78ad28f50ed46 := __obf_015afbee33862a01.__obf_2c3d4aa0a72cea84(__obf_56127cd370854f0b)
	if __obf_3aa78ad28f50ed46 != nil {
		return "", __obf_3aa78ad28f50ed46
	}

	return __obf_39ecaa2d03d54c22(__obf_9b4a5a04bdad2347), nil
}

func (__obf_015afbee33862a01 *Decoder) __obf_2b47e19d74adf37a(__obf_8f0c71619c0382f1 *[]byte) error {
	__obf_145c7a7d25eea2bd, __obf_3aa78ad28f50ed46 := __obf_015afbee33862a01.__obf_4c5f8d154c9027ea()
	if __obf_3aa78ad28f50ed46 != nil {
		return __obf_3aa78ad28f50ed46
	}
	return __obf_015afbee33862a01.__obf_cc93506865e4dbec(__obf_145c7a7d25eea2bd, __obf_8f0c71619c0382f1)
}

func (__obf_015afbee33862a01 *Decoder) __obf_cc93506865e4dbec(__obf_145c7a7d25eea2bd byte, __obf_8f0c71619c0382f1 *[]byte) error {
	__obf_56127cd370854f0b, __obf_3aa78ad28f50ed46 := __obf_015afbee33862a01.__obf_196dd6b18db329d3(__obf_145c7a7d25eea2bd)
	if __obf_3aa78ad28f50ed46 != nil {
		return __obf_3aa78ad28f50ed46
	}
	if __obf_56127cd370854f0b == -1 {
		*__obf_8f0c71619c0382f1 = nil
		return nil
	}

	*__obf_8f0c71619c0382f1, __obf_3aa78ad28f50ed46 = __obf_2c3d4aa0a72cea84(__obf_015afbee33862a01.__obf_9b78235c82fd11c0, *__obf_8f0c71619c0382f1, __obf_56127cd370854f0b)
	return __obf_3aa78ad28f50ed46
}

func (__obf_015afbee33862a01 *Decoder) __obf_6d8f264a5e217b4a(__obf_145c7a7d25eea2bd byte) error {
	__obf_56127cd370854f0b, __obf_3aa78ad28f50ed46 := __obf_015afbee33862a01.__obf_196dd6b18db329d3(__obf_145c7a7d25eea2bd)
	if __obf_3aa78ad28f50ed46 != nil {
		return __obf_3aa78ad28f50ed46
	}
	if __obf_56127cd370854f0b <= 0 {
		return nil
	}
	return __obf_015afbee33862a01.__obf_1a431295b897a1b7(__obf_56127cd370854f0b)
}

func __obf_dcb87a3c2d394964(__obf_015afbee33862a01 *Decoder, __obf_85d270343a0dfe14 reflect.Value) error {
	__obf_145c7a7d25eea2bd, __obf_3aa78ad28f50ed46 := __obf_015afbee33862a01.__obf_4c5f8d154c9027ea()
	if __obf_3aa78ad28f50ed46 != nil {
		return __obf_3aa78ad28f50ed46
	}
	__obf_56127cd370854f0b, __obf_3aa78ad28f50ed46 := __obf_015afbee33862a01.__obf_196dd6b18db329d3(__obf_145c7a7d25eea2bd)
	if __obf_3aa78ad28f50ed46 != nil {
		return __obf_3aa78ad28f50ed46
	}
	if __obf_56127cd370854f0b == -1 {
		return nil
	}
	if __obf_56127cd370854f0b > __obf_85d270343a0dfe14.Len() {
		return fmt.Errorf("%s len is %d, but msgpack has %d elements", __obf_85d270343a0dfe14.Type(), __obf_85d270343a0dfe14.Len(), __obf_56127cd370854f0b)
	}
	__obf_9b4a5a04bdad2347 := __obf_85d270343a0dfe14.Slice(0, __obf_56127cd370854f0b).Bytes()
	return __obf_015afbee33862a01.__obf_8ba8963a8ad56c5c(__obf_9b4a5a04bdad2347)
}
