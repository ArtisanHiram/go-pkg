package __obf_90dd9b56c0f1bd65

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

func GetVideoInfo(__obf_9b7ae180e21623e5 string) (*MediaInfo, error) {
	__obf_9b5717753d5f5860, __obf_17c78879ddf4acdb := context.WithTimeout(context.Background(), 20*time.Second)
	defer __obf_17c78879ddf4acdb()
	var __obf_cdbc8d2c245e05fb struct {
		Format MediaInfo
	}
	if __obf_dadfdd29cd0d4fe8 := ExecJsonStdout(exec.CommandContext(__obf_9b5717753d5f5860, FFProbePath, "-v", "quiet", "-print_format", "json", "-show_format", __obf_9b7ae180e21623e5), &__obf_cdbc8d2c245e05fb); __obf_dadfdd29cd0d4fe8 != nil {
		return nil, fmt.Errorf("error getting JSON from ffprobe output for file '%v': %v", __obf_9b7ae180e21623e5, __obf_dadfdd29cd0d4fe8)
	}

	return &__obf_cdbc8d2c245e05fb.Format, nil
}
