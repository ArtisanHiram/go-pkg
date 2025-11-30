package __obf_5865d4dc61043b11

import (
	"sort"
	"time"
)

// float64ToInt rounds a float64 to an int
func __obf_f602652101083886(__obf_b0bc5503bc2661cc float64) (__obf_37ef2768c239ad7c int) {
	__obf_bc1c068e07610bb4, _ := Round(__obf_b0bc5503bc2661cc, 0)
	return int(__obf_bc1c068e07610bb4)
}

// unixnano returns nanoseconds from UTC epoch
func __obf_d77f0e285f2cf6a1() int64 {
	return time.Now().UTC().UnixNano()
}

// copyslice copies a slice of float64s
func __obf_e27e88b10ae29364(__obf_b0bc5503bc2661cc []float64) []float64 {
	__obf_02161123be315268 := make([]float64, len(__obf_b0bc5503bc2661cc))
	copy(__obf_02161123be315268, __obf_b0bc5503bc2661cc)
	return __obf_02161123be315268
}

// sortedCopy returns a sorted copy of float64s
func __obf_713a257c97207840(__obf_b0bc5503bc2661cc []float64) (copy []float64) {
	copy = __obf_e27e88b10ae29364(__obf_b0bc5503bc2661cc)
	sort.Float64s(copy)
	return
}

// sortedCopyDif returns a sorted copy of float64s
// only if the original data isn't sorted.
// Only use this if returned slice won't be manipulated!
func __obf_87d0671c67fc5a81(__obf_b0bc5503bc2661cc []float64) (copy []float64) {
	if sort.Float64sAreSorted(__obf_b0bc5503bc2661cc) {
		return __obf_b0bc5503bc2661cc
	}
	copy = __obf_e27e88b10ae29364(__obf_b0bc5503bc2661cc)
	sort.Float64s(copy)
	return
}
