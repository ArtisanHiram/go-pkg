package __obf_f60326fd90eb13d9

import (
	"fmt"
	"io"
)

func (__obf_dd04e55ec17d0ebf *HlsServer) GetFrame(__obf_da9c329826f2d541 string, __obf_981ad1a0af9b831b int, __obf_811b158c28965ee0 io.Writer) error {
	__obf_4790050703cb3251 := []string{
		"-timelimit", "15",
		"-loglevel", "error",
		"-ss", fmt.Sprintf("%d.0", __obf_981ad1a0af9b831b),
		"-i", __obf_da9c329826f2d541, "-vf", "scale=320:-1",
		"-frames:v", "1",
		"-f", "image2",
		"-",
	}
	return __obf_dd04e55ec17d0ebf.__obf_e9d8d0cc62a9ece3.Serve(FFMpegPath, __obf_4790050703cb3251, __obf_811b158c28965ee0)
}
