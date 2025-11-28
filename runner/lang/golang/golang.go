package __obf_d3324d866d63b779

import (
	cmd "github.com/ArtisanHiram/go-pkg/runner/cmd"
)

func Run(__obf_6d8e3ce7c3b4be6a, __obf_0fdc0b5e75171ac6 string, __obf_c012407a00b1499d []string) (string, string, error) {
	__obf_276676fec06f4794 := append([]string{"go", "run"}, __obf_c012407a00b1499d...)
	return cmd.RunStdin(__obf_6d8e3ce7c3b4be6a, __obf_0fdc0b5e75171ac6, __obf_276676fec06f4794...)
}
