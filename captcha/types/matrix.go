package __obf_244ef50d49151021

import "math"

// Matrix struct for transformation calculations
type Matrix struct {
	XX, YX, XY, YY, X0, Y0 float64
}

// Translate performs matrix translation calculation
func (__obf_2b033369da82f821 Matrix) Translate(__obf_e487b3618d36e102, __obf_8cbd87ef15403513 float64) Matrix {
	return Matrix{
		1, 0,
		0, 1, __obf_e487b3618d36e102, __obf_8cbd87ef15403513,
	}.Multiply(__obf_2b033369da82f821)
}

// Multiply performs matrix multiplication
func (__obf_2b033369da82f821 Matrix) Multiply(__obf_5d4532d772799c07 Matrix) Matrix {
	return Matrix{__obf_2b033369da82f821.
		XX*__obf_5d4532d772799c07.XX + __obf_2b033369da82f821.YX*__obf_5d4532d772799c07.XY, __obf_2b033369da82f821.
		XX*__obf_5d4532d772799c07.YX + __obf_2b033369da82f821.YX*__obf_5d4532d772799c07.YY, __obf_2b033369da82f821.
		XY*__obf_5d4532d772799c07.XX + __obf_2b033369da82f821.YY*__obf_5d4532d772799c07.XY, __obf_2b033369da82f821.
		XY*__obf_5d4532d772799c07.YX + __obf_2b033369da82f821.YY*__obf_5d4532d772799c07.YY, __obf_2b033369da82f821.
		X0*__obf_5d4532d772799c07.XX + __obf_2b033369da82f821.Y0*__obf_5d4532d772799c07.XY + __obf_5d4532d772799c07.X0, __obf_2b033369da82f821.
		X0*__obf_5d4532d772799c07.YX + __obf_2b033369da82f821.Y0*__obf_5d4532d772799c07.YY + __obf_5d4532d772799c07.Y0,
	}
}

// Rotate performs matrix rotation calculation
func (__obf_2b033369da82f821 Matrix) Rotate(__obf_2b5072b762a61186 float64) Matrix {
	__obf_dbbef660ec878016 := math.Cos(__obf_2b5072b762a61186)
	__obf_964a4421c362d2e8 := math.Sin(__obf_2b5072b762a61186)
	return Matrix{__obf_dbbef660ec878016, __obf_964a4421c362d2e8, -__obf_964a4421c362d2e8, __obf_dbbef660ec878016, 0, 0,
	}.Multiply(__obf_2b033369da82f821)
}
