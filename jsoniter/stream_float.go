package __obf_91620b895eeff9ed

import (
	"fmt"
	"math"
	"strconv"
)

var __obf_20d949f9e88953b9 []uint64

func init() {
	__obf_20d949f9e88953b9 = []uint64{1, 10, 100, 1000, 10000, 100000, 1000000}
}

// WriteFloat32 write float32 to stream
func (__obf_850a7457bb739a32 *Stream) WriteFloat32(__obf_bbfd2ac8618a6f0c float32) {
	if math.IsInf(float64(__obf_bbfd2ac8618a6f0c), 0) || math.IsNaN(float64(__obf_bbfd2ac8618a6f0c)) {
		__obf_850a7457bb739a32.
			Error = fmt.Errorf("unsupported value: %f", __obf_bbfd2ac8618a6f0c)
		return
	}
	__obf_d8816fb3da515227 := math.Abs(float64(__obf_bbfd2ac8618a6f0c))
	fmt := byte('f')
	// Note: Must use float32 comparisons for underlying float32 value to get precise cutoffs right.
	if __obf_d8816fb3da515227 != 0 {
		if float32(__obf_d8816fb3da515227) < 1e-6 || float32(__obf_d8816fb3da515227) >= 1e21 {
			fmt = 'e'
		}
	}
	__obf_850a7457bb739a32.__obf_184433571fa55237 = strconv.AppendFloat(__obf_850a7457bb739a32.__obf_184433571fa55237, float64(__obf_bbfd2ac8618a6f0c), fmt, -1, 32)
	if fmt == 'e' {
		__obf_c80a670e818798d8 := // clean up e-09 to e-9
			len(__obf_850a7457bb739a32.__obf_184433571fa55237)
		if __obf_c80a670e818798d8 >= 4 && __obf_850a7457bb739a32.__obf_184433571fa55237[__obf_c80a670e818798d8-4] == 'e' && __obf_850a7457bb739a32.__obf_184433571fa55237[__obf_c80a670e818798d8-3] == '-' && __obf_850a7457bb739a32.__obf_184433571fa55237[__obf_c80a670e818798d8-2] == '0' {
			__obf_850a7457bb739a32.__obf_184433571fa55237[__obf_c80a670e818798d8-2] = __obf_850a7457bb739a32.__obf_184433571fa55237[__obf_c80a670e818798d8-1]
			__obf_850a7457bb739a32.__obf_184433571fa55237 = __obf_850a7457bb739a32.__obf_184433571fa55237[:__obf_c80a670e818798d8-1]
		}
	}
}

// WriteFloat32Lossy write float32 to stream with ONLY 6 digits precision although much much faster
func (__obf_850a7457bb739a32 *Stream) WriteFloat32Lossy(__obf_bbfd2ac8618a6f0c float32) {
	if math.IsInf(float64(__obf_bbfd2ac8618a6f0c), 0) || math.IsNaN(float64(__obf_bbfd2ac8618a6f0c)) {
		__obf_850a7457bb739a32.
			Error = fmt.Errorf("unsupported value: %f", __obf_bbfd2ac8618a6f0c)
		return
	}
	if __obf_bbfd2ac8618a6f0c < 0 {
		__obf_850a7457bb739a32.__obf_72837f6fe737f078('-')
		__obf_bbfd2ac8618a6f0c = -__obf_bbfd2ac8618a6f0c
	}
	if __obf_bbfd2ac8618a6f0c > 0x4ffffff {
		__obf_850a7457bb739a32.
			WriteFloat32(__obf_bbfd2ac8618a6f0c)
		return
	}
	__obf_b8935e365f3289b0 := 6
	__obf_9f81546fdfc167db := uint64(1000000)
	__obf_2c232eba39d30b33 := // 6
		uint64(float64(__obf_bbfd2ac8618a6f0c)*float64(__obf_9f81546fdfc167db) + 0.5)
	__obf_850a7457bb739a32.
		WriteUint64(__obf_2c232eba39d30b33 / __obf_9f81546fdfc167db)
	__obf_869a957bc19bc26b := __obf_2c232eba39d30b33 % __obf_9f81546fdfc167db
	if __obf_869a957bc19bc26b == 0 {
		return
	}
	__obf_850a7457bb739a32.__obf_72837f6fe737f078('.')
	for __obf_bcd7bec668625a79 := __obf_b8935e365f3289b0 - 1; __obf_bcd7bec668625a79 > 0 && __obf_869a957bc19bc26b < __obf_20d949f9e88953b9[__obf_bcd7bec668625a79]; __obf_bcd7bec668625a79-- {
		__obf_850a7457bb739a32.__obf_72837f6fe737f078('0')
	}
	__obf_850a7457bb739a32.
		WriteUint64(__obf_869a957bc19bc26b)
	for __obf_850a7457bb739a32.__obf_184433571fa55237[len(__obf_850a7457bb739a32.__obf_184433571fa55237)-1] == '0' {
		__obf_850a7457bb739a32.__obf_184433571fa55237 = __obf_850a7457bb739a32.__obf_184433571fa55237[:len(__obf_850a7457bb739a32.

			// WriteFloat64 write float64 to stream
			__obf_184433571fa55237)-1]
	}
}

func (__obf_850a7457bb739a32 *Stream) WriteFloat64(__obf_bbfd2ac8618a6f0c float64) {
	if math.IsInf(__obf_bbfd2ac8618a6f0c, 0) || math.IsNaN(__obf_bbfd2ac8618a6f0c) {
		__obf_850a7457bb739a32.
			Error = fmt.Errorf("unsupported value: %f", __obf_bbfd2ac8618a6f0c)
		return
	}
	__obf_d8816fb3da515227 := math.Abs(__obf_bbfd2ac8618a6f0c)
	fmt := byte('f')
	// Note: Must use float32 comparisons for underlying float32 value to get precise cutoffs right.
	if __obf_d8816fb3da515227 != 0 {
		if __obf_d8816fb3da515227 < 1e-6 || __obf_d8816fb3da515227 >= 1e21 {
			fmt = 'e'
		}
	}
	__obf_850a7457bb739a32.__obf_184433571fa55237 = strconv.AppendFloat(__obf_850a7457bb739a32.__obf_184433571fa55237, float64(__obf_bbfd2ac8618a6f0c), fmt, -1, 64)
	if fmt == 'e' {
		__obf_c80a670e818798d8 := // clean up e-09 to e-9
			len(__obf_850a7457bb739a32.__obf_184433571fa55237)
		if __obf_c80a670e818798d8 >= 4 && __obf_850a7457bb739a32.__obf_184433571fa55237[__obf_c80a670e818798d8-4] == 'e' && __obf_850a7457bb739a32.__obf_184433571fa55237[__obf_c80a670e818798d8-3] == '-' && __obf_850a7457bb739a32.__obf_184433571fa55237[__obf_c80a670e818798d8-2] == '0' {
			__obf_850a7457bb739a32.__obf_184433571fa55237[__obf_c80a670e818798d8-2] = __obf_850a7457bb739a32.__obf_184433571fa55237[__obf_c80a670e818798d8-1]
			__obf_850a7457bb739a32.__obf_184433571fa55237 = __obf_850a7457bb739a32.__obf_184433571fa55237[:__obf_c80a670e818798d8-1]
		}
	}
}

// WriteFloat64Lossy write float64 to stream with ONLY 6 digits precision although much much faster
func (__obf_850a7457bb739a32 *Stream) WriteFloat64Lossy(__obf_bbfd2ac8618a6f0c float64) {
	if math.IsInf(__obf_bbfd2ac8618a6f0c, 0) || math.IsNaN(__obf_bbfd2ac8618a6f0c) {
		__obf_850a7457bb739a32.
			Error = fmt.Errorf("unsupported value: %f", __obf_bbfd2ac8618a6f0c)
		return
	}
	if __obf_bbfd2ac8618a6f0c < 0 {
		__obf_850a7457bb739a32.__obf_72837f6fe737f078('-')
		__obf_bbfd2ac8618a6f0c = -__obf_bbfd2ac8618a6f0c
	}
	if __obf_bbfd2ac8618a6f0c > 0x4ffffff {
		__obf_850a7457bb739a32.
			WriteFloat64(__obf_bbfd2ac8618a6f0c)
		return
	}
	__obf_b8935e365f3289b0 := 6
	__obf_9f81546fdfc167db := uint64(1000000)
	__obf_2c232eba39d30b33 := // 6
		uint64(__obf_bbfd2ac8618a6f0c*float64(__obf_9f81546fdfc167db) + 0.5)
	__obf_850a7457bb739a32.
		WriteUint64(__obf_2c232eba39d30b33 / __obf_9f81546fdfc167db)
	__obf_869a957bc19bc26b := __obf_2c232eba39d30b33 % __obf_9f81546fdfc167db
	if __obf_869a957bc19bc26b == 0 {
		return
	}
	__obf_850a7457bb739a32.__obf_72837f6fe737f078('.')
	for __obf_bcd7bec668625a79 := __obf_b8935e365f3289b0 - 1; __obf_bcd7bec668625a79 > 0 && __obf_869a957bc19bc26b < __obf_20d949f9e88953b9[__obf_bcd7bec668625a79]; __obf_bcd7bec668625a79-- {
		__obf_850a7457bb739a32.__obf_72837f6fe737f078('0')
	}
	__obf_850a7457bb739a32.
		WriteUint64(__obf_869a957bc19bc26b)
	for __obf_850a7457bb739a32.__obf_184433571fa55237[len(__obf_850a7457bb739a32.__obf_184433571fa55237)-1] == '0' {
		__obf_850a7457bb739a32.__obf_184433571fa55237 = __obf_850a7457bb739a32.__obf_184433571fa55237[:len(__obf_850a7457bb739a32.__obf_184433571fa55237)-1]
	}
}
