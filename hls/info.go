package __obf_acea4ab24a824c18

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

func GetVideoInfo(__obf_f2f454838e7353a1 string) (*MediaInfo, error) {
	__obf_efbbc02bedef0e2d, __obf_0e915d14e5e56b99 := context.WithTimeout(context.Background(), 20*time.Second)
	defer __obf_0e915d14e5e56b99()
	var __obf_17e491a79e2a4b27 struct {
		Format MediaInfo
	}
	if __obf_3698cbf06506c070 := ExecJsonStdout(exec.CommandContext(__obf_efbbc02bedef0e2d, FFProbePath, "-v", "quiet", "-print_format", "json", "-show_format", __obf_f2f454838e7353a1), &__obf_17e491a79e2a4b27); __obf_3698cbf06506c070 != nil {
		return nil, fmt.Errorf("error getting JSON from ffprobe output for file '%v': %v", __obf_f2f454838e7353a1, __obf_3698cbf06506c070)
	}

	return &__obf_17e491a79e2a4b27.Format, nil
}
