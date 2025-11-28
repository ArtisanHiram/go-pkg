package __obf_c78eb1e7d5bda6d2

type Money struct {
	__obf_9d9f9d954a562916 float32 // 以分为单位存储
}

func NewMoneyFromYuan(__obf_e6271c614452439c float32) Money {
	return Money{__obf_9d9f9d954a562916: float32(__obf_e6271c614452439c*100 + 0.5)} // 四舍五入
}

func NewMoneyFromCent(__obf_9d9f9d954a562916 float32) Money {
	return Money{__obf_9d9f9d954a562916: __obf_9d9f9d954a562916}
}

func (__obf_f980d5392e8a988b Money) Mul(__obf_c2886e6306c9cd43 float32) Money {
	__obf_99f146607f86c79c := float32(__obf_f980d5392e8a988b.__obf_9d9f9d954a562916) * __obf_c2886e6306c9cd43
	return Money{__obf_9d9f9d954a562916: float32(__obf_99f146607f86c79c + 0.5)} // 四舍五入
}

func (__obf_f980d5392e8a988b Money) MulInt(__obf_c2886e6306c9cd43 int) Money {
	return Money{__obf_9d9f9d954a562916: __obf_f980d5392e8a988b.__obf_9d9f9d954a562916 * float32(__obf_c2886e6306c9cd43)}
}

func (__obf_f980d5392e8a988b Money) Yuan() float32 {
	return float32(__obf_f980d5392e8a988b.__obf_9d9f9d954a562916) / 100.0
}
