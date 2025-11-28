package __obf_c357e2bf526d00f9

import (
	"sort"
	"time"
)

// float64ToInt rounds a float64 to an int
func __obf_5a5f26f165c78990(__obf_60f27e022365452d float64) (__obf_ca93fe7a73dff08f int) {
	__obf_1cecb5e04ab88618, _ := Round(__obf_60f27e022365452d, 0)
	return int(__obf_1cecb5e04ab88618)
}

// unixnano returns nanoseconds from UTC epoch
func __obf_956788a42e61fe78() int64 {
	return time.Now().UTC().UnixNano()
}

// copyslice copies a slice of float64s
func __obf_1f875fa187fe4301(__obf_60f27e022365452d []float64) []float64 {
	__obf_d4e2b66ebe9a2dbf := make([]float64, len(__obf_60f27e022365452d))
	copy(__obf_d4e2b66ebe9a2dbf, __obf_60f27e022365452d)
	return __obf_d4e2b66ebe9a2dbf
}

// sortedCopy returns a sorted copy of float64s
func __obf_a9d139b4fa16b92f(__obf_60f27e022365452d []float64) (copy []float64) {
	copy = __obf_1f875fa187fe4301(__obf_60f27e022365452d)
	sort.Float64s(copy)
	return
}

// sortedCopyDif returns a sorted copy of float64s
// only if the original data isn't sorted.
// Only use this if returned slice won't be manipulated!
func __obf_397c2f24121e735c(__obf_60f27e022365452d []float64) (copy []float64) {
	if sort.Float64sAreSorted(__obf_60f27e022365452d) {
		return __obf_60f27e022365452d
	}
	copy = __obf_1f875fa187fe4301(__obf_60f27e022365452d)
	sort.Float64s(copy)
	return
}
