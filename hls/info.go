package __obf_ab25ed2437cd567a

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

func GetVideoInfo(__obf_77c0216d25f1382f string) (*MediaInfo, error) {
	__obf_378e35ef817a6881, __obf_032b931ab588608d := context.WithTimeout(context.Background(), 20*time.Second)
	defer __obf_032b931ab588608d()
	var __obf_2c72c6ce5ba84cf3 struct {
		Format MediaInfo
	}
	if __obf_9d170493b73636ca := ExecJsonStdout(exec.CommandContext(__obf_378e35ef817a6881, FFProbePath, "-v", "quiet", "-print_format", "json", "-show_format", __obf_77c0216d25f1382f), &__obf_2c72c6ce5ba84cf3); __obf_9d170493b73636ca != nil {
		return nil, fmt.Errorf("error getting JSON from ffprobe output for file '%v': %v", __obf_77c0216d25f1382f, __obf_9d170493b73636ca)
	}

	return &__obf_2c72c6ce5ba84cf3.Format, nil
}
