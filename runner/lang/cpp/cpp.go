package __obf_af216a9d1b79ffcf

import (
	cmd "github.com/ArtisanHiram/go-pkg/runner/cmd"
	"path/filepath"
)

func Run(__obf_d73acaca99ed6a46, __obf_b05b2cecefb9bc88 string, __obf_c34f10e08893d849 []string) (string, string, error) {
	__obf_0963b339ea29c117 := "main.out"
	__obf_0790998721fc66b5 := append([]string{"clang++", "-std=c++11", "-o", __obf_0963b339ea29c117}, __obf_c34f10e08893d849...)
	__obf_0935020bc3c2ea4d, __obf_9a422f1f768368ac, __obf_ea3fb2e295df1556 := cmd.Run(__obf_d73acaca99ed6a46, __obf_0790998721fc66b5...)
	if __obf_ea3fb2e295df1556 != nil {
		return __obf_0935020bc3c2ea4d, __obf_9a422f1f768368ac, __obf_ea3fb2e295df1556
	}
	__obf_c428fa736f3038c9 := filepath.Join(__obf_d73acaca99ed6a46, __obf_0963b339ea29c117)
	return cmd.RunStdin(__obf_d73acaca99ed6a46, __obf_b05b2cecefb9bc88, __obf_c428fa736f3038c9)
}
