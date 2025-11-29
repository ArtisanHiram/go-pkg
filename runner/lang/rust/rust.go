package __obf_6481da4aace47a7f

import (
	cmd "github.com/ArtisanHiram/go-pkg/runner/cmd"
)

func Run(__obf_aba7920988160838 string, __obf_ee28429064a3bd41 string, __obf_533a224507f340eb []string) (string, string, error) {
	__obf_2064011de3b1113f := "main.out"
	__obf_3f491d1ae246d513, __obf_3016679d5d009cf9, __obf_8612c668fa0c30d4 := cmd.Run(__obf_aba7920988160838, "rustc", "-o", __obf_2064011de3b1113f, __obf_533a224507f340eb[0])
	if __obf_8612c668fa0c30d4 != nil {
		return __obf_3f491d1ae246d513, __obf_3016679d5d009cf9, __obf_8612c668fa0c30d4
	}

	return cmd.RunStdin(__obf_aba7920988160838, __obf_ee28429064a3bd41, __obf_2064011de3b1113f)
}
