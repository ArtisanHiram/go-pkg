package __obf_7519c06aae7d41ec

import (
	cmd "github.com/ArtisanHiram/go-pkg/runner/cmd"
	"path/filepath"
)

func Run(__obf_ada1b3b75e7a0fba, __obf_b24bc0a0d5519e59 string, __obf_0a108dceb8cbc315 []string) (string, string, error) {
	__obf_16325cdf74aa9aa9 := "main.out"
	__obf_3bd9b66a613b9f3b := append([]string{"clang++", "-std=c++11", "-o", __obf_16325cdf74aa9aa9}, __obf_0a108dceb8cbc315...)
	__obf_bfb179b242d5be63, __obf_80eca01a49574b66, __obf_cb068f3ec75c4405 := cmd.Run(__obf_ada1b3b75e7a0fba, __obf_3bd9b66a613b9f3b...)
	if __obf_cb068f3ec75c4405 != nil {
		return __obf_bfb179b242d5be63, __obf_80eca01a49574b66, __obf_cb068f3ec75c4405
	}
	__obf_9ec2e68ec34f65c1 := filepath.Join(__obf_ada1b3b75e7a0fba, __obf_16325cdf74aa9aa9)
	return cmd.RunStdin(__obf_ada1b3b75e7a0fba, __obf_b24bc0a0d5519e59, __obf_9ec2e68ec34f65c1)
}
