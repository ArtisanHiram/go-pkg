package __obf_49bab97281532f92

import (
	cmd "github.com/ArtisanHiram/go-pkg/runner/cmd"
)

func Run(__obf_77541e014f0f77d3 string, __obf_806a97214c5d964b string, __obf_0a3c71645021f131 []string) (string, string, error) {
	return cmd.RunStdin(__obf_77541e014f0f77d3, __obf_806a97214c5d964b, "python3", __obf_0a3c71645021f131[0])
}
