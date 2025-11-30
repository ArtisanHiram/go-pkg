package __obf_1eb83f9c3ae53a6a

type Money struct {
	__obf_6c1dbfd659a08b2d float32 // 以分为单位存储
}

func NewMoneyFromYuan(__obf_ab13135a50747d20 float32) Money {
	return Money{__obf_6c1dbfd659a08b2d: float32(__obf_ab13135a50747d20*100 + 0.5)} // 四舍五入
}

func NewMoneyFromCent(__obf_6c1dbfd659a08b2d float32) Money {
	return Money{__obf_6c1dbfd659a08b2d: __obf_6c1dbfd659a08b2d}
}

func (__obf_48767758225f290b Money) Mul(__obf_d3393ba548ba7bd5 float32) Money {
	__obf_3867cb24c167f140 := float32(__obf_48767758225f290b.__obf_6c1dbfd659a08b2d) * __obf_d3393ba548ba7bd5
	return Money{__obf_6c1dbfd659a08b2d: float32(__obf_3867cb24c167f140 + 0.5)} // 四舍五入
}

func (__obf_48767758225f290b Money) MulInt(__obf_d3393ba548ba7bd5 int) Money {
	return Money{__obf_6c1dbfd659a08b2d: __obf_48767758225f290b.__obf_6c1dbfd659a08b2d * float32(__obf_d3393ba548ba7bd5)}
}

func (__obf_48767758225f290b Money) Yuan() float32 {
	return float32(__obf_48767758225f290b.__obf_6c1dbfd659a08b2d) / 100.0
}
