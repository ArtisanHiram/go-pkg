package __obf_8606789be7a3f36a

import (
	payment "github.com/ArtisanHiram/go-pkg/webox/payment"
	program "github.com/ArtisanHiram/go-pkg/webox/program"
)

type Webox struct {
	__obf_53650f24960e8869 *program.Config
	__obf_04d17c3f6db5a7fa *payment.Config
}

func New(__obf_53650f24960e8869 *program.Config, __obf_04d17c3f6db5a7fa *payment.Config) *Webox {
	return &Webox{__obf_53650f24960e8869, __obf_04d17c3f6db5a7fa}
}

func (__obf_78c7db3a2004c151 *Webox) Program() (*program.Program, error) {
	return program.NewProgram(__obf_78c7db3a2004c151.__obf_53650f24960e8869)
}

func (__obf_78c7db3a2004c151 *Webox) Payment() (*payment.Payment, error) {
	return payment.NewPayment(__obf_78c7db3a2004c151.__obf_04d17c3f6db5a7fa)
}
