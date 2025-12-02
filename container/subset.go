package __obf_4a16ef421fb74992

import consistent "github.com/ArtisanHiram/go-pkg/consistent"

func Subset[M consistent.Member](__obf_d3f032c0b695d18e string, __obf_724ec43b234f8ee1 []M, __obf_13d0191a8cbef0c2 int) []M {
	if len(__obf_724ec43b234f8ee1) <= __obf_13d0191a8cbef0c2 {
		return __obf_724ec43b234f8ee1
	}
	__obf_78ca14630fccf3ba := consistent.New[M]()
	__obf_78ca14630fccf3ba.
		NumberOfReplicas = 160
	__obf_78ca14630fccf3ba.
		UseFnv = true
	__obf_78ca14630fccf3ba.
		Set(__obf_724ec43b234f8ee1)

	return __obf_ca1cf556d6dc529b(__obf_78ca14630fccf3ba, __obf_d3f032c0b695d18e, __obf_724ec43b234f8ee1, __obf_13d0191a8cbef0c2)
}

func __obf_ca1cf556d6dc529b[M consistent.Member](__obf_78ca14630fccf3ba *consistent.Consistent[M], __obf_d3f032c0b695d18e string, __obf_724ec43b234f8ee1 []M, __obf_13d0191a8cbef0c2 int) []M {
	__obf_4bafbee6bcca8aef, __obf_fdc252da0d475ab8 := __obf_78ca14630fccf3ba.GetN(__obf_d3f032c0b695d18e, __obf_13d0191a8cbef0c2)
	if __obf_fdc252da0d475ab8 != nil {
		return __obf_724ec43b234f8ee1
	}
	return __obf_4bafbee6bcca8aef
}
