package __obf_eb5e805b9fc230e3

import (
	"fmt"
	"io"
)

func (__obf_f3546e71ce2e1d63 *HlsServer) GetFrame(__obf_6791b0a9cac91e8c string, __obf_c3d7443f5c072fc6 int, __obf_1cab0e86de8129e5 io.Writer) error {
	__obf_9ddc2c49f56f7ff3 := []string{
		"-timelimit", "15",
		"-loglevel", "error",
		"-ss", fmt.Sprintf("%d.0", __obf_c3d7443f5c072fc6),
		"-i", __obf_6791b0a9cac91e8c, "-vf", "scale=320:-1",
		"-frames:v", "1",
		"-f", "image2",
		"-",
	}
	return __obf_f3546e71ce2e1d63.__obf_5468dd2a54dba82c.Serve(FFMpegPath, __obf_9ddc2c49f56f7ff3, __obf_1cab0e86de8129e5)
}
