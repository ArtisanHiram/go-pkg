package __obf_a0e7cee4e3172ca9

import (
	cmd "github.com/ArtisanHiram/go-pkg/runner/cmd"
)

func Run(__obf_01c8f650878aba64 string, __obf_f1dac52dfd8e9ea3 string, __obf_d9273e3c00791174 []string) (string, string, error) {
	return cmd.RunStdin(__obf_01c8f650878aba64, __obf_f1dac52dfd8e9ea3, "python3", __obf_d9273e3c00791174[0])
}
