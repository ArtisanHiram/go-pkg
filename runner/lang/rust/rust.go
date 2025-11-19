package __obf_4122b7e4450aa84e

import (
	cmd "github.com/ArtisanHiram/go-pkg/runner/cmd"
)

func Run(__obf_28956a563dd3b5ea string, __obf_7fd677c2bc84de90 string, __obf_1d5d877252979fd9 []string) (string, string, error) {
	__obf_5643fbce8e45f60e := "main.out"
	__obf_f9d49454b8011b8d, __obf_a7c30a85f1935895, __obf_e5fabf9658c9654a := cmd.Run(__obf_28956a563dd3b5ea, "rustc", "-o", __obf_5643fbce8e45f60e, __obf_1d5d877252979fd9[0])
	if __obf_e5fabf9658c9654a != nil {
		return __obf_f9d49454b8011b8d, __obf_a7c30a85f1935895, __obf_e5fabf9658c9654a
	}

	return cmd.RunStdin(__obf_28956a563dd3b5ea, __obf_7fd677c2bc84de90, __obf_5643fbce8e45f60e)
}
