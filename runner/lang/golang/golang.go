package __obf_35310100e89cec12

import (
	cmd "github.com/ArtisanHiram/go-pkg/runner/cmd"
)

func Run(__obf_bb0275cb504fdb15, __obf_30515a0f2c3ac024 string, __obf_a40422ae501088e6 []string) (string, string, error) {
	__obf_9b7c9b10e7a30dc6 := append([]string{"go", "run"}, __obf_a40422ae501088e6...)
	return cmd.RunStdin(__obf_bb0275cb504fdb15, __obf_30515a0f2c3ac024, __obf_9b7c9b10e7a30dc6...)
}
