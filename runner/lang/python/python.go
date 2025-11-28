package __obf_51e84c3eaf09b10b

import (
	cmd "github.com/ArtisanHiram/go-pkg/runner/cmd"
)

func Run(__obf_686a6eb904f024c1 string, __obf_fc04b12fa546ea88 string, __obf_d72e9ff6ed182a0c []string) (string, string, error) {
	return cmd.RunStdin(__obf_686a6eb904f024c1, __obf_fc04b12fa546ea88, "python3", __obf_d72e9ff6ed182a0c[0])
}
