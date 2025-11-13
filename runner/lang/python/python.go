package __obf_3b0a121a7e2f375d

import (
	cmd "github.com/ArtisanHiram/go-pkg/runner/cmd"
)

func Run(__obf_6281bc8eddaadd20 string, __obf_abba0bfa0fa00988 string, __obf_2973092ce533de46 []string) (string, string, error) {
	return cmd.RunStdin(__obf_6281bc8eddaadd20, __obf_abba0bfa0fa00988, "python3", __obf_2973092ce533de46[0])
}
