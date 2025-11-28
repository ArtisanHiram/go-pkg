package __obf_0c9b364e6d96f4f8

import (
	payment "github.com/ArtisanHiram/go-pkg/webox/payment"
	program "github.com/ArtisanHiram/go-pkg/webox/program"
)

type Webox struct {
	__obf_bfd5106cace80cfa *program.Config
	__obf_21255303eede605b *payment.Config
}

func New(__obf_bfd5106cace80cfa *program.Config, __obf_21255303eede605b *payment.Config) *Webox {
	return &Webox{__obf_bfd5106cace80cfa, __obf_21255303eede605b}
}

func (__obf_fed14cba04f89b53 *Webox) Program() (*program.Program, error) {
	return program.NewProgram(__obf_fed14cba04f89b53.__obf_bfd5106cace80cfa)
}

func (__obf_fed14cba04f89b53 *Webox) Payment() (*payment.Payment, error) {
	return payment.NewPayment(__obf_fed14cba04f89b53.__obf_21255303eede605b)
}
