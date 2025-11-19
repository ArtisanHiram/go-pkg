package __obf_9b24f1d4f0205e14

import (
	cmd "github.com/ArtisanHiram/go-pkg/runner/cmd"
)

func Run(__obf_d8cdc3cfbaf14bf6 string, __obf_4ac07c10a2e90f49 string, __obf_dbef64b08ac691c2 []string) (string, string, error) {
	return cmd.RunStdin(__obf_d8cdc3cfbaf14bf6, __obf_4ac07c10a2e90f49, "python3", __obf_dbef64b08ac691c2[0])
}
