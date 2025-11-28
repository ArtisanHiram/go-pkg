package __obf_0d0d05e597ad9adf

import (
	"sort"
	"time"
)

// float64ToInt rounds a float64 to an int
func __obf_b3d3a21f3716a950(__obf_d044dbbfe899ea30 float64) (__obf_600fe6a3c6b93031 int) {
	__obf_60a2c96891d426c1, _ := Round(__obf_d044dbbfe899ea30, 0)
	return int(__obf_60a2c96891d426c1)
}

// unixnano returns nanoseconds from UTC epoch
func __obf_150443b011911cdc() int64 {
	return time.Now().UTC().UnixNano()
}

// copyslice copies a slice of float64s
func __obf_0a4a2aaf883880d7(__obf_d044dbbfe899ea30 []float64) []float64 {
	__obf_53454715f79fbdf4 := make([]float64, len(__obf_d044dbbfe899ea30))
	copy(__obf_53454715f79fbdf4, __obf_d044dbbfe899ea30)
	return __obf_53454715f79fbdf4
}

// sortedCopy returns a sorted copy of float64s
func __obf_67ec427ad1535581(__obf_d044dbbfe899ea30 []float64) (copy []float64) {
	copy = __obf_0a4a2aaf883880d7(__obf_d044dbbfe899ea30)
	sort.Float64s(copy)
	return
}

// sortedCopyDif returns a sorted copy of float64s
// only if the original data isn't sorted.
// Only use this if returned slice won't be manipulated!
func __obf_d195a604ff666c09(__obf_d044dbbfe899ea30 []float64) (copy []float64) {
	if sort.Float64sAreSorted(__obf_d044dbbfe899ea30) {
		return __obf_d044dbbfe899ea30
	}
	copy = __obf_0a4a2aaf883880d7(__obf_d044dbbfe899ea30)
	sort.Float64s(copy)
	return
}
