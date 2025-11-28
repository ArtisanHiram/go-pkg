package __obf_409180cd94ee825d

import (
	cmd "github.com/ArtisanHiram/go-pkg/runner/cmd"
)

func Run(__obf_3ff602c8f6e55ed5, __obf_90efb2b570ed437a string, __obf_b2b2a4a492b06f38 []string) (string, string, error) {
	__obf_2fc0313ddaad01a8 := append([]string{"groovy"}, __obf_b2b2a4a492b06f38...)
	return cmd.RunStdin(__obf_3ff602c8f6e55ed5, __obf_90efb2b570ed437a, __obf_2fc0313ddaad01a8...)
}
