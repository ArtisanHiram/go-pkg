package __obf_a1fca0dac8708121

import (
	cmd "github.com/ArtisanHiram/go-pkg/runner/cmd"
)

func Run(__obf_1fb7d3e7392415dd string, __obf_b52e073485d94949 string, __obf_5921b537eb12c1b8 []string) (string, string, error) {
	return cmd.RunStdin(__obf_1fb7d3e7392415dd, __obf_b52e073485d94949, "python3", __obf_5921b537eb12c1b8[0])
}
