package __obf_acb51b51b27255f1

import (
	cmd "github.com/ArtisanHiram/go-pkg/runner/cmd"
	"path/filepath"
)

func Run(__obf_472a416e278f6f30 string, __obf_3d475030265ccde7 string, __obf_b08b784032285e4d []string) (string, string, error) {
	__obf_d0b66934a2c4ea7b, __obf_abec7934c7637779, __obf_28c7a9c47b64686a := cmd.Run(__obf_472a416e278f6f30, "javac", __obf_b08b784032285e4d[0])
	if __obf_28c7a9c47b64686a != nil {
		return __obf_d0b66934a2c4ea7b, __obf_abec7934c7637779, __obf_28c7a9c47b64686a
	}
	return cmd.RunStdin(__obf_472a416e278f6f30, __obf_3d475030265ccde7, "java", __obf_696d8d8b1070e27e(__obf_b08b784032285e4d[0]))
}

func __obf_696d8d8b1070e27e(__obf_e4184b60ae1c1c45 string) string {
	__obf_2ad23dd359594ee7 := filepath.Ext(__obf_e4184b60ae1c1c45)
	return __obf_e4184b60ae1c1c45[0 : len(__obf_e4184b60ae1c1c45)-len(__obf_2ad23dd359594ee7)]
}
