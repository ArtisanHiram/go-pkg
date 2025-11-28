package __obf_42bbcad92b7de1a8

import "io"

func (__obf_c7be2ffa97bb9914 *HlsServer) GetCaption(__obf_72ab747420e4c9ee string, __obf_39100c7ef6f93ec8 io.Writer) error {
	__obf_3f936d0e43e30024 := []string{
		"-i", __obf_72ab747420e4c9ee,
		"-f", "webvtt",
		"pipe:",
		// 2> /dev/null
	}
	return __obf_c7be2ffa97bb9914.__obf_6b3042483ee82bf2.Serve(FFMpegPath, __obf_3f936d0e43e30024, __obf_39100c7ef6f93ec8)
}
