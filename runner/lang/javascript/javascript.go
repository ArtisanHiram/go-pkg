package __obf_70dcaf7af4813a40

import (
	cmd "github.com/ArtisanHiram/go-pkg/runner/cmd"
)

func Run(__obf_88e64378b980cc8a string, __obf_76b003cf9974a6ff string, __obf_a96948906f2fc5ba []string) (string, string, error) {
	return cmd.RunStdin(__obf_88e64378b980cc8a, __obf_76b003cf9974a6ff, "node", __obf_a96948906f2fc5ba[0])
}
