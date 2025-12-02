package __obf_eb5e805b9fc230e3

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

func GetVideoInfo(__obf_bd9b77e582496812 string) (*MediaInfo, error) {
	__obf_5f0e28df6b1ffd43, __obf_ac08eb4ed2a9a55f := context.WithTimeout(context.Background(), 20*time.Second)
	defer __obf_ac08eb4ed2a9a55f()
	var __obf_abc498cda9847fed struct {
		Format MediaInfo
	}
	if __obf_2d43222a56faebee := ExecJsonStdout(exec.CommandContext(__obf_5f0e28df6b1ffd43, FFProbePath, "-v", "quiet", "-print_format", "json", "-show_format", __obf_bd9b77e582496812), &__obf_abc498cda9847fed); __obf_2d43222a56faebee != nil {
		return nil, fmt.Errorf("error getting JSON from ffprobe output for file '%v': %v", __obf_bd9b77e582496812, __obf_2d43222a56faebee)
	}

	return &__obf_abc498cda9847fed.Format, nil
}
