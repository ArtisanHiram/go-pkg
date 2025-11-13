package __obf_a20e49a16ecddeac

import (
	cmd "github.com/ArtisanHiram/go-pkg/runner/cmd"
)

func Run(__obf_ccca5ef2c3543522, __obf_03800fb26598a579 string, __obf_806b231d344a0ce5 []string) (string, string, error) {
	__obf_3e2e2d456a4c2774 := append([]string{"go", "run"}, __obf_806b231d344a0ce5...)
	return cmd.RunStdin(__obf_ccca5ef2c3543522, __obf_03800fb26598a579, __obf_3e2e2d456a4c2774...)
}
