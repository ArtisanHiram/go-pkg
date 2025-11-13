package __obf_ab25ed2437cd567a

import "io"

func (__obf_8d3b3856596fc4f4 *HlsServer) GetCaption(__obf_56707f240662d9aa string, __obf_98c2f72bee615664 io.Writer) error {
	__obf_3a7fd6afab4923e8 := []string{
		"-i", __obf_56707f240662d9aa,
		"-f", "webvtt",
		"pipe:",
		// 2> /dev/null
	}
	return __obf_8d3b3856596fc4f4.__obf_57de09c5d2d0d12d.Serve(FFMpegPath, __obf_3a7fd6afab4923e8, __obf_98c2f72bee615664)
}
