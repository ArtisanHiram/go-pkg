package __obf_5822bdcf62d42556

import (
	cmd "github.com/ArtisanHiram/go-pkg/runner/cmd"
	"path/filepath"
)

func Run(__obf_bc1041cfd5aedf7a string, __obf_c6dc68184448cfb8 string, __obf_6370588d58c6ad08 []string) (string, string, error) {
	__obf_78c114c1635b3e33, __obf_cd834910a4d1699f, __obf_39defddfec19008f := cmd.Run(__obf_bc1041cfd5aedf7a, "javac", __obf_6370588d58c6ad08[0])
	if __obf_39defddfec19008f != nil {
		return __obf_78c114c1635b3e33, __obf_cd834910a4d1699f, __obf_39defddfec19008f
	}
	return cmd.RunStdin(__obf_bc1041cfd5aedf7a, __obf_c6dc68184448cfb8, "java", __obf_e6cf0cd6cd858a8c(__obf_6370588d58c6ad08[0]))
}

func __obf_e6cf0cd6cd858a8c(__obf_b6e9177c6182e674 string) string {
	__obf_c3b1eec2c60f9890 := filepath.Ext(__obf_b6e9177c6182e674)
	return __obf_b6e9177c6182e674[0 : len(__obf_b6e9177c6182e674)-len(__obf_c3b1eec2c60f9890)]
}
