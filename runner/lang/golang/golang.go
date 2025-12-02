package __obf_ddec83d9d8f876c1

import (
	cmd "github.com/ArtisanHiram/go-pkg/runner/cmd"
)

func Run(__obf_97782a60a496282b, __obf_2057b8f9000e61c7 string, __obf_cb451cd6f9c1cb78 []string) (string, string, error) {
	__obf_ece408a5339be3af := append([]string{"go", "run"}, __obf_cb451cd6f9c1cb78...)
	return cmd.RunStdin(__obf_97782a60a496282b, __obf_2057b8f9000e61c7, __obf_ece408a5339be3af...)
}
