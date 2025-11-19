package __obf_dc577bef9ac5c6b8

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

func GetVideoInfo(__obf_ef682768d47e5b4e string) (*MediaInfo, error) {
	__obf_5d1af4120f069bb9, __obf_22443b10029c5b8b := context.WithTimeout(context.Background(), 20*time.Second)
	defer __obf_22443b10029c5b8b()
	var __obf_f510363e3a8ebf43 struct {
		Format MediaInfo
	}
	if __obf_14a3b54d08bd697d := ExecJsonStdout(exec.CommandContext(__obf_5d1af4120f069bb9, FFProbePath, "-v", "quiet", "-print_format", "json", "-show_format", __obf_ef682768d47e5b4e), &__obf_f510363e3a8ebf43); __obf_14a3b54d08bd697d != nil {
		return nil, fmt.Errorf("error getting JSON from ffprobe output for file '%v': %v", __obf_ef682768d47e5b4e, __obf_14a3b54d08bd697d)
	}

	return &__obf_f510363e3a8ebf43.Format, nil
}
