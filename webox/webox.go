package __obf_ca688c3e204e86eb

import (
	payment "github.com/ArtisanHiram/go-pkg/webox/payment"
	program "github.com/ArtisanHiram/go-pkg/webox/program"
)

type Webox struct {
	__obf_c12bd46dc9b896a0 *program.Config
	__obf_e6bb2f419ad294a1 *payment.Config
}

func New(__obf_c12bd46dc9b896a0 *program.Config, __obf_e6bb2f419ad294a1 *payment.Config) *Webox {
	return &Webox{__obf_c12bd46dc9b896a0, __obf_e6bb2f419ad294a1}
}

func (__obf_6df5d770465fa1ce *Webox) Program() (*program.Program, error) {
	return program.NewProgram(__obf_6df5d770465fa1ce.__obf_c12bd46dc9b896a0)
}

func (__obf_6df5d770465fa1ce *Webox) Payment() (*payment.Payment, error) {
	return payment.NewPayment(__obf_6df5d770465fa1ce.__obf_e6bb2f419ad294a1)
}
