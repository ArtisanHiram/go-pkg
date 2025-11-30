package __obf_d98d0ee49a41b141

import (
	payment "github.com/ArtisanHiram/go-pkg/webox/payment"
	program "github.com/ArtisanHiram/go-pkg/webox/program"
)

type Webox struct {
	__obf_a41195da92ff3b99 *program.Config
	__obf_4b74b58f7ff7b991 *payment.Config
}

func New(__obf_a41195da92ff3b99 *program.Config, __obf_4b74b58f7ff7b991 *payment.Config) *Webox {
	return &Webox{__obf_a41195da92ff3b99, __obf_4b74b58f7ff7b991}
}

func (__obf_19064d37b2404e78 *Webox) Program() (*program.Program, error) {
	return program.NewProgram(__obf_19064d37b2404e78.__obf_a41195da92ff3b99)
}

func (__obf_19064d37b2404e78 *Webox) Payment() (*payment.Payment, error) {
	return payment.NewPayment(__obf_19064d37b2404e78.__obf_4b74b58f7ff7b991)
}
