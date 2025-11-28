package __obf_0e73be9e4c3ea3fd

// RangeVal .
type Range struct {
	Min, Max int
}

// Size .
type Size struct {
	Width, Height int
}

// Point .
type Point struct {
	X, Y int
}

const (
	DistortNone = iota
	DistortLevel1
	DistortLevel2
	DistortLevel3
	DistortLevel4
	DistortLevel5
)

const (
	QualityNone   = 100
	QualityLevel1 = 95
	QualityLevel2 = 85
	QualityLevel3 = 75
	QualityLevel4 = 65
	QualityLevel5 = 55
)
