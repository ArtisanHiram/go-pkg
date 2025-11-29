package __obf_6ff082bd539c7df0

import (
	"fmt"
	"io"
)

func (__obf_dd5f757557da2c8e *HlsServer) GetFrame(__obf_061397ddd747b96c string, __obf_12ed00a4173fc40d int, __obf_42061c65b3e92c28 io.Writer) error {
	__obf_d4c7dc8800d11aa1 := []string{
		"-timelimit", "15",
		"-loglevel", "error",
		"-ss", fmt.Sprintf("%d.0", __obf_12ed00a4173fc40d),
		"-i", __obf_061397ddd747b96c, "-vf", "scale=320:-1",
		"-frames:v", "1",
		"-f", "image2",
		"-",
	}
	return __obf_dd5f757557da2c8e.__obf_b27f3195132067e8.Serve(FFMpegPath, __obf_d4c7dc8800d11aa1, __obf_42061c65b3e92c28)
}
