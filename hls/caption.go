package __obf_6ff082bd539c7df0

import "io"

func (__obf_dd5f757557da2c8e *HlsServer) GetCaption(__obf_f6e6713da7daf44f string, __obf_42061c65b3e92c28 io.Writer) error {
	__obf_d4c7dc8800d11aa1 := []string{
		"-i", __obf_f6e6713da7daf44f, "-f", "webvtt",
		"pipe:",
		// 2> /dev/null
	}
	return __obf_dd5f757557da2c8e.__obf_e76a5f4630004e4b.Serve(FFMpegPath, __obf_d4c7dc8800d11aa1, __obf_42061c65b3e92c28)
}
