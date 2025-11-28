package __obf_5028a0a829ddcdab

import "io"

func (__obf_883d14d3bb9e62d2 *HlsServer) GetCaption(__obf_8ddbcc776441ebf5 string, __obf_f0018e81daaf1fa9 io.Writer) error {
	__obf_76b5a3882b8f075c := []string{
		"-i", __obf_8ddbcc776441ebf5,
		"-f", "webvtt",
		"pipe:",
		// 2> /dev/null
	}
	return __obf_883d14d3bb9e62d2.__obf_3c22b11f4aa43eed.Serve(FFMpegPath, __obf_76b5a3882b8f075c, __obf_f0018e81daaf1fa9)
}
