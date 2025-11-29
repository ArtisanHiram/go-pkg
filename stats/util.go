package __obf_8cdcb71eb89a0b1a

import (
	"sort"
	"time"
)

// float64ToInt rounds a float64 to an int
func __obf_f090a72098d77103(__obf_12053c429f32e343 float64) (__obf_6be1242c242f91c1 int) {
	__obf_bef96adf07ba8115, _ := Round(__obf_12053c429f32e343, 0)
	return int(__obf_bef96adf07ba8115)
}

// unixnano returns nanoseconds from UTC epoch
func __obf_7d3851aa335f9037() int64 {
	return time.Now().UTC().UnixNano()
}

// copyslice copies a slice of float64s
func __obf_3c935b53d69f7c63(__obf_12053c429f32e343 []float64) []float64 {
	__obf_69ff16df264fa90d := make([]float64, len(__obf_12053c429f32e343))
	copy(__obf_69ff16df264fa90d, __obf_12053c429f32e343)
	return __obf_69ff16df264fa90d
}

// sortedCopy returns a sorted copy of float64s
func __obf_c06efca49e526820(__obf_12053c429f32e343 []float64) (copy []float64) {
	copy = __obf_3c935b53d69f7c63(__obf_12053c429f32e343)
	sort.Float64s(copy)
	return
}

// sortedCopyDif returns a sorted copy of float64s
// only if the original data isn't sorted.
// Only use this if returned slice won't be manipulated!
func __obf_461a714533a4a654(__obf_12053c429f32e343 []float64) (copy []float64) {
	if sort.Float64sAreSorted(__obf_12053c429f32e343) {
		return __obf_12053c429f32e343
	}
	copy = __obf_3c935b53d69f7c63(__obf_12053c429f32e343)
	sort.Float64s(copy)
	return
}
