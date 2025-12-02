package __obf_1ae793f3a03942f7

import (
	cmd "github.com/ArtisanHiram/go-pkg/runner/cmd"
	"path/filepath"
)

func Run(__obf_07d73652328d523f, __obf_6fd1c7558cacd7b1 string, __obf_43a7c73270744d05 []string) (string, string, error) {
	__obf_ab398d3831b78bf0 := "main.out"
	__obf_4d591e8445e476a0 := append([]string{"clang++", "-std=c++11", "-o", __obf_ab398d3831b78bf0}, __obf_43a7c73270744d05...)
	__obf_febe950d79518391, __obf_7483dcee85f22db0, __obf_bcfaa953378d8812 := cmd.Run(__obf_07d73652328d523f, __obf_4d591e8445e476a0...)
	if __obf_bcfaa953378d8812 != nil {
		return __obf_febe950d79518391, __obf_7483dcee85f22db0, __obf_bcfaa953378d8812
	}
	__obf_c43e861762e4fdfa := filepath.Join(__obf_07d73652328d523f, __obf_ab398d3831b78bf0)
	return cmd.RunStdin(__obf_07d73652328d523f, __obf_6fd1c7558cacd7b1, __obf_c43e861762e4fdfa)
}
