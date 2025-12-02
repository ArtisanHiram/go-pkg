package __obf_2865d6713b09b7d6

import (
	cmd "github.com/ArtisanHiram/go-pkg/runner/cmd"
	util "github.com/ArtisanHiram/go-pkg/runner/util"
)

func Run(__obf_c2ecf70c95fa3d78 string, __obf_4fe667c523cacc7e string, __obf_e34019648ee0f294 []string) (string, string, error) {
	__obf_a8e297713abd8649 := "main.js"
	__obf_60c02651e3764b8e := // Find all typescript files and build compile command
		util.FilterByExtension(__obf_e34019648ee0f294, ".ts")
	__obf_3b388461473aeea8 := append([]string{"tsc", "-out", __obf_a8e297713abd8649}, __obf_60c02651e3764b8e...)
	__obf_944677c4c7abc76c,

		// Compile to javascript
		__obf_ea73d5a190a74f0c, __obf_665dbca6612ac353 := cmd.Run(__obf_c2ecf70c95fa3d78, __obf_3b388461473aeea8...)
	if __obf_665dbca6612ac353 != nil {
		return __obf_944677c4c7abc76c, __obf_ea73d5a190a74f0c, __obf_665dbca6612ac353
	}

	return cmd.RunStdin(__obf_c2ecf70c95fa3d78, __obf_4fe667c523cacc7e, "node", __obf_a8e297713abd8649)
}
