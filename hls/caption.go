package __obf_f60326fd90eb13d9

import "io"

func (__obf_dd04e55ec17d0ebf *HlsServer) GetCaption(__obf_2c7a65b2dac56e8a string, __obf_811b158c28965ee0 io.Writer) error {
	__obf_4790050703cb3251 := []string{
		"-i", __obf_2c7a65b2dac56e8a, "-f", "webvtt",
		"pipe:",
		// 2> /dev/null
	}
	return __obf_dd04e55ec17d0ebf.__obf_9b2fbfd16e499a01.Serve(FFMpegPath, __obf_4790050703cb3251, __obf_811b158c28965ee0)
}
