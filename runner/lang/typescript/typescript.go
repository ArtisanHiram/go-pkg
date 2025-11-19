package __obf_8adfbf17478e1f89

import (
	cmd "github.com/ArtisanHiram/go-pkg/runner/cmd"
	util "github.com/ArtisanHiram/go-pkg/runner/util"
)

func Run(__obf_32fbe52d2dc6a92a string, __obf_438d87c4d2bbf204 string, __obf_b2bef1a7869447ca []string) (string, string, error) {
	__obf_94dc0170ba4f438f := "main.js"

	// Find all typescript files and build compile command
	__obf_aa662de27613a631 := util.FilterByExtension(__obf_b2bef1a7869447ca, ".ts")
	__obf_a59f330fa75aa032 := append([]string{"tsc", "-out", __obf_94dc0170ba4f438f}, __obf_aa662de27613a631...)

	// Compile to javascript
	__obf_5a9db0347ebd40a1, __obf_dc162b26552fdd98, __obf_fae34468ff1f4617 := cmd.Run(__obf_32fbe52d2dc6a92a, __obf_a59f330fa75aa032...)
	if __obf_fae34468ff1f4617 != nil {
		return __obf_5a9db0347ebd40a1, __obf_dc162b26552fdd98, __obf_fae34468ff1f4617
	}

	return cmd.RunStdin(__obf_32fbe52d2dc6a92a, __obf_438d87c4d2bbf204, "node", __obf_94dc0170ba4f438f)
}
