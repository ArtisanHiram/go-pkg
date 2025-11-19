package __obf_db3a2adcf9f700bf

import (
	cmd "github.com/ArtisanHiram/go-pkg/runner/cmd"
	"path/filepath"
)

func Run(__obf_511d5deb5f698e16, __obf_a7305547b06d118a string, __obf_a977baec74f93342 []string) (string, string, error) {
	__obf_f2026fc3cea3689a := "main.out"

	__obf_edc62d6b3e6637a6 := append([]string{"clang++", "-std=c++11", "-o", __obf_f2026fc3cea3689a}, __obf_a977baec74f93342...)
	__obf_dd797dcdc10320e0, __obf_025fc615fbfe3ca5, __obf_b99fcc8c31596230 := cmd.Run(__obf_511d5deb5f698e16, __obf_edc62d6b3e6637a6...)
	if __obf_b99fcc8c31596230 != nil {
		return __obf_dd797dcdc10320e0, __obf_025fc615fbfe3ca5, __obf_b99fcc8c31596230
	}

	__obf_b0ffe88338c9ab16 := filepath.Join(__obf_511d5deb5f698e16, __obf_f2026fc3cea3689a)
	return cmd.RunStdin(__obf_511d5deb5f698e16, __obf_a7305547b06d118a, __obf_b0ffe88338c9ab16)
}
