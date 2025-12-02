package __obf_eb5e805b9fc230e3

import "io"

func (__obf_f3546e71ce2e1d63 *HlsServer) GetCaption(__obf_22d6c43ac51329e5 string, __obf_1cab0e86de8129e5 io.Writer) error {
	__obf_9ddc2c49f56f7ff3 := []string{
		"-i", __obf_22d6c43ac51329e5, "-f", "webvtt",
		"pipe:",
		// 2> /dev/null
	}
	return __obf_f3546e71ce2e1d63.__obf_45e21b72f0a6ee88.Serve(FFMpegPath, __obf_9ddc2c49f56f7ff3, __obf_1cab0e86de8129e5)
}
