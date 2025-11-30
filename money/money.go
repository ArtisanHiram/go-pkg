package __obf_cd127ed70989ee72

type Money struct {
	__obf_8a64422747a2b3f3 float32 // 以分为单位存储
}

func NewMoneyFromYuan(__obf_193c6441f2a465e3 float32) Money {
	return Money{__obf_8a64422747a2b3f3: float32(__obf_193c6441f2a465e3*100 + 0.5)} // 四舍五入
}

func NewMoneyFromCent(__obf_8a64422747a2b3f3 float32) Money {
	return Money{__obf_8a64422747a2b3f3: __obf_8a64422747a2b3f3}
}

func (__obf_4fee5dc1ff7b8e3a Money) Mul(__obf_7c2e43e423310448 float32) Money {
	__obf_77379ceb9daabbcf := float32(__obf_4fee5dc1ff7b8e3a.__obf_8a64422747a2b3f3) * __obf_7c2e43e423310448
	return Money{__obf_8a64422747a2b3f3: float32(__obf_77379ceb9daabbcf + 0.5)} // 四舍五入
}

func (__obf_4fee5dc1ff7b8e3a Money) MulInt(__obf_7c2e43e423310448 int) Money {
	return Money{__obf_8a64422747a2b3f3: __obf_4fee5dc1ff7b8e3a.__obf_8a64422747a2b3f3 * float32(__obf_7c2e43e423310448)}
}

func (__obf_4fee5dc1ff7b8e3a Money) Yuan() float32 {
	return float32(__obf_4fee5dc1ff7b8e3a.__obf_8a64422747a2b3f3) / 100.0
}
