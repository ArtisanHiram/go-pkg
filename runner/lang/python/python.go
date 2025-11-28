package __obf_53638c918199ba82

import (
	cmd "github.com/ArtisanHiram/go-pkg/runner/cmd"
)

func Run(__obf_a7b4475e7c8cc428 string, __obf_286acede4def39da string, __obf_a1e0d9696d838bc8 []string) (string, string, error) {
	return cmd.RunStdin(__obf_a7b4475e7c8cc428, __obf_286acede4def39da, "python3", __obf_a1e0d9696d838bc8[0])
}
