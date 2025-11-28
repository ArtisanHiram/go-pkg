package __obf_7e08567a4bdc9954

import (
	cmd "github.com/ArtisanHiram/go-pkg/runner/cmd"
	"path/filepath"
)

func Run(__obf_4d5a7fadb0534422 string, __obf_f950cdbe5ddf115d string, __obf_9f1c5ef2b51ab4b5 []string) (string, string, error) {
	__obf_1d3193fb19efa5f0, __obf_c36e9163431b68d9, __obf_9801fe7fa602ac03 := cmd.Run(__obf_4d5a7fadb0534422, "javac", __obf_9f1c5ef2b51ab4b5[0])
	if __obf_9801fe7fa602ac03 != nil {
		return __obf_1d3193fb19efa5f0, __obf_c36e9163431b68d9, __obf_9801fe7fa602ac03
	}
	return cmd.RunStdin(__obf_4d5a7fadb0534422, __obf_f950cdbe5ddf115d, "java", __obf_53a7ce93732544dd(__obf_9f1c5ef2b51ab4b5[0]))
}

func __obf_53a7ce93732544dd(__obf_eadecbdbce261e73 string) string {
	__obf_e5813e7067ead3c9 := filepath.Ext(__obf_eadecbdbce261e73)
	return __obf_eadecbdbce261e73[0 : len(__obf_eadecbdbce261e73)-len(__obf_e5813e7067ead3c9)]
}
