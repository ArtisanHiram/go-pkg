package __obf_da13203d7bc4eca3

import (
	cmd "github.com/ArtisanHiram/go-pkg/runner/cmd"
	"path/filepath"
)

func Run(__obf_bcfaa4cddd88a281 string, __obf_c90fb5a3b10cf797 string, __obf_2a61165cd6088b1a []string) (string, string, error) {
	__obf_294755a1b1513b5b, __obf_88673492f8e375c5, __obf_2b1e4c8073607094 := cmd.Run(__obf_bcfaa4cddd88a281, "javac", __obf_2a61165cd6088b1a[0])
	if __obf_2b1e4c8073607094 != nil {
		return __obf_294755a1b1513b5b, __obf_88673492f8e375c5, __obf_2b1e4c8073607094
	}
	return cmd.RunStdin(__obf_bcfaa4cddd88a281, __obf_c90fb5a3b10cf797, "java", __obf_64e46d86c8991ffe(__obf_2a61165cd6088b1a[0]))
}

func __obf_64e46d86c8991ffe(__obf_9b5af0726fc9d870 string) string {
	__obf_12448679b8061773 := filepath.Ext(__obf_9b5af0726fc9d870)
	return __obf_9b5af0726fc9d870[0 : len(__obf_9b5af0726fc9d870)-len(__obf_12448679b8061773)]
}
