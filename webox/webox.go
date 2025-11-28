package __obf_d37b430dd5da4a63

import (
	payment "github.com/ArtisanHiram/go-pkg/webox/payment"
	program "github.com/ArtisanHiram/go-pkg/webox/program"
)

type Webox struct {
	__obf_8a6166a8c5f65a69 *program.Config
	__obf_ad23757a16329023 *payment.Config
}

func New(__obf_8a6166a8c5f65a69 *program.Config, __obf_ad23757a16329023 *payment.Config) *Webox {
	return &Webox{__obf_8a6166a8c5f65a69, __obf_ad23757a16329023}
}

func (__obf_c9ff9ad3ffdf1a10 *Webox) Program() (*program.Program, error) {
	return program.NewProgram(__obf_c9ff9ad3ffdf1a10.__obf_8a6166a8c5f65a69)
}

func (__obf_c9ff9ad3ffdf1a10 *Webox) Payment() (*payment.Payment, error) {
	return payment.NewPayment(__obf_c9ff9ad3ffdf1a10.__obf_ad23757a16329023)
}
