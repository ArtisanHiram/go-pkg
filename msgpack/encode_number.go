package __obf_a935eb7f548271a4

import (
	"math"
	"reflect"
)

// EncodeUint8 encodes an uint8 in 2 bytes preserving type of the number.
func (__obf_ed37a34c347049f3 *Encoder) EncodeUint8(__obf_326af9bd942662ac uint8) error {
	return __obf_ed37a34c347049f3.__obf_327ee6f3f671bfd3(Uint8, __obf_326af9bd942662ac)
}

func (__obf_ed37a34c347049f3 *Encoder) __obf_cd27eebc28e38ff5(__obf_326af9bd942662ac uint8) error {
	if __obf_ed37a34c347049f3.__obf_d617f3769ce16c47&__obf_123cc756bbdcb1da != 0 {
		return __obf_ed37a34c347049f3.EncodeUint(uint64(__obf_326af9bd942662ac))
	}
	return __obf_ed37a34c347049f3.EncodeUint8(__obf_326af9bd942662ac)
}

// EncodeUint16 encodes an uint16 in 3 bytes preserving type of the number.
func (__obf_ed37a34c347049f3 *Encoder) EncodeUint16(__obf_326af9bd942662ac uint16) error {
	return __obf_ed37a34c347049f3.__obf_dd92b251146e86fb(Uint16, __obf_326af9bd942662ac)
}

func (__obf_ed37a34c347049f3 *Encoder) __obf_77300210d9b95814(__obf_326af9bd942662ac uint16) error {
	if __obf_ed37a34c347049f3.__obf_d617f3769ce16c47&__obf_123cc756bbdcb1da != 0 {
		return __obf_ed37a34c347049f3.EncodeUint(uint64(__obf_326af9bd942662ac))
	}
	return __obf_ed37a34c347049f3.EncodeUint16(__obf_326af9bd942662ac)
}

// EncodeUint32 encodes an uint16 in 5 bytes preserving type of the number.
func (__obf_ed37a34c347049f3 *Encoder) EncodeUint32(__obf_326af9bd942662ac uint32) error {
	return __obf_ed37a34c347049f3.__obf_21945a55177e9fbd(Uint32, __obf_326af9bd942662ac)
}

func (__obf_ed37a34c347049f3 *Encoder) __obf_2b933f259fdd020b(__obf_326af9bd942662ac uint32) error {
	if __obf_ed37a34c347049f3.__obf_d617f3769ce16c47&__obf_123cc756bbdcb1da != 0 {
		return __obf_ed37a34c347049f3.EncodeUint(uint64(__obf_326af9bd942662ac))
	}
	return __obf_ed37a34c347049f3.EncodeUint32(__obf_326af9bd942662ac)
}

// EncodeUint64 encodes an uint16 in 9 bytes preserving type of the number.
func (__obf_ed37a34c347049f3 *Encoder) EncodeUint64(__obf_326af9bd942662ac uint64) error {
	return __obf_ed37a34c347049f3.__obf_61fc980c1e1f8c83(Uint64, __obf_326af9bd942662ac)
}

func (__obf_ed37a34c347049f3 *Encoder) __obf_71513c11118b96bf(__obf_326af9bd942662ac uint64) error {
	if __obf_ed37a34c347049f3.__obf_d617f3769ce16c47&__obf_123cc756bbdcb1da != 0 {
		return __obf_ed37a34c347049f3.EncodeUint(__obf_326af9bd942662ac)
	}
	return __obf_ed37a34c347049f3.EncodeUint64(__obf_326af9bd942662ac)
}

// EncodeInt8 encodes an int8 in 2 bytes preserving type of the number.
func (__obf_ed37a34c347049f3 *Encoder) EncodeInt8(__obf_326af9bd942662ac int8) error {
	return __obf_ed37a34c347049f3.__obf_327ee6f3f671bfd3(Int8, uint8(__obf_326af9bd942662ac))
}

func (__obf_ed37a34c347049f3 *Encoder) __obf_c5c3e0b1a27006be(__obf_326af9bd942662ac int8) error {
	if __obf_ed37a34c347049f3.__obf_d617f3769ce16c47&__obf_123cc756bbdcb1da != 0 {
		return __obf_ed37a34c347049f3.EncodeInt(int64(__obf_326af9bd942662ac))
	}
	return __obf_ed37a34c347049f3.EncodeInt8(__obf_326af9bd942662ac)
}

// EncodeInt16 encodes an int16 in 3 bytes preserving type of the number.
func (__obf_ed37a34c347049f3 *Encoder) EncodeInt16(__obf_326af9bd942662ac int16) error {
	return __obf_ed37a34c347049f3.__obf_dd92b251146e86fb(Int16, uint16(__obf_326af9bd942662ac))
}

func (__obf_ed37a34c347049f3 *Encoder) __obf_92dc6721cc3bdce1(__obf_326af9bd942662ac int16) error {
	if __obf_ed37a34c347049f3.__obf_d617f3769ce16c47&__obf_123cc756bbdcb1da != 0 {
		return __obf_ed37a34c347049f3.EncodeInt(int64(__obf_326af9bd942662ac))
	}
	return __obf_ed37a34c347049f3.EncodeInt16(__obf_326af9bd942662ac)
}

// EncodeInt32 encodes an int32 in 5 bytes preserving type of the number.
func (__obf_ed37a34c347049f3 *Encoder) EncodeInt32(__obf_326af9bd942662ac int32) error {
	return __obf_ed37a34c347049f3.__obf_21945a55177e9fbd(Int32, uint32(__obf_326af9bd942662ac))
}

func (__obf_ed37a34c347049f3 *Encoder) __obf_53aeb763f6374830(__obf_326af9bd942662ac int32) error {
	if __obf_ed37a34c347049f3.__obf_d617f3769ce16c47&__obf_123cc756bbdcb1da != 0 {
		return __obf_ed37a34c347049f3.EncodeInt(int64(__obf_326af9bd942662ac))
	}
	return __obf_ed37a34c347049f3.EncodeInt32(__obf_326af9bd942662ac)
}

// EncodeInt64 encodes an int64 in 9 bytes preserving type of the number.
func (__obf_ed37a34c347049f3 *Encoder) EncodeInt64(__obf_326af9bd942662ac int64) error {
	return __obf_ed37a34c347049f3.__obf_61fc980c1e1f8c83(Int64, uint64(__obf_326af9bd942662ac))
}

func (__obf_ed37a34c347049f3 *Encoder) __obf_35627de976adab1d(__obf_326af9bd942662ac int64) error {
	if __obf_ed37a34c347049f3.__obf_d617f3769ce16c47&__obf_123cc756bbdcb1da != 0 {
		return __obf_ed37a34c347049f3.EncodeInt(__obf_326af9bd942662ac)
	}
	return __obf_ed37a34c347049f3.EncodeInt64(__obf_326af9bd942662ac)
}

// EncodeUnsignedNumber encodes an uint64 in 1, 2, 3, 5, or 9 bytes.
// Type of the number is lost during encoding.
func (__obf_ed37a34c347049f3 *Encoder) EncodeUint(__obf_326af9bd942662ac uint64) error {
	if __obf_326af9bd942662ac <= math.MaxInt8 {
		return __obf_ed37a34c347049f3.__obf_a3adbf62c8801563.WriteByte(byte(__obf_326af9bd942662ac))
	}
	if __obf_326af9bd942662ac <= math.MaxUint8 {
		return __obf_ed37a34c347049f3.EncodeUint8(uint8(__obf_326af9bd942662ac))
	}
	if __obf_326af9bd942662ac <= math.MaxUint16 {
		return __obf_ed37a34c347049f3.EncodeUint16(uint16(__obf_326af9bd942662ac))
	}
	if __obf_326af9bd942662ac <= math.MaxUint32 {
		return __obf_ed37a34c347049f3.EncodeUint32(uint32(__obf_326af9bd942662ac))
	}
	return __obf_ed37a34c347049f3.EncodeUint64(__obf_326af9bd942662ac)
}

// EncodeNumber encodes an int64 in 1, 2, 3, 5, or 9 bytes.
// Type of the number is lost during encoding.
func (__obf_ed37a34c347049f3 *Encoder) EncodeInt(__obf_326af9bd942662ac int64) error {
	if __obf_326af9bd942662ac >= 0 {
		return __obf_ed37a34c347049f3.EncodeUint(uint64(__obf_326af9bd942662ac))
	}
	if __obf_326af9bd942662ac >= int64(int8(NegFixedNumLow)) {
		return __obf_ed37a34c347049f3.__obf_a3adbf62c8801563.WriteByte(byte(__obf_326af9bd942662ac))
	}
	if __obf_326af9bd942662ac >= math.MinInt8 {
		return __obf_ed37a34c347049f3.EncodeInt8(int8(__obf_326af9bd942662ac))
	}
	if __obf_326af9bd942662ac >= math.MinInt16 {
		return __obf_ed37a34c347049f3.EncodeInt16(int16(__obf_326af9bd942662ac))
	}
	if __obf_326af9bd942662ac >= math.MinInt32 {
		return __obf_ed37a34c347049f3.EncodeInt32(int32(__obf_326af9bd942662ac))
	}
	return __obf_ed37a34c347049f3.EncodeInt64(__obf_326af9bd942662ac)
}

func (__obf_ed37a34c347049f3 *Encoder) EncodeFloat32(__obf_326af9bd942662ac float32) error {
	if __obf_ed37a34c347049f3.__obf_d617f3769ce16c47&__obf_9c276895701d3946 != 0 {
		if float32(int64(__obf_326af9bd942662ac)) == __obf_326af9bd942662ac {
			return __obf_ed37a34c347049f3.EncodeInt(int64(__obf_326af9bd942662ac))
		}
	}
	return __obf_ed37a34c347049f3.__obf_21945a55177e9fbd(Float, math.Float32bits(__obf_326af9bd942662ac))
}

func (__obf_ed37a34c347049f3 *Encoder) EncodeFloat64(__obf_326af9bd942662ac float64) error {
	if __obf_ed37a34c347049f3.__obf_d617f3769ce16c47&__obf_9c276895701d3946 != 0 {
		// Both NaN and Inf convert to int64(-0x8000000000000000)
		// If n is NaN then it never compares true with any other value
		// If n is Inf then it doesn't convert from int64 back to +/-Inf
		// In both cases the comparison works.
		if float64(int64(__obf_326af9bd942662ac)) == __obf_326af9bd942662ac {
			return __obf_ed37a34c347049f3.EncodeInt(int64(__obf_326af9bd942662ac))
		}
	}
	return __obf_ed37a34c347049f3.__obf_61fc980c1e1f8c83(Double, math.Float64bits(__obf_326af9bd942662ac))
}

func (__obf_ed37a34c347049f3 *Encoder) __obf_327ee6f3f671bfd3(__obf_b983039ef2a336eb byte, __obf_326af9bd942662ac uint8) error {
	__obf_ed37a34c347049f3.__obf_a2e16952f247f218 = __obf_ed37a34c347049f3.__obf_a2e16952f247f218[:2]
	__obf_ed37a34c347049f3.__obf_a2e16952f247f218[0] = __obf_b983039ef2a336eb
	__obf_ed37a34c347049f3.__obf_a2e16952f247f218[1] = __obf_326af9bd942662ac
	return __obf_ed37a34c347049f3.__obf_654725fc7d18ae48(__obf_ed37a34c347049f3.__obf_a2e16952f247f218)
}

func (__obf_ed37a34c347049f3 *Encoder) __obf_dd92b251146e86fb(__obf_b983039ef2a336eb byte, __obf_326af9bd942662ac uint16) error {
	__obf_ed37a34c347049f3.__obf_a2e16952f247f218 = __obf_ed37a34c347049f3.__obf_a2e16952f247f218[:3]
	__obf_ed37a34c347049f3.__obf_a2e16952f247f218[0] = __obf_b983039ef2a336eb
	__obf_ed37a34c347049f3.__obf_a2e16952f247f218[1] = byte(__obf_326af9bd942662ac >> 8)
	__obf_ed37a34c347049f3.__obf_a2e16952f247f218[2] = byte(__obf_326af9bd942662ac)
	return __obf_ed37a34c347049f3.__obf_654725fc7d18ae48(__obf_ed37a34c347049f3.__obf_a2e16952f247f218)
}

func (__obf_ed37a34c347049f3 *Encoder) __obf_21945a55177e9fbd(__obf_b983039ef2a336eb byte, __obf_326af9bd942662ac uint32) error {
	__obf_ed37a34c347049f3.__obf_a2e16952f247f218 = __obf_ed37a34c347049f3.__obf_a2e16952f247f218[:5]
	__obf_ed37a34c347049f3.__obf_a2e16952f247f218[0] = __obf_b983039ef2a336eb
	__obf_ed37a34c347049f3.__obf_a2e16952f247f218[1] = byte(__obf_326af9bd942662ac >> 24)
	__obf_ed37a34c347049f3.__obf_a2e16952f247f218[2] = byte(__obf_326af9bd942662ac >> 16)
	__obf_ed37a34c347049f3.__obf_a2e16952f247f218[3] = byte(__obf_326af9bd942662ac >> 8)
	__obf_ed37a34c347049f3.__obf_a2e16952f247f218[4] = byte(__obf_326af9bd942662ac)
	return __obf_ed37a34c347049f3.__obf_654725fc7d18ae48(__obf_ed37a34c347049f3.__obf_a2e16952f247f218)
}

func (__obf_ed37a34c347049f3 *Encoder) __obf_61fc980c1e1f8c83(__obf_b983039ef2a336eb byte, __obf_326af9bd942662ac uint64) error {
	__obf_ed37a34c347049f3.__obf_a2e16952f247f218 = __obf_ed37a34c347049f3.__obf_a2e16952f247f218[:9]
	__obf_ed37a34c347049f3.__obf_a2e16952f247f218[0] = __obf_b983039ef2a336eb
	__obf_ed37a34c347049f3.__obf_a2e16952f247f218[1] = byte(__obf_326af9bd942662ac >> 56)
	__obf_ed37a34c347049f3.__obf_a2e16952f247f218[2] = byte(__obf_326af9bd942662ac >> 48)
	__obf_ed37a34c347049f3.__obf_a2e16952f247f218[3] = byte(__obf_326af9bd942662ac >> 40)
	__obf_ed37a34c347049f3.__obf_a2e16952f247f218[4] = byte(__obf_326af9bd942662ac >> 32)
	__obf_ed37a34c347049f3.__obf_a2e16952f247f218[5] = byte(__obf_326af9bd942662ac >> 24)
	__obf_ed37a34c347049f3.__obf_a2e16952f247f218[6] = byte(__obf_326af9bd942662ac >> 16)
	__obf_ed37a34c347049f3.__obf_a2e16952f247f218[7] = byte(__obf_326af9bd942662ac >> 8)
	__obf_ed37a34c347049f3.__obf_a2e16952f247f218[8] = byte(__obf_326af9bd942662ac)
	return __obf_ed37a34c347049f3.__obf_654725fc7d18ae48(__obf_ed37a34c347049f3.__obf_a2e16952f247f218)
}

func __obf_26b7f97739101864(__obf_ed37a34c347049f3 *Encoder, __obf_6d570581f4b60dbc reflect.Value) error {
	return __obf_ed37a34c347049f3.EncodeUint(__obf_6d570581f4b60dbc.Uint())
}

func __obf_53328fa97f4370bb(__obf_ed37a34c347049f3 *Encoder, __obf_6d570581f4b60dbc reflect.Value) error {
	return __obf_ed37a34c347049f3.EncodeInt(__obf_6d570581f4b60dbc.Int())
}

func __obf_2694089cc89b620b(__obf_ed37a34c347049f3 *Encoder, __obf_6d570581f4b60dbc reflect.Value) error {
	return __obf_ed37a34c347049f3.__obf_cd27eebc28e38ff5(uint8(__obf_6d570581f4b60dbc.Uint()))
}

func __obf_64119742b2dcdc14(__obf_ed37a34c347049f3 *Encoder, __obf_6d570581f4b60dbc reflect.Value) error {
	return __obf_ed37a34c347049f3.__obf_77300210d9b95814(uint16(__obf_6d570581f4b60dbc.Uint()))
}

func __obf_f89278ad27f637c0(__obf_ed37a34c347049f3 *Encoder, __obf_6d570581f4b60dbc reflect.Value) error {
	return __obf_ed37a34c347049f3.__obf_2b933f259fdd020b(uint32(__obf_6d570581f4b60dbc.Uint()))
}

func __obf_2d4225df72c2ec23(__obf_ed37a34c347049f3 *Encoder, __obf_6d570581f4b60dbc reflect.Value) error {
	return __obf_ed37a34c347049f3.__obf_71513c11118b96bf(__obf_6d570581f4b60dbc.Uint())
}

func __obf_17e8f11c01632ce7(__obf_ed37a34c347049f3 *Encoder, __obf_6d570581f4b60dbc reflect.Value) error {
	return __obf_ed37a34c347049f3.__obf_c5c3e0b1a27006be(int8(__obf_6d570581f4b60dbc.Int()))
}

func __obf_aee8774824bbd873(__obf_ed37a34c347049f3 *Encoder, __obf_6d570581f4b60dbc reflect.Value) error {
	return __obf_ed37a34c347049f3.__obf_92dc6721cc3bdce1(int16(__obf_6d570581f4b60dbc.Int()))
}

func __obf_4c356e77af1b6a6b(__obf_ed37a34c347049f3 *Encoder, __obf_6d570581f4b60dbc reflect.Value) error {
	return __obf_ed37a34c347049f3.__obf_53aeb763f6374830(int32(__obf_6d570581f4b60dbc.Int()))
}

func __obf_df3286b177337d63(__obf_ed37a34c347049f3 *Encoder, __obf_6d570581f4b60dbc reflect.Value) error {
	return __obf_ed37a34c347049f3.__obf_35627de976adab1d(__obf_6d570581f4b60dbc.Int())
}

func __obf_f67e8b412fa8016c(__obf_ed37a34c347049f3 *Encoder, __obf_6d570581f4b60dbc reflect.Value) error {
	return __obf_ed37a34c347049f3.EncodeFloat32(float32(__obf_6d570581f4b60dbc.Float()))
}

func __obf_8f648744c64b0585(__obf_ed37a34c347049f3 *Encoder, __obf_6d570581f4b60dbc reflect.Value) error {
	return __obf_ed37a34c347049f3.EncodeFloat64(__obf_6d570581f4b60dbc.Float())
}
