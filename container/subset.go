package __obf_90a4883a02d0b41c

import consistent "github.com/ArtisanHiram/go-pkg/consistent"

func Subset[M consistent.Member](__obf_444f272faed71fcb string, __obf_d4c0ec5cc888c041 []M, __obf_905046bf4fa2eaf9 int) []M {
	if len(__obf_d4c0ec5cc888c041) <= __obf_905046bf4fa2eaf9 {
		return __obf_d4c0ec5cc888c041
	}
	__obf_310b24e69c331a15 := consistent.New[M]()
	__obf_310b24e69c331a15.
		NumberOfReplicas = 160
	__obf_310b24e69c331a15.
		UseFnv = true
	__obf_310b24e69c331a15.
		Set(__obf_d4c0ec5cc888c041)

	return __obf_19a7e9379b654ebd(__obf_310b24e69c331a15, __obf_444f272faed71fcb, __obf_d4c0ec5cc888c041, __obf_905046bf4fa2eaf9)
}

func __obf_19a7e9379b654ebd[M consistent.Member](__obf_310b24e69c331a15 *consistent.Consistent[M], __obf_444f272faed71fcb string, __obf_d4c0ec5cc888c041 []M, __obf_905046bf4fa2eaf9 int) []M {
	__obf_b1cfb94621151fee, __obf_f14c805fc0e38747 := __obf_310b24e69c331a15.GetN(__obf_444f272faed71fcb, __obf_905046bf4fa2eaf9)
	if __obf_f14c805fc0e38747 != nil {
		return __obf_d4c0ec5cc888c041
	}
	return __obf_b1cfb94621151fee
}
