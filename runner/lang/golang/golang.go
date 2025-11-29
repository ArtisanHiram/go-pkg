package __obf_514085d258049d1a

import (
	cmd "github.com/ArtisanHiram/go-pkg/runner/cmd"
)

func Run(__obf_6d398d61eee042a9, __obf_16e590b8bfc31f6f string, __obf_fa221867db25331b []string) (string, string, error) {
	__obf_a0aa1b73d9a26464 := append([]string{"go", "run"}, __obf_fa221867db25331b...)
	return cmd.RunStdin(__obf_6d398d61eee042a9, __obf_16e590b8bfc31f6f, __obf_a0aa1b73d9a26464...)
}
