package __obf_ab25ed2437cd567a

import (
	"fmt"
	"io"
)

func (__obf_8d3b3856596fc4f4 *HlsServer) GetFrame(__obf_a09859f1d205ee4d string, __obf_cfc6c82f5436ba5e int, __obf_98c2f72bee615664 io.Writer) error {
	__obf_3a7fd6afab4923e8 := []string{
		"-timelimit", "15",
		"-loglevel", "error",
		"-ss", fmt.Sprintf("%d.0", __obf_cfc6c82f5436ba5e),
		"-i", __obf_a09859f1d205ee4d,
		"-vf", "scale=320:-1",
		"-frames:v", "1",
		"-f", "image2",
		"-",
	}
	return __obf_8d3b3856596fc4f4.__obf_4f68dab75aa13315.Serve(FFMpegPath, __obf_3a7fd6afab4923e8, __obf_98c2f72bee615664)
}
