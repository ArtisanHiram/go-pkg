package __obf_4c0216572d8cff29

import (
	cmd "github.com/ArtisanHiram/go-pkg/runner/cmd"
	"path/filepath"
)

func Run(__obf_e03e345676972a2a, __obf_28e6fb0ac5eb49bb string, __obf_5dd7a2c000f605ea []string) (string, string, error) {
	__obf_2c67aebf614ce7f1 := "main.out"
	__obf_084b85f67e0844b9 := append([]string{"clang++", "-std=c++11", "-o", __obf_2c67aebf614ce7f1}, __obf_5dd7a2c000f605ea...)
	__obf_d89dd0db288f0685, __obf_ce681b79b06de622, __obf_413c49014e9d6076 := cmd.Run(__obf_e03e345676972a2a, __obf_084b85f67e0844b9...)
	if __obf_413c49014e9d6076 != nil {
		return __obf_d89dd0db288f0685, __obf_ce681b79b06de622, __obf_413c49014e9d6076
	}
	__obf_78bd7d8cf3040a9d := filepath.Join(__obf_e03e345676972a2a, __obf_2c67aebf614ce7f1)
	return cmd.RunStdin(__obf_e03e345676972a2a, __obf_28e6fb0ac5eb49bb, __obf_78bd7d8cf3040a9d)
}
