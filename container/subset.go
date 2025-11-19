package __obf_9861fa13140c30a3

import consistent "github.com/ArtisanHiram/go-pkg/consistent"

func Subset[M consistent.Member](__obf_88598b8460889c5c string, __obf_6dd34d7dc9aa71a5 []M, __obf_8384db2042441bf1 int) []M {
	if len(__obf_6dd34d7dc9aa71a5) <= __obf_8384db2042441bf1 {
		return __obf_6dd34d7dc9aa71a5
	}

	__obf_ce773a87e88fbaa5 := consistent.New[M]()
	__obf_ce773a87e88fbaa5.NumberOfReplicas = 160
	__obf_ce773a87e88fbaa5.UseFnv = true
	__obf_ce773a87e88fbaa5.Set(__obf_6dd34d7dc9aa71a5)

	return __obf_9ef6c7f344225c92(__obf_ce773a87e88fbaa5, __obf_88598b8460889c5c, __obf_6dd34d7dc9aa71a5, __obf_8384db2042441bf1)
}

func __obf_9ef6c7f344225c92[M consistent.Member](__obf_ce773a87e88fbaa5 *consistent.Consistent[M], __obf_88598b8460889c5c string, __obf_6dd34d7dc9aa71a5 []M, __obf_8384db2042441bf1 int) []M {
	__obf_d247953c21c32c89, __obf_d372e4dd7b4813ca := __obf_ce773a87e88fbaa5.GetN(__obf_88598b8460889c5c, __obf_8384db2042441bf1)
	if __obf_d372e4dd7b4813ca != nil {
		return __obf_6dd34d7dc9aa71a5
	}
	return __obf_d247953c21c32c89
}
