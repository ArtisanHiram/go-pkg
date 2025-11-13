package __obf_76599ab96a208083

import consistent "github.com/ArtisanHiram/go-pkg/consistent"

func Subset[M consistent.Member](__obf_16e8235f971fe768 string, __obf_b0d65775420bf7ba []M, __obf_8b642e8339b97a1d int) []M {
	if len(__obf_b0d65775420bf7ba) <= __obf_8b642e8339b97a1d {
		return __obf_b0d65775420bf7ba
	}

	__obf_80f06068a59506dc := consistent.New[M]()
	__obf_80f06068a59506dc.NumberOfReplicas = 160
	__obf_80f06068a59506dc.UseFnv = true
	__obf_80f06068a59506dc.Set(__obf_b0d65775420bf7ba)

	return __obf_5376e06b33e7838a(__obf_80f06068a59506dc, __obf_16e8235f971fe768, __obf_b0d65775420bf7ba, __obf_8b642e8339b97a1d)
}

func __obf_5376e06b33e7838a[M consistent.Member](__obf_80f06068a59506dc *consistent.Consistent[M], __obf_16e8235f971fe768 string, __obf_b0d65775420bf7ba []M, __obf_8b642e8339b97a1d int) []M {
	__obf_8a2cc52ff76ffd6c, __obf_791dcfd55f252396 := __obf_80f06068a59506dc.GetN(__obf_16e8235f971fe768, __obf_8b642e8339b97a1d)
	if __obf_791dcfd55f252396 != nil {
		return __obf_b0d65775420bf7ba
	}
	return __obf_8a2cc52ff76ffd6c
}
