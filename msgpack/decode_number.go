package __obf_a935eb7f548271a4

import (
	"fmt"
	"math"
	"reflect"
)

func (__obf_a21885da2425f2b2 *Decoder) __obf_b2541a0cb78c8e1f(__obf_326af9bd942662ac int) error {
	_, __obf_4d327e1cd40c2e21 := __obf_a21885da2425f2b2.__obf_595cbf51b6653ebf(__obf_326af9bd942662ac)
	return __obf_4d327e1cd40c2e21
}

func (__obf_a21885da2425f2b2 *Decoder) uint8() (uint8, error) {
	__obf_f5df560f4d67421b, __obf_4d327e1cd40c2e21 := __obf_a21885da2425f2b2.__obf_61e5879fcf7b35ce()
	if __obf_4d327e1cd40c2e21 != nil {
		return 0, __obf_4d327e1cd40c2e21
	}
	return __obf_f5df560f4d67421b, nil
}

func (__obf_a21885da2425f2b2 *Decoder) int8() (int8, error) {
	__obf_326af9bd942662ac, __obf_4d327e1cd40c2e21 := __obf_a21885da2425f2b2.uint8()
	return int8(__obf_326af9bd942662ac), __obf_4d327e1cd40c2e21
}

func (__obf_a21885da2425f2b2 *Decoder) uint16() (uint16, error) {
	__obf_f2ca794293605b73, __obf_4d327e1cd40c2e21 := __obf_a21885da2425f2b2.__obf_595cbf51b6653ebf(2)
	if __obf_4d327e1cd40c2e21 != nil {
		return 0, __obf_4d327e1cd40c2e21
	}
	return (uint16(__obf_f2ca794293605b73[0]) << 8) | uint16(__obf_f2ca794293605b73[1]), nil
}

func (__obf_a21885da2425f2b2 *Decoder) int16() (int16, error) {
	__obf_326af9bd942662ac, __obf_4d327e1cd40c2e21 := __obf_a21885da2425f2b2.uint16()
	return int16(__obf_326af9bd942662ac), __obf_4d327e1cd40c2e21
}

func (__obf_a21885da2425f2b2 *Decoder) uint32() (uint32, error) {
	__obf_f2ca794293605b73, __obf_4d327e1cd40c2e21 := __obf_a21885da2425f2b2.__obf_595cbf51b6653ebf(4)
	if __obf_4d327e1cd40c2e21 != nil {
		return 0, __obf_4d327e1cd40c2e21
	}
	__obf_326af9bd942662ac := (uint32(__obf_f2ca794293605b73[0]) << 24) |
		(uint32(__obf_f2ca794293605b73[1]) << 16) |
		(uint32(__obf_f2ca794293605b73[2]) << 8) |
		uint32(__obf_f2ca794293605b73[3])
	return __obf_326af9bd942662ac, nil
}

func (__obf_a21885da2425f2b2 *Decoder) int32() (int32, error) {
	__obf_326af9bd942662ac, __obf_4d327e1cd40c2e21 := __obf_a21885da2425f2b2.uint32()
	return int32(__obf_326af9bd942662ac), __obf_4d327e1cd40c2e21
}

func (__obf_a21885da2425f2b2 *Decoder) uint64() (uint64, error) {
	__obf_f2ca794293605b73, __obf_4d327e1cd40c2e21 := __obf_a21885da2425f2b2.__obf_595cbf51b6653ebf(8)
	if __obf_4d327e1cd40c2e21 != nil {
		return 0, __obf_4d327e1cd40c2e21
	}
	__obf_326af9bd942662ac := (uint64(__obf_f2ca794293605b73[0]) << 56) |
		(uint64(__obf_f2ca794293605b73[1]) << 48) |
		(uint64(__obf_f2ca794293605b73[2]) << 40) |
		(uint64(__obf_f2ca794293605b73[3]) << 32) |
		(uint64(__obf_f2ca794293605b73[4]) << 24) |
		(uint64(__obf_f2ca794293605b73[5]) << 16) |
		(uint64(__obf_f2ca794293605b73[6]) << 8) |
		uint64(__obf_f2ca794293605b73[7])
	return __obf_326af9bd942662ac, nil
}

func (__obf_a21885da2425f2b2 *Decoder) int64() (int64, error) {
	__obf_326af9bd942662ac, __obf_4d327e1cd40c2e21 := __obf_a21885da2425f2b2.uint64()
	return int64(__obf_326af9bd942662ac), __obf_4d327e1cd40c2e21
}

// DecodeUint64 decodes msgpack int8/16/32/64 and uint8/16/32/64
// into Go uint64.
func (__obf_a21885da2425f2b2 *Decoder) DecodeUint64() (uint64, error) {
	__obf_f5df560f4d67421b, __obf_4d327e1cd40c2e21 := __obf_a21885da2425f2b2.__obf_61e5879fcf7b35ce()
	if __obf_4d327e1cd40c2e21 != nil {
		return 0, __obf_4d327e1cd40c2e21
	}
	return __obf_a21885da2425f2b2.uint(__obf_f5df560f4d67421b)
}

func (__obf_a21885da2425f2b2 *Decoder) uint(__obf_f5df560f4d67421b byte) (uint64, error) {
	if __obf_f5df560f4d67421b == Nil {
		return 0, nil
	}
	if IsFixedNum(__obf_f5df560f4d67421b) {
		return uint64(int8(__obf_f5df560f4d67421b)), nil
	}
	switch __obf_f5df560f4d67421b {
	case Uint8:
		__obf_326af9bd942662ac, __obf_4d327e1cd40c2e21 := __obf_a21885da2425f2b2.uint8()
		return uint64(__obf_326af9bd942662ac), __obf_4d327e1cd40c2e21
	case Int8:
		__obf_326af9bd942662ac, __obf_4d327e1cd40c2e21 := __obf_a21885da2425f2b2.int8()
		return uint64(__obf_326af9bd942662ac), __obf_4d327e1cd40c2e21
	case Uint16:
		__obf_326af9bd942662ac, __obf_4d327e1cd40c2e21 := __obf_a21885da2425f2b2.uint16()
		return uint64(__obf_326af9bd942662ac), __obf_4d327e1cd40c2e21
	case Int16:
		__obf_326af9bd942662ac, __obf_4d327e1cd40c2e21 := __obf_a21885da2425f2b2.int16()
		return uint64(__obf_326af9bd942662ac), __obf_4d327e1cd40c2e21
	case Uint32:
		__obf_326af9bd942662ac, __obf_4d327e1cd40c2e21 := __obf_a21885da2425f2b2.uint32()
		return uint64(__obf_326af9bd942662ac), __obf_4d327e1cd40c2e21
	case Int32:
		__obf_326af9bd942662ac, __obf_4d327e1cd40c2e21 := __obf_a21885da2425f2b2.int32()
		return uint64(__obf_326af9bd942662ac), __obf_4d327e1cd40c2e21
	case Uint64, Int64:
		return __obf_a21885da2425f2b2.uint64()
	}
	return 0, fmt.Errorf("msgpack: invalid code=%x decoding uint64", __obf_f5df560f4d67421b)
}

// DecodeInt64 decodes msgpack int8/16/32/64 and uint8/16/32/64
// into Go int64.
func (__obf_a21885da2425f2b2 *Decoder) DecodeInt64() (int64, error) {
	__obf_f5df560f4d67421b, __obf_4d327e1cd40c2e21 := __obf_a21885da2425f2b2.__obf_61e5879fcf7b35ce()
	if __obf_4d327e1cd40c2e21 != nil {
		return 0, __obf_4d327e1cd40c2e21
	}
	return __obf_a21885da2425f2b2.int(__obf_f5df560f4d67421b)
}

func (__obf_a21885da2425f2b2 *Decoder) int(__obf_f5df560f4d67421b byte) (int64, error) {
	if __obf_f5df560f4d67421b == Nil {
		return 0, nil
	}
	if IsFixedNum(__obf_f5df560f4d67421b) {
		return int64(int8(__obf_f5df560f4d67421b)), nil
	}
	switch __obf_f5df560f4d67421b {
	case Uint8:
		__obf_326af9bd942662ac, __obf_4d327e1cd40c2e21 := __obf_a21885da2425f2b2.uint8()
		return int64(__obf_326af9bd942662ac), __obf_4d327e1cd40c2e21
	case Int8:
		__obf_326af9bd942662ac, __obf_4d327e1cd40c2e21 := __obf_a21885da2425f2b2.uint8()
		return int64(int8(__obf_326af9bd942662ac)), __obf_4d327e1cd40c2e21
	case Uint16:
		__obf_326af9bd942662ac, __obf_4d327e1cd40c2e21 := __obf_a21885da2425f2b2.uint16()
		return int64(__obf_326af9bd942662ac), __obf_4d327e1cd40c2e21
	case Int16:
		__obf_326af9bd942662ac, __obf_4d327e1cd40c2e21 := __obf_a21885da2425f2b2.uint16()
		return int64(int16(__obf_326af9bd942662ac)), __obf_4d327e1cd40c2e21
	case Uint32:
		__obf_326af9bd942662ac, __obf_4d327e1cd40c2e21 := __obf_a21885da2425f2b2.uint32()
		return int64(__obf_326af9bd942662ac), __obf_4d327e1cd40c2e21
	case Int32:
		__obf_326af9bd942662ac, __obf_4d327e1cd40c2e21 := __obf_a21885da2425f2b2.uint32()
		return int64(int32(__obf_326af9bd942662ac)), __obf_4d327e1cd40c2e21
	case Uint64, Int64:
		__obf_326af9bd942662ac, __obf_4d327e1cd40c2e21 := __obf_a21885da2425f2b2.uint64()
		return int64(__obf_326af9bd942662ac), __obf_4d327e1cd40c2e21
	}
	return 0, fmt.Errorf("msgpack: invalid code=%x decoding int64", __obf_f5df560f4d67421b)
}

func (__obf_a21885da2425f2b2 *Decoder) DecodeFloat32() (float32, error) {
	__obf_f5df560f4d67421b, __obf_4d327e1cd40c2e21 := __obf_a21885da2425f2b2.__obf_61e5879fcf7b35ce()
	if __obf_4d327e1cd40c2e21 != nil {
		return 0, __obf_4d327e1cd40c2e21
	}
	return __obf_a21885da2425f2b2.float32(__obf_f5df560f4d67421b)
}

func (__obf_a21885da2425f2b2 *Decoder) float32(__obf_f5df560f4d67421b byte) (float32, error) {
	if __obf_f5df560f4d67421b == Float {
		__obf_326af9bd942662ac, __obf_4d327e1cd40c2e21 := __obf_a21885da2425f2b2.uint32()
		if __obf_4d327e1cd40c2e21 != nil {
			return 0, __obf_4d327e1cd40c2e21
		}
		return math.Float32frombits(__obf_326af9bd942662ac), nil
	}
	__obf_326af9bd942662ac, __obf_4d327e1cd40c2e21 := __obf_a21885da2425f2b2.int(__obf_f5df560f4d67421b)
	if __obf_4d327e1cd40c2e21 != nil {
		return 0, fmt.Errorf("msgpack: invalid code=%x decoding float32", __obf_f5df560f4d67421b)
	}
	return float32(__obf_326af9bd942662ac), nil
}

// DecodeFloat64 decodes msgpack float32/64 into Go float64.
func (__obf_a21885da2425f2b2 *Decoder) DecodeFloat64() (float64, error) {
	__obf_f5df560f4d67421b, __obf_4d327e1cd40c2e21 := __obf_a21885da2425f2b2.__obf_61e5879fcf7b35ce()
	if __obf_4d327e1cd40c2e21 != nil {
		return 0, __obf_4d327e1cd40c2e21
	}
	return __obf_a21885da2425f2b2.float64(__obf_f5df560f4d67421b)
}

func (__obf_a21885da2425f2b2 *Decoder) float64(__obf_f5df560f4d67421b byte) (float64, error) {
	switch __obf_f5df560f4d67421b {
	case Float:
		__obf_326af9bd942662ac, __obf_4d327e1cd40c2e21 := __obf_a21885da2425f2b2.float32(__obf_f5df560f4d67421b)
		if __obf_4d327e1cd40c2e21 != nil {
			return 0, __obf_4d327e1cd40c2e21
		}
		return float64(__obf_326af9bd942662ac), nil
	case Double:
		__obf_326af9bd942662ac, __obf_4d327e1cd40c2e21 := __obf_a21885da2425f2b2.uint64()
		if __obf_4d327e1cd40c2e21 != nil {
			return 0, __obf_4d327e1cd40c2e21
		}
		return math.Float64frombits(__obf_326af9bd942662ac), nil
	}
	__obf_326af9bd942662ac, __obf_4d327e1cd40c2e21 := __obf_a21885da2425f2b2.int(__obf_f5df560f4d67421b)
	if __obf_4d327e1cd40c2e21 != nil {
		return 0, fmt.Errorf("msgpack: invalid code=%x decoding float32", __obf_f5df560f4d67421b)
	}
	return float64(__obf_326af9bd942662ac), nil
}

func (__obf_a21885da2425f2b2 *Decoder) DecodeUint() (uint, error) {
	__obf_326af9bd942662ac, __obf_4d327e1cd40c2e21 := __obf_a21885da2425f2b2.DecodeUint64()
	return uint(__obf_326af9bd942662ac), __obf_4d327e1cd40c2e21
}

func (__obf_a21885da2425f2b2 *Decoder) DecodeUint8() (uint8, error) {
	__obf_326af9bd942662ac, __obf_4d327e1cd40c2e21 := __obf_a21885da2425f2b2.DecodeUint64()
	return uint8(__obf_326af9bd942662ac), __obf_4d327e1cd40c2e21
}

func (__obf_a21885da2425f2b2 *Decoder) DecodeUint16() (uint16, error) {
	__obf_326af9bd942662ac, __obf_4d327e1cd40c2e21 := __obf_a21885da2425f2b2.DecodeUint64()
	return uint16(__obf_326af9bd942662ac), __obf_4d327e1cd40c2e21
}

func (__obf_a21885da2425f2b2 *Decoder) DecodeUint32() (uint32, error) {
	__obf_326af9bd942662ac, __obf_4d327e1cd40c2e21 := __obf_a21885da2425f2b2.DecodeUint64()
	return uint32(__obf_326af9bd942662ac), __obf_4d327e1cd40c2e21
}

func (__obf_a21885da2425f2b2 *Decoder) DecodeInt() (int, error) {
	__obf_326af9bd942662ac, __obf_4d327e1cd40c2e21 := __obf_a21885da2425f2b2.DecodeInt64()
	return int(__obf_326af9bd942662ac), __obf_4d327e1cd40c2e21
}

func (__obf_a21885da2425f2b2 *Decoder) DecodeInt8() (int8, error) {
	__obf_326af9bd942662ac, __obf_4d327e1cd40c2e21 := __obf_a21885da2425f2b2.DecodeInt64()
	return int8(__obf_326af9bd942662ac), __obf_4d327e1cd40c2e21
}

func (__obf_a21885da2425f2b2 *Decoder) DecodeInt16() (int16, error) {
	__obf_326af9bd942662ac, __obf_4d327e1cd40c2e21 := __obf_a21885da2425f2b2.DecodeInt64()
	return int16(__obf_326af9bd942662ac), __obf_4d327e1cd40c2e21
}

func (__obf_a21885da2425f2b2 *Decoder) DecodeInt32() (int32, error) {
	__obf_326af9bd942662ac, __obf_4d327e1cd40c2e21 := __obf_a21885da2425f2b2.DecodeInt64()
	return int32(__obf_326af9bd942662ac), __obf_4d327e1cd40c2e21
}

func __obf_79d81581a625ef22(__obf_a21885da2425f2b2 *Decoder, __obf_6d570581f4b60dbc reflect.Value) error {
	__obf_f6dd34b5068d19f3, __obf_4d327e1cd40c2e21 := __obf_a21885da2425f2b2.DecodeFloat32()
	if __obf_4d327e1cd40c2e21 != nil {
		return __obf_4d327e1cd40c2e21
	}
	__obf_6d570581f4b60dbc.
		SetFloat(float64(__obf_f6dd34b5068d19f3))
	return nil
}

func __obf_2c87d512f404034f(__obf_a21885da2425f2b2 *Decoder, __obf_6d570581f4b60dbc reflect.Value) error {
	__obf_f6dd34b5068d19f3, __obf_4d327e1cd40c2e21 := __obf_a21885da2425f2b2.DecodeFloat64()
	if __obf_4d327e1cd40c2e21 != nil {
		return __obf_4d327e1cd40c2e21
	}
	__obf_6d570581f4b60dbc.
		SetFloat(__obf_f6dd34b5068d19f3)
	return nil
}

func __obf_93b2295ef65970e5(__obf_a21885da2425f2b2 *Decoder, __obf_6d570581f4b60dbc reflect.Value) error {
	__obf_326af9bd942662ac, __obf_4d327e1cd40c2e21 := __obf_a21885da2425f2b2.DecodeInt64()
	if __obf_4d327e1cd40c2e21 != nil {
		return __obf_4d327e1cd40c2e21
	}
	__obf_6d570581f4b60dbc.
		SetInt(__obf_326af9bd942662ac)
	return nil
}

func __obf_d5ee3308a7edaa80(__obf_a21885da2425f2b2 *Decoder, __obf_6d570581f4b60dbc reflect.Value) error {
	__obf_326af9bd942662ac, __obf_4d327e1cd40c2e21 := __obf_a21885da2425f2b2.DecodeUint64()
	if __obf_4d327e1cd40c2e21 != nil {
		return __obf_4d327e1cd40c2e21
	}
	__obf_6d570581f4b60dbc.
		SetUint(__obf_326af9bd942662ac)
	return nil
}
