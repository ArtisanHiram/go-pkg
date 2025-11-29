package __obf_6ff082bd539c7df0

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

func GetVideoInfo(__obf_7e66a22be900b14f string) (*MediaInfo, error) {
	__obf_5254a2186cef3f5f, __obf_78c0c95a5ece1ca5 := context.WithTimeout(context.Background(), 20*time.Second)
	defer __obf_78c0c95a5ece1ca5()
	var __obf_c633fcc2cc57bf1c struct {
		Format MediaInfo
	}
	if __obf_ccdf867ff6e9ddb3 := ExecJsonStdout(exec.CommandContext(__obf_5254a2186cef3f5f, FFProbePath, "-v", "quiet", "-print_format", "json", "-show_format", __obf_7e66a22be900b14f), &__obf_c633fcc2cc57bf1c); __obf_ccdf867ff6e9ddb3 != nil {
		return nil, fmt.Errorf("error getting JSON from ffprobe output for file '%v': %v", __obf_7e66a22be900b14f, __obf_ccdf867ff6e9ddb3)
	}

	return &__obf_c633fcc2cc57bf1c.Format, nil
}
