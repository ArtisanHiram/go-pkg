package __obf_65e2c15d0c5b533c

import (
	cmd "github.com/ArtisanHiram/go-pkg/runner/cmd"
)

func Run(__obf_b3f61a8e08b3492e, __obf_699591a0a46f8e74 string, __obf_b19fa761e39ffa4e []string) (string, string, error) {
	__obf_db4f64506645309b := append([]string{"go", "run"}, __obf_b19fa761e39ffa4e...)
	return cmd.RunStdin(__obf_b3f61a8e08b3492e, __obf_699591a0a46f8e74, __obf_db4f64506645309b...)
}
