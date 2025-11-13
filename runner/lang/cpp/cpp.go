package __obf_5a0db05711df14d3

import (
	cmd "github.com/ArtisanHiram/go-pkg/runner/cmd"
	"path/filepath"
)

func Run(__obf_7b494f6a3cb17f3e, __obf_be3f40fa84ad6d53 string, __obf_50506651c383532f []string) (string, string, error) {
	__obf_5639fbfccea9bfe8 := "main.out"

	__obf_ffb3a8d42ea9582c := append([]string{"clang++", "-std=c++11", "-o", __obf_5639fbfccea9bfe8}, __obf_50506651c383532f...)
	__obf_2d255f0e5111391e, __obf_8b884ac68c6fa462, __obf_60f70f238dec4ef2 := cmd.Run(__obf_7b494f6a3cb17f3e, __obf_ffb3a8d42ea9582c...)
	if __obf_60f70f238dec4ef2 != nil {
		return __obf_2d255f0e5111391e, __obf_8b884ac68c6fa462, __obf_60f70f238dec4ef2
	}

	__obf_f465328d12b10bd1 := filepath.Join(__obf_7b494f6a3cb17f3e, __obf_5639fbfccea9bfe8)
	return cmd.RunStdin(__obf_7b494f6a3cb17f3e, __obf_be3f40fa84ad6d53, __obf_f465328d12b10bd1)
}
