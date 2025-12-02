package __obf_e7d13a84deed4934

import (
	"sort"
	"time"
)

// float64ToInt rounds a float64 to an int
func __obf_c967c3a35008f1c8(__obf_31fb5f86ef39e603 float64) (__obf_d06e704e2dbd8d56 int) {
	__obf_2ff15920dfc1b8de, _ := Round(__obf_31fb5f86ef39e603, 0)
	return int(__obf_2ff15920dfc1b8de)
}

// unixnano returns nanoseconds from UTC epoch
func __obf_6f7daec3121cb359() int64 {
	return time.Now().UTC().UnixNano()
}

// copyslice copies a slice of float64s
func __obf_36200b12b4bfb8d7(__obf_31fb5f86ef39e603 []float64) []float64 {
	__obf_1e3975c9e116ae8c := make([]float64, len(__obf_31fb5f86ef39e603))
	copy(__obf_1e3975c9e116ae8c, __obf_31fb5f86ef39e603)
	return __obf_1e3975c9e116ae8c
}

// sortedCopy returns a sorted copy of float64s
func __obf_a9fa2c6c1f72c723(__obf_31fb5f86ef39e603 []float64) (copy []float64) {
	copy = __obf_36200b12b4bfb8d7(__obf_31fb5f86ef39e603)
	sort.Float64s(copy)
	return
}

// sortedCopyDif returns a sorted copy of float64s
// only if the original data isn't sorted.
// Only use this if returned slice won't be manipulated!
func __obf_73c3e1f4104fe140(__obf_31fb5f86ef39e603 []float64) (copy []float64) {
	if sort.Float64sAreSorted(__obf_31fb5f86ef39e603) {
		return __obf_31fb5f86ef39e603
	}
	copy = __obf_36200b12b4bfb8d7(__obf_31fb5f86ef39e603)
	sort.Float64s(copy)
	return
}
