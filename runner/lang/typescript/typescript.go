package __obf_fbebc8f82a04b8b4

import (
	cmd "github.com/ArtisanHiram/go-pkg/runner/cmd"
	util "github.com/ArtisanHiram/go-pkg/runner/util"
)

func Run(__obf_2c95068e6d454a9f string, __obf_e6ca8c4d305b8666 string, __obf_843a4e5031e3a40f []string) (string, string, error) {
	__obf_37269134bddabd34 := "main.js"

	// Find all typescript files and build compile command
	__obf_796739feb9c77f52 := util.FilterByExtension(__obf_843a4e5031e3a40f, ".ts")
	__obf_5d503dd057495cbf := append([]string{"tsc", "-out", __obf_37269134bddabd34}, __obf_796739feb9c77f52...)

	// Compile to javascript
	__obf_159880849dbd6f90, __obf_96222683a03ff27c, __obf_e3d2f690f443f303 := cmd.Run(__obf_2c95068e6d454a9f, __obf_5d503dd057495cbf...)
	if __obf_e3d2f690f443f303 != nil {
		return __obf_159880849dbd6f90, __obf_96222683a03ff27c, __obf_e3d2f690f443f303
	}

	return cmd.RunStdin(__obf_2c95068e6d454a9f, __obf_e6ca8c4d305b8666, "node", __obf_37269134bddabd34)
}
