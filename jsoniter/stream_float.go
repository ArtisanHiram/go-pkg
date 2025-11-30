package __obf_703d23ba520c3295

import (
	"fmt"
	"math"
	"strconv"
)

var __obf_ae9fbf6f0c88dcb9 []uint64

func init() {
	__obf_ae9fbf6f0c88dcb9 = []uint64{1, 10, 100, 1000, 10000, 100000, 1000000}
}

// WriteFloat32 write float32 to stream
func (__obf_9a34f62857fb1d1d *Stream) WriteFloat32(__obf_b924538fffe5fd64 float32) {
	if math.IsInf(float64(__obf_b924538fffe5fd64), 0) || math.IsNaN(float64(__obf_b924538fffe5fd64)) {
		__obf_9a34f62857fb1d1d.
			Error = fmt.Errorf("unsupported value: %f", __obf_b924538fffe5fd64)
		return
	}
	__obf_951b72145031841f := math.Abs(float64(__obf_b924538fffe5fd64))
	fmt := byte('f')
	// Note: Must use float32 comparisons for underlying float32 value to get precise cutoffs right.
	if __obf_951b72145031841f != 0 {
		if float32(__obf_951b72145031841f) < 1e-6 || float32(__obf_951b72145031841f) >= 1e21 {
			fmt = 'e'
		}
	}
	__obf_9a34f62857fb1d1d.__obf_a065f8e0da5f5952 = strconv.AppendFloat(__obf_9a34f62857fb1d1d.__obf_a065f8e0da5f5952, float64(__obf_b924538fffe5fd64), fmt, -1, 32)
	if fmt == 'e' {
		__obf_716729e0f8e419ed := // clean up e-09 to e-9
			len(__obf_9a34f62857fb1d1d.__obf_a065f8e0da5f5952)
		if __obf_716729e0f8e419ed >= 4 && __obf_9a34f62857fb1d1d.__obf_a065f8e0da5f5952[__obf_716729e0f8e419ed-4] == 'e' && __obf_9a34f62857fb1d1d.__obf_a065f8e0da5f5952[__obf_716729e0f8e419ed-3] == '-' && __obf_9a34f62857fb1d1d.__obf_a065f8e0da5f5952[__obf_716729e0f8e419ed-2] == '0' {
			__obf_9a34f62857fb1d1d.__obf_a065f8e0da5f5952[__obf_716729e0f8e419ed-2] = __obf_9a34f62857fb1d1d.__obf_a065f8e0da5f5952[__obf_716729e0f8e419ed-1]
			__obf_9a34f62857fb1d1d.__obf_a065f8e0da5f5952 = __obf_9a34f62857fb1d1d.__obf_a065f8e0da5f5952[:__obf_716729e0f8e419ed-1]
		}
	}
}

// WriteFloat32Lossy write float32 to stream with ONLY 6 digits precision although much much faster
func (__obf_9a34f62857fb1d1d *Stream) WriteFloat32Lossy(__obf_b924538fffe5fd64 float32) {
	if math.IsInf(float64(__obf_b924538fffe5fd64), 0) || math.IsNaN(float64(__obf_b924538fffe5fd64)) {
		__obf_9a34f62857fb1d1d.
			Error = fmt.Errorf("unsupported value: %f", __obf_b924538fffe5fd64)
		return
	}
	if __obf_b924538fffe5fd64 < 0 {
		__obf_9a34f62857fb1d1d.__obf_f7df2a5224462d19('-')
		__obf_b924538fffe5fd64 = -__obf_b924538fffe5fd64
	}
	if __obf_b924538fffe5fd64 > 0x4ffffff {
		__obf_9a34f62857fb1d1d.
			WriteFloat32(__obf_b924538fffe5fd64)
		return
	}
	__obf_a010560b206f5c55 := 6
	__obf_2bfd195538a9cfb2 := uint64(1000000)
	__obf_736ec6eed0c0ab90 := // 6
		uint64(float64(__obf_b924538fffe5fd64)*float64(__obf_2bfd195538a9cfb2) + 0.5)
	__obf_9a34f62857fb1d1d.
		WriteUint64(__obf_736ec6eed0c0ab90 / __obf_2bfd195538a9cfb2)
	__obf_b3f63d8bcb619cfa := __obf_736ec6eed0c0ab90 % __obf_2bfd195538a9cfb2
	if __obf_b3f63d8bcb619cfa == 0 {
		return
	}
	__obf_9a34f62857fb1d1d.__obf_f7df2a5224462d19('.')
	for __obf_f9ed235be8f52eaf := __obf_a010560b206f5c55 - 1; __obf_f9ed235be8f52eaf > 0 && __obf_b3f63d8bcb619cfa < __obf_ae9fbf6f0c88dcb9[__obf_f9ed235be8f52eaf]; __obf_f9ed235be8f52eaf-- {
		__obf_9a34f62857fb1d1d.__obf_f7df2a5224462d19('0')
	}
	__obf_9a34f62857fb1d1d.
		WriteUint64(__obf_b3f63d8bcb619cfa)
	for __obf_9a34f62857fb1d1d.__obf_a065f8e0da5f5952[len(__obf_9a34f62857fb1d1d.__obf_a065f8e0da5f5952)-1] == '0' {
		__obf_9a34f62857fb1d1d.__obf_a065f8e0da5f5952 = __obf_9a34f62857fb1d1d.__obf_a065f8e0da5f5952[:len(__obf_9a34f62857fb1d1d.

			// WriteFloat64 write float64 to stream
			__obf_a065f8e0da5f5952)-1]
	}
}

func (__obf_9a34f62857fb1d1d *Stream) WriteFloat64(__obf_b924538fffe5fd64 float64) {
	if math.IsInf(__obf_b924538fffe5fd64, 0) || math.IsNaN(__obf_b924538fffe5fd64) {
		__obf_9a34f62857fb1d1d.
			Error = fmt.Errorf("unsupported value: %f", __obf_b924538fffe5fd64)
		return
	}
	__obf_951b72145031841f := math.Abs(__obf_b924538fffe5fd64)
	fmt := byte('f')
	// Note: Must use float32 comparisons for underlying float32 value to get precise cutoffs right.
	if __obf_951b72145031841f != 0 {
		if __obf_951b72145031841f < 1e-6 || __obf_951b72145031841f >= 1e21 {
			fmt = 'e'
		}
	}
	__obf_9a34f62857fb1d1d.__obf_a065f8e0da5f5952 = strconv.AppendFloat(__obf_9a34f62857fb1d1d.__obf_a065f8e0da5f5952, float64(__obf_b924538fffe5fd64), fmt, -1, 64)
	if fmt == 'e' {
		__obf_716729e0f8e419ed := // clean up e-09 to e-9
			len(__obf_9a34f62857fb1d1d.__obf_a065f8e0da5f5952)
		if __obf_716729e0f8e419ed >= 4 && __obf_9a34f62857fb1d1d.__obf_a065f8e0da5f5952[__obf_716729e0f8e419ed-4] == 'e' && __obf_9a34f62857fb1d1d.__obf_a065f8e0da5f5952[__obf_716729e0f8e419ed-3] == '-' && __obf_9a34f62857fb1d1d.__obf_a065f8e0da5f5952[__obf_716729e0f8e419ed-2] == '0' {
			__obf_9a34f62857fb1d1d.__obf_a065f8e0da5f5952[__obf_716729e0f8e419ed-2] = __obf_9a34f62857fb1d1d.__obf_a065f8e0da5f5952[__obf_716729e0f8e419ed-1]
			__obf_9a34f62857fb1d1d.__obf_a065f8e0da5f5952 = __obf_9a34f62857fb1d1d.__obf_a065f8e0da5f5952[:__obf_716729e0f8e419ed-1]
		}
	}
}

// WriteFloat64Lossy write float64 to stream with ONLY 6 digits precision although much much faster
func (__obf_9a34f62857fb1d1d *Stream) WriteFloat64Lossy(__obf_b924538fffe5fd64 float64) {
	if math.IsInf(__obf_b924538fffe5fd64, 0) || math.IsNaN(__obf_b924538fffe5fd64) {
		__obf_9a34f62857fb1d1d.
			Error = fmt.Errorf("unsupported value: %f", __obf_b924538fffe5fd64)
		return
	}
	if __obf_b924538fffe5fd64 < 0 {
		__obf_9a34f62857fb1d1d.__obf_f7df2a5224462d19('-')
		__obf_b924538fffe5fd64 = -__obf_b924538fffe5fd64
	}
	if __obf_b924538fffe5fd64 > 0x4ffffff {
		__obf_9a34f62857fb1d1d.
			WriteFloat64(__obf_b924538fffe5fd64)
		return
	}
	__obf_a010560b206f5c55 := 6
	__obf_2bfd195538a9cfb2 := uint64(1000000)
	__obf_736ec6eed0c0ab90 := // 6
		uint64(__obf_b924538fffe5fd64*float64(__obf_2bfd195538a9cfb2) + 0.5)
	__obf_9a34f62857fb1d1d.
		WriteUint64(__obf_736ec6eed0c0ab90 / __obf_2bfd195538a9cfb2)
	__obf_b3f63d8bcb619cfa := __obf_736ec6eed0c0ab90 % __obf_2bfd195538a9cfb2
	if __obf_b3f63d8bcb619cfa == 0 {
		return
	}
	__obf_9a34f62857fb1d1d.__obf_f7df2a5224462d19('.')
	for __obf_f9ed235be8f52eaf := __obf_a010560b206f5c55 - 1; __obf_f9ed235be8f52eaf > 0 && __obf_b3f63d8bcb619cfa < __obf_ae9fbf6f0c88dcb9[__obf_f9ed235be8f52eaf]; __obf_f9ed235be8f52eaf-- {
		__obf_9a34f62857fb1d1d.__obf_f7df2a5224462d19('0')
	}
	__obf_9a34f62857fb1d1d.
		WriteUint64(__obf_b3f63d8bcb619cfa)
	for __obf_9a34f62857fb1d1d.__obf_a065f8e0da5f5952[len(__obf_9a34f62857fb1d1d.__obf_a065f8e0da5f5952)-1] == '0' {
		__obf_9a34f62857fb1d1d.__obf_a065f8e0da5f5952 = __obf_9a34f62857fb1d1d.__obf_a065f8e0da5f5952[:len(__obf_9a34f62857fb1d1d.__obf_a065f8e0da5f5952)-1]
	}
}
