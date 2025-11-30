package __obf_3e0c303610a19bd4

import (
	"fmt"
	"math"
	"reflect"
)

func (__obf_dc35117108ba8439 *Decoder) __obf_19f294acbba68c47(__obf_4909ae60ffbb8e53 int) error {
	_, __obf_8882f6cf6e378ded := __obf_dc35117108ba8439.__obf_b06a4f273442ca29(__obf_4909ae60ffbb8e53)
	return __obf_8882f6cf6e378ded
}

func (__obf_dc35117108ba8439 *Decoder) uint8() (uint8, error) {
	__obf_e46289218af709bf, __obf_8882f6cf6e378ded := __obf_dc35117108ba8439.__obf_5b7d85f9093902c3()
	if __obf_8882f6cf6e378ded != nil {
		return 0, __obf_8882f6cf6e378ded
	}
	return __obf_e46289218af709bf, nil
}

func (__obf_dc35117108ba8439 *Decoder) int8() (int8, error) {
	__obf_4909ae60ffbb8e53, __obf_8882f6cf6e378ded := __obf_dc35117108ba8439.uint8()
	return int8(__obf_4909ae60ffbb8e53), __obf_8882f6cf6e378ded
}

func (__obf_dc35117108ba8439 *Decoder) uint16() (uint16, error) {
	__obf_11bcc66cde095c11, __obf_8882f6cf6e378ded := __obf_dc35117108ba8439.__obf_b06a4f273442ca29(2)
	if __obf_8882f6cf6e378ded != nil {
		return 0, __obf_8882f6cf6e378ded
	}
	return (uint16(__obf_11bcc66cde095c11[0]) << 8) | uint16(__obf_11bcc66cde095c11[1]), nil
}

func (__obf_dc35117108ba8439 *Decoder) int16() (int16, error) {
	__obf_4909ae60ffbb8e53, __obf_8882f6cf6e378ded := __obf_dc35117108ba8439.uint16()
	return int16(__obf_4909ae60ffbb8e53), __obf_8882f6cf6e378ded
}

func (__obf_dc35117108ba8439 *Decoder) uint32() (uint32, error) {
	__obf_11bcc66cde095c11, __obf_8882f6cf6e378ded := __obf_dc35117108ba8439.__obf_b06a4f273442ca29(4)
	if __obf_8882f6cf6e378ded != nil {
		return 0, __obf_8882f6cf6e378ded
	}
	__obf_4909ae60ffbb8e53 := (uint32(__obf_11bcc66cde095c11[0]) << 24) |
		(uint32(__obf_11bcc66cde095c11[1]) << 16) |
		(uint32(__obf_11bcc66cde095c11[2]) << 8) |
		uint32(__obf_11bcc66cde095c11[3])
	return __obf_4909ae60ffbb8e53, nil
}

func (__obf_dc35117108ba8439 *Decoder) int32() (int32, error) {
	__obf_4909ae60ffbb8e53, __obf_8882f6cf6e378ded := __obf_dc35117108ba8439.uint32()
	return int32(__obf_4909ae60ffbb8e53), __obf_8882f6cf6e378ded
}

func (__obf_dc35117108ba8439 *Decoder) uint64() (uint64, error) {
	__obf_11bcc66cde095c11, __obf_8882f6cf6e378ded := __obf_dc35117108ba8439.__obf_b06a4f273442ca29(8)
	if __obf_8882f6cf6e378ded != nil {
		return 0, __obf_8882f6cf6e378ded
	}
	__obf_4909ae60ffbb8e53 := (uint64(__obf_11bcc66cde095c11[0]) << 56) |
		(uint64(__obf_11bcc66cde095c11[1]) << 48) |
		(uint64(__obf_11bcc66cde095c11[2]) << 40) |
		(uint64(__obf_11bcc66cde095c11[3]) << 32) |
		(uint64(__obf_11bcc66cde095c11[4]) << 24) |
		(uint64(__obf_11bcc66cde095c11[5]) << 16) |
		(uint64(__obf_11bcc66cde095c11[6]) << 8) |
		uint64(__obf_11bcc66cde095c11[7])
	return __obf_4909ae60ffbb8e53, nil
}

func (__obf_dc35117108ba8439 *Decoder) int64() (int64, error) {
	__obf_4909ae60ffbb8e53, __obf_8882f6cf6e378ded := __obf_dc35117108ba8439.uint64()
	return int64(__obf_4909ae60ffbb8e53), __obf_8882f6cf6e378ded
}

// DecodeUint64 decodes msgpack int8/16/32/64 and uint8/16/32/64
// into Go uint64.
func (__obf_dc35117108ba8439 *Decoder) DecodeUint64() (uint64, error) {
	__obf_e46289218af709bf, __obf_8882f6cf6e378ded := __obf_dc35117108ba8439.__obf_5b7d85f9093902c3()
	if __obf_8882f6cf6e378ded != nil {
		return 0, __obf_8882f6cf6e378ded
	}
	return __obf_dc35117108ba8439.uint(__obf_e46289218af709bf)
}

func (__obf_dc35117108ba8439 *Decoder) uint(__obf_e46289218af709bf byte) (uint64, error) {
	if __obf_e46289218af709bf == Nil {
		return 0, nil
	}
	if IsFixedNum(__obf_e46289218af709bf) {
		return uint64(int8(__obf_e46289218af709bf)), nil
	}
	switch __obf_e46289218af709bf {
	case Uint8:
		__obf_4909ae60ffbb8e53, __obf_8882f6cf6e378ded := __obf_dc35117108ba8439.uint8()
		return uint64(__obf_4909ae60ffbb8e53), __obf_8882f6cf6e378ded
	case Int8:
		__obf_4909ae60ffbb8e53, __obf_8882f6cf6e378ded := __obf_dc35117108ba8439.int8()
		return uint64(__obf_4909ae60ffbb8e53), __obf_8882f6cf6e378ded
	case Uint16:
		__obf_4909ae60ffbb8e53, __obf_8882f6cf6e378ded := __obf_dc35117108ba8439.uint16()
		return uint64(__obf_4909ae60ffbb8e53), __obf_8882f6cf6e378ded
	case Int16:
		__obf_4909ae60ffbb8e53, __obf_8882f6cf6e378ded := __obf_dc35117108ba8439.int16()
		return uint64(__obf_4909ae60ffbb8e53), __obf_8882f6cf6e378ded
	case Uint32:
		__obf_4909ae60ffbb8e53, __obf_8882f6cf6e378ded := __obf_dc35117108ba8439.uint32()
		return uint64(__obf_4909ae60ffbb8e53), __obf_8882f6cf6e378ded
	case Int32:
		__obf_4909ae60ffbb8e53, __obf_8882f6cf6e378ded := __obf_dc35117108ba8439.int32()
		return uint64(__obf_4909ae60ffbb8e53), __obf_8882f6cf6e378ded
	case Uint64, Int64:
		return __obf_dc35117108ba8439.uint64()
	}
	return 0, fmt.Errorf("msgpack: invalid code=%x decoding uint64", __obf_e46289218af709bf)
}

// DecodeInt64 decodes msgpack int8/16/32/64 and uint8/16/32/64
// into Go int64.
func (__obf_dc35117108ba8439 *Decoder) DecodeInt64() (int64, error) {
	__obf_e46289218af709bf, __obf_8882f6cf6e378ded := __obf_dc35117108ba8439.__obf_5b7d85f9093902c3()
	if __obf_8882f6cf6e378ded != nil {
		return 0, __obf_8882f6cf6e378ded
	}
	return __obf_dc35117108ba8439.int(__obf_e46289218af709bf)
}

func (__obf_dc35117108ba8439 *Decoder) int(__obf_e46289218af709bf byte) (int64, error) {
	if __obf_e46289218af709bf == Nil {
		return 0, nil
	}
	if IsFixedNum(__obf_e46289218af709bf) {
		return int64(int8(__obf_e46289218af709bf)), nil
	}
	switch __obf_e46289218af709bf {
	case Uint8:
		__obf_4909ae60ffbb8e53, __obf_8882f6cf6e378ded := __obf_dc35117108ba8439.uint8()
		return int64(__obf_4909ae60ffbb8e53), __obf_8882f6cf6e378ded
	case Int8:
		__obf_4909ae60ffbb8e53, __obf_8882f6cf6e378ded := __obf_dc35117108ba8439.uint8()
		return int64(int8(__obf_4909ae60ffbb8e53)), __obf_8882f6cf6e378ded
	case Uint16:
		__obf_4909ae60ffbb8e53, __obf_8882f6cf6e378ded := __obf_dc35117108ba8439.uint16()
		return int64(__obf_4909ae60ffbb8e53), __obf_8882f6cf6e378ded
	case Int16:
		__obf_4909ae60ffbb8e53, __obf_8882f6cf6e378ded := __obf_dc35117108ba8439.uint16()
		return int64(int16(__obf_4909ae60ffbb8e53)), __obf_8882f6cf6e378ded
	case Uint32:
		__obf_4909ae60ffbb8e53, __obf_8882f6cf6e378ded := __obf_dc35117108ba8439.uint32()
		return int64(__obf_4909ae60ffbb8e53), __obf_8882f6cf6e378ded
	case Int32:
		__obf_4909ae60ffbb8e53, __obf_8882f6cf6e378ded := __obf_dc35117108ba8439.uint32()
		return int64(int32(__obf_4909ae60ffbb8e53)), __obf_8882f6cf6e378ded
	case Uint64, Int64:
		__obf_4909ae60ffbb8e53, __obf_8882f6cf6e378ded := __obf_dc35117108ba8439.uint64()
		return int64(__obf_4909ae60ffbb8e53), __obf_8882f6cf6e378ded
	}
	return 0, fmt.Errorf("msgpack: invalid code=%x decoding int64", __obf_e46289218af709bf)
}

func (__obf_dc35117108ba8439 *Decoder) DecodeFloat32() (float32, error) {
	__obf_e46289218af709bf, __obf_8882f6cf6e378ded := __obf_dc35117108ba8439.__obf_5b7d85f9093902c3()
	if __obf_8882f6cf6e378ded != nil {
		return 0, __obf_8882f6cf6e378ded
	}
	return __obf_dc35117108ba8439.float32(__obf_e46289218af709bf)
}

func (__obf_dc35117108ba8439 *Decoder) float32(__obf_e46289218af709bf byte) (float32, error) {
	if __obf_e46289218af709bf == Float {
		__obf_4909ae60ffbb8e53, __obf_8882f6cf6e378ded := __obf_dc35117108ba8439.uint32()
		if __obf_8882f6cf6e378ded != nil {
			return 0, __obf_8882f6cf6e378ded
		}
		return math.Float32frombits(__obf_4909ae60ffbb8e53), nil
	}
	__obf_4909ae60ffbb8e53, __obf_8882f6cf6e378ded := __obf_dc35117108ba8439.int(__obf_e46289218af709bf)
	if __obf_8882f6cf6e378ded != nil {
		return 0, fmt.Errorf("msgpack: invalid code=%x decoding float32", __obf_e46289218af709bf)
	}
	return float32(__obf_4909ae60ffbb8e53), nil
}

// DecodeFloat64 decodes msgpack float32/64 into Go float64.
func (__obf_dc35117108ba8439 *Decoder) DecodeFloat64() (float64, error) {
	__obf_e46289218af709bf, __obf_8882f6cf6e378ded := __obf_dc35117108ba8439.__obf_5b7d85f9093902c3()
	if __obf_8882f6cf6e378ded != nil {
		return 0, __obf_8882f6cf6e378ded
	}
	return __obf_dc35117108ba8439.float64(__obf_e46289218af709bf)
}

func (__obf_dc35117108ba8439 *Decoder) float64(__obf_e46289218af709bf byte) (float64, error) {
	switch __obf_e46289218af709bf {
	case Float:
		__obf_4909ae60ffbb8e53, __obf_8882f6cf6e378ded := __obf_dc35117108ba8439.float32(__obf_e46289218af709bf)
		if __obf_8882f6cf6e378ded != nil {
			return 0, __obf_8882f6cf6e378ded
		}
		return float64(__obf_4909ae60ffbb8e53), nil
	case Double:
		__obf_4909ae60ffbb8e53, __obf_8882f6cf6e378ded := __obf_dc35117108ba8439.uint64()
		if __obf_8882f6cf6e378ded != nil {
			return 0, __obf_8882f6cf6e378ded
		}
		return math.Float64frombits(__obf_4909ae60ffbb8e53), nil
	}
	__obf_4909ae60ffbb8e53, __obf_8882f6cf6e378ded := __obf_dc35117108ba8439.int(__obf_e46289218af709bf)
	if __obf_8882f6cf6e378ded != nil {
		return 0, fmt.Errorf("msgpack: invalid code=%x decoding float32", __obf_e46289218af709bf)
	}
	return float64(__obf_4909ae60ffbb8e53), nil
}

func (__obf_dc35117108ba8439 *Decoder) DecodeUint() (uint, error) {
	__obf_4909ae60ffbb8e53, __obf_8882f6cf6e378ded := __obf_dc35117108ba8439.DecodeUint64()
	return uint(__obf_4909ae60ffbb8e53), __obf_8882f6cf6e378ded
}

func (__obf_dc35117108ba8439 *Decoder) DecodeUint8() (uint8, error) {
	__obf_4909ae60ffbb8e53, __obf_8882f6cf6e378ded := __obf_dc35117108ba8439.DecodeUint64()
	return uint8(__obf_4909ae60ffbb8e53), __obf_8882f6cf6e378ded
}

func (__obf_dc35117108ba8439 *Decoder) DecodeUint16() (uint16, error) {
	__obf_4909ae60ffbb8e53, __obf_8882f6cf6e378ded := __obf_dc35117108ba8439.DecodeUint64()
	return uint16(__obf_4909ae60ffbb8e53), __obf_8882f6cf6e378ded
}

func (__obf_dc35117108ba8439 *Decoder) DecodeUint32() (uint32, error) {
	__obf_4909ae60ffbb8e53, __obf_8882f6cf6e378ded := __obf_dc35117108ba8439.DecodeUint64()
	return uint32(__obf_4909ae60ffbb8e53), __obf_8882f6cf6e378ded
}

func (__obf_dc35117108ba8439 *Decoder) DecodeInt() (int, error) {
	__obf_4909ae60ffbb8e53, __obf_8882f6cf6e378ded := __obf_dc35117108ba8439.DecodeInt64()
	return int(__obf_4909ae60ffbb8e53), __obf_8882f6cf6e378ded
}

func (__obf_dc35117108ba8439 *Decoder) DecodeInt8() (int8, error) {
	__obf_4909ae60ffbb8e53, __obf_8882f6cf6e378ded := __obf_dc35117108ba8439.DecodeInt64()
	return int8(__obf_4909ae60ffbb8e53), __obf_8882f6cf6e378ded
}

func (__obf_dc35117108ba8439 *Decoder) DecodeInt16() (int16, error) {
	__obf_4909ae60ffbb8e53, __obf_8882f6cf6e378ded := __obf_dc35117108ba8439.DecodeInt64()
	return int16(__obf_4909ae60ffbb8e53), __obf_8882f6cf6e378ded
}

func (__obf_dc35117108ba8439 *Decoder) DecodeInt32() (int32, error) {
	__obf_4909ae60ffbb8e53, __obf_8882f6cf6e378ded := __obf_dc35117108ba8439.DecodeInt64()
	return int32(__obf_4909ae60ffbb8e53), __obf_8882f6cf6e378ded
}

func __obf_a0985953c5ea413f(__obf_dc35117108ba8439 *Decoder, __obf_63bbcee86d44fdde reflect.Value) error {
	__obf_4add40b0f5dc6b95, __obf_8882f6cf6e378ded := __obf_dc35117108ba8439.DecodeFloat32()
	if __obf_8882f6cf6e378ded != nil {
		return __obf_8882f6cf6e378ded
	}
	__obf_63bbcee86d44fdde.
		SetFloat(float64(__obf_4add40b0f5dc6b95))
	return nil
}

func __obf_f008a12b21b1cd5f(__obf_dc35117108ba8439 *Decoder, __obf_63bbcee86d44fdde reflect.Value) error {
	__obf_4add40b0f5dc6b95, __obf_8882f6cf6e378ded := __obf_dc35117108ba8439.DecodeFloat64()
	if __obf_8882f6cf6e378ded != nil {
		return __obf_8882f6cf6e378ded
	}
	__obf_63bbcee86d44fdde.
		SetFloat(__obf_4add40b0f5dc6b95)
	return nil
}

func __obf_2e7a83f8fd124752(__obf_dc35117108ba8439 *Decoder, __obf_63bbcee86d44fdde reflect.Value) error {
	__obf_4909ae60ffbb8e53, __obf_8882f6cf6e378ded := __obf_dc35117108ba8439.DecodeInt64()
	if __obf_8882f6cf6e378ded != nil {
		return __obf_8882f6cf6e378ded
	}
	__obf_63bbcee86d44fdde.
		SetInt(__obf_4909ae60ffbb8e53)
	return nil
}

func __obf_f1c563c28223f805(__obf_dc35117108ba8439 *Decoder, __obf_63bbcee86d44fdde reflect.Value) error {
	__obf_4909ae60ffbb8e53, __obf_8882f6cf6e378ded := __obf_dc35117108ba8439.DecodeUint64()
	if __obf_8882f6cf6e378ded != nil {
		return __obf_8882f6cf6e378ded
	}
	__obf_63bbcee86d44fdde.
		SetUint(__obf_4909ae60ffbb8e53)
	return nil
}
