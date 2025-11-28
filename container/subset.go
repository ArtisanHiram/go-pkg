package __obf_1fda7fbdeda52f1e

import consistent "github.com/ArtisanHiram/go-pkg/consistent"

func Subset[M consistent.Member](__obf_0319c73a9fcfea78 string, __obf_e1d97ccd4f404a0c []M, __obf_de5e180fe13c970f int) []M {
	if len(__obf_e1d97ccd4f404a0c) <= __obf_de5e180fe13c970f {
		return __obf_e1d97ccd4f404a0c
	}

	__obf_253b4b4b09bc3548 := consistent.New[M]()
	__obf_253b4b4b09bc3548.NumberOfReplicas = 160
	__obf_253b4b4b09bc3548.UseFnv = true
	__obf_253b4b4b09bc3548.Set(__obf_e1d97ccd4f404a0c)

	return __obf_3d27992e137eac13(__obf_253b4b4b09bc3548, __obf_0319c73a9fcfea78, __obf_e1d97ccd4f404a0c, __obf_de5e180fe13c970f)
}

func __obf_3d27992e137eac13[M consistent.Member](__obf_253b4b4b09bc3548 *consistent.Consistent[M], __obf_0319c73a9fcfea78 string, __obf_e1d97ccd4f404a0c []M, __obf_de5e180fe13c970f int) []M {
	__obf_0f4e1927b9296c23, __obf_483c2f6b078721a6 := __obf_253b4b4b09bc3548.GetN(__obf_0319c73a9fcfea78, __obf_de5e180fe13c970f)
	if __obf_483c2f6b078721a6 != nil {
		return __obf_e1d97ccd4f404a0c
	}
	return __obf_0f4e1927b9296c23
}
