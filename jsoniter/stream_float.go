package __obf_c7b79b12b33d8238

import (
	"fmt"
	"math"
	"strconv"
)

var __obf_e075611d0d831082 []uint64

func init() {
	__obf_e075611d0d831082 = []uint64{1, 10, 100, 1000, 10000, 100000, 1000000}
}

// WriteFloat32 write float32 to stream
func (__obf_d8f50bcfe87d8b47 *Stream) WriteFloat32(__obf_35accd096612ebe4 float32) {
	if math.IsInf(float64(__obf_35accd096612ebe4), 0) || math.IsNaN(float64(__obf_35accd096612ebe4)) {
		__obf_d8f50bcfe87d8b47.
			Error = fmt.Errorf("unsupported value: %f", __obf_35accd096612ebe4)
		return
	}
	__obf_b8ed30e32646b77a := math.Abs(float64(__obf_35accd096612ebe4))
	fmt := byte('f')
	// Note: Must use float32 comparisons for underlying float32 value to get precise cutoffs right.
	if __obf_b8ed30e32646b77a != 0 {
		if float32(__obf_b8ed30e32646b77a) < 1e-6 || float32(__obf_b8ed30e32646b77a) >= 1e21 {
			fmt = 'e'
		}
	}
	__obf_d8f50bcfe87d8b47.__obf_684d738bc3ac851a = strconv.AppendFloat(__obf_d8f50bcfe87d8b47.__obf_684d738bc3ac851a, float64(__obf_35accd096612ebe4), fmt, -1, 32)
	if fmt == 'e' {
		__obf_ff4330e73e137d5c := // clean up e-09 to e-9
			len(__obf_d8f50bcfe87d8b47.__obf_684d738bc3ac851a)
		if __obf_ff4330e73e137d5c >= 4 && __obf_d8f50bcfe87d8b47.__obf_684d738bc3ac851a[__obf_ff4330e73e137d5c-4] == 'e' && __obf_d8f50bcfe87d8b47.__obf_684d738bc3ac851a[__obf_ff4330e73e137d5c-3] == '-' && __obf_d8f50bcfe87d8b47.__obf_684d738bc3ac851a[__obf_ff4330e73e137d5c-2] == '0' {
			__obf_d8f50bcfe87d8b47.__obf_684d738bc3ac851a[__obf_ff4330e73e137d5c-2] = __obf_d8f50bcfe87d8b47.__obf_684d738bc3ac851a[__obf_ff4330e73e137d5c-1]
			__obf_d8f50bcfe87d8b47.__obf_684d738bc3ac851a = __obf_d8f50bcfe87d8b47.__obf_684d738bc3ac851a[:__obf_ff4330e73e137d5c-1]
		}
	}
}

// WriteFloat32Lossy write float32 to stream with ONLY 6 digits precision although much much faster
func (__obf_d8f50bcfe87d8b47 *Stream) WriteFloat32Lossy(__obf_35accd096612ebe4 float32) {
	if math.IsInf(float64(__obf_35accd096612ebe4), 0) || math.IsNaN(float64(__obf_35accd096612ebe4)) {
		__obf_d8f50bcfe87d8b47.
			Error = fmt.Errorf("unsupported value: %f", __obf_35accd096612ebe4)
		return
	}
	if __obf_35accd096612ebe4 < 0 {
		__obf_d8f50bcfe87d8b47.__obf_7340077d41996822('-')
		__obf_35accd096612ebe4 = -__obf_35accd096612ebe4
	}
	if __obf_35accd096612ebe4 > 0x4ffffff {
		__obf_d8f50bcfe87d8b47.
			WriteFloat32(__obf_35accd096612ebe4)
		return
	}
	__obf_f2bc39318e336a39 := 6
	__obf_65ebcaf0bb51cc3c := uint64(1000000)
	__obf_f59585f7b6dfba37 := // 6
		uint64(float64(__obf_35accd096612ebe4)*float64(__obf_65ebcaf0bb51cc3c) + 0.5)
	__obf_d8f50bcfe87d8b47.
		WriteUint64(__obf_f59585f7b6dfba37 / __obf_65ebcaf0bb51cc3c)
	__obf_6a9b1eee4897c599 := __obf_f59585f7b6dfba37 % __obf_65ebcaf0bb51cc3c
	if __obf_6a9b1eee4897c599 == 0 {
		return
	}
	__obf_d8f50bcfe87d8b47.__obf_7340077d41996822('.')
	for __obf_c03fbc322c1a9109 := __obf_f2bc39318e336a39 - 1; __obf_c03fbc322c1a9109 > 0 && __obf_6a9b1eee4897c599 < __obf_e075611d0d831082[__obf_c03fbc322c1a9109]; __obf_c03fbc322c1a9109-- {
		__obf_d8f50bcfe87d8b47.__obf_7340077d41996822('0')
	}
	__obf_d8f50bcfe87d8b47.
		WriteUint64(__obf_6a9b1eee4897c599)
	for __obf_d8f50bcfe87d8b47.__obf_684d738bc3ac851a[len(__obf_d8f50bcfe87d8b47.__obf_684d738bc3ac851a)-1] == '0' {
		__obf_d8f50bcfe87d8b47.__obf_684d738bc3ac851a = __obf_d8f50bcfe87d8b47.__obf_684d738bc3ac851a[:len(__obf_d8f50bcfe87d8b47.

			// WriteFloat64 write float64 to stream
			__obf_684d738bc3ac851a)-1]
	}
}

func (__obf_d8f50bcfe87d8b47 *Stream) WriteFloat64(__obf_35accd096612ebe4 float64) {
	if math.IsInf(__obf_35accd096612ebe4, 0) || math.IsNaN(__obf_35accd096612ebe4) {
		__obf_d8f50bcfe87d8b47.
			Error = fmt.Errorf("unsupported value: %f", __obf_35accd096612ebe4)
		return
	}
	__obf_b8ed30e32646b77a := math.Abs(__obf_35accd096612ebe4)
	fmt := byte('f')
	// Note: Must use float32 comparisons for underlying float32 value to get precise cutoffs right.
	if __obf_b8ed30e32646b77a != 0 {
		if __obf_b8ed30e32646b77a < 1e-6 || __obf_b8ed30e32646b77a >= 1e21 {
			fmt = 'e'
		}
	}
	__obf_d8f50bcfe87d8b47.__obf_684d738bc3ac851a = strconv.AppendFloat(__obf_d8f50bcfe87d8b47.__obf_684d738bc3ac851a, float64(__obf_35accd096612ebe4), fmt, -1, 64)
	if fmt == 'e' {
		__obf_ff4330e73e137d5c := // clean up e-09 to e-9
			len(__obf_d8f50bcfe87d8b47.__obf_684d738bc3ac851a)
		if __obf_ff4330e73e137d5c >= 4 && __obf_d8f50bcfe87d8b47.__obf_684d738bc3ac851a[__obf_ff4330e73e137d5c-4] == 'e' && __obf_d8f50bcfe87d8b47.__obf_684d738bc3ac851a[__obf_ff4330e73e137d5c-3] == '-' && __obf_d8f50bcfe87d8b47.__obf_684d738bc3ac851a[__obf_ff4330e73e137d5c-2] == '0' {
			__obf_d8f50bcfe87d8b47.__obf_684d738bc3ac851a[__obf_ff4330e73e137d5c-2] = __obf_d8f50bcfe87d8b47.__obf_684d738bc3ac851a[__obf_ff4330e73e137d5c-1]
			__obf_d8f50bcfe87d8b47.__obf_684d738bc3ac851a = __obf_d8f50bcfe87d8b47.__obf_684d738bc3ac851a[:__obf_ff4330e73e137d5c-1]
		}
	}
}

// WriteFloat64Lossy write float64 to stream with ONLY 6 digits precision although much much faster
func (__obf_d8f50bcfe87d8b47 *Stream) WriteFloat64Lossy(__obf_35accd096612ebe4 float64) {
	if math.IsInf(__obf_35accd096612ebe4, 0) || math.IsNaN(__obf_35accd096612ebe4) {
		__obf_d8f50bcfe87d8b47.
			Error = fmt.Errorf("unsupported value: %f", __obf_35accd096612ebe4)
		return
	}
	if __obf_35accd096612ebe4 < 0 {
		__obf_d8f50bcfe87d8b47.__obf_7340077d41996822('-')
		__obf_35accd096612ebe4 = -__obf_35accd096612ebe4
	}
	if __obf_35accd096612ebe4 > 0x4ffffff {
		__obf_d8f50bcfe87d8b47.
			WriteFloat64(__obf_35accd096612ebe4)
		return
	}
	__obf_f2bc39318e336a39 := 6
	__obf_65ebcaf0bb51cc3c := uint64(1000000)
	__obf_f59585f7b6dfba37 := // 6
		uint64(__obf_35accd096612ebe4*float64(__obf_65ebcaf0bb51cc3c) + 0.5)
	__obf_d8f50bcfe87d8b47.
		WriteUint64(__obf_f59585f7b6dfba37 / __obf_65ebcaf0bb51cc3c)
	__obf_6a9b1eee4897c599 := __obf_f59585f7b6dfba37 % __obf_65ebcaf0bb51cc3c
	if __obf_6a9b1eee4897c599 == 0 {
		return
	}
	__obf_d8f50bcfe87d8b47.__obf_7340077d41996822('.')
	for __obf_c03fbc322c1a9109 := __obf_f2bc39318e336a39 - 1; __obf_c03fbc322c1a9109 > 0 && __obf_6a9b1eee4897c599 < __obf_e075611d0d831082[__obf_c03fbc322c1a9109]; __obf_c03fbc322c1a9109-- {
		__obf_d8f50bcfe87d8b47.__obf_7340077d41996822('0')
	}
	__obf_d8f50bcfe87d8b47.
		WriteUint64(__obf_6a9b1eee4897c599)
	for __obf_d8f50bcfe87d8b47.__obf_684d738bc3ac851a[len(__obf_d8f50bcfe87d8b47.__obf_684d738bc3ac851a)-1] == '0' {
		__obf_d8f50bcfe87d8b47.__obf_684d738bc3ac851a = __obf_d8f50bcfe87d8b47.__obf_684d738bc3ac851a[:len(__obf_d8f50bcfe87d8b47.__obf_684d738bc3ac851a)-1]
	}
}
