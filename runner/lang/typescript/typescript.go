package __obf_c0111000bb9b5937

import (
	cmd "github.com/ArtisanHiram/go-pkg/runner/cmd"
	util "github.com/ArtisanHiram/go-pkg/runner/util"
)

func Run(__obf_a714d32e4983afb1 string, __obf_655f76604e661c4b string, __obf_27712fd49cd6528c []string) (string, string, error) {
	__obf_27bb69b387cf1867 := "main.js"
	__obf_ece178964608ea08 := // Find all typescript files and build compile command
		util.FilterByExtension(__obf_27712fd49cd6528c, ".ts")
	__obf_9b55fa8232b76719 := append([]string{"tsc", "-out", __obf_27bb69b387cf1867}, __obf_ece178964608ea08...)
	__obf_ded1620f43371660,

		// Compile to javascript
		__obf_7d76cb037d2ff522, __obf_28b5f4a1ea14dd49 := cmd.Run(__obf_a714d32e4983afb1, __obf_9b55fa8232b76719...)
	if __obf_28b5f4a1ea14dd49 != nil {
		return __obf_ded1620f43371660, __obf_7d76cb037d2ff522, __obf_28b5f4a1ea14dd49
	}

	return cmd.RunStdin(__obf_a714d32e4983afb1, __obf_655f76604e661c4b, "node", __obf_27bb69b387cf1867)
}
