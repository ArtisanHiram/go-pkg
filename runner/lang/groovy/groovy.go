package __obf_1a1a02cb38ab629e

import (
	cmd "github.com/ArtisanHiram/go-pkg/runner/cmd"
)

func Run(__obf_dd402c7b52c63a49, __obf_d941ac6a881b87f1 string, __obf_da5efe142f4143f4 []string) (string, string, error) {
	__obf_182e0791639e7e0a := append([]string{"groovy"}, __obf_da5efe142f4143f4...)
	return cmd.RunStdin(__obf_dd402c7b52c63a49, __obf_d941ac6a881b87f1, __obf_182e0791639e7e0a...)
}
