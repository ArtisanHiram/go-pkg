package __obf_ed0835c030fa3abe

import (
	cmd "github.com/ArtisanHiram/go-pkg/runner/cmd"
)

func Run(__obf_b63bb62124e92a16 string, __obf_92e0bb6a39e16a01 string, __obf_7569094cd101885a []string) (string, string, error) {
	return cmd.RunStdin(__obf_b63bb62124e92a16, __obf_92e0bb6a39e16a01, "node", __obf_7569094cd101885a[0])
}
