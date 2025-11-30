package __obf_6d820b36fdaf42ad

import (
	cmd "github.com/ArtisanHiram/go-pkg/runner/cmd"
)

func Run(__obf_243a1f004f7f7dc1 string, __obf_1ce56ef8ed96f931 string, __obf_f75c9ec638417924 []string) (string, string, error) {
	return cmd.RunStdin(__obf_243a1f004f7f7dc1, __obf_1ce56ef8ed96f931, "python3", __obf_f75c9ec638417924[0])
}
