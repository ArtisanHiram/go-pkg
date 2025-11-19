package __obf_51237886feadd75e

type Money struct {
	__obf_e05a1fb87d026c9d float32 // 以分为单位存储
}

func NewMoneyFromYuan(__obf_8161fd6bb3739dda float32) Money {
	return Money{__obf_e05a1fb87d026c9d: float32(__obf_8161fd6bb3739dda*100 + 0.5)} // 四舍五入
}

func NewMoneyFromCent(__obf_e05a1fb87d026c9d float32) Money {
	return Money{__obf_e05a1fb87d026c9d: __obf_e05a1fb87d026c9d}
}

func (__obf_b33e474febddc9f0 Money) Mul(__obf_8a6f7c5456e8708f float32) Money {
	__obf_22de033def0894d3 := float32(__obf_b33e474febddc9f0.__obf_e05a1fb87d026c9d) * __obf_8a6f7c5456e8708f
	return Money{__obf_e05a1fb87d026c9d: float32(__obf_22de033def0894d3 + 0.5)} // 四舍五入
}

func (__obf_b33e474febddc9f0 Money) MulInt(__obf_8a6f7c5456e8708f int) Money {
	return Money{__obf_e05a1fb87d026c9d: __obf_b33e474febddc9f0.__obf_e05a1fb87d026c9d * float32(__obf_8a6f7c5456e8708f)}
}

func (__obf_b33e474febddc9f0 Money) Yuan() float32 {
	return float32(__obf_b33e474febddc9f0.__obf_e05a1fb87d026c9d) / 100.0
}
