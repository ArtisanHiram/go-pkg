package __obf_038560a94647875f

import consistent "github.com/ArtisanHiram/go-pkg/consistent"

func Subset[M consistent.Member](__obf_3f292fb80e2ecea2 string, __obf_4752dc15522b0582 []M, __obf_32b938b15d5d4375 int) []M {
	if len(__obf_4752dc15522b0582) <= __obf_32b938b15d5d4375 {
		return __obf_4752dc15522b0582
	}
	__obf_07feaaa91a36a31a := consistent.New[M]()
	__obf_07feaaa91a36a31a.
		NumberOfReplicas = 160
	__obf_07feaaa91a36a31a.
		UseFnv = true
	__obf_07feaaa91a36a31a.
		Set(__obf_4752dc15522b0582)

	return __obf_8b937cd2a06565da(__obf_07feaaa91a36a31a, __obf_3f292fb80e2ecea2, __obf_4752dc15522b0582, __obf_32b938b15d5d4375)
}

func __obf_8b937cd2a06565da[M consistent.Member](__obf_07feaaa91a36a31a *consistent.Consistent[M], __obf_3f292fb80e2ecea2 string, __obf_4752dc15522b0582 []M, __obf_32b938b15d5d4375 int) []M {
	__obf_0d8b959f6c1d21fc, __obf_ad71453048e280e8 := __obf_07feaaa91a36a31a.GetN(__obf_3f292fb80e2ecea2, __obf_32b938b15d5d4375)
	if __obf_ad71453048e280e8 != nil {
		return __obf_4752dc15522b0582
	}
	return __obf_0d8b959f6c1d21fc
}
