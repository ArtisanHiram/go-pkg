package __obf_e16d36bb2331de95

import (
	cmd "github.com/ArtisanHiram/go-pkg/runner/cmd"
)

func Run(__obf_59d87904f42608dd string, __obf_f278d9e3a055f61a string, __obf_023cee25efaa6306 []string) (string, string, error) {
	__obf_be997ee38a1582bf := "main.out"
	__obf_cced19ffccfb0496, __obf_d99c1865fc43642d, __obf_8446dd0680194a21 := cmd.Run(__obf_59d87904f42608dd, "rustc", "-o", __obf_be997ee38a1582bf, __obf_023cee25efaa6306[0])
	if __obf_8446dd0680194a21 != nil {
		return __obf_cced19ffccfb0496, __obf_d99c1865fc43642d, __obf_8446dd0680194a21
	}

	return cmd.RunStdin(__obf_59d87904f42608dd, __obf_f278d9e3a055f61a, __obf_be997ee38a1582bf)
}
