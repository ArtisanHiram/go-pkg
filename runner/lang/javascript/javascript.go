package __obf_8b037fe4424e2dea

import (
	cmd "github.com/ArtisanHiram/go-pkg/runner/cmd"
)

func Run(__obf_71177cc34090925e string, __obf_5341ef77ac40121b string, __obf_7c52d4b25af31f6e []string) (string, string, error) {
	return cmd.RunStdin(__obf_71177cc34090925e, __obf_5341ef77ac40121b, "node", __obf_7c52d4b25af31f6e[0])
}
