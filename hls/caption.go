package __obf_b28324f38df50634

import "io"

func (__obf_2dc28af59fafc546 *HlsServer) GetCaption(__obf_a2e9fb1ef0c6fd12 string, __obf_eb3c523b7bd5ff61 io.Writer) error {
	__obf_afb5d9998846f346 := []string{
		"-i", __obf_a2e9fb1ef0c6fd12,
		"-f", "webvtt",
		"pipe:",
		// 2> /dev/null
	}
	return __obf_2dc28af59fafc546.__obf_31c94ec8882b996a.Serve(FFMpegPath, __obf_afb5d9998846f346, __obf_eb3c523b7bd5ff61)
}
