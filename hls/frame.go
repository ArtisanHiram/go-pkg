package __obf_90dd9b56c0f1bd65

import (
	"fmt"
	"io"
)

func (__obf_073dee7a39e9dbe6 *HlsServer) GetFrame(__obf_3056a50ff16e96dc string, __obf_eb5e480614fe5a41 int, __obf_596cef7dccce398b io.Writer) error {
	__obf_0593c99d6b2a8391 := []string{
		"-timelimit", "15",
		"-loglevel", "error",
		"-ss", fmt.Sprintf("%d.0", __obf_eb5e480614fe5a41),
		"-i", __obf_3056a50ff16e96dc, "-vf", "scale=320:-1",
		"-frames:v", "1",
		"-f", "image2",
		"-",
	}
	return __obf_073dee7a39e9dbe6.__obf_b2fb5b3625c77a34.Serve(FFMpegPath, __obf_0593c99d6b2a8391, __obf_596cef7dccce398b)
}
