package __obf_5441fcd9a319cf59

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

func GetVideoInfo(__obf_75bf03e05fea1b50 string) (*MediaInfo, error) {
	__obf_83fad17274e83f8e, __obf_79bf09d29303d0f4 := context.WithTimeout(context.Background(), 20*time.Second)
	defer __obf_79bf09d29303d0f4()
	var __obf_3753bb63957fd7cb struct {
		Format MediaInfo
	}
	if __obf_e8125dea727cd4c9 := ExecJsonStdout(exec.CommandContext(__obf_83fad17274e83f8e, FFProbePath, "-v", "quiet", "-print_format", "json", "-show_format", __obf_75bf03e05fea1b50), &__obf_3753bb63957fd7cb); __obf_e8125dea727cd4c9 != nil {
		return nil, fmt.Errorf("error getting JSON from ffprobe output for file '%v': %v", __obf_75bf03e05fea1b50, __obf_e8125dea727cd4c9)
	}

	return &__obf_3753bb63957fd7cb.Format, nil
}
