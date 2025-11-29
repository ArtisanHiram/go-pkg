package __obf_fbadb677acc6b6c3

import (
	cmd "github.com/ArtisanHiram/go-pkg/runner/cmd"
	"path/filepath"
)

func Run(__obf_4b18e15edb1d769e string, __obf_42ddc04e5b1cd392 string, __obf_acdceca2dd3b0804 []string) (string, string, error) {
	__obf_3d268fbeded2785a, __obf_53b22c9374b5531d, __obf_540f3360b8383bda := cmd.Run(__obf_4b18e15edb1d769e, "javac", __obf_acdceca2dd3b0804[0])
	if __obf_540f3360b8383bda != nil {
		return __obf_3d268fbeded2785a, __obf_53b22c9374b5531d, __obf_540f3360b8383bda
	}
	return cmd.RunStdin(__obf_4b18e15edb1d769e, __obf_42ddc04e5b1cd392, "java", __obf_b78358bc7d63b095(__obf_acdceca2dd3b0804[0]))
}

func __obf_b78358bc7d63b095(__obf_8d308a1dca0c93c3 string) string {
	__obf_7a223684f3cbddf2 := filepath.Ext(__obf_8d308a1dca0c93c3)
	return __obf_8d308a1dca0c93c3[0 : len(__obf_8d308a1dca0c93c3)-len(__obf_7a223684f3cbddf2)]
}
