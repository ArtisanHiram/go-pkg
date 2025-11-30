package __obf_bae661114bf2a601

import (
	"sort"
	"time"
)

// float64ToInt rounds a float64 to an int
func __obf_91d2ecb5786b6fc2(__obf_c5a1fa438321e828 float64) (__obf_935ae667db647873 int) {
	__obf_a2088af72211e851, _ := Round(__obf_c5a1fa438321e828, 0)
	return int(__obf_a2088af72211e851)
}

// unixnano returns nanoseconds from UTC epoch
func __obf_c0925987c8c2abf1() int64 {
	return time.Now().UTC().UnixNano()
}

// copyslice copies a slice of float64s
func __obf_ea1bc02dcef2d2da(__obf_c5a1fa438321e828 []float64) []float64 {
	__obf_52c56453a486a111 := make([]float64, len(__obf_c5a1fa438321e828))
	copy(__obf_52c56453a486a111, __obf_c5a1fa438321e828)
	return __obf_52c56453a486a111
}

// sortedCopy returns a sorted copy of float64s
func __obf_083545a4cd904204(__obf_c5a1fa438321e828 []float64) (copy []float64) {
	copy = __obf_ea1bc02dcef2d2da(__obf_c5a1fa438321e828)
	sort.Float64s(copy)
	return
}

// sortedCopyDif returns a sorted copy of float64s
// only if the original data isn't sorted.
// Only use this if returned slice won't be manipulated!
func __obf_e0b2761114041739(__obf_c5a1fa438321e828 []float64) (copy []float64) {
	if sort.Float64sAreSorted(__obf_c5a1fa438321e828) {
		return __obf_c5a1fa438321e828
	}
	copy = __obf_ea1bc02dcef2d2da(__obf_c5a1fa438321e828)
	sort.Float64s(copy)
	return
}
