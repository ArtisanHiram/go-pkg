package __obf_d9ce5e9192190e9d

import (
	cmd "github.com/ArtisanHiram/go-pkg/runner/cmd"
)

func Run(__obf_3837ff41e4a05438 string, __obf_c20fc048a1001537 string, __obf_e696e72d806e3c13 []string) (string, string, error) {
	return cmd.RunStdin(__obf_3837ff41e4a05438, __obf_c20fc048a1001537, "python3", __obf_e696e72d806e3c13[0])
}
