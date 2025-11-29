package __obf_b0bebe5eb45b8ad6

import consistent "github.com/ArtisanHiram/go-pkg/consistent"

func Subset[M consistent.Member](__obf_e177826ca867ebb3 string, __obf_c350ddf11727d77a []M, __obf_03c75b3ae8ffc020 int) []M {
	if len(__obf_c350ddf11727d77a) <= __obf_03c75b3ae8ffc020 {
		return __obf_c350ddf11727d77a
	}
	__obf_13be556fca486f9c := consistent.New[M]()
	__obf_13be556fca486f9c.
		NumberOfReplicas = 160
	__obf_13be556fca486f9c.
		UseFnv = true
	__obf_13be556fca486f9c.
		Set(__obf_c350ddf11727d77a)

	return __obf_0ee40a86043cf1b6(__obf_13be556fca486f9c, __obf_e177826ca867ebb3, __obf_c350ddf11727d77a, __obf_03c75b3ae8ffc020)
}

func __obf_0ee40a86043cf1b6[M consistent.Member](__obf_13be556fca486f9c *consistent.Consistent[M], __obf_e177826ca867ebb3 string, __obf_c350ddf11727d77a []M, __obf_03c75b3ae8ffc020 int) []M {
	__obf_dc45440d8379a116, __obf_805eb0b5dd52f7ba := __obf_13be556fca486f9c.GetN(__obf_e177826ca867ebb3, __obf_03c75b3ae8ffc020)
	if __obf_805eb0b5dd52f7ba != nil {
		return __obf_c350ddf11727d77a
	}
	return __obf_dc45440d8379a116
}
