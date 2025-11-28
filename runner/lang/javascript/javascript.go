package __obf_54fa1ddb10038ac2

import (
	cmd "github.com/ArtisanHiram/go-pkg/runner/cmd"
)

func Run(__obf_28bcb607ee417396 string, __obf_4cd941745525199c string, __obf_645829dfd443d4d2 []string) (string, string, error) {
	return cmd.RunStdin(__obf_28bcb607ee417396, __obf_4cd941745525199c, "node", __obf_645829dfd443d4d2[0])
}
