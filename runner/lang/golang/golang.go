package __obf_b272d4798a2d6881

import (
	cmd "github.com/ArtisanHiram/go-pkg/runner/cmd"
)

func Run(__obf_2aa0f7d517f7c6e3, __obf_4b14e53126a4a746 string, __obf_8169d076dab35677 []string) (string, string, error) {
	__obf_435345dda839a9b0 := append([]string{"go", "run"}, __obf_8169d076dab35677...)
	return cmd.RunStdin(__obf_2aa0f7d517f7c6e3, __obf_4b14e53126a4a746, __obf_435345dda839a9b0...)
}
