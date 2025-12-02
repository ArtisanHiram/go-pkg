package __obf_614e0776906f0579

type Money struct {
	__obf_f27cc7d749cc931e float32 // 以分为单位存储
}

func NewMoneyFromYuan(__obf_f3420f103edf305a float32) Money {
	return Money{__obf_f27cc7d749cc931e: float32(__obf_f3420f103edf305a*100 + 0.5)} // 四舍五入
}

func NewMoneyFromCent(__obf_f27cc7d749cc931e float32) Money {
	return Money{__obf_f27cc7d749cc931e: __obf_f27cc7d749cc931e}
}

func (__obf_dc07899a79a63463 Money) Mul(__obf_3f7c5f9a4eca22e8 float32) Money {
	__obf_f5a8a3899a7474ef := float32(__obf_dc07899a79a63463.__obf_f27cc7d749cc931e) * __obf_3f7c5f9a4eca22e8
	return Money{__obf_f27cc7d749cc931e: float32(__obf_f5a8a3899a7474ef + 0.5)} // 四舍五入
}

func (__obf_dc07899a79a63463 Money) MulInt(__obf_3f7c5f9a4eca22e8 int) Money {
	return Money{__obf_f27cc7d749cc931e: __obf_dc07899a79a63463.__obf_f27cc7d749cc931e * float32(__obf_3f7c5f9a4eca22e8)}
}

func (__obf_dc07899a79a63463 Money) Yuan() float32 {
	return float32(__obf_dc07899a79a63463.__obf_f27cc7d749cc931e) / 100.0
}
