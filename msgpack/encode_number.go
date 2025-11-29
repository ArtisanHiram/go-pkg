package __obf_a4aad98aaf3178f4

import (
	"math"
	"reflect"
)

// EncodeUint8 encodes an uint8 in 2 bytes preserving type of the number.
func (__obf_2c8e97779108ab17 *Encoder) EncodeUint8(__obf_99a74e41c9c640c0 uint8) error {
	return __obf_2c8e97779108ab17.__obf_2ebb4be6da23dcc7(Uint8, __obf_99a74e41c9c640c0)
}

func (__obf_2c8e97779108ab17 *Encoder) __obf_79a7337e958b38bd(__obf_99a74e41c9c640c0 uint8) error {
	if __obf_2c8e97779108ab17.__obf_a8400e70a712e9fa&__obf_e0996ed6dc481d61 != 0 {
		return __obf_2c8e97779108ab17.EncodeUint(uint64(__obf_99a74e41c9c640c0))
	}
	return __obf_2c8e97779108ab17.EncodeUint8(__obf_99a74e41c9c640c0)
}

// EncodeUint16 encodes an uint16 in 3 bytes preserving type of the number.
func (__obf_2c8e97779108ab17 *Encoder) EncodeUint16(__obf_99a74e41c9c640c0 uint16) error {
	return __obf_2c8e97779108ab17.__obf_71321d3206264f3b(Uint16, __obf_99a74e41c9c640c0)
}

func (__obf_2c8e97779108ab17 *Encoder) __obf_c70e571779a7939c(__obf_99a74e41c9c640c0 uint16) error {
	if __obf_2c8e97779108ab17.__obf_a8400e70a712e9fa&__obf_e0996ed6dc481d61 != 0 {
		return __obf_2c8e97779108ab17.EncodeUint(uint64(__obf_99a74e41c9c640c0))
	}
	return __obf_2c8e97779108ab17.EncodeUint16(__obf_99a74e41c9c640c0)
}

// EncodeUint32 encodes an uint16 in 5 bytes preserving type of the number.
func (__obf_2c8e97779108ab17 *Encoder) EncodeUint32(__obf_99a74e41c9c640c0 uint32) error {
	return __obf_2c8e97779108ab17.__obf_a40df49dbc3bfc37(Uint32, __obf_99a74e41c9c640c0)
}

func (__obf_2c8e97779108ab17 *Encoder) __obf_24b3505b75405aef(__obf_99a74e41c9c640c0 uint32) error {
	if __obf_2c8e97779108ab17.__obf_a8400e70a712e9fa&__obf_e0996ed6dc481d61 != 0 {
		return __obf_2c8e97779108ab17.EncodeUint(uint64(__obf_99a74e41c9c640c0))
	}
	return __obf_2c8e97779108ab17.EncodeUint32(__obf_99a74e41c9c640c0)
}

// EncodeUint64 encodes an uint16 in 9 bytes preserving type of the number.
func (__obf_2c8e97779108ab17 *Encoder) EncodeUint64(__obf_99a74e41c9c640c0 uint64) error {
	return __obf_2c8e97779108ab17.__obf_99d307e2dd54d2df(Uint64, __obf_99a74e41c9c640c0)
}

func (__obf_2c8e97779108ab17 *Encoder) __obf_a5a421aeac90eb76(__obf_99a74e41c9c640c0 uint64) error {
	if __obf_2c8e97779108ab17.__obf_a8400e70a712e9fa&__obf_e0996ed6dc481d61 != 0 {
		return __obf_2c8e97779108ab17.EncodeUint(__obf_99a74e41c9c640c0)
	}
	return __obf_2c8e97779108ab17.EncodeUint64(__obf_99a74e41c9c640c0)
}

// EncodeInt8 encodes an int8 in 2 bytes preserving type of the number.
func (__obf_2c8e97779108ab17 *Encoder) EncodeInt8(__obf_99a74e41c9c640c0 int8) error {
	return __obf_2c8e97779108ab17.__obf_2ebb4be6da23dcc7(Int8, uint8(__obf_99a74e41c9c640c0))
}

func (__obf_2c8e97779108ab17 *Encoder) __obf_0307600014d27dd3(__obf_99a74e41c9c640c0 int8) error {
	if __obf_2c8e97779108ab17.__obf_a8400e70a712e9fa&__obf_e0996ed6dc481d61 != 0 {
		return __obf_2c8e97779108ab17.EncodeInt(int64(__obf_99a74e41c9c640c0))
	}
	return __obf_2c8e97779108ab17.EncodeInt8(__obf_99a74e41c9c640c0)
}

// EncodeInt16 encodes an int16 in 3 bytes preserving type of the number.
func (__obf_2c8e97779108ab17 *Encoder) EncodeInt16(__obf_99a74e41c9c640c0 int16) error {
	return __obf_2c8e97779108ab17.__obf_71321d3206264f3b(Int16, uint16(__obf_99a74e41c9c640c0))
}

func (__obf_2c8e97779108ab17 *Encoder) __obf_8856870da9fc8fb1(__obf_99a74e41c9c640c0 int16) error {
	if __obf_2c8e97779108ab17.__obf_a8400e70a712e9fa&__obf_e0996ed6dc481d61 != 0 {
		return __obf_2c8e97779108ab17.EncodeInt(int64(__obf_99a74e41c9c640c0))
	}
	return __obf_2c8e97779108ab17.EncodeInt16(__obf_99a74e41c9c640c0)
}

// EncodeInt32 encodes an int32 in 5 bytes preserving type of the number.
func (__obf_2c8e97779108ab17 *Encoder) EncodeInt32(__obf_99a74e41c9c640c0 int32) error {
	return __obf_2c8e97779108ab17.__obf_a40df49dbc3bfc37(Int32, uint32(__obf_99a74e41c9c640c0))
}

func (__obf_2c8e97779108ab17 *Encoder) __obf_6e7af44b46575bda(__obf_99a74e41c9c640c0 int32) error {
	if __obf_2c8e97779108ab17.__obf_a8400e70a712e9fa&__obf_e0996ed6dc481d61 != 0 {
		return __obf_2c8e97779108ab17.EncodeInt(int64(__obf_99a74e41c9c640c0))
	}
	return __obf_2c8e97779108ab17.EncodeInt32(__obf_99a74e41c9c640c0)
}

// EncodeInt64 encodes an int64 in 9 bytes preserving type of the number.
func (__obf_2c8e97779108ab17 *Encoder) EncodeInt64(__obf_99a74e41c9c640c0 int64) error {
	return __obf_2c8e97779108ab17.__obf_99d307e2dd54d2df(Int64, uint64(__obf_99a74e41c9c640c0))
}

func (__obf_2c8e97779108ab17 *Encoder) __obf_2837531178551b8d(__obf_99a74e41c9c640c0 int64) error {
	if __obf_2c8e97779108ab17.__obf_a8400e70a712e9fa&__obf_e0996ed6dc481d61 != 0 {
		return __obf_2c8e97779108ab17.EncodeInt(__obf_99a74e41c9c640c0)
	}
	return __obf_2c8e97779108ab17.EncodeInt64(__obf_99a74e41c9c640c0)
}

// EncodeUnsignedNumber encodes an uint64 in 1, 2, 3, 5, or 9 bytes.
// Type of the number is lost during encoding.
func (__obf_2c8e97779108ab17 *Encoder) EncodeUint(__obf_99a74e41c9c640c0 uint64) error {
	if __obf_99a74e41c9c640c0 <= math.MaxInt8 {
		return __obf_2c8e97779108ab17.__obf_bc988362b9d818fa.WriteByte(byte(__obf_99a74e41c9c640c0))
	}
	if __obf_99a74e41c9c640c0 <= math.MaxUint8 {
		return __obf_2c8e97779108ab17.EncodeUint8(uint8(__obf_99a74e41c9c640c0))
	}
	if __obf_99a74e41c9c640c0 <= math.MaxUint16 {
		return __obf_2c8e97779108ab17.EncodeUint16(uint16(__obf_99a74e41c9c640c0))
	}
	if __obf_99a74e41c9c640c0 <= math.MaxUint32 {
		return __obf_2c8e97779108ab17.EncodeUint32(uint32(__obf_99a74e41c9c640c0))
	}
	return __obf_2c8e97779108ab17.EncodeUint64(__obf_99a74e41c9c640c0)
}

// EncodeNumber encodes an int64 in 1, 2, 3, 5, or 9 bytes.
// Type of the number is lost during encoding.
func (__obf_2c8e97779108ab17 *Encoder) EncodeInt(__obf_99a74e41c9c640c0 int64) error {
	if __obf_99a74e41c9c640c0 >= 0 {
		return __obf_2c8e97779108ab17.EncodeUint(uint64(__obf_99a74e41c9c640c0))
	}
	if __obf_99a74e41c9c640c0 >= int64(int8(NegFixedNumLow)) {
		return __obf_2c8e97779108ab17.__obf_bc988362b9d818fa.WriteByte(byte(__obf_99a74e41c9c640c0))
	}
	if __obf_99a74e41c9c640c0 >= math.MinInt8 {
		return __obf_2c8e97779108ab17.EncodeInt8(int8(__obf_99a74e41c9c640c0))
	}
	if __obf_99a74e41c9c640c0 >= math.MinInt16 {
		return __obf_2c8e97779108ab17.EncodeInt16(int16(__obf_99a74e41c9c640c0))
	}
	if __obf_99a74e41c9c640c0 >= math.MinInt32 {
		return __obf_2c8e97779108ab17.EncodeInt32(int32(__obf_99a74e41c9c640c0))
	}
	return __obf_2c8e97779108ab17.EncodeInt64(__obf_99a74e41c9c640c0)
}

func (__obf_2c8e97779108ab17 *Encoder) EncodeFloat32(__obf_99a74e41c9c640c0 float32) error {
	if __obf_2c8e97779108ab17.__obf_a8400e70a712e9fa&__obf_d003dec3186c9089 != 0 {
		if float32(int64(__obf_99a74e41c9c640c0)) == __obf_99a74e41c9c640c0 {
			return __obf_2c8e97779108ab17.EncodeInt(int64(__obf_99a74e41c9c640c0))
		}
	}
	return __obf_2c8e97779108ab17.__obf_a40df49dbc3bfc37(Float, math.Float32bits(__obf_99a74e41c9c640c0))
}

func (__obf_2c8e97779108ab17 *Encoder) EncodeFloat64(__obf_99a74e41c9c640c0 float64) error {
	if __obf_2c8e97779108ab17.__obf_a8400e70a712e9fa&__obf_d003dec3186c9089 != 0 {
		// Both NaN and Inf convert to int64(-0x8000000000000000)
		// If n is NaN then it never compares true with any other value
		// If n is Inf then it doesn't convert from int64 back to +/-Inf
		// In both cases the comparison works.
		if float64(int64(__obf_99a74e41c9c640c0)) == __obf_99a74e41c9c640c0 {
			return __obf_2c8e97779108ab17.EncodeInt(int64(__obf_99a74e41c9c640c0))
		}
	}
	return __obf_2c8e97779108ab17.__obf_99d307e2dd54d2df(Double, math.Float64bits(__obf_99a74e41c9c640c0))
}

func (__obf_2c8e97779108ab17 *Encoder) __obf_2ebb4be6da23dcc7(__obf_987b95dd4dcfc308 byte, __obf_99a74e41c9c640c0 uint8) error {
	__obf_2c8e97779108ab17.__obf_f5d391a424aa5a3a = __obf_2c8e97779108ab17.__obf_f5d391a424aa5a3a[:2]
	__obf_2c8e97779108ab17.__obf_f5d391a424aa5a3a[0] = __obf_987b95dd4dcfc308
	__obf_2c8e97779108ab17.__obf_f5d391a424aa5a3a[1] = __obf_99a74e41c9c640c0
	return __obf_2c8e97779108ab17.__obf_ab9f7a04111e0451(__obf_2c8e97779108ab17.__obf_f5d391a424aa5a3a)
}

func (__obf_2c8e97779108ab17 *Encoder) __obf_71321d3206264f3b(__obf_987b95dd4dcfc308 byte, __obf_99a74e41c9c640c0 uint16) error {
	__obf_2c8e97779108ab17.__obf_f5d391a424aa5a3a = __obf_2c8e97779108ab17.__obf_f5d391a424aa5a3a[:3]
	__obf_2c8e97779108ab17.__obf_f5d391a424aa5a3a[0] = __obf_987b95dd4dcfc308
	__obf_2c8e97779108ab17.__obf_f5d391a424aa5a3a[1] = byte(__obf_99a74e41c9c640c0 >> 8)
	__obf_2c8e97779108ab17.__obf_f5d391a424aa5a3a[2] = byte(__obf_99a74e41c9c640c0)
	return __obf_2c8e97779108ab17.__obf_ab9f7a04111e0451(__obf_2c8e97779108ab17.__obf_f5d391a424aa5a3a)
}

func (__obf_2c8e97779108ab17 *Encoder) __obf_a40df49dbc3bfc37(__obf_987b95dd4dcfc308 byte, __obf_99a74e41c9c640c0 uint32) error {
	__obf_2c8e97779108ab17.__obf_f5d391a424aa5a3a = __obf_2c8e97779108ab17.__obf_f5d391a424aa5a3a[:5]
	__obf_2c8e97779108ab17.__obf_f5d391a424aa5a3a[0] = __obf_987b95dd4dcfc308
	__obf_2c8e97779108ab17.__obf_f5d391a424aa5a3a[1] = byte(__obf_99a74e41c9c640c0 >> 24)
	__obf_2c8e97779108ab17.__obf_f5d391a424aa5a3a[2] = byte(__obf_99a74e41c9c640c0 >> 16)
	__obf_2c8e97779108ab17.__obf_f5d391a424aa5a3a[3] = byte(__obf_99a74e41c9c640c0 >> 8)
	__obf_2c8e97779108ab17.__obf_f5d391a424aa5a3a[4] = byte(__obf_99a74e41c9c640c0)
	return __obf_2c8e97779108ab17.__obf_ab9f7a04111e0451(__obf_2c8e97779108ab17.__obf_f5d391a424aa5a3a)
}

func (__obf_2c8e97779108ab17 *Encoder) __obf_99d307e2dd54d2df(__obf_987b95dd4dcfc308 byte, __obf_99a74e41c9c640c0 uint64) error {
	__obf_2c8e97779108ab17.__obf_f5d391a424aa5a3a = __obf_2c8e97779108ab17.__obf_f5d391a424aa5a3a[:9]
	__obf_2c8e97779108ab17.__obf_f5d391a424aa5a3a[0] = __obf_987b95dd4dcfc308
	__obf_2c8e97779108ab17.__obf_f5d391a424aa5a3a[1] = byte(__obf_99a74e41c9c640c0 >> 56)
	__obf_2c8e97779108ab17.__obf_f5d391a424aa5a3a[2] = byte(__obf_99a74e41c9c640c0 >> 48)
	__obf_2c8e97779108ab17.__obf_f5d391a424aa5a3a[3] = byte(__obf_99a74e41c9c640c0 >> 40)
	__obf_2c8e97779108ab17.__obf_f5d391a424aa5a3a[4] = byte(__obf_99a74e41c9c640c0 >> 32)
	__obf_2c8e97779108ab17.__obf_f5d391a424aa5a3a[5] = byte(__obf_99a74e41c9c640c0 >> 24)
	__obf_2c8e97779108ab17.__obf_f5d391a424aa5a3a[6] = byte(__obf_99a74e41c9c640c0 >> 16)
	__obf_2c8e97779108ab17.__obf_f5d391a424aa5a3a[7] = byte(__obf_99a74e41c9c640c0 >> 8)
	__obf_2c8e97779108ab17.__obf_f5d391a424aa5a3a[8] = byte(__obf_99a74e41c9c640c0)
	return __obf_2c8e97779108ab17.__obf_ab9f7a04111e0451(__obf_2c8e97779108ab17.__obf_f5d391a424aa5a3a)
}

func __obf_7356d4170a0d75c7(__obf_2c8e97779108ab17 *Encoder, __obf_a1f43267eeb48a1a reflect.Value) error {
	return __obf_2c8e97779108ab17.EncodeUint(__obf_a1f43267eeb48a1a.Uint())
}

func __obf_82e4dec38ff6b0f0(__obf_2c8e97779108ab17 *Encoder, __obf_a1f43267eeb48a1a reflect.Value) error {
	return __obf_2c8e97779108ab17.EncodeInt(__obf_a1f43267eeb48a1a.Int())
}

func __obf_5234f218217cc405(__obf_2c8e97779108ab17 *Encoder, __obf_a1f43267eeb48a1a reflect.Value) error {
	return __obf_2c8e97779108ab17.__obf_79a7337e958b38bd(uint8(__obf_a1f43267eeb48a1a.Uint()))
}

func __obf_707d2811334e2066(__obf_2c8e97779108ab17 *Encoder, __obf_a1f43267eeb48a1a reflect.Value) error {
	return __obf_2c8e97779108ab17.__obf_c70e571779a7939c(uint16(__obf_a1f43267eeb48a1a.Uint()))
}

func __obf_6a77622f3635270d(__obf_2c8e97779108ab17 *Encoder, __obf_a1f43267eeb48a1a reflect.Value) error {
	return __obf_2c8e97779108ab17.__obf_24b3505b75405aef(uint32(__obf_a1f43267eeb48a1a.Uint()))
}

func __obf_2e07f9acaad90e64(__obf_2c8e97779108ab17 *Encoder, __obf_a1f43267eeb48a1a reflect.Value) error {
	return __obf_2c8e97779108ab17.__obf_a5a421aeac90eb76(__obf_a1f43267eeb48a1a.Uint())
}

func __obf_1f3e7325737988bb(__obf_2c8e97779108ab17 *Encoder, __obf_a1f43267eeb48a1a reflect.Value) error {
	return __obf_2c8e97779108ab17.__obf_0307600014d27dd3(int8(__obf_a1f43267eeb48a1a.Int()))
}

func __obf_8438172b55bae72b(__obf_2c8e97779108ab17 *Encoder, __obf_a1f43267eeb48a1a reflect.Value) error {
	return __obf_2c8e97779108ab17.__obf_8856870da9fc8fb1(int16(__obf_a1f43267eeb48a1a.Int()))
}

func __obf_3f87414aacb11a03(__obf_2c8e97779108ab17 *Encoder, __obf_a1f43267eeb48a1a reflect.Value) error {
	return __obf_2c8e97779108ab17.__obf_6e7af44b46575bda(int32(__obf_a1f43267eeb48a1a.Int()))
}

func __obf_b43612dd4961a390(__obf_2c8e97779108ab17 *Encoder, __obf_a1f43267eeb48a1a reflect.Value) error {
	return __obf_2c8e97779108ab17.__obf_2837531178551b8d(__obf_a1f43267eeb48a1a.Int())
}

func __obf_947000f09dd30923(__obf_2c8e97779108ab17 *Encoder, __obf_a1f43267eeb48a1a reflect.Value) error {
	return __obf_2c8e97779108ab17.EncodeFloat32(float32(__obf_a1f43267eeb48a1a.Float()))
}

func __obf_2053638319e34fb7(__obf_2c8e97779108ab17 *Encoder, __obf_a1f43267eeb48a1a reflect.Value) error {
	return __obf_2c8e97779108ab17.EncodeFloat64(__obf_a1f43267eeb48a1a.Float())
}
