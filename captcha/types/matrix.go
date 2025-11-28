package __obf_154ce31cd9492d61

import "math"

// Matrix struct for transformation calculations
type Matrix struct {
	XX, YX, XY, YY, X0, Y0 float64
}

// Translate performs matrix translation calculation
func (__obf_91f28fffbdd83b09 Matrix) Translate(__obf_c87d6c7ba4a30d82, __obf_f8fead3d682d659f float64) Matrix {
	return Matrix{
		1, 0,
		0, 1,
		__obf_c87d6c7ba4a30d82, __obf_f8fead3d682d659f,
	}.Multiply(__obf_91f28fffbdd83b09)
}

// Multiply performs matrix multiplication
func (__obf_91f28fffbdd83b09 Matrix) Multiply(__obf_591e97521b833981 Matrix) Matrix {
	return Matrix{
		__obf_91f28fffbdd83b09.XX*__obf_591e97521b833981.XX + __obf_91f28fffbdd83b09.YX*__obf_591e97521b833981.XY,
		__obf_91f28fffbdd83b09.XX*__obf_591e97521b833981.YX + __obf_91f28fffbdd83b09.YX*__obf_591e97521b833981.YY,
		__obf_91f28fffbdd83b09.XY*__obf_591e97521b833981.XX + __obf_91f28fffbdd83b09.YY*__obf_591e97521b833981.XY,
		__obf_91f28fffbdd83b09.XY*__obf_591e97521b833981.YX + __obf_91f28fffbdd83b09.YY*__obf_591e97521b833981.YY,
		__obf_91f28fffbdd83b09.X0*__obf_591e97521b833981.XX + __obf_91f28fffbdd83b09.Y0*__obf_591e97521b833981.XY + __obf_591e97521b833981.X0,
		__obf_91f28fffbdd83b09.X0*__obf_591e97521b833981.YX + __obf_91f28fffbdd83b09.Y0*__obf_591e97521b833981.YY + __obf_591e97521b833981.Y0,
	}
}

// Rotate performs matrix rotation calculation
func (__obf_91f28fffbdd83b09 Matrix) Rotate(__obf_2c1097e2e9ce1f24 float64) Matrix {
	__obf_ad3a8dc916d9ec04 := math.Cos(__obf_2c1097e2e9ce1f24)
	__obf_a031dfe0cbcad8bb := math.Sin(__obf_2c1097e2e9ce1f24)
	return Matrix{
		__obf_ad3a8dc916d9ec04, __obf_a031dfe0cbcad8bb,
		-__obf_a031dfe0cbcad8bb, __obf_ad3a8dc916d9ec04,
		0, 0,
	}.Multiply(__obf_91f28fffbdd83b09)
}
