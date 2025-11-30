package __obf_0667f663e27c2f92

import (
	cmd "github.com/ArtisanHiram/go-pkg/runner/cmd"
	"path/filepath"
)

func Run(__obf_7691c57473ff1fc4 string, __obf_9142b19c5f0a61e6 string, __obf_5e3f56f4ae5b6f0d []string) (string, string, error) {
	__obf_8b97d4c3eef3ab03, __obf_722d6c94b7dae6e6, __obf_9baddc422e3d18ba := cmd.Run(__obf_7691c57473ff1fc4, "javac", __obf_5e3f56f4ae5b6f0d[0])
	if __obf_9baddc422e3d18ba != nil {
		return __obf_8b97d4c3eef3ab03, __obf_722d6c94b7dae6e6, __obf_9baddc422e3d18ba
	}
	return cmd.RunStdin(__obf_7691c57473ff1fc4, __obf_9142b19c5f0a61e6, "java", __obf_d05461d73bb112f5(__obf_5e3f56f4ae5b6f0d[0]))
}

func __obf_d05461d73bb112f5(__obf_73929f6c587215aa string) string {
	__obf_0be3ba5dd376a61c := filepath.Ext(__obf_73929f6c587215aa)
	return __obf_73929f6c587215aa[0 : len(__obf_73929f6c587215aa)-len(__obf_0be3ba5dd376a61c)]
}
