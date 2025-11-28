package __obf_acea4ab24a824c18

import "io"

func (__obf_c971f1d170edee24 *HlsServer) GetCaption(__obf_944727d4e31f21db string, __obf_7bfc44ad58c48031 io.Writer) error {
	__obf_4e7062996d563de2 := []string{
		"-i", __obf_944727d4e31f21db,
		"-f", "webvtt",
		"pipe:",
		// 2> /dev/null
	}
	return __obf_c971f1d170edee24.__obf_a72f121784d2d8ec.Serve(FFMpegPath, __obf_4e7062996d563de2, __obf_7bfc44ad58c48031)
}
