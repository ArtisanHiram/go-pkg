package __obf_5abac898006950bc

import (
	cmd "github.com/ArtisanHiram/go-pkg/runner/cmd"
)

func Run(__obf_61cdc3d64f3bd933, __obf_936fb58ea015fc72 string, __obf_f5fa54a94e317e31 []string) (string, string, error) {
	__obf_c29ed44ceb6dfcd8 := append([]string{"go", "run"}, __obf_f5fa54a94e317e31...)
	return cmd.RunStdin(__obf_61cdc3d64f3bd933, __obf_936fb58ea015fc72, __obf_c29ed44ceb6dfcd8...)
}
