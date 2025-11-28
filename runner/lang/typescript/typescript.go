package __obf_f30c23e360560102

import (
	cmd "github.com/ArtisanHiram/go-pkg/runner/cmd"
	util "github.com/ArtisanHiram/go-pkg/runner/util"
)

func Run(__obf_22cb330a1b5b3ed1 string, __obf_8a15e7f2e7a47483 string, __obf_9662268cecc4b7be []string) (string, string, error) {
	__obf_40f3e73907ebfb3b := "main.js"

	// Find all typescript files and build compile command
	__obf_6f7ca938ceed3cc0 := util.FilterByExtension(__obf_9662268cecc4b7be, ".ts")
	__obf_168b0d2c3e823042 := append([]string{"tsc", "-out", __obf_40f3e73907ebfb3b}, __obf_6f7ca938ceed3cc0...)

	// Compile to javascript
	__obf_0d622ea4aff688ac, __obf_034d51eac984a419, __obf_2a1fd8d623e76e39 := cmd.Run(__obf_22cb330a1b5b3ed1, __obf_168b0d2c3e823042...)
	if __obf_2a1fd8d623e76e39 != nil {
		return __obf_0d622ea4aff688ac, __obf_034d51eac984a419, __obf_2a1fd8d623e76e39
	}

	return cmd.RunStdin(__obf_22cb330a1b5b3ed1, __obf_8a15e7f2e7a47483, "node", __obf_40f3e73907ebfb3b)
}
