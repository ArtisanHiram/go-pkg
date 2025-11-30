package __obf_187fb4813c1b0c74

type Money struct {
	__obf_6c8cf442c7a25339 float32 // 以分为单位存储
}

func NewMoneyFromYuan(__obf_0f3156395157e8c4 float32) Money {
	return Money{__obf_6c8cf442c7a25339: float32(__obf_0f3156395157e8c4*100 + 0.5)} // 四舍五入
}

func NewMoneyFromCent(__obf_6c8cf442c7a25339 float32) Money {
	return Money{__obf_6c8cf442c7a25339: __obf_6c8cf442c7a25339}
}

func (__obf_bc65b19795f2998e Money) Mul(__obf_177a85a5f3725ee3 float32) Money {
	__obf_e12e0a58baa6fa61 := float32(__obf_bc65b19795f2998e.__obf_6c8cf442c7a25339) * __obf_177a85a5f3725ee3
	return Money{__obf_6c8cf442c7a25339: float32(__obf_e12e0a58baa6fa61 + 0.5)} // 四舍五入
}

func (__obf_bc65b19795f2998e Money) MulInt(__obf_177a85a5f3725ee3 int) Money {
	return Money{__obf_6c8cf442c7a25339: __obf_bc65b19795f2998e.__obf_6c8cf442c7a25339 * float32(__obf_177a85a5f3725ee3)}
}

func (__obf_bc65b19795f2998e Money) Yuan() float32 {
	return float32(__obf_bc65b19795f2998e.__obf_6c8cf442c7a25339) / 100.0
}
