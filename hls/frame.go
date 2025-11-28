package __obf_5028a0a829ddcdab

import (
	"fmt"
	"io"
)

func (__obf_883d14d3bb9e62d2 *HlsServer) GetFrame(__obf_1d59566cc91271b2 string, __obf_b5d1f2c9b628e0a5 int, __obf_f0018e81daaf1fa9 io.Writer) error {
	__obf_76b5a3882b8f075c := []string{
		"-timelimit", "15",
		"-loglevel", "error",
		"-ss", fmt.Sprintf("%d.0", __obf_b5d1f2c9b628e0a5),
		"-i", __obf_1d59566cc91271b2,
		"-vf", "scale=320:-1",
		"-frames:v", "1",
		"-f", "image2",
		"-",
	}
	return __obf_883d14d3bb9e62d2.__obf_3c1630e7a0e9c371.Serve(FFMpegPath, __obf_76b5a3882b8f075c, __obf_f0018e81daaf1fa9)
}
