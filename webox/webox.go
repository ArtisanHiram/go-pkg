package __obf_5cc9a2903338cdd7

import (
	payment "github.com/ArtisanHiram/go-pkg/webox/payment"
	program "github.com/ArtisanHiram/go-pkg/webox/program"
)

type Webox struct {
	__obf_046d9980ac734b80 *program.Config
	__obf_34a55a050daec53b *payment.Config
}

func New(__obf_046d9980ac734b80 *program.Config, __obf_34a55a050daec53b *payment.Config) *Webox {
	return &Webox{__obf_046d9980ac734b80, __obf_34a55a050daec53b}
}

func (__obf_f99f6560e316027a *Webox) Program() (*program.Program, error) {
	return program.NewProgram(__obf_f99f6560e316027a.__obf_046d9980ac734b80)
}

func (__obf_f99f6560e316027a *Webox) Payment() (*payment.Payment, error) {
	return payment.NewPayment(__obf_f99f6560e316027a.__obf_34a55a050daec53b)
}
