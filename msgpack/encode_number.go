package __obf_3e0c303610a19bd4

import (
	"math"
	"reflect"
)

// EncodeUint8 encodes an uint8 in 2 bytes preserving type of the number.
func (__obf_77240dc7776b60b4 *Encoder) EncodeUint8(__obf_4909ae60ffbb8e53 uint8) error {
	return __obf_77240dc7776b60b4.__obf_a5ea8b1a318bf9eb(Uint8, __obf_4909ae60ffbb8e53)
}

func (__obf_77240dc7776b60b4 *Encoder) __obf_4c4fec685e237b13(__obf_4909ae60ffbb8e53 uint8) error {
	if __obf_77240dc7776b60b4.__obf_3cf0882fa5a4cafb&__obf_506e82775380d9f7 != 0 {
		return __obf_77240dc7776b60b4.EncodeUint(uint64(__obf_4909ae60ffbb8e53))
	}
	return __obf_77240dc7776b60b4.EncodeUint8(__obf_4909ae60ffbb8e53)
}

// EncodeUint16 encodes an uint16 in 3 bytes preserving type of the number.
func (__obf_77240dc7776b60b4 *Encoder) EncodeUint16(__obf_4909ae60ffbb8e53 uint16) error {
	return __obf_77240dc7776b60b4.__obf_9ecc220d5defc1d9(Uint16, __obf_4909ae60ffbb8e53)
}

func (__obf_77240dc7776b60b4 *Encoder) __obf_19dd028d0c8c035e(__obf_4909ae60ffbb8e53 uint16) error {
	if __obf_77240dc7776b60b4.__obf_3cf0882fa5a4cafb&__obf_506e82775380d9f7 != 0 {
		return __obf_77240dc7776b60b4.EncodeUint(uint64(__obf_4909ae60ffbb8e53))
	}
	return __obf_77240dc7776b60b4.EncodeUint16(__obf_4909ae60ffbb8e53)
}

// EncodeUint32 encodes an uint16 in 5 bytes preserving type of the number.
func (__obf_77240dc7776b60b4 *Encoder) EncodeUint32(__obf_4909ae60ffbb8e53 uint32) error {
	return __obf_77240dc7776b60b4.__obf_70cbe6e3e07b6347(Uint32, __obf_4909ae60ffbb8e53)
}

func (__obf_77240dc7776b60b4 *Encoder) __obf_d9ac25e44ae7c9f2(__obf_4909ae60ffbb8e53 uint32) error {
	if __obf_77240dc7776b60b4.__obf_3cf0882fa5a4cafb&__obf_506e82775380d9f7 != 0 {
		return __obf_77240dc7776b60b4.EncodeUint(uint64(__obf_4909ae60ffbb8e53))
	}
	return __obf_77240dc7776b60b4.EncodeUint32(__obf_4909ae60ffbb8e53)
}

// EncodeUint64 encodes an uint16 in 9 bytes preserving type of the number.
func (__obf_77240dc7776b60b4 *Encoder) EncodeUint64(__obf_4909ae60ffbb8e53 uint64) error {
	return __obf_77240dc7776b60b4.__obf_4003fdc721833a64(Uint64, __obf_4909ae60ffbb8e53)
}

func (__obf_77240dc7776b60b4 *Encoder) __obf_e32b04a6bb97fb7e(__obf_4909ae60ffbb8e53 uint64) error {
	if __obf_77240dc7776b60b4.__obf_3cf0882fa5a4cafb&__obf_506e82775380d9f7 != 0 {
		return __obf_77240dc7776b60b4.EncodeUint(__obf_4909ae60ffbb8e53)
	}
	return __obf_77240dc7776b60b4.EncodeUint64(__obf_4909ae60ffbb8e53)
}

// EncodeInt8 encodes an int8 in 2 bytes preserving type of the number.
func (__obf_77240dc7776b60b4 *Encoder) EncodeInt8(__obf_4909ae60ffbb8e53 int8) error {
	return __obf_77240dc7776b60b4.__obf_a5ea8b1a318bf9eb(Int8, uint8(__obf_4909ae60ffbb8e53))
}

func (__obf_77240dc7776b60b4 *Encoder) __obf_6ba66788d0a0ac7b(__obf_4909ae60ffbb8e53 int8) error {
	if __obf_77240dc7776b60b4.__obf_3cf0882fa5a4cafb&__obf_506e82775380d9f7 != 0 {
		return __obf_77240dc7776b60b4.EncodeInt(int64(__obf_4909ae60ffbb8e53))
	}
	return __obf_77240dc7776b60b4.EncodeInt8(__obf_4909ae60ffbb8e53)
}

// EncodeInt16 encodes an int16 in 3 bytes preserving type of the number.
func (__obf_77240dc7776b60b4 *Encoder) EncodeInt16(__obf_4909ae60ffbb8e53 int16) error {
	return __obf_77240dc7776b60b4.__obf_9ecc220d5defc1d9(Int16, uint16(__obf_4909ae60ffbb8e53))
}

func (__obf_77240dc7776b60b4 *Encoder) __obf_73b9261f3dd2dce8(__obf_4909ae60ffbb8e53 int16) error {
	if __obf_77240dc7776b60b4.__obf_3cf0882fa5a4cafb&__obf_506e82775380d9f7 != 0 {
		return __obf_77240dc7776b60b4.EncodeInt(int64(__obf_4909ae60ffbb8e53))
	}
	return __obf_77240dc7776b60b4.EncodeInt16(__obf_4909ae60ffbb8e53)
}

// EncodeInt32 encodes an int32 in 5 bytes preserving type of the number.
func (__obf_77240dc7776b60b4 *Encoder) EncodeInt32(__obf_4909ae60ffbb8e53 int32) error {
	return __obf_77240dc7776b60b4.__obf_70cbe6e3e07b6347(Int32, uint32(__obf_4909ae60ffbb8e53))
}

func (__obf_77240dc7776b60b4 *Encoder) __obf_18b40d4bd23a9cae(__obf_4909ae60ffbb8e53 int32) error {
	if __obf_77240dc7776b60b4.__obf_3cf0882fa5a4cafb&__obf_506e82775380d9f7 != 0 {
		return __obf_77240dc7776b60b4.EncodeInt(int64(__obf_4909ae60ffbb8e53))
	}
	return __obf_77240dc7776b60b4.EncodeInt32(__obf_4909ae60ffbb8e53)
}

// EncodeInt64 encodes an int64 in 9 bytes preserving type of the number.
func (__obf_77240dc7776b60b4 *Encoder) EncodeInt64(__obf_4909ae60ffbb8e53 int64) error {
	return __obf_77240dc7776b60b4.__obf_4003fdc721833a64(Int64, uint64(__obf_4909ae60ffbb8e53))
}

func (__obf_77240dc7776b60b4 *Encoder) __obf_c919f124e6c2fb8b(__obf_4909ae60ffbb8e53 int64) error {
	if __obf_77240dc7776b60b4.__obf_3cf0882fa5a4cafb&__obf_506e82775380d9f7 != 0 {
		return __obf_77240dc7776b60b4.EncodeInt(__obf_4909ae60ffbb8e53)
	}
	return __obf_77240dc7776b60b4.EncodeInt64(__obf_4909ae60ffbb8e53)
}

// EncodeUnsignedNumber encodes an uint64 in 1, 2, 3, 5, or 9 bytes.
// Type of the number is lost during encoding.
func (__obf_77240dc7776b60b4 *Encoder) EncodeUint(__obf_4909ae60ffbb8e53 uint64) error {
	if __obf_4909ae60ffbb8e53 <= math.MaxInt8 {
		return __obf_77240dc7776b60b4.__obf_279ea11cfd0ea98d.WriteByte(byte(__obf_4909ae60ffbb8e53))
	}
	if __obf_4909ae60ffbb8e53 <= math.MaxUint8 {
		return __obf_77240dc7776b60b4.EncodeUint8(uint8(__obf_4909ae60ffbb8e53))
	}
	if __obf_4909ae60ffbb8e53 <= math.MaxUint16 {
		return __obf_77240dc7776b60b4.EncodeUint16(uint16(__obf_4909ae60ffbb8e53))
	}
	if __obf_4909ae60ffbb8e53 <= math.MaxUint32 {
		return __obf_77240dc7776b60b4.EncodeUint32(uint32(__obf_4909ae60ffbb8e53))
	}
	return __obf_77240dc7776b60b4.EncodeUint64(__obf_4909ae60ffbb8e53)
}

// EncodeNumber encodes an int64 in 1, 2, 3, 5, or 9 bytes.
// Type of the number is lost during encoding.
func (__obf_77240dc7776b60b4 *Encoder) EncodeInt(__obf_4909ae60ffbb8e53 int64) error {
	if __obf_4909ae60ffbb8e53 >= 0 {
		return __obf_77240dc7776b60b4.EncodeUint(uint64(__obf_4909ae60ffbb8e53))
	}
	if __obf_4909ae60ffbb8e53 >= int64(int8(NegFixedNumLow)) {
		return __obf_77240dc7776b60b4.__obf_279ea11cfd0ea98d.WriteByte(byte(__obf_4909ae60ffbb8e53))
	}
	if __obf_4909ae60ffbb8e53 >= math.MinInt8 {
		return __obf_77240dc7776b60b4.EncodeInt8(int8(__obf_4909ae60ffbb8e53))
	}
	if __obf_4909ae60ffbb8e53 >= math.MinInt16 {
		return __obf_77240dc7776b60b4.EncodeInt16(int16(__obf_4909ae60ffbb8e53))
	}
	if __obf_4909ae60ffbb8e53 >= math.MinInt32 {
		return __obf_77240dc7776b60b4.EncodeInt32(int32(__obf_4909ae60ffbb8e53))
	}
	return __obf_77240dc7776b60b4.EncodeInt64(__obf_4909ae60ffbb8e53)
}

func (__obf_77240dc7776b60b4 *Encoder) EncodeFloat32(__obf_4909ae60ffbb8e53 float32) error {
	if __obf_77240dc7776b60b4.__obf_3cf0882fa5a4cafb&__obf_7ad13a8ba1bf6d52 != 0 {
		if float32(int64(__obf_4909ae60ffbb8e53)) == __obf_4909ae60ffbb8e53 {
			return __obf_77240dc7776b60b4.EncodeInt(int64(__obf_4909ae60ffbb8e53))
		}
	}
	return __obf_77240dc7776b60b4.__obf_70cbe6e3e07b6347(Float, math.Float32bits(__obf_4909ae60ffbb8e53))
}

func (__obf_77240dc7776b60b4 *Encoder) EncodeFloat64(__obf_4909ae60ffbb8e53 float64) error {
	if __obf_77240dc7776b60b4.__obf_3cf0882fa5a4cafb&__obf_7ad13a8ba1bf6d52 != 0 {
		// Both NaN and Inf convert to int64(-0x8000000000000000)
		// If n is NaN then it never compares true with any other value
		// If n is Inf then it doesn't convert from int64 back to +/-Inf
		// In both cases the comparison works.
		if float64(int64(__obf_4909ae60ffbb8e53)) == __obf_4909ae60ffbb8e53 {
			return __obf_77240dc7776b60b4.EncodeInt(int64(__obf_4909ae60ffbb8e53))
		}
	}
	return __obf_77240dc7776b60b4.__obf_4003fdc721833a64(Double, math.Float64bits(__obf_4909ae60ffbb8e53))
}

func (__obf_77240dc7776b60b4 *Encoder) __obf_a5ea8b1a318bf9eb(__obf_545372fefbb733e5 byte, __obf_4909ae60ffbb8e53 uint8) error {
	__obf_77240dc7776b60b4.__obf_27e4a1a33a31e4a7 = __obf_77240dc7776b60b4.__obf_27e4a1a33a31e4a7[:2]
	__obf_77240dc7776b60b4.__obf_27e4a1a33a31e4a7[0] = __obf_545372fefbb733e5
	__obf_77240dc7776b60b4.__obf_27e4a1a33a31e4a7[1] = __obf_4909ae60ffbb8e53
	return __obf_77240dc7776b60b4.__obf_4432f415b3007804(__obf_77240dc7776b60b4.__obf_27e4a1a33a31e4a7)
}

func (__obf_77240dc7776b60b4 *Encoder) __obf_9ecc220d5defc1d9(__obf_545372fefbb733e5 byte, __obf_4909ae60ffbb8e53 uint16) error {
	__obf_77240dc7776b60b4.__obf_27e4a1a33a31e4a7 = __obf_77240dc7776b60b4.__obf_27e4a1a33a31e4a7[:3]
	__obf_77240dc7776b60b4.__obf_27e4a1a33a31e4a7[0] = __obf_545372fefbb733e5
	__obf_77240dc7776b60b4.__obf_27e4a1a33a31e4a7[1] = byte(__obf_4909ae60ffbb8e53 >> 8)
	__obf_77240dc7776b60b4.__obf_27e4a1a33a31e4a7[2] = byte(__obf_4909ae60ffbb8e53)
	return __obf_77240dc7776b60b4.__obf_4432f415b3007804(__obf_77240dc7776b60b4.__obf_27e4a1a33a31e4a7)
}

func (__obf_77240dc7776b60b4 *Encoder) __obf_70cbe6e3e07b6347(__obf_545372fefbb733e5 byte, __obf_4909ae60ffbb8e53 uint32) error {
	__obf_77240dc7776b60b4.__obf_27e4a1a33a31e4a7 = __obf_77240dc7776b60b4.__obf_27e4a1a33a31e4a7[:5]
	__obf_77240dc7776b60b4.__obf_27e4a1a33a31e4a7[0] = __obf_545372fefbb733e5
	__obf_77240dc7776b60b4.__obf_27e4a1a33a31e4a7[1] = byte(__obf_4909ae60ffbb8e53 >> 24)
	__obf_77240dc7776b60b4.__obf_27e4a1a33a31e4a7[2] = byte(__obf_4909ae60ffbb8e53 >> 16)
	__obf_77240dc7776b60b4.__obf_27e4a1a33a31e4a7[3] = byte(__obf_4909ae60ffbb8e53 >> 8)
	__obf_77240dc7776b60b4.__obf_27e4a1a33a31e4a7[4] = byte(__obf_4909ae60ffbb8e53)
	return __obf_77240dc7776b60b4.__obf_4432f415b3007804(__obf_77240dc7776b60b4.__obf_27e4a1a33a31e4a7)
}

func (__obf_77240dc7776b60b4 *Encoder) __obf_4003fdc721833a64(__obf_545372fefbb733e5 byte, __obf_4909ae60ffbb8e53 uint64) error {
	__obf_77240dc7776b60b4.__obf_27e4a1a33a31e4a7 = __obf_77240dc7776b60b4.__obf_27e4a1a33a31e4a7[:9]
	__obf_77240dc7776b60b4.__obf_27e4a1a33a31e4a7[0] = __obf_545372fefbb733e5
	__obf_77240dc7776b60b4.__obf_27e4a1a33a31e4a7[1] = byte(__obf_4909ae60ffbb8e53 >> 56)
	__obf_77240dc7776b60b4.__obf_27e4a1a33a31e4a7[2] = byte(__obf_4909ae60ffbb8e53 >> 48)
	__obf_77240dc7776b60b4.__obf_27e4a1a33a31e4a7[3] = byte(__obf_4909ae60ffbb8e53 >> 40)
	__obf_77240dc7776b60b4.__obf_27e4a1a33a31e4a7[4] = byte(__obf_4909ae60ffbb8e53 >> 32)
	__obf_77240dc7776b60b4.__obf_27e4a1a33a31e4a7[5] = byte(__obf_4909ae60ffbb8e53 >> 24)
	__obf_77240dc7776b60b4.__obf_27e4a1a33a31e4a7[6] = byte(__obf_4909ae60ffbb8e53 >> 16)
	__obf_77240dc7776b60b4.__obf_27e4a1a33a31e4a7[7] = byte(__obf_4909ae60ffbb8e53 >> 8)
	__obf_77240dc7776b60b4.__obf_27e4a1a33a31e4a7[8] = byte(__obf_4909ae60ffbb8e53)
	return __obf_77240dc7776b60b4.__obf_4432f415b3007804(__obf_77240dc7776b60b4.__obf_27e4a1a33a31e4a7)
}

func __obf_c079cfc8e2339c96(__obf_77240dc7776b60b4 *Encoder, __obf_63bbcee86d44fdde reflect.Value) error {
	return __obf_77240dc7776b60b4.EncodeUint(__obf_63bbcee86d44fdde.Uint())
}

func __obf_da00a84d1175f148(__obf_77240dc7776b60b4 *Encoder, __obf_63bbcee86d44fdde reflect.Value) error {
	return __obf_77240dc7776b60b4.EncodeInt(__obf_63bbcee86d44fdde.Int())
}

func __obf_bfc025a7f29928a9(__obf_77240dc7776b60b4 *Encoder, __obf_63bbcee86d44fdde reflect.Value) error {
	return __obf_77240dc7776b60b4.__obf_4c4fec685e237b13(uint8(__obf_63bbcee86d44fdde.Uint()))
}

func __obf_81f7a2d6252d3484(__obf_77240dc7776b60b4 *Encoder, __obf_63bbcee86d44fdde reflect.Value) error {
	return __obf_77240dc7776b60b4.__obf_19dd028d0c8c035e(uint16(__obf_63bbcee86d44fdde.Uint()))
}

func __obf_37b3e9bbb21b1a1b(__obf_77240dc7776b60b4 *Encoder, __obf_63bbcee86d44fdde reflect.Value) error {
	return __obf_77240dc7776b60b4.__obf_d9ac25e44ae7c9f2(uint32(__obf_63bbcee86d44fdde.Uint()))
}

func __obf_f2dc55a2f1f1d57e(__obf_77240dc7776b60b4 *Encoder, __obf_63bbcee86d44fdde reflect.Value) error {
	return __obf_77240dc7776b60b4.__obf_e32b04a6bb97fb7e(__obf_63bbcee86d44fdde.Uint())
}

func __obf_38a64e9c8c30348f(__obf_77240dc7776b60b4 *Encoder, __obf_63bbcee86d44fdde reflect.Value) error {
	return __obf_77240dc7776b60b4.__obf_6ba66788d0a0ac7b(int8(__obf_63bbcee86d44fdde.Int()))
}

func __obf_192fd05a8152c0cf(__obf_77240dc7776b60b4 *Encoder, __obf_63bbcee86d44fdde reflect.Value) error {
	return __obf_77240dc7776b60b4.__obf_73b9261f3dd2dce8(int16(__obf_63bbcee86d44fdde.Int()))
}

func __obf_60d6bd6b5e468734(__obf_77240dc7776b60b4 *Encoder, __obf_63bbcee86d44fdde reflect.Value) error {
	return __obf_77240dc7776b60b4.__obf_18b40d4bd23a9cae(int32(__obf_63bbcee86d44fdde.Int()))
}

func __obf_f9c867edad184432(__obf_77240dc7776b60b4 *Encoder, __obf_63bbcee86d44fdde reflect.Value) error {
	return __obf_77240dc7776b60b4.__obf_c919f124e6c2fb8b(__obf_63bbcee86d44fdde.Int())
}

func __obf_7de64ad2273f4b55(__obf_77240dc7776b60b4 *Encoder, __obf_63bbcee86d44fdde reflect.Value) error {
	return __obf_77240dc7776b60b4.EncodeFloat32(float32(__obf_63bbcee86d44fdde.Float()))
}

func __obf_239ea56f7cf6dcb6(__obf_77240dc7776b60b4 *Encoder, __obf_63bbcee86d44fdde reflect.Value) error {
	return __obf_77240dc7776b60b4.EncodeFloat64(__obf_63bbcee86d44fdde.Float())
}
