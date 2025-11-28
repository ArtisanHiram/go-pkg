package __obf_0e73be9e4c3ea3fd

import "math"

// Matrix struct for transformation calculations
type Matrix struct {
	XX, YX, XY, YY, X0, Y0 float64
}

// Translate performs matrix translation calculation
func (__obf_51d4fa10e4721296 Matrix) Translate(__obf_fe086773828b5535, __obf_44ef04afc5c8241a float64) Matrix {
	return Matrix{
		1, 0,
		0, 1,
		__obf_fe086773828b5535, __obf_44ef04afc5c8241a,
	}.Multiply(__obf_51d4fa10e4721296)
}

// Multiply performs matrix multiplication
func (__obf_51d4fa10e4721296 Matrix) Multiply(__obf_e3b1e50668ab4d43 Matrix) Matrix {
	return Matrix{
		__obf_51d4fa10e4721296.XX*__obf_e3b1e50668ab4d43.XX + __obf_51d4fa10e4721296.YX*__obf_e3b1e50668ab4d43.XY,
		__obf_51d4fa10e4721296.XX*__obf_e3b1e50668ab4d43.YX + __obf_51d4fa10e4721296.YX*__obf_e3b1e50668ab4d43.YY,
		__obf_51d4fa10e4721296.XY*__obf_e3b1e50668ab4d43.XX + __obf_51d4fa10e4721296.YY*__obf_e3b1e50668ab4d43.XY,
		__obf_51d4fa10e4721296.XY*__obf_e3b1e50668ab4d43.YX + __obf_51d4fa10e4721296.YY*__obf_e3b1e50668ab4d43.YY,
		__obf_51d4fa10e4721296.X0*__obf_e3b1e50668ab4d43.XX + __obf_51d4fa10e4721296.Y0*__obf_e3b1e50668ab4d43.XY + __obf_e3b1e50668ab4d43.X0,
		__obf_51d4fa10e4721296.X0*__obf_e3b1e50668ab4d43.YX + __obf_51d4fa10e4721296.Y0*__obf_e3b1e50668ab4d43.YY + __obf_e3b1e50668ab4d43.Y0,
	}
}

// Rotate performs matrix rotation calculation
func (__obf_51d4fa10e4721296 Matrix) Rotate(__obf_b51ed5d77ee8eb95 float64) Matrix {
	__obf_a5de954533461107 := math.Cos(__obf_b51ed5d77ee8eb95)
	__obf_f2c63d2da6f1cab9 := math.Sin(__obf_b51ed5d77ee8eb95)
	return Matrix{
		__obf_a5de954533461107, __obf_f2c63d2da6f1cab9,
		-__obf_f2c63d2da6f1cab9, __obf_a5de954533461107,
		0, 0,
	}.Multiply(__obf_51d4fa10e4721296)
}
