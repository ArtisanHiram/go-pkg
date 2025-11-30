package __obf_e72ce603d10d02a1

import consistent "github.com/ArtisanHiram/go-pkg/consistent"

func Subset[M consistent.Member](__obf_9ae0d9fd67afd0d9 string, __obf_c2ca538d2769cd4b []M, __obf_b686586ed8451139 int) []M {
	if len(__obf_c2ca538d2769cd4b) <= __obf_b686586ed8451139 {
		return __obf_c2ca538d2769cd4b
	}
	__obf_fab71928d52cd4c8 := consistent.New[M]()
	__obf_fab71928d52cd4c8.
		NumberOfReplicas = 160
	__obf_fab71928d52cd4c8.
		UseFnv = true
	__obf_fab71928d52cd4c8.
		Set(__obf_c2ca538d2769cd4b)

	return __obf_e9c2692e9b7cba77(__obf_fab71928d52cd4c8, __obf_9ae0d9fd67afd0d9, __obf_c2ca538d2769cd4b, __obf_b686586ed8451139)
}

func __obf_e9c2692e9b7cba77[M consistent.Member](__obf_fab71928d52cd4c8 *consistent.Consistent[M], __obf_9ae0d9fd67afd0d9 string, __obf_c2ca538d2769cd4b []M, __obf_b686586ed8451139 int) []M {
	__obf_8e19b33164208f18, __obf_3674305c6f99212d := __obf_fab71928d52cd4c8.GetN(__obf_9ae0d9fd67afd0d9, __obf_b686586ed8451139)
	if __obf_3674305c6f99212d != nil {
		return __obf_c2ca538d2769cd4b
	}
	return __obf_8e19b33164208f18
}
