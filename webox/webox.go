package __obf_85d38431f55a80f3

import (
	payment "github.com/ArtisanHiram/go-pkg/webox/payment"
	program "github.com/ArtisanHiram/go-pkg/webox/program"
)

type Webox struct {
	__obf_b408367454de77b6 *program.Config
	__obf_7f3d5b2a9e6a8938 *payment.Config
}

func New(__obf_b408367454de77b6 *program.Config, __obf_7f3d5b2a9e6a8938 *payment.Config) *Webox {
	return &Webox{__obf_b408367454de77b6, __obf_7f3d5b2a9e6a8938}
}

func (__obf_75e3902733ad482c *Webox) Program() (*program.Program, error) {
	return program.NewProgram(__obf_75e3902733ad482c.__obf_b408367454de77b6)
}

func (__obf_75e3902733ad482c *Webox) Payment() (*payment.Payment, error) {
	return payment.NewPayment(__obf_75e3902733ad482c.__obf_7f3d5b2a9e6a8938)
}
