package __obf_aa50dbb4389f0ad9

import (
	"sort"
	"time"
)

// float64ToInt rounds a float64 to an int
func __obf_dfb49e042673d42e(__obf_79ecf1e2364bfac6 float64) (__obf_9f974e6b7e5da208 int) {
	__obf_c9cb4eafe52817bb, _ := Round(__obf_79ecf1e2364bfac6, 0)
	return int(__obf_c9cb4eafe52817bb)
}

// unixnano returns nanoseconds from UTC epoch
func __obf_4d7d278fa8102ad8() int64 {
	return time.Now().UTC().UnixNano()
}

// copyslice copies a slice of float64s
func __obf_a8219eda17838ad2(__obf_79ecf1e2364bfac6 []float64) []float64 {
	__obf_d4b1989cb8216e9b := make([]float64, len(__obf_79ecf1e2364bfac6))
	copy(__obf_d4b1989cb8216e9b, __obf_79ecf1e2364bfac6)
	return __obf_d4b1989cb8216e9b
}

// sortedCopy returns a sorted copy of float64s
func __obf_78a9e1790a5bd3a2(__obf_79ecf1e2364bfac6 []float64) (copy []float64) {
	copy = __obf_a8219eda17838ad2(__obf_79ecf1e2364bfac6)
	sort.Float64s(copy)
	return
}

// sortedCopyDif returns a sorted copy of float64s
// only if the original data isn't sorted.
// Only use this if returned slice won't be manipulated!
func __obf_b39898782f16d8ff(__obf_79ecf1e2364bfac6 []float64) (copy []float64) {
	if sort.Float64sAreSorted(__obf_79ecf1e2364bfac6) {
		return __obf_79ecf1e2364bfac6
	}
	copy = __obf_a8219eda17838ad2(__obf_79ecf1e2364bfac6)
	sort.Float64s(copy)
	return
}
