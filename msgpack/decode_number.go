package __obf_3edaa49e853afa16

import (
	"fmt"
	"math"
	"reflect"
)

func (__obf_015afbee33862a01 *Decoder) __obf_1a431295b897a1b7(__obf_56127cd370854f0b int) error {
	_, __obf_3aa78ad28f50ed46 := __obf_015afbee33862a01.__obf_2c3d4aa0a72cea84(__obf_56127cd370854f0b)
	return __obf_3aa78ad28f50ed46
}

func (__obf_015afbee33862a01 *Decoder) uint8() (uint8, error) {
	__obf_145c7a7d25eea2bd, __obf_3aa78ad28f50ed46 := __obf_015afbee33862a01.__obf_4c5f8d154c9027ea()
	if __obf_3aa78ad28f50ed46 != nil {
		return 0, __obf_3aa78ad28f50ed46
	}
	return __obf_145c7a7d25eea2bd, nil
}

func (__obf_015afbee33862a01 *Decoder) int8() (int8, error) {
	__obf_56127cd370854f0b, __obf_3aa78ad28f50ed46 := __obf_015afbee33862a01.uint8()
	return int8(__obf_56127cd370854f0b), __obf_3aa78ad28f50ed46
}

func (__obf_015afbee33862a01 *Decoder) uint16() (uint16, error) {
	__obf_9b4a5a04bdad2347, __obf_3aa78ad28f50ed46 := __obf_015afbee33862a01.__obf_2c3d4aa0a72cea84(2)
	if __obf_3aa78ad28f50ed46 != nil {
		return 0, __obf_3aa78ad28f50ed46
	}
	return (uint16(__obf_9b4a5a04bdad2347[0]) << 8) | uint16(__obf_9b4a5a04bdad2347[1]), nil
}

func (__obf_015afbee33862a01 *Decoder) int16() (int16, error) {
	__obf_56127cd370854f0b, __obf_3aa78ad28f50ed46 := __obf_015afbee33862a01.uint16()
	return int16(__obf_56127cd370854f0b), __obf_3aa78ad28f50ed46
}

func (__obf_015afbee33862a01 *Decoder) uint32() (uint32, error) {
	__obf_9b4a5a04bdad2347, __obf_3aa78ad28f50ed46 := __obf_015afbee33862a01.__obf_2c3d4aa0a72cea84(4)
	if __obf_3aa78ad28f50ed46 != nil {
		return 0, __obf_3aa78ad28f50ed46
	}
	__obf_56127cd370854f0b := (uint32(__obf_9b4a5a04bdad2347[0]) << 24) |
		(uint32(__obf_9b4a5a04bdad2347[1]) << 16) |
		(uint32(__obf_9b4a5a04bdad2347[2]) << 8) |
		uint32(__obf_9b4a5a04bdad2347[3])
	return __obf_56127cd370854f0b, nil
}

func (__obf_015afbee33862a01 *Decoder) int32() (int32, error) {
	__obf_56127cd370854f0b, __obf_3aa78ad28f50ed46 := __obf_015afbee33862a01.uint32()
	return int32(__obf_56127cd370854f0b), __obf_3aa78ad28f50ed46
}

func (__obf_015afbee33862a01 *Decoder) uint64() (uint64, error) {
	__obf_9b4a5a04bdad2347, __obf_3aa78ad28f50ed46 := __obf_015afbee33862a01.__obf_2c3d4aa0a72cea84(8)
	if __obf_3aa78ad28f50ed46 != nil {
		return 0, __obf_3aa78ad28f50ed46
	}
	__obf_56127cd370854f0b := (uint64(__obf_9b4a5a04bdad2347[0]) << 56) |
		(uint64(__obf_9b4a5a04bdad2347[1]) << 48) |
		(uint64(__obf_9b4a5a04bdad2347[2]) << 40) |
		(uint64(__obf_9b4a5a04bdad2347[3]) << 32) |
		(uint64(__obf_9b4a5a04bdad2347[4]) << 24) |
		(uint64(__obf_9b4a5a04bdad2347[5]) << 16) |
		(uint64(__obf_9b4a5a04bdad2347[6]) << 8) |
		uint64(__obf_9b4a5a04bdad2347[7])
	return __obf_56127cd370854f0b, nil
}

func (__obf_015afbee33862a01 *Decoder) int64() (int64, error) {
	__obf_56127cd370854f0b, __obf_3aa78ad28f50ed46 := __obf_015afbee33862a01.uint64()
	return int64(__obf_56127cd370854f0b), __obf_3aa78ad28f50ed46
}

// DecodeUint64 decodes msgpack int8/16/32/64 and uint8/16/32/64
// into Go uint64.
func (__obf_015afbee33862a01 *Decoder) DecodeUint64() (uint64, error) {
	__obf_145c7a7d25eea2bd, __obf_3aa78ad28f50ed46 := __obf_015afbee33862a01.__obf_4c5f8d154c9027ea()
	if __obf_3aa78ad28f50ed46 != nil {
		return 0, __obf_3aa78ad28f50ed46
	}
	return __obf_015afbee33862a01.uint(__obf_145c7a7d25eea2bd)
}

func (__obf_015afbee33862a01 *Decoder) uint(__obf_145c7a7d25eea2bd byte) (uint64, error) {
	if __obf_145c7a7d25eea2bd == Nil {
		return 0, nil
	}
	if IsFixedNum(__obf_145c7a7d25eea2bd) {
		return uint64(int8(__obf_145c7a7d25eea2bd)), nil
	}
	switch __obf_145c7a7d25eea2bd {
	case Uint8:
		__obf_56127cd370854f0b, __obf_3aa78ad28f50ed46 := __obf_015afbee33862a01.uint8()
		return uint64(__obf_56127cd370854f0b), __obf_3aa78ad28f50ed46
	case Int8:
		__obf_56127cd370854f0b, __obf_3aa78ad28f50ed46 := __obf_015afbee33862a01.int8()
		return uint64(__obf_56127cd370854f0b), __obf_3aa78ad28f50ed46
	case Uint16:
		__obf_56127cd370854f0b, __obf_3aa78ad28f50ed46 := __obf_015afbee33862a01.uint16()
		return uint64(__obf_56127cd370854f0b), __obf_3aa78ad28f50ed46
	case Int16:
		__obf_56127cd370854f0b, __obf_3aa78ad28f50ed46 := __obf_015afbee33862a01.int16()
		return uint64(__obf_56127cd370854f0b), __obf_3aa78ad28f50ed46
	case Uint32:
		__obf_56127cd370854f0b, __obf_3aa78ad28f50ed46 := __obf_015afbee33862a01.uint32()
		return uint64(__obf_56127cd370854f0b), __obf_3aa78ad28f50ed46
	case Int32:
		__obf_56127cd370854f0b, __obf_3aa78ad28f50ed46 := __obf_015afbee33862a01.int32()
		return uint64(__obf_56127cd370854f0b), __obf_3aa78ad28f50ed46
	case Uint64, Int64:
		return __obf_015afbee33862a01.uint64()
	}
	return 0, fmt.Errorf("msgpack: invalid code=%x decoding uint64", __obf_145c7a7d25eea2bd)
}

// DecodeInt64 decodes msgpack int8/16/32/64 and uint8/16/32/64
// into Go int64.
func (__obf_015afbee33862a01 *Decoder) DecodeInt64() (int64, error) {
	__obf_145c7a7d25eea2bd, __obf_3aa78ad28f50ed46 := __obf_015afbee33862a01.__obf_4c5f8d154c9027ea()
	if __obf_3aa78ad28f50ed46 != nil {
		return 0, __obf_3aa78ad28f50ed46
	}
	return __obf_015afbee33862a01.int(__obf_145c7a7d25eea2bd)
}

func (__obf_015afbee33862a01 *Decoder) int(__obf_145c7a7d25eea2bd byte) (int64, error) {
	if __obf_145c7a7d25eea2bd == Nil {
		return 0, nil
	}
	if IsFixedNum(__obf_145c7a7d25eea2bd) {
		return int64(int8(__obf_145c7a7d25eea2bd)), nil
	}
	switch __obf_145c7a7d25eea2bd {
	case Uint8:
		__obf_56127cd370854f0b, __obf_3aa78ad28f50ed46 := __obf_015afbee33862a01.uint8()
		return int64(__obf_56127cd370854f0b), __obf_3aa78ad28f50ed46
	case Int8:
		__obf_56127cd370854f0b, __obf_3aa78ad28f50ed46 := __obf_015afbee33862a01.uint8()
		return int64(int8(__obf_56127cd370854f0b)), __obf_3aa78ad28f50ed46
	case Uint16:
		__obf_56127cd370854f0b, __obf_3aa78ad28f50ed46 := __obf_015afbee33862a01.uint16()
		return int64(__obf_56127cd370854f0b), __obf_3aa78ad28f50ed46
	case Int16:
		__obf_56127cd370854f0b, __obf_3aa78ad28f50ed46 := __obf_015afbee33862a01.uint16()
		return int64(int16(__obf_56127cd370854f0b)), __obf_3aa78ad28f50ed46
	case Uint32:
		__obf_56127cd370854f0b, __obf_3aa78ad28f50ed46 := __obf_015afbee33862a01.uint32()
		return int64(__obf_56127cd370854f0b), __obf_3aa78ad28f50ed46
	case Int32:
		__obf_56127cd370854f0b, __obf_3aa78ad28f50ed46 := __obf_015afbee33862a01.uint32()
		return int64(int32(__obf_56127cd370854f0b)), __obf_3aa78ad28f50ed46
	case Uint64, Int64:
		__obf_56127cd370854f0b, __obf_3aa78ad28f50ed46 := __obf_015afbee33862a01.uint64()
		return int64(__obf_56127cd370854f0b), __obf_3aa78ad28f50ed46
	}
	return 0, fmt.Errorf("msgpack: invalid code=%x decoding int64", __obf_145c7a7d25eea2bd)
}

func (__obf_015afbee33862a01 *Decoder) DecodeFloat32() (float32, error) {
	__obf_145c7a7d25eea2bd, __obf_3aa78ad28f50ed46 := __obf_015afbee33862a01.__obf_4c5f8d154c9027ea()
	if __obf_3aa78ad28f50ed46 != nil {
		return 0, __obf_3aa78ad28f50ed46
	}
	return __obf_015afbee33862a01.float32(__obf_145c7a7d25eea2bd)
}

func (__obf_015afbee33862a01 *Decoder) float32(__obf_145c7a7d25eea2bd byte) (float32, error) {
	if __obf_145c7a7d25eea2bd == Float {
		__obf_56127cd370854f0b, __obf_3aa78ad28f50ed46 := __obf_015afbee33862a01.uint32()
		if __obf_3aa78ad28f50ed46 != nil {
			return 0, __obf_3aa78ad28f50ed46
		}
		return math.Float32frombits(__obf_56127cd370854f0b), nil
	}
	__obf_56127cd370854f0b, __obf_3aa78ad28f50ed46 := __obf_015afbee33862a01.int(__obf_145c7a7d25eea2bd)
	if __obf_3aa78ad28f50ed46 != nil {
		return 0, fmt.Errorf("msgpack: invalid code=%x decoding float32", __obf_145c7a7d25eea2bd)
	}
	return float32(__obf_56127cd370854f0b), nil
}

// DecodeFloat64 decodes msgpack float32/64 into Go float64.
func (__obf_015afbee33862a01 *Decoder) DecodeFloat64() (float64, error) {
	__obf_145c7a7d25eea2bd, __obf_3aa78ad28f50ed46 := __obf_015afbee33862a01.__obf_4c5f8d154c9027ea()
	if __obf_3aa78ad28f50ed46 != nil {
		return 0, __obf_3aa78ad28f50ed46
	}
	return __obf_015afbee33862a01.float64(__obf_145c7a7d25eea2bd)
}

func (__obf_015afbee33862a01 *Decoder) float64(__obf_145c7a7d25eea2bd byte) (float64, error) {
	switch __obf_145c7a7d25eea2bd {
	case Float:
		__obf_56127cd370854f0b, __obf_3aa78ad28f50ed46 := __obf_015afbee33862a01.float32(__obf_145c7a7d25eea2bd)
		if __obf_3aa78ad28f50ed46 != nil {
			return 0, __obf_3aa78ad28f50ed46
		}
		return float64(__obf_56127cd370854f0b), nil
	case Double:
		__obf_56127cd370854f0b, __obf_3aa78ad28f50ed46 := __obf_015afbee33862a01.uint64()
		if __obf_3aa78ad28f50ed46 != nil {
			return 0, __obf_3aa78ad28f50ed46
		}
		return math.Float64frombits(__obf_56127cd370854f0b), nil
	}
	__obf_56127cd370854f0b, __obf_3aa78ad28f50ed46 := __obf_015afbee33862a01.int(__obf_145c7a7d25eea2bd)
	if __obf_3aa78ad28f50ed46 != nil {
		return 0, fmt.Errorf("msgpack: invalid code=%x decoding float32", __obf_145c7a7d25eea2bd)
	}
	return float64(__obf_56127cd370854f0b), nil
}

func (__obf_015afbee33862a01 *Decoder) DecodeUint() (uint, error) {
	__obf_56127cd370854f0b, __obf_3aa78ad28f50ed46 := __obf_015afbee33862a01.DecodeUint64()
	return uint(__obf_56127cd370854f0b), __obf_3aa78ad28f50ed46
}

func (__obf_015afbee33862a01 *Decoder) DecodeUint8() (uint8, error) {
	__obf_56127cd370854f0b, __obf_3aa78ad28f50ed46 := __obf_015afbee33862a01.DecodeUint64()
	return uint8(__obf_56127cd370854f0b), __obf_3aa78ad28f50ed46
}

func (__obf_015afbee33862a01 *Decoder) DecodeUint16() (uint16, error) {
	__obf_56127cd370854f0b, __obf_3aa78ad28f50ed46 := __obf_015afbee33862a01.DecodeUint64()
	return uint16(__obf_56127cd370854f0b), __obf_3aa78ad28f50ed46
}

func (__obf_015afbee33862a01 *Decoder) DecodeUint32() (uint32, error) {
	__obf_56127cd370854f0b, __obf_3aa78ad28f50ed46 := __obf_015afbee33862a01.DecodeUint64()
	return uint32(__obf_56127cd370854f0b), __obf_3aa78ad28f50ed46
}

func (__obf_015afbee33862a01 *Decoder) DecodeInt() (int, error) {
	__obf_56127cd370854f0b, __obf_3aa78ad28f50ed46 := __obf_015afbee33862a01.DecodeInt64()
	return int(__obf_56127cd370854f0b), __obf_3aa78ad28f50ed46
}

func (__obf_015afbee33862a01 *Decoder) DecodeInt8() (int8, error) {
	__obf_56127cd370854f0b, __obf_3aa78ad28f50ed46 := __obf_015afbee33862a01.DecodeInt64()
	return int8(__obf_56127cd370854f0b), __obf_3aa78ad28f50ed46
}

func (__obf_015afbee33862a01 *Decoder) DecodeInt16() (int16, error) {
	__obf_56127cd370854f0b, __obf_3aa78ad28f50ed46 := __obf_015afbee33862a01.DecodeInt64()
	return int16(__obf_56127cd370854f0b), __obf_3aa78ad28f50ed46
}

func (__obf_015afbee33862a01 *Decoder) DecodeInt32() (int32, error) {
	__obf_56127cd370854f0b, __obf_3aa78ad28f50ed46 := __obf_015afbee33862a01.DecodeInt64()
	return int32(__obf_56127cd370854f0b), __obf_3aa78ad28f50ed46
}

func __obf_fe8f515e5461c435(__obf_015afbee33862a01 *Decoder, __obf_85d270343a0dfe14 reflect.Value) error {
	__obf_40a54345a49c3cd5, __obf_3aa78ad28f50ed46 := __obf_015afbee33862a01.DecodeFloat32()
	if __obf_3aa78ad28f50ed46 != nil {
		return __obf_3aa78ad28f50ed46
	}
	__obf_85d270343a0dfe14.
		SetFloat(float64(__obf_40a54345a49c3cd5))
	return nil
}

func __obf_0ccce5c30ace0e82(__obf_015afbee33862a01 *Decoder, __obf_85d270343a0dfe14 reflect.Value) error {
	__obf_40a54345a49c3cd5, __obf_3aa78ad28f50ed46 := __obf_015afbee33862a01.DecodeFloat64()
	if __obf_3aa78ad28f50ed46 != nil {
		return __obf_3aa78ad28f50ed46
	}
	__obf_85d270343a0dfe14.
		SetFloat(__obf_40a54345a49c3cd5)
	return nil
}

func __obf_69ca6ad8831c417d(__obf_015afbee33862a01 *Decoder, __obf_85d270343a0dfe14 reflect.Value) error {
	__obf_56127cd370854f0b, __obf_3aa78ad28f50ed46 := __obf_015afbee33862a01.DecodeInt64()
	if __obf_3aa78ad28f50ed46 != nil {
		return __obf_3aa78ad28f50ed46
	}
	__obf_85d270343a0dfe14.
		SetInt(__obf_56127cd370854f0b)
	return nil
}

func __obf_e4e267de4597ac4a(__obf_015afbee33862a01 *Decoder, __obf_85d270343a0dfe14 reflect.Value) error {
	__obf_56127cd370854f0b, __obf_3aa78ad28f50ed46 := __obf_015afbee33862a01.DecodeUint64()
	if __obf_3aa78ad28f50ed46 != nil {
		return __obf_3aa78ad28f50ed46
	}
	__obf_85d270343a0dfe14.
		SetUint(__obf_56127cd370854f0b)
	return nil
}
