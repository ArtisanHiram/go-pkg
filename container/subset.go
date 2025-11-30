package __obf_9004b47202f9204b

import consistent "github.com/ArtisanHiram/go-pkg/consistent"

func Subset[M consistent.Member](__obf_fa71c403e75a023f string, __obf_d6d9c150ce481970 []M, __obf_c33e6799eff4e4a8 int) []M {
	if len(__obf_d6d9c150ce481970) <= __obf_c33e6799eff4e4a8 {
		return __obf_d6d9c150ce481970
	}
	__obf_4162b39e2a9c29bf := consistent.New[M]()
	__obf_4162b39e2a9c29bf.
		NumberOfReplicas = 160
	__obf_4162b39e2a9c29bf.
		UseFnv = true
	__obf_4162b39e2a9c29bf.
		Set(__obf_d6d9c150ce481970)

	return __obf_6ec7f401c6cf8afd(__obf_4162b39e2a9c29bf, __obf_fa71c403e75a023f, __obf_d6d9c150ce481970, __obf_c33e6799eff4e4a8)
}

func __obf_6ec7f401c6cf8afd[M consistent.Member](__obf_4162b39e2a9c29bf *consistent.Consistent[M], __obf_fa71c403e75a023f string, __obf_d6d9c150ce481970 []M, __obf_c33e6799eff4e4a8 int) []M {
	__obf_2a182f2e0d56f550, __obf_4a0f6bb7d9302e64 := __obf_4162b39e2a9c29bf.GetN(__obf_fa71c403e75a023f, __obf_c33e6799eff4e4a8)
	if __obf_4a0f6bb7d9302e64 != nil {
		return __obf_d6d9c150ce481970
	}
	return __obf_2a182f2e0d56f550
}
