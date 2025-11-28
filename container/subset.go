package __obf_af42fb6cde2beed6

import consistent "github.com/ArtisanHiram/go-pkg/consistent"

func Subset[M consistent.Member](__obf_bf7a39439a0629b6 string, __obf_10409a6d1dcdfe09 []M, __obf_a18c495ebc00d86e int) []M {
	if len(__obf_10409a6d1dcdfe09) <= __obf_a18c495ebc00d86e {
		return __obf_10409a6d1dcdfe09
	}

	__obf_b8edc6183e50f0e0 := consistent.New[M]()
	__obf_b8edc6183e50f0e0.NumberOfReplicas = 160
	__obf_b8edc6183e50f0e0.UseFnv = true
	__obf_b8edc6183e50f0e0.Set(__obf_10409a6d1dcdfe09)

	return __obf_a3e746036282c4af(__obf_b8edc6183e50f0e0, __obf_bf7a39439a0629b6, __obf_10409a6d1dcdfe09, __obf_a18c495ebc00d86e)
}

func __obf_a3e746036282c4af[M consistent.Member](__obf_b8edc6183e50f0e0 *consistent.Consistent[M], __obf_bf7a39439a0629b6 string, __obf_10409a6d1dcdfe09 []M, __obf_a18c495ebc00d86e int) []M {
	__obf_e3ae521dae03af73, __obf_24456bd3f3e29186 := __obf_b8edc6183e50f0e0.GetN(__obf_bf7a39439a0629b6, __obf_a18c495ebc00d86e)
	if __obf_24456bd3f3e29186 != nil {
		return __obf_10409a6d1dcdfe09
	}
	return __obf_e3ae521dae03af73
}
