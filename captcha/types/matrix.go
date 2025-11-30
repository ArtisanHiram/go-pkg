package __obf_738b46210fdb4199

import "math"

// Matrix struct for transformation calculations
type Matrix struct {
	XX, YX, XY, YY, X0, Y0 float64
}

// Translate performs matrix translation calculation
func (__obf_5eeb26a04aed8320 Matrix) Translate(__obf_bd372f9dc9e16d75, __obf_a69b4d01937bfbfc float64) Matrix {
	return Matrix{
		1, 0,
		0, 1, __obf_bd372f9dc9e16d75, __obf_a69b4d01937bfbfc,
	}.Multiply(__obf_5eeb26a04aed8320)
}

// Multiply performs matrix multiplication
func (__obf_5eeb26a04aed8320 Matrix) Multiply(__obf_8a1d048ff9f8df0f Matrix) Matrix {
	return Matrix{__obf_5eeb26a04aed8320.
		XX*__obf_8a1d048ff9f8df0f.XX + __obf_5eeb26a04aed8320.YX*__obf_8a1d048ff9f8df0f.XY, __obf_5eeb26a04aed8320.
		XX*__obf_8a1d048ff9f8df0f.YX + __obf_5eeb26a04aed8320.YX*__obf_8a1d048ff9f8df0f.YY, __obf_5eeb26a04aed8320.
		XY*__obf_8a1d048ff9f8df0f.XX + __obf_5eeb26a04aed8320.YY*__obf_8a1d048ff9f8df0f.XY, __obf_5eeb26a04aed8320.
		XY*__obf_8a1d048ff9f8df0f.YX + __obf_5eeb26a04aed8320.YY*__obf_8a1d048ff9f8df0f.YY, __obf_5eeb26a04aed8320.
		X0*__obf_8a1d048ff9f8df0f.XX + __obf_5eeb26a04aed8320.Y0*__obf_8a1d048ff9f8df0f.XY + __obf_8a1d048ff9f8df0f.X0, __obf_5eeb26a04aed8320.
		X0*__obf_8a1d048ff9f8df0f.YX + __obf_5eeb26a04aed8320.Y0*__obf_8a1d048ff9f8df0f.YY + __obf_8a1d048ff9f8df0f.Y0,
	}
}

// Rotate performs matrix rotation calculation
func (__obf_5eeb26a04aed8320 Matrix) Rotate(__obf_3a2dad55c354b1d6 float64) Matrix {
	__obf_cf060e7b58cbd8c5 := math.Cos(__obf_3a2dad55c354b1d6)
	__obf_da3b85d7c237752d := math.Sin(__obf_3a2dad55c354b1d6)
	return Matrix{__obf_cf060e7b58cbd8c5, __obf_da3b85d7c237752d, -__obf_da3b85d7c237752d, __obf_cf060e7b58cbd8c5, 0, 0,
	}.Multiply(__obf_5eeb26a04aed8320)
}
