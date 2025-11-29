package __obf_bda21a78cb74016a

import "math"

// Matrix struct for transformation calculations
type Matrix struct {
	XX, YX, XY, YY, X0, Y0 float64
}

// Translate performs matrix translation calculation
func (__obf_6ed92d728f70797c Matrix) Translate(__obf_e22e7e43df48995e, __obf_a4dd1bd05990217f float64) Matrix {
	return Matrix{
		1, 0,
		0, 1, __obf_e22e7e43df48995e, __obf_a4dd1bd05990217f,
	}.Multiply(__obf_6ed92d728f70797c)
}

// Multiply performs matrix multiplication
func (__obf_6ed92d728f70797c Matrix) Multiply(__obf_cd225e6d73fa2d13 Matrix) Matrix {
	return Matrix{__obf_6ed92d728f70797c.
		XX*__obf_cd225e6d73fa2d13.XX + __obf_6ed92d728f70797c.YX*__obf_cd225e6d73fa2d13.XY, __obf_6ed92d728f70797c.
		XX*__obf_cd225e6d73fa2d13.YX + __obf_6ed92d728f70797c.YX*__obf_cd225e6d73fa2d13.YY, __obf_6ed92d728f70797c.
		XY*__obf_cd225e6d73fa2d13.XX + __obf_6ed92d728f70797c.YY*__obf_cd225e6d73fa2d13.XY, __obf_6ed92d728f70797c.
		XY*__obf_cd225e6d73fa2d13.YX + __obf_6ed92d728f70797c.YY*__obf_cd225e6d73fa2d13.YY, __obf_6ed92d728f70797c.
		X0*__obf_cd225e6d73fa2d13.XX + __obf_6ed92d728f70797c.Y0*__obf_cd225e6d73fa2d13.XY + __obf_cd225e6d73fa2d13.X0, __obf_6ed92d728f70797c.
		X0*__obf_cd225e6d73fa2d13.YX + __obf_6ed92d728f70797c.Y0*__obf_cd225e6d73fa2d13.YY + __obf_cd225e6d73fa2d13.Y0,
	}
}

// Rotate performs matrix rotation calculation
func (__obf_6ed92d728f70797c Matrix) Rotate(__obf_1ed21c16a991a4c4 float64) Matrix {
	__obf_6ac27ae0c7f24e19 := math.Cos(__obf_1ed21c16a991a4c4)
	__obf_f2659f3da765dd14 := math.Sin(__obf_1ed21c16a991a4c4)
	return Matrix{__obf_6ac27ae0c7f24e19, __obf_f2659f3da765dd14, -__obf_f2659f3da765dd14, __obf_6ac27ae0c7f24e19, 0, 0,
	}.Multiply(__obf_6ed92d728f70797c)
}
