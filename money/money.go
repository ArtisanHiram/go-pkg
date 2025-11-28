package __obf_870946aace603599

type Money struct {
	__obf_9e6555e8dcea7a36 float32 // 以分为单位存储
}

func NewMoneyFromYuan(__obf_b711e680224572cb float32) Money {
	return Money{__obf_9e6555e8dcea7a36: float32(__obf_b711e680224572cb*100 + 0.5)} // 四舍五入
}

func NewMoneyFromCent(__obf_9e6555e8dcea7a36 float32) Money {
	return Money{__obf_9e6555e8dcea7a36: __obf_9e6555e8dcea7a36}
}

func (__obf_6f7babec8ad2a304 Money) Mul(__obf_aee596f687d0a7c9 float32) Money {
	__obf_b536b57f58a94e94 := float32(__obf_6f7babec8ad2a304.__obf_9e6555e8dcea7a36) * __obf_aee596f687d0a7c9
	return Money{__obf_9e6555e8dcea7a36: float32(__obf_b536b57f58a94e94 + 0.5)} // 四舍五入
}

func (__obf_6f7babec8ad2a304 Money) MulInt(__obf_aee596f687d0a7c9 int) Money {
	return Money{__obf_9e6555e8dcea7a36: __obf_6f7babec8ad2a304.__obf_9e6555e8dcea7a36 * float32(__obf_aee596f687d0a7c9)}
}

func (__obf_6f7babec8ad2a304 Money) Yuan() float32 {
	return float32(__obf_6f7babec8ad2a304.__obf_9e6555e8dcea7a36) / 100.0
}
