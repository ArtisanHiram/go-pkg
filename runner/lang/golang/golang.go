package __obf_b52ae7954b3bfd49

import (
	cmd "github.com/ArtisanHiram/go-pkg/runner/cmd"
)

func Run(__obf_e600f9efbe0d153a, __obf_abd1e1aaf1b2fd3c string, __obf_32bc4f10776da8d0 []string) (string, string, error) {
	__obf_cbff54dd37aca28d := append([]string{"go", "run"}, __obf_32bc4f10776da8d0...)
	return cmd.RunStdin(__obf_e600f9efbe0d153a, __obf_abd1e1aaf1b2fd3c, __obf_cbff54dd37aca28d...)
}
