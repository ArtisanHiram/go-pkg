package __obf_029d06f6ab41252a

import (
	cmd "github.com/ArtisanHiram/go-pkg/runner/cmd"
)

func Run(__obf_b563a9cfeb505602 string, __obf_95c2544cf8ec4d55 string, __obf_37ff05725717bbed []string) (string, string, error) {
	__obf_0f14ccf4c376862c := "main.out"
	__obf_95c79405a474e158, __obf_ae29591cd2135faf, __obf_583979ed8d20c435 := cmd.Run(__obf_b563a9cfeb505602, "rustc", "-o", __obf_0f14ccf4c376862c, __obf_37ff05725717bbed[0])
	if __obf_583979ed8d20c435 != nil {
		return __obf_95c79405a474e158, __obf_ae29591cd2135faf, __obf_583979ed8d20c435
	}

	return cmd.RunStdin(__obf_b563a9cfeb505602, __obf_95c2544cf8ec4d55, __obf_0f14ccf4c376862c)
}
