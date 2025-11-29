package __obf_a4aad98aaf3178f4

import (
	"fmt"
	"math"
	"reflect"
)

func (__obf_613397fefdec0ed0 *Decoder) __obf_2bfec4fb91ed09f6(__obf_99a74e41c9c640c0 int) error {
	_, __obf_4061ca0039150c39 := __obf_613397fefdec0ed0.__obf_4429146a43365de4(__obf_99a74e41c9c640c0)
	return __obf_4061ca0039150c39
}

func (__obf_613397fefdec0ed0 *Decoder) uint8() (uint8, error) {
	__obf_6dbe86b3d9d9b037, __obf_4061ca0039150c39 := __obf_613397fefdec0ed0.__obf_46b95b57ed5617b5()
	if __obf_4061ca0039150c39 != nil {
		return 0, __obf_4061ca0039150c39
	}
	return __obf_6dbe86b3d9d9b037, nil
}

func (__obf_613397fefdec0ed0 *Decoder) int8() (int8, error) {
	__obf_99a74e41c9c640c0, __obf_4061ca0039150c39 := __obf_613397fefdec0ed0.uint8()
	return int8(__obf_99a74e41c9c640c0), __obf_4061ca0039150c39
}

func (__obf_613397fefdec0ed0 *Decoder) uint16() (uint16, error) {
	__obf_f57209cfda8e17cf, __obf_4061ca0039150c39 := __obf_613397fefdec0ed0.__obf_4429146a43365de4(2)
	if __obf_4061ca0039150c39 != nil {
		return 0, __obf_4061ca0039150c39
	}
	return (uint16(__obf_f57209cfda8e17cf[0]) << 8) | uint16(__obf_f57209cfda8e17cf[1]), nil
}

func (__obf_613397fefdec0ed0 *Decoder) int16() (int16, error) {
	__obf_99a74e41c9c640c0, __obf_4061ca0039150c39 := __obf_613397fefdec0ed0.uint16()
	return int16(__obf_99a74e41c9c640c0), __obf_4061ca0039150c39
}

func (__obf_613397fefdec0ed0 *Decoder) uint32() (uint32, error) {
	__obf_f57209cfda8e17cf, __obf_4061ca0039150c39 := __obf_613397fefdec0ed0.__obf_4429146a43365de4(4)
	if __obf_4061ca0039150c39 != nil {
		return 0, __obf_4061ca0039150c39
	}
	__obf_99a74e41c9c640c0 := (uint32(__obf_f57209cfda8e17cf[0]) << 24) |
		(uint32(__obf_f57209cfda8e17cf[1]) << 16) |
		(uint32(__obf_f57209cfda8e17cf[2]) << 8) |
		uint32(__obf_f57209cfda8e17cf[3])
	return __obf_99a74e41c9c640c0, nil
}

func (__obf_613397fefdec0ed0 *Decoder) int32() (int32, error) {
	__obf_99a74e41c9c640c0, __obf_4061ca0039150c39 := __obf_613397fefdec0ed0.uint32()
	return int32(__obf_99a74e41c9c640c0), __obf_4061ca0039150c39
}

func (__obf_613397fefdec0ed0 *Decoder) uint64() (uint64, error) {
	__obf_f57209cfda8e17cf, __obf_4061ca0039150c39 := __obf_613397fefdec0ed0.__obf_4429146a43365de4(8)
	if __obf_4061ca0039150c39 != nil {
		return 0, __obf_4061ca0039150c39
	}
	__obf_99a74e41c9c640c0 := (uint64(__obf_f57209cfda8e17cf[0]) << 56) |
		(uint64(__obf_f57209cfda8e17cf[1]) << 48) |
		(uint64(__obf_f57209cfda8e17cf[2]) << 40) |
		(uint64(__obf_f57209cfda8e17cf[3]) << 32) |
		(uint64(__obf_f57209cfda8e17cf[4]) << 24) |
		(uint64(__obf_f57209cfda8e17cf[5]) << 16) |
		(uint64(__obf_f57209cfda8e17cf[6]) << 8) |
		uint64(__obf_f57209cfda8e17cf[7])
	return __obf_99a74e41c9c640c0, nil
}

func (__obf_613397fefdec0ed0 *Decoder) int64() (int64, error) {
	__obf_99a74e41c9c640c0, __obf_4061ca0039150c39 := __obf_613397fefdec0ed0.uint64()
	return int64(__obf_99a74e41c9c640c0), __obf_4061ca0039150c39
}

// DecodeUint64 decodes msgpack int8/16/32/64 and uint8/16/32/64
// into Go uint64.
func (__obf_613397fefdec0ed0 *Decoder) DecodeUint64() (uint64, error) {
	__obf_6dbe86b3d9d9b037, __obf_4061ca0039150c39 := __obf_613397fefdec0ed0.__obf_46b95b57ed5617b5()
	if __obf_4061ca0039150c39 != nil {
		return 0, __obf_4061ca0039150c39
	}
	return __obf_613397fefdec0ed0.uint(__obf_6dbe86b3d9d9b037)
}

func (__obf_613397fefdec0ed0 *Decoder) uint(__obf_6dbe86b3d9d9b037 byte) (uint64, error) {
	if __obf_6dbe86b3d9d9b037 == Nil {
		return 0, nil
	}
	if IsFixedNum(__obf_6dbe86b3d9d9b037) {
		return uint64(int8(__obf_6dbe86b3d9d9b037)), nil
	}
	switch __obf_6dbe86b3d9d9b037 {
	case Uint8:
		__obf_99a74e41c9c640c0, __obf_4061ca0039150c39 := __obf_613397fefdec0ed0.uint8()
		return uint64(__obf_99a74e41c9c640c0), __obf_4061ca0039150c39
	case Int8:
		__obf_99a74e41c9c640c0, __obf_4061ca0039150c39 := __obf_613397fefdec0ed0.int8()
		return uint64(__obf_99a74e41c9c640c0), __obf_4061ca0039150c39
	case Uint16:
		__obf_99a74e41c9c640c0, __obf_4061ca0039150c39 := __obf_613397fefdec0ed0.uint16()
		return uint64(__obf_99a74e41c9c640c0), __obf_4061ca0039150c39
	case Int16:
		__obf_99a74e41c9c640c0, __obf_4061ca0039150c39 := __obf_613397fefdec0ed0.int16()
		return uint64(__obf_99a74e41c9c640c0), __obf_4061ca0039150c39
	case Uint32:
		__obf_99a74e41c9c640c0, __obf_4061ca0039150c39 := __obf_613397fefdec0ed0.uint32()
		return uint64(__obf_99a74e41c9c640c0), __obf_4061ca0039150c39
	case Int32:
		__obf_99a74e41c9c640c0, __obf_4061ca0039150c39 := __obf_613397fefdec0ed0.int32()
		return uint64(__obf_99a74e41c9c640c0), __obf_4061ca0039150c39
	case Uint64, Int64:
		return __obf_613397fefdec0ed0.uint64()
	}
	return 0, fmt.Errorf("msgpack: invalid code=%x decoding uint64", __obf_6dbe86b3d9d9b037)
}

// DecodeInt64 decodes msgpack int8/16/32/64 and uint8/16/32/64
// into Go int64.
func (__obf_613397fefdec0ed0 *Decoder) DecodeInt64() (int64, error) {
	__obf_6dbe86b3d9d9b037, __obf_4061ca0039150c39 := __obf_613397fefdec0ed0.__obf_46b95b57ed5617b5()
	if __obf_4061ca0039150c39 != nil {
		return 0, __obf_4061ca0039150c39
	}
	return __obf_613397fefdec0ed0.int(__obf_6dbe86b3d9d9b037)
}

func (__obf_613397fefdec0ed0 *Decoder) int(__obf_6dbe86b3d9d9b037 byte) (int64, error) {
	if __obf_6dbe86b3d9d9b037 == Nil {
		return 0, nil
	}
	if IsFixedNum(__obf_6dbe86b3d9d9b037) {
		return int64(int8(__obf_6dbe86b3d9d9b037)), nil
	}
	switch __obf_6dbe86b3d9d9b037 {
	case Uint8:
		__obf_99a74e41c9c640c0, __obf_4061ca0039150c39 := __obf_613397fefdec0ed0.uint8()
		return int64(__obf_99a74e41c9c640c0), __obf_4061ca0039150c39
	case Int8:
		__obf_99a74e41c9c640c0, __obf_4061ca0039150c39 := __obf_613397fefdec0ed0.uint8()
		return int64(int8(__obf_99a74e41c9c640c0)), __obf_4061ca0039150c39
	case Uint16:
		__obf_99a74e41c9c640c0, __obf_4061ca0039150c39 := __obf_613397fefdec0ed0.uint16()
		return int64(__obf_99a74e41c9c640c0), __obf_4061ca0039150c39
	case Int16:
		__obf_99a74e41c9c640c0, __obf_4061ca0039150c39 := __obf_613397fefdec0ed0.uint16()
		return int64(int16(__obf_99a74e41c9c640c0)), __obf_4061ca0039150c39
	case Uint32:
		__obf_99a74e41c9c640c0, __obf_4061ca0039150c39 := __obf_613397fefdec0ed0.uint32()
		return int64(__obf_99a74e41c9c640c0), __obf_4061ca0039150c39
	case Int32:
		__obf_99a74e41c9c640c0, __obf_4061ca0039150c39 := __obf_613397fefdec0ed0.uint32()
		return int64(int32(__obf_99a74e41c9c640c0)), __obf_4061ca0039150c39
	case Uint64, Int64:
		__obf_99a74e41c9c640c0, __obf_4061ca0039150c39 := __obf_613397fefdec0ed0.uint64()
		return int64(__obf_99a74e41c9c640c0), __obf_4061ca0039150c39
	}
	return 0, fmt.Errorf("msgpack: invalid code=%x decoding int64", __obf_6dbe86b3d9d9b037)
}

func (__obf_613397fefdec0ed0 *Decoder) DecodeFloat32() (float32, error) {
	__obf_6dbe86b3d9d9b037, __obf_4061ca0039150c39 := __obf_613397fefdec0ed0.__obf_46b95b57ed5617b5()
	if __obf_4061ca0039150c39 != nil {
		return 0, __obf_4061ca0039150c39
	}
	return __obf_613397fefdec0ed0.float32(__obf_6dbe86b3d9d9b037)
}

func (__obf_613397fefdec0ed0 *Decoder) float32(__obf_6dbe86b3d9d9b037 byte) (float32, error) {
	if __obf_6dbe86b3d9d9b037 == Float {
		__obf_99a74e41c9c640c0, __obf_4061ca0039150c39 := __obf_613397fefdec0ed0.uint32()
		if __obf_4061ca0039150c39 != nil {
			return 0, __obf_4061ca0039150c39
		}
		return math.Float32frombits(__obf_99a74e41c9c640c0), nil
	}
	__obf_99a74e41c9c640c0, __obf_4061ca0039150c39 := __obf_613397fefdec0ed0.int(__obf_6dbe86b3d9d9b037)
	if __obf_4061ca0039150c39 != nil {
		return 0, fmt.Errorf("msgpack: invalid code=%x decoding float32", __obf_6dbe86b3d9d9b037)
	}
	return float32(__obf_99a74e41c9c640c0), nil
}

// DecodeFloat64 decodes msgpack float32/64 into Go float64.
func (__obf_613397fefdec0ed0 *Decoder) DecodeFloat64() (float64, error) {
	__obf_6dbe86b3d9d9b037, __obf_4061ca0039150c39 := __obf_613397fefdec0ed0.__obf_46b95b57ed5617b5()
	if __obf_4061ca0039150c39 != nil {
		return 0, __obf_4061ca0039150c39
	}
	return __obf_613397fefdec0ed0.float64(__obf_6dbe86b3d9d9b037)
}

func (__obf_613397fefdec0ed0 *Decoder) float64(__obf_6dbe86b3d9d9b037 byte) (float64, error) {
	switch __obf_6dbe86b3d9d9b037 {
	case Float:
		__obf_99a74e41c9c640c0, __obf_4061ca0039150c39 := __obf_613397fefdec0ed0.float32(__obf_6dbe86b3d9d9b037)
		if __obf_4061ca0039150c39 != nil {
			return 0, __obf_4061ca0039150c39
		}
		return float64(__obf_99a74e41c9c640c0), nil
	case Double:
		__obf_99a74e41c9c640c0, __obf_4061ca0039150c39 := __obf_613397fefdec0ed0.uint64()
		if __obf_4061ca0039150c39 != nil {
			return 0, __obf_4061ca0039150c39
		}
		return math.Float64frombits(__obf_99a74e41c9c640c0), nil
	}
	__obf_99a74e41c9c640c0, __obf_4061ca0039150c39 := __obf_613397fefdec0ed0.int(__obf_6dbe86b3d9d9b037)
	if __obf_4061ca0039150c39 != nil {
		return 0, fmt.Errorf("msgpack: invalid code=%x decoding float32", __obf_6dbe86b3d9d9b037)
	}
	return float64(__obf_99a74e41c9c640c0), nil
}

func (__obf_613397fefdec0ed0 *Decoder) DecodeUint() (uint, error) {
	__obf_99a74e41c9c640c0, __obf_4061ca0039150c39 := __obf_613397fefdec0ed0.DecodeUint64()
	return uint(__obf_99a74e41c9c640c0), __obf_4061ca0039150c39
}

func (__obf_613397fefdec0ed0 *Decoder) DecodeUint8() (uint8, error) {
	__obf_99a74e41c9c640c0, __obf_4061ca0039150c39 := __obf_613397fefdec0ed0.DecodeUint64()
	return uint8(__obf_99a74e41c9c640c0), __obf_4061ca0039150c39
}

func (__obf_613397fefdec0ed0 *Decoder) DecodeUint16() (uint16, error) {
	__obf_99a74e41c9c640c0, __obf_4061ca0039150c39 := __obf_613397fefdec0ed0.DecodeUint64()
	return uint16(__obf_99a74e41c9c640c0), __obf_4061ca0039150c39
}

func (__obf_613397fefdec0ed0 *Decoder) DecodeUint32() (uint32, error) {
	__obf_99a74e41c9c640c0, __obf_4061ca0039150c39 := __obf_613397fefdec0ed0.DecodeUint64()
	return uint32(__obf_99a74e41c9c640c0), __obf_4061ca0039150c39
}

func (__obf_613397fefdec0ed0 *Decoder) DecodeInt() (int, error) {
	__obf_99a74e41c9c640c0, __obf_4061ca0039150c39 := __obf_613397fefdec0ed0.DecodeInt64()
	return int(__obf_99a74e41c9c640c0), __obf_4061ca0039150c39
}

func (__obf_613397fefdec0ed0 *Decoder) DecodeInt8() (int8, error) {
	__obf_99a74e41c9c640c0, __obf_4061ca0039150c39 := __obf_613397fefdec0ed0.DecodeInt64()
	return int8(__obf_99a74e41c9c640c0), __obf_4061ca0039150c39
}

func (__obf_613397fefdec0ed0 *Decoder) DecodeInt16() (int16, error) {
	__obf_99a74e41c9c640c0, __obf_4061ca0039150c39 := __obf_613397fefdec0ed0.DecodeInt64()
	return int16(__obf_99a74e41c9c640c0), __obf_4061ca0039150c39
}

func (__obf_613397fefdec0ed0 *Decoder) DecodeInt32() (int32, error) {
	__obf_99a74e41c9c640c0, __obf_4061ca0039150c39 := __obf_613397fefdec0ed0.DecodeInt64()
	return int32(__obf_99a74e41c9c640c0), __obf_4061ca0039150c39
}

func __obf_60142f5a2be96546(__obf_613397fefdec0ed0 *Decoder, __obf_a1f43267eeb48a1a reflect.Value) error {
	__obf_da4a2296298d318f, __obf_4061ca0039150c39 := __obf_613397fefdec0ed0.DecodeFloat32()
	if __obf_4061ca0039150c39 != nil {
		return __obf_4061ca0039150c39
	}
	__obf_a1f43267eeb48a1a.
		SetFloat(float64(__obf_da4a2296298d318f))
	return nil
}

func __obf_a14f56eca4c741e3(__obf_613397fefdec0ed0 *Decoder, __obf_a1f43267eeb48a1a reflect.Value) error {
	__obf_da4a2296298d318f, __obf_4061ca0039150c39 := __obf_613397fefdec0ed0.DecodeFloat64()
	if __obf_4061ca0039150c39 != nil {
		return __obf_4061ca0039150c39
	}
	__obf_a1f43267eeb48a1a.
		SetFloat(__obf_da4a2296298d318f)
	return nil
}

func __obf_ee0f5bc124a2f6b4(__obf_613397fefdec0ed0 *Decoder, __obf_a1f43267eeb48a1a reflect.Value) error {
	__obf_99a74e41c9c640c0, __obf_4061ca0039150c39 := __obf_613397fefdec0ed0.DecodeInt64()
	if __obf_4061ca0039150c39 != nil {
		return __obf_4061ca0039150c39
	}
	__obf_a1f43267eeb48a1a.
		SetInt(__obf_99a74e41c9c640c0)
	return nil
}

func __obf_4a99d79a1164a943(__obf_613397fefdec0ed0 *Decoder, __obf_a1f43267eeb48a1a reflect.Value) error {
	__obf_99a74e41c9c640c0, __obf_4061ca0039150c39 := __obf_613397fefdec0ed0.DecodeUint64()
	if __obf_4061ca0039150c39 != nil {
		return __obf_4061ca0039150c39
	}
	__obf_a1f43267eeb48a1a.
		SetUint(__obf_99a74e41c9c640c0)
	return nil
}
