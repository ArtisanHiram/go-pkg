package __obf_de86cdc8ae98b45b

import (
	"fmt"
	"math"
	"reflect"
)

func (__obf_dcaad1165bb07af9 *Decoder) __obf_6f540045c1336cac(__obf_2b0247e819b1bf4a int) error {
	_, __obf_0feb3528b7b709ec := __obf_dcaad1165bb07af9.__obf_8f5a813341411779(__obf_2b0247e819b1bf4a)
	return __obf_0feb3528b7b709ec
}

func (__obf_dcaad1165bb07af9 *Decoder) uint8() (uint8, error) {
	__obf_ecf6d2d3253a058d, __obf_0feb3528b7b709ec := __obf_dcaad1165bb07af9.__obf_74ac05d9a9da01f4()
	if __obf_0feb3528b7b709ec != nil {
		return 0, __obf_0feb3528b7b709ec
	}
	return __obf_ecf6d2d3253a058d, nil
}

func (__obf_dcaad1165bb07af9 *Decoder) int8() (int8, error) {
	__obf_2b0247e819b1bf4a, __obf_0feb3528b7b709ec := __obf_dcaad1165bb07af9.uint8()
	return int8(__obf_2b0247e819b1bf4a), __obf_0feb3528b7b709ec
}

func (__obf_dcaad1165bb07af9 *Decoder) uint16() (uint16, error) {
	__obf_a7fd7acf2bd4435f, __obf_0feb3528b7b709ec := __obf_dcaad1165bb07af9.__obf_8f5a813341411779(2)
	if __obf_0feb3528b7b709ec != nil {
		return 0, __obf_0feb3528b7b709ec
	}
	return (uint16(__obf_a7fd7acf2bd4435f[0]) << 8) | uint16(__obf_a7fd7acf2bd4435f[1]), nil
}

func (__obf_dcaad1165bb07af9 *Decoder) int16() (int16, error) {
	__obf_2b0247e819b1bf4a, __obf_0feb3528b7b709ec := __obf_dcaad1165bb07af9.uint16()
	return int16(__obf_2b0247e819b1bf4a), __obf_0feb3528b7b709ec
}

func (__obf_dcaad1165bb07af9 *Decoder) uint32() (uint32, error) {
	__obf_a7fd7acf2bd4435f, __obf_0feb3528b7b709ec := __obf_dcaad1165bb07af9.__obf_8f5a813341411779(4)
	if __obf_0feb3528b7b709ec != nil {
		return 0, __obf_0feb3528b7b709ec
	}
	__obf_2b0247e819b1bf4a := (uint32(__obf_a7fd7acf2bd4435f[0]) << 24) |
		(uint32(__obf_a7fd7acf2bd4435f[1]) << 16) |
		(uint32(__obf_a7fd7acf2bd4435f[2]) << 8) |
		uint32(__obf_a7fd7acf2bd4435f[3])
	return __obf_2b0247e819b1bf4a, nil
}

func (__obf_dcaad1165bb07af9 *Decoder) int32() (int32, error) {
	__obf_2b0247e819b1bf4a, __obf_0feb3528b7b709ec := __obf_dcaad1165bb07af9.uint32()
	return int32(__obf_2b0247e819b1bf4a), __obf_0feb3528b7b709ec
}

func (__obf_dcaad1165bb07af9 *Decoder) uint64() (uint64, error) {
	__obf_a7fd7acf2bd4435f, __obf_0feb3528b7b709ec := __obf_dcaad1165bb07af9.__obf_8f5a813341411779(8)
	if __obf_0feb3528b7b709ec != nil {
		return 0, __obf_0feb3528b7b709ec
	}
	__obf_2b0247e819b1bf4a := (uint64(__obf_a7fd7acf2bd4435f[0]) << 56) |
		(uint64(__obf_a7fd7acf2bd4435f[1]) << 48) |
		(uint64(__obf_a7fd7acf2bd4435f[2]) << 40) |
		(uint64(__obf_a7fd7acf2bd4435f[3]) << 32) |
		(uint64(__obf_a7fd7acf2bd4435f[4]) << 24) |
		(uint64(__obf_a7fd7acf2bd4435f[5]) << 16) |
		(uint64(__obf_a7fd7acf2bd4435f[6]) << 8) |
		uint64(__obf_a7fd7acf2bd4435f[7])
	return __obf_2b0247e819b1bf4a, nil
}

func (__obf_dcaad1165bb07af9 *Decoder) int64() (int64, error) {
	__obf_2b0247e819b1bf4a, __obf_0feb3528b7b709ec := __obf_dcaad1165bb07af9.uint64()
	return int64(__obf_2b0247e819b1bf4a), __obf_0feb3528b7b709ec
}

// DecodeUint64 decodes msgpack int8/16/32/64 and uint8/16/32/64
// into Go uint64.
func (__obf_dcaad1165bb07af9 *Decoder) DecodeUint64() (uint64, error) {
	__obf_ecf6d2d3253a058d, __obf_0feb3528b7b709ec := __obf_dcaad1165bb07af9.__obf_74ac05d9a9da01f4()
	if __obf_0feb3528b7b709ec != nil {
		return 0, __obf_0feb3528b7b709ec
	}
	return __obf_dcaad1165bb07af9.uint(__obf_ecf6d2d3253a058d)
}

func (__obf_dcaad1165bb07af9 *Decoder) uint(__obf_ecf6d2d3253a058d byte) (uint64, error) {
	if __obf_ecf6d2d3253a058d == Nil {
		return 0, nil
	}
	if IsFixedNum(__obf_ecf6d2d3253a058d) {
		return uint64(int8(__obf_ecf6d2d3253a058d)), nil
	}
	switch __obf_ecf6d2d3253a058d {
	case Uint8:
		__obf_2b0247e819b1bf4a, __obf_0feb3528b7b709ec := __obf_dcaad1165bb07af9.uint8()
		return uint64(__obf_2b0247e819b1bf4a), __obf_0feb3528b7b709ec
	case Int8:
		__obf_2b0247e819b1bf4a, __obf_0feb3528b7b709ec := __obf_dcaad1165bb07af9.int8()
		return uint64(__obf_2b0247e819b1bf4a), __obf_0feb3528b7b709ec
	case Uint16:
		__obf_2b0247e819b1bf4a, __obf_0feb3528b7b709ec := __obf_dcaad1165bb07af9.uint16()
		return uint64(__obf_2b0247e819b1bf4a), __obf_0feb3528b7b709ec
	case Int16:
		__obf_2b0247e819b1bf4a, __obf_0feb3528b7b709ec := __obf_dcaad1165bb07af9.int16()
		return uint64(__obf_2b0247e819b1bf4a), __obf_0feb3528b7b709ec
	case Uint32:
		__obf_2b0247e819b1bf4a, __obf_0feb3528b7b709ec := __obf_dcaad1165bb07af9.uint32()
		return uint64(__obf_2b0247e819b1bf4a), __obf_0feb3528b7b709ec
	case Int32:
		__obf_2b0247e819b1bf4a, __obf_0feb3528b7b709ec := __obf_dcaad1165bb07af9.int32()
		return uint64(__obf_2b0247e819b1bf4a), __obf_0feb3528b7b709ec
	case Uint64, Int64:
		return __obf_dcaad1165bb07af9.uint64()
	}
	return 0, fmt.Errorf("msgpack: invalid code=%x decoding uint64", __obf_ecf6d2d3253a058d)
}

// DecodeInt64 decodes msgpack int8/16/32/64 and uint8/16/32/64
// into Go int64.
func (__obf_dcaad1165bb07af9 *Decoder) DecodeInt64() (int64, error) {
	__obf_ecf6d2d3253a058d, __obf_0feb3528b7b709ec := __obf_dcaad1165bb07af9.__obf_74ac05d9a9da01f4()
	if __obf_0feb3528b7b709ec != nil {
		return 0, __obf_0feb3528b7b709ec
	}
	return __obf_dcaad1165bb07af9.int(__obf_ecf6d2d3253a058d)
}

func (__obf_dcaad1165bb07af9 *Decoder) int(__obf_ecf6d2d3253a058d byte) (int64, error) {
	if __obf_ecf6d2d3253a058d == Nil {
		return 0, nil
	}
	if IsFixedNum(__obf_ecf6d2d3253a058d) {
		return int64(int8(__obf_ecf6d2d3253a058d)), nil
	}
	switch __obf_ecf6d2d3253a058d {
	case Uint8:
		__obf_2b0247e819b1bf4a, __obf_0feb3528b7b709ec := __obf_dcaad1165bb07af9.uint8()
		return int64(__obf_2b0247e819b1bf4a), __obf_0feb3528b7b709ec
	case Int8:
		__obf_2b0247e819b1bf4a, __obf_0feb3528b7b709ec := __obf_dcaad1165bb07af9.uint8()
		return int64(int8(__obf_2b0247e819b1bf4a)), __obf_0feb3528b7b709ec
	case Uint16:
		__obf_2b0247e819b1bf4a, __obf_0feb3528b7b709ec := __obf_dcaad1165bb07af9.uint16()
		return int64(__obf_2b0247e819b1bf4a), __obf_0feb3528b7b709ec
	case Int16:
		__obf_2b0247e819b1bf4a, __obf_0feb3528b7b709ec := __obf_dcaad1165bb07af9.uint16()
		return int64(int16(__obf_2b0247e819b1bf4a)), __obf_0feb3528b7b709ec
	case Uint32:
		__obf_2b0247e819b1bf4a, __obf_0feb3528b7b709ec := __obf_dcaad1165bb07af9.uint32()
		return int64(__obf_2b0247e819b1bf4a), __obf_0feb3528b7b709ec
	case Int32:
		__obf_2b0247e819b1bf4a, __obf_0feb3528b7b709ec := __obf_dcaad1165bb07af9.uint32()
		return int64(int32(__obf_2b0247e819b1bf4a)), __obf_0feb3528b7b709ec
	case Uint64, Int64:
		__obf_2b0247e819b1bf4a, __obf_0feb3528b7b709ec := __obf_dcaad1165bb07af9.uint64()
		return int64(__obf_2b0247e819b1bf4a), __obf_0feb3528b7b709ec
	}
	return 0, fmt.Errorf("msgpack: invalid code=%x decoding int64", __obf_ecf6d2d3253a058d)
}

func (__obf_dcaad1165bb07af9 *Decoder) DecodeFloat32() (float32, error) {
	__obf_ecf6d2d3253a058d, __obf_0feb3528b7b709ec := __obf_dcaad1165bb07af9.__obf_74ac05d9a9da01f4()
	if __obf_0feb3528b7b709ec != nil {
		return 0, __obf_0feb3528b7b709ec
	}
	return __obf_dcaad1165bb07af9.float32(__obf_ecf6d2d3253a058d)
}

func (__obf_dcaad1165bb07af9 *Decoder) float32(__obf_ecf6d2d3253a058d byte) (float32, error) {
	if __obf_ecf6d2d3253a058d == Float {
		__obf_2b0247e819b1bf4a, __obf_0feb3528b7b709ec := __obf_dcaad1165bb07af9.uint32()
		if __obf_0feb3528b7b709ec != nil {
			return 0, __obf_0feb3528b7b709ec
		}
		return math.Float32frombits(__obf_2b0247e819b1bf4a), nil
	}
	__obf_2b0247e819b1bf4a, __obf_0feb3528b7b709ec := __obf_dcaad1165bb07af9.int(__obf_ecf6d2d3253a058d)
	if __obf_0feb3528b7b709ec != nil {
		return 0, fmt.Errorf("msgpack: invalid code=%x decoding float32", __obf_ecf6d2d3253a058d)
	}
	return float32(__obf_2b0247e819b1bf4a), nil
}

// DecodeFloat64 decodes msgpack float32/64 into Go float64.
func (__obf_dcaad1165bb07af9 *Decoder) DecodeFloat64() (float64, error) {
	__obf_ecf6d2d3253a058d, __obf_0feb3528b7b709ec := __obf_dcaad1165bb07af9.__obf_74ac05d9a9da01f4()
	if __obf_0feb3528b7b709ec != nil {
		return 0, __obf_0feb3528b7b709ec
	}
	return __obf_dcaad1165bb07af9.float64(__obf_ecf6d2d3253a058d)
}

func (__obf_dcaad1165bb07af9 *Decoder) float64(__obf_ecf6d2d3253a058d byte) (float64, error) {
	switch __obf_ecf6d2d3253a058d {
	case Float:
		__obf_2b0247e819b1bf4a, __obf_0feb3528b7b709ec := __obf_dcaad1165bb07af9.float32(__obf_ecf6d2d3253a058d)
		if __obf_0feb3528b7b709ec != nil {
			return 0, __obf_0feb3528b7b709ec
		}
		return float64(__obf_2b0247e819b1bf4a), nil
	case Double:
		__obf_2b0247e819b1bf4a, __obf_0feb3528b7b709ec := __obf_dcaad1165bb07af9.uint64()
		if __obf_0feb3528b7b709ec != nil {
			return 0, __obf_0feb3528b7b709ec
		}
		return math.Float64frombits(__obf_2b0247e819b1bf4a), nil
	}
	__obf_2b0247e819b1bf4a, __obf_0feb3528b7b709ec := __obf_dcaad1165bb07af9.int(__obf_ecf6d2d3253a058d)
	if __obf_0feb3528b7b709ec != nil {
		return 0, fmt.Errorf("msgpack: invalid code=%x decoding float32", __obf_ecf6d2d3253a058d)
	}
	return float64(__obf_2b0247e819b1bf4a), nil
}

func (__obf_dcaad1165bb07af9 *Decoder) DecodeUint() (uint, error) {
	__obf_2b0247e819b1bf4a, __obf_0feb3528b7b709ec := __obf_dcaad1165bb07af9.DecodeUint64()
	return uint(__obf_2b0247e819b1bf4a), __obf_0feb3528b7b709ec
}

func (__obf_dcaad1165bb07af9 *Decoder) DecodeUint8() (uint8, error) {
	__obf_2b0247e819b1bf4a, __obf_0feb3528b7b709ec := __obf_dcaad1165bb07af9.DecodeUint64()
	return uint8(__obf_2b0247e819b1bf4a), __obf_0feb3528b7b709ec
}

func (__obf_dcaad1165bb07af9 *Decoder) DecodeUint16() (uint16, error) {
	__obf_2b0247e819b1bf4a, __obf_0feb3528b7b709ec := __obf_dcaad1165bb07af9.DecodeUint64()
	return uint16(__obf_2b0247e819b1bf4a), __obf_0feb3528b7b709ec
}

func (__obf_dcaad1165bb07af9 *Decoder) DecodeUint32() (uint32, error) {
	__obf_2b0247e819b1bf4a, __obf_0feb3528b7b709ec := __obf_dcaad1165bb07af9.DecodeUint64()
	return uint32(__obf_2b0247e819b1bf4a), __obf_0feb3528b7b709ec
}

func (__obf_dcaad1165bb07af9 *Decoder) DecodeInt() (int, error) {
	__obf_2b0247e819b1bf4a, __obf_0feb3528b7b709ec := __obf_dcaad1165bb07af9.DecodeInt64()
	return int(__obf_2b0247e819b1bf4a), __obf_0feb3528b7b709ec
}

func (__obf_dcaad1165bb07af9 *Decoder) DecodeInt8() (int8, error) {
	__obf_2b0247e819b1bf4a, __obf_0feb3528b7b709ec := __obf_dcaad1165bb07af9.DecodeInt64()
	return int8(__obf_2b0247e819b1bf4a), __obf_0feb3528b7b709ec
}

func (__obf_dcaad1165bb07af9 *Decoder) DecodeInt16() (int16, error) {
	__obf_2b0247e819b1bf4a, __obf_0feb3528b7b709ec := __obf_dcaad1165bb07af9.DecodeInt64()
	return int16(__obf_2b0247e819b1bf4a), __obf_0feb3528b7b709ec
}

func (__obf_dcaad1165bb07af9 *Decoder) DecodeInt32() (int32, error) {
	__obf_2b0247e819b1bf4a, __obf_0feb3528b7b709ec := __obf_dcaad1165bb07af9.DecodeInt64()
	return int32(__obf_2b0247e819b1bf4a), __obf_0feb3528b7b709ec
}

func __obf_42e61e6f8a4ef035(__obf_dcaad1165bb07af9 *Decoder, __obf_17732590722140e7 reflect.Value) error {
	__obf_cf8245c3f9570af5, __obf_0feb3528b7b709ec := __obf_dcaad1165bb07af9.DecodeFloat32()
	if __obf_0feb3528b7b709ec != nil {
		return __obf_0feb3528b7b709ec
	}
	__obf_17732590722140e7.
		SetFloat(float64(__obf_cf8245c3f9570af5))
	return nil
}

func __obf_57e51fbd6a8c8a03(__obf_dcaad1165bb07af9 *Decoder, __obf_17732590722140e7 reflect.Value) error {
	__obf_cf8245c3f9570af5, __obf_0feb3528b7b709ec := __obf_dcaad1165bb07af9.DecodeFloat64()
	if __obf_0feb3528b7b709ec != nil {
		return __obf_0feb3528b7b709ec
	}
	__obf_17732590722140e7.
		SetFloat(__obf_cf8245c3f9570af5)
	return nil
}

func __obf_58e1420db8e49e5b(__obf_dcaad1165bb07af9 *Decoder, __obf_17732590722140e7 reflect.Value) error {
	__obf_2b0247e819b1bf4a, __obf_0feb3528b7b709ec := __obf_dcaad1165bb07af9.DecodeInt64()
	if __obf_0feb3528b7b709ec != nil {
		return __obf_0feb3528b7b709ec
	}
	__obf_17732590722140e7.
		SetInt(__obf_2b0247e819b1bf4a)
	return nil
}

func __obf_460b1f6481895b57(__obf_dcaad1165bb07af9 *Decoder, __obf_17732590722140e7 reflect.Value) error {
	__obf_2b0247e819b1bf4a, __obf_0feb3528b7b709ec := __obf_dcaad1165bb07af9.DecodeUint64()
	if __obf_0feb3528b7b709ec != nil {
		return __obf_0feb3528b7b709ec
	}
	__obf_17732590722140e7.
		SetUint(__obf_2b0247e819b1bf4a)
	return nil
}
