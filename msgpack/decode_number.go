package __obf_fd770cb9675903b0

import (
	"fmt"
	"math"
	"reflect"
)

func (__obf_5d389b660eefb08c *Decoder) __obf_d93e03df64d057a0(__obf_d774e4844c74bfe9 int) error {
	_, __obf_45342a3a754d12f5 := __obf_5d389b660eefb08c.__obf_31c39140d4d66749(__obf_d774e4844c74bfe9)
	return __obf_45342a3a754d12f5
}

func (__obf_5d389b660eefb08c *Decoder) uint8() (uint8, error) {
	__obf_4148125b350dfea2, __obf_45342a3a754d12f5 := __obf_5d389b660eefb08c.__obf_4786096a37bd32ce()
	if __obf_45342a3a754d12f5 != nil {
		return 0, __obf_45342a3a754d12f5
	}
	return __obf_4148125b350dfea2, nil
}

func (__obf_5d389b660eefb08c *Decoder) int8() (int8, error) {
	__obf_d774e4844c74bfe9, __obf_45342a3a754d12f5 := __obf_5d389b660eefb08c.uint8()
	return int8(__obf_d774e4844c74bfe9), __obf_45342a3a754d12f5
}

func (__obf_5d389b660eefb08c *Decoder) uint16() (uint16, error) {
	__obf_94333af0f0facd65, __obf_45342a3a754d12f5 := __obf_5d389b660eefb08c.__obf_31c39140d4d66749(2)
	if __obf_45342a3a754d12f5 != nil {
		return 0, __obf_45342a3a754d12f5
	}
	return (uint16(__obf_94333af0f0facd65[0]) << 8) | uint16(__obf_94333af0f0facd65[1]), nil
}

func (__obf_5d389b660eefb08c *Decoder) int16() (int16, error) {
	__obf_d774e4844c74bfe9, __obf_45342a3a754d12f5 := __obf_5d389b660eefb08c.uint16()
	return int16(__obf_d774e4844c74bfe9), __obf_45342a3a754d12f5
}

func (__obf_5d389b660eefb08c *Decoder) uint32() (uint32, error) {
	__obf_94333af0f0facd65, __obf_45342a3a754d12f5 := __obf_5d389b660eefb08c.__obf_31c39140d4d66749(4)
	if __obf_45342a3a754d12f5 != nil {
		return 0, __obf_45342a3a754d12f5
	}
	__obf_d774e4844c74bfe9 := (uint32(__obf_94333af0f0facd65[0]) << 24) |
		(uint32(__obf_94333af0f0facd65[1]) << 16) |
		(uint32(__obf_94333af0f0facd65[2]) << 8) |
		uint32(__obf_94333af0f0facd65[3])
	return __obf_d774e4844c74bfe9, nil
}

func (__obf_5d389b660eefb08c *Decoder) int32() (int32, error) {
	__obf_d774e4844c74bfe9, __obf_45342a3a754d12f5 := __obf_5d389b660eefb08c.uint32()
	return int32(__obf_d774e4844c74bfe9), __obf_45342a3a754d12f5
}

func (__obf_5d389b660eefb08c *Decoder) uint64() (uint64, error) {
	__obf_94333af0f0facd65, __obf_45342a3a754d12f5 := __obf_5d389b660eefb08c.__obf_31c39140d4d66749(8)
	if __obf_45342a3a754d12f5 != nil {
		return 0, __obf_45342a3a754d12f5
	}
	__obf_d774e4844c74bfe9 := (uint64(__obf_94333af0f0facd65[0]) << 56) |
		(uint64(__obf_94333af0f0facd65[1]) << 48) |
		(uint64(__obf_94333af0f0facd65[2]) << 40) |
		(uint64(__obf_94333af0f0facd65[3]) << 32) |
		(uint64(__obf_94333af0f0facd65[4]) << 24) |
		(uint64(__obf_94333af0f0facd65[5]) << 16) |
		(uint64(__obf_94333af0f0facd65[6]) << 8) |
		uint64(__obf_94333af0f0facd65[7])
	return __obf_d774e4844c74bfe9, nil
}

func (__obf_5d389b660eefb08c *Decoder) int64() (int64, error) {
	__obf_d774e4844c74bfe9, __obf_45342a3a754d12f5 := __obf_5d389b660eefb08c.uint64()
	return int64(__obf_d774e4844c74bfe9), __obf_45342a3a754d12f5
}

// DecodeUint64 decodes msgpack int8/16/32/64 and uint8/16/32/64
// into Go uint64.
func (__obf_5d389b660eefb08c *Decoder) DecodeUint64() (uint64, error) {
	__obf_4148125b350dfea2, __obf_45342a3a754d12f5 := __obf_5d389b660eefb08c.__obf_4786096a37bd32ce()
	if __obf_45342a3a754d12f5 != nil {
		return 0, __obf_45342a3a754d12f5
	}
	return __obf_5d389b660eefb08c.uint(__obf_4148125b350dfea2)
}

func (__obf_5d389b660eefb08c *Decoder) uint(__obf_4148125b350dfea2 byte) (uint64, error) {
	if __obf_4148125b350dfea2 == Nil {
		return 0, nil
	}
	if IsFixedNum(__obf_4148125b350dfea2) {
		return uint64(int8(__obf_4148125b350dfea2)), nil
	}
	switch __obf_4148125b350dfea2 {
	case Uint8:
		__obf_d774e4844c74bfe9, __obf_45342a3a754d12f5 := __obf_5d389b660eefb08c.uint8()
		return uint64(__obf_d774e4844c74bfe9), __obf_45342a3a754d12f5
	case Int8:
		__obf_d774e4844c74bfe9, __obf_45342a3a754d12f5 := __obf_5d389b660eefb08c.int8()
		return uint64(__obf_d774e4844c74bfe9), __obf_45342a3a754d12f5
	case Uint16:
		__obf_d774e4844c74bfe9, __obf_45342a3a754d12f5 := __obf_5d389b660eefb08c.uint16()
		return uint64(__obf_d774e4844c74bfe9), __obf_45342a3a754d12f5
	case Int16:
		__obf_d774e4844c74bfe9, __obf_45342a3a754d12f5 := __obf_5d389b660eefb08c.int16()
		return uint64(__obf_d774e4844c74bfe9), __obf_45342a3a754d12f5
	case Uint32:
		__obf_d774e4844c74bfe9, __obf_45342a3a754d12f5 := __obf_5d389b660eefb08c.uint32()
		return uint64(__obf_d774e4844c74bfe9), __obf_45342a3a754d12f5
	case Int32:
		__obf_d774e4844c74bfe9, __obf_45342a3a754d12f5 := __obf_5d389b660eefb08c.int32()
		return uint64(__obf_d774e4844c74bfe9), __obf_45342a3a754d12f5
	case Uint64, Int64:
		return __obf_5d389b660eefb08c.uint64()
	}
	return 0, fmt.Errorf("msgpack: invalid code=%x decoding uint64", __obf_4148125b350dfea2)
}

// DecodeInt64 decodes msgpack int8/16/32/64 and uint8/16/32/64
// into Go int64.
func (__obf_5d389b660eefb08c *Decoder) DecodeInt64() (int64, error) {
	__obf_4148125b350dfea2, __obf_45342a3a754d12f5 := __obf_5d389b660eefb08c.__obf_4786096a37bd32ce()
	if __obf_45342a3a754d12f5 != nil {
		return 0, __obf_45342a3a754d12f5
	}
	return __obf_5d389b660eefb08c.int(__obf_4148125b350dfea2)
}

func (__obf_5d389b660eefb08c *Decoder) int(__obf_4148125b350dfea2 byte) (int64, error) {
	if __obf_4148125b350dfea2 == Nil {
		return 0, nil
	}
	if IsFixedNum(__obf_4148125b350dfea2) {
		return int64(int8(__obf_4148125b350dfea2)), nil
	}
	switch __obf_4148125b350dfea2 {
	case Uint8:
		__obf_d774e4844c74bfe9, __obf_45342a3a754d12f5 := __obf_5d389b660eefb08c.uint8()
		return int64(__obf_d774e4844c74bfe9), __obf_45342a3a754d12f5
	case Int8:
		__obf_d774e4844c74bfe9, __obf_45342a3a754d12f5 := __obf_5d389b660eefb08c.uint8()
		return int64(int8(__obf_d774e4844c74bfe9)), __obf_45342a3a754d12f5
	case Uint16:
		__obf_d774e4844c74bfe9, __obf_45342a3a754d12f5 := __obf_5d389b660eefb08c.uint16()
		return int64(__obf_d774e4844c74bfe9), __obf_45342a3a754d12f5
	case Int16:
		__obf_d774e4844c74bfe9, __obf_45342a3a754d12f5 := __obf_5d389b660eefb08c.uint16()
		return int64(int16(__obf_d774e4844c74bfe9)), __obf_45342a3a754d12f5
	case Uint32:
		__obf_d774e4844c74bfe9, __obf_45342a3a754d12f5 := __obf_5d389b660eefb08c.uint32()
		return int64(__obf_d774e4844c74bfe9), __obf_45342a3a754d12f5
	case Int32:
		__obf_d774e4844c74bfe9, __obf_45342a3a754d12f5 := __obf_5d389b660eefb08c.uint32()
		return int64(int32(__obf_d774e4844c74bfe9)), __obf_45342a3a754d12f5
	case Uint64, Int64:
		__obf_d774e4844c74bfe9, __obf_45342a3a754d12f5 := __obf_5d389b660eefb08c.uint64()
		return int64(__obf_d774e4844c74bfe9), __obf_45342a3a754d12f5
	}
	return 0, fmt.Errorf("msgpack: invalid code=%x decoding int64", __obf_4148125b350dfea2)
}

func (__obf_5d389b660eefb08c *Decoder) DecodeFloat32() (float32, error) {
	__obf_4148125b350dfea2, __obf_45342a3a754d12f5 := __obf_5d389b660eefb08c.__obf_4786096a37bd32ce()
	if __obf_45342a3a754d12f5 != nil {
		return 0, __obf_45342a3a754d12f5
	}
	return __obf_5d389b660eefb08c.float32(__obf_4148125b350dfea2)
}

func (__obf_5d389b660eefb08c *Decoder) float32(__obf_4148125b350dfea2 byte) (float32, error) {
	if __obf_4148125b350dfea2 == Float {
		__obf_d774e4844c74bfe9, __obf_45342a3a754d12f5 := __obf_5d389b660eefb08c.uint32()
		if __obf_45342a3a754d12f5 != nil {
			return 0, __obf_45342a3a754d12f5
		}
		return math.Float32frombits(__obf_d774e4844c74bfe9), nil
	}
	__obf_d774e4844c74bfe9, __obf_45342a3a754d12f5 := __obf_5d389b660eefb08c.int(__obf_4148125b350dfea2)
	if __obf_45342a3a754d12f5 != nil {
		return 0, fmt.Errorf("msgpack: invalid code=%x decoding float32", __obf_4148125b350dfea2)
	}
	return float32(__obf_d774e4844c74bfe9), nil
}

// DecodeFloat64 decodes msgpack float32/64 into Go float64.
func (__obf_5d389b660eefb08c *Decoder) DecodeFloat64() (float64, error) {
	__obf_4148125b350dfea2, __obf_45342a3a754d12f5 := __obf_5d389b660eefb08c.__obf_4786096a37bd32ce()
	if __obf_45342a3a754d12f5 != nil {
		return 0, __obf_45342a3a754d12f5
	}
	return __obf_5d389b660eefb08c.float64(__obf_4148125b350dfea2)
}

func (__obf_5d389b660eefb08c *Decoder) float64(__obf_4148125b350dfea2 byte) (float64, error) {
	switch __obf_4148125b350dfea2 {
	case Float:
		__obf_d774e4844c74bfe9, __obf_45342a3a754d12f5 := __obf_5d389b660eefb08c.float32(__obf_4148125b350dfea2)
		if __obf_45342a3a754d12f5 != nil {
			return 0, __obf_45342a3a754d12f5
		}
		return float64(__obf_d774e4844c74bfe9), nil
	case Double:
		__obf_d774e4844c74bfe9, __obf_45342a3a754d12f5 := __obf_5d389b660eefb08c.uint64()
		if __obf_45342a3a754d12f5 != nil {
			return 0, __obf_45342a3a754d12f5
		}
		return math.Float64frombits(__obf_d774e4844c74bfe9), nil
	}
	__obf_d774e4844c74bfe9, __obf_45342a3a754d12f5 := __obf_5d389b660eefb08c.int(__obf_4148125b350dfea2)
	if __obf_45342a3a754d12f5 != nil {
		return 0, fmt.Errorf("msgpack: invalid code=%x decoding float32", __obf_4148125b350dfea2)
	}
	return float64(__obf_d774e4844c74bfe9), nil
}

func (__obf_5d389b660eefb08c *Decoder) DecodeUint() (uint, error) {
	__obf_d774e4844c74bfe9, __obf_45342a3a754d12f5 := __obf_5d389b660eefb08c.DecodeUint64()
	return uint(__obf_d774e4844c74bfe9), __obf_45342a3a754d12f5
}

func (__obf_5d389b660eefb08c *Decoder) DecodeUint8() (uint8, error) {
	__obf_d774e4844c74bfe9, __obf_45342a3a754d12f5 := __obf_5d389b660eefb08c.DecodeUint64()
	return uint8(__obf_d774e4844c74bfe9), __obf_45342a3a754d12f5
}

func (__obf_5d389b660eefb08c *Decoder) DecodeUint16() (uint16, error) {
	__obf_d774e4844c74bfe9, __obf_45342a3a754d12f5 := __obf_5d389b660eefb08c.DecodeUint64()
	return uint16(__obf_d774e4844c74bfe9), __obf_45342a3a754d12f5
}

func (__obf_5d389b660eefb08c *Decoder) DecodeUint32() (uint32, error) {
	__obf_d774e4844c74bfe9, __obf_45342a3a754d12f5 := __obf_5d389b660eefb08c.DecodeUint64()
	return uint32(__obf_d774e4844c74bfe9), __obf_45342a3a754d12f5
}

func (__obf_5d389b660eefb08c *Decoder) DecodeInt() (int, error) {
	__obf_d774e4844c74bfe9, __obf_45342a3a754d12f5 := __obf_5d389b660eefb08c.DecodeInt64()
	return int(__obf_d774e4844c74bfe9), __obf_45342a3a754d12f5
}

func (__obf_5d389b660eefb08c *Decoder) DecodeInt8() (int8, error) {
	__obf_d774e4844c74bfe9, __obf_45342a3a754d12f5 := __obf_5d389b660eefb08c.DecodeInt64()
	return int8(__obf_d774e4844c74bfe9), __obf_45342a3a754d12f5
}

func (__obf_5d389b660eefb08c *Decoder) DecodeInt16() (int16, error) {
	__obf_d774e4844c74bfe9, __obf_45342a3a754d12f5 := __obf_5d389b660eefb08c.DecodeInt64()
	return int16(__obf_d774e4844c74bfe9), __obf_45342a3a754d12f5
}

func (__obf_5d389b660eefb08c *Decoder) DecodeInt32() (int32, error) {
	__obf_d774e4844c74bfe9, __obf_45342a3a754d12f5 := __obf_5d389b660eefb08c.DecodeInt64()
	return int32(__obf_d774e4844c74bfe9), __obf_45342a3a754d12f5
}

func __obf_028e0e88467a2c89(__obf_5d389b660eefb08c *Decoder, __obf_f328a048f76b7256 reflect.Value) error {
	__obf_e25b9e550ea05af3, __obf_45342a3a754d12f5 := __obf_5d389b660eefb08c.DecodeFloat32()
	if __obf_45342a3a754d12f5 != nil {
		return __obf_45342a3a754d12f5
	}
	__obf_f328a048f76b7256.
		SetFloat(float64(__obf_e25b9e550ea05af3))
	return nil
}

func __obf_fdb9efbf3bc19d4f(__obf_5d389b660eefb08c *Decoder, __obf_f328a048f76b7256 reflect.Value) error {
	__obf_e25b9e550ea05af3, __obf_45342a3a754d12f5 := __obf_5d389b660eefb08c.DecodeFloat64()
	if __obf_45342a3a754d12f5 != nil {
		return __obf_45342a3a754d12f5
	}
	__obf_f328a048f76b7256.
		SetFloat(__obf_e25b9e550ea05af3)
	return nil
}

func __obf_7801128403b4eb1b(__obf_5d389b660eefb08c *Decoder, __obf_f328a048f76b7256 reflect.Value) error {
	__obf_d774e4844c74bfe9, __obf_45342a3a754d12f5 := __obf_5d389b660eefb08c.DecodeInt64()
	if __obf_45342a3a754d12f5 != nil {
		return __obf_45342a3a754d12f5
	}
	__obf_f328a048f76b7256.
		SetInt(__obf_d774e4844c74bfe9)
	return nil
}

func __obf_f5ab33312c20fcd4(__obf_5d389b660eefb08c *Decoder, __obf_f328a048f76b7256 reflect.Value) error {
	__obf_d774e4844c74bfe9, __obf_45342a3a754d12f5 := __obf_5d389b660eefb08c.DecodeUint64()
	if __obf_45342a3a754d12f5 != nil {
		return __obf_45342a3a754d12f5
	}
	__obf_f328a048f76b7256.
		SetUint(__obf_d774e4844c74bfe9)
	return nil
}
