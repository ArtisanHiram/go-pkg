package __obf_38b1aa1214ff7dc7

import (
	cmd "github.com/ArtisanHiram/go-pkg/runner/cmd"
	"path/filepath"
)

func Run(__obf_6cf26d45e5410bf0, __obf_df7c86147cd55a32 string, __obf_2f460054267bfc23 []string) (string, string, error) {
	__obf_63ec3df9aa32ee70 := "main.out"

	__obf_290a3632de6ebb2f := append([]string{"clang++", "-std=c++11", "-o", __obf_63ec3df9aa32ee70}, __obf_2f460054267bfc23...)
	__obf_7d2506d701db05e7, __obf_903ae053aeacee9d, __obf_c8ec27a679af29d5 := cmd.Run(__obf_6cf26d45e5410bf0, __obf_290a3632de6ebb2f...)
	if __obf_c8ec27a679af29d5 != nil {
		return __obf_7d2506d701db05e7, __obf_903ae053aeacee9d, __obf_c8ec27a679af29d5
	}

	__obf_da89d9491aa6e5d1 := filepath.Join(__obf_6cf26d45e5410bf0, __obf_63ec3df9aa32ee70)
	return cmd.RunStdin(__obf_6cf26d45e5410bf0, __obf_df7c86147cd55a32, __obf_da89d9491aa6e5d1)
}
