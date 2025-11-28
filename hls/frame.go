package __obf_acea4ab24a824c18

import (
	"fmt"
	"io"
)

func (__obf_c971f1d170edee24 *HlsServer) GetFrame(__obf_53ed20c636479772 string, __obf_637944048afe9262 int, __obf_7bfc44ad58c48031 io.Writer) error {
	__obf_4e7062996d563de2 := []string{
		"-timelimit", "15",
		"-loglevel", "error",
		"-ss", fmt.Sprintf("%d.0", __obf_637944048afe9262),
		"-i", __obf_53ed20c636479772,
		"-vf", "scale=320:-1",
		"-frames:v", "1",
		"-f", "image2",
		"-",
	}
	return __obf_c971f1d170edee24.__obf_32bdc8ca319d4ad7.Serve(FFMpegPath, __obf_4e7062996d563de2, __obf_7bfc44ad58c48031)
}
