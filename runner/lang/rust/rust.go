package __obf_e5a7cf719760b291

import (
	cmd "github.com/ArtisanHiram/go-pkg/runner/cmd"
)

func Run(__obf_49fbbe8b8d27b732 string, __obf_1f1623f5a2133b93 string, __obf_e1310b6b51cc4cf2 []string) (string, string, error) {
	__obf_38629db6ed186c65 := "main.out"
	__obf_158c15e6572cb4ea, __obf_eabe52de728fdc6d, __obf_66431a73e5a91d79 := cmd.Run(__obf_49fbbe8b8d27b732, "rustc", "-o", __obf_38629db6ed186c65, __obf_e1310b6b51cc4cf2[0])
	if __obf_66431a73e5a91d79 != nil {
		return __obf_158c15e6572cb4ea, __obf_eabe52de728fdc6d, __obf_66431a73e5a91d79
	}

	return cmd.RunStdin(__obf_49fbbe8b8d27b732, __obf_1f1623f5a2133b93, __obf_38629db6ed186c65)
}
