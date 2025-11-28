package __obf_082176c85a9c6f78

import (
	cmd "github.com/ArtisanHiram/go-pkg/runner/cmd"
	"path/filepath"
)

func Run(__obf_250976a1b2b5dbdc, __obf_f6e63d5c8d1cd56f string, __obf_26ab64cce8fdf1b5 []string) (string, string, error) {
	__obf_9766bb253d3f7f43 := "main.out"

	__obf_fad74d521b62f2fa := append([]string{"clang++", "-std=c++11", "-o", __obf_9766bb253d3f7f43}, __obf_26ab64cce8fdf1b5...)
	__obf_13dc6c439c145b2b, __obf_dc19227d0de5a8e3, __obf_9af89a3a45080f2b := cmd.Run(__obf_250976a1b2b5dbdc, __obf_fad74d521b62f2fa...)
	if __obf_9af89a3a45080f2b != nil {
		return __obf_13dc6c439c145b2b, __obf_dc19227d0de5a8e3, __obf_9af89a3a45080f2b
	}

	__obf_deed8d81daf9f96d := filepath.Join(__obf_250976a1b2b5dbdc, __obf_9766bb253d3f7f43)
	return cmd.RunStdin(__obf_250976a1b2b5dbdc, __obf_f6e63d5c8d1cd56f, __obf_deed8d81daf9f96d)
}
