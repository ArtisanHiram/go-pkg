package __obf_4f13ac5aa043b5ce

import consistent "github.com/ArtisanHiram/go-pkg/consistent"

func Subset[M consistent.Member](__obf_a23b52dc799eca35 string, __obf_4861b5765fcbc608 []M, __obf_86221501b2c6377b int) []M {
	if len(__obf_4861b5765fcbc608) <= __obf_86221501b2c6377b {
		return __obf_4861b5765fcbc608
	}

	__obf_9c6178a3e5920061 := consistent.New[M]()
	__obf_9c6178a3e5920061.NumberOfReplicas = 160
	__obf_9c6178a3e5920061.UseFnv = true
	__obf_9c6178a3e5920061.Set(__obf_4861b5765fcbc608)

	return __obf_079050093c1fb0ef(__obf_9c6178a3e5920061, __obf_a23b52dc799eca35, __obf_4861b5765fcbc608, __obf_86221501b2c6377b)
}

func __obf_079050093c1fb0ef[M consistent.Member](__obf_9c6178a3e5920061 *consistent.Consistent[M], __obf_a23b52dc799eca35 string, __obf_4861b5765fcbc608 []M, __obf_86221501b2c6377b int) []M {
	__obf_e3689caf271fb55a, __obf_52943f8b1bce7293 := __obf_9c6178a3e5920061.GetN(__obf_a23b52dc799eca35, __obf_86221501b2c6377b)
	if __obf_52943f8b1bce7293 != nil {
		return __obf_4861b5765fcbc608
	}
	return __obf_e3689caf271fb55a
}
