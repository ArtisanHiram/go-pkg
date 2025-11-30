package __obf_fae153e94f216c92

import (
	cmd "github.com/ArtisanHiram/go-pkg/runner/cmd"
	util "github.com/ArtisanHiram/go-pkg/runner/util"
)

func Run(__obf_b1fa45357c1a4360 string, __obf_2c89368ee9395f72 string, __obf_a1c6aaac0dbd4587 []string) (string, string, error) {
	__obf_418169c9fc121690 := "main.js"
	__obf_d96f870f3a0b53ba := // Find all typescript files and build compile command
		util.FilterByExtension(__obf_a1c6aaac0dbd4587, ".ts")
	__obf_1e8e1df8fb93395a := append([]string{"tsc", "-out", __obf_418169c9fc121690}, __obf_d96f870f3a0b53ba...)
	__obf_0e2a6b27ed0d8ac2,

		// Compile to javascript
		__obf_d6b15c39bcb038de, __obf_0f55de9608787e31 := cmd.Run(__obf_b1fa45357c1a4360, __obf_1e8e1df8fb93395a...)
	if __obf_0f55de9608787e31 != nil {
		return __obf_0e2a6b27ed0d8ac2, __obf_d6b15c39bcb038de, __obf_0f55de9608787e31
	}

	return cmd.RunStdin(__obf_b1fa45357c1a4360, __obf_2c89368ee9395f72, "node", __obf_418169c9fc121690)
}
