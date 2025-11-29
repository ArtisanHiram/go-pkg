package __obf_5441fcd9a319cf59

import "io"

func (__obf_bf443cef12bfef60 *HlsServer) GetCaption(__obf_617a682795b315ad string, __obf_9e22338370faf798 io.Writer) error {
	__obf_e5483301bac834be := []string{
		"-i", __obf_617a682795b315ad, "-f", "webvtt",
		"pipe:",
		// 2> /dev/null
	}
	return __obf_bf443cef12bfef60.__obf_b8a98ccb6b0d29bb.Serve(FFMpegPath, __obf_e5483301bac834be, __obf_9e22338370faf798)
}
