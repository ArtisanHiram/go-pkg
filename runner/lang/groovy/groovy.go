package __obf_e4d9832e6fe11f4b

import (
	cmd "github.com/ArtisanHiram/go-pkg/runner/cmd"
)

func Run(__obf_cb3e93ffd7031935, __obf_317bc6a3a8afdad1 string, __obf_65fe968b8b34d70c []string) (string, string, error) {
	__obf_814e87bdd607e069 := append([]string{"groovy"}, __obf_65fe968b8b34d70c...)
	return cmd.RunStdin(__obf_cb3e93ffd7031935, __obf_317bc6a3a8afdad1, __obf_814e87bdd607e069...)
}
