package __obf_0a39789dd8bc2b45

type Money struct {
	__obf_b70d37bcd38a16b9 float32 // 以分为单位存储
}

func NewMoneyFromYuan(__obf_49509bc2cbb69448 float32) Money {
	return Money{__obf_b70d37bcd38a16b9: float32(__obf_49509bc2cbb69448*100 + 0.5)} // 四舍五入
}

func NewMoneyFromCent(__obf_b70d37bcd38a16b9 float32) Money {
	return Money{__obf_b70d37bcd38a16b9: __obf_b70d37bcd38a16b9}
}

func (__obf_47d2da1a173b9ef1 Money) Mul(__obf_f264d27a23260120 float32) Money {
	__obf_35314d12d0899e0a := float32(__obf_47d2da1a173b9ef1.__obf_b70d37bcd38a16b9) * __obf_f264d27a23260120
	return Money{__obf_b70d37bcd38a16b9: float32(__obf_35314d12d0899e0a + 0.5)} // 四舍五入
}

func (__obf_47d2da1a173b9ef1 Money) MulInt(__obf_f264d27a23260120 int) Money {
	return Money{__obf_b70d37bcd38a16b9: __obf_47d2da1a173b9ef1.__obf_b70d37bcd38a16b9 * float32(__obf_f264d27a23260120)}
}

func (__obf_47d2da1a173b9ef1 Money) Yuan() float32 {
	return float32(__obf_47d2da1a173b9ef1.__obf_b70d37bcd38a16b9) / 100.0
}
