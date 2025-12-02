package __obf_89363901d8df63bc

import "math"

// Matrix struct for transformation calculations
type Matrix struct {
	XX, YX, XY, YY, X0, Y0 float64
}

// Translate performs matrix translation calculation
func (__obf_dbe5c645f5262529 Matrix) Translate(__obf_4d548ce157fe2d7b, __obf_e5fce044456d5b92 float64) Matrix {
	return Matrix{
		1, 0,
		0, 1, __obf_4d548ce157fe2d7b, __obf_e5fce044456d5b92,
	}.Multiply(__obf_dbe5c645f5262529)
}

// Multiply performs matrix multiplication
func (__obf_dbe5c645f5262529 Matrix) Multiply(__obf_b1200ec27b8040c7 Matrix) Matrix {
	return Matrix{__obf_dbe5c645f5262529.
		XX*__obf_b1200ec27b8040c7.XX + __obf_dbe5c645f5262529.YX*__obf_b1200ec27b8040c7.XY, __obf_dbe5c645f5262529.
		XX*__obf_b1200ec27b8040c7.YX + __obf_dbe5c645f5262529.YX*__obf_b1200ec27b8040c7.YY, __obf_dbe5c645f5262529.
		XY*__obf_b1200ec27b8040c7.XX + __obf_dbe5c645f5262529.YY*__obf_b1200ec27b8040c7.XY, __obf_dbe5c645f5262529.
		XY*__obf_b1200ec27b8040c7.YX + __obf_dbe5c645f5262529.YY*__obf_b1200ec27b8040c7.YY, __obf_dbe5c645f5262529.
		X0*__obf_b1200ec27b8040c7.XX + __obf_dbe5c645f5262529.Y0*__obf_b1200ec27b8040c7.XY + __obf_b1200ec27b8040c7.X0, __obf_dbe5c645f5262529.
		X0*__obf_b1200ec27b8040c7.YX + __obf_dbe5c645f5262529.Y0*__obf_b1200ec27b8040c7.YY + __obf_b1200ec27b8040c7.Y0,
	}
}

// Rotate performs matrix rotation calculation
func (__obf_dbe5c645f5262529 Matrix) Rotate(__obf_1f7091190512fbb8 float64) Matrix {
	__obf_8f1cd35f635f319d := math.Cos(__obf_1f7091190512fbb8)
	__obf_e2ddfa904e6c7b65 := math.Sin(__obf_1f7091190512fbb8)
	return Matrix{__obf_8f1cd35f635f319d, __obf_e2ddfa904e6c7b65, -__obf_e2ddfa904e6c7b65, __obf_8f1cd35f635f319d, 0, 0,
	}.Multiply(__obf_dbe5c645f5262529)
}
