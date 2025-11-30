package __obf_fd770cb9675903b0

import (
	"math"
	"reflect"
)

// EncodeUint8 encodes an uint8 in 2 bytes preserving type of the number.
func (__obf_e9038cda3b5cf711 *Encoder) EncodeUint8(__obf_d774e4844c74bfe9 uint8) error {
	return __obf_e9038cda3b5cf711.__obf_c58302e6ea7043c4(Uint8, __obf_d774e4844c74bfe9)
}

func (__obf_e9038cda3b5cf711 *Encoder) __obf_3fff860f621acbfb(__obf_d774e4844c74bfe9 uint8) error {
	if __obf_e9038cda3b5cf711.__obf_7185cb62de7af792&__obf_a787b0ded9ae6f01 != 0 {
		return __obf_e9038cda3b5cf711.EncodeUint(uint64(__obf_d774e4844c74bfe9))
	}
	return __obf_e9038cda3b5cf711.EncodeUint8(__obf_d774e4844c74bfe9)
}

// EncodeUint16 encodes an uint16 in 3 bytes preserving type of the number.
func (__obf_e9038cda3b5cf711 *Encoder) EncodeUint16(__obf_d774e4844c74bfe9 uint16) error {
	return __obf_e9038cda3b5cf711.__obf_a98b5ca0c9f03534(Uint16, __obf_d774e4844c74bfe9)
}

func (__obf_e9038cda3b5cf711 *Encoder) __obf_76a0cb980ddece7b(__obf_d774e4844c74bfe9 uint16) error {
	if __obf_e9038cda3b5cf711.__obf_7185cb62de7af792&__obf_a787b0ded9ae6f01 != 0 {
		return __obf_e9038cda3b5cf711.EncodeUint(uint64(__obf_d774e4844c74bfe9))
	}
	return __obf_e9038cda3b5cf711.EncodeUint16(__obf_d774e4844c74bfe9)
}

// EncodeUint32 encodes an uint16 in 5 bytes preserving type of the number.
func (__obf_e9038cda3b5cf711 *Encoder) EncodeUint32(__obf_d774e4844c74bfe9 uint32) error {
	return __obf_e9038cda3b5cf711.__obf_e0e3a3f6b942651e(Uint32, __obf_d774e4844c74bfe9)
}

func (__obf_e9038cda3b5cf711 *Encoder) __obf_caea9e41121095d6(__obf_d774e4844c74bfe9 uint32) error {
	if __obf_e9038cda3b5cf711.__obf_7185cb62de7af792&__obf_a787b0ded9ae6f01 != 0 {
		return __obf_e9038cda3b5cf711.EncodeUint(uint64(__obf_d774e4844c74bfe9))
	}
	return __obf_e9038cda3b5cf711.EncodeUint32(__obf_d774e4844c74bfe9)
}

// EncodeUint64 encodes an uint16 in 9 bytes preserving type of the number.
func (__obf_e9038cda3b5cf711 *Encoder) EncodeUint64(__obf_d774e4844c74bfe9 uint64) error {
	return __obf_e9038cda3b5cf711.__obf_f1fce69f1cdc6606(Uint64, __obf_d774e4844c74bfe9)
}

func (__obf_e9038cda3b5cf711 *Encoder) __obf_c6c82d67d256dc52(__obf_d774e4844c74bfe9 uint64) error {
	if __obf_e9038cda3b5cf711.__obf_7185cb62de7af792&__obf_a787b0ded9ae6f01 != 0 {
		return __obf_e9038cda3b5cf711.EncodeUint(__obf_d774e4844c74bfe9)
	}
	return __obf_e9038cda3b5cf711.EncodeUint64(__obf_d774e4844c74bfe9)
}

// EncodeInt8 encodes an int8 in 2 bytes preserving type of the number.
func (__obf_e9038cda3b5cf711 *Encoder) EncodeInt8(__obf_d774e4844c74bfe9 int8) error {
	return __obf_e9038cda3b5cf711.__obf_c58302e6ea7043c4(Int8, uint8(__obf_d774e4844c74bfe9))
}

func (__obf_e9038cda3b5cf711 *Encoder) __obf_a2d7c83a18736dbc(__obf_d774e4844c74bfe9 int8) error {
	if __obf_e9038cda3b5cf711.__obf_7185cb62de7af792&__obf_a787b0ded9ae6f01 != 0 {
		return __obf_e9038cda3b5cf711.EncodeInt(int64(__obf_d774e4844c74bfe9))
	}
	return __obf_e9038cda3b5cf711.EncodeInt8(__obf_d774e4844c74bfe9)
}

// EncodeInt16 encodes an int16 in 3 bytes preserving type of the number.
func (__obf_e9038cda3b5cf711 *Encoder) EncodeInt16(__obf_d774e4844c74bfe9 int16) error {
	return __obf_e9038cda3b5cf711.__obf_a98b5ca0c9f03534(Int16, uint16(__obf_d774e4844c74bfe9))
}

func (__obf_e9038cda3b5cf711 *Encoder) __obf_b8a6a74732178288(__obf_d774e4844c74bfe9 int16) error {
	if __obf_e9038cda3b5cf711.__obf_7185cb62de7af792&__obf_a787b0ded9ae6f01 != 0 {
		return __obf_e9038cda3b5cf711.EncodeInt(int64(__obf_d774e4844c74bfe9))
	}
	return __obf_e9038cda3b5cf711.EncodeInt16(__obf_d774e4844c74bfe9)
}

// EncodeInt32 encodes an int32 in 5 bytes preserving type of the number.
func (__obf_e9038cda3b5cf711 *Encoder) EncodeInt32(__obf_d774e4844c74bfe9 int32) error {
	return __obf_e9038cda3b5cf711.__obf_e0e3a3f6b942651e(Int32, uint32(__obf_d774e4844c74bfe9))
}

func (__obf_e9038cda3b5cf711 *Encoder) __obf_cb9da83d938d0232(__obf_d774e4844c74bfe9 int32) error {
	if __obf_e9038cda3b5cf711.__obf_7185cb62de7af792&__obf_a787b0ded9ae6f01 != 0 {
		return __obf_e9038cda3b5cf711.EncodeInt(int64(__obf_d774e4844c74bfe9))
	}
	return __obf_e9038cda3b5cf711.EncodeInt32(__obf_d774e4844c74bfe9)
}

// EncodeInt64 encodes an int64 in 9 bytes preserving type of the number.
func (__obf_e9038cda3b5cf711 *Encoder) EncodeInt64(__obf_d774e4844c74bfe9 int64) error {
	return __obf_e9038cda3b5cf711.__obf_f1fce69f1cdc6606(Int64, uint64(__obf_d774e4844c74bfe9))
}

func (__obf_e9038cda3b5cf711 *Encoder) __obf_eac777c6448d3606(__obf_d774e4844c74bfe9 int64) error {
	if __obf_e9038cda3b5cf711.__obf_7185cb62de7af792&__obf_a787b0ded9ae6f01 != 0 {
		return __obf_e9038cda3b5cf711.EncodeInt(__obf_d774e4844c74bfe9)
	}
	return __obf_e9038cda3b5cf711.EncodeInt64(__obf_d774e4844c74bfe9)
}

// EncodeUnsignedNumber encodes an uint64 in 1, 2, 3, 5, or 9 bytes.
// Type of the number is lost during encoding.
func (__obf_e9038cda3b5cf711 *Encoder) EncodeUint(__obf_d774e4844c74bfe9 uint64) error {
	if __obf_d774e4844c74bfe9 <= math.MaxInt8 {
		return __obf_e9038cda3b5cf711.__obf_41c5519e4a6d21b0.WriteByte(byte(__obf_d774e4844c74bfe9))
	}
	if __obf_d774e4844c74bfe9 <= math.MaxUint8 {
		return __obf_e9038cda3b5cf711.EncodeUint8(uint8(__obf_d774e4844c74bfe9))
	}
	if __obf_d774e4844c74bfe9 <= math.MaxUint16 {
		return __obf_e9038cda3b5cf711.EncodeUint16(uint16(__obf_d774e4844c74bfe9))
	}
	if __obf_d774e4844c74bfe9 <= math.MaxUint32 {
		return __obf_e9038cda3b5cf711.EncodeUint32(uint32(__obf_d774e4844c74bfe9))
	}
	return __obf_e9038cda3b5cf711.EncodeUint64(__obf_d774e4844c74bfe9)
}

// EncodeNumber encodes an int64 in 1, 2, 3, 5, or 9 bytes.
// Type of the number is lost during encoding.
func (__obf_e9038cda3b5cf711 *Encoder) EncodeInt(__obf_d774e4844c74bfe9 int64) error {
	if __obf_d774e4844c74bfe9 >= 0 {
		return __obf_e9038cda3b5cf711.EncodeUint(uint64(__obf_d774e4844c74bfe9))
	}
	if __obf_d774e4844c74bfe9 >= int64(int8(NegFixedNumLow)) {
		return __obf_e9038cda3b5cf711.__obf_41c5519e4a6d21b0.WriteByte(byte(__obf_d774e4844c74bfe9))
	}
	if __obf_d774e4844c74bfe9 >= math.MinInt8 {
		return __obf_e9038cda3b5cf711.EncodeInt8(int8(__obf_d774e4844c74bfe9))
	}
	if __obf_d774e4844c74bfe9 >= math.MinInt16 {
		return __obf_e9038cda3b5cf711.EncodeInt16(int16(__obf_d774e4844c74bfe9))
	}
	if __obf_d774e4844c74bfe9 >= math.MinInt32 {
		return __obf_e9038cda3b5cf711.EncodeInt32(int32(__obf_d774e4844c74bfe9))
	}
	return __obf_e9038cda3b5cf711.EncodeInt64(__obf_d774e4844c74bfe9)
}

func (__obf_e9038cda3b5cf711 *Encoder) EncodeFloat32(__obf_d774e4844c74bfe9 float32) error {
	if __obf_e9038cda3b5cf711.__obf_7185cb62de7af792&__obf_19c02dcbfe4a74e9 != 0 {
		if float32(int64(__obf_d774e4844c74bfe9)) == __obf_d774e4844c74bfe9 {
			return __obf_e9038cda3b5cf711.EncodeInt(int64(__obf_d774e4844c74bfe9))
		}
	}
	return __obf_e9038cda3b5cf711.__obf_e0e3a3f6b942651e(Float, math.Float32bits(__obf_d774e4844c74bfe9))
}

func (__obf_e9038cda3b5cf711 *Encoder) EncodeFloat64(__obf_d774e4844c74bfe9 float64) error {
	if __obf_e9038cda3b5cf711.__obf_7185cb62de7af792&__obf_19c02dcbfe4a74e9 != 0 {
		// Both NaN and Inf convert to int64(-0x8000000000000000)
		// If n is NaN then it never compares true with any other value
		// If n is Inf then it doesn't convert from int64 back to +/-Inf
		// In both cases the comparison works.
		if float64(int64(__obf_d774e4844c74bfe9)) == __obf_d774e4844c74bfe9 {
			return __obf_e9038cda3b5cf711.EncodeInt(int64(__obf_d774e4844c74bfe9))
		}
	}
	return __obf_e9038cda3b5cf711.__obf_f1fce69f1cdc6606(Double, math.Float64bits(__obf_d774e4844c74bfe9))
}

func (__obf_e9038cda3b5cf711 *Encoder) __obf_c58302e6ea7043c4(__obf_cde82889ba8e4822 byte, __obf_d774e4844c74bfe9 uint8) error {
	__obf_e9038cda3b5cf711.__obf_099b1f8c6882e66a = __obf_e9038cda3b5cf711.__obf_099b1f8c6882e66a[:2]
	__obf_e9038cda3b5cf711.__obf_099b1f8c6882e66a[0] = __obf_cde82889ba8e4822
	__obf_e9038cda3b5cf711.__obf_099b1f8c6882e66a[1] = __obf_d774e4844c74bfe9
	return __obf_e9038cda3b5cf711.__obf_af6d14a7babbd464(__obf_e9038cda3b5cf711.__obf_099b1f8c6882e66a)
}

func (__obf_e9038cda3b5cf711 *Encoder) __obf_a98b5ca0c9f03534(__obf_cde82889ba8e4822 byte, __obf_d774e4844c74bfe9 uint16) error {
	__obf_e9038cda3b5cf711.__obf_099b1f8c6882e66a = __obf_e9038cda3b5cf711.__obf_099b1f8c6882e66a[:3]
	__obf_e9038cda3b5cf711.__obf_099b1f8c6882e66a[0] = __obf_cde82889ba8e4822
	__obf_e9038cda3b5cf711.__obf_099b1f8c6882e66a[1] = byte(__obf_d774e4844c74bfe9 >> 8)
	__obf_e9038cda3b5cf711.__obf_099b1f8c6882e66a[2] = byte(__obf_d774e4844c74bfe9)
	return __obf_e9038cda3b5cf711.__obf_af6d14a7babbd464(__obf_e9038cda3b5cf711.__obf_099b1f8c6882e66a)
}

func (__obf_e9038cda3b5cf711 *Encoder) __obf_e0e3a3f6b942651e(__obf_cde82889ba8e4822 byte, __obf_d774e4844c74bfe9 uint32) error {
	__obf_e9038cda3b5cf711.__obf_099b1f8c6882e66a = __obf_e9038cda3b5cf711.__obf_099b1f8c6882e66a[:5]
	__obf_e9038cda3b5cf711.__obf_099b1f8c6882e66a[0] = __obf_cde82889ba8e4822
	__obf_e9038cda3b5cf711.__obf_099b1f8c6882e66a[1] = byte(__obf_d774e4844c74bfe9 >> 24)
	__obf_e9038cda3b5cf711.__obf_099b1f8c6882e66a[2] = byte(__obf_d774e4844c74bfe9 >> 16)
	__obf_e9038cda3b5cf711.__obf_099b1f8c6882e66a[3] = byte(__obf_d774e4844c74bfe9 >> 8)
	__obf_e9038cda3b5cf711.__obf_099b1f8c6882e66a[4] = byte(__obf_d774e4844c74bfe9)
	return __obf_e9038cda3b5cf711.__obf_af6d14a7babbd464(__obf_e9038cda3b5cf711.__obf_099b1f8c6882e66a)
}

func (__obf_e9038cda3b5cf711 *Encoder) __obf_f1fce69f1cdc6606(__obf_cde82889ba8e4822 byte, __obf_d774e4844c74bfe9 uint64) error {
	__obf_e9038cda3b5cf711.__obf_099b1f8c6882e66a = __obf_e9038cda3b5cf711.__obf_099b1f8c6882e66a[:9]
	__obf_e9038cda3b5cf711.__obf_099b1f8c6882e66a[0] = __obf_cde82889ba8e4822
	__obf_e9038cda3b5cf711.__obf_099b1f8c6882e66a[1] = byte(__obf_d774e4844c74bfe9 >> 56)
	__obf_e9038cda3b5cf711.__obf_099b1f8c6882e66a[2] = byte(__obf_d774e4844c74bfe9 >> 48)
	__obf_e9038cda3b5cf711.__obf_099b1f8c6882e66a[3] = byte(__obf_d774e4844c74bfe9 >> 40)
	__obf_e9038cda3b5cf711.__obf_099b1f8c6882e66a[4] = byte(__obf_d774e4844c74bfe9 >> 32)
	__obf_e9038cda3b5cf711.__obf_099b1f8c6882e66a[5] = byte(__obf_d774e4844c74bfe9 >> 24)
	__obf_e9038cda3b5cf711.__obf_099b1f8c6882e66a[6] = byte(__obf_d774e4844c74bfe9 >> 16)
	__obf_e9038cda3b5cf711.__obf_099b1f8c6882e66a[7] = byte(__obf_d774e4844c74bfe9 >> 8)
	__obf_e9038cda3b5cf711.__obf_099b1f8c6882e66a[8] = byte(__obf_d774e4844c74bfe9)
	return __obf_e9038cda3b5cf711.__obf_af6d14a7babbd464(__obf_e9038cda3b5cf711.__obf_099b1f8c6882e66a)
}

func __obf_d7a9f99d1b2d37cc(__obf_e9038cda3b5cf711 *Encoder, __obf_f328a048f76b7256 reflect.Value) error {
	return __obf_e9038cda3b5cf711.EncodeUint(__obf_f328a048f76b7256.Uint())
}

func __obf_3944b7d3af00ae61(__obf_e9038cda3b5cf711 *Encoder, __obf_f328a048f76b7256 reflect.Value) error {
	return __obf_e9038cda3b5cf711.EncodeInt(__obf_f328a048f76b7256.Int())
}

func __obf_4a3a44a2fc6f392a(__obf_e9038cda3b5cf711 *Encoder, __obf_f328a048f76b7256 reflect.Value) error {
	return __obf_e9038cda3b5cf711.__obf_3fff860f621acbfb(uint8(__obf_f328a048f76b7256.Uint()))
}

func __obf_2647406311f2e1f7(__obf_e9038cda3b5cf711 *Encoder, __obf_f328a048f76b7256 reflect.Value) error {
	return __obf_e9038cda3b5cf711.__obf_76a0cb980ddece7b(uint16(__obf_f328a048f76b7256.Uint()))
}

func __obf_d68e2f5e7bc5739e(__obf_e9038cda3b5cf711 *Encoder, __obf_f328a048f76b7256 reflect.Value) error {
	return __obf_e9038cda3b5cf711.__obf_caea9e41121095d6(uint32(__obf_f328a048f76b7256.Uint()))
}

func __obf_51c4091845697c82(__obf_e9038cda3b5cf711 *Encoder, __obf_f328a048f76b7256 reflect.Value) error {
	return __obf_e9038cda3b5cf711.__obf_c6c82d67d256dc52(__obf_f328a048f76b7256.Uint())
}

func __obf_dad2c9af88c11156(__obf_e9038cda3b5cf711 *Encoder, __obf_f328a048f76b7256 reflect.Value) error {
	return __obf_e9038cda3b5cf711.__obf_a2d7c83a18736dbc(int8(__obf_f328a048f76b7256.Int()))
}

func __obf_9728683ea899b38d(__obf_e9038cda3b5cf711 *Encoder, __obf_f328a048f76b7256 reflect.Value) error {
	return __obf_e9038cda3b5cf711.__obf_b8a6a74732178288(int16(__obf_f328a048f76b7256.Int()))
}

func __obf_d85919497ce72a5b(__obf_e9038cda3b5cf711 *Encoder, __obf_f328a048f76b7256 reflect.Value) error {
	return __obf_e9038cda3b5cf711.__obf_cb9da83d938d0232(int32(__obf_f328a048f76b7256.Int()))
}

func __obf_546c6c5440e640a7(__obf_e9038cda3b5cf711 *Encoder, __obf_f328a048f76b7256 reflect.Value) error {
	return __obf_e9038cda3b5cf711.__obf_eac777c6448d3606(__obf_f328a048f76b7256.Int())
}

func __obf_709c5fa255b1f690(__obf_e9038cda3b5cf711 *Encoder, __obf_f328a048f76b7256 reflect.Value) error {
	return __obf_e9038cda3b5cf711.EncodeFloat32(float32(__obf_f328a048f76b7256.Float()))
}

func __obf_f288d375e57ad4c1(__obf_e9038cda3b5cf711 *Encoder, __obf_f328a048f76b7256 reflect.Value) error {
	return __obf_e9038cda3b5cf711.EncodeFloat64(__obf_f328a048f76b7256.Float())
}
