package __obf_6f3adfdc12c473d2

import (
	cmd "github.com/ArtisanHiram/go-pkg/runner/cmd"
	"path/filepath"
)

func Run(__obf_144186fe738de92a string, __obf_d9067c94b767e714 string, __obf_d26f0a70bdf271d7 []string) (string, string, error) {
	__obf_d46447c86338f196, __obf_a6c729301135ece7, __obf_81912068d711da91 := cmd.Run(__obf_144186fe738de92a, "javac", __obf_d26f0a70bdf271d7[0])
	if __obf_81912068d711da91 != nil {
		return __obf_d46447c86338f196, __obf_a6c729301135ece7, __obf_81912068d711da91
	}
	return cmd.RunStdin(__obf_144186fe738de92a, __obf_d9067c94b767e714, "java", __obf_4700f6c931fb72ab(__obf_d26f0a70bdf271d7[0]))
}

func __obf_4700f6c931fb72ab(__obf_87d3fbcfa4b5e711 string) string {
	__obf_b31487b939968d9c := filepath.Ext(__obf_87d3fbcfa4b5e711)
	return __obf_87d3fbcfa4b5e711[0 : len(__obf_87d3fbcfa4b5e711)-len(__obf_b31487b939968d9c)]
}
