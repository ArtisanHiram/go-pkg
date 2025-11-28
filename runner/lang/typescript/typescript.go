package __obf_1640d9f4f9c1ece6

import (
	cmd "github.com/ArtisanHiram/go-pkg/runner/cmd"
	util "github.com/ArtisanHiram/go-pkg/runner/util"
)

func Run(__obf_12720fa3ce77b28f string, __obf_0abadd8a3a3bba6f string, __obf_e43fa3221aade64e []string) (string, string, error) {
	__obf_fd800d81e940883f := "main.js"

	// Find all typescript files and build compile command
	__obf_e626d7e3b29019da := util.FilterByExtension(__obf_e43fa3221aade64e, ".ts")
	__obf_296a4288000c72b4 := append([]string{"tsc", "-out", __obf_fd800d81e940883f}, __obf_e626d7e3b29019da...)

	// Compile to javascript
	__obf_3e5cf4b599fe311e, __obf_503187aad500a7fa, __obf_87eb74a913058931 := cmd.Run(__obf_12720fa3ce77b28f, __obf_296a4288000c72b4...)
	if __obf_87eb74a913058931 != nil {
		return __obf_3e5cf4b599fe311e, __obf_503187aad500a7fa, __obf_87eb74a913058931
	}

	return cmd.RunStdin(__obf_12720fa3ce77b28f, __obf_0abadd8a3a3bba6f, "node", __obf_fd800d81e940883f)
}
