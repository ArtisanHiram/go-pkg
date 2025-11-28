package __obf_238061ef2d2dcce0

import (
	cmd "github.com/ArtisanHiram/go-pkg/runner/cmd"
)

func Run(__obf_b5558c95062818ee string, __obf_b7a1ec43071aa14f string, __obf_9a0571a1e637e4a3 []string) (string, string, error) {
	return cmd.RunStdin(__obf_b5558c95062818ee, __obf_b7a1ec43071aa14f, "python3", __obf_9a0571a1e637e4a3[0])
}
