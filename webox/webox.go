package __obf_67818a79f770d3d0

import (
	payment "github.com/ArtisanHiram/go-pkg/webox/payment"
	program "github.com/ArtisanHiram/go-pkg/webox/program"
)

type Webox struct {
	__obf_ad6fddb6c151a691 *program.Config
	__obf_6bbe9f56a1280fb9 *payment.Config
}

func New(__obf_ad6fddb6c151a691 *program.Config, __obf_6bbe9f56a1280fb9 *payment.Config) *Webox {
	return &Webox{__obf_ad6fddb6c151a691, __obf_6bbe9f56a1280fb9}
}

func (__obf_f7d8477278268c80 *Webox) Program() (*program.Program, error) {
	return program.NewProgram(__obf_f7d8477278268c80.__obf_ad6fddb6c151a691)
}

func (__obf_f7d8477278268c80 *Webox) Payment() (*payment.Payment, error) {
	return payment.NewPayment(__obf_f7d8477278268c80.__obf_6bbe9f56a1280fb9)
}
