package __obf_00574e11343a9c9b

type Money struct {
	__obf_bf250de3aae0304a float32 // 以分为单位存储
}

func NewMoneyFromYuan(__obf_1fc9845c109fb37c float32) Money {
	return Money{__obf_bf250de3aae0304a: float32(__obf_1fc9845c109fb37c*100 + 0.5)} // 四舍五入
}

func NewMoneyFromCent(__obf_bf250de3aae0304a float32) Money {
	return Money{__obf_bf250de3aae0304a: __obf_bf250de3aae0304a}
}

func (__obf_45a0b232299aef04 Money) Mul(__obf_c9aa27dd53e235ab float32) Money {
	__obf_bb6c514997aa88d3 := float32(__obf_45a0b232299aef04.__obf_bf250de3aae0304a) * __obf_c9aa27dd53e235ab
	return Money{__obf_bf250de3aae0304a: float32(__obf_bb6c514997aa88d3 + 0.5)} // 四舍五入
}

func (__obf_45a0b232299aef04 Money) MulInt(__obf_c9aa27dd53e235ab int) Money {
	return Money{__obf_bf250de3aae0304a: __obf_45a0b232299aef04.__obf_bf250de3aae0304a * float32(__obf_c9aa27dd53e235ab)}
}

func (__obf_45a0b232299aef04 Money) Yuan() float32 {
	return float32(__obf_45a0b232299aef04.__obf_bf250de3aae0304a) / 100.0
}
