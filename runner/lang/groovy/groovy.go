package __obf_b68f91eece4cab06

import (
	cmd "github.com/ArtisanHiram/go-pkg/runner/cmd"
)

func Run(__obf_e2a93f145cadb913, __obf_7a996ab24db0f364 string, __obf_fd12521cc6efff3b []string) (string, string, error) {
	__obf_d03f7657ebb0eae4 := append([]string{"groovy"}, __obf_fd12521cc6efff3b...)
	return cmd.RunStdin(__obf_e2a93f145cadb913, __obf_7a996ab24db0f364, __obf_d03f7657ebb0eae4...)
}
