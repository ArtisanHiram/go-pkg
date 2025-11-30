package __obf_face5ef149f4f55f

import (
	"sort"
	"time"
)

// float64ToInt rounds a float64 to an int
func __obf_fa1f409c06ca1f1e(__obf_0f75e398d64caf9c float64) (__obf_c1444a6d5e40fa93 int) {
	__obf_8504b33b3faa0149, _ := Round(__obf_0f75e398d64caf9c, 0)
	return int(__obf_8504b33b3faa0149)
}

// unixnano returns nanoseconds from UTC epoch
func __obf_e3b91100e6757d64() int64 {
	return time.Now().UTC().UnixNano()
}

// copyslice copies a slice of float64s
func __obf_2ca28dec38c241a5(__obf_0f75e398d64caf9c []float64) []float64 {
	__obf_73a02b4fdf5ba47a := make([]float64, len(__obf_0f75e398d64caf9c))
	copy(__obf_73a02b4fdf5ba47a, __obf_0f75e398d64caf9c)
	return __obf_73a02b4fdf5ba47a
}

// sortedCopy returns a sorted copy of float64s
func __obf_c61afc841612b113(__obf_0f75e398d64caf9c []float64) (copy []float64) {
	copy = __obf_2ca28dec38c241a5(__obf_0f75e398d64caf9c)
	sort.Float64s(copy)
	return
}

// sortedCopyDif returns a sorted copy of float64s
// only if the original data isn't sorted.
// Only use this if returned slice won't be manipulated!
func __obf_1693538c450ac771(__obf_0f75e398d64caf9c []float64) (copy []float64) {
	if sort.Float64sAreSorted(__obf_0f75e398d64caf9c) {
		return __obf_0f75e398d64caf9c
	}
	copy = __obf_2ca28dec38c241a5(__obf_0f75e398d64caf9c)
	sort.Float64s(copy)
	return
}
