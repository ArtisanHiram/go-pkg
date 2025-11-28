package __obf_2ccff9ea885050de

import (
	cmd "github.com/ArtisanHiram/go-pkg/runner/cmd"
	"path/filepath"
)

func Run(__obf_f894335313a013eb string, __obf_608500b37fa223aa string, __obf_d96ab7cba5e6fce0 []string) (string, string, error) {
	__obf_acd4f27631132f59, __obf_c7463b59df243792, __obf_56b1bcfb48b48ef7 := cmd.Run(__obf_f894335313a013eb, "javac", __obf_d96ab7cba5e6fce0[0])
	if __obf_56b1bcfb48b48ef7 != nil {
		return __obf_acd4f27631132f59, __obf_c7463b59df243792, __obf_56b1bcfb48b48ef7
	}
	return cmd.RunStdin(__obf_f894335313a013eb, __obf_608500b37fa223aa, "java", __obf_58fd77fde630f0d4(__obf_d96ab7cba5e6fce0[0]))
}

func __obf_58fd77fde630f0d4(__obf_7d3bfec93ca81fb9 string) string {
	__obf_175d3de7c5fec3eb := filepath.Ext(__obf_7d3bfec93ca81fb9)
	return __obf_7d3bfec93ca81fb9[0 : len(__obf_7d3bfec93ca81fb9)-len(__obf_175d3de7c5fec3eb)]
}
