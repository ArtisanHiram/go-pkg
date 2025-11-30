package __obf_b8170cecc56ae3b5

import (
	cmd "github.com/ArtisanHiram/go-pkg/runner/cmd"
)

func Run(__obf_dcad6c4b415ef288 string, __obf_28f36c5e2aa7d33e string, __obf_c02ce1dcbf4bf549 []string) (string, string, error) {
	__obf_b607046aa1aac4c2 := "main.out"
	__obf_170a334888f07e49, __obf_56da3358436ea1b1, __obf_876955285d9023f3 := cmd.Run(__obf_dcad6c4b415ef288, "rustc", "-o", __obf_b607046aa1aac4c2, __obf_c02ce1dcbf4bf549[0])
	if __obf_876955285d9023f3 != nil {
		return __obf_170a334888f07e49, __obf_56da3358436ea1b1, __obf_876955285d9023f3
	}

	return cmd.RunStdin(__obf_dcad6c4b415ef288, __obf_28f36c5e2aa7d33e, __obf_b607046aa1aac4c2)
}
