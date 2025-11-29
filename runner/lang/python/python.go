package __obf_d655c4a63064d834

import (
	cmd "github.com/ArtisanHiram/go-pkg/runner/cmd"
)

func Run(__obf_45e9c2c965b3e1c0 string, __obf_9a03bfb6bd20c400 string, __obf_50b654b450792f28 []string) (string, string, error) {
	return cmd.RunStdin(__obf_45e9c2c965b3e1c0, __obf_9a03bfb6bd20c400, "python3", __obf_50b654b450792f28[0])
}
