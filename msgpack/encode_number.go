package __obf_3edaa49e853afa16

import (
	"math"
	"reflect"
)

// EncodeUint8 encodes an uint8 in 2 bytes preserving type of the number.
func (__obf_84d0d31a8288f191 *Encoder) EncodeUint8(__obf_56127cd370854f0b uint8) error {
	return __obf_84d0d31a8288f191.__obf_59b14e1d8bb8dea4(Uint8, __obf_56127cd370854f0b)
}

func (__obf_84d0d31a8288f191 *Encoder) __obf_729a99a7d23def74(__obf_56127cd370854f0b uint8) error {
	if __obf_84d0d31a8288f191.__obf_6aa0efd537415192&__obf_5f5cd069c38159b5 != 0 {
		return __obf_84d0d31a8288f191.EncodeUint(uint64(__obf_56127cd370854f0b))
	}
	return __obf_84d0d31a8288f191.EncodeUint8(__obf_56127cd370854f0b)
}

// EncodeUint16 encodes an uint16 in 3 bytes preserving type of the number.
func (__obf_84d0d31a8288f191 *Encoder) EncodeUint16(__obf_56127cd370854f0b uint16) error {
	return __obf_84d0d31a8288f191.__obf_941ea548b75e478e(Uint16, __obf_56127cd370854f0b)
}

func (__obf_84d0d31a8288f191 *Encoder) __obf_d133025a3c5f654d(__obf_56127cd370854f0b uint16) error {
	if __obf_84d0d31a8288f191.__obf_6aa0efd537415192&__obf_5f5cd069c38159b5 != 0 {
		return __obf_84d0d31a8288f191.EncodeUint(uint64(__obf_56127cd370854f0b))
	}
	return __obf_84d0d31a8288f191.EncodeUint16(__obf_56127cd370854f0b)
}

// EncodeUint32 encodes an uint16 in 5 bytes preserving type of the number.
func (__obf_84d0d31a8288f191 *Encoder) EncodeUint32(__obf_56127cd370854f0b uint32) error {
	return __obf_84d0d31a8288f191.__obf_7636556325987942(Uint32, __obf_56127cd370854f0b)
}

func (__obf_84d0d31a8288f191 *Encoder) __obf_0e704b9f14205f77(__obf_56127cd370854f0b uint32) error {
	if __obf_84d0d31a8288f191.__obf_6aa0efd537415192&__obf_5f5cd069c38159b5 != 0 {
		return __obf_84d0d31a8288f191.EncodeUint(uint64(__obf_56127cd370854f0b))
	}
	return __obf_84d0d31a8288f191.EncodeUint32(__obf_56127cd370854f0b)
}

// EncodeUint64 encodes an uint16 in 9 bytes preserving type of the number.
func (__obf_84d0d31a8288f191 *Encoder) EncodeUint64(__obf_56127cd370854f0b uint64) error {
	return __obf_84d0d31a8288f191.__obf_4c7d6f59732c8908(Uint64, __obf_56127cd370854f0b)
}

func (__obf_84d0d31a8288f191 *Encoder) __obf_dde4b3164695e6f8(__obf_56127cd370854f0b uint64) error {
	if __obf_84d0d31a8288f191.__obf_6aa0efd537415192&__obf_5f5cd069c38159b5 != 0 {
		return __obf_84d0d31a8288f191.EncodeUint(__obf_56127cd370854f0b)
	}
	return __obf_84d0d31a8288f191.EncodeUint64(__obf_56127cd370854f0b)
}

// EncodeInt8 encodes an int8 in 2 bytes preserving type of the number.
func (__obf_84d0d31a8288f191 *Encoder) EncodeInt8(__obf_56127cd370854f0b int8) error {
	return __obf_84d0d31a8288f191.__obf_59b14e1d8bb8dea4(Int8, uint8(__obf_56127cd370854f0b))
}

func (__obf_84d0d31a8288f191 *Encoder) __obf_52e1b5a31e6cdcb7(__obf_56127cd370854f0b int8) error {
	if __obf_84d0d31a8288f191.__obf_6aa0efd537415192&__obf_5f5cd069c38159b5 != 0 {
		return __obf_84d0d31a8288f191.EncodeInt(int64(__obf_56127cd370854f0b))
	}
	return __obf_84d0d31a8288f191.EncodeInt8(__obf_56127cd370854f0b)
}

// EncodeInt16 encodes an int16 in 3 bytes preserving type of the number.
func (__obf_84d0d31a8288f191 *Encoder) EncodeInt16(__obf_56127cd370854f0b int16) error {
	return __obf_84d0d31a8288f191.__obf_941ea548b75e478e(Int16, uint16(__obf_56127cd370854f0b))
}

func (__obf_84d0d31a8288f191 *Encoder) __obf_c87734d4d666474f(__obf_56127cd370854f0b int16) error {
	if __obf_84d0d31a8288f191.__obf_6aa0efd537415192&__obf_5f5cd069c38159b5 != 0 {
		return __obf_84d0d31a8288f191.EncodeInt(int64(__obf_56127cd370854f0b))
	}
	return __obf_84d0d31a8288f191.EncodeInt16(__obf_56127cd370854f0b)
}

// EncodeInt32 encodes an int32 in 5 bytes preserving type of the number.
func (__obf_84d0d31a8288f191 *Encoder) EncodeInt32(__obf_56127cd370854f0b int32) error {
	return __obf_84d0d31a8288f191.__obf_7636556325987942(Int32, uint32(__obf_56127cd370854f0b))
}

func (__obf_84d0d31a8288f191 *Encoder) __obf_6af77158001d8265(__obf_56127cd370854f0b int32) error {
	if __obf_84d0d31a8288f191.__obf_6aa0efd537415192&__obf_5f5cd069c38159b5 != 0 {
		return __obf_84d0d31a8288f191.EncodeInt(int64(__obf_56127cd370854f0b))
	}
	return __obf_84d0d31a8288f191.EncodeInt32(__obf_56127cd370854f0b)
}

// EncodeInt64 encodes an int64 in 9 bytes preserving type of the number.
func (__obf_84d0d31a8288f191 *Encoder) EncodeInt64(__obf_56127cd370854f0b int64) error {
	return __obf_84d0d31a8288f191.__obf_4c7d6f59732c8908(Int64, uint64(__obf_56127cd370854f0b))
}

func (__obf_84d0d31a8288f191 *Encoder) __obf_b746e88d88f1bc44(__obf_56127cd370854f0b int64) error {
	if __obf_84d0d31a8288f191.__obf_6aa0efd537415192&__obf_5f5cd069c38159b5 != 0 {
		return __obf_84d0d31a8288f191.EncodeInt(__obf_56127cd370854f0b)
	}
	return __obf_84d0d31a8288f191.EncodeInt64(__obf_56127cd370854f0b)
}

// EncodeUnsignedNumber encodes an uint64 in 1, 2, 3, 5, or 9 bytes.
// Type of the number is lost during encoding.
func (__obf_84d0d31a8288f191 *Encoder) EncodeUint(__obf_56127cd370854f0b uint64) error {
	if __obf_56127cd370854f0b <= math.MaxInt8 {
		return __obf_84d0d31a8288f191.__obf_0f9492f42c73d099.WriteByte(byte(__obf_56127cd370854f0b))
	}
	if __obf_56127cd370854f0b <= math.MaxUint8 {
		return __obf_84d0d31a8288f191.EncodeUint8(uint8(__obf_56127cd370854f0b))
	}
	if __obf_56127cd370854f0b <= math.MaxUint16 {
		return __obf_84d0d31a8288f191.EncodeUint16(uint16(__obf_56127cd370854f0b))
	}
	if __obf_56127cd370854f0b <= math.MaxUint32 {
		return __obf_84d0d31a8288f191.EncodeUint32(uint32(__obf_56127cd370854f0b))
	}
	return __obf_84d0d31a8288f191.EncodeUint64(__obf_56127cd370854f0b)
}

// EncodeNumber encodes an int64 in 1, 2, 3, 5, or 9 bytes.
// Type of the number is lost during encoding.
func (__obf_84d0d31a8288f191 *Encoder) EncodeInt(__obf_56127cd370854f0b int64) error {
	if __obf_56127cd370854f0b >= 0 {
		return __obf_84d0d31a8288f191.EncodeUint(uint64(__obf_56127cd370854f0b))
	}
	if __obf_56127cd370854f0b >= int64(int8(NegFixedNumLow)) {
		return __obf_84d0d31a8288f191.__obf_0f9492f42c73d099.WriteByte(byte(__obf_56127cd370854f0b))
	}
	if __obf_56127cd370854f0b >= math.MinInt8 {
		return __obf_84d0d31a8288f191.EncodeInt8(int8(__obf_56127cd370854f0b))
	}
	if __obf_56127cd370854f0b >= math.MinInt16 {
		return __obf_84d0d31a8288f191.EncodeInt16(int16(__obf_56127cd370854f0b))
	}
	if __obf_56127cd370854f0b >= math.MinInt32 {
		return __obf_84d0d31a8288f191.EncodeInt32(int32(__obf_56127cd370854f0b))
	}
	return __obf_84d0d31a8288f191.EncodeInt64(__obf_56127cd370854f0b)
}

func (__obf_84d0d31a8288f191 *Encoder) EncodeFloat32(__obf_56127cd370854f0b float32) error {
	if __obf_84d0d31a8288f191.__obf_6aa0efd537415192&__obf_f553c7da4d60e61e != 0 {
		if float32(int64(__obf_56127cd370854f0b)) == __obf_56127cd370854f0b {
			return __obf_84d0d31a8288f191.EncodeInt(int64(__obf_56127cd370854f0b))
		}
	}
	return __obf_84d0d31a8288f191.__obf_7636556325987942(Float, math.Float32bits(__obf_56127cd370854f0b))
}

func (__obf_84d0d31a8288f191 *Encoder) EncodeFloat64(__obf_56127cd370854f0b float64) error {
	if __obf_84d0d31a8288f191.__obf_6aa0efd537415192&__obf_f553c7da4d60e61e != 0 {
		// Both NaN and Inf convert to int64(-0x8000000000000000)
		// If n is NaN then it never compares true with any other value
		// If n is Inf then it doesn't convert from int64 back to +/-Inf
		// In both cases the comparison works.
		if float64(int64(__obf_56127cd370854f0b)) == __obf_56127cd370854f0b {
			return __obf_84d0d31a8288f191.EncodeInt(int64(__obf_56127cd370854f0b))
		}
	}
	return __obf_84d0d31a8288f191.__obf_4c7d6f59732c8908(Double, math.Float64bits(__obf_56127cd370854f0b))
}

func (__obf_84d0d31a8288f191 *Encoder) __obf_59b14e1d8bb8dea4(__obf_508e2d6ff53655c0 byte, __obf_56127cd370854f0b uint8) error {
	__obf_84d0d31a8288f191.__obf_8f3bed0c23610f75 = __obf_84d0d31a8288f191.__obf_8f3bed0c23610f75[:2]
	__obf_84d0d31a8288f191.__obf_8f3bed0c23610f75[0] = __obf_508e2d6ff53655c0
	__obf_84d0d31a8288f191.__obf_8f3bed0c23610f75[1] = __obf_56127cd370854f0b
	return __obf_84d0d31a8288f191.__obf_fe140b946d6a369e(__obf_84d0d31a8288f191.__obf_8f3bed0c23610f75)
}

func (__obf_84d0d31a8288f191 *Encoder) __obf_941ea548b75e478e(__obf_508e2d6ff53655c0 byte, __obf_56127cd370854f0b uint16) error {
	__obf_84d0d31a8288f191.__obf_8f3bed0c23610f75 = __obf_84d0d31a8288f191.__obf_8f3bed0c23610f75[:3]
	__obf_84d0d31a8288f191.__obf_8f3bed0c23610f75[0] = __obf_508e2d6ff53655c0
	__obf_84d0d31a8288f191.__obf_8f3bed0c23610f75[1] = byte(__obf_56127cd370854f0b >> 8)
	__obf_84d0d31a8288f191.__obf_8f3bed0c23610f75[2] = byte(__obf_56127cd370854f0b)
	return __obf_84d0d31a8288f191.__obf_fe140b946d6a369e(__obf_84d0d31a8288f191.__obf_8f3bed0c23610f75)
}

func (__obf_84d0d31a8288f191 *Encoder) __obf_7636556325987942(__obf_508e2d6ff53655c0 byte, __obf_56127cd370854f0b uint32) error {
	__obf_84d0d31a8288f191.__obf_8f3bed0c23610f75 = __obf_84d0d31a8288f191.__obf_8f3bed0c23610f75[:5]
	__obf_84d0d31a8288f191.__obf_8f3bed0c23610f75[0] = __obf_508e2d6ff53655c0
	__obf_84d0d31a8288f191.__obf_8f3bed0c23610f75[1] = byte(__obf_56127cd370854f0b >> 24)
	__obf_84d0d31a8288f191.__obf_8f3bed0c23610f75[2] = byte(__obf_56127cd370854f0b >> 16)
	__obf_84d0d31a8288f191.__obf_8f3bed0c23610f75[3] = byte(__obf_56127cd370854f0b >> 8)
	__obf_84d0d31a8288f191.__obf_8f3bed0c23610f75[4] = byte(__obf_56127cd370854f0b)
	return __obf_84d0d31a8288f191.__obf_fe140b946d6a369e(__obf_84d0d31a8288f191.__obf_8f3bed0c23610f75)
}

func (__obf_84d0d31a8288f191 *Encoder) __obf_4c7d6f59732c8908(__obf_508e2d6ff53655c0 byte, __obf_56127cd370854f0b uint64) error {
	__obf_84d0d31a8288f191.__obf_8f3bed0c23610f75 = __obf_84d0d31a8288f191.__obf_8f3bed0c23610f75[:9]
	__obf_84d0d31a8288f191.__obf_8f3bed0c23610f75[0] = __obf_508e2d6ff53655c0
	__obf_84d0d31a8288f191.__obf_8f3bed0c23610f75[1] = byte(__obf_56127cd370854f0b >> 56)
	__obf_84d0d31a8288f191.__obf_8f3bed0c23610f75[2] = byte(__obf_56127cd370854f0b >> 48)
	__obf_84d0d31a8288f191.__obf_8f3bed0c23610f75[3] = byte(__obf_56127cd370854f0b >> 40)
	__obf_84d0d31a8288f191.__obf_8f3bed0c23610f75[4] = byte(__obf_56127cd370854f0b >> 32)
	__obf_84d0d31a8288f191.__obf_8f3bed0c23610f75[5] = byte(__obf_56127cd370854f0b >> 24)
	__obf_84d0d31a8288f191.__obf_8f3bed0c23610f75[6] = byte(__obf_56127cd370854f0b >> 16)
	__obf_84d0d31a8288f191.__obf_8f3bed0c23610f75[7] = byte(__obf_56127cd370854f0b >> 8)
	__obf_84d0d31a8288f191.__obf_8f3bed0c23610f75[8] = byte(__obf_56127cd370854f0b)
	return __obf_84d0d31a8288f191.__obf_fe140b946d6a369e(__obf_84d0d31a8288f191.__obf_8f3bed0c23610f75)
}

func __obf_ea3307b0a2ac01ef(__obf_84d0d31a8288f191 *Encoder, __obf_85d270343a0dfe14 reflect.Value) error {
	return __obf_84d0d31a8288f191.EncodeUint(__obf_85d270343a0dfe14.Uint())
}

func __obf_2143a6537eb57ebf(__obf_84d0d31a8288f191 *Encoder, __obf_85d270343a0dfe14 reflect.Value) error {
	return __obf_84d0d31a8288f191.EncodeInt(__obf_85d270343a0dfe14.Int())
}

func __obf_55469737ffbcc25a(__obf_84d0d31a8288f191 *Encoder, __obf_85d270343a0dfe14 reflect.Value) error {
	return __obf_84d0d31a8288f191.__obf_729a99a7d23def74(uint8(__obf_85d270343a0dfe14.Uint()))
}

func __obf_a78815c57f0236f5(__obf_84d0d31a8288f191 *Encoder, __obf_85d270343a0dfe14 reflect.Value) error {
	return __obf_84d0d31a8288f191.__obf_d133025a3c5f654d(uint16(__obf_85d270343a0dfe14.Uint()))
}

func __obf_287a9d470d9b65bf(__obf_84d0d31a8288f191 *Encoder, __obf_85d270343a0dfe14 reflect.Value) error {
	return __obf_84d0d31a8288f191.__obf_0e704b9f14205f77(uint32(__obf_85d270343a0dfe14.Uint()))
}

func __obf_b38acace80334530(__obf_84d0d31a8288f191 *Encoder, __obf_85d270343a0dfe14 reflect.Value) error {
	return __obf_84d0d31a8288f191.__obf_dde4b3164695e6f8(__obf_85d270343a0dfe14.Uint())
}

func __obf_408f73c1836414a3(__obf_84d0d31a8288f191 *Encoder, __obf_85d270343a0dfe14 reflect.Value) error {
	return __obf_84d0d31a8288f191.__obf_52e1b5a31e6cdcb7(int8(__obf_85d270343a0dfe14.Int()))
}

func __obf_ad4975ca3a4d2b79(__obf_84d0d31a8288f191 *Encoder, __obf_85d270343a0dfe14 reflect.Value) error {
	return __obf_84d0d31a8288f191.__obf_c87734d4d666474f(int16(__obf_85d270343a0dfe14.Int()))
}

func __obf_7ee966f13d02d236(__obf_84d0d31a8288f191 *Encoder, __obf_85d270343a0dfe14 reflect.Value) error {
	return __obf_84d0d31a8288f191.__obf_6af77158001d8265(int32(__obf_85d270343a0dfe14.Int()))
}

func __obf_33e67bf72b55646f(__obf_84d0d31a8288f191 *Encoder, __obf_85d270343a0dfe14 reflect.Value) error {
	return __obf_84d0d31a8288f191.__obf_b746e88d88f1bc44(__obf_85d270343a0dfe14.Int())
}

func __obf_c3e585f35dc28005(__obf_84d0d31a8288f191 *Encoder, __obf_85d270343a0dfe14 reflect.Value) error {
	return __obf_84d0d31a8288f191.EncodeFloat32(float32(__obf_85d270343a0dfe14.Float()))
}

func __obf_9c643c7031b55a89(__obf_84d0d31a8288f191 *Encoder, __obf_85d270343a0dfe14 reflect.Value) error {
	return __obf_84d0d31a8288f191.EncodeFloat64(__obf_85d270343a0dfe14.Float())
}
