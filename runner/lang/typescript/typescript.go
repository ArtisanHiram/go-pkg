package __obf_d05960dd24a0fbc2

import (
	cmd "github.com/ArtisanHiram/go-pkg/runner/cmd"
	util "github.com/ArtisanHiram/go-pkg/runner/util"
)

func Run(__obf_72ce4b19cd029b03 string, __obf_ddaad23dd45ddc41 string, __obf_8ceb5ce5931dbe7b []string) (string, string, error) {
	__obf_5d933cf75e638bb4 := "main.js"

	// Find all typescript files and build compile command
	__obf_e74f17eda9a1e549 := util.FilterByExtension(__obf_8ceb5ce5931dbe7b, ".ts")
	__obf_2c494392f8098298 := append([]string{"tsc", "-out", __obf_5d933cf75e638bb4}, __obf_e74f17eda9a1e549...)

	// Compile to javascript
	__obf_8dead4522623d9b5, __obf_15f97952eac2ba16, __obf_774cee2ce59cf2c1 := cmd.Run(__obf_72ce4b19cd029b03, __obf_2c494392f8098298...)
	if __obf_774cee2ce59cf2c1 != nil {
		return __obf_8dead4522623d9b5, __obf_15f97952eac2ba16, __obf_774cee2ce59cf2c1
	}

	return cmd.RunStdin(__obf_72ce4b19cd029b03, __obf_ddaad23dd45ddc41, "node", __obf_5d933cf75e638bb4)
}
