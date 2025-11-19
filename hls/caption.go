package __obf_dc577bef9ac5c6b8

import "io"

func (__obf_fdd7931bd5a61fac *HlsServer) GetCaption(__obf_cf3312d5b91d1c43 string, __obf_39c0ed11835c4488 io.Writer) error {
	__obf_89adfe3713e6bab9 := []string{
		"-i", __obf_cf3312d5b91d1c43,
		"-f", "webvtt",
		"pipe:",
		// 2> /dev/null
	}
	return __obf_fdd7931bd5a61fac.__obf_67af47df49797c5b.Serve(FFMpegPath, __obf_89adfe3713e6bab9, __obf_39c0ed11835c4488)
}
