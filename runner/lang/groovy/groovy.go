package __obf_07822a398a143861

import (
	cmd "github.com/ArtisanHiram/go-pkg/runner/cmd"
)

func Run(__obf_89455f33c6bf0ddf, __obf_76734f3540a08606 string, __obf_963add3fb1f47624 []string) (string, string, error) {
	__obf_1af3129532003b80 := append([]string{"groovy"}, __obf_963add3fb1f47624...)
	return cmd.RunStdin(__obf_89455f33c6bf0ddf, __obf_76734f3540a08606, __obf_1af3129532003b80...)
}
