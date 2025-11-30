package __obf_f60326fd90eb13d9

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

func GetVideoInfo(__obf_97f00cf8dd461989 string) (*MediaInfo, error) {
	__obf_b580c2427f47a43f, __obf_36b7b94502851f13 := context.WithTimeout(context.Background(), 20*time.Second)
	defer __obf_36b7b94502851f13()
	var __obf_1fb2eda8309f89d3 struct {
		Format MediaInfo
	}
	if __obf_083e5685f93ba07c := ExecJsonStdout(exec.CommandContext(__obf_b580c2427f47a43f, FFProbePath, "-v", "quiet", "-print_format", "json", "-show_format", __obf_97f00cf8dd461989), &__obf_1fb2eda8309f89d3); __obf_083e5685f93ba07c != nil {
		return nil, fmt.Errorf("error getting JSON from ffprobe output for file '%v': %v", __obf_97f00cf8dd461989, __obf_083e5685f93ba07c)
	}

	return &__obf_1fb2eda8309f89d3.Format, nil
}
