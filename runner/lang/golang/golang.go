package __obf_605b1d248e66b024

import (
	cmd "github.com/ArtisanHiram/go-pkg/runner/cmd"
)

func Run(__obf_d9e9b11cc4a8c6dc, __obf_3e51dfd6ebdbc0c5 string, __obf_32111076fa0352e8 []string) (string, string, error) {
	__obf_5c5dc3723599b438 := append([]string{"go", "run"}, __obf_32111076fa0352e8...)
	return cmd.RunStdin(__obf_d9e9b11cc4a8c6dc, __obf_3e51dfd6ebdbc0c5, __obf_5c5dc3723599b438...)
}
