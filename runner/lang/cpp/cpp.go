package __obf_642e90b3713230d2

import (
	cmd "github.com/ArtisanHiram/go-pkg/runner/cmd"
	"path/filepath"
)

func Run(__obf_db2c3ce67996dcb2, __obf_5a2e477ab920de02 string, __obf_d0c8f701bbc69e2b []string) (string, string, error) {
	__obf_e79eda4bca09e6c0 := "main.out"
	__obf_89dbf2e599434918 := append([]string{"clang++", "-std=c++11", "-o", __obf_e79eda4bca09e6c0}, __obf_d0c8f701bbc69e2b...)
	__obf_8a4c383b1430dc97, __obf_515f19dc6a979a7f, __obf_f55909147cdfcf28 := cmd.Run(__obf_db2c3ce67996dcb2, __obf_89dbf2e599434918...)
	if __obf_f55909147cdfcf28 != nil {
		return __obf_8a4c383b1430dc97, __obf_515f19dc6a979a7f, __obf_f55909147cdfcf28
	}
	__obf_fbd51a75983960ed := filepath.Join(__obf_db2c3ce67996dcb2, __obf_e79eda4bca09e6c0)
	return cmd.RunStdin(__obf_db2c3ce67996dcb2, __obf_5a2e477ab920de02, __obf_fbd51a75983960ed)
}
