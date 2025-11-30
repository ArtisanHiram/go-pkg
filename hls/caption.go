package __obf_08ee9322ff6adb83

import "io"

func (__obf_504413fedee71524 *HlsServer) GetCaption(__obf_489cc4007c169b0b string, __obf_63e18a6991b3d10e io.Writer) error {
	__obf_b8bccd5d97f55f0d := []string{
		"-i", __obf_489cc4007c169b0b, "-f", "webvtt",
		"pipe:",
		// 2> /dev/null
	}
	return __obf_504413fedee71524.__obf_1a9d4d2fd79eae45.Serve(FFMpegPath, __obf_b8bccd5d97f55f0d, __obf_63e18a6991b3d10e)
}
