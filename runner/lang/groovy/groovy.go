package __obf_df686adf42341f37

import (
	cmd "github.com/ArtisanHiram/go-pkg/runner/cmd"
)

func Run(__obf_684b6d1e100dd45c, __obf_60318c8dd38aae2f string, __obf_766b79931eb9fb18 []string) (string, string, error) {
	__obf_272ff09124d27fa1 := append([]string{"groovy"}, __obf_766b79931eb9fb18...)
	return cmd.RunStdin(__obf_684b6d1e100dd45c, __obf_60318c8dd38aae2f, __obf_272ff09124d27fa1...)
}
