package __obf_5b802ce8d9ba56d6

import (
	"fmt"
	"math"
	"strconv"
)

var __obf_532c8e692981b5b4 []uint64

func init() {
	__obf_532c8e692981b5b4 = []uint64{1, 10, 100, 1000, 10000, 100000, 1000000}
}

// WriteFloat32 write float32 to stream
func (__obf_00fc491caa967a74 *Stream) WriteFloat32(__obf_5406b11e33b9d1d3 float32) {
	if math.IsInf(float64(__obf_5406b11e33b9d1d3), 0) || math.IsNaN(float64(__obf_5406b11e33b9d1d3)) {
		__obf_00fc491caa967a74.
			Error = fmt.Errorf("unsupported value: %f", __obf_5406b11e33b9d1d3)
		return
	}
	__obf_643bebdc5d9b3291 := math.Abs(float64(__obf_5406b11e33b9d1d3))
	fmt := byte('f')
	// Note: Must use float32 comparisons for underlying float32 value to get precise cutoffs right.
	if __obf_643bebdc5d9b3291 != 0 {
		if float32(__obf_643bebdc5d9b3291) < 1e-6 || float32(__obf_643bebdc5d9b3291) >= 1e21 {
			fmt = 'e'
		}
	}
	__obf_00fc491caa967a74.__obf_9fc06d9180f0daca = strconv.AppendFloat(__obf_00fc491caa967a74.__obf_9fc06d9180f0daca, float64(__obf_5406b11e33b9d1d3), fmt, -1, 32)
	if fmt == 'e' {
		__obf_2c0ce853cb81f014 := // clean up e-09 to e-9
			len(__obf_00fc491caa967a74.__obf_9fc06d9180f0daca)
		if __obf_2c0ce853cb81f014 >= 4 && __obf_00fc491caa967a74.__obf_9fc06d9180f0daca[__obf_2c0ce853cb81f014-4] == 'e' && __obf_00fc491caa967a74.__obf_9fc06d9180f0daca[__obf_2c0ce853cb81f014-3] == '-' && __obf_00fc491caa967a74.__obf_9fc06d9180f0daca[__obf_2c0ce853cb81f014-2] == '0' {
			__obf_00fc491caa967a74.__obf_9fc06d9180f0daca[__obf_2c0ce853cb81f014-2] = __obf_00fc491caa967a74.__obf_9fc06d9180f0daca[__obf_2c0ce853cb81f014-1]
			__obf_00fc491caa967a74.__obf_9fc06d9180f0daca = __obf_00fc491caa967a74.__obf_9fc06d9180f0daca[:__obf_2c0ce853cb81f014-1]
		}
	}
}

// WriteFloat32Lossy write float32 to stream with ONLY 6 digits precision although much much faster
func (__obf_00fc491caa967a74 *Stream) WriteFloat32Lossy(__obf_5406b11e33b9d1d3 float32) {
	if math.IsInf(float64(__obf_5406b11e33b9d1d3), 0) || math.IsNaN(float64(__obf_5406b11e33b9d1d3)) {
		__obf_00fc491caa967a74.
			Error = fmt.Errorf("unsupported value: %f", __obf_5406b11e33b9d1d3)
		return
	}
	if __obf_5406b11e33b9d1d3 < 0 {
		__obf_00fc491caa967a74.__obf_487ee8035c7dd8f6('-')
		__obf_5406b11e33b9d1d3 = -__obf_5406b11e33b9d1d3
	}
	if __obf_5406b11e33b9d1d3 > 0x4ffffff {
		__obf_00fc491caa967a74.
			WriteFloat32(__obf_5406b11e33b9d1d3)
		return
	}
	__obf_541c7424f0a0f929 := 6
	__obf_66795d43132caf14 := uint64(1000000)
	__obf_0e098d3331337130 := // 6
		uint64(float64(__obf_5406b11e33b9d1d3)*float64(__obf_66795d43132caf14) + 0.5)
	__obf_00fc491caa967a74.
		WriteUint64(__obf_0e098d3331337130 / __obf_66795d43132caf14)
	__obf_07aee296fe4d5abb := __obf_0e098d3331337130 % __obf_66795d43132caf14
	if __obf_07aee296fe4d5abb == 0 {
		return
	}
	__obf_00fc491caa967a74.__obf_487ee8035c7dd8f6('.')
	for __obf_82ddba9040ff8d8c := __obf_541c7424f0a0f929 - 1; __obf_82ddba9040ff8d8c > 0 && __obf_07aee296fe4d5abb < __obf_532c8e692981b5b4[__obf_82ddba9040ff8d8c]; __obf_82ddba9040ff8d8c-- {
		__obf_00fc491caa967a74.__obf_487ee8035c7dd8f6('0')
	}
	__obf_00fc491caa967a74.
		WriteUint64(__obf_07aee296fe4d5abb)
	for __obf_00fc491caa967a74.__obf_9fc06d9180f0daca[len(__obf_00fc491caa967a74.__obf_9fc06d9180f0daca)-1] == '0' {
		__obf_00fc491caa967a74.__obf_9fc06d9180f0daca = __obf_00fc491caa967a74.__obf_9fc06d9180f0daca[:len(__obf_00fc491caa967a74.

			// WriteFloat64 write float64 to stream
			__obf_9fc06d9180f0daca)-1]
	}
}

func (__obf_00fc491caa967a74 *Stream) WriteFloat64(__obf_5406b11e33b9d1d3 float64) {
	if math.IsInf(__obf_5406b11e33b9d1d3, 0) || math.IsNaN(__obf_5406b11e33b9d1d3) {
		__obf_00fc491caa967a74.
			Error = fmt.Errorf("unsupported value: %f", __obf_5406b11e33b9d1d3)
		return
	}
	__obf_643bebdc5d9b3291 := math.Abs(__obf_5406b11e33b9d1d3)
	fmt := byte('f')
	// Note: Must use float32 comparisons for underlying float32 value to get precise cutoffs right.
	if __obf_643bebdc5d9b3291 != 0 {
		if __obf_643bebdc5d9b3291 < 1e-6 || __obf_643bebdc5d9b3291 >= 1e21 {
			fmt = 'e'
		}
	}
	__obf_00fc491caa967a74.__obf_9fc06d9180f0daca = strconv.AppendFloat(__obf_00fc491caa967a74.__obf_9fc06d9180f0daca, float64(__obf_5406b11e33b9d1d3), fmt, -1, 64)
	if fmt == 'e' {
		__obf_2c0ce853cb81f014 := // clean up e-09 to e-9
			len(__obf_00fc491caa967a74.__obf_9fc06d9180f0daca)
		if __obf_2c0ce853cb81f014 >= 4 && __obf_00fc491caa967a74.__obf_9fc06d9180f0daca[__obf_2c0ce853cb81f014-4] == 'e' && __obf_00fc491caa967a74.__obf_9fc06d9180f0daca[__obf_2c0ce853cb81f014-3] == '-' && __obf_00fc491caa967a74.__obf_9fc06d9180f0daca[__obf_2c0ce853cb81f014-2] == '0' {
			__obf_00fc491caa967a74.__obf_9fc06d9180f0daca[__obf_2c0ce853cb81f014-2] = __obf_00fc491caa967a74.__obf_9fc06d9180f0daca[__obf_2c0ce853cb81f014-1]
			__obf_00fc491caa967a74.__obf_9fc06d9180f0daca = __obf_00fc491caa967a74.__obf_9fc06d9180f0daca[:__obf_2c0ce853cb81f014-1]
		}
	}
}

// WriteFloat64Lossy write float64 to stream with ONLY 6 digits precision although much much faster
func (__obf_00fc491caa967a74 *Stream) WriteFloat64Lossy(__obf_5406b11e33b9d1d3 float64) {
	if math.IsInf(__obf_5406b11e33b9d1d3, 0) || math.IsNaN(__obf_5406b11e33b9d1d3) {
		__obf_00fc491caa967a74.
			Error = fmt.Errorf("unsupported value: %f", __obf_5406b11e33b9d1d3)
		return
	}
	if __obf_5406b11e33b9d1d3 < 0 {
		__obf_00fc491caa967a74.__obf_487ee8035c7dd8f6('-')
		__obf_5406b11e33b9d1d3 = -__obf_5406b11e33b9d1d3
	}
	if __obf_5406b11e33b9d1d3 > 0x4ffffff {
		__obf_00fc491caa967a74.
			WriteFloat64(__obf_5406b11e33b9d1d3)
		return
	}
	__obf_541c7424f0a0f929 := 6
	__obf_66795d43132caf14 := uint64(1000000)
	__obf_0e098d3331337130 := // 6
		uint64(__obf_5406b11e33b9d1d3*float64(__obf_66795d43132caf14) + 0.5)
	__obf_00fc491caa967a74.
		WriteUint64(__obf_0e098d3331337130 / __obf_66795d43132caf14)
	__obf_07aee296fe4d5abb := __obf_0e098d3331337130 % __obf_66795d43132caf14
	if __obf_07aee296fe4d5abb == 0 {
		return
	}
	__obf_00fc491caa967a74.__obf_487ee8035c7dd8f6('.')
	for __obf_82ddba9040ff8d8c := __obf_541c7424f0a0f929 - 1; __obf_82ddba9040ff8d8c > 0 && __obf_07aee296fe4d5abb < __obf_532c8e692981b5b4[__obf_82ddba9040ff8d8c]; __obf_82ddba9040ff8d8c-- {
		__obf_00fc491caa967a74.__obf_487ee8035c7dd8f6('0')
	}
	__obf_00fc491caa967a74.
		WriteUint64(__obf_07aee296fe4d5abb)
	for __obf_00fc491caa967a74.__obf_9fc06d9180f0daca[len(__obf_00fc491caa967a74.__obf_9fc06d9180f0daca)-1] == '0' {
		__obf_00fc491caa967a74.__obf_9fc06d9180f0daca = __obf_00fc491caa967a74.__obf_9fc06d9180f0daca[:len(__obf_00fc491caa967a74.__obf_9fc06d9180f0daca)-1]
	}
}
