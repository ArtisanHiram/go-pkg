package __obf_de86cdc8ae98b45b

import (
	"math"
	"reflect"
)

// EncodeUint8 encodes an uint8 in 2 bytes preserving type of the number.
func (__obf_7bae0b613da60dd3 *Encoder) EncodeUint8(__obf_2b0247e819b1bf4a uint8) error {
	return __obf_7bae0b613da60dd3.__obf_fa2b4086fff3f28d(Uint8, __obf_2b0247e819b1bf4a)
}

func (__obf_7bae0b613da60dd3 *Encoder) __obf_d4b5e4b5595c7de8(__obf_2b0247e819b1bf4a uint8) error {
	if __obf_7bae0b613da60dd3.__obf_a15a642720e39c3e&__obf_968e60dffbce5433 != 0 {
		return __obf_7bae0b613da60dd3.EncodeUint(uint64(__obf_2b0247e819b1bf4a))
	}
	return __obf_7bae0b613da60dd3.EncodeUint8(__obf_2b0247e819b1bf4a)
}

// EncodeUint16 encodes an uint16 in 3 bytes preserving type of the number.
func (__obf_7bae0b613da60dd3 *Encoder) EncodeUint16(__obf_2b0247e819b1bf4a uint16) error {
	return __obf_7bae0b613da60dd3.__obf_bd6e685495964ed5(Uint16, __obf_2b0247e819b1bf4a)
}

func (__obf_7bae0b613da60dd3 *Encoder) __obf_db82d42b0f9f6f8f(__obf_2b0247e819b1bf4a uint16) error {
	if __obf_7bae0b613da60dd3.__obf_a15a642720e39c3e&__obf_968e60dffbce5433 != 0 {
		return __obf_7bae0b613da60dd3.EncodeUint(uint64(__obf_2b0247e819b1bf4a))
	}
	return __obf_7bae0b613da60dd3.EncodeUint16(__obf_2b0247e819b1bf4a)
}

// EncodeUint32 encodes an uint16 in 5 bytes preserving type of the number.
func (__obf_7bae0b613da60dd3 *Encoder) EncodeUint32(__obf_2b0247e819b1bf4a uint32) error {
	return __obf_7bae0b613da60dd3.__obf_a95f338c34ae207a(Uint32, __obf_2b0247e819b1bf4a)
}

func (__obf_7bae0b613da60dd3 *Encoder) __obf_4e4467cbc37ccdcb(__obf_2b0247e819b1bf4a uint32) error {
	if __obf_7bae0b613da60dd3.__obf_a15a642720e39c3e&__obf_968e60dffbce5433 != 0 {
		return __obf_7bae0b613da60dd3.EncodeUint(uint64(__obf_2b0247e819b1bf4a))
	}
	return __obf_7bae0b613da60dd3.EncodeUint32(__obf_2b0247e819b1bf4a)
}

// EncodeUint64 encodes an uint16 in 9 bytes preserving type of the number.
func (__obf_7bae0b613da60dd3 *Encoder) EncodeUint64(__obf_2b0247e819b1bf4a uint64) error {
	return __obf_7bae0b613da60dd3.__obf_b239fe264182e9df(Uint64, __obf_2b0247e819b1bf4a)
}

func (__obf_7bae0b613da60dd3 *Encoder) __obf_70cf4747f5711e55(__obf_2b0247e819b1bf4a uint64) error {
	if __obf_7bae0b613da60dd3.__obf_a15a642720e39c3e&__obf_968e60dffbce5433 != 0 {
		return __obf_7bae0b613da60dd3.EncodeUint(__obf_2b0247e819b1bf4a)
	}
	return __obf_7bae0b613da60dd3.EncodeUint64(__obf_2b0247e819b1bf4a)
}

// EncodeInt8 encodes an int8 in 2 bytes preserving type of the number.
func (__obf_7bae0b613da60dd3 *Encoder) EncodeInt8(__obf_2b0247e819b1bf4a int8) error {
	return __obf_7bae0b613da60dd3.__obf_fa2b4086fff3f28d(Int8, uint8(__obf_2b0247e819b1bf4a))
}

func (__obf_7bae0b613da60dd3 *Encoder) __obf_bb6f7cdf9369179d(__obf_2b0247e819b1bf4a int8) error {
	if __obf_7bae0b613da60dd3.__obf_a15a642720e39c3e&__obf_968e60dffbce5433 != 0 {
		return __obf_7bae0b613da60dd3.EncodeInt(int64(__obf_2b0247e819b1bf4a))
	}
	return __obf_7bae0b613da60dd3.EncodeInt8(__obf_2b0247e819b1bf4a)
}

// EncodeInt16 encodes an int16 in 3 bytes preserving type of the number.
func (__obf_7bae0b613da60dd3 *Encoder) EncodeInt16(__obf_2b0247e819b1bf4a int16) error {
	return __obf_7bae0b613da60dd3.__obf_bd6e685495964ed5(Int16, uint16(__obf_2b0247e819b1bf4a))
}

func (__obf_7bae0b613da60dd3 *Encoder) __obf_5b63e7209330087f(__obf_2b0247e819b1bf4a int16) error {
	if __obf_7bae0b613da60dd3.__obf_a15a642720e39c3e&__obf_968e60dffbce5433 != 0 {
		return __obf_7bae0b613da60dd3.EncodeInt(int64(__obf_2b0247e819b1bf4a))
	}
	return __obf_7bae0b613da60dd3.EncodeInt16(__obf_2b0247e819b1bf4a)
}

// EncodeInt32 encodes an int32 in 5 bytes preserving type of the number.
func (__obf_7bae0b613da60dd3 *Encoder) EncodeInt32(__obf_2b0247e819b1bf4a int32) error {
	return __obf_7bae0b613da60dd3.__obf_a95f338c34ae207a(Int32, uint32(__obf_2b0247e819b1bf4a))
}

func (__obf_7bae0b613da60dd3 *Encoder) __obf_fcd295809f29fb18(__obf_2b0247e819b1bf4a int32) error {
	if __obf_7bae0b613da60dd3.__obf_a15a642720e39c3e&__obf_968e60dffbce5433 != 0 {
		return __obf_7bae0b613da60dd3.EncodeInt(int64(__obf_2b0247e819b1bf4a))
	}
	return __obf_7bae0b613da60dd3.EncodeInt32(__obf_2b0247e819b1bf4a)
}

// EncodeInt64 encodes an int64 in 9 bytes preserving type of the number.
func (__obf_7bae0b613da60dd3 *Encoder) EncodeInt64(__obf_2b0247e819b1bf4a int64) error {
	return __obf_7bae0b613da60dd3.__obf_b239fe264182e9df(Int64, uint64(__obf_2b0247e819b1bf4a))
}

func (__obf_7bae0b613da60dd3 *Encoder) __obf_545a172b23c8efda(__obf_2b0247e819b1bf4a int64) error {
	if __obf_7bae0b613da60dd3.__obf_a15a642720e39c3e&__obf_968e60dffbce5433 != 0 {
		return __obf_7bae0b613da60dd3.EncodeInt(__obf_2b0247e819b1bf4a)
	}
	return __obf_7bae0b613da60dd3.EncodeInt64(__obf_2b0247e819b1bf4a)
}

// EncodeUnsignedNumber encodes an uint64 in 1, 2, 3, 5, or 9 bytes.
// Type of the number is lost during encoding.
func (__obf_7bae0b613da60dd3 *Encoder) EncodeUint(__obf_2b0247e819b1bf4a uint64) error {
	if __obf_2b0247e819b1bf4a <= math.MaxInt8 {
		return __obf_7bae0b613da60dd3.__obf_007a7bbeb11f001c.WriteByte(byte(__obf_2b0247e819b1bf4a))
	}
	if __obf_2b0247e819b1bf4a <= math.MaxUint8 {
		return __obf_7bae0b613da60dd3.EncodeUint8(uint8(__obf_2b0247e819b1bf4a))
	}
	if __obf_2b0247e819b1bf4a <= math.MaxUint16 {
		return __obf_7bae0b613da60dd3.EncodeUint16(uint16(__obf_2b0247e819b1bf4a))
	}
	if __obf_2b0247e819b1bf4a <= math.MaxUint32 {
		return __obf_7bae0b613da60dd3.EncodeUint32(uint32(__obf_2b0247e819b1bf4a))
	}
	return __obf_7bae0b613da60dd3.EncodeUint64(__obf_2b0247e819b1bf4a)
}

// EncodeNumber encodes an int64 in 1, 2, 3, 5, or 9 bytes.
// Type of the number is lost during encoding.
func (__obf_7bae0b613da60dd3 *Encoder) EncodeInt(__obf_2b0247e819b1bf4a int64) error {
	if __obf_2b0247e819b1bf4a >= 0 {
		return __obf_7bae0b613da60dd3.EncodeUint(uint64(__obf_2b0247e819b1bf4a))
	}
	if __obf_2b0247e819b1bf4a >= int64(int8(NegFixedNumLow)) {
		return __obf_7bae0b613da60dd3.__obf_007a7bbeb11f001c.WriteByte(byte(__obf_2b0247e819b1bf4a))
	}
	if __obf_2b0247e819b1bf4a >= math.MinInt8 {
		return __obf_7bae0b613da60dd3.EncodeInt8(int8(__obf_2b0247e819b1bf4a))
	}
	if __obf_2b0247e819b1bf4a >= math.MinInt16 {
		return __obf_7bae0b613da60dd3.EncodeInt16(int16(__obf_2b0247e819b1bf4a))
	}
	if __obf_2b0247e819b1bf4a >= math.MinInt32 {
		return __obf_7bae0b613da60dd3.EncodeInt32(int32(__obf_2b0247e819b1bf4a))
	}
	return __obf_7bae0b613da60dd3.EncodeInt64(__obf_2b0247e819b1bf4a)
}

func (__obf_7bae0b613da60dd3 *Encoder) EncodeFloat32(__obf_2b0247e819b1bf4a float32) error {
	if __obf_7bae0b613da60dd3.__obf_a15a642720e39c3e&__obf_7a9e24d108f17fe5 != 0 {
		if float32(int64(__obf_2b0247e819b1bf4a)) == __obf_2b0247e819b1bf4a {
			return __obf_7bae0b613da60dd3.EncodeInt(int64(__obf_2b0247e819b1bf4a))
		}
	}
	return __obf_7bae0b613da60dd3.__obf_a95f338c34ae207a(Float, math.Float32bits(__obf_2b0247e819b1bf4a))
}

func (__obf_7bae0b613da60dd3 *Encoder) EncodeFloat64(__obf_2b0247e819b1bf4a float64) error {
	if __obf_7bae0b613da60dd3.__obf_a15a642720e39c3e&__obf_7a9e24d108f17fe5 != 0 {
		// Both NaN and Inf convert to int64(-0x8000000000000000)
		// If n is NaN then it never compares true with any other value
		// If n is Inf then it doesn't convert from int64 back to +/-Inf
		// In both cases the comparison works.
		if float64(int64(__obf_2b0247e819b1bf4a)) == __obf_2b0247e819b1bf4a {
			return __obf_7bae0b613da60dd3.EncodeInt(int64(__obf_2b0247e819b1bf4a))
		}
	}
	return __obf_7bae0b613da60dd3.__obf_b239fe264182e9df(Double, math.Float64bits(__obf_2b0247e819b1bf4a))
}

func (__obf_7bae0b613da60dd3 *Encoder) __obf_fa2b4086fff3f28d(__obf_34e3ba264a6bb541 byte, __obf_2b0247e819b1bf4a uint8) error {
	__obf_7bae0b613da60dd3.__obf_9c8c6a5e4cc8822c = __obf_7bae0b613da60dd3.__obf_9c8c6a5e4cc8822c[:2]
	__obf_7bae0b613da60dd3.__obf_9c8c6a5e4cc8822c[0] = __obf_34e3ba264a6bb541
	__obf_7bae0b613da60dd3.__obf_9c8c6a5e4cc8822c[1] = __obf_2b0247e819b1bf4a
	return __obf_7bae0b613da60dd3.__obf_b0daaf664cd5cdfb(__obf_7bae0b613da60dd3.__obf_9c8c6a5e4cc8822c)
}

func (__obf_7bae0b613da60dd3 *Encoder) __obf_bd6e685495964ed5(__obf_34e3ba264a6bb541 byte, __obf_2b0247e819b1bf4a uint16) error {
	__obf_7bae0b613da60dd3.__obf_9c8c6a5e4cc8822c = __obf_7bae0b613da60dd3.__obf_9c8c6a5e4cc8822c[:3]
	__obf_7bae0b613da60dd3.__obf_9c8c6a5e4cc8822c[0] = __obf_34e3ba264a6bb541
	__obf_7bae0b613da60dd3.__obf_9c8c6a5e4cc8822c[1] = byte(__obf_2b0247e819b1bf4a >> 8)
	__obf_7bae0b613da60dd3.__obf_9c8c6a5e4cc8822c[2] = byte(__obf_2b0247e819b1bf4a)
	return __obf_7bae0b613da60dd3.__obf_b0daaf664cd5cdfb(__obf_7bae0b613da60dd3.__obf_9c8c6a5e4cc8822c)
}

func (__obf_7bae0b613da60dd3 *Encoder) __obf_a95f338c34ae207a(__obf_34e3ba264a6bb541 byte, __obf_2b0247e819b1bf4a uint32) error {
	__obf_7bae0b613da60dd3.__obf_9c8c6a5e4cc8822c = __obf_7bae0b613da60dd3.__obf_9c8c6a5e4cc8822c[:5]
	__obf_7bae0b613da60dd3.__obf_9c8c6a5e4cc8822c[0] = __obf_34e3ba264a6bb541
	__obf_7bae0b613da60dd3.__obf_9c8c6a5e4cc8822c[1] = byte(__obf_2b0247e819b1bf4a >> 24)
	__obf_7bae0b613da60dd3.__obf_9c8c6a5e4cc8822c[2] = byte(__obf_2b0247e819b1bf4a >> 16)
	__obf_7bae0b613da60dd3.__obf_9c8c6a5e4cc8822c[3] = byte(__obf_2b0247e819b1bf4a >> 8)
	__obf_7bae0b613da60dd3.__obf_9c8c6a5e4cc8822c[4] = byte(__obf_2b0247e819b1bf4a)
	return __obf_7bae0b613da60dd3.__obf_b0daaf664cd5cdfb(__obf_7bae0b613da60dd3.__obf_9c8c6a5e4cc8822c)
}

func (__obf_7bae0b613da60dd3 *Encoder) __obf_b239fe264182e9df(__obf_34e3ba264a6bb541 byte, __obf_2b0247e819b1bf4a uint64) error {
	__obf_7bae0b613da60dd3.__obf_9c8c6a5e4cc8822c = __obf_7bae0b613da60dd3.__obf_9c8c6a5e4cc8822c[:9]
	__obf_7bae0b613da60dd3.__obf_9c8c6a5e4cc8822c[0] = __obf_34e3ba264a6bb541
	__obf_7bae0b613da60dd3.__obf_9c8c6a5e4cc8822c[1] = byte(__obf_2b0247e819b1bf4a >> 56)
	__obf_7bae0b613da60dd3.__obf_9c8c6a5e4cc8822c[2] = byte(__obf_2b0247e819b1bf4a >> 48)
	__obf_7bae0b613da60dd3.__obf_9c8c6a5e4cc8822c[3] = byte(__obf_2b0247e819b1bf4a >> 40)
	__obf_7bae0b613da60dd3.__obf_9c8c6a5e4cc8822c[4] = byte(__obf_2b0247e819b1bf4a >> 32)
	__obf_7bae0b613da60dd3.__obf_9c8c6a5e4cc8822c[5] = byte(__obf_2b0247e819b1bf4a >> 24)
	__obf_7bae0b613da60dd3.__obf_9c8c6a5e4cc8822c[6] = byte(__obf_2b0247e819b1bf4a >> 16)
	__obf_7bae0b613da60dd3.__obf_9c8c6a5e4cc8822c[7] = byte(__obf_2b0247e819b1bf4a >> 8)
	__obf_7bae0b613da60dd3.__obf_9c8c6a5e4cc8822c[8] = byte(__obf_2b0247e819b1bf4a)
	return __obf_7bae0b613da60dd3.__obf_b0daaf664cd5cdfb(__obf_7bae0b613da60dd3.__obf_9c8c6a5e4cc8822c)
}

func __obf_35394d5f1dd6a3c5(__obf_7bae0b613da60dd3 *Encoder, __obf_17732590722140e7 reflect.Value) error {
	return __obf_7bae0b613da60dd3.EncodeUint(__obf_17732590722140e7.Uint())
}

func __obf_0f70427f0c5e77bf(__obf_7bae0b613da60dd3 *Encoder, __obf_17732590722140e7 reflect.Value) error {
	return __obf_7bae0b613da60dd3.EncodeInt(__obf_17732590722140e7.Int())
}

func __obf_d187dcb37efcfe6d(__obf_7bae0b613da60dd3 *Encoder, __obf_17732590722140e7 reflect.Value) error {
	return __obf_7bae0b613da60dd3.__obf_d4b5e4b5595c7de8(uint8(__obf_17732590722140e7.Uint()))
}

func __obf_6148c4456411af75(__obf_7bae0b613da60dd3 *Encoder, __obf_17732590722140e7 reflect.Value) error {
	return __obf_7bae0b613da60dd3.__obf_db82d42b0f9f6f8f(uint16(__obf_17732590722140e7.Uint()))
}

func __obf_baeb98a0e3d5e172(__obf_7bae0b613da60dd3 *Encoder, __obf_17732590722140e7 reflect.Value) error {
	return __obf_7bae0b613da60dd3.__obf_4e4467cbc37ccdcb(uint32(__obf_17732590722140e7.Uint()))
}

func __obf_4722bc9436e891f9(__obf_7bae0b613da60dd3 *Encoder, __obf_17732590722140e7 reflect.Value) error {
	return __obf_7bae0b613da60dd3.__obf_70cf4747f5711e55(__obf_17732590722140e7.Uint())
}

func __obf_8858b34e9db12038(__obf_7bae0b613da60dd3 *Encoder, __obf_17732590722140e7 reflect.Value) error {
	return __obf_7bae0b613da60dd3.__obf_bb6f7cdf9369179d(int8(__obf_17732590722140e7.Int()))
}

func __obf_43983a1b215551a5(__obf_7bae0b613da60dd3 *Encoder, __obf_17732590722140e7 reflect.Value) error {
	return __obf_7bae0b613da60dd3.__obf_5b63e7209330087f(int16(__obf_17732590722140e7.Int()))
}

func __obf_1878d2a59a353162(__obf_7bae0b613da60dd3 *Encoder, __obf_17732590722140e7 reflect.Value) error {
	return __obf_7bae0b613da60dd3.__obf_fcd295809f29fb18(int32(__obf_17732590722140e7.Int()))
}

func __obf_b998d2da39e2780c(__obf_7bae0b613da60dd3 *Encoder, __obf_17732590722140e7 reflect.Value) error {
	return __obf_7bae0b613da60dd3.__obf_545a172b23c8efda(__obf_17732590722140e7.Int())
}

func __obf_ccb4e0a7821f8534(__obf_7bae0b613da60dd3 *Encoder, __obf_17732590722140e7 reflect.Value) error {
	return __obf_7bae0b613da60dd3.EncodeFloat32(float32(__obf_17732590722140e7.Float()))
}

func __obf_30345a7c18493f02(__obf_7bae0b613da60dd3 *Encoder, __obf_17732590722140e7 reflect.Value) error {
	return __obf_7bae0b613da60dd3.EncodeFloat64(__obf_17732590722140e7.Float())
}
