package __obf_dc577bef9ac5c6b8

import (
	"fmt"
	"io"
)

func (__obf_fdd7931bd5a61fac *HlsServer) GetFrame(__obf_1ff51571f6b70e49 string, __obf_90df74c8eee17939 int, __obf_39c0ed11835c4488 io.Writer) error {
	__obf_89adfe3713e6bab9 := []string{
		"-timelimit", "15",
		"-loglevel", "error",
		"-ss", fmt.Sprintf("%d.0", __obf_90df74c8eee17939),
		"-i", __obf_1ff51571f6b70e49,
		"-vf", "scale=320:-1",
		"-frames:v", "1",
		"-f", "image2",
		"-",
	}
	return __obf_fdd7931bd5a61fac.__obf_b69b54cbf0d766ab.Serve(FFMpegPath, __obf_89adfe3713e6bab9, __obf_39c0ed11835c4488)
}
