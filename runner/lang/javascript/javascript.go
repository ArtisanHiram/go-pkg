package __obf_e4b81726977bcc36

import (
	cmd "github.com/ArtisanHiram/go-pkg/runner/cmd"
)

func Run(__obf_e10ecdfac7828b7f string, __obf_3b654155b132ff9c string, __obf_b4f3b81d866b05b8 []string) (string, string, error) {
	return cmd.RunStdin(__obf_e10ecdfac7828b7f, __obf_3b654155b132ff9c, "node", __obf_b4f3b81d866b05b8[0])
}
