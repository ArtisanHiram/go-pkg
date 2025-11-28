package __obf_b28324f38df50634

import (
	"context"
	"fmt"
	"os/exec"
	"time"
)

type MediaInfo struct {
	Size     string `json:"size"`
	BitRate  string `json:"bit_rate"`
	Duration string `json:"duration"`
}

func GetVideoInfo(__obf_ef284eab1db4670a string) (*MediaInfo, error) {
	__obf_c2c8a52eb04b2873, __obf_f6957728c129bb63 := context.WithTimeout(context.Background(), 20*time.Second)
	defer __obf_f6957728c129bb63()
	var __obf_243ecc1b663ec613 struct {
		Format MediaInfo
	}
	if __obf_a4e073529832bad2 := ExecJsonStdout(exec.CommandContext(__obf_c2c8a52eb04b2873, FFProbePath, "-v", "quiet", "-print_format", "json", "-show_format", __obf_ef284eab1db4670a), &__obf_243ecc1b663ec613); __obf_a4e073529832bad2 != nil {
		return nil, fmt.Errorf("error getting JSON from ffprobe output for file '%v': %v", __obf_ef284eab1db4670a, __obf_a4e073529832bad2)
	}

	return &__obf_243ecc1b663ec613.Format, nil
}
