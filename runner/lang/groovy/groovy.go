package __obf_c4f50f83705855c6

import (
	cmd "github.com/ArtisanHiram/go-pkg/runner/cmd"
)

func Run(__obf_0c36a5219081457b, __obf_7d80a7e5c3376105 string, __obf_dd6af5fbd9283448 []string) (string, string, error) {
	__obf_4ac91d1457fdee11 := append([]string{"groovy"}, __obf_dd6af5fbd9283448...)
	return cmd.RunStdin(__obf_0c36a5219081457b, __obf_7d80a7e5c3376105, __obf_4ac91d1457fdee11...)
}
