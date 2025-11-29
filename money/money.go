package __obf_41ba6a5db66b1419

type Money struct {
	__obf_49ae156876d6f526 float32 // 以分为单位存储
}

func NewMoneyFromYuan(__obf_25a9071cb7d3f76d float32) Money {
	return Money{__obf_49ae156876d6f526: float32(__obf_25a9071cb7d3f76d*100 + 0.5)} // 四舍五入
}

func NewMoneyFromCent(__obf_49ae156876d6f526 float32) Money {
	return Money{__obf_49ae156876d6f526: __obf_49ae156876d6f526}
}

func (__obf_f1118cc0ee1bbe62 Money) Mul(__obf_58b1337f4064febc float32) Money {
	__obf_07ad66d96eb9ceb8 := float32(__obf_f1118cc0ee1bbe62.__obf_49ae156876d6f526) * __obf_58b1337f4064febc
	return Money{__obf_49ae156876d6f526: float32(__obf_07ad66d96eb9ceb8 + 0.5)} // 四舍五入
}

func (__obf_f1118cc0ee1bbe62 Money) MulInt(__obf_58b1337f4064febc int) Money {
	return Money{__obf_49ae156876d6f526: __obf_f1118cc0ee1bbe62.__obf_49ae156876d6f526 * float32(__obf_58b1337f4064febc)}
}

func (__obf_f1118cc0ee1bbe62 Money) Yuan() float32 {
	return float32(__obf_f1118cc0ee1bbe62.__obf_49ae156876d6f526) / 100.0
}
