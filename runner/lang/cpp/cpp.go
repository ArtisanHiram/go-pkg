package __obf_157eadd90452810c

import (
	cmd "github.com/ArtisanHiram/go-pkg/runner/cmd"
	"path/filepath"
)

func Run(__obf_245493ca07772d0f, __obf_a97cc330aba12e88 string, __obf_d5b6bd7118b55229 []string) (string, string, error) {
	__obf_1877666572b563b5 := "main.out"

	__obf_8353a79aafe540e0 := append([]string{"clang++", "-std=c++11", "-o", __obf_1877666572b563b5}, __obf_d5b6bd7118b55229...)
	__obf_eaad03328cc643cf, __obf_771045ad20cc85a7, __obf_924a5e519bda93a0 := cmd.Run(__obf_245493ca07772d0f, __obf_8353a79aafe540e0...)
	if __obf_924a5e519bda93a0 != nil {
		return __obf_eaad03328cc643cf, __obf_771045ad20cc85a7, __obf_924a5e519bda93a0
	}

	__obf_857c904636f85dc1 := filepath.Join(__obf_245493ca07772d0f, __obf_1877666572b563b5)
	return cmd.RunStdin(__obf_245493ca07772d0f, __obf_a97cc330aba12e88, __obf_857c904636f85dc1)
}
