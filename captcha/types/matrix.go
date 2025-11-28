package __obf_b0f3614e51931450

import "math"

// Matrix struct for transformation calculations
type Matrix struct {
	XX, YX, XY, YY, X0, Y0 float64
}

// Translate performs matrix translation calculation
func (__obf_bd896093c58f47ac Matrix) Translate(__obf_80fae96c5947b2de, __obf_431a48e7317ece6d float64) Matrix {
	return Matrix{
		1, 0,
		0, 1,
		__obf_80fae96c5947b2de, __obf_431a48e7317ece6d,
	}.Multiply(__obf_bd896093c58f47ac)
}

// Multiply performs matrix multiplication
func (__obf_bd896093c58f47ac Matrix) Multiply(__obf_7a68195488a85895 Matrix) Matrix {
	return Matrix{
		__obf_bd896093c58f47ac.XX*__obf_7a68195488a85895.XX + __obf_bd896093c58f47ac.YX*__obf_7a68195488a85895.XY,
		__obf_bd896093c58f47ac.XX*__obf_7a68195488a85895.YX + __obf_bd896093c58f47ac.YX*__obf_7a68195488a85895.YY,
		__obf_bd896093c58f47ac.XY*__obf_7a68195488a85895.XX + __obf_bd896093c58f47ac.YY*__obf_7a68195488a85895.XY,
		__obf_bd896093c58f47ac.XY*__obf_7a68195488a85895.YX + __obf_bd896093c58f47ac.YY*__obf_7a68195488a85895.YY,
		__obf_bd896093c58f47ac.X0*__obf_7a68195488a85895.XX + __obf_bd896093c58f47ac.Y0*__obf_7a68195488a85895.XY + __obf_7a68195488a85895.X0,
		__obf_bd896093c58f47ac.X0*__obf_7a68195488a85895.YX + __obf_bd896093c58f47ac.Y0*__obf_7a68195488a85895.YY + __obf_7a68195488a85895.Y0,
	}
}

// Rotate performs matrix rotation calculation
func (__obf_bd896093c58f47ac Matrix) Rotate(__obf_31d599a870027db7 float64) Matrix {
	__obf_587b9fceec62883e := math.Cos(__obf_31d599a870027db7)
	__obf_98d8630ad167fcd8 := math.Sin(__obf_31d599a870027db7)
	return Matrix{
		__obf_587b9fceec62883e, __obf_98d8630ad167fcd8,
		-__obf_98d8630ad167fcd8, __obf_587b9fceec62883e,
		0, 0,
	}.Multiply(__obf_bd896093c58f47ac)
}
