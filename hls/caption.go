package __obf_90dd9b56c0f1bd65

import "io"

func (__obf_073dee7a39e9dbe6 *HlsServer) GetCaption(__obf_17a4cba300dc959c string, __obf_596cef7dccce398b io.Writer) error {
	__obf_0593c99d6b2a8391 := []string{
		"-i", __obf_17a4cba300dc959c, "-f", "webvtt",
		"pipe:",
		// 2> /dev/null
	}
	return __obf_073dee7a39e9dbe6.__obf_b62a01073bbe620f.Serve(FFMpegPath, __obf_0593c99d6b2a8391, __obf_596cef7dccce398b)
}
