package __obf_904c6f8cb89fd69c

import (
	cmd "github.com/ArtisanHiram/go-pkg/runner/cmd"
)

func Run(__obf_0e3df79b21f0a183, __obf_960a444e82663488 string, __obf_62e26959de6bd799 []string) (string, string, error) {
	__obf_820537d2789ce2cf := append([]string{"go", "run"}, __obf_62e26959de6bd799...)
	return cmd.RunStdin(__obf_0e3df79b21f0a183, __obf_960a444e82663488, __obf_820537d2789ce2cf...)
}
