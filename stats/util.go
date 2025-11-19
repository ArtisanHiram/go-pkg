package __obf_ce59006f265ba976

import (
	"sort"
	"time"
)

// float64ToInt rounds a float64 to an int
func __obf_d98fcee495fc36b5(__obf_5884ff1bb2ed3b04 float64) (__obf_92afb9d59beb5da0 int) {
	__obf_29df04175ffb0d27, _ := Round(__obf_5884ff1bb2ed3b04, 0)
	return int(__obf_29df04175ffb0d27)
}

// unixnano returns nanoseconds from UTC epoch
func __obf_a9b86570fe3f0149() int64 {
	return time.Now().UTC().UnixNano()
}

// copyslice copies a slice of float64s
func __obf_673895855bd88e8e(__obf_5884ff1bb2ed3b04 []float64) []float64 {
	__obf_65d938f48d53f189 := make([]float64, len(__obf_5884ff1bb2ed3b04))
	copy(__obf_65d938f48d53f189, __obf_5884ff1bb2ed3b04)
	return __obf_65d938f48d53f189
}

// sortedCopy returns a sorted copy of float64s
func __obf_cd4216a3bd4147b5(__obf_5884ff1bb2ed3b04 []float64) (copy []float64) {
	copy = __obf_673895855bd88e8e(__obf_5884ff1bb2ed3b04)
	sort.Float64s(copy)
	return
}

// sortedCopyDif returns a sorted copy of float64s
// only if the original data isn't sorted.
// Only use this if returned slice won't be manipulated!
func __obf_a0c8d74396766b11(__obf_5884ff1bb2ed3b04 []float64) (copy []float64) {
	if sort.Float64sAreSorted(__obf_5884ff1bb2ed3b04) {
		return __obf_5884ff1bb2ed3b04
	}
	copy = __obf_673895855bd88e8e(__obf_5884ff1bb2ed3b04)
	sort.Float64s(copy)
	return
}
