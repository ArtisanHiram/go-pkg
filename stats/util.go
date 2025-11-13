package __obf_9d1e97a8e05fedb6

import (
	"sort"
	"time"
)

// float64ToInt rounds a float64 to an int
func __obf_27e7cc340c52f451(__obf_d4d18bbab73e7cfe float64) (__obf_c4dca37eb66af0ff int) {
	__obf_7e34ae3e1e8f3016, _ := Round(__obf_d4d18bbab73e7cfe, 0)
	return int(__obf_7e34ae3e1e8f3016)
}

// unixnano returns nanoseconds from UTC epoch
func __obf_f3414ea72dde3f60() int64 {
	return time.Now().UTC().UnixNano()
}

// copyslice copies a slice of float64s
func __obf_166de59bf3ffd29b(__obf_d4d18bbab73e7cfe []float64) []float64 {
	__obf_56bdc23e16a7cc2a := make([]float64, len(__obf_d4d18bbab73e7cfe))
	copy(__obf_56bdc23e16a7cc2a, __obf_d4d18bbab73e7cfe)
	return __obf_56bdc23e16a7cc2a
}

// sortedCopy returns a sorted copy of float64s
func __obf_4ab795f70adc8177(__obf_d4d18bbab73e7cfe []float64) (copy []float64) {
	copy = __obf_166de59bf3ffd29b(__obf_d4d18bbab73e7cfe)
	sort.Float64s(copy)
	return
}

// sortedCopyDif returns a sorted copy of float64s
// only if the original data isn't sorted.
// Only use this if returned slice won't be manipulated!
func __obf_bc477c12f4a7a87d(__obf_d4d18bbab73e7cfe []float64) (copy []float64) {
	if sort.Float64sAreSorted(__obf_d4d18bbab73e7cfe) {
		return __obf_d4d18bbab73e7cfe
	}
	copy = __obf_166de59bf3ffd29b(__obf_d4d18bbab73e7cfe)
	sort.Float64s(copy)
	return
}
