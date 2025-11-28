package __obf_aadd419806152164

import (
	payment "github.com/ArtisanHiram/go-pkg/webox/payment"
	program "github.com/ArtisanHiram/go-pkg/webox/program"
)

type Webox struct {
	__obf_a49fa78ca7bb6721 *program.Config
	__obf_1f78b17021904a98 *payment.Config
}

func New(__obf_a49fa78ca7bb6721 *program.Config, __obf_1f78b17021904a98 *payment.Config) *Webox {
	return &Webox{__obf_a49fa78ca7bb6721, __obf_1f78b17021904a98}
}

func (__obf_139b820b3db4e625 *Webox) Program() (*program.Program, error) {
	return program.NewProgram(__obf_139b820b3db4e625.__obf_a49fa78ca7bb6721)
}

func (__obf_139b820b3db4e625 *Webox) Payment() (*payment.Payment, error) {
	return payment.NewPayment(__obf_139b820b3db4e625.__obf_1f78b17021904a98)
}
