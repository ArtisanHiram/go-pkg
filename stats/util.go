package __obf_0c4a280c8ca42808

import (
	"sort"
	"time"
)

// float64ToInt rounds a float64 to an int
func __obf_132364f3f9926888(__obf_b97842801be35eb9 float64) (__obf_c60a0b9d5c2f3b70 int) {
	__obf_2ac22434916fa087, _ := Round(__obf_b97842801be35eb9, 0)
	return int(__obf_2ac22434916fa087)
}

// unixnano returns nanoseconds from UTC epoch
func __obf_c4ad74d538ad4d24() int64 {
	return time.Now().UTC().UnixNano()
}

// copyslice copies a slice of float64s
func __obf_8dd1ba3bfd9a162e(__obf_b97842801be35eb9 []float64) []float64 {
	__obf_232c61af255d55a3 := make([]float64, len(__obf_b97842801be35eb9))
	copy(__obf_232c61af255d55a3, __obf_b97842801be35eb9)
	return __obf_232c61af255d55a3
}

// sortedCopy returns a sorted copy of float64s
func __obf_858f6951a9c2c144(__obf_b97842801be35eb9 []float64) (copy []float64) {
	copy = __obf_8dd1ba3bfd9a162e(__obf_b97842801be35eb9)
	sort.Float64s(copy)
	return
}

// sortedCopyDif returns a sorted copy of float64s
// only if the original data isn't sorted.
// Only use this if returned slice won't be manipulated!
func __obf_3739ecea2c533e8f(__obf_b97842801be35eb9 []float64) (copy []float64) {
	if sort.Float64sAreSorted(__obf_b97842801be35eb9) {
		return __obf_b97842801be35eb9
	}
	copy = __obf_8dd1ba3bfd9a162e(__obf_b97842801be35eb9)
	sort.Float64s(copy)
	return
}
