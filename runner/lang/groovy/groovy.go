package __obf_6b1160169beca6e6

import (
	cmd "github.com/ArtisanHiram/go-pkg/runner/cmd"
)

func Run(__obf_ea5e788fb367d618, __obf_9a4871f25df2716b string, __obf_a6e72077dd75b11d []string) (string, string, error) {
	__obf_9827858cadc12688 := append([]string{"groovy"}, __obf_a6e72077dd75b11d...)
	return cmd.RunStdin(__obf_ea5e788fb367d618, __obf_9a4871f25df2716b, __obf_9827858cadc12688...)
}
