package __obf_6dcb1d06bd949756

import "math"

// Matrix struct for transformation calculations
type Matrix struct {
	XX, YX, XY, YY, X0, Y0 float64
}

// Translate performs matrix translation calculation
func (__obf_59e7114cb8b11964 Matrix) Translate(__obf_57d99e8bd0a3444e, __obf_3376ebf4b35235fa float64) Matrix {
	return Matrix{
		1, 0,
		0, 1, __obf_57d99e8bd0a3444e, __obf_3376ebf4b35235fa,
	}.Multiply(__obf_59e7114cb8b11964)
}

// Multiply performs matrix multiplication
func (__obf_59e7114cb8b11964 Matrix) Multiply(__obf_fa7e21dc009e8a26 Matrix) Matrix {
	return Matrix{__obf_59e7114cb8b11964.
		XX*__obf_fa7e21dc009e8a26.XX + __obf_59e7114cb8b11964.YX*__obf_fa7e21dc009e8a26.XY, __obf_59e7114cb8b11964.
		XX*__obf_fa7e21dc009e8a26.YX + __obf_59e7114cb8b11964.YX*__obf_fa7e21dc009e8a26.YY, __obf_59e7114cb8b11964.
		XY*__obf_fa7e21dc009e8a26.XX + __obf_59e7114cb8b11964.YY*__obf_fa7e21dc009e8a26.XY, __obf_59e7114cb8b11964.
		XY*__obf_fa7e21dc009e8a26.YX + __obf_59e7114cb8b11964.YY*__obf_fa7e21dc009e8a26.YY, __obf_59e7114cb8b11964.
		X0*__obf_fa7e21dc009e8a26.XX + __obf_59e7114cb8b11964.Y0*__obf_fa7e21dc009e8a26.XY + __obf_fa7e21dc009e8a26.X0, __obf_59e7114cb8b11964.
		X0*__obf_fa7e21dc009e8a26.YX + __obf_59e7114cb8b11964.Y0*__obf_fa7e21dc009e8a26.YY + __obf_fa7e21dc009e8a26.Y0,
	}
}

// Rotate performs matrix rotation calculation
func (__obf_59e7114cb8b11964 Matrix) Rotate(__obf_dadbdf6b4f60b255 float64) Matrix {
	__obf_6b14c339de9bea34 := math.Cos(__obf_dadbdf6b4f60b255)
	__obf_496b2fa368070dfc := math.Sin(__obf_dadbdf6b4f60b255)
	return Matrix{__obf_6b14c339de9bea34, __obf_496b2fa368070dfc, -__obf_496b2fa368070dfc, __obf_6b14c339de9bea34, 0, 0,
	}.Multiply(__obf_59e7114cb8b11964)
}
