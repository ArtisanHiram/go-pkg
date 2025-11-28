package __obf_5028a0a829ddcdab

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

func GetVideoInfo(__obf_4b28c4ed43afcd28 string) (*MediaInfo, error) {
	__obf_19b79b6a380e8507, __obf_8c503727c5cd3eea := context.WithTimeout(context.Background(), 20*time.Second)
	defer __obf_8c503727c5cd3eea()
	var __obf_452f0fc3eadce8f6 struct {
		Format MediaInfo
	}
	if __obf_97be5ef7ba0b1a8f := ExecJsonStdout(exec.CommandContext(__obf_19b79b6a380e8507, FFProbePath, "-v", "quiet", "-print_format", "json", "-show_format", __obf_4b28c4ed43afcd28), &__obf_452f0fc3eadce8f6); __obf_97be5ef7ba0b1a8f != nil {
		return nil, fmt.Errorf("error getting JSON from ffprobe output for file '%v': %v", __obf_4b28c4ed43afcd28, __obf_97be5ef7ba0b1a8f)
	}

	return &__obf_452f0fc3eadce8f6.Format, nil
}
