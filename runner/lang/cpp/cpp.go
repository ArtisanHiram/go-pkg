package __obf_9ef5792e02fb702c

import (
	cmd "github.com/ArtisanHiram/go-pkg/runner/cmd"
	"path/filepath"
)

func Run(__obf_d0e19a5d5755352c, __obf_2555c5fb1cda6634 string, __obf_f5875a526676451c []string) (string, string, error) {
	__obf_d21162670dfe9a77 := "main.out"
	__obf_f2d7e882156cfb5d := append([]string{"clang++", "-std=c++11", "-o", __obf_d21162670dfe9a77}, __obf_f5875a526676451c...)
	__obf_7d5d968a722ae6ff, __obf_0051875ddce84106, __obf_4b6ed245ea9e1d3c := cmd.Run(__obf_d0e19a5d5755352c, __obf_f2d7e882156cfb5d...)
	if __obf_4b6ed245ea9e1d3c != nil {
		return __obf_7d5d968a722ae6ff, __obf_0051875ddce84106, __obf_4b6ed245ea9e1d3c
	}
	__obf_942d596088787f6e := filepath.Join(__obf_d0e19a5d5755352c, __obf_d21162670dfe9a77)
	return cmd.RunStdin(__obf_d0e19a5d5755352c, __obf_2555c5fb1cda6634, __obf_942d596088787f6e)
}
