package __obf_f2021de30eba0b59

import (
	payment "github.com/ArtisanHiram/go-pkg/webox/payment"
	program "github.com/ArtisanHiram/go-pkg/webox/program"
)

type Webox struct {
	__obf_b52ded055bc10af4 *program.Config
	__obf_7b737ff37aee73ea *payment.Config
}

func New(__obf_b52ded055bc10af4 *program.Config, __obf_7b737ff37aee73ea *payment.Config) *Webox {
	return &Webox{__obf_b52ded055bc10af4, __obf_7b737ff37aee73ea}
}

func (__obf_351025778a2bceae *Webox) Program() (*program.Program, error) {
	return program.NewProgram(__obf_351025778a2bceae.__obf_b52ded055bc10af4)
}

func (__obf_351025778a2bceae *Webox) Payment() (*payment.Payment, error) {
	return payment.NewPayment(__obf_351025778a2bceae.__obf_7b737ff37aee73ea)
}
