package __obf_62eba4024f8fa381

import consistent "github.com/ArtisanHiram/go-pkg/consistent"

func Subset[M consistent.Member](__obf_0c079c1bfcfbdc7b string, __obf_de86993792205f46 []M, __obf_a5ce6248fdb5a361 int) []M {
	if len(__obf_de86993792205f46) <= __obf_a5ce6248fdb5a361 {
		return __obf_de86993792205f46
	}

	__obf_3e9896b352dd367e := consistent.New[M]()
	__obf_3e9896b352dd367e.NumberOfReplicas = 160
	__obf_3e9896b352dd367e.UseFnv = true
	__obf_3e9896b352dd367e.Set(__obf_de86993792205f46)

	return __obf_3aa35f345159f648(__obf_3e9896b352dd367e, __obf_0c079c1bfcfbdc7b, __obf_de86993792205f46, __obf_a5ce6248fdb5a361)
}

func __obf_3aa35f345159f648[M consistent.Member](__obf_3e9896b352dd367e *consistent.Consistent[M], __obf_0c079c1bfcfbdc7b string, __obf_de86993792205f46 []M, __obf_a5ce6248fdb5a361 int) []M {
	__obf_5697716b3992c6dc, __obf_a4a3a0fe86b36626 := __obf_3e9896b352dd367e.GetN(__obf_0c079c1bfcfbdc7b, __obf_a5ce6248fdb5a361)
	if __obf_a4a3a0fe86b36626 != nil {
		return __obf_de86993792205f46
	}
	return __obf_5697716b3992c6dc
}
