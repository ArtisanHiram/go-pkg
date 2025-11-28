package __obf_b28324f38df50634

import (
	"fmt"
	"io"
)

func (__obf_2dc28af59fafc546 *HlsServer) GetFrame(__obf_bae34d60e029b6cf string, __obf_e59ef9f50c961025 int, __obf_eb3c523b7bd5ff61 io.Writer) error {
	__obf_afb5d9998846f346 := []string{
		"-timelimit", "15",
		"-loglevel", "error",
		"-ss", fmt.Sprintf("%d.0", __obf_e59ef9f50c961025),
		"-i", __obf_bae34d60e029b6cf,
		"-vf", "scale=320:-1",
		"-frames:v", "1",
		"-f", "image2",
		"-",
	}
	return __obf_2dc28af59fafc546.__obf_25e9ffdd6cfa6045.Serve(FFMpegPath, __obf_afb5d9998846f346, __obf_eb3c523b7bd5ff61)
}
