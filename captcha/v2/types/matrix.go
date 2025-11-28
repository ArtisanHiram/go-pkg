package __obf_deb7e65d40f46bd0

import "math"

// Matrix struct for transformation calculations
type Matrix struct {
	XX, YX, XY, YY, X0, Y0 float64
}

// Translate performs matrix translation calculation
func (__obf_a66951353e16056a Matrix) Translate(__obf_ce61f234cbe42808, __obf_7a88a48a3d7db1aa float64) Matrix {
	return Matrix{
		1, 0,
		0, 1,
		__obf_ce61f234cbe42808, __obf_7a88a48a3d7db1aa,
	}.Multiply(__obf_a66951353e16056a)
}

// Multiply performs matrix multiplication
func (__obf_a66951353e16056a Matrix) Multiply(__obf_fb5251c8851279ed Matrix) Matrix {
	return Matrix{
		__obf_a66951353e16056a.XX*__obf_fb5251c8851279ed.XX + __obf_a66951353e16056a.YX*__obf_fb5251c8851279ed.XY,
		__obf_a66951353e16056a.XX*__obf_fb5251c8851279ed.YX + __obf_a66951353e16056a.YX*__obf_fb5251c8851279ed.YY,
		__obf_a66951353e16056a.XY*__obf_fb5251c8851279ed.XX + __obf_a66951353e16056a.YY*__obf_fb5251c8851279ed.XY,
		__obf_a66951353e16056a.XY*__obf_fb5251c8851279ed.YX + __obf_a66951353e16056a.YY*__obf_fb5251c8851279ed.YY,
		__obf_a66951353e16056a.X0*__obf_fb5251c8851279ed.XX + __obf_a66951353e16056a.Y0*__obf_fb5251c8851279ed.XY + __obf_fb5251c8851279ed.X0,
		__obf_a66951353e16056a.X0*__obf_fb5251c8851279ed.YX + __obf_a66951353e16056a.Y0*__obf_fb5251c8851279ed.YY + __obf_fb5251c8851279ed.Y0,
	}
}

// Rotate performs matrix rotation calculation
func (__obf_a66951353e16056a Matrix) Rotate(__obf_0a4f6c9b020fe9b9 float64) Matrix {
	__obf_f012bb226e8e4a2d := math.Cos(__obf_0a4f6c9b020fe9b9)
	__obf_ee0cdc467cd6f854 := math.Sin(__obf_0a4f6c9b020fe9b9)
	return Matrix{
		__obf_f012bb226e8e4a2d, __obf_ee0cdc467cd6f854,
		-__obf_ee0cdc467cd6f854, __obf_f012bb226e8e4a2d,
		0, 0,
	}.Multiply(__obf_a66951353e16056a)
}
