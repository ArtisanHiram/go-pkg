package __obf_596ef75d907b553a

import (
	cmd "github.com/ArtisanHiram/go-pkg/runner/cmd"
)

func Run(__obf_4c0175c4de199457, __obf_2e32bc922cb9bfc5 string, __obf_c62be18179217130 []string) (string, string, error) {
	__obf_d5df21027b89e577 := append([]string{"go", "run"}, __obf_c62be18179217130...)
	return cmd.RunStdin(__obf_4c0175c4de199457, __obf_2e32bc922cb9bfc5, __obf_d5df21027b89e577...)
}
