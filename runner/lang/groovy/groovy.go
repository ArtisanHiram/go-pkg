package __obf_8878052205012bf2

import (
	cmd "github.com/ArtisanHiram/go-pkg/runner/cmd"
)

func Run(__obf_17757802ecb0c677, __obf_575cc3a022b091de string, __obf_96cdbc08136abb57 []string) (string, string, error) {
	__obf_209dc0003ae19154 := append([]string{"groovy"}, __obf_96cdbc08136abb57...)
	return cmd.RunStdin(__obf_17757802ecb0c677, __obf_575cc3a022b091de, __obf_209dc0003ae19154...)
}
