package __obf_13f6e310b0abe7e3

import (
	"sort"
	"time"
)

// float64ToInt rounds a float64 to an int
func __obf_a1c4c9e1b18391fa(__obf_f8ab7230e9a1c193 float64) (__obf_d799f7db20aaee16 int) {
	__obf_cc8569f0e2a39778, _ := Round(__obf_f8ab7230e9a1c193, 0)
	return int(__obf_cc8569f0e2a39778)
}

// unixnano returns nanoseconds from UTC epoch
func __obf_711a707430afdb07() int64 {
	return time.Now().UTC().UnixNano()
}

// copyslice copies a slice of float64s
func __obf_d26975d4f02ddb59(__obf_f8ab7230e9a1c193 []float64) []float64 {
	__obf_516af7fe712d1b0f := make([]float64, len(__obf_f8ab7230e9a1c193))
	copy(__obf_516af7fe712d1b0f, __obf_f8ab7230e9a1c193)
	return __obf_516af7fe712d1b0f
}

// sortedCopy returns a sorted copy of float64s
func __obf_1591f280eef6f330(__obf_f8ab7230e9a1c193 []float64) (copy []float64) {
	copy = __obf_d26975d4f02ddb59(__obf_f8ab7230e9a1c193)
	sort.Float64s(copy)
	return
}

// sortedCopyDif returns a sorted copy of float64s
// only if the original data isn't sorted.
// Only use this if returned slice won't be manipulated!
func __obf_53c5001ec5502100(__obf_f8ab7230e9a1c193 []float64) (copy []float64) {
	if sort.Float64sAreSorted(__obf_f8ab7230e9a1c193) {
		return __obf_f8ab7230e9a1c193
	}
	copy = __obf_d26975d4f02ddb59(__obf_f8ab7230e9a1c193)
	sort.Float64s(copy)
	return
}
