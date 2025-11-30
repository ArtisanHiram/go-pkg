package __obf_08ee9322ff6adb83

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

func GetVideoInfo(__obf_273e2c2474c36944 string) (*MediaInfo, error) {
	__obf_60ee4500ce182178, __obf_2c533782f8daaa17 := context.WithTimeout(context.Background(), 20*time.Second)
	defer __obf_2c533782f8daaa17()
	var __obf_31c328027e6a3134 struct {
		Format MediaInfo
	}
	if __obf_dd2d47498586b1c8 := ExecJsonStdout(exec.CommandContext(__obf_60ee4500ce182178, FFProbePath, "-v", "quiet", "-print_format", "json", "-show_format", __obf_273e2c2474c36944), &__obf_31c328027e6a3134); __obf_dd2d47498586b1c8 != nil {
		return nil, fmt.Errorf("error getting JSON from ffprobe output for file '%v': %v", __obf_273e2c2474c36944, __obf_dd2d47498586b1c8)
	}

	return &__obf_31c328027e6a3134.Format, nil
}
