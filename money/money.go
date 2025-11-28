package __obf_5a25a54b16760cc7

type Money struct {
	__obf_e163b93accad15e8 float32 // 以分为单位存储
}

func NewMoneyFromYuan(__obf_52803f7849b22a85 float32) Money {
	return Money{__obf_e163b93accad15e8: float32(__obf_52803f7849b22a85*100 + 0.5)} // 四舍五入
}

func NewMoneyFromCent(__obf_e163b93accad15e8 float32) Money {
	return Money{__obf_e163b93accad15e8: __obf_e163b93accad15e8}
}

func (__obf_6cacd9da604b2e69 Money) Mul(__obf_d47ff1cf59b56779 float32) Money {
	__obf_6ad144ddc0c6df04 := float32(__obf_6cacd9da604b2e69.__obf_e163b93accad15e8) * __obf_d47ff1cf59b56779
	return Money{__obf_e163b93accad15e8: float32(__obf_6ad144ddc0c6df04 + 0.5)} // 四舍五入
}

func (__obf_6cacd9da604b2e69 Money) MulInt(__obf_d47ff1cf59b56779 int) Money {
	return Money{__obf_e163b93accad15e8: __obf_6cacd9da604b2e69.__obf_e163b93accad15e8 * float32(__obf_d47ff1cf59b56779)}
}

func (__obf_6cacd9da604b2e69 Money) Yuan() float32 {
	return float32(__obf_6cacd9da604b2e69.__obf_e163b93accad15e8) / 100.0
}
