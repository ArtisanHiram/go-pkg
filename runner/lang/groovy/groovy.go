package __obf_0a9f9cbc7f3154be

import (
	cmd "github.com/ArtisanHiram/go-pkg/runner/cmd"
)

func Run(__obf_91699819a7341b3e, __obf_0426cf05ece9e31a string, __obf_f46d56a8f190b1ef []string) (string, string, error) {
	__obf_f286217819c3249e := append([]string{"groovy"}, __obf_f46d56a8f190b1ef...)
	return cmd.RunStdin(__obf_91699819a7341b3e, __obf_0426cf05ece9e31a, __obf_f286217819c3249e...)
}
