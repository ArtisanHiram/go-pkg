package __obf_acea4ab24a824c18

// https://github.com/shimberger/gohls
import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os/exec"
	"syscall"
)

const (
	FFProbePath = "ffprobe"
	FFMpegPath  = "ffmpeg"
)

type Option struct {
	CacheDir            string  `yaml:"cache-dir"`
	SegmentLen          float32 `yaml:"segment-len"`
	WorkersNum          int     `yaml:"workers-num"`
	ExeTimeout          int     `yaml:"exe-timeout"`
	PlaylistContentType string  `yaml:"playlist-content-type"` //  "application/vnd.apple.mpegurl"
	CaptionContentType  string  `yaml:"caption-content-type"`  //  "text/vtt"
}

type HlsServer struct {
	Option                 *Option
	__obf_32bdc8ca319d4ad7 *Worker
	__obf_a72f121784d2d8ec *Worker
	__obf_4abf18d2926be19a *Worker
}

func NewHlsServer(__obf_e2d01abc68aff34f *Option) *HlsServer {
	return &HlsServer{
		__obf_e2d01abc68aff34f,
		NewWorker(__obf_e2d01abc68aff34f.WorkersNum, __obf_e2d01abc68aff34f.ExeTimeout, __obf_e2d01abc68aff34f.CacheDir),
		NewWorker(__obf_e2d01abc68aff34f.WorkersNum, __obf_e2d01abc68aff34f.ExeTimeout, __obf_e2d01abc68aff34f.CacheDir),
		NewWorker(__obf_e2d01abc68aff34f.WorkersNum, __obf_e2d01abc68aff34f.ExeTimeout, __obf_e2d01abc68aff34f.CacheDir),
	}
}

func ExecJsonStdout(__obf_d72bd3c5aff9de25 *exec.Cmd, __obf_3699794f57b2a656 any) error {

	__obf_3317e77c6d5e7689 := new(bytes.Buffer)
	if __obf_3698cbf06506c070 := ExecWriteStdout(__obf_d72bd3c5aff9de25, __obf_3317e77c6d5e7689); __obf_3698cbf06506c070 != nil {
		return __obf_3698cbf06506c070
	}

	return json.Unmarshal(__obf_3317e77c6d5e7689.Bytes(), __obf_3699794f57b2a656)
}

func ExecWriteStdout(__obf_d72bd3c5aff9de25 *exec.Cmd, __obf_7bfc44ad58c48031 io.Writer) error {

	__obf_d9ec524741baccf5, __obf_3698cbf06506c070 := __obf_d72bd3c5aff9de25.StdoutPipe()
	if __obf_3698cbf06506c070 != nil {
		return fmt.Errorf("error opening stdout of command: %v", __obf_3698cbf06506c070)
	}
	defer __obf_d9ec524741baccf5.Close()

	if __obf_3698cbf06506c070 = __obf_d72bd3c5aff9de25.Start(); __obf_3698cbf06506c070 != nil {
		return fmt.Errorf("error starting command: %v", __obf_3698cbf06506c070)
	}

	if _, __obf_3698cbf06506c070 = io.Copy(__obf_7bfc44ad58c48031, __obf_d9ec524741baccf5); __obf_3698cbf06506c070 != nil {
		// Ask the process to exit
		_ = __obf_d72bd3c5aff9de25.Process.Signal(syscall.SIGKILL)
		_, _ = __obf_d72bd3c5aff9de25.Process.Wait()
		return fmt.Errorf("error copying stdout to buffer: %v", __obf_3698cbf06506c070)
	}
	if __obf_3698cbf06506c070 = __obf_d72bd3c5aff9de25.Wait(); __obf_3698cbf06506c070 != nil {
		return fmt.Errorf("command failed %v", __obf_3698cbf06506c070)
	}

	return nil
}
