package __obf_363624866cfb4599

import (
	cmd "github.com/ArtisanHiram/go-pkg/runner/cmd"
)

func Run(__obf_e0cd4cf8fbef9d72 string, __obf_c890c298f87796d7 string, __obf_8d104d86f3cee6d9 []string) (string, string, error) {
	__obf_1e0618b53b89dc62 := "main.out"
	__obf_7a3bb5644ad976eb, __obf_2cda03a3f7ebc86e, __obf_8b4ca127d95d1269 := cmd.Run(__obf_e0cd4cf8fbef9d72, "rustc", "-o", __obf_1e0618b53b89dc62, __obf_8d104d86f3cee6d9[0])
	if __obf_8b4ca127d95d1269 != nil {
		return __obf_7a3bb5644ad976eb, __obf_2cda03a3f7ebc86e, __obf_8b4ca127d95d1269
	}

	return cmd.RunStdin(__obf_e0cd4cf8fbef9d72, __obf_c890c298f87796d7, __obf_1e0618b53b89dc62)
}
