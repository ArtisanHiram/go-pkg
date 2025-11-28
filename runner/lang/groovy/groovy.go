package __obf_6aa911f920c9856e

import (
	cmd "github.com/ArtisanHiram/go-pkg/runner/cmd"
)

func Run(__obf_72f60d417d94ffd0, __obf_dab0e26a7bfc28ee string, __obf_710d68e72a9976c2 []string) (string, string, error) {
	__obf_bb8c45c78ff92109 := append([]string{"groovy"}, __obf_710d68e72a9976c2...)
	return cmd.RunStdin(__obf_72f60d417d94ffd0, __obf_dab0e26a7bfc28ee, __obf_bb8c45c78ff92109...)
}
