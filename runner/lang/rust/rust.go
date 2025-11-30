package __obf_d129d68574ef26fa

import (
	cmd "github.com/ArtisanHiram/go-pkg/runner/cmd"
)

func Run(__obf_5e71d9b09c994bc6 string, __obf_0a601d262683a805 string, __obf_55db9cd486ce9483 []string) (string, string, error) {
	__obf_c9cd9a509e056e80 := "main.out"
	__obf_3fb5dea88cf2ab30, __obf_a6f1e92d73c0d44f, __obf_421970c81b33469c := cmd.Run(__obf_5e71d9b09c994bc6, "rustc", "-o", __obf_c9cd9a509e056e80, __obf_55db9cd486ce9483[0])
	if __obf_421970c81b33469c != nil {
		return __obf_3fb5dea88cf2ab30, __obf_a6f1e92d73c0d44f, __obf_421970c81b33469c
	}

	return cmd.RunStdin(__obf_5e71d9b09c994bc6, __obf_0a601d262683a805, __obf_c9cd9a509e056e80)
}
