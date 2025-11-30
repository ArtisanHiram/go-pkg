package __obf_c3cd04a15c56f32f

import (
	"fmt"
	"math"
	"strconv"
)

var __obf_cef54dda9534fb7d []uint64

func init() {
	__obf_cef54dda9534fb7d = []uint64{1, 10, 100, 1000, 10000, 100000, 1000000}
}

// WriteFloat32 write float32 to stream
func (__obf_2361f5214e84c60f *Stream) WriteFloat32(__obf_d59813f23ad740a8 float32) {
	if math.IsInf(float64(__obf_d59813f23ad740a8), 0) || math.IsNaN(float64(__obf_d59813f23ad740a8)) {
		__obf_2361f5214e84c60f.
			Error = fmt.Errorf("unsupported value: %f", __obf_d59813f23ad740a8)
		return
	}
	__obf_447df7ce5337877c := math.Abs(float64(__obf_d59813f23ad740a8))
	fmt := byte('f')
	// Note: Must use float32 comparisons for underlying float32 value to get precise cutoffs right.
	if __obf_447df7ce5337877c != 0 {
		if float32(__obf_447df7ce5337877c) < 1e-6 || float32(__obf_447df7ce5337877c) >= 1e21 {
			fmt = 'e'
		}
	}
	__obf_2361f5214e84c60f.__obf_ace979f6622823f3 = strconv.AppendFloat(__obf_2361f5214e84c60f.__obf_ace979f6622823f3, float64(__obf_d59813f23ad740a8), fmt, -1, 32)
	if fmt == 'e' {
		__obf_fd4464bb6b13cabd := // clean up e-09 to e-9
			len(__obf_2361f5214e84c60f.__obf_ace979f6622823f3)
		if __obf_fd4464bb6b13cabd >= 4 && __obf_2361f5214e84c60f.__obf_ace979f6622823f3[__obf_fd4464bb6b13cabd-4] == 'e' && __obf_2361f5214e84c60f.__obf_ace979f6622823f3[__obf_fd4464bb6b13cabd-3] == '-' && __obf_2361f5214e84c60f.__obf_ace979f6622823f3[__obf_fd4464bb6b13cabd-2] == '0' {
			__obf_2361f5214e84c60f.__obf_ace979f6622823f3[__obf_fd4464bb6b13cabd-2] = __obf_2361f5214e84c60f.__obf_ace979f6622823f3[__obf_fd4464bb6b13cabd-1]
			__obf_2361f5214e84c60f.__obf_ace979f6622823f3 = __obf_2361f5214e84c60f.__obf_ace979f6622823f3[:__obf_fd4464bb6b13cabd-1]
		}
	}
}

// WriteFloat32Lossy write float32 to stream with ONLY 6 digits precision although much much faster
func (__obf_2361f5214e84c60f *Stream) WriteFloat32Lossy(__obf_d59813f23ad740a8 float32) {
	if math.IsInf(float64(__obf_d59813f23ad740a8), 0) || math.IsNaN(float64(__obf_d59813f23ad740a8)) {
		__obf_2361f5214e84c60f.
			Error = fmt.Errorf("unsupported value: %f", __obf_d59813f23ad740a8)
		return
	}
	if __obf_d59813f23ad740a8 < 0 {
		__obf_2361f5214e84c60f.__obf_c4fec0edfb3875ad('-')
		__obf_d59813f23ad740a8 = -__obf_d59813f23ad740a8
	}
	if __obf_d59813f23ad740a8 > 0x4ffffff {
		__obf_2361f5214e84c60f.
			WriteFloat32(__obf_d59813f23ad740a8)
		return
	}
	__obf_047dd9606d98250f := 6
	__obf_de3ef9b2df6fdb2d := uint64(1000000)
	__obf_e90b59cd02555399 := // 6
		uint64(float64(__obf_d59813f23ad740a8)*float64(__obf_de3ef9b2df6fdb2d) + 0.5)
	__obf_2361f5214e84c60f.
		WriteUint64(__obf_e90b59cd02555399 / __obf_de3ef9b2df6fdb2d)
	__obf_d28ae2e2ffe8774d := __obf_e90b59cd02555399 % __obf_de3ef9b2df6fdb2d
	if __obf_d28ae2e2ffe8774d == 0 {
		return
	}
	__obf_2361f5214e84c60f.__obf_c4fec0edfb3875ad('.')
	for __obf_420b1bf30275e403 := __obf_047dd9606d98250f - 1; __obf_420b1bf30275e403 > 0 && __obf_d28ae2e2ffe8774d < __obf_cef54dda9534fb7d[__obf_420b1bf30275e403]; __obf_420b1bf30275e403-- {
		__obf_2361f5214e84c60f.__obf_c4fec0edfb3875ad('0')
	}
	__obf_2361f5214e84c60f.
		WriteUint64(__obf_d28ae2e2ffe8774d)
	for __obf_2361f5214e84c60f.__obf_ace979f6622823f3[len(__obf_2361f5214e84c60f.__obf_ace979f6622823f3)-1] == '0' {
		__obf_2361f5214e84c60f.__obf_ace979f6622823f3 = __obf_2361f5214e84c60f.__obf_ace979f6622823f3[:len(__obf_2361f5214e84c60f.

			// WriteFloat64 write float64 to stream
			__obf_ace979f6622823f3)-1]
	}
}

func (__obf_2361f5214e84c60f *Stream) WriteFloat64(__obf_d59813f23ad740a8 float64) {
	if math.IsInf(__obf_d59813f23ad740a8, 0) || math.IsNaN(__obf_d59813f23ad740a8) {
		__obf_2361f5214e84c60f.
			Error = fmt.Errorf("unsupported value: %f", __obf_d59813f23ad740a8)
		return
	}
	__obf_447df7ce5337877c := math.Abs(__obf_d59813f23ad740a8)
	fmt := byte('f')
	// Note: Must use float32 comparisons for underlying float32 value to get precise cutoffs right.
	if __obf_447df7ce5337877c != 0 {
		if __obf_447df7ce5337877c < 1e-6 || __obf_447df7ce5337877c >= 1e21 {
			fmt = 'e'
		}
	}
	__obf_2361f5214e84c60f.__obf_ace979f6622823f3 = strconv.AppendFloat(__obf_2361f5214e84c60f.__obf_ace979f6622823f3, float64(__obf_d59813f23ad740a8), fmt, -1, 64)
	if fmt == 'e' {
		__obf_fd4464bb6b13cabd := // clean up e-09 to e-9
			len(__obf_2361f5214e84c60f.__obf_ace979f6622823f3)
		if __obf_fd4464bb6b13cabd >= 4 && __obf_2361f5214e84c60f.__obf_ace979f6622823f3[__obf_fd4464bb6b13cabd-4] == 'e' && __obf_2361f5214e84c60f.__obf_ace979f6622823f3[__obf_fd4464bb6b13cabd-3] == '-' && __obf_2361f5214e84c60f.__obf_ace979f6622823f3[__obf_fd4464bb6b13cabd-2] == '0' {
			__obf_2361f5214e84c60f.__obf_ace979f6622823f3[__obf_fd4464bb6b13cabd-2] = __obf_2361f5214e84c60f.__obf_ace979f6622823f3[__obf_fd4464bb6b13cabd-1]
			__obf_2361f5214e84c60f.__obf_ace979f6622823f3 = __obf_2361f5214e84c60f.__obf_ace979f6622823f3[:__obf_fd4464bb6b13cabd-1]
		}
	}
}

// WriteFloat64Lossy write float64 to stream with ONLY 6 digits precision although much much faster
func (__obf_2361f5214e84c60f *Stream) WriteFloat64Lossy(__obf_d59813f23ad740a8 float64) {
	if math.IsInf(__obf_d59813f23ad740a8, 0) || math.IsNaN(__obf_d59813f23ad740a8) {
		__obf_2361f5214e84c60f.
			Error = fmt.Errorf("unsupported value: %f", __obf_d59813f23ad740a8)
		return
	}
	if __obf_d59813f23ad740a8 < 0 {
		__obf_2361f5214e84c60f.__obf_c4fec0edfb3875ad('-')
		__obf_d59813f23ad740a8 = -__obf_d59813f23ad740a8
	}
	if __obf_d59813f23ad740a8 > 0x4ffffff {
		__obf_2361f5214e84c60f.
			WriteFloat64(__obf_d59813f23ad740a8)
		return
	}
	__obf_047dd9606d98250f := 6
	__obf_de3ef9b2df6fdb2d := uint64(1000000)
	__obf_e90b59cd02555399 := // 6
		uint64(__obf_d59813f23ad740a8*float64(__obf_de3ef9b2df6fdb2d) + 0.5)
	__obf_2361f5214e84c60f.
		WriteUint64(__obf_e90b59cd02555399 / __obf_de3ef9b2df6fdb2d)
	__obf_d28ae2e2ffe8774d := __obf_e90b59cd02555399 % __obf_de3ef9b2df6fdb2d
	if __obf_d28ae2e2ffe8774d == 0 {
		return
	}
	__obf_2361f5214e84c60f.__obf_c4fec0edfb3875ad('.')
	for __obf_420b1bf30275e403 := __obf_047dd9606d98250f - 1; __obf_420b1bf30275e403 > 0 && __obf_d28ae2e2ffe8774d < __obf_cef54dda9534fb7d[__obf_420b1bf30275e403]; __obf_420b1bf30275e403-- {
		__obf_2361f5214e84c60f.__obf_c4fec0edfb3875ad('0')
	}
	__obf_2361f5214e84c60f.
		WriteUint64(__obf_d28ae2e2ffe8774d)
	for __obf_2361f5214e84c60f.__obf_ace979f6622823f3[len(__obf_2361f5214e84c60f.__obf_ace979f6622823f3)-1] == '0' {
		__obf_2361f5214e84c60f.__obf_ace979f6622823f3 = __obf_2361f5214e84c60f.__obf_ace979f6622823f3[:len(__obf_2361f5214e84c60f.__obf_ace979f6622823f3)-1]
	}
}
