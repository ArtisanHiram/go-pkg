package __obf_23e897073e46d7fb

type Money struct {
	__obf_de2c70b9284c08c2 float32 // 以分为单位存储
}

func NewMoneyFromYuan(__obf_2a8c0cd9319a910c float32) Money {
	return Money{__obf_de2c70b9284c08c2: float32(__obf_2a8c0cd9319a910c*100 + 0.5)} // 四舍五入
}

func NewMoneyFromCent(__obf_de2c70b9284c08c2 float32) Money {
	return Money{__obf_de2c70b9284c08c2: __obf_de2c70b9284c08c2}
}

func (__obf_ed2f7249dcfa47db Money) Mul(__obf_df85d8b13f63ee68 float32) Money {
	__obf_d006b26ad66365d0 := float32(__obf_ed2f7249dcfa47db.__obf_de2c70b9284c08c2) * __obf_df85d8b13f63ee68
	return Money{__obf_de2c70b9284c08c2: float32(__obf_d006b26ad66365d0 + 0.5)} // 四舍五入
}

func (__obf_ed2f7249dcfa47db Money) MulInt(__obf_df85d8b13f63ee68 int) Money {
	return Money{__obf_de2c70b9284c08c2: __obf_ed2f7249dcfa47db.__obf_de2c70b9284c08c2 * float32(__obf_df85d8b13f63ee68)}
}

func (__obf_ed2f7249dcfa47db Money) Yuan() float32 {
	return float32(__obf_ed2f7249dcfa47db.__obf_de2c70b9284c08c2) / 100.0
}
