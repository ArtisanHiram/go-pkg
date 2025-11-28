package __obf_879483ff1f2bf10a

import (
	cmd "github.com/ArtisanHiram/go-pkg/runner/cmd"
	"path/filepath"
)

func Run(__obf_8cd7af4effc2c1de string, __obf_c18a7eb76d59d632 string, __obf_714850dc01d2f176 []string) (string, string, error) {
	__obf_e5c65dfc2d9d9094, __obf_28638a8c72859a6d, __obf_4dfa8941a20d4e77 := cmd.Run(__obf_8cd7af4effc2c1de, "javac", __obf_714850dc01d2f176[0])
	if __obf_4dfa8941a20d4e77 != nil {
		return __obf_e5c65dfc2d9d9094, __obf_28638a8c72859a6d, __obf_4dfa8941a20d4e77
	}
	return cmd.RunStdin(__obf_8cd7af4effc2c1de, __obf_c18a7eb76d59d632, "java", __obf_d85808c6e8e23f6b(__obf_714850dc01d2f176[0]))
}

func __obf_d85808c6e8e23f6b(__obf_333b0294d7599854 string) string {
	__obf_96512b00e6a9ee17 := filepath.Ext(__obf_333b0294d7599854)
	return __obf_333b0294d7599854[0 : len(__obf_333b0294d7599854)-len(__obf_96512b00e6a9ee17)]
}
