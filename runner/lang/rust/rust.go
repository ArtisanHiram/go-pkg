package __obf_9ec71e549a9e608c

import (
	cmd "github.com/ArtisanHiram/go-pkg/runner/cmd"
)

func Run(__obf_453ba07448681f71 string, __obf_ba233f730fcd25ad string, __obf_ab63274244f478e7 []string) (string, string, error) {
	__obf_4570f9bb900597c9 := "main.out"
	__obf_3d72140777eec05a, __obf_9532696c22fae7d6, __obf_0a3e11ef497c1b9d := cmd.Run(__obf_453ba07448681f71, "rustc", "-o", __obf_4570f9bb900597c9, __obf_ab63274244f478e7[0])
	if __obf_0a3e11ef497c1b9d != nil {
		return __obf_3d72140777eec05a, __obf_9532696c22fae7d6, __obf_0a3e11ef497c1b9d
	}

	return cmd.RunStdin(__obf_453ba07448681f71, __obf_ba233f730fcd25ad, __obf_4570f9bb900597c9)
}
