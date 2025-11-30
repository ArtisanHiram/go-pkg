package __obf_54406b1a1de84196

import "math"

// Matrix struct for transformation calculations
type Matrix struct {
	XX, YX, XY, YY, X0, Y0 float64
}

// Translate performs matrix translation calculation
func (__obf_23448ffa2f9aadbc Matrix) Translate(__obf_48afb3005cd4a35c, __obf_33a1f511b09ac2af float64) Matrix {
	return Matrix{
		1, 0,
		0, 1, __obf_48afb3005cd4a35c, __obf_33a1f511b09ac2af,
	}.Multiply(__obf_23448ffa2f9aadbc)
}

// Multiply performs matrix multiplication
func (__obf_23448ffa2f9aadbc Matrix) Multiply(__obf_0f18c71e80ef2569 Matrix) Matrix {
	return Matrix{__obf_23448ffa2f9aadbc.
		XX*__obf_0f18c71e80ef2569.XX + __obf_23448ffa2f9aadbc.YX*__obf_0f18c71e80ef2569.XY, __obf_23448ffa2f9aadbc.
		XX*__obf_0f18c71e80ef2569.YX + __obf_23448ffa2f9aadbc.YX*__obf_0f18c71e80ef2569.YY, __obf_23448ffa2f9aadbc.
		XY*__obf_0f18c71e80ef2569.XX + __obf_23448ffa2f9aadbc.YY*__obf_0f18c71e80ef2569.XY, __obf_23448ffa2f9aadbc.
		XY*__obf_0f18c71e80ef2569.YX + __obf_23448ffa2f9aadbc.YY*__obf_0f18c71e80ef2569.YY, __obf_23448ffa2f9aadbc.
		X0*__obf_0f18c71e80ef2569.XX + __obf_23448ffa2f9aadbc.Y0*__obf_0f18c71e80ef2569.XY + __obf_0f18c71e80ef2569.X0, __obf_23448ffa2f9aadbc.
		X0*__obf_0f18c71e80ef2569.YX + __obf_23448ffa2f9aadbc.Y0*__obf_0f18c71e80ef2569.YY + __obf_0f18c71e80ef2569.Y0,
	}
}

// Rotate performs matrix rotation calculation
func (__obf_23448ffa2f9aadbc Matrix) Rotate(__obf_27c01b4675c92ab0 float64) Matrix {
	__obf_a26ee860b3cee35b := math.Cos(__obf_27c01b4675c92ab0)
	__obf_0f0924068fb8be37 := math.Sin(__obf_27c01b4675c92ab0)
	return Matrix{__obf_a26ee860b3cee35b, __obf_0f0924068fb8be37, -__obf_0f0924068fb8be37, __obf_a26ee860b3cee35b, 0, 0,
	}.Multiply(__obf_23448ffa2f9aadbc)
}
