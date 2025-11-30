package __obf_02bec74a885ee521

import (
	payment "github.com/ArtisanHiram/go-pkg/webox/payment"
	program "github.com/ArtisanHiram/go-pkg/webox/program"
)

type Webox struct {
	__obf_a1ac2fbfbe5403a8 *program.Config
	__obf_127edd7a170b15da *payment.Config
}

func New(__obf_a1ac2fbfbe5403a8 *program.Config, __obf_127edd7a170b15da *payment.Config) *Webox {
	return &Webox{__obf_a1ac2fbfbe5403a8, __obf_127edd7a170b15da}
}

func (__obf_abc8161fb1aff46a *Webox) Program() (*program.Program, error) {
	return program.NewProgram(__obf_abc8161fb1aff46a.__obf_a1ac2fbfbe5403a8)
}

func (__obf_abc8161fb1aff46a *Webox) Payment() (*payment.Payment, error) {
	return payment.NewPayment(__obf_abc8161fb1aff46a.__obf_127edd7a170b15da)
}
