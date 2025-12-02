package __obf_6361570abfd596bf

import (
	payment "github.com/ArtisanHiram/go-pkg/webox/payment"
	program "github.com/ArtisanHiram/go-pkg/webox/program"
)

type Webox struct {
	__obf_8b334ee159b59183 *program.Config
	__obf_1ca7c05e731c9ed1 *payment.Config
}

func New(__obf_8b334ee159b59183 *program.Config, __obf_1ca7c05e731c9ed1 *payment.Config) *Webox {
	return &Webox{__obf_8b334ee159b59183, __obf_1ca7c05e731c9ed1}
}

func (__obf_f03e7a7db0960b99 *Webox) Program() (*program.Program, error) {
	return program.NewProgram(__obf_f03e7a7db0960b99.__obf_8b334ee159b59183)
}

func (__obf_f03e7a7db0960b99 *Webox) Payment() (*payment.Payment, error) {
	return payment.NewPayment(__obf_f03e7a7db0960b99.__obf_1ca7c05e731c9ed1)
}
