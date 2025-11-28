package __obf_42bbcad92b7de1a8

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

func GetVideoInfo(__obf_fe2438cb530d0431 string) (*MediaInfo, error) {
	__obf_a8cca90b5a49e090, __obf_93f551cc9971e65e := context.WithTimeout(context.Background(), 20*time.Second)
	defer __obf_93f551cc9971e65e()
	var __obf_9a281637508bffbb struct {
		Format MediaInfo
	}
	if __obf_59b2ec060775896c := ExecJsonStdout(exec.CommandContext(__obf_a8cca90b5a49e090, FFProbePath, "-v", "quiet", "-print_format", "json", "-show_format", __obf_fe2438cb530d0431), &__obf_9a281637508bffbb); __obf_59b2ec060775896c != nil {
		return nil, fmt.Errorf("error getting JSON from ffprobe output for file '%v': %v", __obf_fe2438cb530d0431, __obf_59b2ec060775896c)
	}

	return &__obf_9a281637508bffbb.Format, nil
}
