package __obf_42bbcad92b7de1a8

import (
	"fmt"
	"io"
)

func (__obf_c7be2ffa97bb9914 *HlsServer) GetFrame(__obf_91124693445fffe2 string, __obf_0e23a453545810e5 int, __obf_39100c7ef6f93ec8 io.Writer) error {
	__obf_3f936d0e43e30024 := []string{
		"-timelimit", "15",
		"-loglevel", "error",
		"-ss", fmt.Sprintf("%d.0", __obf_0e23a453545810e5),
		"-i", __obf_91124693445fffe2,
		"-vf", "scale=320:-1",
		"-frames:v", "1",
		"-f", "image2",
		"-",
	}
	return __obf_c7be2ffa97bb9914.__obf_6b0feca1a08b9797.Serve(FFMpegPath, __obf_3f936d0e43e30024, __obf_39100c7ef6f93ec8)
}
