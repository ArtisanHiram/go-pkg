package __obf_08ee9322ff6adb83

import (
	"fmt"
	"io"
)

func (__obf_504413fedee71524 *HlsServer) GetFrame(__obf_5cbee58b9dbb3e28 string, __obf_a7a6e8f733cab444 int, __obf_63e18a6991b3d10e io.Writer) error {
	__obf_b8bccd5d97f55f0d := []string{
		"-timelimit", "15",
		"-loglevel", "error",
		"-ss", fmt.Sprintf("%d.0", __obf_a7a6e8f733cab444),
		"-i", __obf_5cbee58b9dbb3e28, "-vf", "scale=320:-1",
		"-frames:v", "1",
		"-f", "image2",
		"-",
	}
	return __obf_504413fedee71524.__obf_68bd4c2dc42e505f.Serve(FFMpegPath, __obf_b8bccd5d97f55f0d, __obf_63e18a6991b3d10e)
}
